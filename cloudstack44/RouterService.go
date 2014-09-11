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
)

type StartRouterParams struct {
	p map[string]interface{}
}

func (p *StartRouterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *StartRouterParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new StartRouterParams instance,
// as then you are sure you have configured all required params
func (s *RouterService) NewStartRouterParams(id string) *StartRouterParams {
	p := &StartRouterParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Starts a router.
func (s *RouterService) StartRouter(p *StartRouterParams) (*StartRouterResponse, error) {
	resp, err := s.cs.newRequest("startRouter", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r StartRouterResponse
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

		var r StartRouterResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type StartRouterResponse struct {
	JobID               string `json:"jobid,omitempty"`
	Linklocalnetmask    string `json:"linklocalnetmask,omitempty"`
	Gateway             string `json:"gateway,omitempty"`
	Id                  string `json:"id,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Podid               string `json:"podid,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
	Hostname            string `json:"hostname,omitempty"`
	Scriptsversion      string `json:"scriptsversion,omitempty"`
	Publicnetmask       string `json:"publicnetmask,omitempty"`
	Name                string `json:"name,omitempty"`
	Networkdomain       string `json:"networkdomain,omitempty"`
	Publicmacaddress    string `json:"publicmacaddress,omitempty"`
	Isredundantrouter   bool   `json:"isredundantrouter,omitempty"`
	State               string `json:"state,omitempty"`
	Guestnetworkid      string `json:"guestnetworkid,omitempty"`
	Templateid          string `json:"templateid,omitempty"`
	Zoneid              string `json:"zoneid,omitempty"`
	Dns2                string `json:"dns2,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Dns1                string `json:"dns1,omitempty"`
	Requiresupgrade     bool   `json:"requiresupgrade,omitempty"`
	Guestnetmask        string `json:"guestnetmask,omitempty"`
	Nic                 []struct {
		Type         string   `json:"type,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Id           string   `json:"id,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
	} `json:"nic,omitempty"`
	Domainid            string `json:"domainid,omitempty"`
	Ip6dns1             string `json:"ip6dns1,omitempty"`
	Projectid           string `json:"projectid,omitempty"`
	Created             string `json:"created,omitempty"`
	Linklocalnetworkid  string `json:"linklocalnetworkid,omitempty"`
	Guestipaddress      string `json:"guestipaddress,omitempty"`
	Role                string `json:"role,omitempty"`
	Hostid              string `json:"hostid,omitempty"`
	Linklocalip         string `json:"linklocalip,omitempty"`
	Vpcid               string `json:"vpcid,omitempty"`
	Redundantstate      string `json:"redundantstate,omitempty"`
	Project             string `json:"project,omitempty"`
	Publicip            string `json:"publicip,omitempty"`
	Publicnetworkid     string `json:"publicnetworkid,omitempty"`
	Guestmacaddress     string `json:"guestmacaddress,omitempty"`
	Linklocalmacaddress string `json:"linklocalmacaddress,omitempty"`
	Ip6dns2             string `json:"ip6dns2,omitempty"`
	Version             string `json:"version,omitempty"`
	Account             string `json:"account,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
}

type RebootRouterParams struct {
	p map[string]interface{}
}

func (p *RebootRouterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *RebootRouterParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new RebootRouterParams instance,
// as then you are sure you have configured all required params
func (s *RouterService) NewRebootRouterParams(id string) *RebootRouterParams {
	p := &RebootRouterParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Starts a router.
func (s *RouterService) RebootRouter(p *RebootRouterParams) (*RebootRouterResponse, error) {
	resp, err := s.cs.newRequest("rebootRouter", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RebootRouterResponse
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

		var r RebootRouterResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type RebootRouterResponse struct {
	JobID              string `json:"jobid,omitempty"`
	Guestnetmask       string `json:"guestnetmask,omitempty"`
	Hostid             string `json:"hostid,omitempty"`
	Networkdomain      string `json:"networkdomain,omitempty"`
	Publicnetworkid    string `json:"publicnetworkid,omitempty"`
	Version            string `json:"version,omitempty"`
	Gateway            string `json:"gateway,omitempty"`
	Vpcid              string `json:"vpcid,omitempty"`
	Scriptsversion     string `json:"scriptsversion,omitempty"`
	Dns2               string `json:"dns2,omitempty"`
	Zoneid             string `json:"zoneid,omitempty"`
	Domainid           string `json:"domainid,omitempty"`
	Isredundantrouter  bool   `json:"isredundantrouter,omitempty"`
	Templateid         string `json:"templateid,omitempty"`
	Hostname           string `json:"hostname,omitempty"`
	Linklocalnetworkid string `json:"linklocalnetworkid,omitempty"`
	Guestnetworkid     string `json:"guestnetworkid,omitempty"`
	Redundantstate     string `json:"redundantstate,omitempty"`
	Ip6dns1            string `json:"ip6dns1,omitempty"`
	Serviceofferingid  string `json:"serviceofferingid,omitempty"`
	Linklocalnetmask   string `json:"linklocalnetmask,omitempty"`
	Projectid          string `json:"projectid,omitempty"`
	Created            string `json:"created,omitempty"`
	Guestmacaddress    string `json:"guestmacaddress,omitempty"`
	Nic                []struct {
		Traffictype  string   `json:"traffictype,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Id           string   `json:"id,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Type         string   `json:"type,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
	} `json:"nic,omitempty"`
	Ip6dns2             string `json:"ip6dns2,omitempty"`
	Linklocalip         string `json:"linklocalip,omitempty"`
	Guestipaddress      string `json:"guestipaddress,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
	Project             string `json:"project,omitempty"`
	Linklocalmacaddress string `json:"linklocalmacaddress,omitempty"`
	Requiresupgrade     bool   `json:"requiresupgrade,omitempty"`
	Dns1                string `json:"dns1,omitempty"`
	Name                string `json:"name,omitempty"`
	Publicmacaddress    string `json:"publicmacaddress,omitempty"`
	Role                string `json:"role,omitempty"`
	Account             string `json:"account,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	State               string `json:"state,omitempty"`
	Publicip            string `json:"publicip,omitempty"`
	Podid               string `json:"podid,omitempty"`
	Id                  string `json:"id,omitempty"`
	Publicnetmask       string `json:"publicnetmask,omitempty"`
}

type StopRouterParams struct {
	p map[string]interface{}
}

func (p *StopRouterParams) toURLValues() url.Values {
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

func (p *StopRouterParams) SetForced(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forced"] = v
	return
}

func (p *StopRouterParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new StopRouterParams instance,
// as then you are sure you have configured all required params
func (s *RouterService) NewStopRouterParams(id string) *StopRouterParams {
	p := &StopRouterParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Stops a router.
func (s *RouterService) StopRouter(p *StopRouterParams) (*StopRouterResponse, error) {
	resp, err := s.cs.newRequest("stopRouter", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r StopRouterResponse
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

		var r StopRouterResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type StopRouterResponse struct {
	JobID               string `json:"jobid,omitempty"`
	Dns2                string `json:"dns2,omitempty"`
	Linklocalnetmask    string `json:"linklocalnetmask,omitempty"`
	Zoneid              string `json:"zoneid,omitempty"`
	Templateid          string `json:"templateid,omitempty"`
	Created             string `json:"created,omitempty"`
	Guestmacaddress     string `json:"guestmacaddress,omitempty"`
	Linklocalip         string `json:"linklocalip,omitempty"`
	Scriptsversion      string `json:"scriptsversion,omitempty"`
	Projectid           string `json:"projectid,omitempty"`
	Podid               string `json:"podid,omitempty"`
	Ip6dns2             string `json:"ip6dns2,omitempty"`
	Linklocalmacaddress string `json:"linklocalmacaddress,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
	Account             string `json:"account,omitempty"`
	Guestnetworkid      string `json:"guestnetworkid,omitempty"`
	Publicmacaddress    string `json:"publicmacaddress,omitempty"`
	Version             string `json:"version,omitempty"`
	Publicip            string `json:"publicip,omitempty"`
	Nic                 []struct {
		Networkid    string   `json:"networkid,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Type         string   `json:"type,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Id           string   `json:"id,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
	} `json:"nic,omitempty"`
	Hostname            string `json:"hostname,omitempty"`
	Networkdomain       string `json:"networkdomain,omitempty"`
	Isredundantrouter   bool   `json:"isredundantrouter,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Vpcid               string `json:"vpcid,omitempty"`
	Guestipaddress      string `json:"guestipaddress,omitempty"`
	Name                string `json:"name,omitempty"`
	Role                string `json:"role,omitempty"`
	Publicnetmask       string `json:"publicnetmask,omitempty"`
	Publicnetworkid     string `json:"publicnetworkid,omitempty"`
	Guestnetmask        string `json:"guestnetmask,omitempty"`
	Gateway             string `json:"gateway,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
	Domainid            string `json:"domainid,omitempty"`
	State               string `json:"state,omitempty"`
	Id                  string `json:"id,omitempty"`
	Requiresupgrade     bool   `json:"requiresupgrade,omitempty"`
	Ip6dns1             string `json:"ip6dns1,omitempty"`
	Dns1                string `json:"dns1,omitempty"`
	Hostid              string `json:"hostid,omitempty"`
	Project             string `json:"project,omitempty"`
	Redundantstate      string `json:"redundantstate,omitempty"`
	Linklocalnetworkid  string `json:"linklocalnetworkid,omitempty"`
}

type DestroyRouterParams struct {
	p map[string]interface{}
}

func (p *DestroyRouterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DestroyRouterParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DestroyRouterParams instance,
// as then you are sure you have configured all required params
func (s *RouterService) NewDestroyRouterParams(id string) *DestroyRouterParams {
	p := &DestroyRouterParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Destroys a router.
func (s *RouterService) DestroyRouter(p *DestroyRouterParams) (*DestroyRouterResponse, error) {
	resp, err := s.cs.newRequest("destroyRouter", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DestroyRouterResponse
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

		var r DestroyRouterResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type DestroyRouterResponse struct {
	JobID         string `json:"jobid,omitempty"`
	Publicnetmask string `json:"publicnetmask,omitempty"`
	Nic           []struct {
		Isolationuri string   `json:"isolationuri,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Id           string   `json:"id,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Type         string   `json:"type,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
	} `json:"nic,omitempty"`
	Guestmacaddress     string `json:"guestmacaddress,omitempty"`
	Hostid              string `json:"hostid,omitempty"`
	Publicnetworkid     string `json:"publicnetworkid,omitempty"`
	Hostname            string `json:"hostname,omitempty"`
	Networkdomain       string `json:"networkdomain,omitempty"`
	Linklocalnetworkid  string `json:"linklocalnetworkid,omitempty"`
	Linklocalnetmask    string `json:"linklocalnetmask,omitempty"`
	Guestnetworkid      string `json:"guestnetworkid,omitempty"`
	Templateid          string `json:"templateid,omitempty"`
	Ip6dns1             string `json:"ip6dns1,omitempty"`
	Name                string `json:"name,omitempty"`
	Requiresupgrade     bool   `json:"requiresupgrade,omitempty"`
	Created             string `json:"created,omitempty"`
	Redundantstate      string `json:"redundantstate,omitempty"`
	Version             string `json:"version,omitempty"`
	Linklocalmacaddress string `json:"linklocalmacaddress,omitempty"`
	Project             string `json:"project,omitempty"`
	Publicmacaddress    string `json:"publicmacaddress,omitempty"`
	Role                string `json:"role,omitempty"`
	Guestnetmask        string `json:"guestnetmask,omitempty"`
	Projectid           string `json:"projectid,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	State               string `json:"state,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
	Gateway             string `json:"gateway,omitempty"`
	Vpcid               string `json:"vpcid,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Account             string `json:"account,omitempty"`
	Dns1                string `json:"dns1,omitempty"`
	Guestipaddress      string `json:"guestipaddress,omitempty"`
	Isredundantrouter   bool   `json:"isredundantrouter,omitempty"`
	Id                  string `json:"id,omitempty"`
	Domainid            string `json:"domainid,omitempty"`
	Podid               string `json:"podid,omitempty"`
	Publicip            string `json:"publicip,omitempty"`
	Scriptsversion      string `json:"scriptsversion,omitempty"`
	Zoneid              string `json:"zoneid,omitempty"`
	Ip6dns2             string `json:"ip6dns2,omitempty"`
	Dns2                string `json:"dns2,omitempty"`
	Linklocalip         string `json:"linklocalip,omitempty"`
}

type ChangeServiceForRouterParams struct {
	p map[string]interface{}
}

func (p *ChangeServiceForRouterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	return u
}

func (p *ChangeServiceForRouterParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ChangeServiceForRouterParams) SetServiceofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingid"] = v
	return
}

// You should always use this function to get a new ChangeServiceForRouterParams instance,
// as then you are sure you have configured all required params
func (s *RouterService) NewChangeServiceForRouterParams(id string, serviceofferingid string) *ChangeServiceForRouterParams {
	p := &ChangeServiceForRouterParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["serviceofferingid"] = serviceofferingid
	return p
}

// Upgrades domain router to a new service offering
func (s *RouterService) ChangeServiceForRouter(p *ChangeServiceForRouterParams) (*ChangeServiceForRouterResponse, error) {
	resp, err := s.cs.newRequest("changeServiceForRouter", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ChangeServiceForRouterResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ChangeServiceForRouterResponse struct {
	Hostname   string `json:"hostname,omitempty"`
	Templateid string `json:"templateid,omitempty"`
	Nic        []struct {
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Id           string   `json:"id,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Type         string   `json:"type,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
	} `json:"nic,omitempty"`
	Isredundantrouter   bool   `json:"isredundantrouter,omitempty"`
	State               string `json:"state,omitempty"`
	Guestipaddress      string `json:"guestipaddress,omitempty"`
	Projectid           string `json:"projectid,omitempty"`
	Account             string `json:"account,omitempty"`
	Guestmacaddress     string `json:"guestmacaddress,omitempty"`
	Publicmacaddress    string `json:"publicmacaddress,omitempty"`
	Domainid            string `json:"domainid,omitempty"`
	Guestnetworkid      string `json:"guestnetworkid,omitempty"`
	Id                  string `json:"id,omitempty"`
	Guestnetmask        string `json:"guestnetmask,omitempty"`
	Requiresupgrade     bool   `json:"requiresupgrade,omitempty"`
	Publicnetmask       string `json:"publicnetmask,omitempty"`
	Dns2                string `json:"dns2,omitempty"`
	Publicip            string `json:"publicip,omitempty"`
	Version             string `json:"version,omitempty"`
	Name                string `json:"name,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
	Linklocalnetmask    string `json:"linklocalnetmask,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Hostid              string `json:"hostid,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Publicnetworkid     string `json:"publicnetworkid,omitempty"`
	Dns1                string `json:"dns1,omitempty"`
	Podid               string `json:"podid,omitempty"`
	Ip6dns2             string `json:"ip6dns2,omitempty"`
	Linklocalmacaddress string `json:"linklocalmacaddress,omitempty"`
	Linklocalnetworkid  string `json:"linklocalnetworkid,omitempty"`
	Gateway             string `json:"gateway,omitempty"`
	Zoneid              string `json:"zoneid,omitempty"`
	Networkdomain       string `json:"networkdomain,omitempty"`
	Linklocalip         string `json:"linklocalip,omitempty"`
	Redundantstate      string `json:"redundantstate,omitempty"`
	Role                string `json:"role,omitempty"`
	Ip6dns1             string `json:"ip6dns1,omitempty"`
	Scriptsversion      string `json:"scriptsversion,omitempty"`
	Project             string `json:"project,omitempty"`
	Created             string `json:"created,omitempty"`
	Vpcid               string `json:"vpcid,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
}

type ListRoutersParams struct {
	p map[string]interface{}
}

func (p *ListRoutersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["forvpc"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forvpc", vv)
	}
	if v, found := p.p["hostid"]; found {
		u.Set("hostid", v.(string))
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
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
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
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := p.p["version"]; found {
		u.Set("version", v.(string))
	}
	if v, found := p.p["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListRoutersParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListRoutersParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
	return
}

func (p *ListRoutersParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListRoutersParams) SetForvpc(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forvpc"] = v
	return
}

func (p *ListRoutersParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
	return
}

func (p *ListRoutersParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListRoutersParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
	return
}

func (p *ListRoutersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListRoutersParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
	return
}

func (p *ListRoutersParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *ListRoutersParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
	return
}

func (p *ListRoutersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListRoutersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListRoutersParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
	return
}

func (p *ListRoutersParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *ListRoutersParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
	return
}

func (p *ListRoutersParams) SetVersion(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["version"] = v
	return
}

func (p *ListRoutersParams) SetVpcid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcid"] = v
	return
}

func (p *ListRoutersParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new ListRoutersParams instance,
// as then you are sure you have configured all required params
func (s *RouterService) NewListRoutersParams() *ListRoutersParams {
	p := &ListRoutersParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *RouterService) GetRouterID(name string) (string, error) {
	p := &ListRoutersParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	l, err := s.ListRouters(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, name, l)
	}
	return l.Routers[0].Id, nil
}

// List routers.
func (s *RouterService) ListRouters(p *ListRoutersParams) (*ListRoutersResponse, error) {
	resp, err := s.cs.newRequest("listRouters", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListRoutersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListRoutersResponse struct {
	Count   int       `json:"count"`
	Routers []*Router `json:"router"`
}

type Router struct {
	Vpcid               string `json:"vpcid,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Linklocalnetmask    string `json:"linklocalnetmask,omitempty"`
	Publicnetworkid     string `json:"publicnetworkid,omitempty"`
	Domainid            string `json:"domainid,omitempty"`
	Dns2                string `json:"dns2,omitempty"`
	Dns1                string `json:"dns1,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
	Account             string `json:"account,omitempty"`
	Linklocalip         string `json:"linklocalip,omitempty"`
	Linklocalmacaddress string `json:"linklocalmacaddress,omitempty"`
	Ip6dns1             string `json:"ip6dns1,omitempty"`
	Zoneid              string `json:"zoneid,omitempty"`
	Redundantstate      string `json:"redundantstate,omitempty"`
	Publicmacaddress    string `json:"publicmacaddress,omitempty"`
	Role                string `json:"role,omitempty"`
	Publicnetmask       string `json:"publicnetmask,omitempty"`
	Version             string `json:"version,omitempty"`
	Project             string `json:"project,omitempty"`
	Nic                 []struct {
		Type         string   `json:"type,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Id           string   `json:"id,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
	} `json:"nic,omitempty"`
	Projectid           string `json:"projectid,omitempty"`
	Linklocalnetworkid  string `json:"linklocalnetworkid,omitempty"`
	Hostid              string `json:"hostid,omitempty"`
	State               string `json:"state,omitempty"`
	Guestmacaddress     string `json:"guestmacaddress,omitempty"`
	Scriptsversion      string `json:"scriptsversion,omitempty"`
	Ip6dns2             string `json:"ip6dns2,omitempty"`
	Guestnetmask        string `json:"guestnetmask,omitempty"`
	Created             string `json:"created,omitempty"`
	Isredundantrouter   bool   `json:"isredundantrouter,omitempty"`
	Publicip            string `json:"publicip,omitempty"`
	Name                string `json:"name,omitempty"`
	Templateid          string `json:"templateid,omitempty"`
	Requiresupgrade     bool   `json:"requiresupgrade,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Guestipaddress      string `json:"guestipaddress,omitempty"`
	Hostname            string `json:"hostname,omitempty"`
	Id                  string `json:"id,omitempty"`
	Guestnetworkid      string `json:"guestnetworkid,omitempty"`
	Podid               string `json:"podid,omitempty"`
	Gateway             string `json:"gateway,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
	Networkdomain       string `json:"networkdomain,omitempty"`
}

type ListVirtualRouterElementsParams struct {
	p map[string]interface{}
}

func (p *ListVirtualRouterElementsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["enabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("enabled", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["nspid"]; found {
		u.Set("nspid", v.(string))
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

func (p *ListVirtualRouterElementsParams) SetEnabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["enabled"] = v
	return
}

func (p *ListVirtualRouterElementsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListVirtualRouterElementsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListVirtualRouterElementsParams) SetNspid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nspid"] = v
	return
}

func (p *ListVirtualRouterElementsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListVirtualRouterElementsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

// You should always use this function to get a new ListVirtualRouterElementsParams instance,
// as then you are sure you have configured all required params
func (s *RouterService) NewListVirtualRouterElementsParams() *ListVirtualRouterElementsParams {
	p := &ListVirtualRouterElementsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *RouterService) GetVirtualRouterElementID(keyword string) (string, error) {
	p := &ListVirtualRouterElementsParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword

	l, err := s.ListVirtualRouterElements(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, keyword, l)
	}
	return l.VirtualRouterElements[0].Id, nil
}

// Lists all available virtual router elements.
func (s *RouterService) ListVirtualRouterElements(p *ListVirtualRouterElementsParams) (*ListVirtualRouterElementsResponse, error) {
	resp, err := s.cs.newRequest("listVirtualRouterElements", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVirtualRouterElementsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListVirtualRouterElementsResponse struct {
	Count                 int                     `json:"count"`
	VirtualRouterElements []*VirtualRouterElement `json:"virtualrouterelement"`
}

type VirtualRouterElement struct {
	Domain    string `json:"domain,omitempty"`
	Projectid string `json:"projectid,omitempty"`
	Project   string `json:"project,omitempty"`
	Domainid  string `json:"domainid,omitempty"`
	Account   string `json:"account,omitempty"`
	Enabled   bool   `json:"enabled,omitempty"`
	Id        string `json:"id,omitempty"`
	Nspid     string `json:"nspid,omitempty"`
}

type ConfigureVirtualRouterElementParams struct {
	p map[string]interface{}
}

func (p *ConfigureVirtualRouterElementParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["enabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("enabled", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *ConfigureVirtualRouterElementParams) SetEnabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["enabled"] = v
	return
}

func (p *ConfigureVirtualRouterElementParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new ConfigureVirtualRouterElementParams instance,
// as then you are sure you have configured all required params
func (s *RouterService) NewConfigureVirtualRouterElementParams(enabled bool, id string) *ConfigureVirtualRouterElementParams {
	p := &ConfigureVirtualRouterElementParams{}
	p.p = make(map[string]interface{})
	p.p["enabled"] = enabled
	p.p["id"] = id
	return p
}

// Configures a virtual router element.
func (s *RouterService) ConfigureVirtualRouterElement(p *ConfigureVirtualRouterElementParams) (*ConfigureVirtualRouterElementResponse, error) {
	resp, err := s.cs.newRequest("configureVirtualRouterElement", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ConfigureVirtualRouterElementResponse
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

		var r ConfigureVirtualRouterElementResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type ConfigureVirtualRouterElementResponse struct {
	JobID     string `json:"jobid,omitempty"`
	Project   string `json:"project,omitempty"`
	Domainid  string `json:"domainid,omitempty"`
	Nspid     string `json:"nspid,omitempty"`
	Enabled   bool   `json:"enabled,omitempty"`
	Id        string `json:"id,omitempty"`
	Projectid string `json:"projectid,omitempty"`
	Account   string `json:"account,omitempty"`
	Domain    string `json:"domain,omitempty"`
}

type CreateVirtualRouterElementParams struct {
	p map[string]interface{}
}

func (p *CreateVirtualRouterElementParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["nspid"]; found {
		u.Set("nspid", v.(string))
	}
	if v, found := p.p["providertype"]; found {
		u.Set("providertype", v.(string))
	}
	return u
}

func (p *CreateVirtualRouterElementParams) SetNspid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nspid"] = v
	return
}

func (p *CreateVirtualRouterElementParams) SetProvidertype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["providertype"] = v
	return
}

// You should always use this function to get a new CreateVirtualRouterElementParams instance,
// as then you are sure you have configured all required params
func (s *RouterService) NewCreateVirtualRouterElementParams(nspid string) *CreateVirtualRouterElementParams {
	p := &CreateVirtualRouterElementParams{}
	p.p = make(map[string]interface{})
	p.p["nspid"] = nspid
	return p
}

// Create a virtual router element.
func (s *RouterService) CreateVirtualRouterElement(p *CreateVirtualRouterElementParams) (*CreateVirtualRouterElementResponse, error) {
	resp, err := s.cs.newRequest("createVirtualRouterElement", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateVirtualRouterElementResponse
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

		var r CreateVirtualRouterElementResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type CreateVirtualRouterElementResponse struct {
	JobID     string `json:"jobid,omitempty"`
	Enabled   bool   `json:"enabled,omitempty"`
	Domain    string `json:"domain,omitempty"`
	Account   string `json:"account,omitempty"`
	Domainid  string `json:"domainid,omitempty"`
	Id        string `json:"id,omitempty"`
	Project   string `json:"project,omitempty"`
	Projectid string `json:"projectid,omitempty"`
	Nspid     string `json:"nspid,omitempty"`
}
