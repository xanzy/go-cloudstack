//
// Copyright 2014, Sander van Harmelen
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

package cloudstack43

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type ConfigureInternalLoadBalancerElementParams struct {
	p map[string]interface{}
}

func (p *ConfigureInternalLoadBalancerElementParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["enabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("enabled", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *ConfigureInternalLoadBalancerElementParams) SetEnabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["enabled"] = v
	return
}

func (p *ConfigureInternalLoadBalancerElementParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new ConfigureInternalLoadBalancerElementParams instance,
// as then you are sure you have configured all required params
func (s *InternalLBService) NewConfigureInternalLoadBalancerElementParams(enabled bool, id string) *ConfigureInternalLoadBalancerElementParams {
	p := &ConfigureInternalLoadBalancerElementParams{}
	p.p = make(map[string]interface{})
	p.p["enabled"] = enabled
	p.p["id"] = id
	return p
}

// Configures an Internal Load Balancer element.
func (s *InternalLBService) ConfigureInternalLoadBalancerElement(p *ConfigureInternalLoadBalancerElementParams) (*ConfigureInternalLoadBalancerElementResponse, error) {
	resp, err := s.cs.newRequest("configureInternalLoadBalancerElement", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ConfigureInternalLoadBalancerElementResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, warn, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			return nil, err
		}
		// If 'warn' has a value it means the job is running longer than the configured
		// timeout, the resonse will contain the jobid of the running async job
		if warn != nil {
			return &r, warn
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

type ConfigureInternalLoadBalancerElementResponse struct {
	JobID   string `json:"jobid,omitempty"`
	Enabled bool   `json:"enabled,omitempty"`
	Id      string `json:"id,omitempty"`
	Nspid   string `json:"nspid,omitempty"`
}

type CreateInternalLoadBalancerElementParams struct {
	p map[string]interface{}
}

func (p *CreateInternalLoadBalancerElementParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["nspid"]; found {
		u.Set("nspid", v.(string))
	}
	return u
}

func (p *CreateInternalLoadBalancerElementParams) SetNspid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nspid"] = v
	return
}

// You should always use this function to get a new CreateInternalLoadBalancerElementParams instance,
// as then you are sure you have configured all required params
func (s *InternalLBService) NewCreateInternalLoadBalancerElementParams(nspid string) *CreateInternalLoadBalancerElementParams {
	p := &CreateInternalLoadBalancerElementParams{}
	p.p = make(map[string]interface{})
	p.p["nspid"] = nspid
	return p
}

// Create an Internal Load Balancer element.
func (s *InternalLBService) CreateInternalLoadBalancerElement(p *CreateInternalLoadBalancerElementParams) (*CreateInternalLoadBalancerElementResponse, error) {
	resp, err := s.cs.newRequest("createInternalLoadBalancerElement", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateInternalLoadBalancerElementResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, warn, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			return nil, err
		}
		// If 'warn' has a value it means the job is running longer than the configured
		// timeout, the resonse will contain the jobid of the running async job
		if warn != nil {
			return &r, warn
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

type CreateInternalLoadBalancerElementResponse struct {
	JobID   string `json:"jobid,omitempty"`
	Enabled bool   `json:"enabled,omitempty"`
	Nspid   string `json:"nspid,omitempty"`
	Id      string `json:"id,omitempty"`
}

type ListInternalLoadBalancerElementsParams struct {
	p map[string]interface{}
}

func (p *ListInternalLoadBalancerElementsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["enabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("enabled", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["nspid"]; found {
		u.Set("nspid", v.(string))
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

func (p *ListInternalLoadBalancerElementsParams) SetEnabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["enabled"] = v
	return
}

func (p *ListInternalLoadBalancerElementsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListInternalLoadBalancerElementsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListInternalLoadBalancerElementsParams) SetNspid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nspid"] = v
	return
}

func (p *ListInternalLoadBalancerElementsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListInternalLoadBalancerElementsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

// You should always use this function to get a new ListInternalLoadBalancerElementsParams instance,
// as then you are sure you have configured all required params
func (s *InternalLBService) NewListInternalLoadBalancerElementsParams() *ListInternalLoadBalancerElementsParams {
	p := &ListInternalLoadBalancerElementsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *InternalLBService) GetInternalLoadBalancerElementByID(id string) (*InternalLoadBalancerElement, int, error) {
	p := &ListInternalLoadBalancerElementsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	l, err := s.ListInternalLoadBalancerElements(p)
	if err != nil {
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.InternalLoadBalancerElements[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for InternalLoadBalancerElement UUID: %s!", id)
}

// Lists all available Internal Load Balancer elements.
func (s *InternalLBService) ListInternalLoadBalancerElements(p *ListInternalLoadBalancerElementsParams) (*ListInternalLoadBalancerElementsResponse, error) {
	resp, err := s.cs.newRequest("listInternalLoadBalancerElements", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListInternalLoadBalancerElementsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListInternalLoadBalancerElementsResponse struct {
	Count                        int                            `json:"count"`
	InternalLoadBalancerElements []*InternalLoadBalancerElement `json:"internalloadbalancerelement"`
}

type InternalLoadBalancerElement struct {
	Id      string `json:"id,omitempty"`
	Enabled bool   `json:"enabled,omitempty"`
	Nspid   string `json:"nspid,omitempty"`
}

type StopInternalLoadBalancerVMParams struct {
	p map[string]interface{}
}

func (p *StopInternalLoadBalancerVMParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["forced"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forced", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *StopInternalLoadBalancerVMParams) SetForced(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forced"] = v
	return
}

func (p *StopInternalLoadBalancerVMParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new StopInternalLoadBalancerVMParams instance,
// as then you are sure you have configured all required params
func (s *InternalLBService) NewStopInternalLoadBalancerVMParams(id string) *StopInternalLoadBalancerVMParams {
	p := &StopInternalLoadBalancerVMParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Stops an Internal LB vm.
func (s *InternalLBService) StopInternalLoadBalancerVM(p *StopInternalLoadBalancerVMParams) (*StopInternalLoadBalancerVMResponse, error) {
	resp, err := s.cs.newRequest("stopInternalLoadBalancerVM", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r StopInternalLoadBalancerVMResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, warn, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			return nil, err
		}
		// If 'warn' has a value it means the job is running longer than the configured
		// timeout, the resonse will contain the jobid of the running async job
		if warn != nil {
			return &r, warn
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

type StopInternalLoadBalancerVMResponse struct {
	JobID               string `json:"jobid,omitempty"`
	Gateway             string `json:"gateway,omitempty"`
	Dns2                string `json:"dns2,omitempty"`
	Vpcid               string `json:"vpcid,omitempty"`
	Templateid          string `json:"templateid,omitempty"`
	Guestipaddress      string `json:"guestipaddress,omitempty"`
	Isredundantrouter   bool   `json:"isredundantrouter,omitempty"`
	Requiresupgrade     bool   `json:"requiresupgrade,omitempty"`
	Publicmacaddress    string `json:"publicmacaddress,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Projectid           string `json:"projectid,omitempty"`
	Account             string `json:"account,omitempty"`
	Domainid            string `json:"domainid,omitempty"`
	Podid               string `json:"podid,omitempty"`
	Linklocalmacaddress string `json:"linklocalmacaddress,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
	Publicnetworkid     string `json:"publicnetworkid,omitempty"`
	Role                string `json:"role,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
	Linklocalip         string `json:"linklocalip,omitempty"`
	Version             string `json:"version,omitempty"`
	Linklocalnetmask    string `json:"linklocalnetmask,omitempty"`
	Ip6dns2             string `json:"ip6dns2,omitempty"`
	Publicip            string `json:"publicip,omitempty"`
	Hostname            string `json:"hostname,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Project             string `json:"project,omitempty"`
	Scriptsversion      string `json:"scriptsversion,omitempty"`
	Hostid              string `json:"hostid,omitempty"`
	Name                string `json:"name,omitempty"`
	Redundantstate      string `json:"redundantstate,omitempty"`
	Linklocalnetworkid  string `json:"linklocalnetworkid,omitempty"`
	Ip6dns1             string `json:"ip6dns1,omitempty"`
	Guestnetworkid      string `json:"guestnetworkid,omitempty"`
	Created             string `json:"created,omitempty"`
	Zoneid              string `json:"zoneid,omitempty"`
	Guestnetmask        string `json:"guestnetmask,omitempty"`
	State               string `json:"state,omitempty"`
	Networkdomain       string `json:"networkdomain,omitempty"`
	Nic                 []struct {
		Type         string   `json:"type,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Id           string   `json:"id,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
	} `json:"nic,omitempty"`
	Dns1            string `json:"dns1,omitempty"`
	Guestmacaddress string `json:"guestmacaddress,omitempty"`
	Id              string `json:"id,omitempty"`
	Publicnetmask   string `json:"publicnetmask,omitempty"`
}

type StartInternalLoadBalancerVMParams struct {
	p map[string]interface{}
}

func (p *StartInternalLoadBalancerVMParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *StartInternalLoadBalancerVMParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new StartInternalLoadBalancerVMParams instance,
// as then you are sure you have configured all required params
func (s *InternalLBService) NewStartInternalLoadBalancerVMParams(id string) *StartInternalLoadBalancerVMParams {
	p := &StartInternalLoadBalancerVMParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Starts an existing internal lb vm.
func (s *InternalLBService) StartInternalLoadBalancerVM(p *StartInternalLoadBalancerVMParams) (*StartInternalLoadBalancerVMResponse, error) {
	resp, err := s.cs.newRequest("startInternalLoadBalancerVM", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r StartInternalLoadBalancerVMResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, warn, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			return nil, err
		}
		// If 'warn' has a value it means the job is running longer than the configured
		// timeout, the resonse will contain the jobid of the running async job
		if warn != nil {
			return &r, warn
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

type StartInternalLoadBalancerVMResponse struct {
	JobID             string `json:"jobid,omitempty"`
	Name              string `json:"name,omitempty"`
	Linklocalnetmask  string `json:"linklocalnetmask,omitempty"`
	Publicip          string `json:"publicip,omitempty"`
	Scriptsversion    string `json:"scriptsversion,omitempty"`
	Ip6dns2           string `json:"ip6dns2,omitempty"`
	Zoneid            string `json:"zoneid,omitempty"`
	Created           string `json:"created,omitempty"`
	Role              string `json:"role,omitempty"`
	Isredundantrouter bool   `json:"isredundantrouter,omitempty"`
	Guestnetmask      string `json:"guestnetmask,omitempty"`
	Templateid        string `json:"templateid,omitempty"`
	Guestnetworkid    string `json:"guestnetworkid,omitempty"`
	Nic               []struct {
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Id           string   `json:"id,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Type         string   `json:"type,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
	} `json:"nic,omitempty"`
	Publicmacaddress    string `json:"publicmacaddress,omitempty"`
	Redundantstate      string `json:"redundantstate,omitempty"`
	Account             string `json:"account,omitempty"`
	Hostid              string `json:"hostid,omitempty"`
	Dns1                string `json:"dns1,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	State               string `json:"state,omitempty"`
	Guestipaddress      string `json:"guestipaddress,omitempty"`
	Version             string `json:"version,omitempty"`
	Podid               string `json:"podid,omitempty"`
	Projectid           string `json:"projectid,omitempty"`
	Linklocalnetworkid  string `json:"linklocalnetworkid,omitempty"`
	Guestmacaddress     string `json:"guestmacaddress,omitempty"`
	Dns2                string `json:"dns2,omitempty"`
	Linklocalip         string `json:"linklocalip,omitempty"`
	Id                  string `json:"id,omitempty"`
	Ip6dns1             string `json:"ip6dns1,omitempty"`
	Requiresupgrade     bool   `json:"requiresupgrade,omitempty"`
	Domainid            string `json:"domainid,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
	Networkdomain       string `json:"networkdomain,omitempty"`
	Publicnetmask       string `json:"publicnetmask,omitempty"`
	Gateway             string `json:"gateway,omitempty"`
	Project             string `json:"project,omitempty"`
	Hostname            string `json:"hostname,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Linklocalmacaddress string `json:"linklocalmacaddress,omitempty"`
	Vpcid               string `json:"vpcid,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
	Publicnetworkid     string `json:"publicnetworkid,omitempty"`
}

type ListInternalLoadBalancerVMsParams struct {
	p map[string]interface{}
}

func (p *ListInternalLoadBalancerVMsParams) toURLValues() url.Values {
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
	if v, found := p.p["forvpc"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forvpc", vv)
	}
	if v, found := p.p["hostid"]; found {
		u.Set("hostid", v.(string))
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
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := p.p["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListInternalLoadBalancerVMsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListInternalLoadBalancerVMsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListInternalLoadBalancerVMsParams) SetForvpc(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forvpc"] = v
	return
}

func (p *ListInternalLoadBalancerVMsParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
	return
}

func (p *ListInternalLoadBalancerVMsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListInternalLoadBalancerVMsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
	return
}

func (p *ListInternalLoadBalancerVMsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListInternalLoadBalancerVMsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
	return
}

func (p *ListInternalLoadBalancerVMsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *ListInternalLoadBalancerVMsParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
	return
}

func (p *ListInternalLoadBalancerVMsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListInternalLoadBalancerVMsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListInternalLoadBalancerVMsParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
	return
}

func (p *ListInternalLoadBalancerVMsParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *ListInternalLoadBalancerVMsParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
	return
}

func (p *ListInternalLoadBalancerVMsParams) SetVpcid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcid"] = v
	return
}

func (p *ListInternalLoadBalancerVMsParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new ListInternalLoadBalancerVMsParams instance,
// as then you are sure you have configured all required params
func (s *InternalLBService) NewListInternalLoadBalancerVMsParams() *ListInternalLoadBalancerVMsParams {
	p := &ListInternalLoadBalancerVMsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *InternalLBService) GetInternalLoadBalancerVMID(name string) (string, error) {
	p := &ListInternalLoadBalancerVMsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	l, err := s.ListInternalLoadBalancerVMs(p)
	if err != nil {
		return "", err
	}

	if l.Count == 0 {
		return "", fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.InternalLoadBalancerVMs[0].Id, nil
	}

	if l.Count > 1 {
		for _, v := range l.InternalLoadBalancerVMs {
			if v.Name == name {
				return v.Id, nil
			}
		}
	}
	return "", fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *InternalLBService) GetInternalLoadBalancerVMByName(name string) (*InternalLoadBalancerVM, int, error) {
	id, err := s.GetInternalLoadBalancerVMID(name)
	if err != nil {
		return nil, -1, err
	}

	r, count, err := s.GetInternalLoadBalancerVMByID(id)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *InternalLBService) GetInternalLoadBalancerVMByID(id string) (*InternalLoadBalancerVM, int, error) {
	p := &ListInternalLoadBalancerVMsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	l, err := s.ListInternalLoadBalancerVMs(p)
	if err != nil {
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.InternalLoadBalancerVMs[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for InternalLoadBalancerVM UUID: %s!", id)
}

// List internal LB VMs.
func (s *InternalLBService) ListInternalLoadBalancerVMs(p *ListInternalLoadBalancerVMsParams) (*ListInternalLoadBalancerVMsResponse, error) {
	resp, err := s.cs.newRequest("listInternalLoadBalancerVMs", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListInternalLoadBalancerVMsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListInternalLoadBalancerVMsResponse struct {
	Count                   int                       `json:"count"`
	InternalLoadBalancerVMs []*InternalLoadBalancerVM `json:"internalloadbalancervm"`
}

type InternalLoadBalancerVM struct {
	Version             string `json:"version,omitempty"`
	Redundantstate      string `json:"redundantstate,omitempty"`
	Publicnetmask       string `json:"publicnetmask,omitempty"`
	Guestipaddress      string `json:"guestipaddress,omitempty"`
	Vpcid               string `json:"vpcid,omitempty"`
	Linklocalmacaddress string `json:"linklocalmacaddress,omitempty"`
	Dns2                string `json:"dns2,omitempty"`
	Isredundantrouter   bool   `json:"isredundantrouter,omitempty"`
	Role                string `json:"role,omitempty"`
	Guestnetworkid      string `json:"guestnetworkid,omitempty"`
	Ip6dns2             string `json:"ip6dns2,omitempty"`
	Requiresupgrade     bool   `json:"requiresupgrade,omitempty"`
	Linklocalip         string `json:"linklocalip,omitempty"`
	Guestnetmask        string `json:"guestnetmask,omitempty"`
	Linklocalnetmask    string `json:"linklocalnetmask,omitempty"`
	Projectid           string `json:"projectid,omitempty"`
	Gateway             string `json:"gateway,omitempty"`
	Linklocalnetworkid  string `json:"linklocalnetworkid,omitempty"`
	Podid               string `json:"podid,omitempty"`
	State               string `json:"state,omitempty"`
	Publicnetworkid     string `json:"publicnetworkid,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
	Hostid              string `json:"hostid,omitempty"`
	Hostname            string `json:"hostname,omitempty"`
	Scriptsversion      string `json:"scriptsversion,omitempty"`
	Account             string `json:"account,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Project             string `json:"project,omitempty"`
	Nic                 []struct {
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Type         string   `json:"type,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Id           string   `json:"id,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
	} `json:"nic,omitempty"`
	Name             string `json:"name,omitempty"`
	Zonename         string `json:"zonename,omitempty"`
	Templateid       string `json:"templateid,omitempty"`
	Zoneid           string `json:"zoneid,omitempty"`
	Created          string `json:"created,omitempty"`
	Publicip         string `json:"publicip,omitempty"`
	Publicmacaddress string `json:"publicmacaddress,omitempty"`
	Dns1             string `json:"dns1,omitempty"`
	Id               string `json:"id,omitempty"`
	Ip6dns1          string `json:"ip6dns1,omitempty"`
	Guestmacaddress  string `json:"guestmacaddress,omitempty"`
	Networkdomain    string `json:"networkdomain,omitempty"`
	Domainid         string `json:"domainid,omitempty"`
}
