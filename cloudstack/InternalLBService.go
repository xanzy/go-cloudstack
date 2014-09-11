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

package cloudstack

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
	Id      string `json:"id,omitempty"`
	Enabled bool   `json:"enabled,omitempty"`
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
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
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
	Enabled bool   `json:"enabled,omitempty"`
	Nspid   string `json:"nspid,omitempty"`
	Id      string `json:"id,omitempty"`
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
	Requiresupgrade     bool   `json:"requiresupgrade,omitempty"`
	Version             string `json:"version,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Name                string `json:"name,omitempty"`
	Domainid            string `json:"domainid,omitempty"`
	Role                string `json:"role,omitempty"`
	Dns1                string `json:"dns1,omitempty"`
	Linklocalmacaddress string `json:"linklocalmacaddress,omitempty"`
	Guestnetmask        string `json:"guestnetmask,omitempty"`
	Nic                 []struct {
		Networkid    string   `json:"networkid,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Id           string   `json:"id,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Type         string   `json:"type,omitempty"`
	} `json:"nic,omitempty"`
	Hostname           string `json:"hostname,omitempty"`
	Created            string `json:"created,omitempty"`
	Publicmacaddress   string `json:"publicmacaddress,omitempty"`
	Scriptsversion     string `json:"scriptsversion,omitempty"`
	Publicnetmask      string `json:"publicnetmask,omitempty"`
	Zoneid             string `json:"zoneid,omitempty"`
	Templateid         string `json:"templateid,omitempty"`
	Linklocalnetmask   string `json:"linklocalnetmask,omitempty"`
	Gateway            string `json:"gateway,omitempty"`
	Project            string `json:"project,omitempty"`
	Vpcid              string `json:"vpcid,omitempty"`
	Isredundantrouter  bool   `json:"isredundantrouter,omitempty"`
	Networkdomain      string `json:"networkdomain,omitempty"`
	Linklocalnetworkid string `json:"linklocalnetworkid,omitempty"`
	Podid              string `json:"podid,omitempty"`
	Hostid             string `json:"hostid,omitempty"`
	Guestipaddress     string `json:"guestipaddress,omitempty"`
	Domain             string `json:"domain,omitempty"`
	Dns2               string `json:"dns2,omitempty"`
	Ip6dns2            string `json:"ip6dns2,omitempty"`
	Ip6dns1            string `json:"ip6dns1,omitempty"`
	Guestmacaddress    string `json:"guestmacaddress,omitempty"`
	Account            string `json:"account,omitempty"`
	Linklocalip        string `json:"linklocalip,omitempty"`
	Publicip           string `json:"publicip,omitempty"`
	Guestnetworkid     string `json:"guestnetworkid,omitempty"`
	Redundantstate     string `json:"redundantstate,omitempty"`
	Projectid          string `json:"projectid,omitempty"`
	Id                 string `json:"id,omitempty"`
	Zonename           string `json:"zonename,omitempty"`
	Publicnetworkid    string `json:"publicnetworkid,omitempty"`
	State              string `json:"state,omitempty"`
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
	Redundantstate      string `json:"redundantstate,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
	Vpcid               string `json:"vpcid,omitempty"`
	Publicnetworkid     string `json:"publicnetworkid,omitempty"`
	Hostname            string `json:"hostname,omitempty"`
	Id                  string `json:"id,omitempty"`
	Name                string `json:"name,omitempty"`
	Ip6dns1             string `json:"ip6dns1,omitempty"`
	Linklocalnetmask    string `json:"linklocalnetmask,omitempty"`
	Requiresupgrade     bool   `json:"requiresupgrade,omitempty"`
	Linklocalip         string `json:"linklocalip,omitempty"`
	Guestnetworkid      string `json:"guestnetworkid,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
	Version             string `json:"version,omitempty"`
	Hostid              string `json:"hostid,omitempty"`
	Gateway             string `json:"gateway,omitempty"`
	Linklocalmacaddress string `json:"linklocalmacaddress,omitempty"`
	Created             string `json:"created,omitempty"`
	Zoneid              string `json:"zoneid,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Publicmacaddress    string `json:"publicmacaddress,omitempty"`
	Guestnetmask        string `json:"guestnetmask,omitempty"`
	Dns1                string `json:"dns1,omitempty"`
	Podid               string `json:"podid,omitempty"`
	Guestmacaddress     string `json:"guestmacaddress,omitempty"`
	Project             string `json:"project,omitempty"`
	Role                string `json:"role,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Publicip            string `json:"publicip,omitempty"`
	Publicnetmask       string `json:"publicnetmask,omitempty"`
	Ip6dns2             string `json:"ip6dns2,omitempty"`
	Scriptsversion      string `json:"scriptsversion,omitempty"`
	Networkdomain       string `json:"networkdomain,omitempty"`
	Linklocalnetworkid  string `json:"linklocalnetworkid,omitempty"`
	Isredundantrouter   bool   `json:"isredundantrouter,omitempty"`
	Projectid           string `json:"projectid,omitempty"`
	State               string `json:"state,omitempty"`
	Guestipaddress      string `json:"guestipaddress,omitempty"`
	Nic                 []struct {
		Networkid    string   `json:"networkid,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Id           string   `json:"id,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Type         string   `json:"type,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
	} `json:"nic,omitempty"`
	Account    string `json:"account,omitempty"`
	Templateid string `json:"templateid,omitempty"`
	Dns2       string `json:"dns2,omitempty"`
	Domainid   string `json:"domainid,omitempty"`
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
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
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
	Redundantstate    string `json:"redundantstate,omitempty"`
	Isredundantrouter bool   `json:"isredundantrouter,omitempty"`
	Networkdomain     string `json:"networkdomain,omitempty"`
	Ip6dns2           string `json:"ip6dns2,omitempty"`
	Projectid         string `json:"projectid,omitempty"`
	Nic               []struct {
		Networkid    string   `json:"networkid,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Type         string   `json:"type,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Id           string   `json:"id,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
	} `json:"nic,omitempty"`
	Guestipaddress      string `json:"guestipaddress,omitempty"`
	State               string `json:"state,omitempty"`
	Gateway             string `json:"gateway,omitempty"`
	Dns2                string `json:"dns2,omitempty"`
	Id                  string `json:"id,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Version             string `json:"version,omitempty"`
	Podid               string `json:"podid,omitempty"`
	Role                string `json:"role,omitempty"`
	Dns1                string `json:"dns1,omitempty"`
	Vpcid               string `json:"vpcid,omitempty"`
	Publicip            string `json:"publicip,omitempty"`
	Name                string `json:"name,omitempty"`
	Linklocalip         string `json:"linklocalip,omitempty"`
	Linklocalnetworkid  string `json:"linklocalnetworkid,omitempty"`
	Publicnetworkid     string `json:"publicnetworkid,omitempty"`
	Linklocalmacaddress string `json:"linklocalmacaddress,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Publicmacaddress    string `json:"publicmacaddress,omitempty"`
	Guestmacaddress     string `json:"guestmacaddress,omitempty"`
	Scriptsversion      string `json:"scriptsversion,omitempty"`
	Zoneid              string `json:"zoneid,omitempty"`
	Project             string `json:"project,omitempty"`
	Linklocalnetmask    string `json:"linklocalnetmask,omitempty"`
	Hostname            string `json:"hostname,omitempty"`
	Publicnetmask       string `json:"publicnetmask,omitempty"`
	Guestnetworkid      string `json:"guestnetworkid,omitempty"`
	Created             string `json:"created,omitempty"`
	Ip6dns1             string `json:"ip6dns1,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
	Guestnetmask        string `json:"guestnetmask,omitempty"`
	Hostid              string `json:"hostid,omitempty"`
	Requiresupgrade     bool   `json:"requiresupgrade,omitempty"`
	Domainid            string `json:"domainid,omitempty"`
	Templateid          string `json:"templateid,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
	Account             string `json:"account,omitempty"`
}
