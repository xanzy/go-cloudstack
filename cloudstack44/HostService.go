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
	"strings"
)

type AddHostParams struct {
	p map[string]interface{}
}

func (p *AddHostParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["allocationstate"]; found {
		u.Set("allocationstate", v.(string))
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := p.p["clustername"]; found {
		u.Set("clustername", v.(string))
	}
	if v, found := p.p["hosttags"]; found {
		vv := strings.Join(v.([]string), ", ")
		u.Set("hosttags", vv)
	}
	if v, found := p.p["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
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

func (p *AddHostParams) SetAllocationstate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["allocationstate"] = v
	return
}

func (p *AddHostParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
	return
}

func (p *AddHostParams) SetClustername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clustername"] = v
	return
}

func (p *AddHostParams) SetHosttags(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hosttags"] = v
	return
}

func (p *AddHostParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
	return
}

func (p *AddHostParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
	return
}

func (p *AddHostParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
	return
}

func (p *AddHostParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
	return
}

func (p *AddHostParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
	return
}

func (p *AddHostParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new AddHostParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewAddHostParams(hypervisor string, password string, podid string, url string, username string, zoneid string) *AddHostParams {
	p := &AddHostParams{}
	p.p = make(map[string]interface{})
	p.p["hypervisor"] = hypervisor
	p.p["password"] = password
	p.p["podid"] = podid
	p.p["url"] = url
	p.p["username"] = username
	p.p["zoneid"] = zoneid
	return p
}

// Adds a new host.
func (s *HostService) AddHost(p *AddHostParams) (*AddHostResponse, error) {
	resp, err := s.cs.newRequest("addHost", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddHostResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type AddHostResponse struct {
	Cpuused                 string `json:"cpuused,omitempty"`
	Memoryused              int    `json:"memoryused,omitempty"`
	Removed                 string `json:"removed,omitempty"`
	Managementserverid      int    `json:"managementserverid,omitempty"`
	Hasenoughcapacity       bool   `json:"hasenoughcapacity,omitempty"`
	Ipaddress               string `json:"ipaddress,omitempty"`
	Clusterid               string `json:"clusterid,omitempty"`
	Zoneid                  string `json:"zoneid,omitempty"`
	Averageload             int    `json:"averageload,omitempty"`
	Version                 string `json:"version,omitempty"`
	Hypervisorversion       string `json:"hypervisorversion,omitempty"`
	Oscategoryname          string `json:"oscategoryname,omitempty"`
	Networkkbsread          int    `json:"networkkbsread,omitempty"`
	Capabilities            string `json:"capabilities,omitempty"`
	Created                 string `json:"created,omitempty"`
	Clustertype             string `json:"clustertype,omitempty"`
	Type                    string `json:"type,omitempty"`
	Disksizeallocated       int    `json:"disksizeallocated,omitempty"`
	Cpuwithoverprovisioning string `json:"cpuwithoverprovisioning,omitempty"`
	Suitableformigration    bool   `json:"suitableformigration,omitempty"`
	Oscategoryid            string `json:"oscategoryid,omitempty"`
	Disconnected            string `json:"disconnected,omitempty"`
	Memorytotal             int    `json:"memorytotal,omitempty"`
	Podname                 string `json:"podname,omitempty"`
	Cpusockets              int    `json:"cpusockets,omitempty"`
	State                   string `json:"state,omitempty"`
	Networkkbswrite         int    `json:"networkkbswrite,omitempty"`
	Islocalstorageactive    bool   `json:"islocalstorageactive,omitempty"`
	Memoryallocated         int    `json:"memoryallocated,omitempty"`
	Disksizetotal           int    `json:"disksizetotal,omitempty"`
	Cpuallocated            string `json:"cpuallocated,omitempty"`
	Cpunumber               int    `json:"cpunumber,omitempty"`
	Resourcestate           string `json:"resourcestate,omitempty"`
	Name                    string `json:"name,omitempty"`
	Podid                   string `json:"podid,omitempty"`
	Zonename                string `json:"zonename,omitempty"`
	Lastpinged              string `json:"lastpinged,omitempty"`
	Cpuspeed                int    `json:"cpuspeed,omitempty"`
	Hahost                  bool   `json:"hahost,omitempty"`
	Hosttags                string `json:"hosttags,omitempty"`
	Clustername             string `json:"clustername,omitempty"`
	Id                      string `json:"id,omitempty"`
	Events                  string `json:"events,omitempty"`
	Hypervisor              string `json:"hypervisor,omitempty"`
}

type ReconnectHostParams struct {
	p map[string]interface{}
}

func (p *ReconnectHostParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *ReconnectHostParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new ReconnectHostParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewReconnectHostParams(id string) *ReconnectHostParams {
	p := &ReconnectHostParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Reconnects a host.
func (s *HostService) ReconnectHost(p *ReconnectHostParams) (*ReconnectHostResponse, error) {
	resp, err := s.cs.newRequest("reconnectHost", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ReconnectHostResponse
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

		var r ReconnectHostResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type ReconnectHostResponse struct {
	JobID                   string `json:"jobid,omitempty"`
	Memoryused              int    `json:"memoryused,omitempty"`
	Clustername             string `json:"clustername,omitempty"`
	Resourcestate           string `json:"resourcestate,omitempty"`
	Networkkbswrite         int    `json:"networkkbswrite,omitempty"`
	Managementserverid      int    `json:"managementserverid,omitempty"`
	Cpuallocated            string `json:"cpuallocated,omitempty"`
	Disksizeallocated       int    `json:"disksizeallocated,omitempty"`
	Hypervisor              string `json:"hypervisor,omitempty"`
	State                   string `json:"state,omitempty"`
	Clustertype             string `json:"clustertype,omitempty"`
	Disksizetotal           int    `json:"disksizetotal,omitempty"`
	Version                 string `json:"version,omitempty"`
	Cpunumber               int    `json:"cpunumber,omitempty"`
	Created                 string `json:"created,omitempty"`
	Podid                   string `json:"podid,omitempty"`
	Type                    string `json:"type,omitempty"`
	Hasenoughcapacity       bool   `json:"hasenoughcapacity,omitempty"`
	Networkkbsread          int    `json:"networkkbsread,omitempty"`
	Events                  string `json:"events,omitempty"`
	Ipaddress               string `json:"ipaddress,omitempty"`
	Cpuspeed                int    `json:"cpuspeed,omitempty"`
	Memoryallocated         int    `json:"memoryallocated,omitempty"`
	Clusterid               string `json:"clusterid,omitempty"`
	Oscategoryid            string `json:"oscategoryid,omitempty"`
	Cpuwithoverprovisioning string `json:"cpuwithoverprovisioning,omitempty"`
	Capabilities            string `json:"capabilities,omitempty"`
	Removed                 string `json:"removed,omitempty"`
	Lastpinged              string `json:"lastpinged,omitempty"`
	Zoneid                  string `json:"zoneid,omitempty"`
	Podname                 string `json:"podname,omitempty"`
	Islocalstorageactive    bool   `json:"islocalstorageactive,omitempty"`
	Memorytotal             int    `json:"memorytotal,omitempty"`
	Suitableformigration    bool   `json:"suitableformigration,omitempty"`
	Oscategoryname          string `json:"oscategoryname,omitempty"`
	Hahost                  bool   `json:"hahost,omitempty"`
	Hosttags                string `json:"hosttags,omitempty"`
	Averageload             int    `json:"averageload,omitempty"`
	Cpuused                 string `json:"cpuused,omitempty"`
	Name                    string `json:"name,omitempty"`
	Disconnected            string `json:"disconnected,omitempty"`
	Hypervisorversion       string `json:"hypervisorversion,omitempty"`
	Zonename                string `json:"zonename,omitempty"`
	Id                      string `json:"id,omitempty"`
	Cpusockets              int    `json:"cpusockets,omitempty"`
}

type UpdateHostParams struct {
	p map[string]interface{}
}

func (p *UpdateHostParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["allocationstate"]; found {
		u.Set("allocationstate", v.(string))
	}
	if v, found := p.p["hosttags"]; found {
		vv := strings.Join(v.([]string), ", ")
		u.Set("hosttags", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["oscategoryid"]; found {
		u.Set("oscategoryid", v.(string))
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	return u
}

func (p *UpdateHostParams) SetAllocationstate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["allocationstate"] = v
	return
}

func (p *UpdateHostParams) SetHosttags(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hosttags"] = v
	return
}

func (p *UpdateHostParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *UpdateHostParams) SetOscategoryid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["oscategoryid"] = v
	return
}

func (p *UpdateHostParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
	return
}

// You should always use this function to get a new UpdateHostParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewUpdateHostParams(id string) *UpdateHostParams {
	p := &UpdateHostParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a host.
func (s *HostService) UpdateHost(p *UpdateHostParams) (*UpdateHostResponse, error) {
	resp, err := s.cs.newRequest("updateHost", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateHostResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type UpdateHostResponse struct {
	Hypervisorversion       string `json:"hypervisorversion,omitempty"`
	Version                 string `json:"version,omitempty"`
	Cpuwithoverprovisioning string `json:"cpuwithoverprovisioning,omitempty"`
	Id                      string `json:"id,omitempty"`
	Cpuused                 string `json:"cpuused,omitempty"`
	Clusterid               string `json:"clusterid,omitempty"`
	Networkkbswrite         int    `json:"networkkbswrite,omitempty"`
	Networkkbsread          int    `json:"networkkbsread,omitempty"`
	Memoryused              int    `json:"memoryused,omitempty"`
	Name                    string `json:"name,omitempty"`
	Clustername             string `json:"clustername,omitempty"`
	Oscategoryid            string `json:"oscategoryid,omitempty"`
	Resourcestate           string `json:"resourcestate,omitempty"`
	Suitableformigration    bool   `json:"suitableformigration,omitempty"`
	Podname                 string `json:"podname,omitempty"`
	Disconnected            string `json:"disconnected,omitempty"`
	Managementserverid      int    `json:"managementserverid,omitempty"`
	Cpusockets              int    `json:"cpusockets,omitempty"`
	Clustertype             string `json:"clustertype,omitempty"`
	Zonename                string `json:"zonename,omitempty"`
	Memorytotal             int    `json:"memorytotal,omitempty"`
	Removed                 string `json:"removed,omitempty"`
	Podid                   string `json:"podid,omitempty"`
	State                   string `json:"state,omitempty"`
	Memoryallocated         int    `json:"memoryallocated,omitempty"`
	Created                 string `json:"created,omitempty"`
	Cpunumber               int    `json:"cpunumber,omitempty"`
	Cpuallocated            string `json:"cpuallocated,omitempty"`
	Hahost                  bool   `json:"hahost,omitempty"`
	Zoneid                  string `json:"zoneid,omitempty"`
	Hosttags                string `json:"hosttags,omitempty"`
	Ipaddress               string `json:"ipaddress,omitempty"`
	Cpuspeed                int    `json:"cpuspeed,omitempty"`
	Events                  string `json:"events,omitempty"`
	Capabilities            string `json:"capabilities,omitempty"`
	Hasenoughcapacity       bool   `json:"hasenoughcapacity,omitempty"`
	Oscategoryname          string `json:"oscategoryname,omitempty"`
	Disksizetotal           int    `json:"disksizetotal,omitempty"`
	Disksizeallocated       int    `json:"disksizeallocated,omitempty"`
	Averageload             int    `json:"averageload,omitempty"`
	Islocalstorageactive    bool   `json:"islocalstorageactive,omitempty"`
	Type                    string `json:"type,omitempty"`
	Hypervisor              string `json:"hypervisor,omitempty"`
	Lastpinged              string `json:"lastpinged,omitempty"`
}

type DeleteHostParams struct {
	p map[string]interface{}
}

func (p *DeleteHostParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["forced"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forced", vv)
	}
	if v, found := p.p["forcedestroylocalstorage"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forcedestroylocalstorage", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteHostParams) SetForced(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forced"] = v
	return
}

func (p *DeleteHostParams) SetForcedestroylocalstorage(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forcedestroylocalstorage"] = v
	return
}

func (p *DeleteHostParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteHostParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewDeleteHostParams(id string) *DeleteHostParams {
	p := &DeleteHostParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a host.
func (s *HostService) DeleteHost(p *DeleteHostParams) (*DeleteHostResponse, error) {
	resp, err := s.cs.newRequest("deleteHost", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteHostResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type DeleteHostResponse struct {
	Success     bool   `json:"success,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
}

type PrepareHostForMaintenanceParams struct {
	p map[string]interface{}
}

func (p *PrepareHostForMaintenanceParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *PrepareHostForMaintenanceParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new PrepareHostForMaintenanceParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewPrepareHostForMaintenanceParams(id string) *PrepareHostForMaintenanceParams {
	p := &PrepareHostForMaintenanceParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Prepares a host for maintenance.
func (s *HostService) PrepareHostForMaintenance(p *PrepareHostForMaintenanceParams) (*PrepareHostForMaintenanceResponse, error) {
	resp, err := s.cs.newRequest("prepareHostForMaintenance", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r PrepareHostForMaintenanceResponse
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

		var r PrepareHostForMaintenanceResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type PrepareHostForMaintenanceResponse struct {
	JobID                   string `json:"jobid,omitempty"`
	Cpuwithoverprovisioning string `json:"cpuwithoverprovisioning,omitempty"`
	Suitableformigration    bool   `json:"suitableformigration,omitempty"`
	Memorytotal             int    `json:"memorytotal,omitempty"`
	Hosttags                string `json:"hosttags,omitempty"`
	State                   string `json:"state,omitempty"`
	Disksizeallocated       int    `json:"disksizeallocated,omitempty"`
	Podname                 string `json:"podname,omitempty"`
	Averageload             int    `json:"averageload,omitempty"`
	Memoryallocated         int    `json:"memoryallocated,omitempty"`
	Clustername             string `json:"clustername,omitempty"`
	Hahost                  bool   `json:"hahost,omitempty"`
	Disconnected            string `json:"disconnected,omitempty"`
	Capabilities            string `json:"capabilities,omitempty"`
	Ipaddress               string `json:"ipaddress,omitempty"`
	Events                  string `json:"events,omitempty"`
	Oscategoryid            string `json:"oscategoryid,omitempty"`
	Cpuused                 string `json:"cpuused,omitempty"`
	Managementserverid      int    `json:"managementserverid,omitempty"`
	Clustertype             string `json:"clustertype,omitempty"`
	Created                 string `json:"created,omitempty"`
	Cpusockets              int    `json:"cpusockets,omitempty"`
	Cpuallocated            string `json:"cpuallocated,omitempty"`
	Version                 string `json:"version,omitempty"`
	Name                    string `json:"name,omitempty"`
	Memoryused              int    `json:"memoryused,omitempty"`
	Cpunumber               int    `json:"cpunumber,omitempty"`
	Lastpinged              string `json:"lastpinged,omitempty"`
	Podid                   string `json:"podid,omitempty"`
	Oscategoryname          string `json:"oscategoryname,omitempty"`
	Zonename                string `json:"zonename,omitempty"`
	Id                      string `json:"id,omitempty"`
	Zoneid                  string `json:"zoneid,omitempty"`
	Cpuspeed                int    `json:"cpuspeed,omitempty"`
	Resourcestate           string `json:"resourcestate,omitempty"`
	Type                    string `json:"type,omitempty"`
	Hypervisorversion       string `json:"hypervisorversion,omitempty"`
	Hypervisor              string `json:"hypervisor,omitempty"`
	Hasenoughcapacity       bool   `json:"hasenoughcapacity,omitempty"`
	Removed                 string `json:"removed,omitempty"`
	Disksizetotal           int    `json:"disksizetotal,omitempty"`
	Networkkbswrite         int    `json:"networkkbswrite,omitempty"`
	Islocalstorageactive    bool   `json:"islocalstorageactive,omitempty"`
	Clusterid               string `json:"clusterid,omitempty"`
	Networkkbsread          int    `json:"networkkbsread,omitempty"`
}

type CancelHostMaintenanceParams struct {
	p map[string]interface{}
}

func (p *CancelHostMaintenanceParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *CancelHostMaintenanceParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new CancelHostMaintenanceParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewCancelHostMaintenanceParams(id string) *CancelHostMaintenanceParams {
	p := &CancelHostMaintenanceParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Cancels host maintenance.
func (s *HostService) CancelHostMaintenance(p *CancelHostMaintenanceParams) (*CancelHostMaintenanceResponse, error) {
	resp, err := s.cs.newRequest("cancelHostMaintenance", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CancelHostMaintenanceResponse
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

		var r CancelHostMaintenanceResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type CancelHostMaintenanceResponse struct {
	JobID                   string `json:"jobid,omitempty"`
	Networkkbsread          int    `json:"networkkbsread,omitempty"`
	Id                      string `json:"id,omitempty"`
	Averageload             int    `json:"averageload,omitempty"`
	Clustertype             string `json:"clustertype,omitempty"`
	Clusterid               string `json:"clusterid,omitempty"`
	Oscategoryname          string `json:"oscategoryname,omitempty"`
	Resourcestate           string `json:"resourcestate,omitempty"`
	Type                    string `json:"type,omitempty"`
	Created                 string `json:"created,omitempty"`
	Oscategoryid            string `json:"oscategoryid,omitempty"`
	Cpuwithoverprovisioning string `json:"cpuwithoverprovisioning,omitempty"`
	Hypervisor              string `json:"hypervisor,omitempty"`
	Managementserverid      int    `json:"managementserverid,omitempty"`
	Memorytotal             int    `json:"memorytotal,omitempty"`
	Zonename                string `json:"zonename,omitempty"`
	Podid                   string `json:"podid,omitempty"`
	Cpusockets              int    `json:"cpusockets,omitempty"`
	Islocalstorageactive    bool   `json:"islocalstorageactive,omitempty"`
	Lastpinged              string `json:"lastpinged,omitempty"`
	Memoryused              int    `json:"memoryused,omitempty"`
	Memoryallocated         int    `json:"memoryallocated,omitempty"`
	State                   string `json:"state,omitempty"`
	Disksizeallocated       int    `json:"disksizeallocated,omitempty"`
	Cpunumber               int    `json:"cpunumber,omitempty"`
	Networkkbswrite         int    `json:"networkkbswrite,omitempty"`
	Cpuused                 string `json:"cpuused,omitempty"`
	Cpuallocated            string `json:"cpuallocated,omitempty"`
	Events                  string `json:"events,omitempty"`
	Ipaddress               string `json:"ipaddress,omitempty"`
	Cpuspeed                int    `json:"cpuspeed,omitempty"`
	Hosttags                string `json:"hosttags,omitempty"`
	Removed                 string `json:"removed,omitempty"`
	Version                 string `json:"version,omitempty"`
	Capabilities            string `json:"capabilities,omitempty"`
	Disconnected            string `json:"disconnected,omitempty"`
	Zoneid                  string `json:"zoneid,omitempty"`
	Hypervisorversion       string `json:"hypervisorversion,omitempty"`
	Name                    string `json:"name,omitempty"`
	Podname                 string `json:"podname,omitempty"`
	Disksizetotal           int    `json:"disksizetotal,omitempty"`
	Clustername             string `json:"clustername,omitempty"`
	Hahost                  bool   `json:"hahost,omitempty"`
	Suitableformigration    bool   `json:"suitableformigration,omitempty"`
	Hasenoughcapacity       bool   `json:"hasenoughcapacity,omitempty"`
}

type ListHostsParams struct {
	p map[string]interface{}
}

func (p *ListHostsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := p.p["details"]; found {
		vv := strings.Join(v.([]string), ", ")
		u.Set("details", vv)
	}
	if v, found := p.p["hahost"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("hahost", vv)
	}
	if v, found := p.p["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
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
	if v, found := p.p["resourcestate"]; found {
		u.Set("resourcestate", v.(string))
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := p.p["type"]; found {
		u.Set("type", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListHostsParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
	return
}

func (p *ListHostsParams) SetDetails(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
	return
}

func (p *ListHostsParams) SetHahost(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hahost"] = v
	return
}

func (p *ListHostsParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
	return
}

func (p *ListHostsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListHostsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListHostsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *ListHostsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListHostsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListHostsParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
	return
}

func (p *ListHostsParams) SetResourcestate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["resourcestate"] = v
	return
}

func (p *ListHostsParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
	return
}

func (p *ListHostsParams) SetType(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostType"] = v
	return
}

func (p *ListHostsParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
	return
}

func (p *ListHostsParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new ListHostsParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewListHostsParams() *ListHostsParams {
	p := &ListHostsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *HostService) GetHostID(name string) (string, error) {
	p := &ListHostsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	l, err := s.ListHosts(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, name, l)
	}
	return l.Hosts[0].Id, nil
}

// Lists hosts.
func (s *HostService) ListHosts(p *ListHostsParams) (*ListHostsResponse, error) {
	resp, err := s.cs.newRequest("listHosts", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListHostsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListHostsResponse struct {
	Count int     `json:"count"`
	Hosts []*Host `json:"host"`
}

type Host struct {
	Managementserverid      int    `json:"managementserverid,omitempty"`
	Islocalstorageactive    bool   `json:"islocalstorageactive,omitempty"`
	Oscategoryname          string `json:"oscategoryname,omitempty"`
	Ipaddress               string `json:"ipaddress,omitempty"`
	Cpuallocated            string `json:"cpuallocated,omitempty"`
	Averageload             int    `json:"averageload,omitempty"`
	Cpusockets              int    `json:"cpusockets,omitempty"`
	Networkkbsread          int    `json:"networkkbsread,omitempty"`
	Memoryallocated         int    `json:"memoryallocated,omitempty"`
	Type                    string `json:"type,omitempty"`
	Hasenoughcapacity       bool   `json:"hasenoughcapacity,omitempty"`
	Lastpinged              string `json:"lastpinged,omitempty"`
	Capabilities            string `json:"capabilities,omitempty"`
	Podname                 string `json:"podname,omitempty"`
	Zoneid                  string `json:"zoneid,omitempty"`
	Memorytotal             int    `json:"memorytotal,omitempty"`
	Cpuwithoverprovisioning string `json:"cpuwithoverprovisioning,omitempty"`
	Removed                 string `json:"removed,omitempty"`
	Disksizeallocated       int    `json:"disksizeallocated,omitempty"`
	State                   string `json:"state,omitempty"`
	Cpuspeed                int    `json:"cpuspeed,omitempty"`
	Hahost                  bool   `json:"hahost,omitempty"`
	Clustername             string `json:"clustername,omitempty"`
	Id                      string `json:"id,omitempty"`
	Memoryused              int    `json:"memoryused,omitempty"`
	Hypervisor              string `json:"hypervisor,omitempty"`
	Hosttags                string `json:"hosttags,omitempty"`
	Podid                   string `json:"podid,omitempty"`
	Cpunumber               int    `json:"cpunumber,omitempty"`
	Hypervisorversion       string `json:"hypervisorversion,omitempty"`
	Suitableformigration    bool   `json:"suitableformigration,omitempty"`
	Cpuused                 string `json:"cpuused,omitempty"`
	Name                    string `json:"name,omitempty"`
	Created                 string `json:"created,omitempty"`
	Resourcestate           string `json:"resourcestate,omitempty"`
	Clustertype             string `json:"clustertype,omitempty"`
	Networkkbswrite         int    `json:"networkkbswrite,omitempty"`
	Zonename                string `json:"zonename,omitempty"`
	Events                  string `json:"events,omitempty"`
	Disksizetotal           int    `json:"disksizetotal,omitempty"`
	Version                 string `json:"version,omitempty"`
	Oscategoryid            string `json:"oscategoryid,omitempty"`
	Disconnected            string `json:"disconnected,omitempty"`
	Clusterid               string `json:"clusterid,omitempty"`
}

type FindHostsForMigrationParams struct {
	p map[string]interface{}
}

func (p *FindHostsForMigrationParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
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
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *FindHostsForMigrationParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *FindHostsForMigrationParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *FindHostsForMigrationParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *FindHostsForMigrationParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
	return
}

// You should always use this function to get a new FindHostsForMigrationParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewFindHostsForMigrationParams(virtualmachineid string) *FindHostsForMigrationParams {
	p := &FindHostsForMigrationParams{}
	p.p = make(map[string]interface{})
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// Find hosts suitable for migrating a virtual machine.
func (s *HostService) FindHostsForMigration(p *FindHostsForMigrationParams) (*FindHostsForMigrationResponse, error) {
	resp, err := s.cs.newRequest("findHostsForMigration", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r FindHostsForMigrationResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type FindHostsForMigrationResponse struct {
	Managementserverid      int    `json:"managementserverid,omitempty"`
	Id                      string `json:"id,omitempty"`
	Cpuallocated            string `json:"cpuallocated,omitempty"`
	Hasenoughcapacity       bool   `json:"hasenoughcapacity,omitempty"`
	Hypervisor              string `json:"hypervisor,omitempty"`
	Cpunumber               int    `json:"cpunumber,omitempty"`
	Zonename                string `json:"zonename,omitempty"`
	Version                 string `json:"version,omitempty"`
	Oscategoryname          string `json:"oscategoryname,omitempty"`
	Disksizeallocated       int    `json:"disksizeallocated,omitempty"`
	Podname                 string `json:"podname,omitempty"`
	Memoryused              int    `json:"memoryused,omitempty"`
	Hypervisorversion       string `json:"hypervisorversion,omitempty"`
	Name                    string `json:"name,omitempty"`
	Hosttags                string `json:"hosttags,omitempty"`
	Suitableformigration    bool   `json:"suitableformigration,omitempty"`
	Islocalstorageactive    bool   `json:"islocalstorageactive,omitempty"`
	Clusterid               string `json:"clusterid,omitempty"`
	Disksizetotal           int    `json:"disksizetotal,omitempty"`
	Removed                 string `json:"removed,omitempty"`
	Averageload             int    `json:"averageload,omitempty"`
	Cpuspeed                int    `json:"cpuspeed,omitempty"`
	Hahost                  bool   `json:"hahost,omitempty"`
	Cpuwithoverprovisioning string `json:"cpuwithoverprovisioning,omitempty"`
	RequiresStorageMotion   bool   `json:"requiresStorageMotion,omitempty"`
	Disconnected            string `json:"disconnected,omitempty"`
	Networkkbswrite         int    `json:"networkkbswrite,omitempty"`
	Zoneid                  string `json:"zoneid,omitempty"`
	Oscategoryid            string `json:"oscategoryid,omitempty"`
	Clustername             string `json:"clustername,omitempty"`
	Cpuused                 string `json:"cpuused,omitempty"`
	Resourcestate           string `json:"resourcestate,omitempty"`
	Networkkbsread          int    `json:"networkkbsread,omitempty"`
	Type                    string `json:"type,omitempty"`
	Lastpinged              string `json:"lastpinged,omitempty"`
	Memorytotal             int    `json:"memorytotal,omitempty"`
	State                   string `json:"state,omitempty"`
	Events                  string `json:"events,omitempty"`
	Capabilities            string `json:"capabilities,omitempty"`
	Podid                   string `json:"podid,omitempty"`
	Created                 string `json:"created,omitempty"`
	Clustertype             string `json:"clustertype,omitempty"`
	Ipaddress               string `json:"ipaddress,omitempty"`
	Memoryallocated         int    `json:"memoryallocated,omitempty"`
}

type AddSecondaryStorageParams struct {
	p map[string]interface{}
}

func (p *AddSecondaryStorageParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *AddSecondaryStorageParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
	return
}

func (p *AddSecondaryStorageParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new AddSecondaryStorageParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewAddSecondaryStorageParams(url string) *AddSecondaryStorageParams {
	p := &AddSecondaryStorageParams{}
	p.p = make(map[string]interface{})
	p.p["url"] = url
	return p
}

// Adds secondary storage.
func (s *HostService) AddSecondaryStorage(p *AddSecondaryStorageParams) (*AddSecondaryStorageResponse, error) {
	resp, err := s.cs.newRequest("addSecondaryStorage", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddSecondaryStorageResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type AddSecondaryStorageResponse struct {
	Details      []string `json:"details,omitempty"`
	Providername string   `json:"providername,omitempty"`
	Zoneid       string   `json:"zoneid,omitempty"`
	Zonename     string   `json:"zonename,omitempty"`
	Id           string   `json:"id,omitempty"`
	Name         string   `json:"name,omitempty"`
	Url          string   `json:"url,omitempty"`
	Protocol     string   `json:"protocol,omitempty"`
	Scope        string   `json:"scope,omitempty"`
}

type UpdateHostPasswordParams struct {
	p map[string]interface{}
}

func (p *UpdateHostPasswordParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := p.p["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (p *UpdateHostPasswordParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
	return
}

func (p *UpdateHostPasswordParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
	return
}

func (p *UpdateHostPasswordParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
	return
}

func (p *UpdateHostPasswordParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
	return
}

// You should always use this function to get a new UpdateHostPasswordParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewUpdateHostPasswordParams(password string, username string) *UpdateHostPasswordParams {
	p := &UpdateHostPasswordParams{}
	p.p = make(map[string]interface{})
	p.p["password"] = password
	p.p["username"] = username
	return p
}

// Update password of a host/pool on management server.
func (s *HostService) UpdateHostPassword(p *UpdateHostPasswordParams) (*UpdateHostPasswordResponse, error) {
	resp, err := s.cs.newRequest("updateHostPassword", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateHostPasswordResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type UpdateHostPasswordResponse struct {
	Success     bool   `json:"success,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
}

type ReleaseHostReservationParams struct {
	p map[string]interface{}
}

func (p *ReleaseHostReservationParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *ReleaseHostReservationParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new ReleaseHostReservationParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewReleaseHostReservationParams(id string) *ReleaseHostReservationParams {
	p := &ReleaseHostReservationParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Releases host reservation.
func (s *HostService) ReleaseHostReservation(p *ReleaseHostReservationParams) (*ReleaseHostReservationResponse, error) {
	resp, err := s.cs.newRequest("releaseHostReservation", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ReleaseHostReservationResponse
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

		var r ReleaseHostReservationResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type ReleaseHostReservationResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Success     bool   `json:"success,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
}

type AddBaremetalHostParams struct {
	p map[string]interface{}
}

func (p *AddBaremetalHostParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["allocationstate"]; found {
		u.Set("allocationstate", v.(string))
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := p.p["clustername"]; found {
		u.Set("clustername", v.(string))
	}
	if v, found := p.p["hosttags"]; found {
		vv := strings.Join(v.([]string), ", ")
		u.Set("hosttags", vv)
	}
	if v, found := p.p["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := p.p["ipaddress"]; found {
		u.Set("ipaddress", v.(string))
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
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

func (p *AddBaremetalHostParams) SetAllocationstate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["allocationstate"] = v
	return
}

func (p *AddBaremetalHostParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
	return
}

func (p *AddBaremetalHostParams) SetClustername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clustername"] = v
	return
}

func (p *AddBaremetalHostParams) SetHosttags(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hosttags"] = v
	return
}

func (p *AddBaremetalHostParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
	return
}

func (p *AddBaremetalHostParams) SetIpaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddress"] = v
	return
}

func (p *AddBaremetalHostParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
	return
}

func (p *AddBaremetalHostParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
	return
}

func (p *AddBaremetalHostParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
	return
}

func (p *AddBaremetalHostParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
	return
}

func (p *AddBaremetalHostParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new AddBaremetalHostParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewAddBaremetalHostParams(hypervisor string, password string, podid string, url string, username string, zoneid string) *AddBaremetalHostParams {
	p := &AddBaremetalHostParams{}
	p.p = make(map[string]interface{})
	p.p["hypervisor"] = hypervisor
	p.p["password"] = password
	p.p["podid"] = podid
	p.p["url"] = url
	p.p["username"] = username
	p.p["zoneid"] = zoneid
	return p
}

// add a baremetal host
func (s *HostService) AddBaremetalHost(p *AddBaremetalHostParams) (*AddBaremetalHostResponse, error) {
	resp, err := s.cs.newRequest("addBaremetalHost", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddBaremetalHostResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type AddBaremetalHostResponse struct {
	Cpunumber               int    `json:"cpunumber,omitempty"`
	Hasenoughcapacity       bool   `json:"hasenoughcapacity,omitempty"`
	Zonename                string `json:"zonename,omitempty"`
	State                   string `json:"state,omitempty"`
	Resourcestate           string `json:"resourcestate,omitempty"`
	Ipaddress               string `json:"ipaddress,omitempty"`
	Lastpinged              string `json:"lastpinged,omitempty"`
	Cpuwithoverprovisioning string `json:"cpuwithoverprovisioning,omitempty"`
	Clusterid               string `json:"clusterid,omitempty"`
	Capabilities            string `json:"capabilities,omitempty"`
	Memoryused              int    `json:"memoryused,omitempty"`
	Hypervisor              string `json:"hypervisor,omitempty"`
	Removed                 string `json:"removed,omitempty"`
	Events                  string `json:"events,omitempty"`
	Clustername             string `json:"clustername,omitempty"`
	Networkkbswrite         int    `json:"networkkbswrite,omitempty"`
	Podid                   string `json:"podid,omitempty"`
	Networkkbsread          int    `json:"networkkbsread,omitempty"`
	Zoneid                  string `json:"zoneid,omitempty"`
	Disksizeallocated       int    `json:"disksizeallocated,omitempty"`
	Type                    string `json:"type,omitempty"`
	Name                    string `json:"name,omitempty"`
	Islocalstorageactive    bool   `json:"islocalstorageactive,omitempty"`
	Id                      string `json:"id,omitempty"`
	Hosttags                string `json:"hosttags,omitempty"`
	Hypervisorversion       string `json:"hypervisorversion,omitempty"`
	Disconnected            string `json:"disconnected,omitempty"`
	Cpuspeed                int    `json:"cpuspeed,omitempty"`
	Managementserverid      int    `json:"managementserverid,omitempty"`
	Version                 string `json:"version,omitempty"`
	Averageload             int    `json:"averageload,omitempty"`
	Cpusockets              int    `json:"cpusockets,omitempty"`
	Cpuused                 string `json:"cpuused,omitempty"`
	Hahost                  bool   `json:"hahost,omitempty"`
	Podname                 string `json:"podname,omitempty"`
	Oscategoryname          string `json:"oscategoryname,omitempty"`
	Memoryallocated         int    `json:"memoryallocated,omitempty"`
	Suitableformigration    bool   `json:"suitableformigration,omitempty"`
	Clustertype             string `json:"clustertype,omitempty"`
	Oscategoryid            string `json:"oscategoryid,omitempty"`
	Cpuallocated            string `json:"cpuallocated,omitempty"`
	Memorytotal             int    `json:"memorytotal,omitempty"`
	Created                 string `json:"created,omitempty"`
	Disksizetotal           int    `json:"disksizetotal,omitempty"`
}

type DedicateHostParams struct {
	p map[string]interface{}
}

func (p *DedicateHostParams) toURLValues() url.Values {
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
	if v, found := p.p["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	return u
}

func (p *DedicateHostParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *DedicateHostParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *DedicateHostParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
	return
}

// You should always use this function to get a new DedicateHostParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewDedicateHostParams(domainid string, hostid string) *DedicateHostParams {
	p := &DedicateHostParams{}
	p.p = make(map[string]interface{})
	p.p["domainid"] = domainid
	p.p["hostid"] = hostid
	return p
}

// Dedicates a host.
func (s *HostService) DedicateHost(p *DedicateHostParams) (*DedicateHostResponse, error) {
	resp, err := s.cs.newRequest("dedicateHost", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DedicateHostResponse
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

		var r DedicateHostResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type DedicateHostResponse struct {
	JobID           string `json:"jobid,omitempty"`
	Hostname        string `json:"hostname,omitempty"`
	Accountid       string `json:"accountid,omitempty"`
	Hostid          string `json:"hostid,omitempty"`
	Id              string `json:"id,omitempty"`
	Affinitygroupid string `json:"affinitygroupid,omitempty"`
	Domainid        string `json:"domainid,omitempty"`
}

type ReleaseDedicatedHostParams struct {
	p map[string]interface{}
}

func (p *ReleaseDedicatedHostParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	return u
}

func (p *ReleaseDedicatedHostParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
	return
}

// You should always use this function to get a new ReleaseDedicatedHostParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewReleaseDedicatedHostParams(hostid string) *ReleaseDedicatedHostParams {
	p := &ReleaseDedicatedHostParams{}
	p.p = make(map[string]interface{})
	p.p["hostid"] = hostid
	return p
}

// Release the dedication for host
func (s *HostService) ReleaseDedicatedHost(p *ReleaseDedicatedHostParams) (*ReleaseDedicatedHostResponse, error) {
	resp, err := s.cs.newRequest("releaseDedicatedHost", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ReleaseDedicatedHostResponse
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

		var r ReleaseDedicatedHostResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type ReleaseDedicatedHostResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type ListDedicatedHostsParams struct {
	p map[string]interface{}
}

func (p *ListDedicatedHostsParams) toURLValues() url.Values {
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
	if v, found := p.p["hostid"]; found {
		u.Set("hostid", v.(string))
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
	return u
}

func (p *ListDedicatedHostsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListDedicatedHostsParams) SetAffinitygroupid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["affinitygroupid"] = v
	return
}

func (p *ListDedicatedHostsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListDedicatedHostsParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
	return
}

func (p *ListDedicatedHostsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListDedicatedHostsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListDedicatedHostsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

// You should always use this function to get a new ListDedicatedHostsParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewListDedicatedHostsParams() *ListDedicatedHostsParams {
	p := &ListDedicatedHostsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *HostService) GetDedicatedHostID(keyword string) (string, error) {
	p := &ListDedicatedHostsParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword

	l, err := s.ListDedicatedHosts(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, keyword, l)
	}
	return l.DedicatedHosts[0].Id, nil
}

// Lists dedicated hosts.
func (s *HostService) ListDedicatedHosts(p *ListDedicatedHostsParams) (*ListDedicatedHostsResponse, error) {
	resp, err := s.cs.newRequest("listDedicatedHosts", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListDedicatedHostsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListDedicatedHostsResponse struct {
	Count          int              `json:"count"`
	DedicatedHosts []*DedicatedHost `json:"dedicatedhost"`
}

type DedicatedHost struct {
	Accountid       string `json:"accountid,omitempty"`
	Hostid          string `json:"hostid,omitempty"`
	Hostname        string `json:"hostname,omitempty"`
	Id              string `json:"id,omitempty"`
	Affinitygroupid string `json:"affinitygroupid,omitempty"`
	Domainid        string `json:"domainid,omitempty"`
}
