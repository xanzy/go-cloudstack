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
	JobID               string `json:"jobid,omitempty"`
	Forvirtualnetwork   bool   `json:"forvirtualnetwork,omitempty"`
	Password            string `json:"password,omitempty"`
	Haenable            bool   `json:"haenable,omitempty"`
	Templatedisplaytext string `json:"templatedisplaytext,omitempty"`
	Memory              int    `json:"memory,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
	Project             string `json:"project,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
	Projectid           string `json:"projectid,omitempty"`
	Isoname             string `json:"isoname,omitempty"`
	Affinitygroup       []struct {
		Domain            string   `json:"domain,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Name              string   `json:"name,omitempty"`
		Description       string   `json:"description,omitempty"`
		Account           string   `json:"account,omitempty"`
		Type              string   `json:"type,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
		Id                string   `json:"id,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Publicipid   string `json:"publicipid,omitempty"`
	Cpuused      string `json:"cpuused,omitempty"`
	Created      string `json:"created,omitempty"`
	Account      string `json:"account,omitempty"`
	Rootdeviceid int    `json:"rootdeviceid,omitempty"`
	Isoid        string `json:"isoid,omitempty"`
	Diskkbsread  int    `json:"diskkbsread,omitempty"`
	Id           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Keypair      string `json:"keypair,omitempty"`
	Hostid       string `json:"hostid,omitempty"`
	Guestosid    string `json:"guestosid,omitempty"`
	Instancename string `json:"instancename,omitempty"`
	Nic          []struct {
		Ip6address   string   `json:"ip6address,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Type         string   `json:"type,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Id           string   `json:"id,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
	} `json:"nic,omitempty"`
	Displayvm             bool   `json:"displayvm,omitempty"`
	Zoneid                string `json:"zoneid,omitempty"`
	Groupid               string `json:"groupid,omitempty"`
	Isdynamicallyscalable bool   `json:"isdynamicallyscalable,omitempty"`
	Displayname           string `json:"displayname,omitempty"`
	Isodisplaytext        string `json:"isodisplaytext,omitempty"`
	Cpunumber             int    `json:"cpunumber,omitempty"`
	Diskkbswrite          int    `json:"diskkbswrite,omitempty"`
	Templatename          string `json:"templatename,omitempty"`
	Rootdevicetype        string `json:"rootdevicetype,omitempty"`
	Securitygroup         []struct {
		Ingressrule []struct {
			Protocol          string `json:"protocol,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Account           string `json:"account,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
		} `json:"ingressrule,omitempty"`
		Projectid   string `json:"projectid,omitempty"`
		Description string `json:"description,omitempty"`
		Egressrule  []struct {
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Account           string `json:"account,omitempty"`
		} `json:"egressrule,omitempty"`
		Domainid string `json:"domainid,omitempty"`
		Account  string `json:"account,omitempty"`
		Name     string `json:"name,omitempty"`
		Project  string `json:"project,omitempty"`
		Tags     []struct {
			Projectid    string `json:"projectid,omitempty"`
			Key          string `json:"key,omitempty"`
			Account      string `json:"account,omitempty"`
			Project      string `json:"project,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Value        string `json:"value,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Customer     string `json:"customer,omitempty"`
		} `json:"tags,omitempty"`
		Domain string `json:"domain,omitempty"`
		Id     string `json:"id,omitempty"`
	} `json:"securitygroup,omitempty"`
	Hostname string `json:"hostname,omitempty"`
	Tags     []struct {
		Domainid     string `json:"domainid,omitempty"`
		Value        string `json:"value,omitempty"`
		Key          string `json:"key,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Project      string `json:"project,omitempty"`
		Account      string `json:"account,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Domain       string `json:"domain,omitempty"`
	} `json:"tags,omitempty"`
	Publicip        string            `json:"publicip,omitempty"`
	State           string            `json:"state,omitempty"`
	Diskioread      int               `json:"diskioread,omitempty"`
	Details         map[string]string `json:"details,omitempty"`
	Networkkbswrite int               `json:"networkkbswrite,omitempty"`
	Passwordenabled bool              `json:"passwordenabled,omitempty"`
	Domain          string            `json:"domain,omitempty"`
	Cpuspeed        int               `json:"cpuspeed,omitempty"`
	Domainid        string            `json:"domainid,omitempty"`
	Networkkbsread  int               `json:"networkkbsread,omitempty"`
	Hypervisor      string            `json:"hypervisor,omitempty"`
	Diskiowrite     int               `json:"diskiowrite,omitempty"`
	Group           string            `json:"group,omitempty"`
	Templateid      string            `json:"templateid,omitempty"`
	Servicestate    string            `json:"servicestate,omitempty"`
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
	Name        string `json:"name,omitempty"`
	Fingerprint string `json:"fingerprint,omitempty"`
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
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
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
