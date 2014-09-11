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
	"strings"
)

type CreateNetworkACLParams struct {
	p map[string]interface{}
}

func (p *CreateNetworkACLParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["aclid"]; found {
		u.Set("aclid", v.(string))
	}
	if v, found := p.p["action"]; found {
		u.Set("action", v.(string))
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
	if v, found := p.p["number"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("number", vv)
	}
	if v, found := p.p["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := p.p["startport"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("startport", vv)
	}
	if v, found := p.p["traffictype"]; found {
		u.Set("traffictype", v.(string))
	}
	return u
}

func (p *CreateNetworkACLParams) SetAclid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["aclid"] = v
	return
}

func (p *CreateNetworkACLParams) SetAction(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["action"] = v
	return
}

func (p *CreateNetworkACLParams) SetCidrlist(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cidrlist"] = v
	return
}

func (p *CreateNetworkACLParams) SetEndport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["endport"] = v
	return
}

func (p *CreateNetworkACLParams) SetIcmpcode(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["icmpcode"] = v
	return
}

func (p *CreateNetworkACLParams) SetIcmptype(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["icmptype"] = v
	return
}

func (p *CreateNetworkACLParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
	return
}

func (p *CreateNetworkACLParams) SetNumber(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["number"] = v
	return
}

func (p *CreateNetworkACLParams) SetProtocol(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["protocol"] = v
	return
}

func (p *CreateNetworkACLParams) SetStartport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startport"] = v
	return
}

func (p *CreateNetworkACLParams) SetTraffictype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["traffictype"] = v
	return
}

// You should always use this function to get a new CreateNetworkACLParams instance,
// as then you are sure you have configured all required params
func (s *NetworkACLService) NewCreateNetworkACLParams(protocol string) *CreateNetworkACLParams {
	p := &CreateNetworkACLParams{}
	p.p = make(map[string]interface{})
	p.p["protocol"] = protocol
	return p
}

// Creates a ACL rule in the given network (the network has to belong to VPC)
func (s *NetworkACLService) CreateNetworkACL(p *CreateNetworkACLParams) (*CreateNetworkACLResponse, error) {
	resp, err := s.cs.newRequest("createNetworkACL", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateNetworkACLResponse
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

		var r CreateNetworkACLResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type CreateNetworkACLResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Traffictype string `json:"traffictype,omitempty"`
	Cidrlist    string `json:"cidrlist,omitempty"`
	Action      string `json:"action,omitempty"`
	Startport   string `json:"startport,omitempty"`
	Protocol    string `json:"protocol,omitempty"`
	Icmpcode    int    `json:"icmpcode,omitempty"`
	State       string `json:"state,omitempty"`
	Number      int    `json:"number,omitempty"`
	Endport     string `json:"endport,omitempty"`
	Icmptype    int    `json:"icmptype,omitempty"`
	Tags        []struct {
		Key          string `json:"key,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Project      string `json:"project,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Value        string `json:"value,omitempty"`
		Account      string `json:"account,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
	} `json:"tags,omitempty"`
	Aclid string `json:"aclid,omitempty"`
	Id    string `json:"id,omitempty"`
}

type UpdateNetworkACLItemParams struct {
	p map[string]interface{}
}

func (p *UpdateNetworkACLItemParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["action"]; found {
		u.Set("action", v.(string))
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
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["number"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("number", vv)
	}
	if v, found := p.p["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := p.p["startport"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("startport", vv)
	}
	if v, found := p.p["traffictype"]; found {
		u.Set("traffictype", v.(string))
	}
	return u
}

func (p *UpdateNetworkACLItemParams) SetAction(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["action"] = v
	return
}

func (p *UpdateNetworkACLItemParams) SetCidrlist(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cidrlist"] = v
	return
}

func (p *UpdateNetworkACLItemParams) SetEndport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["endport"] = v
	return
}

func (p *UpdateNetworkACLItemParams) SetIcmpcode(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["icmpcode"] = v
	return
}

func (p *UpdateNetworkACLItemParams) SetIcmptype(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["icmptype"] = v
	return
}

func (p *UpdateNetworkACLItemParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *UpdateNetworkACLItemParams) SetNumber(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["number"] = v
	return
}

func (p *UpdateNetworkACLItemParams) SetProtocol(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["protocol"] = v
	return
}

func (p *UpdateNetworkACLItemParams) SetStartport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startport"] = v
	return
}

func (p *UpdateNetworkACLItemParams) SetTraffictype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["traffictype"] = v
	return
}

// You should always use this function to get a new UpdateNetworkACLItemParams instance,
// as then you are sure you have configured all required params
func (s *NetworkACLService) NewUpdateNetworkACLItemParams(id string) *UpdateNetworkACLItemParams {
	p := &UpdateNetworkACLItemParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates ACL Item with specified Id
func (s *NetworkACLService) UpdateNetworkACLItem(p *UpdateNetworkACLItemParams) (*UpdateNetworkACLItemResponse, error) {
	resp, err := s.cs.newRequest("updateNetworkACLItem", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateNetworkACLItemResponse
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

		var r UpdateNetworkACLItemResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type UpdateNetworkACLItemResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Protocol    string `json:"protocol,omitempty"`
	Cidrlist    string `json:"cidrlist,omitempty"`
	Action      string `json:"action,omitempty"`
	Traffictype string `json:"traffictype,omitempty"`
	Tags        []struct {
		Projectid    string `json:"projectid,omitempty"`
		Key          string `json:"key,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Project      string `json:"project,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Value        string `json:"value,omitempty"`
		Account      string `json:"account,omitempty"`
	} `json:"tags,omitempty"`
	Icmptype  int    `json:"icmptype,omitempty"`
	State     string `json:"state,omitempty"`
	Startport string `json:"startport,omitempty"`
	Aclid     string `json:"aclid,omitempty"`
	Icmpcode  int    `json:"icmpcode,omitempty"`
	Number    int    `json:"number,omitempty"`
	Id        string `json:"id,omitempty"`
	Endport   string `json:"endport,omitempty"`
}

type DeleteNetworkACLParams struct {
	p map[string]interface{}
}

func (p *DeleteNetworkACLParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteNetworkACLParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteNetworkACLParams instance,
// as then you are sure you have configured all required params
func (s *NetworkACLService) NewDeleteNetworkACLParams(id string) *DeleteNetworkACLParams {
	p := &DeleteNetworkACLParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a Network ACL
func (s *NetworkACLService) DeleteNetworkACL(p *DeleteNetworkACLParams) (*DeleteNetworkACLResponse, error) {
	resp, err := s.cs.newRequest("deleteNetworkACL", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteNetworkACLResponse
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

		var r DeleteNetworkACLResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type DeleteNetworkACLResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type ListNetworkACLsParams struct {
	p map[string]interface{}
}

func (p *ListNetworkACLsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["aclid"]; found {
		u.Set("aclid", v.(string))
	}
	if v, found := p.p["action"]; found {
		u.Set("action", v.(string))
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
	if v, found := p.p["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := p.p["tags"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["traffictype"]; found {
		u.Set("traffictype", v.(string))
	}
	return u
}

func (p *ListNetworkACLsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListNetworkACLsParams) SetAclid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["aclid"] = v
	return
}

func (p *ListNetworkACLsParams) SetAction(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["action"] = v
	return
}

func (p *ListNetworkACLsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListNetworkACLsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListNetworkACLsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
	return
}

func (p *ListNetworkACLsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListNetworkACLsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
	return
}

func (p *ListNetworkACLsParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
	return
}

func (p *ListNetworkACLsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListNetworkACLsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListNetworkACLsParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *ListNetworkACLsParams) SetProtocol(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["protocol"] = v
	return
}

func (p *ListNetworkACLsParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
	return
}

func (p *ListNetworkACLsParams) SetTraffictype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["traffictype"] = v
	return
}

// You should always use this function to get a new ListNetworkACLsParams instance,
// as then you are sure you have configured all required params
func (s *NetworkACLService) NewListNetworkACLsParams() *ListNetworkACLsParams {
	p := &ListNetworkACLsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkACLService) GetNetworkACLID(keyword string) (string, error) {
	p := &ListNetworkACLsParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword

	l, err := s.ListNetworkACLs(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, keyword, l)
	}
	return l.NetworkACLs[0].Id, nil
}

// Lists all network ACL items
func (s *NetworkACLService) ListNetworkACLs(p *ListNetworkACLsParams) (*ListNetworkACLsResponse, error) {
	resp, err := s.cs.newRequest("listNetworkACLs", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListNetworkACLsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListNetworkACLsResponse struct {
	Count       int           `json:"count"`
	NetworkACLs []*NetworkACL `json:"networkacl"`
}

type NetworkACL struct {
	Startport   string `json:"startport,omitempty"`
	Action      string `json:"action,omitempty"`
	Traffictype string `json:"traffictype,omitempty"`
	Tags        []struct {
		Customer     string `json:"customer,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Account      string `json:"account,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Project      string `json:"project,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Key          string `json:"key,omitempty"`
		Value        string `json:"value,omitempty"`
	} `json:"tags,omitempty"`
	Id       string `json:"id,omitempty"`
	Icmpcode int    `json:"icmpcode,omitempty"`
	Cidrlist string `json:"cidrlist,omitempty"`
	Aclid    string `json:"aclid,omitempty"`
	Number   int    `json:"number,omitempty"`
	Icmptype int    `json:"icmptype,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	State    string `json:"state,omitempty"`
	Endport  string `json:"endport,omitempty"`
}

type CreateNetworkACLListParams struct {
	p map[string]interface{}
}

func (p *CreateNetworkACLListParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	return u
}

func (p *CreateNetworkACLListParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
	return
}

func (p *CreateNetworkACLListParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *CreateNetworkACLListParams) SetVpcid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcid"] = v
	return
}

// You should always use this function to get a new CreateNetworkACLListParams instance,
// as then you are sure you have configured all required params
func (s *NetworkACLService) NewCreateNetworkACLListParams(name string, vpcid string) *CreateNetworkACLListParams {
	p := &CreateNetworkACLListParams{}
	p.p = make(map[string]interface{})
	p.p["name"] = name
	p.p["vpcid"] = vpcid
	return p
}

// Creates a Network ACL for the given VPC
func (s *NetworkACLService) CreateNetworkACLList(p *CreateNetworkACLListParams) (*CreateNetworkACLListResponse, error) {
	resp, err := s.cs.newRequest("createNetworkACLList", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateNetworkACLListResponse
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

		var r CreateNetworkACLListResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type CreateNetworkACLListResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Description string `json:"description,omitempty"`
	Vpcid       string `json:"vpcid,omitempty"`
	Name        string `json:"name,omitempty"`
	Id          string `json:"id,omitempty"`
}

type DeleteNetworkACLListParams struct {
	p map[string]interface{}
}

func (p *DeleteNetworkACLListParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteNetworkACLListParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteNetworkACLListParams instance,
// as then you are sure you have configured all required params
func (s *NetworkACLService) NewDeleteNetworkACLListParams(id string) *DeleteNetworkACLListParams {
	p := &DeleteNetworkACLListParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a Network ACL
func (s *NetworkACLService) DeleteNetworkACLList(p *DeleteNetworkACLListParams) (*DeleteNetworkACLListResponse, error) {
	resp, err := s.cs.newRequest("deleteNetworkACLList", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteNetworkACLListResponse
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

		var r DeleteNetworkACLListResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type DeleteNetworkACLListResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type ReplaceNetworkACLListParams struct {
	p map[string]interface{}
}

func (p *ReplaceNetworkACLListParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["aclid"]; found {
		u.Set("aclid", v.(string))
	}
	if v, found := p.p["gatewayid"]; found {
		u.Set("gatewayid", v.(string))
	}
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	return u
}

func (p *ReplaceNetworkACLListParams) SetAclid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["aclid"] = v
	return
}

func (p *ReplaceNetworkACLListParams) SetGatewayid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gatewayid"] = v
	return
}

func (p *ReplaceNetworkACLListParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
	return
}

// You should always use this function to get a new ReplaceNetworkACLListParams instance,
// as then you are sure you have configured all required params
func (s *NetworkACLService) NewReplaceNetworkACLListParams(aclid string) *ReplaceNetworkACLListParams {
	p := &ReplaceNetworkACLListParams{}
	p.p = make(map[string]interface{})
	p.p["aclid"] = aclid
	return p
}

// Replaces ACL associated with a Network or private gateway
func (s *NetworkACLService) ReplaceNetworkACLList(p *ReplaceNetworkACLListParams) (*ReplaceNetworkACLListResponse, error) {
	resp, err := s.cs.newRequest("replaceNetworkACLList", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ReplaceNetworkACLListResponse
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

		var r ReplaceNetworkACLListResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type ReplaceNetworkACLListResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type ListNetworkACLListsParams struct {
	p map[string]interface{}
}

func (p *ListNetworkACLListsParams) toURLValues() url.Values {
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
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	return u
}

func (p *ListNetworkACLListsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListNetworkACLListsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListNetworkACLListsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListNetworkACLListsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
	return
}

func (p *ListNetworkACLListsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListNetworkACLListsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
	return
}

func (p *ListNetworkACLListsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *ListNetworkACLListsParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
	return
}

func (p *ListNetworkACLListsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListNetworkACLListsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListNetworkACLListsParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *ListNetworkACLListsParams) SetVpcid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcid"] = v
	return
}

// You should always use this function to get a new ListNetworkACLListsParams instance,
// as then you are sure you have configured all required params
func (s *NetworkACLService) NewListNetworkACLListsParams() *ListNetworkACLListsParams {
	p := &ListNetworkACLListsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkACLService) GetNetworkACLListID(name string) (string, error) {
	p := &ListNetworkACLListsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	l, err := s.ListNetworkACLLists(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, name, l)
	}
	return l.NetworkACLLists[0].Id, nil
}

// Lists all network ACLs
func (s *NetworkACLService) ListNetworkACLLists(p *ListNetworkACLListsParams) (*ListNetworkACLListsResponse, error) {
	resp, err := s.cs.newRequest("listNetworkACLLists", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListNetworkACLListsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListNetworkACLListsResponse struct {
	Count           int               `json:"count"`
	NetworkACLLists []*NetworkACLList `json:"networkacllist"`
}

type NetworkACLList struct {
	Name        string `json:"name,omitempty"`
	Vpcid       string `json:"vpcid,omitempty"`
	Id          string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
}
