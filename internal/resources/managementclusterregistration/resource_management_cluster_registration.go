/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package managementclusterregistration

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/pkg/errors"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	k8sClient "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	clustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster"
	managementclusterregistrationmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/managementclusterregistration"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/cluster/manifest"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

type (
	contextMethodKey struct{}
)

func ResourceManagementClusterRegistration() *schema.Resource {
	return &schema.Resource{
		Schema: managementClusterRegistrationSchema,
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			return dataSourceClusterRead(helper.GetContextWithCaller(ctx, helper.DataRead), d, m)
		},
		CreateContext: resourceClusterCreate,
		UpdateContext: resourceClusterInPlaceUpdate,
		DeleteContext: resourceClusterDelete,
		Description:   "Tanzu Mission Control Management Cluster Registration Resource",
	}
}

var managementClusterRegistrationSchema = map[string]*schema.Schema{
	NameKey: {
		Type:        schema.TypeString,
		Description: "Name of this management cluster",
		Required:    true,
		ForceNew:    true,
	},
	OrgIDKey: {
		Type:        schema.TypeString,
		Description: "ID of Organization.",
		Optional:    true,
	},
	common.MetaKey: common.Meta,
	specKey: {
		Type:     schema.TypeList,
		Optional: true,
		MinItems: 1,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				clusterGroupKey: {
					Type:         schema.TypeString,
					Description:  "Cluster group name to be used by default for workload clusters",
					Required:     true,
					ValidateFunc: validation.All(validation.StringIsNotEmpty),
				},
				kubernetesProviderTypeKey: {
					Type:        schema.TypeString,
					Description: "Kubernetes provider type",
					Required:    true,
					ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
						"VMWARE_TANZU_KUBERNETES_GRID_SERVICE", "VMWARE_TANZU_KUBERNETES_GRID",
					}, false)),
				},
				imageRegistryKey: {
					Type:        schema.TypeString,
					Description: "Image registry which is only allowed for TKGm",
					Optional:    true,
				},
				managedWorkloadClusterImageRegistryKey: {
					Type:        schema.TypeString,
					Description: "Managed workload cluster image registry",
					Optional:    true,
				},
				managementClusterProxyNameKey: {
					Type:        schema.TypeString,
					Description: "Management cluster proxy name",
					Optional:    true,
				},
				managedWorkloadClusterProxyNameKey: {
					Type:        schema.TypeString,
					Description: "Managed workload cluster proxy name",
					Optional:    true,
				},
			},
		},
	},
	registerClusterKey: {
		Type:     schema.TypeList,
		Optional: true,
		MinItems: 1,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				registerClusterKubeConfigPathForTKGmKey: {
					Type:        schema.TypeString,
					Description: "Register management cluster KUBECONFIG path for only TKGm",
					ForceNew:    true,
					Optional:    true,
				},
				registerClusterKubeConfigRawForTKGmKey: {
					Type:        schema.TypeString,
					Description: "Register management cluster KUBECONFIG for only TKGm",
					Optional:    true,
					ForceNew:    true,
					Sensitive:   true,
				},
				registerClusterDescriptionForTKGmKey: {
					Type:         schema.TypeString,
					Description:  "Register management cluster description for only TKGm",
					Optional:     true,
					ValidateFunc: validation.StringIsNotWhiteSpace,
				},
			},
		},
	},
	StatusKey: {
		Type:        schema.TypeMap,
		Description: "Status of the management cluster",
		Computed:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
	},
	waitKey: {
		Type:        schema.TypeString,
		Description: "Wait timeout duration.",
		Default:     defaultWaitTimeout.String(),
		Optional:    true,
	},
}

func resourceClusterCreate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config, ok := m.(authctx.TanzuContext)

	if !ok {
		log.Println("[ERROR] error while retrieving Tanzu auth config")
		return diag.Errorf("error while retrieving Tanzu auth config")
	}

	createResponse, createError := createRegistrationResource(config, d)

	if createError != nil {
		return diag.FromErr(errors.Wrapf(createError, "Unable to create Management cluster registration, name : %s", d.Get(NameKey)))
	}

	d.SetId(createResponse.ManagementCluster.Meta.UID)

	if v, ok := d.GetOk(registerClusterKey); ok {
		var (
			kubeClient *k8sClient.Client
			err        error
			kubeConfig interface{}
			manifests  string
		)

		err = validateKubeConfig(v)
		if err != nil {
			return diag.FromErr(err)
		}

		spec := constructSpec(d)
		if *spec.KubernetesProviderType != "VMWARE_TANZU_KUBERNETES_GRID" {
			return diag.Errorf("kubernetes_provider_type must have value VMWARE_TANZU_KUBERNETES_GRID so registration with kubeconfig would be possible")
		}

		isKubeConfigPresent := func(typeKey string) bool {
			if value, ok := d.GetOk(helper.GetFirstElementOf(registerClusterKey, typeKey)); ok {
				if value != nil {
					kubeConfig = value
					return true
				}
			}

			return false
		}

		switch {
		case isKubeConfigPresent(registerClusterKubeConfigPathForTKGmKey):
			kubeConfigFile, _ := kubeConfig.(string)
			kubeClient, err = getK8sClientFromFilePath(kubeConfigFile)

		case isKubeConfigPresent(registerClusterKubeConfigRawForTKGmKey):
			rawKubeConfig, _ := kubeConfig.(string)
			kubeClient, err = getK8sClientFromRawInput(rawKubeConfig)
		}

		if err != nil {
			log.Println("[ERROR] error while creating kubernetes client: ", err.Error())
			return diag.FromErr(err)
		}

		if createResponse.ManagementCluster.Spec.ImageRegistry != "" || createResponse.ManagementCluster.Spec.ProxyName != "" {
			clusterManifest, err := config.TMCConnection.ManagementClusterRegistrationResourceService.ManagementClusterManifestHelperGetManifest(constructFullname(d))
			if err != nil {
				return diag.FromErr(errors.Wrapf(err, "Unable to get manifest (%s), err : %s", clusterManifest.Manifest, err))
			}

			manifests = clusterManifest.Manifest
		} else {
			deploymentManifest, err := manifest.GetK8sManifest(createResponse.ManagementCluster.Status.RegistrationURL)
			if err != nil {
				return append(diags, diag.FromErr(err)...)
			}

			manifests = string(deploymentManifest)
		}

		log.Printf("[INFO] Applying %s manifest objects on to kubernetes cluster", constructFullname(d).ToString())

		err = manifest.Create(kubeClient, manifests, true)
		if err != nil {
			return append(diags, diag.FromErr(err)...)
		}

		log.Printf("[INFO] Cluster registered successfully. Tanzu Mission Control resources(%s) applied successfully", constructFullname(d).ToString())
	}

	return append(diags, dataSourceClusterRead(context.WithValue(ctx, contextMethodKey{}, helper.CreateState), d, m)...)
}

func createRegistrationResource(config authctx.TanzuContext, d *schema.ResourceData) (*managementclusterregistrationmodel.VmwareTanzuManageV1alpha1ManagementclusterCreateManagementClusterResponse, error) {
	statusResponse, _ := config.TMCConnection.ManagementClusterRegistrationResourceService.ManagementClusterResourceServiceGet(constructFullname(d))

	var (
		createResponse *managementclusterregistrationmodel.VmwareTanzuManageV1alpha1ManagementclusterCreateManagementClusterResponse
		createError    error
	)

	registrationRequest := &managementclusterregistrationmodel.VmwareTanzuManageV1alpha1ManagementclusterCreateManagementClusterRequest{
		ManagementCluster: &managementclusterregistrationmodel.VmwareTanzuManageV1alpha1ManagementclusterManagementCluster{
			FullName: constructFullname(d),
			Meta:     common.ConstructMeta(d),
			Spec:     constructSpec(d),
		},
	}

	if statusResponse != nil && statusResponse.ManagementCluster != nil && *statusResponse.ManagementCluster.Status.Phase == managementclusterregistrationmodel.VmwareTanzuManageV1alpha1ManagementclusterPhasePENDING && statusResponse.ManagementCluster.Status.RegistrationURL == "" {
		createResponse, createError = config.TMCConnection.ManagementClusterRegistrationResourceService.ManagementClusterResourceReregisterService(registrationRequest)
	} else {
		createResponse, createError = config.TMCConnection.ManagementClusterRegistrationResourceService.ManagementClusterResourceServiceCreate(registrationRequest)
	}

	return createResponse, createError
}

func validateKubeConfig(value interface{}) error {
	data, _ := value.([]interface{})

	if len(data) == 0 || data[0] == nil {
		return fmt.Errorf("%v is not valid: minimum one valid kube config type is required", registerClusterKey)
	}

	kubeConfigData := data[0].(map[string]interface{})

	kubeConfigTypeFound := make([]string, 0)

	if v, ok := kubeConfigData[registerClusterKubeConfigPathForTKGmKey]; ok {
		if v1, ok := v.(string); ok && len(v1) != 0 {
			kubeConfigTypeFound = append(kubeConfigTypeFound, registerClusterKubeConfigPathForTKGmKey)
		}
	}

	if v, ok := kubeConfigData[registerClusterKubeConfigRawForTKGmKey]; ok {
		if v1, ok := v.(string); ok && len(v1) != 0 {
			kubeConfigTypeFound = append(kubeConfigTypeFound, registerClusterKubeConfigRawForTKGmKey)
		}
	}

	if len(kubeConfigTypeFound) == 0 {
		return fmt.Errorf("no valid kube config type found: minimum one valid kube config type is required")
	} else if len(kubeConfigTypeFound) > 1 {
		return fmt.Errorf("found more than one kube config values")
	}

	return nil
}

func getK8sClientFromRawInput(rawConfiguration string) (*k8sClient.Client, error) {
	if strings.TrimSpace(rawConfiguration) == "" {
		return nil, fmt.Errorf("expected raw kubeconfig to not be an empty string or whitespace")
	}

	return getK8sClient("", rawConfiguration)
}

func getK8sClientFromFilePath(filePath string) (*k8sClient.Client, error) {
	if strings.TrimSpace(filePath) == "" {
		return nil, fmt.Errorf("expected kubeconfig file path to not be an empty string or whitespace")
	}

	return getK8sClient(filePath, "")
}

func getK8sClient(filePath string, rawConfiguration string) (*k8sClient.Client, error) {
	var (
		restConfig *rest.Config
		err        error
	)

	switch {
	case filePath != "":
		restConfig, err = clientcmd.BuildConfigFromFlags("", filePath)
		if err != nil {
			return nil, errors.WithMessagef(err, "Invalid kubeconfig file path provided, filepath : %s", filePath)
		}
	case rawConfiguration != "":
		restConfig, err = clientcmd.RESTConfigFromKubeConfig([]byte(rawConfiguration))
		if err != nil {
			return nil, errors.WithMessagef(err, "Invalid raw kubeconfig provided.")
		}
	}

	if restConfig == nil {
		return nil, errors.WithMessagef(err, "Kubeconfig not provided.")
	}

	restConfig.Timeout = 10 * time.Second

	client, err := k8sClient.New(restConfig, k8sClient.Options{})
	if err != nil {
		return nil, errors.WithMessagef(err, "Error in creating kubernetes client from kubeconfig file provided, filepath : %s", filePath)
	}

	return &client, nil
}

func resourceClusterInPlaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	registrationRequest := &managementclusterregistrationmodel.VmwareTanzuManageV1alpha1ManagementclusterCreateManagementClusterRequest{
		ManagementCluster: &managementclusterregistrationmodel.VmwareTanzuManageV1alpha1ManagementclusterManagementCluster{
			FullName: constructFullname(d),
			Meta:     common.ConstructMeta(d),
			Spec:     constructSpec(d),
		},
	}

	registrationResponse, err := config.TMCConnection.ManagementClusterRegistrationResourceService.ManagementClusterResourceServiceUpdate(registrationRequest)

	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to update Management cluster registration, name : %s", d.Get(NameKey)))
	}

	d.SetId(registrationResponse.ManagementCluster.Meta.UID)

	return dataSourceClusterRead(ctx, d, m)
}

func resourceClusterDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(authctx.TanzuContext)

	err := config.TMCConnection.ManagementClusterRegistrationResourceService.ManagementClusterResourceServiceDelete(constructFullname(d), "false")
	if err != nil && !clienterrors.IsNotFoundError(err) {
		return diag.FromErr(errors.Wrapf(err, "Unable to delete managamement cluster registration entry, name : %s", d.Get(NameKey)))
	}

	var diags diag.Diagnostics

	return diags
}

func constructFullname(d *schema.ResourceData) (fullname *managementclusterregistrationmodel.VmwareTanzuManageV1alpha1ManagementclusterFullName) {
	fullname = &managementclusterregistrationmodel.VmwareTanzuManageV1alpha1ManagementclusterFullName{}

	fullname.Name, _ = d.Get(NameKey).(string)
	fullname.OrgID, _ = d.Get(OrgIDKey).(string)

	return fullname
}

func constructSpec(d *schema.ResourceData) (spec *managementclusterregistrationmodel.VmwareTanzuManageV1alpha1ManagementclusterSpec) {
	spec = &managementclusterregistrationmodel.VmwareTanzuManageV1alpha1ManagementclusterSpec{}

	value, ok := d.GetOk(specKey)
	if !ok {
		return spec
	}

	data, _ := value.([]interface{})

	if len(data) == 0 || data[0] == nil {
		return spec
	}

	specData := data[0].(map[string]interface{})

	if v, ok := specData[clusterGroupKey]; ok {
		spec.DefaultClusterGroup = v.(string)
	}

	if v, ok := specData[managedWorkloadClusterImageRegistryKey]; ok {
		spec.DefaultWorkloadClusterImageRegistry = v.(string)
	}

	if v, ok := specData[managedWorkloadClusterProxyNameKey]; ok {
		spec.DefaultWorkloadClusterProxyName = v.(string)
	}

	if v, ok := specData[imageRegistryKey]; ok {
		spec.ImageRegistry = v.(string)
	}

	if v, ok := specData[kubernetesProviderTypeKey]; ok {
		providerType := clustermodel.VmwareTanzuManageV1alpha1CommonClusterKubernetesProviderType(v.(string))
		spec.KubernetesProviderType = &providerType
	}

	if v, ok := specData[managementClusterProxyNameKey]; ok {
		spec.ProxyName = v.(string)
	}

	return spec
}

func flattenSpec(spec *managementclusterregistrationmodel.VmwareTanzuManageV1alpha1ManagementclusterSpec) (data []interface{}) {
	if spec == nil {
		return data
	}

	flattenSpecData := make(map[string]interface{})

	flattenSpecData[clusterGroupKey] = spec.DefaultClusterGroup

	flattenSpecData[managedWorkloadClusterImageRegistryKey] = spec.DefaultWorkloadClusterImageRegistry

	flattenSpecData[managedWorkloadClusterProxyNameKey] = spec.DefaultWorkloadClusterProxyName

	flattenSpecData[imageRegistryKey] = spec.ImageRegistry

	if spec.KubernetesProviderType != nil {
		flattenSpecData[kubernetesProviderTypeKey] = string(*spec.KubernetesProviderType)
	}

	flattenSpecData[managementClusterProxyNameKey] = spec.ProxyName

	return []interface{}{flattenSpecData}
}
