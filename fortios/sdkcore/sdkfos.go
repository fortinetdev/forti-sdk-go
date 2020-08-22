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

