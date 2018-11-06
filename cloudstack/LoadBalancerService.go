//
// Copyright 2018, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package cloudstack

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type AddF5LoadBalancerParams struct {
	p map[string]interface{}
}

func (p *AddF5LoadBalancerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["networkdevicetype"]; found {
		u.Set("networkdevicetype", v.(string))
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (p *AddF5LoadBalancerParams) SetNetworkdevicetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkdevicetype"] = v
	return
}

func (p *AddF5LoadBalancerParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
	return
}

func (p *AddF5LoadBalancerParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
	return
}

func (p *AddF5LoadBalancerParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
	return
}

func (p *AddF5LoadBalancerParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
	return
}

// You should always use this function to get a new AddF5LoadBalancerParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewAddF5LoadBalancerParams(networkdevicetype string, password string, physicalnetworkid string, url string, username string) *AddF5LoadBalancerParams {
	p := &AddF5LoadBalancerParams{}
	p.p = make(map[string]interface{})
	p.p["networkdevicetype"] = networkdevicetype
	p.p["password"] = password
	p.p["physicalnetworkid"] = physicalnetworkid
	p.p["url"] = url
	p.p["username"] = username
	return p
}

// Adds a F5 BigIP load balancer device
func (s *LoadBalancerService) AddF5LoadBalancer(p *AddF5LoadBalancerParams) (*AddF5LoadBalancerResponse, error) {
	resp, err := s.cs.newRequest("addF5LoadBalancer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddF5LoadBalancerResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.Jobid, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type AddF5LoadBalancerResponse struct {
	Ipaddress         string `json:"ipaddress"`
	Jobid             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Lbdevicecapacity  int64  `json:"lbdevicecapacity"`
	Lbdevicededicated bool   `json:"lbdevicededicated"`
	Lbdeviceid        string `json:"lbdeviceid"`
	Lbdevicename      string `json:"lbdevicename"`
	Lbdevicestate     string `json:"lbdevicestate"`
	Physicalnetworkid string `json:"physicalnetworkid"`
	Privateinterface  string `json:"privateinterface"`
	Provider          string `json:"provider"`
	Publicinterface   string `json:"publicinterface"`
}

type AddNetscalerLoadBalancerParams struct {
	p map[string]interface{}
}

func (p *AddNetscalerLoadBalancerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["gslbprovider"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("gslbprovider", vv)
	}
	if v, found := p.p["gslbproviderprivateip"]; found {
		u.Set("gslbproviderprivateip", v.(string))
	}
	if v, found := p.p["gslbproviderpublicip"]; found {
		u.Set("gslbproviderpublicip", v.(string))
	}
	if v, found := p.p["isexclusivegslbprovider"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isexclusivegslbprovider", vv)
	}
	if v, found := p.p["networkdevicetype"]; found {
		u.Set("networkdevicetype", v.(string))
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (p *AddNetscalerLoadBalancerParams) SetGslbprovider(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gslbprovider"] = v
	return
}

func (p *AddNetscalerLoadBalancerParams) SetGslbproviderprivateip(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gslbproviderprivateip"] = v
	return
}

func (p *AddNetscalerLoadBalancerParams) SetGslbproviderpublicip(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gslbproviderpublicip"] = v
	return
}

func (p *AddNetscalerLoadBalancerParams) SetIsexclusivegslbprovider(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isexclusivegslbprovider"] = v
	return
}

func (p *AddNetscalerLoadBalancerParams) SetNetworkdevicetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkdevicetype"] = v
	return
}

func (p *AddNetscalerLoadBalancerParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
	return
}

func (p *AddNetscalerLoadBalancerParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
	return
}

func (p *AddNetscalerLoadBalancerParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
	return
}

func (p *AddNetscalerLoadBalancerParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
	return
}

// You should always use this function to get a new AddNetscalerLoadBalancerParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewAddNetscalerLoadBalancerParams(networkdevicetype string, password string, physicalnetworkid string, url string, username string) *AddNetscalerLoadBalancerParams {
	p := &AddNetscalerLoadBalancerParams{}
	p.p = make(map[string]interface{})
	p.p["networkdevicetype"] = networkdevicetype
	p.p["password"] = password
	p.p["physicalnetworkid"] = physicalnetworkid
	p.p["url"] = url
	p.p["username"] = username
	return p
}

// Adds a netscaler load balancer device
func (s *LoadBalancerService) AddNetscalerLoadBalancer(p *AddNetscalerLoadBalancerParams) (*AddNetscalerLoadBalancerResponse, error) {
	resp, err := s.cs.newRequest("addNetscalerLoadBalancer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddNetscalerLoadBalancerResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.Jobid, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type AddNetscalerLoadBalancerResponse struct {
	Gslbprovider            bool     `json:"gslbprovider"`
	Gslbproviderprivateip   string   `json:"gslbproviderprivateip"`
	Gslbproviderpublicip    string   `json:"gslbproviderpublicip"`
	Ipaddress               string   `json:"ipaddress"`
	Isexclusivegslbprovider bool     `json:"isexclusivegslbprovider"`
	Jobid                   string   `json:"jobid"`
	Jobstatus               int      `json:"jobstatus"`
	Lbdevicecapacity        int64    `json:"lbdevicecapacity"`
	Lbdevicededicated       bool     `json:"lbdevicededicated"`
	Lbdeviceid              string   `json:"lbdeviceid"`
	Lbdevicename            string   `json:"lbdevicename"`
	Lbdevicestate           string   `json:"lbdevicestate"`
	Physicalnetworkid       string   `json:"physicalnetworkid"`
	Podids                  []string `json:"podids"`
	Privateinterface        string   `json:"privateinterface"`
	Provider                string   `json:"provider"`
	Publicinterface         string   `json:"publicinterface"`
}

type AssignCertToLoadBalancerParams struct {
	p map[string]interface{}
}

func (p *AssignCertToLoadBalancerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["certid"]; found {
		u.Set("certid", v.(string))
	}
	if v, found := p.p["lbruleid"]; found {
		u.Set("lbruleid", v.(string))
	}
	return u
}

func (p *AssignCertToLoadBalancerParams) SetCertid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["certid"] = v
	return
}

func (p *AssignCertToLoadBalancerParams) SetLbruleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbruleid"] = v
	return
}

// You should always use this function to get a new AssignCertToLoadBalancerParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewAssignCertToLoadBalancerParams(certid string, lbruleid string) *AssignCertToLoadBalancerParams {
	p := &AssignCertToLoadBalancerParams{}
	p.p = make(map[string]interface{})
	p.p["certid"] = certid
	p.p["lbruleid"] = lbruleid
	return p
}

// Assigns a certificate to a load balancer rule
func (s *LoadBalancerService) AssignCertToLoadBalancer(p *AssignCertToLoadBalancerParams) (*AssignCertToLoadBalancerResponse, error) {
	resp, err := s.cs.newRequest("assignCertToLoadBalancer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AssignCertToLoadBalancerResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.Jobid, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type AssignCertToLoadBalancerResponse struct {
	Displaytext string `json:"displaytext"`
	Jobid       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type AssignToGlobalLoadBalancerRuleParams struct {
	p map[string]interface{}
}

func (p *AssignToGlobalLoadBalancerRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["gslblbruleweightsmap"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("gslblbruleweightsmap[%d].key", i), k)
			u.Set(fmt.Sprintf("gslblbruleweightsmap[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["loadbalancerrulelist"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("loadbalancerrulelist", vv)
	}
	return u
}

func (p *AssignToGlobalLoadBalancerRuleParams) SetGslblbruleweightsmap(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gslblbruleweightsmap"] = v
	return
}

func (p *AssignToGlobalLoadBalancerRuleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *AssignToGlobalLoadBalancerRuleParams) SetLoadbalancerrulelist(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["loadbalancerrulelist"] = v
	return
}

// You should always use this function to get a new AssignToGlobalLoadBalancerRuleParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewAssignToGlobalLoadBalancerRuleParams(id string, loadbalancerrulelist []string) *AssignToGlobalLoadBalancerRuleParams {
	p := &AssignToGlobalLoadBalancerRuleParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["loadbalancerrulelist"] = loadbalancerrulelist
	return p
}

// Assign load balancer rule or list of load balancer rules to a global load balancer rules.
func (s *LoadBalancerService) AssignToGlobalLoadBalancerRule(p *AssignToGlobalLoadBalancerRuleParams) (*AssignToGlobalLoadBalancerRuleResponse, error) {
	resp, err := s.cs.newRequest("assignToGlobalLoadBalancerRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AssignToGlobalLoadBalancerRuleResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.Jobid, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type AssignToGlobalLoadBalancerRuleResponse struct {
	Displaytext string `json:"displaytext"`
	Jobid       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type AssignToLoadBalancerRuleParams struct {
	p map[string]interface{}
}

func (p *AssignToLoadBalancerRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["virtualmachineids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("virtualmachineids", vv)
	}
	if v, found := p.p["vmidipmap"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("vmidipmap[%d].key", i), k)
			u.Set(fmt.Sprintf("vmidipmap[%d].value", i), vv)
			i++
		}
	}
	return u
}

func (p *AssignToLoadBalancerRuleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *AssignToLoadBalancerRuleParams) SetVirtualmachineids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineids"] = v
	return
}

func (p *AssignToLoadBalancerRuleParams) SetVmidipmap(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vmidipmap"] = v
	return
}

// You should always use this function to get a new AssignToLoadBalancerRuleParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewAssignToLoadBalancerRuleParams(id string) *AssignToLoadBalancerRuleParams {
	p := &AssignToLoadBalancerRuleParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Assigns virtual machine or a list of virtual machines to a load balancer rule.
func (s *LoadBalancerService) AssignToLoadBalancerRule(p *AssignToLoadBalancerRuleParams) (*AssignToLoadBalancerRuleResponse, error) {
	resp, err := s.cs.newRequest("assignToLoadBalancerRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AssignToLoadBalancerRuleResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.Jobid, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type AssignToLoadBalancerRuleResponse struct {
	Displaytext string `json:"displaytext"`
	Jobid       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ConfigureF5LoadBalancerParams struct {
	p map[string]interface{}
}

func (p *ConfigureF5LoadBalancerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["lbdevicecapacity"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("lbdevicecapacity", vv)
	}
	if v, found := p.p["lbdeviceid"]; found {
		u.Set("lbdeviceid", v.(string))
	}
	return u
}

func (p *ConfigureF5LoadBalancerParams) SetLbdevicecapacity(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbdevicecapacity"] = v
	return
}

func (p *ConfigureF5LoadBalancerParams) SetLbdeviceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbdeviceid"] = v
	return
}

// You should always use this function to get a new ConfigureF5LoadBalancerParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewConfigureF5LoadBalancerParams(lbdeviceid string) *ConfigureF5LoadBalancerParams {
	p := &ConfigureF5LoadBalancerParams{}
	p.p = make(map[string]interface{})
	p.p["lbdeviceid"] = lbdeviceid
	return p
}

// configures a F5 load balancer device
func (s *LoadBalancerService) ConfigureF5LoadBalancer(p *ConfigureF5LoadBalancerParams) (*F5LoadBalancerResponse, error) {
	resp, err := s.cs.newRequest("configureF5LoadBalancer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r F5LoadBalancerResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.Jobid, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type F5LoadBalancerResponse struct {
	Ipaddress         string `json:"ipaddress"`
	Jobid             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Lbdevicecapacity  int64  `json:"lbdevicecapacity"`
	Lbdevicededicated bool   `json:"lbdevicededicated"`
	Lbdeviceid        string `json:"lbdeviceid"`
	Lbdevicename      string `json:"lbdevicename"`
	Lbdevicestate     string `json:"lbdevicestate"`
	Physicalnetworkid string `json:"physicalnetworkid"`
	Privateinterface  string `json:"privateinterface"`
	Provider          string `json:"provider"`
	Publicinterface   string `json:"publicinterface"`
}

type ConfigureNetscalerLoadBalancerParams struct {
	p map[string]interface{}
}

func (p *ConfigureNetscalerLoadBalancerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["inline"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("inline", vv)
	}
	if v, found := p.p["lbdevicecapacity"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("lbdevicecapacity", vv)
	}
	if v, found := p.p["lbdevicededicated"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("lbdevicededicated", vv)
	}
	if v, found := p.p["lbdeviceid"]; found {
		u.Set("lbdeviceid", v.(string))
	}
	if v, found := p.p["podids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("podids", vv)
	}
	return u
}

func (p *ConfigureNetscalerLoadBalancerParams) SetInline(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["inline"] = v
	return
}

func (p *ConfigureNetscalerLoadBalancerParams) SetLbdevicecapacity(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbdevicecapacity"] = v
	return
}

func (p *ConfigureNetscalerLoadBalancerParams) SetLbdevicededicated(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbdevicededicated"] = v
	return
}

func (p *ConfigureNetscalerLoadBalancerParams) SetLbdeviceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbdeviceid"] = v
	return
}

func (p *ConfigureNetscalerLoadBalancerParams) SetPodids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podids"] = v
	return
}

// You should always use this function to get a new ConfigureNetscalerLoadBalancerParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewConfigureNetscalerLoadBalancerParams(lbdeviceid string) *ConfigureNetscalerLoadBalancerParams {
	p := &ConfigureNetscalerLoadBalancerParams{}
	p.p = make(map[string]interface{})
	p.p["lbdeviceid"] = lbdeviceid
	return p
}

// configures a netscaler load balancer device
func (s *LoadBalancerService) ConfigureNetscalerLoadBalancer(p *ConfigureNetscalerLoadBalancerParams) (*NetscalerLoadBalancerResponse, error) {
	resp, err := s.cs.newRequest("configureNetscalerLoadBalancer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r NetscalerLoadBalancerResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.Jobid, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type NetscalerLoadBalancerResponse struct {
	Gslbprovider            bool     `json:"gslbprovider"`
	Gslbproviderprivateip   string   `json:"gslbproviderprivateip"`
	Gslbproviderpublicip    string   `json:"gslbproviderpublicip"`
	Ipaddress               string   `json:"ipaddress"`
	Isexclusivegslbprovider bool     `json:"isexclusivegslbprovider"`
	Jobid                   string   `json:"jobid"`
	Jobstatus               int      `json:"jobstatus"`
	Lbdevicecapacity        int64    `json:"lbdevicecapacity"`
	Lbdevicededicated       bool     `json:"lbdevicededicated"`
	Lbdeviceid              string   `json:"lbdeviceid"`
	Lbdevicename            string   `json:"lbdevicename"`
	Lbdevicestate           string   `json:"lbdevicestate"`
	Physicalnetworkid       string   `json:"physicalnetworkid"`
	Podids                  []string `json:"podids"`
	Privateinterface        string   `json:"privateinterface"`
	Provider                string   `json:"provider"`
	Publicinterface         string   `json:"publicinterface"`
}

type CreateGlobalLoadBalancerRuleParams struct {
	p map[string]interface{}
}

func (p *CreateGlobalLoadBalancerRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["gslbdomainname"]; found {
		u.Set("gslbdomainname", v.(string))
	}
	if v, found := p.p["gslblbmethod"]; found {
		u.Set("gslblbmethod", v.(string))
	}
	if v, found := p.p["gslbservicetype"]; found {
		u.Set("gslbservicetype", v.(string))
	}
	if v, found := p.p["gslbstickysessionmethodname"]; found {
		u.Set("gslbstickysessionmethodname", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["regionid"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("regionid", vv)
	}
	return u
}

func (p *CreateGlobalLoadBalancerRuleParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *CreateGlobalLoadBalancerRuleParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
	return
}

func (p *CreateGlobalLoadBalancerRuleParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *CreateGlobalLoadBalancerRuleParams) SetGslbdomainname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gslbdomainname"] = v
	return
}

func (p *CreateGlobalLoadBalancerRuleParams) SetGslblbmethod(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gslblbmethod"] = v
	return
}

func (p *CreateGlobalLoadBalancerRuleParams) SetGslbservicetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gslbservicetype"] = v
	return
}

func (p *CreateGlobalLoadBalancerRuleParams) SetGslbstickysessionmethodname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gslbstickysessionmethodname"] = v
	return
}

func (p *CreateGlobalLoadBalancerRuleParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *CreateGlobalLoadBalancerRuleParams) SetRegionid(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["regionid"] = v
	return
}

// You should always use this function to get a new CreateGlobalLoadBalancerRuleParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewCreateGlobalLoadBalancerRuleParams(gslbdomainname string, gslbservicetype string, name string, regionid int) *CreateGlobalLoadBalancerRuleParams {
	p := &CreateGlobalLoadBalancerRuleParams{}
	p.p = make(map[string]interface{})
	p.p["gslbdomainname"] = gslbdomainname
	p.p["gslbservicetype"] = gslbservicetype
	p.p["name"] = name
	p.p["regionid"] = regionid
	return p
}

// Creates a global load balancer rule
func (s *LoadBalancerService) CreateGlobalLoadBalancerRule(p *CreateGlobalLoadBalancerRuleParams) (*CreateGlobalLoadBalancerRuleResponse, error) {
	resp, err := s.cs.newRequest("createGlobalLoadBalancerRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateGlobalLoadBalancerRuleResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.Jobid, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type CreateGlobalLoadBalancerRuleResponse struct {
	Account                     string `json:"account"`
	Description                 string `json:"description"`
	Domain                      string `json:"domain"`
	Domainid                    string `json:"domainid"`
	Gslbdomainname              string `json:"gslbdomainname"`
	Gslblbmethod                string `json:"gslblbmethod"`
	Gslbservicetype             string `json:"gslbservicetype"`
	Gslbstickysessionmethodname string `json:"gslbstickysessionmethodname"`
	Id                          string `json:"id"`
	Jobid                       string `json:"jobid"`
	Jobstatus                   int    `json:"jobstatus"`
	Loadbalancerrule            []struct {
		Account     string `json:"account"`
		Algorithm   string `json:"algorithm"`
		Cidrlist    string `json:"cidrlist"`
		Description string `json:"description"`
		Domain      string `json:"domain"`
		Domainid    string `json:"domainid"`
		Fordisplay  bool   `json:"fordisplay"`
		Id          string `json:"id"`
		Name        string `json:"name"`
		Networkid   string `json:"networkid"`
		Privateport string `json:"privateport"`
		Project     string `json:"project"`
		Projectid   string `json:"projectid"`
		Protocol    string `json:"protocol"`
		Publicip    string `json:"publicip"`
		Publicipid  string `json:"publicipid"`
		Publicport  string `json:"publicport"`
		State       string `json:"state"`
		Tags        []struct {
			Account      string `json:"account"`
			Customer     string `json:"customer"`
			Domain       string `json:"domain"`
			Domainid     string `json:"domainid"`
			Key          string `json:"key"`
			Project      string `json:"project"`
			Projectid    string `json:"projectid"`
			Resourceid   string `json:"resourceid"`
			Resourcetype string `json:"resourcetype"`
			Value        string `json:"value"`
		} `json:"tags"`
		Zoneid   string `json:"zoneid"`
		Zonename string `json:"zonename"`
	} `json:"loadbalancerrule"`
	Name      string `json:"name"`
	Project   string `json:"project"`
	Projectid string `json:"projectid"`
	Regionid  int    `json:"regionid"`
}

type CreateLBHealthCheckPolicyParams struct {
	p map[string]interface{}
}

func (p *CreateLBHealthCheckPolicyParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["healthythreshold"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("healthythreshold", vv)
	}
	if v, found := p.p["intervaltime"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("intervaltime", vv)
	}
	if v, found := p.p["lbruleid"]; found {
		u.Set("lbruleid", v.(string))
	}
	if v, found := p.p["pingpath"]; found {
		u.Set("pingpath", v.(string))
	}
	if v, found := p.p["responsetimeout"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("responsetimeout", vv)
	}
	if v, found := p.p["unhealthythreshold"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("unhealthythreshold", vv)
	}
	return u
}

func (p *CreateLBHealthCheckPolicyParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
	return
}

func (p *CreateLBHealthCheckPolicyParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
	return
}

func (p *CreateLBHealthCheckPolicyParams) SetHealthythreshold(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["healthythreshold"] = v
	return
}

func (p *CreateLBHealthCheckPolicyParams) SetIntervaltime(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["intervaltime"] = v
	return
}

func (p *CreateLBHealthCheckPolicyParams) SetLbruleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbruleid"] = v
	return
}

func (p *CreateLBHealthCheckPolicyParams) SetPingpath(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pingpath"] = v
	return
}

func (p *CreateLBHealthCheckPolicyParams) SetResponsetimeout(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["responsetimeout"] = v
	return
}

func (p *CreateLBHealthCheckPolicyParams) SetUnhealthythreshold(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["unhealthythreshold"] = v
	return
}

// You should always use this function to get a new CreateLBHealthCheckPolicyParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewCreateLBHealthCheckPolicyParams(lbruleid string) *CreateLBHealthCheckPolicyParams {
	p := &CreateLBHealthCheckPolicyParams{}
	p.p = make(map[string]interface{})
	p.p["lbruleid"] = lbruleid
	return p
}

// Creates a load balancer health check policy
func (s *LoadBalancerService) CreateLBHealthCheckPolicy(p *CreateLBHealthCheckPolicyParams) (*CreateLBHealthCheckPolicyResponse, error) {
	resp, err := s.cs.newRequest("createLBHealthCheckPolicy", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateLBHealthCheckPolicyResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.Jobid, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type CreateLBHealthCheckPolicyResponse struct {
	Account           string `json:"account"`
	Domain            string `json:"domain"`
	Domainid          string `json:"domainid"`
	Healthcheckpolicy []struct {
		Description             string `json:"description"`
		Fordisplay              bool   `json:"fordisplay"`
		Healthcheckinterval     int    `json:"healthcheckinterval"`
		Healthcheckthresshold   int    `json:"healthcheckthresshold"`
		Id                      string `json:"id"`
		Pingpath                string `json:"pingpath"`
		Responsetime            int    `json:"responsetime"`
		State                   string `json:"state"`
		Unhealthcheckthresshold int    `json:"unhealthcheckthresshold"`
	} `json:"healthcheckpolicy"`
	Jobid     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Lbruleid  string `json:"lbruleid"`
	Zoneid    string `json:"zoneid"`
}

type CreateLBStickinessPolicyParams struct {
	p map[string]interface{}
}

func (p *CreateLBStickinessPolicyParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["lbruleid"]; found {
		u.Set("lbruleid", v.(string))
	}
	if v, found := p.p["methodname"]; found {
		u.Set("methodname", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["param"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("param[%d].key", i), k)
			u.Set(fmt.Sprintf("param[%d].value", i), vv)
			i++
		}
	}
	return u
}

func (p *CreateLBStickinessPolicyParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
	return
}

func (p *CreateLBStickinessPolicyParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
	return
}

func (p *CreateLBStickinessPolicyParams) SetLbruleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbruleid"] = v
	return
}

func (p *CreateLBStickinessPolicyParams) SetMethodname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["methodname"] = v
	return
}

func (p *CreateLBStickinessPolicyParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *CreateLBStickinessPolicyParams) SetParam(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["param"] = v
	return
}

// You should always use this function to get a new CreateLBStickinessPolicyParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewCreateLBStickinessPolicyParams(lbruleid string, methodname string, name string) *CreateLBStickinessPolicyParams {
	p := &CreateLBStickinessPolicyParams{}
	p.p = make(map[string]interface{})
	p.p["lbruleid"] = lbruleid
	p.p["methodname"] = methodname
	p.p["name"] = name
	return p
}

// Creates a load balancer stickiness policy
func (s *LoadBalancerService) CreateLBStickinessPolicy(p *CreateLBStickinessPolicyParams) (*CreateLBStickinessPolicyResponse, error) {
	resp, err := s.cs.newRequest("createLBStickinessPolicy", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateLBStickinessPolicyResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.Jobid, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type CreateLBStickinessPolicyResponse struct {
	Account          string `json:"account"`
	Description      string `json:"description"`
	Domain           string `json:"domain"`
	Domainid         string `json:"domainid"`
	Jobid            string `json:"jobid"`
	Jobstatus        int    `json:"jobstatus"`
	Lbruleid         string `json:"lbruleid"`
	Name             string `json:"name"`
	State            string `json:"state"`
	Stickinesspolicy []struct {
		Description string            `json:"description"`
		Fordisplay  bool              `json:"fordisplay"`
		Id          string            `json:"id"`
		Methodname  string            `json:"methodname"`
		Name        string            `json:"name"`
		Params      map[string]string `json:"params"`
		State       string            `json:"state"`
	} `json:"stickinesspolicy"`
	Zoneid string `json:"zoneid"`
}

type CreateLoadBalancerParams struct {
	p map[string]interface{}
}

func (p *CreateLoadBalancerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["algorithm"]; found {
		u.Set("algorithm", v.(string))
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["instanceport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("instanceport", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := p.p["scheme"]; found {
		u.Set("scheme", v.(string))
	}
	if v, found := p.p["sourceipaddress"]; found {
		u.Set("sourceipaddress", v.(string))
	}
	if v, found := p.p["sourceipaddressnetworkid"]; found {
		u.Set("sourceipaddressnetworkid", v.(string))
	}
	if v, found := p.p["sourceport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("sourceport", vv)
	}
	return u
}

func (p *CreateLoadBalancerParams) SetAlgorithm(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["algorithm"] = v
	return
}

func (p *CreateLoadBalancerParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
	return
}

func (p *CreateLoadBalancerParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
	return
}

func (p *CreateLoadBalancerParams) SetInstanceport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["instanceport"] = v
	return
}

func (p *CreateLoadBalancerParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *CreateLoadBalancerParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
	return
}

func (p *CreateLoadBalancerParams) SetScheme(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["scheme"] = v
	return
}

func (p *CreateLoadBalancerParams) SetSourceipaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sourceipaddress"] = v
	return
}

func (p *CreateLoadBalancerParams) SetSourceipaddressnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sourceipaddressnetworkid"] = v
	return
}

func (p *CreateLoadBalancerParams) SetSourceport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sourceport"] = v
	return
}

// You should always use this function to get a new CreateLoadBalancerParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewCreateLoadBalancerParams(algorithm string, instanceport int, name string, networkid string, scheme string, sourceipaddressnetworkid string, sourceport int) *CreateLoadBalancerParams {
	p := &CreateLoadBalancerParams{}
	p.p = make(map[string]interface{})
	p.p["algorithm"] = algorithm
	p.p["instanceport"] = instanceport
	p.p["name"] = name
	p.p["networkid"] = networkid
	p.p["scheme"] = scheme
	p.p["sourceipaddressnetworkid"] = sourceipaddressnetworkid
	p.p["sourceport"] = sourceport
	return p
}

// Creates a load balancer
func (s *LoadBalancerService) CreateLoadBalancer(p *CreateLoadBalancerParams) (*CreateLoadBalancerResponse, error) {
	resp, err := s.cs.newRequest("createLoadBalancer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateLoadBalancerResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.Jobid, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type CreateLoadBalancerResponse struct {
	Account              string `json:"account"`
	Algorithm            string `json:"algorithm"`
	Description          string `json:"description"`
	Domain               string `json:"domain"`
	Domainid             string `json:"domainid"`
	Fordisplay           bool   `json:"fordisplay"`
	Id                   string `json:"id"`
	Jobid                string `json:"jobid"`
	Jobstatus            int    `json:"jobstatus"`
	Loadbalancerinstance []struct {
		Id        string `json:"id"`
		Ipaddress string `json:"ipaddress"`
		Name      string `json:"name"`
		State     string `json:"state"`
	} `json:"loadbalancerinstance"`
	Loadbalancerrule []struct {
		Instanceport int    `json:"instanceport"`
		Sourceport   int    `json:"sourceport"`
		State        string `json:"state"`
	} `json:"loadbalancerrule"`
	Name                     string `json:"name"`
	Networkid                string `json:"networkid"`
	Project                  string `json:"project"`
	Projectid                string `json:"projectid"`
	Sourceipaddress          string `json:"sourceipaddress"`
	Sourceipaddressnetworkid string `json:"sourceipaddressnetworkid"`
	Tags                     []struct {
		Account      string `json:"account"`
		Customer     string `json:"customer"`
		Domain       string `json:"domain"`
		Domainid     string `json:"domainid"`
		Key          string `json:"key"`
		Project      string `json:"project"`
		Projectid    string `json:"projectid"`
		Resourceid   string `json:"resourceid"`
		Resourcetype string `json:"resourcetype"`
		Value        string `json:"value"`
	} `json:"tags"`
}

type CreateLoadBalancerRuleParams struct {
	p map[string]interface{}
}

func (p *CreateLoadBalancerRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["algorithm"]; found {
		u.Set("algorithm", v.(string))
	}
	if v, found := p.p["cidrlist"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("cidrlist", vv)
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := p.p["openfirewall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("openfirewall", vv)
	}
	if v, found := p.p["privateport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("privateport", vv)
	}
	if v, found := p.p["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := p.p["publicipid"]; found {
		u.Set("publicipid", v.(string))
	}
	if v, found := p.p["publicport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("publicport", vv)
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *CreateLoadBalancerRuleParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *CreateLoadBalancerRuleParams) SetAlgorithm(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["algorithm"] = v
	return
}

func (p *CreateLoadBalancerRuleParams) SetCidrlist(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cidrlist"] = v
	return
}

func (p *CreateLoadBalancerRuleParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
	return
}

func (p *CreateLoadBalancerRuleParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *CreateLoadBalancerRuleParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
	return
}

func (p *CreateLoadBalancerRuleParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *CreateLoadBalancerRuleParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
	return
}

func (p *CreateLoadBalancerRuleParams) SetOpenfirewall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["openfirewall"] = v
	return
}

func (p *CreateLoadBalancerRuleParams) SetPrivateport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["privateport"] = v
	return
}

func (p *CreateLoadBalancerRuleParams) SetProtocol(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["protocol"] = v
	return
}

func (p *CreateLoadBalancerRuleParams) SetPublicipid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["publicipid"] = v
	return
}

func (p *CreateLoadBalancerRuleParams) SetPublicport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["publicport"] = v
	return
}

func (p *CreateLoadBalancerRuleParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new CreateLoadBalancerRuleParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewCreateLoadBalancerRuleParams(algorithm string, name string, privateport int, publicport int) *CreateLoadBalancerRuleParams {
	p := &CreateLoadBalancerRuleParams{}
	p.p = make(map[string]interface{})
	p.p["algorithm"] = algorithm
	p.p["name"] = name
	p.p["privateport"] = privateport
	p.p["publicport"] = publicport
	return p
}

// Creates a load balancer rule
func (s *LoadBalancerService) CreateLoadBalancerRule(p *CreateLoadBalancerRuleParams) (*CreateLoadBalancerRuleResponse, error) {
	resp, err := s.cs.newRequest("createLoadBalancerRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateLoadBalancerRuleResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.Jobid, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type CreateLoadBalancerRuleResponse struct {
	Account     string `json:"account"`
	Algorithm   string `json:"algorithm"`
	Cidrlist    string `json:"cidrlist"`
	Description string `json:"description"`
	Domain      string `json:"domain"`
	Domainid    string `json:"domainid"`
	Fordisplay  bool   `json:"fordisplay"`
	Id          string `json:"id"`
	Jobid       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Name        string `json:"name"`
	Networkid   string `json:"networkid"`
	Privateport string `json:"privateport"`
	Project     string `json:"project"`
	Projectid   string `json:"projectid"`
	Protocol    string `json:"protocol"`
	Publicip    string `json:"publicip"`
	Publicipid  string `json:"publicipid"`
	Publicport  string `json:"publicport"`
	State       string `json:"state"`
	Tags        []struct {
		Account      string `json:"account"`
		Customer     string `json:"customer"`
		Domain       string `json:"domain"`
		Domainid     string `json:"domainid"`
		Key          string `json:"key"`
		Project      string `json:"project"`
		Projectid    string `json:"projectid"`
		Resourceid   string `json:"resourceid"`
		Resourcetype string `json:"resourcetype"`
		Value        string `json:"value"`
	} `json:"tags"`
	Zoneid   string `json:"zoneid"`
	Zonename string `json:"zonename"`
}

type DeleteF5LoadBalancerParams struct {
	p map[string]interface{}
}

func (p *DeleteF5LoadBalancerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["lbdeviceid"]; found {
		u.Set("lbdeviceid", v.(string))
	}
	return u
}

func (p *DeleteF5LoadBalancerParams) SetLbdeviceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbdeviceid"] = v
	return
}

// You should always use this function to get a new DeleteF5LoadBalancerParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewDeleteF5LoadBalancerParams(lbdeviceid string) *DeleteF5LoadBalancerParams {
	p := &DeleteF5LoadBalancerParams{}
	p.p = make(map[string]interface{})
	p.p["lbdeviceid"] = lbdeviceid
	return p
}

//  delete a F5 load balancer device
func (s *LoadBalancerService) DeleteF5LoadBalancer(p *DeleteF5LoadBalancerParams) (*DeleteF5LoadBalancerResponse, error) {
	resp, err := s.cs.newRequest("deleteF5LoadBalancer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteF5LoadBalancerResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.Jobid, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type DeleteF5LoadBalancerResponse struct {
	Displaytext string `json:"displaytext"`
	Jobid       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteGlobalLoadBalancerRuleParams struct {
	p map[string]interface{}
}

func (p *DeleteGlobalLoadBalancerRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteGlobalLoadBalancerRuleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteGlobalLoadBalancerRuleParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewDeleteGlobalLoadBalancerRuleParams(id string) *DeleteGlobalLoadBalancerRuleParams {
	p := &DeleteGlobalLoadBalancerRuleParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a global load balancer rule.
func (s *LoadBalancerService) DeleteGlobalLoadBalancerRule(p *DeleteGlobalLoadBalancerRuleParams) (*DeleteGlobalLoadBalancerRuleResponse, error) {
	resp, err := s.cs.newRequest("deleteGlobalLoadBalancerRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteGlobalLoadBalancerRuleResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.Jobid, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type DeleteGlobalLoadBalancerRuleResponse struct {
	Displaytext string `json:"displaytext"`
	Jobid       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteLBHealthCheckPolicyParams struct {
	p map[string]interface{}
}

func (p *DeleteLBHealthCheckPolicyParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteLBHealthCheckPolicyParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteLBHealthCheckPolicyParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewDeleteLBHealthCheckPolicyParams(id string) *DeleteLBHealthCheckPolicyParams {
	p := &DeleteLBHealthCheckPolicyParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a load balancer health check policy.
func (s *LoadBalancerService) DeleteLBHealthCheckPolicy(p *DeleteLBHealthCheckPolicyParams) (*DeleteLBHealthCheckPolicyResponse, error) {
	resp, err := s.cs.newRequest("deleteLBHealthCheckPolicy", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteLBHealthCheckPolicyResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.Jobid, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type DeleteLBHealthCheckPolicyResponse struct {
	Displaytext string `json:"displaytext"`
	Jobid       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteLBStickinessPolicyParams struct {
	p map[string]interface{}
}

func (p *DeleteLBStickinessPolicyParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteLBStickinessPolicyParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteLBStickinessPolicyParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewDeleteLBStickinessPolicyParams(id string) *DeleteLBStickinessPolicyParams {
	p := &DeleteLBStickinessPolicyParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a load balancer stickiness policy.
func (s *LoadBalancerService) DeleteLBStickinessPolicy(p *DeleteLBStickinessPolicyParams) (*DeleteLBStickinessPolicyResponse, error) {
	resp, err := s.cs.newRequest("deleteLBStickinessPolicy", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteLBStickinessPolicyResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.Jobid, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type DeleteLBStickinessPolicyResponse struct {
	Displaytext string `json:"displaytext"`
	Jobid       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteLoadBalancerParams struct {
	p map[string]interface{}
}

func (p *DeleteLoadBalancerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteLoadBalancerParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteLoadBalancerParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewDeleteLoadBalancerParams(id string) *DeleteLoadBalancerParams {
	p := &DeleteLoadBalancerParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a load balancer
func (s *LoadBalancerService) DeleteLoadBalancer(p *DeleteLoadBalancerParams) (*DeleteLoadBalancerResponse, error) {
	resp, err := s.cs.newRequest("deleteLoadBalancer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteLoadBalancerResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.Jobid, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type DeleteLoadBalancerResponse struct {
	Displaytext string `json:"displaytext"`
	Jobid       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteLoadBalancerRuleParams struct {
	p map[string]interface{}
}

func (p *DeleteLoadBalancerRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteLoadBalancerRuleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteLoadBalancerRuleParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewDeleteLoadBalancerRuleParams(id string) *DeleteLoadBalancerRuleParams {
	p := &DeleteLoadBalancerRuleParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a load balancer rule.
func (s *LoadBalancerService) DeleteLoadBalancerRule(p *DeleteLoadBalancerRuleParams) (*DeleteLoadBalancerRuleResponse, error) {
	resp, err := s.cs.newRequest("deleteLoadBalancerRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteLoadBalancerRuleResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.Jobid, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type DeleteLoadBalancerRuleResponse struct {
	Displaytext string `json:"displaytext"`
	Jobid       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteNetscalerLoadBalancerParams struct {
	p map[string]interface{}
}

func (p *DeleteNetscalerLoadBalancerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["lbdeviceid"]; found {
		u.Set("lbdeviceid", v.(string))
	}
	return u
}

func (p *DeleteNetscalerLoadBalancerParams) SetLbdeviceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbdeviceid"] = v
	return
}

// You should always use this function to get a new DeleteNetscalerLoadBalancerParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewDeleteNetscalerLoadBalancerParams(lbdeviceid string) *DeleteNetscalerLoadBalancerParams {
	p := &DeleteNetscalerLoadBalancerParams{}
	p.p = make(map[string]interface{})
	p.p["lbdeviceid"] = lbdeviceid
	return p
}

//  delete a netscaler load balancer device
func (s *LoadBalancerService) DeleteNetscalerLoadBalancer(p *DeleteNetscalerLoadBalancerParams) (*DeleteNetscalerLoadBalancerResponse, error) {
	resp, err := s.cs.newRequest("deleteNetscalerLoadBalancer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteNetscalerLoadBalancerResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.Jobid, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type DeleteNetscalerLoadBalancerResponse struct {
	Displaytext string `json:"displaytext"`
	Jobid       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteSslCertParams struct {
	p map[string]interface{}
}

func (p *DeleteSslCertParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteSslCertParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteSslCertParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewDeleteSslCertParams(id string) *DeleteSslCertParams {
	p := &DeleteSslCertParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Delete a certificate to CloudStack
func (s *LoadBalancerService) DeleteSslCert(p *DeleteSslCertParams) (*DeleteSslCertResponse, error) {
	resp, err := s.cs.newRequest("deleteSslCert", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteSslCertResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteSslCertResponse struct {
	Displaytext string `json:"displaytext"`
	Jobid       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteSslCertResponse) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	if success, ok := m["success"].(string); ok {
		m["success"] = success == "true"
		b, err = json.Marshal(m)
		if err != nil {
			return err
		}
	}

	type alias DeleteSslCertResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ListF5LoadBalancersParams struct {
	p map[string]interface{}
}

func (p *ListF5LoadBalancersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["lbdeviceid"]; found {
		u.Set("lbdeviceid", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	return u
}

func (p *ListF5LoadBalancersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListF5LoadBalancersParams) SetLbdeviceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbdeviceid"] = v
	return
}

func (p *ListF5LoadBalancersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListF5LoadBalancersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListF5LoadBalancersParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
	return
}

// You should always use this function to get a new ListF5LoadBalancersParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewListF5LoadBalancersParams() *ListF5LoadBalancersParams {
	p := &ListF5LoadBalancersParams{}
	p.p = make(map[string]interface{})
	return p
}

// lists F5 load balancer devices
func (s *LoadBalancerService) ListF5LoadBalancers(p *ListF5LoadBalancersParams) (*ListF5LoadBalancersResponse, error) {
	resp, err := s.cs.newRequest("listF5LoadBalancers", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListF5LoadBalancersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListF5LoadBalancersResponse struct {
	Count           int               `json:"count"`
	F5LoadBalancers []*F5LoadBalancer `json:"f5loadbalancer"`
}

type F5LoadBalancer struct {
	Ipaddress         string `json:"ipaddress"`
	Jobid             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Lbdevicecapacity  int64  `json:"lbdevicecapacity"`
	Lbdevicededicated bool   `json:"lbdevicededicated"`
	Lbdeviceid        string `json:"lbdeviceid"`
	Lbdevicename      string `json:"lbdevicename"`
	Lbdevicestate     string `json:"lbdevicestate"`
	Physicalnetworkid string `json:"physicalnetworkid"`
	Privateinterface  string `json:"privateinterface"`
	Provider          string `json:"provider"`
	Publicinterface   string `json:"publicinterface"`
}

type ListGlobalLoadBalancerRulesParams struct {
	p map[string]interface{}
}

func (p *ListGlobalLoadBalancerRulesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["regionid"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("regionid", vv)
	}
	if v, found := p.p["tags"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), vv)
			i++
		}
	}
	return u
}

func (p *ListGlobalLoadBalancerRulesParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListGlobalLoadBalancerRulesParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListGlobalLoadBalancerRulesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListGlobalLoadBalancerRulesParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
	return
}

func (p *ListGlobalLoadBalancerRulesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListGlobalLoadBalancerRulesParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
	return
}

func (p *ListGlobalLoadBalancerRulesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListGlobalLoadBalancerRulesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListGlobalLoadBalancerRulesParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *ListGlobalLoadBalancerRulesParams) SetRegionid(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["regionid"] = v
	return
}

func (p *ListGlobalLoadBalancerRulesParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
	return
}

// You should always use this function to get a new ListGlobalLoadBalancerRulesParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewListGlobalLoadBalancerRulesParams() *ListGlobalLoadBalancerRulesParams {
	p := &ListGlobalLoadBalancerRulesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetGlobalLoadBalancerRuleID(keyword string, opts ...OptionFunc) (string, int, error) {
	p := &ListGlobalLoadBalancerRulesParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListGlobalLoadBalancerRules(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", keyword, l)
	}

	if l.Count == 1 {
		return l.GlobalLoadBalancerRules[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.GlobalLoadBalancerRules {
			if v.Name == keyword {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", keyword, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetGlobalLoadBalancerRuleByName(name string, opts ...OptionFunc) (*GlobalLoadBalancerRule, int, error) {
	id, count, err := s.GetGlobalLoadBalancerRuleID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetGlobalLoadBalancerRuleByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetGlobalLoadBalancerRuleByID(id string, opts ...OptionFunc) (*GlobalLoadBalancerRule, int, error) {
	p := &ListGlobalLoadBalancerRulesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListGlobalLoadBalancerRules(p)
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf(
			"Invalid parameter id value=%s due to incorrect long value format, "+
				"or entity does not exist", id)) {
			return nil, 0, fmt.Errorf("No match found for %s: %+v", id, l)
		}
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.GlobalLoadBalancerRules[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for GlobalLoadBalancerRule UUID: %s!", id)
}

// Lists load balancer rules.
func (s *LoadBalancerService) ListGlobalLoadBalancerRules(p *ListGlobalLoadBalancerRulesParams) (*ListGlobalLoadBalancerRulesResponse, error) {
	resp, err := s.cs.newRequest("listGlobalLoadBalancerRules", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListGlobalLoadBalancerRulesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListGlobalLoadBalancerRulesResponse struct {
	Count                   int                       `json:"count"`
	GlobalLoadBalancerRules []*GlobalLoadBalancerRule `json:"globalloadbalancerrule"`
}

type GlobalLoadBalancerRule struct {
	Account                     string `json:"account"`
	Description                 string `json:"description"`
	Domain                      string `json:"domain"`
	Domainid                    string `json:"domainid"`
	Gslbdomainname              string `json:"gslbdomainname"`
	Gslblbmethod                string `json:"gslblbmethod"`
	Gslbservicetype             string `json:"gslbservicetype"`
	Gslbstickysessionmethodname string `json:"gslbstickysessionmethodname"`
	Id                          string `json:"id"`
	Jobid                       string `json:"jobid"`
	Jobstatus                   int    `json:"jobstatus"`
	Loadbalancerrule            []struct {
		Account     string `json:"account"`
		Algorithm   string `json:"algorithm"`
		Cidrlist    string `json:"cidrlist"`
		Description string `json:"description"`
		Domain      string `json:"domain"`
		Domainid    string `json:"domainid"`
		Fordisplay  bool   `json:"fordisplay"`
		Id          string `json:"id"`
		Name        string `json:"name"`
		Networkid   string `json:"networkid"`
		Privateport string `json:"privateport"`
		Project     string `json:"project"`
		Projectid   string `json:"projectid"`
		Protocol    string `json:"protocol"`
		Publicip    string `json:"publicip"`
		Publicipid  string `json:"publicipid"`
		Publicport  string `json:"publicport"`
		State       string `json:"state"`
		Tags        []struct {
			Account      string `json:"account"`
			Customer     string `json:"customer"`
			Domain       string `json:"domain"`
			Domainid     string `json:"domainid"`
			Key          string `json:"key"`
			Project      string `json:"project"`
			Projectid    string `json:"projectid"`
			Resourceid   string `json:"resourceid"`
			Resourcetype string `json:"resourcetype"`
			Value        string `json:"value"`
		} `json:"tags"`
		Zoneid   string `json:"zoneid"`
		Zonename string `json:"zonename"`
	} `json:"loadbalancerrule"`
	Name      string `json:"name"`
	Project   string `json:"project"`
	Projectid string `json:"projectid"`
	Regionid  int    `json:"regionid"`
}

type ListLBHealthCheckPoliciesParams struct {
	p map[string]interface{}
}

func (p *ListLBHealthCheckPoliciesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["lbruleid"]; found {
		u.Set("lbruleid", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	return u
}

func (p *ListLBHealthCheckPoliciesParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
	return
}

func (p *ListLBHealthCheckPoliciesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListLBHealthCheckPoliciesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListLBHealthCheckPoliciesParams) SetLbruleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbruleid"] = v
	return
}

func (p *ListLBHealthCheckPoliciesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListLBHealthCheckPoliciesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

// You should always use this function to get a new ListLBHealthCheckPoliciesParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewListLBHealthCheckPoliciesParams() *ListLBHealthCheckPoliciesParams {
	p := &ListLBHealthCheckPoliciesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetLBHealthCheckPolicyByID(id string, opts ...OptionFunc) (*LBHealthCheckPolicy, int, error) {
	p := &ListLBHealthCheckPoliciesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListLBHealthCheckPolicies(p)
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf(
			"Invalid parameter id value=%s due to incorrect long value format, "+
				"or entity does not exist", id)) {
			return nil, 0, fmt.Errorf("No match found for %s: %+v", id, l)
		}
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.LBHealthCheckPolicies[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for LBHealthCheckPolicy UUID: %s!", id)
}

// Lists load balancer health check policies.
func (s *LoadBalancerService) ListLBHealthCheckPolicies(p *ListLBHealthCheckPoliciesParams) (*ListLBHealthCheckPoliciesResponse, error) {
	resp, err := s.cs.newRequest("listLBHealthCheckPolicies", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListLBHealthCheckPoliciesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListLBHealthCheckPoliciesResponse struct {
	Count                 int                    `json:"count"`
	LBHealthCheckPolicies []*LBHealthCheckPolicy `json:"lbhealthcheckpolicy"`
}

type LBHealthCheckPolicy struct {
	Account           string `json:"account"`
	Domain            string `json:"domain"`
	Domainid          string `json:"domainid"`
	Healthcheckpolicy []struct {
		Description             string `json:"description"`
		Fordisplay              bool   `json:"fordisplay"`
		Healthcheckinterval     int    `json:"healthcheckinterval"`
		Healthcheckthresshold   int    `json:"healthcheckthresshold"`
		Id                      string `json:"id"`
		Pingpath                string `json:"pingpath"`
		Responsetime            int    `json:"responsetime"`
		State                   string `json:"state"`
		Unhealthcheckthresshold int    `json:"unhealthcheckthresshold"`
	} `json:"healthcheckpolicy"`
	Jobid     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Lbruleid  string `json:"lbruleid"`
	Zoneid    string `json:"zoneid"`
}

type ListLBStickinessPoliciesParams struct {
	p map[string]interface{}
}

func (p *ListLBStickinessPoliciesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["lbruleid"]; found {
		u.Set("lbruleid", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	return u
}

func (p *ListLBStickinessPoliciesParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
	return
}

func (p *ListLBStickinessPoliciesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListLBStickinessPoliciesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListLBStickinessPoliciesParams) SetLbruleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbruleid"] = v
	return
}

func (p *ListLBStickinessPoliciesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListLBStickinessPoliciesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

// You should always use this function to get a new ListLBStickinessPoliciesParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewListLBStickinessPoliciesParams() *ListLBStickinessPoliciesParams {
	p := &ListLBStickinessPoliciesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetLBStickinessPolicyByID(id string, opts ...OptionFunc) (*LBStickinessPolicy, int, error) {
	p := &ListLBStickinessPoliciesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListLBStickinessPolicies(p)
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf(
			"Invalid parameter id value=%s due to incorrect long value format, "+
				"or entity does not exist", id)) {
			return nil, 0, fmt.Errorf("No match found for %s: %+v", id, l)
		}
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.LBStickinessPolicies[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for LBStickinessPolicy UUID: %s!", id)
}

// Lists load balancer stickiness policies.
func (s *LoadBalancerService) ListLBStickinessPolicies(p *ListLBStickinessPoliciesParams) (*ListLBStickinessPoliciesResponse, error) {
	resp, err := s.cs.newRequest("listLBStickinessPolicies", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListLBStickinessPoliciesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListLBStickinessPoliciesResponse struct {
	Count                int                   `json:"count"`
	LBStickinessPolicies []*LBStickinessPolicy `json:"lbstickinesspolicy"`
}

type LBStickinessPolicy struct {
	Account          string `json:"account"`
	Description      string `json:"description"`
	Domain           string `json:"domain"`
	Domainid         string `json:"domainid"`
	Jobid            string `json:"jobid"`
	Jobstatus        int    `json:"jobstatus"`
	Lbruleid         string `json:"lbruleid"`
	Name             string `json:"name"`
	State            string `json:"state"`
	Stickinesspolicy []struct {
		Description string            `json:"description"`
		Fordisplay  bool              `json:"fordisplay"`
		Id          string            `json:"id"`
		Methodname  string            `json:"methodname"`
		Name        string            `json:"name"`
		Params      map[string]string `json:"params"`
		State       string            `json:"state"`
	} `json:"stickinesspolicy"`
	Zoneid string `json:"zoneid"`
}

type ListLoadBalancerRuleInstancesParams struct {
	p map[string]interface{}
}

func (p *ListLoadBalancerRuleInstancesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["applied"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("applied", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["lbvmips"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("lbvmips", vv)
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	return u
}

func (p *ListLoadBalancerRuleInstancesParams) SetApplied(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["applied"] = v
	return
}

func (p *ListLoadBalancerRuleInstancesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListLoadBalancerRuleInstancesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListLoadBalancerRuleInstancesParams) SetLbvmips(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbvmips"] = v
	return
}

func (p *ListLoadBalancerRuleInstancesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListLoadBalancerRuleInstancesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

// You should always use this function to get a new ListLoadBalancerRuleInstancesParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewListLoadBalancerRuleInstancesParams(id string) *ListLoadBalancerRuleInstancesParams {
	p := &ListLoadBalancerRuleInstancesParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetLoadBalancerRuleInstanceByID(id string, opts ...OptionFunc) (*VirtualMachine, int, error) {
	p := &ListLoadBalancerRuleInstancesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListLoadBalancerRuleInstances(p)
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf(
			"Invalid parameter id value=%s due to incorrect long value format, "+
				"or entity does not exist", id)) {
			return nil, 0, fmt.Errorf("No match found for %s: %+v", id, l)
		}
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.LoadBalancerRuleInstances[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for LoadBalancerRuleInstance UUID: %s!", id)
}

// List all virtual machine instances that are assigned to a load balancer rule.
func (s *LoadBalancerService) ListLoadBalancerRuleInstances(p *ListLoadBalancerRuleInstancesParams) (*ListLoadBalancerRuleInstancesResponse, error) {
	resp, err := s.cs.newRequest("listLoadBalancerRuleInstances", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListLoadBalancerRuleInstancesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListLoadBalancerRuleInstancesResponse struct {
	Count                     int                         `json:"count"`
	LBRuleVMIDIPs             []*LoadBalancerRuleInstance `json:"lbrulevmidip"`
	LoadBalancerRuleInstances []*VirtualMachine           `json:"loadbalancerruleinstance"`
}

type LoadBalancerRuleInstance struct {
	Jobid                    string          `json:"jobid"`
	Jobstatus                int             `json:"jobstatus"`
	Lbvmipaddresses          []string        `json:"lbvmipaddresses"`
	Loadbalancerruleinstance *VirtualMachine `json:"loadbalancerruleinstance"`
}

type ListLoadBalancerRulesParams struct {
	p map[string]interface{}
}

func (p *ListLoadBalancerRulesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["publicipid"]; found {
		u.Set("publicipid", v.(string))
	}
	if v, found := p.p["tags"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListLoadBalancerRulesParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListLoadBalancerRulesParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListLoadBalancerRulesParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
	return
}

func (p *ListLoadBalancerRulesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListLoadBalancerRulesParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
	return
}

func (p *ListLoadBalancerRulesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListLoadBalancerRulesParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
	return
}

func (p *ListLoadBalancerRulesParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *ListLoadBalancerRulesParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
	return
}

func (p *ListLoadBalancerRulesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListLoadBalancerRulesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListLoadBalancerRulesParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *ListLoadBalancerRulesParams) SetPublicipid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["publicipid"] = v
	return
}

func (p *ListLoadBalancerRulesParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
	return
}

func (p *ListLoadBalancerRulesParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
	return
}

func (p *ListLoadBalancerRulesParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new ListLoadBalancerRulesParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewListLoadBalancerRulesParams() *ListLoadBalancerRulesParams {
	p := &ListLoadBalancerRulesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetLoadBalancerRuleID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListLoadBalancerRulesParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListLoadBalancerRules(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.LoadBalancerRules[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.LoadBalancerRules {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetLoadBalancerRuleByName(name string, opts ...OptionFunc) (*LoadBalancerRule, int, error) {
	id, count, err := s.GetLoadBalancerRuleID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetLoadBalancerRuleByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetLoadBalancerRuleByID(id string, opts ...OptionFunc) (*LoadBalancerRule, int, error) {
	p := &ListLoadBalancerRulesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListLoadBalancerRules(p)
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf(
			"Invalid parameter id value=%s due to incorrect long value format, "+
				"or entity does not exist", id)) {
			return nil, 0, fmt.Errorf("No match found for %s: %+v", id, l)
		}
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.LoadBalancerRules[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for LoadBalancerRule UUID: %s!", id)
}

// Lists load balancer rules.
func (s *LoadBalancerService) ListLoadBalancerRules(p *ListLoadBalancerRulesParams) (*ListLoadBalancerRulesResponse, error) {
	resp, err := s.cs.newRequest("listLoadBalancerRules", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListLoadBalancerRulesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListLoadBalancerRulesResponse struct {
	Count             int                 `json:"count"`
	LoadBalancerRules []*LoadBalancerRule `json:"loadbalancerrule"`
}

type LoadBalancerRule struct {
	Account     string `json:"account"`
	Algorithm   string `json:"algorithm"`
	Cidrlist    string `json:"cidrlist"`
	Description string `json:"description"`
	Domain      string `json:"domain"`
	Domainid    string `json:"domainid"`
	Fordisplay  bool   `json:"fordisplay"`
	Id          string `json:"id"`
	Jobid       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Name        string `json:"name"`
	Networkid   string `json:"networkid"`
	Privateport string `json:"privateport"`
	Project     string `json:"project"`
	Projectid   string `json:"projectid"`
	Protocol    string `json:"protocol"`
	Publicip    string `json:"publicip"`
	Publicipid  string `json:"publicipid"`
	Publicport  string `json:"publicport"`
	State       string `json:"state"`
	Tags        []struct {
		Account      string `json:"account"`
		Customer     string `json:"customer"`
		Domain       string `json:"domain"`
		Domainid     string `json:"domainid"`
		Key          string `json:"key"`
		Project      string `json:"project"`
		Projectid    string `json:"projectid"`
		Resourceid   string `json:"resourceid"`
		Resourcetype string `json:"resourcetype"`
		Value        string `json:"value"`
	} `json:"tags"`
	Zoneid   string `json:"zoneid"`
	Zonename string `json:"zonename"`
}

type ListLoadBalancersParams struct {
	p map[string]interface{}
}

func (p *ListLoadBalancersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["scheme"]; found {
		u.Set("scheme", v.(string))
	}
	if v, found := p.p["sourceipaddress"]; found {
		u.Set("sourceipaddress", v.(string))
	}
	if v, found := p.p["sourceipaddressnetworkid"]; found {
		u.Set("sourceipaddressnetworkid", v.(string))
	}
	if v, found := p.p["tags"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), vv)
			i++
		}
	}
	return u
}

func (p *ListLoadBalancersParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListLoadBalancersParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListLoadBalancersParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
	return
}

func (p *ListLoadBalancersParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListLoadBalancersParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
	return
}

func (p *ListLoadBalancersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListLoadBalancersParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
	return
}

func (p *ListLoadBalancersParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *ListLoadBalancersParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
	return
}

func (p *ListLoadBalancersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListLoadBalancersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListLoadBalancersParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *ListLoadBalancersParams) SetScheme(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["scheme"] = v
	return
}

func (p *ListLoadBalancersParams) SetSourceipaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sourceipaddress"] = v
	return
}

func (p *ListLoadBalancersParams) SetSourceipaddressnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sourceipaddressnetworkid"] = v
	return
}

func (p *ListLoadBalancersParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
	return
}

// You should always use this function to get a new ListLoadBalancersParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewListLoadBalancersParams() *ListLoadBalancersParams {
	p := &ListLoadBalancersParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetLoadBalancerID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListLoadBalancersParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListLoadBalancers(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.LoadBalancers[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.LoadBalancers {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetLoadBalancerByName(name string, opts ...OptionFunc) (*LoadBalancer, int, error) {
	id, count, err := s.GetLoadBalancerID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetLoadBalancerByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetLoadBalancerByID(id string, opts ...OptionFunc) (*LoadBalancer, int, error) {
	p := &ListLoadBalancersParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListLoadBalancers(p)
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf(
			"Invalid parameter id value=%s due to incorrect long value format, "+
				"or entity does not exist", id)) {
			return nil, 0, fmt.Errorf("No match found for %s: %+v", id, l)
		}
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.LoadBalancers[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for LoadBalancer UUID: %s!", id)
}

// Lists load balancers
func (s *LoadBalancerService) ListLoadBalancers(p *ListLoadBalancersParams) (*ListLoadBalancersResponse, error) {
	resp, err := s.cs.newRequest("listLoadBalancers", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListLoadBalancersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListLoadBalancersResponse struct {
	Count         int             `json:"count"`
	LoadBalancers []*LoadBalancer `json:"loadbalancer"`
}

type LoadBalancer struct {
	Account              string `json:"account"`
	Algorithm            string `json:"algorithm"`
	Description          string `json:"description"`
	Domain               string `json:"domain"`
	Domainid             string `json:"domainid"`
	Fordisplay           bool   `json:"fordisplay"`
	Id                   string `json:"id"`
	Jobid                string `json:"jobid"`
	Jobstatus            int    `json:"jobstatus"`
	Loadbalancerinstance []struct {
		Id        string `json:"id"`
		Ipaddress string `json:"ipaddress"`
		Name      string `json:"name"`
		State     string `json:"state"`
	} `json:"loadbalancerinstance"`
	Loadbalancerrule []struct {
		Instanceport int    `json:"instanceport"`
		Sourceport   int    `json:"sourceport"`
		State        string `json:"state"`
	} `json:"loadbalancerrule"`
	Name                     string `json:"name"`
	Networkid                string `json:"networkid"`
	Project                  string `json:"project"`
	Projectid                string `json:"projectid"`
	Sourceipaddress          string `json:"sourceipaddress"`
	Sourceipaddressnetworkid string `json:"sourceipaddressnetworkid"`
	Tags                     []struct {
		Account      string `json:"account"`
		Customer     string `json:"customer"`
		Domain       string `json:"domain"`
		Domainid     string `json:"domainid"`
		Key          string `json:"key"`
		Project      string `json:"project"`
		Projectid    string `json:"projectid"`
		Resourceid   string `json:"resourceid"`
		Resourcetype string `json:"resourcetype"`
		Value        string `json:"value"`
	} `json:"tags"`
}

type ListNetscalerLoadBalancersParams struct {
	p map[string]interface{}
}

func (p *ListNetscalerLoadBalancersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["lbdeviceid"]; found {
		u.Set("lbdeviceid", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	return u
}

func (p *ListNetscalerLoadBalancersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListNetscalerLoadBalancersParams) SetLbdeviceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbdeviceid"] = v
	return
}

func (p *ListNetscalerLoadBalancersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListNetscalerLoadBalancersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListNetscalerLoadBalancersParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
	return
}

// You should always use this function to get a new ListNetscalerLoadBalancersParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewListNetscalerLoadBalancersParams() *ListNetscalerLoadBalancersParams {
	p := &ListNetscalerLoadBalancersParams{}
	p.p = make(map[string]interface{})
	return p
}

// lists netscaler load balancer devices
func (s *LoadBalancerService) ListNetscalerLoadBalancers(p *ListNetscalerLoadBalancersParams) (*ListNetscalerLoadBalancersResponse, error) {
	resp, err := s.cs.newRequest("listNetscalerLoadBalancers", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListNetscalerLoadBalancersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListNetscalerLoadBalancersResponse struct {
	Count                  int                      `json:"count"`
	NetscalerLoadBalancers []*NetscalerLoadBalancer `json:"netscalerloadbalancer"`
}

type NetscalerLoadBalancer struct {
	Gslbprovider            bool     `json:"gslbprovider"`
	Gslbproviderprivateip   string   `json:"gslbproviderprivateip"`
	Gslbproviderpublicip    string   `json:"gslbproviderpublicip"`
	Ipaddress               string   `json:"ipaddress"`
	Isexclusivegslbprovider bool     `json:"isexclusivegslbprovider"`
	Jobid                   string   `json:"jobid"`
	Jobstatus               int      `json:"jobstatus"`
	Lbdevicecapacity        int64    `json:"lbdevicecapacity"`
	Lbdevicededicated       bool     `json:"lbdevicededicated"`
	Lbdeviceid              string   `json:"lbdeviceid"`
	Lbdevicename            string   `json:"lbdevicename"`
	Lbdevicestate           string   `json:"lbdevicestate"`
	Physicalnetworkid       string   `json:"physicalnetworkid"`
	Podids                  []string `json:"podids"`
	Privateinterface        string   `json:"privateinterface"`
	Provider                string   `json:"provider"`
	Publicinterface         string   `json:"publicinterface"`
}

type ListSslCertsParams struct {
	p map[string]interface{}
}

func (p *ListSslCertsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["accountid"]; found {
		u.Set("accountid", v.(string))
	}
	if v, found := p.p["certid"]; found {
		u.Set("certid", v.(string))
	}
	if v, found := p.p["lbruleid"]; found {
		u.Set("lbruleid", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	return u
}

func (p *ListSslCertsParams) SetAccountid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accountid"] = v
	return
}

func (p *ListSslCertsParams) SetCertid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["certid"] = v
	return
}

func (p *ListSslCertsParams) SetLbruleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbruleid"] = v
	return
}

func (p *ListSslCertsParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

// You should always use this function to get a new ListSslCertsParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewListSslCertsParams() *ListSslCertsParams {
	p := &ListSslCertsParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists SSL certificates
func (s *LoadBalancerService) ListSslCerts(p *ListSslCertsParams) (*ListSslCertsResponse, error) {
	resp, err := s.cs.newRequest("listSslCerts", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListSslCertsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListSslCertsResponse struct {
	Count    int        `json:"count"`
	SslCerts []*SslCert `json:"sslcert"`
}

type SslCert struct {
	Account              string   `json:"account"`
	Certchain            string   `json:"certchain"`
	Certificate          string   `json:"certificate"`
	Domain               string   `json:"domain"`
	Domainid             string   `json:"domainid"`
	Fingerprint          string   `json:"fingerprint"`
	Id                   string   `json:"id"`
	Jobid                string   `json:"jobid"`
	Jobstatus            int      `json:"jobstatus"`
	Loadbalancerrulelist []string `json:"loadbalancerrulelist"`
	Name                 string   `json:"name"`
	Project              string   `json:"project"`
	Projectid            string   `json:"projectid"`
}

type RemoveCertFromLoadBalancerParams struct {
	p map[string]interface{}
}

func (p *RemoveCertFromLoadBalancerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["lbruleid"]; found {
		u.Set("lbruleid", v.(string))
	}
	return u
}

func (p *RemoveCertFromLoadBalancerParams) SetLbruleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbruleid"] = v
	return
}

// You should always use this function to get a new RemoveCertFromLoadBalancerParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewRemoveCertFromLoadBalancerParams(lbruleid string) *RemoveCertFromLoadBalancerParams {
	p := &RemoveCertFromLoadBalancerParams{}
	p.p = make(map[string]interface{})
	p.p["lbruleid"] = lbruleid
	return p
}

// Removes a certificate from a load balancer rule
func (s *LoadBalancerService) RemoveCertFromLoadBalancer(p *RemoveCertFromLoadBalancerParams) (*RemoveCertFromLoadBalancerResponse, error) {
	resp, err := s.cs.newRequest("removeCertFromLoadBalancer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RemoveCertFromLoadBalancerResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.Jobid, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type RemoveCertFromLoadBalancerResponse struct {
	Displaytext string `json:"displaytext"`
	Jobid       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type RemoveFromGlobalLoadBalancerRuleParams struct {
	p map[string]interface{}
}

func (p *RemoveFromGlobalLoadBalancerRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["loadbalancerrulelist"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("loadbalancerrulelist", vv)
	}
	return u
}

func (p *RemoveFromGlobalLoadBalancerRuleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *RemoveFromGlobalLoadBalancerRuleParams) SetLoadbalancerrulelist(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["loadbalancerrulelist"] = v
	return
}

// You should always use this function to get a new RemoveFromGlobalLoadBalancerRuleParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewRemoveFromGlobalLoadBalancerRuleParams(id string, loadbalancerrulelist []string) *RemoveFromGlobalLoadBalancerRuleParams {
	p := &RemoveFromGlobalLoadBalancerRuleParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["loadbalancerrulelist"] = loadbalancerrulelist
	return p
}

// Removes a load balancer rule association with global load balancer rule
func (s *LoadBalancerService) RemoveFromGlobalLoadBalancerRule(p *RemoveFromGlobalLoadBalancerRuleParams) (*RemoveFromGlobalLoadBalancerRuleResponse, error) {
	resp, err := s.cs.newRequest("removeFromGlobalLoadBalancerRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RemoveFromGlobalLoadBalancerRuleResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.Jobid, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type RemoveFromGlobalLoadBalancerRuleResponse struct {
	Displaytext string `json:"displaytext"`
	Jobid       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type RemoveFromLoadBalancerRuleParams struct {
	p map[string]interface{}
}

func (p *RemoveFromLoadBalancerRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["virtualmachineids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("virtualmachineids", vv)
	}
	if v, found := p.p["vmidipmap"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("vmidipmap[%d].key", i), k)
			u.Set(fmt.Sprintf("vmidipmap[%d].value", i), vv)
			i++
		}
	}
	return u
}

func (p *RemoveFromLoadBalancerRuleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *RemoveFromLoadBalancerRuleParams) SetVirtualmachineids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineids"] = v
	return
}

func (p *RemoveFromLoadBalancerRuleParams) SetVmidipmap(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vmidipmap"] = v
	return
}

// You should always use this function to get a new RemoveFromLoadBalancerRuleParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewRemoveFromLoadBalancerRuleParams(id string) *RemoveFromLoadBalancerRuleParams {
	p := &RemoveFromLoadBalancerRuleParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Removes a virtual machine or a list of virtual machines from a load balancer rule.
func (s *LoadBalancerService) RemoveFromLoadBalancerRule(p *RemoveFromLoadBalancerRuleParams) (*RemoveFromLoadBalancerRuleResponse, error) {
	resp, err := s.cs.newRequest("removeFromLoadBalancerRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RemoveFromLoadBalancerRuleResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.Jobid, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type RemoveFromLoadBalancerRuleResponse struct {
	Displaytext string `json:"displaytext"`
	Jobid       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type UpdateGlobalLoadBalancerRuleParams struct {
	p map[string]interface{}
}

func (p *UpdateGlobalLoadBalancerRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["gslblbmethod"]; found {
		u.Set("gslblbmethod", v.(string))
	}
	if v, found := p.p["gslbstickysessionmethodname"]; found {
		u.Set("gslbstickysessionmethodname", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *UpdateGlobalLoadBalancerRuleParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
	return
}

func (p *UpdateGlobalLoadBalancerRuleParams) SetGslblbmethod(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gslblbmethod"] = v
	return
}

func (p *UpdateGlobalLoadBalancerRuleParams) SetGslbstickysessionmethodname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gslbstickysessionmethodname"] = v
	return
}

func (p *UpdateGlobalLoadBalancerRuleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new UpdateGlobalLoadBalancerRuleParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewUpdateGlobalLoadBalancerRuleParams(id string) *UpdateGlobalLoadBalancerRuleParams {
	p := &UpdateGlobalLoadBalancerRuleParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// update global load balancer rules.
func (s *LoadBalancerService) UpdateGlobalLoadBalancerRule(p *UpdateGlobalLoadBalancerRuleParams) (*UpdateGlobalLoadBalancerRuleResponse, error) {
	resp, err := s.cs.newRequest("updateGlobalLoadBalancerRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateGlobalLoadBalancerRuleResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.Jobid, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type UpdateGlobalLoadBalancerRuleResponse struct {
	Account                     string `json:"account"`
	Description                 string `json:"description"`
	Domain                      string `json:"domain"`
	Domainid                    string `json:"domainid"`
	Gslbdomainname              string `json:"gslbdomainname"`
	Gslblbmethod                string `json:"gslblbmethod"`
	Gslbservicetype             string `json:"gslbservicetype"`
	Gslbstickysessionmethodname string `json:"gslbstickysessionmethodname"`
	Id                          string `json:"id"`
	Jobid                       string `json:"jobid"`
	Jobstatus                   int    `json:"jobstatus"`
	Loadbalancerrule            []struct {
		Account     string `json:"account"`
		Algorithm   string `json:"algorithm"`
		Cidrlist    string `json:"cidrlist"`
		Description string `json:"description"`
		Domain      string `json:"domain"`
		Domainid    string `json:"domainid"`
		Fordisplay  bool   `json:"fordisplay"`
		Id          string `json:"id"`
		Name        string `json:"name"`
		Networkid   string `json:"networkid"`
		Privateport string `json:"privateport"`
		Project     string `json:"project"`
		Projectid   string `json:"projectid"`
		Protocol    string `json:"protocol"`
		Publicip    string `json:"publicip"`
		Publicipid  string `json:"publicipid"`
		Publicport  string `json:"publicport"`
		State       string `json:"state"`
		Tags        []struct {
			Account      string `json:"account"`
			Customer     string `json:"customer"`
			Domain       string `json:"domain"`
			Domainid     string `json:"domainid"`
			Key          string `json:"key"`
			Project      string `json:"project"`
			Projectid    string `json:"projectid"`
			Resourceid   string `json:"resourceid"`
			Resourcetype string `json:"resourcetype"`
			Value        string `json:"value"`
		} `json:"tags"`
		Zoneid   string `json:"zoneid"`
		Zonename string `json:"zonename"`
	} `json:"loadbalancerrule"`
	Name      string `json:"name"`
	Project   string `json:"project"`
	Projectid string `json:"projectid"`
	Regionid  int    `json:"regionid"`
}

type UpdateLBHealthCheckPolicyParams struct {
	p map[string]interface{}
}

func (p *UpdateLBHealthCheckPolicyParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["customid"]; found {
		u.Set("customid", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *UpdateLBHealthCheckPolicyParams) SetCustomid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customid"] = v
	return
}

func (p *UpdateLBHealthCheckPolicyParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
	return
}

func (p *UpdateLBHealthCheckPolicyParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new UpdateLBHealthCheckPolicyParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewUpdateLBHealthCheckPolicyParams(id string) *UpdateLBHealthCheckPolicyParams {
	p := &UpdateLBHealthCheckPolicyParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates load balancer health check policy
func (s *LoadBalancerService) UpdateLBHealthCheckPolicy(p *UpdateLBHealthCheckPolicyParams) (*UpdateLBHealthCheckPolicyResponse, error) {
	resp, err := s.cs.newRequest("updateLBHealthCheckPolicy", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateLBHealthCheckPolicyResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.Jobid, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type UpdateLBHealthCheckPolicyResponse struct {
	Account           string `json:"account"`
	Domain            string `json:"domain"`
	Domainid          string `json:"domainid"`
	Healthcheckpolicy []struct {
		Description             string `json:"description"`
		Fordisplay              bool   `json:"fordisplay"`
		Healthcheckinterval     int    `json:"healthcheckinterval"`
		Healthcheckthresshold   int    `json:"healthcheckthresshold"`
		Id                      string `json:"id"`
		Pingpath                string `json:"pingpath"`
		Responsetime            int    `json:"responsetime"`
		State                   string `json:"state"`
		Unhealthcheckthresshold int    `json:"unhealthcheckthresshold"`
	} `json:"healthcheckpolicy"`
	Jobid     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Lbruleid  string `json:"lbruleid"`
	Zoneid    string `json:"zoneid"`
}

type UpdateLBStickinessPolicyParams struct {
	p map[string]interface{}
}

func (p *UpdateLBStickinessPolicyParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["customid"]; found {
		u.Set("customid", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *UpdateLBStickinessPolicyParams) SetCustomid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customid"] = v
	return
}

func (p *UpdateLBStickinessPolicyParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
	return
}

func (p *UpdateLBStickinessPolicyParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new UpdateLBStickinessPolicyParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewUpdateLBStickinessPolicyParams(id string) *UpdateLBStickinessPolicyParams {
	p := &UpdateLBStickinessPolicyParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates load balancer stickiness policy
func (s *LoadBalancerService) UpdateLBStickinessPolicy(p *UpdateLBStickinessPolicyParams) (*UpdateLBStickinessPolicyResponse, error) {
	resp, err := s.cs.newRequest("updateLBStickinessPolicy", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateLBStickinessPolicyResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.Jobid, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type UpdateLBStickinessPolicyResponse struct {
	Account          string `json:"account"`
	Description      string `json:"description"`
	Domain           string `json:"domain"`
	Domainid         string `json:"domainid"`
	Jobid            string `json:"jobid"`
	Jobstatus        int    `json:"jobstatus"`
	Lbruleid         string `json:"lbruleid"`
	Name             string `json:"name"`
	State            string `json:"state"`
	Stickinesspolicy []struct {
		Description string            `json:"description"`
		Fordisplay  bool              `json:"fordisplay"`
		Id          string            `json:"id"`
		Methodname  string            `json:"methodname"`
		Name        string            `json:"name"`
		Params      map[string]string `json:"params"`
		State       string            `json:"state"`
	} `json:"stickinesspolicy"`
	Zoneid string `json:"zoneid"`
}

type UpdateLoadBalancerParams struct {
	p map[string]interface{}
}

func (p *UpdateLoadBalancerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["customid"]; found {
		u.Set("customid", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *UpdateLoadBalancerParams) SetCustomid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customid"] = v
	return
}

func (p *UpdateLoadBalancerParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
	return
}

func (p *UpdateLoadBalancerParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new UpdateLoadBalancerParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewUpdateLoadBalancerParams(id string) *UpdateLoadBalancerParams {
	p := &UpdateLoadBalancerParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a load balancer
func (s *LoadBalancerService) UpdateLoadBalancer(p *UpdateLoadBalancerParams) (*UpdateLoadBalancerResponse, error) {
	resp, err := s.cs.newRequest("updateLoadBalancer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateLoadBalancerResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.Jobid, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type UpdateLoadBalancerResponse struct {
	Account              string `json:"account"`
	Algorithm            string `json:"algorithm"`
	Description          string `json:"description"`
	Domain               string `json:"domain"`
	Domainid             string `json:"domainid"`
	Fordisplay           bool   `json:"fordisplay"`
	Id                   string `json:"id"`
	Jobid                string `json:"jobid"`
	Jobstatus            int    `json:"jobstatus"`
	Loadbalancerinstance []struct {
		Id        string `json:"id"`
		Ipaddress string `json:"ipaddress"`
		Name      string `json:"name"`
		State     string `json:"state"`
	} `json:"loadbalancerinstance"`
	Loadbalancerrule []struct {
		Instanceport int    `json:"instanceport"`
		Sourceport   int    `json:"sourceport"`
		State        string `json:"state"`
	} `json:"loadbalancerrule"`
	Name                     string `json:"name"`
	Networkid                string `json:"networkid"`
	Project                  string `json:"project"`
	Projectid                string `json:"projectid"`
	Sourceipaddress          string `json:"sourceipaddress"`
	Sourceipaddressnetworkid string `json:"sourceipaddressnetworkid"`
	Tags                     []struct {
		Account      string `json:"account"`
		Customer     string `json:"customer"`
		Domain       string `json:"domain"`
		Domainid     string `json:"domainid"`
		Key          string `json:"key"`
		Project      string `json:"project"`
		Projectid    string `json:"projectid"`
		Resourceid   string `json:"resourceid"`
		Resourcetype string `json:"resourcetype"`
		Value        string `json:"value"`
	} `json:"tags"`
}

type UpdateLoadBalancerRuleParams struct {
	p map[string]interface{}
}

func (p *UpdateLoadBalancerRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["algorithm"]; found {
		u.Set("algorithm", v.(string))
	}
	if v, found := p.p["customid"]; found {
		u.Set("customid", v.(string))
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	return u
}

func (p *UpdateLoadBalancerRuleParams) SetAlgorithm(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["algorithm"] = v
	return
}

func (p *UpdateLoadBalancerRuleParams) SetCustomid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customid"] = v
	return
}

func (p *UpdateLoadBalancerRuleParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
	return
}

func (p *UpdateLoadBalancerRuleParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
	return
}

func (p *UpdateLoadBalancerRuleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *UpdateLoadBalancerRuleParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *UpdateLoadBalancerRuleParams) SetProtocol(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["protocol"] = v
	return
}

// You should always use this function to get a new UpdateLoadBalancerRuleParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewUpdateLoadBalancerRuleParams(id string) *UpdateLoadBalancerRuleParams {
	p := &UpdateLoadBalancerRuleParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates load balancer
func (s *LoadBalancerService) UpdateLoadBalancerRule(p *UpdateLoadBalancerRuleParams) (*UpdateLoadBalancerRuleResponse, error) {
	resp, err := s.cs.newRequest("updateLoadBalancerRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateLoadBalancerRuleResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.Jobid, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type UpdateLoadBalancerRuleResponse struct {
	Account     string `json:"account"`
	Algorithm   string `json:"algorithm"`
	Cidrlist    string `json:"cidrlist"`
	Description string `json:"description"`
	Domain      string `json:"domain"`
	Domainid    string `json:"domainid"`
	Fordisplay  bool   `json:"fordisplay"`
	Id          string `json:"id"`
	Jobid       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Name        string `json:"name"`
	Networkid   string `json:"networkid"`
	Privateport string `json:"privateport"`
	Project     string `json:"project"`
	Projectid   string `json:"projectid"`
	Protocol    string `json:"protocol"`
	Publicip    string `json:"publicip"`
	Publicipid  string `json:"publicipid"`
	Publicport  string `json:"publicport"`
	State       string `json:"state"`
	Tags        []struct {
		Account      string `json:"account"`
		Customer     string `json:"customer"`
		Domain       string `json:"domain"`
		Domainid     string `json:"domainid"`
		Key          string `json:"key"`
		Project      string `json:"project"`
		Projectid    string `json:"projectid"`
		Resourceid   string `json:"resourceid"`
		Resourcetype string `json:"resourcetype"`
		Value        string `json:"value"`
	} `json:"tags"`
	Zoneid   string `json:"zoneid"`
	Zonename string `json:"zonename"`
}

type UploadSslCertParams struct {
	p map[string]interface{}
}

func (p *UploadSslCertParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["certchain"]; found {
		u.Set("certchain", v.(string))
	}
	if v, found := p.p["certificate"]; found {
		u.Set("certificate", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["privatekey"]; found {
		u.Set("privatekey", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	return u
}

func (p *UploadSslCertParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *UploadSslCertParams) SetCertchain(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["certchain"] = v
	return
}

func (p *UploadSslCertParams) SetCertificate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["certificate"] = v
	return
}

func (p *UploadSslCertParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *UploadSslCertParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *UploadSslCertParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
	return
}

func (p *UploadSslCertParams) SetPrivatekey(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["privatekey"] = v
	return
}

func (p *UploadSslCertParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

// You should always use this function to get a new UploadSslCertParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewUploadSslCertParams(certificate string, name string, privatekey string) *UploadSslCertParams {
	p := &UploadSslCertParams{}
	p.p = make(map[string]interface{})
	p.p["certificate"] = certificate
	p.p["name"] = name
	p.p["privatekey"] = privatekey
	return p
}

// Upload a certificate to CloudStack
func (s *LoadBalancerService) UploadSslCert(p *UploadSslCertParams) (*UploadSslCertResponse, error) {
	resp, err := s.cs.newRequest("uploadSslCert", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UploadSslCertResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UploadSslCertResponse struct {
	Account              string   `json:"account"`
	Certchain            string   `json:"certchain"`
	Certificate          string   `json:"certificate"`
	Domain               string   `json:"domain"`
	Domainid             string   `json:"domainid"`
	Fingerprint          string   `json:"fingerprint"`
	Id                   string   `json:"id"`
	Jobid                string   `json:"jobid"`
	Jobstatus            int      `json:"jobstatus"`
	Loadbalancerrulelist []string `json:"loadbalancerrulelist"`
	Name                 string   `json:"name"`
	Project              string   `json:"project"`
	Projectid            string   `json:"projectid"`
}
