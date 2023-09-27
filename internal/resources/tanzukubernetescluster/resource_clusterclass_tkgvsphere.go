/*
Copyright 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkc

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	vcenterKey         = "vcenter"
	serverKey          = "server"
	datacenterKey      = "datacenter"
	resourcePoolKey    = "resourcePool"
	folderKey          = "folder"
	datastoreKey       = "datastore"
	cloneModeKey       = "cloneMode"
	tlsThumbprintKey   = "tlsThumbprint"
	storagePolicyIDKey = "storagePolicyID"
	templateKey        = "template"

	identityRefKey = "identityRef"
	kindKey        = "kind"

	userKey              = "user"
	sshAuthorizedKeysKey = "sshAuthorizedKeys"

	aviAPIServerHAProviderKey = "aviAPIServerHAProvider"

	apiServerEndpointKey = "apiServerEndpoint"

	apiServerPortKey = "apiServerPort"

	vipNetworkInterfaceKey = "vipNetworkInterface"

	machineKey       = "machine"
	customVMXKeysKey = "customVMXKeys"
	diskGiBKey       = "diskGiB"
	memoryMiBKey     = "memoryMiB"
	numCPUsKey       = "numCPUs"
	nodeLabelsKey    = "nodeLabels"

	nameserversKey   = "nameservers"
	searchDomainsKey = "searchDomains"

	keyKey = "key"

	workerKey = "worker"

	nodePoolLabelsKey = "nodePoolLabels"

	additionalFQDNKey = "additionalFQDN"

	trustKey = "trust"

	additionalTrustedCAsKey = "additionalTrustedCAs"
	dataKey                 = "data"

	controlPlaneKubeletExtraArgsKey = "controlPlaneKubeletExtraArgs"

	controlPlaneCertificateRotationKey = "controlPlaneCertificateRotation"
	activateKey                        = "activate"
	daysBeforeKey                      = "daysBefore"

	proxyKey      = "proxy"
	httpProxyKey  = "httpProxy"
	httpsProxyKey = "httpsProxy"
	noProxyKey    = "noProxy"
	systemWideKey = "systemWide"

	controlPlaneTaintKey = "controlPlaneTaint"

	pciKey                        = "pci"
	controlPlaneTkgVsphereV100Key = "controlPlane"
	devicesKey                    = "devices"
	hardwareVersionKey            = "hardwareVersion"
	deviceIdKey                   = "deviceId"
	vendorIdKey                   = "vendorId"

	imageRepositoryKey          = "imageRepository"
	hostKey                     = "host"
	tlsCertificateValidationKey = "tlsCertificateValidation"

	kubeControllerManagerExtraArgsKey = "kubeControllerManagerExtraArgs"

	customTDNFRepositoryKey = "customTDNFRepository"
	certificateKey          = "certificate"

	kubeVipLoadBalancerProviderKey = "kubeVipLoadBalancerProvider"

	auditLoggingKey = "auditLogging"
	enabledKey      = "enabled"

	apiServerExtraArgsKey = "apiServerExtraArgs"

	additionalImageRegistriesKey = "additionalImageRegistries"
	caCertKey                    = "caCert"
	skipTlsVerifyKey             = "skipTlsVerify"

	ntpServersKey = "ntpServers"

	workerKubeletExtraArgsKey = "workerKubeletExtraArgs"

	podSecurityStandardKey = "podSecurityStandardKey"
	auditKey               = "audit"
	auditVersionKey        = "auditVersion"
	deactivatedKey         = "deactivated"
	enforceKey             = "enforce"
	enforceVersionKey      = "enforceVersion"
	exemptionsKey          = "exemptions"
	namespacesKey          = "namespaces"
	warnKey                = "warn"
	warnVersionKey         = "warnVersion"

	tlsCipherSuitesKey = "tlsCipherSuites"

	etcdExtraArgsKey = "etcdExtraArgs"

	kubeSchedulerExtraArgsKey = "kubeSchedulerExtraArgs"

	eventRateLimitConfKey = "eventRateLimitConf"

	addressesFromPoolsKey = "addressesFromPools"
	ipv6PrimaryKey        = "ipv6Primary"
	apiGroupKey           = "apiGroup"
)

var tkgVsphereV100Schema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		vcenterKey: {
			Type:        schema.TypeList,
			Description: "",
			Required:    true,
			Elem:        vcenterSchema,
		},
		identityRefKey: {
			Type:        schema.TypeList,
			Description: "",
			Required:    true,
			Elem:        identityRefSchema,
		},
		userKey: {
			Type:        schema.TypeList,
			Description: "",
			Required:    true,
			Elem:        userSchema,
		},
		aviAPIServerHAProviderKey: {
			Type:        schema.TypeBool,
			Description: "",
			Required:    true,
			Default:     false,
		},
		apiServerEndpointKey: {
			Type:        schema.TypeString,
			Description: "",
			Optional:    true,
		},
		apiServerPortKey: {
			Type:        schema.TypeInt,
			Description: "",
			Optional:    true,
		},
		vipNetworkInterfaceKey: {
			Type:        schema.TypeString,
			Description: "",
			Required:    true,
			Default:     "eth0",
		},
		controlPlaneTkgVsphereV100Key: {
			Type:        schema.TypeList,
			Description: "",
			Required:    true,
			Elem:        controlPlaneTkgVsphereV100Schema,
		},
		workerKey: {
			Type:        schema.TypeList,
			Description: "",
			Required:    true,
			Elem:        workerSchema,
		},
		nodePoolLabelsKey: {
			Type:        schema.TypeList,
			Description: "",
			Optional:    true,
			Elem:        nodeLabelsSchema,
		},
		additionalFQDNKey: {
			Type:        schema.TypeSet,
			Description: "",
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		trustKey: {
			Type:        schema.TypeList,
			Description: "",
			Optional:    true,
			Elem:        trustSchema,
		},
		controlPlaneKubeletExtraArgsKey: {
			Type:        schema.TypeMap,
			Description: "",
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		controlPlaneCertificateRotationKey: {
			Type:        schema.TypeList,
			Description: "",
			Optional:    true,
			Elem:        controlPlaneCertificateRotationSchema,
		},
		proxyKey: {
			Type:        schema.TypeList,
			Description: "",
			Optional:    true,
			Elem:        proxySchema,
		},
		controlPlaneTaintKey: {
			Type:        schema.TypeBool,
			Description: "",
			Optional:    true,
			Default:     true,
		},
		pciKey: {
			Type:        schema.TypeList,
			Description: "",
			Optional:    true,
			Elem:        pciSchema,
		},
		imageRepositoryKey: {
			Type:        schema.TypeList,
			Description: "",
			Optional:    true,
			Elem:        imageRepositorySchema,
		},
		kubeControllerManagerExtraArgsKey: {
			Type:        schema.TypeMap,
			Description: "",
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		customTDNFRepositoryKey: {
			Type:        schema.TypeList,
			Description: "",
			Optional:    true,
			Elem:        customTDNFRepositorySchema,
		},
		networkKey: {
			Type:        schema.TypeList,
			Description: "",
			Optional:    true,
			Elem:        networkTkgVsphereV100Schema,
		},
		kubeVipLoadBalancerProviderKey: {
			Type:        schema.TypeBool,
			Description: "",
			Optional:    true,
		},
		auditLoggingKey: {
			Type:        schema.TypeList,
			Description: "",
			Optional:    true,
			Elem:        auditLoggingSchema,
		},
		apiServerExtraArgsKey: {
			Type:        schema.TypeMap,
			Description: "",
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		additionalImageRegistriesKey: {
			Type:        schema.TypeList,
			Description: "",
			Optional:    true,
			Elem:        additionalImageRegistriesSchema,
		},
		ntpServersKey: {
			Type:        schema.TypeSet,
			Description: "",
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		workerKubeletExtraArgsKey: {
			Type:        schema.TypeMap,
			Description: "",
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		podSecurityStandardKey: {
			Type:        schema.TypeList,
			Description: "",
			Optional:    true,
			Elem:        podSecurityStandardSchema,
		},
		tlsCipherSuitesKey: {
			Type:        schema.TypeString,
			Description: "",
			Optional:    true,
		},
		etcdExtraArgsKey: {
			Type:        schema.TypeMap,
			Description: "",
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		kubeSchedulerExtraArgsKey: {
			Type:        schema.TypeMap,
			Description: "",
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		eventRateLimitConfKey: {
			Type:        schema.TypeString,
			Description: "",
			Optional:    true,
		},
	},
}

var vcenterSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		serverKey: {
			Type:        schema.TypeString,
			Description: "",
			Required:    true,
		},
		datacenterKey: {
			Type:        schema.TypeString,
			Description: "",
			Required:    true,
		},
		resourcePoolKey: {
			Type:        schema.TypeString,
			Description: "",
			Required:    true,
		},
		folderKey: {
			Type:        schema.TypeString,
			Description: "",
			Required:    true,
		},
		networkKey: {
			Type:        schema.TypeString,
			Description: "",
			Optional:    true,
			Default:     "VMNetwork",
		},
		datastoreKey: {
			Type:        schema.TypeString,
			Description: "",
			Required:    true,
		},
		cloneModeKey: {
			Type:        schema.TypeString,
			Description: "",
			Optional:    true,
			Default:     "fullClone",
		},
		tlsThumbprintKey: {
			Type:        schema.TypeString,
			Description: "",
			Optional:    true,
		},
		storagePolicyIDKey: {
			Type:        schema.TypeString,
			Description: "",
			Optional:    true,
		},
		templateKey: {
			Type:        schema.TypeString,
			Description: "",
			Optional:    true,
		},
	},
}

var identityRefSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		kindKey: {
			Type:        schema.TypeString,
			Description: "",
			Required:    true,
		},
		nameKey: {
			Type:        schema.TypeString,
			Description: "",
			Required:    true,
		},
	},
}

var userSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		sshAuthorizedKeysKey: {
			Type:        schema.TypeSet,
			Description: "",
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
	},
}

var controlPlaneTkgVsphereV100Schema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		machineKey: {
			Type:        schema.TypeList,
			Description: "",
			Optional:    true,
			Elem:        machineSchema,
		},
		networkKey: {
			Type:        schema.TypeList,
			Description: "",
			Optional:    true,
			Elem:        networkMachineTkgVsphereV100Schema,
		},
		nodeLabelsKey: {
			Type:        schema.TypeList,
			Description: "",
			Optional:    true,
			Elem:        nodeLabelsSchema,
		},
	},
}

var machineSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		customVMXKeysKey: {
			Type:        schema.TypeMap,
			Description: "",
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		diskGiBKey: {
			Type:        schema.TypeString,
			Description: "",
			Optional:    true,
		},
		memoryMiBKey: {
			Type:        schema.TypeString,
			Description: "",
			Optional:    true,
		},
		numCPUsKey: {
			Type:        schema.TypeString,
			Description: "",
			Optional:    true,
		},
	},
}

var networkMachineTkgVsphereV100Schema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		nameserversKey: {
			Type:        schema.TypeSet,
			Description: "",
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		searchDomainsKey: {
			Type:        schema.TypeSet,
			Description: "",
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
	},
}

var nodeLabelsSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		keyKey: {
			Type:        schema.TypeString,
			Description: "",
			Optional:    true,
		},
		valueKey: {
			Type:        schema.TypeString,
			Description: "",
			Optional:    true,
		},
	},
}

var workerSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		machineKey: {
			Type:        schema.TypeList,
			Description: "",
			Optional:    true,
			Elem:        machineSchema,
		},
		networkKey: {
			Type:        schema.TypeList,
			Description: "",
			Optional:    true,
			Elem:        networkMachineTkgVsphereV100Schema,
		},
	},
}

var trustSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		additionalTrustedCAsKey: {
			Type:        schema.TypeList,
			Description: "",
			Optional:    true,
			Elem:        additionalTrustedCAsSchema,
		},
	},
}

var additionalTrustedCAsSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		dataKey: {
			Type:        schema.TypeString,
			Description: "",
			Required:    true,
		},
		nameKey: {
			Type:        schema.TypeString,
			Description: "",
			Required:    true,
		},
	},
}

var controlPlaneCertificateRotationSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		activateKey: {
			Type:        schema.TypeBool,
			Description: "",
			Optional:    true,
		},
		daysBeforeKey: {
			Type:        schema.TypeInt,
			Description: "",
			Optional:    true,
		},
	},
}

var proxySchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		httpProxyKey: {
			Type:        schema.TypeString,
			Description: "",
			Required:    true,
		},
		httpsProxyKey: {
			Type:        schema.TypeString,
			Description: "",
			Required:    true,
		},
		noProxyKey: {
			Type:        schema.TypeSet,
			Description: "",
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		systemWideKey: {
			Type:        schema.TypeBool,
			Description: "",
			Optional:    true,
			Default:     false,
		},
	},
}

var pciSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		controlPlaneTkgVsphereV100Key: {
			Type:        schema.TypeList,
			Description: "",
			Optional:    true,
			Elem:        pciSpecSchema,
		},
		workerKey: {
			Type:        schema.TypeList,
			Description: "",
			Optional:    true,
			Elem:        pciSpecSchema,
		},
	},
}

var pciSpecSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		devicesKey: {
			Type:        schema.TypeList,
			Description: "",
			Optional:    true,
			Elem:        devicesSchema,
		},
		hardwareVersionKey: {
			Type:        schema.TypeString,
			Description: "",
			Optional:    true,
		},
	},
}

var devicesSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		deviceIdKey: {
			Type:        schema.TypeString,
			Description: "",
			Required:    true,
		},
		vendorIdKey: {
			Type:        schema.TypeString,
			Description: "",
			Required:    true,
		},
	},
}

var imageRepositorySchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		hostKey: {
			Type:        schema.TypeString,
			Description: "",
			Optional:    true,
		},
		tlsCertificateValidationKey: {
			Type:        schema.TypeBool,
			Description: "",
			Optional:    true,
		},
	},
}

var customTDNFRepositorySchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		certificateKey: {
			Type:        schema.TypeString,
			Description: "",
			Optional:    true,
		},
	},
}

var networkTkgVsphereV100Schema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		addressesFromPoolsKey: {
			Type:        schema.TypeList,
			Description: "",
			Optional:    true,
			Elem:        addressesFromPoolsSchema,
		},
		ipv6PrimaryKey: {
			Type:        schema.TypeBool,
			Description: "",
			Optional:    true,
		},
	},
}

var addressesFromPoolsSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		apiGroupKey: {
			Type:        schema.TypeString,
			Description: "",
			Required:    true,
		},
		kindKey: {
			Type:        schema.TypeString,
			Description: "",
			Required:    true,
		},
		nameKey: {
			Type:        schema.TypeString,
			Description: "",
			Required:    true,
		},
	},
}

var auditLoggingSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		enabledKey: {
			Type:        schema.TypeBool,
			Description: "",
			Optional:    true,
		},
	},
}

var additionalImageRegistriesSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		caCertKey: {
			Type:        schema.TypeString,
			Description: "",
			Optional:    true,
		},
		hostKey: {
			Type:        schema.TypeString,
			Description: "",
			Required:    true,
		},
		skipTlsVerifyKey: {
			Type:        schema.TypeBool,
			Description: "",
			Optional:    true,
		},
	},
}

var podSecurityStandardSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		auditKey: {
			Type:        schema.TypeString,
			Description: "",
			Optional:    true,
		},
		auditVersionKey: {
			Type:        schema.TypeString,
			Description: "",
			Optional:    true,
		},
		deactivatedKey: {
			Type:        schema.TypeBool,
			Description: "",
			Optional:    true,
		},
		enforceKey: {
			Type:        schema.TypeString,
			Description: "",
			Optional:    true,
		},
		enforceVersionKey: {
			Type:        schema.TypeString,
			Description: "",
			Optional:    true,
		},
		exemptionsKey: {
			Type:        schema.TypeList,
			Description: "",
			Optional:    true,
			Elem:        exemptionsSchema,
		},
		warnKey: {
			Type:        schema.TypeString,
			Description: "",
			Optional:    true,
		},
		warnVersionKey: {
			Type:        schema.TypeString,
			Description: "",
			Optional:    true,
		},
	},
}

var exemptionsSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		namespacesKey: {
			Type:        schema.TypeSet,
			Description: "",
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
	},
}
