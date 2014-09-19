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

type StartRouterResponse struct {
	JobID               string `json:"jobid,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
	Templateid          string `json:"templateid,omitempty"`
	Vpcid               string `json:"vpcid,omitempty"`
	Created             string `json:"created,omitempty"`
	Publicnetmask       string `json:"publicnetmask,omitempty"`
	Dns2                string `json:"dns2,omitempty"`
	Project             string `json:"project,omitempty"`
	Linklocalnetworkid  string `json:"linklocalnetworkid,omitempty"`
	Domainid            string `json:"domainid,omitempty"`
	Podid               string `json:"podid,omitempty"`
	Guestmacaddress     string `json:"guestmacaddress,omitempty"`
	Linklocalnetmask    string `json:"linklocalnetmask,omitempty"`
	State               string `json:"state,omitempty"`
	Linklocalmacaddress string `json:"linklocalmacaddress,omitempty"`
	Isredundantrouter   bool   `json:"isredundantrouter,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Nic                 []struct {
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Type         string   `json:"type,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Id           string   `json:"id,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
	} `json:"nic,omitempty"`
	Guestipaddress    string `json:"guestipaddress,omitempty"`
	Hostid            string `json:"hostid,omitempty"`
	Dns1              string `json:"dns1,omitempty"`
	Guestnetmask      string `json:"guestnetmask,omitempty"`
	Linklocalip       string `json:"linklocalip,omitempty"`
	Zoneid            string `json:"zoneid,omitempty"`
	Publicnetworkid   string `json:"publicnetworkid,omitempty"`
	Name              string `json:"name,omitempty"`
	Publicip          string `json:"publicip,omitempty"`
	Guestnetworkid    string `json:"guestnetworkid,omitempty"`
	Domain            string `json:"domain,omitempty"`
	Ip6dns2           string `json:"ip6dns2,omitempty"`
	Requiresupgrade   bool   `json:"requiresupgrade,omitempty"`
	Role              string `json:"role,omitempty"`
	Projectid         string `json:"projectid,omitempty"`
	Hostname          string `json:"hostname,omitempty"`
	Scriptsversion    string `json:"scriptsversion,omitempty"`
	Redundantstate    string `json:"redundantstate,omitempty"`
	Serviceofferingid string `json:"serviceofferingid,omitempty"`
	Networkdomain     string `json:"networkdomain,omitempty"`
	Id                string `json:"id,omitempty"`
	Ip6dns1           string `json:"ip6dns1,omitempty"`
	Version           string `json:"version,omitempty"`
	Gateway           string `json:"gateway,omitempty"`
	Publicmacaddress  string `json:"publicmacaddress,omitempty"`
	Account           string `json:"account,omitempty"`
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

type RebootRouterResponse struct {
	JobID              string `json:"jobid,omitempty"`
	Networkdomain      string `json:"networkdomain,omitempty"`
	Linklocalnetworkid string `json:"linklocalnetworkid,omitempty"`
	Podid              string `json:"podid,omitempty"`
	Guestnetmask       string `json:"guestnetmask,omitempty"`
	Name               string `json:"name,omitempty"`
	Ip6dns2            string `json:"ip6dns2,omitempty"`
	Nic                []struct {
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Id           string   `json:"id,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Type         string   `json:"type,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
	} `json:"nic,omitempty"`
	Hostname            string `json:"hostname,omitempty"`
	Requiresupgrade     bool   `json:"requiresupgrade,omitempty"`
	Id                  string `json:"id,omitempty"`
	Hostid              string `json:"hostid,omitempty"`
	Linklocalnetmask    string `json:"linklocalnetmask,omitempty"`
	Templateid          string `json:"templateid,omitempty"`
	Gateway             string `json:"gateway,omitempty"`
	Publicip            string `json:"publicip,omitempty"`
	Guestipaddress      string `json:"guestipaddress,omitempty"`
	Publicnetmask       string `json:"publicnetmask,omitempty"`
	Vpcid               string `json:"vpcid,omitempty"`
	Linklocalip         string `json:"linklocalip,omitempty"`
	Dns2                string `json:"dns2,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
	Domainid            string `json:"domainid,omitempty"`
	Project             string `json:"project,omitempty"`
	Guestmacaddress     string `json:"guestmacaddress,omitempty"`
	Scriptsversion      string `json:"scriptsversion,omitempty"`
	Account             string `json:"account,omitempty"`
	Created             string `json:"created,omitempty"`
	Projectid           string `json:"projectid,omitempty"`
	Isredundantrouter   bool   `json:"isredundantrouter,omitempty"`
	Ip6dns1             string `json:"ip6dns1,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
	Zoneid              string `json:"zoneid,omitempty"`
	Publicnetworkid     string `json:"publicnetworkid,omitempty"`
	Dns1                string `json:"dns1,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Linklocalmacaddress string `json:"linklocalmacaddress,omitempty"`
	Publicmacaddress    string `json:"publicmacaddress,omitempty"`
	Role                string `json:"role,omitempty"`
	Version             string `json:"version,omitempty"`
	State               string `json:"state,omitempty"`
	Guestnetworkid      string `json:"guestnetworkid,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Redundantstate      string `json:"redundantstate,omitempty"`
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

type StopRouterResponse struct {
	JobID              string `json:"jobid,omitempty"`
	Publicmacaddress   string `json:"publicmacaddress,omitempty"`
	Redundantstate     string `json:"redundantstate,omitempty"`
	State              string `json:"state,omitempty"`
	Zonename           string `json:"zonename,omitempty"`
	Publicip           string `json:"publicip,omitempty"`
	Linklocalnetworkid string `json:"linklocalnetworkid,omitempty"`
	Zoneid             string `json:"zoneid,omitempty"`
	Account            string `json:"account,omitempty"`
	Created            string `json:"created,omitempty"`
	Publicnetworkid    string `json:"publicnetworkid,omitempty"`
	Vpcid              string `json:"vpcid,omitempty"`
	Isredundantrouter  bool   `json:"isredundantrouter,omitempty"`
	Dns2               string `json:"dns2,omitempty"`
	Guestmacaddress    string `json:"guestmacaddress,omitempty"`
	Dns1               string `json:"dns1,omitempty"`
	Gateway            string `json:"gateway,omitempty"`
	Projectid          string `json:"projectid,omitempty"`
	Id                 string `json:"id,omitempty"`
	Role               string `json:"role,omitempty"`
	Guestipaddress     string `json:"guestipaddress,omitempty"`
	Guestnetworkid     string `json:"guestnetworkid,omitempty"`
	Requiresupgrade    bool   `json:"requiresupgrade,omitempty"`
	Nic                []struct {
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Id           string   `json:"id,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Type         string   `json:"type,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
	} `json:"nic,omitempty"`
	Publicnetmask       string `json:"publicnetmask,omitempty"`
	Scriptsversion      string `json:"scriptsversion,omitempty"`
	Domainid            string `json:"domainid,omitempty"`
	Linklocalmacaddress string `json:"linklocalmacaddress,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
	Podid               string `json:"podid,omitempty"`
	Networkdomain       string `json:"networkdomain,omitempty"`
	Ip6dns2             string `json:"ip6dns2,omitempty"`
	Linklocalnetmask    string `json:"linklocalnetmask,omitempty"`
	Name                string `json:"name,omitempty"`
	Templateid          string `json:"templateid,omitempty"`
	Guestnetmask        string `json:"guestnetmask,omitempty"`
	Ip6dns1             string `json:"ip6dns1,omitempty"`
	Linklocalip         string `json:"linklocalip,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Hostid              string `json:"hostid,omitempty"`
	Project             string `json:"project,omitempty"`
	Version             string `json:"version,omitempty"`
	Hostname            string `json:"hostname,omitempty"`
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

type DestroyRouterResponse struct {
	JobID               string `json:"jobid,omitempty"`
	Hostid              string `json:"hostid,omitempty"`
	Podid               string `json:"podid,omitempty"`
	Publicnetmask       string `json:"publicnetmask,omitempty"`
	State               string `json:"state,omitempty"`
	Hostname            string `json:"hostname,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
	Gateway             string `json:"gateway,omitempty"`
	Linklocalip         string `json:"linklocalip,omitempty"`
	Version             string `json:"version,omitempty"`
	Role                string `json:"role,omitempty"`
	Id                  string `json:"id,omitempty"`
	Linklocalmacaddress string `json:"linklocalmacaddress,omitempty"`
	Publicip            string `json:"publicip,omitempty"`
	Networkdomain       string `json:"networkdomain,omitempty"`
	Project             string `json:"project,omitempty"`
	Guestmacaddress     string `json:"guestmacaddress,omitempty"`
	Linklocalnetmask    string `json:"linklocalnetmask,omitempty"`
	Name                string `json:"name,omitempty"`
	Requiresupgrade     bool   `json:"requiresupgrade,omitempty"`
	Projectid           string `json:"projectid,omitempty"`
	Guestipaddress      string `json:"guestipaddress,omitempty"`
	Guestnetworkid      string `json:"guestnetworkid,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Nic                 []struct {
		Isdefault    bool     `json:"isdefault,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Type         string   `json:"type,omitempty"`
		Id           string   `json:"id,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
	} `json:"nic,omitempty"`
	Redundantstate     string `json:"redundantstate,omitempty"`
	Linklocalnetworkid string `json:"linklocalnetworkid,omitempty"`
	Domain             string `json:"domain,omitempty"`
	Ip6dns1            string `json:"ip6dns1,omitempty"`
	Guestnetmask       string `json:"guestnetmask,omitempty"`
	Templateid         string `json:"templateid,omitempty"`
	Dns2               string `json:"dns2,omitempty"`
	Domainid           string `json:"domainid,omitempty"`
	Created            string `json:"created,omitempty"`
	Isredundantrouter  bool   `json:"isredundantrouter,omitempty"`
	Ip6dns2            string `json:"ip6dns2,omitempty"`
	Account            string `json:"account,omitempty"`
	Dns1               string `json:"dns1,omitempty"`
	Publicnetworkid    string `json:"publicnetworkid,omitempty"`
	Zoneid             string `json:"zoneid,omitempty"`
	Vpcid              string `json:"vpcid,omitempty"`
	Publicmacaddress   string `json:"publicmacaddress,omitempty"`
	Scriptsversion     string `json:"scriptsversion,omitempty"`
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
	Publicip            string `json:"publicip,omitempty"`
	State               string `json:"state,omitempty"`
	Version             string `json:"version,omitempty"`
	Role                string `json:"role,omitempty"`
	Account             string `json:"account,omitempty"`
	Name                string `json:"name,omitempty"`
	Linklocalip         string `json:"linklocalip,omitempty"`
	Ip6dns2             string `json:"ip6dns2,omitempty"`
	Guestnetmask        string `json:"guestnetmask,omitempty"`
	Hostid              string `json:"hostid,omitempty"`
	Guestipaddress      string `json:"guestipaddress,omitempty"`
	Dns2                string `json:"dns2,omitempty"`
	Publicnetworkid     string `json:"publicnetworkid,omitempty"`
	Gateway             string `json:"gateway,omitempty"`
	Requiresupgrade     bool   `json:"requiresupgrade,omitempty"`
	Isredundantrouter   bool   `json:"isredundantrouter,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Created             string `json:"created,omitempty"`
	Networkdomain       string `json:"networkdomain,omitempty"`
	Podid               string `json:"podid,omitempty"`
	Publicnetmask       string `json:"publicnetmask,omitempty"`
	Ip6dns1             string `json:"ip6dns1,omitempty"`
	Zoneid              string `json:"zoneid,omitempty"`
	Guestmacaddress     string `json:"guestmacaddress,omitempty"`
	Redundantstate      string `json:"redundantstate,omitempty"`
	Templateid          string `json:"templateid,omitempty"`
	Id                  string `json:"id,omitempty"`
	Hostname            string `json:"hostname,omitempty"`
	Linklocalmacaddress string `json:"linklocalmacaddress,omitempty"`
	Dns1                string `json:"dns1,omitempty"`
	Linklocalnetmask    string `json:"linklocalnetmask,omitempty"`
	Vpcid               string `json:"vpcid,omitempty"`
	Publicmacaddress    string `json:"publicmacaddress,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
	Scriptsversion      string `json:"scriptsversion,omitempty"`
	Project             string `json:"project,omitempty"`
	Projectid           string `json:"projectid,omitempty"`
	Linklocalnetworkid  string `json:"linklocalnetworkid,omitempty"`
	Guestnetworkid      string `json:"guestnetworkid,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Nic                 []struct {
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Id           string   `json:"id,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Type         string   `json:"type,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
	} `json:"nic,omitempty"`
	Domainid string `json:"domainid,omitempty"`
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

	if l.Count == 0 {
		return "", fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Routers[0].Id, nil
	}

	if l.Count > 1 {
		for _, v := range l.Routers {
			if v.Name == name {
				return v.Id, nil
			}
		}
	}
	return "", fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *RouterService) GetRouterByName(name string) (*Router, int, error) {
	id, err := s.GetRouterID(name)
	if err != nil {
		return nil, -1, err
	}

	r, count, err := s.GetRouterByID(id)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *RouterService) GetRouterByID(id string) (*Router, int, error) {
	p := &ListRoutersParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	l, err := s.ListRouters(p)
	if err != nil {
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.Routers[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Router UUID: %s!", id)
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
	Scriptsversion      string `json:"scriptsversion,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Dns2                string `json:"dns2,omitempty"`
	Redundantstate      string `json:"redundantstate,omitempty"`
	Hostname            string `json:"hostname,omitempty"`
	Linklocalmacaddress string `json:"linklocalmacaddress,omitempty"`
	State               string `json:"state,omitempty"`
	Podid               string `json:"podid,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Name                string `json:"name,omitempty"`
	Isredundantrouter   bool   `json:"isredundantrouter,omitempty"`
	Publicmacaddress    string `json:"publicmacaddress,omitempty"`
	Account             string `json:"account,omitempty"`
	Guestnetmask        string `json:"guestnetmask,omitempty"`
	Guestmacaddress     string `json:"guestmacaddress,omitempty"`
	Ip6dns1             string `json:"ip6dns1,omitempty"`
	Networkdomain       string `json:"networkdomain,omitempty"`
	Gateway             string `json:"gateway,omitempty"`
	Domainid            string `json:"domainid,omitempty"`
	Dns1                string `json:"dns1,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
	Vpcid               string `json:"vpcid,omitempty"`
	Id                  string `json:"id,omitempty"`
	Projectid           string `json:"projectid,omitempty"`
	Publicip            string `json:"publicip,omitempty"`
	Publicnetmask       string `json:"publicnetmask,omitempty"`
	Created             string `json:"created,omitempty"`
	Linklocalnetmask    string `json:"linklocalnetmask,omitempty"`
	Version             string `json:"version,omitempty"`
	Nic                 []struct {
		Gateway      string   `json:"gateway,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Id           string   `json:"id,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Type         string   `json:"type,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Netmask      string   `json:"netmask,omitempty"`
	} `json:"nic,omitempty"`
	Project            string `json:"project,omitempty"`
	Publicnetworkid    string `json:"publicnetworkid,omitempty"`
	Hostid             string `json:"hostid,omitempty"`
	Linklocalnetworkid string `json:"linklocalnetworkid,omitempty"`
	Linklocalip        string `json:"linklocalip,omitempty"`
	Ip6dns2            string `json:"ip6dns2,omitempty"`
	Templateid         string `json:"templateid,omitempty"`
	Zoneid             string `json:"zoneid,omitempty"`
	Role               string `json:"role,omitempty"`
	Guestipaddress     string `json:"guestipaddress,omitempty"`
	Requiresupgrade    bool   `json:"requiresupgrade,omitempty"`
	Serviceofferingid  string `json:"serviceofferingid,omitempty"`
	Guestnetworkid     string `json:"guestnetworkid,omitempty"`
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
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
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
func (s *RouterService) GetVirtualRouterElementByID(id string) (*VirtualRouterElement, int, error) {
	p := &ListVirtualRouterElementsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	l, err := s.ListVirtualRouterElements(p)
	if err != nil {
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.VirtualRouterElements[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for VirtualRouterElement UUID: %s!", id)
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
	Project   string `json:"project,omitempty"`
	Account   string `json:"account,omitempty"`
	Domain    string `json:"domain,omitempty"`
	Nspid     string `json:"nspid,omitempty"`
	Projectid string `json:"projectid,omitempty"`
	Enabled   bool   `json:"enabled,omitempty"`
	Id        string `json:"id,omitempty"`
	Domainid  string `json:"domainid,omitempty"`
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

type ConfigureVirtualRouterElementResponse struct {
	JobID     string `json:"jobid,omitempty"`
	Nspid     string `json:"nspid,omitempty"`
	Domainid  string `json:"domainid,omitempty"`
	Account   string `json:"account,omitempty"`
	Project   string `json:"project,omitempty"`
	Projectid string `json:"projectid,omitempty"`
	Id        string `json:"id,omitempty"`
	Domain    string `json:"domain,omitempty"`
	Enabled   bool   `json:"enabled,omitempty"`
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

type CreateVirtualRouterElementResponse struct {
	JobID     string `json:"jobid,omitempty"`
	Enabled   bool   `json:"enabled,omitempty"`
	Account   string `json:"account,omitempty"`
	Domainid  string `json:"domainid,omitempty"`
	Projectid string `json:"projectid,omitempty"`
	Nspid     string `json:"nspid,omitempty"`
	Id        string `json:"id,omitempty"`
	Domain    string `json:"domain,omitempty"`
	Project   string `json:"project,omitempty"`
}
