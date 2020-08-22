// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Description: SDK for FortiOS Provider

package forticlient

type creatUpdateOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}



// UpdateSystemGlobal API operation for FortiOS updates the specified Global.
// Returns the index value of the Global and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - global chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemGlobal(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/global"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemGlobal API operation for FortiOS deletes the specified Global.
// Returns error for service API and SDK errors.
// See the system - global chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemGlobal(mkey string) (err error) {

	//No unset API for system - global
	return
}

// ReadSystemGlobal API operation for FortiOS gets the Global
// with the specified index value.
// Returns the requested Global value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - global chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemGlobal(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/global"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemAccprofile API operation for FortiOS creates a new Accprofile.
// Returns the index value of the Accprofile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - accprofile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemAccprofile(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/accprofile"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemAccprofile API operation for FortiOS updates the specified Accprofile.
// Returns the index value of the Accprofile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - accprofile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAccprofile(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/accprofile"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemAccprofile API operation for FortiOS deletes the specified Accprofile.
// Returns error for service API and SDK errors.
// See the system - accprofile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAccprofile(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/accprofile"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemAccprofile API operation for FortiOS gets the Accprofile
// with the specified index value.
// Returns the requested Accprofile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - accprofile chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAccprofile(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/accprofile"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemObjectTagging API operation for FortiOS creates a new Object Tagging.
// Returns the index value of the Object Tagging and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - object-tagging chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemObjectTagging(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/object-tagging"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemObjectTagging API operation for FortiOS updates the specified Object Tagging.
// Returns the index value of the Object Tagging and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - object-tagging chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemObjectTagging(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/object-tagging"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemObjectTagging API operation for FortiOS deletes the specified Object Tagging.
// Returns error for service API and SDK errors.
// See the system - object-tagging chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemObjectTagging(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/object-tagging"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemObjectTagging API operation for FortiOS gets the Object Tagging
// with the specified index value.
// Returns the requested Object Tagging value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - object-tagging chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemObjectTagging(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/object-tagging"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemInterface API operation for FortiOS creates a new Interface.
// Returns the index value of the Interface and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemInterface(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/interface"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemInterface API operation for FortiOS updates the specified Interface.
// Returns the index value of the Interface and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemInterface(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/interface"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemInterface API operation for FortiOS deletes the specified Interface.
// Returns error for service API and SDK errors.
// See the system - interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemInterface(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/interface"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemInterface API operation for FortiOS gets the Interface
// with the specified index value.
// Returns the requested Interface value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemInterface(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/interface"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateSystemPasswordPolicy API operation for FortiOS updates the specified Password Policy.
// Returns the index value of the Password Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - password-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemPasswordPolicy(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/password-policy"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemPasswordPolicy API operation for FortiOS deletes the specified Password Policy.
// Returns error for service API and SDK errors.
// See the system - password-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemPasswordPolicy(mkey string) (err error) {

	//No unset API for system - password-policy
	return
}

// ReadSystemPasswordPolicy API operation for FortiOS gets the Password Policy
// with the specified index value.
// Returns the requested Password Policy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - password-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemPasswordPolicy(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/password-policy"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateSystemPasswordPolicyGuestAdmin API operation for FortiOS updates the specified Password Policy Guest Admin.
// Returns the index value of the Password Policy Guest Admin and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - password-policy-guest-admin chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemPasswordPolicyGuestAdmin(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/password-policy-guest-admin"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemPasswordPolicyGuestAdmin API operation for FortiOS deletes the specified Password Policy Guest Admin.
// Returns error for service API and SDK errors.
// See the system - password-policy-guest-admin chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemPasswordPolicyGuestAdmin(mkey string) (err error) {

	//No unset API for system - password-policy-guest-admin
	return
}

// ReadSystemPasswordPolicyGuestAdmin API operation for FortiOS gets the Password Policy Guest Admin
// with the specified index value.
// Returns the requested Password Policy Guest Admin value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - password-policy-guest-admin chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemPasswordPolicyGuestAdmin(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/password-policy-guest-admin"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemSmsServer API operation for FortiOS creates a new Sms Server.
// Returns the index value of the Sms Server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sms-server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemSmsServer(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/sms-server"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemSmsServer API operation for FortiOS updates the specified Sms Server.
// Returns the index value of the Sms Server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sms-server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSmsServer(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/sms-server"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemSmsServer API operation for FortiOS deletes the specified Sms Server.
// Returns error for service API and SDK errors.
// See the system - sms-server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSmsServer(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/sms-server"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemSmsServer API operation for FortiOS gets the Sms Server
// with the specified index value.
// Returns the requested Sms Server value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sms-server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSmsServer(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/sms-server"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemAdmin API operation for FortiOS creates a new Admin.
// Returns the index value of the Admin and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemAdmin(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/admin"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemAdmin API operation for FortiOS updates the specified Admin.
// Returns the index value of the Admin and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAdmin(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/admin"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemAdmin API operation for FortiOS deletes the specified Admin.
// Returns error for service API and SDK errors.
// See the system - admin chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAdmin(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/admin"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemAdmin API operation for FortiOS gets the Admin
// with the specified index value.
// Returns the requested Admin value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAdmin(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/admin"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemApiUser API operation for FortiOS creates a new Api User.
// Returns the index value of the Api User and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - api-user chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemApiUser(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/api-user"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemApiUser API operation for FortiOS updates the specified Api User.
// Returns the index value of the Api User and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - api-user chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemApiUser(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/api-user"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemApiUser API operation for FortiOS deletes the specified Api User.
// Returns error for service API and SDK errors.
// See the system - api-user chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemApiUser(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/api-user"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemApiUser API operation for FortiOS gets the Api User
// with the specified index value.
// Returns the requested Api User value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - api-user chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemApiUser(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/api-user"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemSitTunnel API operation for FortiOS creates a new Sit Tunnel.
// Returns the index value of the Sit Tunnel and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sit-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemSitTunnel(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/sit-tunnel"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemSitTunnel API operation for FortiOS updates the specified Sit Tunnel.
// Returns the index value of the Sit Tunnel and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sit-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSitTunnel(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/sit-tunnel"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemSitTunnel API operation for FortiOS deletes the specified Sit Tunnel.
// Returns error for service API and SDK errors.
// See the system - sit-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSitTunnel(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/sit-tunnel"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemSitTunnel API operation for FortiOS gets the Sit Tunnel
// with the specified index value.
// Returns the requested Sit Tunnel value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sit-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSitTunnel(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/sit-tunnel"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateSystemFssoPolling API operation for FortiOS updates the specified Fsso Polling.
// Returns the index value of the Fsso Polling and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fsso-polling chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemFssoPolling(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/fsso-polling"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemFssoPolling API operation for FortiOS deletes the specified Fsso Polling.
// Returns error for service API and SDK errors.
// See the system - fsso-polling chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemFssoPolling(mkey string) (err error) {

	//No unset API for system - fsso-polling
	return
}

// ReadSystemFssoPolling API operation for FortiOS gets the Fsso Polling
// with the specified index value.
// Returns the requested Fsso Polling value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fsso-polling chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemFssoPolling(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/fsso-polling"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateSystemHa API operation for FortiOS updates the specified Ha.
// Returns the index value of the Ha and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ha chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemHa(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/ha"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemHa API operation for FortiOS deletes the specified Ha.
// Returns error for service API and SDK errors.
// See the system - ha chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemHa(mkey string) (err error) {

	//No unset API for system - ha
	return
}

// ReadSystemHa API operation for FortiOS gets the Ha
// with the specified index value.
// Returns the requested Ha value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ha chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemHa(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/ha"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateSystemHaMonitor API operation for FortiOS updates the specified Ha Monitor.
// Returns the index value of the Ha Monitor and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ha-monitor chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemHaMonitor(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/ha-monitor"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemHaMonitor API operation for FortiOS deletes the specified Ha Monitor.
// Returns error for service API and SDK errors.
// See the system - ha-monitor chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemHaMonitor(mkey string) (err error) {

	//No unset API for system - ha-monitor
	return
}

// ReadSystemHaMonitor API operation for FortiOS gets the Ha Monitor
// with the specified index value.
// Returns the requested Ha Monitor value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ha-monitor chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemHaMonitor(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/ha-monitor"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemArpTable API operation for FortiOS creates a new Arp Table.
// Returns the index value of the Arp Table and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - arp-table chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemArpTable(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/arp-table"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemArpTable API operation for FortiOS updates the specified Arp Table.
// Returns the index value of the Arp Table and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - arp-table chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemArpTable(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/arp-table"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemArpTable API operation for FortiOS deletes the specified Arp Table.
// Returns error for service API and SDK errors.
// See the system - arp-table chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemArpTable(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/arp-table"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemArpTable API operation for FortiOS gets the Arp Table
// with the specified index value.
// Returns the requested Arp Table value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - arp-table chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemArpTable(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/arp-table"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemIpv6NeighborCache API operation for FortiOS creates a new Ipv6 Neighbor Cache.
// Returns the index value of the Ipv6 Neighbor Cache and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ipv6-neighbor-cache chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemIpv6NeighborCache(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/ipv6-neighbor-cache"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemIpv6NeighborCache API operation for FortiOS updates the specified Ipv6 Neighbor Cache.
// Returns the index value of the Ipv6 Neighbor Cache and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ipv6-neighbor-cache chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemIpv6NeighborCache(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/ipv6-neighbor-cache"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemIpv6NeighborCache API operation for FortiOS deletes the specified Ipv6 Neighbor Cache.
// Returns error for service API and SDK errors.
// See the system - ipv6-neighbor-cache chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemIpv6NeighborCache(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/ipv6-neighbor-cache"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemIpv6NeighborCache API operation for FortiOS gets the Ipv6 Neighbor Cache
// with the specified index value.
// Returns the requested Ipv6 Neighbor Cache value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ipv6-neighbor-cache chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemIpv6NeighborCache(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/ipv6-neighbor-cache"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateSystemDns API operation for FortiOS updates the specified Dns.
// Returns the index value of the Dns and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - dns chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemDns(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/dns"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemDns API operation for FortiOS deletes the specified Dns.
// Returns error for service API and SDK errors.
// See the system - dns chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemDns(mkey string) (err error) {

	//No unset API for system - dns
	return
}

// ReadSystemDns API operation for FortiOS gets the Dns
// with the specified index value.
// Returns the requested Dns value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - dns chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemDns(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/dns"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemDdns API operation for FortiOS creates a new Ddns.
// Returns the index value of the Ddns and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ddns chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemDdns(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/ddns"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemDdns API operation for FortiOS updates the specified Ddns.
// Returns the index value of the Ddns and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ddns chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemDdns(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/ddns"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemDdns API operation for FortiOS deletes the specified Ddns.
// Returns error for service API and SDK errors.
// See the system - ddns chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemDdns(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/ddns"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemDdns API operation for FortiOS gets the Ddns
// with the specified index value.
// Returns the requested Ddns value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ddns chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemDdns(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/ddns"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateSystemSflow API operation for FortiOS updates the specified Sflow.
// Returns the index value of the Sflow and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sflow chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSflow(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/sflow"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemSflow API operation for FortiOS deletes the specified Sflow.
// Returns error for service API and SDK errors.
// See the system - sflow chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSflow(mkey string) (err error) {

	//No unset API for system - sflow
	return
}

// ReadSystemSflow API operation for FortiOS gets the Sflow
// with the specified index value.
// Returns the requested Sflow value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sflow chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSflow(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/sflow"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateSystemVdomSflow API operation for FortiOS updates the specified Vdom Sflow.
// Returns the index value of the Vdom Sflow and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vdom-sflow chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemVdomSflow(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/vdom-sflow"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemVdomSflow API operation for FortiOS deletes the specified Vdom Sflow.
// Returns error for service API and SDK errors.
// See the system - vdom-sflow chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemVdomSflow(mkey string) (err error) {

	//No unset API for system - vdom-sflow
	return
}

// ReadSystemVdomSflow API operation for FortiOS gets the Vdom Sflow
// with the specified index value.
// Returns the requested Vdom Sflow value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vdom-sflow chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemVdomSflow(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/vdom-sflow"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateSystemNetflow API operation for FortiOS updates the specified Netflow.
// Returns the index value of the Netflow and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - netflow chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemNetflow(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/netflow"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemNetflow API operation for FortiOS deletes the specified Netflow.
// Returns error for service API and SDK errors.
// See the system - netflow chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemNetflow(mkey string) (err error) {

	//No unset API for system - netflow
	return
}

// ReadSystemNetflow API operation for FortiOS gets the Netflow
// with the specified index value.
// Returns the requested Netflow value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - netflow chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemNetflow(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/netflow"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateSystemVdomNetflow API operation for FortiOS updates the specified Vdom Netflow.
// Returns the index value of the Vdom Netflow and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vdom-netflow chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemVdomNetflow(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/vdom-netflow"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemVdomNetflow API operation for FortiOS deletes the specified Vdom Netflow.
// Returns error for service API and SDK errors.
// See the system - vdom-netflow chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemVdomNetflow(mkey string) (err error) {

	//No unset API for system - vdom-netflow
	return
}

// ReadSystemVdomNetflow API operation for FortiOS gets the Vdom Netflow
// with the specified index value.
// Returns the requested Vdom Netflow value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vdom-netflow chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemVdomNetflow(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/vdom-netflow"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemReplacemsgImage API operation for FortiOS creates a new Replacemsg Image.
// Returns the index value of the Replacemsg Image and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - replacemsg-image chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemReplacemsgImage(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/replacemsg-image"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemReplacemsgImage API operation for FortiOS updates the specified Replacemsg Image.
// Returns the index value of the Replacemsg Image and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - replacemsg-image chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemReplacemsgImage(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/replacemsg-image"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemReplacemsgImage API operation for FortiOS deletes the specified Replacemsg Image.
// Returns error for service API and SDK errors.
// See the system - replacemsg-image chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemReplacemsgImage(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/replacemsg-image"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemReplacemsgImage API operation for FortiOS gets the Replacemsg Image
// with the specified index value.
// Returns the requested Replacemsg Image value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - replacemsg-image chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemReplacemsgImage(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/replacemsg-image"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemReplacemsgGroup API operation for FortiOS creates a new Replacemsg Group.
// Returns the index value of the Replacemsg Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - replacemsg-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemReplacemsgGroup(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/replacemsg-group"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemReplacemsgGroup API operation for FortiOS updates the specified Replacemsg Group.
// Returns the index value of the Replacemsg Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - replacemsg-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemReplacemsgGroup(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/replacemsg-group"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemReplacemsgGroup API operation for FortiOS deletes the specified Replacemsg Group.
// Returns error for service API and SDK errors.
// See the system - replacemsg-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemReplacemsgGroup(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/replacemsg-group"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemReplacemsgGroup API operation for FortiOS gets the Replacemsg Group
// with the specified index value.
// Returns the requested Replacemsg Group value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - replacemsg-group chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemReplacemsgGroup(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/replacemsg-group"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateSystemSnmpSysinfo API operation for FortiOS updates the specified Sysinfo.
// Returns the index value of the Sysinfo and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.snmp - sysinfo chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSnmpSysinfo(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.snmp/sysinfo"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemSnmpSysinfo API operation for FortiOS deletes the specified Sysinfo.
// Returns error for service API and SDK errors.
// See the system.snmp - sysinfo chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSnmpSysinfo(mkey string) (err error) {

	//No unset API for system.snmp - sysinfo
	return
}

// ReadSystemSnmpSysinfo API operation for FortiOS gets the Sysinfo
// with the specified index value.
// Returns the requested Sysinfo value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.snmp - sysinfo chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSnmpSysinfo(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.snmp/sysinfo"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemSnmpCommunity API operation for FortiOS creates a new Community.
// Returns the index value of the Community and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.snmp - community chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemSnmpCommunity(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system.snmp/community"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemSnmpCommunity API operation for FortiOS updates the specified Community.
// Returns the index value of the Community and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.snmp - community chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSnmpCommunity(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.snmp/community"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemSnmpCommunity API operation for FortiOS deletes the specified Community.
// Returns error for service API and SDK errors.
// See the system.snmp - community chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSnmpCommunity(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system.snmp/community"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemSnmpCommunity API operation for FortiOS gets the Community
// with the specified index value.
// Returns the requested Community value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.snmp - community chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSnmpCommunity(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.snmp/community"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemSnmpUser API operation for FortiOS creates a new User.
// Returns the index value of the User and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.snmp - user chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemSnmpUser(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system.snmp/user"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemSnmpUser API operation for FortiOS updates the specified User.
// Returns the index value of the User and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.snmp - user chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSnmpUser(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.snmp/user"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemSnmpUser API operation for FortiOS deletes the specified User.
// Returns error for service API and SDK errors.
// See the system.snmp - user chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSnmpUser(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system.snmp/user"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemSnmpUser API operation for FortiOS gets the User
// with the specified index value.
// Returns the requested User value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.snmp - user chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSnmpUser(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.snmp/user"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateSystemAutoupdatePushUpdate API operation for FortiOS updates the specified Push Update.
// Returns the index value of the Push Update and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.autoupdate - push-update chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAutoupdatePushUpdate(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.autoupdate/push-update"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemAutoupdatePushUpdate API operation for FortiOS deletes the specified Push Update.
// Returns error for service API and SDK errors.
// See the system.autoupdate - push-update chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAutoupdatePushUpdate(mkey string) (err error) {

	//No unset API for system.autoupdate - push-update
	return
}

// ReadSystemAutoupdatePushUpdate API operation for FortiOS gets the Push Update
// with the specified index value.
// Returns the requested Push Update value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.autoupdate - push-update chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAutoupdatePushUpdate(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.autoupdate/push-update"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateSystemAutoupdateSchedule API operation for FortiOS updates the specified Schedule.
// Returns the index value of the Schedule and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.autoupdate - schedule chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAutoupdateSchedule(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.autoupdate/schedule"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemAutoupdateSchedule API operation for FortiOS deletes the specified Schedule.
// Returns error for service API and SDK errors.
// See the system.autoupdate - schedule chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAutoupdateSchedule(mkey string) (err error) {

	//No unset API for system.autoupdate - schedule
	return
}

// ReadSystemAutoupdateSchedule API operation for FortiOS gets the Schedule
// with the specified index value.
// Returns the requested Schedule value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.autoupdate - schedule chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAutoupdateSchedule(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.autoupdate/schedule"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateSystemAutoupdateTunneling API operation for FortiOS updates the specified Tunneling.
// Returns the index value of the Tunneling and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.autoupdate - tunneling chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAutoupdateTunneling(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.autoupdate/tunneling"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemAutoupdateTunneling API operation for FortiOS deletes the specified Tunneling.
// Returns error for service API and SDK errors.
// See the system.autoupdate - tunneling chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAutoupdateTunneling(mkey string) (err error) {

	//No unset API for system.autoupdate - tunneling
	return
}

// ReadSystemAutoupdateTunneling API operation for FortiOS gets the Tunneling
// with the specified index value.
// Returns the requested Tunneling value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.autoupdate - tunneling chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAutoupdateTunneling(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.autoupdate/tunneling"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateSystemSessionTtl API operation for FortiOS updates the specified Session Ttl.
// Returns the index value of the Session Ttl and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - session-ttl chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSessionTtl(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/session-ttl"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemSessionTtl API operation for FortiOS deletes the specified Session Ttl.
// Returns error for service API and SDK errors.
// See the system - session-ttl chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSessionTtl(mkey string) (err error) {

	//No unset API for system - session-ttl
	return
}

// ReadSystemSessionTtl API operation for FortiOS gets the Session Ttl
// with the specified index value.
// Returns the requested Session Ttl value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - session-ttl chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSessionTtl(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/session-ttl"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemDhcpServer API operation for FortiOS creates a new Server.
// Returns the index value of the Server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.dhcp - server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemDhcpServer(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system.dhcp/server"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemDhcpServer API operation for FortiOS updates the specified Server.
// Returns the index value of the Server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.dhcp - server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemDhcpServer(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.dhcp/server"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemDhcpServer API operation for FortiOS deletes the specified Server.
// Returns error for service API and SDK errors.
// See the system.dhcp - server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemDhcpServer(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system.dhcp/server"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemDhcpServer API operation for FortiOS gets the Server
// with the specified index value.
// Returns the requested Server value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.dhcp - server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemDhcpServer(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.dhcp/server"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemAlias API operation for FortiOS creates a new Alias.
// Returns the index value of the Alias and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - alias chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemAlias(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/alias"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemAlias API operation for FortiOS updates the specified Alias.
// Returns the index value of the Alias and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - alias chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAlias(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/alias"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemAlias API operation for FortiOS deletes the specified Alias.
// Returns error for service API and SDK errors.
// See the system - alias chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAlias(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/alias"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemAlias API operation for FortiOS gets the Alias
// with the specified index value.
// Returns the requested Alias value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - alias chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAlias(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/alias"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemAutoScript API operation for FortiOS creates a new Auto Script.
// Returns the index value of the Auto Script and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - auto-script chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemAutoScript(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/auto-script"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemAutoScript API operation for FortiOS updates the specified Auto Script.
// Returns the index value of the Auto Script and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - auto-script chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAutoScript(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/auto-script"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemAutoScript API operation for FortiOS deletes the specified Auto Script.
// Returns error for service API and SDK errors.
// See the system - auto-script chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAutoScript(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/auto-script"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemAutoScript API operation for FortiOS gets the Auto Script
// with the specified index value.
// Returns the requested Auto Script value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - auto-script chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAutoScript(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/auto-script"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateSystemManagementTunnel API operation for FortiOS updates the specified Management Tunnel.
// Returns the index value of the Management Tunnel and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - management-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemManagementTunnel(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/management-tunnel"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemManagementTunnel API operation for FortiOS deletes the specified Management Tunnel.
// Returns error for service API and SDK errors.
// See the system - management-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemManagementTunnel(mkey string) (err error) {

	//No unset API for system - management-tunnel
	return
}

// ReadSystemManagementTunnel API operation for FortiOS gets the Management Tunnel
// with the specified index value.
// Returns the requested Management Tunnel value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - management-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemManagementTunnel(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/management-tunnel"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateSystemFortimanager API operation for FortiOS updates the specified Fortimanager.
// Returns the index value of the Fortimanager and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fortimanager chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemFortimanager(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/fortimanager"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemFortimanager API operation for FortiOS deletes the specified Fortimanager.
// Returns error for service API and SDK errors.
// See the system - fortimanager chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemFortimanager(mkey string) (err error) {

	//No unset API for system - fortimanager
	return
}

// ReadSystemFortimanager API operation for FortiOS gets the Fortimanager
// with the specified index value.
// Returns the requested Fortimanager value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fortimanager chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemFortimanager(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/fortimanager"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateSystemFm API operation for FortiOS updates the specified Fm.
// Returns the index value of the Fm and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fm chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemFm(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/fm"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemFm API operation for FortiOS deletes the specified Fm.
// Returns error for service API and SDK errors.
// See the system - fm chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemFm(mkey string) (err error) {

	//No unset API for system - fm
	return
}

// ReadSystemFm API operation for FortiOS gets the Fm
// with the specified index value.
// Returns the requested Fm value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fm chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemFm(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/fm"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateSystemCentralManagement API operation for FortiOS updates the specified Central Management.
// Returns the index value of the Central Management and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - central-management chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemCentralManagement(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/central-management"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemCentralManagement API operation for FortiOS deletes the specified Central Management.
// Returns error for service API and SDK errors.
// See the system - central-management chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemCentralManagement(mkey string) (err error) {

	//No unset API for system - central-management
	return
}

// ReadSystemCentralManagement API operation for FortiOS gets the Central Management
// with the specified index value.
// Returns the requested Central Management value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - central-management chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemCentralManagement(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/central-management"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemZone API operation for FortiOS creates a new Zone.
// Returns the index value of the Zone and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - zone chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemZone(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/zone"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemZone API operation for FortiOS updates the specified Zone.
// Returns the index value of the Zone and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - zone chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemZone(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/zone"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemZone API operation for FortiOS deletes the specified Zone.
// Returns error for service API and SDK errors.
// See the system - zone chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemZone(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/zone"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemZone API operation for FortiOS gets the Zone
// with the specified index value.
// Returns the requested Zone value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - zone chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemZone(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/zone"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemSdnConnector API operation for FortiOS creates a new Sdn Connector.
// Returns the index value of the Sdn Connector and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sdn-connector chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemSdnConnector(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/sdn-connector"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemSdnConnector API operation for FortiOS updates the specified Sdn Connector.
// Returns the index value of the Sdn Connector and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sdn-connector chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSdnConnector(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/sdn-connector"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemSdnConnector API operation for FortiOS deletes the specified Sdn Connector.
// Returns error for service API and SDK errors.
// See the system - sdn-connector chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSdnConnector(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/sdn-connector"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemSdnConnector API operation for FortiOS gets the Sdn Connector
// with the specified index value.
// Returns the requested Sdn Connector value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sdn-connector chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSdnConnector(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/sdn-connector"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemIpv6Tunnel API operation for FortiOS creates a new Ipv6 Tunnel.
// Returns the index value of the Ipv6 Tunnel and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ipv6-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemIpv6Tunnel(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/ipv6-tunnel"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemIpv6Tunnel API operation for FortiOS updates the specified Ipv6 Tunnel.
// Returns the index value of the Ipv6 Tunnel and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ipv6-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemIpv6Tunnel(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/ipv6-tunnel"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemIpv6Tunnel API operation for FortiOS deletes the specified Ipv6 Tunnel.
// Returns error for service API and SDK errors.
// See the system - ipv6-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemIpv6Tunnel(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/ipv6-tunnel"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemIpv6Tunnel API operation for FortiOS gets the Ipv6 Tunnel
// with the specified index value.
// Returns the requested Ipv6 Tunnel value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ipv6-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemIpv6Tunnel(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/ipv6-tunnel"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemExternalResource API operation for FortiOS creates a new External Resource.
// Returns the index value of the External Resource and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - external-resource chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemExternalResource(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/external-resource"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemExternalResource API operation for FortiOS updates the specified External Resource.
// Returns the index value of the External Resource and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - external-resource chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemExternalResource(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/external-resource"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemExternalResource API operation for FortiOS deletes the specified External Resource.
// Returns error for service API and SDK errors.
// See the system - external-resource chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemExternalResource(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/external-resource"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemExternalResource API operation for FortiOS gets the External Resource
// with the specified index value.
// Returns the requested External Resource value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - external-resource chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemExternalResource(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/external-resource"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateSystemNetworkVisibility API operation for FortiOS updates the specified Network Visibility.
// Returns the index value of the Network Visibility and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - network-visibility chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemNetworkVisibility(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/network-visibility"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemNetworkVisibility API operation for FortiOS deletes the specified Network Visibility.
// Returns error for service API and SDK errors.
// See the system - network-visibility chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemNetworkVisibility(mkey string) (err error) {

	//No unset API for system - network-visibility
	return
}

// ReadSystemNetworkVisibility API operation for FortiOS gets the Network Visibility
// with the specified index value.
// Returns the requested Network Visibility value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - network-visibility chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemNetworkVisibility(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/network-visibility"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemGreTunnel API operation for FortiOS creates a new Gre Tunnel.
// Returns the index value of the Gre Tunnel and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - gre-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemGreTunnel(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/gre-tunnel"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemGreTunnel API operation for FortiOS updates the specified Gre Tunnel.
// Returns the index value of the Gre Tunnel and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - gre-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemGreTunnel(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/gre-tunnel"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemGreTunnel API operation for FortiOS deletes the specified Gre Tunnel.
// Returns error for service API and SDK errors.
// See the system - gre-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemGreTunnel(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/gre-tunnel"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemGreTunnel API operation for FortiOS gets the Gre Tunnel
// with the specified index value.
// Returns the requested Gre Tunnel value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - gre-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemGreTunnel(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/gre-tunnel"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemIpsecAggregate API operation for FortiOS creates a new Ipsec Aggregate.
// Returns the index value of the Ipsec Aggregate and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ipsec-aggregate chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemIpsecAggregate(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/ipsec-aggregate"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemIpsecAggregate API operation for FortiOS updates the specified Ipsec Aggregate.
// Returns the index value of the Ipsec Aggregate and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ipsec-aggregate chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemIpsecAggregate(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/ipsec-aggregate"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemIpsecAggregate API operation for FortiOS deletes the specified Ipsec Aggregate.
// Returns error for service API and SDK errors.
// See the system - ipsec-aggregate chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemIpsecAggregate(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/ipsec-aggregate"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemIpsecAggregate API operation for FortiOS gets the Ipsec Aggregate
// with the specified index value.
// Returns the requested Ipsec Aggregate value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ipsec-aggregate chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemIpsecAggregate(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/ipsec-aggregate"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemIpipTunnel API operation for FortiOS creates a new Ipip Tunnel.
// Returns the index value of the Ipip Tunnel and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ipip-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemIpipTunnel(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/ipip-tunnel"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemIpipTunnel API operation for FortiOS updates the specified Ipip Tunnel.
// Returns the index value of the Ipip Tunnel and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ipip-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemIpipTunnel(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/ipip-tunnel"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemIpipTunnel API operation for FortiOS deletes the specified Ipip Tunnel.
// Returns error for service API and SDK errors.
// See the system - ipip-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemIpipTunnel(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/ipip-tunnel"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemIpipTunnel API operation for FortiOS gets the Ipip Tunnel
// with the specified index value.
// Returns the requested Ipip Tunnel value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ipip-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemIpipTunnel(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/ipip-tunnel"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemMobileTunnel API operation for FortiOS creates a new Mobile Tunnel.
// Returns the index value of the Mobile Tunnel and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - mobile-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemMobileTunnel(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/mobile-tunnel"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemMobileTunnel API operation for FortiOS updates the specified Mobile Tunnel.
// Returns the index value of the Mobile Tunnel and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - mobile-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemMobileTunnel(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/mobile-tunnel"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemMobileTunnel API operation for FortiOS deletes the specified Mobile Tunnel.
// Returns error for service API and SDK errors.
// See the system - mobile-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemMobileTunnel(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/mobile-tunnel"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemMobileTunnel API operation for FortiOS gets the Mobile Tunnel
// with the specified index value.
// Returns the requested Mobile Tunnel value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - mobile-tunnel chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemMobileTunnel(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/mobile-tunnel"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemPppoeInterface API operation for FortiOS creates a new Pppoe Interface.
// Returns the index value of the Pppoe Interface and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - pppoe-interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemPppoeInterface(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/pppoe-interface"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemPppoeInterface API operation for FortiOS updates the specified Pppoe Interface.
// Returns the index value of the Pppoe Interface and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - pppoe-interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemPppoeInterface(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/pppoe-interface"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemPppoeInterface API operation for FortiOS deletes the specified Pppoe Interface.
// Returns error for service API and SDK errors.
// See the system - pppoe-interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemPppoeInterface(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/pppoe-interface"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemPppoeInterface API operation for FortiOS gets the Pppoe Interface
// with the specified index value.
// Returns the requested Pppoe Interface value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - pppoe-interface chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemPppoeInterface(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/pppoe-interface"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemVxlan API operation for FortiOS creates a new Vxlan.
// Returns the index value of the Vxlan and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vxlan chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemVxlan(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/vxlan"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemVxlan API operation for FortiOS updates the specified Vxlan.
// Returns the index value of the Vxlan and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vxlan chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemVxlan(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/vxlan"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemVxlan API operation for FortiOS deletes the specified Vxlan.
// Returns error for service API and SDK errors.
// See the system - vxlan chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemVxlan(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/vxlan"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemVxlan API operation for FortiOS gets the Vxlan
// with the specified index value.
// Returns the requested Vxlan value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vxlan chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemVxlan(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/vxlan"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemDnsDatabase API operation for FortiOS creates a new Dns Database.
// Returns the index value of the Dns Database and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - dns-database chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemDnsDatabase(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/dns-database"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemDnsDatabase API operation for FortiOS updates the specified Dns Database.
// Returns the index value of the Dns Database and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - dns-database chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemDnsDatabase(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/dns-database"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemDnsDatabase API operation for FortiOS deletes the specified Dns Database.
// Returns error for service API and SDK errors.
// See the system - dns-database chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemDnsDatabase(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/dns-database"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemDnsDatabase API operation for FortiOS gets the Dns Database
// with the specified index value.
// Returns the requested Dns Database value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - dns-database chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemDnsDatabase(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/dns-database"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemDnsServer API operation for FortiOS creates a new Dns Server.
// Returns the index value of the Dns Server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - dns-server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemDnsServer(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/dns-server"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemDnsServer API operation for FortiOS updates the specified Dns Server.
// Returns the index value of the Dns Server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - dns-server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemDnsServer(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/dns-server"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemDnsServer API operation for FortiOS deletes the specified Dns Server.
// Returns error for service API and SDK errors.
// See the system - dns-server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemDnsServer(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/dns-server"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemDnsServer API operation for FortiOS gets the Dns Server
// with the specified index value.
// Returns the requested Dns Server value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - dns-server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemDnsServer(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/dns-server"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateSystemResourceLimits API operation for FortiOS updates the specified Resource Limits.
// Returns the index value of the Resource Limits and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - resource-limits chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemResourceLimits(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/resource-limits"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemResourceLimits API operation for FortiOS deletes the specified Resource Limits.
// Returns error for service API and SDK errors.
// See the system - resource-limits chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemResourceLimits(mkey string) (err error) {

	//No unset API for system - resource-limits
	return
}

// ReadSystemResourceLimits API operation for FortiOS gets the Resource Limits
// with the specified index value.
// Returns the requested Resource Limits value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - resource-limits chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemResourceLimits(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/resource-limits"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateSystemVirtualWanLink API operation for FortiOS updates the specified Virtual Wan Link.
// Returns the index value of the Virtual Wan Link and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - virtual-wan-link chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemVirtualWanLink(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/virtual-wan-link"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemVirtualWanLink API operation for FortiOS deletes the specified Virtual Wan Link.
// Returns error for service API and SDK errors.
// See the system - virtual-wan-link chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemVirtualWanLink(mkey string) (err error) {

	//No unset API for system - virtual-wan-link
	return
}

// ReadSystemVirtualWanLink API operation for FortiOS gets the Virtual Wan Link
// with the specified index value.
// Returns the requested Virtual Wan Link value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - virtual-wan-link chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemVirtualWanLink(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/virtual-wan-link"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemLldpNetworkPolicy API operation for FortiOS creates a new Network Policy.
// Returns the index value of the Network Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.lldp - network-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemLldpNetworkPolicy(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system.lldp/network-policy"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemLldpNetworkPolicy API operation for FortiOS updates the specified Network Policy.
// Returns the index value of the Network Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.lldp - network-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLldpNetworkPolicy(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.lldp/network-policy"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemLldpNetworkPolicy API operation for FortiOS deletes the specified Network Policy.
// Returns error for service API and SDK errors.
// See the system.lldp - network-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLldpNetworkPolicy(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system.lldp/network-policy"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemLldpNetworkPolicy API operation for FortiOS gets the Network Policy
// with the specified index value.
// Returns the requested Network Policy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.lldp - network-policy chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLldpNetworkPolicy(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.lldp/network-policy"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemClusterSync API operation for FortiOS creates a new Cluster Sync.
// Returns the index value of the Cluster Sync and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - cluster-sync chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemClusterSync(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/cluster-sync"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemClusterSync API operation for FortiOS updates the specified Cluster Sync.
// Returns the index value of the Cluster Sync and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - cluster-sync chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemClusterSync(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/cluster-sync"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemClusterSync API operation for FortiOS deletes the specified Cluster Sync.
// Returns error for service API and SDK errors.
// See the system - cluster-sync chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemClusterSync(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/cluster-sync"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemClusterSync API operation for FortiOS gets the Cluster Sync
// with the specified index value.
// Returns the requested Cluster Sync value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - cluster-sync chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemClusterSync(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/cluster-sync"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}


// UpdateSystemFortiguard API operation for FortiOS updates the specified Fortiguard.
// Returns the index value of the Fortiguard and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fortiguard chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemFortiguard(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/fortiguard"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemFortiguard API operation for FortiOS deletes the specified Fortiguard.
// Returns error for service API and SDK errors.
// See the system - fortiguard chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemFortiguard(mkey string) (err error) {

	//No unset API for system - fortiguard
	return
}

// ReadSystemFortiguard API operation for FortiOS gets the Fortiguard
// with the specified index value.
// Returns the requested Fortiguard value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fortiguard chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemFortiguard(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/fortiguard"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}


// UpdateSystemEmailServer API operation for FortiOS updates the specified Email Server.
// Returns the index value of the Email Server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - email-server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemEmailServer(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/email-server"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemEmailServer API operation for FortiOS deletes the specified Email Server.
// Returns error for service API and SDK errors.
// See the system - email-server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemEmailServer(mkey string) (err error) {

	//No unset API for system - email-server
	return
}

// ReadSystemEmailServer API operation for FortiOS gets the Email Server
// with the specified index value.
// Returns the requested Email Server value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - email-server chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemEmailServer(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/email-server"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

