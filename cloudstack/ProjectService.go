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

type CreateProjectParams struct {
	p map[string]interface{}
}

func (p *CreateProjectParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	return u
}

func (p *CreateProjectParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *CreateProjectParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
	return
}

func (p *CreateProjectParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *CreateProjectParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

// You should always use this function to get a new CreateProjectParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewCreateProjectParams(displaytext string, name string) *CreateProjectParams {
	p := &CreateProjectParams{}
	p.p = make(map[string]interface{})
	p.p["displaytext"] = displaytext
	p.p["name"] = name
	return p
}

// Creates a project
func (s *ProjectService) CreateProject(p *CreateProjectParams) (*CreateProjectResponse, error) {
	resp, err := s.cs.newRequest("createProject", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateProjectResponse
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

type CreateProjectResponse struct {
	JobID                   string `json:"jobid,omitempty"`
	Secondarystoragelimit   string `json:"secondarystoragelimit,omitempty"`
	Vmtotal                 int    `json:"vmtotal,omitempty"`
	Networklimit            string `json:"networklimit,omitempty"`
	Primarystoragetotal     int    `json:"primarystoragetotal,omitempty"`
	Volumetotal             int    `json:"volumetotal,omitempty"`
	Domainid                string `json:"domainid,omitempty"`
	Snapshotlimit           string `json:"snapshotlimit,omitempty"`
	Vmlimit                 string `json:"vmlimit,omitempty"`
	Account                 string `json:"account,omitempty"`
	Id                      string `json:"id,omitempty"`
	Templatetotal           int    `json:"templatetotal,omitempty"`
	Ipavailable             string `json:"ipavailable,omitempty"`
	Secondarystoragetotal   int    `json:"secondarystoragetotal,omitempty"`
	Volumelimit             string `json:"volumelimit,omitempty"`
	Vpclimit                string `json:"vpclimit,omitempty"`
	Primarystorageavailable string `json:"primarystorageavailable,omitempty"`
	Cpulimit                string `json:"cpulimit,omitempty"`
	Templateavailable       string `json:"templateavailable,omitempty"`
	Name                    string `json:"name,omitempty"`
	Cpuavailable            string `json:"cpuavailable,omitempty"`
	Tags                    []struct {
		Projectid    string `json:"projectid,omitempty"`
		Value        string `json:"value,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Key          string `json:"key,omitempty"`
		Project      string `json:"project,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Account      string `json:"account,omitempty"`
	} `json:"tags,omitempty"`
	Vpctotal                  int    `json:"vpctotal,omitempty"`
	Memorytotal               int    `json:"memorytotal,omitempty"`
	Vpcavailable              string `json:"vpcavailable,omitempty"`
	State                     string `json:"state,omitempty"`
	Vmrunning                 int    `json:"vmrunning,omitempty"`
	Volumeavailable           string `json:"volumeavailable,omitempty"`
	Displaytext               string `json:"displaytext,omitempty"`
	Snapshotavailable         string `json:"snapshotavailable,omitempty"`
	Memorylimit               string `json:"memorylimit,omitempty"`
	Snapshottotal             int    `json:"snapshottotal,omitempty"`
	Networktotal              int    `json:"networktotal,omitempty"`
	Cputotal                  int    `json:"cputotal,omitempty"`
	Primarystoragelimit       string `json:"primarystoragelimit,omitempty"`
	Templatelimit             string `json:"templatelimit,omitempty"`
	Memoryavailable           string `json:"memoryavailable,omitempty"`
	Iptotal                   int    `json:"iptotal,omitempty"`
	Vmavailable               string `json:"vmavailable,omitempty"`
	Networkavailable          string `json:"networkavailable,omitempty"`
	Secondarystorageavailable string `json:"secondarystorageavailable,omitempty"`
	Iplimit                   string `json:"iplimit,omitempty"`
	Vmstopped                 int    `json:"vmstopped,omitempty"`
	Domain                    string `json:"domain,omitempty"`
}

type DeleteProjectParams struct {
	p map[string]interface{}
}

func (p *DeleteProjectParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteProjectParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteProjectParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewDeleteProjectParams(id string) *DeleteProjectParams {
	p := &DeleteProjectParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a project
func (s *ProjectService) DeleteProject(p *DeleteProjectParams) (*DeleteProjectResponse, error) {
	resp, err := s.cs.newRequest("deleteProject", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteProjectResponse
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

type DeleteProjectResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Success     bool   `json:"success,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
}

type UpdateProjectParams struct {
	p map[string]interface{}
}

func (p *UpdateProjectParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *UpdateProjectParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *UpdateProjectParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
	return
}

func (p *UpdateProjectParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new UpdateProjectParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewUpdateProjectParams(id string) *UpdateProjectParams {
	p := &UpdateProjectParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a project
func (s *ProjectService) UpdateProject(p *UpdateProjectParams) (*UpdateProjectResponse, error) {
	resp, err := s.cs.newRequest("updateProject", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateProjectResponse
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

type UpdateProjectResponse struct {
	JobID    string `json:"jobid,omitempty"`
	Vpctotal int    `json:"vpctotal,omitempty"`
	Tags     []struct {
		Resourcetype string `json:"resourcetype,omitempty"`
		Value        string `json:"value,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Project      string `json:"project,omitempty"`
		Key          string `json:"key,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Account      string `json:"account,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Customer     string `json:"customer,omitempty"`
	} `json:"tags,omitempty"`
	Memorylimit               string `json:"memorylimit,omitempty"`
	Secondarystoragetotal     int    `json:"secondarystoragetotal,omitempty"`
	Vmstopped                 int    `json:"vmstopped,omitempty"`
	Cpuavailable              string `json:"cpuavailable,omitempty"`
	Memoryavailable           string `json:"memoryavailable,omitempty"`
	Iptotal                   int    `json:"iptotal,omitempty"`
	Secondarystorageavailable string `json:"secondarystorageavailable,omitempty"`
	Vmtotal                   int    `json:"vmtotal,omitempty"`
	Snapshotavailable         string `json:"snapshotavailable,omitempty"`
	Iplimit                   string `json:"iplimit,omitempty"`
	Account                   string `json:"account,omitempty"`
	Vpclimit                  string `json:"vpclimit,omitempty"`
	Domain                    string `json:"domain,omitempty"`
	Templatelimit             string `json:"templatelimit,omitempty"`
	Volumelimit               string `json:"volumelimit,omitempty"`
	Templateavailable         string `json:"templateavailable,omitempty"`
	Secondarystoragelimit     string `json:"secondarystoragelimit,omitempty"`
	Snapshotlimit             string `json:"snapshotlimit,omitempty"`
	Vmlimit                   string `json:"vmlimit,omitempty"`
	Networklimit              string `json:"networklimit,omitempty"`
	Primarystoragetotal       int    `json:"primarystoragetotal,omitempty"`
	Templatetotal             int    `json:"templatetotal,omitempty"`
	Vpcavailable              string `json:"vpcavailable,omitempty"`
	Cputotal                  int    `json:"cputotal,omitempty"`
	Ipavailable               string `json:"ipavailable,omitempty"`
	Volumeavailable           string `json:"volumeavailable,omitempty"`
	Vmavailable               string `json:"vmavailable,omitempty"`
	Networktotal              int    `json:"networktotal,omitempty"`
	Memorytotal               int    `json:"memorytotal,omitempty"`
	Primarystorageavailable   string `json:"primarystorageavailable,omitempty"`
	Vmrunning                 int    `json:"vmrunning,omitempty"`
	Primarystoragelimit       string `json:"primarystoragelimit,omitempty"`
	Name                      string `json:"name,omitempty"`
	Id                        string `json:"id,omitempty"`
	Cpulimit                  string `json:"cpulimit,omitempty"`
	Domainid                  string `json:"domainid,omitempty"`
	Snapshottotal             int    `json:"snapshottotal,omitempty"`
	Networkavailable          string `json:"networkavailable,omitempty"`
	Displaytext               string `json:"displaytext,omitempty"`
	State                     string `json:"state,omitempty"`
	Volumetotal               int    `json:"volumetotal,omitempty"`
}

type ActivateProjectParams struct {
	p map[string]interface{}
}

func (p *ActivateProjectParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *ActivateProjectParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new ActivateProjectParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewActivateProjectParams(id string) *ActivateProjectParams {
	p := &ActivateProjectParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Activates a project
func (s *ProjectService) ActivateProject(p *ActivateProjectParams) (*ActivateProjectResponse, error) {
	resp, err := s.cs.newRequest("activateProject", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ActivateProjectResponse
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

type ActivateProjectResponse struct {
	JobID                     string `json:"jobid,omitempty"`
	Snapshotlimit             string `json:"snapshotlimit,omitempty"`
	Templatetotal             int    `json:"templatetotal,omitempty"`
	Ipavailable               string `json:"ipavailable,omitempty"`
	Vmlimit                   string `json:"vmlimit,omitempty"`
	Vpclimit                  string `json:"vpclimit,omitempty"`
	Vmtotal                   int    `json:"vmtotal,omitempty"`
	Templateavailable         string `json:"templateavailable,omitempty"`
	Snapshottotal             int    `json:"snapshottotal,omitempty"`
	Networktotal              int    `json:"networktotal,omitempty"`
	Account                   string `json:"account,omitempty"`
	Memorytotal               int    `json:"memorytotal,omitempty"`
	Domainid                  string `json:"domainid,omitempty"`
	Primarystoragelimit       string `json:"primarystoragelimit,omitempty"`
	Cpuavailable              string `json:"cpuavailable,omitempty"`
	Primarystoragetotal       int    `json:"primarystoragetotal,omitempty"`
	Name                      string `json:"name,omitempty"`
	Domain                    string `json:"domain,omitempty"`
	Vpcavailable              string `json:"vpcavailable,omitempty"`
	Vmavailable               string `json:"vmavailable,omitempty"`
	Volumelimit               string `json:"volumelimit,omitempty"`
	Memoryavailable           string `json:"memoryavailable,omitempty"`
	Primarystorageavailable   string `json:"primarystorageavailable,omitempty"`
	Templatelimit             string `json:"templatelimit,omitempty"`
	Secondarystoragetotal     int    `json:"secondarystoragetotal,omitempty"`
	Networkavailable          string `json:"networkavailable,omitempty"`
	State                     string `json:"state,omitempty"`
	Vmrunning                 int    `json:"vmrunning,omitempty"`
	Id                        string `json:"id,omitempty"`
	Cputotal                  int    `json:"cputotal,omitempty"`
	Secondarystoragelimit     string `json:"secondarystoragelimit,omitempty"`
	Cpulimit                  string `json:"cpulimit,omitempty"`
	Iptotal                   int    `json:"iptotal,omitempty"`
	Snapshotavailable         string `json:"snapshotavailable,omitempty"`
	Secondarystorageavailable string `json:"secondarystorageavailable,omitempty"`
	Tags                      []struct {
		Project      string `json:"project,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Value        string `json:"value,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Account      string `json:"account,omitempty"`
		Key          string `json:"key,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
	} `json:"tags,omitempty"`
	Iplimit         string `json:"iplimit,omitempty"`
	Volumeavailable string `json:"volumeavailable,omitempty"`
	Volumetotal     int    `json:"volumetotal,omitempty"`
	Vmstopped       int    `json:"vmstopped,omitempty"`
	Networklimit    string `json:"networklimit,omitempty"`
	Vpctotal        int    `json:"vpctotal,omitempty"`
	Displaytext     string `json:"displaytext,omitempty"`
	Memorylimit     string `json:"memorylimit,omitempty"`
}

type SuspendProjectParams struct {
	p map[string]interface{}
}

func (p *SuspendProjectParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *SuspendProjectParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new SuspendProjectParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewSuspendProjectParams(id string) *SuspendProjectParams {
	p := &SuspendProjectParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Suspends a project
func (s *ProjectService) SuspendProject(p *SuspendProjectParams) (*SuspendProjectResponse, error) {
	resp, err := s.cs.newRequest("suspendProject", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r SuspendProjectResponse
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

type SuspendProjectResponse struct {
	JobID                   string `json:"jobid,omitempty"`
	Account                 string `json:"account,omitempty"`
	Secondarystoragelimit   string `json:"secondarystoragelimit,omitempty"`
	Domain                  string `json:"domain,omitempty"`
	Vpclimit                string `json:"vpclimit,omitempty"`
	Primarystorageavailable string `json:"primarystorageavailable,omitempty"`
	Templatelimit           string `json:"templatelimit,omitempty"`
	Memorylimit             string `json:"memorylimit,omitempty"`
	Primarystoragetotal     int    `json:"primarystoragetotal,omitempty"`
	Snapshotlimit           string `json:"snapshotlimit,omitempty"`
	Vpcavailable            string `json:"vpcavailable,omitempty"`
	Vmrunning               int    `json:"vmrunning,omitempty"`
	Id                      string `json:"id,omitempty"`
	Volumelimit             string `json:"volumelimit,omitempty"`
	Iptotal                 int    `json:"iptotal,omitempty"`
	Networktotal            int    `json:"networktotal,omitempty"`
	Cpulimit                string `json:"cpulimit,omitempty"`
	Templateavailable       string `json:"templateavailable,omitempty"`
	Vpctotal                int    `json:"vpctotal,omitempty"`
	Cpuavailable            string `json:"cpuavailable,omitempty"`
	State                   string `json:"state,omitempty"`
	Networkavailable        string `json:"networkavailable,omitempty"`
	Displaytext             string `json:"displaytext,omitempty"`
	Templatetotal           int    `json:"templatetotal,omitempty"`
	Domainid                string `json:"domainid,omitempty"`
	Memoryavailable         string `json:"memoryavailable,omitempty"`
	Vmstopped               int    `json:"vmstopped,omitempty"`
	Vmavailable             string `json:"vmavailable,omitempty"`
	Tags                    []struct {
		Resourcetype string `json:"resourcetype,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Key          string `json:"key,omitempty"`
		Project      string `json:"project,omitempty"`
		Value        string `json:"value,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Account      string `json:"account,omitempty"`
		Domain       string `json:"domain,omitempty"`
	} `json:"tags,omitempty"`
	Volumeavailable           string `json:"volumeavailable,omitempty"`
	Vmtotal                   int    `json:"vmtotal,omitempty"`
	Secondarystoragetotal     int    `json:"secondarystoragetotal,omitempty"`
	Secondarystorageavailable string `json:"secondarystorageavailable,omitempty"`
	Volumetotal               int    `json:"volumetotal,omitempty"`
	Iplimit                   string `json:"iplimit,omitempty"`
	Snapshotavailable         string `json:"snapshotavailable,omitempty"`
	Vmlimit                   string `json:"vmlimit,omitempty"`
	Ipavailable               string `json:"ipavailable,omitempty"`
	Cputotal                  int    `json:"cputotal,omitempty"`
	Networklimit              string `json:"networklimit,omitempty"`
	Name                      string `json:"name,omitempty"`
	Snapshottotal             int    `json:"snapshottotal,omitempty"`
	Memorytotal               int    `json:"memorytotal,omitempty"`
	Primarystoragelimit       string `json:"primarystoragelimit,omitempty"`
}

type ListProjectsParams struct {
	p map[string]interface{}
}

func (p *ListProjectsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
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
	return u
}

func (p *ListProjectsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListProjectsParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
	return
}

func (p *ListProjectsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListProjectsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListProjectsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
	return
}

func (p *ListProjectsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListProjectsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
	return
}

func (p *ListProjectsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *ListProjectsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListProjectsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListProjectsParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
	return
}

func (p *ListProjectsParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
	return
}

// You should always use this function to get a new ListProjectsParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewListProjectsParams() *ListProjectsParams {
	p := &ListProjectsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ProjectService) GetProjectID(name string) (string, error) {
	p := &ListProjectsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	l, err := s.ListProjects(p)
	if err != nil {
		return "", err
	}

	if l.Count == 0 {
		return "", fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Projects[0].Id, nil
	}

	if l.Count > 1 {
		for _, v := range l.Projects {
			if v.Name == name {
				return v.Id, nil
			}
		}
	}
	return "", fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ProjectService) GetProjectByName(name string) (*Project, int, error) {
	id, err := s.GetProjectID(name)
	if err != nil {
		return nil, -1, err
	}

	r, count, err := s.GetProjectByID(id)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ProjectService) GetProjectByID(id string) (*Project, int, error) {
	p := &ListProjectsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	l, err := s.ListProjects(p)
	if err != nil {
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.Projects[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Project UUID: %s!", id)
}

// Lists projects and provides detailed information for listed projects
func (s *ProjectService) ListProjects(p *ListProjectsParams) (*ListProjectsResponse, error) {
	resp, err := s.cs.newRequest("listProjects", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListProjectsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListProjectsResponse struct {
	Count    int        `json:"count"`
	Projects []*Project `json:"project"`
}

type Project struct {
	Primarystoragelimit string `json:"primarystoragelimit,omitempty"`
	Account             string `json:"account,omitempty"`
	Primarystoragetotal int    `json:"primarystoragetotal,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Id                  string `json:"id,omitempty"`
	Volumelimit         string `json:"volumelimit,omitempty"`
	Vmlimit             string `json:"vmlimit,omitempty"`
	Templatetotal       int    `json:"templatetotal,omitempty"`
	Displaytext         string `json:"displaytext,omitempty"`
	Vmstopped           int    `json:"vmstopped,omitempty"`
	Volumeavailable     string `json:"volumeavailable,omitempty"`
	Templateavailable   string `json:"templateavailable,omitempty"`
	Snapshottotal       int    `json:"snapshottotal,omitempty"`
	Vpcavailable        string `json:"vpcavailable,omitempty"`
	Iplimit             string `json:"iplimit,omitempty"`
	Memorylimit         string `json:"memorylimit,omitempty"`
	Vpctotal            int    `json:"vpctotal,omitempty"`
	Templatelimit       string `json:"templatelimit,omitempty"`
	Domainid            string `json:"domainid,omitempty"`
	Vmtotal             int    `json:"vmtotal,omitempty"`
	Networktotal        int    `json:"networktotal,omitempty"`
	Tags                []struct {
		Domainid     string `json:"domainid,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Value        string `json:"value,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Account      string `json:"account,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Key          string `json:"key,omitempty"`
		Project      string `json:"project,omitempty"`
	} `json:"tags,omitempty"`
	Secondarystoragelimit     string `json:"secondarystoragelimit,omitempty"`
	State                     string `json:"state,omitempty"`
	Memoryavailable           string `json:"memoryavailable,omitempty"`
	Ipavailable               string `json:"ipavailable,omitempty"`
	Snapshotlimit             string `json:"snapshotlimit,omitempty"`
	Networkavailable          string `json:"networkavailable,omitempty"`
	Name                      string `json:"name,omitempty"`
	Networklimit              string `json:"networklimit,omitempty"`
	Cpulimit                  string `json:"cpulimit,omitempty"`
	Cpuavailable              string `json:"cpuavailable,omitempty"`
	Vmrunning                 int    `json:"vmrunning,omitempty"`
	Snapshotavailable         string `json:"snapshotavailable,omitempty"`
	Vmavailable               string `json:"vmavailable,omitempty"`
	Secondarystorageavailable string `json:"secondarystorageavailable,omitempty"`
	Volumetotal               int    `json:"volumetotal,omitempty"`
	Primarystorageavailable   string `json:"primarystorageavailable,omitempty"`
	Vpclimit                  string `json:"vpclimit,omitempty"`
	Cputotal                  int    `json:"cputotal,omitempty"`
	Secondarystoragetotal     int    `json:"secondarystoragetotal,omitempty"`
	Iptotal                   int    `json:"iptotal,omitempty"`
	Memorytotal               int    `json:"memorytotal,omitempty"`
}

type ListProjectInvitationsParams struct {
	p map[string]interface{}
}

func (p *ListProjectInvitationsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["activeonly"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("activeonly", vv)
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
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
	return u
}

func (p *ListProjectInvitationsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListProjectInvitationsParams) SetActiveonly(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["activeonly"] = v
	return
}

func (p *ListProjectInvitationsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListProjectInvitationsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListProjectInvitationsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
	return
}

func (p *ListProjectInvitationsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListProjectInvitationsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
	return
}

func (p *ListProjectInvitationsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListProjectInvitationsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListProjectInvitationsParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *ListProjectInvitationsParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
	return
}

// You should always use this function to get a new ListProjectInvitationsParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewListProjectInvitationsParams() *ListProjectInvitationsParams {
	p := &ListProjectInvitationsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ProjectService) GetProjectInvitationByID(id string) (*ProjectInvitation, int, error) {
	p := &ListProjectInvitationsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	l, err := s.ListProjectInvitations(p)
	if err != nil {
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.ProjectInvitations[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for ProjectInvitation UUID: %s!", id)
}

// Lists projects and provides detailed information for listed projects
func (s *ProjectService) ListProjectInvitations(p *ListProjectInvitationsParams) (*ListProjectInvitationsResponse, error) {
	resp, err := s.cs.newRequest("listProjectInvitations", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListProjectInvitationsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListProjectInvitationsResponse struct {
	Count              int                  `json:"count"`
	ProjectInvitations []*ProjectInvitation `json:"projectinvitation"`
}

type ProjectInvitation struct {
	State     string `json:"state,omitempty"`
	Account   string `json:"account,omitempty"`
	Domain    string `json:"domain,omitempty"`
	Projectid string `json:"projectid,omitempty"`
	Id        string `json:"id,omitempty"`
	Email     string `json:"email,omitempty"`
	Domainid  string `json:"domainid,omitempty"`
	Project   string `json:"project,omitempty"`
}

type UpdateProjectInvitationParams struct {
	p map[string]interface{}
}

func (p *UpdateProjectInvitationParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["accept"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("accept", vv)
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["token"]; found {
		u.Set("token", v.(string))
	}
	return u
}

func (p *UpdateProjectInvitationParams) SetAccept(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accept"] = v
	return
}

func (p *UpdateProjectInvitationParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *UpdateProjectInvitationParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *UpdateProjectInvitationParams) SetToken(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["token"] = v
	return
}

// You should always use this function to get a new UpdateProjectInvitationParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewUpdateProjectInvitationParams(projectid string) *UpdateProjectInvitationParams {
	p := &UpdateProjectInvitationParams{}
	p.p = make(map[string]interface{})
	p.p["projectid"] = projectid
	return p
}

// Accepts or declines project invitation
func (s *ProjectService) UpdateProjectInvitation(p *UpdateProjectInvitationParams) (*UpdateProjectInvitationResponse, error) {
	resp, err := s.cs.newRequest("updateProjectInvitation", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateProjectInvitationResponse
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

type UpdateProjectInvitationResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Success     bool   `json:"success,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
}

type DeleteProjectInvitationParams struct {
	p map[string]interface{}
}

func (p *DeleteProjectInvitationParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteProjectInvitationParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteProjectInvitationParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewDeleteProjectInvitationParams(id string) *DeleteProjectInvitationParams {
	p := &DeleteProjectInvitationParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Accepts or declines project invitation
func (s *ProjectService) DeleteProjectInvitation(p *DeleteProjectInvitationParams) (*DeleteProjectInvitationResponse, error) {
	resp, err := s.cs.newRequest("deleteProjectInvitation", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteProjectInvitationResponse
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

type DeleteProjectInvitationResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Success     bool   `json:"success,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
}
