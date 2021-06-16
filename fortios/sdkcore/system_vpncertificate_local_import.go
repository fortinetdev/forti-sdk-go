package forticlient

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type JSONSystemVpnCertificateLocalImport struct {
	Name        string `json:"certname"`
	Type        string `json:"type"`
	Certificate string `json:"file_content"`
	Password    string `json:"password,omitempty"`
	PrivateKey  string `json:"key_file_content,omitempty"`
	Scope       string `json:"scope"`
}

// Uploads certificate to Fortigate
func (c *FortiSDKClient) CreateSystemVpnCertificateLocalImport(params *JSONSystemVpnCertificateLocalImport) (res string, err error) {
	HTTPMethod := "POST"
	path := "/api/v2/monitor/vpn-certificate/local/import"
	params.Certificate = base64.StdEncoding.EncodeToString([]byte(params.Certificate))
	if params.PrivateKey != "" {
		params.PrivateKey = base64.StdEncoding.EncodeToString([]byte(params.PrivateKey))
	}
	locJSON, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	bytes := bytes.NewBuffer(locJSON)
	req := c.NewRequest(HTTPMethod, path, nil, bytes)
	err = req.Send()
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %s", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close()

	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return
	}
	log.Printf("FOS-fortios reading response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)
	if result["status"] == "error" {
		err = fortiAPIErrorFormat(result, string(body))
		if err != nil {
			err = fmt.Errorf("error importing certificate: %s", err)
			return
		}
	}

	res = string(body)

	return
}
