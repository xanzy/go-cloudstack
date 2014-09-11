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

package cloudstack44

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
	JobID                 string            `json:"jobid,omitempty"`
	Isoid                 string            `json:"isoid,omitempty"`
	Groupid               string            `json:"groupid,omitempty"`
	Memory                int               `json:"memory,omitempty"`
	Details               map[string]string `json:"details,omitempty"`
	Domainid              string            `json:"domainid,omitempty"`
	Networkkbsread        int               `json:"networkkbsread,omitempty"`
	Serviceofferingname   string            `json:"serviceofferingname,omitempty"`
	Domain                string            `json:"domain,omitempty"`
	Isdynamicallyscalable bool              `json:"isdynamicallyscalable,omitempty"`
	Guestosid             string            `json:"guestosid,omitempty"`
	Securitygroup         []struct {
		Domain    string `json:"domain,omitempty"`
		Project   string `json:"project,omitempty"`
		Projectid string `json:"projectid,omitempty"`
		Tags      []struct {
			Domainid     string `json:"domainid,omitempty"`
			Project      string `json:"project,omitempty"`
			Account      string `json:"account,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Value        string `json:"value,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Key          string `json:"key,omitempty"`
		} `json:"tags,omitempty"`
		Account    string `json:"account,omitempty"`
		Domainid   string `json:"domainid,omitempty"`
		Egressrule []struct {
			Account           string `json:"account,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
		} `json:"egressrule,omitempty"`
		Id          string `json:"id,omitempty"`
		Description string `json:"description,omitempty"`
		Ingressrule []struct {
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Account           string `json:"account,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
		} `json:"ingressrule,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"securitygroup,omitempty"`
	Zoneid              string `json:"zoneid,omitempty"`
	Id                  string `json:"id,omitempty"`
	Account             string `json:"account,omitempty"`
	Name                string `json:"name,omitempty"`
	Project             string `json:"project,omitempty"`
	Cpuused             string `json:"cpuused,omitempty"`
	Diskiowrite         int    `json:"diskiowrite,omitempty"`
	State               string `json:"state,omitempty"`
	Hostname            string `json:"hostname,omitempty"`
	Templatedisplaytext string `json:"templatedisplaytext,omitempty"`
	Passwordenabled     bool   `json:"passwordenabled,omitempty"`
	Isoname             string `json:"isoname,omitempty"`
	Cpuspeed            int    `json:"cpuspeed,omitempty"`
	Displayname         string `json:"displayname,omitempty"`
	Diskioread          int    `json:"diskioread,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
	Rootdevicetype      string `json:"rootdevicetype,omitempty"`
	Instancename        string `json:"instancename,omitempty"`
	Displayvm           bool   `json:"displayvm,omitempty"`
	Rootdeviceid        int    `json:"rootdeviceid,omitempty"`
	Publicip            string `json:"publicip,omitempty"`
	Haenable            bool   `json:"haenable,omitempty"`
	Templateid          string `json:"templateid,omitempty"`
	Affinitygroup       []struct {
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Type              string   `json:"type,omitempty"`
		Id                string   `json:"id,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		Account           string   `json:"account,omitempty"`
		Name              string   `json:"name,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
		Description       string   `json:"description,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Password string `json:"password,omitempty"`
	Keypair  string `json:"keypair,omitempty"`
	Zonename string `json:"zonename,omitempty"`
	Tags     []struct {
		Key          string `json:"key,omitempty"`
		Project      string `json:"project,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Account      string `json:"account,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Value        string `json:"value,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
	} `json:"tags,omitempty"`
	Hypervisor      string `json:"hypervisor,omitempty"`
	Networkkbswrite int    `json:"networkkbswrite,omitempty"`
	Isodisplaytext  string `json:"isodisplaytext,omitempty"`
	Hostid          string `json:"hostid,omitempty"`
	Diskkbsread     int    `json:"diskkbsread,omitempty"`
	Group           string `json:"group,omitempty"`
	Projectid       string `json:"projectid,omitempty"`
	Nic             []struct {
		Netmask      string   `json:"netmask,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Type         string   `json:"type,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Id           string   `json:"id,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
	} `json:"nic,omitempty"`
	Created           string `json:"created,omitempty"`
	Forvirtualnetwork bool   `json:"forvirtualnetwork,omitempty"`
	Cpunumber         int    `json:"cpunumber,omitempty"`
	Templatename      string `json:"templatename,omitempty"`
	Diskkbswrite      int    `json:"diskkbswrite,omitempty"`
	Servicestate      string `json:"servicestate,omitempty"`
	Publicipid        string `json:"publicipid,omitempty"`
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
	Created               string `json:"created,omitempty"`
	Networkkbswrite       int    `json:"networkkbswrite,omitempty"`
	Haenable              bool   `json:"haenable,omitempty"`
	Cpuspeed              int    `json:"cpuspeed,omitempty"`
	Hypervisor            string `json:"hypervisor,omitempty"`
	Isodisplaytext        string `json:"isodisplaytext,omitempty"`
	Templatedisplaytext   string `json:"templatedisplaytext,omitempty"`
	Groupid               string `json:"groupid,omitempty"`
	Isdynamicallyscalable bool   `json:"isdynamicallyscalable,omitempty"`
	State                 string `json:"state,omitempty"`
	Diskioread            int    `json:"diskioread,omitempty"`
	Serviceofferingname   string `json:"serviceofferingname,omitempty"`
	Networkkbsread        int    `json:"networkkbsread,omitempty"`
	Project               string `json:"project,omitempty"`
	Cpunumber             int    `json:"cpunumber,omitempty"`
	Account               string `json:"account,omitempty"`
	Cpuused               string `json:"cpuused,omitempty"`
	Password              string `json:"password,omitempty"`
	Serviceofferingid     string `json:"serviceofferingid,omitempty"`
	Displayname           string `json:"displayname,omitempty"`
	Diskiowrite           int    `json:"diskiowrite,omitempty"`
	Diskkbsread           int    `json:"diskkbsread,omitempty"`
	Hostid                string `json:"hostid,omitempty"`
	Publicip              string `json:"publicip,omitempty"`
	Nic                   []struct {
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Id           string   `json:"id,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Type         string   `json:"type,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
	} `json:"nic,omitempty"`
	Id           string `json:"id,omitempty"`
	Rootdeviceid int    `json:"rootdeviceid,omitempty"`
	Tags         []struct {
		Resourcetype string `json:"resourcetype,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Project      string `json:"project,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Account      string `json:"account,omitempty"`
		Key          string `json:"key,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Value        string `json:"value,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
	} `json:"tags,omitempty"`
	Zonename          string            `json:"zonename,omitempty"`
	Templatename      string            `json:"templatename,omitempty"`
	Isoname           string            `json:"isoname,omitempty"`
	Passwordenabled   bool              `json:"passwordenabled,omitempty"`
	Rootdevicetype    string            `json:"rootdevicetype,omitempty"`
	Memory            int               `json:"memory,omitempty"`
	Displayvm         bool              `json:"displayvm,omitempty"`
	Domainid          string            `json:"domainid,omitempty"`
	Guestosid         string            `json:"guestosid,omitempty"`
	Diskkbswrite      int               `json:"diskkbswrite,omitempty"`
	Forvirtualnetwork bool              `json:"forvirtualnetwork,omitempty"`
	Details           map[string]string `json:"details,omitempty"`
	Servicestate      string            `json:"servicestate,omitempty"`
	Hostname          string            `json:"hostname,omitempty"`
	Securitygroup     []struct {
		Project     string `json:"project,omitempty"`
		Domain      string `json:"domain,omitempty"`
		Id          string `json:"id,omitempty"`
		Description string `json:"description,omitempty"`
		Name        string `json:"name,omitempty"`
		Egressrule  []struct {
			Account           string `json:"account,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
		} `json:"egressrule,omitempty"`
		Account string `json:"account,omitempty"`
		Tags    []struct {
			Domain       string `json:"domain,omitempty"`
			Account      string `json:"account,omitempty"`
			Project      string `json:"project,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
			Value        string `json:"value,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Key          string `json:"key,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
		} `json:"tags,omitempty"`
		Ingressrule []struct {
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Account           string `json:"account,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
		} `json:"ingressrule,omitempty"`
		Domainid  string `json:"domainid,omitempty"`
		Projectid string `json:"projectid,omitempty"`
	} `json:"securitygroup,omitempty"`
	Keypair       string `json:"keypair,omitempty"`
	Name          string `json:"name,omitempty"`
	Domain        string `json:"domain,omitempty"`
	Publicipid    string `json:"publicipid,omitempty"`
	Affinitygroup []struct {
		Domainid          string   `json:"domainid,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		Name              string   `json:"name,omitempty"`
		Account           string   `json:"account,omitempty"`
		Description       string   `json:"description,omitempty"`
		Id                string   `json:"id,omitempty"`
		Type              string   `json:"type,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Isoid        string `json:"isoid,omitempty"`
	Templateid   string `json:"templateid,omitempty"`
	Group        string `json:"group,omitempty"`
	Zoneid       string `json:"zoneid,omitempty"`
	Projectid    string `json:"projectid,omitempty"`
	Instancename string `json:"instancename,omitempty"`
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
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
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
	Bootable bool   `json:"bootable,omitempty"`
	Id       string `json:"id,omitempty"`
	Tags     []struct {
		Key          string `json:"key,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Project      string `json:"project,omitempty"`
		Account      string `json:"account,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Value        string `json:"value,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
	} `json:"tags,omitempty"`
	Zoneid                string            `json:"zoneid,omitempty"`
	Isready               bool              `json:"isready,omitempty"`
	Hostname              string            `json:"hostname,omitempty"`
	Checksum              string            `json:"checksum,omitempty"`
	Isfeatured            bool              `json:"isfeatured,omitempty"`
	Domainid              string            `json:"domainid,omitempty"`
	Status                string            `json:"status,omitempty"`
	Hypervisor            string            `json:"hypervisor,omitempty"`
	Sshkeyenabled         bool              `json:"sshkeyenabled,omitempty"`
	Templatetag           string            `json:"templatetag,omitempty"`
	Format                string            `json:"format,omitempty"`
	Ostypename            string            `json:"ostypename,omitempty"`
	Account               string            `json:"account,omitempty"`
	Templatetype          string            `json:"templatetype,omitempty"`
	Zonename              string            `json:"zonename,omitempty"`
	Hostid                string            `json:"hostid,omitempty"`
	Project               string            `json:"project,omitempty"`
	Details               map[string]string `json:"details,omitempty"`
	Isextractable         bool              `json:"isextractable,omitempty"`
	Ostypeid              string            `json:"ostypeid,omitempty"`
	Displaytext           string            `json:"displaytext,omitempty"`
	Name                  string            `json:"name,omitempty"`
	Removed               string            `json:"removed,omitempty"`
	Created               string            `json:"created,omitempty"`
	Accountid             string            `json:"accountid,omitempty"`
	Passwordenabled       bool              `json:"passwordenabled,omitempty"`
	CrossZones            bool              `json:"crossZones,omitempty"`
	Domain                string            `json:"domain,omitempty"`
	Isdynamicallyscalable bool              `json:"isdynamicallyscalable,omitempty"`
	Ispublic              bool              `json:"ispublic,omitempty"`
	Size                  int               `json:"size,omitempty"`
	Projectid             string            `json:"projectid,omitempty"`
	Sourcetemplateid      string            `json:"sourcetemplateid,omitempty"`
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
	Templatetype          string            `json:"templatetype,omitempty"`
	Bootable              bool              `json:"bootable,omitempty"`
	Ispublic              bool              `json:"ispublic,omitempty"`
	CrossZones            bool              `json:"crossZones,omitempty"`
	Isready               bool              `json:"isready,omitempty"`
	Created               string            `json:"created,omitempty"`
	Projectid             string            `json:"projectid,omitempty"`
	Isdynamicallyscalable bool              `json:"isdynamicallyscalable,omitempty"`
	Size                  int               `json:"size,omitempty"`
	Domainid              string            `json:"domainid,omitempty"`
	Hostname              string            `json:"hostname,omitempty"`
	Domain                string            `json:"domain,omitempty"`
	Details               map[string]string `json:"details,omitempty"`
	Templatetag           string            `json:"templatetag,omitempty"`
	Passwordenabled       bool              `json:"passwordenabled,omitempty"`
	Checksum              string            `json:"checksum,omitempty"`
	Zoneid                string            `json:"zoneid,omitempty"`
	Status                string            `json:"status,omitempty"`
	Account               string            `json:"account,omitempty"`
	Removed               string            `json:"removed,omitempty"`
	Ostypename            string            `json:"ostypename,omitempty"`
	Id                    string            `json:"id,omitempty"`
	Sshkeyenabled         bool              `json:"sshkeyenabled,omitempty"`
	Project               string            `json:"project,omitempty"`
	Name                  string            `json:"name,omitempty"`
	Isfeatured            bool              `json:"isfeatured,omitempty"`
	Sourcetemplateid      string            `json:"sourcetemplateid,omitempty"`
	Ostypeid              string            `json:"ostypeid,omitempty"`
	Accountid             string            `json:"accountid,omitempty"`
	Tags                  []struct {
		Account      string `json:"account,omitempty"`
		Project      string `json:"project,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Value        string `json:"value,omitempty"`
		Key          string `json:"key,omitempty"`
	} `json:"tags,omitempty"`
	Isextractable bool   `json:"isextractable,omitempty"`
	Format        string `json:"format,omitempty"`
	Hypervisor    string `json:"hypervisor,omitempty"`
	Hostid        string `json:"hostid,omitempty"`
	Displaytext   string `json:"displaytext,omitempty"`
	Zonename      string `json:"zonename,omitempty"`
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
		vv := strconv.FormatInt(v.(int64), 10)
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
	Templatetag      string            `json:"templatetag,omitempty"`
	Ispublic         bool              `json:"ispublic,omitempty"`
	Domainid         string            `json:"domainid,omitempty"`
	Details          map[string]string `json:"details,omitempty"`
	Hostid           string            `json:"hostid,omitempty"`
	Created          string            `json:"created,omitempty"`
	Status           string            `json:"status,omitempty"`
	Checksum         string            `json:"checksum,omitempty"`
	Sourcetemplateid string            `json:"sourcetemplateid,omitempty"`
	Bootable         bool              `json:"bootable,omitempty"`
	Sshkeyenabled    bool              `json:"sshkeyenabled,omitempty"`
	Tags             []struct {
		Resourceid   string `json:"resourceid,omitempty"`
		Key          string `json:"key,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Project      string `json:"project,omitempty"`
		Account      string `json:"account,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Value        string `json:"value,omitempty"`
		Customer     string `json:"customer,omitempty"`
	} `json:"tags,omitempty"`
	Size                  int    `json:"size,omitempty"`
	Removed               string `json:"removed,omitempty"`
	Id                    string `json:"id,omitempty"`
	Accountid             string `json:"accountid,omitempty"`
	Ostypename            string `json:"ostypename,omitempty"`
	Passwordenabled       bool   `json:"passwordenabled,omitempty"`
	Name                  string `json:"name,omitempty"`
	Domain                string `json:"domain,omitempty"`
	Isdynamicallyscalable bool   `json:"isdynamicallyscalable,omitempty"`
	Projectid             string `json:"projectid,omitempty"`
	Templatetype          string `json:"templatetype,omitempty"`
	Isfeatured            bool   `json:"isfeatured,omitempty"`
	Format                string `json:"format,omitempty"`
	Zonename              string `json:"zonename,omitempty"`
	Hostname              string `json:"hostname,omitempty"`
	CrossZones            bool   `json:"crossZones,omitempty"`
	Isready               bool   `json:"isready,omitempty"`
	Hypervisor            string `json:"hypervisor,omitempty"`
	Zoneid                string `json:"zoneid,omitempty"`
	Project               string `json:"project,omitempty"`
	Ostypeid              string `json:"ostypeid,omitempty"`
	Isextractable         bool   `json:"isextractable,omitempty"`
	Account               string `json:"account,omitempty"`
	Displaytext           string `json:"displaytext,omitempty"`
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
	Displaytext           string `json:"displaytext,omitempty"`
	Status                string `json:"status,omitempty"`
	Isfeatured            bool   `json:"isfeatured,omitempty"`
	Zoneid                string `json:"zoneid,omitempty"`
	Id                    string `json:"id,omitempty"`
	Sourcetemplateid      string `json:"sourcetemplateid,omitempty"`
	Checksum              string `json:"checksum,omitempty"`
	Project               string `json:"project,omitempty"`
	Size                  int    `json:"size,omitempty"`
	Hypervisor            string `json:"hypervisor,omitempty"`
	Domainid              string `json:"domainid,omitempty"`
	Projectid             string `json:"projectid,omitempty"`
	Isdynamicallyscalable bool   `json:"isdynamicallyscalable,omitempty"`
	Ispublic              bool   `json:"ispublic,omitempty"`
	Zonename              string `json:"zonename,omitempty"`
	Templatetag           string `json:"templatetag,omitempty"`
	Sshkeyenabled         bool   `json:"sshkeyenabled,omitempty"`
	Name                  string `json:"name,omitempty"`
	Templatetype          string `json:"templatetype,omitempty"`
	Passwordenabled       bool   `json:"passwordenabled,omitempty"`
	Accountid             string `json:"accountid,omitempty"`
	CrossZones            bool   `json:"crossZones,omitempty"`
	Isextractable         bool   `json:"isextractable,omitempty"`
	Ostypename            string `json:"ostypename,omitempty"`
	Domain                string `json:"domain,omitempty"`
	Isready               bool   `json:"isready,omitempty"`
	Created               string `json:"created,omitempty"`
	Ostypeid              string `json:"ostypeid,omitempty"`
	Tags                  []struct {
		Project      string `json:"project,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Key          string `json:"key,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Value        string `json:"value,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Account      string `json:"account,omitempty"`
	} `json:"tags,omitempty"`
	Bootable bool              `json:"bootable,omitempty"`
	Hostid   string            `json:"hostid,omitempty"`
	Hostname string            `json:"hostname,omitempty"`
	Account  string            `json:"account,omitempty"`
	Details  map[string]string `json:"details,omitempty"`
	Format   string            `json:"format,omitempty"`
	Removed  string            `json:"removed,omitempty"`
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
	Success     bool   `json:"success,omitempty"`
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
	Projectids []string `json:"projectids,omitempty"`
	Domainid   string   `json:"domainid,omitempty"`
	Account    []string `json:"account,omitempty"`
	Id         string   `json:"id,omitempty"`
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
	Storagetype      string `json:"storagetype,omitempty"`
	Zonename         string `json:"zonename,omitempty"`
	Name             string `json:"name,omitempty"`
	Url              string `json:"url,omitempty"`
	ExtractMode      string `json:"extractMode,omitempty"`
	Resultstring     string `json:"resultstring,omitempty"`
	Status           string `json:"status,omitempty"`
	State            string `json:"state,omitempty"`
	Uploadpercentage int    `json:"uploadpercentage,omitempty"`
	Zoneid           string `json:"zoneid,omitempty"`
	ExtractId        string `json:"extractId,omitempty"`
	Id               string `json:"id,omitempty"`
	Accountid        string `json:"accountid,omitempty"`
	Created          string `json:"created,omitempty"`
}
