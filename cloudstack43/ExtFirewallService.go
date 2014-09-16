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

type AddExternalFirewallParams struct {
	p map[string]interface{}
}

func (p *AddExternalFirewallParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *AddExternalFirewallParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
	return
}

func (p *AddExternalFirewallParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
	return
}

func (p *AddExternalFirewallParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
	return
}

func (p *AddExternalFirewallParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new AddExternalFirewallParams instance,
// as then you are sure you have configured all required params
func (s *ExtFirewallService) NewAddExternalFirewallParams(password string, url string, username string, zoneid string) *AddExternalFirewallParams {
	p := &AddExternalFirewallParams{}
	p.p = make(map[string]interface{})
	p.p["password"] = password
	p.p["url"] = url
	p.p["username"] = username
	p.p["zoneid"] = zoneid
	return p
}

// Adds an external firewall appliance
func (s *ExtFirewallService) AddExternalFirewall(p *AddExternalFirewallParams) (*AddExternalFirewallResponse, error) {
	resp, err := s.cs.newRequest("addExternalFirewall", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddExternalFirewallResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type AddExternalFirewallResponse struct {
	Timeout          string `json:"timeout,omitempty"`
	Username         string `json:"username,omitempty"`
	Numretries       string `json:"numretries,omitempty"`
	Ipaddress        string `json:"ipaddress,omitempty"`
	Zoneid           string `json:"zoneid,omitempty"`
	Privateinterface string `json:"privateinterface,omitempty"`
	Privatezone      string `json:"privatezone,omitempty"`
	Usageinterface   string `json:"usageinterface,omitempty"`
	Publicinterface  string `json:"publicinterface,omitempty"`
	Id               string `json:"id,omitempty"`
	Publiczone       string `json:"publiczone,omitempty"`
}

type DeleteExternalFirewallParams struct {
	p map[string]interface{}
}

func (p *DeleteExternalFirewallParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteExternalFirewallParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteExternalFirewallParams instance,
// as then you are sure you have configured all required params
func (s *ExtFirewallService) NewDeleteExternalFirewallParams(id string) *DeleteExternalFirewallParams {
	p := &DeleteExternalFirewallParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes an external firewall appliance.
func (s *ExtFirewallService) DeleteExternalFirewall(p *DeleteExternalFirewallParams) (*DeleteExternalFirewallResponse, error) {
	resp, err := s.cs.newRequest("deleteExternalFirewall", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteExternalFirewallResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type DeleteExternalFirewallResponse struct {
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type ListExternalFirewallsParams struct {
	p map[string]interface{}
}

func (p *ListExternalFirewallsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListExternalFirewallsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListExternalFirewallsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListExternalFirewallsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListExternalFirewallsParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new ListExternalFirewallsParams instance,
// as then you are sure you have configured all required params
func (s *ExtFirewallService) NewListExternalFirewallsParams(zoneid string) *ListExternalFirewallsParams {
	p := &ListExternalFirewallsParams{}
	p.p = make(map[string]interface{})
	p.p["zoneid"] = zoneid
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ExtFirewallService) GetExternalFirewallID(keyword string, zoneid string) (string, error) {
	p := &ListExternalFirewallsParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword
	p.p["zoneid"] = zoneid

	l, err := s.ListExternalFirewalls(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, keyword, l)
	}
	return l.ExternalFirewalls[0].Id, nil
}

// List external firewall appliances.
func (s *ExtFirewallService) ListExternalFirewalls(p *ListExternalFirewallsParams) (*ListExternalFirewallsResponse, error) {
	resp, err := s.cs.newRequest("listExternalFirewalls", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListExternalFirewallsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListExternalFirewallsResponse struct {
	Count             int                 `json:"count"`
	ExternalFirewalls []*ExternalFirewall `json:"externalfirewall"`
}

type ExternalFirewall struct {
	Privatezone      string `json:"privatezone,omitempty"`
	Zoneid           string `json:"zoneid,omitempty"`
	Publicinterface  string `json:"publicinterface,omitempty"`
	Numretries       string `json:"numretries,omitempty"`
	Publiczone       string `json:"publiczone,omitempty"`
	Usageinterface   string `json:"usageinterface,omitempty"`
	Privateinterface string `json:"privateinterface,omitempty"`
	Ipaddress        string `json:"ipaddress,omitempty"`
	Username         string `json:"username,omitempty"`
	Timeout          string `json:"timeout,omitempty"`
	Id               string `json:"id,omitempty"`
}
