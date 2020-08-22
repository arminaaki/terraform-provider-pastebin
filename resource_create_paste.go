package main

import (
	"log"
	"regexp"

	"context"

	"github.com/arminaaki/gopastebin"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceCreatePasteCreate(d *schema.ResourceData, m interface{}) error {
	client, _ := gopastebin.NewClient(&gopastebin.AccountRequest{
		APIDevKey:       m.(*Config).ApiDevKey,
		APIUserName:     m.(*Config).ApiUserName,
		APIUserPassword: m.(*Config).ApiUserPassword,
	})

	result, _, err := client.Paste.Create(context.TODO(), createPasteAPIParams(d))
	if err != nil {
		panic(err)
	}

	d.SetId(result.PasteURL)

	return nil
}

func resourceCreatePasteRead(d *schema.ResourceData, m interface{}) error {

	client, _ := gopastebin.NewClient(&gopastebin.AccountRequest{
		APIDevKey:       m.(*Config).ApiDevKey,
		APIUserName:     m.(*Config).ApiUserName,
		APIUserPassword: m.(*Config).ApiUserPassword,
	})
	_, _, err := client.Paste.GetRaw(context.TODO(), &gopastebin.PasteGetRawRequest{})
	if err != nil {
		d.SetId("")
		log.Println("[DEBUG] invalid paste url")
	}

	return nil
}

func resourceCreatePasteDelete(d *schema.ResourceData, m interface{}) error {
	client, _ := gopastebin.NewClient(&gopastebin.AccountRequest{
		APIDevKey:       m.(*Config).ApiDevKey,
		APIUserName:     m.(*Config).ApiUserName,
		APIUserPassword: m.(*Config).ApiUserPassword,
	})

	_, err := client.Paste.Delete(context.TODO(), deletePasteAPIParams(d))
	if err != nil {
		return err
	}

	return nil
}

func resourceCreatePaste() *schema.Resource {
	return &schema.Resource{
		Create: resourceCreatePasteCreate,
		Read:   resourceCreatePasteRead,
		Delete: resourceCreatePasteDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "default",
			},
			"api_dev_key": {
				Type:      schema.TypeString,
				Optional:  true,
				ForceNew:  true,
				Sensitive: true,
			},
			"api_paste_code": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"api_paste_private": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"api_paste_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"api_paste_expire_date": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"api_paste_format": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"api_user_key": {
				Type:      schema.TypeString,
				Optional:  true,
				ForceNew:  true,
				Sensitive: true,
			},
		},
	}
}

func createPasteAPIParams(d *schema.ResourceData) *gopastebin.PasteRequest {

	return &gopastebin.PasteRequest{
		APIPasteName:       d.Get("api_paste_name").(string),
		APIPasteCode:       d.Get("api_paste_code").(string),
		APIPasteFormat:     d.Get("api_paste_format").(string),
		APIPastePrivate:    d.Get("api_paste_private").(string),
		APIPasteExpireDate: d.Get("api_paste_expire_date").(string),
	}
}

func deletePasteAPIParams(d *schema.ResourceData) *gopastebin.PasteDeleteRequest {
	re := regexp.MustCompile(`https:\/\/pastebin\.com\/(.+)`)

	return &gopastebin.PasteDeleteRequest{
		APIPasteKey: re.FindStringSubmatch(d.Id())[1],
	}
}
