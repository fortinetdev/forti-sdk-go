// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet),
// Liang Liu (@MaxxLiu22), Yue Wang (@yuew-ftnt)

// Description: Description: SDK for FortiAnalyzer Provider

package forticlient

import (
	"fmt"
)



// UpdateDvmCmdAddDevice API operation for FortiAnalyzer updates the specified CmdAddDevice.
// Returns the index value of the CmdAddDevice and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dvm - cmd add device chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateDvmCmdAddDevice(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/dvm/cmd/add/device"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "exec", params, false)
	return
}




// UpdateDvmCmdDelDevice API operation for FortiAnalyzer updates the specified CmdDelDevice.
// Returns the index value of the CmdDelDevice and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dvm - cmd del device chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateDvmCmdDelDevice(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/dvm/cmd/del/device"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "exec", params, false)
	return
}



// CreateDvmdbAdom API operation for FortiAnalyzer creates a new Adom.
// Returns the index value of the Adom and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dvmdb - adom chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) CreateDvmdbAdom(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/dvmdb/adom"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output, err = createUpdate(c, globaladom, path, "add", params, false)
	return
}

// UpdateDvmdbAdom API operation for FortiAnalyzer updates the specified Adom.
// Returns the index value of the Adom and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dvmdb - adom chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateDvmdbAdom(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/dvmdb/adom"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteDvmdbAdom API operation for FortiAnalyzer deletes the specified Adom.
// Returns error for service API and SDK errors.
// See the dvmdb - adom chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteDvmdbAdom(globaladom, mkey string, paralist []string) (err error) {
	path := "/dvmdb/adom"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadDvmdbAdom API operation for FortiAnalyzer gets the Adom
// with the specified index value.
// Returns the requested Adom value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dvmdb - adom chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadDvmdbAdom(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/dvmdb/adom"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateDvmdbGroup API operation for FortiAnalyzer creates a new Group.
// Returns the index value of the Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dvmdb - group chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) CreateDvmdbGroup(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/dvmdb/[*]/group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output, err = createUpdate(c, globaladom, path, "add", params, false)
	return
}

// UpdateDvmdbGroup API operation for FortiAnalyzer updates the specified Group.
// Returns the index value of the Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dvmdb - group chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateDvmdbGroup(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/dvmdb/[*]/group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteDvmdbGroup API operation for FortiAnalyzer deletes the specified Group.
// Returns error for service API and SDK errors.
// See the dvmdb - group chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteDvmdbGroup(globaladom, mkey string, paralist []string) (err error) {
	path := "/dvmdb/[*]/group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadDvmdbGroup API operation for FortiAnalyzer gets the Group
// with the specified index value.
// Returns the requested Group value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dvmdb - group chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadDvmdbGroup(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/dvmdb/[*]/group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateFmupdateAnalyzerVirusreport API operation for FortiAnalyzer updates the specified AnalyzerVirusreport.
// Returns the index value of the AnalyzerVirusreport and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - analyzer virusreport chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdateAnalyzerVirusreport(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/analyzer/virusreport"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteFmupdateAnalyzerVirusreport API operation for FortiAnalyzer deletes the specified AnalyzerVirusreport.
// Returns error for service API and SDK errors.
// See the fmupdate - analyzer virusreport chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdateAnalyzerVirusreport(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for fmupdate - analyzer virusreport
	return
}

// ReadFmupdateAnalyzerVirusreport API operation for FortiAnalyzer gets the AnalyzerVirusreport
// with the specified index value.
// Returns the requested AnalyzerVirusreport value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - analyzer virusreport chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdateAnalyzerVirusreport(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/analyzer/virusreport"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateFmupdateAvIpsAdvancedLog API operation for FortiAnalyzer updates the specified Av IpsAdvanced Log.
// Returns the index value of the Av IpsAdvanced Log and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - av-ips advanced-log chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdateAvIpsAdvancedLog(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/av-ips/advanced-log"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteFmupdateAvIpsAdvancedLog API operation for FortiAnalyzer deletes the specified Av IpsAdvanced Log.
// Returns error for service API and SDK errors.
// See the fmupdate - av-ips advanced-log chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdateAvIpsAdvancedLog(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for fmupdate - av-ips advanced-log
	return
}

// ReadFmupdateAvIpsAdvancedLog API operation for FortiAnalyzer gets the Av IpsAdvanced Log
// with the specified index value.
// Returns the requested Av IpsAdvanced Log value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - av-ips advanced-log chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdateAvIpsAdvancedLog(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/av-ips/advanced-log"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateFmupdateAvIpsWebProxy API operation for FortiAnalyzer updates the specified Av IpsWeb Proxy.
// Returns the index value of the Av IpsWeb Proxy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - av-ips web-proxy chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdateAvIpsWebProxy(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/av-ips/web-proxy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteFmupdateAvIpsWebProxy API operation for FortiAnalyzer deletes the specified Av IpsWeb Proxy.
// Returns error for service API and SDK errors.
// See the fmupdate - av-ips web-proxy chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdateAvIpsWebProxy(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for fmupdate - av-ips web-proxy
	return
}

// ReadFmupdateAvIpsWebProxy API operation for FortiAnalyzer gets the Av IpsWeb Proxy
// with the specified index value.
// Returns the requested Av IpsWeb Proxy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - av-ips web-proxy chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdateAvIpsWebProxy(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/av-ips/web-proxy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateFmupdateCustomUrlList API operation for FortiAnalyzer updates the specified Custom Url List.
// Returns the index value of the Custom Url List and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - custom-url-list chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdateCustomUrlList(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/custom-url-list"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteFmupdateCustomUrlList API operation for FortiAnalyzer deletes the specified Custom Url List.
// Returns error for service API and SDK errors.
// See the fmupdate - custom-url-list chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdateCustomUrlList(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for fmupdate - custom-url-list
	return
}

// ReadFmupdateCustomUrlList API operation for FortiAnalyzer gets the Custom Url List
// with the specified index value.
// Returns the requested Custom Url List value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - custom-url-list chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdateCustomUrlList(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/custom-url-list"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateFmupdateDiskQuota API operation for FortiAnalyzer updates the specified Disk Quota.
// Returns the index value of the Disk Quota and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - disk-quota chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdateDiskQuota(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/disk-quota"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteFmupdateDiskQuota API operation for FortiAnalyzer deletes the specified Disk Quota.
// Returns error for service API and SDK errors.
// See the fmupdate - disk-quota chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdateDiskQuota(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for fmupdate - disk-quota
	return
}

// ReadFmupdateDiskQuota API operation for FortiAnalyzer gets the Disk Quota
// with the specified index value.
// Returns the requested Disk Quota value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - disk-quota chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdateDiskQuota(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/disk-quota"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateFmupdateFctServices API operation for FortiAnalyzer updates the specified Fct Services.
// Returns the index value of the Fct Services and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - fct-services chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdateFctServices(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/fct-services"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteFmupdateFctServices API operation for FortiAnalyzer deletes the specified Fct Services.
// Returns error for service API and SDK errors.
// See the fmupdate - fct-services chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdateFctServices(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for fmupdate - fct-services
	return
}

// ReadFmupdateFctServices API operation for FortiAnalyzer gets the Fct Services
// with the specified index value.
// Returns the requested Fct Services value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - fct-services chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdateFctServices(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/fct-services"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateFmupdateFdsSetting API operation for FortiAnalyzer updates the specified Fds Setting.
// Returns the index value of the Fds Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - fds-setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdateFdsSetting(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/fds-setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteFmupdateFdsSetting API operation for FortiAnalyzer deletes the specified Fds Setting.
// Returns error for service API and SDK errors.
// See the fmupdate - fds-setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdateFdsSetting(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for fmupdate - fds-setting
	return
}

// ReadFmupdateFdsSetting API operation for FortiAnalyzer gets the Fds Setting
// with the specified index value.
// Returns the requested Fds Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - fds-setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdateFdsSetting(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/fds-setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateFmupdateFdsSettingPushOverride API operation for FortiAnalyzer updates the specified Fds SettingPush Override.
// Returns the index value of the Fds SettingPush Override and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - fds-setting push-override chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdateFdsSettingPushOverride(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/fds-setting/push-override"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteFmupdateFdsSettingPushOverride API operation for FortiAnalyzer deletes the specified Fds SettingPush Override.
// Returns error for service API and SDK errors.
// See the fmupdate - fds-setting push-override chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdateFdsSettingPushOverride(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for fmupdate - fds-setting push-override
	return
}

// ReadFmupdateFdsSettingPushOverride API operation for FortiAnalyzer gets the Fds SettingPush Override
// with the specified index value.
// Returns the requested Fds SettingPush Override value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - fds-setting push-override chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdateFdsSettingPushOverride(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/fds-setting/push-override"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateFmupdateFdsSettingPushOverrideToClient API operation for FortiAnalyzer updates the specified Fds SettingPush Override To Client.
// Returns the index value of the Fds SettingPush Override To Client and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - fds-setting push-override-to-client chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdateFdsSettingPushOverrideToClient(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/fds-setting/push-override-to-client"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteFmupdateFdsSettingPushOverrideToClient API operation for FortiAnalyzer deletes the specified Fds SettingPush Override To Client.
// Returns error for service API and SDK errors.
// See the fmupdate - fds-setting push-override-to-client chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdateFdsSettingPushOverrideToClient(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for fmupdate - fds-setting push-override-to-client
	return
}

// ReadFmupdateFdsSettingPushOverrideToClient API operation for FortiAnalyzer gets the Fds SettingPush Override To Client
// with the specified index value.
// Returns the requested Fds SettingPush Override To Client value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - fds-setting push-override-to-client chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdateFdsSettingPushOverrideToClient(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/fds-setting/push-override-to-client"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateFmupdateFdsSettingServerOverride API operation for FortiAnalyzer updates the specified Fds SettingServer Override.
// Returns the index value of the Fds SettingServer Override and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - fds-setting server-override chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdateFdsSettingServerOverride(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/fds-setting/server-override"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteFmupdateFdsSettingServerOverride API operation for FortiAnalyzer deletes the specified Fds SettingServer Override.
// Returns error for service API and SDK errors.
// See the fmupdate - fds-setting server-override chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdateFdsSettingServerOverride(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for fmupdate - fds-setting server-override
	return
}

// ReadFmupdateFdsSettingServerOverride API operation for FortiAnalyzer gets the Fds SettingServer Override
// with the specified index value.
// Returns the requested Fds SettingServer Override value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - fds-setting server-override chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdateFdsSettingServerOverride(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/fds-setting/server-override"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateFmupdateFdsSettingUpdateSchedule API operation for FortiAnalyzer updates the specified Fds SettingUpdate Schedule.
// Returns the index value of the Fds SettingUpdate Schedule and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - fds-setting update-schedule chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdateFdsSettingUpdateSchedule(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/fds-setting/update-schedule"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteFmupdateFdsSettingUpdateSchedule API operation for FortiAnalyzer deletes the specified Fds SettingUpdate Schedule.
// Returns error for service API and SDK errors.
// See the fmupdate - fds-setting update-schedule chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdateFdsSettingUpdateSchedule(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for fmupdate - fds-setting update-schedule
	return
}

// ReadFmupdateFdsSettingUpdateSchedule API operation for FortiAnalyzer gets the Fds SettingUpdate Schedule
// with the specified index value.
// Returns the requested Fds SettingUpdate Schedule value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - fds-setting update-schedule chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdateFdsSettingUpdateSchedule(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/fds-setting/update-schedule"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateFmupdateFwmSetting API operation for FortiAnalyzer updates the specified Fwm Setting.
// Returns the index value of the Fwm Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - fwm-setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdateFwmSetting(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/fwm-setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteFmupdateFwmSetting API operation for FortiAnalyzer deletes the specified Fwm Setting.
// Returns error for service API and SDK errors.
// See the fmupdate - fwm-setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdateFwmSetting(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for fmupdate - fwm-setting
	return
}

// ReadFmupdateFwmSetting API operation for FortiAnalyzer gets the Fwm Setting
// with the specified index value.
// Returns the requested Fwm Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - fwm-setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdateFwmSetting(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/fwm-setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateFmupdateFwmSettingUpgradeTimeout API operation for FortiAnalyzer updates the specified Fwm SettingUpgrade Timeout.
// Returns the index value of the Fwm SettingUpgrade Timeout and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - fwm-setting upgrade-timeout chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdateFwmSettingUpgradeTimeout(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/fwm-setting/upgrade-timeout"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteFmupdateFwmSettingUpgradeTimeout API operation for FortiAnalyzer deletes the specified Fwm SettingUpgrade Timeout.
// Returns error for service API and SDK errors.
// See the fmupdate - fwm-setting upgrade-timeout chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdateFwmSettingUpgradeTimeout(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for fmupdate - fwm-setting upgrade-timeout
	return
}

// ReadFmupdateFwmSettingUpgradeTimeout API operation for FortiAnalyzer gets the Fwm SettingUpgrade Timeout
// with the specified index value.
// Returns the requested Fwm SettingUpgrade Timeout value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - fwm-setting upgrade-timeout chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdateFwmSettingUpgradeTimeout(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/fwm-setting/upgrade-timeout"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateFmupdateMultilayer API operation for FortiAnalyzer updates the specified Multilayer.
// Returns the index value of the Multilayer and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - multilayer chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdateMultilayer(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/multilayer"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteFmupdateMultilayer API operation for FortiAnalyzer deletes the specified Multilayer.
// Returns error for service API and SDK errors.
// See the fmupdate - multilayer chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdateMultilayer(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for fmupdate - multilayer
	return
}

// ReadFmupdateMultilayer API operation for FortiAnalyzer gets the Multilayer
// with the specified index value.
// Returns the requested Multilayer value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - multilayer chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdateMultilayer(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/multilayer"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateFmupdatePublicnetwork API operation for FortiAnalyzer updates the specified Publicnetwork.
// Returns the index value of the Publicnetwork and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - publicnetwork chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdatePublicnetwork(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/publicnetwork"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteFmupdatePublicnetwork API operation for FortiAnalyzer deletes the specified Publicnetwork.
// Returns error for service API and SDK errors.
// See the fmupdate - publicnetwork chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdatePublicnetwork(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for fmupdate - publicnetwork
	return
}

// ReadFmupdatePublicnetwork API operation for FortiAnalyzer gets the Publicnetwork
// with the specified index value.
// Returns the requested Publicnetwork value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - publicnetwork chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdatePublicnetwork(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/publicnetwork"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateFmupdateServerAccessPriorities API operation for FortiAnalyzer updates the specified Server Access Priorities.
// Returns the index value of the Server Access Priorities and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - server-access-priorities chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdateServerAccessPriorities(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/server-access-priorities"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteFmupdateServerAccessPriorities API operation for FortiAnalyzer deletes the specified Server Access Priorities.
// Returns error for service API and SDK errors.
// See the fmupdate - server-access-priorities chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdateServerAccessPriorities(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for fmupdate - server-access-priorities
	return
}

// ReadFmupdateServerAccessPriorities API operation for FortiAnalyzer gets the Server Access Priorities
// with the specified index value.
// Returns the requested Server Access Priorities value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - server-access-priorities chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdateServerAccessPriorities(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/server-access-priorities"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateFmupdateServerOverrideStatus API operation for FortiAnalyzer updates the specified Server Override Status.
// Returns the index value of the Server Override Status and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - server-override-status chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdateServerOverrideStatus(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/server-override-status"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteFmupdateServerOverrideStatus API operation for FortiAnalyzer deletes the specified Server Override Status.
// Returns error for service API and SDK errors.
// See the fmupdate - server-override-status chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdateServerOverrideStatus(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for fmupdate - server-override-status
	return
}

// ReadFmupdateServerOverrideStatus API operation for FortiAnalyzer gets the Server Override Status
// with the specified index value.
// Returns the requested Server Override Status value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - server-override-status chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdateServerOverrideStatus(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/server-override-status"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateFmupdateService API operation for FortiAnalyzer updates the specified Service.
// Returns the index value of the Service and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - service chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdateService(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/service"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteFmupdateService API operation for FortiAnalyzer deletes the specified Service.
// Returns error for service API and SDK errors.
// See the fmupdate - service chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdateService(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for fmupdate - service
	return
}

// ReadFmupdateService API operation for FortiAnalyzer gets the Service
// with the specified index value.
// Returns the requested Service value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - service chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdateService(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/service"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateFmupdateWebSpamFgdSetting API operation for FortiAnalyzer updates the specified Web SpamFgd Setting.
// Returns the index value of the Web SpamFgd Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - web-spam fgd-setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdateWebSpamFgdSetting(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/web-spam/fgd-setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteFmupdateWebSpamFgdSetting API operation for FortiAnalyzer deletes the specified Web SpamFgd Setting.
// Returns error for service API and SDK errors.
// See the fmupdate - web-spam fgd-setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdateWebSpamFgdSetting(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for fmupdate - web-spam fgd-setting
	return
}

// ReadFmupdateWebSpamFgdSetting API operation for FortiAnalyzer gets the Web SpamFgd Setting
// with the specified index value.
// Returns the requested Web SpamFgd Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - web-spam fgd-setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdateWebSpamFgdSetting(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/web-spam/fgd-setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateFmupdateWebSpamWebProxy API operation for FortiAnalyzer updates the specified Web SpamWeb Proxy.
// Returns the index value of the Web SpamWeb Proxy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - web-spam web-proxy chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdateWebSpamWebProxy(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/web-spam/web-proxy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteFmupdateWebSpamWebProxy API operation for FortiAnalyzer deletes the specified Web SpamWeb Proxy.
// Returns error for service API and SDK errors.
// See the fmupdate - web-spam web-proxy chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdateWebSpamWebProxy(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for fmupdate - web-spam web-proxy
	return
}

// ReadFmupdateWebSpamWebProxy API operation for FortiAnalyzer gets the Web SpamWeb Proxy
// with the specified index value.
// Returns the requested Web SpamWeb Proxy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - web-spam web-proxy chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdateWebSpamWebProxy(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/web-spam/web-proxy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemAdminGroup API operation for FortiAnalyzer creates a new AdminGroup.
// Returns the index value of the AdminGroup and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin group chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemAdminGroup(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/admin/group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output, err = createUpdate(c, globaladom, path, "add", params, false)
	return
}

// UpdateSystemAdminGroup API operation for FortiAnalyzer updates the specified AdminGroup.
// Returns the index value of the AdminGroup and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin group chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAdminGroup(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/admin/group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemAdminGroup API operation for FortiAnalyzer deletes the specified AdminGroup.
// Returns error for service API and SDK errors.
// See the system - admin group chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAdminGroup(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/admin/group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemAdminGroup API operation for FortiAnalyzer gets the AdminGroup
// with the specified index value.
// Returns the requested AdminGroup value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin group chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAdminGroup(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/admin/group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemAdminLdap API operation for FortiAnalyzer creates a new AdminLdap.
// Returns the index value of the AdminLdap and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin ldap chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemAdminLdap(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/admin/ldap"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output, err = createUpdate(c, globaladom, path, "add", params, false)
	return
}

// UpdateSystemAdminLdap API operation for FortiAnalyzer updates the specified AdminLdap.
// Returns the index value of the AdminLdap and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin ldap chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAdminLdap(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/admin/ldap"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemAdminLdap API operation for FortiAnalyzer deletes the specified AdminLdap.
// Returns error for service API and SDK errors.
// See the system - admin ldap chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAdminLdap(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/admin/ldap"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemAdminLdap API operation for FortiAnalyzer gets the AdminLdap
// with the specified index value.
// Returns the requested AdminLdap value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin ldap chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAdminLdap(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/admin/ldap"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemAdminProfile API operation for FortiAnalyzer creates a new AdminProfile.
// Returns the index value of the AdminProfile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin profile chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemAdminProfile(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/admin/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output, err = createUpdate(c, globaladom, path, "add", params, false)
	return
}

// UpdateSystemAdminProfile API operation for FortiAnalyzer updates the specified AdminProfile.
// Returns the index value of the AdminProfile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin profile chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAdminProfile(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/admin/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemAdminProfile API operation for FortiAnalyzer deletes the specified AdminProfile.
// Returns error for service API and SDK errors.
// See the system - admin profile chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAdminProfile(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/admin/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemAdminProfile API operation for FortiAnalyzer gets the AdminProfile
// with the specified index value.
// Returns the requested AdminProfile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin profile chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAdminProfile(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/admin/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemAdminRadius API operation for FortiAnalyzer creates a new AdminRadius.
// Returns the index value of the AdminRadius and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin radius chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemAdminRadius(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/admin/radius"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output, err = createUpdate(c, globaladom, path, "add", params, false)
	return
}

// UpdateSystemAdminRadius API operation for FortiAnalyzer updates the specified AdminRadius.
// Returns the index value of the AdminRadius and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin radius chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAdminRadius(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/admin/radius"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemAdminRadius API operation for FortiAnalyzer deletes the specified AdminRadius.
// Returns error for service API and SDK errors.
// See the system - admin radius chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAdminRadius(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/admin/radius"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemAdminRadius API operation for FortiAnalyzer gets the AdminRadius
// with the specified index value.
// Returns the requested AdminRadius value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin radius chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAdminRadius(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/admin/radius"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemAdminSetting API operation for FortiAnalyzer updates the specified AdminSetting.
// Returns the index value of the AdminSetting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAdminSetting(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/admin/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemAdminSetting API operation for FortiAnalyzer deletes the specified AdminSetting.
// Returns error for service API and SDK errors.
// See the system - admin setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAdminSetting(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - admin setting
	return
}

// ReadSystemAdminSetting API operation for FortiAnalyzer gets the AdminSetting
// with the specified index value.
// Returns the requested AdminSetting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAdminSetting(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/admin/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemAdminTacacs API operation for FortiAnalyzer creates a new AdminTacacs.
// Returns the index value of the AdminTacacs and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin tacacs chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemAdminTacacs(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/admin/tacacs"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output, err = createUpdate(c, globaladom, path, "add", params, false)
	return
}

// UpdateSystemAdminTacacs API operation for FortiAnalyzer updates the specified AdminTacacs.
// Returns the index value of the AdminTacacs and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin tacacs chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAdminTacacs(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/admin/tacacs"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemAdminTacacs API operation for FortiAnalyzer deletes the specified AdminTacacs.
// Returns error for service API and SDK errors.
// See the system - admin tacacs chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAdminTacacs(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/admin/tacacs"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemAdminTacacs API operation for FortiAnalyzer gets the AdminTacacs
// with the specified index value.
// Returns the requested AdminTacacs value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin tacacs chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAdminTacacs(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/admin/tacacs"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemAdminUser API operation for FortiAnalyzer creates a new AdminUser.
// Returns the index value of the AdminUser and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin user chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemAdminUser(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/admin/user"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output, err = createUpdate(c, globaladom, path, "add", params, false)
	return
}

// UpdateSystemAdminUser API operation for FortiAnalyzer updates the specified AdminUser.
// Returns the index value of the AdminUser and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin user chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAdminUser(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/admin/user"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemAdminUser API operation for FortiAnalyzer deletes the specified AdminUser.
// Returns error for service API and SDK errors.
// See the system - admin user chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAdminUser(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/admin/user"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemAdminUser API operation for FortiAnalyzer gets the AdminUser
// with the specified index value.
// Returns the requested AdminUser value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin user chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAdminUser(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/admin/user"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemAlertConsole API operation for FortiAnalyzer updates the specified Alert Console.
// Returns the index value of the Alert Console and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - alert-console chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAlertConsole(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/alert-console"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemAlertConsole API operation for FortiAnalyzer deletes the specified Alert Console.
// Returns error for service API and SDK errors.
// See the system - alert-console chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAlertConsole(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - alert-console
	return
}

// ReadSystemAlertConsole API operation for FortiAnalyzer gets the Alert Console
// with the specified index value.
// Returns the requested Alert Console value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - alert-console chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAlertConsole(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/alert-console"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemAlertEvent API operation for FortiAnalyzer creates a new Alert Event.
// Returns the index value of the Alert Event and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - alert-event chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemAlertEvent(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/alert-event"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output, err = createUpdate(c, globaladom, path, "add", params, false)
	return
}

// UpdateSystemAlertEvent API operation for FortiAnalyzer updates the specified Alert Event.
// Returns the index value of the Alert Event and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - alert-event chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAlertEvent(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/alert-event"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemAlertEvent API operation for FortiAnalyzer deletes the specified Alert Event.
// Returns error for service API and SDK errors.
// See the system - alert-event chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAlertEvent(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/alert-event"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemAlertEvent API operation for FortiAnalyzer gets the Alert Event
// with the specified index value.
// Returns the requested Alert Event value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - alert-event chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAlertEvent(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/alert-event"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemAlertemail API operation for FortiAnalyzer updates the specified Alertemail.
// Returns the index value of the Alertemail and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - alertemail chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAlertemail(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/alertemail"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemAlertemail API operation for FortiAnalyzer deletes the specified Alertemail.
// Returns error for service API and SDK errors.
// See the system - alertemail chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAlertemail(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - alertemail
	return
}

// ReadSystemAlertemail API operation for FortiAnalyzer gets the Alertemail
// with the specified index value.
// Returns the requested Alertemail value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - alertemail chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAlertemail(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/alertemail"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemAutoDelete API operation for FortiAnalyzer updates the specified Auto Delete.
// Returns the index value of the Auto Delete and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - auto-delete chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAutoDelete(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/auto-delete"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemAutoDelete API operation for FortiAnalyzer deletes the specified Auto Delete.
// Returns error for service API and SDK errors.
// See the system - auto-delete chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAutoDelete(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - auto-delete
	return
}

// ReadSystemAutoDelete API operation for FortiAnalyzer gets the Auto Delete
// with the specified index value.
// Returns the requested Auto Delete value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - auto-delete chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAutoDelete(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/auto-delete"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemAutoDeleteDlpFilesAutoDeletion API operation for FortiAnalyzer updates the specified Auto DeleteDlp Files Auto Deletion.
// Returns the index value of the Auto DeleteDlp Files Auto Deletion and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - auto-delete dlp-files-auto-deletion chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAutoDeleteDlpFilesAutoDeletion(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/auto-delete/dlp-files-auto-deletion"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemAutoDeleteDlpFilesAutoDeletion API operation for FortiAnalyzer deletes the specified Auto DeleteDlp Files Auto Deletion.
// Returns error for service API and SDK errors.
// See the system - auto-delete dlp-files-auto-deletion chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAutoDeleteDlpFilesAutoDeletion(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - auto-delete dlp-files-auto-deletion
	return
}

// ReadSystemAutoDeleteDlpFilesAutoDeletion API operation for FortiAnalyzer gets the Auto DeleteDlp Files Auto Deletion
// with the specified index value.
// Returns the requested Auto DeleteDlp Files Auto Deletion value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - auto-delete dlp-files-auto-deletion chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAutoDeleteDlpFilesAutoDeletion(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/auto-delete/dlp-files-auto-deletion"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemAutoDeleteLogAutoDeletion API operation for FortiAnalyzer updates the specified Auto DeleteLog Auto Deletion.
// Returns the index value of the Auto DeleteLog Auto Deletion and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - auto-delete log-auto-deletion chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAutoDeleteLogAutoDeletion(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/auto-delete/log-auto-deletion"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemAutoDeleteLogAutoDeletion API operation for FortiAnalyzer deletes the specified Auto DeleteLog Auto Deletion.
// Returns error for service API and SDK errors.
// See the system - auto-delete log-auto-deletion chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAutoDeleteLogAutoDeletion(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - auto-delete log-auto-deletion
	return
}

// ReadSystemAutoDeleteLogAutoDeletion API operation for FortiAnalyzer gets the Auto DeleteLog Auto Deletion
// with the specified index value.
// Returns the requested Auto DeleteLog Auto Deletion value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - auto-delete log-auto-deletion chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAutoDeleteLogAutoDeletion(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/auto-delete/log-auto-deletion"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemAutoDeleteQuarantineFilesAutoDeletion API operation for FortiAnalyzer updates the specified Auto DeleteQuarantine Files Auto Deletion.
// Returns the index value of the Auto DeleteQuarantine Files Auto Deletion and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - auto-delete quarantine-files-auto-deletion chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAutoDeleteQuarantineFilesAutoDeletion(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/auto-delete/quarantine-files-auto-deletion"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemAutoDeleteQuarantineFilesAutoDeletion API operation for FortiAnalyzer deletes the specified Auto DeleteQuarantine Files Auto Deletion.
// Returns error for service API and SDK errors.
// See the system - auto-delete quarantine-files-auto-deletion chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAutoDeleteQuarantineFilesAutoDeletion(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - auto-delete quarantine-files-auto-deletion
	return
}

// ReadSystemAutoDeleteQuarantineFilesAutoDeletion API operation for FortiAnalyzer gets the Auto DeleteQuarantine Files Auto Deletion
// with the specified index value.
// Returns the requested Auto DeleteQuarantine Files Auto Deletion value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - auto-delete quarantine-files-auto-deletion chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAutoDeleteQuarantineFilesAutoDeletion(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/auto-delete/quarantine-files-auto-deletion"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemAutoDeleteReportAutoDeletion API operation for FortiAnalyzer updates the specified Auto DeleteReport Auto Deletion.
// Returns the index value of the Auto DeleteReport Auto Deletion and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - auto-delete report-auto-deletion chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAutoDeleteReportAutoDeletion(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/auto-delete/report-auto-deletion"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemAutoDeleteReportAutoDeletion API operation for FortiAnalyzer deletes the specified Auto DeleteReport Auto Deletion.
// Returns error for service API and SDK errors.
// See the system - auto-delete report-auto-deletion chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAutoDeleteReportAutoDeletion(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - auto-delete report-auto-deletion
	return
}

// ReadSystemAutoDeleteReportAutoDeletion API operation for FortiAnalyzer gets the Auto DeleteReport Auto Deletion
// with the specified index value.
// Returns the requested Auto DeleteReport Auto Deletion value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - auto-delete report-auto-deletion chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAutoDeleteReportAutoDeletion(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/auto-delete/report-auto-deletion"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemBackupAllSettings API operation for FortiAnalyzer updates the specified BackupAll Settings.
// Returns the index value of the BackupAll Settings and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - backup all-settings chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemBackupAllSettings(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/backup/all-settings"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemBackupAllSettings API operation for FortiAnalyzer deletes the specified BackupAll Settings.
// Returns error for service API and SDK errors.
// See the system - backup all-settings chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemBackupAllSettings(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - backup all-settings
	return
}

// ReadSystemBackupAllSettings API operation for FortiAnalyzer gets the BackupAll Settings
// with the specified index value.
// Returns the requested BackupAll Settings value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - backup all-settings chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemBackupAllSettings(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/backup/all-settings"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemCentralManagement API operation for FortiAnalyzer updates the specified Central Management.
// Returns the index value of the Central Management and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - central-management chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemCentralManagement(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/central-management"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemCentralManagement API operation for FortiAnalyzer deletes the specified Central Management.
// Returns error for service API and SDK errors.
// See the system - central-management chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemCentralManagement(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - central-management
	return
}

// ReadSystemCentralManagement API operation for FortiAnalyzer gets the Central Management
// with the specified index value.
// Returns the requested Central Management value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - central-management chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemCentralManagement(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/central-management"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemCertificateCa API operation for FortiAnalyzer creates a new CertificateCa.
// Returns the index value of the CertificateCa and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - certificate ca chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemCertificateCa(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/certificate/ca"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output, err = createUpdate(c, globaladom, path, "add", params, false)
	return
}

// UpdateSystemCertificateCa API operation for FortiAnalyzer updates the specified CertificateCa.
// Returns the index value of the CertificateCa and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - certificate ca chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemCertificateCa(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/certificate/ca"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemCertificateCa API operation for FortiAnalyzer deletes the specified CertificateCa.
// Returns error for service API and SDK errors.
// See the system - certificate ca chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemCertificateCa(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/certificate/ca"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemCertificateCa API operation for FortiAnalyzer gets the CertificateCa
// with the specified index value.
// Returns the requested CertificateCa value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - certificate ca chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemCertificateCa(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/certificate/ca"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemCertificateCrl API operation for FortiAnalyzer creates a new CertificateCrl.
// Returns the index value of the CertificateCrl and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - certificate crl chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemCertificateCrl(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/certificate/crl"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output, err = createUpdate(c, globaladom, path, "add", params, false)
	return
}

// UpdateSystemCertificateCrl API operation for FortiAnalyzer updates the specified CertificateCrl.
// Returns the index value of the CertificateCrl and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - certificate crl chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemCertificateCrl(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/certificate/crl"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemCertificateCrl API operation for FortiAnalyzer deletes the specified CertificateCrl.
// Returns error for service API and SDK errors.
// See the system - certificate crl chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemCertificateCrl(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/certificate/crl"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemCertificateCrl API operation for FortiAnalyzer gets the CertificateCrl
// with the specified index value.
// Returns the requested CertificateCrl value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - certificate crl chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemCertificateCrl(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/certificate/crl"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemCertificateLocal API operation for FortiAnalyzer creates a new CertificateLocal.
// Returns the index value of the CertificateLocal and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - certificate local chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemCertificateLocal(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/certificate/local"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output, err = createUpdate(c, globaladom, path, "add", params, false)
	return
}

// UpdateSystemCertificateLocal API operation for FortiAnalyzer updates the specified CertificateLocal.
// Returns the index value of the CertificateLocal and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - certificate local chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemCertificateLocal(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/certificate/local"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemCertificateLocal API operation for FortiAnalyzer deletes the specified CertificateLocal.
// Returns error for service API and SDK errors.
// See the system - certificate local chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemCertificateLocal(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/certificate/local"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemCertificateLocal API operation for FortiAnalyzer gets the CertificateLocal
// with the specified index value.
// Returns the requested CertificateLocal value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - certificate local chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemCertificateLocal(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/certificate/local"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemCertificateOftp API operation for FortiAnalyzer updates the specified CertificateOftp.
// Returns the index value of the CertificateOftp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - certificate oftp chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemCertificateOftp(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/certificate/oftp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemCertificateOftp API operation for FortiAnalyzer deletes the specified CertificateOftp.
// Returns error for service API and SDK errors.
// See the system - certificate oftp chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemCertificateOftp(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - certificate oftp
	return
}

// ReadSystemCertificateOftp API operation for FortiAnalyzer gets the CertificateOftp
// with the specified index value.
// Returns the requested CertificateOftp value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - certificate oftp chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemCertificateOftp(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/certificate/oftp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemCertificateRemote API operation for FortiAnalyzer creates a new CertificateRemote.
// Returns the index value of the CertificateRemote and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - certificate remote chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemCertificateRemote(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/certificate/remote"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output, err = createUpdate(c, globaladom, path, "add", params, false)
	return
}

// UpdateSystemCertificateRemote API operation for FortiAnalyzer updates the specified CertificateRemote.
// Returns the index value of the CertificateRemote and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - certificate remote chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemCertificateRemote(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/certificate/remote"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemCertificateRemote API operation for FortiAnalyzer deletes the specified CertificateRemote.
// Returns error for service API and SDK errors.
// See the system - certificate remote chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemCertificateRemote(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/certificate/remote"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemCertificateRemote API operation for FortiAnalyzer gets the CertificateRemote
// with the specified index value.
// Returns the requested CertificateRemote value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - certificate remote chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemCertificateRemote(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/certificate/remote"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemCertificateSsh API operation for FortiAnalyzer creates a new CertificateSsh.
// Returns the index value of the CertificateSsh and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - certificate ssh chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemCertificateSsh(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/certificate/ssh"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output, err = createUpdate(c, globaladom, path, "add", params, false)
	return
}

// UpdateSystemCertificateSsh API operation for FortiAnalyzer updates the specified CertificateSsh.
// Returns the index value of the CertificateSsh and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - certificate ssh chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemCertificateSsh(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/certificate/ssh"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemCertificateSsh API operation for FortiAnalyzer deletes the specified CertificateSsh.
// Returns error for service API and SDK errors.
// See the system - certificate ssh chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemCertificateSsh(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/certificate/ssh"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemCertificateSsh API operation for FortiAnalyzer gets the CertificateSsh
// with the specified index value.
// Returns the requested CertificateSsh value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - certificate ssh chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemCertificateSsh(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/certificate/ssh"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemConnector API operation for FortiAnalyzer updates the specified Connector.
// Returns the index value of the Connector and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - connector chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemConnector(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/connector"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemConnector API operation for FortiAnalyzer deletes the specified Connector.
// Returns error for service API and SDK errors.
// See the system - connector chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemConnector(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - connector
	return
}

// ReadSystemConnector API operation for FortiAnalyzer gets the Connector
// with the specified index value.
// Returns the requested Connector value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - connector chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemConnector(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/connector"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemDns API operation for FortiAnalyzer updates the specified Dns.
// Returns the index value of the Dns and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - dns chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemDns(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/dns"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemDns API operation for FortiAnalyzer deletes the specified Dns.
// Returns error for service API and SDK errors.
// See the system - dns chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemDns(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - dns
	return
}

// ReadSystemDns API operation for FortiAnalyzer gets the Dns
// with the specified index value.
// Returns the requested Dns value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - dns chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemDns(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/dns"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemDocker API operation for FortiAnalyzer updates the specified Docker.
// Returns the index value of the Docker and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - docker chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemDocker(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/docker"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemDocker API operation for FortiAnalyzer deletes the specified Docker.
// Returns error for service API and SDK errors.
// See the system - docker chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemDocker(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - docker
	return
}

// ReadSystemDocker API operation for FortiAnalyzer gets the Docker
// with the specified index value.
// Returns the requested Docker value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - docker chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemDocker(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/docker"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemFips API operation for FortiAnalyzer updates the specified Fips.
// Returns the index value of the Fips and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fips chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemFips(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/fips"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemFips API operation for FortiAnalyzer deletes the specified Fips.
// Returns error for service API and SDK errors.
// See the system - fips chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemFips(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - fips
	return
}

// ReadSystemFips API operation for FortiAnalyzer gets the Fips
// with the specified index value.
// Returns the requested Fips value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fips chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemFips(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/fips"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemFortiviewAutoCache API operation for FortiAnalyzer updates the specified FortiviewAuto Cache.
// Returns the index value of the FortiviewAuto Cache and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fortiview auto-cache chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemFortiviewAutoCache(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/fortiview/auto-cache"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemFortiviewAutoCache API operation for FortiAnalyzer deletes the specified FortiviewAuto Cache.
// Returns error for service API and SDK errors.
// See the system - fortiview auto-cache chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemFortiviewAutoCache(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - fortiview auto-cache
	return
}

// ReadSystemFortiviewAutoCache API operation for FortiAnalyzer gets the FortiviewAuto Cache
// with the specified index value.
// Returns the requested FortiviewAuto Cache value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fortiview auto-cache chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemFortiviewAutoCache(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/fortiview/auto-cache"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemFortiviewSetting API operation for FortiAnalyzer updates the specified FortiviewSetting.
// Returns the index value of the FortiviewSetting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fortiview setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemFortiviewSetting(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/fortiview/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemFortiviewSetting API operation for FortiAnalyzer deletes the specified FortiviewSetting.
// Returns error for service API and SDK errors.
// See the system - fortiview setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemFortiviewSetting(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - fortiview setting
	return
}

// ReadSystemFortiviewSetting API operation for FortiAnalyzer gets the FortiviewSetting
// with the specified index value.
// Returns the requested FortiviewSetting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fortiview setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemFortiviewSetting(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/fortiview/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemGlobal API operation for FortiAnalyzer updates the specified Global.
// Returns the index value of the Global and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - global chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemGlobal(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/global"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemGlobal API operation for FortiAnalyzer deletes the specified Global.
// Returns error for service API and SDK errors.
// See the system - global chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemGlobal(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - global
	return
}

// ReadSystemGlobal API operation for FortiAnalyzer gets the Global
// with the specified index value.
// Returns the requested Global value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - global chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemGlobal(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/global"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemGlobalSslCipherSuites API operation for FortiAnalyzer creates a new GlobalSsl Cipher Suites.
// Returns the index value of the GlobalSsl Cipher Suites and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - global ssl-cipher-suites chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemGlobalSslCipherSuites(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/global/ssl-cipher-suites"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output, err = createUpdate(c, globaladom, path, "add", params, false)
	return
}

// UpdateSystemGlobalSslCipherSuites API operation for FortiAnalyzer updates the specified GlobalSsl Cipher Suites.
// Returns the index value of the GlobalSsl Cipher Suites and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - global ssl-cipher-suites chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemGlobalSslCipherSuites(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/global/ssl-cipher-suites"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemGlobalSslCipherSuites API operation for FortiAnalyzer deletes the specified GlobalSsl Cipher Suites.
// Returns error for service API and SDK errors.
// See the system - global ssl-cipher-suites chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemGlobalSslCipherSuites(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/global/ssl-cipher-suites"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemGlobalSslCipherSuites API operation for FortiAnalyzer gets the GlobalSsl Cipher Suites
// with the specified index value.
// Returns the requested GlobalSsl Cipher Suites value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - global ssl-cipher-suites chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemGlobalSslCipherSuites(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/global/ssl-cipher-suites"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemGuiact API operation for FortiAnalyzer updates the specified Guiact.
// Returns the index value of the Guiact and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - guiact chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemGuiact(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/guiact"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemGuiact API operation for FortiAnalyzer deletes the specified Guiact.
// Returns error for service API and SDK errors.
// See the system - guiact chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemGuiact(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - guiact
	return
}

// ReadSystemGuiact API operation for FortiAnalyzer gets the Guiact
// with the specified index value.
// Returns the requested Guiact value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - guiact chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemGuiact(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/guiact"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemHa API operation for FortiAnalyzer updates the specified Ha.
// Returns the index value of the Ha and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ha chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemHa(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/ha"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemHa API operation for FortiAnalyzer deletes the specified Ha.
// Returns error for service API and SDK errors.
// See the system - ha chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemHa(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - ha
	return
}

// ReadSystemHa API operation for FortiAnalyzer gets the Ha
// with the specified index value.
// Returns the requested Ha value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ha chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemHa(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/ha"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemHaPeer API operation for FortiAnalyzer creates a new HaPeer.
// Returns the index value of the HaPeer and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ha peer chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemHaPeer(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/ha/peer"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output, err = createUpdate(c, globaladom, path, "add", params, false)
	return
}

// UpdateSystemHaPeer API operation for FortiAnalyzer updates the specified HaPeer.
// Returns the index value of the HaPeer and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ha peer chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemHaPeer(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/ha/peer"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemHaPeer API operation for FortiAnalyzer deletes the specified HaPeer.
// Returns error for service API and SDK errors.
// See the system - ha peer chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemHaPeer(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/ha/peer"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemHaPeer API operation for FortiAnalyzer gets the HaPeer
// with the specified index value.
// Returns the requested HaPeer value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ha peer chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemHaPeer(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/ha/peer"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemHaPrivatePeer API operation for FortiAnalyzer creates a new HaPrivate Peer.
// Returns the index value of the HaPrivate Peer and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ha private-peer chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemHaPrivatePeer(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/ha/private-peer"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output, err = createUpdate(c, globaladom, path, "add", params, false)
	return
}

// UpdateSystemHaPrivatePeer API operation for FortiAnalyzer updates the specified HaPrivate Peer.
// Returns the index value of the HaPrivate Peer and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ha private-peer chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemHaPrivatePeer(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/ha/private-peer"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemHaPrivatePeer API operation for FortiAnalyzer deletes the specified HaPrivate Peer.
// Returns error for service API and SDK errors.
// See the system - ha private-peer chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemHaPrivatePeer(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/ha/private-peer"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemHaPrivatePeer API operation for FortiAnalyzer gets the HaPrivate Peer
// with the specified index value.
// Returns the requested HaPrivate Peer value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ha private-peer chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemHaPrivatePeer(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/ha/private-peer"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemHaVip API operation for FortiAnalyzer creates a new HaVip.
// Returns the index value of the HaVip and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ha vip chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemHaVip(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/ha/vip"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output, err = createUpdate(c, globaladom, path, "add", params, false)
	return
}

// UpdateSystemHaVip API operation for FortiAnalyzer updates the specified HaVip.
// Returns the index value of the HaVip and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ha vip chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemHaVip(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/ha/vip"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemHaVip API operation for FortiAnalyzer deletes the specified HaVip.
// Returns error for service API and SDK errors.
// See the system - ha vip chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemHaVip(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/ha/vip"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemHaVip API operation for FortiAnalyzer gets the HaVip
// with the specified index value.
// Returns the requested HaVip value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ha vip chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemHaVip(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/ha/vip"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemInterface API operation for FortiAnalyzer creates a new Interface.
// Returns the index value of the Interface and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - interface chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemInterface(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/interface"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output, err = createUpdate(c, globaladom, path, "add", params, false)
	return
}

// UpdateSystemInterface API operation for FortiAnalyzer updates the specified Interface.
// Returns the index value of the Interface and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - interface chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemInterface(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/interface"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemInterface API operation for FortiAnalyzer deletes the specified Interface.
// Returns error for service API and SDK errors.
// See the system - interface chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemInterface(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/interface"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemInterface API operation for FortiAnalyzer gets the Interface
// with the specified index value.
// Returns the requested Interface value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - interface chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemInterface(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/interface"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemLocalInPolicy API operation for FortiAnalyzer creates a new Local In Policy.
// Returns the index value of the Local In Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - local-in-policy chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemLocalInPolicy(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/local-in-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output, err = createUpdate(c, globaladom, path, "add", params, false)
	return
}

// UpdateSystemLocalInPolicy API operation for FortiAnalyzer updates the specified Local In Policy.
// Returns the index value of the Local In Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - local-in-policy chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLocalInPolicy(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/local-in-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemLocalInPolicy API operation for FortiAnalyzer deletes the specified Local In Policy.
// Returns error for service API and SDK errors.
// See the system - local-in-policy chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLocalInPolicy(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/local-in-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemLocalInPolicy API operation for FortiAnalyzer gets the Local In Policy
// with the specified index value.
// Returns the requested Local In Policy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - local-in-policy chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLocalInPolicy(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/local-in-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemLocalInPolicy6 API operation for FortiAnalyzer creates a new Local In Policy6.
// Returns the index value of the Local In Policy6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - local-in-policy6 chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemLocalInPolicy6(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/local-in-policy6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output, err = createUpdate(c, globaladom, path, "add", params, false)
	return
}

// UpdateSystemLocalInPolicy6 API operation for FortiAnalyzer updates the specified Local In Policy6.
// Returns the index value of the Local In Policy6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - local-in-policy6 chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLocalInPolicy6(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/local-in-policy6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemLocalInPolicy6 API operation for FortiAnalyzer deletes the specified Local In Policy6.
// Returns error for service API and SDK errors.
// See the system - local-in-policy6 chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLocalInPolicy6(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/local-in-policy6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemLocalInPolicy6 API operation for FortiAnalyzer gets the Local In Policy6
// with the specified index value.
// Returns the requested Local In Policy6 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - local-in-policy6 chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLocalInPolicy6(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/local-in-policy6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLocallogDiskFilter API operation for FortiAnalyzer updates the specified LocallogDiskFilter.
// Returns the index value of the LocallogDiskFilter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog disk filter chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLocallogDiskFilter(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/disk/filter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemLocallogDiskFilter API operation for FortiAnalyzer deletes the specified LocallogDiskFilter.
// Returns error for service API and SDK errors.
// See the system - locallog disk filter chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLocallogDiskFilter(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - locallog disk filter
	return
}

// ReadSystemLocallogDiskFilter API operation for FortiAnalyzer gets the LocallogDiskFilter
// with the specified index value.
// Returns the requested LocallogDiskFilter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog disk filter chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLocallogDiskFilter(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/disk/filter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLocallogDiskSetting API operation for FortiAnalyzer updates the specified LocallogDiskSetting.
// Returns the index value of the LocallogDiskSetting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog disk setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLocallogDiskSetting(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/disk/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemLocallogDiskSetting API operation for FortiAnalyzer deletes the specified LocallogDiskSetting.
// Returns error for service API and SDK errors.
// See the system - locallog disk setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLocallogDiskSetting(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - locallog disk setting
	return
}

// ReadSystemLocallogDiskSetting API operation for FortiAnalyzer gets the LocallogDiskSetting
// with the specified index value.
// Returns the requested LocallogDiskSetting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog disk setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLocallogDiskSetting(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/disk/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLocallogFortianalyzer2Filter API operation for FortiAnalyzer updates the specified LocallogFortianalyzer2Filter.
// Returns the index value of the LocallogFortianalyzer2Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog fortianalyzer2 filter chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLocallogFortianalyzer2Filter(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/fortianalyzer2/filter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemLocallogFortianalyzer2Filter API operation for FortiAnalyzer deletes the specified LocallogFortianalyzer2Filter.
// Returns error for service API and SDK errors.
// See the system - locallog fortianalyzer2 filter chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLocallogFortianalyzer2Filter(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - locallog fortianalyzer2 filter
	return
}

// ReadSystemLocallogFortianalyzer2Filter API operation for FortiAnalyzer gets the LocallogFortianalyzer2Filter
// with the specified index value.
// Returns the requested LocallogFortianalyzer2Filter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog fortianalyzer2 filter chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLocallogFortianalyzer2Filter(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/fortianalyzer2/filter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLocallogFortianalyzer2Setting API operation for FortiAnalyzer updates the specified LocallogFortianalyzer2Setting.
// Returns the index value of the LocallogFortianalyzer2Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog fortianalyzer2 setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLocallogFortianalyzer2Setting(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/fortianalyzer2/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemLocallogFortianalyzer2Setting API operation for FortiAnalyzer deletes the specified LocallogFortianalyzer2Setting.
// Returns error for service API and SDK errors.
// See the system - locallog fortianalyzer2 setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLocallogFortianalyzer2Setting(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - locallog fortianalyzer2 setting
	return
}

// ReadSystemLocallogFortianalyzer2Setting API operation for FortiAnalyzer gets the LocallogFortianalyzer2Setting
// with the specified index value.
// Returns the requested LocallogFortianalyzer2Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog fortianalyzer2 setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLocallogFortianalyzer2Setting(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/fortianalyzer2/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLocallogFortianalyzer3Filter API operation for FortiAnalyzer updates the specified LocallogFortianalyzer3Filter.
// Returns the index value of the LocallogFortianalyzer3Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog fortianalyzer3 filter chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLocallogFortianalyzer3Filter(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/fortianalyzer3/filter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemLocallogFortianalyzer3Filter API operation for FortiAnalyzer deletes the specified LocallogFortianalyzer3Filter.
// Returns error for service API and SDK errors.
// See the system - locallog fortianalyzer3 filter chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLocallogFortianalyzer3Filter(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - locallog fortianalyzer3 filter
	return
}

// ReadSystemLocallogFortianalyzer3Filter API operation for FortiAnalyzer gets the LocallogFortianalyzer3Filter
// with the specified index value.
// Returns the requested LocallogFortianalyzer3Filter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog fortianalyzer3 filter chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLocallogFortianalyzer3Filter(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/fortianalyzer3/filter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLocallogFortianalyzer3Setting API operation for FortiAnalyzer updates the specified LocallogFortianalyzer3Setting.
// Returns the index value of the LocallogFortianalyzer3Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog fortianalyzer3 setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLocallogFortianalyzer3Setting(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/fortianalyzer3/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemLocallogFortianalyzer3Setting API operation for FortiAnalyzer deletes the specified LocallogFortianalyzer3Setting.
// Returns error for service API and SDK errors.
// See the system - locallog fortianalyzer3 setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLocallogFortianalyzer3Setting(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - locallog fortianalyzer3 setting
	return
}

// ReadSystemLocallogFortianalyzer3Setting API operation for FortiAnalyzer gets the LocallogFortianalyzer3Setting
// with the specified index value.
// Returns the requested LocallogFortianalyzer3Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog fortianalyzer3 setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLocallogFortianalyzer3Setting(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/fortianalyzer3/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLocallogFortianalyzerFilter API operation for FortiAnalyzer updates the specified LocallogFortianalyzerFilter.
// Returns the index value of the LocallogFortianalyzerFilter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog fortianalyzer filter chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLocallogFortianalyzerFilter(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/fortianalyzer/filter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemLocallogFortianalyzerFilter API operation for FortiAnalyzer deletes the specified LocallogFortianalyzerFilter.
// Returns error for service API and SDK errors.
// See the system - locallog fortianalyzer filter chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLocallogFortianalyzerFilter(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - locallog fortianalyzer filter
	return
}

// ReadSystemLocallogFortianalyzerFilter API operation for FortiAnalyzer gets the LocallogFortianalyzerFilter
// with the specified index value.
// Returns the requested LocallogFortianalyzerFilter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog fortianalyzer filter chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLocallogFortianalyzerFilter(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/fortianalyzer/filter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLocallogFortianalyzerSetting API operation for FortiAnalyzer updates the specified LocallogFortianalyzerSetting.
// Returns the index value of the LocallogFortianalyzerSetting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog fortianalyzer setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLocallogFortianalyzerSetting(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/fortianalyzer/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemLocallogFortianalyzerSetting API operation for FortiAnalyzer deletes the specified LocallogFortianalyzerSetting.
// Returns error for service API and SDK errors.
// See the system - locallog fortianalyzer setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLocallogFortianalyzerSetting(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - locallog fortianalyzer setting
	return
}

// ReadSystemLocallogFortianalyzerSetting API operation for FortiAnalyzer gets the LocallogFortianalyzerSetting
// with the specified index value.
// Returns the requested LocallogFortianalyzerSetting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog fortianalyzer setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLocallogFortianalyzerSetting(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/fortianalyzer/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLocallogMemoryFilter API operation for FortiAnalyzer updates the specified LocallogMemoryFilter.
// Returns the index value of the LocallogMemoryFilter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog memory filter chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLocallogMemoryFilter(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/memory/filter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemLocallogMemoryFilter API operation for FortiAnalyzer deletes the specified LocallogMemoryFilter.
// Returns error for service API and SDK errors.
// See the system - locallog memory filter chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLocallogMemoryFilter(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - locallog memory filter
	return
}

// ReadSystemLocallogMemoryFilter API operation for FortiAnalyzer gets the LocallogMemoryFilter
// with the specified index value.
// Returns the requested LocallogMemoryFilter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog memory filter chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLocallogMemoryFilter(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/memory/filter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLocallogMemorySetting API operation for FortiAnalyzer updates the specified LocallogMemorySetting.
// Returns the index value of the LocallogMemorySetting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog memory setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLocallogMemorySetting(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/memory/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemLocallogMemorySetting API operation for FortiAnalyzer deletes the specified LocallogMemorySetting.
// Returns error for service API and SDK errors.
// See the system - locallog memory setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLocallogMemorySetting(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - locallog memory setting
	return
}

// ReadSystemLocallogMemorySetting API operation for FortiAnalyzer gets the LocallogMemorySetting
// with the specified index value.
// Returns the requested LocallogMemorySetting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog memory setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLocallogMemorySetting(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/memory/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLocallogSetting API operation for FortiAnalyzer updates the specified LocallogSetting.
// Returns the index value of the LocallogSetting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLocallogSetting(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemLocallogSetting API operation for FortiAnalyzer deletes the specified LocallogSetting.
// Returns error for service API and SDK errors.
// See the system - locallog setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLocallogSetting(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - locallog setting
	return
}

// ReadSystemLocallogSetting API operation for FortiAnalyzer gets the LocallogSetting
// with the specified index value.
// Returns the requested LocallogSetting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLocallogSetting(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLocallogSyslogd2Filter API operation for FortiAnalyzer updates the specified LocallogSyslogd2Filter.
// Returns the index value of the LocallogSyslogd2Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog syslogd2 filter chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLocallogSyslogd2Filter(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/syslogd2/filter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemLocallogSyslogd2Filter API operation for FortiAnalyzer deletes the specified LocallogSyslogd2Filter.
// Returns error for service API and SDK errors.
// See the system - locallog syslogd2 filter chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLocallogSyslogd2Filter(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - locallog syslogd2 filter
	return
}

// ReadSystemLocallogSyslogd2Filter API operation for FortiAnalyzer gets the LocallogSyslogd2Filter
// with the specified index value.
// Returns the requested LocallogSyslogd2Filter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog syslogd2 filter chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLocallogSyslogd2Filter(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/syslogd2/filter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLocallogSyslogd2Setting API operation for FortiAnalyzer updates the specified LocallogSyslogd2Setting.
// Returns the index value of the LocallogSyslogd2Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog syslogd2 setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLocallogSyslogd2Setting(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/syslogd2/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemLocallogSyslogd2Setting API operation for FortiAnalyzer deletes the specified LocallogSyslogd2Setting.
// Returns error for service API and SDK errors.
// See the system - locallog syslogd2 setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLocallogSyslogd2Setting(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - locallog syslogd2 setting
	return
}

// ReadSystemLocallogSyslogd2Setting API operation for FortiAnalyzer gets the LocallogSyslogd2Setting
// with the specified index value.
// Returns the requested LocallogSyslogd2Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog syslogd2 setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLocallogSyslogd2Setting(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/syslogd2/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLocallogSyslogd3Filter API operation for FortiAnalyzer updates the specified LocallogSyslogd3Filter.
// Returns the index value of the LocallogSyslogd3Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog syslogd3 filter chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLocallogSyslogd3Filter(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/syslogd3/filter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemLocallogSyslogd3Filter API operation for FortiAnalyzer deletes the specified LocallogSyslogd3Filter.
// Returns error for service API and SDK errors.
// See the system - locallog syslogd3 filter chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLocallogSyslogd3Filter(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - locallog syslogd3 filter
	return
}

// ReadSystemLocallogSyslogd3Filter API operation for FortiAnalyzer gets the LocallogSyslogd3Filter
// with the specified index value.
// Returns the requested LocallogSyslogd3Filter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog syslogd3 filter chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLocallogSyslogd3Filter(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/syslogd3/filter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLocallogSyslogd3Setting API operation for FortiAnalyzer updates the specified LocallogSyslogd3Setting.
// Returns the index value of the LocallogSyslogd3Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog syslogd3 setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLocallogSyslogd3Setting(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/syslogd3/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemLocallogSyslogd3Setting API operation for FortiAnalyzer deletes the specified LocallogSyslogd3Setting.
// Returns error for service API and SDK errors.
// See the system - locallog syslogd3 setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLocallogSyslogd3Setting(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - locallog syslogd3 setting
	return
}

// ReadSystemLocallogSyslogd3Setting API operation for FortiAnalyzer gets the LocallogSyslogd3Setting
// with the specified index value.
// Returns the requested LocallogSyslogd3Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog syslogd3 setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLocallogSyslogd3Setting(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/syslogd3/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLocallogSyslogdFilter API operation for FortiAnalyzer updates the specified LocallogSyslogdFilter.
// Returns the index value of the LocallogSyslogdFilter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog syslogd filter chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLocallogSyslogdFilter(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/syslogd/filter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemLocallogSyslogdFilter API operation for FortiAnalyzer deletes the specified LocallogSyslogdFilter.
// Returns error for service API and SDK errors.
// See the system - locallog syslogd filter chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLocallogSyslogdFilter(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - locallog syslogd filter
	return
}

// ReadSystemLocallogSyslogdFilter API operation for FortiAnalyzer gets the LocallogSyslogdFilter
// with the specified index value.
// Returns the requested LocallogSyslogdFilter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog syslogd filter chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLocallogSyslogdFilter(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/syslogd/filter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLocallogSyslogdSetting API operation for FortiAnalyzer updates the specified LocallogSyslogdSetting.
// Returns the index value of the LocallogSyslogdSetting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog syslogd setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLocallogSyslogdSetting(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/syslogd/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemLocallogSyslogdSetting API operation for FortiAnalyzer deletes the specified LocallogSyslogdSetting.
// Returns error for service API and SDK errors.
// See the system - locallog syslogd setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLocallogSyslogdSetting(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - locallog syslogd setting
	return
}

// ReadSystemLocallogSyslogdSetting API operation for FortiAnalyzer gets the LocallogSyslogdSetting
// with the specified index value.
// Returns the requested LocallogSyslogdSetting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog syslogd setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLocallogSyslogdSetting(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/syslogd/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemLogFetchClientProfile API operation for FortiAnalyzer creates a new Log FetchClient Profile.
// Returns the index value of the Log FetchClient Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log-fetch client-profile chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemLogFetchClientProfile(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/log-fetch/client-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output, err = createUpdate(c, globaladom, path, "add", params, false)
	return
}

// UpdateSystemLogFetchClientProfile API operation for FortiAnalyzer updates the specified Log FetchClient Profile.
// Returns the index value of the Log FetchClient Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log-fetch client-profile chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLogFetchClientProfile(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/log-fetch/client-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemLogFetchClientProfile API operation for FortiAnalyzer deletes the specified Log FetchClient Profile.
// Returns error for service API and SDK errors.
// See the system - log-fetch client-profile chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLogFetchClientProfile(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/log-fetch/client-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemLogFetchClientProfile API operation for FortiAnalyzer gets the Log FetchClient Profile
// with the specified index value.
// Returns the requested Log FetchClient Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log-fetch client-profile chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLogFetchClientProfile(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/log-fetch/client-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLogFetchServerSettings API operation for FortiAnalyzer updates the specified Log FetchServer Settings.
// Returns the index value of the Log FetchServer Settings and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log-fetch server-settings chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLogFetchServerSettings(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/log-fetch/server-settings"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemLogFetchServerSettings API operation for FortiAnalyzer deletes the specified Log FetchServer Settings.
// Returns error for service API and SDK errors.
// See the system - log-fetch server-settings chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLogFetchServerSettings(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - log-fetch server-settings
	return
}

// ReadSystemLogFetchServerSettings API operation for FortiAnalyzer gets the Log FetchServer Settings
// with the specified index value.
// Returns the requested Log FetchServer Settings value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log-fetch server-settings chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLogFetchServerSettings(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/log-fetch/server-settings"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemLogForward API operation for FortiAnalyzer creates a new Log Forward.
// Returns the index value of the Log Forward and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log-forward chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemLogForward(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/log-forward"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output, err = createUpdate(c, globaladom, path, "add", params, false)
	return
}

// UpdateSystemLogForward API operation for FortiAnalyzer updates the specified Log Forward.
// Returns the index value of the Log Forward and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log-forward chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLogForward(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/log-forward"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemLogForward API operation for FortiAnalyzer deletes the specified Log Forward.
// Returns error for service API and SDK errors.
// See the system - log-forward chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLogForward(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/log-forward"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemLogForward API operation for FortiAnalyzer gets the Log Forward
// with the specified index value.
// Returns the requested Log Forward value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log-forward chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLogForward(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/log-forward"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLogForwardService API operation for FortiAnalyzer updates the specified Log Forward Service.
// Returns the index value of the Log Forward Service and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log-forward-service chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLogForwardService(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/log-forward-service"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemLogForwardService API operation for FortiAnalyzer deletes the specified Log Forward Service.
// Returns error for service API and SDK errors.
// See the system - log-forward-service chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLogForwardService(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - log-forward-service
	return
}

// ReadSystemLogForwardService API operation for FortiAnalyzer gets the Log Forward Service
// with the specified index value.
// Returns the requested Log Forward Service value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log-forward-service chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLogForwardService(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/log-forward-service"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLogAlert API operation for FortiAnalyzer updates the specified LogAlert.
// Returns the index value of the LogAlert and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log alert chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLogAlert(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/alert"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemLogAlert API operation for FortiAnalyzer deletes the specified LogAlert.
// Returns error for service API and SDK errors.
// See the system - log alert chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLogAlert(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - log alert
	return
}

// ReadSystemLogAlert API operation for FortiAnalyzer gets the LogAlert
// with the specified index value.
// Returns the requested LogAlert value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log alert chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLogAlert(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/alert"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemLogDeviceDisable API operation for FortiAnalyzer creates a new LogDevice Disable.
// Returns the index value of the LogDevice Disable and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log device-disable chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemLogDeviceDisable(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/device-disable"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output, err = createUpdate(c, globaladom, path, "add", params, false)
	return
}

// UpdateSystemLogDeviceDisable API operation for FortiAnalyzer updates the specified LogDevice Disable.
// Returns the index value of the LogDevice Disable and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log device-disable chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLogDeviceDisable(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/device-disable"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemLogDeviceDisable API operation for FortiAnalyzer deletes the specified LogDevice Disable.
// Returns error for service API and SDK errors.
// See the system - log device-disable chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLogDeviceDisable(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/log/device-disable"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemLogDeviceDisable API operation for FortiAnalyzer gets the LogDevice Disable
// with the specified index value.
// Returns the requested LogDevice Disable value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log device-disable chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLogDeviceDisable(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/device-disable"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLogFosPolicyStats API operation for FortiAnalyzer updates the specified LogFos Policy Stats.
// Returns the index value of the LogFos Policy Stats and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log fos-policy-stats chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLogFosPolicyStats(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/fos-policy-stats"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemLogFosPolicyStats API operation for FortiAnalyzer deletes the specified LogFos Policy Stats.
// Returns error for service API and SDK errors.
// See the system - log fos-policy-stats chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLogFosPolicyStats(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - log fos-policy-stats
	return
}

// ReadSystemLogFosPolicyStats API operation for FortiAnalyzer gets the LogFos Policy Stats
// with the specified index value.
// Returns the requested LogFos Policy Stats value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log fos-policy-stats chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLogFosPolicyStats(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/fos-policy-stats"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLogInterfaceStats API operation for FortiAnalyzer updates the specified LogInterface Stats.
// Returns the index value of the LogInterface Stats and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log interface-stats chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLogInterfaceStats(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/interface-stats"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemLogInterfaceStats API operation for FortiAnalyzer deletes the specified LogInterface Stats.
// Returns error for service API and SDK errors.
// See the system - log interface-stats chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLogInterfaceStats(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - log interface-stats
	return
}

// ReadSystemLogInterfaceStats API operation for FortiAnalyzer gets the LogInterface Stats
// with the specified index value.
// Returns the requested LogInterface Stats value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log interface-stats chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLogInterfaceStats(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/interface-stats"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLogIoc API operation for FortiAnalyzer updates the specified LogIoc.
// Returns the index value of the LogIoc and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log ioc chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLogIoc(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/ioc"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemLogIoc API operation for FortiAnalyzer deletes the specified LogIoc.
// Returns error for service API and SDK errors.
// See the system - log ioc chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLogIoc(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - log ioc
	return
}

// ReadSystemLogIoc API operation for FortiAnalyzer gets the LogIoc
// with the specified index value.
// Returns the requested LogIoc value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log ioc chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLogIoc(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/ioc"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemLogMailDomain API operation for FortiAnalyzer creates a new LogMail Domain.
// Returns the index value of the LogMail Domain and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log mail-domain chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemLogMailDomain(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/mail-domain"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output, err = createUpdate(c, globaladom, path, "add", params, false)
	return
}

// UpdateSystemLogMailDomain API operation for FortiAnalyzer updates the specified LogMail Domain.
// Returns the index value of the LogMail Domain and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log mail-domain chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLogMailDomain(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/mail-domain"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemLogMailDomain API operation for FortiAnalyzer deletes the specified LogMail Domain.
// Returns error for service API and SDK errors.
// See the system - log mail-domain chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLogMailDomain(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/log/mail-domain"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemLogMailDomain API operation for FortiAnalyzer gets the LogMail Domain
// with the specified index value.
// Returns the requested LogMail Domain value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log mail-domain chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLogMailDomain(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/mail-domain"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLogRatelimit API operation for FortiAnalyzer updates the specified LogRatelimit.
// Returns the index value of the LogRatelimit and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log ratelimit chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLogRatelimit(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/ratelimit"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemLogRatelimit API operation for FortiAnalyzer deletes the specified LogRatelimit.
// Returns error for service API and SDK errors.
// See the system - log ratelimit chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLogRatelimit(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - log ratelimit
	return
}

// ReadSystemLogRatelimit API operation for FortiAnalyzer gets the LogRatelimit
// with the specified index value.
// Returns the requested LogRatelimit value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log ratelimit chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLogRatelimit(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/ratelimit"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemLogRatelimitDevice API operation for FortiAnalyzer creates a new LogRatelimitDevice.
// Returns the index value of the LogRatelimitDevice and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log ratelimit device chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemLogRatelimitDevice(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/ratelimit/device"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output, err = createUpdate(c, globaladom, path, "add", params, false)
	return
}

// UpdateSystemLogRatelimitDevice API operation for FortiAnalyzer updates the specified LogRatelimitDevice.
// Returns the index value of the LogRatelimitDevice and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log ratelimit device chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLogRatelimitDevice(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/ratelimit/device"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemLogRatelimitDevice API operation for FortiAnalyzer deletes the specified LogRatelimitDevice.
// Returns error for service API and SDK errors.
// See the system - log ratelimit device chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLogRatelimitDevice(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/log/ratelimit/device"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemLogRatelimitDevice API operation for FortiAnalyzer gets the LogRatelimitDevice
// with the specified index value.
// Returns the requested LogRatelimitDevice value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log ratelimit device chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLogRatelimitDevice(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/ratelimit/device"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemLogRatelimitRatelimits API operation for FortiAnalyzer creates a new LogRatelimitRatelimits.
// Returns the index value of the LogRatelimitRatelimits and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log ratelimit ratelimits chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemLogRatelimitRatelimits(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/ratelimit/ratelimits"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output, err = createUpdate(c, globaladom, path, "add", params, false)
	return
}

// UpdateSystemLogRatelimitRatelimits API operation for FortiAnalyzer updates the specified LogRatelimitRatelimits.
// Returns the index value of the LogRatelimitRatelimits and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log ratelimit ratelimits chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLogRatelimitRatelimits(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/ratelimit/ratelimits"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemLogRatelimitRatelimits API operation for FortiAnalyzer deletes the specified LogRatelimitRatelimits.
// Returns error for service API and SDK errors.
// See the system - log ratelimit ratelimits chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLogRatelimitRatelimits(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/log/ratelimit/ratelimits"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemLogRatelimitRatelimits API operation for FortiAnalyzer gets the LogRatelimitRatelimits
// with the specified index value.
// Returns the requested LogRatelimitRatelimits value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log ratelimit ratelimits chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLogRatelimitRatelimits(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/ratelimit/ratelimits"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLogSettings API operation for FortiAnalyzer updates the specified LogSettings.
// Returns the index value of the LogSettings and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log settings chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLogSettings(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/settings"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemLogSettings API operation for FortiAnalyzer deletes the specified LogSettings.
// Returns error for service API and SDK errors.
// See the system - log settings chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLogSettings(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - log settings
	return
}

// ReadSystemLogSettings API operation for FortiAnalyzer gets the LogSettings
// with the specified index value.
// Returns the requested LogSettings value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log settings chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLogSettings(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/settings"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLogSettingsRollingAnalyzer API operation for FortiAnalyzer updates the specified LogSettingsRolling Analyzer.
// Returns the index value of the LogSettingsRolling Analyzer and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log settings rolling-analyzer chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLogSettingsRollingAnalyzer(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/settings/rolling-analyzer"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemLogSettingsRollingAnalyzer API operation for FortiAnalyzer deletes the specified LogSettingsRolling Analyzer.
// Returns error for service API and SDK errors.
// See the system - log settings rolling-analyzer chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLogSettingsRollingAnalyzer(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - log settings rolling-analyzer
	return
}

// ReadSystemLogSettingsRollingAnalyzer API operation for FortiAnalyzer gets the LogSettingsRolling Analyzer
// with the specified index value.
// Returns the requested LogSettingsRolling Analyzer value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log settings rolling-analyzer chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLogSettingsRollingAnalyzer(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/settings/rolling-analyzer"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLogSettingsRollingLocal API operation for FortiAnalyzer updates the specified LogSettingsRolling Local.
// Returns the index value of the LogSettingsRolling Local and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log settings rolling-local chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLogSettingsRollingLocal(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/settings/rolling-local"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemLogSettingsRollingLocal API operation for FortiAnalyzer deletes the specified LogSettingsRolling Local.
// Returns error for service API and SDK errors.
// See the system - log settings rolling-local chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLogSettingsRollingLocal(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - log settings rolling-local
	return
}

// ReadSystemLogSettingsRollingLocal API operation for FortiAnalyzer gets the LogSettingsRolling Local
// with the specified index value.
// Returns the requested LogSettingsRolling Local value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log settings rolling-local chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLogSettingsRollingLocal(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/settings/rolling-local"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLogSettingsRollingRegular API operation for FortiAnalyzer updates the specified LogSettingsRolling Regular.
// Returns the index value of the LogSettingsRolling Regular and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log settings rolling-regular chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLogSettingsRollingRegular(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/settings/rolling-regular"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemLogSettingsRollingRegular API operation for FortiAnalyzer deletes the specified LogSettingsRolling Regular.
// Returns error for service API and SDK errors.
// See the system - log settings rolling-regular chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLogSettingsRollingRegular(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - log settings rolling-regular
	return
}

// ReadSystemLogSettingsRollingRegular API operation for FortiAnalyzer gets the LogSettingsRolling Regular
// with the specified index value.
// Returns the requested LogSettingsRolling Regular value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log settings rolling-regular chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLogSettingsRollingRegular(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/settings/rolling-regular"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLogTopology API operation for FortiAnalyzer updates the specified LogTopology.
// Returns the index value of the LogTopology and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log topology chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLogTopology(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/topology"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemLogTopology API operation for FortiAnalyzer deletes the specified LogTopology.
// Returns error for service API and SDK errors.
// See the system - log topology chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLogTopology(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - log topology
	return
}

// ReadSystemLogTopology API operation for FortiAnalyzer gets the LogTopology
// with the specified index value.
// Returns the requested LogTopology value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log topology chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLogTopology(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/topology"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemMail API operation for FortiAnalyzer creates a new Mail.
// Returns the index value of the Mail and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - mail chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemMail(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/mail"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output, err = createUpdate(c, globaladom, path, "add", params, false)
	return
}

// UpdateSystemMail API operation for FortiAnalyzer updates the specified Mail.
// Returns the index value of the Mail and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - mail chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemMail(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/mail"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemMail API operation for FortiAnalyzer deletes the specified Mail.
// Returns error for service API and SDK errors.
// See the system - mail chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemMail(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/mail"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemMail API operation for FortiAnalyzer gets the Mail
// with the specified index value.
// Returns the requested Mail value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - mail chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemMail(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/mail"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemMetadataAdmins API operation for FortiAnalyzer creates a new MetadataAdmins.
// Returns the index value of the MetadataAdmins and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - metadata admins chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemMetadataAdmins(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/metadata/admins"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output, err = createUpdate(c, globaladom, path, "add", params, false)
	return
}

// UpdateSystemMetadataAdmins API operation for FortiAnalyzer updates the specified MetadataAdmins.
// Returns the index value of the MetadataAdmins and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - metadata admins chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemMetadataAdmins(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/metadata/admins"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemMetadataAdmins API operation for FortiAnalyzer deletes the specified MetadataAdmins.
// Returns error for service API and SDK errors.
// See the system - metadata admins chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemMetadataAdmins(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/metadata/admins"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemMetadataAdmins API operation for FortiAnalyzer gets the MetadataAdmins
// with the specified index value.
// Returns the requested MetadataAdmins value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - metadata admins chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemMetadataAdmins(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/metadata/admins"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemNtp API operation for FortiAnalyzer updates the specified Ntp.
// Returns the index value of the Ntp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ntp chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemNtp(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/ntp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemNtp API operation for FortiAnalyzer deletes the specified Ntp.
// Returns error for service API and SDK errors.
// See the system - ntp chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemNtp(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - ntp
	return
}

// ReadSystemNtp API operation for FortiAnalyzer gets the Ntp
// with the specified index value.
// Returns the requested Ntp value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ntp chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemNtp(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/ntp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemNtpNtpserver API operation for FortiAnalyzer creates a new NtpNtpserver.
// Returns the index value of the NtpNtpserver and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ntp ntpserver chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemNtpNtpserver(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/ntp/ntpserver"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output, err = createUpdate(c, globaladom, path, "add", params, false)
	return
}

// UpdateSystemNtpNtpserver API operation for FortiAnalyzer updates the specified NtpNtpserver.
// Returns the index value of the NtpNtpserver and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ntp ntpserver chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemNtpNtpserver(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/ntp/ntpserver"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemNtpNtpserver API operation for FortiAnalyzer deletes the specified NtpNtpserver.
// Returns error for service API and SDK errors.
// See the system - ntp ntpserver chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemNtpNtpserver(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/ntp/ntpserver"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemNtpNtpserver API operation for FortiAnalyzer gets the NtpNtpserver
// with the specified index value.
// Returns the requested NtpNtpserver value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ntp ntpserver chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemNtpNtpserver(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/ntp/ntpserver"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemPasswordPolicy API operation for FortiAnalyzer updates the specified Password Policy.
// Returns the index value of the Password Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - password-policy chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemPasswordPolicy(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/password-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemPasswordPolicy API operation for FortiAnalyzer deletes the specified Password Policy.
// Returns error for service API and SDK errors.
// See the system - password-policy chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemPasswordPolicy(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - password-policy
	return
}

// ReadSystemPasswordPolicy API operation for FortiAnalyzer gets the Password Policy
// with the specified index value.
// Returns the requested Password Policy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - password-policy chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemPasswordPolicy(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/password-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemReportAutoCache API operation for FortiAnalyzer updates the specified ReportAuto Cache.
// Returns the index value of the ReportAuto Cache and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - report auto-cache chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemReportAutoCache(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/report/auto-cache"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemReportAutoCache API operation for FortiAnalyzer deletes the specified ReportAuto Cache.
// Returns error for service API and SDK errors.
// See the system - report auto-cache chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemReportAutoCache(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - report auto-cache
	return
}

// ReadSystemReportAutoCache API operation for FortiAnalyzer gets the ReportAuto Cache
// with the specified index value.
// Returns the requested ReportAuto Cache value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - report auto-cache chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemReportAutoCache(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/report/auto-cache"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemReportEstBrowseTime API operation for FortiAnalyzer updates the specified ReportEst Browse Time.
// Returns the index value of the ReportEst Browse Time and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - report est-browse-time chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemReportEstBrowseTime(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/report/est-browse-time"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemReportEstBrowseTime API operation for FortiAnalyzer deletes the specified ReportEst Browse Time.
// Returns error for service API and SDK errors.
// See the system - report est-browse-time chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemReportEstBrowseTime(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - report est-browse-time
	return
}

// ReadSystemReportEstBrowseTime API operation for FortiAnalyzer gets the ReportEst Browse Time
// with the specified index value.
// Returns the requested ReportEst Browse Time value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - report est-browse-time chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemReportEstBrowseTime(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/report/est-browse-time"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemReportSetting API operation for FortiAnalyzer updates the specified ReportSetting.
// Returns the index value of the ReportSetting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - report setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemReportSetting(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/report/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemReportSetting API operation for FortiAnalyzer deletes the specified ReportSetting.
// Returns error for service API and SDK errors.
// See the system - report setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemReportSetting(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - report setting
	return
}

// ReadSystemReportSetting API operation for FortiAnalyzer gets the ReportSetting
// with the specified index value.
// Returns the requested ReportSetting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - report setting chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemReportSetting(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/report/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemRoute API operation for FortiAnalyzer creates a new Route.
// Returns the index value of the Route and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - route chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemRoute(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/route"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output, err = createUpdate(c, globaladom, path, "add", params, false)
	return
}

// UpdateSystemRoute API operation for FortiAnalyzer updates the specified Route.
// Returns the index value of the Route and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - route chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemRoute(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/route"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemRoute API operation for FortiAnalyzer deletes the specified Route.
// Returns error for service API and SDK errors.
// See the system - route chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemRoute(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/route"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemRoute API operation for FortiAnalyzer gets the Route
// with the specified index value.
// Returns the requested Route value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - route chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemRoute(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/route"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemRoute6 API operation for FortiAnalyzer creates a new Route6.
// Returns the index value of the Route6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - route6 chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemRoute6(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/route6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output, err = createUpdate(c, globaladom, path, "add", params, false)
	return
}

// UpdateSystemRoute6 API operation for FortiAnalyzer updates the specified Route6.
// Returns the index value of the Route6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - route6 chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemRoute6(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/route6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemRoute6 API operation for FortiAnalyzer deletes the specified Route6.
// Returns error for service API and SDK errors.
// See the system - route6 chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemRoute6(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/route6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemRoute6 API operation for FortiAnalyzer gets the Route6
// with the specified index value.
// Returns the requested Route6 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - route6 chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemRoute6(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/route6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemSaml API operation for FortiAnalyzer updates the specified Saml.
// Returns the index value of the Saml and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - saml chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSaml(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/saml"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemSaml API operation for FortiAnalyzer deletes the specified Saml.
// Returns error for service API and SDK errors.
// See the system - saml chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSaml(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - saml
	return
}

// ReadSystemSaml API operation for FortiAnalyzer gets the Saml
// with the specified index value.
// Returns the requested Saml value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - saml chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSaml(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/saml"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemSamlFabricIdp API operation for FortiAnalyzer creates a new SamlFabric Idp.
// Returns the index value of the SamlFabric Idp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - saml fabric-idp chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemSamlFabricIdp(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/saml/fabric-idp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output, err = createUpdate(c, globaladom, path, "add", params, false)
	return
}

// UpdateSystemSamlFabricIdp API operation for FortiAnalyzer updates the specified SamlFabric Idp.
// Returns the index value of the SamlFabric Idp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - saml fabric-idp chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSamlFabricIdp(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/saml/fabric-idp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemSamlFabricIdp API operation for FortiAnalyzer deletes the specified SamlFabric Idp.
// Returns error for service API and SDK errors.
// See the system - saml fabric-idp chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSamlFabricIdp(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/saml/fabric-idp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemSamlFabricIdp API operation for FortiAnalyzer gets the SamlFabric Idp
// with the specified index value.
// Returns the requested SamlFabric Idp value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - saml fabric-idp chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSamlFabricIdp(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/saml/fabric-idp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemSamlServiceProviders API operation for FortiAnalyzer creates a new SamlService Providers.
// Returns the index value of the SamlService Providers and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - saml service-providers chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemSamlServiceProviders(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/saml/service-providers"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output, err = createUpdate(c, globaladom, path, "add", params, false)
	return
}

// UpdateSystemSamlServiceProviders API operation for FortiAnalyzer updates the specified SamlService Providers.
// Returns the index value of the SamlService Providers and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - saml service-providers chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSamlServiceProviders(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/saml/service-providers"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemSamlServiceProviders API operation for FortiAnalyzer deletes the specified SamlService Providers.
// Returns error for service API and SDK errors.
// See the system - saml service-providers chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSamlServiceProviders(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/saml/service-providers"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemSamlServiceProviders API operation for FortiAnalyzer gets the SamlService Providers
// with the specified index value.
// Returns the requested SamlService Providers value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - saml service-providers chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSamlServiceProviders(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/saml/service-providers"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemSniffer API operation for FortiAnalyzer creates a new Sniffer.
// Returns the index value of the Sniffer and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sniffer chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemSniffer(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/sniffer"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output, err = createUpdate(c, globaladom, path, "add", params, false)
	return
}

// UpdateSystemSniffer API operation for FortiAnalyzer updates the specified Sniffer.
// Returns the index value of the Sniffer and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sniffer chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSniffer(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/sniffer"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemSniffer API operation for FortiAnalyzer deletes the specified Sniffer.
// Returns error for service API and SDK errors.
// See the system - sniffer chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSniffer(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/sniffer"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemSniffer API operation for FortiAnalyzer gets the Sniffer
// with the specified index value.
// Returns the requested Sniffer value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sniffer chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSniffer(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/sniffer"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemSnmpCommunity API operation for FortiAnalyzer creates a new SnmpCommunity.
// Returns the index value of the SnmpCommunity and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - snmp community chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemSnmpCommunity(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/snmp/community"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output, err = createUpdate(c, globaladom, path, "add", params, false)
	return
}

// UpdateSystemSnmpCommunity API operation for FortiAnalyzer updates the specified SnmpCommunity.
// Returns the index value of the SnmpCommunity and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - snmp community chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSnmpCommunity(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/snmp/community"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemSnmpCommunity API operation for FortiAnalyzer deletes the specified SnmpCommunity.
// Returns error for service API and SDK errors.
// See the system - snmp community chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSnmpCommunity(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/snmp/community"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemSnmpCommunity API operation for FortiAnalyzer gets the SnmpCommunity
// with the specified index value.
// Returns the requested SnmpCommunity value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - snmp community chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSnmpCommunity(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/snmp/community"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemSnmpSysinfo API operation for FortiAnalyzer updates the specified SnmpSysinfo.
// Returns the index value of the SnmpSysinfo and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - snmp sysinfo chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSnmpSysinfo(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/snmp/sysinfo"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemSnmpSysinfo API operation for FortiAnalyzer deletes the specified SnmpSysinfo.
// Returns error for service API and SDK errors.
// See the system - snmp sysinfo chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSnmpSysinfo(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - snmp sysinfo
	return
}

// ReadSystemSnmpSysinfo API operation for FortiAnalyzer gets the SnmpSysinfo
// with the specified index value.
// Returns the requested SnmpSysinfo value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - snmp sysinfo chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSnmpSysinfo(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/snmp/sysinfo"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemSnmpUser API operation for FortiAnalyzer creates a new SnmpUser.
// Returns the index value of the SnmpUser and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - snmp user chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemSnmpUser(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/snmp/user"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output, err = createUpdate(c, globaladom, path, "add", params, false)
	return
}

// UpdateSystemSnmpUser API operation for FortiAnalyzer updates the specified SnmpUser.
// Returns the index value of the SnmpUser and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - snmp user chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSnmpUser(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/snmp/user"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemSnmpUser API operation for FortiAnalyzer deletes the specified SnmpUser.
// Returns error for service API and SDK errors.
// See the system - snmp user chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSnmpUser(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/snmp/user"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemSnmpUser API operation for FortiAnalyzer gets the SnmpUser
// with the specified index value.
// Returns the requested SnmpUser value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - snmp user chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSnmpUser(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/snmp/user"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemSocFabric API operation for FortiAnalyzer updates the specified Soc Fabric.
// Returns the index value of the Soc Fabric and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - soc-fabric chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSocFabric(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/soc-fabric"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemSocFabric API operation for FortiAnalyzer deletes the specified Soc Fabric.
// Returns error for service API and SDK errors.
// See the system - soc-fabric chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSocFabric(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - soc-fabric
	return
}

// ReadSystemSocFabric API operation for FortiAnalyzer gets the Soc Fabric
// with the specified index value.
// Returns the requested Soc Fabric value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - soc-fabric chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSocFabric(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/soc-fabric"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemSql API operation for FortiAnalyzer updates the specified Sql.
// Returns the index value of the Sql and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sql chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSql(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/sql"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemSql API operation for FortiAnalyzer deletes the specified Sql.
// Returns error for service API and SDK errors.
// See the system - sql chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSql(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - sql
	return
}

// ReadSystemSql API operation for FortiAnalyzer gets the Sql
// with the specified index value.
// Returns the requested Sql value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sql chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSql(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/sql"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemSqlCustomIndex API operation for FortiAnalyzer creates a new SqlCustom Index.
// Returns the index value of the SqlCustom Index and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sql custom-index chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemSqlCustomIndex(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/sql/custom-index"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output, err = createUpdate(c, globaladom, path, "add", params, false)
	return
}

// UpdateSystemSqlCustomIndex API operation for FortiAnalyzer updates the specified SqlCustom Index.
// Returns the index value of the SqlCustom Index and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sql custom-index chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSqlCustomIndex(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/sql/custom-index"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemSqlCustomIndex API operation for FortiAnalyzer deletes the specified SqlCustom Index.
// Returns error for service API and SDK errors.
// See the system - sql custom-index chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSqlCustomIndex(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/sql/custom-index"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemSqlCustomIndex API operation for FortiAnalyzer gets the SqlCustom Index
// with the specified index value.
// Returns the requested SqlCustom Index value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sql custom-index chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSqlCustomIndex(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/sql/custom-index"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemSqlCustomSkipidx API operation for FortiAnalyzer creates a new SqlCustom Skipidx.
// Returns the index value of the SqlCustom Skipidx and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sql custom-skipidx chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemSqlCustomSkipidx(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/sql/custom-skipidx"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output, err = createUpdate(c, globaladom, path, "add", params, false)
	return
}

// UpdateSystemSqlCustomSkipidx API operation for FortiAnalyzer updates the specified SqlCustom Skipidx.
// Returns the index value of the SqlCustom Skipidx and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sql custom-skipidx chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSqlCustomSkipidx(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/sql/custom-skipidx"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemSqlCustomSkipidx API operation for FortiAnalyzer deletes the specified SqlCustom Skipidx.
// Returns error for service API and SDK errors.
// See the system - sql custom-skipidx chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSqlCustomSkipidx(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/sql/custom-skipidx"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemSqlCustomSkipidx API operation for FortiAnalyzer gets the SqlCustom Skipidx
// with the specified index value.
// Returns the requested SqlCustom Skipidx value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sql custom-skipidx chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSqlCustomSkipidx(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/sql/custom-skipidx"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemSqlTsIndexField API operation for FortiAnalyzer creates a new SqlTs Index Field.
// Returns the index value of the SqlTs Index Field and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sql ts-index-field chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemSqlTsIndexField(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/sql/ts-index-field"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output, err = createUpdate(c, globaladom, path, "add", params, false)
	return
}

// UpdateSystemSqlTsIndexField API operation for FortiAnalyzer updates the specified SqlTs Index Field.
// Returns the index value of the SqlTs Index Field and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sql ts-index-field chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSqlTsIndexField(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/sql/ts-index-field"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemSqlTsIndexField API operation for FortiAnalyzer deletes the specified SqlTs Index Field.
// Returns error for service API and SDK errors.
// See the system - sql ts-index-field chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSqlTsIndexField(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/sql/ts-index-field"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemSqlTsIndexField API operation for FortiAnalyzer gets the SqlTs Index Field
// with the specified index value.
// Returns the requested SqlTs Index Field value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sql ts-index-field chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSqlTsIndexField(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/sql/ts-index-field"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemSyslog API operation for FortiAnalyzer creates a new Syslog.
// Returns the index value of the Syslog and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - syslog chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemSyslog(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/syslog"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output, err = createUpdate(c, globaladom, path, "add", params, false)
	return
}

// UpdateSystemSyslog API operation for FortiAnalyzer updates the specified Syslog.
// Returns the index value of the Syslog and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - syslog chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSyslog(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/syslog"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemSyslog API operation for FortiAnalyzer deletes the specified Syslog.
// Returns error for service API and SDK errors.
// See the system - syslog chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSyslog(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/syslog"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemSyslog API operation for FortiAnalyzer gets the Syslog
// with the specified index value.
// Returns the requested Syslog value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - syslog chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSyslog(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/syslog"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemWebProxy API operation for FortiAnalyzer updates the specified Web Proxy.
// Returns the index value of the Web Proxy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - web-proxy chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemWebProxy(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/web-proxy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemWebProxy API operation for FortiAnalyzer deletes the specified Web Proxy.
// Returns error for service API and SDK errors.
// See the system - web-proxy chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemWebProxy(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - web-proxy
	return
}

// ReadSystemWebProxy API operation for FortiAnalyzer gets the Web Proxy
// with the specified index value.
// Returns the requested Web Proxy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - web-proxy chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemWebProxy(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/web-proxy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemWorkflowApprovalMatrix API operation for FortiAnalyzer creates a new WorkflowApproval Matrix.
// Returns the index value of the WorkflowApproval Matrix and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - workflow approval-matrix chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemWorkflowApprovalMatrix(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/workflow/approval-matrix"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output, err = createUpdate(c, globaladom, path, "add", params, false)
	return
}

// UpdateSystemWorkflowApprovalMatrix API operation for FortiAnalyzer updates the specified WorkflowApproval Matrix.
// Returns the index value of the WorkflowApproval Matrix and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - workflow approval-matrix chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemWorkflowApprovalMatrix(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/workflow/approval-matrix"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	output, err = createUpdate(c, globaladom, path, "set", params, false)
	return
}

// DeleteSystemWorkflowApprovalMatrix API operation for FortiAnalyzer deletes the specified WorkflowApproval Matrix.
// Returns error for service API and SDK errors.
// See the system - workflow approval-matrix chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemWorkflowApprovalMatrix(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/workflow/approval-matrix"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemWorkflowApprovalMatrix API operation for FortiAnalyzer gets the WorkflowApproval Matrix
// with the specified index value.
// Returns the requested WorkflowApproval Matrix value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - workflow approval-matrix chapter in the FortiAnalyzer Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemWorkflowApprovalMatrix(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/workflow/approval-matrix"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}
