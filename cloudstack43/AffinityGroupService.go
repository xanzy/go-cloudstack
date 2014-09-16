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
	"strings"
)

type CreateAffinityGroupParams struct {
	p map[string]interface{}
}

func (p *CreateAffinityGroupParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["type"]; found {
		u.Set("type", v.(string))
	}
	return u
}

func (p *CreateAffinityGroupParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *CreateAffinityGroupParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
	return
}

func (p *CreateAffinityGroupParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *CreateAffinityGroupParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *CreateAffinityGroupParams) SetType(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["affinityGroupType"] = v
	return
}

// You should always use this function to get a new CreateAffinityGroupParams instance,
// as then you are sure you have configured all required params
func (s *AffinityGroupService) NewCreateAffinityGroupParams(name string, affinityGroupType string) *CreateAffinityGroupParams {
	p := &CreateAffinityGroupParams{}
	p.p = make(map[string]interface{})
	p.p["name"] = name
	p.p["affinityGroupType"] = affinityGroupType
	return p
}

// Creates an affinity/anti-affinity group
func (s *AffinityGroupService) CreateAffinityGroup(p *CreateAffinityGroupParams) (*CreateAffinityGroupResponse, error) {
	resp, err := s.cs.newRequest("createAffinityGroup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateAffinityGroupResponse
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

		var r CreateAffinityGroupResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type CreateAffinityGroupResponse struct {
	JobID             string   `json:"jobid,omitempty"`
	Domain            string   `json:"domain,omitempty"`
	Account           string   `json:"account,omitempty"`
	Description       string   `json:"description,omitempty"`
	Type              string   `json:"type,omitempty"`
	Domainid          string   `json:"domainid,omitempty"`
	Name              string   `json:"name,omitempty"`
	VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
	Id                string   `json:"id,omitempty"`
}

type DeleteAffinityGroupParams struct {
	p map[string]interface{}
}

func (p *DeleteAffinityGroupParams) toURLValues() url.Values {
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
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	return u
}

func (p *DeleteAffinityGroupParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *DeleteAffinityGroupParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *DeleteAffinityGroupParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *DeleteAffinityGroupParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

// You should always use this function to get a new DeleteAffinityGroupParams instance,
// as then you are sure you have configured all required params
func (s *AffinityGroupService) NewDeleteAffinityGroupParams() *DeleteAffinityGroupParams {
	p := &DeleteAffinityGroupParams{}
	p.p = make(map[string]interface{})
	return p
}

// Deletes affinity group
func (s *AffinityGroupService) DeleteAffinityGroup(p *DeleteAffinityGroupParams) (*DeleteAffinityGroupResponse, error) {
	resp, err := s.cs.newRequest("deleteAffinityGroup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteAffinityGroupResponse
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

		var r DeleteAffinityGroupResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type DeleteAffinityGroupResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Success     bool   `json:"success,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
}

type ListAffinityGroupsParams struct {
	p map[string]interface{}
}

func (p *ListAffinityGroupsParams) toURLValues() url.Values {
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
	if v, found := p.p["type"]; found {
		u.Set("type", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *ListAffinityGroupsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListAffinityGroupsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListAffinityGroupsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListAffinityGroupsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
	return
}

func (p *ListAffinityGroupsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListAffinityGroupsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
	return
}

func (p *ListAffinityGroupsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *ListAffinityGroupsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListAffinityGroupsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListAffinityGroupsParams) SetType(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["affinityGroupType"] = v
	return
}

func (p *ListAffinityGroupsParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
	return
}

// You should always use this function to get a new ListAffinityGroupsParams instance,
// as then you are sure you have configured all required params
func (s *AffinityGroupService) NewListAffinityGroupsParams() *ListAffinityGroupsParams {
	p := &ListAffinityGroupsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AffinityGroupService) GetAffinityGroupID(name string) (string, error) {
	p := &ListAffinityGroupsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	l, err := s.ListAffinityGroups(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, name, l)
	}
	return l.AffinityGroups[0].Id, nil
}

// Lists affinity groups
func (s *AffinityGroupService) ListAffinityGroups(p *ListAffinityGroupsParams) (*ListAffinityGroupsResponse, error) {
	resp, err := s.cs.newRequest("listAffinityGroups", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListAffinityGroupsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListAffinityGroupsResponse struct {
	Count          int              `json:"count"`
	AffinityGroups []*AffinityGroup `json:"affinitygroup"`
}

type AffinityGroup struct {
	Type              string   `json:"type,omitempty"`
	Description       string   `json:"description,omitempty"`
	Domainid          string   `json:"domainid,omitempty"`
	Id                string   `json:"id,omitempty"`
	Domain            string   `json:"domain,omitempty"`
	Account           string   `json:"account,omitempty"`
	Name              string   `json:"name,omitempty"`
	VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
}

type UpdateVMAffinityGroupParams struct {
	p map[string]interface{}
}

func (p *UpdateVMAffinityGroupParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["affinitygroupids"]; found {
		vv := strings.Join(v.([]string), ", ")
		u.Set("affinitygroupids", vv)
	}
	if v, found := p.p["affinitygroupnames"]; found {
		vv := strings.Join(v.([]string), ", ")
		u.Set("affinitygroupnames", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *UpdateVMAffinityGroupParams) SetAffinitygroupids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["affinitygroupids"] = v
	return
}

func (p *UpdateVMAffinityGroupParams) SetAffinitygroupnames(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["affinitygroupnames"] = v
	return
}

func (p *UpdateVMAffinityGroupParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new UpdateVMAffinityGroupParams instance,
// as then you are sure you have configured all required params
func (s *AffinityGroupService) NewUpdateVMAffinityGroupParams(id string) *UpdateVMAffinityGroupParams {
	p := &UpdateVMAffinityGroupParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates the affinity/anti-affinity group associations of a virtual machine. The VM has to be stopped and restarted for the new properties to take effect.
func (s *AffinityGroupService) UpdateVMAffinityGroup(p *UpdateVMAffinityGroupParams) (*UpdateVMAffinityGroupResponse, error) {
	resp, err := s.cs.newRequest("updateVMAffinityGroup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateVMAffinityGroupResponse
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

		var r UpdateVMAffinityGroupResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type UpdateVMAffinityGroupResponse struct {
	JobID         string `json:"jobid,omitempty"`
	Cpunumber     int    `json:"cpunumber,omitempty"`
	Domainid      string `json:"domainid,omitempty"`
	Created       string `json:"created,omitempty"`
	Securitygroup []struct {
		Ingressrule []struct {
			Endport           int    `json:"endport,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Account           string `json:"account,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
		} `json:"ingressrule,omitempty"`
		Description string `json:"description,omitempty"`
		Name        string `json:"name,omitempty"`
		Tags        []struct {
			Account      string `json:"account,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
			Project      string `json:"project,omitempty"`
			Key          string `json:"key,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Value        string `json:"value,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
		} `json:"tags,omitempty"`
		Project    string `json:"project,omitempty"`
		Account    string `json:"account,omitempty"`
		Projectid  string `json:"projectid,omitempty"`
		Domain     string `json:"domain,omitempty"`
		Id         string `json:"id,omitempty"`
		Domainid   string `json:"domainid,omitempty"`
		Egressrule []struct {
			Protocol          string `json:"protocol,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Account           string `json:"account,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
		} `json:"egressrule,omitempty"`
	} `json:"securitygroup,omitempty"`
	Haenable              bool   `json:"haenable,omitempty"`
	Networkkbswrite       int    `json:"networkkbswrite,omitempty"`
	Hypervisor            string `json:"hypervisor,omitempty"`
	Rootdevicetype        string `json:"rootdevicetype,omitempty"`
	Publicipid            string `json:"publicipid,omitempty"`
	Servicestate          string `json:"servicestate,omitempty"`
	Displayvm             bool   `json:"displayvm,omitempty"`
	Projectid             string `json:"projectid,omitempty"`
	Rootdeviceid          int    `json:"rootdeviceid,omitempty"`
	Serviceofferingid     string `json:"serviceofferingid,omitempty"`
	Guestosid             string `json:"guestosid,omitempty"`
	Instancename          string `json:"instancename,omitempty"`
	Diskiowrite           int    `json:"diskiowrite,omitempty"`
	Account               string `json:"account,omitempty"`
	Isdynamicallyscalable bool   `json:"isdynamicallyscalable,omitempty"`
	Hostid                string `json:"hostid,omitempty"`
	Isodisplaytext        string `json:"isodisplaytext,omitempty"`
	Domain                string `json:"domain,omitempty"`
	Displayname           string `json:"displayname,omitempty"`
	Nic                   []struct {
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Id           string   `json:"id,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Type         string   `json:"type,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
	} `json:"nic,omitempty"`
	Templatedisplaytext string `json:"templatedisplaytext,omitempty"`
	Groupid             string `json:"groupid,omitempty"`
	Cpuused             string `json:"cpuused,omitempty"`
	Diskioread          int    `json:"diskioread,omitempty"`
	Keypair             string `json:"keypair,omitempty"`
	Password            string `json:"password,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Isoname             string `json:"isoname,omitempty"`
	Templateid          string `json:"templateid,omitempty"`
	Publicip            string `json:"publicip,omitempty"`
	Group               string `json:"group,omitempty"`
	Zoneid              string `json:"zoneid,omitempty"`
	Project             string `json:"project,omitempty"`
	Networkkbsread      int    `json:"networkkbsread,omitempty"`
	Diskkbswrite        int    `json:"diskkbswrite,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
	Tags                []struct {
		Projectid    string `json:"projectid,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Account      string `json:"account,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Value        string `json:"value,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Project      string `json:"project,omitempty"`
		Key          string `json:"key,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
	} `json:"tags,omitempty"`
	Affinitygroup []struct {
		Id                string   `json:"id,omitempty"`
		Name              string   `json:"name,omitempty"`
		Type              string   `json:"type,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
		Account           string   `json:"account,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Description       string   `json:"description,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Isoid             string            `json:"isoid,omitempty"`
	Name              string            `json:"name,omitempty"`
	Templatename      string            `json:"templatename,omitempty"`
	Id                string            `json:"id,omitempty"`
	State             string            `json:"state,omitempty"`
	Details           map[string]string `json:"details,omitempty"`
	Forvirtualnetwork bool              `json:"forvirtualnetwork,omitempty"`
	Diskkbsread       int               `json:"diskkbsread,omitempty"`
	Cpuspeed          int               `json:"cpuspeed,omitempty"`
	Passwordenabled   bool              `json:"passwordenabled,omitempty"`
	Memory            int               `json:"memory,omitempty"`
	Hostname          string            `json:"hostname,omitempty"`
}

type ListAffinityGroupTypesParams struct {
	p map[string]interface{}
}

func (p *ListAffinityGroupTypesParams) toURLValues() url.Values {
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
	return u
}

func (p *ListAffinityGroupTypesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListAffinityGroupTypesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListAffinityGroupTypesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

// You should always use this function to get a new ListAffinityGroupTypesParams instance,
// as then you are sure you have configured all required params
func (s *AffinityGroupService) NewListAffinityGroupTypesParams() *ListAffinityGroupTypesParams {
	p := &ListAffinityGroupTypesParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists affinity group types available
func (s *AffinityGroupService) ListAffinityGroupTypes(p *ListAffinityGroupTypesParams) (*ListAffinityGroupTypesResponse, error) {
	resp, err := s.cs.newRequest("listAffinityGroupTypes", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListAffinityGroupTypesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListAffinityGroupTypesResponse struct {
	Count              int                  `json:"count"`
	AffinityGroupTypes []*AffinityGroupType `json:"affinitygrouptype"`
}

type AffinityGroupType struct {
	Type string `json:"type,omitempty"`
}
