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
		vv := strconv.Itoa(v.(int))
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
	Iscleanuprequired   bool              `json:"iscleanuprequired,omitempty"`
	Id                  string            `json:"id,omitempty"`
	Snapshotlimit       string            `json:"snapshotlimit,omitempty"`
	Memorytotal         int               `json:"memorytotal,omitempty"`
	Primarystoragelimit string            `json:"primarystoragelimit,omitempty"`
	Volumelimit         string            `json:"volumelimit,omitempty"`
	Vmlimit             string            `json:"vmlimit,omitempty"`
	Name                string            `json:"name,omitempty"`
	Accountdetails      map[string]string `json:"accountdetails,omitempty"`
	Templatelimit       string            `json:"templatelimit,omitempty"`
	Memorylimit         string            `json:"memorylimit,omitempty"`
	Vmtotal             int               `json:"vmtotal,omitempty"`
	Networklimit        string            `json:"networklimit,omitempty"`
	User                []struct {
		Secretkey           string `json:"secretkey,omitempty"`
		Isdefault           bool   `json:"isdefault,omitempty"`
		Username            string `json:"username,omitempty"`
		State               string `json:"state,omitempty"`
		Account             string `json:"account,omitempty"`
		Email               string `json:"email,omitempty"`
		Firstname           string `json:"firstname,omitempty"`
		Accounttype         int    `json:"accounttype,omitempty"`
		Accountid           string `json:"accountid,omitempty"`
		Domainid            string `json:"domainid,omitempty"`
		Id                  string `json:"id,omitempty"`
		Timezone            string `json:"timezone,omitempty"`
		Created             string `json:"created,omitempty"`
		Domain              string `json:"domain,omitempty"`
		Apikey              string `json:"apikey,omitempty"`
		Lastname            string `json:"lastname,omitempty"`
		Iscallerchilddomain bool   `json:"iscallerchilddomain,omitempty"`
	} `json:"user,omitempty"`
	Vpcavailable              string `json:"vpcavailable,omitempty"`
	Vmstopped                 int    `json:"vmstopped,omitempty"`
	Memoryavailable           string `json:"memoryavailable,omitempty"`
	Projectavailable          string `json:"projectavailable,omitempty"`
	Vpclimit                  string `json:"vpclimit,omitempty"`
	Projecttotal              int    `json:"projecttotal,omitempty"`
	Snapshottotal             int    `json:"snapshottotal,omitempty"`
	Vpctotal                  int    `json:"vpctotal,omitempty"`
	Isdefault                 bool   `json:"isdefault,omitempty"`
	Cpulimit                  string `json:"cpulimit,omitempty"`
	Vmrunning                 int    `json:"vmrunning,omitempty"`
	Snapshotavailable         string `json:"snapshotavailable,omitempty"`
	Primarystoragetotal       int    `json:"primarystoragetotal,omitempty"`
	Ipavailable               string `json:"ipavailable,omitempty"`
	Networktotal              int    `json:"networktotal,omitempty"`
	Networkdomain             string `json:"networkdomain,omitempty"`
	Secondarystorageavailable string `json:"secondarystorageavailable,omitempty"`
	Iplimit                   string `json:"iplimit,omitempty"`
	Projectlimit              string `json:"projectlimit,omitempty"`
	Volumetotal               int    `json:"volumetotal,omitempty"`
	Receivedbytes             int    `json:"receivedbytes,omitempty"`
	Primarystorageavailable   string `json:"primarystorageavailable,omitempty"`
	Secondarystoragelimit     string `json:"secondarystoragelimit,omitempty"`
	Cpuavailable              string `json:"cpuavailable,omitempty"`
	Vmavailable               string `json:"vmavailable,omitempty"`
	Accounttype               int    `json:"accounttype,omitempty"`
	Domainid                  string `json:"domainid,omitempty"`
	Volumeavailable           string `json:"volumeavailable,omitempty"`
	Templatetotal             int    `json:"templatetotal,omitempty"`
	Defaultzoneid             string `json:"defaultzoneid,omitempty"`
	Networkavailable          string `json:"networkavailable,omitempty"`
	State                     string `json:"state,omitempty"`
	Templateavailable         string `json:"templateavailable,omitempty"`
	Cputotal                  int    `json:"cputotal,omitempty"`
	Domain                    string `json:"domain,omitempty"`
	Sentbytes                 int    `json:"sentbytes,omitempty"`
	Iptotal                   int    `json:"iptotal,omitempty"`
	Secondarystoragetotal     int    `json:"secondarystoragetotal,omitempty"`
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

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type DeleteAccountResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
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
	Memorylimit           string `json:"memorylimit,omitempty"`
	Id                    string `json:"id,omitempty"`
	Vpctotal              int    `json:"vpctotal,omitempty"`
	Snapshottotal         int    `json:"snapshottotal,omitempty"`
	Vmrunning             int    `json:"vmrunning,omitempty"`
	Secondarystoragelimit string `json:"secondarystoragelimit,omitempty"`
	User                  []struct {
		Accountid           string `json:"accountid,omitempty"`
		Firstname           string `json:"firstname,omitempty"`
		Timezone            string `json:"timezone,omitempty"`
		Accounttype         int    `json:"accounttype,omitempty"`
		Secretkey           string `json:"secretkey,omitempty"`
		Lastname            string `json:"lastname,omitempty"`
		Isdefault           bool   `json:"isdefault,omitempty"`
		Created             string `json:"created,omitempty"`
		State               string `json:"state,omitempty"`
		Domainid            string `json:"domainid,omitempty"`
		Account             string `json:"account,omitempty"`
		Domain              string `json:"domain,omitempty"`
		Email               string `json:"email,omitempty"`
		Apikey              string `json:"apikey,omitempty"`
		Username            string `json:"username,omitempty"`
		Iscallerchilddomain bool   `json:"iscallerchilddomain,omitempty"`
		Id                  string `json:"id,omitempty"`
	} `json:"user,omitempty"`
	Vmlimit                   string            `json:"vmlimit,omitempty"`
	Vmavailable               string            `json:"vmavailable,omitempty"`
	Isdefault                 bool              `json:"isdefault,omitempty"`
	Networkavailable          string            `json:"networkavailable,omitempty"`
	Secondarystoragetotal     int               `json:"secondarystoragetotal,omitempty"`
	Vpcavailable              string            `json:"vpcavailable,omitempty"`
	Snapshotlimit             string            `json:"snapshotlimit,omitempty"`
	Volumetotal               int               `json:"volumetotal,omitempty"`
	Cpulimit                  string            `json:"cpulimit,omitempty"`
	Snapshotavailable         string            `json:"snapshotavailable,omitempty"`
	Cputotal                  int               `json:"cputotal,omitempty"`
	Cpuavailable              string            `json:"cpuavailable,omitempty"`
	Primarystoragelimit       string            `json:"primarystoragelimit,omitempty"`
	Defaultzoneid             string            `json:"defaultzoneid,omitempty"`
	Projectavailable          string            `json:"projectavailable,omitempty"`
	Vmstopped                 int               `json:"vmstopped,omitempty"`
	Templatetotal             int               `json:"templatetotal,omitempty"`
	Projectlimit              string            `json:"projectlimit,omitempty"`
	Domainid                  string            `json:"domainid,omitempty"`
	Iplimit                   string            `json:"iplimit,omitempty"`
	Templatelimit             string            `json:"templatelimit,omitempty"`
	Memoryavailable           string            `json:"memoryavailable,omitempty"`
	Sentbytes                 int               `json:"sentbytes,omitempty"`
	Primarystoragetotal       int               `json:"primarystoragetotal,omitempty"`
	Iptotal                   int               `json:"iptotal,omitempty"`
	Iscleanuprequired         bool              `json:"iscleanuprequired,omitempty"`
	Networktotal              int               `json:"networktotal,omitempty"`
	Accountdetails            map[string]string `json:"accountdetails,omitempty"`
	Primarystorageavailable   string            `json:"primarystorageavailable,omitempty"`
	Projecttotal              int               `json:"projecttotal,omitempty"`
	Networklimit              string            `json:"networklimit,omitempty"`
	Vpclimit                  string            `json:"vpclimit,omitempty"`
	Vmtotal                   int               `json:"vmtotal,omitempty"`
	Templateavailable         string            `json:"templateavailable,omitempty"`
	Networkdomain             string            `json:"networkdomain,omitempty"`
	Receivedbytes             int               `json:"receivedbytes,omitempty"`
	Ipavailable               string            `json:"ipavailable,omitempty"`
	Memorytotal               int               `json:"memorytotal,omitempty"`
	Accounttype               int               `json:"accounttype,omitempty"`
	Domain                    string            `json:"domain,omitempty"`
	Volumeavailable           string            `json:"volumeavailable,omitempty"`
	Secondarystorageavailable string            `json:"secondarystorageavailable,omitempty"`
	Name                      string            `json:"name,omitempty"`
	State                     string            `json:"state,omitempty"`
	Volumelimit               string            `json:"volumelimit,omitempty"`
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

type DisableAccountResponse struct {
	JobID                 string `json:"jobid,omitempty"`
	Iplimit               string `json:"iplimit,omitempty"`
	Vmtotal               int    `json:"vmtotal,omitempty"`
	Volumeavailable       string `json:"volumeavailable,omitempty"`
	Networkavailable      string `json:"networkavailable,omitempty"`
	Snapshotlimit         string `json:"snapshotlimit,omitempty"`
	Secondarystoragetotal int    `json:"secondarystoragetotal,omitempty"`
	Vpcavailable          string `json:"vpcavailable,omitempty"`
	Receivedbytes         int    `json:"receivedbytes,omitempty"`
	Cpuavailable          string `json:"cpuavailable,omitempty"`
	Name                  string `json:"name,omitempty"`
	Memoryavailable       string `json:"memoryavailable,omitempty"`
	Memorytotal           int    `json:"memorytotal,omitempty"`
	Volumetotal           int    `json:"volumetotal,omitempty"`
	Domain                string `json:"domain,omitempty"`
	Projecttotal          int    `json:"projecttotal,omitempty"`
	Snapshottotal         int    `json:"snapshottotal,omitempty"`
	Volumelimit           string `json:"volumelimit,omitempty"`
	Networklimit          string `json:"networklimit,omitempty"`
	Templatetotal         int    `json:"templatetotal,omitempty"`
	Iptotal               int    `json:"iptotal,omitempty"`
	Projectavailable      string `json:"projectavailable,omitempty"`
	Secondarystoragelimit string `json:"secondarystoragelimit,omitempty"`
	Iscleanuprequired     bool   `json:"iscleanuprequired,omitempty"`
	Primarystoragelimit   string `json:"primarystoragelimit,omitempty"`
	Projectlimit          string `json:"projectlimit,omitempty"`
	Networktotal          int    `json:"networktotal,omitempty"`
	Memorylimit           string `json:"memorylimit,omitempty"`
	State                 string `json:"state,omitempty"`
	Cputotal              int    `json:"cputotal,omitempty"`
	User                  []struct {
		Domainid            string `json:"domainid,omitempty"`
		Username            string `json:"username,omitempty"`
		Account             string `json:"account,omitempty"`
		Id                  string `json:"id,omitempty"`
		Lastname            string `json:"lastname,omitempty"`
		Timezone            string `json:"timezone,omitempty"`
		Created             string `json:"created,omitempty"`
		Domain              string `json:"domain,omitempty"`
		State               string `json:"state,omitempty"`
		Email               string `json:"email,omitempty"`
		Accounttype         int    `json:"accounttype,omitempty"`
		Accountid           string `json:"accountid,omitempty"`
		Apikey              string `json:"apikey,omitempty"`
		Secretkey           string `json:"secretkey,omitempty"`
		Firstname           string `json:"firstname,omitempty"`
		Isdefault           bool   `json:"isdefault,omitempty"`
		Iscallerchilddomain bool   `json:"iscallerchilddomain,omitempty"`
	} `json:"user,omitempty"`
	Vpctotal                  int               `json:"vpctotal,omitempty"`
	Accountdetails            map[string]string `json:"accountdetails,omitempty"`
	Primarystoragetotal       int               `json:"primarystoragetotal,omitempty"`
	Accounttype               int               `json:"accounttype,omitempty"`
	Defaultzoneid             string            `json:"defaultzoneid,omitempty"`
	Networkdomain             string            `json:"networkdomain,omitempty"`
	Domainid                  string            `json:"domainid,omitempty"`
	Sentbytes                 int               `json:"sentbytes,omitempty"`
	Templateavailable         string            `json:"templateavailable,omitempty"`
	Ipavailable               string            `json:"ipavailable,omitempty"`
	Vmavailable               string            `json:"vmavailable,omitempty"`
	Primarystorageavailable   string            `json:"primarystorageavailable,omitempty"`
	Vpclimit                  string            `json:"vpclimit,omitempty"`
	Secondarystorageavailable string            `json:"secondarystorageavailable,omitempty"`
	Id                        string            `json:"id,omitempty"`
	Snapshotavailable         string            `json:"snapshotavailable,omitempty"`
	Vmstopped                 int               `json:"vmstopped,omitempty"`
	Isdefault                 bool              `json:"isdefault,omitempty"`
	Vmlimit                   string            `json:"vmlimit,omitempty"`
	Cpulimit                  string            `json:"cpulimit,omitempty"`
	Templatelimit             string            `json:"templatelimit,omitempty"`
	Vmrunning                 int               `json:"vmrunning,omitempty"`
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
	Snapshottotal       int    `json:"snapshottotal,omitempty"`
	Cpulimit            string `json:"cpulimit,omitempty"`
	Networkavailable    string `json:"networkavailable,omitempty"`
	Projectavailable    string `json:"projectavailable,omitempty"`
	Ipavailable         string `json:"ipavailable,omitempty"`
	Iptotal             int    `json:"iptotal,omitempty"`
	Primarystoragelimit string `json:"primarystoragelimit,omitempty"`
	Cpuavailable        string `json:"cpuavailable,omitempty"`
	Vmlimit             string `json:"vmlimit,omitempty"`
	User                []struct {
		Firstname           string `json:"firstname,omitempty"`
		Domain              string `json:"domain,omitempty"`
		Isdefault           bool   `json:"isdefault,omitempty"`
		Email               string `json:"email,omitempty"`
		Secretkey           string `json:"secretkey,omitempty"`
		State               string `json:"state,omitempty"`
		Accountid           string `json:"accountid,omitempty"`
		Iscallerchilddomain bool   `json:"iscallerchilddomain,omitempty"`
		Lastname            string `json:"lastname,omitempty"`
		Timezone            string `json:"timezone,omitempty"`
		Accounttype         int    `json:"accounttype,omitempty"`
		Id                  string `json:"id,omitempty"`
		Username            string `json:"username,omitempty"`
		Created             string `json:"created,omitempty"`
		Apikey              string `json:"apikey,omitempty"`
		Account             string `json:"account,omitempty"`
		Domainid            string `json:"domainid,omitempty"`
	} `json:"user,omitempty"`
	Secondarystoragelimit     string            `json:"secondarystoragelimit,omitempty"`
	Volumelimit               string            `json:"volumelimit,omitempty"`
	Networklimit              string            `json:"networklimit,omitempty"`
	Templatelimit             string            `json:"templatelimit,omitempty"`
	Snapshotlimit             string            `json:"snapshotlimit,omitempty"`
	Memorytotal               int               `json:"memorytotal,omitempty"`
	Networkdomain             string            `json:"networkdomain,omitempty"`
	Secondarystoragetotal     int               `json:"secondarystoragetotal,omitempty"`
	Projecttotal              int               `json:"projecttotal,omitempty"`
	Secondarystorageavailable string            `json:"secondarystorageavailable,omitempty"`
	Vmstopped                 int               `json:"vmstopped,omitempty"`
	Id                        string            `json:"id,omitempty"`
	Accountdetails            map[string]string `json:"accountdetails,omitempty"`
	Receivedbytes             int               `json:"receivedbytes,omitempty"`
	Templatetotal             int               `json:"templatetotal,omitempty"`
	Iplimit                   string            `json:"iplimit,omitempty"`
	Vpctotal                  int               `json:"vpctotal,omitempty"`
	Vpclimit                  string            `json:"vpclimit,omitempty"`
	Sentbytes                 int               `json:"sentbytes,omitempty"`
	Memorylimit               string            `json:"memorylimit,omitempty"`
	Primarystoragetotal       int               `json:"primarystoragetotal,omitempty"`
	Cputotal                  int               `json:"cputotal,omitempty"`
	Vmrunning                 int               `json:"vmrunning,omitempty"`
	Vmavailable               string            `json:"vmavailable,omitempty"`
	Iscleanuprequired         bool              `json:"iscleanuprequired,omitempty"`
	Networktotal              int               `json:"networktotal,omitempty"`
	Volumeavailable           string            `json:"volumeavailable,omitempty"`
	Templateavailable         string            `json:"templateavailable,omitempty"`
	Name                      string            `json:"name,omitempty"`
	Vmtotal                   int               `json:"vmtotal,omitempty"`
	Accounttype               int               `json:"accounttype,omitempty"`
	Domain                    string            `json:"domain,omitempty"`
	Defaultzoneid             string            `json:"defaultzoneid,omitempty"`
	Volumetotal               int               `json:"volumetotal,omitempty"`
	Memoryavailable           string            `json:"memoryavailable,omitempty"`
	Projectlimit              string            `json:"projectlimit,omitempty"`
	State                     string            `json:"state,omitempty"`
	Isdefault                 bool              `json:"isdefault,omitempty"`
	Snapshotavailable         string            `json:"snapshotavailable,omitempty"`
	Primarystorageavailable   string            `json:"primarystorageavailable,omitempty"`
	Vpcavailable              string            `json:"vpcavailable,omitempty"`
	Domainid                  string            `json:"domainid,omitempty"`
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
	Templatetotal int    `json:"templatetotal,omitempty"`
	Vmlimit       string `json:"vmlimit,omitempty"`
	Name          string `json:"name,omitempty"`
	User          []struct {
		Domainid            string `json:"domainid,omitempty"`
		Accountid           string `json:"accountid,omitempty"`
		Secretkey           string `json:"secretkey,omitempty"`
		Domain              string `json:"domain,omitempty"`
		Created             string `json:"created,omitempty"`
		Email               string `json:"email,omitempty"`
		Id                  string `json:"id,omitempty"`
		Timezone            string `json:"timezone,omitempty"`
		Isdefault           bool   `json:"isdefault,omitempty"`
		Accounttype         int    `json:"accounttype,omitempty"`
		Apikey              string `json:"apikey,omitempty"`
		Firstname           string `json:"firstname,omitempty"`
		State               string `json:"state,omitempty"`
		Account             string `json:"account,omitempty"`
		Lastname            string `json:"lastname,omitempty"`
		Iscallerchilddomain bool   `json:"iscallerchilddomain,omitempty"`
		Username            string `json:"username,omitempty"`
	} `json:"user,omitempty"`
	Vpcavailable              string            `json:"vpcavailable,omitempty"`
	Memorytotal               int               `json:"memorytotal,omitempty"`
	Id                        string            `json:"id,omitempty"`
	Sentbytes                 int               `json:"sentbytes,omitempty"`
	Snapshotavailable         string            `json:"snapshotavailable,omitempty"`
	Networktotal              int               `json:"networktotal,omitempty"`
	State                     string            `json:"state,omitempty"`
	Vmrunning                 int               `json:"vmrunning,omitempty"`
	Cpulimit                  string            `json:"cpulimit,omitempty"`
	Secondarystoragetotal     int               `json:"secondarystoragetotal,omitempty"`
	Domainid                  string            `json:"domainid,omitempty"`
	Projecttotal              int               `json:"projecttotal,omitempty"`
	Cpuavailable              string            `json:"cpuavailable,omitempty"`
	Vmavailable               string            `json:"vmavailable,omitempty"`
	Volumelimit               string            `json:"volumelimit,omitempty"`
	Defaultzoneid             string            `json:"defaultzoneid,omitempty"`
	Memorylimit               string            `json:"memorylimit,omitempty"`
	Iscleanuprequired         bool              `json:"iscleanuprequired,omitempty"`
	Projectavailable          string            `json:"projectavailable,omitempty"`
	Cputotal                  int               `json:"cputotal,omitempty"`
	Networkdomain             string            `json:"networkdomain,omitempty"`
	Primarystorageavailable   string            `json:"primarystorageavailable,omitempty"`
	Iplimit                   string            `json:"iplimit,omitempty"`
	Secondarystoragelimit     string            `json:"secondarystoragelimit,omitempty"`
	Accountdetails            map[string]string `json:"accountdetails,omitempty"`
	Primarystoragelimit       string            `json:"primarystoragelimit,omitempty"`
	Secondarystorageavailable string            `json:"secondarystorageavailable,omitempty"`
	Volumetotal               int               `json:"volumetotal,omitempty"`
	Templateavailable         string            `json:"templateavailable,omitempty"`
	Volumeavailable           string            `json:"volumeavailable,omitempty"`
	Accounttype               int               `json:"accounttype,omitempty"`
	Networklimit              string            `json:"networklimit,omitempty"`
	Domain                    string            `json:"domain,omitempty"`
	Templatelimit             string            `json:"templatelimit,omitempty"`
	Ipavailable               string            `json:"ipavailable,omitempty"`
	Isdefault                 bool              `json:"isdefault,omitempty"`
	Memoryavailable           string            `json:"memoryavailable,omitempty"`
	Vmstopped                 int               `json:"vmstopped,omitempty"`
	Projectlimit              string            `json:"projectlimit,omitempty"`
	Primarystoragetotal       int               `json:"primarystoragetotal,omitempty"`
	Snapshottotal             int               `json:"snapshottotal,omitempty"`
	Vmtotal                   int               `json:"vmtotal,omitempty"`
	Snapshotlimit             string            `json:"snapshotlimit,omitempty"`
	Receivedbytes             int               `json:"receivedbytes,omitempty"`
	Vpctotal                  int               `json:"vpctotal,omitempty"`
	Networkavailable          string            `json:"networkavailable,omitempty"`
	Vpclimit                  string            `json:"vpclimit,omitempty"`
	Iptotal                   int               `json:"iptotal,omitempty"`
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
		vv := strconv.Itoa(v.(int))
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
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
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

	if l.Count == 0 {
		return "", fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Accounts[0].Id, nil
	}

	if l.Count > 1 {
		for _, v := range l.Accounts {
			if v.Name == name {
				return v.Id, nil
			}
		}
	}
	return "", fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AccountService) GetAccountByName(name string) (*Account, int, error) {
	id, err := s.GetAccountID(name)
	if err != nil {
		return nil, -1, err
	}

	r, count, err := s.GetAccountByID(id)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AccountService) GetAccountByID(id string) (*Account, int, error) {
	p := &ListAccountsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	l, err := s.ListAccounts(p)
	if err != nil {
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.Accounts[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Account UUID: %s!", id)
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
	Cpulimit                string `json:"cpulimit,omitempty"`
	Vpcavailable            string `json:"vpcavailable,omitempty"`
	Templateavailable       string `json:"templateavailable,omitempty"`
	Accounttype             int    `json:"accounttype,omitempty"`
	Vmavailable             string `json:"vmavailable,omitempty"`
	Vmstopped               int    `json:"vmstopped,omitempty"`
	Cputotal                int    `json:"cputotal,omitempty"`
	Primarystoragetotal     int    `json:"primarystoragetotal,omitempty"`
	Vpclimit                string `json:"vpclimit,omitempty"`
	Templatetotal           int    `json:"templatetotal,omitempty"`
	Receivedbytes           int    `json:"receivedbytes,omitempty"`
	Networklimit            string `json:"networklimit,omitempty"`
	Volumelimit             string `json:"volumelimit,omitempty"`
	Memorytotal             int    `json:"memorytotal,omitempty"`
	Primarystorageavailable string `json:"primarystorageavailable,omitempty"`
	Domain                  string `json:"domain,omitempty"`
	Snapshotlimit           string `json:"snapshotlimit,omitempty"`
	User                    []struct {
		Timezone            string `json:"timezone,omitempty"`
		Username            string `json:"username,omitempty"`
		Account             string `json:"account,omitempty"`
		Id                  string `json:"id,omitempty"`
		Lastname            string `json:"lastname,omitempty"`
		Isdefault           bool   `json:"isdefault,omitempty"`
		Apikey              string `json:"apikey,omitempty"`
		Created             string `json:"created,omitempty"`
		Iscallerchilddomain bool   `json:"iscallerchilddomain,omitempty"`
		Secretkey           string `json:"secretkey,omitempty"`
		Accounttype         int    `json:"accounttype,omitempty"`
		State               string `json:"state,omitempty"`
		Domainid            string `json:"domainid,omitempty"`
		Email               string `json:"email,omitempty"`
		Accountid           string `json:"accountid,omitempty"`
		Firstname           string `json:"firstname,omitempty"`
		Domain              string `json:"domain,omitempty"`
	} `json:"user,omitempty"`
	Vpctotal                  int               `json:"vpctotal,omitempty"`
	Accountdetails            map[string]string `json:"accountdetails,omitempty"`
	Iscleanuprequired         bool              `json:"iscleanuprequired,omitempty"`
	Snapshottotal             int               `json:"snapshottotal,omitempty"`
	Isdefault                 bool              `json:"isdefault,omitempty"`
	Vmtotal                   int               `json:"vmtotal,omitempty"`
	Memoryavailable           string            `json:"memoryavailable,omitempty"`
	Secondarystoragetotal     int               `json:"secondarystoragetotal,omitempty"`
	Iplimit                   string            `json:"iplimit,omitempty"`
	Memorylimit               string            `json:"memorylimit,omitempty"`
	Projectavailable          string            `json:"projectavailable,omitempty"`
	Vmrunning                 int               `json:"vmrunning,omitempty"`
	Projecttotal              int               `json:"projecttotal,omitempty"`
	Iptotal                   int               `json:"iptotal,omitempty"`
	Networkavailable          string            `json:"networkavailable,omitempty"`
	Secondarystoragelimit     string            `json:"secondarystoragelimit,omitempty"`
	Name                      string            `json:"name,omitempty"`
	Primarystoragelimit       string            `json:"primarystoragelimit,omitempty"`
	Defaultzoneid             string            `json:"defaultzoneid,omitempty"`
	Networkdomain             string            `json:"networkdomain,omitempty"`
	Networktotal              int               `json:"networktotal,omitempty"`
	Snapshotavailable         string            `json:"snapshotavailable,omitempty"`
	Volumetotal               int               `json:"volumetotal,omitempty"`
	Secondarystorageavailable string            `json:"secondarystorageavailable,omitempty"`
	Domainid                  string            `json:"domainid,omitempty"`
	Cpuavailable              string            `json:"cpuavailable,omitempty"`
	Vmlimit                   string            `json:"vmlimit,omitempty"`
	State                     string            `json:"state,omitempty"`
	Templatelimit             string            `json:"templatelimit,omitempty"`
	Sentbytes                 int               `json:"sentbytes,omitempty"`
	Id                        string            `json:"id,omitempty"`
	Ipavailable               string            `json:"ipavailable,omitempty"`
	Projectlimit              string            `json:"projectlimit,omitempty"`
	Volumeavailable           string            `json:"volumeavailable,omitempty"`
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

type MarkDefaultZoneForAccountResponse struct {
	JobID                     string            `json:"jobid,omitempty"`
	Projectlimit              string            `json:"projectlimit,omitempty"`
	Cputotal                  int               `json:"cputotal,omitempty"`
	Accountdetails            map[string]string `json:"accountdetails,omitempty"`
	Receivedbytes             int               `json:"receivedbytes,omitempty"`
	Vmrunning                 int               `json:"vmrunning,omitempty"`
	Id                        string            `json:"id,omitempty"`
	Volumetotal               int               `json:"volumetotal,omitempty"`
	Primarystorageavailable   string            `json:"primarystorageavailable,omitempty"`
	Isdefault                 bool              `json:"isdefault,omitempty"`
	Name                      string            `json:"name,omitempty"`
	Iplimit                   string            `json:"iplimit,omitempty"`
	Domainid                  string            `json:"domainid,omitempty"`
	Accounttype               int               `json:"accounttype,omitempty"`
	Networktotal              int               `json:"networktotal,omitempty"`
	Vmtotal                   int               `json:"vmtotal,omitempty"`
	Primarystoragelimit       string            `json:"primarystoragelimit,omitempty"`
	State                     string            `json:"state,omitempty"`
	Projecttotal              int               `json:"projecttotal,omitempty"`
	Domain                    string            `json:"domain,omitempty"`
	Vmlimit                   string            `json:"vmlimit,omitempty"`
	Memorytotal               int               `json:"memorytotal,omitempty"`
	Memoryavailable           string            `json:"memoryavailable,omitempty"`
	Networkdomain             string            `json:"networkdomain,omitempty"`
	Snapshotlimit             string            `json:"snapshotlimit,omitempty"`
	Vmavailable               string            `json:"vmavailable,omitempty"`
	Volumeavailable           string            `json:"volumeavailable,omitempty"`
	Iscleanuprequired         bool              `json:"iscleanuprequired,omitempty"`
	Vmstopped                 int               `json:"vmstopped,omitempty"`
	Volumelimit               string            `json:"volumelimit,omitempty"`
	Networkavailable          string            `json:"networkavailable,omitempty"`
	Secondarystorageavailable string            `json:"secondarystorageavailable,omitempty"`
	Templateavailable         string            `json:"templateavailable,omitempty"`
	Sentbytes                 int               `json:"sentbytes,omitempty"`
	Iptotal                   int               `json:"iptotal,omitempty"`
	Secondarystoragetotal     int               `json:"secondarystoragetotal,omitempty"`
	Snapshotavailable         string            `json:"snapshotavailable,omitempty"`
	Cpulimit                  string            `json:"cpulimit,omitempty"`
	Templatelimit             string            `json:"templatelimit,omitempty"`
	Vpctotal                  int               `json:"vpctotal,omitempty"`
	Vpcavailable              string            `json:"vpcavailable,omitempty"`
	Defaultzoneid             string            `json:"defaultzoneid,omitempty"`
	Networklimit              string            `json:"networklimit,omitempty"`
	Memorylimit               string            `json:"memorylimit,omitempty"`
	Vpclimit                  string            `json:"vpclimit,omitempty"`
	Ipavailable               string            `json:"ipavailable,omitempty"`
	Cpuavailable              string            `json:"cpuavailable,omitempty"`
	User                      []struct {
		Email               string `json:"email,omitempty"`
		Firstname           string `json:"firstname,omitempty"`
		Isdefault           bool   `json:"isdefault,omitempty"`
		Accounttype         int    `json:"accounttype,omitempty"`
		State               string `json:"state,omitempty"`
		Domain              string `json:"domain,omitempty"`
		Timezone            string `json:"timezone,omitempty"`
		Username            string `json:"username,omitempty"`
		Account             string `json:"account,omitempty"`
		Apikey              string `json:"apikey,omitempty"`
		Domainid            string `json:"domainid,omitempty"`
		Iscallerchilddomain bool   `json:"iscallerchilddomain,omitempty"`
		Created             string `json:"created,omitempty"`
		Id                  string `json:"id,omitempty"`
		Secretkey           string `json:"secretkey,omitempty"`
		Accountid           string `json:"accountid,omitempty"`
		Lastname            string `json:"lastname,omitempty"`
	} `json:"user,omitempty"`
	Projectavailable      string `json:"projectavailable,omitempty"`
	Secondarystoragelimit string `json:"secondarystoragelimit,omitempty"`
	Primarystoragetotal   int    `json:"primarystoragetotal,omitempty"`
	Snapshottotal         int    `json:"snapshottotal,omitempty"`
	Templatetotal         int    `json:"templatetotal,omitempty"`
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

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type AddAccountToProjectResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Success     bool   `json:"success,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
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

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type DeleteAccountFromProjectResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Success     bool   `json:"success,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
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
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
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

	if l.Count == 0 {
		return "", fmt.Errorf("No match found for %s: %+v", keyword, l)
	}

	if l.Count == 1 {
		return l.ProjectAccounts[0].Id, nil
	}

	if l.Count > 1 {
		for _, v := range l.ProjectAccounts {
			if v.Name == keyword {
				return v.Id, nil
			}
		}
	}
	return "", fmt.Errorf("Could not find an exact match for %s: %+v", keyword, l)
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
	Vmtotal               int    `json:"vmtotal,omitempty"`
	Account               string `json:"account,omitempty"`
	Vmrunning             int    `json:"vmrunning,omitempty"`
	Snapshotlimit         string `json:"snapshotlimit,omitempty"`
	Cpuavailable          string `json:"cpuavailable,omitempty"`
	Volumetotal           int    `json:"volumetotal,omitempty"`
	Volumelimit           string `json:"volumelimit,omitempty"`
	Name                  string `json:"name,omitempty"`
	Snapshottotal         int    `json:"snapshottotal,omitempty"`
	Networkavailable      string `json:"networkavailable,omitempty"`
	Vpctotal              int    `json:"vpctotal,omitempty"`
	Vmlimit               string `json:"vmlimit,omitempty"`
	Templatetotal         int    `json:"templatetotal,omitempty"`
	Iptotal               int    `json:"iptotal,omitempty"`
	Id                    string `json:"id,omitempty"`
	Memorylimit           string `json:"memorylimit,omitempty"`
	Vpcavailable          string `json:"vpcavailable,omitempty"`
	Primarystoragetotal   int    `json:"primarystoragetotal,omitempty"`
	Snapshotavailable     string `json:"snapshotavailable,omitempty"`
	Vmavailable           string `json:"vmavailable,omitempty"`
	Displaytext           string `json:"displaytext,omitempty"`
	Networklimit          string `json:"networklimit,omitempty"`
	Secondarystoragelimit string `json:"secondarystoragelimit,omitempty"`
	Volumeavailable       string `json:"volumeavailable,omitempty"`
	Tags                  []struct {
		Account      string `json:"account,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Key          string `json:"key,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Project      string `json:"project,omitempty"`
		Value        string `json:"value,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
	} `json:"tags,omitempty"`
	Cpulimit                  string `json:"cpulimit,omitempty"`
	Ipavailable               string `json:"ipavailable,omitempty"`
	Primarystorageavailable   string `json:"primarystorageavailable,omitempty"`
	Templateavailable         string `json:"templateavailable,omitempty"`
	Vpclimit                  string `json:"vpclimit,omitempty"`
	Iplimit                   string `json:"iplimit,omitempty"`
	Secondarystoragetotal     int    `json:"secondarystoragetotal,omitempty"`
	State                     string `json:"state,omitempty"`
	Domain                    string `json:"domain,omitempty"`
	Domainid                  string `json:"domainid,omitempty"`
	Networktotal              int    `json:"networktotal,omitempty"`
	Templatelimit             string `json:"templatelimit,omitempty"`
	Cputotal                  int    `json:"cputotal,omitempty"`
	Memorytotal               int    `json:"memorytotal,omitempty"`
	Memoryavailable           string `json:"memoryavailable,omitempty"`
	Vmstopped                 int    `json:"vmstopped,omitempty"`
	Primarystoragelimit       string `json:"primarystoragelimit,omitempty"`
	Secondarystorageavailable string `json:"secondarystorageavailable,omitempty"`
}
