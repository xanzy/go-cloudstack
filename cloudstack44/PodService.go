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

package cloudstack44

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type CreatePodParams struct {
	p map[string]interface{}
}

func (p *CreatePodParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["allocationstate"]; found {
		u.Set("allocationstate", v.(string))
	}
	if v, found := p.p["endip"]; found {
		u.Set("endip", v.(string))
	}
	if v, found := p.p["gateway"]; found {
		u.Set("gateway", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["netmask"]; found {
		u.Set("netmask", v.(string))
	}
	if v, found := p.p["startip"]; found {
		u.Set("startip", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *CreatePodParams) SetAllocationstate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["allocationstate"] = v
	return
}

func (p *CreatePodParams) SetEndip(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["endip"] = v
	return
}

func (p *CreatePodParams) SetGateway(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gateway"] = v
	return
}

func (p *CreatePodParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *CreatePodParams) SetNetmask(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["netmask"] = v
	return
}

func (p *CreatePodParams) SetStartip(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startip"] = v
	return
}

func (p *CreatePodParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new CreatePodParams instance,
// as then you are sure you have configured all required params
func (s *PodService) NewCreatePodParams(gateway string, name string, netmask string, startip string, zoneid string) *CreatePodParams {
	p := &CreatePodParams{}
	p.p = make(map[string]interface{})
	p.p["gateway"] = gateway
	p.p["name"] = name
	p.p["netmask"] = netmask
	p.p["startip"] = startip
	p.p["zoneid"] = zoneid
	return p
}

// Creates a new Pod.
func (s *PodService) CreatePod(p *CreatePodParams) (*CreatePodResponse, error) {
	resp, err := s.cs.newRequest("createPod", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreatePodResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type CreatePodResponse struct {
	Name            string `json:"name,omitempty"`
	Netmask         string `json:"netmask,omitempty"`
	Gateway         string `json:"gateway,omitempty"`
	Allocationstate string `json:"allocationstate,omitempty"`
	Endip           string `json:"endip,omitempty"`
	Id              string `json:"id,omitempty"`
	Zoneid          string `json:"zoneid,omitempty"`
	Capacity        []struct {
		Capacityused  int    `json:"capacityused,omitempty"`
		Zoneid        string `json:"zoneid,omitempty"`
		Podid         string `json:"podid,omitempty"`
		Type          int    `json:"type,omitempty"`
		Zonename      string `json:"zonename,omitempty"`
		Podname       string `json:"podname,omitempty"`
		Clusterid     string `json:"clusterid,omitempty"`
		Capacitytotal int    `json:"capacitytotal,omitempty"`
		Clustername   string `json:"clustername,omitempty"`
		Percentused   string `json:"percentused,omitempty"`
	} `json:"capacity,omitempty"`
	Zonename string `json:"zonename,omitempty"`
	Startip  string `json:"startip,omitempty"`
}

type UpdatePodParams struct {
	p map[string]interface{}
}

func (p *UpdatePodParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["allocationstate"]; found {
		u.Set("allocationstate", v.(string))
	}
	if v, found := p.p["endip"]; found {
		u.Set("endip", v.(string))
	}
	if v, found := p.p["gateway"]; found {
		u.Set("gateway", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["netmask"]; found {
		u.Set("netmask", v.(string))
	}
	if v, found := p.p["startip"]; found {
		u.Set("startip", v.(string))
	}
	return u
}

func (p *UpdatePodParams) SetAllocationstate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["allocationstate"] = v
	return
}

func (p *UpdatePodParams) SetEndip(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["endip"] = v
	return
}

func (p *UpdatePodParams) SetGateway(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gateway"] = v
	return
}

func (p *UpdatePodParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *UpdatePodParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *UpdatePodParams) SetNetmask(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["netmask"] = v
	return
}

func (p *UpdatePodParams) SetStartip(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startip"] = v
	return
}

// You should always use this function to get a new UpdatePodParams instance,
// as then you are sure you have configured all required params
func (s *PodService) NewUpdatePodParams(id string) *UpdatePodParams {
	p := &UpdatePodParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a Pod.
func (s *PodService) UpdatePod(p *UpdatePodParams) (*UpdatePodResponse, error) {
	resp, err := s.cs.newRequest("updatePod", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdatePodResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type UpdatePodResponse struct {
	Zoneid          string `json:"zoneid,omitempty"`
	Id              string `json:"id,omitempty"`
	Netmask         string `json:"netmask,omitempty"`
	Name            string `json:"name,omitempty"`
	Endip           string `json:"endip,omitempty"`
	Startip         string `json:"startip,omitempty"`
	Gateway         string `json:"gateway,omitempty"`
	Zonename        string `json:"zonename,omitempty"`
	Allocationstate string `json:"allocationstate,omitempty"`
	Capacity        []struct {
		Capacitytotal int    `json:"capacitytotal,omitempty"`
		Capacityused  int    `json:"capacityused,omitempty"`
		Clustername   string `json:"clustername,omitempty"`
		Type          int    `json:"type,omitempty"`
		Zoneid        string `json:"zoneid,omitempty"`
		Clusterid     string `json:"clusterid,omitempty"`
		Podname       string `json:"podname,omitempty"`
		Zonename      string `json:"zonename,omitempty"`
		Percentused   string `json:"percentused,omitempty"`
		Podid         string `json:"podid,omitempty"`
	} `json:"capacity,omitempty"`
}

type DeletePodParams struct {
	p map[string]interface{}
}

func (p *DeletePodParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeletePodParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeletePodParams instance,
// as then you are sure you have configured all required params
func (s *PodService) NewDeletePodParams(id string) *DeletePodParams {
	p := &DeletePodParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a Pod.
func (s *PodService) DeletePod(p *DeletePodParams) (*DeletePodResponse, error) {
	resp, err := s.cs.newRequest("deletePod", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeletePodResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type DeletePodResponse struct {
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type ListPodsParams struct {
	p map[string]interface{}
}

func (p *ListPodsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["allocationstate"]; found {
		u.Set("allocationstate", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["showcapacities"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("showcapacities", vv)
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListPodsParams) SetAllocationstate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["allocationstate"] = v
	return
}

func (p *ListPodsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListPodsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListPodsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *ListPodsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListPodsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListPodsParams) SetShowcapacities(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["showcapacities"] = v
	return
}

func (p *ListPodsParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new ListPodsParams instance,
// as then you are sure you have configured all required params
func (s *PodService) NewListPodsParams() *ListPodsParams {
	p := &ListPodsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *PodService) GetPodID(name string) (string, error) {
	p := &ListPodsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	l, err := s.ListPods(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, name, l)
	}
	return l.Pods[0].Id, nil
}

// Lists all Pods.
func (s *PodService) ListPods(p *ListPodsParams) (*ListPodsResponse, error) {
	resp, err := s.cs.newRequest("listPods", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListPodsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListPodsResponse struct {
	Count int    `json:"count"`
	Pods  []*Pod `json:"pod"`
}

type Pod struct {
	Zoneid   string `json:"zoneid,omitempty"`
	Capacity []struct {
		Clustername   string `json:"clustername,omitempty"`
		Zonename      string `json:"zonename,omitempty"`
		Percentused   string `json:"percentused,omitempty"`
		Capacityused  int    `json:"capacityused,omitempty"`
		Podname       string `json:"podname,omitempty"`
		Type          int    `json:"type,omitempty"`
		Clusterid     string `json:"clusterid,omitempty"`
		Zoneid        string `json:"zoneid,omitempty"`
		Podid         string `json:"podid,omitempty"`
		Capacitytotal int    `json:"capacitytotal,omitempty"`
	} `json:"capacity,omitempty"`
	Gateway         string `json:"gateway,omitempty"`
	Allocationstate string `json:"allocationstate,omitempty"`
	Zonename        string `json:"zonename,omitempty"`
	Id              string `json:"id,omitempty"`
	Netmask         string `json:"netmask,omitempty"`
	Endip           string `json:"endip,omitempty"`
	Startip         string `json:"startip,omitempty"`
	Name            string `json:"name,omitempty"`
}

type DedicatePodParams struct {
	p map[string]interface{}
}

func (p *DedicatePodParams) toURLValues() url.Values {
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
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	return u
}

func (p *DedicatePodParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *DedicatePodParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *DedicatePodParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
	return
}

// You should always use this function to get a new DedicatePodParams instance,
// as then you are sure you have configured all required params
func (s *PodService) NewDedicatePodParams(domainid string, podid string) *DedicatePodParams {
	p := &DedicatePodParams{}
	p.p = make(map[string]interface{})
	p.p["domainid"] = domainid
	p.p["podid"] = podid
	return p
}

// Dedicates a Pod.
func (s *PodService) DedicatePod(p *DedicatePodParams) (*DedicatePodResponse, error) {
	resp, err := s.cs.newRequest("dedicatePod", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DedicatePodResponse
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

		var r DedicatePodResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type DedicatePodResponse struct {
	JobID           string `json:"jobid,omitempty"`
	Affinitygroupid string `json:"affinitygroupid,omitempty"`
	Accountid       string `json:"accountid,omitempty"`
	Id              string `json:"id,omitempty"`
	Podid           string `json:"podid,omitempty"`
	Domainid        string `json:"domainid,omitempty"`
	Podname         string `json:"podname,omitempty"`
}

type ReleaseDedicatedPodParams struct {
	p map[string]interface{}
}

func (p *ReleaseDedicatedPodParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	return u
}

func (p *ReleaseDedicatedPodParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
	return
}

// You should always use this function to get a new ReleaseDedicatedPodParams instance,
// as then you are sure you have configured all required params
func (s *PodService) NewReleaseDedicatedPodParams(podid string) *ReleaseDedicatedPodParams {
	p := &ReleaseDedicatedPodParams{}
	p.p = make(map[string]interface{})
	p.p["podid"] = podid
	return p
}

// Release the dedication for the pod
func (s *PodService) ReleaseDedicatedPod(p *ReleaseDedicatedPodParams) (*ReleaseDedicatedPodResponse, error) {
	resp, err := s.cs.newRequest("releaseDedicatedPod", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ReleaseDedicatedPodResponse
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

		var r ReleaseDedicatedPodResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type ReleaseDedicatedPodResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type ListDedicatedPodsParams struct {
	p map[string]interface{}
}

func (p *ListDedicatedPodsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["affinitygroupid"]; found {
		u.Set("affinitygroupid", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
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
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	return u
}

func (p *ListDedicatedPodsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListDedicatedPodsParams) SetAffinitygroupid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["affinitygroupid"] = v
	return
}

func (p *ListDedicatedPodsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListDedicatedPodsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListDedicatedPodsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListDedicatedPodsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListDedicatedPodsParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
	return
}

// You should always use this function to get a new ListDedicatedPodsParams instance,
// as then you are sure you have configured all required params
func (s *PodService) NewListDedicatedPodsParams() *ListDedicatedPodsParams {
	p := &ListDedicatedPodsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *PodService) GetDedicatedPodID(keyword string) (string, error) {
	p := &ListDedicatedPodsParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword

	l, err := s.ListDedicatedPods(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, keyword, l)
	}
	return l.DedicatedPods[0].Id, nil
}

// Lists dedicated pods.
func (s *PodService) ListDedicatedPods(p *ListDedicatedPodsParams) (*ListDedicatedPodsResponse, error) {
	resp, err := s.cs.newRequest("listDedicatedPods", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListDedicatedPodsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListDedicatedPodsResponse struct {
	Count         int             `json:"count"`
	DedicatedPods []*DedicatedPod `json:"dedicatedpod"`
}

type DedicatedPod struct {
	Accountid       string `json:"accountid,omitempty"`
	Id              string `json:"id,omitempty"`
	Domainid        string `json:"domainid,omitempty"`
	Podid           string `json:"podid,omitempty"`
	Affinitygroupid string `json:"affinitygroupid,omitempty"`
	Podname         string `json:"podname,omitempty"`
}
