/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package spec

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	releaseclustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/helmrelease/cluster"
)

const (
	ChartRefKey                = "chart_ref"
	GitRepositoryKey           = "git_repository"
	HelmRepositorykey          = "helm_repository"
	RepositoryNameKey          = "repository_name"
	RepositoryNamespaceNameKey = "repository_namespace"
	ChartPathKey               = "chart_path"
	ChartNameKey               = "chart_name"
	VersionKey                 = "version"
	IntervalKey                = "interval"
	InlineConfigKey            = "inline_config"
	TargetNamespaceKey         = "target_namespace"
	SpecKey                    = "spec"
)

var SpecSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Spec for the Repository.",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			InlineConfigKey: {
				Type:        schema.TypeString,
				Description: "Inline values in yaml format.",
				Optional:    true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					newInlineConfig, err := readYamlFile(new)
					if err != nil {
						return false
					}

					return old == newInlineConfig
				},
			},
			TargetNamespaceKey: {
				Type:        schema.TypeString,
				Description: "Name of target namespace.",
				Optional:    true,
			},
			IntervalKey: {
				Type:        schema.TypeString,
				Description: "Interval at which to reconcile the Helm release. This is the interval at which Tanzu Mission Control will attempt to reconcile changes in the helm release to the cluster. A sync interval of 0 would result in no future syncs. If no value is entered, a default interval of 5 minutes will be applied as `5m`.",
				Optional:    true,
				Default:     "5m",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					durationInScript, err := time.ParseDuration(old)
					if err != nil {
						return false
					}

					durationInState, err := time.ParseDuration(new)
					if err != nil {
						return false
					}

					return durationInScript.Seconds() == durationInState.Seconds()
				},
			},
			ChartRefKey: refSchema,
		},
	},
}

var refSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Reference to the chart which will be installed.",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			GitRepositoryKey:  gitRepoSchema,
			HelmRepositorykey: helmRepoSchema,
		},
	},
}

var gitRepoSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Git repository type.",
	Optional:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			RepositoryNameKey: {
				Type:        schema.TypeString,
				Description: "Repository name.",
				Required:    true,
			},
			RepositoryNamespaceNameKey: {
				Type:        schema.TypeString,
				Description: "Repository namespace.",
				Required:    true,
			},
			ChartPathKey: {
				Type:        schema.TypeString,
				Description: "Path of the chart in the git repository.",
				Required:    true,
			},
		},
	},
}

var helmRepoSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Helm repository type.",
	Optional:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			RepositoryNameKey: {
				Type:        schema.TypeString,
				Description: "Repository name.",
				Required:    true,
			},
			RepositoryNamespaceNameKey: {
				Type:        schema.TypeString,
				Description: "Repository namespace.",
				Required:    true,
			},
			ChartNameKey: {
				Type:        schema.TypeString,
				Description: "Name of the chart in the helm repository.",
				Required:    true,
			},
			VersionKey: {
				Type:        schema.TypeString,
				Description: "Chart version, applicable for helm repository type.",
				Required:    true,
			},
		},
	},
}

func expandRef(data []interface{}) (specRef *releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseChartRef) {
	if len(data) == 0 || data[0] == nil {
		return specRef
	}

	refData, _ := data[0].(map[string]interface{})

	specRef = &releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseChartRef{}

	if ref, ok := refData[GitRepositoryKey]; ok {
		if v1, ok := ref.([]interface{}); ok && len(v1) != 0 {
			data := v1[0].(map[string]interface{})

			var repositoryName, repositoryNamespace, chartPath string

			if v, ok := data[RepositoryNameKey]; ok {
				repositoryName = v.(string)
			}

			if v, ok := data[RepositoryNamespaceNameKey]; ok {
				repositoryNamespace = v.(string)
			}

			if v, ok := data[ChartPathKey]; ok {
				chartPath = v.(string)
			}

			specRef.Chart = chartPath
			specRef.RepositoryName = repositoryName
			specRef.RepositoryNamespace = repositoryNamespace
			specRef.RepositoryType = releaseclustermodel.NewVmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryType(
				releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryTypeGIT,
			)

			return specRef
		}
	}

	if ref, ok := refData[HelmRepositorykey]; ok {
		if v1, ok := ref.([]interface{}); ok && len(v1) != 0 {
			data := v1[0].(map[string]interface{})

			var repositoryName, repositoryNamespace, chartName, version string

			if v, ok := data[RepositoryNameKey]; ok {
				repositoryName = v.(string)
			}

			if v, ok := data[RepositoryNamespaceNameKey]; ok {
				repositoryNamespace = v.(string)
			}

			if v, ok := data[ChartNameKey]; ok {
				chartName = v.(string)
			}

			if v, ok := data[VersionKey]; ok {
				version = v.(string)
			}

			specRef.Chart = chartName
			specRef.Version = version
			specRef.RepositoryName = repositoryName
			specRef.RepositoryNamespace = repositoryNamespace
			specRef.RepositoryType = releaseclustermodel.NewVmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryType(
				releaseclustermodel.VmwareTanzuManageV1alpha1ClusterNamespaceFluxcdHelmReleaseRepositoryTypeHELM,
			)

			return specRef
		}
	}

	return specRef
}

func HasSpecChanged(d *schema.ResourceData) bool {
	updateRequired := false

	switch {
	case d.HasChange(helper.GetFirstElementOf(SpecKey, IntervalKey)):
		fallthrough
	case d.HasChange(helper.GetFirstElementOf(SpecKey, InlineConfigKey)):
		fallthrough
	case d.HasChange(helper.GetFirstElementOf(SpecKey, TargetNamespaceKey)):
		fallthrough
	case d.HasChange(helper.GetFirstElementOf(SpecKey, ChartRefKey, GitRepositoryKey, RepositoryNameKey)):
		fallthrough
	case d.HasChange(helper.GetFirstElementOf(SpecKey, ChartRefKey, GitRepositoryKey, RepositoryNamespaceNameKey)):
		fallthrough
	case d.HasChange(helper.GetFirstElementOf(SpecKey, ChartRefKey, GitRepositoryKey, ChartPathKey)):
		fallthrough
	case d.HasChange(helper.GetFirstElementOf(SpecKey, ChartRefKey, HelmRepositorykey, RepositoryNameKey)):
		fallthrough
	case d.HasChange(helper.GetFirstElementOf(SpecKey, ChartRefKey, HelmRepositorykey, RepositoryNamespaceNameKey)):
		fallthrough
	case d.HasChange(helper.GetFirstElementOf(SpecKey, ChartRefKey, HelmRepositorykey, ChartNameKey)):
		fallthrough
	case d.HasChange(helper.GetFirstElementOf(SpecKey, ChartRefKey, HelmRepositorykey, VersionKey)):
		updateRequired = true
	}

	return updateRequired
}

func readYamlFile(fileName string) (string, error) {
	inputFile, err := os.Open(fileName)
	if err != nil {
		return "", errors.WithMessage(err, fmt.Sprintf("Error opening the %s file.", fileName))
	}

	defer inputFile.Close()

	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, inputFile)

	if err != nil {
		return "", err
	}

	_, err = yaml.Marshal(buf.String())
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func fileExists(filepath string) bool {
	fileinfo, err := os.Stat(filepath)

	if os.IsNotExist(err) {
		log.Println("[ERROR] file does not exists.")
		return false
	}

	return !fileinfo.IsDir()
}
