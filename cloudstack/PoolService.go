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
	"strings"
)

type ListStoragePoolsParams struct {
	p map[string]interface{}
}

func (p *ListStoragePoolsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["ipaddress"]; found {
		u.Set("ipaddress", v.(string))
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
	if v, found := p.p["path"]; found {
		u.Set("path", v.(string))
	}
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := p.p["scope"]; found {
		u.Set("scope", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListStoragePoolsParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
	return
}

func (p *ListStoragePoolsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListStoragePoolsParams) SetIpaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddress"] = v
	return
}

func (p *ListStoragePoolsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListStoragePoolsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *ListStoragePoolsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListStoragePoolsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListStoragePoolsParams) SetPath(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["path"] = v
	return
}

func (p *ListStoragePoolsParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
	return
}

func (p *ListStoragePoolsParams) SetScope(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["scope"] = v
	return
}

func (p *ListStoragePoolsParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new ListStoragePoolsParams instance,
// as then you are sure you have configured all required params
func (s *PoolService) NewListStoragePoolsParams() *ListStoragePoolsParams {
	p := &ListStoragePoolsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *PoolService) GetStoragePoolID(name string) (string, error) {
	p := &ListStoragePoolsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	l, err := s.ListStoragePools(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, name, l)
	}
	return l.StoragePools[0].Id, nil
}

// Lists storage pools.
func (s *PoolService) ListStoragePools(p *ListStoragePoolsParams) (*ListStoragePoolsResponse, error) {
	resp, err := s.cs.newRequest("listStoragePools", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListStoragePoolsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListStoragePoolsResponse struct {
	Count        int            `json:"count"`
	StoragePools []*StoragePool `json:"storagepool"`
}

type StoragePool struct {
	Storagecapabilities  map[string]string `json:"storagecapabilities,omitempty"`
	Tags                 string            `json:"tags,omitempty"`
	Hypervisor           string            `json:"hypervisor,omitempty"`
	Podid                string            `json:"podid,omitempty"`
	Clustername          string            `json:"clustername,omitempty"`
	Disksizetotal        int               `json:"disksizetotal,omitempty"`
	Zoneid               string            `json:"zoneid,omitempty"`
	Clusterid            string            `json:"clusterid,omitempty"`
	Path                 string            `json:"path,omitempty"`
	State                string            `json:"state,omitempty"`
	Capacityiops         int               `json:"capacityiops,omitempty"`
	Suitableformigration bool              `json:"suitableformigration,omitempty"`
	Created              string            `json:"created,omitempty"`
	Type                 string            `json:"type,omitempty"`
	Scope                string            `json:"scope,omitempty"`
	Id                   string            `json:"id,omitempty"`
	Podname              string            `json:"podname,omitempty"`
	Zonename             string            `json:"zonename,omitempty"`
	Disksizeallocated    int               `json:"disksizeallocated,omitempty"`
	Ipaddress            string            `json:"ipaddress,omitempty"`
	Disksizeused         int               `json:"disksizeused,omitempty"`
	Name                 string            `json:"name,omitempty"`
}

type CreateStoragePoolParams struct {
	p map[string]interface{}
}

func (p *CreateStoragePoolParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["capacitybytes"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("capacitybytes", vv)
	}
	if v, found := p.p["capacityiops"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("capacityiops", vv)
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := p.p["details"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("details[%d].key", i), k)
			u.Set(fmt.Sprintf("details[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := p.p["managed"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("managed", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := p.p["provider"]; found {
		u.Set("provider", v.(string))
	}
	if v, found := p.p["scope"]; found {
		u.Set("scope", v.(string))
	}
	if v, found := p.p["tags"]; found {
		u.Set("tags", v.(string))
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *CreateStoragePoolParams) SetCapacitybytes(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["capacitybytes"] = v
	return
}

func (p *CreateStoragePoolParams) SetCapacityiops(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["capacityiops"] = v
	return
}

func (p *CreateStoragePoolParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
	return
}

func (p *CreateStoragePoolParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
	return
}

func (p *CreateStoragePoolParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
	return
}

func (p *CreateStoragePoolParams) SetManaged(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["managed"] = v
	return
}

func (p *CreateStoragePoolParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *CreateStoragePoolParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
	return
}

func (p *CreateStoragePoolParams) SetProvider(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["provider"] = v
	return
}

func (p *CreateStoragePoolParams) SetScope(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["scope"] = v
	return
}

func (p *CreateStoragePoolParams) SetTags(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
	return
}

func (p *CreateStoragePoolParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
	return
}

func (p *CreateStoragePoolParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new CreateStoragePoolParams instance,
// as then you are sure you have configured all required params
func (s *PoolService) NewCreateStoragePoolParams(name string, url string, zoneid string) *CreateStoragePoolParams {
	p := &CreateStoragePoolParams{}
	p.p = make(map[string]interface{})
	p.p["name"] = name
	p.p["url"] = url
	p.p["zoneid"] = zoneid
	return p
}

// Creates a storage pool.
func (s *PoolService) CreateStoragePool(p *CreateStoragePoolParams) (*CreateStoragePoolResponse, error) {
	resp, err := s.cs.newRequest("createStoragePool", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateStoragePoolResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type CreateStoragePoolResponse struct {
	Path                 string            `json:"path,omitempty"`
	Capacityiops         int               `json:"capacityiops,omitempty"`
	Zonename             string            `json:"zonename,omitempty"`
	Storagecapabilities  map[string]string `json:"storagecapabilities,omitempty"`
	Tags                 string            `json:"tags,omitempty"`
	State                string            `json:"state,omitempty"`
	Hypervisor           string            `json:"hypervisor,omitempty"`
	Disksizeallocated    int               `json:"disksizeallocated,omitempty"`
	Podid                string            `json:"podid,omitempty"`
	Ipaddress            string            `json:"ipaddress,omitempty"`
	Type                 string            `json:"type,omitempty"`
	Clusterid            string            `json:"clusterid,omitempty"`
	Id                   string            `json:"id,omitempty"`
	Clustername          string            `json:"clustername,omitempty"`
	Created              string            `json:"created,omitempty"`
	Disksizetotal        int               `json:"disksizetotal,omitempty"`
	Zoneid               string            `json:"zoneid,omitempty"`
	Scope                string            `json:"scope,omitempty"`
	Disksizeused         int               `json:"disksizeused,omitempty"`
	Podname              string            `json:"podname,omitempty"`
	Name                 string            `json:"name,omitempty"`
	Suitableformigration bool              `json:"suitableformigration,omitempty"`
}

type UpdateStoragePoolParams struct {
	p map[string]interface{}
}

func (p *UpdateStoragePoolParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["capacitybytes"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("capacitybytes", vv)
	}
	if v, found := p.p["capacityiops"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("capacityiops", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["tags"]; found {
		vv := strings.Join(v.([]string), ", ")
		u.Set("tags", vv)
	}
	return u
}

func (p *UpdateStoragePoolParams) SetCapacitybytes(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["capacitybytes"] = v
	return
}

func (p *UpdateStoragePoolParams) SetCapacityiops(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["capacityiops"] = v
	return
}

func (p *UpdateStoragePoolParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *UpdateStoragePoolParams) SetTags(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
	return
}

// You should always use this function to get a new UpdateStoragePoolParams instance,
// as then you are sure you have configured all required params
func (s *PoolService) NewUpdateStoragePoolParams(id string) *UpdateStoragePoolParams {
	p := &UpdateStoragePoolParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a storage pool.
func (s *PoolService) UpdateStoragePool(p *UpdateStoragePoolParams) (*UpdateStoragePoolResponse, error) {
	resp, err := s.cs.newRequest("updateStoragePool", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateStoragePoolResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type UpdateStoragePoolResponse struct {
	Storagecapabilities  map[string]string `json:"storagecapabilities,omitempty"`
	Zonename             string            `json:"zonename,omitempty"`
	Podname              string            `json:"podname,omitempty"`
	Podid                string            `json:"podid,omitempty"`
	Tags                 string            `json:"tags,omitempty"`
	Name                 string            `json:"name,omitempty"`
	Hypervisor           string            `json:"hypervisor,omitempty"`
	Ipaddress            string            `json:"ipaddress,omitempty"`
	Suitableformigration bool              `json:"suitableformigration,omitempty"`
	Type                 string            `json:"type,omitempty"`
	Disksizetotal        int               `json:"disksizetotal,omitempty"`
	Disksizeallocated    int               `json:"disksizeallocated,omitempty"`
	Id                   string            `json:"id,omitempty"`
	Disksizeused         int               `json:"disksizeused,omitempty"`
	Path                 string            `json:"path,omitempty"`
	Scope                string            `json:"scope,omitempty"`
	Zoneid               string            `json:"zoneid,omitempty"`
	Clustername          string            `json:"clustername,omitempty"`
	Capacityiops         int               `json:"capacityiops,omitempty"`
	Created              string            `json:"created,omitempty"`
	State                string            `json:"state,omitempty"`
	Clusterid            string            `json:"clusterid,omitempty"`
}

type DeleteStoragePoolParams struct {
	p map[string]interface{}
}

func (p *DeleteStoragePoolParams) toURLValues() url.Values {
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

func (p *DeleteStoragePoolParams) SetForced(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forced"] = v
	return
}

func (p *DeleteStoragePoolParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteStoragePoolParams instance,
// as then you are sure you have configured all required params
func (s *PoolService) NewDeleteStoragePoolParams(id string) *DeleteStoragePoolParams {
	p := &DeleteStoragePoolParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a storage pool.
func (s *PoolService) DeleteStoragePool(p *DeleteStoragePoolParams) (*DeleteStoragePoolResponse, error) {
	resp, err := s.cs.newRequest("deleteStoragePool", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteStoragePoolResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type DeleteStoragePoolResponse struct {
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type FindStoragePoolsForMigrationParams struct {
	p map[string]interface{}
}

func (p *FindStoragePoolsForMigrationParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
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

func (p *FindStoragePoolsForMigrationParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *FindStoragePoolsForMigrationParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *FindStoragePoolsForMigrationParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *FindStoragePoolsForMigrationParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

// You should always use this function to get a new FindStoragePoolsForMigrationParams instance,
// as then you are sure you have configured all required params
func (s *PoolService) NewFindStoragePoolsForMigrationParams(id string) *FindStoragePoolsForMigrationParams {
	p := &FindStoragePoolsForMigrationParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Lists storage pools available for migration of a volume.
func (s *PoolService) FindStoragePoolsForMigration(p *FindStoragePoolsForMigrationParams) (*FindStoragePoolsForMigrationResponse, error) {
	resp, err := s.cs.newRequest("findStoragePoolsForMigration", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r FindStoragePoolsForMigrationResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type FindStoragePoolsForMigrationResponse struct {
	State                string            `json:"state,omitempty"`
	Name                 string            `json:"name,omitempty"`
	Id                   string            `json:"id,omitempty"`
	Zonename             string            `json:"zonename,omitempty"`
	Storagecapabilities  map[string]string `json:"storagecapabilities,omitempty"`
	Tags                 string            `json:"tags,omitempty"`
	Ipaddress            string            `json:"ipaddress,omitempty"`
	Disksizeused         int               `json:"disksizeused,omitempty"`
	Hypervisor           string            `json:"hypervisor,omitempty"`
	Scope                string            `json:"scope,omitempty"`
	Zoneid               string            `json:"zoneid,omitempty"`
	Clusterid            string            `json:"clusterid,omitempty"`
	Created              string            `json:"created,omitempty"`
	Type                 string            `json:"type,omitempty"`
	Disksizeallocated    int               `json:"disksizeallocated,omitempty"`
	Capacityiops         int               `json:"capacityiops,omitempty"`
	Podname              string            `json:"podname,omitempty"`
	Path                 string            `json:"path,omitempty"`
	Clustername          string            `json:"clustername,omitempty"`
	Podid                string            `json:"podid,omitempty"`
	Disksizetotal        int               `json:"disksizetotal,omitempty"`
	Suitableformigration bool              `json:"suitableformigration,omitempty"`
}
