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

		var r DeployVirtualMachineResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type DeployVirtualMachineResponse struct {
	JobID                 string `json:"jobid,omitempty"`
	Rootdeviceid          int    `json:"rootdeviceid,omitempty"`
	Servicestate          string `json:"servicestate,omitempty"`
	Displayvm             bool   `json:"displayvm,omitempty"`
	Passwordenabled       bool   `json:"passwordenabled,omitempty"`
	State                 string `json:"state,omitempty"`
	Hostname              string `json:"hostname,omitempty"`
	Networkkbswrite       int    `json:"networkkbswrite,omitempty"`
	Templateid            string `json:"templateid,omitempty"`
	Publicipid            string `json:"publicipid,omitempty"`
	Rootdevicetype        string `json:"rootdevicetype,omitempty"`
	Keypair               string `json:"keypair,omitempty"`
	Networkkbsread        int    `json:"networkkbsread,omitempty"`
	Instancename          string `json:"instancename,omitempty"`
	Isdynamicallyscalable bool   `json:"isdynamicallyscalable,omitempty"`
	Id                    string `json:"id,omitempty"`
	Affinitygroup         []struct {
		Domain            string   `json:"domain,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
		Type              string   `json:"type,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Description       string   `json:"description,omitempty"`
		Account           string   `json:"account,omitempty"`
		Name              string   `json:"name,omitempty"`
		Id                string   `json:"id,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Isoname             string            `json:"isoname,omitempty"`
	Guestosid           string            `json:"guestosid,omitempty"`
	Publicip            string            `json:"publicip,omitempty"`
	Name                string            `json:"name,omitempty"`
	Domainid            string            `json:"domainid,omitempty"`
	Zonename            string            `json:"zonename,omitempty"`
	Details             map[string]string `json:"details,omitempty"`
	Serviceofferingname string            `json:"serviceofferingname,omitempty"`
	Tags                []struct {
		Domain       string `json:"domain,omitempty"`
		Project      string `json:"project,omitempty"`
		Key          string `json:"key,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Account      string `json:"account,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Value        string `json:"value,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
	} `json:"tags,omitempty"`
	Created             string `json:"created,omitempty"`
	Diskkbswrite        int    `json:"diskkbswrite,omitempty"`
	Cpuspeed            int    `json:"cpuspeed,omitempty"`
	Haenable            bool   `json:"haenable,omitempty"`
	Forvirtualnetwork   bool   `json:"forvirtualnetwork,omitempty"`
	Hostid              string `json:"hostid,omitempty"`
	Projectid           string `json:"projectid,omitempty"`
	Cpunumber           int    `json:"cpunumber,omitempty"`
	Templatedisplaytext string `json:"templatedisplaytext,omitempty"`
	Password            string `json:"password,omitempty"`
	Isoid               string `json:"isoid,omitempty"`
	Displayname         string `json:"displayname,omitempty"`
	Groupid             string `json:"groupid,omitempty"`
	Group               string `json:"group,omitempty"`
	Account             string `json:"account,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
	Zoneid              string `json:"zoneid,omitempty"`
	Nic                 []struct {
		Type         string   `json:"type,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Id           string   `json:"id,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
	} `json:"nic,omitempty"`
	Templatename   string `json:"templatename,omitempty"`
	Memory         int    `json:"memory,omitempty"`
	Diskkbsread    int    `json:"diskkbsread,omitempty"`
	Hypervisor     string `json:"hypervisor,omitempty"`
	Project        string `json:"project,omitempty"`
	Isodisplaytext string `json:"isodisplaytext,omitempty"`
	Securitygroup  []struct {
		Description string `json:"description,omitempty"`
		Ingressrule []struct {
			Icmptype          int    `json:"icmptype,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Account           string `json:"account,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
		} `json:"ingressrule,omitempty"`
		Tags []struct {
			Resourceid   string `json:"resourceid,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Value        string `json:"value,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Project      string `json:"project,omitempty"`
			Key          string `json:"key,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Account      string `json:"account,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
		} `json:"tags,omitempty"`
		Name       string `json:"name,omitempty"`
		Projectid  string `json:"projectid,omitempty"`
		Account    string `json:"account,omitempty"`
		Domainid   string `json:"domainid,omitempty"`
		Id         string `json:"id,omitempty"`
		Egressrule []struct {
			Cidr              string `json:"cidr,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Account           string `json:"account,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Startport         int    `json:"startport,omitempty"`
		} `json:"egressrule,omitempty"`
		Domain  string `json:"domain,omitempty"`
		Project string `json:"project,omitempty"`
	} `json:"securitygroup,omitempty"`
	Cpuused     string `json:"cpuused,omitempty"`
	Diskioread  int    `json:"diskioread,omitempty"`
	Diskiowrite int    `json:"diskiowrite,omitempty"`
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

		var r DestroyVirtualMachineResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type DestroyVirtualMachineResponse struct {
	JobID                 string `json:"jobid,omitempty"`
	Rootdevicetype        string `json:"rootdevicetype,omitempty"`
	Password              string `json:"password,omitempty"`
	Templatename          string `json:"templatename,omitempty"`
	Domainid              string `json:"domainid,omitempty"`
	Diskkbsread           int    `json:"diskkbsread,omitempty"`
	Servicestate          string `json:"servicestate,omitempty"`
	Memory                int    `json:"memory,omitempty"`
	Isodisplaytext        string `json:"isodisplaytext,omitempty"`
	Isdynamicallyscalable bool   `json:"isdynamicallyscalable,omitempty"`
	Hypervisor            string `json:"hypervisor,omitempty"`
	Serviceofferingid     string `json:"serviceofferingid,omitempty"`
	Templatedisplaytext   string `json:"templatedisplaytext,omitempty"`
	Keypair               string `json:"keypair,omitempty"`
	Diskkbswrite          int    `json:"diskkbswrite,omitempty"`
	Nic                   []struct {
		Id           string   `json:"id,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Type         string   `json:"type,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
	} `json:"nic,omitempty"`
	Networkkbsread      int               `json:"networkkbsread,omitempty"`
	Isoid               string            `json:"isoid,omitempty"`
	Created             string            `json:"created,omitempty"`
	Haenable            bool              `json:"haenable,omitempty"`
	Templateid          string            `json:"templateid,omitempty"`
	Details             map[string]string `json:"details,omitempty"`
	Groupid             string            `json:"groupid,omitempty"`
	Hostid              string            `json:"hostid,omitempty"`
	Publicip            string            `json:"publicip,omitempty"`
	Group               string            `json:"group,omitempty"`
	Passwordenabled     bool              `json:"passwordenabled,omitempty"`
	Serviceofferingname string            `json:"serviceofferingname,omitempty"`
	Cpuspeed            int               `json:"cpuspeed,omitempty"`
	Securitygroup       []struct {
		Account string `json:"account,omitempty"`
		Name    string `json:"name,omitempty"`
		Tags    []struct {
			Resourcetype string `json:"resourcetype,omitempty"`
			Key          string `json:"key,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
			Project      string `json:"project,omitempty"`
			Account      string `json:"account,omitempty"`
			Value        string `json:"value,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
		} `json:"tags,omitempty"`
		Egressrule []struct {
			Ruleid            string `json:"ruleid,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Account           string `json:"account,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
		} `json:"egressrule,omitempty"`
		Id          string `json:"id,omitempty"`
		Project     string `json:"project,omitempty"`
		Projectid   string `json:"projectid,omitempty"`
		Domain      string `json:"domain,omitempty"`
		Domainid    string `json:"domainid,omitempty"`
		Ingressrule []struct {
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Account           string `json:"account,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
		} `json:"ingressrule,omitempty"`
		Description string `json:"description,omitempty"`
	} `json:"securitygroup,omitempty"`
	Guestosid       string `json:"guestosid,omitempty"`
	Displayvm       bool   `json:"displayvm,omitempty"`
	Diskioread      int    `json:"diskioread,omitempty"`
	Project         string `json:"project,omitempty"`
	Isoname         string `json:"isoname,omitempty"`
	Networkkbswrite int    `json:"networkkbswrite,omitempty"`
	Displayname     string `json:"displayname,omitempty"`
	Diskiowrite     int    `json:"diskiowrite,omitempty"`
	Id              string `json:"id,omitempty"`
	Affinitygroup   []struct {
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Account           string   `json:"account,omitempty"`
		Name              string   `json:"name,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		Description       string   `json:"description,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
		Id                string   `json:"id,omitempty"`
		Type              string   `json:"type,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Cpunumber int `json:"cpunumber,omitempty"`
	Tags      []struct {
		Resourceid   string `json:"resourceid,omitempty"`
		Key          string `json:"key,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Account      string `json:"account,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Value        string `json:"value,omitempty"`
		Project      string `json:"project,omitempty"`
	} `json:"tags,omitempty"`
	Domain            string `json:"domain,omitempty"`
	Cpuused           string `json:"cpuused,omitempty"`
	State             string `json:"state,omitempty"`
	Name              string `json:"name,omitempty"`
	Account           string `json:"account,omitempty"`
	Hostname          string `json:"hostname,omitempty"`
	Projectid         string `json:"projectid,omitempty"`
	Publicipid        string `json:"publicipid,omitempty"`
	Instancename      string `json:"instancename,omitempty"`
	Rootdeviceid      int    `json:"rootdeviceid,omitempty"`
	Zoneid            string `json:"zoneid,omitempty"`
	Zonename          string `json:"zonename,omitempty"`
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

		var r RebootVirtualMachineResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type RebootVirtualMachineResponse struct {
	JobID               string `json:"jobid,omitempty"`
	Isodisplaytext      string `json:"isodisplaytext,omitempty"`
	Keypair             string `json:"keypair,omitempty"`
	Cpuused             string `json:"cpuused,omitempty"`
	Passwordenabled     bool   `json:"passwordenabled,omitempty"`
	Groupid             string `json:"groupid,omitempty"`
	Password            string `json:"password,omitempty"`
	Name                string `json:"name,omitempty"`
	Created             string `json:"created,omitempty"`
	Cpuspeed            int    `json:"cpuspeed,omitempty"`
	Hostid              string `json:"hostid,omitempty"`
	Isoname             string `json:"isoname,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Nic                 []struct {
		Id           string   `json:"id,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Type         string   `json:"type,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
	} `json:"nic,omitempty"`
	Displayvm             bool              `json:"displayvm,omitempty"`
	Domainid              string            `json:"domainid,omitempty"`
	Diskioread            int               `json:"diskioread,omitempty"`
	Rootdevicetype        string            `json:"rootdevicetype,omitempty"`
	Isoid                 string            `json:"isoid,omitempty"`
	Diskkbsread           int               `json:"diskkbsread,omitempty"`
	Hypervisor            string            `json:"hypervisor,omitempty"`
	Forvirtualnetwork     bool              `json:"forvirtualnetwork,omitempty"`
	Domain                string            `json:"domain,omitempty"`
	Publicipid            string            `json:"publicipid,omitempty"`
	Networkkbsread        int               `json:"networkkbsread,omitempty"`
	Details               map[string]string `json:"details,omitempty"`
	Group                 string            `json:"group,omitempty"`
	State                 string            `json:"state,omitempty"`
	Diskkbswrite          int               `json:"diskkbswrite,omitempty"`
	Isdynamicallyscalable bool              `json:"isdynamicallyscalable,omitempty"`
	Serviceofferingid     string            `json:"serviceofferingid,omitempty"`
	Project               string            `json:"project,omitempty"`
	Account               string            `json:"account,omitempty"`
	Affinitygroup         []struct {
		Domainid          string   `json:"domainid,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Account           string   `json:"account,omitempty"`
		Description       string   `json:"description,omitempty"`
		Id                string   `json:"id,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		Type              string   `json:"type,omitempty"`
		Name              string   `json:"name,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Instancename        string `json:"instancename,omitempty"`
	Diskiowrite         int    `json:"diskiowrite,omitempty"`
	Cpunumber           int    `json:"cpunumber,omitempty"`
	Servicestate        string `json:"servicestate,omitempty"`
	Templatedisplaytext string `json:"templatedisplaytext,omitempty"`
	Guestosid           string `json:"guestosid,omitempty"`
	Projectid           string `json:"projectid,omitempty"`
	Memory              int    `json:"memory,omitempty"`
	Templatename        string `json:"templatename,omitempty"`
	Haenable            bool   `json:"haenable,omitempty"`
	Rootdeviceid        int    `json:"rootdeviceid,omitempty"`
	Publicip            string `json:"publicip,omitempty"`
	Zoneid              string `json:"zoneid,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
	Templateid          string `json:"templateid,omitempty"`
	Id                  string `json:"id,omitempty"`
	Securitygroup       []struct {
		Project string `json:"project,omitempty"`
		Tags    []struct {
			Domainid     string `json:"domainid,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Project      string `json:"project,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
			Account      string `json:"account,omitempty"`
			Value        string `json:"value,omitempty"`
			Key          string `json:"key,omitempty"`
		} `json:"tags,omitempty"`
		Domain      string `json:"domain,omitempty"`
		Description string `json:"description,omitempty"`
		Egressrule  []struct {
			Account           string `json:"account,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
		} `json:"egressrule,omitempty"`
		Domainid    string `json:"domainid,omitempty"`
		Id          string `json:"id,omitempty"`
		Account     string `json:"account,omitempty"`
		Projectid   string `json:"projectid,omitempty"`
		Ingressrule []struct {
			Account           string `json:"account,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Endport           int    `json:"endport,omitempty"`
		} `json:"ingressrule,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"securitygroup,omitempty"`
	Networkkbswrite int `json:"networkkbswrite,omitempty"`
	Tags            []struct {
		Account      string `json:"account,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Key          string `json:"key,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Value        string `json:"value,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Project      string `json:"project,omitempty"`
	} `json:"tags,omitempty"`
	Displayname string `json:"displayname,omitempty"`
	Hostname    string `json:"hostname,omitempty"`
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

		var r StartVirtualMachineResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type StartVirtualMachineResponse struct {
	JobID             string `json:"jobid,omitempty"`
	Isoid             string `json:"isoid,omitempty"`
	Forvirtualnetwork bool   `json:"forvirtualnetwork,omitempty"`
	Tags              []struct {
		Customer     string `json:"customer,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Key          string `json:"key,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Project      string `json:"project,omitempty"`
		Value        string `json:"value,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Account      string `json:"account,omitempty"`
	} `json:"tags,omitempty"`
	Projectid           string            `json:"projectid,omitempty"`
	Networkkbsread      int               `json:"networkkbsread,omitempty"`
	Details             map[string]string `json:"details,omitempty"`
	Haenable            bool              `json:"haenable,omitempty"`
	Publicip            string            `json:"publicip,omitempty"`
	Servicestate        string            `json:"servicestate,omitempty"`
	Name                string            `json:"name,omitempty"`
	Memory              int               `json:"memory,omitempty"`
	Zonename            string            `json:"zonename,omitempty"`
	Displayname         string            `json:"displayname,omitempty"`
	Rootdevicetype      string            `json:"rootdevicetype,omitempty"`
	Diskkbswrite        int               `json:"diskkbswrite,omitempty"`
	Zoneid              string            `json:"zoneid,omitempty"`
	Domain              string            `json:"domain,omitempty"`
	State               string            `json:"state,omitempty"`
	Isodisplaytext      string            `json:"isodisplaytext,omitempty"`
	Diskiowrite         int               `json:"diskiowrite,omitempty"`
	Diskioread          int               `json:"diskioread,omitempty"`
	Serviceofferingname string            `json:"serviceofferingname,omitempty"`
	Securitygroup       []struct {
		Egressrule []struct {
			Icmptype          int    `json:"icmptype,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Account           string `json:"account,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
		} `json:"egressrule,omitempty"`
		Project string `json:"project,omitempty"`
		Tags    []struct {
			Projectid    string `json:"projectid,omitempty"`
			Value        string `json:"value,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Project      string `json:"project,omitempty"`
			Account      string `json:"account,omitempty"`
			Key          string `json:"key,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
		} `json:"tags,omitempty"`
		Account     string `json:"account,omitempty"`
		Description string `json:"description,omitempty"`
		Projectid   string `json:"projectid,omitempty"`
		Name        string `json:"name,omitempty"`
		Id          string `json:"id,omitempty"`
		Domain      string `json:"domain,omitempty"`
		Ingressrule []struct {
			Startport         int    `json:"startport,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Account           string `json:"account,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
		} `json:"ingressrule,omitempty"`
		Domainid string `json:"domainid,omitempty"`
	} `json:"securitygroup,omitempty"`
	Instancename        string `json:"instancename,omitempty"`
	Diskkbsread         int    `json:"diskkbsread,omitempty"`
	Account             string `json:"account,omitempty"`
	Templatedisplaytext string `json:"templatedisplaytext,omitempty"`
	Id                  string `json:"id,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
	Rootdeviceid        int    `json:"rootdeviceid,omitempty"`
	Hostname            string `json:"hostname,omitempty"`
	Guestosid           string `json:"guestosid,omitempty"`
	Cpuspeed            int    `json:"cpuspeed,omitempty"`
	Cpunumber           int    `json:"cpunumber,omitempty"`
	Templateid          string `json:"templateid,omitempty"`
	Groupid             string `json:"groupid,omitempty"`
	Cpuused             string `json:"cpuused,omitempty"`
	Hypervisor          string `json:"hypervisor,omitempty"`
	Group               string `json:"group,omitempty"`
	Nic                 []struct {
		Networkname  string   `json:"networkname,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Id           string   `json:"id,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Type         string   `json:"type,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
	} `json:"nic,omitempty"`
	Networkkbswrite       int    `json:"networkkbswrite,omitempty"`
	Isoname               string `json:"isoname,omitempty"`
	Templatename          string `json:"templatename,omitempty"`
	Isdynamicallyscalable bool   `json:"isdynamicallyscalable,omitempty"`
	Passwordenabled       bool   `json:"passwordenabled,omitempty"`
	Keypair               string `json:"keypair,omitempty"`
	Domainid              string `json:"domainid,omitempty"`
	Project               string `json:"project,omitempty"`
	Password              string `json:"password,omitempty"`
	Affinitygroup         []struct {
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Id                string   `json:"id,omitempty"`
		Type              string   `json:"type,omitempty"`
		Name              string   `json:"name,omitempty"`
		Description       string   `json:"description,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		Account           string   `json:"account,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Created    string `json:"created,omitempty"`
	Hostid     string `json:"hostid,omitempty"`
	Publicipid string `json:"publicipid,omitempty"`
	Displayvm  bool   `json:"displayvm,omitempty"`
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

		var r StopVirtualMachineResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type StopVirtualMachineResponse struct {
	JobID         string `json:"jobid,omitempty"`
	Securitygroup []struct {
		Id   string `json:"id,omitempty"`
		Tags []struct {
			Key          string `json:"key,omitempty"`
			Project      string `json:"project,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Account      string `json:"account,omitempty"`
			Value        string `json:"value,omitempty"`
		} `json:"tags,omitempty"`
		Name       string `json:"name,omitempty"`
		Account    string `json:"account,omitempty"`
		Egressrule []struct {
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Account           string `json:"account,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
		} `json:"egressrule,omitempty"`
		Domainid    string `json:"domainid,omitempty"`
		Project     string `json:"project,omitempty"`
		Projectid   string `json:"projectid,omitempty"`
		Description string `json:"description,omitempty"`
		Ingressrule []struct {
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Account           string `json:"account,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
		} `json:"ingressrule,omitempty"`
		Domain string `json:"domain,omitempty"`
	} `json:"securitygroup,omitempty"`
	Isoid                 string            `json:"isoid,omitempty"`
	Networkkbsread        int               `json:"networkkbsread,omitempty"`
	Diskkbswrite          int               `json:"diskkbswrite,omitempty"`
	Domain                string            `json:"domain,omitempty"`
	Diskioread            int               `json:"diskioread,omitempty"`
	Displayname           string            `json:"displayname,omitempty"`
	Serviceofferingid     string            `json:"serviceofferingid,omitempty"`
	Servicestate          string            `json:"servicestate,omitempty"`
	Cpuspeed              int               `json:"cpuspeed,omitempty"`
	Templatedisplaytext   string            `json:"templatedisplaytext,omitempty"`
	Templatename          string            `json:"templatename,omitempty"`
	Diskkbsread           int               `json:"diskkbsread,omitempty"`
	Instancename          string            `json:"instancename,omitempty"`
	Isodisplaytext        string            `json:"isodisplaytext,omitempty"`
	Templateid            string            `json:"templateid,omitempty"`
	Isoname               string            `json:"isoname,omitempty"`
	Publicipid            string            `json:"publicipid,omitempty"`
	Cpuused               string            `json:"cpuused,omitempty"`
	Id                    string            `json:"id,omitempty"`
	Diskiowrite           int               `json:"diskiowrite,omitempty"`
	Rootdevicetype        string            `json:"rootdevicetype,omitempty"`
	Groupid               string            `json:"groupid,omitempty"`
	State                 string            `json:"state,omitempty"`
	Projectid             string            `json:"projectid,omitempty"`
	Guestosid             string            `json:"guestosid,omitempty"`
	Isdynamicallyscalable bool              `json:"isdynamicallyscalable,omitempty"`
	Serviceofferingname   string            `json:"serviceofferingname,omitempty"`
	Zoneid                string            `json:"zoneid,omitempty"`
	Password              string            `json:"password,omitempty"`
	Haenable              bool              `json:"haenable,omitempty"`
	Hostid                string            `json:"hostid,omitempty"`
	Forvirtualnetwork     bool              `json:"forvirtualnetwork,omitempty"`
	Publicip              string            `json:"publicip,omitempty"`
	Details               map[string]string `json:"details,omitempty"`
	Group                 string            `json:"group,omitempty"`
	Name                  string            `json:"name,omitempty"`
	Hypervisor            string            `json:"hypervisor,omitempty"`
	Displayvm             bool              `json:"displayvm,omitempty"`
	Rootdeviceid          int               `json:"rootdeviceid,omitempty"`
	Tags                  []struct {
		Project      string `json:"project,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Account      string `json:"account,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Value        string `json:"value,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Key          string `json:"key,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
	} `json:"tags,omitempty"`
	Keypair  string `json:"keypair,omitempty"`
	Domainid string `json:"domainid,omitempty"`
	Zonename string `json:"zonename,omitempty"`
	Account  string `json:"account,omitempty"`
	Nic      []struct {
		Netmask      string   `json:"netmask,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Id           string   `json:"id,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Type         string   `json:"type,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
	} `json:"nic,omitempty"`
	Affinitygroup []struct {
		Account           string   `json:"account,omitempty"`
		Description       string   `json:"description,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		Name              string   `json:"name,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
		Id                string   `json:"id,omitempty"`
		Type              string   `json:"type,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Passwordenabled bool   `json:"passwordenabled,omitempty"`
	Project         string `json:"project,omitempty"`
	Networkkbswrite int    `json:"networkkbswrite,omitempty"`
	Hostname        string `json:"hostname,omitempty"`
	Memory          int    `json:"memory,omitempty"`
	Created         string `json:"created,omitempty"`
	Cpunumber       int    `json:"cpunumber,omitempty"`
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

		var r ResetPasswordForVirtualMachineResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type ResetPasswordForVirtualMachineResponse struct {
	JobID                 string `json:"jobid,omitempty"`
	Passwordenabled       bool   `json:"passwordenabled,omitempty"`
	Diskioread            int    `json:"diskioread,omitempty"`
	Serviceofferingid     string `json:"serviceofferingid,omitempty"`
	Group                 string `json:"group,omitempty"`
	Displayvm             bool   `json:"displayvm,omitempty"`
	Keypair               string `json:"keypair,omitempty"`
	Account               string `json:"account,omitempty"`
	Cpunumber             int    `json:"cpunumber,omitempty"`
	Isoname               string `json:"isoname,omitempty"`
	Instancename          string `json:"instancename,omitempty"`
	Cpuspeed              int    `json:"cpuspeed,omitempty"`
	Hostname              string `json:"hostname,omitempty"`
	Serviceofferingname   string `json:"serviceofferingname,omitempty"`
	Guestosid             string `json:"guestosid,omitempty"`
	Cpuused               string `json:"cpuused,omitempty"`
	State                 string `json:"state,omitempty"`
	Networkkbswrite       int    `json:"networkkbswrite,omitempty"`
	Isdynamicallyscalable bool   `json:"isdynamicallyscalable,omitempty"`
	Name                  string `json:"name,omitempty"`
	Templateid            string `json:"templateid,omitempty"`
	Nic                   []struct {
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Type         string   `json:"type,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Id           string   `json:"id,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
	} `json:"nic,omitempty"`
	Publicipid     string `json:"publicipid,omitempty"`
	Networkkbsread int    `json:"networkkbsread,omitempty"`
	Rootdevicetype string `json:"rootdevicetype,omitempty"`
	Haenable       bool   `json:"haenable,omitempty"`
	Groupid        string `json:"groupid,omitempty"`
	Diskkbsread    int    `json:"diskkbsread,omitempty"`
	Isodisplaytext string `json:"isodisplaytext,omitempty"`
	Isoid          string `json:"isoid,omitempty"`
	Displayname    string `json:"displayname,omitempty"`
	Securitygroup  []struct {
		Tags []struct {
			Value        string `json:"value,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Project      string `json:"project,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Key          string `json:"key,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Account      string `json:"account,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Customer     string `json:"customer,omitempty"`
		} `json:"tags,omitempty"`
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
		Id          string `json:"id,omitempty"`
		Projectid   string `json:"projectid,omitempty"`
		Description string `json:"description,omitempty"`
		Domain      string `json:"domain,omitempty"`
		Account     string `json:"account,omitempty"`
		Domainid    string `json:"domainid,omitempty"`
		Name        string `json:"name,omitempty"`
		Egressrule  []struct {
			Protocol          string `json:"protocol,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Account           string `json:"account,omitempty"`
		} `json:"egressrule,omitempty"`
		Project string `json:"project,omitempty"`
	} `json:"securitygroup,omitempty"`
	Memory              int               `json:"memory,omitempty"`
	Templatename        string            `json:"templatename,omitempty"`
	Password            string            `json:"password,omitempty"`
	Hostid              string            `json:"hostid,omitempty"`
	Zoneid              string            `json:"zoneid,omitempty"`
	Id                  string            `json:"id,omitempty"`
	Created             string            `json:"created,omitempty"`
	Forvirtualnetwork   bool              `json:"forvirtualnetwork,omitempty"`
	Details             map[string]string `json:"details,omitempty"`
	Diskiowrite         int               `json:"diskiowrite,omitempty"`
	Templatedisplaytext string            `json:"templatedisplaytext,omitempty"`
	Hypervisor          string            `json:"hypervisor,omitempty"`
	Tags                []struct {
		Domain       string `json:"domain,omitempty"`
		Key          string `json:"key,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Value        string `json:"value,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Project      string `json:"project,omitempty"`
		Account      string `json:"account,omitempty"`
	} `json:"tags,omitempty"`
	Diskkbswrite  int    `json:"diskkbswrite,omitempty"`
	Publicip      string `json:"publicip,omitempty"`
	Affinitygroup []struct {
		Id                string   `json:"id,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Account           string   `json:"account,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
		Name              string   `json:"name,omitempty"`
		Type              string   `json:"type,omitempty"`
		Description       string   `json:"description,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Domain       string `json:"domain,omitempty"`
	Rootdeviceid int    `json:"rootdeviceid,omitempty"`
	Zonename     string `json:"zonename,omitempty"`
	Domainid     string `json:"domainid,omitempty"`
	Servicestate string `json:"servicestate,omitempty"`
	Projectid    string `json:"projectid,omitempty"`
	Project      string `json:"project,omitempty"`
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
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Diskkbswrite        int    `json:"diskkbswrite,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
	Hostname            string `json:"hostname,omitempty"`
	Cpuused             string `json:"cpuused,omitempty"`
	Templatename        string `json:"templatename,omitempty"`
	Id                  string `json:"id,omitempty"`
	Nic                 []struct {
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Type         string   `json:"type,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Id           string   `json:"id,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
	} `json:"nic,omitempty"`
	Keypair           string `json:"keypair,omitempty"`
	Forvirtualnetwork bool   `json:"forvirtualnetwork,omitempty"`
	Hypervisor        string `json:"hypervisor,omitempty"`
	Rootdevicetype    string `json:"rootdevicetype,omitempty"`
	Servicestate      string `json:"servicestate,omitempty"`
	Guestosid         string `json:"guestosid,omitempty"`
	Password          string `json:"password,omitempty"`
	Account           string `json:"account,omitempty"`
	Securitygroup     []struct {
		Egressrule []struct {
			Account           string `json:"account,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
		} `json:"egressrule,omitempty"`
		Name      string `json:"name,omitempty"`
		Id        string `json:"id,omitempty"`
		Domainid  string `json:"domainid,omitempty"`
		Projectid string `json:"projectid,omitempty"`
		Domain    string `json:"domain,omitempty"`
		Tags      []struct {
			Resourcetype string `json:"resourcetype,omitempty"`
			Key          string `json:"key,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Value        string `json:"value,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Project      string `json:"project,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
			Account      string `json:"account,omitempty"`
		} `json:"tags,omitempty"`
		Project     string `json:"project,omitempty"`
		Ingressrule []struct {
			Startport         int    `json:"startport,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Account           string `json:"account,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
		} `json:"ingressrule,omitempty"`
		Description string `json:"description,omitempty"`
		Account     string `json:"account,omitempty"`
	} `json:"securitygroup,omitempty"`
	Displayvm             bool              `json:"displayvm,omitempty"`
	Group                 string            `json:"group,omitempty"`
	Zoneid                string            `json:"zoneid,omitempty"`
	Details               map[string]string `json:"details,omitempty"`
	Templatedisplaytext   string            `json:"templatedisplaytext,omitempty"`
	Projectid             string            `json:"projectid,omitempty"`
	Project               string            `json:"project,omitempty"`
	Diskiowrite           int               `json:"diskiowrite,omitempty"`
	Hostid                string            `json:"hostid,omitempty"`
	Displayname           string            `json:"displayname,omitempty"`
	Cpuspeed              int               `json:"cpuspeed,omitempty"`
	Name                  string            `json:"name,omitempty"`
	Cpunumber             int               `json:"cpunumber,omitempty"`
	Publicip              string            `json:"publicip,omitempty"`
	Instancename          string            `json:"instancename,omitempty"`
	Templateid            string            `json:"templateid,omitempty"`
	Memory                int               `json:"memory,omitempty"`
	Diskkbsread           int               `json:"diskkbsread,omitempty"`
	Publicipid            string            `json:"publicipid,omitempty"`
	Isdynamicallyscalable bool              `json:"isdynamicallyscalable,omitempty"`
	Domain                string            `json:"domain,omitempty"`
	Diskioread            int               `json:"diskioread,omitempty"`
	Zonename              string            `json:"zonename,omitempty"`
	Haenable              bool              `json:"haenable,omitempty"`
	Networkkbsread        int               `json:"networkkbsread,omitempty"`
	Isoname               string            `json:"isoname,omitempty"`
	Passwordenabled       bool              `json:"passwordenabled,omitempty"`
	Isodisplaytext        string            `json:"isodisplaytext,omitempty"`
	Rootdeviceid          int               `json:"rootdeviceid,omitempty"`
	Isoid                 string            `json:"isoid,omitempty"`
	Domainid              string            `json:"domainid,omitempty"`
	Groupid               string            `json:"groupid,omitempty"`
	Networkkbswrite       int               `json:"networkkbswrite,omitempty"`
	Affinitygroup         []struct {
		Type              string   `json:"type,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
		Account           string   `json:"account,omitempty"`
		Id                string   `json:"id,omitempty"`
		Description       string   `json:"description,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		Name              string   `json:"name,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Created string `json:"created,omitempty"`
	State   string `json:"state,omitempty"`
	Tags    []struct {
		Domainid     string `json:"domainid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Project      string `json:"project,omitempty"`
		Key          string `json:"key,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Account      string `json:"account,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Value        string `json:"value,omitempty"`
		Domain       string `json:"domain,omitempty"`
	} `json:"tags,omitempty"`
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
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, name, l)
	}
	return l.VirtualMachines[0].Id, nil
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
	Hostname          string `json:"hostname,omitempty"`
	Rootdevicetype    string `json:"rootdevicetype,omitempty"`
	Domain            string `json:"domain,omitempty"`
	Isodisplaytext    string `json:"isodisplaytext,omitempty"`
	Cpuspeed          int    `json:"cpuspeed,omitempty"`
	Rootdeviceid      int    `json:"rootdeviceid,omitempty"`
	Projectid         string `json:"projectid,omitempty"`
	Diskiowrite       int    `json:"diskiowrite,omitempty"`
	Displayname       string `json:"displayname,omitempty"`
	Instancename      string `json:"instancename,omitempty"`
	Project           string `json:"project,omitempty"`
	Displayvm         bool   `json:"displayvm,omitempty"`
	Templateid        string `json:"templateid,omitempty"`
	Groupid           string `json:"groupid,omitempty"`
	Id                string `json:"id,omitempty"`
	Publicipid        string `json:"publicipid,omitempty"`
	State             string `json:"state,omitempty"`
	Memory            int    `json:"memory,omitempty"`
	Servicestate      string `json:"servicestate,omitempty"`
	Publicip          string `json:"publicip,omitempty"`
	Name              string `json:"name,omitempty"`
	Hostid            string `json:"hostid,omitempty"`
	Serviceofferingid string `json:"serviceofferingid,omitempty"`
	Diskkbswrite      int    `json:"diskkbswrite,omitempty"`
	Diskkbsread       int    `json:"diskkbsread,omitempty"`
	Templatename      string `json:"templatename,omitempty"`
	Isoname           string `json:"isoname,omitempty"`
	Isoid             string `json:"isoid,omitempty"`
	Networkkbsread    int    `json:"networkkbsread,omitempty"`
	Group             string `json:"group,omitempty"`
	Diskioread        int    `json:"diskioread,omitempty"`
	Zonename          string `json:"zonename,omitempty"`
	Affinitygroup     []struct {
		Type              string   `json:"type,omitempty"`
		Account           string   `json:"account,omitempty"`
		Name              string   `json:"name,omitempty"`
		Description       string   `json:"description,omitempty"`
		Id                string   `json:"id,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Securitygroup []struct {
		Tags []struct {
			Projectid    string `json:"projectid,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Value        string `json:"value,omitempty"`
			Project      string `json:"project,omitempty"`
			Key          string `json:"key,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Account      string `json:"account,omitempty"`
		} `json:"tags,omitempty"`
		Egressrule []struct {
			Protocol          string `json:"protocol,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Account           string `json:"account,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
		} `json:"egressrule,omitempty"`
		Domain      string `json:"domain,omitempty"`
		Project     string `json:"project,omitempty"`
		Name        string `json:"name,omitempty"`
		Projectid   string `json:"projectid,omitempty"`
		Id          string `json:"id,omitempty"`
		Account     string `json:"account,omitempty"`
		Ingressrule []struct {
			Startport         int    `json:"startport,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Account           string `json:"account,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
		} `json:"ingressrule,omitempty"`
		Description string `json:"description,omitempty"`
		Domainid    string `json:"domainid,omitempty"`
	} `json:"securitygroup,omitempty"`
	Nic []struct {
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Id           string   `json:"id,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Type         string   `json:"type,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
	} `json:"nic,omitempty"`
	Cpuused             string            `json:"cpuused,omitempty"`
	Hypervisor          string            `json:"hypervisor,omitempty"`
	Networkkbswrite     int               `json:"networkkbswrite,omitempty"`
	Passwordenabled     bool              `json:"passwordenabled,omitempty"`
	Serviceofferingname string            `json:"serviceofferingname,omitempty"`
	Guestosid           string            `json:"guestosid,omitempty"`
	Domainid            string            `json:"domainid,omitempty"`
	Templatedisplaytext string            `json:"templatedisplaytext,omitempty"`
	Details             map[string]string `json:"details,omitempty"`
	Zoneid              string            `json:"zoneid,omitempty"`
	Password            string            `json:"password,omitempty"`
	Keypair             string            `json:"keypair,omitempty"`
	Tags                []struct {
		Customer     string `json:"customer,omitempty"`
		Account      string `json:"account,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Value        string `json:"value,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Project      string `json:"project,omitempty"`
		Key          string `json:"key,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
	} `json:"tags,omitempty"`
	Account               string `json:"account,omitempty"`
	Cpunumber             int    `json:"cpunumber,omitempty"`
	Created               string `json:"created,omitempty"`
	Haenable              bool   `json:"haenable,omitempty"`
	Isdynamicallyscalable bool   `json:"isdynamicallyscalable,omitempty"`
	Forvirtualnetwork     bool   `json:"forvirtualnetwork,omitempty"`
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

		var r RestoreVirtualMachineResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type RestoreVirtualMachineResponse struct {
	JobID             string            `json:"jobid,omitempty"`
	Passwordenabled   bool              `json:"passwordenabled,omitempty"`
	Rootdevicetype    string            `json:"rootdevicetype,omitempty"`
	Diskioread        int               `json:"diskioread,omitempty"`
	Forvirtualnetwork bool              `json:"forvirtualnetwork,omitempty"`
	Serviceofferingid string            `json:"serviceofferingid,omitempty"`
	Cpunumber         int               `json:"cpunumber,omitempty"`
	Domainid          string            `json:"domainid,omitempty"`
	Displayvm         bool              `json:"displayvm,omitempty"`
	Details           map[string]string `json:"details,omitempty"`
	Diskkbswrite      int               `json:"diskkbswrite,omitempty"`
	Project           string            `json:"project,omitempty"`
	Keypair           string            `json:"keypair,omitempty"`
	Hypervisor        string            `json:"hypervisor,omitempty"`
	Password          string            `json:"password,omitempty"`
	Securitygroup     []struct {
		Name        string `json:"name,omitempty"`
		Projectid   string `json:"projectid,omitempty"`
		Description string `json:"description,omitempty"`
		Domain      string `json:"domain,omitempty"`
		Account     string `json:"account,omitempty"`
		Tags        []struct {
			Project      string `json:"project,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Account      string `json:"account,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Value        string `json:"value,omitempty"`
			Key          string `json:"key,omitempty"`
			Customer     string `json:"customer,omitempty"`
		} `json:"tags,omitempty"`
		Egressrule []struct {
			Startport         int    `json:"startport,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Account           string `json:"account,omitempty"`
		} `json:"egressrule,omitempty"`
		Ingressrule []struct {
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Account           string `json:"account,omitempty"`
		} `json:"ingressrule,omitempty"`
		Domainid string `json:"domainid,omitempty"`
		Project  string `json:"project,omitempty"`
		Id       string `json:"id,omitempty"`
	} `json:"securitygroup,omitempty"`
	Name                  string `json:"name,omitempty"`
	Templatename          string `json:"templatename,omitempty"`
	Templateid            string `json:"templateid,omitempty"`
	Hostname              string `json:"hostname,omitempty"`
	Rootdeviceid          int    `json:"rootdeviceid,omitempty"`
	Templatedisplaytext   string `json:"templatedisplaytext,omitempty"`
	Zoneid                string `json:"zoneid,omitempty"`
	Diskkbsread           int    `json:"diskkbsread,omitempty"`
	Isoid                 string `json:"isoid,omitempty"`
	Account               string `json:"account,omitempty"`
	Diskiowrite           int    `json:"diskiowrite,omitempty"`
	Zonename              string `json:"zonename,omitempty"`
	Domain                string `json:"domain,omitempty"`
	Isoname               string `json:"isoname,omitempty"`
	Servicestate          string `json:"servicestate,omitempty"`
	Isdynamicallyscalable bool   `json:"isdynamicallyscalable,omitempty"`
	Created               string `json:"created,omitempty"`
	Id                    string `json:"id,omitempty"`
	Cpuspeed              int    `json:"cpuspeed,omitempty"`
	Memory                int    `json:"memory,omitempty"`
	Cpuused               string `json:"cpuused,omitempty"`
	Isodisplaytext        string `json:"isodisplaytext,omitempty"`
	State                 string `json:"state,omitempty"`
	Group                 string `json:"group,omitempty"`
	Instancename          string `json:"instancename,omitempty"`
	Affinitygroup         []struct {
		Account           string   `json:"account,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
		Id                string   `json:"id,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Name              string   `json:"name,omitempty"`
		Type              string   `json:"type,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		Description       string   `json:"description,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Networkkbsread      int    `json:"networkkbsread,omitempty"`
	Hostid              string `json:"hostid,omitempty"`
	Networkkbswrite     int    `json:"networkkbswrite,omitempty"`
	Groupid             string `json:"groupid,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Nic                 []struct {
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Type         string   `json:"type,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Id           string   `json:"id,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
	} `json:"nic,omitempty"`
	Projectid  string `json:"projectid,omitempty"`
	Publicipid string `json:"publicipid,omitempty"`
	Haenable   bool   `json:"haenable,omitempty"`
	Guestosid  string `json:"guestosid,omitempty"`
	Publicip   string `json:"publicip,omitempty"`
	Tags       []struct {
		Project      string `json:"project,omitempty"`
		Account      string `json:"account,omitempty"`
		Key          string `json:"key,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Value        string `json:"value,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Domain       string `json:"domain,omitempty"`
	} `json:"tags,omitempty"`
	Displayname string `json:"displayname,omitempty"`
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
	Templateid        string `json:"templateid,omitempty"`
	Cpuspeed          int    `json:"cpuspeed,omitempty"`
	Guestosid         string `json:"guestosid,omitempty"`
	Serviceofferingid string `json:"serviceofferingid,omitempty"`
	Hostid            string `json:"hostid,omitempty"`
	Diskkbsread       int    `json:"diskkbsread,omitempty"`
	Group             string `json:"group,omitempty"`
	Rootdeviceid      int    `json:"rootdeviceid,omitempty"`
	Securitygroup     []struct {
		Project    string `json:"project,omitempty"`
		Domainid   string `json:"domainid,omitempty"`
		Egressrule []struct {
			Account           string `json:"account,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
		} `json:"egressrule,omitempty"`
		Id   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Tags []struct {
			Domainid     string `json:"domainid,omitempty"`
			Project      string `json:"project,omitempty"`
			Account      string `json:"account,omitempty"`
			Key          string `json:"key,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Value        string `json:"value,omitempty"`
		} `json:"tags,omitempty"`
		Projectid   string `json:"projectid,omitempty"`
		Description string `json:"description,omitempty"`
		Ingressrule []struct {
			Protocol          string `json:"protocol,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Account           string `json:"account,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Startport         int    `json:"startport,omitempty"`
		} `json:"ingressrule,omitempty"`
		Domain  string `json:"domain,omitempty"`
		Account string `json:"account,omitempty"`
	} `json:"securitygroup,omitempty"`
	Zonename      string `json:"zonename,omitempty"`
	Project       string `json:"project,omitempty"`
	Affinitygroup []struct {
		Domain            string   `json:"domain,omitempty"`
		Name              string   `json:"name,omitempty"`
		Type              string   `json:"type,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
		Description       string   `json:"description,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Account           string   `json:"account,omitempty"`
		Id                string   `json:"id,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Isodisplaytext string            `json:"isodisplaytext,omitempty"`
	Password       string            `json:"password,omitempty"`
	Details        map[string]string `json:"details,omitempty"`
	Id             string            `json:"id,omitempty"`
	Diskkbswrite   int               `json:"diskkbswrite,omitempty"`
	Created        string            `json:"created,omitempty"`
	Nic            []struct {
		Netmask      string   `json:"netmask,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Id           string   `json:"id,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Type         string   `json:"type,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
	} `json:"nic,omitempty"`
	Templatename          string `json:"templatename,omitempty"`
	Haenable              bool   `json:"haenable,omitempty"`
	Serviceofferingname   string `json:"serviceofferingname,omitempty"`
	Templatedisplaytext   string `json:"templatedisplaytext,omitempty"`
	Zoneid                string `json:"zoneid,omitempty"`
	Account               string `json:"account,omitempty"`
	Domain                string `json:"domain,omitempty"`
	Servicestate          string `json:"servicestate,omitempty"`
	Isoid                 string `json:"isoid,omitempty"`
	Name                  string `json:"name,omitempty"`
	Displayname           string `json:"displayname,omitempty"`
	Projectid             string `json:"projectid,omitempty"`
	Domainid              string `json:"domainid,omitempty"`
	Forvirtualnetwork     bool   `json:"forvirtualnetwork,omitempty"`
	Hypervisor            string `json:"hypervisor,omitempty"`
	State                 string `json:"state,omitempty"`
	Isoname               string `json:"isoname,omitempty"`
	Instancename          string `json:"instancename,omitempty"`
	Cpunumber             int    `json:"cpunumber,omitempty"`
	Diskioread            int    `json:"diskioread,omitempty"`
	Diskiowrite           int    `json:"diskiowrite,omitempty"`
	Networkkbswrite       int    `json:"networkkbswrite,omitempty"`
	Isdynamicallyscalable bool   `json:"isdynamicallyscalable,omitempty"`
	Groupid               string `json:"groupid,omitempty"`
	Keypair               string `json:"keypair,omitempty"`
	Tags                  []struct {
		Value        string `json:"value,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Project      string `json:"project,omitempty"`
		Account      string `json:"account,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Key          string `json:"key,omitempty"`
		Customer     string `json:"customer,omitempty"`
	} `json:"tags,omitempty"`
	Passwordenabled bool   `json:"passwordenabled,omitempty"`
	Memory          int    `json:"memory,omitempty"`
	Networkkbsread  int    `json:"networkkbsread,omitempty"`
	Publicipid      string `json:"publicipid,omitempty"`
	Cpuused         string `json:"cpuused,omitempty"`
	Displayvm       bool   `json:"displayvm,omitempty"`
	Rootdevicetype  string `json:"rootdevicetype,omitempty"`
	Hostname        string `json:"hostname,omitempty"`
	Publicip        string `json:"publicip,omitempty"`
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

		var r ScaleVirtualMachineResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type ScaleVirtualMachineResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
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
	Cpunumber           int    `json:"cpunumber,omitempty"`
	Project             string `json:"project,omitempty"`
	Name                string `json:"name,omitempty"`
	Passwordenabled     bool   `json:"passwordenabled,omitempty"`
	Networkkbsread      int    `json:"networkkbsread,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
	Created             string `json:"created,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Isodisplaytext      string `json:"isodisplaytext,omitempty"`
	Domainid            string `json:"domainid,omitempty"`
	Guestosid           string `json:"guestosid,omitempty"`
	Projectid           string `json:"projectid,omitempty"`
	Publicipid          string `json:"publicipid,omitempty"`
	Displayname         string `json:"displayname,omitempty"`
	Tags                []struct {
		Domainid     string `json:"domainid,omitempty"`
		Value        string `json:"value,omitempty"`
		Key          string `json:"key,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Project      string `json:"project,omitempty"`
		Account      string `json:"account,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
	} `json:"tags,omitempty"`
	Displayvm      bool              `json:"displayvm,omitempty"`
	Servicestate   string            `json:"servicestate,omitempty"`
	Account        string            `json:"account,omitempty"`
	Diskiowrite    int               `json:"diskiowrite,omitempty"`
	State          string            `json:"state,omitempty"`
	Details        map[string]string `json:"details,omitempty"`
	Diskkbsread    int               `json:"diskkbsread,omitempty"`
	Rootdevicetype string            `json:"rootdevicetype,omitempty"`
	Securitygroup  []struct {
		Domainid string `json:"domainid,omitempty"`
		Domain   string `json:"domain,omitempty"`
		Tags     []struct {
			Account      string `json:"account,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Key          string `json:"key,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Project      string `json:"project,omitempty"`
			Value        string `json:"value,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
		} `json:"tags,omitempty"`
		Projectid   string `json:"projectid,omitempty"`
		Ingressrule []struct {
			Cidr              string `json:"cidr,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Account           string `json:"account,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Startport         int    `json:"startport,omitempty"`
		} `json:"ingressrule,omitempty"`
		Account    string `json:"account,omitempty"`
		Id         string `json:"id,omitempty"`
		Egressrule []struct {
			Endport           int    `json:"endport,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Account           string `json:"account,omitempty"`
		} `json:"egressrule,omitempty"`
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
		Project     string `json:"project,omitempty"`
	} `json:"securitygroup,omitempty"`
	Keypair               string `json:"keypair,omitempty"`
	Isdynamicallyscalable bool   `json:"isdynamicallyscalable,omitempty"`
	Haenable              bool   `json:"haenable,omitempty"`
	Networkkbswrite       int    `json:"networkkbswrite,omitempty"`
	Hypervisor            string `json:"hypervisor,omitempty"`
	Rootdeviceid          int    `json:"rootdeviceid,omitempty"`
	Password              string `json:"password,omitempty"`
	Publicip              string `json:"publicip,omitempty"`
	Cpuspeed              int    `json:"cpuspeed,omitempty"`
	Isoname               string `json:"isoname,omitempty"`
	Hostname              string `json:"hostname,omitempty"`
	Hostid                string `json:"hostid,omitempty"`
	Group                 string `json:"group,omitempty"`
	Cpuused               string `json:"cpuused,omitempty"`
	Affinitygroup         []struct {
		Description       string   `json:"description,omitempty"`
		Id                string   `json:"id,omitempty"`
		Name              string   `json:"name,omitempty"`
		Account           string   `json:"account,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Type              string   `json:"type,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Isoid               string `json:"isoid,omitempty"`
	Templatename        string `json:"templatename,omitempty"`
	Templatedisplaytext string `json:"templatedisplaytext,omitempty"`
	Diskioread          int    `json:"diskioread,omitempty"`
	Instancename        string `json:"instancename,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
	Forvirtualnetwork   bool   `json:"forvirtualnetwork,omitempty"`
	Id                  string `json:"id,omitempty"`
	Memory              int    `json:"memory,omitempty"`
	Templateid          string `json:"templateid,omitempty"`
	Groupid             string `json:"groupid,omitempty"`
	Nic                 []struct {
		Isdefault    bool     `json:"isdefault,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Id           string   `json:"id,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Type         string   `json:"type,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
	} `json:"nic,omitempty"`
	Zoneid       string `json:"zoneid,omitempty"`
	Diskkbswrite int    `json:"diskkbswrite,omitempty"`
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

		var r MigrateVirtualMachineResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type MigrateVirtualMachineResponse struct {
	JobID               string `json:"jobid,omitempty"`
	Hostid              string `json:"hostid,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Affinitygroup       []struct {
		Account           string   `json:"account,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Type              string   `json:"type,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
		Description       string   `json:"description,omitempty"`
		Id                string   `json:"id,omitempty"`
		Name              string   `json:"name,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Group         string `json:"group,omitempty"`
	Securitygroup []struct {
		Tags []struct {
			Projectid    string `json:"projectid,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Project      string `json:"project,omitempty"`
			Account      string `json:"account,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Value        string `json:"value,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Key          string `json:"key,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
		} `json:"tags,omitempty"`
		Ingressrule []struct {
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Account           string `json:"account,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
		} `json:"ingressrule,omitempty"`
		Project     string `json:"project,omitempty"`
		Description string `json:"description,omitempty"`
		Id          string `json:"id,omitempty"`
		Name        string `json:"name,omitempty"`
		Domain      string `json:"domain,omitempty"`
		Egressrule  []struct {
			Endport           int    `json:"endport,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Account           string `json:"account,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
		} `json:"egressrule,omitempty"`
		Account   string `json:"account,omitempty"`
		Domainid  string `json:"domainid,omitempty"`
		Projectid string `json:"projectid,omitempty"`
	} `json:"securitygroup,omitempty"`
	Servicestate      string `json:"servicestate,omitempty"`
	Isodisplaytext    string `json:"isodisplaytext,omitempty"`
	Diskiowrite       int    `json:"diskiowrite,omitempty"`
	Isoid             string `json:"isoid,omitempty"`
	Hostname          string `json:"hostname,omitempty"`
	Projectid         string `json:"projectid,omitempty"`
	Forvirtualnetwork bool   `json:"forvirtualnetwork,omitempty"`
	Id                string `json:"id,omitempty"`
	Created           string `json:"created,omitempty"`
	Haenable          bool   `json:"haenable,omitempty"`
	Serviceofferingid string `json:"serviceofferingid,omitempty"`
	Keypair           string `json:"keypair,omitempty"`
	Cpunumber         int    `json:"cpunumber,omitempty"`
	Networkkbswrite   int    `json:"networkkbswrite,omitempty"`
	Tags              []struct {
		Resourcetype string `json:"resourcetype,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Key          string `json:"key,omitempty"`
		Project      string `json:"project,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Account      string `json:"account,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Value        string `json:"value,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
	} `json:"tags,omitempty"`
	Templateid     string `json:"templateid,omitempty"`
	Diskkbsread    int    `json:"diskkbsread,omitempty"`
	Domainid       string `json:"domainid,omitempty"`
	Zonename       string `json:"zonename,omitempty"`
	Cpuspeed       int    `json:"cpuspeed,omitempty"`
	Rootdevicetype string `json:"rootdevicetype,omitempty"`
	Account        string `json:"account,omitempty"`
	Displayname    string `json:"displayname,omitempty"`
	Name           string `json:"name,omitempty"`
	Password       string `json:"password,omitempty"`
	Nic            []struct {
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Id           string   `json:"id,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Type         string   `json:"type,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
	} `json:"nic,omitempty"`
	Guestosid             string            `json:"guestosid,omitempty"`
	Publicipid            string            `json:"publicipid,omitempty"`
	Instancename          string            `json:"instancename,omitempty"`
	Zoneid                string            `json:"zoneid,omitempty"`
	Details               map[string]string `json:"details,omitempty"`
	Networkkbsread        int               `json:"networkkbsread,omitempty"`
	Diskioread            int               `json:"diskioread,omitempty"`
	Templatedisplaytext   string            `json:"templatedisplaytext,omitempty"`
	Groupid               string            `json:"groupid,omitempty"`
	Passwordenabled       bool              `json:"passwordenabled,omitempty"`
	Publicip              string            `json:"publicip,omitempty"`
	Templatename          string            `json:"templatename,omitempty"`
	Diskkbswrite          int               `json:"diskkbswrite,omitempty"`
	Displayvm             bool              `json:"displayvm,omitempty"`
	Isdynamicallyscalable bool              `json:"isdynamicallyscalable,omitempty"`
	Memory                int               `json:"memory,omitempty"`
	Project               string            `json:"project,omitempty"`
	Cpuused               string            `json:"cpuused,omitempty"`
	Rootdeviceid          int               `json:"rootdeviceid,omitempty"`
	Hypervisor            string            `json:"hypervisor,omitempty"`
	State                 string            `json:"state,omitempty"`
	Isoname               string            `json:"isoname,omitempty"`
	Domain                string            `json:"domain,omitempty"`
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

		var r MigrateVirtualMachineWithVolumeResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type MigrateVirtualMachineWithVolumeResponse struct {
	JobID               string `json:"jobid,omitempty"`
	Account             string `json:"account,omitempty"`
	Networkkbsread      int    `json:"networkkbsread,omitempty"`
	Groupid             string `json:"groupid,omitempty"`
	Hypervisor          string `json:"hypervisor,omitempty"`
	Diskioread          int    `json:"diskioread,omitempty"`
	Servicestate        string `json:"servicestate,omitempty"`
	Memory              int    `json:"memory,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Templatename        string `json:"templatename,omitempty"`
	Publicipid          string `json:"publicipid,omitempty"`
	State               string `json:"state,omitempty"`
	Diskiowrite         int    `json:"diskiowrite,omitempty"`
	Hostid              string `json:"hostid,omitempty"`
	Affinitygroup       []struct {
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Type              string   `json:"type,omitempty"`
		Description       string   `json:"description,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		Account           string   `json:"account,omitempty"`
		Name              string   `json:"name,omitempty"`
		Id                string   `json:"id,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Serviceofferingid     string `json:"serviceofferingid,omitempty"`
	Created               string `json:"created,omitempty"`
	Haenable              bool   `json:"haenable,omitempty"`
	Rootdevicetype        string `json:"rootdevicetype,omitempty"`
	Diskkbsread           int    `json:"diskkbsread,omitempty"`
	Cpunumber             int    `json:"cpunumber,omitempty"`
	Isodisplaytext        string `json:"isodisplaytext,omitempty"`
	Diskkbswrite          int    `json:"diskkbswrite,omitempty"`
	Project               string `json:"project,omitempty"`
	Forvirtualnetwork     bool   `json:"forvirtualnetwork,omitempty"`
	Publicip              string `json:"publicip,omitempty"`
	Cpuused               string `json:"cpuused,omitempty"`
	Displayvm             bool   `json:"displayvm,omitempty"`
	Isdynamicallyscalable bool   `json:"isdynamicallyscalable,omitempty"`
	Projectid             string `json:"projectid,omitempty"`
	Rootdeviceid          int    `json:"rootdeviceid,omitempty"`
	Id                    string `json:"id,omitempty"`
	Securitygroup         []struct {
		Project     string `json:"project,omitempty"`
		Projectid   string `json:"projectid,omitempty"`
		Ingressrule []struct {
			Account           string `json:"account,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Startport         int    `json:"startport,omitempty"`
		} `json:"ingressrule,omitempty"`
		Account    string `json:"account,omitempty"`
		Egressrule []struct {
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Account           string `json:"account,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
		} `json:"egressrule,omitempty"`
		Description string `json:"description,omitempty"`
		Name        string `json:"name,omitempty"`
		Id          string `json:"id,omitempty"`
		Tags        []struct {
			Domainid     string `json:"domainid,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Account      string `json:"account,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Value        string `json:"value,omitempty"`
			Key          string `json:"key,omitempty"`
			Project      string `json:"project,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
		} `json:"tags,omitempty"`
		Domain   string `json:"domain,omitempty"`
		Domainid string `json:"domainid,omitempty"`
	} `json:"securitygroup,omitempty"`
	Instancename string `json:"instancename,omitempty"`
	Isoname      string `json:"isoname,omitempty"`
	Zonename     string `json:"zonename,omitempty"`
	Tags         []struct {
		Resourcetype string `json:"resourcetype,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Value        string `json:"value,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Key          string `json:"key,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Account      string `json:"account,omitempty"`
		Project      string `json:"project,omitempty"`
	} `json:"tags,omitempty"`
	Cpuspeed  int               `json:"cpuspeed,omitempty"`
	Password  string            `json:"password,omitempty"`
	Zoneid    string            `json:"zoneid,omitempty"`
	Details   map[string]string `json:"details,omitempty"`
	Guestosid string            `json:"guestosid,omitempty"`
	Name      string            `json:"name,omitempty"`
	Keypair   string            `json:"keypair,omitempty"`
	Domainid  string            `json:"domainid,omitempty"`
	Nic       []struct {
		Ip6address   string   `json:"ip6address,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Id           string   `json:"id,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Type         string   `json:"type,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
	} `json:"nic,omitempty"`
	Passwordenabled     bool   `json:"passwordenabled,omitempty"`
	Group               string `json:"group,omitempty"`
	Hostname            string `json:"hostname,omitempty"`
	Networkkbswrite     int    `json:"networkkbswrite,omitempty"`
	Isoid               string `json:"isoid,omitempty"`
	Displayname         string `json:"displayname,omitempty"`
	Templateid          string `json:"templateid,omitempty"`
	Templatedisplaytext string `json:"templatedisplaytext,omitempty"`
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
	Isdynamicallyscalable bool   `json:"isdynamicallyscalable,omitempty"`
	Group                 string `json:"group,omitempty"`
	Rootdeviceid          int    `json:"rootdeviceid,omitempty"`
	State                 string `json:"state,omitempty"`
	Nic                   []struct {
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Type         string   `json:"type,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Id           string   `json:"id,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
	} `json:"nic,omitempty"`
	Instancename        string            `json:"instancename,omitempty"`
	Templatedisplaytext string            `json:"templatedisplaytext,omitempty"`
	Details             map[string]string `json:"details,omitempty"`
	Serviceofferingname string            `json:"serviceofferingname,omitempty"`
	Diskiowrite         int               `json:"diskiowrite,omitempty"`
	Hypervisor          string            `json:"hypervisor,omitempty"`
	Password            string            `json:"password,omitempty"`
	Displayvm           bool              `json:"displayvm,omitempty"`
	Diskkbswrite        int               `json:"diskkbswrite,omitempty"`
	Created             string            `json:"created,omitempty"`
	Memory              int               `json:"memory,omitempty"`
	Isoid               string            `json:"isoid,omitempty"`
	Diskioread          int               `json:"diskioread,omitempty"`
	Id                  string            `json:"id,omitempty"`
	Project             string            `json:"project,omitempty"`
	Passwordenabled     bool              `json:"passwordenabled,omitempty"`
	Serviceofferingid   string            `json:"serviceofferingid,omitempty"`
	Haenable            bool              `json:"haenable,omitempty"`
	Servicestate        string            `json:"servicestate,omitempty"`
	Affinitygroup       []struct {
		Id                string   `json:"id,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		Account           string   `json:"account,omitempty"`
		Name              string   `json:"name,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
		Description       string   `json:"description,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Type              string   `json:"type,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Rootdevicetype    string `json:"rootdevicetype,omitempty"`
	Isoname           string `json:"isoname,omitempty"`
	Isodisplaytext    string `json:"isodisplaytext,omitempty"`
	Networkkbsread    int    `json:"networkkbsread,omitempty"`
	Hostname          string `json:"hostname,omitempty"`
	Templatename      string `json:"templatename,omitempty"`
	Projectid         string `json:"projectid,omitempty"`
	Domainid          string `json:"domainid,omitempty"`
	Zonename          string `json:"zonename,omitempty"`
	Forvirtualnetwork bool   `json:"forvirtualnetwork,omitempty"`
	Securitygroup     []struct {
		Name       string `json:"name,omitempty"`
		Domain     string `json:"domain,omitempty"`
		Egressrule []struct {
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Account           string `json:"account,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
		} `json:"egressrule,omitempty"`
		Tags []struct {
			Domainid     string `json:"domainid,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Project      string `json:"project,omitempty"`
			Value        string `json:"value,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Key          string `json:"key,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Account      string `json:"account,omitempty"`
		} `json:"tags,omitempty"`
		Project     string `json:"project,omitempty"`
		Account     string `json:"account,omitempty"`
		Ingressrule []struct {
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Account           string `json:"account,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
		} `json:"ingressrule,omitempty"`
		Id          string `json:"id,omitempty"`
		Domainid    string `json:"domainid,omitempty"`
		Projectid   string `json:"projectid,omitempty"`
		Description string `json:"description,omitempty"`
	} `json:"securitygroup,omitempty"`
	Cpuspeed    int    `json:"cpuspeed,omitempty"`
	Zoneid      string `json:"zoneid,omitempty"`
	Templateid  string `json:"templateid,omitempty"`
	Publicip    string `json:"publicip,omitempty"`
	Diskkbsread int    `json:"diskkbsread,omitempty"`
	Hostid      string `json:"hostid,omitempty"`
	Guestosid   string `json:"guestosid,omitempty"`
	Displayname string `json:"displayname,omitempty"`
	Name        string `json:"name,omitempty"`
	Domain      string `json:"domain,omitempty"`
	Account     string `json:"account,omitempty"`
	Groupid     string `json:"groupid,omitempty"`
	Publicipid  string `json:"publicipid,omitempty"`
	Cpunumber   int    `json:"cpunumber,omitempty"`
	Tags        []struct {
		Domain       string `json:"domain,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Value        string `json:"value,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Account      string `json:"account,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Project      string `json:"project,omitempty"`
		Key          string `json:"key,omitempty"`
	} `json:"tags,omitempty"`
	Keypair         string `json:"keypair,omitempty"`
	Cpuused         string `json:"cpuused,omitempty"`
	Networkkbswrite int    `json:"networkkbswrite,omitempty"`
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

		var r ExpungeVirtualMachineResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
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

		var r CleanVMReservationsResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
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

		var r AddNicToVirtualMachineResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type AddNicToVirtualMachineResponse struct {
	JobID               string `json:"jobid,omitempty"`
	Templatedisplaytext string `json:"templatedisplaytext,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Groupid             string `json:"groupid,omitempty"`
	Networkkbsread      int    `json:"networkkbsread,omitempty"`
	Zoneid              string `json:"zoneid,omitempty"`
	Cpuspeed            int    `json:"cpuspeed,omitempty"`
	Hostname            string `json:"hostname,omitempty"`
	Id                  string `json:"id,omitempty"`
	Securitygroup       []struct {
		Egressrule []struct {
			Endport           int    `json:"endport,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Account           string `json:"account,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
		} `json:"egressrule,omitempty"`
		Id          string `json:"id,omitempty"`
		Domain      string `json:"domain,omitempty"`
		Name        string `json:"name,omitempty"`
		Domainid    string `json:"domainid,omitempty"`
		Description string `json:"description,omitempty"`
		Ingressrule []struct {
			Endport           int    `json:"endport,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Account           string `json:"account,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
		} `json:"ingressrule,omitempty"`
		Projectid string `json:"projectid,omitempty"`
		Account   string `json:"account,omitempty"`
		Project   string `json:"project,omitempty"`
		Tags      []struct {
			Projectid    string `json:"projectid,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Key          string `json:"key,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Project      string `json:"project,omitempty"`
			Account      string `json:"account,omitempty"`
			Value        string `json:"value,omitempty"`
		} `json:"tags,omitempty"`
	} `json:"securitygroup,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Nic                 []struct {
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Id           string   `json:"id,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Type         string   `json:"type,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
	} `json:"nic,omitempty"`
	Account       string `json:"account,omitempty"`
	Domainid      string `json:"domainid,omitempty"`
	Memory        int    `json:"memory,omitempty"`
	Affinitygroup []struct {
		Domainid          string   `json:"domainid,omitempty"`
		Type              string   `json:"type,omitempty"`
		Account           string   `json:"account,omitempty"`
		Name              string   `json:"name,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		Description       string   `json:"description,omitempty"`
		Id                string   `json:"id,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Password          string `json:"password,omitempty"`
	Hostid            string `json:"hostid,omitempty"`
	Diskkbsread       int    `json:"diskkbsread,omitempty"`
	Created           string `json:"created,omitempty"`
	Networkkbswrite   int    `json:"networkkbswrite,omitempty"`
	Group             string `json:"group,omitempty"`
	Rootdeviceid      int    `json:"rootdeviceid,omitempty"`
	Forvirtualnetwork bool   `json:"forvirtualnetwork,omitempty"`
	Publicip          string `json:"publicip,omitempty"`
	Passwordenabled   bool   `json:"passwordenabled,omitempty"`
	Diskioread        int    `json:"diskioread,omitempty"`
	Tags              []struct {
		Domain       string `json:"domain,omitempty"`
		Key          string `json:"key,omitempty"`
		Account      string `json:"account,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Project      string `json:"project,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Value        string `json:"value,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
	} `json:"tags,omitempty"`
	Instancename          string            `json:"instancename,omitempty"`
	Keypair               string            `json:"keypair,omitempty"`
	Details               map[string]string `json:"details,omitempty"`
	Isoid                 string            `json:"isoid,omitempty"`
	Isodisplaytext        string            `json:"isodisplaytext,omitempty"`
	Cpuused               string            `json:"cpuused,omitempty"`
	Haenable              bool              `json:"haenable,omitempty"`
	Serviceofferingid     string            `json:"serviceofferingid,omitempty"`
	Diskkbswrite          int               `json:"diskkbswrite,omitempty"`
	Templateid            string            `json:"templateid,omitempty"`
	Displayname           string            `json:"displayname,omitempty"`
	State                 string            `json:"state,omitempty"`
	Publicipid            string            `json:"publicipid,omitempty"`
	Displayvm             bool              `json:"displayvm,omitempty"`
	Rootdevicetype        string            `json:"rootdevicetype,omitempty"`
	Servicestate          string            `json:"servicestate,omitempty"`
	Diskiowrite           int               `json:"diskiowrite,omitempty"`
	Projectid             string            `json:"projectid,omitempty"`
	Hypervisor            string            `json:"hypervisor,omitempty"`
	Cpunumber             int               `json:"cpunumber,omitempty"`
	Name                  string            `json:"name,omitempty"`
	Isoname               string            `json:"isoname,omitempty"`
	Project               string            `json:"project,omitempty"`
	Templatename          string            `json:"templatename,omitempty"`
	Zonename              string            `json:"zonename,omitempty"`
	Guestosid             string            `json:"guestosid,omitempty"`
	Isdynamicallyscalable bool              `json:"isdynamicallyscalable,omitempty"`
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

		var r RemoveNicFromVirtualMachineResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type RemoveNicFromVirtualMachineResponse struct {
	JobID               string            `json:"jobid,omitempty"`
	Forvirtualnetwork   bool              `json:"forvirtualnetwork,omitempty"`
	Isoid               string            `json:"isoid,omitempty"`
	Id                  string            `json:"id,omitempty"`
	Diskkbsread         int               `json:"diskkbsread,omitempty"`
	Passwordenabled     bool              `json:"passwordenabled,omitempty"`
	Projectid           string            `json:"projectid,omitempty"`
	Domainid            string            `json:"domainid,omitempty"`
	Domain              string            `json:"domain,omitempty"`
	Details             map[string]string `json:"details,omitempty"`
	Templatename        string            `json:"templatename,omitempty"`
	Displayname         string            `json:"displayname,omitempty"`
	Publicip            string            `json:"publicip,omitempty"`
	Templatedisplaytext string            `json:"templatedisplaytext,omitempty"`
	Instancename        string            `json:"instancename,omitempty"`
	Cpuspeed            int               `json:"cpuspeed,omitempty"`
	Diskioread          int               `json:"diskioread,omitempty"`
	State               string            `json:"state,omitempty"`
	Cpuused             string            `json:"cpuused,omitempty"`
	Serviceofferingid   string            `json:"serviceofferingid,omitempty"`
	Cpunumber           int               `json:"cpunumber,omitempty"`
	Networkkbswrite     int               `json:"networkkbswrite,omitempty"`
	Group               string            `json:"group,omitempty"`
	Isodisplaytext      string            `json:"isodisplaytext,omitempty"`
	Created             string            `json:"created,omitempty"`
	Isoname             string            `json:"isoname,omitempty"`
	Tags                []struct {
		Domain       string `json:"domain,omitempty"`
		Account      string `json:"account,omitempty"`
		Project      string `json:"project,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Key          string `json:"key,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Value        string `json:"value,omitempty"`
	} `json:"tags,omitempty"`
	Displayvm           bool   `json:"displayvm,omitempty"`
	Servicestate        string `json:"servicestate,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Diskiowrite         int    `json:"diskiowrite,omitempty"`
	Nic                 []struct {
		Ip6address   string   `json:"ip6address,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Id           string   `json:"id,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Type         string   `json:"type,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
	} `json:"nic,omitempty"`
	Rootdevicetype string `json:"rootdevicetype,omitempty"`
	Guestosid      string `json:"guestosid,omitempty"`
	Haenable       bool   `json:"haenable,omitempty"`
	Zonename       string `json:"zonename,omitempty"`
	Affinitygroup  []struct {
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		Description       string   `json:"description,omitempty"`
		Id                string   `json:"id,omitempty"`
		Type              string   `json:"type,omitempty"`
		Account           string   `json:"account,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
		Name              string   `json:"name,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Groupid        string `json:"groupid,omitempty"`
	Templateid     string `json:"templateid,omitempty"`
	Keypair        string `json:"keypair,omitempty"`
	Networkkbsread int    `json:"networkkbsread,omitempty"`
	Publicipid     string `json:"publicipid,omitempty"`
	Password       string `json:"password,omitempty"`
	Rootdeviceid   int    `json:"rootdeviceid,omitempty"`
	Hypervisor     string `json:"hypervisor,omitempty"`
	Securitygroup  []struct {
		Egressrule []struct {
			Account           string `json:"account,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Endport           int    `json:"endport,omitempty"`
		} `json:"egressrule,omitempty"`
		Ingressrule []struct {
			Account           string `json:"account,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
		} `json:"ingressrule,omitempty"`
		Name    string `json:"name,omitempty"`
		Domain  string `json:"domain,omitempty"`
		Project string `json:"project,omitempty"`
		Tags    []struct {
			Projectid    string `json:"projectid,omitempty"`
			Key          string `json:"key,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Project      string `json:"project,omitempty"`
			Value        string `json:"value,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Account      string `json:"account,omitempty"`
		} `json:"tags,omitempty"`
		Account     string `json:"account,omitempty"`
		Description string `json:"description,omitempty"`
		Domainid    string `json:"domainid,omitempty"`
		Projectid   string `json:"projectid,omitempty"`
		Id          string `json:"id,omitempty"`
	} `json:"securitygroup,omitempty"`
	Name                  string `json:"name,omitempty"`
	Hostid                string `json:"hostid,omitempty"`
	Project               string `json:"project,omitempty"`
	Account               string `json:"account,omitempty"`
	Diskkbswrite          int    `json:"diskkbswrite,omitempty"`
	Hostname              string `json:"hostname,omitempty"`
	Isdynamicallyscalable bool   `json:"isdynamicallyscalable,omitempty"`
	Zoneid                string `json:"zoneid,omitempty"`
	Memory                int    `json:"memory,omitempty"`
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

		var r UpdateDefaultNicForVirtualMachineResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type UpdateDefaultNicForVirtualMachineResponse struct {
	JobID           string `json:"jobid,omitempty"`
	Memory          int    `json:"memory,omitempty"`
	Cpuused         string `json:"cpuused,omitempty"`
	Networkkbsread  int    `json:"networkkbsread,omitempty"`
	Networkkbswrite int    `json:"networkkbswrite,omitempty"`
	Hostid          string `json:"hostid,omitempty"`
	State           string `json:"state,omitempty"`
	Passwordenabled bool   `json:"passwordenabled,omitempty"`
	Rootdevicetype  string `json:"rootdevicetype,omitempty"`
	Nic             []struct {
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Type         string   `json:"type,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Id           string   `json:"id,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
	} `json:"nic,omitempty"`
	Guestosid         string            `json:"guestosid,omitempty"`
	Isodisplaytext    string            `json:"isodisplaytext,omitempty"`
	Domainid          string            `json:"domainid,omitempty"`
	Zoneid            string            `json:"zoneid,omitempty"`
	Account           string            `json:"account,omitempty"`
	Hostname          string            `json:"hostname,omitempty"`
	Forvirtualnetwork bool              `json:"forvirtualnetwork,omitempty"`
	Hypervisor        string            `json:"hypervisor,omitempty"`
	Rootdeviceid      int               `json:"rootdeviceid,omitempty"`
	Instancename      string            `json:"instancename,omitempty"`
	Password          string            `json:"password,omitempty"`
	Group             string            `json:"group,omitempty"`
	Name              string            `json:"name,omitempty"`
	Haenable          bool              `json:"haenable,omitempty"`
	Details           map[string]string `json:"details,omitempty"`
	Created           string            `json:"created,omitempty"`
	Templateid        string            `json:"templateid,omitempty"`
	Domain            string            `json:"domain,omitempty"`
	Servicestate      string            `json:"servicestate,omitempty"`
	Cpuspeed          int               `json:"cpuspeed,omitempty"`
	Id                string            `json:"id,omitempty"`
	Tags              []struct {
		Resourcetype string `json:"resourcetype,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Project      string `json:"project,omitempty"`
		Value        string `json:"value,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Key          string `json:"key,omitempty"`
		Account      string `json:"account,omitempty"`
	} `json:"tags,omitempty"`
	Isoid                 string `json:"isoid,omitempty"`
	Templatename          string `json:"templatename,omitempty"`
	Diskkbswrite          int    `json:"diskkbswrite,omitempty"`
	Projectid             string `json:"projectid,omitempty"`
	Diskkbsread           int    `json:"diskkbsread,omitempty"`
	Groupid               string `json:"groupid,omitempty"`
	Templatedisplaytext   string `json:"templatedisplaytext,omitempty"`
	Displayname           string `json:"displayname,omitempty"`
	Serviceofferingid     string `json:"serviceofferingid,omitempty"`
	Cpunumber             int    `json:"cpunumber,omitempty"`
	Isdynamicallyscalable bool   `json:"isdynamicallyscalable,omitempty"`
	Securitygroup         []struct {
		Domainid    string `json:"domainid,omitempty"`
		Ingressrule []struct {
			Endport           int    `json:"endport,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Account           string `json:"account,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
		} `json:"ingressrule,omitempty"`
		Project   string `json:"project,omitempty"`
		Projectid string `json:"projectid,omitempty"`
		Tags      []struct {
			Key          string `json:"key,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
			Project      string `json:"project,omitempty"`
			Value        string `json:"value,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Account      string `json:"account,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
		} `json:"tags,omitempty"`
		Description string `json:"description,omitempty"`
		Egressrule  []struct {
			Cidr              string `json:"cidr,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Account           string `json:"account,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
		} `json:"egressrule,omitempty"`
		Domain  string `json:"domain,omitempty"`
		Account string `json:"account,omitempty"`
		Id      string `json:"id,omitempty"`
		Name    string `json:"name,omitempty"`
	} `json:"securitygroup,omitempty"`
	Keypair             string `json:"keypair,omitempty"`
	Project             string `json:"project,omitempty"`
	Isoname             string `json:"isoname,omitempty"`
	Publicipid          string `json:"publicipid,omitempty"`
	Displayvm           bool   `json:"displayvm,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Diskiowrite         int    `json:"diskiowrite,omitempty"`
	Diskioread          int    `json:"diskioread,omitempty"`
	Affinitygroup       []struct {
		Name              string   `json:"name,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Description       string   `json:"description,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		Id                string   `json:"id,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
		Account           string   `json:"account,omitempty"`
		Type              string   `json:"type,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Zonename string `json:"zonename,omitempty"`
	Publicip string `json:"publicip,omitempty"`
}
