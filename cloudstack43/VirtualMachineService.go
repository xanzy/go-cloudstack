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

type DeployVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *DeployVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["affinitygroupids"]; found {
		vv := strings.Join(v.([]string), ", ")
		u.Set("affinitygroupids", vv)
	}
	if v, found := p.p["affinitygroupnames"]; found {
		vv := strings.Join(v.([]string), ", ")
		u.Set("affinitygroupnames", vv)
	}
	if v, found := p.p["details"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("details[%d].key", i), k)
			u.Set(fmt.Sprintf("details[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["diskofferingid"]; found {
		u.Set("diskofferingid", v.(string))
	}
	if v, found := p.p["displayname"]; found {
		u.Set("displayname", v.(string))
	}
	if v, found := p.p["displayvm"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("displayvm", vv)
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["group"]; found {
		u.Set("group", v.(string))
	}
	if v, found := p.p["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	if v, found := p.p["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := p.p["ip6address"]; found {
		u.Set("ip6address", v.(string))
	}
	if v, found := p.p["ipaddress"]; found {
		u.Set("ipaddress", v.(string))
	}
	if v, found := p.p["iptonetworklist"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("iptonetworklist[%d].key", i), k)
			u.Set(fmt.Sprintf("iptonetworklist[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["keyboard"]; found {
		u.Set("keyboard", v.(string))
	}
	if v, found := p.p["keypair"]; found {
		u.Set("keypair", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["networkids"]; found {
		vv := strings.Join(v.([]string), ", ")
		u.Set("networkids", vv)
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["securitygroupids"]; found {
		vv := strings.Join(v.([]string), ", ")
		u.Set("securitygroupids", vv)
	}
	if v, found := p.p["securitygroupnames"]; found {
		vv := strings.Join(v.([]string), ", ")
		u.Set("securitygroupnames", vv)
	}
	if v, found := p.p["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	if v, found := p.p["size"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("size", vv)
	}
	if v, found := p.p["startvm"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("startvm", vv)
	}
	if v, found := p.p["templateid"]; found {
		u.Set("templateid", v.(string))
	}
	if v, found := p.p["userdata"]; found {
		u.Set("userdata", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *DeployVirtualMachineParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *DeployVirtualMachineParams) SetAffinitygroupids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["affinitygroupids"] = v
	return
}

func (p *DeployVirtualMachineParams) SetAffinitygroupnames(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["affinitygroupnames"] = v
	return
}

func (p *DeployVirtualMachineParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
	return
}

func (p *DeployVirtualMachineParams) SetDiskofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["diskofferingid"] = v
	return
}

func (p *DeployVirtualMachineParams) SetDisplayname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displayname"] = v
	return
}

func (p *DeployVirtualMachineParams) SetDisplayvm(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displayvm"] = v
	return
}

func (p *DeployVirtualMachineParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *DeployVirtualMachineParams) SetGroup(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["group"] = v
	return
}

func (p *DeployVirtualMachineParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
	return
}

func (p *DeployVirtualMachineParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
	return
}

func (p *DeployVirtualMachineParams) SetIp6address(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ip6address"] = v
	return
}

func (p *DeployVirtualMachineParams) SetIpaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddress"] = v
	return
}

func (p *DeployVirtualMachineParams) SetIptonetworklist(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["iptonetworklist"] = v
	return
}

func (p *DeployVirtualMachineParams) SetKeyboard(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyboard"] = v
	return
}

func (p *DeployVirtualMachineParams) SetKeypair(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keypair"] = v
	return
}

func (p *DeployVirtualMachineParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *DeployVirtualMachineParams) SetNetworkids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkids"] = v
	return
}

func (p *DeployVirtualMachineParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *DeployVirtualMachineParams) SetSecuritygroupids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["securitygroupids"] = v
	return
}

func (p *DeployVirtualMachineParams) SetSecuritygroupnames(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["securitygroupnames"] = v
	return
}

func (p *DeployVirtualMachineParams) SetServiceofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingid"] = v
	return
}

func (p *DeployVirtualMachineParams) SetSize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["size"] = v
	return
}

func (p *DeployVirtualMachineParams) SetStartvm(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startvm"] = v
	return
}

func (p *DeployVirtualMachineParams) SetTemplateid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templateid"] = v
	return
}

func (p *DeployVirtualMachineParams) SetUserdata(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userdata"] = v
	return
}

func (p *DeployVirtualMachineParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new DeployVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewDeployVirtualMachineParams(serviceofferingid string, templateid string, zoneid string) *DeployVirtualMachineParams {
	p := &DeployVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["serviceofferingid"] = serviceofferingid
	p.p["templateid"] = templateid
	p.p["zoneid"] = zoneid
	return p
}

// Creates and automatically starts a virtual machine based on a service offering, disk offering, and template.
func (s *VirtualMachineService) DeployVirtualMachine(p *DeployVirtualMachineParams) (*DeployVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("deployVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeployVirtualMachineResponse
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

type DeployVirtualMachineResponse struct {
	JobID  string `json:"jobid,omitempty"`
	Hostid string `json:"hostid,omitempty"`
	Nic    []struct {
		Isolationuri string   `json:"isolationuri,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Type         string   `json:"type,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Id           string   `json:"id,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
	} `json:"nic,omitempty"`
	Haenable              bool   `json:"haenable,omitempty"`
	Cpuused               string `json:"cpuused,omitempty"`
	Displayname           string `json:"displayname,omitempty"`
	Networkkbsread        int    `json:"networkkbsread,omitempty"`
	Diskkbsread           int    `json:"diskkbsread,omitempty"`
	Diskiowrite           int    `json:"diskiowrite,omitempty"`
	Rootdeviceid          int    `json:"rootdeviceid,omitempty"`
	Isoid                 string `json:"isoid,omitempty"`
	Hypervisor            string `json:"hypervisor,omitempty"`
	Domainid              string `json:"domainid,omitempty"`
	Memory                int    `json:"memory,omitempty"`
	Servicestate          string `json:"servicestate,omitempty"`
	Zoneid                string `json:"zoneid,omitempty"`
	Passwordenabled       bool   `json:"passwordenabled,omitempty"`
	Hostname              string `json:"hostname,omitempty"`
	State                 string `json:"state,omitempty"`
	Group                 string `json:"group,omitempty"`
	Templatedisplaytext   string `json:"templatedisplaytext,omitempty"`
	Isdynamicallyscalable bool   `json:"isdynamicallyscalable,omitempty"`
	Isodisplaytext        string `json:"isodisplaytext,omitempty"`
	Affinitygroup         []struct {
		Name              string   `json:"name,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Type              string   `json:"type,omitempty"`
		Description       string   `json:"description,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		Id                string   `json:"id,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
		Account           string   `json:"account,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Templatename      string            `json:"templatename,omitempty"`
	Account           string            `json:"account,omitempty"`
	Password          string            `json:"password,omitempty"`
	Serviceofferingid string            `json:"serviceofferingid,omitempty"`
	Created           string            `json:"created,omitempty"`
	Name              string            `json:"name,omitempty"`
	Details           map[string]string `json:"details,omitempty"`
	Isoname           string            `json:"isoname,omitempty"`
	Id                string            `json:"id,omitempty"`
	Cpuspeed          int               `json:"cpuspeed,omitempty"`
	Securitygroup     []struct {
		Id        string `json:"id,omitempty"`
		Name      string `json:"name,omitempty"`
		Projectid string `json:"projectid,omitempty"`
		Tags      []struct {
			Projectid    string `json:"projectid,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Key          string `json:"key,omitempty"`
			Account      string `json:"account,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Value        string `json:"value,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Project      string `json:"project,omitempty"`
		} `json:"tags,omitempty"`
		Project    string `json:"project,omitempty"`
		Egressrule []struct {
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Account           string `json:"account,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Endport           int    `json:"endport,omitempty"`
		} `json:"egressrule,omitempty"`
		Ingressrule []struct {
			Endport           int    `json:"endport,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Account           string `json:"account,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
		} `json:"ingressrule,omitempty"`
		Domain      string `json:"domain,omitempty"`
		Description string `json:"description,omitempty"`
		Account     string `json:"account,omitempty"`
		Domainid    string `json:"domainid,omitempty"`
	} `json:"securitygroup,omitempty"`
	Diskkbswrite      int    `json:"diskkbswrite,omitempty"`
	Networkkbswrite   int    `json:"networkkbswrite,omitempty"`
	Displayvm         bool   `json:"displayvm,omitempty"`
	Keypair           string `json:"keypair,omitempty"`
	Forvirtualnetwork bool   `json:"forvirtualnetwork,omitempty"`
	Publicipid        string `json:"publicipid,omitempty"`
	Domain            string `json:"domain,omitempty"`
	Tags              []struct {
		Projectid    string `json:"projectid,omitempty"`
		Account      string `json:"account,omitempty"`
		Value        string `json:"value,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Project      string `json:"project,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Key          string `json:"key,omitempty"`
	} `json:"tags,omitempty"`
	Diskioread          int    `json:"diskioread,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Guestosid           string `json:"guestosid,omitempty"`
	Project             string `json:"project,omitempty"`
	Publicip            string `json:"publicip,omitempty"`
	Templateid          string `json:"templateid,omitempty"`
	Rootdevicetype      string `json:"rootdevicetype,omitempty"`
	Cpunumber           int    `json:"cpunumber,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
	Projectid           string `json:"projectid,omitempty"`
	Groupid             string `json:"groupid,omitempty"`
	Instancename        string `json:"instancename,omitempty"`
}

type DestroyVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *DestroyVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["expunge"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("expunge", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DestroyVirtualMachineParams) SetExpunge(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["expunge"] = v
	return
}

func (p *DestroyVirtualMachineParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DestroyVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewDestroyVirtualMachineParams(id string) *DestroyVirtualMachineParams {
	p := &DestroyVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Destroys a virtual machine. Once destroyed, only the administrator can recover it.
func (s *VirtualMachineService) DestroyVirtualMachine(p *DestroyVirtualMachineParams) (*DestroyVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("destroyVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DestroyVirtualMachineResponse
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

type DestroyVirtualMachineResponse struct {
	JobID        string `json:"jobid,omitempty"`
	Group        string `json:"group,omitempty"`
	Project      string `json:"project,omitempty"`
	Rootdeviceid int    `json:"rootdeviceid,omitempty"`
	Hostid       string `json:"hostid,omitempty"`
	Created      string `json:"created,omitempty"`
	Zoneid       string `json:"zoneid,omitempty"`
	Nic          []struct {
		Type         string   `json:"type,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Id           string   `json:"id,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
	} `json:"nic,omitempty"`
	Diskiowrite           int    `json:"diskiowrite,omitempty"`
	Isdynamicallyscalable bool   `json:"isdynamicallyscalable,omitempty"`
	Templateid            string `json:"templateid,omitempty"`
	Domain                string `json:"domain,omitempty"`
	Displayname           string `json:"displayname,omitempty"`
	Templatedisplaytext   string `json:"templatedisplaytext,omitempty"`
	Memory                int    `json:"memory,omitempty"`
	Securitygroup         []struct {
		Description string `json:"description,omitempty"`
		Id          string `json:"id,omitempty"`
		Domain      string `json:"domain,omitempty"`
		Egressrule  []struct {
			Endport           int    `json:"endport,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Account           string `json:"account,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
		} `json:"egressrule,omitempty"`
		Name        string `json:"name,omitempty"`
		Account     string `json:"account,omitempty"`
		Project     string `json:"project,omitempty"`
		Domainid    string `json:"domainid,omitempty"`
		Projectid   string `json:"projectid,omitempty"`
		Ingressrule []struct {
			Icmptype          int    `json:"icmptype,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Account           string `json:"account,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Endport           int    `json:"endport,omitempty"`
		} `json:"ingressrule,omitempty"`
		Tags []struct {
			Account      string `json:"account,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Value        string `json:"value,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Project      string `json:"project,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Key          string `json:"key,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
		} `json:"tags,omitempty"`
	} `json:"securitygroup,omitempty"`
	Name      string `json:"name,omitempty"`
	Projectid string `json:"projectid,omitempty"`
	Tags      []struct {
		Customer     string `json:"customer,omitempty"`
		Value        string `json:"value,omitempty"`
		Project      string `json:"project,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Key          string `json:"key,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Account      string `json:"account,omitempty"`
	} `json:"tags,omitempty"`
	Guestosid           string            `json:"guestosid,omitempty"`
	Hostname            string            `json:"hostname,omitempty"`
	Rootdevicetype      string            `json:"rootdevicetype,omitempty"`
	Zonename            string            `json:"zonename,omitempty"`
	Id                  string            `json:"id,omitempty"`
	Instancename        string            `json:"instancename,omitempty"`
	Isodisplaytext      string            `json:"isodisplaytext,omitempty"`
	Isoname             string            `json:"isoname,omitempty"`
	Diskkbsread         int               `json:"diskkbsread,omitempty"`
	Isoid               string            `json:"isoid,omitempty"`
	Diskkbswrite        int               `json:"diskkbswrite,omitempty"`
	State               string            `json:"state,omitempty"`
	Passwordenabled     bool              `json:"passwordenabled,omitempty"`
	Servicestate        string            `json:"servicestate,omitempty"`
	Cpunumber           int               `json:"cpunumber,omitempty"`
	Details             map[string]string `json:"details,omitempty"`
	Groupid             string            `json:"groupid,omitempty"`
	Serviceofferingname string            `json:"serviceofferingname,omitempty"`
	Cpuspeed            int               `json:"cpuspeed,omitempty"`
	Domainid            string            `json:"domainid,omitempty"`
	Haenable            bool              `json:"haenable,omitempty"`
	Account             string            `json:"account,omitempty"`
	Diskioread          int               `json:"diskioread,omitempty"`
	Hypervisor          string            `json:"hypervisor,omitempty"`
	Publicipid          string            `json:"publicipid,omitempty"`
	Publicip            string            `json:"publicip,omitempty"`
	Cpuused             string            `json:"cpuused,omitempty"`
	Networkkbswrite     int               `json:"networkkbswrite,omitempty"`
	Serviceofferingid   string            `json:"serviceofferingid,omitempty"`
	Displayvm           bool              `json:"displayvm,omitempty"`
	Affinitygroup       []struct {
		Domain            string   `json:"domain,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
		Type              string   `json:"type,omitempty"`
		Id                string   `json:"id,omitempty"`
		Description       string   `json:"description,omitempty"`
		Account           string   `json:"account,omitempty"`
		Name              string   `json:"name,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Templatename      string `json:"templatename,omitempty"`
	Networkkbsread    int    `json:"networkkbsread,omitempty"`
	Password          string `json:"password,omitempty"`
	Keypair           string `json:"keypair,omitempty"`
	Forvirtualnetwork bool   `json:"forvirtualnetwork,omitempty"`
}

type RebootVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *RebootVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *RebootVirtualMachineParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new RebootVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewRebootVirtualMachineParams(id string) *RebootVirtualMachineParams {
	p := &RebootVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Reboots a virtual machine.
func (s *VirtualMachineService) RebootVirtualMachine(p *RebootVirtualMachineParams) (*RebootVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("rebootVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RebootVirtualMachineResponse
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

type RebootVirtualMachineResponse struct {
	JobID                 string            `json:"jobid,omitempty"`
	Details               map[string]string `json:"details,omitempty"`
	Project               string            `json:"project,omitempty"`
	Zonename              string            `json:"zonename,omitempty"`
	Displayname           string            `json:"displayname,omitempty"`
	Isoname               string            `json:"isoname,omitempty"`
	Networkkbswrite       int               `json:"networkkbswrite,omitempty"`
	Hostid                string            `json:"hostid,omitempty"`
	Forvirtualnetwork     bool              `json:"forvirtualnetwork,omitempty"`
	Created               string            `json:"created,omitempty"`
	Id                    string            `json:"id,omitempty"`
	Hostname              string            `json:"hostname,omitempty"`
	State                 string            `json:"state,omitempty"`
	Cpunumber             int               `json:"cpunumber,omitempty"`
	Zoneid                string            `json:"zoneid,omitempty"`
	Domain                string            `json:"domain,omitempty"`
	Passwordenabled       bool              `json:"passwordenabled,omitempty"`
	Guestosid             string            `json:"guestosid,omitempty"`
	Groupid               string            `json:"groupid,omitempty"`
	Displayvm             bool              `json:"displayvm,omitempty"`
	Templateid            string            `json:"templateid,omitempty"`
	Templatedisplaytext   string            `json:"templatedisplaytext,omitempty"`
	Cpuspeed              int               `json:"cpuspeed,omitempty"`
	Account               string            `json:"account,omitempty"`
	Networkkbsread        int               `json:"networkkbsread,omitempty"`
	Cpuused               string            `json:"cpuused,omitempty"`
	Isdynamicallyscalable bool              `json:"isdynamicallyscalable,omitempty"`
	Hypervisor            string            `json:"hypervisor,omitempty"`
	Isoid                 string            `json:"isoid,omitempty"`
	Servicestate          string            `json:"servicestate,omitempty"`
	Haenable              bool              `json:"haenable,omitempty"`
	Diskkbswrite          int               `json:"diskkbswrite,omitempty"`
	Isodisplaytext        string            `json:"isodisplaytext,omitempty"`
	Diskioread            int               `json:"diskioread,omitempty"`
	Memory                int               `json:"memory,omitempty"`
	Domainid              string            `json:"domainid,omitempty"`
	Rootdevicetype        string            `json:"rootdevicetype,omitempty"`
	Name                  string            `json:"name,omitempty"`
	Projectid             string            `json:"projectid,omitempty"`
	Publicipid            string            `json:"publicipid,omitempty"`
	Templatename          string            `json:"templatename,omitempty"`
	Serviceofferingid     string            `json:"serviceofferingid,omitempty"`
	Keypair               string            `json:"keypair,omitempty"`
	Serviceofferingname   string            `json:"serviceofferingname,omitempty"`
	Diskkbsread           int               `json:"diskkbsread,omitempty"`
	Nic                   []struct {
		Networkname  string   `json:"networkname,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Id           string   `json:"id,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Type         string   `json:"type,omitempty"`
	} `json:"nic,omitempty"`
	Affinitygroup []struct {
		Domainid          string   `json:"domainid,omitempty"`
		Account           string   `json:"account,omitempty"`
		Description       string   `json:"description,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Type              string   `json:"type,omitempty"`
		Id                string   `json:"id,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		Name              string   `json:"name,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Group         string `json:"group,omitempty"`
	Securitygroup []struct {
		Domainid    string `json:"domainid,omitempty"`
		Account     string `json:"account,omitempty"`
		Description string `json:"description,omitempty"`
		Tags        []struct {
			Value        string `json:"value,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Project      string `json:"project,omitempty"`
			Account      string `json:"account,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Key          string `json:"key,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
		} `json:"tags,omitempty"`
		Projectid  string `json:"projectid,omitempty"`
		Egressrule []struct {
			Cidr              string `json:"cidr,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Account           string `json:"account,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Endport           int    `json:"endport,omitempty"`
		} `json:"egressrule,omitempty"`
		Name        string `json:"name,omitempty"`
		Ingressrule []struct {
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Account           string `json:"account,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
		} `json:"ingressrule,omitempty"`
		Domain  string `json:"domain,omitempty"`
		Project string `json:"project,omitempty"`
		Id      string `json:"id,omitempty"`
	} `json:"securitygroup,omitempty"`
	Password     string `json:"password,omitempty"`
	Rootdeviceid int    `json:"rootdeviceid,omitempty"`
	Publicip     string `json:"publicip,omitempty"`
	Diskiowrite  int    `json:"diskiowrite,omitempty"`
	Tags         []struct {
		Resourceid   string `json:"resourceid,omitempty"`
		Project      string `json:"project,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Account      string `json:"account,omitempty"`
		Key          string `json:"key,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Value        string `json:"value,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Customer     string `json:"customer,omitempty"`
	} `json:"tags,omitempty"`
	Instancename string `json:"instancename,omitempty"`
}

type StartVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *StartVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *StartVirtualMachineParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
	return
}

func (p *StartVirtualMachineParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new StartVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewStartVirtualMachineParams(id string) *StartVirtualMachineParams {
	p := &StartVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Starts a virtual machine.
func (s *VirtualMachineService) StartVirtualMachine(p *StartVirtualMachineParams) (*StartVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("startVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r StartVirtualMachineResponse
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

type StartVirtualMachineResponse struct {
	JobID    string `json:"jobid,omitempty"`
	Cpuspeed int    `json:"cpuspeed,omitempty"`
	Nic      []struct {
		Traffictype  string   `json:"traffictype,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Type         string   `json:"type,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Id           string   `json:"id,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
	} `json:"nic,omitempty"`
	Forvirtualnetwork     bool   `json:"forvirtualnetwork,omitempty"`
	Templatedisplaytext   string `json:"templatedisplaytext,omitempty"`
	Isdynamicallyscalable bool   `json:"isdynamicallyscalable,omitempty"`
	Templatename          string `json:"templatename,omitempty"`
	Displayvm             bool   `json:"displayvm,omitempty"`
	Isoname               string `json:"isoname,omitempty"`
	Name                  string `json:"name,omitempty"`
	Diskkbsread           int    `json:"diskkbsread,omitempty"`
	Hostname              string `json:"hostname,omitempty"`
	Account               string `json:"account,omitempty"`
	Memory                int    `json:"memory,omitempty"`
	Groupid               string `json:"groupid,omitempty"`
	Password              string `json:"password,omitempty"`
	Domain                string `json:"domain,omitempty"`
	Publicip              string `json:"publicip,omitempty"`
	Servicestate          string `json:"servicestate,omitempty"`
	Haenable              bool   `json:"haenable,omitempty"`
	Created               string `json:"created,omitempty"`
	State                 string `json:"state,omitempty"`
	Diskioread            int    `json:"diskioread,omitempty"`
	Cpuused               string `json:"cpuused,omitempty"`
	Cpunumber             int    `json:"cpunumber,omitempty"`
	Diskiowrite           int    `json:"diskiowrite,omitempty"`
	Domainid              string `json:"domainid,omitempty"`
	Projectid             string `json:"projectid,omitempty"`
	Serviceofferingname   string `json:"serviceofferingname,omitempty"`
	Displayname           string `json:"displayname,omitempty"`
	Passwordenabled       bool   `json:"passwordenabled,omitempty"`
	Hostid                string `json:"hostid,omitempty"`
	Zonename              string `json:"zonename,omitempty"`
	Networkkbsread        int    `json:"networkkbsread,omitempty"`
	Publicipid            string `json:"publicipid,omitempty"`
	Serviceofferingid     string `json:"serviceofferingid,omitempty"`
	Project               string `json:"project,omitempty"`
	Rootdevicetype        string `json:"rootdevicetype,omitempty"`
	Instancename          string `json:"instancename,omitempty"`
	Group                 string `json:"group,omitempty"`
	Zoneid                string `json:"zoneid,omitempty"`
	Hypervisor            string `json:"hypervisor,omitempty"`
	Isoid                 string `json:"isoid,omitempty"`
	Isodisplaytext        string `json:"isodisplaytext,omitempty"`
	Diskkbswrite          int    `json:"diskkbswrite,omitempty"`
	Tags                  []struct {
		Value        string `json:"value,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Account      string `json:"account,omitempty"`
		Project      string `json:"project,omitempty"`
		Key          string `json:"key,omitempty"`
	} `json:"tags,omitempty"`
	Networkkbswrite int `json:"networkkbswrite,omitempty"`
	Affinitygroup   []struct {
		Account           string   `json:"account,omitempty"`
		Description       string   `json:"description,omitempty"`
		Type              string   `json:"type,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
		Id                string   `json:"id,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Name              string   `json:"name,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Keypair       string `json:"keypair,omitempty"`
	Securitygroup []struct {
		Ingressrule []struct {
			Account           string `json:"account,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Startport         int    `json:"startport,omitempty"`
		} `json:"ingressrule,omitempty"`
		Domainid string `json:"domainid,omitempty"`
		Project  string `json:"project,omitempty"`
		Tags     []struct {
			Domain       string `json:"domain,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Key          string `json:"key,omitempty"`
			Project      string `json:"project,omitempty"`
			Account      string `json:"account,omitempty"`
			Value        string `json:"value,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
		} `json:"tags,omitempty"`
		Description string `json:"description,omitempty"`
		Name        string `json:"name,omitempty"`
		Egressrule  []struct {
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Account           string `json:"account,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
		} `json:"egressrule,omitempty"`
		Account   string `json:"account,omitempty"`
		Projectid string `json:"projectid,omitempty"`
		Id        string `json:"id,omitempty"`
		Domain    string `json:"domain,omitempty"`
	} `json:"securitygroup,omitempty"`
	Id           string            `json:"id,omitempty"`
	Guestosid    string            `json:"guestosid,omitempty"`
	Rootdeviceid int               `json:"rootdeviceid,omitempty"`
	Details      map[string]string `json:"details,omitempty"`
	Templateid   string            `json:"templateid,omitempty"`
}

type StopVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *StopVirtualMachineParams) toURLValues() url.Values {
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

func (p *StopVirtualMachineParams) SetForced(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forced"] = v
	return
}

func (p *StopVirtualMachineParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new StopVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewStopVirtualMachineParams(id string) *StopVirtualMachineParams {
	p := &StopVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Stops a virtual machine.
func (s *VirtualMachineService) StopVirtualMachine(p *StopVirtualMachineParams) (*StopVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("stopVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r StopVirtualMachineResponse
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

type StopVirtualMachineResponse struct {
	JobID         string `json:"jobid,omitempty"`
	Projectid     string `json:"projectid,omitempty"`
	Diskioread    int    `json:"diskioread,omitempty"`
	Displayname   string `json:"displayname,omitempty"`
	Hostname      string `json:"hostname,omitempty"`
	Groupid       string `json:"groupid,omitempty"`
	Isoname       string `json:"isoname,omitempty"`
	Rootdeviceid  int    `json:"rootdeviceid,omitempty"`
	Id            string `json:"id,omitempty"`
	Affinitygroup []struct {
		Id                string   `json:"id,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Description       string   `json:"description,omitempty"`
		Account           string   `json:"account,omitempty"`
		Type              string   `json:"type,omitempty"`
		Name              string   `json:"name,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Isoid        string            `json:"isoid,omitempty"`
	Templateid   string            `json:"templateid,omitempty"`
	Diskiowrite  int               `json:"diskiowrite,omitempty"`
	Details      map[string]string `json:"details,omitempty"`
	Servicestate string            `json:"servicestate,omitempty"`
	Password     string            `json:"password,omitempty"`
	Nic          []struct {
		Networkname  string   `json:"networkname,omitempty"`
		Type         string   `json:"type,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Id           string   `json:"id,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
	} `json:"nic,omitempty"`
	Publicipid        string `json:"publicipid,omitempty"`
	Memory            int    `json:"memory,omitempty"`
	Displayvm         bool   `json:"displayvm,omitempty"`
	Serviceofferingid string `json:"serviceofferingid,omitempty"`
	Tags              []struct {
		Account      string `json:"account,omitempty"`
		Key          string `json:"key,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Value        string `json:"value,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Project      string `json:"project,omitempty"`
	} `json:"tags,omitempty"`
	Publicip      string `json:"publicip,omitempty"`
	Domainid      string `json:"domainid,omitempty"`
	Created       string `json:"created,omitempty"`
	Securitygroup []struct {
		Ingressrule []struct {
			Account           string `json:"account,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
		} `json:"ingressrule,omitempty"`
		Domainid  string `json:"domainid,omitempty"`
		Projectid string `json:"projectid,omitempty"`
		Domain    string `json:"domain,omitempty"`
		Name      string `json:"name,omitempty"`
		Id        string `json:"id,omitempty"`
		Tags      []struct {
			Domain       string `json:"domain,omitempty"`
			Project      string `json:"project,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Value        string `json:"value,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Key          string `json:"key,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
			Account      string `json:"account,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
		} `json:"tags,omitempty"`
		Project    string `json:"project,omitempty"`
		Egressrule []struct {
			Protocol          string `json:"protocol,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Account           string `json:"account,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Endport           int    `json:"endport,omitempty"`
		} `json:"egressrule,omitempty"`
		Description string `json:"description,omitempty"`
		Account     string `json:"account,omitempty"`
	} `json:"securitygroup,omitempty"`
	Account               string `json:"account,omitempty"`
	Networkkbsread        int    `json:"networkkbsread,omitempty"`
	Networkkbswrite       int    `json:"networkkbswrite,omitempty"`
	Instancename          string `json:"instancename,omitempty"`
	Hostid                string `json:"hostid,omitempty"`
	Name                  string `json:"name,omitempty"`
	Group                 string `json:"group,omitempty"`
	Serviceofferingname   string `json:"serviceofferingname,omitempty"`
	Rootdevicetype        string `json:"rootdevicetype,omitempty"`
	Project               string `json:"project,omitempty"`
	Haenable              bool   `json:"haenable,omitempty"`
	Diskkbswrite          int    `json:"diskkbswrite,omitempty"`
	Keypair               string `json:"keypair,omitempty"`
	Zonename              string `json:"zonename,omitempty"`
	Zoneid                string `json:"zoneid,omitempty"`
	Isdynamicallyscalable bool   `json:"isdynamicallyscalable,omitempty"`
	Hypervisor            string `json:"hypervisor,omitempty"`
	Guestosid             string `json:"guestosid,omitempty"`
	Isodisplaytext        string `json:"isodisplaytext,omitempty"`
	Passwordenabled       bool   `json:"passwordenabled,omitempty"`
	Cpuspeed              int    `json:"cpuspeed,omitempty"`
	Cpuused               string `json:"cpuused,omitempty"`
	State                 string `json:"state,omitempty"`
	Diskkbsread           int    `json:"diskkbsread,omitempty"`
	Cpunumber             int    `json:"cpunumber,omitempty"`
	Forvirtualnetwork     bool   `json:"forvirtualnetwork,omitempty"`
	Templatename          string `json:"templatename,omitempty"`
	Templatedisplaytext   string `json:"templatedisplaytext,omitempty"`
	Domain                string `json:"domain,omitempty"`
}

type ResetPasswordForVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *ResetPasswordForVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *ResetPasswordForVirtualMachineParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new ResetPasswordForVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewResetPasswordForVirtualMachineParams(id string) *ResetPasswordForVirtualMachineParams {
	p := &ResetPasswordForVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Resets the password for virtual machine. The virtual machine must be in a "Stopped" state and the template must already support this feature for this command to take effect. [async]
func (s *VirtualMachineService) ResetPasswordForVirtualMachine(p *ResetPasswordForVirtualMachineParams) (*ResetPasswordForVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("resetPasswordForVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ResetPasswordForVirtualMachineResponse
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

type ResetPasswordForVirtualMachineResponse struct {
	JobID               string `json:"jobid,omitempty"`
	State               string `json:"state,omitempty"`
	Networkkbswrite     int    `json:"networkkbswrite,omitempty"`
	Isoid               string `json:"isoid,omitempty"`
	Templatedisplaytext string `json:"templatedisplaytext,omitempty"`
	Hostname            string `json:"hostname,omitempty"`
	Tags                []struct {
		Customer     string `json:"customer,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Value        string `json:"value,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Key          string `json:"key,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Account      string `json:"account,omitempty"`
		Project      string `json:"project,omitempty"`
	} `json:"tags,omitempty"`
	Servicestate          string `json:"servicestate,omitempty"`
	Serviceofferingname   string `json:"serviceofferingname,omitempty"`
	Account               string `json:"account,omitempty"`
	Keypair               string `json:"keypair,omitempty"`
	Isdynamicallyscalable bool   `json:"isdynamicallyscalable,omitempty"`
	Isodisplaytext        string `json:"isodisplaytext,omitempty"`
	Templatename          string `json:"templatename,omitempty"`
	Groupid               string `json:"groupid,omitempty"`
	Cpuspeed              int    `json:"cpuspeed,omitempty"`
	Memory                int    `json:"memory,omitempty"`
	Guestosid             string `json:"guestosid,omitempty"`
	Instancename          string `json:"instancename,omitempty"`
	Created               string `json:"created,omitempty"`
	Domainid              string `json:"domainid,omitempty"`
	Name                  string `json:"name,omitempty"`
	Haenable              bool   `json:"haenable,omitempty"`
	Publicipid            string `json:"publicipid,omitempty"`
	Nic                   []struct {
		Gateway      string   `json:"gateway,omitempty"`
		Type         string   `json:"type,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Id           string   `json:"id,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
	} `json:"nic,omitempty"`
	Cpunumber         int    `json:"cpunumber,omitempty"`
	Diskiowrite       int    `json:"diskiowrite,omitempty"`
	Displayvm         bool   `json:"displayvm,omitempty"`
	Project           string `json:"project,omitempty"`
	Diskkbsread       int    `json:"diskkbsread,omitempty"`
	Passwordenabled   bool   `json:"passwordenabled,omitempty"`
	Publicip          string `json:"publicip,omitempty"`
	Displayname       string `json:"displayname,omitempty"`
	Isoname           string `json:"isoname,omitempty"`
	Hypervisor        string `json:"hypervisor,omitempty"`
	Templateid        string `json:"templateid,omitempty"`
	Forvirtualnetwork bool   `json:"forvirtualnetwork,omitempty"`
	Rootdeviceid      int    `json:"rootdeviceid,omitempty"`
	Affinitygroup     []struct {
		Description       string   `json:"description,omitempty"`
		Id                string   `json:"id,omitempty"`
		Account           string   `json:"account,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Name              string   `json:"name,omitempty"`
		Type              string   `json:"type,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Serviceofferingid string `json:"serviceofferingid,omitempty"`
	Networkkbsread    int    `json:"networkkbsread,omitempty"`
	Group             string `json:"group,omitempty"`
	Zonename          string `json:"zonename,omitempty"`
	Projectid         string `json:"projectid,omitempty"`
	Password          string `json:"password,omitempty"`
	Domain            string `json:"domain,omitempty"`
	Cpuused           string `json:"cpuused,omitempty"`
	Zoneid            string `json:"zoneid,omitempty"`
	Rootdevicetype    string `json:"rootdevicetype,omitempty"`
	Diskioread        int    `json:"diskioread,omitempty"`
	Diskkbswrite      int    `json:"diskkbswrite,omitempty"`
	Securitygroup     []struct {
		Domainid    string `json:"domainid,omitempty"`
		Id          string `json:"id,omitempty"`
		Domain      string `json:"domain,omitempty"`
		Ingressrule []struct {
			Startport         int    `json:"startport,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Account           string `json:"account,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
		} `json:"ingressrule,omitempty"`
		Name string `json:"name,omitempty"`
		Tags []struct {
			Domain       string `json:"domain,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Key          string `json:"key,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
			Project      string `json:"project,omitempty"`
			Account      string `json:"account,omitempty"`
			Value        string `json:"value,omitempty"`
		} `json:"tags,omitempty"`
		Description string `json:"description,omitempty"`
		Projectid   string `json:"projectid,omitempty"`
		Egressrule  []struct {
			Cidr              string `json:"cidr,omitempty"`
			Account           string `json:"account,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
		} `json:"egressrule,omitempty"`
		Project string `json:"project,omitempty"`
		Account string `json:"account,omitempty"`
	} `json:"securitygroup,omitempty"`
	Details map[string]string `json:"details,omitempty"`
	Id      string            `json:"id,omitempty"`
	Hostid  string            `json:"hostid,omitempty"`
}

type UpdateVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *UpdateVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["displayname"]; found {
		u.Set("displayname", v.(string))
	}
	if v, found := p.p["displayvm"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("displayvm", vv)
	}
	if v, found := p.p["group"]; found {
		u.Set("group", v.(string))
	}
	if v, found := p.p["haenable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("haenable", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["isdynamicallyscalable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isdynamicallyscalable", vv)
	}
	if v, found := p.p["ostypeid"]; found {
		u.Set("ostypeid", v.(string))
	}
	if v, found := p.p["userdata"]; found {
		u.Set("userdata", v.(string))
	}
	return u
}

func (p *UpdateVirtualMachineParams) SetDisplayname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displayname"] = v
	return
}

func (p *UpdateVirtualMachineParams) SetDisplayvm(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displayvm"] = v
	return
}

func (p *UpdateVirtualMachineParams) SetGroup(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["group"] = v
	return
}

func (p *UpdateVirtualMachineParams) SetHaenable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["haenable"] = v
	return
}

func (p *UpdateVirtualMachineParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *UpdateVirtualMachineParams) SetIsdynamicallyscalable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isdynamicallyscalable"] = v
	return
}

func (p *UpdateVirtualMachineParams) SetOstypeid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ostypeid"] = v
	return
}

func (p *UpdateVirtualMachineParams) SetUserdata(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userdata"] = v
	return
}

// You should always use this function to get a new UpdateVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewUpdateVirtualMachineParams(id string) *UpdateVirtualMachineParams {
	p := &UpdateVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates properties of a virtual machine. The VM has to be stopped and restarted for the new properties to take effect. UpdateVirtualMachine does not first check whether the VM is stopped. Therefore, stop the VM manually before issuing this call.
func (s *VirtualMachineService) UpdateVirtualMachine(p *UpdateVirtualMachineParams) (*UpdateVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("updateVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateVirtualMachineResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type UpdateVirtualMachineResponse struct {
	Hypervisor    string `json:"hypervisor,omitempty"`
	Cpuused       string `json:"cpuused,omitempty"`
	Securitygroup []struct {
		Tags []struct {
			Key          string `json:"key,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Project      string `json:"project,omitempty"`
			Value        string `json:"value,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
			Account      string `json:"account,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
		} `json:"tags,omitempty"`
		Projectid   string `json:"projectid,omitempty"`
		Account     string `json:"account,omitempty"`
		Description string `json:"description,omitempty"`
		Egressrule  []struct {
			Cidr              string `json:"cidr,omitempty"`
			Account           string `json:"account,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
		} `json:"egressrule,omitempty"`
		Domain      string `json:"domain,omitempty"`
		Id          string `json:"id,omitempty"`
		Domainid    string `json:"domainid,omitempty"`
		Project     string `json:"project,omitempty"`
		Ingressrule []struct {
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Account           string `json:"account,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
		} `json:"ingressrule,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"securitygroup,omitempty"`
	Hostid            string `json:"hostid,omitempty"`
	Forvirtualnetwork bool   `json:"forvirtualnetwork,omitempty"`
	Diskkbsread       int    `json:"diskkbsread,omitempty"`
	Account           string `json:"account,omitempty"`
	Domain            string `json:"domain,omitempty"`
	Memory            int    `json:"memory,omitempty"`
	Project           string `json:"project,omitempty"`
	Templateid        string `json:"templateid,omitempty"`
	Password          string `json:"password,omitempty"`
	Nic               []struct {
		Id           string   `json:"id,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Type         string   `json:"type,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
	} `json:"nic,omitempty"`
	Isodisplaytext string `json:"isodisplaytext,omitempty"`
	Id             string `json:"id,omitempty"`
	Rootdeviceid   int    `json:"rootdeviceid,omitempty"`
	Displayvm      bool   `json:"displayvm,omitempty"`
	Publicipid     string `json:"publicipid,omitempty"`
	Affinitygroup  []struct {
		Type              string   `json:"type,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		Id                string   `json:"id,omitempty"`
		Description       string   `json:"description,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
		Name              string   `json:"name,omitempty"`
		Account           string   `json:"account,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Rootdevicetype string `json:"rootdevicetype,omitempty"`
	State          string `json:"state,omitempty"`
	Diskiowrite    int    `json:"diskiowrite,omitempty"`
	Name           string `json:"name,omitempty"`
	Zoneid         string `json:"zoneid,omitempty"`
	Diskkbswrite   int    `json:"diskkbswrite,omitempty"`
	Diskioread     int    `json:"diskioread,omitempty"`
	Servicestate   string `json:"servicestate,omitempty"`
	Zonename       string `json:"zonename,omitempty"`
	Groupid        string `json:"groupid,omitempty"`
	Cpuspeed       int    `json:"cpuspeed,omitempty"`
	Projectid      string `json:"projectid,omitempty"`
	Tags           []struct {
		Domain       string `json:"domain,omitempty"`
		Value        string `json:"value,omitempty"`
		Account      string `json:"account,omitempty"`
		Key          string `json:"key,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Project      string `json:"project,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Customer     string `json:"customer,omitempty"`
	} `json:"tags,omitempty"`
	Keypair               string            `json:"keypair,omitempty"`
	Serviceofferingid     string            `json:"serviceofferingid,omitempty"`
	Instancename          string            `json:"instancename,omitempty"`
	Isdynamicallyscalable bool              `json:"isdynamicallyscalable,omitempty"`
	Templatename          string            `json:"templatename,omitempty"`
	Templatedisplaytext   string            `json:"templatedisplaytext,omitempty"`
	Serviceofferingname   string            `json:"serviceofferingname,omitempty"`
	Guestosid             string            `json:"guestosid,omitempty"`
	Hostname              string            `json:"hostname,omitempty"`
	Publicip              string            `json:"publicip,omitempty"`
	Haenable              bool              `json:"haenable,omitempty"`
	Isoname               string            `json:"isoname,omitempty"`
	Networkkbswrite       int               `json:"networkkbswrite,omitempty"`
	Displayname           string            `json:"displayname,omitempty"`
	Created               string            `json:"created,omitempty"`
	Domainid              string            `json:"domainid,omitempty"`
	Isoid                 string            `json:"isoid,omitempty"`
	Cpunumber             int               `json:"cpunumber,omitempty"`
	Group                 string            `json:"group,omitempty"`
	Details               map[string]string `json:"details,omitempty"`
	Passwordenabled       bool              `json:"passwordenabled,omitempty"`
	Networkkbsread        int               `json:"networkkbsread,omitempty"`
}

type ListVirtualMachinesParams struct {
	p map[string]interface{}
}

func (p *ListVirtualMachinesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["affinitygroupid"]; found {
		u.Set("affinitygroupid", v.(string))
	}
	if v, found := p.p["details"]; found {
		vv := strings.Join(v.([]string), ", ")
		u.Set("details", vv)
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["forvirtualnetwork"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forvirtualnetwork", vv)
	}
	if v, found := p.p["groupid"]; found {
		u.Set("groupid", v.(string))
	}
	if v, found := p.p["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	if v, found := p.p["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["isoid"]; found {
		u.Set("isoid", v.(string))
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
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := p.p["storageid"]; found {
		u.Set("storageid", v.(string))
	}
	if v, found := p.p["tags"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["templateid"]; found {
		u.Set("templateid", v.(string))
	}
	if v, found := p.p["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListVirtualMachinesParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListVirtualMachinesParams) SetAffinitygroupid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["affinitygroupid"] = v
	return
}

func (p *ListVirtualMachinesParams) SetDetails(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
	return
}

func (p *ListVirtualMachinesParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListVirtualMachinesParams) SetForvirtualnetwork(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forvirtualnetwork"] = v
	return
}

func (p *ListVirtualMachinesParams) SetGroupid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["groupid"] = v
	return
}

func (p *ListVirtualMachinesParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
	return
}

func (p *ListVirtualMachinesParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
	return
}

func (p *ListVirtualMachinesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListVirtualMachinesParams) SetIsoid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isoid"] = v
	return
}

func (p *ListVirtualMachinesParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
	return
}

func (p *ListVirtualMachinesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListVirtualMachinesParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
	return
}

func (p *ListVirtualMachinesParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *ListVirtualMachinesParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
	return
}

func (p *ListVirtualMachinesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListVirtualMachinesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListVirtualMachinesParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
	return
}

func (p *ListVirtualMachinesParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *ListVirtualMachinesParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
	return
}

func (p *ListVirtualMachinesParams) SetStorageid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["storageid"] = v
	return
}

func (p *ListVirtualMachinesParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
	return
}

func (p *ListVirtualMachinesParams) SetTemplateid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templateid"] = v
	return
}

func (p *ListVirtualMachinesParams) SetVpcid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcid"] = v
	return
}

func (p *ListVirtualMachinesParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new ListVirtualMachinesParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewListVirtualMachinesParams() *ListVirtualMachinesParams {
	p := &ListVirtualMachinesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VirtualMachineService) GetVirtualMachineID(name string) (string, error) {
	p := &ListVirtualMachinesParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	l, err := s.ListVirtualMachines(p)
	if err != nil {
		return "", err
	}

	if l.Count == 0 {
		return "", fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.VirtualMachines[0].Id, nil
	}

	if l.Count > 1 {
		for _, v := range l.VirtualMachines {
			if v.Name == name {
				return v.Id, nil
			}
		}
	}
	return "", fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VirtualMachineService) GetVirtualMachineByName(name string) (*VirtualMachine, int, error) {
	id, err := s.GetVirtualMachineID(name)
	if err != nil {
		return nil, -1, err
	}

	r, count, err := s.GetVirtualMachineByID(id)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VirtualMachineService) GetVirtualMachineByID(id string) (*VirtualMachine, int, error) {
	p := &ListVirtualMachinesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	l, err := s.ListVirtualMachines(p)
	if err != nil {
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.VirtualMachines[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for VirtualMachine UUID: %s!", id)
}

// List the virtual machines owned by the account.
func (s *VirtualMachineService) ListVirtualMachines(p *ListVirtualMachinesParams) (*ListVirtualMachinesResponse, error) {
	resp, err := s.cs.newRequest("listVirtualMachines", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVirtualMachinesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListVirtualMachinesResponse struct {
	Count           int               `json:"count"`
	VirtualMachines []*VirtualMachine `json:"virtualmachine"`
}

type VirtualMachine struct {
	Tags []struct {
		Resourceid   string `json:"resourceid,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Project      string `json:"project,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Account      string `json:"account,omitempty"`
		Value        string `json:"value,omitempty"`
		Key          string `json:"key,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
	} `json:"tags,omitempty"`
	Isodisplaytext string            `json:"isodisplaytext,omitempty"`
	Account        string            `json:"account,omitempty"`
	Publicipid     string            `json:"publicipid,omitempty"`
	Details        map[string]string `json:"details,omitempty"`
	Isoid          string            `json:"isoid,omitempty"`
	Nic            []struct {
		Macaddress   string   `json:"macaddress,omitempty"`
		Type         string   `json:"type,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Id           string   `json:"id,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
	} `json:"nic,omitempty"`
	Groupid       string `json:"groupid,omitempty"`
	Memory        int    `json:"memory,omitempty"`
	Haenable      bool   `json:"haenable,omitempty"`
	Hostname      string `json:"hostname,omitempty"`
	Guestosid     string `json:"guestosid,omitempty"`
	Rootdeviceid  int    `json:"rootdeviceid,omitempty"`
	Templateid    string `json:"templateid,omitempty"`
	Group         string `json:"group,omitempty"`
	Domain        string `json:"domain,omitempty"`
	Projectid     string `json:"projectid,omitempty"`
	Diskkbswrite  int    `json:"diskkbswrite,omitempty"`
	Affinitygroup []struct {
		Domainid          string   `json:"domainid,omitempty"`
		Id                string   `json:"id,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Description       string   `json:"description,omitempty"`
		Type              string   `json:"type,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		Account           string   `json:"account,omitempty"`
		Name              string   `json:"name,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Displayname           string `json:"displayname,omitempty"`
	Networkkbswrite       int    `json:"networkkbswrite,omitempty"`
	Hypervisor            string `json:"hypervisor,omitempty"`
	Displayvm             bool   `json:"displayvm,omitempty"`
	Isdynamicallyscalable bool   `json:"isdynamicallyscalable,omitempty"`
	Project               string `json:"project,omitempty"`
	Hostid                string `json:"hostid,omitempty"`
	Diskkbsread           int    `json:"diskkbsread,omitempty"`
	Domainid              string `json:"domainid,omitempty"`
	Name                  string `json:"name,omitempty"`
	Passwordenabled       bool   `json:"passwordenabled,omitempty"`
	Templatedisplaytext   string `json:"templatedisplaytext,omitempty"`
	Zonename              string `json:"zonename,omitempty"`
	Cpunumber             int    `json:"cpunumber,omitempty"`
	Keypair               string `json:"keypair,omitempty"`
	Networkkbsread        int    `json:"networkkbsread,omitempty"`
	Created               string `json:"created,omitempty"`
	Forvirtualnetwork     bool   `json:"forvirtualnetwork,omitempty"`
	State                 string `json:"state,omitempty"`
	Serviceofferingname   string `json:"serviceofferingname,omitempty"`
	Publicip              string `json:"publicip,omitempty"`
	Serviceofferingid     string `json:"serviceofferingid,omitempty"`
	Id                    string `json:"id,omitempty"`
	Cpuused               string `json:"cpuused,omitempty"`
	Instancename          string `json:"instancename,omitempty"`
	Password              string `json:"password,omitempty"`
	Servicestate          string `json:"servicestate,omitempty"`
	Cpuspeed              int    `json:"cpuspeed,omitempty"`
	Zoneid                string `json:"zoneid,omitempty"`
	Rootdevicetype        string `json:"rootdevicetype,omitempty"`
	Securitygroup         []struct {
		Description string `json:"description,omitempty"`
		Projectid   string `json:"projectid,omitempty"`
		Id          string `json:"id,omitempty"`
		Tags        []struct {
			Domain       string `json:"domain,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Key          string `json:"key,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
			Value        string `json:"value,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Project      string `json:"project,omitempty"`
			Account      string `json:"account,omitempty"`
		} `json:"tags,omitempty"`
		Egressrule []struct {
			Cidr              string `json:"cidr,omitempty"`
			Account           string `json:"account,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
		} `json:"egressrule,omitempty"`
		Domainid    string `json:"domainid,omitempty"`
		Ingressrule []struct {
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Account           string `json:"account,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
		} `json:"ingressrule,omitempty"`
		Project string `json:"project,omitempty"`
		Domain  string `json:"domain,omitempty"`
		Account string `json:"account,omitempty"`
		Name    string `json:"name,omitempty"`
	} `json:"securitygroup,omitempty"`
	Diskioread   int    `json:"diskioread,omitempty"`
	Isoname      string `json:"isoname,omitempty"`
	Templatename string `json:"templatename,omitempty"`
	Diskiowrite  int    `json:"diskiowrite,omitempty"`
}

type GetVMPasswordParams struct {
	p map[string]interface{}
}

func (p *GetVMPasswordParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *GetVMPasswordParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new GetVMPasswordParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewGetVMPasswordParams(id string) *GetVMPasswordParams {
	p := &GetVMPasswordParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Returns an encrypted password for the VM
func (s *VirtualMachineService) GetVMPassword(p *GetVMPasswordParams) (*GetVMPasswordResponse, error) {
	resp, err := s.cs.newRequest("getVMPassword", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r GetVMPasswordResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type GetVMPasswordResponse struct {
	Encryptedpassword string `json:"encryptedpassword,omitempty"`
}

type RestoreVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *RestoreVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["templateid"]; found {
		u.Set("templateid", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *RestoreVirtualMachineParams) SetTemplateid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templateid"] = v
	return
}

func (p *RestoreVirtualMachineParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
	return
}

// You should always use this function to get a new RestoreVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewRestoreVirtualMachineParams(virtualmachineid string) *RestoreVirtualMachineParams {
	p := &RestoreVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// Restore a VM to original template/ISO or new template/ISO
func (s *VirtualMachineService) RestoreVirtualMachine(p *RestoreVirtualMachineParams) (*RestoreVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("restoreVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RestoreVirtualMachineResponse
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

type RestoreVirtualMachineResponse struct {
	JobID         string `json:"jobid,omitempty"`
	Zoneid        string `json:"zoneid,omitempty"`
	Affinitygroup []struct {
		Domain            string   `json:"domain,omitempty"`
		Type              string   `json:"type,omitempty"`
		Account           string   `json:"account,omitempty"`
		Description       string   `json:"description,omitempty"`
		Name              string   `json:"name,omitempty"`
		Id                string   `json:"id,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Instancename      string `json:"instancename,omitempty"`
	Diskioread        int    `json:"diskioread,omitempty"`
	Haenable          bool   `json:"haenable,omitempty"`
	Guestosid         string `json:"guestosid,omitempty"`
	Group             string `json:"group,omitempty"`
	Groupid           string `json:"groupid,omitempty"`
	Publicipid        string `json:"publicipid,omitempty"`
	Diskkbsread       int    `json:"diskkbsread,omitempty"`
	State             string `json:"state,omitempty"`
	Rootdeviceid      int    `json:"rootdeviceid,omitempty"`
	Memory            int    `json:"memory,omitempty"`
	Diskkbswrite      int    `json:"diskkbswrite,omitempty"`
	Isodisplaytext    string `json:"isodisplaytext,omitempty"`
	Keypair           string `json:"keypair,omitempty"`
	Created           string `json:"created,omitempty"`
	Id                string `json:"id,omitempty"`
	Displayvm         bool   `json:"displayvm,omitempty"`
	Account           string `json:"account,omitempty"`
	Hostname          string `json:"hostname,omitempty"`
	Domainid          string `json:"domainid,omitempty"`
	Serviceofferingid string `json:"serviceofferingid,omitempty"`
	Templateid        string `json:"templateid,omitempty"`
	Password          string `json:"password,omitempty"`
	Publicip          string `json:"publicip,omitempty"`
	Cpuspeed          int    `json:"cpuspeed,omitempty"`
	Cpuused           string `json:"cpuused,omitempty"`
	Displayname       string `json:"displayname,omitempty"`
	Isoname           string `json:"isoname,omitempty"`
	Passwordenabled   bool   `json:"passwordenabled,omitempty"`
	Tags              []struct {
		Account      string `json:"account,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Key          string `json:"key,omitempty"`
		Value        string `json:"value,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Project      string `json:"project,omitempty"`
	} `json:"tags,omitempty"`
	Isoid                 string            `json:"isoid,omitempty"`
	Networkkbsread        int               `json:"networkkbsread,omitempty"`
	Isdynamicallyscalable bool              `json:"isdynamicallyscalable,omitempty"`
	Details               map[string]string `json:"details,omitempty"`
	Name                  string            `json:"name,omitempty"`
	Hypervisor            string            `json:"hypervisor,omitempty"`
	Securitygroup         []struct {
		Projectid string `json:"projectid,omitempty"`
		Account   string `json:"account,omitempty"`
		Id        string `json:"id,omitempty"`
		Name      string `json:"name,omitempty"`
		Tags      []struct {
			Resourcetype string `json:"resourcetype,omitempty"`
			Value        string `json:"value,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Project      string `json:"project,omitempty"`
			Account      string `json:"account,omitempty"`
			Key          string `json:"key,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
		} `json:"tags,omitempty"`
		Egressrule []struct {
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Account           string `json:"account,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
		} `json:"egressrule,omitempty"`
		Ingressrule []struct {
			Account           string `json:"account,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
		} `json:"ingressrule,omitempty"`
		Domainid    string `json:"domainid,omitempty"`
		Description string `json:"description,omitempty"`
		Domain      string `json:"domain,omitempty"`
		Project     string `json:"project,omitempty"`
	} `json:"securitygroup,omitempty"`
	Diskiowrite         int    `json:"diskiowrite,omitempty"`
	Networkkbswrite     int    `json:"networkkbswrite,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
	Templatename        string `json:"templatename,omitempty"`
	Hostid              string `json:"hostid,omitempty"`
	Forvirtualnetwork   bool   `json:"forvirtualnetwork,omitempty"`
	Templatedisplaytext string `json:"templatedisplaytext,omitempty"`
	Projectid           string `json:"projectid,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Rootdevicetype      string `json:"rootdevicetype,omitempty"`
	Project             string `json:"project,omitempty"`
	Nic                 []struct {
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Type         string   `json:"type,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Id           string   `json:"id,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
	} `json:"nic,omitempty"`
	Servicestate string `json:"servicestate,omitempty"`
	Cpunumber    int    `json:"cpunumber,omitempty"`
}

type ChangeServiceForVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *ChangeServiceForVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["details"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("details[%d].key", i), k)
			u.Set(fmt.Sprintf("details[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	return u
}

func (p *ChangeServiceForVirtualMachineParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
	return
}

func (p *ChangeServiceForVirtualMachineParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ChangeServiceForVirtualMachineParams) SetServiceofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingid"] = v
	return
}

// You should always use this function to get a new ChangeServiceForVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewChangeServiceForVirtualMachineParams(id string, serviceofferingid string) *ChangeServiceForVirtualMachineParams {
	p := &ChangeServiceForVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["serviceofferingid"] = serviceofferingid
	return p
}

// Changes the service offering for a virtual machine. The virtual machine must be in a "Stopped" state for this command to take effect.
func (s *VirtualMachineService) ChangeServiceForVirtualMachine(p *ChangeServiceForVirtualMachineParams) (*ChangeServiceForVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("changeServiceForVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ChangeServiceForVirtualMachineResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ChangeServiceForVirtualMachineResponse struct {
	Id              string `json:"id,omitempty"`
	Cpuspeed        int    `json:"cpuspeed,omitempty"`
	Groupid         string `json:"groupid,omitempty"`
	Rootdevicetype  string `json:"rootdevicetype,omitempty"`
	Passwordenabled bool   `json:"passwordenabled,omitempty"`
	Domain          string `json:"domain,omitempty"`
	Hostname        string `json:"hostname,omitempty"`
	Isoid           string `json:"isoid,omitempty"`
	Keypair         string `json:"keypair,omitempty"`
	Domainid        string `json:"domainid,omitempty"`
	Guestosid       string `json:"guestosid,omitempty"`
	Cpuused         string `json:"cpuused,omitempty"`
	Diskioread      int    `json:"diskioread,omitempty"`
	Templateid      string `json:"templateid,omitempty"`
	Networkkbsread  int    `json:"networkkbsread,omitempty"`
	Rootdeviceid    int    `json:"rootdeviceid,omitempty"`
	Password        string `json:"password,omitempty"`
	Affinitygroup   []struct {
		Account           string   `json:"account,omitempty"`
		Description       string   `json:"description,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
		Type              string   `json:"type,omitempty"`
		Name              string   `json:"name,omitempty"`
		Id                string   `json:"id,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Name                  string `json:"name,omitempty"`
	Serviceofferingid     string `json:"serviceofferingid,omitempty"`
	Displayvm             bool   `json:"displayvm,omitempty"`
	Created               string `json:"created,omitempty"`
	Publicip              string `json:"publicip,omitempty"`
	Displayname           string `json:"displayname,omitempty"`
	Instancename          string `json:"instancename,omitempty"`
	Servicestate          string `json:"servicestate,omitempty"`
	Isdynamicallyscalable bool   `json:"isdynamicallyscalable,omitempty"`
	Publicipid            string `json:"publicipid,omitempty"`
	Diskkbsread           int    `json:"diskkbsread,omitempty"`
	Tags                  []struct {
		Value        string `json:"value,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Key          string `json:"key,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Account      string `json:"account,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Project      string `json:"project,omitempty"`
	} `json:"tags,omitempty"`
	Serviceofferingname string            `json:"serviceofferingname,omitempty"`
	Projectid           string            `json:"projectid,omitempty"`
	Networkkbswrite     int               `json:"networkkbswrite,omitempty"`
	State               string            `json:"state,omitempty"`
	Details             map[string]string `json:"details,omitempty"`
	Forvirtualnetwork   bool              `json:"forvirtualnetwork,omitempty"`
	Hypervisor          string            `json:"hypervisor,omitempty"`
	Project             string            `json:"project,omitempty"`
	Memory              int               `json:"memory,omitempty"`
	Haenable            bool              `json:"haenable,omitempty"`
	Templatename        string            `json:"templatename,omitempty"`
	Diskiowrite         int               `json:"diskiowrite,omitempty"`
	Templatedisplaytext string            `json:"templatedisplaytext,omitempty"`
	Hostid              string            `json:"hostid,omitempty"`
	Nic                 []struct {
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Type         string   `json:"type,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Id           string   `json:"id,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
	} `json:"nic,omitempty"`
	Securitygroup []struct {
		Egressrule []struct {
			Endport           int    `json:"endport,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Account           string `json:"account,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
		} `json:"egressrule,omitempty"`
		Ingressrule []struct {
			Ruleid            string `json:"ruleid,omitempty"`
			Account           string `json:"account,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Startport         int    `json:"startport,omitempty"`
		} `json:"ingressrule,omitempty"`
		Name string `json:"name,omitempty"`
		Tags []struct {
			Projectid    string `json:"projectid,omitempty"`
			Project      string `json:"project,omitempty"`
			Account      string `json:"account,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Value        string `json:"value,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Key          string `json:"key,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Domain       string `json:"domain,omitempty"`
		} `json:"tags,omitempty"`
		Domain      string `json:"domain,omitempty"`
		Id          string `json:"id,omitempty"`
		Account     string `json:"account,omitempty"`
		Domainid    string `json:"domainid,omitempty"`
		Project     string `json:"project,omitempty"`
		Description string `json:"description,omitempty"`
		Projectid   string `json:"projectid,omitempty"`
	} `json:"securitygroup,omitempty"`
	Zonename       string `json:"zonename,omitempty"`
	Account        string `json:"account,omitempty"`
	Group          string `json:"group,omitempty"`
	Isodisplaytext string `json:"isodisplaytext,omitempty"`
	Zoneid         string `json:"zoneid,omitempty"`
	Isoname        string `json:"isoname,omitempty"`
	Cpunumber      int    `json:"cpunumber,omitempty"`
	Diskkbswrite   int    `json:"diskkbswrite,omitempty"`
}

type ScaleVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *ScaleVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["details"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("details[%d].key", i), k)
			u.Set(fmt.Sprintf("details[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	return u
}

func (p *ScaleVirtualMachineParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
	return
}

func (p *ScaleVirtualMachineParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ScaleVirtualMachineParams) SetServiceofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingid"] = v
	return
}

// You should always use this function to get a new ScaleVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewScaleVirtualMachineParams(id string, serviceofferingid string) *ScaleVirtualMachineParams {
	p := &ScaleVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["serviceofferingid"] = serviceofferingid
	return p
}

// Scales the virtual machine to a new service offering.
func (s *VirtualMachineService) ScaleVirtualMachine(p *ScaleVirtualMachineParams) (*ScaleVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("scaleVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ScaleVirtualMachineResponse
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

type ScaleVirtualMachineResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Success     bool   `json:"success,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
}

type AssignVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *AssignVirtualMachineParams) toURLValues() url.Values {
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
	if v, found := p.p["networkids"]; found {
		vv := strings.Join(v.([]string), ", ")
		u.Set("networkids", vv)
	}
	if v, found := p.p["securitygroupids"]; found {
		vv := strings.Join(v.([]string), ", ")
		u.Set("securitygroupids", vv)
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *AssignVirtualMachineParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *AssignVirtualMachineParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *AssignVirtualMachineParams) SetNetworkids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkids"] = v
	return
}

func (p *AssignVirtualMachineParams) SetSecuritygroupids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["securitygroupids"] = v
	return
}

func (p *AssignVirtualMachineParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
	return
}

// You should always use this function to get a new AssignVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewAssignVirtualMachineParams(account string, domainid string, virtualmachineid string) *AssignVirtualMachineParams {
	p := &AssignVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["account"] = account
	p.p["domainid"] = domainid
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// Change ownership of a VM from one account to another. This API is available for Basic zones with security groups and Advanced zones with guest networks. A root administrator can reassign a VM from any account to any other account in any domain. A domain administrator can reassign a VM to any account in the same domain.
func (s *VirtualMachineService) AssignVirtualMachine(p *AssignVirtualMachineParams) (*AssignVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("assignVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AssignVirtualMachineResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type AssignVirtualMachineResponse struct {
	Displayvm           bool   `json:"displayvm,omitempty"`
	Memory              int    `json:"memory,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Hostname            string `json:"hostname,omitempty"`
	Created             string `json:"created,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
	Domainid            string `json:"domainid,omitempty"`
	Affinitygroup       []struct {
		Account           string   `json:"account,omitempty"`
		Description       string   `json:"description,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		Id                string   `json:"id,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
		Name              string   `json:"name,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Type              string   `json:"type,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Isoid               string `json:"isoid,omitempty"`
	Group               string `json:"group,omitempty"`
	Isoname             string `json:"isoname,omitempty"`
	Templatedisplaytext string `json:"templatedisplaytext,omitempty"`
	Hypervisor          string `json:"hypervisor,omitempty"`
	Rootdevicetype      string `json:"rootdevicetype,omitempty"`
	Diskkbswrite        int    `json:"diskkbswrite,omitempty"`
	Zoneid              string `json:"zoneid,omitempty"`
	Diskioread          int    `json:"diskioread,omitempty"`
	Publicipid          string `json:"publicipid,omitempty"`
	Account             string `json:"account,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Name                string `json:"name,omitempty"`
	Templateid          string `json:"templateid,omitempty"`
	Cpuspeed            int    `json:"cpuspeed,omitempty"`
	Servicestate        string `json:"servicestate,omitempty"`
	Projectid           string `json:"projectid,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
	Cpunumber           int    `json:"cpunumber,omitempty"`
	Rootdeviceid        int    `json:"rootdeviceid,omitempty"`
	Guestosid           string `json:"guestosid,omitempty"`
	Tags                []struct {
		Project      string `json:"project,omitempty"`
		Value        string `json:"value,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Account      string `json:"account,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Key          string `json:"key,omitempty"`
	} `json:"tags,omitempty"`
	Templatename      string            `json:"templatename,omitempty"`
	Cpuused           string            `json:"cpuused,omitempty"`
	Hostid            string            `json:"hostid,omitempty"`
	Id                string            `json:"id,omitempty"`
	Details           map[string]string `json:"details,omitempty"`
	Diskiowrite       int               `json:"diskiowrite,omitempty"`
	Password          string            `json:"password,omitempty"`
	Publicip          string            `json:"publicip,omitempty"`
	Forvirtualnetwork bool              `json:"forvirtualnetwork,omitempty"`
	Groupid           string            `json:"groupid,omitempty"`
	State             string            `json:"state,omitempty"`
	Instancename      string            `json:"instancename,omitempty"`
	Networkkbswrite   int               `json:"networkkbswrite,omitempty"`
	Haenable          bool              `json:"haenable,omitempty"`
	Isodisplaytext    string            `json:"isodisplaytext,omitempty"`
	Securitygroup     []struct {
		Description string `json:"description,omitempty"`
		Id          string `json:"id,omitempty"`
		Project     string `json:"project,omitempty"`
		Account     string `json:"account,omitempty"`
		Projectid   string `json:"projectid,omitempty"`
		Domainid    string `json:"domainid,omitempty"`
		Tags        []struct {
			Account      string `json:"account,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
			Project      string `json:"project,omitempty"`
			Value        string `json:"value,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Key          string `json:"key,omitempty"`
		} `json:"tags,omitempty"`
		Domain      string `json:"domain,omitempty"`
		Ingressrule []struct {
			Cidr              string `json:"cidr,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Account           string `json:"account,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
		} `json:"ingressrule,omitempty"`
		Egressrule []struct {
			Startport         int    `json:"startport,omitempty"`
			Account           string `json:"account,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
		} `json:"egressrule,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"securitygroup,omitempty"`
	Displayname     string `json:"displayname,omitempty"`
	Keypair         string `json:"keypair,omitempty"`
	Diskkbsread     int    `json:"diskkbsread,omitempty"`
	Passwordenabled bool   `json:"passwordenabled,omitempty"`
	Project         string `json:"project,omitempty"`
	Nic             []struct {
		Type         string   `json:"type,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Id           string   `json:"id,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
	} `json:"nic,omitempty"`
	Networkkbsread        int  `json:"networkkbsread,omitempty"`
	Isdynamicallyscalable bool `json:"isdynamicallyscalable,omitempty"`
}

type MigrateVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *MigrateVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	if v, found := p.p["storageid"]; found {
		u.Set("storageid", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *MigrateVirtualMachineParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
	return
}

func (p *MigrateVirtualMachineParams) SetStorageid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["storageid"] = v
	return
}

func (p *MigrateVirtualMachineParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
	return
}

// You should always use this function to get a new MigrateVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewMigrateVirtualMachineParams(virtualmachineid string) *MigrateVirtualMachineParams {
	p := &MigrateVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// Attempts Migration of a VM to a different host or Root volume of the vm to a different storage pool
func (s *VirtualMachineService) MigrateVirtualMachine(p *MigrateVirtualMachineParams) (*MigrateVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("migrateVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r MigrateVirtualMachineResponse
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

type MigrateVirtualMachineResponse struct {
	JobID         string `json:"jobid,omitempty"`
	Hostname      string `json:"hostname,omitempty"`
	Name          string `json:"name,omitempty"`
	Securitygroup []struct {
		Egressrule []struct {
			Icmptype          int    `json:"icmptype,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Account           string `json:"account,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Startport         int    `json:"startport,omitempty"`
		} `json:"egressrule,omitempty"`
		Ingressrule []struct {
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Account           string `json:"account,omitempty"`
		} `json:"ingressrule,omitempty"`
		Description string `json:"description,omitempty"`
		Domain      string `json:"domain,omitempty"`
		Project     string `json:"project,omitempty"`
		Projectid   string `json:"projectid,omitempty"`
		Account     string `json:"account,omitempty"`
		Name        string `json:"name,omitempty"`
		Tags        []struct {
			Value        string `json:"value,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Project      string `json:"project,omitempty"`
			Key          string `json:"key,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Account      string `json:"account,omitempty"`
		} `json:"tags,omitempty"`
		Domainid string `json:"domainid,omitempty"`
		Id       string `json:"id,omitempty"`
	} `json:"securitygroup,omitempty"`
	Publicip string `json:"publicip,omitempty"`
	Nic      []struct {
		Isdefault    bool     `json:"isdefault,omitempty"`
		Type         string   `json:"type,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Id           string   `json:"id,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
	} `json:"nic,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Publicipid          string `json:"publicipid,omitempty"`
	Tags                []struct {
		Resourcetype string `json:"resourcetype,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Key          string `json:"key,omitempty"`
		Project      string `json:"project,omitempty"`
		Value        string `json:"value,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Account      string `json:"account,omitempty"`
	} `json:"tags,omitempty"`
	Networkkbswrite       int               `json:"networkkbswrite,omitempty"`
	Zoneid                string            `json:"zoneid,omitempty"`
	Projectid             string            `json:"projectid,omitempty"`
	Groupid               string            `json:"groupid,omitempty"`
	Keypair               string            `json:"keypair,omitempty"`
	Diskkbswrite          int               `json:"diskkbswrite,omitempty"`
	Serviceofferingid     string            `json:"serviceofferingid,omitempty"`
	Diskioread            int               `json:"diskioread,omitempty"`
	Rootdevicetype        string            `json:"rootdevicetype,omitempty"`
	Templatename          string            `json:"templatename,omitempty"`
	Domainid              string            `json:"domainid,omitempty"`
	Cpuused               string            `json:"cpuused,omitempty"`
	Cpunumber             int               `json:"cpunumber,omitempty"`
	Diskiowrite           int               `json:"diskiowrite,omitempty"`
	Forvirtualnetwork     bool              `json:"forvirtualnetwork,omitempty"`
	Networkkbsread        int               `json:"networkkbsread,omitempty"`
	Guestosid             string            `json:"guestosid,omitempty"`
	Haenable              bool              `json:"haenable,omitempty"`
	Rootdeviceid          int               `json:"rootdeviceid,omitempty"`
	Displayname           string            `json:"displayname,omitempty"`
	Templateid            string            `json:"templateid,omitempty"`
	Hostid                string            `json:"hostid,omitempty"`
	Displayvm             bool              `json:"displayvm,omitempty"`
	Zonename              string            `json:"zonename,omitempty"`
	Passwordenabled       bool              `json:"passwordenabled,omitempty"`
	Id                    string            `json:"id,omitempty"`
	Domain                string            `json:"domain,omitempty"`
	Instancename          string            `json:"instancename,omitempty"`
	Account               string            `json:"account,omitempty"`
	Project               string            `json:"project,omitempty"`
	Isdynamicallyscalable bool              `json:"isdynamicallyscalable,omitempty"`
	Cpuspeed              int               `json:"cpuspeed,omitempty"`
	Created               string            `json:"created,omitempty"`
	Templatedisplaytext   string            `json:"templatedisplaytext,omitempty"`
	Isoid                 string            `json:"isoid,omitempty"`
	Memory                int               `json:"memory,omitempty"`
	Isodisplaytext        string            `json:"isodisplaytext,omitempty"`
	Details               map[string]string `json:"details,omitempty"`
	Password              string            `json:"password,omitempty"`
	State                 string            `json:"state,omitempty"`
	Affinitygroup         []struct {
		Type              string   `json:"type,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
		Id                string   `json:"id,omitempty"`
		Account           string   `json:"account,omitempty"`
		Description       string   `json:"description,omitempty"`
		Name              string   `json:"name,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Group        string `json:"group,omitempty"`
	Diskkbsread  int    `json:"diskkbsread,omitempty"`
	Isoname      string `json:"isoname,omitempty"`
	Servicestate string `json:"servicestate,omitempty"`
	Hypervisor   string `json:"hypervisor,omitempty"`
}

type MigrateVirtualMachineWithVolumeParams struct {
	p map[string]interface{}
}

func (p *MigrateVirtualMachineWithVolumeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	if v, found := p.p["migrateto"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("migrateto[%d].key", i), k)
			u.Set(fmt.Sprintf("migrateto[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *MigrateVirtualMachineWithVolumeParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
	return
}

func (p *MigrateVirtualMachineWithVolumeParams) SetMigrateto(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["migrateto"] = v
	return
}

func (p *MigrateVirtualMachineWithVolumeParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
	return
}

// You should always use this function to get a new MigrateVirtualMachineWithVolumeParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewMigrateVirtualMachineWithVolumeParams(hostid string, virtualmachineid string) *MigrateVirtualMachineWithVolumeParams {
	p := &MigrateVirtualMachineWithVolumeParams{}
	p.p = make(map[string]interface{})
	p.p["hostid"] = hostid
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// Attempts Migration of a VM with its volumes to a different host
func (s *VirtualMachineService) MigrateVirtualMachineWithVolume(p *MigrateVirtualMachineWithVolumeParams) (*MigrateVirtualMachineWithVolumeResponse, error) {
	resp, err := s.cs.newRequest("migrateVirtualMachineWithVolume", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r MigrateVirtualMachineWithVolumeResponse
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

type MigrateVirtualMachineWithVolumeResponse struct {
	JobID             string `json:"jobid,omitempty"`
	Hypervisor        string `json:"hypervisor,omitempty"`
	State             string `json:"state,omitempty"`
	Isoname           string `json:"isoname,omitempty"`
	Templateid        string `json:"templateid,omitempty"`
	Memory            int    `json:"memory,omitempty"`
	Groupid           string `json:"groupid,omitempty"`
	Networkkbsread    int    `json:"networkkbsread,omitempty"`
	Id                string `json:"id,omitempty"`
	Isodisplaytext    string `json:"isodisplaytext,omitempty"`
	Diskioread        int    `json:"diskioread,omitempty"`
	Forvirtualnetwork bool   `json:"forvirtualnetwork,omitempty"`
	Publicipid        string `json:"publicipid,omitempty"`
	Affinitygroup     []struct {
		Name              string   `json:"name,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		Type              string   `json:"type,omitempty"`
		Id                string   `json:"id,omitempty"`
		Description       string   `json:"description,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
		Account           string   `json:"account,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Account       string `json:"account,omitempty"`
	Haenable      bool   `json:"haenable,omitempty"`
	Cpuused       string `json:"cpuused,omitempty"`
	Name          string `json:"name,omitempty"`
	Displayvm     bool   `json:"displayvm,omitempty"`
	Project       string `json:"project,omitempty"`
	Instancename  string `json:"instancename,omitempty"`
	Securitygroup []struct {
		Egressrule []struct {
			Icmptype          int    `json:"icmptype,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Account           string `json:"account,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
		} `json:"egressrule,omitempty"`
		Account     string `json:"account,omitempty"`
		Domain      string `json:"domain,omitempty"`
		Description string `json:"description,omitempty"`
		Project     string `json:"project,omitempty"`
		Tags        []struct {
			Key          string `json:"key,omitempty"`
			Account      string `json:"account,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Project      string `json:"project,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
			Value        string `json:"value,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
		} `json:"tags,omitempty"`
		Domainid    string `json:"domainid,omitempty"`
		Id          string `json:"id,omitempty"`
		Ingressrule []struct {
			Startport         int    `json:"startport,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Account           string `json:"account,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
		} `json:"ingressrule,omitempty"`
		Projectid string `json:"projectid,omitempty"`
		Name      string `json:"name,omitempty"`
	} `json:"securitygroup,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Isoid               string `json:"isoid,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
	Diskkbsread         int    `json:"diskkbsread,omitempty"`
	Diskiowrite         int    `json:"diskiowrite,omitempty"`
	Nic                 []struct {
		Id           string   `json:"id,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Type         string   `json:"type,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
	} `json:"nic,omitempty"`
	Zoneid   string            `json:"zoneid,omitempty"`
	Details  map[string]string `json:"details,omitempty"`
	Cpuspeed int               `json:"cpuspeed,omitempty"`
	Hostid   string            `json:"hostid,omitempty"`
	Tags     []struct {
		Domainid     string `json:"domainid,omitempty"`
		Project      string `json:"project,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Value        string `json:"value,omitempty"`
		Account      string `json:"account,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Key          string `json:"key,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
	} `json:"tags,omitempty"`
	Cpunumber             int    `json:"cpunumber,omitempty"`
	Hostname              string `json:"hostname,omitempty"`
	Rootdeviceid          int    `json:"rootdeviceid,omitempty"`
	Keypair               string `json:"keypair,omitempty"`
	Templatename          string `json:"templatename,omitempty"`
	Group                 string `json:"group,omitempty"`
	Created               string `json:"created,omitempty"`
	Domainid              string `json:"domainid,omitempty"`
	Publicip              string `json:"publicip,omitempty"`
	Zonename              string `json:"zonename,omitempty"`
	Displayname           string `json:"displayname,omitempty"`
	Guestosid             string `json:"guestosid,omitempty"`
	Passwordenabled       bool   `json:"passwordenabled,omitempty"`
	Networkkbswrite       int    `json:"networkkbswrite,omitempty"`
	Diskkbswrite          int    `json:"diskkbswrite,omitempty"`
	Templatedisplaytext   string `json:"templatedisplaytext,omitempty"`
	Projectid             string `json:"projectid,omitempty"`
	Rootdevicetype        string `json:"rootdevicetype,omitempty"`
	Domain                string `json:"domain,omitempty"`
	Servicestate          string `json:"servicestate,omitempty"`
	Isdynamicallyscalable bool   `json:"isdynamicallyscalable,omitempty"`
	Password              string `json:"password,omitempty"`
}

type RecoverVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *RecoverVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *RecoverVirtualMachineParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new RecoverVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewRecoverVirtualMachineParams(id string) *RecoverVirtualMachineParams {
	p := &RecoverVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Recovers a virtual machine.
func (s *VirtualMachineService) RecoverVirtualMachine(p *RecoverVirtualMachineParams) (*RecoverVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("recoverVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RecoverVirtualMachineResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type RecoverVirtualMachineResponse struct {
	Publicip            string `json:"publicip,omitempty"`
	Created             string `json:"created,omitempty"`
	Publicipid          string `json:"publicipid,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Networkkbswrite     int    `json:"networkkbswrite,omitempty"`
	Name                string `json:"name,omitempty"`
	Nic                 []struct {
		Isolationuri string   `json:"isolationuri,omitempty"`
		Id           string   `json:"id,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Type         string   `json:"type,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
	} `json:"nic,omitempty"`
	Isoid                 string `json:"isoid,omitempty"`
	Cpuused               string `json:"cpuused,omitempty"`
	Displayvm             bool   `json:"displayvm,omitempty"`
	Hypervisor            string `json:"hypervisor,omitempty"`
	Servicestate          string `json:"servicestate,omitempty"`
	Keypair               string `json:"keypair,omitempty"`
	Id                    string `json:"id,omitempty"`
	Cpuspeed              int    `json:"cpuspeed,omitempty"`
	Group                 string `json:"group,omitempty"`
	Memory                int    `json:"memory,omitempty"`
	Templatedisplaytext   string `json:"templatedisplaytext,omitempty"`
	Hostid                string `json:"hostid,omitempty"`
	Displayname           string `json:"displayname,omitempty"`
	Cpunumber             int    `json:"cpunumber,omitempty"`
	Rootdeviceid          int    `json:"rootdeviceid,omitempty"`
	Projectid             string `json:"projectid,omitempty"`
	Guestosid             string `json:"guestosid,omitempty"`
	Account               string `json:"account,omitempty"`
	Haenable              bool   `json:"haenable,omitempty"`
	Templatename          string `json:"templatename,omitempty"`
	Hostname              string `json:"hostname,omitempty"`
	Domainid              string `json:"domainid,omitempty"`
	Isdynamicallyscalable bool   `json:"isdynamicallyscalable,omitempty"`
	Isoname               string `json:"isoname,omitempty"`
	Forvirtualnetwork     bool   `json:"forvirtualnetwork,omitempty"`
	Zoneid                string `json:"zoneid,omitempty"`
	State                 string `json:"state,omitempty"`
	Networkkbsread        int    `json:"networkkbsread,omitempty"`
	Serviceofferingid     string `json:"serviceofferingid,omitempty"`
	Password              string `json:"password,omitempty"`
	Diskkbsread           int    `json:"diskkbsread,omitempty"`
	Securitygroup         []struct {
		Domain      string `json:"domain,omitempty"`
		Description string `json:"description,omitempty"`
		Egressrule  []struct {
			Endport           int    `json:"endport,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Account           string `json:"account,omitempty"`
			Startport         int    `json:"startport,omitempty"`
		} `json:"egressrule,omitempty"`
		Name        string `json:"name,omitempty"`
		Ingressrule []struct {
			Account           string `json:"account,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Startport         int    `json:"startport,omitempty"`
		} `json:"ingressrule,omitempty"`
		Domainid  string `json:"domainid,omitempty"`
		Projectid string `json:"projectid,omitempty"`
		Account   string `json:"account,omitempty"`
		Id        string `json:"id,omitempty"`
		Project   string `json:"project,omitempty"`
		Tags      []struct {
			Domain       string `json:"domain,omitempty"`
			Key          string `json:"key,omitempty"`
			Project      string `json:"project,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Account      string `json:"account,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Value        string `json:"value,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
		} `json:"tags,omitempty"`
	} `json:"securitygroup,omitempty"`
	Diskioread     int               `json:"diskioread,omitempty"`
	Details        map[string]string `json:"details,omitempty"`
	Groupid        string            `json:"groupid,omitempty"`
	Isodisplaytext string            `json:"isodisplaytext,omitempty"`
	Diskkbswrite   int               `json:"diskkbswrite,omitempty"`
	Tags           []struct {
		Domainid     string `json:"domainid,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Account      string `json:"account,omitempty"`
		Value        string `json:"value,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Project      string `json:"project,omitempty"`
		Key          string `json:"key,omitempty"`
	} `json:"tags,omitempty"`
	Passwordenabled bool   `json:"passwordenabled,omitempty"`
	Domain          string `json:"domain,omitempty"`
	Rootdevicetype  string `json:"rootdevicetype,omitempty"`
	Project         string `json:"project,omitempty"`
	Templateid      string `json:"templateid,omitempty"`
	Instancename    string `json:"instancename,omitempty"`
	Affinitygroup   []struct {
		Name              string   `json:"name,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Id                string   `json:"id,omitempty"`
		Type              string   `json:"type,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
		Description       string   `json:"description,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		Account           string   `json:"account,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Diskiowrite int    `json:"diskiowrite,omitempty"`
	Zonename    string `json:"zonename,omitempty"`
}

type ExpungeVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *ExpungeVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *ExpungeVirtualMachineParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new ExpungeVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewExpungeVirtualMachineParams(id string) *ExpungeVirtualMachineParams {
	p := &ExpungeVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Expunge a virtual machine. Once expunged, it cannot be recoverd.
func (s *VirtualMachineService) ExpungeVirtualMachine(p *ExpungeVirtualMachineParams) (*ExpungeVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("expungeVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ExpungeVirtualMachineResponse
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

type ExpungeVirtualMachineResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Success     bool   `json:"success,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
}

type CleanVMReservationsParams struct {
	p map[string]interface{}
}

func (p *CleanVMReservationsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	return u
}

// You should always use this function to get a new CleanVMReservationsParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewCleanVMReservationsParams() *CleanVMReservationsParams {
	p := &CleanVMReservationsParams{}
	p.p = make(map[string]interface{})
	return p
}

// Cleanups VM reservations in the database.
func (s *VirtualMachineService) CleanVMReservations(p *CleanVMReservationsParams) (*CleanVMReservationsResponse, error) {
	resp, err := s.cs.newRequest("cleanVMReservations", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CleanVMReservationsResponse
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

type CleanVMReservationsResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type AddNicToVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *AddNicToVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["ipaddress"]; found {
		u.Set("ipaddress", v.(string))
	}
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *AddNicToVirtualMachineParams) SetIpaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddress"] = v
	return
}

func (p *AddNicToVirtualMachineParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
	return
}

func (p *AddNicToVirtualMachineParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
	return
}

// You should always use this function to get a new AddNicToVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewAddNicToVirtualMachineParams(networkid string, virtualmachineid string) *AddNicToVirtualMachineParams {
	p := &AddNicToVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["networkid"] = networkid
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// Adds VM to specified network by creating a NIC
func (s *VirtualMachineService) AddNicToVirtualMachine(p *AddNicToVirtualMachineParams) (*AddNicToVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("addNicToVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddNicToVirtualMachineResponse
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

type AddNicToVirtualMachineResponse struct {
	JobID          string `json:"jobid,omitempty"`
	Password       string `json:"password,omitempty"`
	Networkkbsread int    `json:"networkkbsread,omitempty"`
	Nic            []struct {
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Type         string   `json:"type,omitempty"`
		Id           string   `json:"id,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
	} `json:"nic,omitempty"`
	Displayname           string            `json:"displayname,omitempty"`
	Diskiowrite           int               `json:"diskiowrite,omitempty"`
	Isdynamicallyscalable bool              `json:"isdynamicallyscalable,omitempty"`
	Diskioread            int               `json:"diskioread,omitempty"`
	Hostid                string            `json:"hostid,omitempty"`
	Id                    string            `json:"id,omitempty"`
	Guestosid             string            `json:"guestosid,omitempty"`
	Displayvm             bool              `json:"displayvm,omitempty"`
	Cpuspeed              int               `json:"cpuspeed,omitempty"`
	Details               map[string]string `json:"details,omitempty"`
	Templateid            string            `json:"templateid,omitempty"`
	Serviceofferingname   string            `json:"serviceofferingname,omitempty"`
	Instancename          string            `json:"instancename,omitempty"`
	Hostname              string            `json:"hostname,omitempty"`
	Created               string            `json:"created,omitempty"`
	Templatename          string            `json:"templatename,omitempty"`
	Account               string            `json:"account,omitempty"`
	Project               string            `json:"project,omitempty"`
	Zonename              string            `json:"zonename,omitempty"`
	State                 string            `json:"state,omitempty"`
	Hypervisor            string            `json:"hypervisor,omitempty"`
	Haenable              bool              `json:"haenable,omitempty"`
	Domain                string            `json:"domain,omitempty"`
	Rootdeviceid          int               `json:"rootdeviceid,omitempty"`
	Rootdevicetype        string            `json:"rootdevicetype,omitempty"`
	Diskkbsread           int               `json:"diskkbsread,omitempty"`
	Projectid             string            `json:"projectid,omitempty"`
	Isodisplaytext        string            `json:"isodisplaytext,omitempty"`
	Cpunumber             int               `json:"cpunumber,omitempty"`
	Publicipid            string            `json:"publicipid,omitempty"`
	Keypair               string            `json:"keypair,omitempty"`
	Tags                  []struct {
		Domainid     string `json:"domainid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Project      string `json:"project,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Account      string `json:"account,omitempty"`
		Value        string `json:"value,omitempty"`
		Key          string `json:"key,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
	} `json:"tags,omitempty"`
	Isoname       string `json:"isoname,omitempty"`
	Affinitygroup []struct {
		Account           string   `json:"account,omitempty"`
		Name              string   `json:"name,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Id                string   `json:"id,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
		Description       string   `json:"description,omitempty"`
		Type              string   `json:"type,omitempty"`
		Domain            string   `json:"domain,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Isoid               string `json:"isoid,omitempty"`
	Cpuused             string `json:"cpuused,omitempty"`
	Passwordenabled     bool   `json:"passwordenabled,omitempty"`
	Groupid             string `json:"groupid,omitempty"`
	Memory              int    `json:"memory,omitempty"`
	Templatedisplaytext string `json:"templatedisplaytext,omitempty"`
	Domainid            string `json:"domainid,omitempty"`
	Diskkbswrite        int    `json:"diskkbswrite,omitempty"`
	Name                string `json:"name,omitempty"`
	Servicestate        string `json:"servicestate,omitempty"`
	Securitygroup       []struct {
		Id         string `json:"id,omitempty"`
		Account    string `json:"account,omitempty"`
		Projectid  string `json:"projectid,omitempty"`
		Name       string `json:"name,omitempty"`
		Egressrule []struct {
			Cidr              string `json:"cidr,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Account           string `json:"account,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Startport         int    `json:"startport,omitempty"`
		} `json:"egressrule,omitempty"`
		Domain      string `json:"domain,omitempty"`
		Description string `json:"description,omitempty"`
		Tags        []struct {
			Domain       string `json:"domain,omitempty"`
			Account      string `json:"account,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Value        string `json:"value,omitempty"`
			Key          string `json:"key,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Project      string `json:"project,omitempty"`
		} `json:"tags,omitempty"`
		Ingressrule []struct {
			Account           string `json:"account,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
		} `json:"ingressrule,omitempty"`
		Domainid string `json:"domainid,omitempty"`
		Project  string `json:"project,omitempty"`
	} `json:"securitygroup,omitempty"`
	Group             string `json:"group,omitempty"`
	Serviceofferingid string `json:"serviceofferingid,omitempty"`
	Networkkbswrite   int    `json:"networkkbswrite,omitempty"`
	Forvirtualnetwork bool   `json:"forvirtualnetwork,omitempty"`
	Zoneid            string `json:"zoneid,omitempty"`
	Publicip          string `json:"publicip,omitempty"`
}

type RemoveNicFromVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *RemoveNicFromVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["nicid"]; found {
		u.Set("nicid", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *RemoveNicFromVirtualMachineParams) SetNicid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nicid"] = v
	return
}

func (p *RemoveNicFromVirtualMachineParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
	return
}

// You should always use this function to get a new RemoveNicFromVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewRemoveNicFromVirtualMachineParams(nicid string, virtualmachineid string) *RemoveNicFromVirtualMachineParams {
	p := &RemoveNicFromVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["nicid"] = nicid
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// Removes VM from specified network by deleting a NIC
func (s *VirtualMachineService) RemoveNicFromVirtualMachine(p *RemoveNicFromVirtualMachineParams) (*RemoveNicFromVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("removeNicFromVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RemoveNicFromVirtualMachineResponse
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

type RemoveNicFromVirtualMachineResponse struct {
	JobID          string            `json:"jobid,omitempty"`
	Isoname        string            `json:"isoname,omitempty"`
	Zoneid         string            `json:"zoneid,omitempty"`
	Guestosid      string            `json:"guestosid,omitempty"`
	Isoid          string            `json:"isoid,omitempty"`
	Publicipid     string            `json:"publicipid,omitempty"`
	Password       string            `json:"password,omitempty"`
	Account        string            `json:"account,omitempty"`
	Isodisplaytext string            `json:"isodisplaytext,omitempty"`
	Networkkbsread int               `json:"networkkbsread,omitempty"`
	Created        string            `json:"created,omitempty"`
	Rootdevicetype string            `json:"rootdevicetype,omitempty"`
	Memory         int               `json:"memory,omitempty"`
	Projectid      string            `json:"projectid,omitempty"`
	Details        map[string]string `json:"details,omitempty"`
	Diskkbsread    int               `json:"diskkbsread,omitempty"`
	Cpuused        string            `json:"cpuused,omitempty"`
	Cpuspeed       int               `json:"cpuspeed,omitempty"`
	Templatename   string            `json:"templatename,omitempty"`
	Displayname    string            `json:"displayname,omitempty"`
	Diskioread     int               `json:"diskioread,omitempty"`
	Haenable       bool              `json:"haenable,omitempty"`
	Publicip       string            `json:"publicip,omitempty"`
	Nic            []struct {
		Gateway      string   `json:"gateway,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Id           string   `json:"id,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Type         string   `json:"type,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
	} `json:"nic,omitempty"`
	Serviceofferingname   string `json:"serviceofferingname,omitempty"`
	Name                  string `json:"name,omitempty"`
	Keypair               string `json:"keypair,omitempty"`
	Hypervisor            string `json:"hypervisor,omitempty"`
	Rootdeviceid          int    `json:"rootdeviceid,omitempty"`
	Templateid            string `json:"templateid,omitempty"`
	Cpunumber             int    `json:"cpunumber,omitempty"`
	Diskiowrite           int    `json:"diskiowrite,omitempty"`
	Isdynamicallyscalable bool   `json:"isdynamicallyscalable,omitempty"`
	Id                    string `json:"id,omitempty"`
	Instancename          string `json:"instancename,omitempty"`
	State                 string `json:"state,omitempty"`
	Forvirtualnetwork     bool   `json:"forvirtualnetwork,omitempty"`
	Domainid              string `json:"domainid,omitempty"`
	Servicestate          string `json:"servicestate,omitempty"`
	Hostname              string `json:"hostname,omitempty"`
	Displayvm             bool   `json:"displayvm,omitempty"`
	Diskkbswrite          int    `json:"diskkbswrite,omitempty"`
	Project               string `json:"project,omitempty"`
	Serviceofferingid     string `json:"serviceofferingid,omitempty"`
	Groupid               string `json:"groupid,omitempty"`
	Group                 string `json:"group,omitempty"`
	Zonename              string `json:"zonename,omitempty"`
	Passwordenabled       bool   `json:"passwordenabled,omitempty"`
	Affinitygroup         []struct {
		Domainid          string   `json:"domainid,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Account           string   `json:"account,omitempty"`
		Description       string   `json:"description,omitempty"`
		Name              string   `json:"name,omitempty"`
		Type              string   `json:"type,omitempty"`
		Id                string   `json:"id,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Networkkbswrite int `json:"networkkbswrite,omitempty"`
	Securitygroup   []struct {
		Tags []struct {
			Projectid    string `json:"projectid,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Key          string `json:"key,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Account      string `json:"account,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Project      string `json:"project,omitempty"`
			Value        string `json:"value,omitempty"`
		} `json:"tags,omitempty"`
		Egressrule []struct {
			Endport           int    `json:"endport,omitempty"`
			Account           string `json:"account,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Startport         int    `json:"startport,omitempty"`
		} `json:"egressrule,omitempty"`
		Domainid    string `json:"domainid,omitempty"`
		Projectid   string `json:"projectid,omitempty"`
		Domain      string `json:"domain,omitempty"`
		Description string `json:"description,omitempty"`
		Ingressrule []struct {
			Protocol          string `json:"protocol,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Account           string `json:"account,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
		} `json:"ingressrule,omitempty"`
		Project string `json:"project,omitempty"`
		Name    string `json:"name,omitempty"`
		Account string `json:"account,omitempty"`
		Id      string `json:"id,omitempty"`
	} `json:"securitygroup,omitempty"`
	Hostid string `json:"hostid,omitempty"`
	Tags   []struct {
		Domainid     string `json:"domainid,omitempty"`
		Project      string `json:"project,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Account      string `json:"account,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Value        string `json:"value,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Key          string `json:"key,omitempty"`
	} `json:"tags,omitempty"`
	Templatedisplaytext string `json:"templatedisplaytext,omitempty"`
	Domain              string `json:"domain,omitempty"`
}

type UpdateDefaultNicForVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *UpdateDefaultNicForVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["nicid"]; found {
		u.Set("nicid", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *UpdateDefaultNicForVirtualMachineParams) SetNicid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nicid"] = v
	return
}

func (p *UpdateDefaultNicForVirtualMachineParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
	return
}

// You should always use this function to get a new UpdateDefaultNicForVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewUpdateDefaultNicForVirtualMachineParams(nicid string, virtualmachineid string) *UpdateDefaultNicForVirtualMachineParams {
	p := &UpdateDefaultNicForVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["nicid"] = nicid
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// Changes the default NIC on a VM
func (s *VirtualMachineService) UpdateDefaultNicForVirtualMachine(p *UpdateDefaultNicForVirtualMachineParams) (*UpdateDefaultNicForVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("updateDefaultNicForVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateDefaultNicForVirtualMachineResponse
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

type UpdateDefaultNicForVirtualMachineResponse struct {
	JobID string `json:"jobid,omitempty"`
	Nic   []struct {
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Id           string   `json:"id,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Type         string   `json:"type,omitempty"`
	} `json:"nic,omitempty"`
	Affinitygroup []struct {
		Account           string   `json:"account,omitempty"`
		Id                string   `json:"id,omitempty"`
		Name              string   `json:"name,omitempty"`
		Description       string   `json:"description,omitempty"`
		Type              string   `json:"type,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Securitygroup []struct {
		Name        string `json:"name,omitempty"`
		Project     string `json:"project,omitempty"`
		Account     string `json:"account,omitempty"`
		Domain      string `json:"domain,omitempty"`
		Ingressrule []struct {
			Icmptype          int    `json:"icmptype,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Account           string `json:"account,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Startport         int    `json:"startport,omitempty"`
		} `json:"ingressrule,omitempty"`
		Domainid string `json:"domainid,omitempty"`
		Tags     []struct {
			Resourceid   string `json:"resourceid,omitempty"`
			Key          string `json:"key,omitempty"`
			Account      string `json:"account,omitempty"`
			Value        string `json:"value,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Project      string `json:"project,omitempty"`
		} `json:"tags,omitempty"`
		Projectid   string `json:"projectid,omitempty"`
		Description string `json:"description,omitempty"`
		Id          string `json:"id,omitempty"`
		Egressrule  []struct {
			Account           string `json:"account,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Startport         int    `json:"startport,omitempty"`
		} `json:"egressrule,omitempty"`
	} `json:"securitygroup,omitempty"`
	Isoname               string            `json:"isoname,omitempty"`
	Domainid              string            `json:"domainid,omitempty"`
	Cpunumber             int               `json:"cpunumber,omitempty"`
	Zoneid                string            `json:"zoneid,omitempty"`
	Forvirtualnetwork     bool              `json:"forvirtualnetwork,omitempty"`
	Projectid             string            `json:"projectid,omitempty"`
	Templateid            string            `json:"templateid,omitempty"`
	Servicestate          string            `json:"servicestate,omitempty"`
	Serviceofferingid     string            `json:"serviceofferingid,omitempty"`
	Haenable              bool              `json:"haenable,omitempty"`
	Cpuspeed              int               `json:"cpuspeed,omitempty"`
	Passwordenabled       bool              `json:"passwordenabled,omitempty"`
	Diskiowrite           int               `json:"diskiowrite,omitempty"`
	Hypervisor            string            `json:"hypervisor,omitempty"`
	Serviceofferingname   string            `json:"serviceofferingname,omitempty"`
	Networkkbsread        int               `json:"networkkbsread,omitempty"`
	Templatedisplaytext   string            `json:"templatedisplaytext,omitempty"`
	State                 string            `json:"state,omitempty"`
	Isdynamicallyscalable bool              `json:"isdynamicallyscalable,omitempty"`
	Networkkbswrite       int               `json:"networkkbswrite,omitempty"`
	Publicip              string            `json:"publicip,omitempty"`
	Id                    string            `json:"id,omitempty"`
	Instancename          string            `json:"instancename,omitempty"`
	Details               map[string]string `json:"details,omitempty"`
	Isodisplaytext        string            `json:"isodisplaytext,omitempty"`
	Domain                string            `json:"domain,omitempty"`
	Diskkbswrite          int               `json:"diskkbswrite,omitempty"`
	Memory                int               `json:"memory,omitempty"`
	Rootdevicetype        string            `json:"rootdevicetype,omitempty"`
	Diskkbsread           int               `json:"diskkbsread,omitempty"`
	Keypair               string            `json:"keypair,omitempty"`
	Hostname              string            `json:"hostname,omitempty"`
	Created               string            `json:"created,omitempty"`
	Rootdeviceid          int               `json:"rootdeviceid,omitempty"`
	Isoid                 string            `json:"isoid,omitempty"`
	Cpuused               string            `json:"cpuused,omitempty"`
	Project               string            `json:"project,omitempty"`
	Groupid               string            `json:"groupid,omitempty"`
	Password              string            `json:"password,omitempty"`
	Publicipid            string            `json:"publicipid,omitempty"`
	Hostid                string            `json:"hostid,omitempty"`
	Account               string            `json:"account,omitempty"`
	Guestosid             string            `json:"guestosid,omitempty"`
	Diskioread            int               `json:"diskioread,omitempty"`
	Templatename          string            `json:"templatename,omitempty"`
	Displayvm             bool              `json:"displayvm,omitempty"`
	Group                 string            `json:"group,omitempty"`
	Zonename              string            `json:"zonename,omitempty"`
	Tags                  []struct {
		Domainid     string `json:"domainid,omitempty"`
		Project      string `json:"project,omitempty"`
		Account      string `json:"account,omitempty"`
		Key          string `json:"key,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Value        string `json:"value,omitempty"`
	} `json:"tags,omitempty"`
	Displayname string `json:"displayname,omitempty"`
	Name        string `json:"name,omitempty"`
}
