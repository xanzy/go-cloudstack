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
	Ipaddress               string `json:"ipaddress,omitempty"`
	Hypervisor              string `json:"hypervisor,omitempty"`
	Managementserverid      int    `json:"managementserverid,omitempty"`
	Podname                 string `json:"podname,omitempty"`
	State                   string `json:"state,omitempty"`
	Clustername             string `json:"clustername,omitempty"`
	Disksizeallocated       int    `json:"disksizeallocated,omitempty"`
	Disksizetotal           int    `json:"disksizetotal,omitempty"`
	Memorytotal             int    `json:"memorytotal,omitempty"`
	Cpunumber               int    `json:"cpunumber,omitempty"`
	Hahost                  bool   `json:"hahost,omitempty"`
	Networkkbsread          int    `json:"networkkbsread,omitempty"`
	Name                    string `json:"name,omitempty"`
	Hosttags                string `json:"hosttags,omitempty"`
	Hasenoughcapacity       bool   `json:"hasenoughcapacity,omitempty"`
	Memoryused              int    `json:"memoryused,omitempty"`
	Resourcestate           string `json:"resourcestate,omitempty"`
	Hypervisorversion       string `json:"hypervisorversion,omitempty"`
	Removed                 string `json:"removed,omitempty"`
	Capabilities            string `json:"capabilities,omitempty"`
	Memoryallocated         int    `json:"memoryallocated,omitempty"`
	Oscategoryname          string `json:"oscategoryname,omitempty"`
	Podid                   string `json:"podid,omitempty"`
	Averageload             int    `json:"averageload,omitempty"`
	Islocalstorageactive    bool   `json:"islocalstorageactive,omitempty"`
	Clusterid               string `json:"clusterid,omitempty"`
	Cpuused                 string `json:"cpuused,omitempty"`
	Cpuwithoverprovisioning string `json:"cpuwithoverprovisioning,omitempty"`
	Oscategoryid            string `json:"oscategoryid,omitempty"`
	Suitableformigration    bool   `json:"suitableformigration,omitempty"`
	Created                 string `json:"created,omitempty"`
	Type                    string `json:"type,omitempty"`
	Cpuspeed                int    `json:"cpuspeed,omitempty"`
	Events                  string `json:"events,omitempty"`
	Clustertype             string `json:"clustertype,omitempty"`
	Version                 string `json:"version,omitempty"`
	Zoneid                  string `json:"zoneid,omitempty"`
	Disconnected            string `json:"disconnected,omitempty"`
	Zonename                string `json:"zonename,omitempty"`
	Cpuallocated            string `json:"cpuallocated,omitempty"`
	Cpusockets              int    `json:"cpusockets,omitempty"`
	Networkkbswrite         int    `json:"networkkbswrite,omitempty"`
	Lastpinged              string `json:"lastpinged,omitempty"`
	Id                      string `json:"id,omitempty"`
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
	State                   string `json:"state,omitempty"`
	Hypervisor              string `json:"hypervisor,omitempty"`
	Type                    string `json:"type,omitempty"`
	Memoryallocated         int    `json:"memoryallocated,omitempty"`
	Name                    string `json:"name,omitempty"`
	Podid                   string `json:"podid,omitempty"`
	Suitableformigration    bool   `json:"suitableformigration,omitempty"`
	Capabilities            string `json:"capabilities,omitempty"`
	Memoryused              int    `json:"memoryused,omitempty"`
	Networkkbsread          int    `json:"networkkbsread,omitempty"`
	Managementserverid      int    `json:"managementserverid,omitempty"`
	Zonename                string `json:"zonename,omitempty"`
	Hahost                  bool   `json:"hahost,omitempty"`
	Events                  string `json:"events,omitempty"`
	Clustername             string `json:"clustername,omitempty"`
	Disksizetotal           int    `json:"disksizetotal,omitempty"`
	Podname                 string `json:"podname,omitempty"`
	Cpunumber               int    `json:"cpunumber,omitempty"`
	Resourcestate           string `json:"resourcestate,omitempty"`
	Islocalstorageactive    bool   `json:"islocalstorageactive,omitempty"`
	Ipaddress               string `json:"ipaddress,omitempty"`
	Lastpinged              string `json:"lastpinged,omitempty"`
	Created                 string `json:"created,omitempty"`
	Cpuallocated            string `json:"cpuallocated,omitempty"`
	Cpuspeed                int    `json:"cpuspeed,omitempty"`
	Zoneid                  string `json:"zoneid,omitempty"`
	Hasenoughcapacity       bool   `json:"hasenoughcapacity,omitempty"`
	Cpuwithoverprovisioning string `json:"cpuwithoverprovisioning,omitempty"`
	Networkkbswrite         int    `json:"networkkbswrite,omitempty"`
	Clustertype             string `json:"clustertype,omitempty"`
	Averageload             int    `json:"averageload,omitempty"`
	Id                      string `json:"id,omitempty"`
	Cpuused                 string `json:"cpuused,omitempty"`
	Memorytotal             int    `json:"memorytotal,omitempty"`
	Removed                 string `json:"removed,omitempty"`
	Version                 string `json:"version,omitempty"`
	Disconnected            string `json:"disconnected,omitempty"`
	Hosttags                string `json:"hosttags,omitempty"`
	Oscategoryid            string `json:"oscategoryid,omitempty"`
	Clusterid               string `json:"clusterid,omitempty"`
	Hypervisorversion       string `json:"hypervisorversion,omitempty"`
	Cpusockets              int    `json:"cpusockets,omitempty"`
	Disksizeallocated       int    `json:"disksizeallocated,omitempty"`
	Oscategoryname          string `json:"oscategoryname,omitempty"`
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
	Disksizeallocated       int    `json:"disksizeallocated,omitempty"`
	Name                    string `json:"name,omitempty"`
	Hahost                  bool   `json:"hahost,omitempty"`
	Managementserverid      int    `json:"managementserverid,omitempty"`
	Memorytotal             int    `json:"memorytotal,omitempty"`
	Disksizetotal           int    `json:"disksizetotal,omitempty"`
	Version                 string `json:"version,omitempty"`
	Zoneid                  string `json:"zoneid,omitempty"`
	Memoryused              int    `json:"memoryused,omitempty"`
	Networkkbswrite         int    `json:"networkkbswrite,omitempty"`
	Lastpinged              string `json:"lastpinged,omitempty"`
	State                   string `json:"state,omitempty"`
	Cpunumber               int    `json:"cpunumber,omitempty"`
	Hypervisorversion       string `json:"hypervisorversion,omitempty"`
	Id                      string `json:"id,omitempty"`
	Networkkbsread          int    `json:"networkkbsread,omitempty"`
	Cpuallocated            string `json:"cpuallocated,omitempty"`
	Capabilities            string `json:"capabilities,omitempty"`
	Zonename                string `json:"zonename,omitempty"`
	Ipaddress               string `json:"ipaddress,omitempty"`
	Cpuused                 string `json:"cpuused,omitempty"`
	Events                  string `json:"events,omitempty"`
	Podname                 string `json:"podname,omitempty"`
	Disconnected            string `json:"disconnected,omitempty"`
	Oscategoryname          string `json:"oscategoryname,omitempty"`
	Cpuwithoverprovisioning string `json:"cpuwithoverprovisioning,omitempty"`
	Cpuspeed                int    `json:"cpuspeed,omitempty"`
	Memoryallocated         int    `json:"memoryallocated,omitempty"`
	Cpusockets              int    `json:"cpusockets,omitempty"`
	Clustertype             string `json:"clustertype,omitempty"`
	Podid                   string `json:"podid,omitempty"`
	Clusterid               string `json:"clusterid,omitempty"`
	Resourcestate           string `json:"resourcestate,omitempty"`
	Islocalstorageactive    bool   `json:"islocalstorageactive,omitempty"`
	Hypervisor              string `json:"hypervisor,omitempty"`
	Averageload             int    `json:"averageload,omitempty"`
	Hasenoughcapacity       bool   `json:"hasenoughcapacity,omitempty"`
	Removed                 string `json:"removed,omitempty"`
	Type                    string `json:"type,omitempty"`
	Clustername             string `json:"clustername,omitempty"`
	Suitableformigration    bool   `json:"suitableformigration,omitempty"`
	Hosttags                string `json:"hosttags,omitempty"`
	Created                 string `json:"created,omitempty"`
	Oscategoryid            string `json:"oscategoryid,omitempty"`
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
	Hosttags                string `json:"hosttags,omitempty"`
	Disconnected            string `json:"disconnected,omitempty"`
	Podname                 string `json:"podname,omitempty"`
	Suitableformigration    bool   `json:"suitableformigration,omitempty"`
	Networkkbswrite         int    `json:"networkkbswrite,omitempty"`
	Hypervisorversion       string `json:"hypervisorversion,omitempty"`
	Cpuused                 string `json:"cpuused,omitempty"`
	Averageload             int    `json:"averageload,omitempty"`
	Cpuwithoverprovisioning string `json:"cpuwithoverprovisioning,omitempty"`
	Cpunumber               int    `json:"cpunumber,omitempty"`
	Capabilities            string `json:"capabilities,omitempty"`
	Memoryallocated         int    `json:"memoryallocated,omitempty"`
	Removed                 string `json:"removed,omitempty"`
	Lastpinged              string `json:"lastpinged,omitempty"`
	Memoryused              int    `json:"memoryused,omitempty"`
	Podid                   string `json:"podid,omitempty"`
	Resourcestate           string `json:"resourcestate,omitempty"`
	Clusterid               string `json:"clusterid,omitempty"`
	Cpuspeed                int    `json:"cpuspeed,omitempty"`
	Oscategoryid            string `json:"oscategoryid,omitempty"`
	Memorytotal             int    `json:"memorytotal,omitempty"`
	State                   string `json:"state,omitempty"`
	Created                 string `json:"created,omitempty"`
	Version                 string `json:"version,omitempty"`
	Cpuallocated            string `json:"cpuallocated,omitempty"`
	Name                    string `json:"name,omitempty"`
	Ipaddress               string `json:"ipaddress,omitempty"`
	Disksizetotal           int    `json:"disksizetotal,omitempty"`
	Clustername             string `json:"clustername,omitempty"`
	Hasenoughcapacity       bool   `json:"hasenoughcapacity,omitempty"`
	Cpusockets              int    `json:"cpusockets,omitempty"`
	Clustertype             string `json:"clustertype,omitempty"`
	Disksizeallocated       int    `json:"disksizeallocated,omitempty"`
	Id                      string `json:"id,omitempty"`
	Islocalstorageactive    bool   `json:"islocalstorageactive,omitempty"`
	Zoneid                  string `json:"zoneid,omitempty"`
	Hypervisor              string `json:"hypervisor,omitempty"`
	Events                  string `json:"events,omitempty"`
	Managementserverid      int    `json:"managementserverid,omitempty"`
	Oscategoryname          string `json:"oscategoryname,omitempty"`
	Type                    string `json:"type,omitempty"`
	Hahost                  bool   `json:"hahost,omitempty"`
	Networkkbsread          int    `json:"networkkbsread,omitempty"`
	Zonename                string `json:"zonename,omitempty"`
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
	Zoneid                  string `json:"zoneid,omitempty"`
	Id                      string `json:"id,omitempty"`
	Cpunumber               int    `json:"cpunumber,omitempty"`
	Disksizetotal           int    `json:"disksizetotal,omitempty"`
	Hypervisorversion       string `json:"hypervisorversion,omitempty"`
	Resourcestate           string `json:"resourcestate,omitempty"`
	Cpuspeed                int    `json:"cpuspeed,omitempty"`
	Podid                   string `json:"podid,omitempty"`
	Lastpinged              string `json:"lastpinged,omitempty"`
	Managementserverid      int    `json:"managementserverid,omitempty"`
	Memoryused              int    `json:"memoryused,omitempty"`
	Networkkbsread          int    `json:"networkkbsread,omitempty"`
	Cpuallocated            string `json:"cpuallocated,omitempty"`
	Suitableformigration    bool   `json:"suitableformigration,omitempty"`
	Memorytotal             int    `json:"memorytotal,omitempty"`
	Hosttags                string `json:"hosttags,omitempty"`
	Hahost                  bool   `json:"hahost,omitempty"`
	Clustername             string `json:"clustername,omitempty"`
	Zonename                string `json:"zonename,omitempty"`
	Clustertype             string `json:"clustertype,omitempty"`
	Cpusockets              int    `json:"cpusockets,omitempty"`
	Oscategoryname          string `json:"oscategoryname,omitempty"`
	Type                    string `json:"type,omitempty"`
	Memoryallocated         int    `json:"memoryallocated,omitempty"`
	State                   string `json:"state,omitempty"`
	Disconnected            string `json:"disconnected,omitempty"`
	Clusterid               string `json:"clusterid,omitempty"`
	Ipaddress               string `json:"ipaddress,omitempty"`
	Oscategoryid            string `json:"oscategoryid,omitempty"`
	Cpuwithoverprovisioning string `json:"cpuwithoverprovisioning,omitempty"`
	Cpuused                 string `json:"cpuused,omitempty"`
	Version                 string `json:"version,omitempty"`
	Name                    string `json:"name,omitempty"`
	Removed                 string `json:"removed,omitempty"`
	Hasenoughcapacity       bool   `json:"hasenoughcapacity,omitempty"`
	Capabilities            string `json:"capabilities,omitempty"`
	Networkkbswrite         int    `json:"networkkbswrite,omitempty"`
	Events                  string `json:"events,omitempty"`
	Islocalstorageactive    bool   `json:"islocalstorageactive,omitempty"`
	Created                 string `json:"created,omitempty"`
	Averageload             int    `json:"averageload,omitempty"`
	Podname                 string `json:"podname,omitempty"`
	Hypervisor              string `json:"hypervisor,omitempty"`
	Disksizeallocated       int    `json:"disksizeallocated,omitempty"`
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
	Created                 string `json:"created,omitempty"`
	Podid                   string `json:"podid,omitempty"`
	Cpunumber               int    `json:"cpunumber,omitempty"`
	Hypervisorversion       string `json:"hypervisorversion,omitempty"`
	Capabilities            string `json:"capabilities,omitempty"`
	Removed                 string `json:"removed,omitempty"`
	Id                      string `json:"id,omitempty"`
	Podname                 string `json:"podname,omitempty"`
	Hahost                  bool   `json:"hahost,omitempty"`
	Cpuused                 string `json:"cpuused,omitempty"`
	Cpuallocated            string `json:"cpuallocated,omitempty"`
	Zoneid                  string `json:"zoneid,omitempty"`
	Disksizetotal           int    `json:"disksizetotal,omitempty"`
	Oscategoryid            string `json:"oscategoryid,omitempty"`
	Networkkbswrite         int    `json:"networkkbswrite,omitempty"`
	Networkkbsread          int    `json:"networkkbsread,omitempty"`
	Managementserverid      int    `json:"managementserverid,omitempty"`
	Hypervisor              string `json:"hypervisor,omitempty"`
	Clustername             string `json:"clustername,omitempty"`
	Cpuwithoverprovisioning string `json:"cpuwithoverprovisioning,omitempty"`
	Clustertype             string `json:"clustertype,omitempty"`
	Cpusockets              int    `json:"cpusockets,omitempty"`
	Zonename                string `json:"zonename,omitempty"`
	Memorytotal             int    `json:"memorytotal,omitempty"`
	Memoryallocated         int    `json:"memoryallocated,omitempty"`
	Suitableformigration    bool   `json:"suitableformigration,omitempty"`
	Hosttags                string `json:"hosttags,omitempty"`
	Type                    string `json:"type,omitempty"`
	Memoryused              int    `json:"memoryused,omitempty"`
	Islocalstorageactive    bool   `json:"islocalstorageactive,omitempty"`
	Averageload             int    `json:"averageload,omitempty"`
	Events                  string `json:"events,omitempty"`
	Ipaddress               string `json:"ipaddress,omitempty"`
	Hasenoughcapacity       bool   `json:"hasenoughcapacity,omitempty"`
	Name                    string `json:"name,omitempty"`
	Version                 string `json:"version,omitempty"`
	Clusterid               string `json:"clusterid,omitempty"`
	Lastpinged              string `json:"lastpinged,omitempty"`
	Disksizeallocated       int    `json:"disksizeallocated,omitempty"`
	Oscategoryname          string `json:"oscategoryname,omitempty"`
	State                   string `json:"state,omitempty"`
	Cpuspeed                int    `json:"cpuspeed,omitempty"`
	Resourcestate           string `json:"resourcestate,omitempty"`
	Disconnected            string `json:"disconnected,omitempty"`
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
	Name                    string `json:"name,omitempty"`
	Hypervisorversion       string `json:"hypervisorversion,omitempty"`
	Ipaddress               string `json:"ipaddress,omitempty"`
	Clustertype             string `json:"clustertype,omitempty"`
	Cpuwithoverprovisioning string `json:"cpuwithoverprovisioning,omitempty"`
	Created                 string `json:"created,omitempty"`
	Resourcestate           string `json:"resourcestate,omitempty"`
	Hypervisor              string `json:"hypervisor,omitempty"`
	Hahost                  bool   `json:"hahost,omitempty"`
	Disksizetotal           int    `json:"disksizetotal,omitempty"`
	Events                  string `json:"events,omitempty"`
	Hasenoughcapacity       bool   `json:"hasenoughcapacity,omitempty"`
	Version                 string `json:"version,omitempty"`
	Capabilities            string `json:"capabilities,omitempty"`
	Oscategoryname          string `json:"oscategoryname,omitempty"`
	Id                      string `json:"id,omitempty"`
	Networkkbsread          int    `json:"networkkbsread,omitempty"`
	Averageload             int    `json:"averageload,omitempty"`
	Oscategoryid            string `json:"oscategoryid,omitempty"`
	Zoneid                  string `json:"zoneid,omitempty"`
	Cpuused                 string `json:"cpuused,omitempty"`
	Networkkbswrite         int    `json:"networkkbswrite,omitempty"`
	Suitableformigration    bool   `json:"suitableformigration,omitempty"`
	RequiresStorageMotion   bool   `json:"requiresStorageMotion,omitempty"`
	Cpuspeed                int    `json:"cpuspeed,omitempty"`
	Disksizeallocated       int    `json:"disksizeallocated,omitempty"`
	Podid                   string `json:"podid,omitempty"`
	Cpunumber               int    `json:"cpunumber,omitempty"`
	Hosttags                string `json:"hosttags,omitempty"`
	Clusterid               string `json:"clusterid,omitempty"`
	Memorytotal             int    `json:"memorytotal,omitempty"`
	Managementserverid      int    `json:"managementserverid,omitempty"`
	Zonename                string `json:"zonename,omitempty"`
	Type                    string `json:"type,omitempty"`
	Disconnected            string `json:"disconnected,omitempty"`
	Memoryallocated         int    `json:"memoryallocated,omitempty"`
	Podname                 string `json:"podname,omitempty"`
	Clustername             string `json:"clustername,omitempty"`
	Lastpinged              string `json:"lastpinged,omitempty"`
	Islocalstorageactive    bool   `json:"islocalstorageactive,omitempty"`
	Cpuallocated            string `json:"cpuallocated,omitempty"`
	Removed                 string `json:"removed,omitempty"`
	State                   string `json:"state,omitempty"`
	Memoryused              int    `json:"memoryused,omitempty"`
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
	Url          string   `json:"url,omitempty"`
	Scope        string   `json:"scope,omitempty"`
	Details      []string `json:"details,omitempty"`
	Protocol     string   `json:"protocol,omitempty"`
	Zonename     string   `json:"zonename,omitempty"`
	Providername string   `json:"providername,omitempty"`
	Id           string   `json:"id,omitempty"`
	Zoneid       string   `json:"zoneid,omitempty"`
	Name         string   `json:"name,omitempty"`
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
	Cpuspeed                int    `json:"cpuspeed,omitempty"`
	Hypervisorversion       string `json:"hypervisorversion,omitempty"`
	Ipaddress               string `json:"ipaddress,omitempty"`
	Oscategoryname          string `json:"oscategoryname,omitempty"`
	Cpunumber               int    `json:"cpunumber,omitempty"`
	Cpuwithoverprovisioning string `json:"cpuwithoverprovisioning,omitempty"`
	Networkkbswrite         int    `json:"networkkbswrite,omitempty"`
	Islocalstorageactive    bool   `json:"islocalstorageactive,omitempty"`
	Cpusockets              int    `json:"cpusockets,omitempty"`
	Suitableformigration    bool   `json:"suitableformigration,omitempty"`
	Clustername             string `json:"clustername,omitempty"`
	Resourcestate           string `json:"resourcestate,omitempty"`
	Disksizetotal           int    `json:"disksizetotal,omitempty"`
	Memoryused              int    `json:"memoryused,omitempty"`
	Oscategoryid            string `json:"oscategoryid,omitempty"`
	Hosttags                string `json:"hosttags,omitempty"`
	Version                 string `json:"version,omitempty"`
	Cpuused                 string `json:"cpuused,omitempty"`
	Hypervisor              string `json:"hypervisor,omitempty"`
	Events                  string `json:"events,omitempty"`
	Name                    string `json:"name,omitempty"`
	Created                 string `json:"created,omitempty"`
	Id                      string `json:"id,omitempty"`
	Type                    string `json:"type,omitempty"`
	Memoryallocated         int    `json:"memoryallocated,omitempty"`
	Zonename                string `json:"zonename,omitempty"`
	Removed                 string `json:"removed,omitempty"`
	Networkkbsread          int    `json:"networkkbsread,omitempty"`
	Hasenoughcapacity       bool   `json:"hasenoughcapacity,omitempty"`
	Capabilities            string `json:"capabilities,omitempty"`
	Podid                   string `json:"podid,omitempty"`
	Memorytotal             int    `json:"memorytotal,omitempty"`
	Disconnected            string `json:"disconnected,omitempty"`
	State                   string `json:"state,omitempty"`
	Averageload             int    `json:"averageload,omitempty"`
	Lastpinged              string `json:"lastpinged,omitempty"`
	Cpuallocated            string `json:"cpuallocated,omitempty"`
	Clustertype             string `json:"clustertype,omitempty"`
	Clusterid               string `json:"clusterid,omitempty"`
	Disksizeallocated       int    `json:"disksizeallocated,omitempty"`
	Podname                 string `json:"podname,omitempty"`
	Hahost                  bool   `json:"hahost,omitempty"`
	Zoneid                  string `json:"zoneid,omitempty"`
	Managementserverid      int    `json:"managementserverid,omitempty"`
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
	Id              string `json:"id,omitempty"`
	Affinitygroupid string `json:"affinitygroupid,omitempty"`
	Accountid       string `json:"accountid,omitempty"`
	Hostname        string `json:"hostname,omitempty"`
	Domainid        string `json:"domainid,omitempty"`
	Hostid          string `json:"hostid,omitempty"`
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
	Success     bool   `json:"success,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
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
	Hostname        string `json:"hostname,omitempty"`
	Affinitygroupid string `json:"affinitygroupid,omitempty"`
	Hostid          string `json:"hostid,omitempty"`
	Domainid        string `json:"domainid,omitempty"`
	Id              string `json:"id,omitempty"`
}
