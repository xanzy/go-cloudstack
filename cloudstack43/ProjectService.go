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
	JobID               string `json:"jobid,omitempty"`
	Vmstopped           int    `json:"vmstopped,omitempty"`
	Primarystoragelimit string `json:"primarystoragelimit,omitempty"`
	Networktotal        int    `json:"networktotal,omitempty"`
	Templatelimit       string `json:"templatelimit,omitempty"`
	Tags                []struct {
		Account      string `json:"account,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Value        string `json:"value,omitempty"`
		Key          string `json:"key,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Project      string `json:"project,omitempty"`
	} `json:"tags,omitempty"`
	Domainid                  string `json:"domainid,omitempty"`
	Cpulimit                  string `json:"cpulimit,omitempty"`
	Templateavailable         string `json:"templateavailable,omitempty"`
	Volumetotal               int    `json:"volumetotal,omitempty"`
	State                     string `json:"state,omitempty"`
	Iplimit                   string `json:"iplimit,omitempty"`
	Vmavailable               string `json:"vmavailable,omitempty"`
	Displaytext               string `json:"displaytext,omitempty"`
	Snapshotlimit             string `json:"snapshotlimit,omitempty"`
	Iptotal                   int    `json:"iptotal,omitempty"`
	Vmrunning                 int    `json:"vmrunning,omitempty"`
	Cpuavailable              string `json:"cpuavailable,omitempty"`
	Snapshotavailable         string `json:"snapshotavailable,omitempty"`
	Secondarystoragetotal     int    `json:"secondarystoragetotal,omitempty"`
	Domain                    string `json:"domain,omitempty"`
	Secondarystorageavailable string `json:"secondarystorageavailable,omitempty"`
	Volumelimit               string `json:"volumelimit,omitempty"`
	Networklimit              string `json:"networklimit,omitempty"`
	Ipavailable               string `json:"ipavailable,omitempty"`
	Networkavailable          string `json:"networkavailable,omitempty"`
	Snapshottotal             int    `json:"snapshottotal,omitempty"`
	Vmtotal                   int    `json:"vmtotal,omitempty"`
	Memorytotal               int    `json:"memorytotal,omitempty"`
	Name                      string `json:"name,omitempty"`
	Vpclimit                  string `json:"vpclimit,omitempty"`
	Id                        string `json:"id,omitempty"`
	Primarystoragetotal       int    `json:"primarystoragetotal,omitempty"`
	Volumeavailable           string `json:"volumeavailable,omitempty"`
	Secondarystoragelimit     string `json:"secondarystoragelimit,omitempty"`
	Memorylimit               string `json:"memorylimit,omitempty"`
	Vpctotal                  int    `json:"vpctotal,omitempty"`
	Vmlimit                   string `json:"vmlimit,omitempty"`
	Cputotal                  int    `json:"cputotal,omitempty"`
	Memoryavailable           string `json:"memoryavailable,omitempty"`
	Primarystorageavailable   string `json:"primarystorageavailable,omitempty"`
	Account                   string `json:"account,omitempty"`
	Templatetotal             int    `json:"templatetotal,omitempty"`
	Vpcavailable              string `json:"vpcavailable,omitempty"`
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
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
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
	JobID string `json:"jobid,omitempty"`
	Tags  []struct {
		Value        string `json:"value,omitempty"`
		Key          string `json:"key,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Project      string `json:"project,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Account      string `json:"account,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
	} `json:"tags,omitempty"`
	Cpuavailable              string `json:"cpuavailable,omitempty"`
	Memoryavailable           string `json:"memoryavailable,omitempty"`
	Secondarystoragelimit     string `json:"secondarystoragelimit,omitempty"`
	Name                      string `json:"name,omitempty"`
	Vpcavailable              string `json:"vpcavailable,omitempty"`
	Primarystoragelimit       string `json:"primarystoragelimit,omitempty"`
	Domainid                  string `json:"domainid,omitempty"`
	Domain                    string `json:"domain,omitempty"`
	Templateavailable         string `json:"templateavailable,omitempty"`
	Memorylimit               string `json:"memorylimit,omitempty"`
	Snapshotlimit             string `json:"snapshotlimit,omitempty"`
	Memorytotal               int    `json:"memorytotal,omitempty"`
	State                     string `json:"state,omitempty"`
	Volumelimit               string `json:"volumelimit,omitempty"`
	Networktotal              int    `json:"networktotal,omitempty"`
	Networkavailable          string `json:"networkavailable,omitempty"`
	Cputotal                  int    `json:"cputotal,omitempty"`
	Primarystoragetotal       int    `json:"primarystoragetotal,omitempty"`
	Templatetotal             int    `json:"templatetotal,omitempty"`
	Iptotal                   int    `json:"iptotal,omitempty"`
	Cpulimit                  string `json:"cpulimit,omitempty"`
	Id                        string `json:"id,omitempty"`
	Vmstopped                 int    `json:"vmstopped,omitempty"`
	Vmlimit                   string `json:"vmlimit,omitempty"`
	Volumeavailable           string `json:"volumeavailable,omitempty"`
	Networklimit              string `json:"networklimit,omitempty"`
	Iplimit                   string `json:"iplimit,omitempty"`
	Vpctotal                  int    `json:"vpctotal,omitempty"`
	Volumetotal               int    `json:"volumetotal,omitempty"`
	Vpclimit                  string `json:"vpclimit,omitempty"`
	Displaytext               string `json:"displaytext,omitempty"`
	Vmtotal                   int    `json:"vmtotal,omitempty"`
	Snapshottotal             int    `json:"snapshottotal,omitempty"`
	Snapshotavailable         string `json:"snapshotavailable,omitempty"`
	Ipavailable               string `json:"ipavailable,omitempty"`
	Account                   string `json:"account,omitempty"`
	Templatelimit             string `json:"templatelimit,omitempty"`
	Primarystorageavailable   string `json:"primarystorageavailable,omitempty"`
	Secondarystoragetotal     int    `json:"secondarystoragetotal,omitempty"`
	Vmrunning                 int    `json:"vmrunning,omitempty"`
	Vmavailable               string `json:"vmavailable,omitempty"`
	Secondarystorageavailable string `json:"secondarystorageavailable,omitempty"`
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
	JobID                     string `json:"jobid,omitempty"`
	Cpulimit                  string `json:"cpulimit,omitempty"`
	Vpcavailable              string `json:"vpcavailable,omitempty"`
	Primarystoragelimit       string `json:"primarystoragelimit,omitempty"`
	Account                   string `json:"account,omitempty"`
	Vpclimit                  string `json:"vpclimit,omitempty"`
	Primarystorageavailable   string `json:"primarystorageavailable,omitempty"`
	Domain                    string `json:"domain,omitempty"`
	Volumelimit               string `json:"volumelimit,omitempty"`
	Displaytext               string `json:"displaytext,omitempty"`
	Templatelimit             string `json:"templatelimit,omitempty"`
	Secondarystorageavailable string `json:"secondarystorageavailable,omitempty"`
	Vpctotal                  int    `json:"vpctotal,omitempty"`
	State                     string `json:"state,omitempty"`
	Tags                      []struct {
		Customer     string `json:"customer,omitempty"`
		Account      string `json:"account,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Value        string `json:"value,omitempty"`
		Key          string `json:"key,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Project      string `json:"project,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Domain       string `json:"domain,omitempty"`
	} `json:"tags,omitempty"`
	Primarystoragetotal   int    `json:"primarystoragetotal,omitempty"`
	Templateavailable     string `json:"templateavailable,omitempty"`
	Snapshottotal         int    `json:"snapshottotal,omitempty"`
	Domainid              string `json:"domainid,omitempty"`
	Vmavailable           string `json:"vmavailable,omitempty"`
	Secondarystoragelimit string `json:"secondarystoragelimit,omitempty"`
	Iptotal               int    `json:"iptotal,omitempty"`
	Vmtotal               int    `json:"vmtotal,omitempty"`
	Id                    string `json:"id,omitempty"`
	Memorytotal           int    `json:"memorytotal,omitempty"`
	Snapshotlimit         string `json:"snapshotlimit,omitempty"`
	Secondarystoragetotal int    `json:"secondarystoragetotal,omitempty"`
	Networktotal          int    `json:"networktotal,omitempty"`
	Cpuavailable          string `json:"cpuavailable,omitempty"`
	Vmrunning             int    `json:"vmrunning,omitempty"`
	Vmlimit               string `json:"vmlimit,omitempty"`
	Networkavailable      string `json:"networkavailable,omitempty"`
	Vmstopped             int    `json:"vmstopped,omitempty"`
	Memorylimit           string `json:"memorylimit,omitempty"`
	Memoryavailable       string `json:"memoryavailable,omitempty"`
	Networklimit          string `json:"networklimit,omitempty"`
	Iplimit               string `json:"iplimit,omitempty"`
	Volumetotal           int    `json:"volumetotal,omitempty"`
	Ipavailable           string `json:"ipavailable,omitempty"`
	Cputotal              int    `json:"cputotal,omitempty"`
	Snapshotavailable     string `json:"snapshotavailable,omitempty"`
	Name                  string `json:"name,omitempty"`
	Volumeavailable       string `json:"volumeavailable,omitempty"`
	Templatetotal         int    `json:"templatetotal,omitempty"`
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
	Vmavailable               string `json:"vmavailable,omitempty"`
	Vpctotal                  int    `json:"vpctotal,omitempty"`
	Memoryavailable           string `json:"memoryavailable,omitempty"`
	Cpulimit                  string `json:"cpulimit,omitempty"`
	Vmlimit                   string `json:"vmlimit,omitempty"`
	Volumeavailable           string `json:"volumeavailable,omitempty"`
	Volumelimit               string `json:"volumelimit,omitempty"`
	Snapshotavailable         string `json:"snapshotavailable,omitempty"`
	Snapshotlimit             string `json:"snapshotlimit,omitempty"`
	Templatelimit             string `json:"templatelimit,omitempty"`
	Snapshottotal             int    `json:"snapshottotal,omitempty"`
	Iptotal                   int    `json:"iptotal,omitempty"`
	Templateavailable         string `json:"templateavailable,omitempty"`
	Vpclimit                  string `json:"vpclimit,omitempty"`
	Domain                    string `json:"domain,omitempty"`
	Vmrunning                 int    `json:"vmrunning,omitempty"`
	Secondarystorageavailable string `json:"secondarystorageavailable,omitempty"`
	Networkavailable          string `json:"networkavailable,omitempty"`
	Templatetotal             int    `json:"templatetotal,omitempty"`
	Iplimit                   string `json:"iplimit,omitempty"`
	Vmstopped                 int    `json:"vmstopped,omitempty"`
	Volumetotal               int    `json:"volumetotal,omitempty"`
	Networktotal              int    `json:"networktotal,omitempty"`
	State                     string `json:"state,omitempty"`
	Displaytext               string `json:"displaytext,omitempty"`
	Memorylimit               string `json:"memorylimit,omitempty"`
	Secondarystoragetotal     int    `json:"secondarystoragetotal,omitempty"`
	Id                        string `json:"id,omitempty"`
	Vpcavailable              string `json:"vpcavailable,omitempty"`
	Cpuavailable              string `json:"cpuavailable,omitempty"`
	Primarystoragelimit       string `json:"primarystoragelimit,omitempty"`
	Secondarystoragelimit     string `json:"secondarystoragelimit,omitempty"`
	Name                      string `json:"name,omitempty"`
	Primarystoragetotal       int    `json:"primarystoragetotal,omitempty"`
	Vmtotal                   int    `json:"vmtotal,omitempty"`
	Tags                      []struct {
		Resourceid   string `json:"resourceid,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Value        string `json:"value,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Project      string `json:"project,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Key          string `json:"key,omitempty"`
		Account      string `json:"account,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Customer     string `json:"customer,omitempty"`
	} `json:"tags,omitempty"`
	Primarystorageavailable string `json:"primarystorageavailable,omitempty"`
	Domainid                string `json:"domainid,omitempty"`
	Ipavailable             string `json:"ipavailable,omitempty"`
	Account                 string `json:"account,omitempty"`
	Networklimit            string `json:"networklimit,omitempty"`
	Memorytotal             int    `json:"memorytotal,omitempty"`
	Cputotal                int    `json:"cputotal,omitempty"`
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
	Vpclimit                string `json:"vpclimit,omitempty"`
	Displaytext             string `json:"displaytext,omitempty"`
	Snapshotlimit           string `json:"snapshotlimit,omitempty"`
	Vmstopped               int    `json:"vmstopped,omitempty"`
	Templateavailable       string `json:"templateavailable,omitempty"`
	State                   string `json:"state,omitempty"`
	Snapshottotal           int    `json:"snapshottotal,omitempty"`
	Account                 string `json:"account,omitempty"`
	Domainid                string `json:"domainid,omitempty"`
	Vmavailable             string `json:"vmavailable,omitempty"`
	Domain                  string `json:"domain,omitempty"`
	Iptotal                 int    `json:"iptotal,omitempty"`
	Primarystoragetotal     int    `json:"primarystoragetotal,omitempty"`
	Secondarystoragelimit   string `json:"secondarystoragelimit,omitempty"`
	Secondarystoragetotal   int    `json:"secondarystoragetotal,omitempty"`
	Vmtotal                 int    `json:"vmtotal,omitempty"`
	Primarystorageavailable string `json:"primarystorageavailable,omitempty"`
	Ipavailable             string `json:"ipavailable,omitempty"`
	Templatetotal           int    `json:"templatetotal,omitempty"`
	Name                    string `json:"name,omitempty"`
	Cpulimit                string `json:"cpulimit,omitempty"`
	Tags                    []struct {
		Value        string `json:"value,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Account      string `json:"account,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Project      string `json:"project,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Key          string `json:"key,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Domain       string `json:"domain,omitempty"`
	} `json:"tags,omitempty"`
	Vpctotal                  int    `json:"vpctotal,omitempty"`
	Id                        string `json:"id,omitempty"`
	Volumelimit               string `json:"volumelimit,omitempty"`
	Networktotal              int    `json:"networktotal,omitempty"`
	Secondarystorageavailable string `json:"secondarystorageavailable,omitempty"`
	Volumeavailable           string `json:"volumeavailable,omitempty"`
	Primarystoragelimit       string `json:"primarystoragelimit,omitempty"`
	Templatelimit             string `json:"templatelimit,omitempty"`
	Vmrunning                 int    `json:"vmrunning,omitempty"`
	Vpcavailable              string `json:"vpcavailable,omitempty"`
	Volumetotal               int    `json:"volumetotal,omitempty"`
	Vmlimit                   string `json:"vmlimit,omitempty"`
	Networkavailable          string `json:"networkavailable,omitempty"`
	Cputotal                  int    `json:"cputotal,omitempty"`
	Cpuavailable              string `json:"cpuavailable,omitempty"`
	Snapshotavailable         string `json:"snapshotavailable,omitempty"`
	Memorylimit               string `json:"memorylimit,omitempty"`
	Networklimit              string `json:"networklimit,omitempty"`
	Memorytotal               int    `json:"memorytotal,omitempty"`
	Iplimit                   string `json:"iplimit,omitempty"`
	Memoryavailable           string `json:"memoryavailable,omitempty"`
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
	State     string `json:"state,omitempty"`
	Project   string `json:"project,omitempty"`
	Domainid  string `json:"domainid,omitempty"`
	Projectid string `json:"projectid,omitempty"`
	Domain    string `json:"domain,omitempty"`
	Id        string `json:"id,omitempty"`
	Account   string `json:"account,omitempty"`
	Email     string `json:"email,omitempty"`
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
