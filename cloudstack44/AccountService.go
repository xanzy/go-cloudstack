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
	Cpuavailable string `json:"cpuavailable,omitempty"`
	User         []struct {
		Isdefault           bool   `json:"isdefault,omitempty"`
		Firstname           string `json:"firstname,omitempty"`
		Account             string `json:"account,omitempty"`
		Domainid            string `json:"domainid,omitempty"`
		Timezone            string `json:"timezone,omitempty"`
		Id                  string `json:"id,omitempty"`
		Accountid           string `json:"accountid,omitempty"`
		Iscallerchilddomain bool   `json:"iscallerchilddomain,omitempty"`
		Accounttype         int    `json:"accounttype,omitempty"`
		Domain              string `json:"domain,omitempty"`
		Apikey              string `json:"apikey,omitempty"`
		Lastname            string `json:"lastname,omitempty"`
		Username            string `json:"username,omitempty"`
		Email               string `json:"email,omitempty"`
		State               string `json:"state,omitempty"`
		Created             string `json:"created,omitempty"`
		Secretkey           string `json:"secretkey,omitempty"`
	} `json:"user,omitempty"`
	Templateavailable         string            `json:"templateavailable,omitempty"`
	Primarystoragelimit       string            `json:"primarystoragelimit,omitempty"`
	Snapshottotal             int               `json:"snapshottotal,omitempty"`
	Vpcavailable              string            `json:"vpcavailable,omitempty"`
	Memoryavailable           string            `json:"memoryavailable,omitempty"`
	Projectavailable          string            `json:"projectavailable,omitempty"`
	Defaultzoneid             string            `json:"defaultzoneid,omitempty"`
	State                     string            `json:"state,omitempty"`
	Iplimit                   string            `json:"iplimit,omitempty"`
	Vmrunning                 int               `json:"vmrunning,omitempty"`
	Cpulimit                  string            `json:"cpulimit,omitempty"`
	Id                        string            `json:"id,omitempty"`
	Networkavailable          string            `json:"networkavailable,omitempty"`
	Vmavailable               string            `json:"vmavailable,omitempty"`
	Ipavailable               string            `json:"ipavailable,omitempty"`
	Volumeavailable           string            `json:"volumeavailable,omitempty"`
	Domainid                  string            `json:"domainid,omitempty"`
	Networklimit              string            `json:"networklimit,omitempty"`
	Projectlimit              string            `json:"projectlimit,omitempty"`
	Networkdomain             string            `json:"networkdomain,omitempty"`
	Vmtotal                   int               `json:"vmtotal,omitempty"`
	Primarystorageavailable   string            `json:"primarystorageavailable,omitempty"`
	Cputotal                  int               `json:"cputotal,omitempty"`
	Snapshotlimit             string            `json:"snapshotlimit,omitempty"`
	Templatetotal             int               `json:"templatetotal,omitempty"`
	Vpctotal                  int               `json:"vpctotal,omitempty"`
	Secondarystoragelimit     string            `json:"secondarystoragelimit,omitempty"`
	Snapshotavailable         string            `json:"snapshotavailable,omitempty"`
	Vpclimit                  string            `json:"vpclimit,omitempty"`
	Memorytotal               int               `json:"memorytotal,omitempty"`
	Accountdetails            map[string]string `json:"accountdetails,omitempty"`
	Templatelimit             string            `json:"templatelimit,omitempty"`
	Iscleanuprequired         bool              `json:"iscleanuprequired,omitempty"`
	Name                      string            `json:"name,omitempty"`
	Primarystoragetotal       int               `json:"primarystoragetotal,omitempty"`
	Receivedbytes             int               `json:"receivedbytes,omitempty"`
	Networktotal              int               `json:"networktotal,omitempty"`
	Volumetotal               int               `json:"volumetotal,omitempty"`
	Vmlimit                   string            `json:"vmlimit,omitempty"`
	Domain                    string            `json:"domain,omitempty"`
	Secondarystoragetotal     int               `json:"secondarystoragetotal,omitempty"`
	Isdefault                 bool              `json:"isdefault,omitempty"`
	Memorylimit               string            `json:"memorylimit,omitempty"`
	Vmstopped                 int               `json:"vmstopped,omitempty"`
	Iptotal                   int               `json:"iptotal,omitempty"`
	Volumelimit               string            `json:"volumelimit,omitempty"`
	Projecttotal              int               `json:"projecttotal,omitempty"`
	Sentbytes                 int               `json:"sentbytes,omitempty"`
	Accounttype               int               `json:"accounttype,omitempty"`
	Secondarystorageavailable string            `json:"secondarystorageavailable,omitempty"`
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
	Cputotal                int               `json:"cputotal,omitempty"`
	Memorytotal             int               `json:"memorytotal,omitempty"`
	Defaultzoneid           string            `json:"defaultzoneid,omitempty"`
	Networkdomain           string            `json:"networkdomain,omitempty"`
	Templateavailable       string            `json:"templateavailable,omitempty"`
	Name                    string            `json:"name,omitempty"`
	Cpulimit                string            `json:"cpulimit,omitempty"`
	Vpctotal                int               `json:"vpctotal,omitempty"`
	Iplimit                 string            `json:"iplimit,omitempty"`
	Memoryavailable         string            `json:"memoryavailable,omitempty"`
	Iptotal                 int               `json:"iptotal,omitempty"`
	Vmstopped               int               `json:"vmstopped,omitempty"`
	Templatetotal           int               `json:"templatetotal,omitempty"`
	Volumeavailable         string            `json:"volumeavailable,omitempty"`
	Snapshotavailable       string            `json:"snapshotavailable,omitempty"`
	Receivedbytes           int               `json:"receivedbytes,omitempty"`
	Accountdetails          map[string]string `json:"accountdetails,omitempty"`
	Vmrunning               int               `json:"vmrunning,omitempty"`
	Domainid                string            `json:"domainid,omitempty"`
	Sentbytes               int               `json:"sentbytes,omitempty"`
	Projectlimit            string            `json:"projectlimit,omitempty"`
	Secondarystoragelimit   string            `json:"secondarystoragelimit,omitempty"`
	Networktotal            int               `json:"networktotal,omitempty"`
	Snapshotlimit           string            `json:"snapshotlimit,omitempty"`
	Projectavailable        string            `json:"projectavailable,omitempty"`
	Vpclimit                string            `json:"vpclimit,omitempty"`
	Volumetotal             int               `json:"volumetotal,omitempty"`
	Memorylimit             string            `json:"memorylimit,omitempty"`
	Id                      string            `json:"id,omitempty"`
	Ipavailable             string            `json:"ipavailable,omitempty"`
	Vpcavailable            string            `json:"vpcavailable,omitempty"`
	State                   string            `json:"state,omitempty"`
	Snapshottotal           int               `json:"snapshottotal,omitempty"`
	Domain                  string            `json:"domain,omitempty"`
	Templatelimit           string            `json:"templatelimit,omitempty"`
	Vmlimit                 string            `json:"vmlimit,omitempty"`
	Projecttotal            int               `json:"projecttotal,omitempty"`
	Primarystorageavailable string            `json:"primarystorageavailable,omitempty"`
	Accounttype             int               `json:"accounttype,omitempty"`
	Vmavailable             string            `json:"vmavailable,omitempty"`
	Cpuavailable            string            `json:"cpuavailable,omitempty"`
	Secondarystoragetotal   int               `json:"secondarystoragetotal,omitempty"`
	Iscleanuprequired       bool              `json:"iscleanuprequired,omitempty"`
	Networklimit            string            `json:"networklimit,omitempty"`
	Volumelimit             string            `json:"volumelimit,omitempty"`
	Primarystoragetotal     int               `json:"primarystoragetotal,omitempty"`
	Isdefault               bool              `json:"isdefault,omitempty"`
	User                    []struct {
		Account             string `json:"account,omitempty"`
		Lastname            string `json:"lastname,omitempty"`
		Timezone            string `json:"timezone,omitempty"`
		Domainid            string `json:"domainid,omitempty"`
		Isdefault           bool   `json:"isdefault,omitempty"`
		State               string `json:"state,omitempty"`
		Accounttype         int    `json:"accounttype,omitempty"`
		Id                  string `json:"id,omitempty"`
		Domain              string `json:"domain,omitempty"`
		Email               string `json:"email,omitempty"`
		Firstname           string `json:"firstname,omitempty"`
		Apikey              string `json:"apikey,omitempty"`
		Secretkey           string `json:"secretkey,omitempty"`
		Accountid           string `json:"accountid,omitempty"`
		Created             string `json:"created,omitempty"`
		Iscallerchilddomain bool   `json:"iscallerchilddomain,omitempty"`
		Username            string `json:"username,omitempty"`
	} `json:"user,omitempty"`
	Primarystoragelimit       string `json:"primarystoragelimit,omitempty"`
	Secondarystorageavailable string `json:"secondarystorageavailable,omitempty"`
	Networkavailable          string `json:"networkavailable,omitempty"`
	Vmtotal                   int    `json:"vmtotal,omitempty"`
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
	JobID                     string            `json:"jobid,omitempty"`
	Memoryavailable           string            `json:"memoryavailable,omitempty"`
	Memorylimit               string            `json:"memorylimit,omitempty"`
	Defaultzoneid             string            `json:"defaultzoneid,omitempty"`
	Cpuavailable              string            `json:"cpuavailable,omitempty"`
	Projectlimit              string            `json:"projectlimit,omitempty"`
	Vmstopped                 int               `json:"vmstopped,omitempty"`
	Domain                    string            `json:"domain,omitempty"`
	Receivedbytes             int               `json:"receivedbytes,omitempty"`
	Templatelimit             string            `json:"templatelimit,omitempty"`
	Primarystoragetotal       int               `json:"primarystoragetotal,omitempty"`
	Templatetotal             int               `json:"templatetotal,omitempty"`
	Projecttotal              int               `json:"projecttotal,omitempty"`
	Primarystoragelimit       string            `json:"primarystoragelimit,omitempty"`
	Snapshotavailable         string            `json:"snapshotavailable,omitempty"`
	Iptotal                   int               `json:"iptotal,omitempty"`
	Networklimit              string            `json:"networklimit,omitempty"`
	Domainid                  string            `json:"domainid,omitempty"`
	Networktotal              int               `json:"networktotal,omitempty"`
	Volumelimit               string            `json:"volumelimit,omitempty"`
	Accounttype               int               `json:"accounttype,omitempty"`
	Accountdetails            map[string]string `json:"accountdetails,omitempty"`
	Memorytotal               int               `json:"memorytotal,omitempty"`
	Templateavailable         string            `json:"templateavailable,omitempty"`
	Iplimit                   string            `json:"iplimit,omitempty"`
	Vpcavailable              string            `json:"vpcavailable,omitempty"`
	Iscleanuprequired         bool              `json:"iscleanuprequired,omitempty"`
	Primarystorageavailable   string            `json:"primarystorageavailable,omitempty"`
	Secondarystorageavailable string            `json:"secondarystorageavailable,omitempty"`
	Projectavailable          string            `json:"projectavailable,omitempty"`
	Volumetotal               int               `json:"volumetotal,omitempty"`
	Secondarystoragetotal     int               `json:"secondarystoragetotal,omitempty"`
	Ipavailable               string            `json:"ipavailable,omitempty"`
	Id                        string            `json:"id,omitempty"`
	Secondarystoragelimit     string            `json:"secondarystoragelimit,omitempty"`
	Vmavailable               string            `json:"vmavailable,omitempty"`
	State                     string            `json:"state,omitempty"`
	Name                      string            `json:"name,omitempty"`
	Vpctotal                  int               `json:"vpctotal,omitempty"`
	Cputotal                  int               `json:"cputotal,omitempty"`
	User                      []struct {
		Email               string `json:"email,omitempty"`
		Apikey              string `json:"apikey,omitempty"`
		Secretkey           string `json:"secretkey,omitempty"`
		Lastname            string `json:"lastname,omitempty"`
		Id                  string `json:"id,omitempty"`
		Domainid            string `json:"domainid,omitempty"`
		Iscallerchilddomain bool   `json:"iscallerchilddomain,omitempty"`
		Account             string `json:"account,omitempty"`
		Isdefault           bool   `json:"isdefault,omitempty"`
		Timezone            string `json:"timezone,omitempty"`
		Domain              string `json:"domain,omitempty"`
		Accountid           string `json:"accountid,omitempty"`
		Username            string `json:"username,omitempty"`
		Created             string `json:"created,omitempty"`
		State               string `json:"state,omitempty"`
		Firstname           string `json:"firstname,omitempty"`
		Accounttype         int    `json:"accounttype,omitempty"`
	} `json:"user,omitempty"`
	Vpclimit         string `json:"vpclimit,omitempty"`
	Vmtotal          int    `json:"vmtotal,omitempty"`
	Isdefault        bool   `json:"isdefault,omitempty"`
	Networkdomain    string `json:"networkdomain,omitempty"`
	Vmrunning        int    `json:"vmrunning,omitempty"`
	Snapshotlimit    string `json:"snapshotlimit,omitempty"`
	Cpulimit         string `json:"cpulimit,omitempty"`
	Snapshottotal    int    `json:"snapshottotal,omitempty"`
	Sentbytes        int    `json:"sentbytes,omitempty"`
	Networkavailable string `json:"networkavailable,omitempty"`
	Volumeavailable  string `json:"volumeavailable,omitempty"`
	Vmlimit          string `json:"vmlimit,omitempty"`
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
	Projectlimit              string `json:"projectlimit,omitempty"`
	Secondarystorageavailable string `json:"secondarystorageavailable,omitempty"`
	Domain                    string `json:"domain,omitempty"`
	Vmrunning                 int    `json:"vmrunning,omitempty"`
	Secondarystoragelimit     string `json:"secondarystoragelimit,omitempty"`
	Memoryavailable           string `json:"memoryavailable,omitempty"`
	Projectavailable          string `json:"projectavailable,omitempty"`
	Volumeavailable           string `json:"volumeavailable,omitempty"`
	Vpclimit                  string `json:"vpclimit,omitempty"`
	Accounttype               int    `json:"accounttype,omitempty"`
	Memorytotal               int    `json:"memorytotal,omitempty"`
	Primarystoragelimit       string `json:"primarystoragelimit,omitempty"`
	Networktotal              int    `json:"networktotal,omitempty"`
	Volumelimit               string `json:"volumelimit,omitempty"`
	Secondarystoragetotal     int    `json:"secondarystoragetotal,omitempty"`
	Templateavailable         string `json:"templateavailable,omitempty"`
	State                     string `json:"state,omitempty"`
	Vmlimit                   string `json:"vmlimit,omitempty"`
	Snapshotlimit             string `json:"snapshotlimit,omitempty"`
	Volumetotal               int    `json:"volumetotal,omitempty"`
	Networkavailable          string `json:"networkavailable,omitempty"`
	Vmstopped                 int    `json:"vmstopped,omitempty"`
	Projecttotal              int    `json:"projecttotal,omitempty"`
	Cpuavailable              string `json:"cpuavailable,omitempty"`
	Receivedbytes             int    `json:"receivedbytes,omitempty"`
	Vmtotal                   int    `json:"vmtotal,omitempty"`
	Vpctotal                  int    `json:"vpctotal,omitempty"`
	Domainid                  string `json:"domainid,omitempty"`
	Isdefault                 bool   `json:"isdefault,omitempty"`
	Iplimit                   string `json:"iplimit,omitempty"`
	User                      []struct {
		Firstname           string `json:"firstname,omitempty"`
		Email               string `json:"email,omitempty"`
		Account             string `json:"account,omitempty"`
		Timezone            string `json:"timezone,omitempty"`
		Username            string `json:"username,omitempty"`
		Domainid            string `json:"domainid,omitempty"`
		Accounttype         int    `json:"accounttype,omitempty"`
		Accountid           string `json:"accountid,omitempty"`
		Apikey              string `json:"apikey,omitempty"`
		Id                  string `json:"id,omitempty"`
		Created             string `json:"created,omitempty"`
		Domain              string `json:"domain,omitempty"`
		Secretkey           string `json:"secretkey,omitempty"`
		State               string `json:"state,omitempty"`
		Isdefault           bool   `json:"isdefault,omitempty"`
		Iscallerchilddomain bool   `json:"iscallerchilddomain,omitempty"`
		Lastname            string `json:"lastname,omitempty"`
	} `json:"user,omitempty"`
	Ipavailable             string            `json:"ipavailable,omitempty"`
	Vmavailable             string            `json:"vmavailable,omitempty"`
	Accountdetails          map[string]string `json:"accountdetails,omitempty"`
	Networklimit            string            `json:"networklimit,omitempty"`
	Cpulimit                string            `json:"cpulimit,omitempty"`
	Templatelimit           string            `json:"templatelimit,omitempty"`
	Name                    string            `json:"name,omitempty"`
	Snapshotavailable       string            `json:"snapshotavailable,omitempty"`
	Sentbytes               int               `json:"sentbytes,omitempty"`
	Primarystorageavailable string            `json:"primarystorageavailable,omitempty"`
	Iptotal                 int               `json:"iptotal,omitempty"`
	Vpcavailable            string            `json:"vpcavailable,omitempty"`
	Iscleanuprequired       bool              `json:"iscleanuprequired,omitempty"`
	Memorylimit             string            `json:"memorylimit,omitempty"`
	Primarystoragetotal     int               `json:"primarystoragetotal,omitempty"`
	Templatetotal           int               `json:"templatetotal,omitempty"`
	Snapshottotal           int               `json:"snapshottotal,omitempty"`
	Cputotal                int               `json:"cputotal,omitempty"`
	Defaultzoneid           string            `json:"defaultzoneid,omitempty"`
	Id                      string            `json:"id,omitempty"`
	Networkdomain           string            `json:"networkdomain,omitempty"`
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
	Memoryavailable         string `json:"memoryavailable,omitempty"`
	Vmtotal                 int    `json:"vmtotal,omitempty"`
	Receivedbytes           int    `json:"receivedbytes,omitempty"`
	Sentbytes               int    `json:"sentbytes,omitempty"`
	Memorytotal             int    `json:"memorytotal,omitempty"`
	Primarystorageavailable string `json:"primarystorageavailable,omitempty"`
	Iplimit                 string `json:"iplimit,omitempty"`
	Primarystoragelimit     string `json:"primarystoragelimit,omitempty"`
	Projectavailable        string `json:"projectavailable,omitempty"`
	Vpclimit                string `json:"vpclimit,omitempty"`
	Name                    string `json:"name,omitempty"`
	Iscleanuprequired       bool   `json:"iscleanuprequired,omitempty"`
	Projectlimit            string `json:"projectlimit,omitempty"`
	Cpulimit                string `json:"cpulimit,omitempty"`
	Projecttotal            int    `json:"projecttotal,omitempty"`
	Snapshottotal           int    `json:"snapshottotal,omitempty"`
	Domainid                string `json:"domainid,omitempty"`
	Snapshotavailable       string `json:"snapshotavailable,omitempty"`
	Memorylimit             string `json:"memorylimit,omitempty"`
	Templatetotal           int    `json:"templatetotal,omitempty"`
	Templatelimit           string `json:"templatelimit,omitempty"`
	Snapshotlimit           string `json:"snapshotlimit,omitempty"`
	User                    []struct {
		Accountid           string `json:"accountid,omitempty"`
		Isdefault           bool   `json:"isdefault,omitempty"`
		State               string `json:"state,omitempty"`
		Domain              string `json:"domain,omitempty"`
		Firstname           string `json:"firstname,omitempty"`
		Created             string `json:"created,omitempty"`
		Secretkey           string `json:"secretkey,omitempty"`
		Account             string `json:"account,omitempty"`
		Lastname            string `json:"lastname,omitempty"`
		Domainid            string `json:"domainid,omitempty"`
		Username            string `json:"username,omitempty"`
		Apikey              string `json:"apikey,omitempty"`
		Email               string `json:"email,omitempty"`
		Timezone            string `json:"timezone,omitempty"`
		Iscallerchilddomain bool   `json:"iscallerchilddomain,omitempty"`
		Accounttype         int    `json:"accounttype,omitempty"`
		Id                  string `json:"id,omitempty"`
	} `json:"user,omitempty"`
	Id                        string            `json:"id,omitempty"`
	Secondarystorageavailable string            `json:"secondarystorageavailable,omitempty"`
	Templateavailable         string            `json:"templateavailable,omitempty"`
	Secondarystoragetotal     int               `json:"secondarystoragetotal,omitempty"`
	Vmrunning                 int               `json:"vmrunning,omitempty"`
	Vmlimit                   string            `json:"vmlimit,omitempty"`
	Volumetotal               int               `json:"volumetotal,omitempty"`
	Ipavailable               string            `json:"ipavailable,omitempty"`
	Vmavailable               string            `json:"vmavailable,omitempty"`
	Vpctotal                  int               `json:"vpctotal,omitempty"`
	Vmstopped                 int               `json:"vmstopped,omitempty"`
	Networklimit              string            `json:"networklimit,omitempty"`
	Iptotal                   int               `json:"iptotal,omitempty"`
	Networkdomain             string            `json:"networkdomain,omitempty"`
	Volumeavailable           string            `json:"volumeavailable,omitempty"`
	Domain                    string            `json:"domain,omitempty"`
	Cputotal                  int               `json:"cputotal,omitempty"`
	Accounttype               int               `json:"accounttype,omitempty"`
	Volumelimit               string            `json:"volumelimit,omitempty"`
	Networkavailable          string            `json:"networkavailable,omitempty"`
	Networktotal              int               `json:"networktotal,omitempty"`
	Isdefault                 bool              `json:"isdefault,omitempty"`
	State                     string            `json:"state,omitempty"`
	Accountdetails            map[string]string `json:"accountdetails,omitempty"`
	Cpuavailable              string            `json:"cpuavailable,omitempty"`
	Vpcavailable              string            `json:"vpcavailable,omitempty"`
	Defaultzoneid             string            `json:"defaultzoneid,omitempty"`
	Secondarystoragelimit     string            `json:"secondarystoragelimit,omitempty"`
	Primarystoragetotal       int               `json:"primarystoragetotal,omitempty"`
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
	Cpuavailable            string `json:"cpuavailable,omitempty"`
	Iptotal                 int    `json:"iptotal,omitempty"`
	Name                    string `json:"name,omitempty"`
	Cputotal                int    `json:"cputotal,omitempty"`
	Iscleanuprequired       bool   `json:"iscleanuprequired,omitempty"`
	Projecttotal            int    `json:"projecttotal,omitempty"`
	Snapshottotal           int    `json:"snapshottotal,omitempty"`
	Receivedbytes           int    `json:"receivedbytes,omitempty"`
	Networkavailable        string `json:"networkavailable,omitempty"`
	Primarystorageavailable string `json:"primarystorageavailable,omitempty"`
	Volumelimit             string `json:"volumelimit,omitempty"`
	Primarystoragetotal     int    `json:"primarystoragetotal,omitempty"`
	Snapshotlimit           string `json:"snapshotlimit,omitempty"`
	Projectavailable        string `json:"projectavailable,omitempty"`
	Memorylimit             string `json:"memorylimit,omitempty"`
	Sentbytes               int    `json:"sentbytes,omitempty"`
	Domainid                string `json:"domainid,omitempty"`
	Volumetotal             int    `json:"volumetotal,omitempty"`
	Isdefault               bool   `json:"isdefault,omitempty"`
	Cpulimit                string `json:"cpulimit,omitempty"`
	Networkdomain           string `json:"networkdomain,omitempty"`
	User                    []struct {
		Timezone            string `json:"timezone,omitempty"`
		Apikey              string `json:"apikey,omitempty"`
		Accountid           string `json:"accountid,omitempty"`
		Username            string `json:"username,omitempty"`
		Domainid            string `json:"domainid,omitempty"`
		Firstname           string `json:"firstname,omitempty"`
		Email               string `json:"email,omitempty"`
		Iscallerchilddomain bool   `json:"iscallerchilddomain,omitempty"`
		Secretkey           string `json:"secretkey,omitempty"`
		State               string `json:"state,omitempty"`
		Id                  string `json:"id,omitempty"`
		Isdefault           bool   `json:"isdefault,omitempty"`
		Domain              string `json:"domain,omitempty"`
		Lastname            string `json:"lastname,omitempty"`
		Accounttype         int    `json:"accounttype,omitempty"`
		Created             string `json:"created,omitempty"`
		Account             string `json:"account,omitempty"`
	} `json:"user,omitempty"`
	Secondarystoragelimit     string            `json:"secondarystoragelimit,omitempty"`
	Networktotal              int               `json:"networktotal,omitempty"`
	Primarystoragelimit       string            `json:"primarystoragelimit,omitempty"`
	Vmrunning                 int               `json:"vmrunning,omitempty"`
	Templateavailable         string            `json:"templateavailable,omitempty"`
	State                     string            `json:"state,omitempty"`
	Vpcavailable              string            `json:"vpcavailable,omitempty"`
	Vmtotal                   int               `json:"vmtotal,omitempty"`
	Domain                    string            `json:"domain,omitempty"`
	Accounttype               int               `json:"accounttype,omitempty"`
	Iplimit                   string            `json:"iplimit,omitempty"`
	Networklimit              string            `json:"networklimit,omitempty"`
	Projectlimit              string            `json:"projectlimit,omitempty"`
	Vpclimit                  string            `json:"vpclimit,omitempty"`
	Secondarystoragetotal     int               `json:"secondarystoragetotal,omitempty"`
	Defaultzoneid             string            `json:"defaultzoneid,omitempty"`
	Templatetotal             int               `json:"templatetotal,omitempty"`
	Volumeavailable           string            `json:"volumeavailable,omitempty"`
	Templatelimit             string            `json:"templatelimit,omitempty"`
	Vpctotal                  int               `json:"vpctotal,omitempty"`
	Vmstopped                 int               `json:"vmstopped,omitempty"`
	Id                        string            `json:"id,omitempty"`
	Snapshotavailable         string            `json:"snapshotavailable,omitempty"`
	Secondarystorageavailable string            `json:"secondarystorageavailable,omitempty"`
	Memoryavailable           string            `json:"memoryavailable,omitempty"`
	Vmlimit                   string            `json:"vmlimit,omitempty"`
	Vmavailable               string            `json:"vmavailable,omitempty"`
	Memorytotal               int               `json:"memorytotal,omitempty"`
	Ipavailable               string            `json:"ipavailable,omitempty"`
	Accountdetails            map[string]string `json:"accountdetails,omitempty"`
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
	Secondarystoragetotal     int               `json:"secondarystoragetotal,omitempty"`
	Sentbytes                 int               `json:"sentbytes,omitempty"`
	Cpuavailable              string            `json:"cpuavailable,omitempty"`
	Defaultzoneid             string            `json:"defaultzoneid,omitempty"`
	Vmavailable               string            `json:"vmavailable,omitempty"`
	Vmstopped                 int               `json:"vmstopped,omitempty"`
	Secondarystorageavailable string            `json:"secondarystorageavailable,omitempty"`
	Volumeavailable           string            `json:"volumeavailable,omitempty"`
	Name                      string            `json:"name,omitempty"`
	Domain                    string            `json:"domain,omitempty"`
	Projecttotal              int               `json:"projecttotal,omitempty"`
	Networklimit              string            `json:"networklimit,omitempty"`
	Vpctotal                  int               `json:"vpctotal,omitempty"`
	Vmtotal                   int               `json:"vmtotal,omitempty"`
	Vpcavailable              string            `json:"vpcavailable,omitempty"`
	State                     string            `json:"state,omitempty"`
	Cpulimit                  string            `json:"cpulimit,omitempty"`
	Projectlimit              string            `json:"projectlimit,omitempty"`
	Volumetotal               int               `json:"volumetotal,omitempty"`
	Memorytotal               int               `json:"memorytotal,omitempty"`
	Templatelimit             string            `json:"templatelimit,omitempty"`
	Accounttype               int               `json:"accounttype,omitempty"`
	Vmlimit                   string            `json:"vmlimit,omitempty"`
	Ipavailable               string            `json:"ipavailable,omitempty"`
	Projectavailable          string            `json:"projectavailable,omitempty"`
	Isdefault                 bool              `json:"isdefault,omitempty"`
	Memorylimit               string            `json:"memorylimit,omitempty"`
	Domainid                  string            `json:"domainid,omitempty"`
	Primarystoragelimit       string            `json:"primarystoragelimit,omitempty"`
	Accountdetails            map[string]string `json:"accountdetails,omitempty"`
	Networkdomain             string            `json:"networkdomain,omitempty"`
	Secondarystoragelimit     string            `json:"secondarystoragelimit,omitempty"`
	Templateavailable         string            `json:"templateavailable,omitempty"`
	Networkavailable          string            `json:"networkavailable,omitempty"`
	Volumelimit               string            `json:"volumelimit,omitempty"`
	Iplimit                   string            `json:"iplimit,omitempty"`
	Id                        string            `json:"id,omitempty"`
	Snapshotavailable         string            `json:"snapshotavailable,omitempty"`
	User                      []struct {
		Created             string `json:"created,omitempty"`
		Timezone            string `json:"timezone,omitempty"`
		Iscallerchilddomain bool   `json:"iscallerchilddomain,omitempty"`
		Username            string `json:"username,omitempty"`
		Isdefault           bool   `json:"isdefault,omitempty"`
		Accountid           string `json:"accountid,omitempty"`
		Email               string `json:"email,omitempty"`
		Firstname           string `json:"firstname,omitempty"`
		Accounttype         int    `json:"accounttype,omitempty"`
		Secretkey           string `json:"secretkey,omitempty"`
		State               string `json:"state,omitempty"`
		Domain              string `json:"domain,omitempty"`
		Account             string `json:"account,omitempty"`
		Id                  string `json:"id,omitempty"`
		Lastname            string `json:"lastname,omitempty"`
		Domainid            string `json:"domainid,omitempty"`
		Apikey              string `json:"apikey,omitempty"`
	} `json:"user,omitempty"`
	Vmrunning               int    `json:"vmrunning,omitempty"`
	Snapshotlimit           string `json:"snapshotlimit,omitempty"`
	Primarystoragetotal     int    `json:"primarystoragetotal,omitempty"`
	Iscleanuprequired       bool   `json:"iscleanuprequired,omitempty"`
	Receivedbytes           int    `json:"receivedbytes,omitempty"`
	Vpclimit                string `json:"vpclimit,omitempty"`
	Memoryavailable         string `json:"memoryavailable,omitempty"`
	Templatetotal           int    `json:"templatetotal,omitempty"`
	Snapshottotal           int    `json:"snapshottotal,omitempty"`
	Primarystorageavailable string `json:"primarystorageavailable,omitempty"`
	Networktotal            int    `json:"networktotal,omitempty"`
	Iptotal                 int    `json:"iptotal,omitempty"`
	Cputotal                int    `json:"cputotal,omitempty"`
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
	Networklimit string `json:"networklimit,omitempty"`
	Id           string `json:"id,omitempty"`
	Volumelimit  string `json:"volumelimit,omitempty"`
	Vpclimit     string `json:"vpclimit,omitempty"`
	Vmlimit      string `json:"vmlimit,omitempty"`
	Ipavailable  string `json:"ipavailable,omitempty"`
	Name         string `json:"name,omitempty"`
	Displaytext  string `json:"displaytext,omitempty"`
	Tags         []struct {
		Domain       string `json:"domain,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Project      string `json:"project,omitempty"`
		Account      string `json:"account,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Value        string `json:"value,omitempty"`
		Key          string `json:"key,omitempty"`
	} `json:"tags,omitempty"`
	Volumeavailable           string `json:"volumeavailable,omitempty"`
	Domain                    string `json:"domain,omitempty"`
	Snapshotavailable         string `json:"snapshotavailable,omitempty"`
	Iplimit                   string `json:"iplimit,omitempty"`
	Primarystoragetotal       int    `json:"primarystoragetotal,omitempty"`
	Networktotal              int    `json:"networktotal,omitempty"`
	Iptotal                   int    `json:"iptotal,omitempty"`
	Secondarystoragelimit     string `json:"secondarystoragelimit,omitempty"`
	Vpctotal                  int    `json:"vpctotal,omitempty"`
	Volumetotal               int    `json:"volumetotal,omitempty"`
	Vmstopped                 int    `json:"vmstopped,omitempty"`
	Templateavailable         string `json:"templateavailable,omitempty"`
	Primarystorageavailable   string `json:"primarystorageavailable,omitempty"`
	Vmavailable               string `json:"vmavailable,omitempty"`
	Secondarystorageavailable string `json:"secondarystorageavailable,omitempty"`
	Cputotal                  int    `json:"cputotal,omitempty"`
	Primarystoragelimit       string `json:"primarystoragelimit,omitempty"`
	Snapshotlimit             string `json:"snapshotlimit,omitempty"`
	Secondarystoragetotal     int    `json:"secondarystoragetotal,omitempty"`
	Networkavailable          string `json:"networkavailable,omitempty"`
	Templatetotal             int    `json:"templatetotal,omitempty"`
	Cpuavailable              string `json:"cpuavailable,omitempty"`
	Cpulimit                  string `json:"cpulimit,omitempty"`
	Memorytotal               int    `json:"memorytotal,omitempty"`
	State                     string `json:"state,omitempty"`
	Templatelimit             string `json:"templatelimit,omitempty"`
	Memorylimit               string `json:"memorylimit,omitempty"`
	Vmtotal                   int    `json:"vmtotal,omitempty"`
	Account                   string `json:"account,omitempty"`
	Memoryavailable           string `json:"memoryavailable,omitempty"`
	Vmrunning                 int    `json:"vmrunning,omitempty"`
	Domainid                  string `json:"domainid,omitempty"`
	Snapshottotal             int    `json:"snapshottotal,omitempty"`
	Vpcavailable              string `json:"vpcavailable,omitempty"`
}
