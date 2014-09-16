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

type CreateDiskOfferingParams struct {
	p map[string]interface{}
}

func (p *CreateDiskOfferingParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["bytesreadrate"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("bytesreadrate", vv)
	}
	if v, found := p.p["byteswriterate"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("byteswriterate", vv)
	}
	if v, found := p.p["customized"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("customized", vv)
	}
	if v, found := p.p["customizediops"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("customizediops", vv)
	}
	if v, found := p.p["disksize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("disksize", vv)
	}
	if v, found := p.p["displayoffering"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("displayoffering", vv)
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["hypervisorsnapshotreserve"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("hypervisorsnapshotreserve", vv)
	}
	if v, found := p.p["iopsreadrate"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("iopsreadrate", vv)
	}
	if v, found := p.p["iopswriterate"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("iopswriterate", vv)
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
	if v, found := p.p["storagetype"]; found {
		u.Set("storagetype", v.(string))
	}
	if v, found := p.p["tags"]; found {
		u.Set("tags", v.(string))
	}
	return u
}

func (p *CreateDiskOfferingParams) SetBytesreadrate(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bytesreadrate"] = v
	return
}

func (p *CreateDiskOfferingParams) SetByteswriterate(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["byteswriterate"] = v
	return
}

func (p *CreateDiskOfferingParams) SetCustomized(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customized"] = v
	return
}

func (p *CreateDiskOfferingParams) SetCustomizediops(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customizediops"] = v
	return
}

func (p *CreateDiskOfferingParams) SetDisksize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["disksize"] = v
	return
}

func (p *CreateDiskOfferingParams) SetDisplayoffering(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displayoffering"] = v
	return
}

func (p *CreateDiskOfferingParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
	return
}

func (p *CreateDiskOfferingParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *CreateDiskOfferingParams) SetHypervisorsnapshotreserve(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisorsnapshotreserve"] = v
	return
}

func (p *CreateDiskOfferingParams) SetIopsreadrate(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["iopsreadrate"] = v
	return
}

func (p *CreateDiskOfferingParams) SetIopswriterate(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["iopswriterate"] = v
	return
}

func (p *CreateDiskOfferingParams) SetMaxiops(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["maxiops"] = v
	return
}

func (p *CreateDiskOfferingParams) SetMiniops(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["miniops"] = v
	return
}

func (p *CreateDiskOfferingParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *CreateDiskOfferingParams) SetStoragetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["storagetype"] = v
	return
}

func (p *CreateDiskOfferingParams) SetTags(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
	return
}

// You should always use this function to get a new CreateDiskOfferingParams instance,
// as then you are sure you have configured all required params
func (s *DiskOfferingService) NewCreateDiskOfferingParams(displaytext string, name string) *CreateDiskOfferingParams {
	p := &CreateDiskOfferingParams{}
	p.p = make(map[string]interface{})
	p.p["displaytext"] = displaytext
	p.p["name"] = name
	return p
}

// Creates a disk offering.
func (s *DiskOfferingService) CreateDiskOffering(p *CreateDiskOfferingParams) (*CreateDiskOfferingResponse, error) {
	resp, err := s.cs.newRequest("createDiskOffering", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateDiskOfferingResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type CreateDiskOfferingResponse struct {
	Created            string `json:"created,omitempty"`
	Id                 string `json:"id,omitempty"`
	DiskBytesReadRate  int    `json:"diskBytesReadRate,omitempty"`
	Displaytext        string `json:"displaytext,omitempty"`
	DiskIopsReadRate   int    `json:"diskIopsReadRate,omitempty"`
	Domainid           string `json:"domainid,omitempty"`
	Miniops            int    `json:"miniops,omitempty"`
	Iscustomizediops   bool   `json:"iscustomizediops,omitempty"`
	Maxiops            int    `json:"maxiops,omitempty"`
	Iscustomized       bool   `json:"iscustomized,omitempty"`
	Tags               string `json:"tags,omitempty"`
	Disksize           int    `json:"disksize,omitempty"`
	DiskIopsWriteRate  int    `json:"diskIopsWriteRate,omitempty"`
	DiskBytesWriteRate int    `json:"diskBytesWriteRate,omitempty"`
	Displayoffering    bool   `json:"displayoffering,omitempty"`
	Storagetype        string `json:"storagetype,omitempty"`
	Domain             string `json:"domain,omitempty"`
	Name               string `json:"name,omitempty"`
}

type UpdateDiskOfferingParams struct {
	p map[string]interface{}
}

func (p *UpdateDiskOfferingParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["displayoffering"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("displayoffering", vv)
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["sortkey"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("sortkey", vv)
	}
	return u
}

func (p *UpdateDiskOfferingParams) SetDisplayoffering(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displayoffering"] = v
	return
}

func (p *UpdateDiskOfferingParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
	return
}

func (p *UpdateDiskOfferingParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *UpdateDiskOfferingParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *UpdateDiskOfferingParams) SetSortkey(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sortkey"] = v
	return
}

// You should always use this function to get a new UpdateDiskOfferingParams instance,
// as then you are sure you have configured all required params
func (s *DiskOfferingService) NewUpdateDiskOfferingParams(id string) *UpdateDiskOfferingParams {
	p := &UpdateDiskOfferingParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a disk offering.
func (s *DiskOfferingService) UpdateDiskOffering(p *UpdateDiskOfferingParams) (*UpdateDiskOfferingResponse, error) {
	resp, err := s.cs.newRequest("updateDiskOffering", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateDiskOfferingResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type UpdateDiskOfferingResponse struct {
	Displaytext        string `json:"displaytext,omitempty"`
	DiskBytesReadRate  int    `json:"diskBytesReadRate,omitempty"`
	Disksize           int    `json:"disksize,omitempty"`
	Domain             string `json:"domain,omitempty"`
	Name               string `json:"name,omitempty"`
	Id                 string `json:"id,omitempty"`
	Tags               string `json:"tags,omitempty"`
	Domainid           string `json:"domainid,omitempty"`
	Miniops            int    `json:"miniops,omitempty"`
	Storagetype        string `json:"storagetype,omitempty"`
	Maxiops            int    `json:"maxiops,omitempty"`
	Displayoffering    bool   `json:"displayoffering,omitempty"`
	DiskBytesWriteRate int    `json:"diskBytesWriteRate,omitempty"`
	Created            string `json:"created,omitempty"`
	DiskIopsWriteRate  int    `json:"diskIopsWriteRate,omitempty"`
	DiskIopsReadRate   int    `json:"diskIopsReadRate,omitempty"`
	Iscustomized       bool   `json:"iscustomized,omitempty"`
	Iscustomizediops   bool   `json:"iscustomizediops,omitempty"`
}

type DeleteDiskOfferingParams struct {
	p map[string]interface{}
}

func (p *DeleteDiskOfferingParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteDiskOfferingParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteDiskOfferingParams instance,
// as then you are sure you have configured all required params
func (s *DiskOfferingService) NewDeleteDiskOfferingParams(id string) *DeleteDiskOfferingParams {
	p := &DeleteDiskOfferingParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a disk offering.
func (s *DiskOfferingService) DeleteDiskOffering(p *DeleteDiskOfferingParams) (*DeleteDiskOfferingResponse, error) {
	resp, err := s.cs.newRequest("deleteDiskOffering", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteDiskOfferingResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type DeleteDiskOfferingResponse struct {
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type ListDiskOfferingsParams struct {
	p map[string]interface{}
}

func (p *ListDiskOfferingsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
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
	return u
}

func (p *ListDiskOfferingsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListDiskOfferingsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListDiskOfferingsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListDiskOfferingsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *ListDiskOfferingsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListDiskOfferingsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

// You should always use this function to get a new ListDiskOfferingsParams instance,
// as then you are sure you have configured all required params
func (s *DiskOfferingService) NewListDiskOfferingsParams() *ListDiskOfferingsParams {
	p := &ListDiskOfferingsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *DiskOfferingService) GetDiskOfferingID(name string) (string, error) {
	p := &ListDiskOfferingsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	l, err := s.ListDiskOfferings(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, name, l)
	}
	return l.DiskOfferings[0].Id, nil
}

// Lists all available disk offerings.
func (s *DiskOfferingService) ListDiskOfferings(p *ListDiskOfferingsParams) (*ListDiskOfferingsResponse, error) {
	resp, err := s.cs.newRequest("listDiskOfferings", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListDiskOfferingsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListDiskOfferingsResponse struct {
	Count         int             `json:"count"`
	DiskOfferings []*DiskOffering `json:"diskoffering"`
}

type DiskOffering struct {
	Displayoffering    bool   `json:"displayoffering,omitempty"`
	Disksize           int    `json:"disksize,omitempty"`
	Name               string `json:"name,omitempty"`
	Displaytext        string `json:"displaytext,omitempty"`
	Iscustomized       bool   `json:"iscustomized,omitempty"`
	Miniops            int    `json:"miniops,omitempty"`
	Maxiops            int    `json:"maxiops,omitempty"`
	DiskIopsWriteRate  int    `json:"diskIopsWriteRate,omitempty"`
	DiskBytesReadRate  int    `json:"diskBytesReadRate,omitempty"`
	Domainid           string `json:"domainid,omitempty"`
	Iscustomizediops   bool   `json:"iscustomizediops,omitempty"`
	Id                 string `json:"id,omitempty"`
	Domain             string `json:"domain,omitempty"`
	Created            string `json:"created,omitempty"`
	Tags               string `json:"tags,omitempty"`
	DiskIopsReadRate   int    `json:"diskIopsReadRate,omitempty"`
	DiskBytesWriteRate int    `json:"diskBytesWriteRate,omitempty"`
	Storagetype        string `json:"storagetype,omitempty"`
}
