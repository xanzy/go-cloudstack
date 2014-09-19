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

type AttachIsoParams struct {
	p map[string]interface{}
}

func (p *AttachIsoParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *AttachIsoParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *AttachIsoParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
	return
}

// You should always use this function to get a new AttachIsoParams instance,
// as then you are sure you have configured all required params
func (s *ISOService) NewAttachIsoParams(id string, virtualmachineid string) *AttachIsoParams {
	p := &AttachIsoParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// Attaches an ISO to a virtual machine.
func (s *ISOService) AttachIso(p *AttachIsoParams) (*AttachIsoResponse, error) {
	resp, err := s.cs.newRequest("attachIso", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AttachIsoResponse
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

type AttachIsoResponse struct {
	JobID             string `json:"jobid,omitempty"`
	Group             string `json:"group,omitempty"`
	Servicestate      string `json:"servicestate,omitempty"`
	Displayvm         bool   `json:"displayvm,omitempty"`
	Diskkbswrite      int    `json:"diskkbswrite,omitempty"`
	Project           string `json:"project,omitempty"`
	Publicip          string `json:"publicip,omitempty"`
	Zonename          string `json:"zonename,omitempty"`
	Id                string `json:"id,omitempty"`
	Groupid           string `json:"groupid,omitempty"`
	Haenable          bool   `json:"haenable,omitempty"`
	Created           string `json:"created,omitempty"`
	Password          string `json:"password,omitempty"`
	Hostid            string `json:"hostid,omitempty"`
	Serviceofferingid string `json:"serviceofferingid,omitempty"`
	Diskkbsread       int    `json:"diskkbsread,omitempty"`
	Forvirtualnetwork bool   `json:"forvirtualnetwork,omitempty"`
	Isoname           string `json:"isoname,omitempty"`
	Nic               []struct {
		Networkname  string   `json:"networkname,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Id           string   `json:"id,omitempty"`
		Type         string   `json:"type,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
	} `json:"nic,omitempty"`
	Domain         string `json:"domain,omitempty"`
	Templateid     string `json:"templateid,omitempty"`
	Keypair        string `json:"keypair,omitempty"`
	Diskiowrite    int    `json:"diskiowrite,omitempty"`
	Rootdevicetype string `json:"rootdevicetype,omitempty"`
	Name           string `json:"name,omitempty"`
	Instancename   string `json:"instancename,omitempty"`
	Isoid          string `json:"isoid,omitempty"`
	Publicipid     string `json:"publicipid,omitempty"`
	Tags           []struct {
		Account      string `json:"account,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Project      string `json:"project,omitempty"`
		Value        string `json:"value,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Key          string `json:"key,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
	} `json:"tags,omitempty"`
	Cpuspeed              int    `json:"cpuspeed,omitempty"`
	Cpunumber             int    `json:"cpunumber,omitempty"`
	Isdynamicallyscalable bool   `json:"isdynamicallyscalable,omitempty"`
	State                 string `json:"state,omitempty"`
	Memory                int    `json:"memory,omitempty"`
	Affinitygroup         []struct {
		Account           string   `json:"account,omitempty"`
		Name              string   `json:"name,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
		Type              string   `json:"type,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		Description       string   `json:"description,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Id                string   `json:"id,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Templatename        string            `json:"templatename,omitempty"`
	Zoneid              string            `json:"zoneid,omitempty"`
	Projectid           string            `json:"projectid,omitempty"`
	Networkkbsread      int               `json:"networkkbsread,omitempty"`
	Isodisplaytext      string            `json:"isodisplaytext,omitempty"`
	Rootdeviceid        int               `json:"rootdeviceid,omitempty"`
	Details             map[string]string `json:"details,omitempty"`
	Displayname         string            `json:"displayname,omitempty"`
	Serviceofferingname string            `json:"serviceofferingname,omitempty"`
	Templatedisplaytext string            `json:"templatedisplaytext,omitempty"`
	Domainid            string            `json:"domainid,omitempty"`
	Networkkbswrite     int               `json:"networkkbswrite,omitempty"`
	Cpuused             string            `json:"cpuused,omitempty"`
	Diskioread          int               `json:"diskioread,omitempty"`
	Hostname            string            `json:"hostname,omitempty"`
	Securitygroup       []struct {
		Project     string `json:"project,omitempty"`
		Description string `json:"description,omitempty"`
		Id          string `json:"id,omitempty"`
		Name        string `json:"name,omitempty"`
		Ingressrule []struct {
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Account           string `json:"account,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
		} `json:"ingressrule,omitempty"`
		Tags []struct {
			Domainid     string `json:"domainid,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Project      string `json:"project,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
			Value        string `json:"value,omitempty"`
			Key          string `json:"key,omitempty"`
			Account      string `json:"account,omitempty"`
		} `json:"tags,omitempty"`
		Domain     string `json:"domain,omitempty"`
		Egressrule []struct {
			Icmptype          int    `json:"icmptype,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Account           string `json:"account,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Startport         int    `json:"startport,omitempty"`
		} `json:"egressrule,omitempty"`
		Projectid string `json:"projectid,omitempty"`
		Account   string `json:"account,omitempty"`
		Domainid  string `json:"domainid,omitempty"`
	} `json:"securitygroup,omitempty"`
	Passwordenabled bool   `json:"passwordenabled,omitempty"`
	Hypervisor      string `json:"hypervisor,omitempty"`
	Guestosid       string `json:"guestosid,omitempty"`
	Account         string `json:"account,omitempty"`
}

type DetachIsoParams struct {
	p map[string]interface{}
}

func (p *DetachIsoParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *DetachIsoParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
	return
}

// You should always use this function to get a new DetachIsoParams instance,
// as then you are sure you have configured all required params
func (s *ISOService) NewDetachIsoParams(virtualmachineid string) *DetachIsoParams {
	p := &DetachIsoParams{}
	p.p = make(map[string]interface{})
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// Detaches any ISO file (if any) currently attached to a virtual machine.
func (s *ISOService) DetachIso(p *DetachIsoParams) (*DetachIsoResponse, error) {
	resp, err := s.cs.newRequest("detachIso", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DetachIsoResponse
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

type DetachIsoResponse struct {
	JobID        string `json:"jobid,omitempty"`
	Keypair      string `json:"keypair,omitempty"`
	Cpuused      string `json:"cpuused,omitempty"`
	State        string `json:"state,omitempty"`
	Domainid     string `json:"domainid,omitempty"`
	Diskiowrite  int    `json:"diskiowrite,omitempty"`
	Groupid      string `json:"groupid,omitempty"`
	Templatename string `json:"templatename,omitempty"`
	Tags         []struct {
		Account      string `json:"account,omitempty"`
		Project      string `json:"project,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Value        string `json:"value,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Key          string `json:"key,omitempty"`
	} `json:"tags,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Networkkbsread      int    `json:"networkkbsread,omitempty"`
	Id                  string `json:"id,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Publicip            string `json:"publicip,omitempty"`
	Isoid               string `json:"isoid,omitempty"`
	Securitygroup       []struct {
		Project    string `json:"project,omitempty"`
		Domainid   string `json:"domainid,omitempty"`
		Egressrule []struct {
			Startport         int    `json:"startport,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Account           string `json:"account,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
		} `json:"egressrule,omitempty"`
		Id          string `json:"id,omitempty"`
		Description string `json:"description,omitempty"`
		Name        string `json:"name,omitempty"`
		Ingressrule []struct {
			Cidr              string `json:"cidr,omitempty"`
			Account           string `json:"account,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Startport         int    `json:"startport,omitempty"`
		} `json:"ingressrule,omitempty"`
		Projectid string `json:"projectid,omitempty"`
		Account   string `json:"account,omitempty"`
		Tags      []struct {
			Value        string `json:"value,omitempty"`
			Project      string `json:"project,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
			Account      string `json:"account,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Key          string `json:"key,omitempty"`
		} `json:"tags,omitempty"`
		Domain string `json:"domain,omitempty"`
	} `json:"securitygroup,omitempty"`
	Account             string `json:"account,omitempty"`
	Hypervisor          string `json:"hypervisor,omitempty"`
	Displayvm           bool   `json:"displayvm,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
	Projectid           string `json:"projectid,omitempty"`
	Password            string `json:"password,omitempty"`
	Created             string `json:"created,omitempty"`
	Hostid              string `json:"hostid,omitempty"`
	Diskkbsread         int    `json:"diskkbsread,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
	Group               string `json:"group,omitempty"`
	Templatedisplaytext string `json:"templatedisplaytext,omitempty"`
	Cpuspeed            int    `json:"cpuspeed,omitempty"`
	Name                string `json:"name,omitempty"`
	Zoneid              string `json:"zoneid,omitempty"`
	Displayname         string `json:"displayname,omitempty"`
	Memory              int    `json:"memory,omitempty"`
	Diskioread          int    `json:"diskioread,omitempty"`
	Hostname            string `json:"hostname,omitempty"`
	Affinitygroup       []struct {
		Type              string   `json:"type,omitempty"`
		Name              string   `json:"name,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
		Id                string   `json:"id,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		Description       string   `json:"description,omitempty"`
		Account           string   `json:"account,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Templateid string `json:"templateid,omitempty"`
	Nic        []struct {
		Macaddress   string   `json:"macaddress,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Id           string   `json:"id,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Type         string   `json:"type,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
	} `json:"nic,omitempty"`
	Rootdeviceid          int               `json:"rootdeviceid,omitempty"`
	Publicipid            string            `json:"publicipid,omitempty"`
	Forvirtualnetwork     bool              `json:"forvirtualnetwork,omitempty"`
	Isodisplaytext        string            `json:"isodisplaytext,omitempty"`
	Haenable              bool              `json:"haenable,omitempty"`
	Details               map[string]string `json:"details,omitempty"`
	Instancename          string            `json:"instancename,omitempty"`
	Passwordenabled       bool              `json:"passwordenabled,omitempty"`
	Diskkbswrite          int               `json:"diskkbswrite,omitempty"`
	Guestosid             string            `json:"guestosid,omitempty"`
	Servicestate          string            `json:"servicestate,omitempty"`
	Networkkbswrite       int               `json:"networkkbswrite,omitempty"`
	Isoname               string            `json:"isoname,omitempty"`
	Project               string            `json:"project,omitempty"`
	Isdynamicallyscalable bool              `json:"isdynamicallyscalable,omitempty"`
	Rootdevicetype        string            `json:"rootdevicetype,omitempty"`
	Cpunumber             int               `json:"cpunumber,omitempty"`
}

type ListIsosParams struct {
	p map[string]interface{}
}

func (p *ListIsosParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["bootable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("bootable", vv)
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["isofilter"]; found {
		u.Set("isofilter", v.(string))
	}
	if v, found := p.p["ispublic"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("ispublic", vv)
	}
	if v, found := p.p["isready"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isready", vv)
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
	if v, found := p.p["showremoved"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("showremoved", vv)
	}
	if v, found := p.p["tags"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListIsosParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListIsosParams) SetBootable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bootable"] = v
	return
}

func (p *ListIsosParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListIsosParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
	return
}

func (p *ListIsosParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListIsosParams) SetIsofilter(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isofilter"] = v
	return
}

func (p *ListIsosParams) SetIspublic(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ispublic"] = v
	return
}

func (p *ListIsosParams) SetIsready(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isready"] = v
	return
}

func (p *ListIsosParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
	return
}

func (p *ListIsosParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListIsosParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
	return
}

func (p *ListIsosParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *ListIsosParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListIsosParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListIsosParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *ListIsosParams) SetShowremoved(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["showremoved"] = v
	return
}

func (p *ListIsosParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
	return
}

func (p *ListIsosParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new ListIsosParams instance,
// as then you are sure you have configured all required params
func (s *ISOService) NewListIsosParams() *ListIsosParams {
	p := &ListIsosParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ISOService) GetIsoID(name string) (string, error) {
	p := &ListIsosParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	l, err := s.ListIsos(p)
	if err != nil {
		return "", err
	}

	if l.Count == 0 {
		return "", fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Isos[0].Id, nil
	}

	if l.Count > 1 {
		for _, v := range l.Isos {
			if v.Name == name {
				return v.Id, nil
			}
		}
	}
	return "", fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ISOService) GetIsoByName(name string) (*Iso, int, error) {
	id, err := s.GetIsoID(name)
	if err != nil {
		return nil, -1, err
	}

	r, count, err := s.GetIsoByID(id)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ISOService) GetIsoByID(id string) (*Iso, int, error) {
	p := &ListIsosParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	l, err := s.ListIsos(p)
	if err != nil {
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.Isos[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Iso UUID: %s!", id)
}

// Lists all available ISO files.
func (s *ISOService) ListIsos(p *ListIsosParams) (*ListIsosResponse, error) {
	resp, err := s.cs.newRequest("listIsos", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListIsosResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListIsosResponse struct {
	Count int    `json:"count"`
	Isos  []*Iso `json:"iso"`
}

type Iso struct {
	Isdynamicallyscalable bool              `json:"isdynamicallyscalable,omitempty"`
	Isextractable         bool              `json:"isextractable,omitempty"`
	Passwordenabled       bool              `json:"passwordenabled,omitempty"`
	Details               map[string]string `json:"details,omitempty"`
	Id                    string            `json:"id,omitempty"`
	Templatetag           string            `json:"templatetag,omitempty"`
	Sourcetemplateid      string            `json:"sourcetemplateid,omitempty"`
	Created               string            `json:"created,omitempty"`
	Ispublic              bool              `json:"ispublic,omitempty"`
	Zoneid                string            `json:"zoneid,omitempty"`
	Status                string            `json:"status,omitempty"`
	Hostname              string            `json:"hostname,omitempty"`
	Bootable              bool              `json:"bootable,omitempty"`
	Zonename              string            `json:"zonename,omitempty"`
	Sshkeyenabled         bool              `json:"sshkeyenabled,omitempty"`
	Ostypeid              string            `json:"ostypeid,omitempty"`
	Hypervisor            string            `json:"hypervisor,omitempty"`
	Size                  int               `json:"size,omitempty"`
	Hostid                string            `json:"hostid,omitempty"`
	Projectid             string            `json:"projectid,omitempty"`
	Domainid              string            `json:"domainid,omitempty"`
	Isready               bool              `json:"isready,omitempty"`
	Name                  string            `json:"name,omitempty"`
	Tags                  []struct {
		Key          string `json:"key,omitempty"`
		Value        string `json:"value,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Project      string `json:"project,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Account      string `json:"account,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
	} `json:"tags,omitempty"`
	Project      string `json:"project,omitempty"`
	Isfeatured   bool   `json:"isfeatured,omitempty"`
	Templatetype string `json:"templatetype,omitempty"`
	CrossZones   bool   `json:"crossZones,omitempty"`
	Domain       string `json:"domain,omitempty"`
	Checksum     string `json:"checksum,omitempty"`
	Displaytext  string `json:"displaytext,omitempty"`
	Removed      string `json:"removed,omitempty"`
	Format       string `json:"format,omitempty"`
	Ostypename   string `json:"ostypename,omitempty"`
	Accountid    string `json:"accountid,omitempty"`
	Account      string `json:"account,omitempty"`
}

type RegisterIsoParams struct {
	p map[string]interface{}
}

func (p *RegisterIsoParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["bootable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("bootable", vv)
	}
	if v, found := p.p["checksum"]; found {
		u.Set("checksum", v.(string))
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["imagestoreuuid"]; found {
		u.Set("imagestoreuuid", v.(string))
	}
	if v, found := p.p["isdynamicallyscalable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isdynamicallyscalable", vv)
	}
	if v, found := p.p["isextractable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isextractable", vv)
	}
	if v, found := p.p["isfeatured"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isfeatured", vv)
	}
	if v, found := p.p["ispublic"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("ispublic", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["ostypeid"]; found {
		u.Set("ostypeid", v.(string))
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

func (p *RegisterIsoParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *RegisterIsoParams) SetBootable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bootable"] = v
	return
}

func (p *RegisterIsoParams) SetChecksum(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["checksum"] = v
	return
}

func (p *RegisterIsoParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
	return
}

func (p *RegisterIsoParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *RegisterIsoParams) SetImagestoreuuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["imagestoreuuid"] = v
	return
}

func (p *RegisterIsoParams) SetIsdynamicallyscalable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isdynamicallyscalable"] = v
	return
}

func (p *RegisterIsoParams) SetIsextractable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isextractable"] = v
	return
}

func (p *RegisterIsoParams) SetIsfeatured(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isfeatured"] = v
	return
}

func (p *RegisterIsoParams) SetIspublic(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ispublic"] = v
	return
}

func (p *RegisterIsoParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *RegisterIsoParams) SetOstypeid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ostypeid"] = v
	return
}

func (p *RegisterIsoParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *RegisterIsoParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
	return
}

func (p *RegisterIsoParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new RegisterIsoParams instance,
// as then you are sure you have configured all required params
func (s *ISOService) NewRegisterIsoParams(displaytext string, name string, url string, zoneid string) *RegisterIsoParams {
	p := &RegisterIsoParams{}
	p.p = make(map[string]interface{})
	p.p["displaytext"] = displaytext
	p.p["name"] = name
	p.p["url"] = url
	p.p["zoneid"] = zoneid
	return p
}

// Registers an existing ISO into the CloudStack Cloud.
func (s *ISOService) RegisterIso(p *RegisterIsoParams) (*RegisterIsoResponse, error) {
	resp, err := s.cs.newRequest("registerIso", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RegisterIsoResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type RegisterIsoResponse struct {
	Projectid             string            `json:"projectid,omitempty"`
	Zoneid                string            `json:"zoneid,omitempty"`
	Ostypeid              string            `json:"ostypeid,omitempty"`
	Domainid              string            `json:"domainid,omitempty"`
	Sourcetemplateid      string            `json:"sourcetemplateid,omitempty"`
	Created               string            `json:"created,omitempty"`
	Hostid                string            `json:"hostid,omitempty"`
	Details               map[string]string `json:"details,omitempty"`
	Ispublic              bool              `json:"ispublic,omitempty"`
	Name                  string            `json:"name,omitempty"`
	Displaytext           string            `json:"displaytext,omitempty"`
	Project               string            `json:"project,omitempty"`
	Removed               string            `json:"removed,omitempty"`
	Isdynamicallyscalable bool              `json:"isdynamicallyscalable,omitempty"`
	Accountid             string            `json:"accountid,omitempty"`
	Format                string            `json:"format,omitempty"`
	Zonename              string            `json:"zonename,omitempty"`
	Templatetype          string            `json:"templatetype,omitempty"`
	Hostname              string            `json:"hostname,omitempty"`
	Checksum              string            `json:"checksum,omitempty"`
	Id                    string            `json:"id,omitempty"`
	CrossZones            bool              `json:"crossZones,omitempty"`
	Tags                  []struct {
		Resourcetype string `json:"resourcetype,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Project      string `json:"project,omitempty"`
		Account      string `json:"account,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Key          string `json:"key,omitempty"`
		Value        string `json:"value,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
	} `json:"tags,omitempty"`
	Bootable        bool   `json:"bootable,omitempty"`
	Isready         bool   `json:"isready,omitempty"`
	Templatetag     string `json:"templatetag,omitempty"`
	Passwordenabled bool   `json:"passwordenabled,omitempty"`
	Domain          string `json:"domain,omitempty"`
	Account         string `json:"account,omitempty"`
	Isextractable   bool   `json:"isextractable,omitempty"`
	Sshkeyenabled   bool   `json:"sshkeyenabled,omitempty"`
	Size            int    `json:"size,omitempty"`
	Ostypename      string `json:"ostypename,omitempty"`
	Status          string `json:"status,omitempty"`
	Isfeatured      bool   `json:"isfeatured,omitempty"`
	Hypervisor      string `json:"hypervisor,omitempty"`
}

type UpdateIsoParams struct {
	p map[string]interface{}
}

func (p *UpdateIsoParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["bootable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("bootable", vv)
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["format"]; found {
		u.Set("format", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["isdynamicallyscalable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isdynamicallyscalable", vv)
	}
	if v, found := p.p["isrouting"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrouting", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["ostypeid"]; found {
		u.Set("ostypeid", v.(string))
	}
	if v, found := p.p["passwordenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("passwordenabled", vv)
	}
	if v, found := p.p["sortkey"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("sortkey", vv)
	}
	return u
}

func (p *UpdateIsoParams) SetBootable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bootable"] = v
	return
}

func (p *UpdateIsoParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
	return
}

func (p *UpdateIsoParams) SetFormat(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["format"] = v
	return
}

func (p *UpdateIsoParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *UpdateIsoParams) SetIsdynamicallyscalable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isdynamicallyscalable"] = v
	return
}

func (p *UpdateIsoParams) SetIsrouting(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrouting"] = v
	return
}

func (p *UpdateIsoParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *UpdateIsoParams) SetOstypeid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ostypeid"] = v
	return
}

func (p *UpdateIsoParams) SetPasswordenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["passwordenabled"] = v
	return
}

func (p *UpdateIsoParams) SetSortkey(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sortkey"] = v
	return
}

// You should always use this function to get a new UpdateIsoParams instance,
// as then you are sure you have configured all required params
func (s *ISOService) NewUpdateIsoParams(id string) *UpdateIsoParams {
	p := &UpdateIsoParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates an ISO file.
func (s *ISOService) UpdateIso(p *UpdateIsoParams) (*UpdateIsoResponse, error) {
	resp, err := s.cs.newRequest("updateIso", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateIsoResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type UpdateIsoResponse struct {
	Hypervisor            string            `json:"hypervisor,omitempty"`
	Isextractable         bool              `json:"isextractable,omitempty"`
	Projectid             string            `json:"projectid,omitempty"`
	Account               string            `json:"account,omitempty"`
	Zoneid                string            `json:"zoneid,omitempty"`
	Sshkeyenabled         bool              `json:"sshkeyenabled,omitempty"`
	Hostname              string            `json:"hostname,omitempty"`
	Domainid              string            `json:"domainid,omitempty"`
	Name                  string            `json:"name,omitempty"`
	Id                    string            `json:"id,omitempty"`
	Zonename              string            `json:"zonename,omitempty"`
	CrossZones            bool              `json:"crossZones,omitempty"`
	Templatetag           string            `json:"templatetag,omitempty"`
	Format                string            `json:"format,omitempty"`
	Removed               string            `json:"removed,omitempty"`
	Bootable              bool              `json:"bootable,omitempty"`
	Templatetype          string            `json:"templatetype,omitempty"`
	Ostypeid              string            `json:"ostypeid,omitempty"`
	Checksum              string            `json:"checksum,omitempty"`
	Size                  int               `json:"size,omitempty"`
	Ostypename            string            `json:"ostypename,omitempty"`
	Accountid             string            `json:"accountid,omitempty"`
	Isdynamicallyscalable bool              `json:"isdynamicallyscalable,omitempty"`
	Created               string            `json:"created,omitempty"`
	Status                string            `json:"status,omitempty"`
	Details               map[string]string `json:"details,omitempty"`
	Isready               bool              `json:"isready,omitempty"`
	Ispublic              bool              `json:"ispublic,omitempty"`
	Displaytext           string            `json:"displaytext,omitempty"`
	Domain                string            `json:"domain,omitempty"`
	Project               string            `json:"project,omitempty"`
	Hostid                string            `json:"hostid,omitempty"`
	Isfeatured            bool              `json:"isfeatured,omitempty"`
	Tags                  []struct {
		Resourcetype string `json:"resourcetype,omitempty"`
		Project      string `json:"project,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Account      string `json:"account,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Value        string `json:"value,omitempty"`
		Key          string `json:"key,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
	} `json:"tags,omitempty"`
	Sourcetemplateid string `json:"sourcetemplateid,omitempty"`
	Passwordenabled  bool   `json:"passwordenabled,omitempty"`
}

type DeleteIsoParams struct {
	p map[string]interface{}
}

func (p *DeleteIsoParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *DeleteIsoParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *DeleteIsoParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new DeleteIsoParams instance,
// as then you are sure you have configured all required params
func (s *ISOService) NewDeleteIsoParams(id string) *DeleteIsoParams {
	p := &DeleteIsoParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes an ISO file.
func (s *ISOService) DeleteIso(p *DeleteIsoParams) (*DeleteIsoResponse, error) {
	resp, err := s.cs.newRequest("deleteIso", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteIsoResponse
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

type DeleteIsoResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Success     bool   `json:"success,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
}

type CopyIsoParams struct {
	p map[string]interface{}
}

func (p *CopyIsoParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["destzoneid"]; found {
		u.Set("destzoneid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["sourcezoneid"]; found {
		u.Set("sourcezoneid", v.(string))
	}
	return u
}

func (p *CopyIsoParams) SetDestzoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["destzoneid"] = v
	return
}

func (p *CopyIsoParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *CopyIsoParams) SetSourcezoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sourcezoneid"] = v
	return
}

// You should always use this function to get a new CopyIsoParams instance,
// as then you are sure you have configured all required params
func (s *ISOService) NewCopyIsoParams(destzoneid string, id string) *CopyIsoParams {
	p := &CopyIsoParams{}
	p.p = make(map[string]interface{})
	p.p["destzoneid"] = destzoneid
	p.p["id"] = id
	return p
}

// Copies an iso from one zone to another.
func (s *ISOService) CopyIso(p *CopyIsoParams) (*CopyIsoResponse, error) {
	resp, err := s.cs.newRequest("copyIso", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CopyIsoResponse
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

type CopyIsoResponse struct {
	JobID         string            `json:"jobid,omitempty"`
	Hostname      string            `json:"hostname,omitempty"`
	Ispublic      bool              `json:"ispublic,omitempty"`
	Id            string            `json:"id,omitempty"`
	Status        string            `json:"status,omitempty"`
	Templatetag   string            `json:"templatetag,omitempty"`
	Zonename      string            `json:"zonename,omitempty"`
	Size          int               `json:"size,omitempty"`
	Removed       string            `json:"removed,omitempty"`
	Domainid      string            `json:"domainid,omitempty"`
	Name          string            `json:"name,omitempty"`
	Zoneid        string            `json:"zoneid,omitempty"`
	Sshkeyenabled bool              `json:"sshkeyenabled,omitempty"`
	Displaytext   string            `json:"displaytext,omitempty"`
	Ostypeid      string            `json:"ostypeid,omitempty"`
	Checksum      string            `json:"checksum,omitempty"`
	Project       string            `json:"project,omitempty"`
	Ostypename    string            `json:"ostypename,omitempty"`
	Isready       bool              `json:"isready,omitempty"`
	Details       map[string]string `json:"details,omitempty"`
	Isfeatured    bool              `json:"isfeatured,omitempty"`
	Accountid     string            `json:"accountid,omitempty"`
	Domain        string            `json:"domain,omitempty"`
	Created       string            `json:"created,omitempty"`
	Projectid     string            `json:"projectid,omitempty"`
	Hypervisor    string            `json:"hypervisor,omitempty"`
	Bootable      bool              `json:"bootable,omitempty"`
	Tags          []struct {
		Customer     string `json:"customer,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Key          string `json:"key,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Account      string `json:"account,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Project      string `json:"project,omitempty"`
		Value        string `json:"value,omitempty"`
	} `json:"tags,omitempty"`
	Account               string `json:"account,omitempty"`
	Format                string `json:"format,omitempty"`
	Isextractable         bool   `json:"isextractable,omitempty"`
	Isdynamicallyscalable bool   `json:"isdynamicallyscalable,omitempty"`
	Templatetype          string `json:"templatetype,omitempty"`
	CrossZones            bool   `json:"crossZones,omitempty"`
	Passwordenabled       bool   `json:"passwordenabled,omitempty"`
	Sourcetemplateid      string `json:"sourcetemplateid,omitempty"`
	Hostid                string `json:"hostid,omitempty"`
}

type UpdateIsoPermissionsParams struct {
	p map[string]interface{}
}

func (p *UpdateIsoPermissionsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["accounts"]; found {
		vv := strings.Join(v.([]string), ", ")
		u.Set("accounts", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["isextractable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isextractable", vv)
	}
	if v, found := p.p["isfeatured"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isfeatured", vv)
	}
	if v, found := p.p["ispublic"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("ispublic", vv)
	}
	if v, found := p.p["op"]; found {
		u.Set("op", v.(string))
	}
	if v, found := p.p["projectids"]; found {
		vv := strings.Join(v.([]string), ", ")
		u.Set("projectids", vv)
	}
	return u
}

func (p *UpdateIsoPermissionsParams) SetAccounts(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accounts"] = v
	return
}

func (p *UpdateIsoPermissionsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *UpdateIsoPermissionsParams) SetIsextractable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isextractable"] = v
	return
}

func (p *UpdateIsoPermissionsParams) SetIsfeatured(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isfeatured"] = v
	return
}

func (p *UpdateIsoPermissionsParams) SetIspublic(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ispublic"] = v
	return
}

func (p *UpdateIsoPermissionsParams) SetOp(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["op"] = v
	return
}

func (p *UpdateIsoPermissionsParams) SetProjectids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectids"] = v
	return
}

// You should always use this function to get a new UpdateIsoPermissionsParams instance,
// as then you are sure you have configured all required params
func (s *ISOService) NewUpdateIsoPermissionsParams(id string) *UpdateIsoPermissionsParams {
	p := &UpdateIsoPermissionsParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates iso permissions
func (s *ISOService) UpdateIsoPermissions(p *UpdateIsoPermissionsParams) (*UpdateIsoPermissionsResponse, error) {
	resp, err := s.cs.newRequest("updateIsoPermissions", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateIsoPermissionsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type UpdateIsoPermissionsResponse struct {
	Success     string `json:"success,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
}

type ListIsoPermissionsParams struct {
	p map[string]interface{}
}

func (p *ListIsoPermissionsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *ListIsoPermissionsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new ListIsoPermissionsParams instance,
// as then you are sure you have configured all required params
func (s *ISOService) NewListIsoPermissionsParams(id string) *ListIsoPermissionsParams {
	p := &ListIsoPermissionsParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ISOService) GetIsoPermissionByID(id string) (*IsoPermission, int, error) {
	p := &ListIsoPermissionsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id
	p.p["id"] = id

	l, err := s.ListIsoPermissions(p)
	if err != nil {
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.IsoPermissions[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for IsoPermission UUID: %s!", id)
}

// List iso visibility and all accounts that have permissions to view this iso.
func (s *ISOService) ListIsoPermissions(p *ListIsoPermissionsParams) (*ListIsoPermissionsResponse, error) {
	resp, err := s.cs.newRequest("listIsoPermissions", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListIsoPermissionsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListIsoPermissionsResponse struct {
	Count          int              `json:"count"`
	IsoPermissions []*IsoPermission `json:"isopermission"`
}

type IsoPermission struct {
	Ispublic   bool     `json:"ispublic,omitempty"`
	Id         string   `json:"id,omitempty"`
	Domainid   string   `json:"domainid,omitempty"`
	Projectids []string `json:"projectids,omitempty"`
	Account    []string `json:"account,omitempty"`
}

type ExtractIsoParams struct {
	p map[string]interface{}
}

func (p *ExtractIsoParams) toURLValues() url.Values {
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

func (p *ExtractIsoParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ExtractIsoParams) SetMode(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["mode"] = v
	return
}

func (p *ExtractIsoParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
	return
}

func (p *ExtractIsoParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new ExtractIsoParams instance,
// as then you are sure you have configured all required params
func (s *ISOService) NewExtractIsoParams(id string, mode string) *ExtractIsoParams {
	p := &ExtractIsoParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["mode"] = mode
	return p
}

// Extracts an ISO
func (s *ISOService) ExtractIso(p *ExtractIsoParams) (*ExtractIsoResponse, error) {
	resp, err := s.cs.newRequest("extractIso", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ExtractIsoResponse
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

type ExtractIsoResponse struct {
	JobID            string `json:"jobid,omitempty"`
	Name             string `json:"name,omitempty"`
	Zoneid           string `json:"zoneid,omitempty"`
	ExtractMode      string `json:"extractMode,omitempty"`
	ExtractId        string `json:"extractId,omitempty"`
	Accountid        string `json:"accountid,omitempty"`
	Created          string `json:"created,omitempty"`
	Status           string `json:"status,omitempty"`
	Url              string `json:"url,omitempty"`
	State            string `json:"state,omitempty"`
	Resultstring     string `json:"resultstring,omitempty"`
	Id               string `json:"id,omitempty"`
	Zonename         string `json:"zonename,omitempty"`
	Uploadpercentage int    `json:"uploadpercentage,omitempty"`
	Storagetype      string `json:"storagetype,omitempty"`
}
