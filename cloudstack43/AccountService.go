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

type CreateAccountParams struct {
	p map[string]interface{}
}

func (p *CreateAccountParams) toURLValues() url.Values {
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
	if v, found := p.p["email"]; found {
		u.Set("email", v.(string))
	}
	if v, found := p.p["firstname"]; found {
		u.Set("firstname", v.(string))
	}
	if v, found := p.p["lastname"]; found {
		u.Set("lastname", v.(string))
	}
	if v, found := p.p["networkdomain"]; found {
		u.Set("networkdomain", v.(string))
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
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

func (p *CreateAccountParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *CreateAccountParams) SetAccountdetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accountdetails"] = v
	return
}

func (p *CreateAccountParams) SetAccountid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accountid"] = v
	return
}

func (p *CreateAccountParams) SetAccounttype(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accounttype"] = v
	return
}

func (p *CreateAccountParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *CreateAccountParams) SetEmail(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["email"] = v
	return
}

func (p *CreateAccountParams) SetFirstname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["firstname"] = v
	return
}

func (p *CreateAccountParams) SetLastname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lastname"] = v
	return
}

func (p *CreateAccountParams) SetNetworkdomain(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkdomain"] = v
	return
}

func (p *CreateAccountParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
	return
}

func (p *CreateAccountParams) SetTimezone(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["timezone"] = v
	return
}

func (p *CreateAccountParams) SetUserid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userid"] = v
	return
}

func (p *CreateAccountParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
	return
}

// You should always use this function to get a new CreateAccountParams instance,
// as then you are sure you have configured all required params
func (s *AccountService) NewCreateAccountParams(accounttype int, email string, firstname string, lastname string, password string, username string) *CreateAccountParams {
	p := &CreateAccountParams{}
	p.p = make(map[string]interface{})
	p.p["accounttype"] = accounttype
	p.p["email"] = email
	p.p["firstname"] = firstname
	p.p["lastname"] = lastname
	p.p["password"] = password
	p.p["username"] = username
	return p
}

// Creates an account
func (s *AccountService) CreateAccount(p *CreateAccountParams) (*CreateAccountResponse, error) {
	resp, err := s.cs.newRequest("createAccount", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateAccountResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type CreateAccountResponse struct {
	Isdefault             bool   `json:"isdefault,omitempty"`
	Snapshotlimit         string `json:"snapshotlimit,omitempty"`
	Networktotal          int    `json:"networktotal,omitempty"`
	Vmtotal               int    `json:"vmtotal,omitempty"`
	Domain                string `json:"domain,omitempty"`
	Vmavailable           string `json:"vmavailable,omitempty"`
	Iscleanuprequired     bool   `json:"iscleanuprequired,omitempty"`
	Vmrunning             int    `json:"vmrunning,omitempty"`
	Primarystoragelimit   string `json:"primarystoragelimit,omitempty"`
	Networkdomain         string `json:"networkdomain,omitempty"`
	Name                  string `json:"name,omitempty"`
	Templateavailable     string `json:"templateavailable,omitempty"`
	Sentbytes             int    `json:"sentbytes,omitempty"`
	Id                    string `json:"id,omitempty"`
	Secondarystoragetotal int    `json:"secondarystoragetotal,omitempty"`
	Cputotal              int    `json:"cputotal,omitempty"`
	Defaultzoneid         string `json:"defaultzoneid,omitempty"`
	Memoryavailable       string `json:"memoryavailable,omitempty"`
	Volumelimit           string `json:"volumelimit,omitempty"`
	Receivedbytes         int    `json:"receivedbytes,omitempty"`
	Networklimit          string `json:"networklimit,omitempty"`
	Secondarystoragelimit string `json:"secondarystoragelimit,omitempty"`
	Snapshotavailable     string `json:"snapshotavailable,omitempty"`
	User                  []struct {
		Created             string `json:"created,omitempty"`
		State               string `json:"state,omitempty"`
		Id                  string `json:"id,omitempty"`
		Firstname           string `json:"firstname,omitempty"`
		Domainid            string `json:"domainid,omitempty"`
		Timezone            string `json:"timezone,omitempty"`
		Isdefault           bool   `json:"isdefault,omitempty"`
		Iscallerchilddomain bool   `json:"iscallerchilddomain,omitempty"`
		Domain              string `json:"domain,omitempty"`
		Apikey              string `json:"apikey,omitempty"`
		Username            string `json:"username,omitempty"`
		Accounttype         int    `json:"accounttype,omitempty"`
		Accountid           string `json:"accountid,omitempty"`
		Secretkey           string `json:"secretkey,omitempty"`
		Lastname            string `json:"lastname,omitempty"`
		Account             string `json:"account,omitempty"`
		Email               string `json:"email,omitempty"`
	} `json:"user,omitempty"`
	Memorylimit               string            `json:"memorylimit,omitempty"`
	Accounttype               int               `json:"accounttype,omitempty"`
	Accountdetails            map[string]string `json:"accountdetails,omitempty"`
	Domainid                  string            `json:"domainid,omitempty"`
	Primarystorageavailable   string            `json:"primarystorageavailable,omitempty"`
	State                     string            `json:"state,omitempty"`
	Projectavailable          string            `json:"projectavailable,omitempty"`
	Volumeavailable           string            `json:"volumeavailable,omitempty"`
	Ipavailable               string            `json:"ipavailable,omitempty"`
	Vmlimit                   string            `json:"vmlimit,omitempty"`
	Cpuavailable              string            `json:"cpuavailable,omitempty"`
	Iptotal                   int               `json:"iptotal,omitempty"`
	Projectlimit              string            `json:"projectlimit,omitempty"`
	Volumetotal               int               `json:"volumetotal,omitempty"`
	Templatelimit             string            `json:"templatelimit,omitempty"`
	Cpulimit                  string            `json:"cpulimit,omitempty"`
	Primarystoragetotal       int               `json:"primarystoragetotal,omitempty"`
	Vmstopped                 int               `json:"vmstopped,omitempty"`
	Memorytotal               int               `json:"memorytotal,omitempty"`
	Snapshottotal             int               `json:"snapshottotal,omitempty"`
	Networkavailable          string            `json:"networkavailable,omitempty"`
	Iplimit                   string            `json:"iplimit,omitempty"`
	Projecttotal              int               `json:"projecttotal,omitempty"`
	Vpclimit                  string            `json:"vpclimit,omitempty"`
	Vpcavailable              string            `json:"vpcavailable,omitempty"`
	Templatetotal             int               `json:"templatetotal,omitempty"`
	Secondarystorageavailable string            `json:"secondarystorageavailable,omitempty"`
	Vpctotal                  int               `json:"vpctotal,omitempty"`
}

type DeleteAccountParams struct {
	p map[string]interface{}
}

func (p *DeleteAccountParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteAccountParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteAccountParams instance,
// as then you are sure you have configured all required params
func (s *AccountService) NewDeleteAccountParams(id string) *DeleteAccountParams {
	p := &DeleteAccountParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a account, and all users associated with this account
func (s *AccountService) DeleteAccount(p *DeleteAccountParams) (*DeleteAccountResponse, error) {
	resp, err := s.cs.newRequest("deleteAccount", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteAccountResponse
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

		var r DeleteAccountResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type DeleteAccountResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Success     bool   `json:"success,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
}

type UpdateAccountParams struct {
	p map[string]interface{}
}

func (p *UpdateAccountParams) toURLValues() url.Values {
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
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["networkdomain"]; found {
		u.Set("networkdomain", v.(string))
	}
	if v, found := p.p["newname"]; found {
		u.Set("newname", v.(string))
	}
	return u
}

func (p *UpdateAccountParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *UpdateAccountParams) SetAccountdetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accountdetails"] = v
	return
}

func (p *UpdateAccountParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *UpdateAccountParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *UpdateAccountParams) SetNetworkdomain(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkdomain"] = v
	return
}

func (p *UpdateAccountParams) SetNewname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["newname"] = v
	return
}

// You should always use this function to get a new UpdateAccountParams instance,
// as then you are sure you have configured all required params
func (s *AccountService) NewUpdateAccountParams(newname string) *UpdateAccountParams {
	p := &UpdateAccountParams{}
	p.p = make(map[string]interface{})
	p.p["newname"] = newname
	return p
}

// Updates account information for the authenticated user
func (s *AccountService) UpdateAccount(p *UpdateAccountParams) (*UpdateAccountResponse, error) {
	resp, err := s.cs.newRequest("updateAccount", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateAccountResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type UpdateAccountResponse struct {
	Networkdomain             string            `json:"networkdomain,omitempty"`
	Networktotal              int               `json:"networktotal,omitempty"`
	Vpctotal                  int               `json:"vpctotal,omitempty"`
	Secondarystorageavailable string            `json:"secondarystorageavailable,omitempty"`
	Accountdetails            map[string]string `json:"accountdetails,omitempty"`
	Templatelimit             string            `json:"templatelimit,omitempty"`
	Iscleanuprequired         bool              `json:"iscleanuprequired,omitempty"`
	Vmavailable               string            `json:"vmavailable,omitempty"`
	Memoryavailable           string            `json:"memoryavailable,omitempty"`
	Vmstopped                 int               `json:"vmstopped,omitempty"`
	Name                      string            `json:"name,omitempty"`
	Networklimit              string            `json:"networklimit,omitempty"`
	Templatetotal             int               `json:"templatetotal,omitempty"`
	Primarystoragetotal       int               `json:"primarystoragetotal,omitempty"`
	Cputotal                  int               `json:"cputotal,omitempty"`
	Primarystorageavailable   string            `json:"primarystorageavailable,omitempty"`
	Receivedbytes             int               `json:"receivedbytes,omitempty"`
	Vmrunning                 int               `json:"vmrunning,omitempty"`
	Ipavailable               string            `json:"ipavailable,omitempty"`
	Vmlimit                   string            `json:"vmlimit,omitempty"`
	Memorylimit               string            `json:"memorylimit,omitempty"`
	Templateavailable         string            `json:"templateavailable,omitempty"`
	Projectavailable          string            `json:"projectavailable,omitempty"`
	Accounttype               int               `json:"accounttype,omitempty"`
	Projectlimit              string            `json:"projectlimit,omitempty"`
	Domain                    string            `json:"domain,omitempty"`
	Vmtotal                   int               `json:"vmtotal,omitempty"`
	Iplimit                   string            `json:"iplimit,omitempty"`
	Secondarystoragetotal     int               `json:"secondarystoragetotal,omitempty"`
	Cpuavailable              string            `json:"cpuavailable,omitempty"`
	Snapshotavailable         string            `json:"snapshotavailable,omitempty"`
	Projecttotal              int               `json:"projecttotal,omitempty"`
	Id                        string            `json:"id,omitempty"`
	Snapshotlimit             string            `json:"snapshotlimit,omitempty"`
	Volumeavailable           string            `json:"volumeavailable,omitempty"`
	Memorytotal               int               `json:"memorytotal,omitempty"`
	Primarystoragelimit       string            `json:"primarystoragelimit,omitempty"`
	Cpulimit                  string            `json:"cpulimit,omitempty"`
	State                     string            `json:"state,omitempty"`
	Volumelimit               string            `json:"volumelimit,omitempty"`
	User                      []struct {
		Secretkey           string `json:"secretkey,omitempty"`
		Username            string `json:"username,omitempty"`
		State               string `json:"state,omitempty"`
		Accounttype         int    `json:"accounttype,omitempty"`
		Domain              string `json:"domain,omitempty"`
		Timezone            string `json:"timezone,omitempty"`
		Iscallerchilddomain bool   `json:"iscallerchilddomain,omitempty"`
		Isdefault           bool   `json:"isdefault,omitempty"`
		Domainid            string `json:"domainid,omitempty"`
		Email               string `json:"email,omitempty"`
		Id                  string `json:"id,omitempty"`
		Firstname           string `json:"firstname,omitempty"`
		Lastname            string `json:"lastname,omitempty"`
		Created             string `json:"created,omitempty"`
		Accountid           string `json:"accountid,omitempty"`
		Account             string `json:"account,omitempty"`
		Apikey              string `json:"apikey,omitempty"`
	} `json:"user,omitempty"`
	Defaultzoneid         string `json:"defaultzoneid,omitempty"`
	Iptotal               int    `json:"iptotal,omitempty"`
	Secondarystoragelimit string `json:"secondarystoragelimit,omitempty"`
	Sentbytes             int    `json:"sentbytes,omitempty"`
	Snapshottotal         int    `json:"snapshottotal,omitempty"`
	Volumetotal           int    `json:"volumetotal,omitempty"`
	Isdefault             bool   `json:"isdefault,omitempty"`
	Vpcavailable          string `json:"vpcavailable,omitempty"`
	Vpclimit              string `json:"vpclimit,omitempty"`
	Networkavailable      string `json:"networkavailable,omitempty"`
	Domainid              string `json:"domainid,omitempty"`
}

type DisableAccountParams struct {
	p map[string]interface{}
}

func (p *DisableAccountParams) toURLValues() url.Values {
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
	if v, found := p.p["lock"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("lock", vv)
	}
	return u
}

func (p *DisableAccountParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *DisableAccountParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *DisableAccountParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *DisableAccountParams) SetLock(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lock"] = v
	return
}

// You should always use this function to get a new DisableAccountParams instance,
// as then you are sure you have configured all required params
func (s *AccountService) NewDisableAccountParams(lock bool) *DisableAccountParams {
	p := &DisableAccountParams{}
	p.p = make(map[string]interface{})
	p.p["lock"] = lock
	return p
}

// Disables an account
func (s *AccountService) DisableAccount(p *DisableAccountParams) (*DisableAccountResponse, error) {
	resp, err := s.cs.newRequest("disableAccount", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DisableAccountResponse
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

		var r DisableAccountResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type DisableAccountResponse struct {
	JobID string `json:"jobid,omitempty"`
	User  []struct {
		Lastname            string `json:"lastname,omitempty"`
		Email               string `json:"email,omitempty"`
		Accounttype         int    `json:"accounttype,omitempty"`
		Domainid            string `json:"domainid,omitempty"`
		Account             string `json:"account,omitempty"`
		Domain              string `json:"domain,omitempty"`
		Iscallerchilddomain bool   `json:"iscallerchilddomain,omitempty"`
		Accountid           string `json:"accountid,omitempty"`
		Apikey              string `json:"apikey,omitempty"`
		Username            string `json:"username,omitempty"`
		Id                  string `json:"id,omitempty"`
		Isdefault           bool   `json:"isdefault,omitempty"`
		Created             string `json:"created,omitempty"`
		Firstname           string `json:"firstname,omitempty"`
		State               string `json:"state,omitempty"`
		Timezone            string `json:"timezone,omitempty"`
		Secretkey           string `json:"secretkey,omitempty"`
	} `json:"user,omitempty"`
	Vmavailable               string            `json:"vmavailable,omitempty"`
	Domain                    string            `json:"domain,omitempty"`
	Projectavailable          string            `json:"projectavailable,omitempty"`
	Iptotal                   int               `json:"iptotal,omitempty"`
	Memorytotal               int               `json:"memorytotal,omitempty"`
	State                     string            `json:"state,omitempty"`
	Iscleanuprequired         bool              `json:"iscleanuprequired,omitempty"`
	Id                        string            `json:"id,omitempty"`
	Ipavailable               string            `json:"ipavailable,omitempty"`
	Defaultzoneid             string            `json:"defaultzoneid,omitempty"`
	Vpctotal                  int               `json:"vpctotal,omitempty"`
	Accountdetails            map[string]string `json:"accountdetails,omitempty"`
	Vmlimit                   string            `json:"vmlimit,omitempty"`
	Name                      string            `json:"name,omitempty"`
	Iplimit                   string            `json:"iplimit,omitempty"`
	Vpcavailable              string            `json:"vpcavailable,omitempty"`
	Networkavailable          string            `json:"networkavailable,omitempty"`
	Cputotal                  int               `json:"cputotal,omitempty"`
	Secondarystoragelimit     string            `json:"secondarystoragelimit,omitempty"`
	Memorylimit               string            `json:"memorylimit,omitempty"`
	Networkdomain             string            `json:"networkdomain,omitempty"`
	Primarystoragelimit       string            `json:"primarystoragelimit,omitempty"`
	Secondarystorageavailable string            `json:"secondarystorageavailable,omitempty"`
	Networktotal              int               `json:"networktotal,omitempty"`
	Vmtotal                   int               `json:"vmtotal,omitempty"`
	Memoryavailable           string            `json:"memoryavailable,omitempty"`
	Sentbytes                 int               `json:"sentbytes,omitempty"`
	Volumetotal               int               `json:"volumetotal,omitempty"`
	Accounttype               int               `json:"accounttype,omitempty"`
	Volumelimit               string            `json:"volumelimit,omitempty"`
	Vpclimit                  string            `json:"vpclimit,omitempty"`
	Networklimit              string            `json:"networklimit,omitempty"`
	Snapshotlimit             string            `json:"snapshotlimit,omitempty"`
	Domainid                  string            `json:"domainid,omitempty"`
	Templatetotal             int               `json:"templatetotal,omitempty"`
	Snapshottotal             int               `json:"snapshottotal,omitempty"`
	Isdefault                 bool              `json:"isdefault,omitempty"`
	Cpulimit                  string            `json:"cpulimit,omitempty"`
	Cpuavailable              string            `json:"cpuavailable,omitempty"`
	Vmrunning                 int               `json:"vmrunning,omitempty"`
	Snapshotavailable         string            `json:"snapshotavailable,omitempty"`
	Projectlimit              string            `json:"projectlimit,omitempty"`
	Templateavailable         string            `json:"templateavailable,omitempty"`
	Volumeavailable           string            `json:"volumeavailable,omitempty"`
	Receivedbytes             int               `json:"receivedbytes,omitempty"`
	Templatelimit             string            `json:"templatelimit,omitempty"`
	Projecttotal              int               `json:"projecttotal,omitempty"`
	Secondarystoragetotal     int               `json:"secondarystoragetotal,omitempty"`
	Vmstopped                 int               `json:"vmstopped,omitempty"`
	Primarystorageavailable   string            `json:"primarystorageavailable,omitempty"`
	Primarystoragetotal       int               `json:"primarystoragetotal,omitempty"`
}

type EnableAccountParams struct {
	p map[string]interface{}
}

func (p *EnableAccountParams) toURLValues() url.Values {
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
	return u
}

func (p *EnableAccountParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *EnableAccountParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *EnableAccountParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new EnableAccountParams instance,
// as then you are sure you have configured all required params
func (s *AccountService) NewEnableAccountParams() *EnableAccountParams {
	p := &EnableAccountParams{}
	p.p = make(map[string]interface{})
	return p
}

// Enables an account
func (s *AccountService) EnableAccount(p *EnableAccountParams) (*EnableAccountResponse, error) {
	resp, err := s.cs.newRequest("enableAccount", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r EnableAccountResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type EnableAccountResponse struct {
	State                     string `json:"state,omitempty"`
	Name                      string `json:"name,omitempty"`
	Volumetotal               int    `json:"volumetotal,omitempty"`
	Templatelimit             string `json:"templatelimit,omitempty"`
	Templateavailable         string `json:"templateavailable,omitempty"`
	Vpctotal                  int    `json:"vpctotal,omitempty"`
	Vpcavailable              string `json:"vpcavailable,omitempty"`
	Primarystoragelimit       string `json:"primarystoragelimit,omitempty"`
	Domainid                  string `json:"domainid,omitempty"`
	Sentbytes                 int    `json:"sentbytes,omitempty"`
	Networklimit              string `json:"networklimit,omitempty"`
	Memorylimit               string `json:"memorylimit,omitempty"`
	Primarystoragetotal       int    `json:"primarystoragetotal,omitempty"`
	Primarystorageavailable   string `json:"primarystorageavailable,omitempty"`
	Networktotal              int    `json:"networktotal,omitempty"`
	Cpulimit                  string `json:"cpulimit,omitempty"`
	Defaultzoneid             string `json:"defaultzoneid,omitempty"`
	Receivedbytes             int    `json:"receivedbytes,omitempty"`
	Vmstopped                 int    `json:"vmstopped,omitempty"`
	Projecttotal              int    `json:"projecttotal,omitempty"`
	Accounttype               int    `json:"accounttype,omitempty"`
	Networkavailable          string `json:"networkavailable,omitempty"`
	Cputotal                  int    `json:"cputotal,omitempty"`
	Snapshotlimit             string `json:"snapshotlimit,omitempty"`
	Volumeavailable           string `json:"volumeavailable,omitempty"`
	Vpclimit                  string `json:"vpclimit,omitempty"`
	Isdefault                 bool   `json:"isdefault,omitempty"`
	Snapshotavailable         string `json:"snapshotavailable,omitempty"`
	Cpuavailable              string `json:"cpuavailable,omitempty"`
	Iscleanuprequired         bool   `json:"iscleanuprequired,omitempty"`
	Domain                    string `json:"domain,omitempty"`
	Memoryavailable           string `json:"memoryavailable,omitempty"`
	Secondarystorageavailable string `json:"secondarystorageavailable,omitempty"`
	Templatetotal             int    `json:"templatetotal,omitempty"`
	Vmtotal                   int    `json:"vmtotal,omitempty"`
	Memorytotal               int    `json:"memorytotal,omitempty"`
	Id                        string `json:"id,omitempty"`
	Volumelimit               string `json:"volumelimit,omitempty"`
	Ipavailable               string `json:"ipavailable,omitempty"`
	Snapshottotal             int    `json:"snapshottotal,omitempty"`
	Vmlimit                   string `json:"vmlimit,omitempty"`
	Vmavailable               string `json:"vmavailable,omitempty"`
	Iptotal                   int    `json:"iptotal,omitempty"`
	Secondarystoragelimit     string `json:"secondarystoragelimit,omitempty"`
	User                      []struct {
		Isdefault           bool   `json:"isdefault,omitempty"`
		Lastname            string `json:"lastname,omitempty"`
		Account             string `json:"account,omitempty"`
		Domain              string `json:"domain,omitempty"`
		Username            string `json:"username,omitempty"`
		Firstname           string `json:"firstname,omitempty"`
		Secretkey           string `json:"secretkey,omitempty"`
		State               string `json:"state,omitempty"`
		Accounttype         int    `json:"accounttype,omitempty"`
		Domainid            string `json:"domainid,omitempty"`
		Timezone            string `json:"timezone,omitempty"`
		Apikey              string `json:"apikey,omitempty"`
		Accountid           string `json:"accountid,omitempty"`
		Created             string `json:"created,omitempty"`
		Id                  string `json:"id,omitempty"`
		Email               string `json:"email,omitempty"`
		Iscallerchilddomain bool   `json:"iscallerchilddomain,omitempty"`
	} `json:"user,omitempty"`
	Networkdomain         string            `json:"networkdomain,omitempty"`
	Vmrunning             int               `json:"vmrunning,omitempty"`
	Accountdetails        map[string]string `json:"accountdetails,omitempty"`
	Secondarystoragetotal int               `json:"secondarystoragetotal,omitempty"`
	Iplimit               string            `json:"iplimit,omitempty"`
	Projectavailable      string            `json:"projectavailable,omitempty"`
	Projectlimit          string            `json:"projectlimit,omitempty"`
}

type LockAccountParams struct {
	p map[string]interface{}
}

func (p *LockAccountParams) toURLValues() url.Values {
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
	return u
}

func (p *LockAccountParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *LockAccountParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

// You should always use this function to get a new LockAccountParams instance,
// as then you are sure you have configured all required params
func (s *AccountService) NewLockAccountParams(account string, domainid string) *LockAccountParams {
	p := &LockAccountParams{}
	p.p = make(map[string]interface{})
	p.p["account"] = account
	p.p["domainid"] = domainid
	return p
}

// Locks an account
func (s *AccountService) LockAccount(p *LockAccountParams) (*LockAccountResponse, error) {
	resp, err := s.cs.newRequest("lockAccount", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r LockAccountResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type LockAccountResponse struct {
	Volumetotal               int               `json:"volumetotal,omitempty"`
	Networkdomain             string            `json:"networkdomain,omitempty"`
	Primarystoragetotal       int               `json:"primarystoragetotal,omitempty"`
	Ipavailable               string            `json:"ipavailable,omitempty"`
	Secondarystorageavailable string            `json:"secondarystorageavailable,omitempty"`
	Projectavailable          string            `json:"projectavailable,omitempty"`
	Vmrunning                 int               `json:"vmrunning,omitempty"`
	Snapshottotal             int               `json:"snapshottotal,omitempty"`
	Domain                    string            `json:"domain,omitempty"`
	Networklimit              string            `json:"networklimit,omitempty"`
	Primarystorageavailable   string            `json:"primarystorageavailable,omitempty"`
	Name                      string            `json:"name,omitempty"`
	Vmlimit                   string            `json:"vmlimit,omitempty"`
	Secondarystoragetotal     int               `json:"secondarystoragetotal,omitempty"`
	Memorytotal               int               `json:"memorytotal,omitempty"`
	Vmtotal                   int               `json:"vmtotal,omitempty"`
	Snapshotavailable         string            `json:"snapshotavailable,omitempty"`
	Iscleanuprequired         bool              `json:"iscleanuprequired,omitempty"`
	Isdefault                 bool              `json:"isdefault,omitempty"`
	Vpcavailable              string            `json:"vpcavailable,omitempty"`
	Templatelimit             string            `json:"templatelimit,omitempty"`
	Accountdetails            map[string]string `json:"accountdetails,omitempty"`
	Templatetotal             int               `json:"templatetotal,omitempty"`
	Projecttotal              int               `json:"projecttotal,omitempty"`
	Iplimit                   string            `json:"iplimit,omitempty"`
	Templateavailable         string            `json:"templateavailable,omitempty"`
	Networktotal              int               `json:"networktotal,omitempty"`
	Memorylimit               string            `json:"memorylimit,omitempty"`
	Domainid                  string            `json:"domainid,omitempty"`
	Primarystoragelimit       string            `json:"primarystoragelimit,omitempty"`
	Vpctotal                  int               `json:"vpctotal,omitempty"`
	Iptotal                   int               `json:"iptotal,omitempty"`
	Cpulimit                  string            `json:"cpulimit,omitempty"`
	Cpuavailable              string            `json:"cpuavailable,omitempty"`
	Volumeavailable           string            `json:"volumeavailable,omitempty"`
	Defaultzoneid             string            `json:"defaultzoneid,omitempty"`
	Vmstopped                 int               `json:"vmstopped,omitempty"`
	State                     string            `json:"state,omitempty"`
	Accounttype               int               `json:"accounttype,omitempty"`
	Vmavailable               string            `json:"vmavailable,omitempty"`
	Volumelimit               string            `json:"volumelimit,omitempty"`
	Projectlimit              string            `json:"projectlimit,omitempty"`
	Snapshotlimit             string            `json:"snapshotlimit,omitempty"`
	Networkavailable          string            `json:"networkavailable,omitempty"`
	Vpclimit                  string            `json:"vpclimit,omitempty"`
	User                      []struct {
		Domain              string `json:"domain,omitempty"`
		State               string `json:"state,omitempty"`
		Id                  string `json:"id,omitempty"`
		Username            string `json:"username,omitempty"`
		Account             string `json:"account,omitempty"`
		Domainid            string `json:"domainid,omitempty"`
		Accountid           string `json:"accountid,omitempty"`
		Iscallerchilddomain bool   `json:"iscallerchilddomain,omitempty"`
		Accounttype         int    `json:"accounttype,omitempty"`
		Firstname           string `json:"firstname,omitempty"`
		Lastname            string `json:"lastname,omitempty"`
		Timezone            string `json:"timezone,omitempty"`
		Isdefault           bool   `json:"isdefault,omitempty"`
		Apikey              string `json:"apikey,omitempty"`
		Email               string `json:"email,omitempty"`
		Secretkey           string `json:"secretkey,omitempty"`
		Created             string `json:"created,omitempty"`
	} `json:"user,omitempty"`
	Sentbytes             int    `json:"sentbytes,omitempty"`
	Memoryavailable       string `json:"memoryavailable,omitempty"`
	Receivedbytes         int    `json:"receivedbytes,omitempty"`
	Cputotal              int    `json:"cputotal,omitempty"`
	Id                    string `json:"id,omitempty"`
	Secondarystoragelimit string `json:"secondarystoragelimit,omitempty"`
}

type ListAccountsParams struct {
	p map[string]interface{}
}

func (p *ListAccountsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["accounttype"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("accounttype", vv)
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["iscleanuprequired"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("iscleanuprequired", vv)
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
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	return u
}

func (p *ListAccountsParams) SetAccounttype(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accounttype"] = v
	return
}

func (p *ListAccountsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListAccountsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListAccountsParams) SetIscleanuprequired(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["iscleanuprequired"] = v
	return
}

func (p *ListAccountsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
	return
}

func (p *ListAccountsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListAccountsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
	return
}

func (p *ListAccountsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *ListAccountsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListAccountsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListAccountsParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
	return
}

// You should always use this function to get a new ListAccountsParams instance,
// as then you are sure you have configured all required params
func (s *AccountService) NewListAccountsParams() *ListAccountsParams {
	p := &ListAccountsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AccountService) GetAccountID(name string) (string, error) {
	p := &ListAccountsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	l, err := s.ListAccounts(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, name, l)
	}
	return l.Accounts[0].Id, nil
}

// Lists accounts and provides detailed account information for listed accounts
func (s *AccountService) ListAccounts(p *ListAccountsParams) (*ListAccountsResponse, error) {
	resp, err := s.cs.newRequest("listAccounts", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListAccountsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListAccountsResponse struct {
	Count    int        `json:"count"`
	Accounts []*Account `json:"account"`
}

type Account struct {
	Snapshottotal         int    `json:"snapshottotal,omitempty"`
	Sentbytes             int    `json:"sentbytes,omitempty"`
	Name                  string `json:"name,omitempty"`
	Vmtotal               int    `json:"vmtotal,omitempty"`
	Cpuavailable          string `json:"cpuavailable,omitempty"`
	Iscleanuprequired     bool   `json:"iscleanuprequired,omitempty"`
	Vmrunning             int    `json:"vmrunning,omitempty"`
	Networklimit          string `json:"networklimit,omitempty"`
	Projecttotal          int    `json:"projecttotal,omitempty"`
	Memorytotal           int    `json:"memorytotal,omitempty"`
	Secondarystoragetotal int    `json:"secondarystoragetotal,omitempty"`
	Templatetotal         int    `json:"templatetotal,omitempty"`
	Networkavailable      string `json:"networkavailable,omitempty"`
	Templateavailable     string `json:"templateavailable,omitempty"`
	Snapshotlimit         string `json:"snapshotlimit,omitempty"`
	Cpulimit              string `json:"cpulimit,omitempty"`
	User                  []struct {
		Apikey              string `json:"apikey,omitempty"`
		Domainid            string `json:"domainid,omitempty"`
		Lastname            string `json:"lastname,omitempty"`
		Account             string `json:"account,omitempty"`
		State               string `json:"state,omitempty"`
		Isdefault           bool   `json:"isdefault,omitempty"`
		Firstname           string `json:"firstname,omitempty"`
		Accounttype         int    `json:"accounttype,omitempty"`
		Secretkey           string `json:"secretkey,omitempty"`
		Id                  string `json:"id,omitempty"`
		Domain              string `json:"domain,omitempty"`
		Username            string `json:"username,omitempty"`
		Timezone            string `json:"timezone,omitempty"`
		Iscallerchilddomain bool   `json:"iscallerchilddomain,omitempty"`
		Email               string `json:"email,omitempty"`
		Accountid           string `json:"accountid,omitempty"`
		Created             string `json:"created,omitempty"`
	} `json:"user,omitempty"`
	Vmlimit                   string            `json:"vmlimit,omitempty"`
	Volumeavailable           string            `json:"volumeavailable,omitempty"`
	Vpctotal                  int               `json:"vpctotal,omitempty"`
	Cputotal                  int               `json:"cputotal,omitempty"`
	Accountdetails            map[string]string `json:"accountdetails,omitempty"`
	Networkdomain             string            `json:"networkdomain,omitempty"`
	Projectlimit              string            `json:"projectlimit,omitempty"`
	Secondarystorageavailable string            `json:"secondarystorageavailable,omitempty"`
	Memoryavailable           string            `json:"memoryavailable,omitempty"`
	Primarystoragelimit       string            `json:"primarystoragelimit,omitempty"`
	Iplimit                   string            `json:"iplimit,omitempty"`
	State                     string            `json:"state,omitempty"`
	Memorylimit               string            `json:"memorylimit,omitempty"`
	Vpclimit                  string            `json:"vpclimit,omitempty"`
	Id                        string            `json:"id,omitempty"`
	Secondarystoragelimit     string            `json:"secondarystoragelimit,omitempty"`
	Snapshotavailable         string            `json:"snapshotavailable,omitempty"`
	Volumetotal               int               `json:"volumetotal,omitempty"`
	Isdefault                 bool              `json:"isdefault,omitempty"`
	Templatelimit             string            `json:"templatelimit,omitempty"`
	Domain                    string            `json:"domain,omitempty"`
	Receivedbytes             int               `json:"receivedbytes,omitempty"`
	Vmstopped                 int               `json:"vmstopped,omitempty"`
	Networktotal              int               `json:"networktotal,omitempty"`
	Projectavailable          string            `json:"projectavailable,omitempty"`
	Primarystorageavailable   string            `json:"primarystorageavailable,omitempty"`
	Domainid                  string            `json:"domainid,omitempty"`
	Iptotal                   int               `json:"iptotal,omitempty"`
	Vpcavailable              string            `json:"vpcavailable,omitempty"`
	Ipavailable               string            `json:"ipavailable,omitempty"`
	Volumelimit               string            `json:"volumelimit,omitempty"`
	Vmavailable               string            `json:"vmavailable,omitempty"`
	Accounttype               int               `json:"accounttype,omitempty"`
	Defaultzoneid             string            `json:"defaultzoneid,omitempty"`
	Primarystoragetotal       int               `json:"primarystoragetotal,omitempty"`
}

type MarkDefaultZoneForAccountParams struct {
	p map[string]interface{}
}

func (p *MarkDefaultZoneForAccountParams) toURLValues() url.Values {
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
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *MarkDefaultZoneForAccountParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *MarkDefaultZoneForAccountParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *MarkDefaultZoneForAccountParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new MarkDefaultZoneForAccountParams instance,
// as then you are sure you have configured all required params
func (s *AccountService) NewMarkDefaultZoneForAccountParams(account string, domainid string, zoneid string) *MarkDefaultZoneForAccountParams {
	p := &MarkDefaultZoneForAccountParams{}
	p.p = make(map[string]interface{})
	p.p["account"] = account
	p.p["domainid"] = domainid
	p.p["zoneid"] = zoneid
	return p
}

// Marks a default zone for this account
func (s *AccountService) MarkDefaultZoneForAccount(p *MarkDefaultZoneForAccountParams) (*MarkDefaultZoneForAccountResponse, error) {
	resp, err := s.cs.newRequest("markDefaultZoneForAccount", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r MarkDefaultZoneForAccountResponse
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

		var r MarkDefaultZoneForAccountResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type MarkDefaultZoneForAccountResponse struct {
	JobID                     string            `json:"jobid,omitempty"`
	Memorylimit               string            `json:"memorylimit,omitempty"`
	Cpuavailable              string            `json:"cpuavailable,omitempty"`
	Projectlimit              string            `json:"projectlimit,omitempty"`
	Ipavailable               string            `json:"ipavailable,omitempty"`
	Defaultzoneid             string            `json:"defaultzoneid,omitempty"`
	Projecttotal              int               `json:"projecttotal,omitempty"`
	Memorytotal               int               `json:"memorytotal,omitempty"`
	Secondarystoragetotal     int               `json:"secondarystoragetotal,omitempty"`
	Cpulimit                  string            `json:"cpulimit,omitempty"`
	Vmrunning                 int               `json:"vmrunning,omitempty"`
	Vmtotal                   int               `json:"vmtotal,omitempty"`
	Vpclimit                  string            `json:"vpclimit,omitempty"`
	Snapshotlimit             string            `json:"snapshotlimit,omitempty"`
	Primarystoragetotal       int               `json:"primarystoragetotal,omitempty"`
	Secondarystorageavailable string            `json:"secondarystorageavailable,omitempty"`
	Id                        string            `json:"id,omitempty"`
	Sentbytes                 int               `json:"sentbytes,omitempty"`
	Templatelimit             string            `json:"templatelimit,omitempty"`
	Isdefault                 bool              `json:"isdefault,omitempty"`
	Accounttype               int               `json:"accounttype,omitempty"`
	Cputotal                  int               `json:"cputotal,omitempty"`
	Accountdetails            map[string]string `json:"accountdetails,omitempty"`
	Templatetotal             int               `json:"templatetotal,omitempty"`
	Networkdomain             string            `json:"networkdomain,omitempty"`
	Iptotal                   int               `json:"iptotal,omitempty"`
	Primarystorageavailable   string            `json:"primarystorageavailable,omitempty"`
	Iscleanuprequired         bool              `json:"iscleanuprequired,omitempty"`
	User                      []struct {
		Accountid           string `json:"accountid,omitempty"`
		Email               string `json:"email,omitempty"`
		Domain              string `json:"domain,omitempty"`
		Firstname           string `json:"firstname,omitempty"`
		Id                  string `json:"id,omitempty"`
		Lastname            string `json:"lastname,omitempty"`
		State               string `json:"state,omitempty"`
		Iscallerchilddomain bool   `json:"iscallerchilddomain,omitempty"`
		Accounttype         int    `json:"accounttype,omitempty"`
		Isdefault           bool   `json:"isdefault,omitempty"`
		Username            string `json:"username,omitempty"`
		Timezone            string `json:"timezone,omitempty"`
		Domainid            string `json:"domainid,omitempty"`
		Apikey              string `json:"apikey,omitempty"`
		Account             string `json:"account,omitempty"`
		Secretkey           string `json:"secretkey,omitempty"`
		Created             string `json:"created,omitempty"`
	} `json:"user,omitempty"`
	Volumelimit           string `json:"volumelimit,omitempty"`
	Primarystoragelimit   string `json:"primarystoragelimit,omitempty"`
	Networkavailable      string `json:"networkavailable,omitempty"`
	Projectavailable      string `json:"projectavailable,omitempty"`
	Networktotal          int    `json:"networktotal,omitempty"`
	Volumetotal           int    `json:"volumetotal,omitempty"`
	Vpcavailable          string `json:"vpcavailable,omitempty"`
	State                 string `json:"state,omitempty"`
	Domain                string `json:"domain,omitempty"`
	Secondarystoragelimit string `json:"secondarystoragelimit,omitempty"`
	Vmlimit               string `json:"vmlimit,omitempty"`
	Vpctotal              int    `json:"vpctotal,omitempty"`
	Templateavailable     string `json:"templateavailable,omitempty"`
	Snapshottotal         int    `json:"snapshottotal,omitempty"`
	Receivedbytes         int    `json:"receivedbytes,omitempty"`
	Memoryavailable       string `json:"memoryavailable,omitempty"`
	Iplimit               string `json:"iplimit,omitempty"`
	Vmstopped             int    `json:"vmstopped,omitempty"`
	Vmavailable           string `json:"vmavailable,omitempty"`
	Domainid              string `json:"domainid,omitempty"`
	Volumeavailable       string `json:"volumeavailable,omitempty"`
	Name                  string `json:"name,omitempty"`
	Networklimit          string `json:"networklimit,omitempty"`
	Snapshotavailable     string `json:"snapshotavailable,omitempty"`
}

type AddAccountToProjectParams struct {
	p map[string]interface{}
}

func (p *AddAccountToProjectParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["email"]; found {
		u.Set("email", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	return u
}

func (p *AddAccountToProjectParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *AddAccountToProjectParams) SetEmail(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["email"] = v
	return
}

func (p *AddAccountToProjectParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

// You should always use this function to get a new AddAccountToProjectParams instance,
// as then you are sure you have configured all required params
func (s *AccountService) NewAddAccountToProjectParams(projectid string) *AddAccountToProjectParams {
	p := &AddAccountToProjectParams{}
	p.p = make(map[string]interface{})
	p.p["projectid"] = projectid
	return p
}

// Adds acoount to a project
func (s *AccountService) AddAccountToProject(p *AddAccountToProjectParams) (*AddAccountToProjectResponse, error) {
	resp, err := s.cs.newRequest("addAccountToProject", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddAccountToProjectResponse
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

		var r AddAccountToProjectResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type AddAccountToProjectResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type DeleteAccountFromProjectParams struct {
	p map[string]interface{}
}

func (p *DeleteAccountFromProjectParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	return u
}

func (p *DeleteAccountFromProjectParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *DeleteAccountFromProjectParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

// You should always use this function to get a new DeleteAccountFromProjectParams instance,
// as then you are sure you have configured all required params
func (s *AccountService) NewDeleteAccountFromProjectParams(account string, projectid string) *DeleteAccountFromProjectParams {
	p := &DeleteAccountFromProjectParams{}
	p.p = make(map[string]interface{})
	p.p["account"] = account
	p.p["projectid"] = projectid
	return p
}

// Deletes account from the project
func (s *AccountService) DeleteAccountFromProject(p *DeleteAccountFromProjectParams) (*DeleteAccountFromProjectResponse, error) {
	resp, err := s.cs.newRequest("deleteAccountFromProject", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteAccountFromProjectResponse
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

		var r DeleteAccountFromProjectResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type DeleteAccountFromProjectResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type ListProjectAccountsParams struct {
	p map[string]interface{}
}

func (p *ListProjectAccountsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
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
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["role"]; found {
		u.Set("role", v.(string))
	}
	return u
}

func (p *ListProjectAccountsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListProjectAccountsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListProjectAccountsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListProjectAccountsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListProjectAccountsParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *ListProjectAccountsParams) SetRole(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["role"] = v
	return
}

// You should always use this function to get a new ListProjectAccountsParams instance,
// as then you are sure you have configured all required params
func (s *AccountService) NewListProjectAccountsParams(projectid string) *ListProjectAccountsParams {
	p := &ListProjectAccountsParams{}
	p.p = make(map[string]interface{})
	p.p["projectid"] = projectid
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AccountService) GetProjectAccountID(keyword string, projectid string) (string, error) {
	p := &ListProjectAccountsParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword
	p.p["projectid"] = projectid

	l, err := s.ListProjectAccounts(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, keyword, l)
	}
	return l.ProjectAccounts[0].Id, nil
}

// Lists project's accounts
func (s *AccountService) ListProjectAccounts(p *ListProjectAccountsParams) (*ListProjectAccountsResponse, error) {
	resp, err := s.cs.newRequest("listProjectAccounts", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListProjectAccountsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListProjectAccountsResponse struct {
	Count           int               `json:"count"`
	ProjectAccounts []*ProjectAccount `json:"projectaccount"`
}

type ProjectAccount struct {
	Networkavailable          string `json:"networkavailable,omitempty"`
	Volumeavailable           string `json:"volumeavailable,omitempty"`
	Vmtotal                   int    `json:"vmtotal,omitempty"`
	Vmavailable               string `json:"vmavailable,omitempty"`
	Cputotal                  int    `json:"cputotal,omitempty"`
	Vpclimit                  string `json:"vpclimit,omitempty"`
	Volumetotal               int    `json:"volumetotal,omitempty"`
	Vmlimit                   string `json:"vmlimit,omitempty"`
	Secondarystoragelimit     string `json:"secondarystoragelimit,omitempty"`
	Volumelimit               string `json:"volumelimit,omitempty"`
	State                     string `json:"state,omitempty"`
	Snapshotlimit             string `json:"snapshotlimit,omitempty"`
	Domain                    string `json:"domain,omitempty"`
	Displaytext               string `json:"displaytext,omitempty"`
	Snapshotavailable         string `json:"snapshotavailable,omitempty"`
	Id                        string `json:"id,omitempty"`
	Iptotal                   int    `json:"iptotal,omitempty"`
	Domainid                  string `json:"domainid,omitempty"`
	Ipavailable               string `json:"ipavailable,omitempty"`
	Cpulimit                  string `json:"cpulimit,omitempty"`
	Vpcavailable              string `json:"vpcavailable,omitempty"`
	Primarystoragetotal       int    `json:"primarystoragetotal,omitempty"`
	Secondarystorageavailable string `json:"secondarystorageavailable,omitempty"`
	Snapshottotal             int    `json:"snapshottotal,omitempty"`
	Memorylimit               string `json:"memorylimit,omitempty"`
	Iplimit                   string `json:"iplimit,omitempty"`
	Templateavailable         string `json:"templateavailable,omitempty"`
	Name                      string `json:"name,omitempty"`
	Templatetotal             int    `json:"templatetotal,omitempty"`
	Primarystorageavailable   string `json:"primarystorageavailable,omitempty"`
	Networktotal              int    `json:"networktotal,omitempty"`
	Vmstopped                 int    `json:"vmstopped,omitempty"`
	Tags                      []struct {
		Domain       string `json:"domain,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Project      string `json:"project,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Value        string `json:"value,omitempty"`
		Account      string `json:"account,omitempty"`
		Key          string `json:"key,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
	} `json:"tags,omitempty"`
	Vpctotal              int    `json:"vpctotal,omitempty"`
	Memoryavailable       string `json:"memoryavailable,omitempty"`
	Primarystoragelimit   string `json:"primarystoragelimit,omitempty"`
	Templatelimit         string `json:"templatelimit,omitempty"`
	Account               string `json:"account,omitempty"`
	Memorytotal           int    `json:"memorytotal,omitempty"`
	Vmrunning             int    `json:"vmrunning,omitempty"`
	Networklimit          string `json:"networklimit,omitempty"`
	Cpuavailable          string `json:"cpuavailable,omitempty"`
	Secondarystoragetotal int    `json:"secondarystoragetotal,omitempty"`
}
