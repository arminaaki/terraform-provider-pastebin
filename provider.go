package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_dev_key": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The api_dev_key for API operations.",
			},
			"api_user_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The api_user_name for API operations.",
			},
			"api_user_password": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The api_user_password for API operations.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"pastebin_api_user_key": resourceAPIUserKey(),
			"pastebin_create_paste": resourceCreatePaste(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		ApiUserName:     d.Get("api_user_name").(string),
		ApiUserPassword: d.Get("api_user_password").(string),
		ApiDevKey:       d.Get("api_dev_key").(string),
		BaseUrl:         "https://pastebin.com/",
	}

	return &config, nil
}
