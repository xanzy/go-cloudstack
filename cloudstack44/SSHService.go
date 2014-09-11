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
	"net/url"
	"strconv"
)

type ResetSSHKeyForVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *ResetSSHKeyForVirtualMachineParams) toURLValues() url.Values {
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
	if v, found := p.p["keypair"]; found {
		u.Set("keypair", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	return u
}

func (p *ResetSSHKeyForVirtualMachineParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ResetSSHKeyForVirtualMachineParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ResetSSHKeyForVirtualMachineParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ResetSSHKeyForVirtualMachineParams) SetKeypair(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keypair"] = v
	return
}

func (p *ResetSSHKeyForVirtualMachineParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

// You should always use this function to get a new ResetSSHKeyForVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *SSHService) NewResetSSHKeyForVirtualMachineParams(id string, keypair string) *ResetSSHKeyForVirtualMachineParams {
	p := &ResetSSHKeyForVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["keypair"] = keypair
	return p
}

// Resets the SSH Key for virtual machine. The virtual machine must be in a "Stopped" state. [async]
func (s *SSHService) ResetSSHKeyForVirtualMachine(p *ResetSSHKeyForVirtualMachineParams) (*ResetSSHKeyForVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("resetSSHKeyForVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ResetSSHKeyForVirtualMachineResponse
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

		var r ResetSSHKeyForVirtualMachineResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type ResetSSHKeyForVirtualMachineResponse struct {
	JobID                 string            `json:"jobid,omitempty"`
	Isoname               string            `json:"isoname,omitempty"`
	Hypervisor            string            `json:"hypervisor,omitempty"`
	Hostname              string            `json:"hostname,omitempty"`
	Created               string            `json:"created,omitempty"`
	Cpunumber             int               `json:"cpunumber,omitempty"`
	Projectid             string            `json:"projectid,omitempty"`
	Networkkbswrite       int               `json:"networkkbswrite,omitempty"`
	Account               string            `json:"account,omitempty"`
	Serviceofferingid     string            `json:"serviceofferingid,omitempty"`
	Diskkbswrite          int               `json:"diskkbswrite,omitempty"`
	Isdynamicallyscalable bool              `json:"isdynamicallyscalable,omitempty"`
	Rootdevicetype        string            `json:"rootdevicetype,omitempty"`
	Details               map[string]string `json:"details,omitempty"`
	Rootdeviceid          int               `json:"rootdeviceid,omitempty"`
	Isoid                 string            `json:"isoid,omitempty"`
	Templateid            string            `json:"templateid,omitempty"`
	Cpuused               string            `json:"cpuused,omitempty"`
	Project               string            `json:"project,omitempty"`
	Displayvm             bool              `json:"displayvm,omitempty"`
	Passwordenabled       bool              `json:"passwordenabled,omitempty"`
	Domainid              string            `json:"domainid,omitempty"`
	Cpuspeed              int               `json:"cpuspeed,omitempty"`
	Diskioread            int               `json:"diskioread,omitempty"`
	Tags                  []struct {
		Account      string `json:"account,omitempty"`
		Value        string `json:"value,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Project      string `json:"project,omitempty"`
		Key          string `json:"key,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Domain       string `json:"domain,omitempty"`
	} `json:"tags,omitempty"`
	Instancename      string `json:"instancename,omitempty"`
	Id                string `json:"id,omitempty"`
	Forvirtualnetwork bool   `json:"forvirtualnetwork,omitempty"`
	Keypair           string `json:"keypair,omitempty"`
	Securitygroup     []struct {
		Id          string `json:"id,omitempty"`
		Projectid   string `json:"projectid,omitempty"`
		Ingressrule []struct {
			Startport         int    `json:"startport,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Account           string `json:"account,omitempty"`
		} `json:"ingressrule,omitempty"`
		Name        string `json:"name,omitempty"`
		Account     string `json:"account,omitempty"`
		Description string `json:"description,omitempty"`
		Project     string `json:"project,omitempty"`
		Egressrule  []struct {
			Ruleid            string `json:"ruleid,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Account           string `json:"account,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
		} `json:"egressrule,omitempty"`
		Tags []struct {
			Projectid    string `json:"projectid,omitempty"`
			Value        string `json:"value,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Account      string `json:"account,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Key          string `json:"key,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Project      string `json:"project,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
		} `json:"tags,omitempty"`
		Domain   string `json:"domain,omitempty"`
		Domainid string `json:"domainid,omitempty"`
	} `json:"securitygroup,omitempty"`
	Publicip   string `json:"publicip,omitempty"`
	Publicipid string `json:"publicipid,omitempty"`
	Name       string `json:"name,omitempty"`
	Zoneid     string `json:"zoneid,omitempty"`
	Guestosid  string `json:"guestosid,omitempty"`
	Group      string `json:"group,omitempty"`
	Nic        []struct {
		Gateway      string   `json:"gateway,omitempty"`
		Type         string   `json:"type,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Id           string   `json:"id,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
	} `json:"nic,omitempty"`
	Networkkbsread      int    `json:"networkkbsread,omitempty"`
	Password            string `json:"password,omitempty"`
	Isodisplaytext      string `json:"isodisplaytext,omitempty"`
	Groupid             string `json:"groupid,omitempty"`
	Haenable            bool   `json:"haenable,omitempty"`
	Hostid              string `json:"hostid,omitempty"`
	State               string `json:"state,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Templatedisplaytext string `json:"templatedisplaytext,omitempty"`
	Displayname         string `json:"displayname,omitempty"`
	Templatename        string `json:"templatename,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
	Affinitygroup       []struct {
		Account           string   `json:"account,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		Description       string   `json:"description,omitempty"`
		Name              string   `json:"name,omitempty"`
		Type              string   `json:"type,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Id                string   `json:"id,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Diskiowrite  int    `json:"diskiowrite,omitempty"`
	Domain       string `json:"domain,omitempty"`
	Servicestate string `json:"servicestate,omitempty"`
	Diskkbsread  int    `json:"diskkbsread,omitempty"`
	Memory       int    `json:"memory,omitempty"`
}

type RegisterSSHKeyPairParams struct {
	p map[string]interface{}
}

func (p *RegisterSSHKeyPairParams) toURLValues() url.Values {
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
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["publickey"]; found {
		u.Set("publickey", v.(string))
	}
	return u
}

func (p *RegisterSSHKeyPairParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *RegisterSSHKeyPairParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *RegisterSSHKeyPairParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *RegisterSSHKeyPairParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *RegisterSSHKeyPairParams) SetPublickey(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["publickey"] = v
	return
}

// You should always use this function to get a new RegisterSSHKeyPairParams instance,
// as then you are sure you have configured all required params
func (s *SSHService) NewRegisterSSHKeyPairParams(name string, publickey string) *RegisterSSHKeyPairParams {
	p := &RegisterSSHKeyPairParams{}
	p.p = make(map[string]interface{})
	p.p["name"] = name
	p.p["publickey"] = publickey
	return p
}

// Register a public key in a keypair under a certain name
func (s *SSHService) RegisterSSHKeyPair(p *RegisterSSHKeyPairParams) (*RegisterSSHKeyPairResponse, error) {
	resp, err := s.cs.newRequest("registerSSHKeyPair", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RegisterSSHKeyPairResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type RegisterSSHKeyPairResponse struct {
	Fingerprint string `json:"fingerprint,omitempty"`
	Name        string `json:"name,omitempty"`
}

type CreateSSHKeyPairParams struct {
	p map[string]interface{}
}

func (p *CreateSSHKeyPairParams) toURLValues() url.Values {
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
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	return u
}

func (p *CreateSSHKeyPairParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *CreateSSHKeyPairParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *CreateSSHKeyPairParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *CreateSSHKeyPairParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

// You should always use this function to get a new CreateSSHKeyPairParams instance,
// as then you are sure you have configured all required params
func (s *SSHService) NewCreateSSHKeyPairParams(name string) *CreateSSHKeyPairParams {
	p := &CreateSSHKeyPairParams{}
	p.p = make(map[string]interface{})
	p.p["name"] = name
	return p
}

// Create a new keypair and returns the private key
func (s *SSHService) CreateSSHKeyPair(p *CreateSSHKeyPairParams) (*CreateSSHKeyPairResponse, error) {
	resp, err := s.cs.newRequest("createSSHKeyPair", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateSSHKeyPairResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type CreateSSHKeyPairResponse struct {
	Privatekey string `json:"privatekey,omitempty"`
}

type DeleteSSHKeyPairParams struct {
	p map[string]interface{}
}

func (p *DeleteSSHKeyPairParams) toURLValues() url.Values {
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
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	return u
}

func (p *DeleteSSHKeyPairParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *DeleteSSHKeyPairParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *DeleteSSHKeyPairParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *DeleteSSHKeyPairParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

// You should always use this function to get a new DeleteSSHKeyPairParams instance,
// as then you are sure you have configured all required params
func (s *SSHService) NewDeleteSSHKeyPairParams(name string) *DeleteSSHKeyPairParams {
	p := &DeleteSSHKeyPairParams{}
	p.p = make(map[string]interface{})
	p.p["name"] = name
	return p
}

// Deletes a keypair by name
func (s *SSHService) DeleteSSHKeyPair(p *DeleteSSHKeyPairParams) (*DeleteSSHKeyPairResponse, error) {
	resp, err := s.cs.newRequest("deleteSSHKeyPair", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteSSHKeyPairResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type DeleteSSHKeyPairResponse struct {
	Success     bool   `json:"success,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
}

type ListSSHKeyPairsParams struct {
	p map[string]interface{}
}

func (p *ListSSHKeyPairsParams) toURLValues() url.Values {
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
	if v, found := p.p["fingerprint"]; found {
		u.Set("fingerprint", v.(string))
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
	return u
}

func (p *ListSSHKeyPairsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListSSHKeyPairsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListSSHKeyPairsParams) SetFingerprint(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fingerprint"] = v
	return
}

func (p *ListSSHKeyPairsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
	return
}

func (p *ListSSHKeyPairsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListSSHKeyPairsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
	return
}

func (p *ListSSHKeyPairsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *ListSSHKeyPairsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListSSHKeyPairsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListSSHKeyPairsParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

// You should always use this function to get a new ListSSHKeyPairsParams instance,
// as then you are sure you have configured all required params
func (s *SSHService) NewListSSHKeyPairsParams() *ListSSHKeyPairsParams {
	p := &ListSSHKeyPairsParams{}
	p.p = make(map[string]interface{})
	return p
}

// List registered keypairs
func (s *SSHService) ListSSHKeyPairs(p *ListSSHKeyPairsParams) (*ListSSHKeyPairsResponse, error) {
	resp, err := s.cs.newRequest("listSSHKeyPairs", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListSSHKeyPairsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListSSHKeyPairsResponse struct {
	Count       int           `json:"count"`
	SSHKeyPairs []*SSHKeyPair `json:"sshkeypair"`
}

type SSHKeyPair struct {
	Fingerprint string `json:"fingerprint,omitempty"`
	Name        string `json:"name,omitempty"`
}
