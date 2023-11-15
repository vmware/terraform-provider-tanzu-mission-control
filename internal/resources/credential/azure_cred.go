package credential

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var azureCredSpec = &schema.Schema{
	Type:        schema.TypeList,
	Optional:    true,
	MaxItems:    1,
	Description: "Azure credential",
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			servicePrincipalKey:         azureServicePrincipalSpec,
			servicePrincipalWithCertKey: azureServicePrincipalWithCertSpec,
		},
	},
}

var azureServicePrincipalSpec = &schema.Schema{
	Type:        schema.TypeList,
	Optional:    true,
	MaxItems:    1,
	Description: "Azure service principal",
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			subscriptionIDKey: {
				Description: "Subscription ID of the Azure credential",
				Type:        schema.TypeString,
				Required:    true,
			},
			tenantIDKey: {
				Description: "Tenant ID of the Azure credential",
				Type:        schema.TypeString,
				Required:    true,
			},
			resourceGroupKey: {
				Description:  "Resource Group name",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 90),
			},
			clientIDKey: {
				Description: "Client ID of the Service Principal",
				Type:        schema.TypeString,
				Required:    true,
			},
			clientSecretKey: {
				Description: "Client Secret of the Service Principal",
				Type:        schema.TypeString,
				Optional:    true,
			},
			azureCloudNameKey: {
				Description: "Azure Cloud name",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	},
}

var azureServicePrincipalWithCertSpec = &schema.Schema{
	Type:        schema.TypeList,
	Optional:    true,
	MaxItems:    1,
	Description: "Azure service principal",
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			subscriptionIDKey: {
				Description: "Subscription ID of the Azure credential",
				Type:        schema.TypeString,
				Required:    true,
			},
			tenantIDKey: {
				Description: "Tenant ID of the Azure credential",
				Type:        schema.TypeString,
				Required:    true,
			},
			clientIDKey: {
				Description: "Client ID of the Service Principal",
				Type:        schema.TypeString,
				Required:    true,
			},
			clientCertificateKey: {
				Description: "Client certificate of the Service Principal",
				Type:        schema.TypeString,
				Required:    true,
			},
			azureCloudNameKey: {
				Description: "Azure Cloud name",
				Type:        schema.TypeString,
				Optional:    true,
			},
			managedSubscriptionsKey: {
				Description: "IDs of the Azure Subscriptions that the Service Principal can manage",
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	},
}
