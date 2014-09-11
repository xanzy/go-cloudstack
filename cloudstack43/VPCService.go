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

type CreateVPCParams struct {
	p map[string]interface{}
}

func (p *CreateVPCParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["cidr"]; found {
		u.Set("cidr", v.(string))
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
	if v, found := p.p["networkdomain"]; found {
		u.Set("networkdomain", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["start"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("start", vv)
	}
	if v, found := p.p["vpcofferingid"]; found {
		u.Set("vpcofferingid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *CreateVPCParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *CreateVPCParams) SetCidr(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cidr"] = v
	return
}

func (p *CreateVPCParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
	return
}

func (p *CreateVPCParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *CreateVPCParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *CreateVPCParams) SetNetworkdomain(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkdomain"] = v
	return
}

func (p *CreateVPCParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *CreateVPCParams) SetStart(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["start"] = v
	return
}

func (p *CreateVPCParams) SetVpcofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcofferingid"] = v
	return
}

func (p *CreateVPCParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new CreateVPCParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewCreateVPCParams(cidr string, displaytext string, name string, vpcofferingid string, zoneid string) *CreateVPCParams {
	p := &CreateVPCParams{}
	p.p = make(map[string]interface{})
	p.p["cidr"] = cidr
	p.p["displaytext"] = displaytext
	p.p["name"] = name
	p.p["vpcofferingid"] = vpcofferingid
	p.p["zoneid"] = zoneid
	return p
}

// Creates a VPC
func (s *VPCService) CreateVPC(p *CreateVPCParams) (*CreateVPCResponse, error) {
	resp, err := s.cs.newRequest("createVPC", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateVPCResponse
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

		var r CreateVPCResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type CreateVPCResponse struct {
	JobID  string `json:"jobid,omitempty"`
	Zoneid string `json:"zoneid,omitempty"`
	Tags   []struct {
		Resourceid   string `json:"resourceid,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Account      string `json:"account,omitempty"`
		Key          string `json:"key,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Project      string `json:"project,omitempty"`
		Value        string `json:"value,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
	} `json:"tags,omitempty"`
	Restartrequired bool   `json:"restartrequired,omitempty"`
	Id              string `json:"id,omitempty"`
	Displaytext     string `json:"displaytext,omitempty"`
	Projectid       string `json:"projectid,omitempty"`
	Networkdomain   string `json:"networkdomain,omitempty"`
	Project         string `json:"project,omitempty"`
	Domain          string `json:"domain,omitempty"`
	Name            string `json:"name,omitempty"`
	Vpcofferingid   string `json:"vpcofferingid,omitempty"`
	State           string `json:"state,omitempty"`
	Created         string `json:"created,omitempty"`
	Service         []struct {
		Name     string `json:"name,omitempty"`
		Provider []struct {
			Name                         string   `json:"name,omitempty"`
			Canenableindividualservice   bool     `json:"canenableindividualservice,omitempty"`
			Servicelist                  []string `json:"servicelist,omitempty"`
			State                        string   `json:"state,omitempty"`
			Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid,omitempty"`
			Physicalnetworkid            string   `json:"physicalnetworkid,omitempty"`
			Id                           string   `json:"id,omitempty"`
		} `json:"provider,omitempty"`
		Capability []struct {
			Canchooseservicecapability bool   `json:"canchooseservicecapability,omitempty"`
			Value                      string `json:"value,omitempty"`
			Name                       string `json:"name,omitempty"`
		} `json:"capability,omitempty"`
	} `json:"service,omitempty"`
	Network []struct {
		Reservediprange   string `json:"reservediprange,omitempty"`
		Networkcidr       string `json:"networkcidr,omitempty"`
		Id                string `json:"id,omitempty"`
		Vlan              string `json:"vlan,omitempty"`
		Zonename          string `json:"zonename,omitempty"`
		Ispersistent      bool   `json:"ispersistent,omitempty"`
		Acltype           string `json:"acltype,omitempty"`
		Physicalnetworkid string `json:"physicalnetworkid,omitempty"`
		Ip6gateway        string `json:"ip6gateway,omitempty"`
		Tags              []struct {
			Resourceid   string `json:"resourceid,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Value        string `json:"value,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Project      string `json:"project,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Account      string `json:"account,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Key          string `json:"key,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
		} `json:"tags,omitempty"`
		Account                     string `json:"account,omitempty"`
		Displaynetwork              bool   `json:"displaynetwork,omitempty"`
		Canusefordeploy             bool   `json:"canusefordeploy,omitempty"`
		Projectid                   string `json:"projectid,omitempty"`
		Networkofferingdisplaytext  string `json:"networkofferingdisplaytext,omitempty"`
		Gateway                     string `json:"gateway,omitempty"`
		Isdefault                   bool   `json:"isdefault,omitempty"`
		Ip6cidr                     string `json:"ip6cidr,omitempty"`
		Restartrequired             bool   `json:"restartrequired,omitempty"`
		Networkdomain               string `json:"networkdomain,omitempty"`
		Aclid                       string `json:"aclid,omitempty"`
		Broadcasturi                string `json:"broadcasturi,omitempty"`
		Subdomainaccess             bool   `json:"subdomainaccess,omitempty"`
		Issystem                    bool   `json:"issystem,omitempty"`
		Displaytext                 string `json:"displaytext,omitempty"`
		Vpcid                       string `json:"vpcid,omitempty"`
		Zoneid                      string `json:"zoneid,omitempty"`
		Networkofferingname         string `json:"networkofferingname,omitempty"`
		Related                     string `json:"related,omitempty"`
		State                       string `json:"state,omitempty"`
		Specifyipranges             bool   `json:"specifyipranges,omitempty"`
		Networkofferingid           string `json:"networkofferingid,omitempty"`
		Netmask                     string `json:"netmask,omitempty"`
		Traffictype                 string `json:"traffictype,omitempty"`
		Networkofferingconservemode bool   `json:"networkofferingconservemode,omitempty"`
		Name                        string `json:"name,omitempty"`
		Domainid                    string `json:"domainid,omitempty"`
		Broadcastdomaintype         string `json:"broadcastdomaintype,omitempty"`
		Dns2                        string `json:"dns2,omitempty"`
		Dns1                        string `json:"dns1,omitempty"`
		Cidr                        string `json:"cidr,omitempty"`
		Service                     []struct {
			Name     string `json:"name,omitempty"`
			Provider []struct {
				State                        string   `json:"state,omitempty"`
				Physicalnetworkid            string   `json:"physicalnetworkid,omitempty"`
				Canenableindividualservice   bool     `json:"canenableindividualservice,omitempty"`
				Name                         string   `json:"name,omitempty"`
				Servicelist                  []string `json:"servicelist,omitempty"`
				Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid,omitempty"`
				Id                           string   `json:"id,omitempty"`
			} `json:"provider,omitempty"`
			Capability []struct {
				Name                       string `json:"name,omitempty"`
				Value                      string `json:"value,omitempty"`
				Canchooseservicecapability bool   `json:"canchooseservicecapability,omitempty"`
			} `json:"capability,omitempty"`
		} `json:"service,omitempty"`
		Domain                      string `json:"domain,omitempty"`
		Networkofferingavailability string `json:"networkofferingavailability,omitempty"`
		Type                        string `json:"type,omitempty"`
		Project                     string `json:"project,omitempty"`
	} `json:"network,omitempty"`
	Zonename string `json:"zonename,omitempty"`
	Domainid string `json:"domainid,omitempty"`
	Cidr     string `json:"cidr,omitempty"`
	Account  string `json:"account,omitempty"`
}

type ListVPCsParams struct {
	p map[string]interface{}
}

func (p *ListVPCsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["cidr"]; found {
		u.Set("cidr", v.(string))
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
	if v, found := p.p["restartrequired"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("restartrequired", vv)
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := p.p["supportedservices"]; found {
		vv := strings.Join(v.([]string), ", ")
		u.Set("supportedservices", vv)
	}
	if v, found := p.p["tags"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["vpcofferingid"]; found {
		u.Set("vpcofferingid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListVPCsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListVPCsParams) SetCidr(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cidr"] = v
	return
}

func (p *ListVPCsParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
	return
}

func (p *ListVPCsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListVPCsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListVPCsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
	return
}

func (p *ListVPCsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListVPCsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
	return
}

func (p *ListVPCsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *ListVPCsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListVPCsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListVPCsParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *ListVPCsParams) SetRestartrequired(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["restartrequired"] = v
	return
}

func (p *ListVPCsParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
	return
}

func (p *ListVPCsParams) SetSupportedservices(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["supportedservices"] = v
	return
}

func (p *ListVPCsParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
	return
}

func (p *ListVPCsParams) SetVpcofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcofferingid"] = v
	return
}

func (p *ListVPCsParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new ListVPCsParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewListVPCsParams() *ListVPCsParams {
	p := &ListVPCsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VPCService) GetVPCID(name string) (string, error) {
	p := &ListVPCsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	l, err := s.ListVPCs(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, name, l)
	}
	return l.VPCs[0].Id, nil
}

// Lists VPCs
func (s *VPCService) ListVPCs(p *ListVPCsParams) (*ListVPCsResponse, error) {
	resp, err := s.cs.newRequest("listVPCs", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVPCsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListVPCsResponse struct {
	Count int    `json:"count"`
	VPCs  []*VPC `json:"vpc"`
}

type VPC struct {
	Id      string `json:"id,omitempty"`
	Service []struct {
		Provider []struct {
			Name                         string   `json:"name,omitempty"`
			Id                           string   `json:"id,omitempty"`
			Canenableindividualservice   bool     `json:"canenableindividualservice,omitempty"`
			Servicelist                  []string `json:"servicelist,omitempty"`
			Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid,omitempty"`
			State                        string   `json:"state,omitempty"`
			Physicalnetworkid            string   `json:"physicalnetworkid,omitempty"`
		} `json:"provider,omitempty"`
		Capability []struct {
			Name                       string `json:"name,omitempty"`
			Value                      string `json:"value,omitempty"`
			Canchooseservicecapability bool   `json:"canchooseservicecapability,omitempty"`
		} `json:"capability,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"service,omitempty"`
	Domain  string `json:"domain,omitempty"`
	Cidr    string `json:"cidr,omitempty"`
	Created string `json:"created,omitempty"`
	Tags    []struct {
		Customer     string `json:"customer,omitempty"`
		Project      string `json:"project,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Key          string `json:"key,omitempty"`
		Account      string `json:"account,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Value        string `json:"value,omitempty"`
	} `json:"tags,omitempty"`
	Zoneid        string `json:"zoneid,omitempty"`
	Name          string `json:"name,omitempty"`
	Displaytext   string `json:"displaytext,omitempty"`
	Project       string `json:"project,omitempty"`
	Projectid     string `json:"projectid,omitempty"`
	Domainid      string `json:"domainid,omitempty"`
	Zonename      string `json:"zonename,omitempty"`
	Networkdomain string `json:"networkdomain,omitempty"`
	Network       []struct {
		Domain                      string `json:"domain,omitempty"`
		State                       string `json:"state,omitempty"`
		Networkdomain               string `json:"networkdomain,omitempty"`
		Vlan                        string `json:"vlan,omitempty"`
		Id                          string `json:"id,omitempty"`
		Networkofferingavailability string `json:"networkofferingavailability,omitempty"`
		Displaytext                 string `json:"displaytext,omitempty"`
		Networkofferingname         string `json:"networkofferingname,omitempty"`
		Physicalnetworkid           string `json:"physicalnetworkid,omitempty"`
		Traffictype                 string `json:"traffictype,omitempty"`
		Tags                        []struct {
			Value        string `json:"value,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
			Account      string `json:"account,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Project      string `json:"project,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Key          string `json:"key,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
		} `json:"tags,omitempty"`
		Specifyipranges             bool   `json:"specifyipranges,omitempty"`
		Reservediprange             string `json:"reservediprange,omitempty"`
		Broadcasturi                string `json:"broadcasturi,omitempty"`
		Dns1                        string `json:"dns1,omitempty"`
		Zonename                    string `json:"zonename,omitempty"`
		Vpcid                       string `json:"vpcid,omitempty"`
		Related                     string `json:"related,omitempty"`
		Type                        string `json:"type,omitempty"`
		Networkofferingconservemode bool   `json:"networkofferingconservemode,omitempty"`
		Restartrequired             bool   `json:"restartrequired,omitempty"`
		Networkofferingdisplaytext  string `json:"networkofferingdisplaytext,omitempty"`
		Canusefordeploy             bool   `json:"canusefordeploy,omitempty"`
		Networkofferingid           string `json:"networkofferingid,omitempty"`
		Networkcidr                 string `json:"networkcidr,omitempty"`
		Broadcastdomaintype         string `json:"broadcastdomaintype,omitempty"`
		Domainid                    string `json:"domainid,omitempty"`
		Netmask                     string `json:"netmask,omitempty"`
		Dns2                        string `json:"dns2,omitempty"`
		Projectid                   string `json:"projectid,omitempty"`
		Cidr                        string `json:"cidr,omitempty"`
		Ip6cidr                     string `json:"ip6cidr,omitempty"`
		Name                        string `json:"name,omitempty"`
		Service                     []struct {
			Capability []struct {
				Value                      string `json:"value,omitempty"`
				Canchooseservicecapability bool   `json:"canchooseservicecapability,omitempty"`
				Name                       string `json:"name,omitempty"`
			} `json:"capability,omitempty"`
			Name     string `json:"name,omitempty"`
			Provider []struct {
				Name                         string   `json:"name,omitempty"`
				Servicelist                  []string `json:"servicelist,omitempty"`
				State                        string   `json:"state,omitempty"`
				Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid,omitempty"`
				Physicalnetworkid            string   `json:"physicalnetworkid,omitempty"`
				Canenableindividualservice   bool     `json:"canenableindividualservice,omitempty"`
				Id                           string   `json:"id,omitempty"`
			} `json:"provider,omitempty"`
		} `json:"service,omitempty"`
		Acltype         string `json:"acltype,omitempty"`
		Aclid           string `json:"aclid,omitempty"`
		Gateway         string `json:"gateway,omitempty"`
		Account         string `json:"account,omitempty"`
		Zoneid          string `json:"zoneid,omitempty"`
		Issystem        bool   `json:"issystem,omitempty"`
		Ispersistent    bool   `json:"ispersistent,omitempty"`
		Isdefault       bool   `json:"isdefault,omitempty"`
		Project         string `json:"project,omitempty"`
		Subdomainaccess bool   `json:"subdomainaccess,omitempty"`
		Ip6gateway      string `json:"ip6gateway,omitempty"`
		Displaynetwork  bool   `json:"displaynetwork,omitempty"`
	} `json:"network,omitempty"`
	Vpcofferingid   string `json:"vpcofferingid,omitempty"`
	Account         string `json:"account,omitempty"`
	Restartrequired bool   `json:"restartrequired,omitempty"`
	State           string `json:"state,omitempty"`
}

type DeleteVPCParams struct {
	p map[string]interface{}
}

func (p *DeleteVPCParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteVPCParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteVPCParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewDeleteVPCParams(id string) *DeleteVPCParams {
	p := &DeleteVPCParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a VPC
func (s *VPCService) DeleteVPC(p *DeleteVPCParams) (*DeleteVPCResponse, error) {
	resp, err := s.cs.newRequest("deleteVPC", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteVPCResponse
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

		var r DeleteVPCResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type DeleteVPCResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Success     bool   `json:"success,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
}

type UpdateVPCParams struct {
	p map[string]interface{}
}

func (p *UpdateVPCParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	return u
}

func (p *UpdateVPCParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
	return
}

func (p *UpdateVPCParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *UpdateVPCParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

// You should always use this function to get a new UpdateVPCParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewUpdateVPCParams(id string, name string) *UpdateVPCParams {
	p := &UpdateVPCParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["name"] = name
	return p
}

// Updates a VPC
func (s *VPCService) UpdateVPC(p *UpdateVPCParams) (*UpdateVPCResponse, error) {
	resp, err := s.cs.newRequest("updateVPC", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateVPCResponse
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

		var r UpdateVPCResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type UpdateVPCResponse struct {
	JobID         string `json:"jobid,omitempty"`
	Domainid      string `json:"domainid,omitempty"`
	Zonename      string `json:"zonename,omitempty"`
	Displaytext   string `json:"displaytext,omitempty"`
	Networkdomain string `json:"networkdomain,omitempty"`
	Domain        string `json:"domain,omitempty"`
	Service       []struct {
		Name       string `json:"name,omitempty"`
		Capability []struct {
			Name                       string `json:"name,omitempty"`
			Value                      string `json:"value,omitempty"`
			Canchooseservicecapability bool   `json:"canchooseservicecapability,omitempty"`
		} `json:"capability,omitempty"`
		Provider []struct {
			Servicelist                  []string `json:"servicelist,omitempty"`
			Id                           string   `json:"id,omitempty"`
			Name                         string   `json:"name,omitempty"`
			Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid,omitempty"`
			Physicalnetworkid            string   `json:"physicalnetworkid,omitempty"`
			Canenableindividualservice   bool     `json:"canenableindividualservice,omitempty"`
			State                        string   `json:"state,omitempty"`
		} `json:"provider,omitempty"`
	} `json:"service,omitempty"`
	State         string `json:"state,omitempty"`
	Vpcofferingid string `json:"vpcofferingid,omitempty"`
	Projectid     string `json:"projectid,omitempty"`
	Network       []struct {
		Account         string `json:"account,omitempty"`
		Zonename        string `json:"zonename,omitempty"`
		Ispersistent    bool   `json:"ispersistent,omitempty"`
		Isdefault       bool   `json:"isdefault,omitempty"`
		Dns2            string `json:"dns2,omitempty"`
		Vpcid           string `json:"vpcid,omitempty"`
		Traffictype     string `json:"traffictype,omitempty"`
		Related         string `json:"related,omitempty"`
		Reservediprange string `json:"reservediprange,omitempty"`
		Ip6gateway      string `json:"ip6gateway,omitempty"`
		Canusefordeploy bool   `json:"canusefordeploy,omitempty"`
		Service         []struct {
			Name       string `json:"name,omitempty"`
			Capability []struct {
				Name                       string `json:"name,omitempty"`
				Canchooseservicecapability bool   `json:"canchooseservicecapability,omitempty"`
				Value                      string `json:"value,omitempty"`
			} `json:"capability,omitempty"`
			Provider []struct {
				Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid,omitempty"`
				Name                         string   `json:"name,omitempty"`
				Id                           string   `json:"id,omitempty"`
				Servicelist                  []string `json:"servicelist,omitempty"`
				State                        string   `json:"state,omitempty"`
				Canenableindividualservice   bool     `json:"canenableindividualservice,omitempty"`
				Physicalnetworkid            string   `json:"physicalnetworkid,omitempty"`
			} `json:"provider,omitempty"`
		} `json:"service,omitempty"`
		Domain                      string `json:"domain,omitempty"`
		Broadcastdomaintype         string `json:"broadcastdomaintype,omitempty"`
		Networkofferingavailability string `json:"networkofferingavailability,omitempty"`
		Dns1                        string `json:"dns1,omitempty"`
		Issystem                    bool   `json:"issystem,omitempty"`
		Specifyipranges             bool   `json:"specifyipranges,omitempty"`
		Networkofferingid           string `json:"networkofferingid,omitempty"`
		Project                     string `json:"project,omitempty"`
		Domainid                    string `json:"domainid,omitempty"`
		Acltype                     string `json:"acltype,omitempty"`
		Displaytext                 string `json:"displaytext,omitempty"`
		Ip6cidr                     string `json:"ip6cidr,omitempty"`
		Restartrequired             bool   `json:"restartrequired,omitempty"`
		Projectid                   string `json:"projectid,omitempty"`
		Networkdomain               string `json:"networkdomain,omitempty"`
		Zoneid                      string `json:"zoneid,omitempty"`
		Networkcidr                 string `json:"networkcidr,omitempty"`
		Id                          string `json:"id,omitempty"`
		Subdomainaccess             bool   `json:"subdomainaccess,omitempty"`
		Displaynetwork              bool   `json:"displaynetwork,omitempty"`
		Vlan                        string `json:"vlan,omitempty"`
		Netmask                     string `json:"netmask,omitempty"`
		Type                        string `json:"type,omitempty"`
		Name                        string `json:"name,omitempty"`
		Gateway                     string `json:"gateway,omitempty"`
		Tags                        []struct {
			Value        string `json:"value,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
			Key          string `json:"key,omitempty"`
			Project      string `json:"project,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Account      string `json:"account,omitempty"`
		} `json:"tags,omitempty"`
		Networkofferingdisplaytext  string `json:"networkofferingdisplaytext,omitempty"`
		Networkofferingconservemode bool   `json:"networkofferingconservemode,omitempty"`
		Cidr                        string `json:"cidr,omitempty"`
		Networkofferingname         string `json:"networkofferingname,omitempty"`
		Physicalnetworkid           string `json:"physicalnetworkid,omitempty"`
		Aclid                       string `json:"aclid,omitempty"`
		State                       string `json:"state,omitempty"`
		Broadcasturi                string `json:"broadcasturi,omitempty"`
	} `json:"network,omitempty"`
	Name    string `json:"name,omitempty"`
	Project string `json:"project,omitempty"`
	Id      string `json:"id,omitempty"`
	Created string `json:"created,omitempty"`
	Tags    []struct {
		Resourcetype string `json:"resourcetype,omitempty"`
		Value        string `json:"value,omitempty"`
		Account      string `json:"account,omitempty"`
		Key          string `json:"key,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Project      string `json:"project,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
	} `json:"tags,omitempty"`
	Cidr            string `json:"cidr,omitempty"`
	Account         string `json:"account,omitempty"`
	Restartrequired bool   `json:"restartrequired,omitempty"`
	Zoneid          string `json:"zoneid,omitempty"`
}

type RestartVPCParams struct {
	p map[string]interface{}
}

func (p *RestartVPCParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *RestartVPCParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new RestartVPCParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewRestartVPCParams(id string) *RestartVPCParams {
	p := &RestartVPCParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Restarts a VPC
func (s *VPCService) RestartVPC(p *RestartVPCParams) (*RestartVPCResponse, error) {
	resp, err := s.cs.newRequest("restartVPC", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RestartVPCResponse
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

		var r RestartVPCResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type RestartVPCResponse struct {
	JobID         string `json:"jobid,omitempty"`
	Networkdomain string `json:"networkdomain,omitempty"`
	Zoneid        string `json:"zoneid,omitempty"`
	Service       []struct {
		Name     string `json:"name,omitempty"`
		Provider []struct {
			Id                           string   `json:"id,omitempty"`
			Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid,omitempty"`
			Canenableindividualservice   bool     `json:"canenableindividualservice,omitempty"`
			Physicalnetworkid            string   `json:"physicalnetworkid,omitempty"`
			Name                         string   `json:"name,omitempty"`
			State                        string   `json:"state,omitempty"`
			Servicelist                  []string `json:"servicelist,omitempty"`
		} `json:"provider,omitempty"`
		Capability []struct {
			Value                      string `json:"value,omitempty"`
			Canchooseservicecapability bool   `json:"canchooseservicecapability,omitempty"`
			Name                       string `json:"name,omitempty"`
		} `json:"capability,omitempty"`
	} `json:"service,omitempty"`
	Cidr      string `json:"cidr,omitempty"`
	Name      string `json:"name,omitempty"`
	State     string `json:"state,omitempty"`
	Project   string `json:"project,omitempty"`
	Projectid string `json:"projectid,omitempty"`
	Domain    string `json:"domain,omitempty"`
	Id        string `json:"id,omitempty"`
	Tags      []struct {
		Domainid     string `json:"domainid,omitempty"`
		Value        string `json:"value,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Key          string `json:"key,omitempty"`
		Project      string `json:"project,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Account      string `json:"account,omitempty"`
	} `json:"tags,omitempty"`
	Domainid        string `json:"domainid,omitempty"`
	Restartrequired bool   `json:"restartrequired,omitempty"`
	Account         string `json:"account,omitempty"`
	Created         string `json:"created,omitempty"`
	Displaytext     string `json:"displaytext,omitempty"`
	Vpcofferingid   string `json:"vpcofferingid,omitempty"`
	Zonename        string `json:"zonename,omitempty"`
	Network         []struct {
		Dns2                        string `json:"dns2,omitempty"`
		Domainid                    string `json:"domainid,omitempty"`
		Related                     string `json:"related,omitempty"`
		Traffictype                 string `json:"traffictype,omitempty"`
		Networkofferingconservemode bool   `json:"networkofferingconservemode,omitempty"`
		Issystem                    bool   `json:"issystem,omitempty"`
		Displaynetwork              bool   `json:"displaynetwork,omitempty"`
		Networkofferingdisplaytext  string `json:"networkofferingdisplaytext,omitempty"`
		Aclid                       string `json:"aclid,omitempty"`
		Reservediprange             string `json:"reservediprange,omitempty"`
		Zoneid                      string `json:"zoneid,omitempty"`
		Displaytext                 string `json:"displaytext,omitempty"`
		Vlan                        string `json:"vlan,omitempty"`
		Restartrequired             bool   `json:"restartrequired,omitempty"`
		Ip6cidr                     string `json:"ip6cidr,omitempty"`
		Isdefault                   bool   `json:"isdefault,omitempty"`
		Domain                      string `json:"domain,omitempty"`
		Id                          string `json:"id,omitempty"`
		Networkdomain               string `json:"networkdomain,omitempty"`
		Physicalnetworkid           string `json:"physicalnetworkid,omitempty"`
		Service                     []struct {
			Provider []struct {
				Id                           string   `json:"id,omitempty"`
				Servicelist                  []string `json:"servicelist,omitempty"`
				Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid,omitempty"`
				Canenableindividualservice   bool     `json:"canenableindividualservice,omitempty"`
				State                        string   `json:"state,omitempty"`
				Name                         string   `json:"name,omitempty"`
				Physicalnetworkid            string   `json:"physicalnetworkid,omitempty"`
			} `json:"provider,omitempty"`
			Capability []struct {
				Value                      string `json:"value,omitempty"`
				Canchooseservicecapability bool   `json:"canchooseservicecapability,omitempty"`
				Name                       string `json:"name,omitempty"`
			} `json:"capability,omitempty"`
			Name string `json:"name,omitempty"`
		} `json:"service,omitempty"`
		Canusefordeploy     bool   `json:"canusefordeploy,omitempty"`
		Type                string `json:"type,omitempty"`
		Broadcasturi        string `json:"broadcasturi,omitempty"`
		Broadcastdomaintype string `json:"broadcastdomaintype,omitempty"`
		Networkofferingid   string `json:"networkofferingid,omitempty"`
		Networkofferingname string `json:"networkofferingname,omitempty"`
		Projectid           string `json:"projectid,omitempty"`
		Tags                []struct {
			Domain       string `json:"domain,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Project      string `json:"project,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Value        string `json:"value,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Account      string `json:"account,omitempty"`
			Key          string `json:"key,omitempty"`
		} `json:"tags,omitempty"`
		Specifyipranges             bool   `json:"specifyipranges,omitempty"`
		Gateway                     string `json:"gateway,omitempty"`
		Netmask                     string `json:"netmask,omitempty"`
		Subdomainaccess             bool   `json:"subdomainaccess,omitempty"`
		Acltype                     string `json:"acltype,omitempty"`
		Ispersistent                bool   `json:"ispersistent,omitempty"`
		Vpcid                       string `json:"vpcid,omitempty"`
		Project                     string `json:"project,omitempty"`
		Name                        string `json:"name,omitempty"`
		Ip6gateway                  string `json:"ip6gateway,omitempty"`
		Networkofferingavailability string `json:"networkofferingavailability,omitempty"`
		Account                     string `json:"account,omitempty"`
		Cidr                        string `json:"cidr,omitempty"`
		Zonename                    string `json:"zonename,omitempty"`
		Networkcidr                 string `json:"networkcidr,omitempty"`
		Dns1                        string `json:"dns1,omitempty"`
		State                       string `json:"state,omitempty"`
	} `json:"network,omitempty"`
}

type CreateVPCOfferingParams struct {
	p map[string]interface{}
}

func (p *CreateVPCOfferingParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	if v, found := p.p["serviceproviderlist"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("serviceproviderlist[%d].key", i), k)
			u.Set(fmt.Sprintf("serviceproviderlist[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["supportedservices"]; found {
		vv := strings.Join(v.([]string), ", ")
		u.Set("supportedservices", vv)
	}
	return u
}

func (p *CreateVPCOfferingParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
	return
}

func (p *CreateVPCOfferingParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *CreateVPCOfferingParams) SetServiceofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingid"] = v
	return
}

func (p *CreateVPCOfferingParams) SetServiceproviderlist(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceproviderlist"] = v
	return
}

func (p *CreateVPCOfferingParams) SetSupportedservices(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["supportedservices"] = v
	return
}

// You should always use this function to get a new CreateVPCOfferingParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewCreateVPCOfferingParams(displaytext string, name string, supportedservices []string) *CreateVPCOfferingParams {
	p := &CreateVPCOfferingParams{}
	p.p = make(map[string]interface{})
	p.p["displaytext"] = displaytext
	p.p["name"] = name
	p.p["supportedservices"] = supportedservices
	return p
}

// Creates VPC offering
func (s *VPCService) CreateVPCOffering(p *CreateVPCOfferingParams) (*CreateVPCOfferingResponse, error) {
	resp, err := s.cs.newRequest("createVPCOffering", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateVPCOfferingResponse
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

		var r CreateVPCOfferingResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type CreateVPCOfferingResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	State       string `json:"state,omitempty"`
	Service     []struct {
		Provider []struct {
			Physicalnetworkid            string   `json:"physicalnetworkid,omitempty"`
			Servicelist                  []string `json:"servicelist,omitempty"`
			Name                         string   `json:"name,omitempty"`
			Canenableindividualservice   bool     `json:"canenableindividualservice,omitempty"`
			Id                           string   `json:"id,omitempty"`
			Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid,omitempty"`
			State                        string   `json:"state,omitempty"`
		} `json:"provider,omitempty"`
		Name       string `json:"name,omitempty"`
		Capability []struct {
			Canchooseservicecapability bool   `json:"canchooseservicecapability,omitempty"`
			Name                       string `json:"name,omitempty"`
			Value                      string `json:"value,omitempty"`
		} `json:"capability,omitempty"`
	} `json:"service,omitempty"`
	Isdefault bool   `json:"isdefault,omitempty"`
	Created   string `json:"created,omitempty"`
	Name      string `json:"name,omitempty"`
	Id        string `json:"id,omitempty"`
}

type UpdateVPCOfferingParams struct {
	p map[string]interface{}
}

func (p *UpdateVPCOfferingParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	return u
}

func (p *UpdateVPCOfferingParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
	return
}

func (p *UpdateVPCOfferingParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *UpdateVPCOfferingParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *UpdateVPCOfferingParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
	return
}

// You should always use this function to get a new UpdateVPCOfferingParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewUpdateVPCOfferingParams(id string) *UpdateVPCOfferingParams {
	p := &UpdateVPCOfferingParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates VPC offering
func (s *VPCService) UpdateVPCOffering(p *UpdateVPCOfferingParams) (*UpdateVPCOfferingResponse, error) {
	resp, err := s.cs.newRequest("updateVPCOffering", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateVPCOfferingResponse
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

		var r UpdateVPCOfferingResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type UpdateVPCOfferingResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Name        string `json:"name,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	State       string `json:"state,omitempty"`
	Isdefault   bool   `json:"isdefault,omitempty"`
	Created     string `json:"created,omitempty"`
	Id          string `json:"id,omitempty"`
	Service     []struct {
		Capability []struct {
			Name                       string `json:"name,omitempty"`
			Value                      string `json:"value,omitempty"`
			Canchooseservicecapability bool   `json:"canchooseservicecapability,omitempty"`
		} `json:"capability,omitempty"`
		Name     string `json:"name,omitempty"`
		Provider []struct {
			Id                           string   `json:"id,omitempty"`
			Name                         string   `json:"name,omitempty"`
			Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid,omitempty"`
			Physicalnetworkid            string   `json:"physicalnetworkid,omitempty"`
			Servicelist                  []string `json:"servicelist,omitempty"`
			Canenableindividualservice   bool     `json:"canenableindividualservice,omitempty"`
			State                        string   `json:"state,omitempty"`
		} `json:"provider,omitempty"`
	} `json:"service,omitempty"`
}

type DeleteVPCOfferingParams struct {
	p map[string]interface{}
}

func (p *DeleteVPCOfferingParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteVPCOfferingParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteVPCOfferingParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewDeleteVPCOfferingParams(id string) *DeleteVPCOfferingParams {
	p := &DeleteVPCOfferingParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes VPC offering
func (s *VPCService) DeleteVPCOffering(p *DeleteVPCOfferingParams) (*DeleteVPCOfferingResponse, error) {
	resp, err := s.cs.newRequest("deleteVPCOffering", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteVPCOfferingResponse
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

		var r DeleteVPCOfferingResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type DeleteVPCOfferingResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type ListVPCOfferingsParams struct {
	p map[string]interface{}
}

func (p *ListVPCOfferingsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["isdefault"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isdefault", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
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
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := p.p["supportedservices"]; found {
		vv := strings.Join(v.([]string), ", ")
		u.Set("supportedservices", vv)
	}
	return u
}

func (p *ListVPCOfferingsParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
	return
}

func (p *ListVPCOfferingsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListVPCOfferingsParams) SetIsdefault(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isdefault"] = v
	return
}

func (p *ListVPCOfferingsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListVPCOfferingsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *ListVPCOfferingsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListVPCOfferingsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListVPCOfferingsParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
	return
}

func (p *ListVPCOfferingsParams) SetSupportedservices(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["supportedservices"] = v
	return
}

// You should always use this function to get a new ListVPCOfferingsParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewListVPCOfferingsParams() *ListVPCOfferingsParams {
	p := &ListVPCOfferingsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VPCService) GetVPCOfferingID(name string) (string, error) {
	p := &ListVPCOfferingsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	l, err := s.ListVPCOfferings(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, name, l)
	}
	return l.VPCOfferings[0].Id, nil
}

// Lists VPC offerings
func (s *VPCService) ListVPCOfferings(p *ListVPCOfferingsParams) (*ListVPCOfferingsResponse, error) {
	resp, err := s.cs.newRequest("listVPCOfferings", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVPCOfferingsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListVPCOfferingsResponse struct {
	Count        int            `json:"count"`
	VPCOfferings []*VPCOffering `json:"vpcoffering"`
}

type VPCOffering struct {
	State     string `json:"state,omitempty"`
	Name      string `json:"name,omitempty"`
	Isdefault bool   `json:"isdefault,omitempty"`
	Service   []struct {
		Provider []struct {
			Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid,omitempty"`
			Physicalnetworkid            string   `json:"physicalnetworkid,omitempty"`
			Servicelist                  []string `json:"servicelist,omitempty"`
			Id                           string   `json:"id,omitempty"`
			State                        string   `json:"state,omitempty"`
			Canenableindividualservice   bool     `json:"canenableindividualservice,omitempty"`
			Name                         string   `json:"name,omitempty"`
		} `json:"provider,omitempty"`
		Capability []struct {
			Value                      string `json:"value,omitempty"`
			Canchooseservicecapability bool   `json:"canchooseservicecapability,omitempty"`
			Name                       string `json:"name,omitempty"`
		} `json:"capability,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"service,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Id          string `json:"id,omitempty"`
	Created     string `json:"created,omitempty"`
}

type CreatePrivateGatewayParams struct {
	p map[string]interface{}
}

func (p *CreatePrivateGatewayParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["aclid"]; found {
		u.Set("aclid", v.(string))
	}
	if v, found := p.p["gateway"]; found {
		u.Set("gateway", v.(string))
	}
	if v, found := p.p["ipaddress"]; found {
		u.Set("ipaddress", v.(string))
	}
	if v, found := p.p["netmask"]; found {
		u.Set("netmask", v.(string))
	}
	if v, found := p.p["networkofferingid"]; found {
		u.Set("networkofferingid", v.(string))
	}
	if v, found := p.p["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	if v, found := p.p["sourcenatsupported"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("sourcenatsupported", vv)
	}
	if v, found := p.p["vlan"]; found {
		u.Set("vlan", v.(string))
	}
	if v, found := p.p["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	return u
}

func (p *CreatePrivateGatewayParams) SetAclid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["aclid"] = v
	return
}

func (p *CreatePrivateGatewayParams) SetGateway(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gateway"] = v
	return
}

func (p *CreatePrivateGatewayParams) SetIpaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddress"] = v
	return
}

func (p *CreatePrivateGatewayParams) SetNetmask(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["netmask"] = v
	return
}

func (p *CreatePrivateGatewayParams) SetNetworkofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkofferingid"] = v
	return
}

func (p *CreatePrivateGatewayParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
	return
}

func (p *CreatePrivateGatewayParams) SetSourcenatsupported(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sourcenatsupported"] = v
	return
}

func (p *CreatePrivateGatewayParams) SetVlan(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vlan"] = v
	return
}

func (p *CreatePrivateGatewayParams) SetVpcid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcid"] = v
	return
}

// You should always use this function to get a new CreatePrivateGatewayParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewCreatePrivateGatewayParams(gateway string, ipaddress string, netmask string, vlan string, vpcid string) *CreatePrivateGatewayParams {
	p := &CreatePrivateGatewayParams{}
	p.p = make(map[string]interface{})
	p.p["gateway"] = gateway
	p.p["ipaddress"] = ipaddress
	p.p["netmask"] = netmask
	p.p["vlan"] = vlan
	p.p["vpcid"] = vpcid
	return p
}

// Creates a private gateway
func (s *VPCService) CreatePrivateGateway(p *CreatePrivateGatewayParams) (*CreatePrivateGatewayResponse, error) {
	resp, err := s.cs.newRequest("createPrivateGateway", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreatePrivateGatewayResponse
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

		var r CreatePrivateGatewayResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type CreatePrivateGatewayResponse struct {
	JobID              string `json:"jobid,omitempty"`
	Account            string `json:"account,omitempty"`
	Project            string `json:"project,omitempty"`
	State              string `json:"state,omitempty"`
	Ipaddress          string `json:"ipaddress,omitempty"`
	Zonename           string `json:"zonename,omitempty"`
	Domain             string `json:"domain,omitempty"`
	Sourcenatsupported bool   `json:"sourcenatsupported,omitempty"`
	Zoneid             string `json:"zoneid,omitempty"`
	Aclid              string `json:"aclid,omitempty"`
	Netmask            string `json:"netmask,omitempty"`
	Id                 string `json:"id,omitempty"`
	Vlan               string `json:"vlan,omitempty"`
	Physicalnetworkid  string `json:"physicalnetworkid,omitempty"`
	Domainid           string `json:"domainid,omitempty"`
	Projectid          string `json:"projectid,omitempty"`
	Gateway            string `json:"gateway,omitempty"`
	Vpcid              string `json:"vpcid,omitempty"`
}

type ListPrivateGatewaysParams struct {
	p map[string]interface{}
}

func (p *ListPrivateGatewaysParams) toURLValues() url.Values {
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
	if v, found := p.p["ipaddress"]; found {
		u.Set("ipaddress", v.(string))
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
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := p.p["vlan"]; found {
		u.Set("vlan", v.(string))
	}
	if v, found := p.p["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	return u
}

func (p *ListPrivateGatewaysParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListPrivateGatewaysParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListPrivateGatewaysParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListPrivateGatewaysParams) SetIpaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddress"] = v
	return
}

func (p *ListPrivateGatewaysParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
	return
}

func (p *ListPrivateGatewaysParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListPrivateGatewaysParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
	return
}

func (p *ListPrivateGatewaysParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListPrivateGatewaysParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListPrivateGatewaysParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *ListPrivateGatewaysParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
	return
}

func (p *ListPrivateGatewaysParams) SetVlan(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vlan"] = v
	return
}

func (p *ListPrivateGatewaysParams) SetVpcid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcid"] = v
	return
}

// You should always use this function to get a new ListPrivateGatewaysParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewListPrivateGatewaysParams() *ListPrivateGatewaysParams {
	p := &ListPrivateGatewaysParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VPCService) GetPrivateGatewayID(keyword string) (string, error) {
	p := &ListPrivateGatewaysParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword

	l, err := s.ListPrivateGateways(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, keyword, l)
	}
	return l.PrivateGateways[0].Id, nil
}

// List private gateways
func (s *VPCService) ListPrivateGateways(p *ListPrivateGatewaysParams) (*ListPrivateGatewaysResponse, error) {
	resp, err := s.cs.newRequest("listPrivateGateways", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListPrivateGatewaysResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListPrivateGatewaysResponse struct {
	Count           int               `json:"count"`
	PrivateGateways []*PrivateGateway `json:"privategateway"`
}

type PrivateGateway struct {
	Gateway            string `json:"gateway,omitempty"`
	Vlan               string `json:"vlan,omitempty"`
	Physicalnetworkid  string `json:"physicalnetworkid,omitempty"`
	Account            string `json:"account,omitempty"`
	Vpcid              string `json:"vpcid,omitempty"`
	Zoneid             string `json:"zoneid,omitempty"`
	Project            string `json:"project,omitempty"`
	Domainid           string `json:"domainid,omitempty"`
	Zonename           string `json:"zonename,omitempty"`
	Aclid              string `json:"aclid,omitempty"`
	Netmask            string `json:"netmask,omitempty"`
	Projectid          string `json:"projectid,omitempty"`
	Sourcenatsupported bool   `json:"sourcenatsupported,omitempty"`
	State              string `json:"state,omitempty"`
	Domain             string `json:"domain,omitempty"`
	Id                 string `json:"id,omitempty"`
	Ipaddress          string `json:"ipaddress,omitempty"`
}

type DeletePrivateGatewayParams struct {
	p map[string]interface{}
}

func (p *DeletePrivateGatewayParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeletePrivateGatewayParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeletePrivateGatewayParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewDeletePrivateGatewayParams(id string) *DeletePrivateGatewayParams {
	p := &DeletePrivateGatewayParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a Private gateway
func (s *VPCService) DeletePrivateGateway(p *DeletePrivateGatewayParams) (*DeletePrivateGatewayResponse, error) {
	resp, err := s.cs.newRequest("deletePrivateGateway", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeletePrivateGatewayResponse
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

		var r DeletePrivateGatewayResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type DeletePrivateGatewayResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Success     bool   `json:"success,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
}

type CreateStaticRouteParams struct {
	p map[string]interface{}
}

func (p *CreateStaticRouteParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["cidr"]; found {
		u.Set("cidr", v.(string))
	}
	if v, found := p.p["gatewayid"]; found {
		u.Set("gatewayid", v.(string))
	}
	return u
}

func (p *CreateStaticRouteParams) SetCidr(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cidr"] = v
	return
}

func (p *CreateStaticRouteParams) SetGatewayid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gatewayid"] = v
	return
}

// You should always use this function to get a new CreateStaticRouteParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewCreateStaticRouteParams(cidr string, gatewayid string) *CreateStaticRouteParams {
	p := &CreateStaticRouteParams{}
	p.p = make(map[string]interface{})
	p.p["cidr"] = cidr
	p.p["gatewayid"] = gatewayid
	return p
}

// Creates a static route
func (s *VPCService) CreateStaticRoute(p *CreateStaticRouteParams) (*CreateStaticRouteResponse, error) {
	resp, err := s.cs.newRequest("createStaticRoute", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateStaticRouteResponse
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

		var r CreateStaticRouteResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type CreateStaticRouteResponse struct {
	JobID string `json:"jobid,omitempty"`
	Tags  []struct {
		Project      string `json:"project,omitempty"`
		Key          string `json:"key,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Account      string `json:"account,omitempty"`
		Value        string `json:"value,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
	} `json:"tags,omitempty"`
	Domainid  string `json:"domainid,omitempty"`
	Project   string `json:"project,omitempty"`
	Account   string `json:"account,omitempty"`
	State     string `json:"state,omitempty"`
	Domain    string `json:"domain,omitempty"`
	Cidr      string `json:"cidr,omitempty"`
	Gatewayid string `json:"gatewayid,omitempty"`
	Projectid string `json:"projectid,omitempty"`
	Vpcid     string `json:"vpcid,omitempty"`
	Id        string `json:"id,omitempty"`
}

type DeleteStaticRouteParams struct {
	p map[string]interface{}
}

func (p *DeleteStaticRouteParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteStaticRouteParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteStaticRouteParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewDeleteStaticRouteParams(id string) *DeleteStaticRouteParams {
	p := &DeleteStaticRouteParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a static route
func (s *VPCService) DeleteStaticRoute(p *DeleteStaticRouteParams) (*DeleteStaticRouteResponse, error) {
	resp, err := s.cs.newRequest("deleteStaticRoute", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteStaticRouteResponse
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

		var r DeleteStaticRouteResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type DeleteStaticRouteResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Success     bool   `json:"success,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
}

type ListStaticRoutesParams struct {
	p map[string]interface{}
}

func (p *ListStaticRoutesParams) toURLValues() url.Values {
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
	if v, found := p.p["gatewayid"]; found {
		u.Set("gatewayid", v.(string))
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
	if v, found := p.p["tags"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	return u
}

func (p *ListStaticRoutesParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListStaticRoutesParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListStaticRoutesParams) SetGatewayid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gatewayid"] = v
	return
}

func (p *ListStaticRoutesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListStaticRoutesParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
	return
}

func (p *ListStaticRoutesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListStaticRoutesParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
	return
}

func (p *ListStaticRoutesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListStaticRoutesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListStaticRoutesParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *ListStaticRoutesParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
	return
}

func (p *ListStaticRoutesParams) SetVpcid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcid"] = v
	return
}

// You should always use this function to get a new ListStaticRoutesParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewListStaticRoutesParams() *ListStaticRoutesParams {
	p := &ListStaticRoutesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VPCService) GetStaticRouteID(keyword string) (string, error) {
	p := &ListStaticRoutesParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword

	l, err := s.ListStaticRoutes(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, keyword, l)
	}
	return l.StaticRoutes[0].Id, nil
}

// Lists all static routes
func (s *VPCService) ListStaticRoutes(p *ListStaticRoutesParams) (*ListStaticRoutesResponse, error) {
	resp, err := s.cs.newRequest("listStaticRoutes", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListStaticRoutesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListStaticRoutesResponse struct {
	Count        int            `json:"count"`
	StaticRoutes []*StaticRoute `json:"staticroute"`
}

type StaticRoute struct {
	Projectid string `json:"projectid,omitempty"`
	Tags      []struct {
		Domain       string `json:"domain,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Project      string `json:"project,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Value        string `json:"value,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Account      string `json:"account,omitempty"`
		Key          string `json:"key,omitempty"`
	} `json:"tags,omitempty"`
	Cidr      string `json:"cidr,omitempty"`
	Gatewayid string `json:"gatewayid,omitempty"`
	Domain    string `json:"domain,omitempty"`
	State     string `json:"state,omitempty"`
	Domainid  string `json:"domainid,omitempty"`
	Id        string `json:"id,omitempty"`
	Account   string `json:"account,omitempty"`
	Vpcid     string `json:"vpcid,omitempty"`
	Project   string `json:"project,omitempty"`
}
