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

		var r AttachIsoResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type AttachIsoResponse struct {
	JobID             string `json:"jobid,omitempty"`
	Name              string `json:"name,omitempty"`
	Diskioread        int    `json:"diskioread,omitempty"`
	Instancename      string `json:"instancename,omitempty"`
	Project           string `json:"project,omitempty"`
	Haenable          bool   `json:"haenable,omitempty"`
	Servicestate      string `json:"servicestate,omitempty"`
	Serviceofferingid string `json:"serviceofferingid,omitempty"`
	Passwordenabled   bool   `json:"passwordenabled,omitempty"`
	Password          string `json:"password,omitempty"`
	Diskkbswrite      int    `json:"diskkbswrite,omitempty"`
	Nic               []struct {
		Netmask      string   `json:"netmask,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Id           string   `json:"id,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Type         string   `json:"type,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
	} `json:"nic,omitempty"`
	Isoid               string `json:"isoid,omitempty"`
	Hypervisor          string `json:"hypervisor,omitempty"`
	Templatename        string `json:"templatename,omitempty"`
	Memory              int    `json:"memory,omitempty"`
	Displayname         string `json:"displayname,omitempty"`
	Publicipid          string `json:"publicipid,omitempty"`
	Networkkbsread      int    `json:"networkkbsread,omitempty"`
	Hostname            string `json:"hostname,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Securitygroup       []struct {
		Id       string `json:"id,omitempty"`
		Domainid string `json:"domainid,omitempty"`
		Account  string `json:"account,omitempty"`
		Tags     []struct {
			Domain       string `json:"domain,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Key          string `json:"key,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Value        string `json:"value,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Account      string `json:"account,omitempty"`
			Project      string `json:"project,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
		} `json:"tags,omitempty"`
		Name        string `json:"name,omitempty"`
		Project     string `json:"project,omitempty"`
		Projectid   string `json:"projectid,omitempty"`
		Ingressrule []struct {
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Account           string `json:"account,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
		} `json:"ingressrule,omitempty"`
		Egressrule []struct {
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Account           string `json:"account,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
		} `json:"egressrule,omitempty"`
		Domain      string `json:"domain,omitempty"`
		Description string `json:"description,omitempty"`
	} `json:"securitygroup,omitempty"`
	Groupid             string            `json:"groupid,omitempty"`
	Networkkbswrite     int               `json:"networkkbswrite,omitempty"`
	Cpunumber           int               `json:"cpunumber,omitempty"`
	Guestosid           string            `json:"guestosid,omitempty"`
	Details             map[string]string `json:"details,omitempty"`
	Zonename            string            `json:"zonename,omitempty"`
	Created             string            `json:"created,omitempty"`
	Id                  string            `json:"id,omitempty"`
	Templatedisplaytext string            `json:"templatedisplaytext,omitempty"`
	Rootdeviceid        int               `json:"rootdeviceid,omitempty"`
	Domain              string            `json:"domain,omitempty"`
	Diskiowrite         int               `json:"diskiowrite,omitempty"`
	Isodisplaytext      string            `json:"isodisplaytext,omitempty"`
	Affinitygroup       []struct {
		Description       string   `json:"description,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
		Type              string   `json:"type,omitempty"`
		Account           string   `json:"account,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Name              string   `json:"name,omitempty"`
		Id                string   `json:"id,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Tags []struct {
		Account      string `json:"account,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Value        string `json:"value,omitempty"`
		Key          string `json:"key,omitempty"`
		Project      string `json:"project,omitempty"`
	} `json:"tags,omitempty"`
	Isoname               string `json:"isoname,omitempty"`
	Zoneid                string `json:"zoneid,omitempty"`
	Templateid            string `json:"templateid,omitempty"`
	Domainid              string `json:"domainid,omitempty"`
	Rootdevicetype        string `json:"rootdevicetype,omitempty"`
	Displayvm             bool   `json:"displayvm,omitempty"`
	Publicip              string `json:"publicip,omitempty"`
	Cpuspeed              int    `json:"cpuspeed,omitempty"`
	Diskkbsread           int    `json:"diskkbsread,omitempty"`
	Account               string `json:"account,omitempty"`
	Keypair               string `json:"keypair,omitempty"`
	Isdynamicallyscalable bool   `json:"isdynamicallyscalable,omitempty"`
	Cpuused               string `json:"cpuused,omitempty"`
	State                 string `json:"state,omitempty"`
	Projectid             string `json:"projectid,omitempty"`
	Group                 string `json:"group,omitempty"`
	Forvirtualnetwork     bool   `json:"forvirtualnetwork,omitempty"`
	Hostid                string `json:"hostid,omitempty"`
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

		var r DetachIsoResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type DetachIsoResponse struct {
	JobID                 string `json:"jobid,omitempty"`
	Passwordenabled       bool   `json:"passwordenabled,omitempty"`
	Groupid               string `json:"groupid,omitempty"`
	Memory                int    `json:"memory,omitempty"`
	Networkkbsread        int    `json:"networkkbsread,omitempty"`
	Domainid              string `json:"domainid,omitempty"`
	Templatedisplaytext   string `json:"templatedisplaytext,omitempty"`
	Name                  string `json:"name,omitempty"`
	Instancename          string `json:"instancename,omitempty"`
	Account               string `json:"account,omitempty"`
	Isodisplaytext        string `json:"isodisplaytext,omitempty"`
	Publicipid            string `json:"publicipid,omitempty"`
	Isdynamicallyscalable bool   `json:"isdynamicallyscalable,omitempty"`
	Publicip              string `json:"publicip,omitempty"`
	Rootdeviceid          int    `json:"rootdeviceid,omitempty"`
	Zoneid                string `json:"zoneid,omitempty"`
	Nic                   []struct {
		Networkname  string   `json:"networkname,omitempty"`
		Id           string   `json:"id,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Type         string   `json:"type,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
	} `json:"nic,omitempty"`
	Haenable       bool              `json:"haenable,omitempty"`
	Diskkbsread    int               `json:"diskkbsread,omitempty"`
	Details        map[string]string `json:"details,omitempty"`
	State          string            `json:"state,omitempty"`
	Keypair        string            `json:"keypair,omitempty"`
	Rootdevicetype string            `json:"rootdevicetype,omitempty"`
	Hypervisor     string            `json:"hypervisor,omitempty"`
	Affinitygroup  []struct {
		Domain            string   `json:"domain,omitempty"`
		Account           string   `json:"account,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
		Id                string   `json:"id,omitempty"`
		Description       string   `json:"description,omitempty"`
		Type              string   `json:"type,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Name              string   `json:"name,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Forvirtualnetwork   bool   `json:"forvirtualnetwork,omitempty"`
	Isoname             string `json:"isoname,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
	Cpuspeed            int    `json:"cpuspeed,omitempty"`
	Password            string `json:"password,omitempty"`
	Project             string `json:"project,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Displayvm           bool   `json:"displayvm,omitempty"`
	Diskkbswrite        int    `json:"diskkbswrite,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Isoid               string `json:"isoid,omitempty"`
	Hostid              string `json:"hostid,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
	Group               string `json:"group,omitempty"`
	Templateid          string `json:"templateid,omitempty"`
	Tags                []struct {
		Resourceid   string `json:"resourceid,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Account      string `json:"account,omitempty"`
		Project      string `json:"project,omitempty"`
		Value        string `json:"value,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Key          string `json:"key,omitempty"`
	} `json:"tags,omitempty"`
	Securitygroup []struct {
		Domain string `json:"domain,omitempty"`
		Name   string `json:"name,omitempty"`
		Tags   []struct {
			Domainid     string `json:"domainid,omitempty"`
			Account      string `json:"account,omitempty"`
			Value        string `json:"value,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Key          string `json:"key,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Project      string `json:"project,omitempty"`
		} `json:"tags,omitempty"`
		Egressrule []struct {
			Ruleid            string `json:"ruleid,omitempty"`
			Account           string `json:"account,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
		} `json:"egressrule,omitempty"`
		Ingressrule []struct {
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Account           string `json:"account,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
		} `json:"ingressrule,omitempty"`
		Id          string `json:"id,omitempty"`
		Projectid   string `json:"projectid,omitempty"`
		Domainid    string `json:"domainid,omitempty"`
		Account     string `json:"account,omitempty"`
		Description string `json:"description,omitempty"`
		Project     string `json:"project,omitempty"`
	} `json:"securitygroup,omitempty"`
	Cpunumber       int    `json:"cpunumber,omitempty"`
	Diskioread      int    `json:"diskioread,omitempty"`
	Networkkbswrite int    `json:"networkkbswrite,omitempty"`
	Templatename    string `json:"templatename,omitempty"`
	Servicestate    string `json:"servicestate,omitempty"`
	Displayname     string `json:"displayname,omitempty"`
	Guestosid       string `json:"guestosid,omitempty"`
	Created         string `json:"created,omitempty"`
	Diskiowrite     int    `json:"diskiowrite,omitempty"`
	Projectid       string `json:"projectid,omitempty"`
	Cpuused         string `json:"cpuused,omitempty"`
	Id              string `json:"id,omitempty"`
	Hostname        string `json:"hostname,omitempty"`
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
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, name, l)
	}
	return l.Isos[0].Id, nil
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
	Isdynamicallyscalable bool `json:"isdynamicallyscalable,omitempty"`
	Sshkeyenabled         bool `json:"sshkeyenabled,omitempty"`
	Ispublic              bool `json:"ispublic,omitempty"`
	Bootable              bool `json:"bootable,omitempty"`
	Tags                  []struct {
		Domain       string `json:"domain,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Key          string `json:"key,omitempty"`
		Account      string `json:"account,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Project      string `json:"project,omitempty"`
		Value        string `json:"value,omitempty"`
	} `json:"tags,omitempty"`
	Displaytext      string            `json:"displaytext,omitempty"`
	Zoneid           string            `json:"zoneid,omitempty"`
	Size             int               `json:"size,omitempty"`
	Ostypename       string            `json:"ostypename,omitempty"`
	Details          map[string]string `json:"details,omitempty"`
	Sourcetemplateid string            `json:"sourcetemplateid,omitempty"`
	Format           string            `json:"format,omitempty"`
	Isready          bool              `json:"isready,omitempty"`
	Removed          string            `json:"removed,omitempty"`
	Templatetag      string            `json:"templatetag,omitempty"`
	Zonename         string            `json:"zonename,omitempty"`
	CrossZones       bool              `json:"crossZones,omitempty"`
	Isfeatured       bool              `json:"isfeatured,omitempty"`
	Passwordenabled  bool              `json:"passwordenabled,omitempty"`
	Templatetype     string            `json:"templatetype,omitempty"`
	Id               string            `json:"id,omitempty"`
	Isextractable    bool              `json:"isextractable,omitempty"`
	Ostypeid         string            `json:"ostypeid,omitempty"`
	Account          string            `json:"account,omitempty"`
	Hostid           string            `json:"hostid,omitempty"`
	Checksum         string            `json:"checksum,omitempty"`
	Projectid        string            `json:"projectid,omitempty"`
	Accountid        string            `json:"accountid,omitempty"`
	Name             string            `json:"name,omitempty"`
	Domain           string            `json:"domain,omitempty"`
	Hypervisor       string            `json:"hypervisor,omitempty"`
	Status           string            `json:"status,omitempty"`
	Domainid         string            `json:"domainid,omitempty"`
	Created          string            `json:"created,omitempty"`
	Hostname         string            `json:"hostname,omitempty"`
	Project          string            `json:"project,omitempty"`
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
	Accountid       string `json:"accountid,omitempty"`
	Isfeatured      bool   `json:"isfeatured,omitempty"`
	Account         string `json:"account,omitempty"`
	Removed         string `json:"removed,omitempty"`
	Templatetag     string `json:"templatetag,omitempty"`
	Hostid          string `json:"hostid,omitempty"`
	Size            int    `json:"size,omitempty"`
	Passwordenabled bool   `json:"passwordenabled,omitempty"`
	Sshkeyenabled   bool   `json:"sshkeyenabled,omitempty"`
	CrossZones      bool   `json:"crossZones,omitempty"`
	Format          string `json:"format,omitempty"`
	Domain          string `json:"domain,omitempty"`
	Displaytext     string `json:"displaytext,omitempty"`
	Templatetype    string `json:"templatetype,omitempty"`
	Zonename        string `json:"zonename,omitempty"`
	Hostname        string `json:"hostname,omitempty"`
	Isextractable   bool   `json:"isextractable,omitempty"`
	Tags            []struct {
		Domain       string `json:"domain,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Key          string `json:"key,omitempty"`
		Project      string `json:"project,omitempty"`
		Account      string `json:"account,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Value        string `json:"value,omitempty"`
	} `json:"tags,omitempty"`
	Bootable              bool              `json:"bootable,omitempty"`
	Id                    string            `json:"id,omitempty"`
	Ostypeid              string            `json:"ostypeid,omitempty"`
	Project               string            `json:"project,omitempty"`
	Isdynamicallyscalable bool              `json:"isdynamicallyscalable,omitempty"`
	Ispublic              bool              `json:"ispublic,omitempty"`
	Status                string            `json:"status,omitempty"`
	Domainid              string            `json:"domainid,omitempty"`
	Zoneid                string            `json:"zoneid,omitempty"`
	Projectid             string            `json:"projectid,omitempty"`
	Ostypename            string            `json:"ostypename,omitempty"`
	Checksum              string            `json:"checksum,omitempty"`
	Name                  string            `json:"name,omitempty"`
	Hypervisor            string            `json:"hypervisor,omitempty"`
	Isready               bool              `json:"isready,omitempty"`
	Details               map[string]string `json:"details,omitempty"`
	Sourcetemplateid      string            `json:"sourcetemplateid,omitempty"`
	Created               string            `json:"created,omitempty"`
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
	Sshkeyenabled         bool              `json:"sshkeyenabled,omitempty"`
	Displaytext           string            `json:"displaytext,omitempty"`
	Domain                string            `json:"domain,omitempty"`
	Ostypeid              string            `json:"ostypeid,omitempty"`
	Accountid             string            `json:"accountid,omitempty"`
	Passwordenabled       bool              `json:"passwordenabled,omitempty"`
	Removed               string            `json:"removed,omitempty"`
	Sourcetemplateid      string            `json:"sourcetemplateid,omitempty"`
	Format                string            `json:"format,omitempty"`
	Id                    string            `json:"id,omitempty"`
	Ostypename            string            `json:"ostypename,omitempty"`
	Zoneid                string            `json:"zoneid,omitempty"`
	Details               map[string]string `json:"details,omitempty"`
	Ispublic              bool              `json:"ispublic,omitempty"`
	Hypervisor            string            `json:"hypervisor,omitempty"`
	Project               string            `json:"project,omitempty"`
	CrossZones            bool              `json:"crossZones,omitempty"`
	Domainid              string            `json:"domainid,omitempty"`
	Isdynamicallyscalable bool              `json:"isdynamicallyscalable,omitempty"`
	Name                  string            `json:"name,omitempty"`
	Projectid             string            `json:"projectid,omitempty"`
	Templatetype          string            `json:"templatetype,omitempty"`
	Templatetag           string            `json:"templatetag,omitempty"`
	Checksum              string            `json:"checksum,omitempty"`
	Zonename              string            `json:"zonename,omitempty"`
	Status                string            `json:"status,omitempty"`
	Isextractable         bool              `json:"isextractable,omitempty"`
	Account               string            `json:"account,omitempty"`
	Bootable              bool              `json:"bootable,omitempty"`
	Isready               bool              `json:"isready,omitempty"`
	Hostname              string            `json:"hostname,omitempty"`
	Hostid                string            `json:"hostid,omitempty"`
	Created               string            `json:"created,omitempty"`
	Isfeatured            bool              `json:"isfeatured,omitempty"`
	Tags                  []struct {
		Key          string `json:"key,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Value        string `json:"value,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Project      string `json:"project,omitempty"`
		Account      string `json:"account,omitempty"`
		Domain       string `json:"domain,omitempty"`
	} `json:"tags,omitempty"`
	Size int `json:"size,omitempty"`
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

		var r DeleteIsoResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
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

		var r CopyIsoResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type CopyIsoResponse struct {
	JobID                 string `json:"jobid,omitempty"`
	Isdynamicallyscalable bool   `json:"isdynamicallyscalable,omitempty"`
	Hypervisor            string `json:"hypervisor,omitempty"`
	Account               string `json:"account,omitempty"`
	Checksum              string `json:"checksum,omitempty"`
	Passwordenabled       bool   `json:"passwordenabled,omitempty"`
	Isready               bool   `json:"isready,omitempty"`
	Isfeatured            bool   `json:"isfeatured,omitempty"`
	Isextractable         bool   `json:"isextractable,omitempty"`
	Project               string `json:"project,omitempty"`
	Templatetag           string `json:"templatetag,omitempty"`
	Zonename              string `json:"zonename,omitempty"`
	Format                string `json:"format,omitempty"`
	Ostypename            string `json:"ostypename,omitempty"`
	Templatetype          string `json:"templatetype,omitempty"`
	Name                  string `json:"name,omitempty"`
	Tags                  []struct {
		Account      string `json:"account,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Key          string `json:"key,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Project      string `json:"project,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Value        string `json:"value,omitempty"`
	} `json:"tags,omitempty"`
	CrossZones       bool              `json:"crossZones,omitempty"`
	Id               string            `json:"id,omitempty"`
	Hostid           string            `json:"hostid,omitempty"`
	Projectid        string            `json:"projectid,omitempty"`
	Displaytext      string            `json:"displaytext,omitempty"`
	Zoneid           string            `json:"zoneid,omitempty"`
	Domainid         string            `json:"domainid,omitempty"`
	Removed          string            `json:"removed,omitempty"`
	Sshkeyenabled    bool              `json:"sshkeyenabled,omitempty"`
	Ostypeid         string            `json:"ostypeid,omitempty"`
	Bootable         bool              `json:"bootable,omitempty"`
	Size             int               `json:"size,omitempty"`
	Ispublic         bool              `json:"ispublic,omitempty"`
	Status           string            `json:"status,omitempty"`
	Sourcetemplateid string            `json:"sourcetemplateid,omitempty"`
	Hostname         string            `json:"hostname,omitempty"`
	Domain           string            `json:"domain,omitempty"`
	Accountid        string            `json:"accountid,omitempty"`
	Created          string            `json:"created,omitempty"`
	Details          map[string]string `json:"details,omitempty"`
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
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
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
	Projectids []string `json:"projectids,omitempty"`
	Id         string   `json:"id,omitempty"`
	Account    []string `json:"account,omitempty"`
	Ispublic   bool     `json:"ispublic,omitempty"`
	Domainid   string   `json:"domainid,omitempty"`
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

		var r ExtractIsoResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type ExtractIsoResponse struct {
	JobID            string `json:"jobid,omitempty"`
	Url              string `json:"url,omitempty"`
	Resultstring     string `json:"resultstring,omitempty"`
	ExtractId        string `json:"extractId,omitempty"`
	Status           string `json:"status,omitempty"`
	Created          string `json:"created,omitempty"`
	Storagetype      string `json:"storagetype,omitempty"`
	Uploadpercentage int    `json:"uploadpercentage,omitempty"`
	Zonename         string `json:"zonename,omitempty"`
	Name             string `json:"name,omitempty"`
	ExtractMode      string `json:"extractMode,omitempty"`
	Id               string `json:"id,omitempty"`
	Accountid        string `json:"accountid,omitempty"`
	Zoneid           string `json:"zoneid,omitempty"`
	State            string `json:"state,omitempty"`
}
