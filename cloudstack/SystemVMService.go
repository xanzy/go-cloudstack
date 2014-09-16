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

type StartSystemVmParams struct {
	p map[string]interface{}
}

func (p *StartSystemVmParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *StartSystemVmParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new StartSystemVmParams instance,
// as then you are sure you have configured all required params
func (s *SystemVMService) NewStartSystemVmParams(id string) *StartSystemVmParams {
	p := &StartSystemVmParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Starts a system virtual machine.
func (s *SystemVMService) StartSystemVm(p *StartSystemVmParams) (*StartSystemVmResponse, error) {
	resp, err := s.cs.newRequest("startSystemVm", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r StartSystemVmResponse
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

		var r StartSystemVmResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type StartSystemVmResponse struct {
	JobID                string `json:"jobid,omitempty"`
	Jobstatus            int    `json:"jobstatus,omitempty"`
	Dns1                 string `json:"dns1,omitempty"`
	Privatenetmask       string `json:"privatenetmask,omitempty"`
	Hostname             string `json:"hostname,omitempty"`
	State                string `json:"state,omitempty"`
	Networkdomain        string `json:"networkdomain,omitempty"`
	Privatemacaddress    string `json:"privatemacaddress,omitempty"`
	Linklocalmacaddress  string `json:"linklocalmacaddress,omitempty"`
	Created              string `json:"created,omitempty"`
	Zoneid               string `json:"zoneid,omitempty"`
	Publicnetmask        string `json:"publicnetmask,omitempty"`
	Jobid                string `json:"jobid,omitempty"`
	Systemvmtype         string `json:"systemvmtype,omitempty"`
	Name                 string `json:"name,omitempty"`
	Privateip            string `json:"privateip,omitempty"`
	Linklocalnetmask     string `json:"linklocalnetmask,omitempty"`
	Publicip             string `json:"publicip,omitempty"`
	Activeviewersessions int    `json:"activeviewersessions,omitempty"`
	Podid                string `json:"podid,omitempty"`
	Templateid           string `json:"templateid,omitempty"`
	Hostid               string `json:"hostid,omitempty"`
	Dns2                 string `json:"dns2,omitempty"`
	Publicmacaddress     string `json:"publicmacaddress,omitempty"`
	Gateway              string `json:"gateway,omitempty"`
	Id                   string `json:"id,omitempty"`
	Zonename             string `json:"zonename,omitempty"`
	Linklocalip          string `json:"linklocalip,omitempty"`
}

type RebootSystemVmParams struct {
	p map[string]interface{}
}

func (p *RebootSystemVmParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *RebootSystemVmParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new RebootSystemVmParams instance,
// as then you are sure you have configured all required params
func (s *SystemVMService) NewRebootSystemVmParams(id string) *RebootSystemVmParams {
	p := &RebootSystemVmParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Reboots a system VM.
func (s *SystemVMService) RebootSystemVm(p *RebootSystemVmParams) (*RebootSystemVmResponse, error) {
	resp, err := s.cs.newRequest("rebootSystemVm", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RebootSystemVmResponse
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

		var r RebootSystemVmResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type RebootSystemVmResponse struct {
	JobID                string `json:"jobid,omitempty"`
	Name                 string `json:"name,omitempty"`
	Publicip             string `json:"publicip,omitempty"`
	Privatemacaddress    string `json:"privatemacaddress,omitempty"`
	Templateid           string `json:"templateid,omitempty"`
	Publicmacaddress     string `json:"publicmacaddress,omitempty"`
	Jobid                string `json:"jobid,omitempty"`
	Linklocalmacaddress  string `json:"linklocalmacaddress,omitempty"`
	Zoneid               string `json:"zoneid,omitempty"`
	Jobstatus            int    `json:"jobstatus,omitempty"`
	Networkdomain        string `json:"networkdomain,omitempty"`
	Linklocalip          string `json:"linklocalip,omitempty"`
	Podid                string `json:"podid,omitempty"`
	Gateway              string `json:"gateway,omitempty"`
	Systemvmtype         string `json:"systemvmtype,omitempty"`
	Linklocalnetmask     string `json:"linklocalnetmask,omitempty"`
	Zonename             string `json:"zonename,omitempty"`
	Id                   string `json:"id,omitempty"`
	Dns2                 string `json:"dns2,omitempty"`
	Hostname             string `json:"hostname,omitempty"`
	Dns1                 string `json:"dns1,omitempty"`
	State                string `json:"state,omitempty"`
	Activeviewersessions int    `json:"activeviewersessions,omitempty"`
	Publicnetmask        string `json:"publicnetmask,omitempty"`
	Hostid               string `json:"hostid,omitempty"`
	Privateip            string `json:"privateip,omitempty"`
	Created              string `json:"created,omitempty"`
	Privatenetmask       string `json:"privatenetmask,omitempty"`
}

type StopSystemVmParams struct {
	p map[string]interface{}
}

func (p *StopSystemVmParams) toURLValues() url.Values {
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

func (p *StopSystemVmParams) SetForced(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forced"] = v
	return
}

func (p *StopSystemVmParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new StopSystemVmParams instance,
// as then you are sure you have configured all required params
func (s *SystemVMService) NewStopSystemVmParams(id string) *StopSystemVmParams {
	p := &StopSystemVmParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Stops a system VM.
func (s *SystemVMService) StopSystemVm(p *StopSystemVmParams) (*StopSystemVmResponse, error) {
	resp, err := s.cs.newRequest("stopSystemVm", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r StopSystemVmResponse
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

		var r StopSystemVmResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type StopSystemVmResponse struct {
	JobID                string `json:"jobid,omitempty"`
	Jobid                string `json:"jobid,omitempty"`
	Privatenetmask       string `json:"privatenetmask,omitempty"`
	Linklocalmacaddress  string `json:"linklocalmacaddress,omitempty"`
	Name                 string `json:"name,omitempty"`
	Id                   string `json:"id,omitempty"`
	Publicip             string `json:"publicip,omitempty"`
	Publicnetmask        string `json:"publicnetmask,omitempty"`
	Jobstatus            int    `json:"jobstatus,omitempty"`
	Templateid           string `json:"templateid,omitempty"`
	Hostname             string `json:"hostname,omitempty"`
	Created              string `json:"created,omitempty"`
	Systemvmtype         string `json:"systemvmtype,omitempty"`
	Linklocalip          string `json:"linklocalip,omitempty"`
	Publicmacaddress     string `json:"publicmacaddress,omitempty"`
	Gateway              string `json:"gateway,omitempty"`
	Zoneid               string `json:"zoneid,omitempty"`
	Activeviewersessions int    `json:"activeviewersessions,omitempty"`
	Podid                string `json:"podid,omitempty"`
	Dns1                 string `json:"dns1,omitempty"`
	Zonename             string `json:"zonename,omitempty"`
	State                string `json:"state,omitempty"`
	Hostid               string `json:"hostid,omitempty"`
	Dns2                 string `json:"dns2,omitempty"`
	Linklocalnetmask     string `json:"linklocalnetmask,omitempty"`
	Privatemacaddress    string `json:"privatemacaddress,omitempty"`
	Networkdomain        string `json:"networkdomain,omitempty"`
	Privateip            string `json:"privateip,omitempty"`
}

type DestroySystemVmParams struct {
	p map[string]interface{}
}

func (p *DestroySystemVmParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DestroySystemVmParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DestroySystemVmParams instance,
// as then you are sure you have configured all required params
func (s *SystemVMService) NewDestroySystemVmParams(id string) *DestroySystemVmParams {
	p := &DestroySystemVmParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Destroyes a system virtual machine.
func (s *SystemVMService) DestroySystemVm(p *DestroySystemVmParams) (*DestroySystemVmResponse, error) {
	resp, err := s.cs.newRequest("destroySystemVm", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DestroySystemVmResponse
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

		var r DestroySystemVmResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type DestroySystemVmResponse struct {
	JobID                string `json:"jobid,omitempty"`
	Linklocalnetmask     string `json:"linklocalnetmask,omitempty"`
	Publicmacaddress     string `json:"publicmacaddress,omitempty"`
	Systemvmtype         string `json:"systemvmtype,omitempty"`
	Hostname             string `json:"hostname,omitempty"`
	Id                   string `json:"id,omitempty"`
	Jobid                string `json:"jobid,omitempty"`
	Privatemacaddress    string `json:"privatemacaddress,omitempty"`
	Linklocalmacaddress  string `json:"linklocalmacaddress,omitempty"`
	Activeviewersessions int    `json:"activeviewersessions,omitempty"`
	Privatenetmask       string `json:"privatenetmask,omitempty"`
	Publicip             string `json:"publicip,omitempty"`
	State                string `json:"state,omitempty"`
	Publicnetmask        string `json:"publicnetmask,omitempty"`
	Hostid               string `json:"hostid,omitempty"`
	Privateip            string `json:"privateip,omitempty"`
	Templateid           string `json:"templateid,omitempty"`
	Dns2                 string `json:"dns2,omitempty"`
	Networkdomain        string `json:"networkdomain,omitempty"`
	Zonename             string `json:"zonename,omitempty"`
	Podid                string `json:"podid,omitempty"`
	Linklocalip          string `json:"linklocalip,omitempty"`
	Created              string `json:"created,omitempty"`
	Gateway              string `json:"gateway,omitempty"`
	Name                 string `json:"name,omitempty"`
	Dns1                 string `json:"dns1,omitempty"`
	Zoneid               string `json:"zoneid,omitempty"`
	Jobstatus            int    `json:"jobstatus,omitempty"`
}

type ListSystemVmsParams struct {
	p map[string]interface{}
}

func (p *ListSystemVmsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["hostid"]; found {
		u.Set("hostid", v.(string))
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
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := p.p["storageid"]; found {
		u.Set("storageid", v.(string))
	}
	if v, found := p.p["systemvmtype"]; found {
		u.Set("systemvmtype", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListSystemVmsParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
	return
}

func (p *ListSystemVmsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListSystemVmsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListSystemVmsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *ListSystemVmsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListSystemVmsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListSystemVmsParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
	return
}

func (p *ListSystemVmsParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
	return
}

func (p *ListSystemVmsParams) SetStorageid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["storageid"] = v
	return
}

func (p *ListSystemVmsParams) SetSystemvmtype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["systemvmtype"] = v
	return
}

func (p *ListSystemVmsParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new ListSystemVmsParams instance,
// as then you are sure you have configured all required params
func (s *SystemVMService) NewListSystemVmsParams() *ListSystemVmsParams {
	p := &ListSystemVmsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SystemVMService) GetSystemVmID(name string) (string, error) {
	p := &ListSystemVmsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	l, err := s.ListSystemVms(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, name, l)
	}
	return l.SystemVms[0].Id, nil
}

// List system virtual machines.
func (s *SystemVMService) ListSystemVms(p *ListSystemVmsParams) (*ListSystemVmsResponse, error) {
	resp, err := s.cs.newRequest("listSystemVms", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListSystemVmsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListSystemVmsResponse struct {
	Count     int         `json:"count"`
	SystemVms []*SystemVm `json:"systemvm"`
}

type SystemVm struct {
	Publicip             string `json:"publicip,omitempty"`
	Privatemacaddress    string `json:"privatemacaddress,omitempty"`
	Activeviewersessions int    `json:"activeviewersessions,omitempty"`
	Privateip            string `json:"privateip,omitempty"`
	Podid                string `json:"podid,omitempty"`
	State                string `json:"state,omitempty"`
	Created              string `json:"created,omitempty"`
	Linklocalmacaddress  string `json:"linklocalmacaddress,omitempty"`
	Networkdomain        string `json:"networkdomain,omitempty"`
	Hostname             string `json:"hostname,omitempty"`
	Name                 string `json:"name,omitempty"`
	Jobstatus            int    `json:"jobstatus,omitempty"`
	Gateway              string `json:"gateway,omitempty"`
	Id                   string `json:"id,omitempty"`
	Linklocalip          string `json:"linklocalip,omitempty"`
	Dns1                 string `json:"dns1,omitempty"`
	Systemvmtype         string `json:"systemvmtype,omitempty"`
	Linklocalnetmask     string `json:"linklocalnetmask,omitempty"`
	Templateid           string `json:"templateid,omitempty"`
	Dns2                 string `json:"dns2,omitempty"`
	Publicnetmask        string `json:"publicnetmask,omitempty"`
	Zonename             string `json:"zonename,omitempty"`
	Hostid               string `json:"hostid,omitempty"`
	Publicmacaddress     string `json:"publicmacaddress,omitempty"`
	Jobid                string `json:"jobid,omitempty"`
	Privatenetmask       string `json:"privatenetmask,omitempty"`
	Zoneid               string `json:"zoneid,omitempty"`
}

type MigrateSystemVmParams struct {
	p map[string]interface{}
}

func (p *MigrateSystemVmParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *MigrateSystemVmParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
	return
}

func (p *MigrateSystemVmParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
	return
}

// You should always use this function to get a new MigrateSystemVmParams instance,
// as then you are sure you have configured all required params
func (s *SystemVMService) NewMigrateSystemVmParams(hostid string, virtualmachineid string) *MigrateSystemVmParams {
	p := &MigrateSystemVmParams{}
	p.p = make(map[string]interface{})
	p.p["hostid"] = hostid
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// Attempts Migration of a system virtual machine to the host specified.
func (s *SystemVMService) MigrateSystemVm(p *MigrateSystemVmParams) (*MigrateSystemVmResponse, error) {
	resp, err := s.cs.newRequest("migrateSystemVm", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r MigrateSystemVmResponse
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

		var r MigrateSystemVmResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type MigrateSystemVmResponse struct {
	JobID                string `json:"jobid,omitempty"`
	Networkdomain        string `json:"networkdomain,omitempty"`
	Hostid               string `json:"hostid,omitempty"`
	Dns1                 string `json:"dns1,omitempty"`
	Linklocalmacaddress  string `json:"linklocalmacaddress,omitempty"`
	Podid                string `json:"podid,omitempty"`
	Activeviewersessions int    `json:"activeviewersessions,omitempty"`
	Publicnetmask        string `json:"publicnetmask,omitempty"`
	Gateway              string `json:"gateway,omitempty"`
	Privatenetmask       string `json:"privatenetmask,omitempty"`
	Dns2                 string `json:"dns2,omitempty"`
	Linklocalip          string `json:"linklocalip,omitempty"`
	State                string `json:"state,omitempty"`
	Systemvmtype         string `json:"systemvmtype,omitempty"`
	Jobstatus            int    `json:"jobstatus,omitempty"`
	Created              string `json:"created,omitempty"`
	Linklocalnetmask     string `json:"linklocalnetmask,omitempty"`
	Jobid                string `json:"jobid,omitempty"`
	Publicip             string `json:"publicip,omitempty"`
	Id                   string `json:"id,omitempty"`
	Zonename             string `json:"zonename,omitempty"`
	Publicmacaddress     string `json:"publicmacaddress,omitempty"`
	Zoneid               string `json:"zoneid,omitempty"`
	Hostname             string `json:"hostname,omitempty"`
	Privatemacaddress    string `json:"privatemacaddress,omitempty"`
	Name                 string `json:"name,omitempty"`
	Privateip            string `json:"privateip,omitempty"`
	Templateid           string `json:"templateid,omitempty"`
}

type ChangeServiceForSystemVmParams struct {
	p map[string]interface{}
}

func (p *ChangeServiceForSystemVmParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["details"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("details[%d].key", i), k)
			u.Set(fmt.Sprintf("details[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	return u
}

func (p *ChangeServiceForSystemVmParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
	return
}

func (p *ChangeServiceForSystemVmParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ChangeServiceForSystemVmParams) SetServiceofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingid"] = v
	return
}

// You should always use this function to get a new ChangeServiceForSystemVmParams instance,
// as then you are sure you have configured all required params
func (s *SystemVMService) NewChangeServiceForSystemVmParams(id string, serviceofferingid string) *ChangeServiceForSystemVmParams {
	p := &ChangeServiceForSystemVmParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["serviceofferingid"] = serviceofferingid
	return p
}

// Changes the service offering for a system vm (console proxy or secondary storage). The system vm must be in a "Stopped" state for this command to take effect.
func (s *SystemVMService) ChangeServiceForSystemVm(p *ChangeServiceForSystemVmParams) (*ChangeServiceForSystemVmResponse, error) {
	resp, err := s.cs.newRequest("changeServiceForSystemVm", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ChangeServiceForSystemVmResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ChangeServiceForSystemVmResponse struct {
	Publicnetmask        string `json:"publicnetmask,omitempty"`
	Name                 string `json:"name,omitempty"`
	Zoneid               string `json:"zoneid,omitempty"`
	Created              string `json:"created,omitempty"`
	Podid                string `json:"podid,omitempty"`
	Dns1                 string `json:"dns1,omitempty"`
	Gateway              string `json:"gateway,omitempty"`
	Hostname             string `json:"hostname,omitempty"`
	Activeviewersessions int    `json:"activeviewersessions,omitempty"`
	Jobstatus            int    `json:"jobstatus,omitempty"`
	Linklocalip          string `json:"linklocalip,omitempty"`
	Linklocalnetmask     string `json:"linklocalnetmask,omitempty"`
	Systemvmtype         string `json:"systemvmtype,omitempty"`
	Templateid           string `json:"templateid,omitempty"`
	Id                   string `json:"id,omitempty"`
	Privatenetmask       string `json:"privatenetmask,omitempty"`
	Publicip             string `json:"publicip,omitempty"`
	Zonename             string `json:"zonename,omitempty"`
	Privateip            string `json:"privateip,omitempty"`
	Linklocalmacaddress  string `json:"linklocalmacaddress,omitempty"`
	Networkdomain        string `json:"networkdomain,omitempty"`
	State                string `json:"state,omitempty"`
	Dns2                 string `json:"dns2,omitempty"`
	Publicmacaddress     string `json:"publicmacaddress,omitempty"`
	Hostid               string `json:"hostid,omitempty"`
	Jobid                string `json:"jobid,omitempty"`
	Privatemacaddress    string `json:"privatemacaddress,omitempty"`
}

type ScaleSystemVmParams struct {
	p map[string]interface{}
}

func (p *ScaleSystemVmParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["details"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("details[%d].key", i), k)
			u.Set(fmt.Sprintf("details[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	return u
}

func (p *ScaleSystemVmParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
	return
}

func (p *ScaleSystemVmParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ScaleSystemVmParams) SetServiceofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingid"] = v
	return
}

// You should always use this function to get a new ScaleSystemVmParams instance,
// as then you are sure you have configured all required params
func (s *SystemVMService) NewScaleSystemVmParams(id string, serviceofferingid string) *ScaleSystemVmParams {
	p := &ScaleSystemVmParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["serviceofferingid"] = serviceofferingid
	return p
}

// Scale the service offering for a system vm (console proxy or secondary storage). The system vm must be in a "Stopped" state for this command to take effect.
func (s *SystemVMService) ScaleSystemVm(p *ScaleSystemVmParams) (*ScaleSystemVmResponse, error) {
	resp, err := s.cs.newRequest("scaleSystemVm", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ScaleSystemVmResponse
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

		var r ScaleSystemVmResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type ScaleSystemVmResponse struct {
	JobID                string `json:"jobid,omitempty"`
	Privatenetmask       string `json:"privatenetmask,omitempty"`
	Publicip             string `json:"publicip,omitempty"`
	Hostname             string `json:"hostname,omitempty"`
	Jobid                string `json:"jobid,omitempty"`
	Dns1                 string `json:"dns1,omitempty"`
	Linklocalip          string `json:"linklocalip,omitempty"`
	State                string `json:"state,omitempty"`
	Zoneid               string `json:"zoneid,omitempty"`
	Privatemacaddress    string `json:"privatemacaddress,omitempty"`
	Name                 string `json:"name,omitempty"`
	Systemvmtype         string `json:"systemvmtype,omitempty"`
	Dns2                 string `json:"dns2,omitempty"`
	Networkdomain        string `json:"networkdomain,omitempty"`
	Linklocalmacaddress  string `json:"linklocalmacaddress,omitempty"`
	Id                   string `json:"id,omitempty"`
	Privateip            string `json:"privateip,omitempty"`
	Created              string `json:"created,omitempty"`
	Linklocalnetmask     string `json:"linklocalnetmask,omitempty"`
	Publicnetmask        string `json:"publicnetmask,omitempty"`
	Jobstatus            int    `json:"jobstatus,omitempty"`
	Podid                string `json:"podid,omitempty"`
	Templateid           string `json:"templateid,omitempty"`
	Gateway              string `json:"gateway,omitempty"`
	Hostid               string `json:"hostid,omitempty"`
	Activeviewersessions int    `json:"activeviewersessions,omitempty"`
	Zonename             string `json:"zonename,omitempty"`
	Publicmacaddress     string `json:"publicmacaddress,omitempty"`
}
