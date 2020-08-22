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

