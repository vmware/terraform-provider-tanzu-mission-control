/*
Copyright 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package ekscluster

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/authctx"
	clienterrors "github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/errors"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	eksmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/ekscluster"
	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/common"
)

type (
	contextMethodKey struct{}
)

var ignoredTagsPrefix = "tmc.cloud.vmware.com/"

const defaultTimeout = 3 * time.Minute

func ResourceTMCEKSCluster() *schema.Resource {
	return &schema.Resource{
		Schema:        clusterSchema,
		CreateContext: resourceClusterCreate,
		ReadContext:   dataSourceTMCEKSClusterRead,
		UpdateContext: resourceClusterInPlaceUpdate,
		DeleteContext: resourceClusterDelete,
		Description:   "Tanzu Mission Control EKS Cluster Resource",
	}
}

var clusterSchema = map[string]*schema.Schema{
	CredentialNameKey: {
		Type:        schema.TypeString,
		Description: "Name of the AWS Crendential in Tanzu Mission Control",
		Required:    true,
		ForceNew:    true,
	},
	RegionKey: {
		Type:        schema.TypeString,
		Description: "AWS Region of the this cluster",
		Required:    true,
		ForceNew:    true,
	},
	NameKey: {
		Type:        schema.TypeString,
		Description: "Name of this cluster",
		Required:    true,
		ForceNew:    true,
	},
	common.MetaKey: common.Meta,
	specKey:        clusterSpecSchema,
	StatusKey: {
		Type:        schema.TypeMap,
		Description: "Status of the cluster",
		Computed:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
	},
	waitKey: {
		Type:        schema.TypeString,
		Description: "Wait timeout duration until cluster resource reaches READY state. Accepted timeout duration values like 5s, 45m, or 3h, higher than zero",
		Default:     "default",
		Optional:    true,
	},
}

var clusterSpecSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Spec for the cluster",
	Optional:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			clusterGroupKey: {
				Type:        schema.TypeString,
				Description: "Name of the cluster group to which this cluster belongs",
				Default:     clusterGroupDefaultValue,
				Optional:    true,
			},
			proxyNameKey: {
				Type:        schema.TypeString,
				Description: "Optional proxy name is the name of the Proxy Config to be used for the cluster",
				Optional:    true,
			},
			configKey: configSchema,
			nodepoolKey: {
				Type:        schema.TypeList,
				Description: "Nodepool definitions for the cluster",
				Required:    true,
				MinItems:    1,
				Elem:        nodepoolDefinitionSchema,
			},
		},
	},
}

var configSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "EKS config for the cluster control plane",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			roleArnKey: {
				Type:        schema.TypeString,
				Description: "ARN of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations",
				Required:    true,
				ForceNew:    true,
			},
			kubernetesVersionKey: {
				Type:        schema.TypeString,
				Description: "Kubernetes version of the cluster",
				Required:    true,
			},
			tagsKey: {
				Type:        schema.TypeMap,
				Description: "The metadata to apply to the cluster to assist with categorization and organization",
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return strings.Contains(k, ignoredTagsPrefix)
				},
			},
			kubernetesNetworkConfigKey: {
				Type:        schema.TypeList,
				Description: "Kubernetes Network Config",
				Optional:    true,
				ForceNew:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						serviceCidrKey: {
							Type:        schema.TypeString,
							Description: "Service CIDR for Kubernetes services",
							Required:    true,
							ForceNew:    true,
						},
					},
				},
			},
			loggingKey: {
				Type:        schema.TypeList,
				Description: "EKS logging configuration",
				Optional:    true,
				ForceNew:    false,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						apiServerKey: {
							Type:        schema.TypeBool,
							Description: "Enable API server logs",
							Optional:    true,
							ForceNew:    false,
						},
						auditKey: {
							Type:        schema.TypeBool,
							Description: "Enable audit logs",
							Optional:    true,
							ForceNew:    false,
						},
						authenticatorKey: {
							Type:        schema.TypeBool,
							Description: "Enable authenticator logs",
							Optional:    true,
							ForceNew:    false,
						},
						controllerManagerKey: {
							Type:        schema.TypeBool,
							Description: "Enable controller manager logs",
							Optional:    true,
							ForceNew:    false,
						},
						schedulerKey: {
							Type:        schema.TypeBool,
							Description: "Enable scheduler logs",
							Optional:    true,
							ForceNew:    false,
						},
					},
				},
			},
			vpcKey: vpcSchema,
		},
	},
}

var vpcSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "VPC config",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			enablePrivateAccessKey: {
				Type:        schema.TypeBool,
				Description: "Enable private access on the cluster",
				Optional:    true,
				ForceNew:    false,
			},
			enablePublicAccessKey: {
				Type:        schema.TypeBool,
				Description: "Enable public access on the cluster",
				Optional:    true,
				ForceNew:    false,
			},
			publicAccessCidrsKey: {
				Type:        schema.TypeSet,
				Description: "Public access cidrs",
				Optional:    true,
				ForceNew:    false,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			securityGroupsKey: {
				Type:        schema.TypeSet,
				Description: "Security groups for the cluster VMs",
				Optional:    true,
				ForceNew:    true,
				MaxItems:    5,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			subnetIdsKey: {
				Type:        schema.TypeSet,
				Description: "Subnet ids used by the cluster",
				Required:    true,
				ForceNew:    true,
				MinItems:    2,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	},
}

func resourceClusterCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config, ok := m.(authctx.TanzuContext)
	if !ok {
		log.Println("[ERROR] error while retrieving Tanzu auth config")
		return diag.Errorf("error while retrieving Tanzu auth config")
	}

	clusterReq := &eksmodel.VmwareTanzuManageV1alpha1EksclusterCreateUpdateEksClusterRequest{
		EksCluster: &eksmodel.VmwareTanzuManageV1alpha1EksclusterEksCluster{
			FullName: constructFullname(d),
			Meta:     common.ConstructMeta(d),
			Spec:     constructSpec(d),
		},
	}

	clusterResponse, err := config.TMCConnection.EKSClusterResourceService.EksClusterResourceServiceCreate(clusterReq)
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to create Tanzu Mission Control EKS cluster entry, name : %s", d.Get(NameKey)))
	}

	d.SetId(clusterResponse.EksCluster.Meta.UID)

	return dataSourceTMCEKSClusterRead(context.WithValue(ctx, contextMethodKey{}, "create"), d, m)
}

func resourceClusterDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(authctx.TanzuContext)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	err := config.TMCConnection.EKSClusterResourceService.EksClusterResourceServiceDelete(constructFullname(d), "false")
	if err != nil && !clienterrors.IsNotFoundError(err) {
		return diag.FromErr(errors.Wrapf(err, "Unable to delete Tanzu Mission Control EKS cluster entry, name : %s", d.Get(NameKey)))
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	getClusterResourceRetryableFn := func() (retry bool, err error) {
		_, err = config.TMCConnection.EKSClusterResourceService.EksClusterResourceServiceGet(constructFullname(d))
		if err == nil {
			return true, errors.New("cluster deletion in progress")
		}

		if !clienterrors.IsNotFoundError(err) {
			return true, err
		}

		return false, nil
	}

	timeoutDuration := getRetryTimeout(d)

	_, err = helper.RetryUntilTimeout(getClusterResourceRetryableFn, 10*time.Second, timeoutDuration)
	if err != nil {
		diag.FromErr(errors.Wrapf(err, "verify %s EKS cluster resource clean up", d.Get(NameKey)))
	}

	return diags
}

func resourceClusterInPlaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	config := m.(authctx.TanzuContext)

	// Get call to initialise the cluster struct
	getResp, err := config.TMCConnection.EKSClusterResourceService.EksClusterResourceServiceGet(constructFullname(d))
	if err != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to get Tanzu Mission Control EKS cluster entry, name : %s", d.Get(NameKey)))
	}

	opsRetryTimeout := getRetryTimeout(d)

	clusterSpec := constructSpec(d)

	// EKS cluster update API on TMC side ignores nodepools passed to it.
	// The nodepools have to be updated via separate nodepool API, hence we
	// deal with them separately.
	errnp := handleNodepoolDiffs(config, opsRetryTimeout, getResp.EksCluster.FullName, clusterSpec.NodePools)

	clusterSpec.NodePools = nil

	errcl := handleClusterDiff(config, getResp.EksCluster, common.ConstructMeta(d), clusterSpec)
	if errcl != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to update Tanzu Mission Control EKS cluster entry, name : %s", d.Get(NameKey)))
	}

	// this is moved here so as to not bail on the cluster update
	// when there is a nodepool update error
	if errnp != nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to update Tanzu Mission Control EKS cluster's nodepools, name : %s", d.Get(NameKey)))
	}

	log.Printf("[INFO] cluster update successful")

	return dataSourceTMCEKSClusterRead(ctx, d, m)
}

func handleClusterDiff(config authctx.TanzuContext, tmcCluster *eksmodel.VmwareTanzuManageV1alpha1EksclusterEksCluster, meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta, clusterSpec *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec) error {
	updateCluster := false

	if meta.Description != tmcCluster.Meta.Description ||
		!reflect.DeepEqual(meta.Labels, tmcCluster.Meta.Labels) {
		updateCluster = true
		tmcCluster.Meta.Description = meta.Description
		tmcCluster.Meta.Labels = meta.Labels
	}

	if !reflect.DeepEqual(clusterSpec, tmcCluster.Spec) {
		updateCluster = true
	}

	// The TF update request was only for nodepools.
	// No need to update Cluster, it will error out.
	if !updateCluster {
		return nil
	}

	// there is some translation error, which results
	// in mismatch on the server.
	tmcCluster.Meta.CreationTime = strfmt.DateTime{}

	newCluster := &eksmodel.VmwareTanzuManageV1alpha1EksclusterEksCluster{
		FullName: tmcCluster.FullName,
		Meta:     tmcCluster.Meta,
		Spec:     clusterSpec,
	}

	_, err := config.TMCConnection.EKSClusterResourceService.EksClusterResourceServiceUpdate(
		&eksmodel.VmwareTanzuManageV1alpha1EksclusterCreateUpdateEksClusterRequest{
			EksCluster: newCluster,
		},
	)

	if err != nil {
		return errors.Wrapf(err, "Unable to update Tanzu Mission Control EKS cluster entry, name : %s", tmcCluster.FullName.Name)
	}

	return nil
}

func handleNodepoolDiffs(config authctx.TanzuContext, opsRetryTimeout time.Duration, clusterFn *eksmodel.VmwareTanzuManageV1alpha1EksclusterFullName, nodepools []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition) error {
	npresp, err := config.TMCConnection.EKSNodePoolResourceService.EksNodePoolResourceServiceList(clusterFn)
	if err != nil {
		return errors.Wrapf(err, "failed to list nodepools for cluster: %s", clusterFn)
	}

	npPosMap := nodepoolPosMap(nodepools)
	tmcNps := map[string]*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolNodepool{}

	npUpdate := []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition{}
	npCreate := []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition{}
	npDelete := []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolFullName{}

	for _, tmcNp := range npresp.Nodepools {
		tmcNps[tmcNp.FullName.Name] = tmcNp

		if pos, ok := npPosMap[tmcNp.FullName.Name]; ok {
			// np exisits in both TMC and TF
			newNp := nodepools[pos]
			fillTMCSetValues(tmcNp.Spec, newNp.Spec)

			if checkNodepoolUpdate(tmcNp, newNp) {
				fmt.Printf("np %s diff: %s", tmcNp.FullName.Name, cmp.Diff(tmcNp.Spec, newNp.Spec))
				npUpdate = append(npUpdate, newNp)
			}
		} else {
			// np exisits in TMC but not in TF
			npDelete = append(npDelete, tmcNp.FullName)
		}
	}

	for _, tfNp := range nodepools {
		if _, ok := tmcNps[tfNp.Info.Name]; !ok {
			npCreate = append(npCreate, tfNp)
		}
	}

	err = handleNodepoolCreates(config, opsRetryTimeout, clusterFn, npCreate)
	if err != nil {
		return errors.Wrap(err, "failed to create nodepools that are not present in TMC")
	}

	err = handleNodepoolUpdates(config, opsRetryTimeout, tmcNps, npUpdate)
	if err != nil {
		return errors.Wrapf(err, "failed to update existing nodepools")
	}

	err = handleNodepoolDeletes(config, opsRetryTimeout, npDelete)
	if err != nil {
		return errors.Wrapf(err, "failed to delete nodepools")
	}

	return nil
}

func handleNodepoolDeletes(config authctx.TanzuContext, opsRetryTimeout time.Duration, npFns []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolFullName) error {
	for _, npFn := range npFns {
		err := config.TMCConnection.EKSNodePoolResourceService.EksNodePoolResourceServiceDelete(npFn)
		if err != nil {
			return errors.Wrap(err, "delete api call failed")
		}

		getNodepoolResourceRetryableFn := func() (retry bool, err error) {
			_, err = config.TMCConnection.EKSNodePoolResourceService.EksNodePoolResourceServiceGet(npFn)
			if err == nil {
				// we don't want to fail deletion if the deletion is not
				// completed within the expected time
				return true, nil
			}

			if !clienterrors.IsNotFoundError(err) {
				return true, err
			}

			return false, nil
		}

		_, err = helper.RetryUntilTimeout(getNodepoolResourceRetryableFn, 10*time.Second, opsRetryTimeout)
		if err != nil {
			return errors.Wrapf(err, "failed to verify EKS nodepool resource(%s) clean up", npFn.Name)
		}
	}

	return nil
}

func handleNodepoolUpdates(config authctx.TanzuContext, opsRetryTimeout time.Duration, tmcNps map[string]*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolNodepool, nps []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition) error {
	for _, np := range nps {
		tmcNp := tmcNps[np.Info.Name]

		req := &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolAPIRequest{
			Nodepool: &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolNodepool{
				FullName: tmcNp.FullName,
				Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
					Annotations:      tmcNp.Meta.Annotations,
					Description:      np.Info.Description,
					Labels:           tmcNp.Meta.Labels,
					ParentReferences: tmcNp.Meta.ParentReferences,
					ResourceVersion:  tmcNp.Meta.ResourceVersion,
					UID:              tmcNp.Meta.UID,
				},
				Spec: np.Spec,
			},
		}

		_, err := config.TMCConnection.EKSNodePoolResourceService.EksNodePoolResourceServiceUpdate(req)
		if err != nil {
			return errors.Wrapf(err, "failed to update nodepool %s", np.Info.Name)
		}

		getNodepoolResourceRetryableFn := getWaitForNodepoolReadyFn(config, tmcNp.FullName)

		_, err = helper.RetryUntilTimeout(getNodepoolResourceRetryableFn, 10*time.Second, opsRetryTimeout)
		if err != nil {
			return errors.Wrapf(err, "failed to verify EKS nodepool resource(%s) creation", np.Info.Name)
		}
	}

	return nil
}

func fillTMCSetValues(tmcNpSpec *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec, npSpec *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec) {
	if npSpec.AmiType == "" {
		npSpec.AmiType = tmcNpSpec.AmiType
	}

	if npSpec.CapacityType == "" {
		npSpec.CapacityType = tmcNpSpec.CapacityType
	}
}

func handleNodepoolCreates(config authctx.TanzuContext, opsRetryTimeout time.Duration, clusterFn *eksmodel.VmwareTanzuManageV1alpha1EksclusterFullName, nps []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition) error {
	for _, np := range nps {
		npFn := &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolFullName{
			CredentialName: clusterFn.CredentialName,
			Region:         clusterFn.Region,
			EksClusterName: clusterFn.Name,
			Name:           np.Info.Name,
		}

		req := &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolAPIRequest{
			Nodepool: &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolNodepool{
				FullName: npFn,
				Meta: &objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta{
					Description: np.Info.Description,
				},
				Spec: np.Spec,
			},
		}

		_, err := config.TMCConnection.EKSNodePoolResourceService.EksNodePoolResourceServiceCreate(req)
		if err != nil {
			return errors.Wrapf(err, "failed to create nodepool %s", np.Info.Name)
		}

		getNodepoolResourceRetryableFn := getWaitForNodepoolReadyFn(config, npFn)

		_, err = helper.RetryUntilTimeout(getNodepoolResourceRetryableFn, 10*time.Second, opsRetryTimeout)
		if err != nil {
			return errors.Wrapf(err, "failed to verify EKS nodepool resource(%s) creation", npFn.Name)
		}
	}

	return nil
}

func getWaitForNodepoolReadyFn(config authctx.TanzuContext, npFn *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolFullName) func() (retry bool, err error) {
	return func() (retry bool, err error) {
		resp, err := config.TMCConnection.EKSNodePoolResourceService.EksNodePoolResourceServiceGet(npFn)
		if err != nil {
			return true, errors.Wrapf(err, "Unable to get Tanzu Mission Control EKS nodepoool entry, name : %s", npFn.Name)
		}

		if resp.Nodepool.Status.Phase != nil &&
			*resp.Nodepool.Status.Phase != eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolStatusPhaseREADY {
			return true, nil
		}

		return false, nil
	}
}

func checkNodepoolUpdate(oldNp *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolNodepool, newNp *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition) bool {
	return oldNp.Meta.Description != newNp.Info.Description ||
		!nodepoolSpecEqual(oldNp.Spec, newNp.Spec)
}

func constructFullname(d *schema.ResourceData) (fullname *eksmodel.VmwareTanzuManageV1alpha1EksclusterFullName) {
	fullname = &eksmodel.VmwareTanzuManageV1alpha1EksclusterFullName{}

	fullname.CredentialName, _ = d.Get(CredentialNameKey).(string)
	fullname.Region, _ = d.Get(RegionKey).(string)
	fullname.Name, _ = d.Get(NameKey).(string)

	return fullname
}

func getRetryTimeout(d *schema.ResourceData) time.Duration {
	timeoutValueData, _ := d.Get(waitKey).(string)
	if timeoutValueData != "default" {
		providedDuration, parseErr := time.ParseDuration(timeoutValueData)
		if parseErr == nil {
			return providedDuration
		}
	}

	return defaultTimeout
}
