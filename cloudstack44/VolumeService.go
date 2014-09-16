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

type AttachVolumeParams struct {
	p map[string]interface{}
}

func (p *AttachVolumeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["deviceid"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("deviceid", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *AttachVolumeParams) SetDeviceid(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["deviceid"] = v
	return
}

func (p *AttachVolumeParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *AttachVolumeParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
	return
}

// You should always use this function to get a new AttachVolumeParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewAttachVolumeParams(id string, virtualmachineid string) *AttachVolumeParams {
	p := &AttachVolumeParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// Attaches a disk volume to a virtual machine.
func (s *VolumeService) AttachVolume(p *AttachVolumeParams) (*AttachVolumeResponse, error) {
	resp, err := s.cs.newRequest("attachVolume", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AttachVolumeResponse
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

		var r AttachVolumeResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type AttachVolumeResponse struct {
	JobID                      string `json:"jobid,omitempty"`
	Serviceofferingname        string `json:"serviceofferingname,omitempty"`
	Diskofferingid             string `json:"diskofferingid,omitempty"`
	Id                         string `json:"id,omitempty"`
	Serviceofferingid          string `json:"serviceofferingid,omitempty"`
	Diskofferingname           string `json:"diskofferingname,omitempty"`
	Deviceid                   int    `json:"deviceid,omitempty"`
	DiskBytesReadRate          int    `json:"diskBytesReadRate,omitempty"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext,omitempty"`
	Name                       string `json:"name,omitempty"`
	Vmdisplayname              string `json:"vmdisplayname,omitempty"`
	Zonename                   string `json:"zonename,omitempty"`
	State                      string `json:"state,omitempty"`
	Path                       string `json:"path,omitempty"`
	DiskIopsReadRate           int    `json:"diskIopsReadRate,omitempty"`
	Quiescevm                  bool   `json:"quiescevm,omitempty"`
	Isextractable              bool   `json:"isextractable,omitempty"`
	DiskIopsWriteRate          int    `json:"diskIopsWriteRate,omitempty"`
	Projectid                  string `json:"projectid,omitempty"`
	Project                    string `json:"project,omitempty"`
	Virtualmachineid           string `json:"virtualmachineid,omitempty"`
	Attached                   string `json:"attached,omitempty"`
	Vmname                     string `json:"vmname,omitempty"`
	Size                       int    `json:"size,omitempty"`
	DiskBytesWriteRate         int    `json:"diskBytesWriteRate,omitempty"`
	Maxiops                    int    `json:"maxiops,omitempty"`
	Storagetype                string `json:"storagetype,omitempty"`
	Displayvolume              bool   `json:"displayvolume,omitempty"`
	Vmstate                    string `json:"vmstate,omitempty"`
	Destroyed                  bool   `json:"destroyed,omitempty"`
	Snapshotid                 string `json:"snapshotid,omitempty"`
	Storage                    string `json:"storage,omitempty"`
	Hypervisor                 string `json:"hypervisor,omitempty"`
	Domain                     string `json:"domain,omitempty"`
	Status                     string `json:"status,omitempty"`
	Zoneid                     string `json:"zoneid,omitempty"`
	Tags                       []struct {
		Resourceid   string `json:"resourceid,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Project      string `json:"project,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Key          string `json:"key,omitempty"`
		Account      string `json:"account,omitempty"`
		Value        string `json:"value,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
	} `json:"tags,omitempty"`
	Storageid               string `json:"storageid,omitempty"`
	Diskofferingdisplaytext string `json:"diskofferingdisplaytext,omitempty"`
	Miniops                 int    `json:"miniops,omitempty"`
	Domainid                string `json:"domainid,omitempty"`
	Created                 string `json:"created,omitempty"`
	Account                 string `json:"account,omitempty"`
	Type                    string `json:"type,omitempty"`
}

type UploadVolumeParams struct {
	p map[string]interface{}
}

func (p *UploadVolumeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["checksum"]; found {
		u.Set("checksum", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["format"]; found {
		u.Set("format", v.(string))
	}
	if v, found := p.p["imagestoreuuid"]; found {
		u.Set("imagestoreuuid", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *UploadVolumeParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *UploadVolumeParams) SetChecksum(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["checksum"] = v
	return
}

func (p *UploadVolumeParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *UploadVolumeParams) SetFormat(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["format"] = v
	return
}

func (p *UploadVolumeParams) SetImagestoreuuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["imagestoreuuid"] = v
	return
}

func (p *UploadVolumeParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *UploadVolumeParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *UploadVolumeParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
	return
}

func (p *UploadVolumeParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new UploadVolumeParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewUploadVolumeParams(format string, name string, url string, zoneid string) *UploadVolumeParams {
	p := &UploadVolumeParams{}
	p.p = make(map[string]interface{})
	p.p["format"] = format
	p.p["name"] = name
	p.p["url"] = url
	p.p["zoneid"] = zoneid
	return p
}

// Uploads a data disk.
func (s *VolumeService) UploadVolume(p *UploadVolumeParams) (*UploadVolumeResponse, error) {
	resp, err := s.cs.newRequest("uploadVolume", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UploadVolumeResponse
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

		var r UploadVolumeResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type UploadVolumeResponse struct {
	JobID                      string `json:"jobid,omitempty"`
	Account                    string `json:"account,omitempty"`
	Created                    string `json:"created,omitempty"`
	Projectid                  string `json:"projectid,omitempty"`
	Vmstate                    string `json:"vmstate,omitempty"`
	Domain                     string `json:"domain,omitempty"`
	Hypervisor                 string `json:"hypervisor,omitempty"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext,omitempty"`
	DiskBytesWriteRate         int    `json:"diskBytesWriteRate,omitempty"`
	Attached                   string `json:"attached,omitempty"`
	Miniops                    int    `json:"miniops,omitempty"`
	Storageid                  string `json:"storageid,omitempty"`
	Zonename                   string `json:"zonename,omitempty"`
	Path                       string `json:"path,omitempty"`
	Snapshotid                 string `json:"snapshotid,omitempty"`
	Size                       int    `json:"size,omitempty"`
	DiskIopsWriteRate          int    `json:"diskIopsWriteRate,omitempty"`
	Storage                    string `json:"storage,omitempty"`
	Status                     string `json:"status,omitempty"`
	Id                         string `json:"id,omitempty"`
	Type                       string `json:"type,omitempty"`
	State                      string `json:"state,omitempty"`
	Zoneid                     string `json:"zoneid,omitempty"`
	Deviceid                   int    `json:"deviceid,omitempty"`
	Project                    string `json:"project,omitempty"`
	Diskofferingdisplaytext    string `json:"diskofferingdisplaytext,omitempty"`
	DiskIopsReadRate           int    `json:"diskIopsReadRate,omitempty"`
	Vmdisplayname              string `json:"vmdisplayname,omitempty"`
	Quiescevm                  bool   `json:"quiescevm,omitempty"`
	Serviceofferingname        string `json:"serviceofferingname,omitempty"`
	Isextractable              bool   `json:"isextractable,omitempty"`
	Diskofferingid             string `json:"diskofferingid,omitempty"`
	DiskBytesReadRate          int    `json:"diskBytesReadRate,omitempty"`
	Tags                       []struct {
		Domainid     string `json:"domainid,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Account      string `json:"account,omitempty"`
		Project      string `json:"project,omitempty"`
		Value        string `json:"value,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Key          string `json:"key,omitempty"`
	} `json:"tags,omitempty"`
	Vmname            string `json:"vmname,omitempty"`
	Domainid          string `json:"domainid,omitempty"`
	Virtualmachineid  string `json:"virtualmachineid,omitempty"`
	Storagetype       string `json:"storagetype,omitempty"`
	Destroyed         bool   `json:"destroyed,omitempty"`
	Maxiops           int    `json:"maxiops,omitempty"`
	Name              string `json:"name,omitempty"`
	Displayvolume     bool   `json:"displayvolume,omitempty"`
	Serviceofferingid string `json:"serviceofferingid,omitempty"`
	Diskofferingname  string `json:"diskofferingname,omitempty"`
}

type DetachVolumeParams struct {
	p map[string]interface{}
}

func (p *DetachVolumeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["deviceid"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("deviceid", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *DetachVolumeParams) SetDeviceid(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["deviceid"] = v
	return
}

func (p *DetachVolumeParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *DetachVolumeParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
	return
}

// You should always use this function to get a new DetachVolumeParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewDetachVolumeParams() *DetachVolumeParams {
	p := &DetachVolumeParams{}
	p.p = make(map[string]interface{})
	return p
}

// Detaches a disk volume from a virtual machine.
func (s *VolumeService) DetachVolume(p *DetachVolumeParams) (*DetachVolumeResponse, error) {
	resp, err := s.cs.newRequest("detachVolume", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DetachVolumeResponse
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

		var r DetachVolumeResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type DetachVolumeResponse struct {
	JobID               string `json:"jobid,omitempty"`
	Vmname              string `json:"vmname,omitempty"`
	Isextractable       bool   `json:"isextractable,omitempty"`
	Status              string `json:"status,omitempty"`
	Destroyed           bool   `json:"destroyed,omitempty"`
	Type                string `json:"type,omitempty"`
	Created             string `json:"created,omitempty"`
	Path                string `json:"path,omitempty"`
	Maxiops             int    `json:"maxiops,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Quiescevm           bool   `json:"quiescevm,omitempty"`
	DiskBytesReadRate   int    `json:"diskBytesReadRate,omitempty"`
	Storageid           string `json:"storageid,omitempty"`
	Displayvolume       bool   `json:"displayvolume,omitempty"`
	Snapshotid          string `json:"snapshotid,omitempty"`
	Account             string `json:"account,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
	Size                int    `json:"size,omitempty"`
	Deviceid            int    `json:"deviceid,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	DiskIopsWriteRate   int    `json:"diskIopsWriteRate,omitempty"`
	State               string `json:"state,omitempty"`
	Miniops             int    `json:"miniops,omitempty"`
	Vmstate             string `json:"vmstate,omitempty"`
	Virtualmachineid    string `json:"virtualmachineid,omitempty"`
	Id                  string `json:"id,omitempty"`
	Project             string `json:"project,omitempty"`
	Domainid            string `json:"domainid,omitempty"`
	Storage             string `json:"storage,omitempty"`
	Diskofferingname    string `json:"diskofferingname,omitempty"`
	Projectid           string `json:"projectid,omitempty"`
	Tags                []struct {
		Key          string `json:"key,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Project      string `json:"project,omitempty"`
		Account      string `json:"account,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Value        string `json:"value,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
	} `json:"tags,omitempty"`
	Name                       string `json:"name,omitempty"`
	Diskofferingid             string `json:"diskofferingid,omitempty"`
	DiskBytesWriteRate         int    `json:"diskBytesWriteRate,omitempty"`
	Hypervisor                 string `json:"hypervisor,omitempty"`
	Zoneid                     string `json:"zoneid,omitempty"`
	DiskIopsReadRate           int    `json:"diskIopsReadRate,omitempty"`
	Storagetype                string `json:"storagetype,omitempty"`
	Attached                   string `json:"attached,omitempty"`
	Diskofferingdisplaytext    string `json:"diskofferingdisplaytext,omitempty"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext,omitempty"`
	Vmdisplayname              string `json:"vmdisplayname,omitempty"`
}

type CreateVolumeParams struct {
	p map[string]interface{}
}

func (p *CreateVolumeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["diskofferingid"]; found {
		u.Set("diskofferingid", v.(string))
	}
	if v, found := p.p["displayvolume"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("displayvolume", vv)
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["maxiops"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("maxiops", vv)
	}
	if v, found := p.p["miniops"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("miniops", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["size"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("size", vv)
	}
	if v, found := p.p["snapshotid"]; found {
		u.Set("snapshotid", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *CreateVolumeParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *CreateVolumeParams) SetDiskofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["diskofferingid"] = v
	return
}

func (p *CreateVolumeParams) SetDisplayvolume(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displayvolume"] = v
	return
}

func (p *CreateVolumeParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *CreateVolumeParams) SetMaxiops(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["maxiops"] = v
	return
}

func (p *CreateVolumeParams) SetMiniops(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["miniops"] = v
	return
}

func (p *CreateVolumeParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *CreateVolumeParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *CreateVolumeParams) SetSize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["size"] = v
	return
}

func (p *CreateVolumeParams) SetSnapshotid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["snapshotid"] = v
	return
}

func (p *CreateVolumeParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
	return
}

func (p *CreateVolumeParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new CreateVolumeParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewCreateVolumeParams(name string) *CreateVolumeParams {
	p := &CreateVolumeParams{}
	p.p = make(map[string]interface{})
	p.p["name"] = name
	return p
}

// Creates a disk volume from a disk offering. This disk volume must still be attached to a virtual machine to make use of it.
func (s *VolumeService) CreateVolume(p *CreateVolumeParams) (*CreateVolumeResponse, error) {
	resp, err := s.cs.newRequest("createVolume", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateVolumeResponse
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

		var r CreateVolumeResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type CreateVolumeResponse struct {
	JobID                   string `json:"jobid,omitempty"`
	Vmname                  string `json:"vmname,omitempty"`
	Vmstate                 string `json:"vmstate,omitempty"`
	DiskIopsReadRate        int    `json:"diskIopsReadRate,omitempty"`
	Zonename                string `json:"zonename,omitempty"`
	Created                 string `json:"created,omitempty"`
	Domainid                string `json:"domainid,omitempty"`
	DiskBytesWriteRate      int    `json:"diskBytesWriteRate,omitempty"`
	Diskofferingdisplaytext string `json:"diskofferingdisplaytext,omitempty"`
	Diskofferingname        string `json:"diskofferingname,omitempty"`
	Diskofferingid          string `json:"diskofferingid,omitempty"`
	Hypervisor              string `json:"hypervisor,omitempty"`
	Serviceofferingid       string `json:"serviceofferingid,omitempty"`
	Storage                 string `json:"storage,omitempty"`
	Virtualmachineid        string `json:"virtualmachineid,omitempty"`
	Quiescevm               bool   `json:"quiescevm,omitempty"`
	Snapshotid              string `json:"snapshotid,omitempty"`
	Storagetype             string `json:"storagetype,omitempty"`
	DiskIopsWriteRate       int    `json:"diskIopsWriteRate,omitempty"`
	Isextractable           bool   `json:"isextractable,omitempty"`
	Type                    string `json:"type,omitempty"`
	Miniops                 int    `json:"miniops,omitempty"`
	Storageid               string `json:"storageid,omitempty"`
	Deviceid                int    `json:"deviceid,omitempty"`
	State                   string `json:"state,omitempty"`
	Project                 string `json:"project,omitempty"`
	Size                    int    `json:"size,omitempty"`
	Vmdisplayname           string `json:"vmdisplayname,omitempty"`
	Attached                string `json:"attached,omitempty"`
	Zoneid                  string `json:"zoneid,omitempty"`
	Tags                    []struct {
		Project      string `json:"project,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Value        string `json:"value,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Key          string `json:"key,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Account      string `json:"account,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
	} `json:"tags,omitempty"`
	Name                       string `json:"name,omitempty"`
	Domain                     string `json:"domain,omitempty"`
	Id                         string `json:"id,omitempty"`
	Account                    string `json:"account,omitempty"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext,omitempty"`
	Serviceofferingname        string `json:"serviceofferingname,omitempty"`
	Projectid                  string `json:"projectid,omitempty"`
	Path                       string `json:"path,omitempty"`
	Destroyed                  bool   `json:"destroyed,omitempty"`
	Displayvolume              bool   `json:"displayvolume,omitempty"`
	Status                     string `json:"status,omitempty"`
	Maxiops                    int    `json:"maxiops,omitempty"`
	DiskBytesReadRate          int    `json:"diskBytesReadRate,omitempty"`
}

type DeleteVolumeParams struct {
	p map[string]interface{}
}

func (p *DeleteVolumeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteVolumeParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteVolumeParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewDeleteVolumeParams(id string) *DeleteVolumeParams {
	p := &DeleteVolumeParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a detached disk volume.
func (s *VolumeService) DeleteVolume(p *DeleteVolumeParams) (*DeleteVolumeResponse, error) {
	resp, err := s.cs.newRequest("deleteVolume", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteVolumeResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type DeleteVolumeResponse struct {
	Success     bool   `json:"success,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
}

type ListVolumesParams struct {
	p map[string]interface{}
}

func (p *ListVolumesParams) toURLValues() url.Values {
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
	if v, found := p.p["storageid"]; found {
		u.Set("storageid", v.(string))
	}
	if v, found := p.p["tags"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), vv)
			i++
		}
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

func (p *ListVolumesParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListVolumesParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListVolumesParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
	return
}

func (p *ListVolumesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListVolumesParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
	return
}

func (p *ListVolumesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListVolumesParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
	return
}

func (p *ListVolumesParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *ListVolumesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListVolumesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListVolumesParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
	return
}

func (p *ListVolumesParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *ListVolumesParams) SetStorageid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["storageid"] = v
	return
}

func (p *ListVolumesParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
	return
}

func (p *ListVolumesParams) SetType(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["volumeType"] = v
	return
}

func (p *ListVolumesParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
	return
}

func (p *ListVolumesParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new ListVolumesParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewListVolumesParams() *ListVolumesParams {
	p := &ListVolumesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VolumeService) GetVolumeID(name string) (string, error) {
	p := &ListVolumesParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	l, err := s.ListVolumes(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, name, l)
	}
	return l.Volumes[0].Id, nil
}

// Lists all volumes.
func (s *VolumeService) ListVolumes(p *ListVolumesParams) (*ListVolumesResponse, error) {
	resp, err := s.cs.newRequest("listVolumes", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVolumesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListVolumesResponse struct {
	Count   int       `json:"count"`
	Volumes []*Volume `json:"volume"`
}

type Volume struct {
	Destroyed                  bool   `json:"destroyed,omitempty"`
	Isextractable              bool   `json:"isextractable,omitempty"`
	Diskofferingname           string `json:"diskofferingname,omitempty"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext,omitempty"`
	Size                       int    `json:"size,omitempty"`
	DiskIopsReadRate           int    `json:"diskIopsReadRate,omitempty"`
	Zoneid                     string `json:"zoneid,omitempty"`
	Vmstate                    string `json:"vmstate,omitempty"`
	Name                       string `json:"name,omitempty"`
	Path                       string `json:"path,omitempty"`
	Serviceofferingid          string `json:"serviceofferingid,omitempty"`
	Maxiops                    int    `json:"maxiops,omitempty"`
	Vmname                     string `json:"vmname,omitempty"`
	Diskofferingdisplaytext    string `json:"diskofferingdisplaytext,omitempty"`
	Id                         string `json:"id,omitempty"`
	Projectid                  string `json:"projectid,omitempty"`
	State                      string `json:"state,omitempty"`
	Type                       string `json:"type,omitempty"`
	Diskofferingid             string `json:"diskofferingid,omitempty"`
	Deviceid                   int    `json:"deviceid,omitempty"`
	Created                    string `json:"created,omitempty"`
	Tags                       []struct {
		Projectid    string `json:"projectid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Key          string `json:"key,omitempty"`
		Account      string `json:"account,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Project      string `json:"project,omitempty"`
		Value        string `json:"value,omitempty"`
		Customer     string `json:"customer,omitempty"`
	} `json:"tags,omitempty"`
	Hypervisor          string `json:"hypervisor,omitempty"`
	DiskIopsWriteRate   int    `json:"diskIopsWriteRate,omitempty"`
	DiskBytesWriteRate  int    `json:"diskBytesWriteRate,omitempty"`
	Domainid            string `json:"domainid,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Snapshotid          string `json:"snapshotid,omitempty"`
	Displayvolume       bool   `json:"displayvolume,omitempty"`
	Vmdisplayname       string `json:"vmdisplayname,omitempty"`
	Virtualmachineid    string `json:"virtualmachineid,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Account             string `json:"account,omitempty"`
	Storage             string `json:"storage,omitempty"`
	Attached            string `json:"attached,omitempty"`
	Storageid           string `json:"storageid,omitempty"`
	Project             string `json:"project,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
	Storagetype         string `json:"storagetype,omitempty"`
	Status              string `json:"status,omitempty"`
	Quiescevm           bool   `json:"quiescevm,omitempty"`
	DiskBytesReadRate   int    `json:"diskBytesReadRate,omitempty"`
	Miniops             int    `json:"miniops,omitempty"`
}

type ExtractVolumeParams struct {
	p map[string]interface{}
}

func (p *ExtractVolumeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["mode"]; found {
		u.Set("mode", v.(string))
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ExtractVolumeParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ExtractVolumeParams) SetMode(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["mode"] = v
	return
}

func (p *ExtractVolumeParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
	return
}

func (p *ExtractVolumeParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new ExtractVolumeParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewExtractVolumeParams(id string, mode string, zoneid string) *ExtractVolumeParams {
	p := &ExtractVolumeParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["mode"] = mode
	p.p["zoneid"] = zoneid
	return p
}

// Extracts volume
func (s *VolumeService) ExtractVolume(p *ExtractVolumeParams) (*ExtractVolumeResponse, error) {
	resp, err := s.cs.newRequest("extractVolume", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ExtractVolumeResponse
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

		var r ExtractVolumeResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type ExtractVolumeResponse struct {
	JobID            string `json:"jobid,omitempty"`
	ExtractId        string `json:"extractId,omitempty"`
	State            string `json:"state,omitempty"`
	Zoneid           string `json:"zoneid,omitempty"`
	Storagetype      string `json:"storagetype,omitempty"`
	Id               string `json:"id,omitempty"`
	Zonename         string `json:"zonename,omitempty"`
	Created          string `json:"created,omitempty"`
	ExtractMode      string `json:"extractMode,omitempty"`
	Name             string `json:"name,omitempty"`
	Accountid        string `json:"accountid,omitempty"`
	Resultstring     string `json:"resultstring,omitempty"`
	Status           string `json:"status,omitempty"`
	Uploadpercentage int    `json:"uploadpercentage,omitempty"`
	Url              string `json:"url,omitempty"`
}

type MigrateVolumeParams struct {
	p map[string]interface{}
}

func (p *MigrateVolumeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["livemigrate"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("livemigrate", vv)
	}
	if v, found := p.p["storageid"]; found {
		u.Set("storageid", v.(string))
	}
	if v, found := p.p["volumeid"]; found {
		u.Set("volumeid", v.(string))
	}
	return u
}

func (p *MigrateVolumeParams) SetLivemigrate(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["livemigrate"] = v
	return
}

func (p *MigrateVolumeParams) SetStorageid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["storageid"] = v
	return
}

func (p *MigrateVolumeParams) SetVolumeid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["volumeid"] = v
	return
}

// You should always use this function to get a new MigrateVolumeParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewMigrateVolumeParams(storageid string, volumeid string) *MigrateVolumeParams {
	p := &MigrateVolumeParams{}
	p.p = make(map[string]interface{})
	p.p["storageid"] = storageid
	p.p["volumeid"] = volumeid
	return p
}

// Migrate volume
func (s *VolumeService) MigrateVolume(p *MigrateVolumeParams) (*MigrateVolumeResponse, error) {
	resp, err := s.cs.newRequest("migrateVolume", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r MigrateVolumeResponse
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

		var r MigrateVolumeResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type MigrateVolumeResponse struct {
	JobID                      string `json:"jobid,omitempty"`
	Isextractable              bool   `json:"isextractable,omitempty"`
	Storage                    string `json:"storage,omitempty"`
	Name                       string `json:"name,omitempty"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext,omitempty"`
	Serviceofferingname        string `json:"serviceofferingname,omitempty"`
	Snapshotid                 string `json:"snapshotid,omitempty"`
	Account                    string `json:"account,omitempty"`
	Tags                       []struct {
		Resourcetype string `json:"resourcetype,omitempty"`
		Value        string `json:"value,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Project      string `json:"project,omitempty"`
		Key          string `json:"key,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Account      string `json:"account,omitempty"`
	} `json:"tags,omitempty"`
	Zoneid                  string `json:"zoneid,omitempty"`
	Attached                string `json:"attached,omitempty"`
	Serviceofferingid       string `json:"serviceofferingid,omitempty"`
	DiskBytesReadRate       int    `json:"diskBytesReadRate,omitempty"`
	DiskIopsReadRate        int    `json:"diskIopsReadRate,omitempty"`
	Projectid               string `json:"projectid,omitempty"`
	DiskIopsWriteRate       int    `json:"diskIopsWriteRate,omitempty"`
	Storageid               string `json:"storageid,omitempty"`
	Diskofferingname        string `json:"diskofferingname,omitempty"`
	Status                  string `json:"status,omitempty"`
	Vmdisplayname           string `json:"vmdisplayname,omitempty"`
	Created                 string `json:"created,omitempty"`
	Size                    int    `json:"size,omitempty"`
	DiskBytesWriteRate      int    `json:"diskBytesWriteRate,omitempty"`
	Project                 string `json:"project,omitempty"`
	Destroyed               bool   `json:"destroyed,omitempty"`
	State                   string `json:"state,omitempty"`
	Storagetype             string `json:"storagetype,omitempty"`
	Maxiops                 int    `json:"maxiops,omitempty"`
	Zonename                string `json:"zonename,omitempty"`
	Diskofferingdisplaytext string `json:"diskofferingdisplaytext,omitempty"`
	Type                    string `json:"type,omitempty"`
	Domainid                string `json:"domainid,omitempty"`
	Displayvolume           bool   `json:"displayvolume,omitempty"`
	Diskofferingid          string `json:"diskofferingid,omitempty"`
	Hypervisor              string `json:"hypervisor,omitempty"`
	Vmname                  string `json:"vmname,omitempty"`
	Domain                  string `json:"domain,omitempty"`
	Vmstate                 string `json:"vmstate,omitempty"`
	Deviceid                int    `json:"deviceid,omitempty"`
	Virtualmachineid        string `json:"virtualmachineid,omitempty"`
	Quiescevm               bool   `json:"quiescevm,omitempty"`
	Miniops                 int    `json:"miniops,omitempty"`
	Path                    string `json:"path,omitempty"`
	Id                      string `json:"id,omitempty"`
}

type ResizeVolumeParams struct {
	p map[string]interface{}
}

func (p *ResizeVolumeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["diskofferingid"]; found {
		u.Set("diskofferingid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["shrinkok"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("shrinkok", vv)
	}
	if v, found := p.p["size"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("size", vv)
	}
	return u
}

func (p *ResizeVolumeParams) SetDiskofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["diskofferingid"] = v
	return
}

func (p *ResizeVolumeParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ResizeVolumeParams) SetShrinkok(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["shrinkok"] = v
	return
}

func (p *ResizeVolumeParams) SetSize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["size"] = v
	return
}

// You should always use this function to get a new ResizeVolumeParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewResizeVolumeParams() *ResizeVolumeParams {
	p := &ResizeVolumeParams{}
	p.p = make(map[string]interface{})
	return p
}

// Resizes a volume
func (s *VolumeService) ResizeVolume(p *ResizeVolumeParams) (*ResizeVolumeResponse, error) {
	resp, err := s.cs.newRequest("resizeVolume", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ResizeVolumeResponse
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

		var r ResizeVolumeResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type ResizeVolumeResponse struct {
	JobID                   string `json:"jobid,omitempty"`
	Type                    string `json:"type,omitempty"`
	Diskofferingname        string `json:"diskofferingname,omitempty"`
	Serviceofferingid       string `json:"serviceofferingid,omitempty"`
	Deviceid                int    `json:"deviceid,omitempty"`
	Account                 string `json:"account,omitempty"`
	DiskIopsWriteRate       int    `json:"diskIopsWriteRate,omitempty"`
	Hypervisor              string `json:"hypervisor,omitempty"`
	Quiescevm               bool   `json:"quiescevm,omitempty"`
	DiskBytesWriteRate      int    `json:"diskBytesWriteRate,omitempty"`
	Domainid                string `json:"domainid,omitempty"`
	Zonename                string `json:"zonename,omitempty"`
	Storageid               string `json:"storageid,omitempty"`
	Name                    string `json:"name,omitempty"`
	Vmdisplayname           string `json:"vmdisplayname,omitempty"`
	Snapshotid              string `json:"snapshotid,omitempty"`
	DiskIopsReadRate        int    `json:"diskIopsReadRate,omitempty"`
	Status                  string `json:"status,omitempty"`
	Destroyed               bool   `json:"destroyed,omitempty"`
	Size                    int    `json:"size,omitempty"`
	Path                    string `json:"path,omitempty"`
	Vmstate                 string `json:"vmstate,omitempty"`
	Serviceofferingname     string `json:"serviceofferingname,omitempty"`
	State                   string `json:"state,omitempty"`
	Attached                string `json:"attached,omitempty"`
	Diskofferingdisplaytext string `json:"diskofferingdisplaytext,omitempty"`
	Created                 string `json:"created,omitempty"`
	Domain                  string `json:"domain,omitempty"`
	Tags                    []struct {
		Key          string `json:"key,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Account      string `json:"account,omitempty"`
		Project      string `json:"project,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Value        string `json:"value,omitempty"`
	} `json:"tags,omitempty"`
	Isextractable              bool   `json:"isextractable,omitempty"`
	Storagetype                string `json:"storagetype,omitempty"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext,omitempty"`
	DiskBytesReadRate          int    `json:"diskBytesReadRate,omitempty"`
	Maxiops                    int    `json:"maxiops,omitempty"`
	Projectid                  string `json:"projectid,omitempty"`
	Miniops                    int    `json:"miniops,omitempty"`
	Diskofferingid             string `json:"diskofferingid,omitempty"`
	Project                    string `json:"project,omitempty"`
	Virtualmachineid           string `json:"virtualmachineid,omitempty"`
	Id                         string `json:"id,omitempty"`
	Zoneid                     string `json:"zoneid,omitempty"`
	Storage                    string `json:"storage,omitempty"`
	Displayvolume              bool   `json:"displayvolume,omitempty"`
	Vmname                     string `json:"vmname,omitempty"`
}

type UpdateVolumeParams struct {
	p map[string]interface{}
}

func (p *UpdateVolumeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["displayvolume"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("displayvolume", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["path"]; found {
		u.Set("path", v.(string))
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := p.p["storageid"]; found {
		u.Set("storageid", v.(string))
	}
	return u
}

func (p *UpdateVolumeParams) SetDisplayvolume(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displayvolume"] = v
	return
}

func (p *UpdateVolumeParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *UpdateVolumeParams) SetPath(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["path"] = v
	return
}

func (p *UpdateVolumeParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
	return
}

func (p *UpdateVolumeParams) SetStorageid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["storageid"] = v
	return
}

// You should always use this function to get a new UpdateVolumeParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewUpdateVolumeParams() *UpdateVolumeParams {
	p := &UpdateVolumeParams{}
	p.p = make(map[string]interface{})
	return p
}

// Updates the volume.
func (s *VolumeService) UpdateVolume(p *UpdateVolumeParams) (*UpdateVolumeResponse, error) {
	resp, err := s.cs.newRequest("updateVolume", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateVolumeResponse
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

		var r UpdateVolumeResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type UpdateVolumeResponse struct {
	JobID                      string `json:"jobid,omitempty"`
	Created                    string `json:"created,omitempty"`
	Status                     string `json:"status,omitempty"`
	Isextractable              bool   `json:"isextractable,omitempty"`
	Maxiops                    int    `json:"maxiops,omitempty"`
	Size                       int    `json:"size,omitempty"`
	Diskofferingdisplaytext    string `json:"diskofferingdisplaytext,omitempty"`
	Diskofferingname           string `json:"diskofferingname,omitempty"`
	Quiescevm                  bool   `json:"quiescevm,omitempty"`
	Serviceofferingid          string `json:"serviceofferingid,omitempty"`
	Snapshotid                 string `json:"snapshotid,omitempty"`
	Projectid                  string `json:"projectid,omitempty"`
	Serviceofferingname        string `json:"serviceofferingname,omitempty"`
	Zonename                   string `json:"zonename,omitempty"`
	Domain                     string `json:"domain,omitempty"`
	Destroyed                  bool   `json:"destroyed,omitempty"`
	Diskofferingid             string `json:"diskofferingid,omitempty"`
	Hypervisor                 string `json:"hypervisor,omitempty"`
	Account                    string `json:"account,omitempty"`
	DiskBytesWriteRate         int    `json:"diskBytesWriteRate,omitempty"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext,omitempty"`
	Miniops                    int    `json:"miniops,omitempty"`
	DiskIopsReadRate           int    `json:"diskIopsReadRate,omitempty"`
	Displayvolume              bool   `json:"displayvolume,omitempty"`
	Virtualmachineid           string `json:"virtualmachineid,omitempty"`
	DiskIopsWriteRate          int    `json:"diskIopsWriteRate,omitempty"`
	Storage                    string `json:"storage,omitempty"`
	Zoneid                     string `json:"zoneid,omitempty"`
	Tags                       []struct {
		Value        string `json:"value,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Project      string `json:"project,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Key          string `json:"key,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Account      string `json:"account,omitempty"`
	} `json:"tags,omitempty"`
	Vmdisplayname     string `json:"vmdisplayname,omitempty"`
	Id                string `json:"id,omitempty"`
	DiskBytesReadRate int    `json:"diskBytesReadRate,omitempty"`
	Name              string `json:"name,omitempty"`
	Attached          string `json:"attached,omitempty"`
	Vmstate           string `json:"vmstate,omitempty"`
	Domainid          string `json:"domainid,omitempty"`
	Storagetype       string `json:"storagetype,omitempty"`
	Type              string `json:"type,omitempty"`
	State             string `json:"state,omitempty"`
	Storageid         string `json:"storageid,omitempty"`
	Project           string `json:"project,omitempty"`
	Path              string `json:"path,omitempty"`
	Deviceid          int    `json:"deviceid,omitempty"`
	Vmname            string `json:"vmname,omitempty"`
}
