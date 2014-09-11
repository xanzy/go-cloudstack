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
		vv := strconv.FormatInt(v.(int64), 10)
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
	Snapshotlimit             string `json:"snapshotlimit,omitempty"`
	Cpulimit                  string `json:"cpulimit,omitempty"`
	Secondarystorageavailable string `json:"secondarystorageavailable,omitempty"`
	Primarystoragelimit       string `json:"primarystoragelimit,omitempty"`
	Volumelimit               string `json:"volumelimit,omitempty"`
	Snapshotavailable         string `json:"snapshotavailable,omitempty"`
	Projectavailable          string `json:"projectavailable,omitempty"`
	Receivedbytes             int    `json:"receivedbytes,omitempty"`
	Sentbytes                 int    `json:"sentbytes,omitempty"`
	Vmavailable               string `json:"vmavailable,omitempty"`
	Networkavailable          string `json:"networkavailable,omitempty"`
	Vpcavailable              string `json:"vpcavailable,omitempty"`
	Networklimit              string `json:"networklimit,omitempty"`
	Name                      string `json:"name,omitempty"`
	Domainid                  string `json:"domainid,omitempty"`
	Iscleanuprequired         bool   `json:"iscleanuprequired,omitempty"`
	Vmstopped                 int    `json:"vmstopped,omitempty"`
	Templateavailable         string `json:"templateavailable,omitempty"`
	Secondarystoragetotal     int    `json:"secondarystoragetotal,omitempty"`
	Projecttotal              int    `json:"projecttotal,omitempty"`
	Vpctotal                  int    `json:"vpctotal,omitempty"`
	Iplimit                   string `json:"iplimit,omitempty"`
	Vmrunning                 int    `json:"vmrunning,omitempty"`
	Memoryavailable           string `json:"memoryavailable,omitempty"`
	Secondarystoragelimit     string `json:"secondarystoragelimit,omitempty"`
	Id                        string `json:"id,omitempty"`
	Domain                    string `json:"domain,omitempty"`
	Volumetotal               int    `json:"volumetotal,omitempty"`
	Defaultzoneid             string `json:"defaultzoneid,omitempty"`
	Accounttype               int    `json:"accounttype,omitempty"`
	Isdefault                 bool   `json:"isdefault,omitempty"`
	Templatelimit             string `json:"templatelimit,omitempty"`
	Templatetotal             int    `json:"templatetotal,omitempty"`
	Ipavailable               string `json:"ipavailable,omitempty"`
	Networktotal              int    `json:"networktotal,omitempty"`
	Primarystoragetotal       int    `json:"primarystoragetotal,omitempty"`
	Vmlimit                   string `json:"vmlimit,omitempty"`
	Cpuavailable              string `json:"cpuavailable,omitempty"`
	Vpclimit                  string `json:"vpclimit,omitempty"`
	Memorylimit               string `json:"memorylimit,omitempty"`
	Memorytotal               int    `json:"memorytotal,omitempty"`
	User                      []struct {
		Accounttype         int    `json:"accounttype,omitempty"`
		Apikey              string `json:"apikey,omitempty"`
		State               string `json:"state,omitempty"`
		Secretkey           string `json:"secretkey,omitempty"`
		Email               string `json:"email,omitempty"`
		Isdefault           bool   `json:"isdefault,omitempty"`
		Account             string `json:"account,omitempty"`
		Domain              string `json:"domain,omitempty"`
		Username            string `json:"username,omitempty"`
		Firstname           string `json:"firstname,omitempty"`
		Created             string `json:"created,omitempty"`
		Lastname            string `json:"lastname,omitempty"`
		Timezone            string `json:"timezone,omitempty"`
		Accountid           string `json:"accountid,omitempty"`
		Iscallerchilddomain bool   `json:"iscallerchilddomain,omitempty"`
		Domainid            string `json:"domainid,omitempty"`
		Id                  string `json:"id,omitempty"`
	} `json:"user,omitempty"`
	Primarystorageavailable string            `json:"primarystorageavailable,omitempty"`
	Accountdetails          map[string]string `json:"accountdetails,omitempty"`
	Iptotal                 int               `json:"iptotal,omitempty"`
	Projectlimit            string            `json:"projectlimit,omitempty"`
	Volumeavailable         string            `json:"volumeavailable,omitempty"`
	Cputotal                int               `json:"cputotal,omitempty"`
	Networkdomain           string            `json:"networkdomain,omitempty"`
	Snapshottotal           int               `json:"snapshottotal,omitempty"`
	Vmtotal                 int               `json:"vmtotal,omitempty"`
	State                   string            `json:"state,omitempty"`
}
