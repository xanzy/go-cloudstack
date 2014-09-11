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
		vv := strconv.FormatInt(v.(int64), 10)
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
	Maxiops                    int    `json:"maxiops,omitempty"`
	Vmstate                    string `json:"vmstate,omitempty"`
	Miniops                    int    `json:"miniops,omitempty"`
	Status                     string `json:"status,omitempty"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext,omitempty"`
	Size                       int    `json:"size,omitempty"`
	Diskofferingname           string `json:"diskofferingname,omitempty"`
	DiskBytesWriteRate         int    `json:"diskBytesWriteRate,omitempty"`
	Created                    string `json:"created,omitempty"`
	Deviceid                   int    `json:"deviceid,omitempty"`
	Id                         string `json:"id,omitempty"`
	Name                       string `json:"name,omitempty"`
	Isextractable              bool   `json:"isextractable,omitempty"`
	Destroyed                  bool   `json:"destroyed,omitempty"`
	Storagetype                string `json:"storagetype,omitempty"`
	Serviceofferingname        string `json:"serviceofferingname,omitempty"`
	Snapshotid                 string `json:"snapshotid,omitempty"`
	Project                    string `json:"project,omitempty"`
	Storageid                  string `json:"storageid,omitempty"`
	Vmdisplayname              string `json:"vmdisplayname,omitempty"`
	Serviceofferingid          string `json:"serviceofferingid,omitempty"`
	Vmname                     string `json:"vmname,omitempty"`
	Zoneid                     string `json:"zoneid,omitempty"`
	DiskIopsWriteRate          int    `json:"diskIopsWriteRate,omitempty"`
	Zonename                   string `json:"zonename,omitempty"`
	Attached                   string `json:"attached,omitempty"`
	Domainid                   string `json:"domainid,omitempty"`
	Type                       string `json:"type,omitempty"`
	DiskIopsReadRate           int    `json:"diskIopsReadRate,omitempty"`
	Diskofferingdisplaytext    string `json:"diskofferingdisplaytext,omitempty"`
	Account                    string `json:"account,omitempty"`
	Projectid                  string `json:"projectid,omitempty"`
	Virtualmachineid           string `json:"virtualmachineid,omitempty"`
	State                      string `json:"state,omitempty"`
	Storage                    string `json:"storage,omitempty"`
	Domain                     string `json:"domain,omitempty"`
	Quiescevm                  bool   `json:"quiescevm,omitempty"`
	Path                       string `json:"path,omitempty"`
	Hypervisor                 string `json:"hypervisor,omitempty"`
	Tags                       []struct {
		Domainid     string `json:"domainid,omitempty"`
		Account      string `json:"account,omitempty"`
		Project      string `json:"project,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Value        string `json:"value,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Key          string `json:"key,omitempty"`
	} `json:"tags,omitempty"`
	DiskBytesReadRate int    `json:"diskBytesReadRate,omitempty"`
	Displayvolume     bool   `json:"displayvolume,omitempty"`
	Diskofferingid    string `json:"diskofferingid,omitempty"`
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
	JobID             string `json:"jobid,omitempty"`
	Domain            string `json:"domain,omitempty"`
	Project           string `json:"project,omitempty"`
	DiskIopsWriteRate int    `json:"diskIopsWriteRate,omitempty"`
	Projectid         string `json:"projectid,omitempty"`
	Path              string `json:"path,omitempty"`
	Isextractable     bool   `json:"isextractable,omitempty"`
	Hypervisor        string `json:"hypervisor,omitempty"`
	Vmname            string `json:"vmname,omitempty"`
	Deviceid          int    `json:"deviceid,omitempty"`
	Zoneid            string `json:"zoneid,omitempty"`
	Virtualmachineid  string `json:"virtualmachineid,omitempty"`
	Storageid         string `json:"storageid,omitempty"`
	Size              int    `json:"size,omitempty"`
	Destroyed         bool   `json:"destroyed,omitempty"`
	Tags              []struct {
		Project      string `json:"project,omitempty"`
		Account      string `json:"account,omitempty"`
		Key          string `json:"key,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Value        string `json:"value,omitempty"`
	} `json:"tags,omitempty"`
	Storagetype                string `json:"storagetype,omitempty"`
	Diskofferingdisplaytext    string `json:"diskofferingdisplaytext,omitempty"`
	Maxiops                    int    `json:"maxiops,omitempty"`
	Type                       string `json:"type,omitempty"`
	Serviceofferingname        string `json:"serviceofferingname,omitempty"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext,omitempty"`
	Miniops                    int    `json:"miniops,omitempty"`
	Name                       string `json:"name,omitempty"`
	Vmstate                    string `json:"vmstate,omitempty"`
	Quiescevm                  bool   `json:"quiescevm,omitempty"`
	Id                         string `json:"id,omitempty"`
	Attached                   string `json:"attached,omitempty"`
	Vmdisplayname              string `json:"vmdisplayname,omitempty"`
	Snapshotid                 string `json:"snapshotid,omitempty"`
	Storage                    string `json:"storage,omitempty"`
	Serviceofferingid          string `json:"serviceofferingid,omitempty"`
	Status                     string `json:"status,omitempty"`
	Diskofferingname           string `json:"diskofferingname,omitempty"`
	Diskofferingid             string `json:"diskofferingid,omitempty"`
	DiskBytesWriteRate         int    `json:"diskBytesWriteRate,omitempty"`
	Created                    string `json:"created,omitempty"`
	DiskBytesReadRate          int    `json:"diskBytesReadRate,omitempty"`
	Account                    string `json:"account,omitempty"`
	State                      string `json:"state,omitempty"`
	DiskIopsReadRate           int    `json:"diskIopsReadRate,omitempty"`
	Displayvolume              bool   `json:"displayvolume,omitempty"`
	Zonename                   string `json:"zonename,omitempty"`
	Domainid                   string `json:"domainid,omitempty"`
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
		vv := strconv.FormatInt(v.(int64), 10)
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
	JobID                      string `json:"jobid,omitempty"`
	Attached                   string `json:"attached,omitempty"`
	Maxiops                    int    `json:"maxiops,omitempty"`
	Status                     string `json:"status,omitempty"`
	Storage                    string `json:"storage,omitempty"`
	Storageid                  string `json:"storageid,omitempty"`
	Path                       string `json:"path,omitempty"`
	Deviceid                   int    `json:"deviceid,omitempty"`
	Hypervisor                 string `json:"hypervisor,omitempty"`
	Storagetype                string `json:"storagetype,omitempty"`
	DiskBytesReadRate          int    `json:"diskBytesReadRate,omitempty"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext,omitempty"`
	Diskofferingdisplaytext    string `json:"diskofferingdisplaytext,omitempty"`
	Projectid                  string `json:"projectid,omitempty"`
	Domain                     string `json:"domain,omitempty"`
	Project                    string `json:"project,omitempty"`
	Serviceofferingid          string `json:"serviceofferingid,omitempty"`
	Vmname                     string `json:"vmname,omitempty"`
	Account                    string `json:"account,omitempty"`
	Vmdisplayname              string `json:"vmdisplayname,omitempty"`
	Type                       string `json:"type,omitempty"`
	Snapshotid                 string `json:"snapshotid,omitempty"`
	Zoneid                     string `json:"zoneid,omitempty"`
	DiskIopsReadRate           int    `json:"diskIopsReadRate,omitempty"`
	Displayvolume              bool   `json:"displayvolume,omitempty"`
	Vmstate                    string `json:"vmstate,omitempty"`
	Id                         string `json:"id,omitempty"`
	Domainid                   string `json:"domainid,omitempty"`
	Name                       string `json:"name,omitempty"`
	Created                    string `json:"created,omitempty"`
	Virtualmachineid           string `json:"virtualmachineid,omitempty"`
	Quiescevm                  bool   `json:"quiescevm,omitempty"`
	Zonename                   string `json:"zonename,omitempty"`
	Tags                       []struct {
		Value        string `json:"value,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Account      string `json:"account,omitempty"`
		Key          string `json:"key,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Project      string `json:"project,omitempty"`
	} `json:"tags,omitempty"`
	State               string `json:"state,omitempty"`
	Isextractable       bool   `json:"isextractable,omitempty"`
	DiskBytesWriteRate  int    `json:"diskBytesWriteRate,omitempty"`
	Size                int    `json:"size,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Diskofferingname    string `json:"diskofferingname,omitempty"`
	Destroyed           bool   `json:"destroyed,omitempty"`
	DiskIopsWriteRate   int    `json:"diskIopsWriteRate,omitempty"`
	Diskofferingid      string `json:"diskofferingid,omitempty"`
	Miniops             int    `json:"miniops,omitempty"`
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
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("maxiops", vv)
	}
	if v, found := p.p["miniops"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("miniops", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["size"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
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
	Serviceofferingid       string `json:"serviceofferingid,omitempty"`
	Displayvolume           bool   `json:"displayvolume,omitempty"`
	Zonename                string `json:"zonename,omitempty"`
	Hypervisor              string `json:"hypervisor,omitempty"`
	Size                    int    `json:"size,omitempty"`
	Id                      string `json:"id,omitempty"`
	Name                    string `json:"name,omitempty"`
	Projectid               string `json:"projectid,omitempty"`
	Miniops                 int    `json:"miniops,omitempty"`
	Diskofferingdisplaytext string `json:"diskofferingdisplaytext,omitempty"`
	Destroyed               bool   `json:"destroyed,omitempty"`
	Created                 string `json:"created,omitempty"`
	DiskIopsWriteRate       int    `json:"diskIopsWriteRate,omitempty"`
	Project                 string `json:"project,omitempty"`
	Account                 string `json:"account,omitempty"`
	Type                    string `json:"type,omitempty"`
	Maxiops                 int    `json:"maxiops,omitempty"`
	Storagetype             string `json:"storagetype,omitempty"`
	Vmname                  string `json:"vmname,omitempty"`
	DiskIopsReadRate        int    `json:"diskIopsReadRate,omitempty"`
	Status                  string `json:"status,omitempty"`
	Isextractable           bool   `json:"isextractable,omitempty"`
	Tags                    []struct {
		Domain       string `json:"domain,omitempty"`
		Key          string `json:"key,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Project      string `json:"project,omitempty"`
		Value        string `json:"value,omitempty"`
		Account      string `json:"account,omitempty"`
	} `json:"tags,omitempty"`
	DiskBytesWriteRate         int    `json:"diskBytesWriteRate,omitempty"`
	Diskofferingname           string `json:"diskofferingname,omitempty"`
	Zoneid                     string `json:"zoneid,omitempty"`
	Path                       string `json:"path,omitempty"`
	Attached                   string `json:"attached,omitempty"`
	Quiescevm                  bool   `json:"quiescevm,omitempty"`
	Domain                     string `json:"domain,omitempty"`
	DiskBytesReadRate          int    `json:"diskBytesReadRate,omitempty"`
	Serviceofferingname        string `json:"serviceofferingname,omitempty"`
	Diskofferingid             string `json:"diskofferingid,omitempty"`
	State                      string `json:"state,omitempty"`
	Snapshotid                 string `json:"snapshotid,omitempty"`
	Storage                    string `json:"storage,omitempty"`
	Domainid                   string `json:"domainid,omitempty"`
	Storageid                  string `json:"storageid,omitempty"`
	Deviceid                   int    `json:"deviceid,omitempty"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext,omitempty"`
	Vmstate                    string `json:"vmstate,omitempty"`
	Virtualmachineid           string `json:"virtualmachineid,omitempty"`
	Vmdisplayname              string `json:"vmdisplayname,omitempty"`
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
	Tags []struct {
		Key          string `json:"key,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Value        string `json:"value,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Account      string `json:"account,omitempty"`
		Project      string `json:"project,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Customer     string `json:"customer,omitempty"`
	} `json:"tags,omitempty"`
	Vmname                     string `json:"vmname,omitempty"`
	Vmstate                    string `json:"vmstate,omitempty"`
	Projectid                  string `json:"projectid,omitempty"`
	Attached                   string `json:"attached,omitempty"`
	Hypervisor                 string `json:"hypervisor,omitempty"`
	Virtualmachineid           string `json:"virtualmachineid,omitempty"`
	Path                       string `json:"path,omitempty"`
	Zonename                   string `json:"zonename,omitempty"`
	Miniops                    int    `json:"miniops,omitempty"`
	Isextractable              bool   `json:"isextractable,omitempty"`
	Snapshotid                 string `json:"snapshotid,omitempty"`
	DiskIopsWriteRate          int    `json:"diskIopsWriteRate,omitempty"`
	Id                         string `json:"id,omitempty"`
	Status                     string `json:"status,omitempty"`
	Storagetype                string `json:"storagetype,omitempty"`
	Zoneid                     string `json:"zoneid,omitempty"`
	Diskofferingid             string `json:"diskofferingid,omitempty"`
	State                      string `json:"state,omitempty"`
	Account                    string `json:"account,omitempty"`
	Size                       int    `json:"size,omitempty"`
	DiskBytesWriteRate         int    `json:"diskBytesWriteRate,omitempty"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext,omitempty"`
	Diskofferingdisplaytext    string `json:"diskofferingdisplaytext,omitempty"`
	Domain                     string `json:"domain,omitempty"`
	Name                       string `json:"name,omitempty"`
	DiskBytesReadRate          int    `json:"diskBytesReadRate,omitempty"`
	Created                    string `json:"created,omitempty"`
	DiskIopsReadRate           int    `json:"diskIopsReadRate,omitempty"`
	Storageid                  string `json:"storageid,omitempty"`
	Serviceofferingid          string `json:"serviceofferingid,omitempty"`
	Storage                    string `json:"storage,omitempty"`
	Project                    string `json:"project,omitempty"`
	Domainid                   string `json:"domainid,omitempty"`
	Vmdisplayname              string `json:"vmdisplayname,omitempty"`
	Destroyed                  bool   `json:"destroyed,omitempty"`
	Diskofferingname           string `json:"diskofferingname,omitempty"`
	Displayvolume              bool   `json:"displayvolume,omitempty"`
	Quiescevm                  bool   `json:"quiescevm,omitempty"`
	Deviceid                   int    `json:"deviceid,omitempty"`
	Maxiops                    int    `json:"maxiops,omitempty"`
	Serviceofferingname        string `json:"serviceofferingname,omitempty"`
	Type                       string `json:"type,omitempty"`
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
	Id               string `json:"id,omitempty"`
	Name             string `json:"name,omitempty"`
	Status           string `json:"status,omitempty"`
	Resultstring     string `json:"resultstring,omitempty"`
	Uploadpercentage int    `json:"uploadpercentage,omitempty"`
	Zonename         string `json:"zonename,omitempty"`
	Created          string `json:"created,omitempty"`
	State            string `json:"state,omitempty"`
	Storagetype      string `json:"storagetype,omitempty"`
	ExtractId        string `json:"extractId,omitempty"`
	ExtractMode      string `json:"extractMode,omitempty"`
	Accountid        string `json:"accountid,omitempty"`
	Url              string `json:"url,omitempty"`
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
	Type                       string `json:"type,omitempty"`
	Attached                   string `json:"attached,omitempty"`
	Snapshotid                 string `json:"snapshotid,omitempty"`
	Account                    string `json:"account,omitempty"`
	Status                     string `json:"status,omitempty"`
	State                      string `json:"state,omitempty"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext,omitempty"`
	Tags                       []struct {
		Projectid    string `json:"projectid,omitempty"`
		Value        string `json:"value,omitempty"`
		Account      string `json:"account,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Project      string `json:"project,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Key          string `json:"key,omitempty"`
	} `json:"tags,omitempty"`
	Serviceofferingname     string `json:"serviceofferingname,omitempty"`
	DiskBytesReadRate       int    `json:"diskBytesReadRate,omitempty"`
	Virtualmachineid        string `json:"virtualmachineid,omitempty"`
	Storageid               string `json:"storageid,omitempty"`
	Maxiops                 int    `json:"maxiops,omitempty"`
	DiskIopsWriteRate       int    `json:"diskIopsWriteRate,omitempty"`
	Hypervisor              string `json:"hypervisor,omitempty"`
	Projectid               string `json:"projectid,omitempty"`
	Destroyed               bool   `json:"destroyed,omitempty"`
	Name                    string `json:"name,omitempty"`
	Domainid                string `json:"domainid,omitempty"`
	Storage                 string `json:"storage,omitempty"`
	Isextractable           bool   `json:"isextractable,omitempty"`
	Diskofferingdisplaytext string `json:"diskofferingdisplaytext,omitempty"`
	Quiescevm               bool   `json:"quiescevm,omitempty"`
	Created                 string `json:"created,omitempty"`
	Serviceofferingid       string `json:"serviceofferingid,omitempty"`
	Domain                  string `json:"domain,omitempty"`
	Miniops                 int    `json:"miniops,omitempty"`
	DiskIopsReadRate        int    `json:"diskIopsReadRate,omitempty"`
	Diskofferingname        string `json:"diskofferingname,omitempty"`
	Project                 string `json:"project,omitempty"`
	Zonename                string `json:"zonename,omitempty"`
	Vmname                  string `json:"vmname,omitempty"`
	Path                    string `json:"path,omitempty"`
	Deviceid                int    `json:"deviceid,omitempty"`
	Displayvolume           bool   `json:"displayvolume,omitempty"`
	Diskofferingid          string `json:"diskofferingid,omitempty"`
	Storagetype             string `json:"storagetype,omitempty"`
	Zoneid                  string `json:"zoneid,omitempty"`
	Id                      string `json:"id,omitempty"`
	DiskBytesWriteRate      int    `json:"diskBytesWriteRate,omitempty"`
	Size                    int    `json:"size,omitempty"`
	Vmstate                 string `json:"vmstate,omitempty"`
	Vmdisplayname           string `json:"vmdisplayname,omitempty"`
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
		vv := strconv.FormatInt(v.(int64), 10)
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
	JobID                      string `json:"jobid,omitempty"`
	Attached                   string `json:"attached,omitempty"`
	Zonename                   string `json:"zonename,omitempty"`
	Snapshotid                 string `json:"snapshotid,omitempty"`
	Type                       string `json:"type,omitempty"`
	Name                       string `json:"name,omitempty"`
	Deviceid                   int    `json:"deviceid,omitempty"`
	Id                         string `json:"id,omitempty"`
	Diskofferingid             string `json:"diskofferingid,omitempty"`
	Created                    string `json:"created,omitempty"`
	DiskBytesWriteRate         int    `json:"diskBytesWriteRate,omitempty"`
	Displayvolume              bool   `json:"displayvolume,omitempty"`
	Projectid                  string `json:"projectid,omitempty"`
	DiskIopsReadRate           int    `json:"diskIopsReadRate,omitempty"`
	Storage                    string `json:"storage,omitempty"`
	Size                       int    `json:"size,omitempty"`
	Domain                     string `json:"domain,omitempty"`
	Diskofferingname           string `json:"diskofferingname,omitempty"`
	Vmstate                    string `json:"vmstate,omitempty"`
	Diskofferingdisplaytext    string `json:"diskofferingdisplaytext,omitempty"`
	Project                    string `json:"project,omitempty"`
	Storageid                  string `json:"storageid,omitempty"`
	State                      string `json:"state,omitempty"`
	Destroyed                  bool   `json:"destroyed,omitempty"`
	Maxiops                    int    `json:"maxiops,omitempty"`
	Status                     string `json:"status,omitempty"`
	Quiescevm                  bool   `json:"quiescevm,omitempty"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext,omitempty"`
	Zoneid                     string `json:"zoneid,omitempty"`
	Account                    string `json:"account,omitempty"`
	Tags                       []struct {
		Resourceid   string `json:"resourceid,omitempty"`
		Value        string `json:"value,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Project      string `json:"project,omitempty"`
		Key          string `json:"key,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Account      string `json:"account,omitempty"`
	} `json:"tags,omitempty"`
	Domainid            string `json:"domainid,omitempty"`
	DiskBytesReadRate   int    `json:"diskBytesReadRate,omitempty"`
	Vmdisplayname       string `json:"vmdisplayname,omitempty"`
	Hypervisor          string `json:"hypervisor,omitempty"`
	Path                string `json:"path,omitempty"`
	Vmname              string `json:"vmname,omitempty"`
	Miniops             int    `json:"miniops,omitempty"`
	Virtualmachineid    string `json:"virtualmachineid,omitempty"`
	Storagetype         string `json:"storagetype,omitempty"`
	Isextractable       bool   `json:"isextractable,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	DiskIopsWriteRate   int    `json:"diskIopsWriteRate,omitempty"`
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
	Domain                     string `json:"domain,omitempty"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext,omitempty"`
	Id                         string `json:"id,omitempty"`
	Storagetype                string `json:"storagetype,omitempty"`
	Vmname                     string `json:"vmname,omitempty"`
	Attached                   string `json:"attached,omitempty"`
	Miniops                    int    `json:"miniops,omitempty"`
	Account                    string `json:"account,omitempty"`
	Diskofferingid             string `json:"diskofferingid,omitempty"`
	Vmstate                    string `json:"vmstate,omitempty"`
	Zonename                   string `json:"zonename,omitempty"`
	Path                       string `json:"path,omitempty"`
	Vmdisplayname              string `json:"vmdisplayname,omitempty"`
	Diskofferingdisplaytext    string `json:"diskofferingdisplaytext,omitempty"`
	Size                       int    `json:"size,omitempty"`
	Projectid                  string `json:"projectid,omitempty"`
	Name                       string `json:"name,omitempty"`
	Domainid                   string `json:"domainid,omitempty"`
	Created                    string `json:"created,omitempty"`
	Displayvolume              bool   `json:"displayvolume,omitempty"`
	Snapshotid                 string `json:"snapshotid,omitempty"`
	Maxiops                    int    `json:"maxiops,omitempty"`
	Virtualmachineid           string `json:"virtualmachineid,omitempty"`
	Serviceofferingid          string `json:"serviceofferingid,omitempty"`
	Diskofferingname           string `json:"diskofferingname,omitempty"`
	Type                       string `json:"type,omitempty"`
	Hypervisor                 string `json:"hypervisor,omitempty"`
	Storageid                  string `json:"storageid,omitempty"`
	DiskBytesWriteRate         int    `json:"diskBytesWriteRate,omitempty"`
	Storage                    string `json:"storage,omitempty"`
	Isextractable              bool   `json:"isextractable,omitempty"`
	Destroyed                  bool   `json:"destroyed,omitempty"`
	DiskIopsWriteRate          int    `json:"diskIopsWriteRate,omitempty"`
	Quiescevm                  bool   `json:"quiescevm,omitempty"`
	Deviceid                   int    `json:"deviceid,omitempty"`
	Zoneid                     string `json:"zoneid,omitempty"`
	DiskIopsReadRate           int    `json:"diskIopsReadRate,omitempty"`
	Status                     string `json:"status,omitempty"`
	Project                    string `json:"project,omitempty"`
	State                      string `json:"state,omitempty"`
	Serviceofferingname        string `json:"serviceofferingname,omitempty"`
	Tags                       []struct {
		Projectid    string `json:"projectid,omitempty"`
		Key          string `json:"key,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Value        string `json:"value,omitempty"`
		Account      string `json:"account,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Project      string `json:"project,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
	} `json:"tags,omitempty"`
	DiskBytesReadRate int `json:"diskBytesReadRate,omitempty"`
}
