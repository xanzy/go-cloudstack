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

		var r ConfigureInternalLoadBalancerElementResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type ConfigureInternalLoadBalancerElementResponse struct {
	JobID   string `json:"jobid,omitempty"`
	Nspid   string `json:"nspid,omitempty"`
	Id      string `json:"id,omitempty"`
	Enabled bool   `json:"enabled,omitempty"`
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

		var r CreateInternalLoadBalancerElementResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type CreateInternalLoadBalancerElementResponse struct {
	JobID   string `json:"jobid,omitempty"`
	Nspid   string `json:"nspid,omitempty"`
	Enabled bool   `json:"enabled,omitempty"`
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
func (s *InternalLBService) GetInternalLoadBalancerElementID(keyword string) (string, error) {
	p := &ListInternalLoadBalancerElementsParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword

	l, err := s.ListInternalLoadBalancerElements(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, keyword, l)
	}
	return l.InternalLoadBalancerElements[0].Id, nil
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
	Nspid   string `json:"nspid,omitempty"`
	Enabled bool   `json:"enabled,omitempty"`
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

		var r StopInternalLoadBalancerVMResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type StopInternalLoadBalancerVMResponse struct {
	JobID               string `json:"jobid,omitempty"`
	Ip6dns1             string `json:"ip6dns1,omitempty"`
	Podid               string `json:"podid,omitempty"`
	Account             string `json:"account,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
	Hostname            string `json:"hostname,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Hostid              string `json:"hostid,omitempty"`
	Name                string `json:"name,omitempty"`
	Scriptsversion      string `json:"scriptsversion,omitempty"`
	Dns1                string `json:"dns1,omitempty"`
	Zoneid              string `json:"zoneid,omitempty"`
	Linklocalnetmask    string `json:"linklocalnetmask,omitempty"`
	Isredundantrouter   bool   `json:"isredundantrouter,omitempty"`
	Publicip            string `json:"publicip,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Publicnetmask       string `json:"publicnetmask,omitempty"`
	Role                string `json:"role,omitempty"`
	Publicmacaddress    string `json:"publicmacaddress,omitempty"`
	Guestipaddress      string `json:"guestipaddress,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
	Linklocalip         string `json:"linklocalip,omitempty"`
	Dns2                string `json:"dns2,omitempty"`
	Vpcid               string `json:"vpcid,omitempty"`
	Ip6dns2             string `json:"ip6dns2,omitempty"`
	Linklocalmacaddress string `json:"linklocalmacaddress,omitempty"`
	Publicnetworkid     string `json:"publicnetworkid,omitempty"`
	Gateway             string `json:"gateway,omitempty"`
	Guestnetmask        string `json:"guestnetmask,omitempty"`
	Domainid            string `json:"domainid,omitempty"`
	Linklocalnetworkid  string `json:"linklocalnetworkid,omitempty"`
	Created             string `json:"created,omitempty"`
	Networkdomain       string `json:"networkdomain,omitempty"`
	Id                  string `json:"id,omitempty"`
	Guestmacaddress     string `json:"guestmacaddress,omitempty"`
	Version             string `json:"version,omitempty"`
	Guestnetworkid      string `json:"guestnetworkid,omitempty"`
	Project             string `json:"project,omitempty"`
	Projectid           string `json:"projectid,omitempty"`
	Templateid          string `json:"templateid,omitempty"`
	Redundantstate      string `json:"redundantstate,omitempty"`
	Requiresupgrade     bool   `json:"requiresupgrade,omitempty"`
	Nic                 []struct {
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Type         string   `json:"type,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Id           string   `json:"id,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
	} `json:"nic,omitempty"`
	State string `json:"state,omitempty"`
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

		var r StartInternalLoadBalancerVMResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type StartInternalLoadBalancerVMResponse struct {
	JobID               string `json:"jobid,omitempty"`
	Vpcid               string `json:"vpcid,omitempty"`
	Hostid              string `json:"hostid,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Project             string `json:"project,omitempty"`
	Publicmacaddress    string `json:"publicmacaddress,omitempty"`
	Guestnetmask        string `json:"guestnetmask,omitempty"`
	Redundantstate      string `json:"redundantstate,omitempty"`
	Gateway             string `json:"gateway,omitempty"`
	Linklocalnetworkid  string `json:"linklocalnetworkid,omitempty"`
	Projectid           string `json:"projectid,omitempty"`
	Guestmacaddress     string `json:"guestmacaddress,omitempty"`
	Ip6dns1             string `json:"ip6dns1,omitempty"`
	Publicip            string `json:"publicip,omitempty"`
	State               string `json:"state,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
	Linklocalmacaddress string `json:"linklocalmacaddress,omitempty"`
	Role                string `json:"role,omitempty"`
	Guestnetworkid      string `json:"guestnetworkid,omitempty"`
	Dns2                string `json:"dns2,omitempty"`
	Created             string `json:"created,omitempty"`
	Name                string `json:"name,omitempty"`
	Nic                 []struct {
		Networkid    string   `json:"networkid,omitempty"`
		Type         string   `json:"type,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Id           string   `json:"id,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
	} `json:"nic,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Publicnetworkid     string `json:"publicnetworkid,omitempty"`
	Templateid          string `json:"templateid,omitempty"`
	Publicnetmask       string `json:"publicnetmask,omitempty"`
	Hostname            string `json:"hostname,omitempty"`
	Zoneid              string `json:"zoneid,omitempty"`
	Isredundantrouter   bool   `json:"isredundantrouter,omitempty"`
	Version             string `json:"version,omitempty"`
	Requiresupgrade     bool   `json:"requiresupgrade,omitempty"`
	Account             string `json:"account,omitempty"`
	Dns1                string `json:"dns1,omitempty"`
	Linklocalnetmask    string `json:"linklocalnetmask,omitempty"`
	Scriptsversion      string `json:"scriptsversion,omitempty"`
	Networkdomain       string `json:"networkdomain,omitempty"`
	Guestipaddress      string `json:"guestipaddress,omitempty"`
	Id                  string `json:"id,omitempty"`
	Ip6dns2             string `json:"ip6dns2,omitempty"`
	Podid               string `json:"podid,omitempty"`
	Domainid            string `json:"domainid,omitempty"`
	Linklocalip         string `json:"linklocalip,omitempty"`
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
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, name, l)
	}
	return l.InternalLoadBalancerVMs[0].Id, nil
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
	Projectid          string `json:"projectid,omitempty"`
	Publicip           string `json:"publicip,omitempty"`
	Created            string `json:"created,omitempty"`
	Linklocalnetworkid string `json:"linklocalnetworkid,omitempty"`
	Hostname           string `json:"hostname,omitempty"`
	Nic                []struct {
		Gateway      string   `json:"gateway,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Id           string   `json:"id,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Type         string   `json:"type,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
	} `json:"nic,omitempty"`
	Role                string `json:"role,omitempty"`
	Podid               string `json:"podid,omitempty"`
	Linklocalip         string `json:"linklocalip,omitempty"`
	Isredundantrouter   bool   `json:"isredundantrouter,omitempty"`
	Guestipaddress      string `json:"guestipaddress,omitempty"`
	Domainid            string `json:"domainid,omitempty"`
	Publicmacaddress    string `json:"publicmacaddress,omitempty"`
	Gateway             string `json:"gateway,omitempty"`
	Dns2                string `json:"dns2,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Name                string `json:"name,omitempty"`
	Account             string `json:"account,omitempty"`
	Zoneid              string `json:"zoneid,omitempty"`
	Id                  string `json:"id,omitempty"`
	Ip6dns1             string `json:"ip6dns1,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
	Publicnetmask       string `json:"publicnetmask,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
	Publicnetworkid     string `json:"publicnetworkid,omitempty"`
	Scriptsversion      string `json:"scriptsversion,omitempty"`
	Guestnetmask        string `json:"guestnetmask,omitempty"`
	Ip6dns2             string `json:"ip6dns2,omitempty"`
	State               string `json:"state,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Hostid              string `json:"hostid,omitempty"`
	Project             string `json:"project,omitempty"`
	Guestnetworkid      string `json:"guestnetworkid,omitempty"`
	Networkdomain       string `json:"networkdomain,omitempty"`
	Vpcid               string `json:"vpcid,omitempty"`
	Linklocalmacaddress string `json:"linklocalmacaddress,omitempty"`
	Guestmacaddress     string `json:"guestmacaddress,omitempty"`
	Redundantstate      string `json:"redundantstate,omitempty"`
	Dns1                string `json:"dns1,omitempty"`
	Linklocalnetmask    string `json:"linklocalnetmask,omitempty"`
	Requiresupgrade     bool   `json:"requiresupgrade,omitempty"`
	Templateid          string `json:"templateid,omitempty"`
	Version             string `json:"version,omitempty"`
}
