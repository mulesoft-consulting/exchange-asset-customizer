package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"

	"github.com/monaco-io/request"
	auth "github.com/mulesoft-consulting/cloudhub-client-go/authorization"
)

type ExchangeClient struct {
	ctx          *context.Context
	username     string
	pwd          string
	access_token string
	baseurl      string
}

func getServer(index int) string {
	if index == 0 {
		return "https://anypoint.mulesoft.com"
	} else if index == 1 {
		return "https://eu1.anypoint.mulesoft.com"
	}
	return ""
}

func newExchangeClient(ctx context.Context, username string, pwd string) *ExchangeClient {
	return &ExchangeClient{
		ctx:      &ctx,
		username: username,
		pwd:      pwd,
		baseurl:  getServer(ctx.Value(auth.ContextServerIndex).(int)),
	}
}

func (c *ExchangeClient) login() error {
	creds := auth.NewUserPwdCredentials()
	creds.SetUsername(c.username)
	creds.SetPassword(c.pwd)
	//authenticate
	cfgauth := auth.NewConfiguration()
	authclient := auth.NewAPIClient(cfgauth)
	req := authclient.DefaultApi.LoginPost(*c.ctx).UserPwdCredentials(*creds)
	authres, httpr, err := req.Execute()

	if err != nil {
		var details string
		if httpr != nil {
			b, _ := ioutil.ReadAll(httpr.Body)
			details = string(b)
		} else {
			details = err.Error()
		}
		return errors.New(details)
	}
	defer httpr.Body.Close()
	c.access_token = authres.GetAccessToken()

	return nil
}

func (c *ExchangeClient) addCustomField(organizationId string, dataType string, displayName string, tagKey string) error {
	u, err := url.Parse(fmt.Sprintf("%s/exchange/api/v2/organizations/%s/fields", c.baseurl, url.PathEscape(organizationId)))
	if err != nil {
		return err
	}

	client := request.Client{
		Context: *c.ctx,
		URL:     u.String(),
		Method:  request.POST,
		Bearer:  c.access_token,
		JSON: CustomFieldKeyBody{
			DataType:    dataType,
			DisplayName: displayName,
			TagKey:      tagKey,
		},
	}
	resp := client.Send()
	if !resp.OK() {
		return resp.Error()
	}
	var result interface{}
	if err := resp.Scan(&result).Error(); err != nil {
		return err
	}
	respBody, _ := json.Marshal(result)
	statusCode := resp.Response().StatusCode
	if statusCode == 409 {
		//already exists skipping
		return nil
	} else if statusCode >= 400 {
		var details string
		message := "[ERROR] Post custom-field responded with status " + fmt.Sprint(statusCode)
		details = string(respBody)
		if details != "" {
			message += "\ndetails: \n" + details
		}
		return errors.New(message)
	}
	defer resp.Close()

	return nil
}

func (c *ExchangeClient) putCustomFieldValue(organizationId string, assetName string, version string, key string, value string) error {
	u, err := url.Parse(fmt.Sprintf("%s/exchange/api/v2/assets/%s/%s/%s/tags/fields/%s", c.baseurl, url.PathEscape(organizationId), url.PathEscape(assetName), url.PathEscape(version), url.PathEscape(key)))
	if err != nil {
		return err
	}

	client := request.Client{
		Context: *c.ctx,
		URL:     u.String(),
		Method:  request.PUT,
		Bearer:  c.access_token,
		JSON: &CustomFieldValueBody{
			TagValue: value,
		},
	}
	resp := client.Send()
	if !resp.OK() {
		return resp.Error()
	}
	var result interface{}
	if err := resp.Scan(&result).Error(); err != nil {
		return err
	}
	respBody, _ := json.Marshal(result)

	if statusCode := resp.Response().StatusCode; statusCode >= 400 {
		var details string
		message := "[ERROR] PUT custom-field-value responded with status " + fmt.Sprint(statusCode)
		details = string(respBody)
		if details != "" {
			message += "\ndetails : \n" + details
		}
		return errors.New(message)
	}
	defer resp.Close()

	return nil
}

func (c *ExchangeClient) createCustomCategory(organizationId string, assetName string, version string, key string, value []string) error {
	u, err := url.Parse(fmt.Sprintf("%s/exchange/api/v2/assets/%s/%s/%s/tags/categories/%s", c.baseurl, url.PathEscape(organizationId), url.PathEscape(assetName), url.PathEscape(version), url.PathEscape(key)))
	if err != nil {
		return err
	}
	client := request.Client{
		Context: *c.ctx,
		URL:     u.String(),
		Method:  request.PUT,
		Bearer:  c.access_token,
		JSON: &CustomCategoryValueBody{
			TagValue: value,
		},
	}
	resp := client.Send()
	if !resp.OK() {
		return resp.Error()
	}
	var result interface{}
	if err := resp.Scan(&result).Error(); err != nil {
		return err
	}
	respBody, _ := json.Marshal(result)
	if statusCode := resp.Response().StatusCode; statusCode >= 400 {
		var details string
		message := "[ERROR] PUT custom-field-value responded with status " + fmt.Sprint(statusCode)
		details = string(respBody)
		if details != "" {
			message += "\ndetails : \n" + details
		}
		return errors.New(message)
	}
	defer resp.Close()

	return nil
}

func (c *ExchangeClient) patchAssetAttributes(organizationId string, assetName string, key string, value string) error {
	u, err := url.Parse(fmt.Sprintf("%s/exchange/api/v2/assets/%s/%s", c.baseurl, url.PathEscape(organizationId), url.PathEscape(assetName)))
	if err != nil {
		return err
	}
	client := request.Client{
		Context: *c.ctx,
		URL:     u.String(),
		Method:  request.PATCH,
		Bearer:  c.access_token,
		JSON:    "{\"" + key + "\":\"" + value + "\"}",
	}
	resp := client.Send()
	if !resp.OK() {
		return resp.Error()
	}
	var result interface{}
	if err := resp.Scan(&result).Error(); err != nil {
		return err
	}
	respBody, _ := json.Marshal(result)
	if statusCode := resp.Response().StatusCode; statusCode >= 400 {
		var details string
		message := "[ERROR] PATCH attributes responded with status " + fmt.Sprint(statusCode)
		details = string(respBody)
		if details != "" {
			message += "\ndetails : \n" + details
		}
		return errors.New(message)
	}
	defer resp.Close()

	return nil
}
