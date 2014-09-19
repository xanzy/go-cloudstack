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

type CreateAffinityGroupResponse struct {
	JobID             string   `json:"jobid,omitempty"`
	Name              string   `json:"name,omitempty"`
	Type              string   `json:"type,omitempty"`
	Account           string   `json:"account,omitempty"`
	Description       string   `json:"description,omitempty"`
	Id                string   `json:"id,omitempty"`
	Domain            string   `json:"domain,omitempty"`
	VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
	Domainid          string   `json:"domainid,omitempty"`
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

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
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

	if l.Count == 0 {
		return "", fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.AffinityGroups[0].Id, nil
	}

	if l.Count > 1 {
		for _, v := range l.AffinityGroups {
			if v.Name == name {
				return v.Id, nil
			}
		}
	}
	return "", fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AffinityGroupService) GetAffinityGroupByName(name string) (*AffinityGroup, int, error) {
	id, err := s.GetAffinityGroupID(name)
	if err != nil {
		return nil, -1, err
	}

	r, count, err := s.GetAffinityGroupByID(id)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AffinityGroupService) GetAffinityGroupByID(id string) (*AffinityGroup, int, error) {
	p := &ListAffinityGroupsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	l, err := s.ListAffinityGroups(p)
	if err != nil {
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.AffinityGroups[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for AffinityGroup UUID: %s!", id)
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
	Account           string   `json:"account,omitempty"`
	VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
	Domain            string   `json:"domain,omitempty"`
	Type              string   `json:"type,omitempty"`
	Name              string   `json:"name,omitempty"`
	Domainid          string   `json:"domainid,omitempty"`
	Id                string   `json:"id,omitempty"`
	Description       string   `json:"description,omitempty"`
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

type UpdateVMAffinityGroupResponse struct {
	JobID                 string `json:"jobid,omitempty"`
	Publicipid            string `json:"publicipid,omitempty"`
	Passwordenabled       bool   `json:"passwordenabled,omitempty"`
	Isdynamicallyscalable bool   `json:"isdynamicallyscalable,omitempty"`
	Zonename              string `json:"zonename,omitempty"`
	Diskkbsread           int    `json:"diskkbsread,omitempty"`
	Group                 string `json:"group,omitempty"`
	Forvirtualnetwork     bool   `json:"forvirtualnetwork,omitempty"`
	Templateid            string `json:"templateid,omitempty"`
	Isoname               string `json:"isoname,omitempty"`
	Networkkbsread        int    `json:"networkkbsread,omitempty"`
	Affinitygroup         []struct {
		Name              string   `json:"name,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
		Type              string   `json:"type,omitempty"`
		Account           string   `json:"account,omitempty"`
		Id                string   `json:"id,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Description       string   `json:"description,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Keypair      string `json:"keypair,omitempty"`
	Groupid      string `json:"groupid,omitempty"`
	Projectid    string `json:"projectid,omitempty"`
	Diskkbswrite int    `json:"diskkbswrite,omitempty"`
	Password     string `json:"password,omitempty"`
	Account      string `json:"account,omitempty"`
	Tags         []struct {
		Customer     string `json:"customer,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Value        string `json:"value,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Key          string `json:"key,omitempty"`
		Account      string `json:"account,omitempty"`
		Project      string `json:"project,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
	} `json:"tags,omitempty"`
	Haenable            bool   `json:"haenable,omitempty"`
	Hypervisor          string `json:"hypervisor,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Publicip            string `json:"publicip,omitempty"`
	Cpuspeed            int    `json:"cpuspeed,omitempty"`
	Displayvm           bool   `json:"displayvm,omitempty"`
	Created             string `json:"created,omitempty"`
	Id                  string `json:"id,omitempty"`
	Cpuused             string `json:"cpuused,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
	Project             string `json:"project,omitempty"`
	State               string `json:"state,omitempty"`
	Diskioread          int    `json:"diskioread,omitempty"`
	Guestosid           string `json:"guestosid,omitempty"`
	Cpunumber           int    `json:"cpunumber,omitempty"`
	Diskiowrite         int    `json:"diskiowrite,omitempty"`
	Rootdevicetype      string `json:"rootdevicetype,omitempty"`
	Isoid               string `json:"isoid,omitempty"`
	Hostid              string `json:"hostid,omitempty"`
	Instancename        string `json:"instancename,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Zoneid              string `json:"zoneid,omitempty"`
	Rootdeviceid        int    `json:"rootdeviceid,omitempty"`
	Displayname         string `json:"displayname,omitempty"`
	Securitygroup       []struct {
		Projectid string `json:"projectid,omitempty"`
		Tags      []struct {
			Account      string `json:"account,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
			Value        string `json:"value,omitempty"`
			Key          string `json:"key,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Project      string `json:"project,omitempty"`
		} `json:"tags,omitempty"`
		Domainid    string `json:"domainid,omitempty"`
		Project     string `json:"project,omitempty"`
		Ingressrule []struct {
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Account           string `json:"account,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
		} `json:"ingressrule,omitempty"`
		Id         string `json:"id,omitempty"`
		Egressrule []struct {
			Cidr              string `json:"cidr,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Account           string `json:"account,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
		} `json:"egressrule,omitempty"`
		Domain      string `json:"domain,omitempty"`
		Account     string `json:"account,omitempty"`
		Description string `json:"description,omitempty"`
		Name        string `json:"name,omitempty"`
	} `json:"securitygroup,omitempty"`
	Templatedisplaytext string            `json:"templatedisplaytext,omitempty"`
	Memory              int               `json:"memory,omitempty"`
	Networkkbswrite     int               `json:"networkkbswrite,omitempty"`
	Name                string            `json:"name,omitempty"`
	Hostname            string            `json:"hostname,omitempty"`
	Templatename        string            `json:"templatename,omitempty"`
	Details             map[string]string `json:"details,omitempty"`
	Isodisplaytext      string            `json:"isodisplaytext,omitempty"`
	Servicestate        string            `json:"servicestate,omitempty"`
	Domainid            string            `json:"domainid,omitempty"`
	Nic                 []struct {
		Ip6address   string   `json:"ip6address,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Type         string   `json:"type,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Id           string   `json:"id,omitempty"`
	} `json:"nic,omitempty"`
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
