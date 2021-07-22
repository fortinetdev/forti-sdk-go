package forticlient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type PolicyLookupResponse struct {
	HTTP_Method string `json:"http_method"`
}

type PolicyLookupResult struct {
	Success  bool   `json:"success"`
	PolicyID int    `json:"policy_id"`
	DestIp   string `json:"destip"`
}
type JSONPolicyLookupOutput struct {
	Vdom       string             `json:"vdom"`
	Mkey       string             `json:"mkey"`
	Status     string             `json:"status"`
	HTTPStatus float64            `json:"http_status"`
	Results    PolicyLookupResult `json:"results"`
}

type PolicyLookupRequest struct {
	Destination     string `json:"dest"`
	DestPort        string `json:"destport"`
	IPVersion       string `json:"ipVersion"`
	Protocol        string `json:"protocol"`
	SourceIP        string `json:"sourceip"`
	SourceInterface string `json:"srcintf"`
}

func (c *FortiSDKClient) ReadFirewallPolicyLookup(params *PolicyLookupRequest) (lookupResult *PolicyLookupResult, err error) {
	HTTPMethod := "GET"

	path := "/api/v2/monitor/firewall/policy-lookup"
	path += "?dest=" + EscapeURLString(params.Destination)
	path += "&destport=" + EscapeURLString(params.DestPort)
	path += "&ipVersion=ipv4&ipv6=false&protocol=tcp&sourceip=" + EscapeURLString(params.SourceIP)
	path += "&srcintf=" + EscapeURLString(params.SourceInterface)

	lookupResult = &PolicyLookupResult{}

	if err != nil {
		log.Fatal(err)
		return
	}

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.Send()
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

	//var result map[string]interface{}
	var result JSONPolicyLookupOutput
	json.Unmarshal([]byte(string(body)), &result)
	if result.Status != "success" {
		return
	}

	return &result.Results, nil
}
