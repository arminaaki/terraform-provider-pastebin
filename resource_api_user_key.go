package main

import (
	"errors"
	"fmt"
	"log"
	"regexp"

	"github.com/hashicorp/terraform/helper/schema"
	"gopkg.in/resty.v1"
)

func resourceAPIUserKeyCreate(d *schema.ResourceData, m interface{}) error {

	config := m.(*Config)
	apiKey, responseError := createAPIKey(config.ApiDevKey, config.ApiUserName, config.ApiUserPassword, config.BaseUrl)

	if responseError != nil {
		return responseError
	}
	d.SetId(d.Get("name").(string))

	err := d.Set("api_user_key", apiKey)
	if err != nil {
		return err
	}

	return resourceAPIUserKeyRead(d, m)
}

func resourceAPIUserKeyRead(d *schema.ResourceData, m interface{}) error {

	config := m.(*Config)
	APIUserKey := d.Get("api_user_key").(string)
	if APIUserKey == "" {
		log.Println("Emty api_user_key")
		return nil
	}
	listPastesParams := map[string]string{
		"api_dev_key":       config.ApiDevKey,
		"api_user_key":      APIUserKey,
		"api_results_limit": "1",
		"api_option":        "list",
	}
	resp, err := resty.SetRetryCount(3).
		R().
		SetFormData(listPastesParams).
		Post(urlBuilder("api/api_post.php", config.BaseUrl))

	if err != nil {
		log.Printf("Request error: %s\n", err)
		return err
	}

	err2 := validateAPIDevKey(string(resp.Body()))
	if err2 != nil {
		log.Println(err2)
		return err2
	}

	// If api_user_key is invalid, trigger creation of the key
	err3 := validateAPIUserKey(string(resp.Body()))
	if err3 != nil {
		log.Println(err3)
		d.SetId("")
		return nil
	}

	return nil
}

func resourceAPIUserKeyUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceAPIUserKeyRead(d, m)
}

func resourceAPIUserKeyDelete(d *schema.ResourceData, m interface{}) error {
	//Invalidate the existing api_user_key by creating a new one
	config := m.(*Config)
	_, err := createAPIKey(config.ApiDevKey, config.ApiUserName, config.ApiUserPassword, config.BaseUrl)

	if err != nil {
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

func validateAPIDevKey(response string) error {
	r := regexp.MustCompile("^Bad API request, invalid api_dev_key")
	// Check response body:
	if r.MatchString(response) {
		return errors.New(response)
	}

	return nil
}

func validateAPIUserKey(response string) error {
	r := regexp.MustCompile("^Bad API request, invalid api_user_key")
	// Check response body:
	if r.MatchString(response) {
		return errors.New(response)
	}

	return nil
}

func urlBuilder(path string, BaseURL string) string {
	return fmt.Sprintf("%s/%s", BaseURL, path)
}

func createAPIKey(APIDevKey string, APIUserName string, APIUserPassword string, BaseURL string) (string, error) {
	createAPIKeyParams := map[string]string{
		"api_dev_key":       APIDevKey,
		"api_user_name":     APIUserName,
		"api_user_password": APIUserPassword,
	}

	resp, err := resty.SetRetryCount(3).
		R().
		SetFormData(createAPIKeyParams).
		Post(urlBuilder("api/api_login.php", BaseURL))

	if err != nil {
		log.Printf("Request error: %s\n", err)
		return "", err
	}

	responseBodyString := string(resp.Body())

	err2 := validateAPIDevKey(responseBodyString)
	if err2 != nil {
		log.Println(err2)
		return "", err2
	}

	// If api_user_key is invalid, trigger creation of the key
	err3 := validateAPIUserKey(responseBodyString)
	if err3 != nil {
		log.Println(err3)
		return "", nil
	}

	return responseBodyString, nil
}
