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

type CreateDomainParams struct {
	p map[string]interface{}
}

func (p *CreateDomainParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["networkdomain"]; found {
		u.Set("networkdomain", v.(string))
	}
	if v, found := p.p["parentdomainid"]; found {
		u.Set("parentdomainid", v.(string))
	}
	return u
}

func (p *CreateDomainParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *CreateDomainParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *CreateDomainParams) SetNetworkdomain(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkdomain"] = v
	return
}

func (p *CreateDomainParams) SetParentdomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["parentdomainid"] = v
	return
}

// You should always use this function to get a new CreateDomainParams instance,
// as then you are sure you have configured all required params
func (s *DomainService) NewCreateDomainParams(name string) *CreateDomainParams {
	p := &CreateDomainParams{}
	p.p = make(map[string]interface{})
	p.p["name"] = name
	return p
}

// Creates a domain
func (s *DomainService) CreateDomain(p *CreateDomainParams) (*CreateDomainResponse, error) {
	resp, err := s.cs.newRequest("createDomain", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateDomainResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type CreateDomainResponse struct {
	Networkdomain    string `json:"networkdomain,omitempty"`
	Haschild         bool   `json:"haschild,omitempty"`
	Path             string `json:"path,omitempty"`
	Parentdomainid   string `json:"parentdomainid,omitempty"`
	Name             string `json:"name,omitempty"`
	Id               string `json:"id,omitempty"`
	Level            int    `json:"level,omitempty"`
	Parentdomainname string `json:"parentdomainname,omitempty"`
}

type UpdateDomainParams struct {
	p map[string]interface{}
}

func (p *UpdateDomainParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["networkdomain"]; found {
		u.Set("networkdomain", v.(string))
	}
	return u
}

func (p *UpdateDomainParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *UpdateDomainParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *UpdateDomainParams) SetNetworkdomain(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkdomain"] = v
	return
}

// You should always use this function to get a new UpdateDomainParams instance,
// as then you are sure you have configured all required params
func (s *DomainService) NewUpdateDomainParams(id string) *UpdateDomainParams {
	p := &UpdateDomainParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a domain with a new name
func (s *DomainService) UpdateDomain(p *UpdateDomainParams) (*UpdateDomainResponse, error) {
	resp, err := s.cs.newRequest("updateDomain", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateDomainResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type UpdateDomainResponse struct {
	Haschild         bool   `json:"haschild,omitempty"`
	Parentdomainname string `json:"parentdomainname,omitempty"`
	Name             string `json:"name,omitempty"`
	Path             string `json:"path,omitempty"`
	Networkdomain    string `json:"networkdomain,omitempty"`
	Id               string `json:"id,omitempty"`
	Parentdomainid   string `json:"parentdomainid,omitempty"`
	Level            int    `json:"level,omitempty"`
}

type DeleteDomainParams struct {
	p map[string]interface{}
}

func (p *DeleteDomainParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["cleanup"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("cleanup", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteDomainParams) SetCleanup(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cleanup"] = v
	return
}

func (p *DeleteDomainParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteDomainParams instance,
// as then you are sure you have configured all required params
func (s *DomainService) NewDeleteDomainParams(id string) *DeleteDomainParams {
	p := &DeleteDomainParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a specified domain
func (s *DomainService) DeleteDomain(p *DeleteDomainParams) (*DeleteDomainResponse, error) {
	resp, err := s.cs.newRequest("deleteDomain", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteDomainResponse
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

		var r DeleteDomainResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type DeleteDomainResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type ListDomainsParams struct {
	p map[string]interface{}
}

func (p *ListDomainsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["level"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("level", vv)
	}
	if v, found := p.p["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
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

func (p *ListDomainsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListDomainsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListDomainsParams) SetLevel(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["level"] = v
	return
}

func (p *ListDomainsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
	return
}

func (p *ListDomainsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *ListDomainsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListDomainsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

// You should always use this function to get a new ListDomainsParams instance,
// as then you are sure you have configured all required params
func (s *DomainService) NewListDomainsParams() *ListDomainsParams {
	p := &ListDomainsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *DomainService) GetDomainID(name string) (string, error) {
	p := &ListDomainsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	l, err := s.ListDomains(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, name, l)
	}
	return l.Domains[0].Id, nil
}

// Lists domains and provides detailed information for listed domains
func (s *DomainService) ListDomains(p *ListDomainsParams) (*ListDomainsResponse, error) {
	resp, err := s.cs.newRequest("listDomains", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListDomainsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListDomainsResponse struct {
	Count   int       `json:"count"`
	Domains []*Domain `json:"domain"`
}

type Domain struct {
	Path             string `json:"path,omitempty"`
	Name             string `json:"name,omitempty"`
	Level            int    `json:"level,omitempty"`
	Parentdomainname string `json:"parentdomainname,omitempty"`
	Parentdomainid   string `json:"parentdomainid,omitempty"`
	Id               string `json:"id,omitempty"`
	Networkdomain    string `json:"networkdomain,omitempty"`
	Haschild         bool   `json:"haschild,omitempty"`
}

type ListDomainChildrenParams struct {
	p map[string]interface{}
}

func (p *ListDomainChildrenParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
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

func (p *ListDomainChildrenParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListDomainChildrenParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
	return
}

func (p *ListDomainChildrenParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListDomainChildrenParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
	return
}

func (p *ListDomainChildrenParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *ListDomainChildrenParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListDomainChildrenParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

// You should always use this function to get a new ListDomainChildrenParams instance,
// as then you are sure you have configured all required params
func (s *DomainService) NewListDomainChildrenParams() *ListDomainChildrenParams {
	p := &ListDomainChildrenParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *DomainService) GetDomainChildrenID(name string) (string, error) {
	p := &ListDomainChildrenParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	l, err := s.ListDomainChildren(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, name, l)
	}
	return l.DomainChildren[0].Id, nil
}

// Lists all children domains belonging to a specified domain
func (s *DomainService) ListDomainChildren(p *ListDomainChildrenParams) (*ListDomainChildrenResponse, error) {
	resp, err := s.cs.newRequest("listDomainChildren", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListDomainChildrenResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListDomainChildrenResponse struct {
	Count          int               `json:"count"`
	DomainChildren []*DomainChildren `json:"domainchildren"`
}

type DomainChildren struct {
	Parentdomainid   string `json:"parentdomainid,omitempty"`
	Name             string `json:"name,omitempty"`
	Level            int    `json:"level,omitempty"`
	Path             string `json:"path,omitempty"`
	Networkdomain    string `json:"networkdomain,omitempty"`
	Haschild         bool   `json:"haschild,omitempty"`
	Id               string `json:"id,omitempty"`
	Parentdomainname string `json:"parentdomainname,omitempty"`
}
