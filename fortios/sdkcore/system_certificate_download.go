package forticlient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type JSONSystemCertificateDownload struct {
	Mkey string `json:"mkey"`
	Type string `json:"type"`
}

// Downloads certificate from Fortigate
func (c *FortiSDKClient) ReadSystemCertificateDownload(data *JSONSystemCertificateDownload, vdomparam string) (res string, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/monitor/system/certificate/download"
	mkey := data.Mkey
	cert_type := data.Type

	params := make(map[string][]string)
	params["mkey"] = []string{mkey}
	params["type"] = []string{cert_type}

	req := c.NewRequest(HTTPMethod, path, &params, nil)
	err = req.Send3(vdomparam)
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %s", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close() //#

	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return
	}
	log.Printf("FOS-fortios reading response: %s", string(body))

	// if success returns string containing certificate so we want to parse and check it
	// didn't fail such as if requesting a cert that doesn't exist
	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)
	if result["status"] == "error" {
		err = fortiAPIErrorFormat(result, string(body))
		if err != nil {
			err = fmt.Errorf("error downloading certificate: %s", err)
			return
		}
	}

	res = string(body)

	return
}
