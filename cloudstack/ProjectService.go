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

		var r CreateProjectResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type CreateProjectResponse struct {
	JobID                   string `json:"jobid,omitempty"`
	Account                 string `json:"account,omitempty"`
	Name                    string `json:"name,omitempty"`
	State                   string `json:"state,omitempty"`
	Secondarystoragetotal   int    `json:"secondarystoragetotal,omitempty"`
	Snapshotlimit           string `json:"snapshotlimit,omitempty"`
	Vmtotal                 int    `json:"vmtotal,omitempty"`
	Vmavailable             string `json:"vmavailable,omitempty"`
	Cpulimit                string `json:"cpulimit,omitempty"`
	Memorylimit             string `json:"memorylimit,omitempty"`
	Primarystorageavailable string `json:"primarystorageavailable,omitempty"`
	Vmrunning               int    `json:"vmrunning,omitempty"`
	Networklimit            string `json:"networklimit,omitempty"`
	Templatetotal           int    `json:"templatetotal,omitempty"`
	Networktotal            int    `json:"networktotal,omitempty"`
	Primarystoragetotal     int    `json:"primarystoragetotal,omitempty"`
	Tags                    []struct {
		Project      string `json:"project,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Value        string `json:"value,omitempty"`
		Key          string `json:"key,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Account      string `json:"account,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
	} `json:"tags,omitempty"`
	Displaytext               string `json:"displaytext,omitempty"`
	Vpclimit                  string `json:"vpclimit,omitempty"`
	Volumetotal               int    `json:"volumetotal,omitempty"`
	Volumelimit               string `json:"volumelimit,omitempty"`
	Snapshottotal             int    `json:"snapshottotal,omitempty"`
	Vpcavailable              string `json:"vpcavailable,omitempty"`
	Vmlimit                   string `json:"vmlimit,omitempty"`
	Snapshotavailable         string `json:"snapshotavailable,omitempty"`
	Templateavailable         string `json:"templateavailable,omitempty"`
	Volumeavailable           string `json:"volumeavailable,omitempty"`
	Vmstopped                 int    `json:"vmstopped,omitempty"`
	Vpctotal                  int    `json:"vpctotal,omitempty"`
	Templatelimit             string `json:"templatelimit,omitempty"`
	Memoryavailable           string `json:"memoryavailable,omitempty"`
	Ipavailable               string `json:"ipavailable,omitempty"`
	Id                        string `json:"id,omitempty"`
	Cpuavailable              string `json:"cpuavailable,omitempty"`
	Primarystoragelimit       string `json:"primarystoragelimit,omitempty"`
	Iplimit                   string `json:"iplimit,omitempty"`
	Domainid                  string `json:"domainid,omitempty"`
	Memorytotal               int    `json:"memorytotal,omitempty"`
	Cputotal                  int    `json:"cputotal,omitempty"`
	Secondarystoragelimit     string `json:"secondarystoragelimit,omitempty"`
	Networkavailable          string `json:"networkavailable,omitempty"`
	Secondarystorageavailable string `json:"secondarystorageavailable,omitempty"`
	Domain                    string `json:"domain,omitempty"`
	Iptotal                   int    `json:"iptotal,omitempty"`
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

		var r DeleteProjectResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
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

		var r UpdateProjectResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type UpdateProjectResponse struct {
	JobID                     string `json:"jobid,omitempty"`
	Volumetotal               int    `json:"volumetotal,omitempty"`
	Templatetotal             int    `json:"templatetotal,omitempty"`
	Vmlimit                   string `json:"vmlimit,omitempty"`
	Memoryavailable           string `json:"memoryavailable,omitempty"`
	Vmrunning                 int    `json:"vmrunning,omitempty"`
	Networklimit              string `json:"networklimit,omitempty"`
	Iplimit                   string `json:"iplimit,omitempty"`
	Secondarystoragelimit     string `json:"secondarystoragelimit,omitempty"`
	Vmstopped                 int    `json:"vmstopped,omitempty"`
	Ipavailable               string `json:"ipavailable,omitempty"`
	Domainid                  string `json:"domainid,omitempty"`
	Vpcavailable              string `json:"vpcavailable,omitempty"`
	Volumelimit               string `json:"volumelimit,omitempty"`
	Name                      string `json:"name,omitempty"`
	Volumeavailable           string `json:"volumeavailable,omitempty"`
	Secondarystorageavailable string `json:"secondarystorageavailable,omitempty"`
	Vmtotal                   int    `json:"vmtotal,omitempty"`
	Cpulimit                  string `json:"cpulimit,omitempty"`
	Secondarystoragetotal     int    `json:"secondarystoragetotal,omitempty"`
	Iptotal                   int    `json:"iptotal,omitempty"`
	Tags                      []struct {
		Domainid     string `json:"domainid,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Key          string `json:"key,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Account      string `json:"account,omitempty"`
		Value        string `json:"value,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Project      string `json:"project,omitempty"`
	} `json:"tags,omitempty"`
	Memorylimit             string `json:"memorylimit,omitempty"`
	Primarystoragelimit     string `json:"primarystoragelimit,omitempty"`
	Templateavailable       string `json:"templateavailable,omitempty"`
	Templatelimit           string `json:"templatelimit,omitempty"`
	Vmavailable             string `json:"vmavailable,omitempty"`
	Networkavailable        string `json:"networkavailable,omitempty"`
	Primarystorageavailable string `json:"primarystorageavailable,omitempty"`
	State                   string `json:"state,omitempty"`
	Memorytotal             int    `json:"memorytotal,omitempty"`
	Snapshotlimit           string `json:"snapshotlimit,omitempty"`
	Account                 string `json:"account,omitempty"`
	Snapshottotal           int    `json:"snapshottotal,omitempty"`
	Vpctotal                int    `json:"vpctotal,omitempty"`
	Vpclimit                string `json:"vpclimit,omitempty"`
	Displaytext             string `json:"displaytext,omitempty"`
	Primarystoragetotal     int    `json:"primarystoragetotal,omitempty"`
	Snapshotavailable       string `json:"snapshotavailable,omitempty"`
	Id                      string `json:"id,omitempty"`
	Networktotal            int    `json:"networktotal,omitempty"`
	Cputotal                int    `json:"cputotal,omitempty"`
	Cpuavailable            string `json:"cpuavailable,omitempty"`
	Domain                  string `json:"domain,omitempty"`
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

		var r ActivateProjectResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type ActivateProjectResponse struct {
	JobID                 string `json:"jobid,omitempty"`
	Secondarystoragelimit string `json:"secondarystoragelimit,omitempty"`
	Snapshottotal         int    `json:"snapshottotal,omitempty"`
	Domainid              string `json:"domainid,omitempty"`
	Templatetotal         int    `json:"templatetotal,omitempty"`
	Networkavailable      string `json:"networkavailable,omitempty"`
	Ipavailable           string `json:"ipavailable,omitempty"`
	Networklimit          string `json:"networklimit,omitempty"`
	Tags                  []struct {
		Domainid     string `json:"domainid,omitempty"`
		Value        string `json:"value,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Account      string `json:"account,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Project      string `json:"project,omitempty"`
		Key          string `json:"key,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
	} `json:"tags,omitempty"`
	Primarystorageavailable   string `json:"primarystorageavailable,omitempty"`
	Vmlimit                   string `json:"vmlimit,omitempty"`
	Primarystoragelimit       string `json:"primarystoragelimit,omitempty"`
	Volumeavailable           string `json:"volumeavailable,omitempty"`
	Iptotal                   int    `json:"iptotal,omitempty"`
	Iplimit                   string `json:"iplimit,omitempty"`
	Account                   string `json:"account,omitempty"`
	Volumetotal               int    `json:"volumetotal,omitempty"`
	Cpuavailable              string `json:"cpuavailable,omitempty"`
	Memorytotal               int    `json:"memorytotal,omitempty"`
	Secondarystorageavailable string `json:"secondarystorageavailable,omitempty"`
	Displaytext               string `json:"displaytext,omitempty"`
	Vmavailable               string `json:"vmavailable,omitempty"`
	Name                      string `json:"name,omitempty"`
	Memoryavailable           string `json:"memoryavailable,omitempty"`
	State                     string `json:"state,omitempty"`
	Snapshotlimit             string `json:"snapshotlimit,omitempty"`
	Templateavailable         string `json:"templateavailable,omitempty"`
	Secondarystoragetotal     int    `json:"secondarystoragetotal,omitempty"`
	Domain                    string `json:"domain,omitempty"`
	Vmrunning                 int    `json:"vmrunning,omitempty"`
	Cpulimit                  string `json:"cpulimit,omitempty"`
	Vmstopped                 int    `json:"vmstopped,omitempty"`
	Templatelimit             string `json:"templatelimit,omitempty"`
	Vpctotal                  int    `json:"vpctotal,omitempty"`
	Id                        string `json:"id,omitempty"`
	Memorylimit               string `json:"memorylimit,omitempty"`
	Vpcavailable              string `json:"vpcavailable,omitempty"`
	Primarystoragetotal       int    `json:"primarystoragetotal,omitempty"`
	Vmtotal                   int    `json:"vmtotal,omitempty"`
	Networktotal              int    `json:"networktotal,omitempty"`
	Cputotal                  int    `json:"cputotal,omitempty"`
	Snapshotavailable         string `json:"snapshotavailable,omitempty"`
	Volumelimit               string `json:"volumelimit,omitempty"`
	Vpclimit                  string `json:"vpclimit,omitempty"`
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

		var r SuspendProjectResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type SuspendProjectResponse struct {
	JobID                     string `json:"jobid,omitempty"`
	Iplimit                   string `json:"iplimit,omitempty"`
	Secondarystoragetotal     int    `json:"secondarystoragetotal,omitempty"`
	Vpclimit                  string `json:"vpclimit,omitempty"`
	Templateavailable         string `json:"templateavailable,omitempty"`
	Networktotal              int    `json:"networktotal,omitempty"`
	Volumetotal               int    `json:"volumetotal,omitempty"`
	State                     string `json:"state,omitempty"`
	Memorylimit               string `json:"memorylimit,omitempty"`
	Vmrunning                 int    `json:"vmrunning,omitempty"`
	Snapshotlimit             string `json:"snapshotlimit,omitempty"`
	Vmtotal                   int    `json:"vmtotal,omitempty"`
	Primarystoragelimit       string `json:"primarystoragelimit,omitempty"`
	Primarystoragetotal       int    `json:"primarystoragetotal,omitempty"`
	Vmlimit                   string `json:"vmlimit,omitempty"`
	Volumelimit               string `json:"volumelimit,omitempty"`
	Secondarystorageavailable string `json:"secondarystorageavailable,omitempty"`
	Domain                    string `json:"domain,omitempty"`
	Templatetotal             int    `json:"templatetotal,omitempty"`
	Networkavailable          string `json:"networkavailable,omitempty"`
	Domainid                  string `json:"domainid,omitempty"`
	Iptotal                   int    `json:"iptotal,omitempty"`
	Vpctotal                  int    `json:"vpctotal,omitempty"`
	Cpulimit                  string `json:"cpulimit,omitempty"`
	Id                        string `json:"id,omitempty"`
	Memoryavailable           string `json:"memoryavailable,omitempty"`
	Primarystorageavailable   string `json:"primarystorageavailable,omitempty"`
	Networklimit              string `json:"networklimit,omitempty"`
	Vmavailable               string `json:"vmavailable,omitempty"`
	Volumeavailable           string `json:"volumeavailable,omitempty"`
	Snapshotavailable         string `json:"snapshotavailable,omitempty"`
	Secondarystoragelimit     string `json:"secondarystoragelimit,omitempty"`
	Cpuavailable              string `json:"cpuavailable,omitempty"`
	Memorytotal               int    `json:"memorytotal,omitempty"`
	Ipavailable               string `json:"ipavailable,omitempty"`
	Vmstopped                 int    `json:"vmstopped,omitempty"`
	Templatelimit             string `json:"templatelimit,omitempty"`
	Vpcavailable              string `json:"vpcavailable,omitempty"`
	Displaytext               string `json:"displaytext,omitempty"`
	Tags                      []struct {
		Customer     string `json:"customer,omitempty"`
		Project      string `json:"project,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Value        string `json:"value,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Key          string `json:"key,omitempty"`
		Account      string `json:"account,omitempty"`
	} `json:"tags,omitempty"`
	Account       string `json:"account,omitempty"`
	Name          string `json:"name,omitempty"`
	Snapshottotal int    `json:"snapshottotal,omitempty"`
	Cputotal      int    `json:"cputotal,omitempty"`
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
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, name, l)
	}
	return l.Projects[0].Id, nil
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
	Vmavailable               string `json:"vmavailable,omitempty"`
	Cpulimit                  string `json:"cpulimit,omitempty"`
	Snapshotavailable         string `json:"snapshotavailable,omitempty"`
	Ipavailable               string `json:"ipavailable,omitempty"`
	Secondarystoragetotal     int    `json:"secondarystoragetotal,omitempty"`
	Account                   string `json:"account,omitempty"`
	Volumeavailable           string `json:"volumeavailable,omitempty"`
	Volumetotal               int    `json:"volumetotal,omitempty"`
	Vmtotal                   int    `json:"vmtotal,omitempty"`
	Vmstopped                 int    `json:"vmstopped,omitempty"`
	Templateavailable         string `json:"templateavailable,omitempty"`
	Snapshotlimit             string `json:"snapshotlimit,omitempty"`
	Primarystoragelimit       string `json:"primarystoragelimit,omitempty"`
	Iplimit                   string `json:"iplimit,omitempty"`
	Memoryavailable           string `json:"memoryavailable,omitempty"`
	Domain                    string `json:"domain,omitempty"`
	Primarystorageavailable   string `json:"primarystorageavailable,omitempty"`
	Id                        string `json:"id,omitempty"`
	Memorylimit               string `json:"memorylimit,omitempty"`
	Vmrunning                 int    `json:"vmrunning,omitempty"`
	State                     string `json:"state,omitempty"`
	Vpclimit                  string `json:"vpclimit,omitempty"`
	Memorytotal               int    `json:"memorytotal,omitempty"`
	Templatelimit             string `json:"templatelimit,omitempty"`
	Networkavailable          string `json:"networkavailable,omitempty"`
	Networktotal              int    `json:"networktotal,omitempty"`
	Vpctotal                  int    `json:"vpctotal,omitempty"`
	Name                      string `json:"name,omitempty"`
	Domainid                  string `json:"domainid,omitempty"`
	Secondarystoragelimit     string `json:"secondarystoragelimit,omitempty"`
	Networklimit              string `json:"networklimit,omitempty"`
	Secondarystorageavailable string `json:"secondarystorageavailable,omitempty"`
	Iptotal                   int    `json:"iptotal,omitempty"`
	Vpcavailable              string `json:"vpcavailable,omitempty"`
	Snapshottotal             int    `json:"snapshottotal,omitempty"`
	Displaytext               string `json:"displaytext,omitempty"`
	Primarystoragetotal       int    `json:"primarystoragetotal,omitempty"`
	Tags                      []struct {
		Domain       string `json:"domain,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Value        string `json:"value,omitempty"`
		Key          string `json:"key,omitempty"`
		Account      string `json:"account,omitempty"`
		Project      string `json:"project,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
	} `json:"tags,omitempty"`
	Cpuavailable  string `json:"cpuavailable,omitempty"`
	Volumelimit   string `json:"volumelimit,omitempty"`
	Cputotal      int    `json:"cputotal,omitempty"`
	Templatetotal int    `json:"templatetotal,omitempty"`
	Vmlimit       string `json:"vmlimit,omitempty"`
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
func (s *ProjectService) GetProjectInvitationID(keyword string) (string, error) {
	p := &ListProjectInvitationsParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword

	l, err := s.ListProjectInvitations(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, keyword, l)
	}
	return l.ProjectInvitations[0].Id, nil
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
	Domain    string `json:"domain,omitempty"`
	State     string `json:"state,omitempty"`
	Projectid string `json:"projectid,omitempty"`
	Id        string `json:"id,omitempty"`
	Email     string `json:"email,omitempty"`
	Project   string `json:"project,omitempty"`
	Domainid  string `json:"domainid,omitempty"`
	Account   string `json:"account,omitempty"`
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

		var r UpdateProjectInvitationResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
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

		var r DeleteProjectInvitationResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type DeleteProjectInvitationResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}
