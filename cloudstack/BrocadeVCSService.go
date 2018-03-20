//
// Copyright 2018, Sander van Harmelen
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
	"net/url"
	"strconv"
)

type AddBrocadeVcsDeviceParams struct {
	p map[string]interface{}
}

func (p *AddBrocadeVcsDeviceParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["hostname"]; found {
		u.Set("hostname", v.(string))
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (p *AddBrocadeVcsDeviceParams) SetHostname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostname"] = v
	return
}

func (p *AddBrocadeVcsDeviceParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
	return
}

func (p *AddBrocadeVcsDeviceParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
	return
}

func (p *AddBrocadeVcsDeviceParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
	return
}

// You should always use this function to get a new AddBrocadeVcsDeviceParams instance,
// as then you are sure you have configured all required params
func (s *BrocadeVCSService) NewAddBrocadeVcsDeviceParams(hostname string, password string, physicalnetworkid string, username string) *AddBrocadeVcsDeviceParams {
	p := &AddBrocadeVcsDeviceParams{}
	p.p = make(map[string]interface{})
	p.p["hostname"] = hostname
	p.p["password"] = password
	p.p["physicalnetworkid"] = physicalnetworkid
	p.p["username"] = username
	return p
}

// Adds a Brocade VCS Switch
func (s *BrocadeVCSService) AddBrocadeVcsDevice(p *AddBrocadeVcsDeviceParams) (*AddBrocadeVcsDeviceResponse, error) {
	resp, err := s.cs.newRequest("addBrocadeVcsDevice", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddBrocadeVcsDeviceResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
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

type AddBrocadeVcsDeviceResponse struct {
	JobID             string `json:"jobid"`
	Brocadedevicename string `json:"brocadedevicename"`
	Hostname          string `json:"hostname"`
	Physicalnetworkid string `json:"physicalnetworkid"`
	Provider          string `json:"provider"`
	Vcsdeviceid       string `json:"vcsdeviceid"`
}

type DeleteBrocadeVcsDeviceParams struct {
	p map[string]interface{}
}

func (p *DeleteBrocadeVcsDeviceParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["vcsdeviceid"]; found {
		u.Set("vcsdeviceid", v.(string))
	}
	return u
}

func (p *DeleteBrocadeVcsDeviceParams) SetVcsdeviceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vcsdeviceid"] = v
	return
}

// You should always use this function to get a new DeleteBrocadeVcsDeviceParams instance,
// as then you are sure you have configured all required params
func (s *BrocadeVCSService) NewDeleteBrocadeVcsDeviceParams(vcsdeviceid string) *DeleteBrocadeVcsDeviceParams {
	p := &DeleteBrocadeVcsDeviceParams{}
	p.p = make(map[string]interface{})
	p.p["vcsdeviceid"] = vcsdeviceid
	return p
}

//  delete a Brocade VCS Switch
func (s *BrocadeVCSService) DeleteBrocadeVcsDevice(p *DeleteBrocadeVcsDeviceParams) (*DeleteBrocadeVcsDeviceResponse, error) {
	resp, err := s.cs.newRequest("deleteBrocadeVcsDevice", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteBrocadeVcsDeviceResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type DeleteBrocadeVcsDeviceResponse struct {
	JobID       string `json:"jobid"`
	Displaytext string `json:"displaytext"`
	Success     bool   `json:"success"`
}

type ListBrocadeVcsDevicesParams struct {
	p map[string]interface{}
}

func (p *ListBrocadeVcsDevicesParams) toURLValues() url.Values {
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
	if v, found := p.p["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	if v, found := p.p["vcsdeviceid"]; found {
		u.Set("vcsdeviceid", v.(string))
	}
	return u
}

func (p *ListBrocadeVcsDevicesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListBrocadeVcsDevicesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListBrocadeVcsDevicesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListBrocadeVcsDevicesParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
	return
}

func (p *ListBrocadeVcsDevicesParams) SetVcsdeviceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vcsdeviceid"] = v
	return
}

// You should always use this function to get a new ListBrocadeVcsDevicesParams instance,
// as then you are sure you have configured all required params
func (s *BrocadeVCSService) NewListBrocadeVcsDevicesParams() *ListBrocadeVcsDevicesParams {
	p := &ListBrocadeVcsDevicesParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists Brocade VCS Switches
func (s *BrocadeVCSService) ListBrocadeVcsDevices(p *ListBrocadeVcsDevicesParams) (*ListBrocadeVcsDevicesResponse, error) {
	resp, err := s.cs.newRequest("listBrocadeVcsDevices", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListBrocadeVcsDevicesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListBrocadeVcsDevicesResponse struct {
	Count             int                 `json:"count"`
	BrocadeVcsDevices []*BrocadeVcsDevice `json:"brocadevcsdevice"`
}

type BrocadeVcsDevice struct {
	Brocadedevicename string `json:"brocadedevicename"`
	Hostname          string `json:"hostname"`
	Physicalnetworkid string `json:"physicalnetworkid"`
	Provider          string `json:"provider"`
	Vcsdeviceid       string `json:"vcsdeviceid"`
}
