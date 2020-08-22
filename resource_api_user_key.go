package main

import (
	"context"
	"log"

	"github.com/arminaaki/gopastebin"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAPIUserKeyCreate(d *schema.ResourceData, m interface{}) error {

	client, err := gopastebin.NewClient(&gopastebin.AccountRequest{
		APIDevKey:       m.(*Config).ApiDevKey,
		APIUserName:     m.(*Config).ApiUserName,
		APIUserPassword: m.(*Config).ApiUserPassword,
	})
	if err != nil {
		return err
	}
	d.SetId(d.Get("name").(string))
	if err := d.Set("api_user_key", client.APIUserKey); err != nil {
		return err
	}

	return resourceAPIUserKeyRead(d, m)
}

func resourceAPIUserKeyRead(d *schema.ResourceData, m interface{}) error {

	client, err := gopastebin.NewClient(&gopastebin.AccountRequest{
		APIDevKey:       m.(*Config).ApiDevKey,
		APIUserName:     m.(*Config).ApiUserName,
		APIUserPassword: m.(*Config).ApiUserPassword,
	})
	if err != nil {
		return err
	}

	APIUserKey := d.Get("api_user_key").(string)
	if APIUserKey == "" {
		log.Println("Emty api_user_key")
		return nil
	}

	if _, _, err := client.Paste.List(context.TODO(), &gopastebin.PasteListRequest{}); err != nil {
		return err
	}

	return nil
}

func resourceAPIUserKeyUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceAPIUserKeyRead(d, m)
}

func resourceAPIUserKeyDelete(d *schema.ResourceData, m interface{}) error {
	if _, err := gopastebin.NewClient(&gopastebin.AccountRequest{
		APIDevKey:       m.(*Config).ApiDevKey,
		APIUserName:     m.(*Config).ApiUserName,
		APIUserPassword: m.(*Config).ApiUserPassword,
	}); err != nil {
		return err
	}

	return nil
}

func resourceAPIUserKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceAPIUserKeyCreate,
		Read:   resourceAPIUserKeyRead,
		Update: resourceAPIUserKeyUpdate,
		Delete: resourceAPIUserKeyDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default",
			},
			"api_user_key": &schema.Schema{
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
		},
	}
}
