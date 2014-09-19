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

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type AttachVolumeResponse struct {
	JobID                      string `json:"jobid,omitempty"`
	Snapshotid                 string `json:"snapshotid,omitempty"`
	Vmdisplayname              string `json:"vmdisplayname,omitempty"`
	Path                       string `json:"path,omitempty"`
	Serviceofferingname        string `json:"serviceofferingname,omitempty"`
	Id                         string `json:"id,omitempty"`
	Quiescevm                  bool   `json:"quiescevm,omitempty"`
	DiskBytesWriteRate         int    `json:"diskBytesWriteRate,omitempty"`
	DiskBytesReadRate          int    `json:"diskBytesReadRate,omitempty"`
	DiskIopsReadRate           int    `json:"diskIopsReadRate,omitempty"`
	Destroyed                  bool   `json:"destroyed,omitempty"`
	Maxiops                    int    `json:"maxiops,omitempty"`
	Storagetype                string `json:"storagetype,omitempty"`
	Zoneid                     string `json:"zoneid,omitempty"`
	Diskofferingdisplaytext    string `json:"diskofferingdisplaytext,omitempty"`
	Hypervisor                 string `json:"hypervisor,omitempty"`
	Status                     string `json:"status,omitempty"`
	Diskofferingname           string `json:"diskofferingname,omitempty"`
	Domainid                   string `json:"domainid,omitempty"`
	Vmname                     string `json:"vmname,omitempty"`
	Displayvolume              bool   `json:"displayvolume,omitempty"`
	Project                    string `json:"project,omitempty"`
	Vmstate                    string `json:"vmstate,omitempty"`
	Attached                   string `json:"attached,omitempty"`
	Projectid                  string `json:"projectid,omitempty"`
	Serviceofferingid          string `json:"serviceofferingid,omitempty"`
	Storageid                  string `json:"storageid,omitempty"`
	Virtualmachineid           string `json:"virtualmachineid,omitempty"`
	Isextractable              bool   `json:"isextractable,omitempty"`
	Account                    string `json:"account,omitempty"`
	Miniops                    int    `json:"miniops,omitempty"`
	Size                       int    `json:"size,omitempty"`
	Type                       string `json:"type,omitempty"`
	Zonename                   string `json:"zonename,omitempty"`
	Domain                     string `json:"domain,omitempty"`
	Diskofferingid             string `json:"diskofferingid,omitempty"`
	Name                       string `json:"name,omitempty"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext,omitempty"`
	DiskIopsWriteRate          int    `json:"diskIopsWriteRate,omitempty"`
	Created                    string `json:"created,omitempty"`
	State                      string `json:"state,omitempty"`
	Tags                       []struct {
		Customer     string `json:"customer,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Value        string `json:"value,omitempty"`
		Project      string `json:"project,omitempty"`
		Key          string `json:"key,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Account      string `json:"account,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
	} `json:"tags,omitempty"`
	Storage  string `json:"storage,omitempty"`
	Deviceid int    `json:"deviceid,omitempty"`
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

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type UploadVolumeResponse struct {
	JobID                   string `json:"jobid,omitempty"`
	DiskIopsReadRate        int    `json:"diskIopsReadRate,omitempty"`
	Serviceofferingname     string `json:"serviceofferingname,omitempty"`
	Account                 string `json:"account,omitempty"`
	Attached                string `json:"attached,omitempty"`
	Displayvolume           bool   `json:"displayvolume,omitempty"`
	Size                    int    `json:"size,omitempty"`
	Storagetype             string `json:"storagetype,omitempty"`
	Domain                  string `json:"domain,omitempty"`
	Path                    string `json:"path,omitempty"`
	Diskofferingdisplaytext string `json:"diskofferingdisplaytext,omitempty"`
	Vmname                  string `json:"vmname,omitempty"`
	Tags                    []struct {
		Resourcetype string `json:"resourcetype,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Project      string `json:"project,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Account      string `json:"account,omitempty"`
		Value        string `json:"value,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Key          string `json:"key,omitempty"`
	} `json:"tags,omitempty"`
	Zonename                   string `json:"zonename,omitempty"`
	Created                    string `json:"created,omitempty"`
	Vmstate                    string `json:"vmstate,omitempty"`
	Serviceofferingid          string `json:"serviceofferingid,omitempty"`
	Snapshotid                 string `json:"snapshotid,omitempty"`
	Hypervisor                 string `json:"hypervisor,omitempty"`
	Storage                    string `json:"storage,omitempty"`
	Name                       string `json:"name,omitempty"`
	Type                       string `json:"type,omitempty"`
	Id                         string `json:"id,omitempty"`
	State                      string `json:"state,omitempty"`
	Project                    string `json:"project,omitempty"`
	DiskBytesReadRate          int    `json:"diskBytesReadRate,omitempty"`
	DiskBytesWriteRate         int    `json:"diskBytesWriteRate,omitempty"`
	Diskofferingid             string `json:"diskofferingid,omitempty"`
	Storageid                  string `json:"storageid,omitempty"`
	Isextractable              bool   `json:"isextractable,omitempty"`
	Maxiops                    int    `json:"maxiops,omitempty"`
	Zoneid                     string `json:"zoneid,omitempty"`
	Destroyed                  bool   `json:"destroyed,omitempty"`
	Status                     string `json:"status,omitempty"`
	Virtualmachineid           string `json:"virtualmachineid,omitempty"`
	Domainid                   string `json:"domainid,omitempty"`
	Quiescevm                  bool   `json:"quiescevm,omitempty"`
	Deviceid                   int    `json:"deviceid,omitempty"`
	Vmdisplayname              string `json:"vmdisplayname,omitempty"`
	Projectid                  string `json:"projectid,omitempty"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext,omitempty"`
	Diskofferingname           string `json:"diskofferingname,omitempty"`
	Miniops                    int    `json:"miniops,omitempty"`
	DiskIopsWriteRate          int    `json:"diskIopsWriteRate,omitempty"`
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

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type DetachVolumeResponse struct {
	JobID                      string `json:"jobid,omitempty"`
	DiskBytesWriteRate         int    `json:"diskBytesWriteRate,omitempty"`
	Status                     string `json:"status,omitempty"`
	Type                       string `json:"type,omitempty"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext,omitempty"`
	Vmname                     string `json:"vmname,omitempty"`
	Domainid                   string `json:"domainid,omitempty"`
	Zonename                   string `json:"zonename,omitempty"`
	Account                    string `json:"account,omitempty"`
	Storageid                  string `json:"storageid,omitempty"`
	Snapshotid                 string `json:"snapshotid,omitempty"`
	DiskIopsReadRate           int    `json:"diskIopsReadRate,omitempty"`
	Serviceofferingname        string `json:"serviceofferingname,omitempty"`
	Attached                   string `json:"attached,omitempty"`
	Id                         string `json:"id,omitempty"`
	Vmstate                    string `json:"vmstate,omitempty"`
	Tags                       []struct {
		Project      string `json:"project,omitempty"`
		Account      string `json:"account,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Key          string `json:"key,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Value        string `json:"value,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
	} `json:"tags,omitempty"`
	Isextractable           bool   `json:"isextractable,omitempty"`
	DiskIopsWriteRate       int    `json:"diskIopsWriteRate,omitempty"`
	Miniops                 int    `json:"miniops,omitempty"`
	Vmdisplayname           string `json:"vmdisplayname,omitempty"`
	Serviceofferingid       string `json:"serviceofferingid,omitempty"`
	Diskofferingname        string `json:"diskofferingname,omitempty"`
	Zoneid                  string `json:"zoneid,omitempty"`
	Project                 string `json:"project,omitempty"`
	Hypervisor              string `json:"hypervisor,omitempty"`
	Projectid               string `json:"projectid,omitempty"`
	Size                    int    `json:"size,omitempty"`
	Path                    string `json:"path,omitempty"`
	Quiescevm               bool   `json:"quiescevm,omitempty"`
	Domain                  string `json:"domain,omitempty"`
	Maxiops                 int    `json:"maxiops,omitempty"`
	Diskofferingid          string `json:"diskofferingid,omitempty"`
	Deviceid                int    `json:"deviceid,omitempty"`
	Displayvolume           bool   `json:"displayvolume,omitempty"`
	Storage                 string `json:"storage,omitempty"`
	Virtualmachineid        string `json:"virtualmachineid,omitempty"`
	Created                 string `json:"created,omitempty"`
	State                   string `json:"state,omitempty"`
	Diskofferingdisplaytext string `json:"diskofferingdisplaytext,omitempty"`
	Storagetype             string `json:"storagetype,omitempty"`
	Destroyed               bool   `json:"destroyed,omitempty"`
	DiskBytesReadRate       int    `json:"diskBytesReadRate,omitempty"`
	Name                    string `json:"name,omitempty"`
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

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type CreateVolumeResponse struct {
	JobID            string `json:"jobid,omitempty"`
	Virtualmachineid string `json:"virtualmachineid,omitempty"`
	Storagetype      string `json:"storagetype,omitempty"`
	Destroyed        bool   `json:"destroyed,omitempty"`
	Domainid         string `json:"domainid,omitempty"`
	Quiescevm        bool   `json:"quiescevm,omitempty"`
	Isextractable    bool   `json:"isextractable,omitempty"`
	Vmstate          string `json:"vmstate,omitempty"`
	Tags             []struct {
		Project      string `json:"project,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Key          string `json:"key,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Value        string `json:"value,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Account      string `json:"account,omitempty"`
	} `json:"tags,omitempty"`
	Serviceofferingid          string `json:"serviceofferingid,omitempty"`
	Diskofferingname           string `json:"diskofferingname,omitempty"`
	Vmdisplayname              string `json:"vmdisplayname,omitempty"`
	DiskBytesReadRate          int    `json:"diskBytesReadRate,omitempty"`
	DiskIopsReadRate           int    `json:"diskIopsReadRate,omitempty"`
	Size                       int    `json:"size,omitempty"`
	DiskIopsWriteRate          int    `json:"diskIopsWriteRate,omitempty"`
	Vmname                     string `json:"vmname,omitempty"`
	Name                       string `json:"name,omitempty"`
	Created                    string `json:"created,omitempty"`
	Maxiops                    int    `json:"maxiops,omitempty"`
	Deviceid                   int    `json:"deviceid,omitempty"`
	Serviceofferingname        string `json:"serviceofferingname,omitempty"`
	DiskBytesWriteRate         int    `json:"diskBytesWriteRate,omitempty"`
	Storageid                  string `json:"storageid,omitempty"`
	Account                    string `json:"account,omitempty"`
	Project                    string `json:"project,omitempty"`
	Type                       string `json:"type,omitempty"`
	State                      string `json:"state,omitempty"`
	Diskofferingdisplaytext    string `json:"diskofferingdisplaytext,omitempty"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext,omitempty"`
	Attached                   string `json:"attached,omitempty"`
	Snapshotid                 string `json:"snapshotid,omitempty"`
	Path                       string `json:"path,omitempty"`
	Storage                    string `json:"storage,omitempty"`
	Miniops                    int    `json:"miniops,omitempty"`
	Projectid                  string `json:"projectid,omitempty"`
	Domain                     string `json:"domain,omitempty"`
	Zonename                   string `json:"zonename,omitempty"`
	Id                         string `json:"id,omitempty"`
	Status                     string `json:"status,omitempty"`
	Hypervisor                 string `json:"hypervisor,omitempty"`
	Diskofferingid             string `json:"diskofferingid,omitempty"`
	Zoneid                     string `json:"zoneid,omitempty"`
	Displayvolume              bool   `json:"displayvolume,omitempty"`
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
	Displaytext string `json:"displaytext,omitempty"`
	Success     string `json:"success,omitempty"`
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

	if l.Count == 0 {
		return "", fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Volumes[0].Id, nil
	}

	if l.Count > 1 {
		for _, v := range l.Volumes {
			if v.Name == name {
				return v.Id, nil
			}
		}
	}
	return "", fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VolumeService) GetVolumeByName(name string) (*Volume, int, error) {
	id, err := s.GetVolumeID(name)
	if err != nil {
		return nil, -1, err
	}

	r, count, err := s.GetVolumeByID(id)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VolumeService) GetVolumeByID(id string) (*Volume, int, error) {
	p := &ListVolumesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	l, err := s.ListVolumes(p)
	if err != nil {
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.Volumes[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Volume UUID: %s!", id)
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
	DiskBytesReadRate int    `json:"diskBytesReadRate,omitempty"`
	Size              int    `json:"size,omitempty"`
	DiskIopsReadRate  int    `json:"diskIopsReadRate,omitempty"`
	Id                string `json:"id,omitempty"`
	Name              string `json:"name,omitempty"`
	Type              string `json:"type,omitempty"`
	Storageid         string `json:"storageid,omitempty"`
	Created           string `json:"created,omitempty"`
	Tags              []struct {
		Projectid    string `json:"projectid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Value        string `json:"value,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Project      string `json:"project,omitempty"`
		Account      string `json:"account,omitempty"`
		Key          string `json:"key,omitempty"`
	} `json:"tags,omitempty"`
	DiskBytesWriteRate         int    `json:"diskBytesWriteRate,omitempty"`
	Attached                   string `json:"attached,omitempty"`
	Path                       string `json:"path,omitempty"`
	Storage                    string `json:"storage,omitempty"`
	Projectid                  string `json:"projectid,omitempty"`
	Storagetype                string `json:"storagetype,omitempty"`
	Vmdisplayname              string `json:"vmdisplayname,omitempty"`
	Snapshotid                 string `json:"snapshotid,omitempty"`
	Miniops                    int    `json:"miniops,omitempty"`
	Domain                     string `json:"domain,omitempty"`
	DiskIopsWriteRate          int    `json:"diskIopsWriteRate,omitempty"`
	Quiescevm                  bool   `json:"quiescevm,omitempty"`
	Diskofferingdisplaytext    string `json:"diskofferingdisplaytext,omitempty"`
	Vmstate                    string `json:"vmstate,omitempty"`
	Destroyed                  bool   `json:"destroyed,omitempty"`
	Zoneid                     string `json:"zoneid,omitempty"`
	Zonename                   string `json:"zonename,omitempty"`
	Hypervisor                 string `json:"hypervisor,omitempty"`
	Virtualmachineid           string `json:"virtualmachineid,omitempty"`
	Displayvolume              bool   `json:"displayvolume,omitempty"`
	State                      string `json:"state,omitempty"`
	Maxiops                    int    `json:"maxiops,omitempty"`
	Vmname                     string `json:"vmname,omitempty"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext,omitempty"`
	Diskofferingname           string `json:"diskofferingname,omitempty"`
	Diskofferingid             string `json:"diskofferingid,omitempty"`
	Serviceofferingid          string `json:"serviceofferingid,omitempty"`
	Account                    string `json:"account,omitempty"`
	Project                    string `json:"project,omitempty"`
	Domainid                   string `json:"domainid,omitempty"`
	Isextractable              bool   `json:"isextractable,omitempty"`
	Status                     string `json:"status,omitempty"`
	Serviceofferingname        string `json:"serviceofferingname,omitempty"`
	Deviceid                   int    `json:"deviceid,omitempty"`
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

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type ExtractVolumeResponse struct {
	JobID            string `json:"jobid,omitempty"`
	Url              string `json:"url,omitempty"`
	ExtractMode      string `json:"extractMode,omitempty"`
	Id               string `json:"id,omitempty"`
	Storagetype      string `json:"storagetype,omitempty"`
	Zonename         string `json:"zonename,omitempty"`
	Resultstring     string `json:"resultstring,omitempty"`
	Accountid        string `json:"accountid,omitempty"`
	ExtractId        string `json:"extractId,omitempty"`
	Uploadpercentage int    `json:"uploadpercentage,omitempty"`
	Created          string `json:"created,omitempty"`
	Name             string `json:"name,omitempty"`
	Status           string `json:"status,omitempty"`
	State            string `json:"state,omitempty"`
	Zoneid           string `json:"zoneid,omitempty"`
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

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type MigrateVolumeResponse struct {
	JobID             string `json:"jobid,omitempty"`
	Status            string `json:"status,omitempty"`
	Vmname            string `json:"vmname,omitempty"`
	Quiescevm         bool   `json:"quiescevm,omitempty"`
	Id                string `json:"id,omitempty"`
	Projectid         string `json:"projectid,omitempty"`
	Attached          string `json:"attached,omitempty"`
	Snapshotid        string `json:"snapshotid,omitempty"`
	Domainid          string `json:"domainid,omitempty"`
	Storage           string `json:"storage,omitempty"`
	Account           string `json:"account,omitempty"`
	Virtualmachineid  string `json:"virtualmachineid,omitempty"`
	Vmdisplayname     string `json:"vmdisplayname,omitempty"`
	Serviceofferingid string `json:"serviceofferingid,omitempty"`
	State             string `json:"state,omitempty"`
	Zoneid            string `json:"zoneid,omitempty"`
	Destroyed         bool   `json:"destroyed,omitempty"`
	Type              string `json:"type,omitempty"`
	Vmstate           string `json:"vmstate,omitempty"`
	Tags              []struct {
		Value        string `json:"value,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Key          string `json:"key,omitempty"`
		Account      string `json:"account,omitempty"`
		Project      string `json:"project,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
	} `json:"tags,omitempty"`
	DiskIopsWriteRate          int    `json:"diskIopsWriteRate,omitempty"`
	Serviceofferingname        string `json:"serviceofferingname,omitempty"`
	Storageid                  string `json:"storageid,omitempty"`
	Maxiops                    int    `json:"maxiops,omitempty"`
	Displayvolume              bool   `json:"displayvolume,omitempty"`
	Project                    string `json:"project,omitempty"`
	Hypervisor                 string `json:"hypervisor,omitempty"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext,omitempty"`
	Diskofferingdisplaytext    string `json:"diskofferingdisplaytext,omitempty"`
	Diskofferingid             string `json:"diskofferingid,omitempty"`
	Deviceid                   int    `json:"deviceid,omitempty"`
	Domain                     string `json:"domain,omitempty"`
	Zonename                   string `json:"zonename,omitempty"`
	DiskBytesReadRate          int    `json:"diskBytesReadRate,omitempty"`
	Size                       int    `json:"size,omitempty"`
	Name                       string `json:"name,omitempty"`
	Path                       string `json:"path,omitempty"`
	Miniops                    int    `json:"miniops,omitempty"`
	Isextractable              bool   `json:"isextractable,omitempty"`
	Created                    string `json:"created,omitempty"`
	Diskofferingname           string `json:"diskofferingname,omitempty"`
	DiskIopsReadRate           int    `json:"diskIopsReadRate,omitempty"`
	DiskBytesWriteRate         int    `json:"diskBytesWriteRate,omitempty"`
	Storagetype                string `json:"storagetype,omitempty"`
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

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type ResizeVolumeResponse struct {
	JobID                      string `json:"jobid,omitempty"`
	Account                    string `json:"account,omitempty"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext,omitempty"`
	Miniops                    int    `json:"miniops,omitempty"`
	Virtualmachineid           string `json:"virtualmachineid,omitempty"`
	DiskIopsWriteRate          int    `json:"diskIopsWriteRate,omitempty"`
	Hypervisor                 string `json:"hypervisor,omitempty"`
	Vmdisplayname              string `json:"vmdisplayname,omitempty"`
	Displayvolume              bool   `json:"displayvolume,omitempty"`
	Attached                   string `json:"attached,omitempty"`
	Snapshotid                 string `json:"snapshotid,omitempty"`
	Size                       int    `json:"size,omitempty"`
	DiskBytesReadRate          int    `json:"diskBytesReadRate,omitempty"`
	Zoneid                     string `json:"zoneid,omitempty"`
	Projectid                  string `json:"projectid,omitempty"`
	Storageid                  string `json:"storageid,omitempty"`
	Serviceofferingid          string `json:"serviceofferingid,omitempty"`
	Domainid                   string `json:"domainid,omitempty"`
	Isextractable              bool   `json:"isextractable,omitempty"`
	Destroyed                  bool   `json:"destroyed,omitempty"`
	Serviceofferingname        string `json:"serviceofferingname,omitempty"`
	Id                         string `json:"id,omitempty"`
	Created                    string `json:"created,omitempty"`
	Diskofferingid             string `json:"diskofferingid,omitempty"`
	Path                       string `json:"path,omitempty"`
	Diskofferingdisplaytext    string `json:"diskofferingdisplaytext,omitempty"`
	Name                       string `json:"name,omitempty"`
	Vmstate                    string `json:"vmstate,omitempty"`
	Diskofferingname           string `json:"diskofferingname,omitempty"`
	Status                     string `json:"status,omitempty"`
	Project                    string `json:"project,omitempty"`
	Storage                    string `json:"storage,omitempty"`
	State                      string `json:"state,omitempty"`
	Vmname                     string `json:"vmname,omitempty"`
	Zonename                   string `json:"zonename,omitempty"`
	DiskIopsReadRate           int    `json:"diskIopsReadRate,omitempty"`
	DiskBytesWriteRate         int    `json:"diskBytesWriteRate,omitempty"`
	Maxiops                    int    `json:"maxiops,omitempty"`
	Quiescevm                  bool   `json:"quiescevm,omitempty"`
	Type                       string `json:"type,omitempty"`
	Tags                       []struct {
		Domainid     string `json:"domainid,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Account      string `json:"account,omitempty"`
		Project      string `json:"project,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Value        string `json:"value,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Key          string `json:"key,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
	} `json:"tags,omitempty"`
	Storagetype string `json:"storagetype,omitempty"`
	Deviceid    int    `json:"deviceid,omitempty"`
	Domain      string `json:"domain,omitempty"`
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

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type UpdateVolumeResponse struct {
	JobID                      string `json:"jobid,omitempty"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext,omitempty"`
	Status                     string `json:"status,omitempty"`
	Type                       string `json:"type,omitempty"`
	Size                       int    `json:"size,omitempty"`
	Zonename                   string `json:"zonename,omitempty"`
	Id                         string `json:"id,omitempty"`
	Serviceofferingid          string `json:"serviceofferingid,omitempty"`
	Projectid                  string `json:"projectid,omitempty"`
	DiskIopsReadRate           int    `json:"diskIopsReadRate,omitempty"`
	Storageid                  string `json:"storageid,omitempty"`
	Maxiops                    int    `json:"maxiops,omitempty"`
	Path                       string `json:"path,omitempty"`
	State                      string `json:"state,omitempty"`
	Zoneid                     string `json:"zoneid,omitempty"`
	Storagetype                string `json:"storagetype,omitempty"`
	Displayvolume              bool   `json:"displayvolume,omitempty"`
	Storage                    string `json:"storage,omitempty"`
	Vmname                     string `json:"vmname,omitempty"`
	Tags                       []struct {
		Customer     string `json:"customer,omitempty"`
		Key          string `json:"key,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Value        string `json:"value,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Account      string `json:"account,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Project      string `json:"project,omitempty"`
	} `json:"tags,omitempty"`
	Account                 string `json:"account,omitempty"`
	Serviceofferingname     string `json:"serviceofferingname,omitempty"`
	Diskofferingdisplaytext string `json:"diskofferingdisplaytext,omitempty"`
	Name                    string `json:"name,omitempty"`
	Miniops                 int    `json:"miniops,omitempty"`
	Project                 string `json:"project,omitempty"`
	Domain                  string `json:"domain,omitempty"`
	DiskBytesWriteRate      int    `json:"diskBytesWriteRate,omitempty"`
	DiskIopsWriteRate       int    `json:"diskIopsWriteRate,omitempty"`
	Virtualmachineid        string `json:"virtualmachineid,omitempty"`
	Destroyed               bool   `json:"destroyed,omitempty"`
	Domainid                string `json:"domainid,omitempty"`
	Diskofferingid          string `json:"diskofferingid,omitempty"`
	Diskofferingname        string `json:"diskofferingname,omitempty"`
	Created                 string `json:"created,omitempty"`
	Isextractable           bool   `json:"isextractable,omitempty"`
	Hypervisor              string `json:"hypervisor,omitempty"`
	Quiescevm               bool   `json:"quiescevm,omitempty"`
	Deviceid                int    `json:"deviceid,omitempty"`
	Vmdisplayname           string `json:"vmdisplayname,omitempty"`
	Attached                string `json:"attached,omitempty"`
	Vmstate                 string `json:"vmstate,omitempty"`
	Snapshotid              string `json:"snapshotid,omitempty"`
	DiskBytesReadRate       int    `json:"diskBytesReadRate,omitempty"`
}
