package fmgclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
    "time"
)

// Request represents a JSON-RPC request sent by a client.
type Request struct {
	Id      uint64         `json:"id"`
	Method  string         `json:"method"`
	Session string         `json:"session"`
	Params  [1]interface{} `json:"params"`
}

// FmgSDKClient is for communicating with FortiManager
type FmgSDKClient struct {
	Ipaddr string
	User   string
	Passwd string
	Debug  string
	Client *http.Client
	SessionString	string
	DebugNum	int
	Init	bool
}

// NewClient is for creating new client
// Input:
//   @ip: ipaddress of fortimanager
//   @user: username of fortimanager
//   @passwd: passwd of fortimanager
//   @client: http client used
// Output:
//   @FmgSDKClient: fmg client
func NewClient(ip, user, passwd string, client *http.Client) *FmgSDKClient {
	d := os.Getenv("TRACEDEBUG")

	return &FmgSDKClient{
		Ipaddr: ip,
		User:   user,
		Passwd: passwd,
		Client: client,
		Debug:  d,
		Init: false,
	}
}

// NewEmptyClient initializes a new global empty plugin FmgSDKClient
// It returns the created client object
func NewEmptyClient() *FmgSDKClient {
	return &FmgSDKClient{
		Init: false,
	}
}


// Execute is for sending the http request to FortiManager
// Input:
//   @req: http request
// Output:
//   @result: http response result
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) Execute(req *Request) (result map[string]interface{}, err error) {
	j, _ := json.Marshal(req)

	if c.Debug == "ON" || c.Debug == "on" {
		log.Printf("[TRACEDEBUG] ==> request: %s", j)
	}

	httpResp, err := c.Client.Post("https://"+c.Ipaddr+"/jsonrpc", "application/json", bytes.NewBuffer(j))
	if err != nil {
		err = fmt.Errorf("Login failed: %s", err)
		return
	}
	defer httpResp.Body.Close()

	// check response
	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %s", err)
		return
	}

	if c.Debug == "ON" || c.Debug == "on" {
		log.Printf("[TRACEDEBUG] ==> result: %s", body)
	}

	result = map[string]interface{}{}
	json.Unmarshal([]byte(string(body)), &result)

	if result != nil {
		if len(result) == 0 {
			err = fmt.Errorf("No result got, details:\n%s", body)
			return
		}

		if id := uint64(result["id"].(float64)); id != req.Id {
			err = fmt.Errorf("id not match, should be 1, but is %d", id)
			return
		}

		if result["result"] != nil {
			status := (result["result"].([]interface{}))[0].(map[string]interface{})["status"].(map[string]interface{})

			c := uint(status["code"].(float64))
			m := status["message"].(string)
			if c != 0 || m != "OK" {
				err = fmt.Errorf("status not right: code is %d, message is %s", c, m)
				return
			}
		} else {
			err = fmt.Errorf("can't get response status: %s", err)
			return
		}
	}

	return
}


func getEnvSession() (session string) {
	fmgTestacc := os.Getenv("FORTIOS_FMG_TESTACC")

	if fmgTestacc == "" {
		session = ""
		return
	}

	sessionString := os.Getenv("FORTIOS_FMG_SESSION")
	sessionTimeStamp := os.Getenv("FORTIOS_FMG_SESSIONTIMESTAMP")
	if sessionString == "" || sessionTimeStamp == "" {
		session = ""
		return
	}

	lasttime, err := time.Parse(time.RFC3339, sessionTimeStamp)
	if err != nil {
		session = ""
		return
	}

	now := time.Now()
	subM := now.Sub(lasttime)

	if subM.Minutes() > 15 {
		session = ""
		return
	}

	return sessionString
}

func setEnvSession(session string) {
	fmgTestacc := os.Getenv("FORTIOS_FMG_TESTACC")

	if fmgTestacc == "" {
		return
	}

	now := time.Now()

	os.Setenv("FORTIOS_FMG_SESSION", session)
	os.Setenv("FORTIOS_FMG_SESSIONTIMESTAMP", now.Format(time.RFC3339))
}

// Login is for logging in
// Output:
//   @session: login session
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) Login() (session string, err error) {
	sessionEnvString := getEnvSession()
	if sessionEnvString != "" {
		session = sessionEnvString
		return
	}

	params := map[string]interface{}{
		"data": map[string]string{
			"user":   c.User,
			"passwd": c.Passwd,
		},
		"url": "/sys/login/user",
	}

	req := &Request{
		Id:     1,
		Method: "exec",
		Params: [1]interface{}{params},
	}

	result, err := c.Execute(req)
	if err != nil {
		return "", fmt.Errorf("login failed:%s", err)
	}

	session = result["session"].(string)

	setEnvSession(session)
	return
}


// Logout is for logging out
// Input:
//   @s: login session
// Output:
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) Logout(s string) (err error) {
	params := map[string]interface{}{
		"url": "/sys/logout",
	}

	req := &Request{
		Id:      1,
		Method:  "exec",
		Params:  [1]interface{}{params},
		Session: s,
	}

	_, err = c.Execute(req)
	if err != nil {
		err = fmt.Errorf("logout failed:%s", err)
		return
	}
	return
}

// Do is for executing the related request
// Input:
//   @method: operation method
//   @params: request infor needed
// Output:
//   @output: result infor got back
//   @err: error details if failure, and nil if success
func (c *FmgSDKClient) Do(method string, params map[string]interface{}) (output map[string]interface{}, err error) {
	if c.Init == false {
		err = fmt.Errorf("FortiManager connection did not initialize successfully!")
		return nil, err
	}
	req := &Request{
		Id:      1,
		Method:  method,
		Params:  [1]interface{}{params},
		Session: c.SessionString,//session,
	}

	output, err = c.Execute(req)
	return
}

// Trace is for debugging
// Input:
//   @s: function name to be traced
// Output:
//   @func()
func (f *FmgSDKClient) Trace(s string) func() {
	if f.Debug == "ON" {
		log.Printf("[TRACEDEBUG] -> Enter %s <-", s)
		return func() { log.Printf("[TRACEDEBUG]    -> Leave %s <-", s) }
	}

	return func() {}
}
