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

type CreateSnapshotParams struct {
	p map[string]interface{}
}

func (p *CreateSnapshotParams) toURLValues() url.Values {
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
	if v, found := p.p["policyid"]; found {
		u.Set("policyid", v.(string))
	}
	if v, found := p.p["quiescevm"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("quiescevm", vv)
	}
	if v, found := p.p["volumeid"]; found {
		u.Set("volumeid", v.(string))
	}
	return u
}

func (p *CreateSnapshotParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *CreateSnapshotParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *CreateSnapshotParams) SetPolicyid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["policyid"] = v
	return
}

func (p *CreateSnapshotParams) SetQuiescevm(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["quiescevm"] = v
	return
}

func (p *CreateSnapshotParams) SetVolumeid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["volumeid"] = v
	return
}

// You should always use this function to get a new CreateSnapshotParams instance,
// as then you are sure you have configured all required params
func (s *SnapshotService) NewCreateSnapshotParams(volumeid string) *CreateSnapshotParams {
	p := &CreateSnapshotParams{}
	p.p = make(map[string]interface{})
	p.p["volumeid"] = volumeid
	return p
}

// Creates an instant snapshot of a volume.
func (s *SnapshotService) CreateSnapshot(p *CreateSnapshotParams) (*CreateSnapshotResponse, error) {
	resp, err := s.cs.newRequest("createSnapshot", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateSnapshotResponse
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

type CreateSnapshotResponse struct {
	JobID        string `json:"jobid,omitempty"`
	Created      string `json:"created,omitempty"`
	Volumeid     string `json:"volumeid,omitempty"`
	Revertable   bool   `json:"revertable,omitempty"`
	Intervaltype string `json:"intervaltype,omitempty"`
	Account      string `json:"account,omitempty"`
	Zoneid       string `json:"zoneid,omitempty"`
	Snapshottype string `json:"snapshottype,omitempty"`
	Name         string `json:"name,omitempty"`
	Tags         []struct {
		Project      string `json:"project,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Value        string `json:"value,omitempty"`
		Account      string `json:"account,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Key          string `json:"key,omitempty"`
	} `json:"tags,omitempty"`
	State      string `json:"state,omitempty"`
	Volumetype string `json:"volumetype,omitempty"`
	Domainid   string `json:"domainid,omitempty"`
	Domain     string `json:"domain,omitempty"`
	Projectid  string `json:"projectid,omitempty"`
	Id         string `json:"id,omitempty"`
	Project    string `json:"project,omitempty"`
	Volumename string `json:"volumename,omitempty"`
}

type ListSnapshotsParams struct {
	p map[string]interface{}
}

func (p *ListSnapshotsParams) toURLValues() url.Values {
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
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["intervaltype"]; found {
		u.Set("intervaltype", v.(string))
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
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["snapshottype"]; found {
		u.Set("snapshottype", v.(string))
	}
	if v, found := p.p["tags"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["volumeid"]; found {
		u.Set("volumeid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListSnapshotsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListSnapshotsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListSnapshotsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListSnapshotsParams) SetIntervaltype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["intervaltype"] = v
	return
}

func (p *ListSnapshotsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
	return
}

func (p *ListSnapshotsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListSnapshotsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
	return
}

func (p *ListSnapshotsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *ListSnapshotsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListSnapshotsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListSnapshotsParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *ListSnapshotsParams) SetSnapshottype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["snapshottype"] = v
	return
}

func (p *ListSnapshotsParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
	return
}

func (p *ListSnapshotsParams) SetVolumeid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["volumeid"] = v
	return
}

func (p *ListSnapshotsParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new ListSnapshotsParams instance,
// as then you are sure you have configured all required params
func (s *SnapshotService) NewListSnapshotsParams() *ListSnapshotsParams {
	p := &ListSnapshotsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SnapshotService) GetSnapshotID(name string) (string, error) {
	p := &ListSnapshotsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	l, err := s.ListSnapshots(p)
	if err != nil {
		return "", err
	}

	if l.Count == 0 {
		return "", fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Snapshots[0].Id, nil
	}

	if l.Count > 1 {
		for _, v := range l.Snapshots {
			if v.Name == name {
				return v.Id, nil
			}
		}
	}
	return "", fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SnapshotService) GetSnapshotByName(name string) (*Snapshot, int, error) {
	id, err := s.GetSnapshotID(name)
	if err != nil {
		return nil, -1, err
	}

	r, count, err := s.GetSnapshotByID(id)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SnapshotService) GetSnapshotByID(id string) (*Snapshot, int, error) {
	p := &ListSnapshotsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	l, err := s.ListSnapshots(p)
	if err != nil {
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.Snapshots[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Snapshot UUID: %s!", id)
}

// Lists all available snapshots for the account.
func (s *SnapshotService) ListSnapshots(p *ListSnapshotsParams) (*ListSnapshotsResponse, error) {
	resp, err := s.cs.newRequest("listSnapshots", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListSnapshotsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListSnapshotsResponse struct {
	Count     int         `json:"count"`
	Snapshots []*Snapshot `json:"snapshot"`
}

type Snapshot struct {
	Tags []struct {
		Domainid     string `json:"domainid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Project      string `json:"project,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Value        string `json:"value,omitempty"`
		Key          string `json:"key,omitempty"`
		Account      string `json:"account,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
	} `json:"tags,omitempty"`
	Domainid     string `json:"domainid,omitempty"`
	Projectid    string `json:"projectid,omitempty"`
	Revertable   bool   `json:"revertable,omitempty"`
	Volumename   string `json:"volumename,omitempty"`
	Volumeid     string `json:"volumeid,omitempty"`
	Volumetype   string `json:"volumetype,omitempty"`
	Project      string `json:"project,omitempty"`
	Intervaltype string `json:"intervaltype,omitempty"`
	State        string `json:"state,omitempty"`
	Created      string `json:"created,omitempty"`
	Account      string `json:"account,omitempty"`
	Snapshottype string `json:"snapshottype,omitempty"`
	Zoneid       string `json:"zoneid,omitempty"`
	Name         string `json:"name,omitempty"`
	Domain       string `json:"domain,omitempty"`
	Id           string `json:"id,omitempty"`
}

type DeleteSnapshotParams struct {
	p map[string]interface{}
}

func (p *DeleteSnapshotParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteSnapshotParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteSnapshotParams instance,
// as then you are sure you have configured all required params
func (s *SnapshotService) NewDeleteSnapshotParams(id string) *DeleteSnapshotParams {
	p := &DeleteSnapshotParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a snapshot of a disk volume.
func (s *SnapshotService) DeleteSnapshot(p *DeleteSnapshotParams) (*DeleteSnapshotResponse, error) {
	resp, err := s.cs.newRequest("deleteSnapshot", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteSnapshotResponse
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

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type DeleteSnapshotResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type CreateSnapshotPolicyParams struct {
	p map[string]interface{}
}

func (p *CreateSnapshotPolicyParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["intervaltype"]; found {
		u.Set("intervaltype", v.(string))
	}
	if v, found := p.p["maxsnaps"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("maxsnaps", vv)
	}
	if v, found := p.p["schedule"]; found {
		u.Set("schedule", v.(string))
	}
	if v, found := p.p["timezone"]; found {
		u.Set("timezone", v.(string))
	}
	if v, found := p.p["volumeid"]; found {
		u.Set("volumeid", v.(string))
	}
	return u
}

func (p *CreateSnapshotPolicyParams) SetIntervaltype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["intervaltype"] = v
	return
}

func (p *CreateSnapshotPolicyParams) SetMaxsnaps(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["maxsnaps"] = v
	return
}

func (p *CreateSnapshotPolicyParams) SetSchedule(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["schedule"] = v
	return
}

func (p *CreateSnapshotPolicyParams) SetTimezone(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["timezone"] = v
	return
}

func (p *CreateSnapshotPolicyParams) SetVolumeid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["volumeid"] = v
	return
}

// You should always use this function to get a new CreateSnapshotPolicyParams instance,
// as then you are sure you have configured all required params
func (s *SnapshotService) NewCreateSnapshotPolicyParams(intervaltype string, maxsnaps int, schedule string, timezone string, volumeid string) *CreateSnapshotPolicyParams {
	p := &CreateSnapshotPolicyParams{}
	p.p = make(map[string]interface{})
	p.p["intervaltype"] = intervaltype
	p.p["maxsnaps"] = maxsnaps
	p.p["schedule"] = schedule
	p.p["timezone"] = timezone
	p.p["volumeid"] = volumeid
	return p
}

// Creates a snapshot policy for the account.
func (s *SnapshotService) CreateSnapshotPolicy(p *CreateSnapshotPolicyParams) (*CreateSnapshotPolicyResponse, error) {
	resp, err := s.cs.newRequest("createSnapshotPolicy", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateSnapshotPolicyResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type CreateSnapshotPolicyResponse struct {
	Schedule     string `json:"schedule,omitempty"`
	Maxsnaps     int    `json:"maxsnaps,omitempty"`
	Intervaltype int    `json:"intervaltype,omitempty"`
	Volumeid     string `json:"volumeid,omitempty"`
	Timezone     string `json:"timezone,omitempty"`
	Id           string `json:"id,omitempty"`
}

type DeleteSnapshotPoliciesParams struct {
	p map[string]interface{}
}

func (p *DeleteSnapshotPoliciesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["ids"]; found {
		vv := strings.Join(v.([]string), ", ")
		u.Set("ids", vv)
	}
	return u
}

func (p *DeleteSnapshotPoliciesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *DeleteSnapshotPoliciesParams) SetIds(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ids"] = v
	return
}

// You should always use this function to get a new DeleteSnapshotPoliciesParams instance,
// as then you are sure you have configured all required params
func (s *SnapshotService) NewDeleteSnapshotPoliciesParams() *DeleteSnapshotPoliciesParams {
	p := &DeleteSnapshotPoliciesParams{}
	p.p = make(map[string]interface{})
	return p
}

// Deletes snapshot policies for the account.
func (s *SnapshotService) DeleteSnapshotPolicies(p *DeleteSnapshotPoliciesParams) (*DeleteSnapshotPoliciesResponse, error) {
	resp, err := s.cs.newRequest("deleteSnapshotPolicies", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteSnapshotPoliciesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type DeleteSnapshotPoliciesResponse struct {
	Success     string `json:"success,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
}

type ListSnapshotPoliciesParams struct {
	p map[string]interface{}
}

func (p *ListSnapshotPoliciesParams) toURLValues() url.Values {
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
	if v, found := p.p["volumeid"]; found {
		u.Set("volumeid", v.(string))
	}
	return u
}

func (p *ListSnapshotPoliciesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListSnapshotPoliciesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListSnapshotPoliciesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListSnapshotPoliciesParams) SetVolumeid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["volumeid"] = v
	return
}

// You should always use this function to get a new ListSnapshotPoliciesParams instance,
// as then you are sure you have configured all required params
func (s *SnapshotService) NewListSnapshotPoliciesParams(volumeid string) *ListSnapshotPoliciesParams {
	p := &ListSnapshotPoliciesParams{}
	p.p = make(map[string]interface{})
	p.p["volumeid"] = volumeid
	return p
}

// Lists snapshot policies.
func (s *SnapshotService) ListSnapshotPolicies(p *ListSnapshotPoliciesParams) (*ListSnapshotPoliciesResponse, error) {
	resp, err := s.cs.newRequest("listSnapshotPolicies", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListSnapshotPoliciesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListSnapshotPoliciesResponse struct {
	Count            int               `json:"count"`
	SnapshotPolicies []*SnapshotPolicy `json:"snapshotpolicy"`
}

type SnapshotPolicy struct {
	Schedule     string `json:"schedule,omitempty"`
	Timezone     string `json:"timezone,omitempty"`
	Volumeid     string `json:"volumeid,omitempty"`
	Maxsnaps     int    `json:"maxsnaps,omitempty"`
	Id           string `json:"id,omitempty"`
	Intervaltype int    `json:"intervaltype,omitempty"`
}

type RevertSnapshotParams struct {
	p map[string]interface{}
}

func (p *RevertSnapshotParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *RevertSnapshotParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new RevertSnapshotParams instance,
// as then you are sure you have configured all required params
func (s *SnapshotService) NewRevertSnapshotParams(id string) *RevertSnapshotParams {
	p := &RevertSnapshotParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// revert a volume snapshot.
func (s *SnapshotService) RevertSnapshot(p *RevertSnapshotParams) (*RevertSnapshotResponse, error) {
	resp, err := s.cs.newRequest("revertSnapshot", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RevertSnapshotResponse
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

type RevertSnapshotResponse struct {
	JobID        string `json:"jobid,omitempty"`
	Volumetype   string `json:"volumetype,omitempty"`
	Id           string `json:"id,omitempty"`
	Snapshottype string `json:"snapshottype,omitempty"`
	Tags         []struct {
		Domain       string `json:"domain,omitempty"`
		Value        string `json:"value,omitempty"`
		Key          string `json:"key,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Project      string `json:"project,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Account      string `json:"account,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
	} `json:"tags,omitempty"`
	Project      string `json:"project,omitempty"`
	Revertable   bool   `json:"revertable,omitempty"`
	Volumeid     string `json:"volumeid,omitempty"`
	Account      string `json:"account,omitempty"`
	Domainid     string `json:"domainid,omitempty"`
	Intervaltype string `json:"intervaltype,omitempty"`
	Created      string `json:"created,omitempty"`
	State        string `json:"state,omitempty"`
	Name         string `json:"name,omitempty"`
	Domain       string `json:"domain,omitempty"`
	Volumename   string `json:"volumename,omitempty"`
	Zoneid       string `json:"zoneid,omitempty"`
	Projectid    string `json:"projectid,omitempty"`
}

type ListVMSnapshotParams struct {
	p map[string]interface{}
}

func (p *ListVMSnapshotParams) toURLValues() url.Values {
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
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := p.p["tags"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	if v, found := p.p["vmsnapshotid"]; found {
		u.Set("vmsnapshotid", v.(string))
	}
	return u
}

func (p *ListVMSnapshotParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListVMSnapshotParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListVMSnapshotParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
	return
}

func (p *ListVMSnapshotParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListVMSnapshotParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
	return
}

func (p *ListVMSnapshotParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *ListVMSnapshotParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListVMSnapshotParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListVMSnapshotParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *ListVMSnapshotParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
	return
}

func (p *ListVMSnapshotParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
	return
}

func (p *ListVMSnapshotParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
	return
}

func (p *ListVMSnapshotParams) SetVmsnapshotid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vmsnapshotid"] = v
	return
}

// You should always use this function to get a new ListVMSnapshotParams instance,
// as then you are sure you have configured all required params
func (s *SnapshotService) NewListVMSnapshotParams() *ListVMSnapshotParams {
	p := &ListVMSnapshotParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SnapshotService) GetVMSnapshotID(name string) (string, error) {
	p := &ListVMSnapshotParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	l, err := s.ListVMSnapshot(p)
	if err != nil {
		return "", err
	}

	if l.Count == 0 {
		return "", fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.VMSnapshot[0].Id, nil
	}

	if l.Count > 1 {
		for _, v := range l.VMSnapshot {
			if v.Name == name {
				return v.Id, nil
			}
		}
	}
	return "", fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// List virtual machine snapshot by conditions
func (s *SnapshotService) ListVMSnapshot(p *ListVMSnapshotParams) (*ListVMSnapshotResponse, error) {
	resp, err := s.cs.newRequest("listVMSnapshot", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVMSnapshotResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListVMSnapshotResponse struct {
	Count      int           `json:"count"`
	VMSnapshot []*VMSnapshot `json:"vmsnapshot"`
}

type VMSnapshot struct {
	Zoneid           string `json:"zoneid,omitempty"`
	Id               string `json:"id,omitempty"`
	Account          string `json:"account,omitempty"`
	State            string `json:"state,omitempty"`
	Displayname      string `json:"displayname,omitempty"`
	Domain           string `json:"domain,omitempty"`
	Domainid         string `json:"domainid,omitempty"`
	Projectid        string `json:"projectid,omitempty"`
	Parent           string `json:"parent,omitempty"`
	Current          bool   `json:"current,omitempty"`
	Project          string `json:"project,omitempty"`
	Virtualmachineid string `json:"virtualmachineid,omitempty"`
	Created          string `json:"created,omitempty"`
	Description      string `json:"description,omitempty"`
	ParentName       string `json:"parentName,omitempty"`
	Name             string `json:"name,omitempty"`
	Type             string `json:"type,omitempty"`
}

type CreateVMSnapshotParams struct {
	p map[string]interface{}
}

func (p *CreateVMSnapshotParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["quiescevm"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("quiescevm", vv)
	}
	if v, found := p.p["snapshotmemory"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("snapshotmemory", vv)
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *CreateVMSnapshotParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
	return
}

func (p *CreateVMSnapshotParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *CreateVMSnapshotParams) SetQuiescevm(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["quiescevm"] = v
	return
}

func (p *CreateVMSnapshotParams) SetSnapshotmemory(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["snapshotmemory"] = v
	return
}

func (p *CreateVMSnapshotParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
	return
}

// You should always use this function to get a new CreateVMSnapshotParams instance,
// as then you are sure you have configured all required params
func (s *SnapshotService) NewCreateVMSnapshotParams(virtualmachineid string) *CreateVMSnapshotParams {
	p := &CreateVMSnapshotParams{}
	p.p = make(map[string]interface{})
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// Creates snapshot for a vm.
func (s *SnapshotService) CreateVMSnapshot(p *CreateVMSnapshotParams) (*CreateVMSnapshotResponse, error) {
	resp, err := s.cs.newRequest("createVMSnapshot", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateVMSnapshotResponse
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

type CreateVMSnapshotResponse struct {
	JobID            string `json:"jobid,omitempty"`
	Description      string `json:"description,omitempty"`
	Domain           string `json:"domain,omitempty"`
	Id               string `json:"id,omitempty"`
	Created          string `json:"created,omitempty"`
	Displayname      string `json:"displayname,omitempty"`
	Name             string `json:"name,omitempty"`
	State            string `json:"state,omitempty"`
	Project          string `json:"project,omitempty"`
	ParentName       string `json:"parentName,omitempty"`
	Current          bool   `json:"current,omitempty"`
	Zoneid           string `json:"zoneid,omitempty"`
	Virtualmachineid string `json:"virtualmachineid,omitempty"`
	Type             string `json:"type,omitempty"`
	Projectid        string `json:"projectid,omitempty"`
	Account          string `json:"account,omitempty"`
	Parent           string `json:"parent,omitempty"`
	Domainid         string `json:"domainid,omitempty"`
}

type DeleteVMSnapshotParams struct {
	p map[string]interface{}
}

func (p *DeleteVMSnapshotParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["vmsnapshotid"]; found {
		u.Set("vmsnapshotid", v.(string))
	}
	return u
}

func (p *DeleteVMSnapshotParams) SetVmsnapshotid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vmsnapshotid"] = v
	return
}

// You should always use this function to get a new DeleteVMSnapshotParams instance,
// as then you are sure you have configured all required params
func (s *SnapshotService) NewDeleteVMSnapshotParams(vmsnapshotid string) *DeleteVMSnapshotParams {
	p := &DeleteVMSnapshotParams{}
	p.p = make(map[string]interface{})
	p.p["vmsnapshotid"] = vmsnapshotid
	return p
}

// Deletes a vmsnapshot.
func (s *SnapshotService) DeleteVMSnapshot(p *DeleteVMSnapshotParams) (*DeleteVMSnapshotResponse, error) {
	resp, err := s.cs.newRequest("deleteVMSnapshot", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteVMSnapshotResponse
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

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type DeleteVMSnapshotResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type RevertToVMSnapshotParams struct {
	p map[string]interface{}
}

func (p *RevertToVMSnapshotParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["vmsnapshotid"]; found {
		u.Set("vmsnapshotid", v.(string))
	}
	return u
}

func (p *RevertToVMSnapshotParams) SetVmsnapshotid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vmsnapshotid"] = v
	return
}

// You should always use this function to get a new RevertToVMSnapshotParams instance,
// as then you are sure you have configured all required params
func (s *SnapshotService) NewRevertToVMSnapshotParams(vmsnapshotid string) *RevertToVMSnapshotParams {
	p := &RevertToVMSnapshotParams{}
	p.p = make(map[string]interface{})
	p.p["vmsnapshotid"] = vmsnapshotid
	return p
}

// Revert VM from a vmsnapshot.
func (s *SnapshotService) RevertToVMSnapshot(p *RevertToVMSnapshotParams) (*RevertToVMSnapshotResponse, error) {
	resp, err := s.cs.newRequest("revertToVMSnapshot", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RevertToVMSnapshotResponse
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

type RevertToVMSnapshotResponse struct {
	JobID        string `json:"jobid,omitempty"`
	Guestosid    string `json:"guestosid,omitempty"`
	Instancename string `json:"instancename,omitempty"`
	Domainid     string `json:"domainid,omitempty"`
	Templateid   string `json:"templateid,omitempty"`
	Diskkbswrite int    `json:"diskkbswrite,omitempty"`
	Tags         []struct {
		Customer     string `json:"customer,omitempty"`
		Account      string `json:"account,omitempty"`
		Value        string `json:"value,omitempty"`
		Key          string `json:"key,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Project      string `json:"project,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Domain       string `json:"domain,omitempty"`
	} `json:"tags,omitempty"`
	Id                    string            `json:"id,omitempty"`
	Serviceofferingid     string            `json:"serviceofferingid,omitempty"`
	Details               map[string]string `json:"details,omitempty"`
	Diskioread            int               `json:"diskioread,omitempty"`
	Project               string            `json:"project,omitempty"`
	Cpunumber             int               `json:"cpunumber,omitempty"`
	Templatename          string            `json:"templatename,omitempty"`
	Forvirtualnetwork     bool              `json:"forvirtualnetwork,omitempty"`
	Zonename              string            `json:"zonename,omitempty"`
	Keypair               string            `json:"keypair,omitempty"`
	Cpuspeed              int               `json:"cpuspeed,omitempty"`
	Account               string            `json:"account,omitempty"`
	Serviceofferingname   string            `json:"serviceofferingname,omitempty"`
	Name                  string            `json:"name,omitempty"`
	Haenable              bool              `json:"haenable,omitempty"`
	Hostname              string            `json:"hostname,omitempty"`
	Isdynamicallyscalable bool              `json:"isdynamicallyscalable,omitempty"`
	Memory                int               `json:"memory,omitempty"`
	Password              string            `json:"password,omitempty"`
	Displayname           string            `json:"displayname,omitempty"`
	Created               string            `json:"created,omitempty"`
	Displayvm             bool              `json:"displayvm,omitempty"`
	Projectid             string            `json:"projectid,omitempty"`
	Isodisplaytext        string            `json:"isodisplaytext,omitempty"`
	Rootdevicetype        string            `json:"rootdevicetype,omitempty"`
	Group                 string            `json:"group,omitempty"`
	Domain                string            `json:"domain,omitempty"`
	Publicipid            string            `json:"publicipid,omitempty"`
	Servicestate          string            `json:"servicestate,omitempty"`
	Isoid                 string            `json:"isoid,omitempty"`
	Hostid                string            `json:"hostid,omitempty"`
	Affinitygroup         []struct {
		Domainid          string   `json:"domainid,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		Name              string   `json:"name,omitempty"`
		Account           string   `json:"account,omitempty"`
		Description       string   `json:"description,omitempty"`
		Type              string   `json:"type,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Id                string   `json:"id,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Passwordenabled bool `json:"passwordenabled,omitempty"`
	Rootdeviceid    int  `json:"rootdeviceid,omitempty"`
	Nic             []struct {
		Networkname  string   `json:"networkname,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Type         string   `json:"type,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Id           string   `json:"id,omitempty"`
	} `json:"nic,omitempty"`
	Securitygroup []struct {
		Account string `json:"account,omitempty"`
		Project string `json:"project,omitempty"`
		Tags    []struct {
			Resourcetype string `json:"resourcetype,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Value        string `json:"value,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
			Account      string `json:"account,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Project      string `json:"project,omitempty"`
			Key          string `json:"key,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
		} `json:"tags,omitempty"`
		Domainid   string `json:"domainid,omitempty"`
		Egressrule []struct {
			Icmptype          int    `json:"icmptype,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Account           string `json:"account,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
		} `json:"egressrule,omitempty"`
		Id          string `json:"id,omitempty"`
		Name        string `json:"name,omitempty"`
		Projectid   string `json:"projectid,omitempty"`
		Ingressrule []struct {
			Icmptype          int    `json:"icmptype,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Account           string `json:"account,omitempty"`
		} `json:"ingressrule,omitempty"`
		Description string `json:"description,omitempty"`
		Domain      string `json:"domain,omitempty"`
	} `json:"securitygroup,omitempty"`
	State               string `json:"state,omitempty"`
	Templatedisplaytext string `json:"templatedisplaytext,omitempty"`
	Publicip            string `json:"publicip,omitempty"`
	Isoname             string `json:"isoname,omitempty"`
	Zoneid              string `json:"zoneid,omitempty"`
	Networkkbswrite     int    `json:"networkkbswrite,omitempty"`
	Hypervisor          string `json:"hypervisor,omitempty"`
	Diskiowrite         int    `json:"diskiowrite,omitempty"`
	Groupid             string `json:"groupid,omitempty"`
	Cpuused             string `json:"cpuused,omitempty"`
	Diskkbsread         int    `json:"diskkbsread,omitempty"`
	Networkkbsread      int    `json:"networkkbsread,omitempty"`
}
