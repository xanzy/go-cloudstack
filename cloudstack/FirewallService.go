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
	"strings"
)

type ListPortForwardingRulesParams struct {
	p map[string]interface{}
}

func (p *ListPortForwardingRulesParams) toURLValues() url.Values {
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
	if v, found := p.p["ipaddressid"]; found {
		u.Set("ipaddressid", v.(string))
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
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
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

func (p *ListPortForwardingRulesParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListPortForwardingRulesParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListPortForwardingRulesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListPortForwardingRulesParams) SetIpaddressid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddressid"] = v
	return
}

func (p *ListPortForwardingRulesParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
	return
}

func (p *ListPortForwardingRulesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListPortForwardingRulesParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
	return
}

func (p *ListPortForwardingRulesParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
	return
}

func (p *ListPortForwardingRulesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListPortForwardingRulesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListPortForwardingRulesParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *ListPortForwardingRulesParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
	return
}

// You should always use this function to get a new ListPortForwardingRulesParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewListPortForwardingRulesParams() *ListPortForwardingRulesParams {
	p := &ListPortForwardingRulesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *FirewallService) GetPortForwardingRuleID(keyword string) (string, error) {
	p := &ListPortForwardingRulesParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword

	l, err := s.ListPortForwardingRules(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, keyword, l)
	}
	return l.PortForwardingRules[0].Id, nil
}

// Lists all port forwarding rules for an IP address.
func (s *FirewallService) ListPortForwardingRules(p *ListPortForwardingRulesParams) (*ListPortForwardingRulesResponse, error) {
	resp, err := s.cs.newRequest("listPortForwardingRules", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListPortForwardingRulesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListPortForwardingRulesResponse struct {
	Count               int                   `json:"count"`
	PortForwardingRules []*PortForwardingRule `json:"portforwardingrule"`
}

type PortForwardingRule struct {
	Privateport               string `json:"privateport,omitempty"`
	Virtualmachinedisplayname string `json:"virtualmachinedisplayname,omitempty"`
	Tags                      []struct {
		Value        string `json:"value,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Project      string `json:"project,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Account      string `json:"account,omitempty"`
		Key          string `json:"key,omitempty"`
		Customer     string `json:"customer,omitempty"`
	} `json:"tags,omitempty"`
	Id                 string `json:"id,omitempty"`
	Cidrlist           string `json:"cidrlist,omitempty"`
	Publicport         string `json:"publicport,omitempty"`
	Virtualmachineid   string `json:"virtualmachineid,omitempty"`
	Privateendport     string `json:"privateendport,omitempty"`
	Ipaddress          string `json:"ipaddress,omitempty"`
	Vmguestip          string `json:"vmguestip,omitempty"`
	Protocol           string `json:"protocol,omitempty"`
	Ipaddressid        string `json:"ipaddressid,omitempty"`
	Virtualmachinename string `json:"virtualmachinename,omitempty"`
	Networkid          string `json:"networkid,omitempty"`
	Publicendport      string `json:"publicendport,omitempty"`
	State              string `json:"state,omitempty"`
}

type CreatePortForwardingRuleParams struct {
	p map[string]interface{}
}

func (p *CreatePortForwardingRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["cidrlist"]; found {
		vv := strings.Join(v.([]string), ", ")
		u.Set("cidrlist", vv)
	}
	if v, found := p.p["ipaddressid"]; found {
		u.Set("ipaddressid", v.(string))
	}
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := p.p["openfirewall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("openfirewall", vv)
	}
	if v, found := p.p["privateendport"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("privateendport", vv)
	}
	if v, found := p.p["privateport"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("privateport", vv)
	}
	if v, found := p.p["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := p.p["publicendport"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("publicendport", vv)
	}
	if v, found := p.p["publicport"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("publicport", vv)
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	if v, found := p.p["vmguestip"]; found {
		u.Set("vmguestip", v.(string))
	}
	return u
}

func (p *CreatePortForwardingRuleParams) SetCidrlist(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cidrlist"] = v
	return
}

func (p *CreatePortForwardingRuleParams) SetIpaddressid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddressid"] = v
	return
}

func (p *CreatePortForwardingRuleParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
	return
}

func (p *CreatePortForwardingRuleParams) SetOpenfirewall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["openfirewall"] = v
	return
}

func (p *CreatePortForwardingRuleParams) SetPrivateendport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["privateendport"] = v
	return
}

func (p *CreatePortForwardingRuleParams) SetPrivateport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["privateport"] = v
	return
}

func (p *CreatePortForwardingRuleParams) SetProtocol(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["protocol"] = v
	return
}

func (p *CreatePortForwardingRuleParams) SetPublicendport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["publicendport"] = v
	return
}

func (p *CreatePortForwardingRuleParams) SetPublicport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["publicport"] = v
	return
}

func (p *CreatePortForwardingRuleParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
	return
}

func (p *CreatePortForwardingRuleParams) SetVmguestip(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vmguestip"] = v
	return
}

// You should always use this function to get a new CreatePortForwardingRuleParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewCreatePortForwardingRuleParams(ipaddressid string, privateport int, protocol string, publicport int, virtualmachineid string) *CreatePortForwardingRuleParams {
	p := &CreatePortForwardingRuleParams{}
	p.p = make(map[string]interface{})
	p.p["ipaddressid"] = ipaddressid
	p.p["privateport"] = privateport
	p.p["protocol"] = protocol
	p.p["publicport"] = publicport
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// Creates a port forwarding rule
func (s *FirewallService) CreatePortForwardingRule(p *CreatePortForwardingRuleParams) (*CreatePortForwardingRuleResponse, error) {
	resp, err := s.cs.newRequest("createPortForwardingRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreatePortForwardingRuleResponse
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

		var r CreatePortForwardingRuleResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type CreatePortForwardingRuleResponse struct {
	JobID string `json:"jobid,omitempty"`
	Tags  []struct {
		Project      string `json:"project,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Key          string `json:"key,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Value        string `json:"value,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Account      string `json:"account,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
	} `json:"tags,omitempty"`
	Protocol                  string `json:"protocol,omitempty"`
	Virtualmachinedisplayname string `json:"virtualmachinedisplayname,omitempty"`
	Virtualmachineid          string `json:"virtualmachineid,omitempty"`
	Vmguestip                 string `json:"vmguestip,omitempty"`
	Privateport               string `json:"privateport,omitempty"`
	Privateendport            string `json:"privateendport,omitempty"`
	Publicport                string `json:"publicport,omitempty"`
	Ipaddressid               string `json:"ipaddressid,omitempty"`
	Cidrlist                  string `json:"cidrlist,omitempty"`
	State                     string `json:"state,omitempty"`
	Ipaddress                 string `json:"ipaddress,omitempty"`
	Virtualmachinename        string `json:"virtualmachinename,omitempty"`
	Id                        string `json:"id,omitempty"`
	Publicendport             string `json:"publicendport,omitempty"`
	Networkid                 string `json:"networkid,omitempty"`
}

type DeletePortForwardingRuleParams struct {
	p map[string]interface{}
}

func (p *DeletePortForwardingRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeletePortForwardingRuleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeletePortForwardingRuleParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewDeletePortForwardingRuleParams(id string) *DeletePortForwardingRuleParams {
	p := &DeletePortForwardingRuleParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a port forwarding rule
func (s *FirewallService) DeletePortForwardingRule(p *DeletePortForwardingRuleParams) (*DeletePortForwardingRuleResponse, error) {
	resp, err := s.cs.newRequest("deletePortForwardingRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeletePortForwardingRuleResponse
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

		var r DeletePortForwardingRuleResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type DeletePortForwardingRuleResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type UpdatePortForwardingRuleParams struct {
	p map[string]interface{}
}

func (p *UpdatePortForwardingRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["ipaddressid"]; found {
		u.Set("ipaddressid", v.(string))
	}
	if v, found := p.p["privateip"]; found {
		u.Set("privateip", v.(string))
	}
	if v, found := p.p["privateport"]; found {
		u.Set("privateport", v.(string))
	}
	if v, found := p.p["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := p.p["publicport"]; found {
		u.Set("publicport", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *UpdatePortForwardingRuleParams) SetIpaddressid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddressid"] = v
	return
}

func (p *UpdatePortForwardingRuleParams) SetPrivateip(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["privateip"] = v
	return
}

func (p *UpdatePortForwardingRuleParams) SetPrivateport(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["privateport"] = v
	return
}

func (p *UpdatePortForwardingRuleParams) SetProtocol(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["protocol"] = v
	return
}

func (p *UpdatePortForwardingRuleParams) SetPublicport(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["publicport"] = v
	return
}

func (p *UpdatePortForwardingRuleParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
	return
}

// You should always use this function to get a new UpdatePortForwardingRuleParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewUpdatePortForwardingRuleParams(ipaddressid string, privateport string, protocol string, publicport string) *UpdatePortForwardingRuleParams {
	p := &UpdatePortForwardingRuleParams{}
	p.p = make(map[string]interface{})
	p.p["ipaddressid"] = ipaddressid
	p.p["privateport"] = privateport
	p.p["protocol"] = protocol
	p.p["publicport"] = publicport
	return p
}

// Updates a port forwarding rule.  Only the private port and the virtual machine can be updated.
func (s *FirewallService) UpdatePortForwardingRule(p *UpdatePortForwardingRuleParams) (*UpdatePortForwardingRuleResponse, error) {
	resp, err := s.cs.newRequest("updatePortForwardingRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdatePortForwardingRuleResponse
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

		var r UpdatePortForwardingRuleResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type UpdatePortForwardingRuleResponse struct {
	JobID                     string `json:"jobid,omitempty"`
	Ipaddress                 string `json:"ipaddress,omitempty"`
	Vmguestip                 string `json:"vmguestip,omitempty"`
	Virtualmachinedisplayname string `json:"virtualmachinedisplayname,omitempty"`
	Tags                      []struct {
		Value        string `json:"value,omitempty"`
		Account      string `json:"account,omitempty"`
		Project      string `json:"project,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Key          string `json:"key,omitempty"`
	} `json:"tags,omitempty"`
	State              string `json:"state,omitempty"`
	Networkid          string `json:"networkid,omitempty"`
	Id                 string `json:"id,omitempty"`
	Publicendport      string `json:"publicendport,omitempty"`
	Virtualmachinename string `json:"virtualmachinename,omitempty"`
	Cidrlist           string `json:"cidrlist,omitempty"`
	Privateport        string `json:"privateport,omitempty"`
	Protocol           string `json:"protocol,omitempty"`
	Ipaddressid        string `json:"ipaddressid,omitempty"`
	Publicport         string `json:"publicport,omitempty"`
	Privateendport     string `json:"privateendport,omitempty"`
	Virtualmachineid   string `json:"virtualmachineid,omitempty"`
}

type CreateFirewallRuleParams struct {
	p map[string]interface{}
}

func (p *CreateFirewallRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["cidrlist"]; found {
		vv := strings.Join(v.([]string), ", ")
		u.Set("cidrlist", vv)
	}
	if v, found := p.p["endport"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("endport", vv)
	}
	if v, found := p.p["icmpcode"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("icmpcode", vv)
	}
	if v, found := p.p["icmptype"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("icmptype", vv)
	}
	if v, found := p.p["ipaddressid"]; found {
		u.Set("ipaddressid", v.(string))
	}
	if v, found := p.p["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := p.p["startport"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("startport", vv)
	}
	if v, found := p.p["type"]; found {
		u.Set("type", v.(string))
	}
	return u
}

func (p *CreateFirewallRuleParams) SetCidrlist(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cidrlist"] = v
	return
}

func (p *CreateFirewallRuleParams) SetEndport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["endport"] = v
	return
}

func (p *CreateFirewallRuleParams) SetIcmpcode(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["icmpcode"] = v
	return
}

func (p *CreateFirewallRuleParams) SetIcmptype(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["icmptype"] = v
	return
}

func (p *CreateFirewallRuleParams) SetIpaddressid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddressid"] = v
	return
}

func (p *CreateFirewallRuleParams) SetProtocol(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["protocol"] = v
	return
}

func (p *CreateFirewallRuleParams) SetStartport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startport"] = v
	return
}

func (p *CreateFirewallRuleParams) SetType(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["firewallType"] = v
	return
}

// You should always use this function to get a new CreateFirewallRuleParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewCreateFirewallRuleParams(ipaddressid string, protocol string) *CreateFirewallRuleParams {
	p := &CreateFirewallRuleParams{}
	p.p = make(map[string]interface{})
	p.p["ipaddressid"] = ipaddressid
	p.p["protocol"] = protocol
	return p
}

// Creates a firewall rule for a given ip address
func (s *FirewallService) CreateFirewallRule(p *CreateFirewallRuleParams) (*CreateFirewallRuleResponse, error) {
	resp, err := s.cs.newRequest("createFirewallRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateFirewallRuleResponse
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

		var r CreateFirewallRuleResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type CreateFirewallRuleResponse struct {
	JobID    string `json:"jobid,omitempty"`
	Cidrlist string `json:"cidrlist,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	Icmptype int    `json:"icmptype,omitempty"`
	State    string `json:"state,omitempty"`
	Tags     []struct {
		Project      string `json:"project,omitempty"`
		Key          string `json:"key,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Value        string `json:"value,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Account      string `json:"account,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
	} `json:"tags,omitempty"`
	Icmpcode    int    `json:"icmpcode,omitempty"`
	Startport   string `json:"startport,omitempty"`
	Ipaddress   string `json:"ipaddress,omitempty"`
	Ipaddressid string `json:"ipaddressid,omitempty"`
	Id          string `json:"id,omitempty"`
	Networkid   string `json:"networkid,omitempty"`
	Endport     string `json:"endport,omitempty"`
}

type DeleteFirewallRuleParams struct {
	p map[string]interface{}
}

func (p *DeleteFirewallRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteFirewallRuleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteFirewallRuleParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewDeleteFirewallRuleParams(id string) *DeleteFirewallRuleParams {
	p := &DeleteFirewallRuleParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a firewall rule
func (s *FirewallService) DeleteFirewallRule(p *DeleteFirewallRuleParams) (*DeleteFirewallRuleResponse, error) {
	resp, err := s.cs.newRequest("deleteFirewallRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteFirewallRuleResponse
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

		var r DeleteFirewallRuleResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type DeleteFirewallRuleResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type ListFirewallRulesParams struct {
	p map[string]interface{}
}

func (p *ListFirewallRulesParams) toURLValues() url.Values {
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
	if v, found := p.p["ipaddressid"]; found {
		u.Set("ipaddressid", v.(string))
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
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
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

func (p *ListFirewallRulesParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListFirewallRulesParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListFirewallRulesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListFirewallRulesParams) SetIpaddressid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddressid"] = v
	return
}

func (p *ListFirewallRulesParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
	return
}

func (p *ListFirewallRulesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListFirewallRulesParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
	return
}

func (p *ListFirewallRulesParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
	return
}

func (p *ListFirewallRulesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListFirewallRulesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListFirewallRulesParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *ListFirewallRulesParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
	return
}

// You should always use this function to get a new ListFirewallRulesParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewListFirewallRulesParams() *ListFirewallRulesParams {
	p := &ListFirewallRulesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *FirewallService) GetFirewallRuleID(keyword string) (string, error) {
	p := &ListFirewallRulesParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword

	l, err := s.ListFirewallRules(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, keyword, l)
	}
	return l.FirewallRules[0].Id, nil
}

// Lists all firewall rules for an IP address.
func (s *FirewallService) ListFirewallRules(p *ListFirewallRulesParams) (*ListFirewallRulesResponse, error) {
	resp, err := s.cs.newRequest("listFirewallRules", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListFirewallRulesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListFirewallRulesResponse struct {
	Count         int             `json:"count"`
	FirewallRules []*FirewallRule `json:"firewallrule"`
}

type FirewallRule struct {
	State       string `json:"state,omitempty"`
	Ipaddress   string `json:"ipaddress,omitempty"`
	Id          string `json:"id,omitempty"`
	Ipaddressid string `json:"ipaddressid,omitempty"`
	Networkid   string `json:"networkid,omitempty"`
	Startport   string `json:"startport,omitempty"`
	Tags        []struct {
		Domain       string `json:"domain,omitempty"`
		Key          string `json:"key,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Account      string `json:"account,omitempty"`
		Value        string `json:"value,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Project      string `json:"project,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
	} `json:"tags,omitempty"`
	Endport  string `json:"endport,omitempty"`
	Icmpcode int    `json:"icmpcode,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	Icmptype int    `json:"icmptype,omitempty"`
	Cidrlist string `json:"cidrlist,omitempty"`
}

type CreateEgressFirewallRuleParams struct {
	p map[string]interface{}
}

func (p *CreateEgressFirewallRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["cidrlist"]; found {
		vv := strings.Join(v.([]string), ", ")
		u.Set("cidrlist", vv)
	}
	if v, found := p.p["endport"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("endport", vv)
	}
	if v, found := p.p["icmpcode"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("icmpcode", vv)
	}
	if v, found := p.p["icmptype"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("icmptype", vv)
	}
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := p.p["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := p.p["startport"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("startport", vv)
	}
	if v, found := p.p["type"]; found {
		u.Set("type", v.(string))
	}
	return u
}

func (p *CreateEgressFirewallRuleParams) SetCidrlist(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cidrlist"] = v
	return
}

func (p *CreateEgressFirewallRuleParams) SetEndport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["endport"] = v
	return
}

func (p *CreateEgressFirewallRuleParams) SetIcmpcode(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["icmpcode"] = v
	return
}

func (p *CreateEgressFirewallRuleParams) SetIcmptype(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["icmptype"] = v
	return
}

func (p *CreateEgressFirewallRuleParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
	return
}

func (p *CreateEgressFirewallRuleParams) SetProtocol(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["protocol"] = v
	return
}

func (p *CreateEgressFirewallRuleParams) SetStartport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startport"] = v
	return
}

func (p *CreateEgressFirewallRuleParams) SetType(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["firewallType"] = v
	return
}

// You should always use this function to get a new CreateEgressFirewallRuleParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewCreateEgressFirewallRuleParams(networkid string, protocol string) *CreateEgressFirewallRuleParams {
	p := &CreateEgressFirewallRuleParams{}
	p.p = make(map[string]interface{})
	p.p["networkid"] = networkid
	p.p["protocol"] = protocol
	return p
}

// Creates a egress firewall rule for a given network
func (s *FirewallService) CreateEgressFirewallRule(p *CreateEgressFirewallRuleParams) (*CreateEgressFirewallRuleResponse, error) {
	resp, err := s.cs.newRequest("createEgressFirewallRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateEgressFirewallRuleResponse
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

		var r CreateEgressFirewallRuleResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type CreateEgressFirewallRuleResponse struct {
	JobID     string `json:"jobid,omitempty"`
	Id        string `json:"id,omitempty"`
	Protocol  string `json:"protocol,omitempty"`
	State     string `json:"state,omitempty"`
	Startport string `json:"startport,omitempty"`
	Icmpcode  int    `json:"icmpcode,omitempty"`
	Cidrlist  string `json:"cidrlist,omitempty"`
	Ipaddress string `json:"ipaddress,omitempty"`
	Tags      []struct {
		Domain       string `json:"domain,omitempty"`
		Key          string `json:"key,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Project      string `json:"project,omitempty"`
		Account      string `json:"account,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Value        string `json:"value,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
	} `json:"tags,omitempty"`
	Icmptype    int    `json:"icmptype,omitempty"`
	Endport     string `json:"endport,omitempty"`
	Networkid   string `json:"networkid,omitempty"`
	Ipaddressid string `json:"ipaddressid,omitempty"`
}

type DeleteEgressFirewallRuleParams struct {
	p map[string]interface{}
}

func (p *DeleteEgressFirewallRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteEgressFirewallRuleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteEgressFirewallRuleParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewDeleteEgressFirewallRuleParams(id string) *DeleteEgressFirewallRuleParams {
	p := &DeleteEgressFirewallRuleParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes an ggress firewall rule
func (s *FirewallService) DeleteEgressFirewallRule(p *DeleteEgressFirewallRuleParams) (*DeleteEgressFirewallRuleResponse, error) {
	resp, err := s.cs.newRequest("deleteEgressFirewallRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteEgressFirewallRuleResponse
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

		var r DeleteEgressFirewallRuleResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type DeleteEgressFirewallRuleResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type ListEgressFirewallRulesParams struct {
	p map[string]interface{}
}

func (p *ListEgressFirewallRulesParams) toURLValues() url.Values {
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
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["ipaddressid"]; found {
		u.Set("ipaddressid", v.(string))
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
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
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
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
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

func (p *ListEgressFirewallRulesParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListEgressFirewallRulesParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListEgressFirewallRulesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListEgressFirewallRulesParams) SetIpaddressid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddressid"] = v
	return
}

func (p *ListEgressFirewallRulesParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
	return
}

func (p *ListEgressFirewallRulesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListEgressFirewallRulesParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
	return
}

func (p *ListEgressFirewallRulesParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
	return
}

func (p *ListEgressFirewallRulesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListEgressFirewallRulesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListEgressFirewallRulesParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *ListEgressFirewallRulesParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
	return
}

// You should always use this function to get a new ListEgressFirewallRulesParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewListEgressFirewallRulesParams() *ListEgressFirewallRulesParams {
	p := &ListEgressFirewallRulesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *FirewallService) GetEgressFirewallRuleID(keyword string) (string, error) {
	p := &ListEgressFirewallRulesParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword

	l, err := s.ListEgressFirewallRules(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, keyword, l)
	}
	return l.EgressFirewallRules[0].Id, nil
}

// Lists all egress firewall rules for network id.
func (s *FirewallService) ListEgressFirewallRules(p *ListEgressFirewallRulesParams) (*ListEgressFirewallRulesResponse, error) {
	resp, err := s.cs.newRequest("listEgressFirewallRules", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListEgressFirewallRulesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListEgressFirewallRulesResponse struct {
	Count               int                   `json:"count"`
	EgressFirewallRules []*EgressFirewallRule `json:"egressfirewallrule"`
}

type EgressFirewallRule struct {
	Icmptype    int    `json:"icmptype,omitempty"`
	Id          string `json:"id,omitempty"`
	Ipaddress   string `json:"ipaddress,omitempty"`
	Startport   string `json:"startport,omitempty"`
	Endport     string `json:"endport,omitempty"`
	Cidrlist    string `json:"cidrlist,omitempty"`
	Networkid   string `json:"networkid,omitempty"`
	Ipaddressid string `json:"ipaddressid,omitempty"`
	Tags        []struct {
		Resourceid   string `json:"resourceid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Account      string `json:"account,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Project      string `json:"project,omitempty"`
		Value        string `json:"value,omitempty"`
		Key          string `json:"key,omitempty"`
		Domain       string `json:"domain,omitempty"`
	} `json:"tags,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	State    string `json:"state,omitempty"`
	Icmpcode int    `json:"icmpcode,omitempty"`
}

type AddPaloAltoFirewallParams struct {
	p map[string]interface{}
}

func (p *AddPaloAltoFirewallParams) toURLValues() url.Values {
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

func (p *AddPaloAltoFirewallParams) SetNetworkdevicetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkdevicetype"] = v
	return
}

func (p *AddPaloAltoFirewallParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
	return
}

func (p *AddPaloAltoFirewallParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
	return
}

func (p *AddPaloAltoFirewallParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
	return
}

func (p *AddPaloAltoFirewallParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
	return
}

// You should always use this function to get a new AddPaloAltoFirewallParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewAddPaloAltoFirewallParams(networkdevicetype string, password string, physicalnetworkid string, url string, username string) *AddPaloAltoFirewallParams {
	p := &AddPaloAltoFirewallParams{}
	p.p = make(map[string]interface{})
	p.p["networkdevicetype"] = networkdevicetype
	p.p["password"] = password
	p.p["physicalnetworkid"] = physicalnetworkid
	p.p["url"] = url
	p.p["username"] = username
	return p
}

// Adds a Palo Alto firewall device
func (s *FirewallService) AddPaloAltoFirewall(p *AddPaloAltoFirewallParams) (*AddPaloAltoFirewallResponse, error) {
	resp, err := s.cs.newRequest("addPaloAltoFirewall", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddPaloAltoFirewallResponse
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

		var r AddPaloAltoFirewallResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type AddPaloAltoFirewallResponse struct {
	JobID             string `json:"jobid,omitempty"`
	Privateinterface  string `json:"privateinterface,omitempty"`
	Provider          string `json:"provider,omitempty"`
	Privatezone       string `json:"privatezone,omitempty"`
	Fwdevicecapacity  int    `json:"fwdevicecapacity,omitempty"`
	Usageinterface    string `json:"usageinterface,omitempty"`
	Fwdevicename      string `json:"fwdevicename,omitempty"`
	Ipaddress         string `json:"ipaddress,omitempty"`
	Physicalnetworkid string `json:"physicalnetworkid,omitempty"`
	Publicinterface   string `json:"publicinterface,omitempty"`
	Publiczone        string `json:"publiczone,omitempty"`
	Timeout           string `json:"timeout,omitempty"`
	Fwdevicestate     string `json:"fwdevicestate,omitempty"`
	Zoneid            string `json:"zoneid,omitempty"`
	Numretries        string `json:"numretries,omitempty"`
	Username          string `json:"username,omitempty"`
	Fwdeviceid        string `json:"fwdeviceid,omitempty"`
}

type DeletePaloAltoFirewallParams struct {
	p map[string]interface{}
}

func (p *DeletePaloAltoFirewallParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["fwdeviceid"]; found {
		u.Set("fwdeviceid", v.(string))
	}
	return u
}

func (p *DeletePaloAltoFirewallParams) SetFwdeviceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fwdeviceid"] = v
	return
}

// You should always use this function to get a new DeletePaloAltoFirewallParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewDeletePaloAltoFirewallParams(fwdeviceid string) *DeletePaloAltoFirewallParams {
	p := &DeletePaloAltoFirewallParams{}
	p.p = make(map[string]interface{})
	p.p["fwdeviceid"] = fwdeviceid
	return p
}

//  delete a Palo Alto firewall device
func (s *FirewallService) DeletePaloAltoFirewall(p *DeletePaloAltoFirewallParams) (*DeletePaloAltoFirewallResponse, error) {
	resp, err := s.cs.newRequest("deletePaloAltoFirewall", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeletePaloAltoFirewallResponse
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

		var r DeletePaloAltoFirewallResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type DeletePaloAltoFirewallResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Success     bool   `json:"success,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
}

type ConfigurePaloAltoFirewallParams struct {
	p map[string]interface{}
}

func (p *ConfigurePaloAltoFirewallParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["fwdevicecapacity"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("fwdevicecapacity", vv)
	}
	if v, found := p.p["fwdeviceid"]; found {
		u.Set("fwdeviceid", v.(string))
	}
	return u
}

func (p *ConfigurePaloAltoFirewallParams) SetFwdevicecapacity(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fwdevicecapacity"] = v
	return
}

func (p *ConfigurePaloAltoFirewallParams) SetFwdeviceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fwdeviceid"] = v
	return
}

// You should always use this function to get a new ConfigurePaloAltoFirewallParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewConfigurePaloAltoFirewallParams(fwdeviceid string) *ConfigurePaloAltoFirewallParams {
	p := &ConfigurePaloAltoFirewallParams{}
	p.p = make(map[string]interface{})
	p.p["fwdeviceid"] = fwdeviceid
	return p
}

// Configures a Palo Alto firewall device
func (s *FirewallService) ConfigurePaloAltoFirewall(p *ConfigurePaloAltoFirewallParams) (*ConfigurePaloAltoFirewallResponse, error) {
	resp, err := s.cs.newRequest("configurePaloAltoFirewall", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ConfigurePaloAltoFirewallResponse
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

		var r ConfigurePaloAltoFirewallResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type ConfigurePaloAltoFirewallResponse struct {
	JobID             string `json:"jobid,omitempty"`
	Zoneid            string `json:"zoneid,omitempty"`
	Timeout           string `json:"timeout,omitempty"`
	Provider          string `json:"provider,omitempty"`
	Physicalnetworkid string `json:"physicalnetworkid,omitempty"`
	Username          string `json:"username,omitempty"`
	Fwdevicename      string `json:"fwdevicename,omitempty"`
	Fwdeviceid        string `json:"fwdeviceid,omitempty"`
	Fwdevicestate     string `json:"fwdevicestate,omitempty"`
	Numretries        string `json:"numretries,omitempty"`
	Fwdevicecapacity  int    `json:"fwdevicecapacity,omitempty"`
	Usageinterface    string `json:"usageinterface,omitempty"`
	Privateinterface  string `json:"privateinterface,omitempty"`
	Privatezone       string `json:"privatezone,omitempty"`
	Ipaddress         string `json:"ipaddress,omitempty"`
	Publiczone        string `json:"publiczone,omitempty"`
	Publicinterface   string `json:"publicinterface,omitempty"`
}

type ListPaloAltoFirewallsParams struct {
	p map[string]interface{}
}

func (p *ListPaloAltoFirewallsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["fwdeviceid"]; found {
		u.Set("fwdeviceid", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("pagesize", vv)
	}
	if v, found := p.p["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	return u
}

func (p *ListPaloAltoFirewallsParams) SetFwdeviceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fwdeviceid"] = v
	return
}

func (p *ListPaloAltoFirewallsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListPaloAltoFirewallsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListPaloAltoFirewallsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListPaloAltoFirewallsParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
	return
}

// You should always use this function to get a new ListPaloAltoFirewallsParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewListPaloAltoFirewallsParams() *ListPaloAltoFirewallsParams {
	p := &ListPaloAltoFirewallsParams{}
	p.p = make(map[string]interface{})
	return p
}

// lists Palo Alto firewall devices in a physical network
func (s *FirewallService) ListPaloAltoFirewalls(p *ListPaloAltoFirewallsParams) (*ListPaloAltoFirewallsResponse, error) {
	resp, err := s.cs.newRequest("listPaloAltoFirewalls", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListPaloAltoFirewallsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListPaloAltoFirewallsResponse struct {
	Count             int                 `json:"count"`
	PaloAltoFirewalls []*PaloAltoFirewall `json:"paloaltofirewall"`
}

type PaloAltoFirewall struct {
	Username          string `json:"username,omitempty"`
	Privateinterface  string `json:"privateinterface,omitempty"`
	Usageinterface    string `json:"usageinterface,omitempty"`
	Physicalnetworkid string `json:"physicalnetworkid,omitempty"`
	Numretries        string `json:"numretries,omitempty"`
	Publicinterface   string `json:"publicinterface,omitempty"`
	Privatezone       string `json:"privatezone,omitempty"`
	Provider          string `json:"provider,omitempty"`
	Zoneid            string `json:"zoneid,omitempty"`
	Fwdevicecapacity  int    `json:"fwdevicecapacity,omitempty"`
	Timeout           string `json:"timeout,omitempty"`
	Fwdeviceid        string `json:"fwdeviceid,omitempty"`
	Fwdevicename      string `json:"fwdevicename,omitempty"`
	Fwdevicestate     string `json:"fwdevicestate,omitempty"`
	Publiczone        string `json:"publiczone,omitempty"`
	Ipaddress         string `json:"ipaddress,omitempty"`
}
