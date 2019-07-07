package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/hashicorp/terraform/helper/schema"
	"gopkg.in/resty.v1"
)

func resourceCreatePasteCreate(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	log.Println("HI FROM resourceCreatePasteCreate LALALALALALA")
	// apiKey, responseError := createAPIKey(config.ApiDevKey, config.ApiUserName, config.ApiUserPassword, config.BaseUrl)
	log.Println(createPasteAPIParams(d))

	pasteURL, responseError := pasteOperation(createPasteAPIParams(d), config.BaseUrl)
	if responseError != nil {
		return responseError
	}
	d.SetId(pasteURL)

	return nil
}

func resourceCreatePasteRead(d *schema.ResourceData, m interface{}) error {
	_, responseError := readPaste(d.Id())
	if responseError != nil {
		d.SetId("")
		log.Println("[DEBUG] invalid paste url")
	}

	return nil
}

func resourceCreatePasteUpdate(d *schema.ResourceData, m interface{}) error { return nil }

func resourceCreatePaste() *schema.Resource {
	return &schema.Resource{
		Create: resourceCreatePasteCreate,
		Read:   resourceCreatePasteRead,
		Update: resourceCreatePasteUpdate,
		Delete: resourceCreatePasteDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default",
			},
			"api_dev_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"api_paste_code": {
				Type:     schema.TypeString,
				Required: true,
			},
			"api_paste_private": {
				Type:     schema.TypeString,
				Required: true,
			},
			"api_paste_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"api_paste_expire_date": {
				Type:     schema.TypeString,
				Required: true,
			},
			"api_paste_format": {
				Type:     schema.TypeString,
				Required: true,
			},
			"api_user_key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createPasteAPIParams(d *schema.ResourceData) map[string]string {
	return map[string]string{
		"api_user_key":          d.Get("api_user_key").(string),
		"api_dev_key":           d.Get("api_dev_key").(string),
		"api_paste_code":        d.Get("api_paste_code").(string),
		"api_paste_private":     d.Get("api_paste_private").(string),
		"api_paste_name":        d.Get("api_paste_name").(string),
		"api_paste_expire_date": d.Get("api_paste_expire_date").(string),
		"api_paste_format":      d.Get("api_paste_format").(string),
		"api_option":            "paste",
	}
}

func pasteOperation(createPasteAPIParamas map[string]string, baseURL string) (string, error) {
	resp, err := resty.SetRetryCount(3).
		R().
		SetFormData(createPasteAPIParamas).
		Post(urlBuilder("api/api_post.php", baseURL))

	if err != nil {
		log.Printf("Request error: %s\n", err)
		return "", err
	}

	return string(resp.Body()), nil
}

func readPaste(pasteURL string) (string, error) {
	resp, err := resty.SetRetryCount(3).
		R().
		Get(pasteURL)

	if err != nil {
		log.Printf("Request error: %s\n", err)
		return "", err
	}

	if resp.StatusCode() == http.StatusNotFound {
		return "", errors.New("Paste Not Found")

	}

	return string(resp.Body()), nil
}
