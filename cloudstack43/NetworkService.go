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

type DedicatePublicIpRangeParams struct {
	p map[string]interface{}
}

func (p *DedicatePublicIpRangeParams) toURLValues() url.Values {
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
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	return u
}

func (p *DedicatePublicIpRangeParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *DedicatePublicIpRangeParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *DedicatePublicIpRangeParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *DedicatePublicIpRangeParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

// You should always use this function to get a new DedicatePublicIpRangeParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewDedicatePublicIpRangeParams(account string, domainid string, id string) *DedicatePublicIpRangeParams {
	p := &DedicatePublicIpRangeParams{}
	p.p = make(map[string]interface{})
	p.p["account"] = account
	p.p["domainid"] = domainid
	p.p["id"] = id
	return p
}

// Dedicates a Public IP range to an account
func (s *NetworkService) DedicatePublicIpRange(p *DedicatePublicIpRangeParams) (*DedicatePublicIpRangeResponse, error) {
	resp, err := s.cs.newRequest("dedicatePublicIpRange", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DedicatePublicIpRangeResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type DedicatePublicIpRangeResponse struct {
	Startipv6         string `json:"startipv6,omitempty"`
	Projectid         string `json:"projectid,omitempty"`
	Project           string `json:"project,omitempty"`
	Physicalnetworkid string `json:"physicalnetworkid,omitempty"`
	Account           string `json:"account,omitempty"`
	Gateway           string `json:"gateway,omitempty"`
	Zoneid            string `json:"zoneid,omitempty"`
	Domain            string `json:"domain,omitempty"`
	Podid             string `json:"podid,omitempty"`
	Netmask           string `json:"netmask,omitempty"`
	Networkid         string `json:"networkid,omitempty"`
	Endip             string `json:"endip,omitempty"`
	Ip6cidr           string `json:"ip6cidr,omitempty"`
	Domainid          string `json:"domainid,omitempty"`
	Description       string `json:"description,omitempty"`
	Forvirtualnetwork bool   `json:"forvirtualnetwork,omitempty"`
	Vlan              string `json:"vlan,omitempty"`
	Id                string `json:"id,omitempty"`
	Podname           string `json:"podname,omitempty"`
	Endipv6           string `json:"endipv6,omitempty"`
	Ip6gateway        string `json:"ip6gateway,omitempty"`
	Startip           string `json:"startip,omitempty"`
}

type ReleasePublicIpRangeParams struct {
	p map[string]interface{}
}

func (p *ReleasePublicIpRangeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *ReleasePublicIpRangeParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new ReleasePublicIpRangeParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewReleasePublicIpRangeParams(id string) *ReleasePublicIpRangeParams {
	p := &ReleasePublicIpRangeParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Releases a Public IP range back to the system pool
func (s *NetworkService) ReleasePublicIpRange(p *ReleasePublicIpRangeParams) (*ReleasePublicIpRangeResponse, error) {
	resp, err := s.cs.newRequest("releasePublicIpRange", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ReleasePublicIpRangeResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ReleasePublicIpRangeResponse struct {
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type CreateNetworkParams struct {
	p map[string]interface{}
}

func (p *CreateNetworkParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["aclid"]; found {
		u.Set("aclid", v.(string))
	}
	if v, found := p.p["acltype"]; found {
		u.Set("acltype", v.(string))
	}
	if v, found := p.p["displaynetwork"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("displaynetwork", vv)
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["endip"]; found {
		u.Set("endip", v.(string))
	}
	if v, found := p.p["endipv6"]; found {
		u.Set("endipv6", v.(string))
	}
	if v, found := p.p["gateway"]; found {
		u.Set("gateway", v.(string))
	}
	if v, found := p.p["ip6cidr"]; found {
		u.Set("ip6cidr", v.(string))
	}
	if v, found := p.p["ip6gateway"]; found {
		u.Set("ip6gateway", v.(string))
	}
	if v, found := p.p["isolatedpvlan"]; found {
		u.Set("isolatedpvlan", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["netmask"]; found {
		u.Set("netmask", v.(string))
	}
	if v, found := p.p["networkdomain"]; found {
		u.Set("networkdomain", v.(string))
	}
	if v, found := p.p["networkofferingid"]; found {
		u.Set("networkofferingid", v.(string))
	}
	if v, found := p.p["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["startip"]; found {
		u.Set("startip", v.(string))
	}
	if v, found := p.p["startipv6"]; found {
		u.Set("startipv6", v.(string))
	}
	if v, found := p.p["subdomainaccess"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("subdomainaccess", vv)
	}
	if v, found := p.p["vlan"]; found {
		u.Set("vlan", v.(string))
	}
	if v, found := p.p["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *CreateNetworkParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *CreateNetworkParams) SetAclid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["aclid"] = v
	return
}

func (p *CreateNetworkParams) SetAcltype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["acltype"] = v
	return
}

func (p *CreateNetworkParams) SetDisplaynetwork(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaynetwork"] = v
	return
}

func (p *CreateNetworkParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
	return
}

func (p *CreateNetworkParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *CreateNetworkParams) SetEndip(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["endip"] = v
	return
}

func (p *CreateNetworkParams) SetEndipv6(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["endipv6"] = v
	return
}

func (p *CreateNetworkParams) SetGateway(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gateway"] = v
	return
}

func (p *CreateNetworkParams) SetIp6cidr(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ip6cidr"] = v
	return
}

func (p *CreateNetworkParams) SetIp6gateway(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ip6gateway"] = v
	return
}

func (p *CreateNetworkParams) SetIsolatedpvlan(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isolatedpvlan"] = v
	return
}

func (p *CreateNetworkParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *CreateNetworkParams) SetNetmask(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["netmask"] = v
	return
}

func (p *CreateNetworkParams) SetNetworkdomain(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkdomain"] = v
	return
}

func (p *CreateNetworkParams) SetNetworkofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkofferingid"] = v
	return
}

func (p *CreateNetworkParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
	return
}

func (p *CreateNetworkParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *CreateNetworkParams) SetStartip(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startip"] = v
	return
}

func (p *CreateNetworkParams) SetStartipv6(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startipv6"] = v
	return
}

func (p *CreateNetworkParams) SetSubdomainaccess(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["subdomainaccess"] = v
	return
}

func (p *CreateNetworkParams) SetVlan(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vlan"] = v
	return
}

func (p *CreateNetworkParams) SetVpcid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcid"] = v
	return
}

func (p *CreateNetworkParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new CreateNetworkParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewCreateNetworkParams(displaytext string, name string, networkofferingid string, zoneid string) *CreateNetworkParams {
	p := &CreateNetworkParams{}
	p.p = make(map[string]interface{})
	p.p["displaytext"] = displaytext
	p.p["name"] = name
	p.p["networkofferingid"] = networkofferingid
	p.p["zoneid"] = zoneid
	return p
}

// Creates a network
func (s *NetworkService) CreateNetwork(p *CreateNetworkParams) (*CreateNetworkResponse, error) {
	resp, err := s.cs.newRequest("createNetwork", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateNetworkResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type CreateNetworkResponse struct {
	Related                     string `json:"related,omitempty"`
	Account                     string `json:"account,omitempty"`
	Physicalnetworkid           string `json:"physicalnetworkid,omitempty"`
	Project                     string `json:"project,omitempty"`
	Networkofferingname         string `json:"networkofferingname,omitempty"`
	Ip6cidr                     string `json:"ip6cidr,omitempty"`
	Ip6gateway                  string `json:"ip6gateway,omitempty"`
	Displaytext                 string `json:"displaytext,omitempty"`
	Isdefault                   bool   `json:"isdefault,omitempty"`
	Cidr                        string `json:"cidr,omitempty"`
	Vlan                        string `json:"vlan,omitempty"`
	Canusefordeploy             bool   `json:"canusefordeploy,omitempty"`
	Zonename                    string `json:"zonename,omitempty"`
	Gateway                     string `json:"gateway,omitempty"`
	Networkofferingdisplaytext  string `json:"networkofferingdisplaytext,omitempty"`
	Projectid                   string `json:"projectid,omitempty"`
	Zoneid                      string `json:"zoneid,omitempty"`
	Restartrequired             bool   `json:"restartrequired,omitempty"`
	Acltype                     string `json:"acltype,omitempty"`
	Domain                      string `json:"domain,omitempty"`
	Issystem                    bool   `json:"issystem,omitempty"`
	Networkcidr                 string `json:"networkcidr,omitempty"`
	Networkofferingavailability string `json:"networkofferingavailability,omitempty"`
	Reservediprange             string `json:"reservediprange,omitempty"`
	Traffictype                 string `json:"traffictype,omitempty"`
	Tags                        []struct {
		Project      string `json:"project,omitempty"`
		Value        string `json:"value,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Key          string `json:"key,omitempty"`
		Account      string `json:"account,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
	} `json:"tags,omitempty"`
	Service []struct {
		Name     string `json:"name,omitempty"`
		Provider []struct {
			Physicalnetworkid            string   `json:"physicalnetworkid,omitempty"`
			Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid,omitempty"`
			Canenableindividualservice   bool     `json:"canenableindividualservice,omitempty"`
			Id                           string   `json:"id,omitempty"`
			Name                         string   `json:"name,omitempty"`
			Servicelist                  []string `json:"servicelist,omitempty"`
			State                        string   `json:"state,omitempty"`
		} `json:"provider,omitempty"`
		Capability []struct {
			Canchooseservicecapability bool   `json:"canchooseservicecapability,omitempty"`
			Value                      string `json:"value,omitempty"`
			Name                       string `json:"name,omitempty"`
		} `json:"capability,omitempty"`
	} `json:"service,omitempty"`
	Vpcid                       string `json:"vpcid,omitempty"`
	Specifyipranges             bool   `json:"specifyipranges,omitempty"`
	State                       string `json:"state,omitempty"`
	Displaynetwork              bool   `json:"displaynetwork,omitempty"`
	Aclid                       string `json:"aclid,omitempty"`
	Ispersistent                bool   `json:"ispersistent,omitempty"`
	Networkofferingid           string `json:"networkofferingid,omitempty"`
	Networkofferingconservemode bool   `json:"networkofferingconservemode,omitempty"`
	Domainid                    string `json:"domainid,omitempty"`
	Broadcastdomaintype         string `json:"broadcastdomaintype,omitempty"`
	Dns2                        string `json:"dns2,omitempty"`
	Subdomainaccess             bool   `json:"subdomainaccess,omitempty"`
	Id                          string `json:"id,omitempty"`
	Networkdomain               string `json:"networkdomain,omitempty"`
	Name                        string `json:"name,omitempty"`
	Broadcasturi                string `json:"broadcasturi,omitempty"`
	Type                        string `json:"type,omitempty"`
	Dns1                        string `json:"dns1,omitempty"`
	Netmask                     string `json:"netmask,omitempty"`
}

type DeleteNetworkParams struct {
	p map[string]interface{}
}

func (p *DeleteNetworkParams) toURLValues() url.Values {
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

func (p *DeleteNetworkParams) SetForced(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forced"] = v
	return
}

func (p *DeleteNetworkParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteNetworkParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewDeleteNetworkParams(id string) *DeleteNetworkParams {
	p := &DeleteNetworkParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a network
func (s *NetworkService) DeleteNetwork(p *DeleteNetworkParams) (*DeleteNetworkResponse, error) {
	resp, err := s.cs.newRequest("deleteNetwork", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteNetworkResponse
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

		var r DeleteNetworkResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type DeleteNetworkResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Success     bool   `json:"success,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
}

type ListNetworksParams struct {
	p map[string]interface{}
}

func (p *ListNetworksParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["acltype"]; found {
		u.Set("acltype", v.(string))
	}
	if v, found := p.p["canusefordeploy"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("canusefordeploy", vv)
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["forvpc"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forvpc", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := p.p["issystem"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("issystem", vv)
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
	if v, found := p.p["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["restartrequired"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("restartrequired", vv)
	}
	if v, found := p.p["specifyipranges"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("specifyipranges", vv)
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
	if v, found := p.p["traffictype"]; found {
		u.Set("traffictype", v.(string))
	}
	if v, found := p.p["type"]; found {
		u.Set("type", v.(string))
	}
	if v, found := p.p["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListNetworksParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListNetworksParams) SetAcltype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["acltype"] = v
	return
}

func (p *ListNetworksParams) SetCanusefordeploy(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["canusefordeploy"] = v
	return
}

func (p *ListNetworksParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListNetworksParams) SetForvpc(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forvpc"] = v
	return
}

func (p *ListNetworksParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListNetworksParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
	return
}

func (p *ListNetworksParams) SetIssystem(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["issystem"] = v
	return
}

func (p *ListNetworksParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListNetworksParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
	return
}

func (p *ListNetworksParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListNetworksParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListNetworksParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
	return
}

func (p *ListNetworksParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *ListNetworksParams) SetRestartrequired(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["restartrequired"] = v
	return
}

func (p *ListNetworksParams) SetSpecifyipranges(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["specifyipranges"] = v
	return
}

func (p *ListNetworksParams) SetSupportedservices(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["supportedservices"] = v
	return
}

func (p *ListNetworksParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
	return
}

func (p *ListNetworksParams) SetTraffictype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["traffictype"] = v
	return
}

func (p *ListNetworksParams) SetType(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkType"] = v
	return
}

func (p *ListNetworksParams) SetVpcid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcid"] = v
	return
}

func (p *ListNetworksParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new ListNetworksParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewListNetworksParams() *ListNetworksParams {
	p := &ListNetworksParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkService) GetNetworkID(keyword string) (string, error) {
	p := &ListNetworksParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword

	l, err := s.ListNetworks(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, keyword, l)
	}
	return l.Networks[0].Id, nil
}

// Lists all available networks.
func (s *NetworkService) ListNetworks(p *ListNetworksParams) (*ListNetworksResponse, error) {
	resp, err := s.cs.newRequest("listNetworks", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListNetworksResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListNetworksResponse struct {
	Count    int        `json:"count"`
	Networks []*Network `json:"network"`
}

type Network struct {
	Service []struct {
		Capability []struct {
			Canchooseservicecapability bool   `json:"canchooseservicecapability,omitempty"`
			Value                      string `json:"value,omitempty"`
			Name                       string `json:"name,omitempty"`
		} `json:"capability,omitempty"`
		Provider []struct {
			Physicalnetworkid            string   `json:"physicalnetworkid,omitempty"`
			Canenableindividualservice   bool     `json:"canenableindividualservice,omitempty"`
			Servicelist                  []string `json:"servicelist,omitempty"`
			Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid,omitempty"`
			Name                         string   `json:"name,omitempty"`
			Id                           string   `json:"id,omitempty"`
			State                        string   `json:"state,omitempty"`
		} `json:"provider,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"service,omitempty"`
	Gateway                    string `json:"gateway,omitempty"`
	Vlan                       string `json:"vlan,omitempty"`
	Account                    string `json:"account,omitempty"`
	Zonename                   string `json:"zonename,omitempty"`
	Traffictype                string `json:"traffictype,omitempty"`
	Cidr                       string `json:"cidr,omitempty"`
	Networkofferingdisplaytext string `json:"networkofferingdisplaytext,omitempty"`
	Aclid                      string `json:"aclid,omitempty"`
	Isdefault                  bool   `json:"isdefault,omitempty"`
	Networkcidr                string `json:"networkcidr,omitempty"`
	Vpcid                      string `json:"vpcid,omitempty"`
	Project                    string `json:"project,omitempty"`
	Netmask                    string `json:"netmask,omitempty"`
	Issystem                   bool   `json:"issystem,omitempty"`
	Type                       string `json:"type,omitempty"`
	Broadcastdomaintype        string `json:"broadcastdomaintype,omitempty"`
	Dns1                       string `json:"dns1,omitempty"`
	Related                    string `json:"related,omitempty"`
	Domain                     string `json:"domain,omitempty"`
	Projectid                  string `json:"projectid,omitempty"`
	Acltype                    string `json:"acltype,omitempty"`
	Displaynetwork             bool   `json:"displaynetwork,omitempty"`
	Networkdomain              string `json:"networkdomain,omitempty"`
	Broadcasturi               string `json:"broadcasturi,omitempty"`
	Canusefordeploy            bool   `json:"canusefordeploy,omitempty"`
	Tags                       []struct {
		Account      string `json:"account,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Value        string `json:"value,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Key          string `json:"key,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Project      string `json:"project,omitempty"`
	} `json:"tags,omitempty"`
	Id                          string `json:"id,omitempty"`
	Networkofferingavailability string `json:"networkofferingavailability,omitempty"`
	Domainid                    string `json:"domainid,omitempty"`
	Reservediprange             string `json:"reservediprange,omitempty"`
	Subdomainaccess             bool   `json:"subdomainaccess,omitempty"`
	Displaytext                 string `json:"displaytext,omitempty"`
	Dns2                        string `json:"dns2,omitempty"`
	Ip6gateway                  string `json:"ip6gateway,omitempty"`
	Physicalnetworkid           string `json:"physicalnetworkid,omitempty"`
	Ip6cidr                     string `json:"ip6cidr,omitempty"`
	Ispersistent                bool   `json:"ispersistent,omitempty"`
	Networkofferingconservemode bool   `json:"networkofferingconservemode,omitempty"`
	State                       string `json:"state,omitempty"`
	Zoneid                      string `json:"zoneid,omitempty"`
	Specifyipranges             bool   `json:"specifyipranges,omitempty"`
	Restartrequired             bool   `json:"restartrequired,omitempty"`
	Networkofferingid           string `json:"networkofferingid,omitempty"`
	Name                        string `json:"name,omitempty"`
	Networkofferingname         string `json:"networkofferingname,omitempty"`
}

type RestartNetworkParams struct {
	p map[string]interface{}
}

func (p *RestartNetworkParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["cleanup"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("cleanup", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *RestartNetworkParams) SetCleanup(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cleanup"] = v
	return
}

func (p *RestartNetworkParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new RestartNetworkParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewRestartNetworkParams(id string) *RestartNetworkParams {
	p := &RestartNetworkParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Restarts the network; includes 1) restarting network elements - virtual routers, dhcp servers 2) reapplying all public ips 3) reapplying loadBalancing/portForwarding rules
func (s *NetworkService) RestartNetwork(p *RestartNetworkParams) (*RestartNetworkResponse, error) {
	resp, err := s.cs.newRequest("restartNetwork", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RestartNetworkResponse
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

		var r RestartNetworkResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type RestartNetworkResponse struct {
	JobID string `json:"jobid,omitempty"`
	Tags  []struct {
		Domain       string `json:"domain,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Value        string `json:"value,omitempty"`
		Key          string `json:"key,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Account      string `json:"account,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Project      string `json:"project,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
	} `json:"tags,omitempty"`
	Issystem                  bool   `json:"issystem,omitempty"`
	Domain                    string `json:"domain,omitempty"`
	Projectid                 string `json:"projectid,omitempty"`
	Forvirtualnetwork         bool   `json:"forvirtualnetwork,omitempty"`
	Associatednetworkname     string `json:"associatednetworkname,omitempty"`
	Networkid                 string `json:"networkid,omitempty"`
	State                     string `json:"state,omitempty"`
	Vpcid                     string `json:"vpcid,omitempty"`
	Vlanid                    string `json:"vlanid,omitempty"`
	Physicalnetworkid         string `json:"physicalnetworkid,omitempty"`
	Virtualmachinename        string `json:"virtualmachinename,omitempty"`
	Project                   string `json:"project,omitempty"`
	Purpose                   string `json:"purpose,omitempty"`
	Virtualmachineid          string `json:"virtualmachineid,omitempty"`
	Allocated                 string `json:"allocated,omitempty"`
	Vmipaddress               string `json:"vmipaddress,omitempty"`
	Domainid                  string `json:"domainid,omitempty"`
	Vlanname                  string `json:"vlanname,omitempty"`
	Virtualmachinedisplayname string `json:"virtualmachinedisplayname,omitempty"`
	Isstaticnat               bool   `json:"isstaticnat,omitempty"`
	Account                   string `json:"account,omitempty"`
	Id                        string `json:"id,omitempty"`
	Associatednetworkid       string `json:"associatednetworkid,omitempty"`
	Ipaddress                 string `json:"ipaddress,omitempty"`
	Zonename                  string `json:"zonename,omitempty"`
	Zoneid                    string `json:"zoneid,omitempty"`
	Issourcenat               bool   `json:"issourcenat,omitempty"`
	Isportable                bool   `json:"isportable,omitempty"`
}

type UpdateNetworkParams struct {
	p map[string]interface{}
}

func (p *UpdateNetworkParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["changecidr"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("changecidr", vv)
	}
	if v, found := p.p["displaynetwork"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("displaynetwork", vv)
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["guestvmcidr"]; found {
		u.Set("guestvmcidr", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["networkdomain"]; found {
		u.Set("networkdomain", v.(string))
	}
	if v, found := p.p["networkofferingid"]; found {
		u.Set("networkofferingid", v.(string))
	}
	return u
}

func (p *UpdateNetworkParams) SetChangecidr(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["changecidr"] = v
	return
}

func (p *UpdateNetworkParams) SetDisplaynetwork(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaynetwork"] = v
	return
}

func (p *UpdateNetworkParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
	return
}

func (p *UpdateNetworkParams) SetGuestvmcidr(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["guestvmcidr"] = v
	return
}

func (p *UpdateNetworkParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *UpdateNetworkParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *UpdateNetworkParams) SetNetworkdomain(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkdomain"] = v
	return
}

func (p *UpdateNetworkParams) SetNetworkofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkofferingid"] = v
	return
}

// You should always use this function to get a new UpdateNetworkParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewUpdateNetworkParams(id string) *UpdateNetworkParams {
	p := &UpdateNetworkParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a network
func (s *NetworkService) UpdateNetwork(p *UpdateNetworkParams) (*UpdateNetworkResponse, error) {
	resp, err := s.cs.newRequest("updateNetwork", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateNetworkResponse
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

		var r UpdateNetworkResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type UpdateNetworkResponse struct {
	JobID                       string `json:"jobid,omitempty"`
	Displaynetwork              bool   `json:"displaynetwork,omitempty"`
	Traffictype                 string `json:"traffictype,omitempty"`
	Networkcidr                 string `json:"networkcidr,omitempty"`
	Networkofferingname         string `json:"networkofferingname,omitempty"`
	Broadcasturi                string `json:"broadcasturi,omitempty"`
	Dns1                        string `json:"dns1,omitempty"`
	Dns2                        string `json:"dns2,omitempty"`
	Vlan                        string `json:"vlan,omitempty"`
	Ip6cidr                     string `json:"ip6cidr,omitempty"`
	Networkofferingconservemode bool   `json:"networkofferingconservemode,omitempty"`
	Ip6gateway                  string `json:"ip6gateway,omitempty"`
	Specifyipranges             bool   `json:"specifyipranges,omitempty"`
	Netmask                     string `json:"netmask,omitempty"`
	Networkdomain               string `json:"networkdomain,omitempty"`
	Name                        string `json:"name,omitempty"`
	Domainid                    string `json:"domainid,omitempty"`
	Account                     string `json:"account,omitempty"`
	Project                     string `json:"project,omitempty"`
	State                       string `json:"state,omitempty"`
	Ispersistent                bool   `json:"ispersistent,omitempty"`
	Reservediprange             string `json:"reservediprange,omitempty"`
	Projectid                   string `json:"projectid,omitempty"`
	Cidr                        string `json:"cidr,omitempty"`
	Displaytext                 string `json:"displaytext,omitempty"`
	Issystem                    bool   `json:"issystem,omitempty"`
	Type                        string `json:"type,omitempty"`
	Zonename                    string `json:"zonename,omitempty"`
	Vpcid                       string `json:"vpcid,omitempty"`
	Subdomainaccess             bool   `json:"subdomainaccess,omitempty"`
	Zoneid                      string `json:"zoneid,omitempty"`
	Aclid                       string `json:"aclid,omitempty"`
	Isdefault                   bool   `json:"isdefault,omitempty"`
	Related                     string `json:"related,omitempty"`
	Networkofferingavailability string `json:"networkofferingavailability,omitempty"`
	Service                     []struct {
		Provider []struct {
			Servicelist                  []string `json:"servicelist,omitempty"`
			Id                           string   `json:"id,omitempty"`
			State                        string   `json:"state,omitempty"`
			Name                         string   `json:"name,omitempty"`
			Physicalnetworkid            string   `json:"physicalnetworkid,omitempty"`
			Canenableindividualservice   bool     `json:"canenableindividualservice,omitempty"`
			Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid,omitempty"`
		} `json:"provider,omitempty"`
		Capability []struct {
			Value                      string `json:"value,omitempty"`
			Canchooseservicecapability bool   `json:"canchooseservicecapability,omitempty"`
			Name                       string `json:"name,omitempty"`
		} `json:"capability,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"service,omitempty"`
	Gateway                    string `json:"gateway,omitempty"`
	Broadcastdomaintype        string `json:"broadcastdomaintype,omitempty"`
	Id                         string `json:"id,omitempty"`
	Restartrequired            bool   `json:"restartrequired,omitempty"`
	Acltype                    string `json:"acltype,omitempty"`
	Canusefordeploy            bool   `json:"canusefordeploy,omitempty"`
	Networkofferingdisplaytext string `json:"networkofferingdisplaytext,omitempty"`
	Physicalnetworkid          string `json:"physicalnetworkid,omitempty"`
	Domain                     string `json:"domain,omitempty"`
	Networkofferingid          string `json:"networkofferingid,omitempty"`
	Tags                       []struct {
		Resourceid   string `json:"resourceid,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Value        string `json:"value,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Account      string `json:"account,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Project      string `json:"project,omitempty"`
		Key          string `json:"key,omitempty"`
		Domain       string `json:"domain,omitempty"`
	} `json:"tags,omitempty"`
}

type CreatePhysicalNetworkParams struct {
	p map[string]interface{}
}

func (p *CreatePhysicalNetworkParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["broadcastdomainrange"]; found {
		u.Set("broadcastdomainrange", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["isolationmethods"]; found {
		vv := strings.Join(v.([]string), ", ")
		u.Set("isolationmethods", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["networkspeed"]; found {
		u.Set("networkspeed", v.(string))
	}
	if v, found := p.p["tags"]; found {
		vv := strings.Join(v.([]string), ", ")
		u.Set("tags", vv)
	}
	if v, found := p.p["vlan"]; found {
		u.Set("vlan", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *CreatePhysicalNetworkParams) SetBroadcastdomainrange(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["broadcastdomainrange"] = v
	return
}

func (p *CreatePhysicalNetworkParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *CreatePhysicalNetworkParams) SetIsolationmethods(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isolationmethods"] = v
	return
}

func (p *CreatePhysicalNetworkParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *CreatePhysicalNetworkParams) SetNetworkspeed(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkspeed"] = v
	return
}

func (p *CreatePhysicalNetworkParams) SetTags(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
	return
}

func (p *CreatePhysicalNetworkParams) SetVlan(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vlan"] = v
	return
}

func (p *CreatePhysicalNetworkParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new CreatePhysicalNetworkParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewCreatePhysicalNetworkParams(name string, zoneid string) *CreatePhysicalNetworkParams {
	p := &CreatePhysicalNetworkParams{}
	p.p = make(map[string]interface{})
	p.p["name"] = name
	p.p["zoneid"] = zoneid
	return p
}

// Creates a physical network
func (s *NetworkService) CreatePhysicalNetwork(p *CreatePhysicalNetworkParams) (*CreatePhysicalNetworkResponse, error) {
	resp, err := s.cs.newRequest("createPhysicalNetwork", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreatePhysicalNetworkResponse
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

		var r CreatePhysicalNetworkResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type CreatePhysicalNetworkResponse struct {
	JobID                string `json:"jobid,omitempty"`
	Isolationmethods     string `json:"isolationmethods,omitempty"`
	State                string `json:"state,omitempty"`
	Id                   string `json:"id,omitempty"`
	Broadcastdomainrange string `json:"broadcastdomainrange,omitempty"`
	Domainid             string `json:"domainid,omitempty"`
	Zoneid               string `json:"zoneid,omitempty"`
	Networkspeed         string `json:"networkspeed,omitempty"`
	Name                 string `json:"name,omitempty"`
	Tags                 string `json:"tags,omitempty"`
	Vlan                 string `json:"vlan,omitempty"`
}

type DeletePhysicalNetworkParams struct {
	p map[string]interface{}
}

func (p *DeletePhysicalNetworkParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeletePhysicalNetworkParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeletePhysicalNetworkParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewDeletePhysicalNetworkParams(id string) *DeletePhysicalNetworkParams {
	p := &DeletePhysicalNetworkParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a Physical Network.
func (s *NetworkService) DeletePhysicalNetwork(p *DeletePhysicalNetworkParams) (*DeletePhysicalNetworkResponse, error) {
	resp, err := s.cs.newRequest("deletePhysicalNetwork", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeletePhysicalNetworkResponse
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

		var r DeletePhysicalNetworkResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type DeletePhysicalNetworkResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type ListPhysicalNetworksParams struct {
	p map[string]interface{}
}

func (p *ListPhysicalNetworksParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
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
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListPhysicalNetworksParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListPhysicalNetworksParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListPhysicalNetworksParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *ListPhysicalNetworksParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListPhysicalNetworksParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListPhysicalNetworksParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new ListPhysicalNetworksParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewListPhysicalNetworksParams() *ListPhysicalNetworksParams {
	p := &ListPhysicalNetworksParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkService) GetPhysicalNetworkID(name string) (string, error) {
	p := &ListPhysicalNetworksParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	l, err := s.ListPhysicalNetworks(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, name, l)
	}
	return l.PhysicalNetworks[0].Id, nil
}

// Lists physical networks
func (s *NetworkService) ListPhysicalNetworks(p *ListPhysicalNetworksParams) (*ListPhysicalNetworksResponse, error) {
	resp, err := s.cs.newRequest("listPhysicalNetworks", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListPhysicalNetworksResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListPhysicalNetworksResponse struct {
	Count            int                `json:"count"`
	PhysicalNetworks []*PhysicalNetwork `json:"physicalnetwork"`
}

type PhysicalNetwork struct {
	Zoneid               string `json:"zoneid,omitempty"`
	Domainid             string `json:"domainid,omitempty"`
	State                string `json:"state,omitempty"`
	Name                 string `json:"name,omitempty"`
	Broadcastdomainrange string `json:"broadcastdomainrange,omitempty"`
	Networkspeed         string `json:"networkspeed,omitempty"`
	Id                   string `json:"id,omitempty"`
	Vlan                 string `json:"vlan,omitempty"`
	Isolationmethods     string `json:"isolationmethods,omitempty"`
	Tags                 string `json:"tags,omitempty"`
}

type UpdatePhysicalNetworkParams struct {
	p map[string]interface{}
}

func (p *UpdatePhysicalNetworkParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["networkspeed"]; found {
		u.Set("networkspeed", v.(string))
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := p.p["tags"]; found {
		vv := strings.Join(v.([]string), ", ")
		u.Set("tags", vv)
	}
	if v, found := p.p["vlan"]; found {
		u.Set("vlan", v.(string))
	}
	return u
}

func (p *UpdatePhysicalNetworkParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *UpdatePhysicalNetworkParams) SetNetworkspeed(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkspeed"] = v
	return
}

func (p *UpdatePhysicalNetworkParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
	return
}

func (p *UpdatePhysicalNetworkParams) SetTags(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
	return
}

func (p *UpdatePhysicalNetworkParams) SetVlan(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vlan"] = v
	return
}

// You should always use this function to get a new UpdatePhysicalNetworkParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewUpdatePhysicalNetworkParams(id string) *UpdatePhysicalNetworkParams {
	p := &UpdatePhysicalNetworkParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a physical network
func (s *NetworkService) UpdatePhysicalNetwork(p *UpdatePhysicalNetworkParams) (*UpdatePhysicalNetworkResponse, error) {
	resp, err := s.cs.newRequest("updatePhysicalNetwork", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdatePhysicalNetworkResponse
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

		var r UpdatePhysicalNetworkResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type UpdatePhysicalNetworkResponse struct {
	JobID                string `json:"jobid,omitempty"`
	Tags                 string `json:"tags,omitempty"`
	Name                 string `json:"name,omitempty"`
	Isolationmethods     string `json:"isolationmethods,omitempty"`
	Vlan                 string `json:"vlan,omitempty"`
	Networkspeed         string `json:"networkspeed,omitempty"`
	State                string `json:"state,omitempty"`
	Domainid             string `json:"domainid,omitempty"`
	Id                   string `json:"id,omitempty"`
	Zoneid               string `json:"zoneid,omitempty"`
	Broadcastdomainrange string `json:"broadcastdomainrange,omitempty"`
}

type ListSupportedNetworkServicesParams struct {
	p map[string]interface{}
}

func (p *ListSupportedNetworkServicesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("pagesize", vv)
	}
	if v, found := p.p["provider"]; found {
		u.Set("provider", v.(string))
	}
	if v, found := p.p["service"]; found {
		u.Set("service", v.(string))
	}
	return u
}

func (p *ListSupportedNetworkServicesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListSupportedNetworkServicesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListSupportedNetworkServicesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListSupportedNetworkServicesParams) SetProvider(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["provider"] = v
	return
}

func (p *ListSupportedNetworkServicesParams) SetService(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["service"] = v
	return
}

// You should always use this function to get a new ListSupportedNetworkServicesParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewListSupportedNetworkServicesParams() *ListSupportedNetworkServicesParams {
	p := &ListSupportedNetworkServicesParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists all network services provided by CloudStack or for the given Provider.
func (s *NetworkService) ListSupportedNetworkServices(p *ListSupportedNetworkServicesParams) (*ListSupportedNetworkServicesResponse, error) {
	resp, err := s.cs.newRequest("listSupportedNetworkServices", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListSupportedNetworkServicesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListSupportedNetworkServicesResponse struct {
	Count                    int                        `json:"count"`
	SupportedNetworkServices []*SupportedNetworkService `json:"supportednetworkservice"`
}

type SupportedNetworkService struct {
	Name       string `json:"name,omitempty"`
	Capability []struct {
		Value                      string `json:"value,omitempty"`
		Name                       string `json:"name,omitempty"`
		Canchooseservicecapability bool   `json:"canchooseservicecapability,omitempty"`
	} `json:"capability,omitempty"`
	Provider []struct {
		Name                         string   `json:"name,omitempty"`
		Id                           string   `json:"id,omitempty"`
		Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid,omitempty"`
		Physicalnetworkid            string   `json:"physicalnetworkid,omitempty"`
		Canenableindividualservice   bool     `json:"canenableindividualservice,omitempty"`
		State                        string   `json:"state,omitempty"`
		Servicelist                  []string `json:"servicelist,omitempty"`
	} `json:"provider,omitempty"`
}

type AddNetworkServiceProviderParams struct {
	p map[string]interface{}
}

func (p *AddNetworkServiceProviderParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["destinationphysicalnetworkid"]; found {
		u.Set("destinationphysicalnetworkid", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	if v, found := p.p["servicelist"]; found {
		vv := strings.Join(v.([]string), ", ")
		u.Set("servicelist", vv)
	}
	return u
}

func (p *AddNetworkServiceProviderParams) SetDestinationphysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["destinationphysicalnetworkid"] = v
	return
}

func (p *AddNetworkServiceProviderParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *AddNetworkServiceProviderParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
	return
}

func (p *AddNetworkServiceProviderParams) SetServicelist(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["servicelist"] = v
	return
}

// You should always use this function to get a new AddNetworkServiceProviderParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewAddNetworkServiceProviderParams(name string, physicalnetworkid string) *AddNetworkServiceProviderParams {
	p := &AddNetworkServiceProviderParams{}
	p.p = make(map[string]interface{})
	p.p["name"] = name
	p.p["physicalnetworkid"] = physicalnetworkid
	return p
}

// Adds a network serviceProvider to a physical network
func (s *NetworkService) AddNetworkServiceProvider(p *AddNetworkServiceProviderParams) (*AddNetworkServiceProviderResponse, error) {
	resp, err := s.cs.newRequest("addNetworkServiceProvider", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddNetworkServiceProviderResponse
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

		var r AddNetworkServiceProviderResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type AddNetworkServiceProviderResponse struct {
	JobID                        string   `json:"jobid,omitempty"`
	Canenableindividualservice   bool     `json:"canenableindividualservice,omitempty"`
	Servicelist                  []string `json:"servicelist,omitempty"`
	Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid,omitempty"`
	State                        string   `json:"state,omitempty"`
	Id                           string   `json:"id,omitempty"`
	Physicalnetworkid            string   `json:"physicalnetworkid,omitempty"`
	Name                         string   `json:"name,omitempty"`
}

type DeleteNetworkServiceProviderParams struct {
	p map[string]interface{}
}

func (p *DeleteNetworkServiceProviderParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteNetworkServiceProviderParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteNetworkServiceProviderParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewDeleteNetworkServiceProviderParams(id string) *DeleteNetworkServiceProviderParams {
	p := &DeleteNetworkServiceProviderParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a Network Service Provider.
func (s *NetworkService) DeleteNetworkServiceProvider(p *DeleteNetworkServiceProviderParams) (*DeleteNetworkServiceProviderResponse, error) {
	resp, err := s.cs.newRequest("deleteNetworkServiceProvider", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteNetworkServiceProviderResponse
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

		var r DeleteNetworkServiceProviderResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type DeleteNetworkServiceProviderResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Success     bool   `json:"success,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
}

type ListNetworkServiceProvidersParams struct {
	p map[string]interface{}
}

func (p *ListNetworkServiceProvidersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
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
	if v, found := p.p["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	return u
}

func (p *ListNetworkServiceProvidersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListNetworkServiceProvidersParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *ListNetworkServiceProvidersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListNetworkServiceProvidersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListNetworkServiceProvidersParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
	return
}

func (p *ListNetworkServiceProvidersParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
	return
}

// You should always use this function to get a new ListNetworkServiceProvidersParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewListNetworkServiceProvidersParams() *ListNetworkServiceProvidersParams {
	p := &ListNetworkServiceProvidersParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkService) GetNetworkServiceProviderID(name string) (string, error) {
	p := &ListNetworkServiceProvidersParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	l, err := s.ListNetworkServiceProviders(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, name, l)
	}
	return l.NetworkServiceProviders[0].Id, nil
}

// Lists network serviceproviders for a given physical network.
func (s *NetworkService) ListNetworkServiceProviders(p *ListNetworkServiceProvidersParams) (*ListNetworkServiceProvidersResponse, error) {
	resp, err := s.cs.newRequest("listNetworkServiceProviders", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListNetworkServiceProvidersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListNetworkServiceProvidersResponse struct {
	Count                   int                       `json:"count"`
	NetworkServiceProviders []*NetworkServiceProvider `json:"networkserviceprovider"`
}

type NetworkServiceProvider struct {
	Canenableindividualservice   bool     `json:"canenableindividualservice,omitempty"`
	Physicalnetworkid            string   `json:"physicalnetworkid,omitempty"`
	Servicelist                  []string `json:"servicelist,omitempty"`
	State                        string   `json:"state,omitempty"`
	Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid,omitempty"`
	Name                         string   `json:"name,omitempty"`
	Id                           string   `json:"id,omitempty"`
}

type UpdateNetworkServiceProviderParams struct {
	p map[string]interface{}
}

func (p *UpdateNetworkServiceProviderParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["servicelist"]; found {
		vv := strings.Join(v.([]string), ", ")
		u.Set("servicelist", vv)
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	return u
}

func (p *UpdateNetworkServiceProviderParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *UpdateNetworkServiceProviderParams) SetServicelist(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["servicelist"] = v
	return
}

func (p *UpdateNetworkServiceProviderParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
	return
}

// You should always use this function to get a new UpdateNetworkServiceProviderParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewUpdateNetworkServiceProviderParams(id string) *UpdateNetworkServiceProviderParams {
	p := &UpdateNetworkServiceProviderParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a network serviceProvider of a physical network
func (s *NetworkService) UpdateNetworkServiceProvider(p *UpdateNetworkServiceProviderParams) (*UpdateNetworkServiceProviderResponse, error) {
	resp, err := s.cs.newRequest("updateNetworkServiceProvider", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateNetworkServiceProviderResponse
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

		var r UpdateNetworkServiceProviderResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type UpdateNetworkServiceProviderResponse struct {
	JobID                        string   `json:"jobid,omitempty"`
	Name                         string   `json:"name,omitempty"`
	Id                           string   `json:"id,omitempty"`
	Physicalnetworkid            string   `json:"physicalnetworkid,omitempty"`
	Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid,omitempty"`
	State                        string   `json:"state,omitempty"`
	Servicelist                  []string `json:"servicelist,omitempty"`
	Canenableindividualservice   bool     `json:"canenableindividualservice,omitempty"`
}

type CreateStorageNetworkIpRangeParams struct {
	p map[string]interface{}
}

func (p *CreateStorageNetworkIpRangeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["endip"]; found {
		u.Set("endip", v.(string))
	}
	if v, found := p.p["gateway"]; found {
		u.Set("gateway", v.(string))
	}
	if v, found := p.p["netmask"]; found {
		u.Set("netmask", v.(string))
	}
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := p.p["startip"]; found {
		u.Set("startip", v.(string))
	}
	if v, found := p.p["vlan"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("vlan", vv)
	}
	return u
}

func (p *CreateStorageNetworkIpRangeParams) SetEndip(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["endip"] = v
	return
}

func (p *CreateStorageNetworkIpRangeParams) SetGateway(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gateway"] = v
	return
}

func (p *CreateStorageNetworkIpRangeParams) SetNetmask(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["netmask"] = v
	return
}

func (p *CreateStorageNetworkIpRangeParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
	return
}

func (p *CreateStorageNetworkIpRangeParams) SetStartip(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startip"] = v
	return
}

func (p *CreateStorageNetworkIpRangeParams) SetVlan(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vlan"] = v
	return
}

// You should always use this function to get a new CreateStorageNetworkIpRangeParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewCreateStorageNetworkIpRangeParams(gateway string, netmask string, podid string, startip string) *CreateStorageNetworkIpRangeParams {
	p := &CreateStorageNetworkIpRangeParams{}
	p.p = make(map[string]interface{})
	p.p["gateway"] = gateway
	p.p["netmask"] = netmask
	p.p["podid"] = podid
	p.p["startip"] = startip
	return p
}

// Creates a Storage network IP range.
func (s *NetworkService) CreateStorageNetworkIpRange(p *CreateStorageNetworkIpRangeParams) (*CreateStorageNetworkIpRangeResponse, error) {
	resp, err := s.cs.newRequest("createStorageNetworkIpRange", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateStorageNetworkIpRangeResponse
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

		var r CreateStorageNetworkIpRangeResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type CreateStorageNetworkIpRangeResponse struct {
	JobID     string `json:"jobid,omitempty"`
	Networkid string `json:"networkid,omitempty"`
	Netmask   string `json:"netmask,omitempty"`
	Podid     string `json:"podid,omitempty"`
	Startip   string `json:"startip,omitempty"`
	Vlan      int    `json:"vlan,omitempty"`
	Zoneid    string `json:"zoneid,omitempty"`
	Endip     string `json:"endip,omitempty"`
	Gateway   string `json:"gateway,omitempty"`
	Id        string `json:"id,omitempty"`
}

type DeleteStorageNetworkIpRangeParams struct {
	p map[string]interface{}
}

func (p *DeleteStorageNetworkIpRangeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteStorageNetworkIpRangeParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteStorageNetworkIpRangeParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewDeleteStorageNetworkIpRangeParams(id string) *DeleteStorageNetworkIpRangeParams {
	p := &DeleteStorageNetworkIpRangeParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a storage network IP Range.
func (s *NetworkService) DeleteStorageNetworkIpRange(p *DeleteStorageNetworkIpRangeParams) (*DeleteStorageNetworkIpRangeResponse, error) {
	resp, err := s.cs.newRequest("deleteStorageNetworkIpRange", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteStorageNetworkIpRangeResponse
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

		var r DeleteStorageNetworkIpRangeResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type DeleteStorageNetworkIpRangeResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type ListStorageNetworkIpRangeParams struct {
	p map[string]interface{}
}

func (p *ListStorageNetworkIpRangeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("pagesize", vv)
	}
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListStorageNetworkIpRangeParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListStorageNetworkIpRangeParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListStorageNetworkIpRangeParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListStorageNetworkIpRangeParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListStorageNetworkIpRangeParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
	return
}

func (p *ListStorageNetworkIpRangeParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new ListStorageNetworkIpRangeParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewListStorageNetworkIpRangeParams() *ListStorageNetworkIpRangeParams {
	p := &ListStorageNetworkIpRangeParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkService) GetStorageNetworkIpRangeID(keyword string) (string, error) {
	p := &ListStorageNetworkIpRangeParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword

	l, err := s.ListStorageNetworkIpRange(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, keyword, l)
	}
	return l.StorageNetworkIpRange[0].Id, nil
}

// List a storage network IP range.
func (s *NetworkService) ListStorageNetworkIpRange(p *ListStorageNetworkIpRangeParams) (*ListStorageNetworkIpRangeResponse, error) {
	resp, err := s.cs.newRequest("listStorageNetworkIpRange", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListStorageNetworkIpRangeResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListStorageNetworkIpRangeResponse struct {
	Count                 int                      `json:"count"`
	StorageNetworkIpRange []*StorageNetworkIpRange `json:"storagenetworkiprange"`
}

type StorageNetworkIpRange struct {
	Netmask   string `json:"netmask,omitempty"`
	Startip   string `json:"startip,omitempty"`
	Networkid string `json:"networkid,omitempty"`
	Vlan      int    `json:"vlan,omitempty"`
	Podid     string `json:"podid,omitempty"`
	Gateway   string `json:"gateway,omitempty"`
	Endip     string `json:"endip,omitempty"`
	Zoneid    string `json:"zoneid,omitempty"`
	Id        string `json:"id,omitempty"`
}

type UpdateStorageNetworkIpRangeParams struct {
	p map[string]interface{}
}

func (p *UpdateStorageNetworkIpRangeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["endip"]; found {
		u.Set("endip", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["netmask"]; found {
		u.Set("netmask", v.(string))
	}
	if v, found := p.p["startip"]; found {
		u.Set("startip", v.(string))
	}
	if v, found := p.p["vlan"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("vlan", vv)
	}
	return u
}

func (p *UpdateStorageNetworkIpRangeParams) SetEndip(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["endip"] = v
	return
}

func (p *UpdateStorageNetworkIpRangeParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *UpdateStorageNetworkIpRangeParams) SetNetmask(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["netmask"] = v
	return
}

func (p *UpdateStorageNetworkIpRangeParams) SetStartip(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startip"] = v
	return
}

func (p *UpdateStorageNetworkIpRangeParams) SetVlan(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vlan"] = v
	return
}

// You should always use this function to get a new UpdateStorageNetworkIpRangeParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewUpdateStorageNetworkIpRangeParams(id string) *UpdateStorageNetworkIpRangeParams {
	p := &UpdateStorageNetworkIpRangeParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Update a Storage network IP range, only allowed when no IPs in this range have been allocated.
func (s *NetworkService) UpdateStorageNetworkIpRange(p *UpdateStorageNetworkIpRangeParams) (*UpdateStorageNetworkIpRangeResponse, error) {
	resp, err := s.cs.newRequest("updateStorageNetworkIpRange", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateStorageNetworkIpRangeResponse
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

		var r UpdateStorageNetworkIpRangeResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type UpdateStorageNetworkIpRangeResponse struct {
	JobID     string `json:"jobid,omitempty"`
	Endip     string `json:"endip,omitempty"`
	Zoneid    string `json:"zoneid,omitempty"`
	Gateway   string `json:"gateway,omitempty"`
	Netmask   string `json:"netmask,omitempty"`
	Networkid string `json:"networkid,omitempty"`
	Podid     string `json:"podid,omitempty"`
	Vlan      int    `json:"vlan,omitempty"`
	Startip   string `json:"startip,omitempty"`
	Id        string `json:"id,omitempty"`
}

type ListF5LoadBalancerNetworksParams struct {
	p map[string]interface{}
}

func (p *ListF5LoadBalancerNetworksParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["lbdeviceid"]; found {
		u.Set("lbdeviceid", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("pagesize", vv)
	}
	return u
}

func (p *ListF5LoadBalancerNetworksParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListF5LoadBalancerNetworksParams) SetLbdeviceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbdeviceid"] = v
	return
}

func (p *ListF5LoadBalancerNetworksParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListF5LoadBalancerNetworksParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

// You should always use this function to get a new ListF5LoadBalancerNetworksParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewListF5LoadBalancerNetworksParams(lbdeviceid string) *ListF5LoadBalancerNetworksParams {
	p := &ListF5LoadBalancerNetworksParams{}
	p.p = make(map[string]interface{})
	p.p["lbdeviceid"] = lbdeviceid
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkService) GetF5LoadBalancerNetworkID(keyword string, lbdeviceid string) (string, error) {
	p := &ListF5LoadBalancerNetworksParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword
	p.p["lbdeviceid"] = lbdeviceid

	l, err := s.ListF5LoadBalancerNetworks(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, keyword, l)
	}
	return l.F5LoadBalancerNetworks[0].Id, nil
}

// lists network that are using a F5 load balancer device
func (s *NetworkService) ListF5LoadBalancerNetworks(p *ListF5LoadBalancerNetworksParams) (*ListF5LoadBalancerNetworksResponse, error) {
	resp, err := s.cs.newRequest("listF5LoadBalancerNetworks", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListF5LoadBalancerNetworksResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListF5LoadBalancerNetworksResponse struct {
	Count                  int                      `json:"count"`
	F5LoadBalancerNetworks []*F5LoadBalancerNetwork `json:"f5loadbalancernetwork"`
}

type F5LoadBalancerNetwork struct {
	Subdomainaccess             bool   `json:"subdomainaccess,omitempty"`
	Isdefault                   bool   `json:"isdefault,omitempty"`
	Networkcidr                 string `json:"networkcidr,omitempty"`
	Networkofferingavailability string `json:"networkofferingavailability,omitempty"`
	Ispersistent                bool   `json:"ispersistent,omitempty"`
	Restartrequired             bool   `json:"restartrequired,omitempty"`
	Project                     string `json:"project,omitempty"`
	Id                          string `json:"id,omitempty"`
	Issystem                    bool   `json:"issystem,omitempty"`
	Broadcasturi                string `json:"broadcasturi,omitempty"`
	Tags                        []struct {
		Value        string `json:"value,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Key          string `json:"key,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Project      string `json:"project,omitempty"`
		Account      string `json:"account,omitempty"`
	} `json:"tags,omitempty"`
	Traffictype                 string `json:"traffictype,omitempty"`
	Domain                      string `json:"domain,omitempty"`
	Networkdomain               string `json:"networkdomain,omitempty"`
	Networkofferingname         string `json:"networkofferingname,omitempty"`
	Acltype                     string `json:"acltype,omitempty"`
	Networkofferingconservemode bool   `json:"networkofferingconservemode,omitempty"`
	Vpcid                       string `json:"vpcid,omitempty"`
	Broadcastdomaintype         string `json:"broadcastdomaintype,omitempty"`
	Displaynetwork              bool   `json:"displaynetwork,omitempty"`
	Zoneid                      string `json:"zoneid,omitempty"`
	Specifyipranges             bool   `json:"specifyipranges,omitempty"`
	Dns2                        string `json:"dns2,omitempty"`
	Canusefordeploy             bool   `json:"canusefordeploy,omitempty"`
	Cidr                        string `json:"cidr,omitempty"`
	Name                        string `json:"name,omitempty"`
	Ip6cidr                     string `json:"ip6cidr,omitempty"`
	Projectid                   string `json:"projectid,omitempty"`
	Gateway                     string `json:"gateway,omitempty"`
	Service                     []struct {
		Capability []struct {
			Name                       string `json:"name,omitempty"`
			Canchooseservicecapability bool   `json:"canchooseservicecapability,omitempty"`
			Value                      string `json:"value,omitempty"`
		} `json:"capability,omitempty"`
		Name     string `json:"name,omitempty"`
		Provider []struct {
			Servicelist                  []string `json:"servicelist,omitempty"`
			Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid,omitempty"`
			Id                           string   `json:"id,omitempty"`
			Name                         string   `json:"name,omitempty"`
			Canenableindividualservice   bool     `json:"canenableindividualservice,omitempty"`
			State                        string   `json:"state,omitempty"`
			Physicalnetworkid            string   `json:"physicalnetworkid,omitempty"`
		} `json:"provider,omitempty"`
	} `json:"service,omitempty"`
	Zonename                   string `json:"zonename,omitempty"`
	State                      string `json:"state,omitempty"`
	Physicalnetworkid          string `json:"physicalnetworkid,omitempty"`
	Related                    string `json:"related,omitempty"`
	Dns1                       string `json:"dns1,omitempty"`
	Domainid                   string `json:"domainid,omitempty"`
	Netmask                    string `json:"netmask,omitempty"`
	Type                       string `json:"type,omitempty"`
	Aclid                      string `json:"aclid,omitempty"`
	Reservediprange            string `json:"reservediprange,omitempty"`
	Networkofferingid          string `json:"networkofferingid,omitempty"`
	Vlan                       string `json:"vlan,omitempty"`
	Ip6gateway                 string `json:"ip6gateway,omitempty"`
	Account                    string `json:"account,omitempty"`
	Displaytext                string `json:"displaytext,omitempty"`
	Networkofferingdisplaytext string `json:"networkofferingdisplaytext,omitempty"`
}

type ListSrxFirewallNetworksParams struct {
	p map[string]interface{}
}

func (p *ListSrxFirewallNetworksParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["lbdeviceid"]; found {
		u.Set("lbdeviceid", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("pagesize", vv)
	}
	return u
}

func (p *ListSrxFirewallNetworksParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListSrxFirewallNetworksParams) SetLbdeviceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbdeviceid"] = v
	return
}

func (p *ListSrxFirewallNetworksParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListSrxFirewallNetworksParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

// You should always use this function to get a new ListSrxFirewallNetworksParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewListSrxFirewallNetworksParams(lbdeviceid string) *ListSrxFirewallNetworksParams {
	p := &ListSrxFirewallNetworksParams{}
	p.p = make(map[string]interface{})
	p.p["lbdeviceid"] = lbdeviceid
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkService) GetSrxFirewallNetworkID(keyword string, lbdeviceid string) (string, error) {
	p := &ListSrxFirewallNetworksParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword
	p.p["lbdeviceid"] = lbdeviceid

	l, err := s.ListSrxFirewallNetworks(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, keyword, l)
	}
	return l.SrxFirewallNetworks[0].Id, nil
}

// lists network that are using SRX firewall device
func (s *NetworkService) ListSrxFirewallNetworks(p *ListSrxFirewallNetworksParams) (*ListSrxFirewallNetworksResponse, error) {
	resp, err := s.cs.newRequest("listSrxFirewallNetworks", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListSrxFirewallNetworksResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListSrxFirewallNetworksResponse struct {
	Count               int                   `json:"count"`
	SrxFirewallNetworks []*SrxFirewallNetwork `json:"srxfirewallnetwork"`
}

type SrxFirewallNetwork struct {
	Vlan                string `json:"vlan,omitempty"`
	Physicalnetworkid   string `json:"physicalnetworkid,omitempty"`
	Zoneid              string `json:"zoneid,omitempty"`
	Displaytext         string `json:"displaytext,omitempty"`
	Networkdomain       string `json:"networkdomain,omitempty"`
	Id                  string `json:"id,omitempty"`
	Networkofferingname string `json:"networkofferingname,omitempty"`
	Canusefordeploy     bool   `json:"canusefordeploy,omitempty"`
	Networkcidr         string `json:"networkcidr,omitempty"`
	Tags                []struct {
		Account      string `json:"account,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Project      string `json:"project,omitempty"`
		Value        string `json:"value,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Key          string `json:"key,omitempty"`
	} `json:"tags,omitempty"`
	Project                     string `json:"project,omitempty"`
	Specifyipranges             bool   `json:"specifyipranges,omitempty"`
	Networkofferingavailability string `json:"networkofferingavailability,omitempty"`
	Broadcastdomaintype         string `json:"broadcastdomaintype,omitempty"`
	Aclid                       string `json:"aclid,omitempty"`
	Isdefault                   bool   `json:"isdefault,omitempty"`
	Traffictype                 string `json:"traffictype,omitempty"`
	Name                        string `json:"name,omitempty"`
	Type                        string `json:"type,omitempty"`
	Projectid                   string `json:"projectid,omitempty"`
	Ip6gateway                  string `json:"ip6gateway,omitempty"`
	Service                     []struct {
		Capability []struct {
			Name                       string `json:"name,omitempty"`
			Value                      string `json:"value,omitempty"`
			Canchooseservicecapability bool   `json:"canchooseservicecapability,omitempty"`
		} `json:"capability,omitempty"`
		Provider []struct {
			Servicelist                  []string `json:"servicelist,omitempty"`
			Physicalnetworkid            string   `json:"physicalnetworkid,omitempty"`
			Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid,omitempty"`
			State                        string   `json:"state,omitempty"`
			Canenableindividualservice   bool     `json:"canenableindividualservice,omitempty"`
			Name                         string   `json:"name,omitempty"`
			Id                           string   `json:"id,omitempty"`
		} `json:"provider,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"service,omitempty"`
	Gateway                     string `json:"gateway,omitempty"`
	Networkofferingconservemode bool   `json:"networkofferingconservemode,omitempty"`
	Netmask                     string `json:"netmask,omitempty"`
	Cidr                        string `json:"cidr,omitempty"`
	Issystem                    bool   `json:"issystem,omitempty"`
	Displaynetwork              bool   `json:"displaynetwork,omitempty"`
	Subdomainaccess             bool   `json:"subdomainaccess,omitempty"`
	Zonename                    string `json:"zonename,omitempty"`
	Ip6cidr                     string `json:"ip6cidr,omitempty"`
	Dns1                        string `json:"dns1,omitempty"`
	Vpcid                       string `json:"vpcid,omitempty"`
	Broadcasturi                string `json:"broadcasturi,omitempty"`
	Domain                      string `json:"domain,omitempty"`
	Dns2                        string `json:"dns2,omitempty"`
	Reservediprange             string `json:"reservediprange,omitempty"`
	Restartrequired             bool   `json:"restartrequired,omitempty"`
	Networkofferingdisplaytext  string `json:"networkofferingdisplaytext,omitempty"`
	State                       string `json:"state,omitempty"`
	Ispersistent                bool   `json:"ispersistent,omitempty"`
	Networkofferingid           string `json:"networkofferingid,omitempty"`
	Related                     string `json:"related,omitempty"`
	Account                     string `json:"account,omitempty"`
	Acltype                     string `json:"acltype,omitempty"`
	Domainid                    string `json:"domainid,omitempty"`
}

type ListPaloAltoFirewallNetworksParams struct {
	p map[string]interface{}
}

func (p *ListPaloAltoFirewallNetworksParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["lbdeviceid"]; found {
		u.Set("lbdeviceid", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("pagesize", vv)
	}
	return u
}

func (p *ListPaloAltoFirewallNetworksParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListPaloAltoFirewallNetworksParams) SetLbdeviceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbdeviceid"] = v
	return
}

func (p *ListPaloAltoFirewallNetworksParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListPaloAltoFirewallNetworksParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

// You should always use this function to get a new ListPaloAltoFirewallNetworksParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewListPaloAltoFirewallNetworksParams(lbdeviceid string) *ListPaloAltoFirewallNetworksParams {
	p := &ListPaloAltoFirewallNetworksParams{}
	p.p = make(map[string]interface{})
	p.p["lbdeviceid"] = lbdeviceid
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkService) GetPaloAltoFirewallNetworkID(keyword string, lbdeviceid string) (string, error) {
	p := &ListPaloAltoFirewallNetworksParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword
	p.p["lbdeviceid"] = lbdeviceid

	l, err := s.ListPaloAltoFirewallNetworks(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, keyword, l)
	}
	return l.PaloAltoFirewallNetworks[0].Id, nil
}

// lists network that are using Palo Alto firewall device
func (s *NetworkService) ListPaloAltoFirewallNetworks(p *ListPaloAltoFirewallNetworksParams) (*ListPaloAltoFirewallNetworksResponse, error) {
	resp, err := s.cs.newRequest("listPaloAltoFirewallNetworks", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListPaloAltoFirewallNetworksResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListPaloAltoFirewallNetworksResponse struct {
	Count                    int                        `json:"count"`
	PaloAltoFirewallNetworks []*PaloAltoFirewallNetwork `json:"paloaltofirewallnetwork"`
}

type PaloAltoFirewallNetwork struct {
	Vlan                       string `json:"vlan,omitempty"`
	Subdomainaccess            bool   `json:"subdomainaccess,omitempty"`
	Project                    string `json:"project,omitempty"`
	Type                       string `json:"type,omitempty"`
	Networkofferingdisplaytext string `json:"networkofferingdisplaytext,omitempty"`
	Restartrequired            bool   `json:"restartrequired,omitempty"`
	Networkofferingid          string `json:"networkofferingid,omitempty"`
	Acltype                    string `json:"acltype,omitempty"`
	Tags                       []struct {
		Account      string `json:"account,omitempty"`
		Project      string `json:"project,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Value        string `json:"value,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Key          string `json:"key,omitempty"`
	} `json:"tags,omitempty"`
	Specifyipranges             bool   `json:"specifyipranges,omitempty"`
	Name                        string `json:"name,omitempty"`
	Related                     string `json:"related,omitempty"`
	Isdefault                   bool   `json:"isdefault,omitempty"`
	Netmask                     string `json:"netmask,omitempty"`
	Displaynetwork              bool   `json:"displaynetwork,omitempty"`
	Dns2                        string `json:"dns2,omitempty"`
	Vpcid                       string `json:"vpcid,omitempty"`
	Traffictype                 string `json:"traffictype,omitempty"`
	Displaytext                 string `json:"displaytext,omitempty"`
	Domainid                    string `json:"domainid,omitempty"`
	Dns1                        string `json:"dns1,omitempty"`
	Physicalnetworkid           string `json:"physicalnetworkid,omitempty"`
	Gateway                     string `json:"gateway,omitempty"`
	Aclid                       string `json:"aclid,omitempty"`
	Ip6cidr                     string `json:"ip6cidr,omitempty"`
	Account                     string `json:"account,omitempty"`
	Projectid                   string `json:"projectid,omitempty"`
	Issystem                    bool   `json:"issystem,omitempty"`
	State                       string `json:"state,omitempty"`
	Zonename                    string `json:"zonename,omitempty"`
	Cidr                        string `json:"cidr,omitempty"`
	Ispersistent                bool   `json:"ispersistent,omitempty"`
	Canusefordeploy             bool   `json:"canusefordeploy,omitempty"`
	Networkofferingconservemode bool   `json:"networkofferingconservemode,omitempty"`
	Reservediprange             string `json:"reservediprange,omitempty"`
	Service                     []struct {
		Provider []struct {
			State                        string   `json:"state,omitempty"`
			Name                         string   `json:"name,omitempty"`
			Canenableindividualservice   bool     `json:"canenableindividualservice,omitempty"`
			Id                           string   `json:"id,omitempty"`
			Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid,omitempty"`
			Physicalnetworkid            string   `json:"physicalnetworkid,omitempty"`
			Servicelist                  []string `json:"servicelist,omitempty"`
		} `json:"provider,omitempty"`
		Name       string `json:"name,omitempty"`
		Capability []struct {
			Canchooseservicecapability bool   `json:"canchooseservicecapability,omitempty"`
			Name                       string `json:"name,omitempty"`
			Value                      string `json:"value,omitempty"`
		} `json:"capability,omitempty"`
	} `json:"service,omitempty"`
	Domain                      string `json:"domain,omitempty"`
	Broadcastdomaintype         string `json:"broadcastdomaintype,omitempty"`
	Networkcidr                 string `json:"networkcidr,omitempty"`
	Networkofferingname         string `json:"networkofferingname,omitempty"`
	Broadcasturi                string `json:"broadcasturi,omitempty"`
	Networkofferingavailability string `json:"networkofferingavailability,omitempty"`
	Zoneid                      string `json:"zoneid,omitempty"`
	Ip6gateway                  string `json:"ip6gateway,omitempty"`
	Id                          string `json:"id,omitempty"`
	Networkdomain               string `json:"networkdomain,omitempty"`
}

type ListNetscalerLoadBalancerNetworksParams struct {
	p map[string]interface{}
}

func (p *ListNetscalerLoadBalancerNetworksParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["lbdeviceid"]; found {
		u.Set("lbdeviceid", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("pagesize", vv)
	}
	return u
}

func (p *ListNetscalerLoadBalancerNetworksParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListNetscalerLoadBalancerNetworksParams) SetLbdeviceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbdeviceid"] = v
	return
}

func (p *ListNetscalerLoadBalancerNetworksParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListNetscalerLoadBalancerNetworksParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

// You should always use this function to get a new ListNetscalerLoadBalancerNetworksParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewListNetscalerLoadBalancerNetworksParams(lbdeviceid string) *ListNetscalerLoadBalancerNetworksParams {
	p := &ListNetscalerLoadBalancerNetworksParams{}
	p.p = make(map[string]interface{})
	p.p["lbdeviceid"] = lbdeviceid
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkService) GetNetscalerLoadBalancerNetworkID(keyword string, lbdeviceid string) (string, error) {
	p := &ListNetscalerLoadBalancerNetworksParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword
	p.p["lbdeviceid"] = lbdeviceid

	l, err := s.ListNetscalerLoadBalancerNetworks(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, keyword, l)
	}
	return l.NetscalerLoadBalancerNetworks[0].Id, nil
}

// lists network that are using a netscaler load balancer device
func (s *NetworkService) ListNetscalerLoadBalancerNetworks(p *ListNetscalerLoadBalancerNetworksParams) (*ListNetscalerLoadBalancerNetworksResponse, error) {
	resp, err := s.cs.newRequest("listNetscalerLoadBalancerNetworks", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListNetscalerLoadBalancerNetworksResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListNetscalerLoadBalancerNetworksResponse struct {
	Count                         int                             `json:"count"`
	NetscalerLoadBalancerNetworks []*NetscalerLoadBalancerNetwork `json:"netscalerloadbalancernetwork"`
}

type NetscalerLoadBalancerNetwork struct {
	Physicalnetworkid           string `json:"physicalnetworkid,omitempty"`
	Domainid                    string `json:"domainid,omitempty"`
	Aclid                       string `json:"aclid,omitempty"`
	Netmask                     string `json:"netmask,omitempty"`
	Zoneid                      string `json:"zoneid,omitempty"`
	Ip6gateway                  string `json:"ip6gateway,omitempty"`
	Projectid                   string `json:"projectid,omitempty"`
	Acltype                     string `json:"acltype,omitempty"`
	Broadcastdomaintype         string `json:"broadcastdomaintype,omitempty"`
	Networkofferingavailability string `json:"networkofferingavailability,omitempty"`
	Zonename                    string `json:"zonename,omitempty"`
	Networkcidr                 string `json:"networkcidr,omitempty"`
	Networkofferingconservemode bool   `json:"networkofferingconservemode,omitempty"`
	Networkdomain               string `json:"networkdomain,omitempty"`
	Networkofferingdisplaytext  string `json:"networkofferingdisplaytext,omitempty"`
	Cidr                        string `json:"cidr,omitempty"`
	Restartrequired             bool   `json:"restartrequired,omitempty"`
	Networkofferingid           string `json:"networkofferingid,omitempty"`
	Ip6cidr                     string `json:"ip6cidr,omitempty"`
	Id                          string `json:"id,omitempty"`
	Displaynetwork              bool   `json:"displaynetwork,omitempty"`
	Specifyipranges             bool   `json:"specifyipranges,omitempty"`
	Domain                      string `json:"domain,omitempty"`
	Related                     string `json:"related,omitempty"`
	Account                     string `json:"account,omitempty"`
	Vlan                        string `json:"vlan,omitempty"`
	Traffictype                 string `json:"traffictype,omitempty"`
	Isdefault                   bool   `json:"isdefault,omitempty"`
	Subdomainaccess             bool   `json:"subdomainaccess,omitempty"`
	Project                     string `json:"project,omitempty"`
	Dns2                        string `json:"dns2,omitempty"`
	Name                        string `json:"name,omitempty"`
	Issystem                    bool   `json:"issystem,omitempty"`
	Ispersistent                bool   `json:"ispersistent,omitempty"`
	Reservediprange             string `json:"reservediprange,omitempty"`
	Displaytext                 string `json:"displaytext,omitempty"`
	Broadcasturi                string `json:"broadcasturi,omitempty"`
	Type                        string `json:"type,omitempty"`
	Dns1                        string `json:"dns1,omitempty"`
	Vpcid                       string `json:"vpcid,omitempty"`
	Networkofferingname         string `json:"networkofferingname,omitempty"`
	Gateway                     string `json:"gateway,omitempty"`
	Tags                        []struct {
		Project      string `json:"project,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Account      string `json:"account,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Value        string `json:"value,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Key          string `json:"key,omitempty"`
	} `json:"tags,omitempty"`
	Canusefordeploy bool `json:"canusefordeploy,omitempty"`
	Service         []struct {
		Name       string `json:"name,omitempty"`
		Capability []struct {
			Canchooseservicecapability bool   `json:"canchooseservicecapability,omitempty"`
			Value                      string `json:"value,omitempty"`
			Name                       string `json:"name,omitempty"`
		} `json:"capability,omitempty"`
		Provider []struct {
			Name                         string   `json:"name,omitempty"`
			Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid,omitempty"`
			Physicalnetworkid            string   `json:"physicalnetworkid,omitempty"`
			Canenableindividualservice   bool     `json:"canenableindividualservice,omitempty"`
			Servicelist                  []string `json:"servicelist,omitempty"`
			Id                           string   `json:"id,omitempty"`
			State                        string   `json:"state,omitempty"`
		} `json:"provider,omitempty"`
	} `json:"service,omitempty"`
	State string `json:"state,omitempty"`
}

type ListNiciraNvpDeviceNetworksParams struct {
	p map[string]interface{}
}

func (p *ListNiciraNvpDeviceNetworksParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["nvpdeviceid"]; found {
		u.Set("nvpdeviceid", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("pagesize", vv)
	}
	return u
}

func (p *ListNiciraNvpDeviceNetworksParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListNiciraNvpDeviceNetworksParams) SetNvpdeviceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nvpdeviceid"] = v
	return
}

func (p *ListNiciraNvpDeviceNetworksParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListNiciraNvpDeviceNetworksParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

// You should always use this function to get a new ListNiciraNvpDeviceNetworksParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewListNiciraNvpDeviceNetworksParams(nvpdeviceid string) *ListNiciraNvpDeviceNetworksParams {
	p := &ListNiciraNvpDeviceNetworksParams{}
	p.p = make(map[string]interface{})
	p.p["nvpdeviceid"] = nvpdeviceid
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkService) GetNiciraNvpDeviceNetworkID(keyword string, nvpdeviceid string) (string, error) {
	p := &ListNiciraNvpDeviceNetworksParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword
	p.p["nvpdeviceid"] = nvpdeviceid

	l, err := s.ListNiciraNvpDeviceNetworks(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, keyword, l)
	}
	return l.NiciraNvpDeviceNetworks[0].Id, nil
}

// lists network that are using a nicira nvp device
func (s *NetworkService) ListNiciraNvpDeviceNetworks(p *ListNiciraNvpDeviceNetworksParams) (*ListNiciraNvpDeviceNetworksResponse, error) {
	resp, err := s.cs.newRequest("listNiciraNvpDeviceNetworks", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListNiciraNvpDeviceNetworksResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListNiciraNvpDeviceNetworksResponse struct {
	Count                   int                       `json:"count"`
	NiciraNvpDeviceNetworks []*NiciraNvpDeviceNetwork `json:"niciranvpdevicenetwork"`
}

type NiciraNvpDeviceNetwork struct {
	Ispersistent                bool   `json:"ispersistent,omitempty"`
	Networkofferingdisplaytext  string `json:"networkofferingdisplaytext,omitempty"`
	Cidr                        string `json:"cidr,omitempty"`
	Type                        string `json:"type,omitempty"`
	Traffictype                 string `json:"traffictype,omitempty"`
	Physicalnetworkid           string `json:"physicalnetworkid,omitempty"`
	Acltype                     string `json:"acltype,omitempty"`
	Specifyipranges             bool   `json:"specifyipranges,omitempty"`
	Ip6gateway                  string `json:"ip6gateway,omitempty"`
	Dns2                        string `json:"dns2,omitempty"`
	Networkofferingname         string `json:"networkofferingname,omitempty"`
	Ip6cidr                     string `json:"ip6cidr,omitempty"`
	Networkofferingconservemode bool   `json:"networkofferingconservemode,omitempty"`
	Dns1                        string `json:"dns1,omitempty"`
	Reservediprange             string `json:"reservediprange,omitempty"`
	Zonename                    string `json:"zonename,omitempty"`
	Networkdomain               string `json:"networkdomain,omitempty"`
	Networkofferingavailability string `json:"networkofferingavailability,omitempty"`
	Aclid                       string `json:"aclid,omitempty"`
	Domain                      string `json:"domain,omitempty"`
	Service                     []struct {
		Provider []struct {
			Canenableindividualservice   bool     `json:"canenableindividualservice,omitempty"`
			Id                           string   `json:"id,omitempty"`
			Physicalnetworkid            string   `json:"physicalnetworkid,omitempty"`
			Servicelist                  []string `json:"servicelist,omitempty"`
			Name                         string   `json:"name,omitempty"`
			State                        string   `json:"state,omitempty"`
			Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid,omitempty"`
		} `json:"provider,omitempty"`
		Capability []struct {
			Canchooseservicecapability bool   `json:"canchooseservicecapability,omitempty"`
			Value                      string `json:"value,omitempty"`
			Name                       string `json:"name,omitempty"`
		} `json:"capability,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"service,omitempty"`
	Related             string `json:"related,omitempty"`
	Vlan                string `json:"vlan,omitempty"`
	Isdefault           bool   `json:"isdefault,omitempty"`
	Broadcastdomaintype string `json:"broadcastdomaintype,omitempty"`
	Tags                []struct {
		Key          string `json:"key,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Account      string `json:"account,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Project      string `json:"project,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Value        string `json:"value,omitempty"`
	} `json:"tags,omitempty"`
	Displaytext       string `json:"displaytext,omitempty"`
	Networkofferingid string `json:"networkofferingid,omitempty"`
	Subdomainaccess   bool   `json:"subdomainaccess,omitempty"`
	Zoneid            string `json:"zoneid,omitempty"`
	Name              string `json:"name,omitempty"`
	Id                string `json:"id,omitempty"`
	Netmask           string `json:"netmask,omitempty"`
	Projectid         string `json:"projectid,omitempty"`
	Canusefordeploy   bool   `json:"canusefordeploy,omitempty"`
	Networkcidr       string `json:"networkcidr,omitempty"`
	Issystem          bool   `json:"issystem,omitempty"`
	Displaynetwork    bool   `json:"displaynetwork,omitempty"`
	State             string `json:"state,omitempty"`
	Broadcasturi      string `json:"broadcasturi,omitempty"`
	Restartrequired   bool   `json:"restartrequired,omitempty"`
	Vpcid             string `json:"vpcid,omitempty"`
	Domainid          string `json:"domainid,omitempty"`
	Project           string `json:"project,omitempty"`
	Account           string `json:"account,omitempty"`
	Gateway           string `json:"gateway,omitempty"`
}

type ListNetworkIsolationMethodsParams struct {
	p map[string]interface{}
}

func (p *ListNetworkIsolationMethodsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("pagesize", vv)
	}
	return u
}

func (p *ListNetworkIsolationMethodsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListNetworkIsolationMethodsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListNetworkIsolationMethodsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

// You should always use this function to get a new ListNetworkIsolationMethodsParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewListNetworkIsolationMethodsParams() *ListNetworkIsolationMethodsParams {
	p := &ListNetworkIsolationMethodsParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists supported methods of network isolation
func (s *NetworkService) ListNetworkIsolationMethods(p *ListNetworkIsolationMethodsParams) (*ListNetworkIsolationMethodsResponse, error) {
	resp, err := s.cs.newRequest("listNetworkIsolationMethods", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListNetworkIsolationMethodsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListNetworkIsolationMethodsResponse struct {
	Count                   int                       `json:"count"`
	NetworkIsolationMethods []*NetworkIsolationMethod `json:"networkisolationmethod"`
}

type NetworkIsolationMethod struct {
	Name string `json:"name,omitempty"`
}
