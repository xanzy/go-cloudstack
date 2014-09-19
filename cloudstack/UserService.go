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

type CreateUserParams struct {
	p map[string]interface{}
}

func (p *CreateUserParams) toURLValues() url.Values {
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
	if v, found := p.p["email"]; found {
		u.Set("email", v.(string))
	}
	if v, found := p.p["firstname"]; found {
		u.Set("firstname", v.(string))
	}
	if v, found := p.p["lastname"]; found {
		u.Set("lastname", v.(string))
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

func (p *CreateUserParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *CreateUserParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *CreateUserParams) SetEmail(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["email"] = v
	return
}

func (p *CreateUserParams) SetFirstname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["firstname"] = v
	return
}

func (p *CreateUserParams) SetLastname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lastname"] = v
	return
}

func (p *CreateUserParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
	return
}

func (p *CreateUserParams) SetTimezone(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["timezone"] = v
	return
}

func (p *CreateUserParams) SetUserid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userid"] = v
	return
}

func (p *CreateUserParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
	return
}

// You should always use this function to get a new CreateUserParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewCreateUserParams(account string, email string, firstname string, lastname string, password string, username string) *CreateUserParams {
	p := &CreateUserParams{}
	p.p = make(map[string]interface{})
	p.p["account"] = account
	p.p["email"] = email
	p.p["firstname"] = firstname
	p.p["lastname"] = lastname
	p.p["password"] = password
	p.p["username"] = username
	return p
}

// Creates a user for an account that already exists
func (s *UserService) CreateUser(p *CreateUserParams) (*CreateUserResponse, error) {
	resp, err := s.cs.newRequest("createUser", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateUserResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type CreateUserResponse struct {
	Email               string `json:"email,omitempty"`
	Created             string `json:"created,omitempty"`
	Domainid            string `json:"domainid,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Accounttype         int    `json:"accounttype,omitempty"`
	Timezone            string `json:"timezone,omitempty"`
	Isdefault           bool   `json:"isdefault,omitempty"`
	Iscallerchilddomain bool   `json:"iscallerchilddomain,omitempty"`
	Secretkey           string `json:"secretkey,omitempty"`
	Firstname           string `json:"firstname,omitempty"`
	Username            string `json:"username,omitempty"`
	State               string `json:"state,omitempty"`
	Id                  string `json:"id,omitempty"`
	Lastname            string `json:"lastname,omitempty"`
	Account             string `json:"account,omitempty"`
	Apikey              string `json:"apikey,omitempty"`
	Accountid           string `json:"accountid,omitempty"`
}

type DeleteUserParams struct {
	p map[string]interface{}
}

func (p *DeleteUserParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteUserParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteUserParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewDeleteUserParams(id string) *DeleteUserParams {
	p := &DeleteUserParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a user for an account
func (s *UserService) DeleteUser(p *DeleteUserParams) (*DeleteUserResponse, error) {
	resp, err := s.cs.newRequest("deleteUser", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteUserResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type DeleteUserResponse struct {
	Displaytext string `json:"displaytext,omitempty"`
	Success     string `json:"success,omitempty"`
}

type UpdateUserParams struct {
	p map[string]interface{}
}

func (p *UpdateUserParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["email"]; found {
		u.Set("email", v.(string))
	}
	if v, found := p.p["firstname"]; found {
		u.Set("firstname", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["lastname"]; found {
		u.Set("lastname", v.(string))
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["timezone"]; found {
		u.Set("timezone", v.(string))
	}
	if v, found := p.p["userapikey"]; found {
		u.Set("userapikey", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	if v, found := p.p["usersecretkey"]; found {
		u.Set("usersecretkey", v.(string))
	}
	return u
}

func (p *UpdateUserParams) SetEmail(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["email"] = v
	return
}

func (p *UpdateUserParams) SetFirstname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["firstname"] = v
	return
}

func (p *UpdateUserParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *UpdateUserParams) SetLastname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lastname"] = v
	return
}

func (p *UpdateUserParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
	return
}

func (p *UpdateUserParams) SetTimezone(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["timezone"] = v
	return
}

func (p *UpdateUserParams) SetUserapikey(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userapikey"] = v
	return
}

func (p *UpdateUserParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
	return
}

func (p *UpdateUserParams) SetUsersecretkey(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["usersecretkey"] = v
	return
}

// You should always use this function to get a new UpdateUserParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewUpdateUserParams(id string) *UpdateUserParams {
	p := &UpdateUserParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a user account
func (s *UserService) UpdateUser(p *UpdateUserParams) (*UpdateUserResponse, error) {
	resp, err := s.cs.newRequest("updateUser", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateUserResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type UpdateUserResponse struct {
	State               string `json:"state,omitempty"`
	Timezone            string `json:"timezone,omitempty"`
	Created             string `json:"created,omitempty"`
	Domainid            string `json:"domainid,omitempty"`
	Firstname           string `json:"firstname,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Lastname            string `json:"lastname,omitempty"`
	Apikey              string `json:"apikey,omitempty"`
	Accountid           string `json:"accountid,omitempty"`
	Email               string `json:"email,omitempty"`
	Iscallerchilddomain bool   `json:"iscallerchilddomain,omitempty"`
	Account             string `json:"account,omitempty"`
	Id                  string `json:"id,omitempty"`
	Secretkey           string `json:"secretkey,omitempty"`
	Isdefault           bool   `json:"isdefault,omitempty"`
	Username            string `json:"username,omitempty"`
	Accounttype         int    `json:"accounttype,omitempty"`
}

type ListUsersParams struct {
	p map[string]interface{}
}

func (p *ListUsersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
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
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (p *ListUsersParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListUsersParams) SetAccounttype(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accounttype"] = v
	return
}

func (p *ListUsersParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListUsersParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListUsersParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
	return
}

func (p *ListUsersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListUsersParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
	return
}

func (p *ListUsersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListUsersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListUsersParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
	return
}

func (p *ListUsersParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
	return
}

// You should always use this function to get a new ListUsersParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewListUsersParams() *ListUsersParams {
	p := &ListUsersParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *UserService) GetUserByID(id string) (*User, int, error) {
	p := &ListUsersParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	l, err := s.ListUsers(p)
	if err != nil {
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.Users[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for User UUID: %s!", id)
}

// Lists user accounts
func (s *UserService) ListUsers(p *ListUsersParams) (*ListUsersResponse, error) {
	resp, err := s.cs.newRequest("listUsers", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListUsersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListUsersResponse struct {
	Count int     `json:"count"`
	Users []*User `json:"user"`
}

type User struct {
	Id                  string `json:"id,omitempty"`
	Isdefault           bool   `json:"isdefault,omitempty"`
	Accounttype         int    `json:"accounttype,omitempty"`
	Created             string `json:"created,omitempty"`
	Firstname           string `json:"firstname,omitempty"`
	State               string `json:"state,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Domainid            string `json:"domainid,omitempty"`
	Account             string `json:"account,omitempty"`
	Secretkey           string `json:"secretkey,omitempty"`
	Lastname            string `json:"lastname,omitempty"`
	Accountid           string `json:"accountid,omitempty"`
	Username            string `json:"username,omitempty"`
	Timezone            string `json:"timezone,omitempty"`
	Email               string `json:"email,omitempty"`
	Apikey              string `json:"apikey,omitempty"`
	Iscallerchilddomain bool   `json:"iscallerchilddomain,omitempty"`
}

type LockUserParams struct {
	p map[string]interface{}
}

func (p *LockUserParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *LockUserParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new LockUserParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewLockUserParams(id string) *LockUserParams {
	p := &LockUserParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Locks a user account
func (s *UserService) LockUser(p *LockUserParams) (*LockUserResponse, error) {
	resp, err := s.cs.newRequest("lockUser", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r LockUserResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type LockUserResponse struct {
	Firstname           string `json:"firstname,omitempty"`
	Username            string `json:"username,omitempty"`
	Lastname            string `json:"lastname,omitempty"`
	Apikey              string `json:"apikey,omitempty"`
	Isdefault           bool   `json:"isdefault,omitempty"`
	State               string `json:"state,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Email               string `json:"email,omitempty"`
	Domainid            string `json:"domainid,omitempty"`
	Created             string `json:"created,omitempty"`
	Secretkey           string `json:"secretkey,omitempty"`
	Timezone            string `json:"timezone,omitempty"`
	Id                  string `json:"id,omitempty"`
	Accounttype         int    `json:"accounttype,omitempty"`
	Iscallerchilddomain bool   `json:"iscallerchilddomain,omitempty"`
	Account             string `json:"account,omitempty"`
	Accountid           string `json:"accountid,omitempty"`
}

type DisableUserParams struct {
	p map[string]interface{}
}

func (p *DisableUserParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DisableUserParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DisableUserParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewDisableUserParams(id string) *DisableUserParams {
	p := &DisableUserParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Disables a user account
func (s *UserService) DisableUser(p *DisableUserParams) (*DisableUserResponse, error) {
	resp, err := s.cs.newRequest("disableUser", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DisableUserResponse
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

type DisableUserResponse struct {
	JobID               string `json:"jobid,omitempty"`
	Created             string `json:"created,omitempty"`
	Iscallerchilddomain bool   `json:"iscallerchilddomain,omitempty"`
	Domainid            string `json:"domainid,omitempty"`
	Timezone            string `json:"timezone,omitempty"`
	Account             string `json:"account,omitempty"`
	Accountid           string `json:"accountid,omitempty"`
	Username            string `json:"username,omitempty"`
	Email               string `json:"email,omitempty"`
	Apikey              string `json:"apikey,omitempty"`
	Accounttype         int    `json:"accounttype,omitempty"`
	Firstname           string `json:"firstname,omitempty"`
	Id                  string `json:"id,omitempty"`
	Lastname            string `json:"lastname,omitempty"`
	State               string `json:"state,omitempty"`
	Isdefault           bool   `json:"isdefault,omitempty"`
	Secretkey           string `json:"secretkey,omitempty"`
	Domain              string `json:"domain,omitempty"`
}

type EnableUserParams struct {
	p map[string]interface{}
}

func (p *EnableUserParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *EnableUserParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new EnableUserParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewEnableUserParams(id string) *EnableUserParams {
	p := &EnableUserParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Enables a user account
func (s *UserService) EnableUser(p *EnableUserParams) (*EnableUserResponse, error) {
	resp, err := s.cs.newRequest("enableUser", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r EnableUserResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type EnableUserResponse struct {
	Created             string `json:"created,omitempty"`
	Isdefault           bool   `json:"isdefault,omitempty"`
	Accountid           string `json:"accountid,omitempty"`
	Firstname           string `json:"firstname,omitempty"`
	Id                  string `json:"id,omitempty"`
	Apikey              string `json:"apikey,omitempty"`
	Username            string `json:"username,omitempty"`
	Email               string `json:"email,omitempty"`
	State               string `json:"state,omitempty"`
	Secretkey           string `json:"secretkey,omitempty"`
	Iscallerchilddomain bool   `json:"iscallerchilddomain,omitempty"`
	Domainid            string `json:"domainid,omitempty"`
	Account             string `json:"account,omitempty"`
	Timezone            string `json:"timezone,omitempty"`
	Lastname            string `json:"lastname,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Accounttype         int    `json:"accounttype,omitempty"`
}

type GetUserParams struct {
	p map[string]interface{}
}

func (p *GetUserParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["userapikey"]; found {
		u.Set("userapikey", v.(string))
	}
	return u
}

func (p *GetUserParams) SetUserapikey(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userapikey"] = v
	return
}

// You should always use this function to get a new GetUserParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewGetUserParams(userapikey string) *GetUserParams {
	p := &GetUserParams{}
	p.p = make(map[string]interface{})
	p.p["userapikey"] = userapikey
	return p
}

// Find user account by API key
func (s *UserService) GetUser(p *GetUserParams) (*GetUserResponse, error) {
	resp, err := s.cs.newRequest("getUser", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r GetUserResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type GetUserResponse struct {
	Domainid            string `json:"domainid,omitempty"`
	Username            string `json:"username,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Lastname            string `json:"lastname,omitempty"`
	Isdefault           bool   `json:"isdefault,omitempty"`
	State               string `json:"state,omitempty"`
	Created             string `json:"created,omitempty"`
	Account             string `json:"account,omitempty"`
	Accounttype         int    `json:"accounttype,omitempty"`
	Id                  string `json:"id,omitempty"`
	Email               string `json:"email,omitempty"`
	Timezone            string `json:"timezone,omitempty"`
	Apikey              string `json:"apikey,omitempty"`
	Firstname           string `json:"firstname,omitempty"`
	Iscallerchilddomain bool   `json:"iscallerchilddomain,omitempty"`
	Secretkey           string `json:"secretkey,omitempty"`
	Accountid           string `json:"accountid,omitempty"`
}

type RegisterUserKeysParams struct {
	p map[string]interface{}
}

func (p *RegisterUserKeysParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *RegisterUserKeysParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new RegisterUserKeysParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewRegisterUserKeysParams(id string) *RegisterUserKeysParams {
	p := &RegisterUserKeysParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// This command allows a user to register for the developer API, returning a secret key and an API key. This request is made through the integration API port, so it is a privileged command and must be made on behalf of a user. It is up to the implementer just how the username and password are entered, and then how that translates to an integration API request. Both secret key and API key should be returned to the user
func (s *UserService) RegisterUserKeys(p *RegisterUserKeysParams) (*RegisterUserKeysResponse, error) {
	resp, err := s.cs.newRequest("registerUserKeys", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RegisterUserKeysResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type RegisterUserKeysResponse struct {
	Apikey    string `json:"apikey,omitempty"`
	Secretkey string `json:"secretkey,omitempty"`
}

type ListLdapUsersParams struct {
	p map[string]interface{}
}

func (p *ListLdapUsersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["listtype"]; found {
		u.Set("listtype", v.(string))
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

func (p *ListLdapUsersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListLdapUsersParams) SetListtype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listtype"] = v
	return
}

func (p *ListLdapUsersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListLdapUsersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

// You should always use this function to get a new ListLdapUsersParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewListLdapUsersParams() *ListLdapUsersParams {
	p := &ListLdapUsersParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists all LDAP Users
func (s *UserService) ListLdapUsers(p *ListLdapUsersParams) (*ListLdapUsersResponse, error) {
	resp, err := s.cs.newRequest("listLdapUsers", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListLdapUsersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListLdapUsersResponse struct {
	Count     int         `json:"count"`
	LdapUsers []*LdapUser `json:"ldapuser"`
}

type LdapUser struct {
	Username  string `json:"username,omitempty"`
	Domain    string `json:"domain,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Email     string `json:"email,omitempty"`
	Principal string `json:"principal,omitempty"`
}

type ImportLdapUsersParams struct {
	p map[string]interface{}
}

func (p *ImportLdapUsersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["accountdetails"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("accountdetails[%d].key", i), k)
			u.Set(fmt.Sprintf("accountdetails[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["accounttype"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("accounttype", vv)
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["group"]; found {
		u.Set("group", v.(string))
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
	if v, found := p.p["timezone"]; found {
		u.Set("timezone", v.(string))
	}
	return u
}

func (p *ImportLdapUsersParams) SetAccountdetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accountdetails"] = v
	return
}

func (p *ImportLdapUsersParams) SetAccounttype(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accounttype"] = v
	return
}

func (p *ImportLdapUsersParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ImportLdapUsersParams) SetGroup(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["group"] = v
	return
}

func (p *ImportLdapUsersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ImportLdapUsersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ImportLdapUsersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ImportLdapUsersParams) SetTimezone(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["timezone"] = v
	return
}

// You should always use this function to get a new ImportLdapUsersParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewImportLdapUsersParams(accounttype int) *ImportLdapUsersParams {
	p := &ImportLdapUsersParams{}
	p.p = make(map[string]interface{})
	p.p["accounttype"] = accounttype
	return p
}

// Import LDAP users
func (s *UserService) ImportLdapUsers(p *ImportLdapUsersParams) (*ImportLdapUsersResponse, error) {
	resp, err := s.cs.newRequest("importLdapUsers", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ImportLdapUsersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ImportLdapUsersResponse struct {
	Email     string `json:"email,omitempty"`
	Domain    string `json:"domain,omitempty"`
	Username  string `json:"username,omitempty"`
	Principal string `json:"principal,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Firstname string `json:"firstname,omitempty"`
}
