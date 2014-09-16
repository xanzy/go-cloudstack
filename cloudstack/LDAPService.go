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

type LdapCreateAccountParams struct {
	p map[string]interface{}
}

func (p *LdapCreateAccountParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["accountdetails"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("accountdetails[%d].key", i), k)
			u.Set(fmt.Sprintf("accountdetails[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["accountid"]; found {
		u.Set("accountid", v.(string))
	}
	if v, found := p.p["accounttype"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("accounttype", vv)
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["networkdomain"]; found {
		u.Set("networkdomain", v.(string))
	}
	if v, found := p.p["timezone"]; found {
		u.Set("timezone", v.(string))
	}
	if v, found := p.p["userid"]; found {
		u.Set("userid", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (p *LdapCreateAccountParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *LdapCreateAccountParams) SetAccountdetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accountdetails"] = v
	return
}

func (p *LdapCreateAccountParams) SetAccountid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accountid"] = v
	return
}

func (p *LdapCreateAccountParams) SetAccounttype(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accounttype"] = v
	return
}

func (p *LdapCreateAccountParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *LdapCreateAccountParams) SetNetworkdomain(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkdomain"] = v
	return
}

func (p *LdapCreateAccountParams) SetTimezone(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["timezone"] = v
	return
}

func (p *LdapCreateAccountParams) SetUserid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userid"] = v
	return
}

func (p *LdapCreateAccountParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
	return
}

// You should always use this function to get a new LdapCreateAccountParams instance,
// as then you are sure you have configured all required params
func (s *LDAPService) NewLdapCreateAccountParams(accounttype int, username string) *LdapCreateAccountParams {
	p := &LdapCreateAccountParams{}
	p.p = make(map[string]interface{})
	p.p["accounttype"] = accounttype
	p.p["username"] = username
	return p
}

// Creates an account from an LDAP user
func (s *LDAPService) LdapCreateAccount(p *LdapCreateAccountParams) (*LdapCreateAccountResponse, error) {
	resp, err := s.cs.newRequest("ldapCreateAccount", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r LdapCreateAccountResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type LdapCreateAccountResponse struct {
	Vmtotal                 int    `json:"vmtotal,omitempty"`
	Primarystoragelimit     string `json:"primarystoragelimit,omitempty"`
	Cputotal                int    `json:"cputotal,omitempty"`
	Primarystorageavailable string `json:"primarystorageavailable,omitempty"`
	Isdefault               bool   `json:"isdefault,omitempty"`
	User                    []struct {
		Id                  string `json:"id,omitempty"`
		Account             string `json:"account,omitempty"`
		Apikey              string `json:"apikey,omitempty"`
		Domain              string `json:"domain,omitempty"`
		Accounttype         int    `json:"accounttype,omitempty"`
		Lastname            string `json:"lastname,omitempty"`
		Isdefault           bool   `json:"isdefault,omitempty"`
		Username            string `json:"username,omitempty"`
		Created             string `json:"created,omitempty"`
		State               string `json:"state,omitempty"`
		Iscallerchilddomain bool   `json:"iscallerchilddomain,omitempty"`
		Firstname           string `json:"firstname,omitempty"`
		Email               string `json:"email,omitempty"`
		Domainid            string `json:"domainid,omitempty"`
		Accountid           string `json:"accountid,omitempty"`
		Timezone            string `json:"timezone,omitempty"`
		Secretkey           string `json:"secretkey,omitempty"`
	} `json:"user,omitempty"`
	Projectlimit              string            `json:"projectlimit,omitempty"`
	Snapshottotal             int               `json:"snapshottotal,omitempty"`
	Networklimit              string            `json:"networklimit,omitempty"`
	Accountdetails            map[string]string `json:"accountdetails,omitempty"`
	Snapshotavailable         string            `json:"snapshotavailable,omitempty"`
	Iscleanuprequired         bool              `json:"iscleanuprequired,omitempty"`
	Volumeavailable           string            `json:"volumeavailable,omitempty"`
	Projecttotal              int               `json:"projecttotal,omitempty"`
	Networktotal              int               `json:"networktotal,omitempty"`
	Vpctotal                  int               `json:"vpctotal,omitempty"`
	Cpuavailable              string            `json:"cpuavailable,omitempty"`
	Accounttype               int               `json:"accounttype,omitempty"`
	Iplimit                   string            `json:"iplimit,omitempty"`
	Networkavailable          string            `json:"networkavailable,omitempty"`
	Memorytotal               int               `json:"memorytotal,omitempty"`
	Sentbytes                 int               `json:"sentbytes,omitempty"`
	Secondarystoragelimit     string            `json:"secondarystoragelimit,omitempty"`
	Primarystoragetotal       int               `json:"primarystoragetotal,omitempty"`
	Templateavailable         string            `json:"templateavailable,omitempty"`
	Cpulimit                  string            `json:"cpulimit,omitempty"`
	Vmrunning                 int               `json:"vmrunning,omitempty"`
	Ipavailable               string            `json:"ipavailable,omitempty"`
	Receivedbytes             int               `json:"receivedbytes,omitempty"`
	Name                      string            `json:"name,omitempty"`
	Vpclimit                  string            `json:"vpclimit,omitempty"`
	Templatelimit             string            `json:"templatelimit,omitempty"`
	Vmlimit                   string            `json:"vmlimit,omitempty"`
	Volumelimit               string            `json:"volumelimit,omitempty"`
	Secondarystoragetotal     int               `json:"secondarystoragetotal,omitempty"`
	Templatetotal             int               `json:"templatetotal,omitempty"`
	Defaultzoneid             string            `json:"defaultzoneid,omitempty"`
	Id                        string            `json:"id,omitempty"`
	Projectavailable          string            `json:"projectavailable,omitempty"`
	Iptotal                   int               `json:"iptotal,omitempty"`
	Vmavailable               string            `json:"vmavailable,omitempty"`
	Vpcavailable              string            `json:"vpcavailable,omitempty"`
	Vmstopped                 int               `json:"vmstopped,omitempty"`
	Snapshotlimit             string            `json:"snapshotlimit,omitempty"`
	Networkdomain             string            `json:"networkdomain,omitempty"`
	Domain                    string            `json:"domain,omitempty"`
	Memoryavailable           string            `json:"memoryavailable,omitempty"`
	Domainid                  string            `json:"domainid,omitempty"`
	Memorylimit               string            `json:"memorylimit,omitempty"`
	Secondarystorageavailable string            `json:"secondarystorageavailable,omitempty"`
	State                     string            `json:"state,omitempty"`
	Volumetotal               int               `json:"volumetotal,omitempty"`
}
