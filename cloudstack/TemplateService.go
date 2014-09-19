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
	"strings"
)

type CreateTemplateParams struct {
	p map[string]interface{}
}

func (p *CreateTemplateParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["bits"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("bits", vv)
	}
	if v, found := p.p["details"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("details[%d].key", i), k)
			u.Set(fmt.Sprintf("details[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["isdynamicallyscalable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isdynamicallyscalable", vv)
	}
	if v, found := p.p["isfeatured"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isfeatured", vv)
	}
	if v, found := p.p["ispublic"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("ispublic", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["ostypeid"]; found {
		u.Set("ostypeid", v.(string))
	}
	if v, found := p.p["passwordenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("passwordenabled", vv)
	}
	if v, found := p.p["requireshvm"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("requireshvm", vv)
	}
	if v, found := p.p["snapshotid"]; found {
		u.Set("snapshotid", v.(string))
	}
	if v, found := p.p["templatetag"]; found {
		u.Set("templatetag", v.(string))
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	if v, found := p.p["volumeid"]; found {
		u.Set("volumeid", v.(string))
	}
	return u
}

func (p *CreateTemplateParams) SetBits(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bits"] = v
	return
}

func (p *CreateTemplateParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
	return
}

func (p *CreateTemplateParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
	return
}

func (p *CreateTemplateParams) SetIsdynamicallyscalable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isdynamicallyscalable"] = v
	return
}

func (p *CreateTemplateParams) SetIsfeatured(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isfeatured"] = v
	return
}

func (p *CreateTemplateParams) SetIspublic(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ispublic"] = v
	return
}

func (p *CreateTemplateParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *CreateTemplateParams) SetOstypeid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ostypeid"] = v
	return
}

func (p *CreateTemplateParams) SetPasswordenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["passwordenabled"] = v
	return
}

func (p *CreateTemplateParams) SetRequireshvm(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["requireshvm"] = v
	return
}

func (p *CreateTemplateParams) SetSnapshotid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["snapshotid"] = v
	return
}

func (p *CreateTemplateParams) SetTemplatetag(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templatetag"] = v
	return
}

func (p *CreateTemplateParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
	return
}

func (p *CreateTemplateParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
	return
}

func (p *CreateTemplateParams) SetVolumeid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["volumeid"] = v
	return
}

// You should always use this function to get a new CreateTemplateParams instance,
// as then you are sure you have configured all required params
func (s *TemplateService) NewCreateTemplateParams(displaytext string, name string, ostypeid string) *CreateTemplateParams {
	p := &CreateTemplateParams{}
	p.p = make(map[string]interface{})
	p.p["displaytext"] = displaytext
	p.p["name"] = name
	p.p["ostypeid"] = ostypeid
	return p
}

// Creates a template of a virtual machine. The virtual machine must be in a STOPPED state. A template created from this command is automatically designated as a private template visible to the account that created it.
func (s *TemplateService) CreateTemplate(p *CreateTemplateParams) (*CreateTemplateResponse, error) {
	resp, err := s.cs.newRequest("createTemplate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateTemplateResponse
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

type CreateTemplateResponse struct {
	JobID                 string `json:"jobid,omitempty"`
	Created               string `json:"created,omitempty"`
	CrossZones            bool   `json:"crossZones,omitempty"`
	Name                  string `json:"name,omitempty"`
	Isdynamicallyscalable bool   `json:"isdynamicallyscalable,omitempty"`
	Id                    string `json:"id,omitempty"`
	Size                  int    `json:"size,omitempty"`
	Templatetype          string `json:"templatetype,omitempty"`
	Domain                string `json:"domain,omitempty"`
	Accountid             string `json:"accountid,omitempty"`
	Isextractable         bool   `json:"isextractable,omitempty"`
	Project               string `json:"project,omitempty"`
	Isready               bool   `json:"isready,omitempty"`
	Checksum              string `json:"checksum,omitempty"`
	Displaytext           string `json:"displaytext,omitempty"`
	Projectid             string `json:"projectid,omitempty"`
	Domainid              string `json:"domainid,omitempty"`
	Removed               string `json:"removed,omitempty"`
	Account               string `json:"account,omitempty"`
	Status                string `json:"status,omitempty"`
	Zonename              string `json:"zonename,omitempty"`
	Ostypeid              string `json:"ostypeid,omitempty"`
	Bootable              bool   `json:"bootable,omitempty"`
	Templatetag           string `json:"templatetag,omitempty"`
	Tags                  []struct {
		Resourceid   string `json:"resourceid,omitempty"`
		Account      string `json:"account,omitempty"`
		Value        string `json:"value,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Project      string `json:"project,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Key          string `json:"key,omitempty"`
		Customer     string `json:"customer,omitempty"`
	} `json:"tags,omitempty"`
	Ostypename       string            `json:"ostypename,omitempty"`
	Zoneid           string            `json:"zoneid,omitempty"`
	Format           string            `json:"format,omitempty"`
	Ispublic         bool              `json:"ispublic,omitempty"`
	Sshkeyenabled    bool              `json:"sshkeyenabled,omitempty"`
	Details          map[string]string `json:"details,omitempty"`
	Hostid           string            `json:"hostid,omitempty"`
	Hypervisor       string            `json:"hypervisor,omitempty"`
	Hostname         string            `json:"hostname,omitempty"`
	Sourcetemplateid string            `json:"sourcetemplateid,omitempty"`
	Passwordenabled  bool              `json:"passwordenabled,omitempty"`
	Isfeatured       bool              `json:"isfeatured,omitempty"`
}

type RegisterTemplateParams struct {
	p map[string]interface{}
}

func (p *RegisterTemplateParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["bits"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("bits", vv)
	}
	if v, found := p.p["checksum"]; found {
		u.Set("checksum", v.(string))
	}
	if v, found := p.p["details"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("details[%d].key", i), k)
			u.Set(fmt.Sprintf("details[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["format"]; found {
		u.Set("format", v.(string))
	}
	if v, found := p.p["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := p.p["isdynamicallyscalable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isdynamicallyscalable", vv)
	}
	if v, found := p.p["isextractable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isextractable", vv)
	}
	if v, found := p.p["isfeatured"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isfeatured", vv)
	}
	if v, found := p.p["ispublic"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("ispublic", vv)
	}
	if v, found := p.p["isrouting"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrouting", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["ostypeid"]; found {
		u.Set("ostypeid", v.(string))
	}
	if v, found := p.p["passwordenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("passwordenabled", vv)
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["requireshvm"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("requireshvm", vv)
	}
	if v, found := p.p["sshkeyenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("sshkeyenabled", vv)
	}
	if v, found := p.p["templatetag"]; found {
		u.Set("templatetag", v.(string))
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *RegisterTemplateParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *RegisterTemplateParams) SetBits(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bits"] = v
	return
}

func (p *RegisterTemplateParams) SetChecksum(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["checksum"] = v
	return
}

func (p *RegisterTemplateParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
	return
}

func (p *RegisterTemplateParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
	return
}

func (p *RegisterTemplateParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *RegisterTemplateParams) SetFormat(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["format"] = v
	return
}

func (p *RegisterTemplateParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
	return
}

func (p *RegisterTemplateParams) SetIsdynamicallyscalable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isdynamicallyscalable"] = v
	return
}

func (p *RegisterTemplateParams) SetIsextractable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isextractable"] = v
	return
}

func (p *RegisterTemplateParams) SetIsfeatured(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isfeatured"] = v
	return
}

func (p *RegisterTemplateParams) SetIspublic(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ispublic"] = v
	return
}

func (p *RegisterTemplateParams) SetIsrouting(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrouting"] = v
	return
}

func (p *RegisterTemplateParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *RegisterTemplateParams) SetOstypeid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ostypeid"] = v
	return
}

func (p *RegisterTemplateParams) SetPasswordenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["passwordenabled"] = v
	return
}

func (p *RegisterTemplateParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *RegisterTemplateParams) SetRequireshvm(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["requireshvm"] = v
	return
}

func (p *RegisterTemplateParams) SetSshkeyenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sshkeyenabled"] = v
	return
}

func (p *RegisterTemplateParams) SetTemplatetag(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templatetag"] = v
	return
}

func (p *RegisterTemplateParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
	return
}

func (p *RegisterTemplateParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new RegisterTemplateParams instance,
// as then you are sure you have configured all required params
func (s *TemplateService) NewRegisterTemplateParams(displaytext string, format string, hypervisor string, name string, ostypeid string, url string, zoneid string) *RegisterTemplateParams {
	p := &RegisterTemplateParams{}
	p.p = make(map[string]interface{})
	p.p["displaytext"] = displaytext
	p.p["format"] = format
	p.p["hypervisor"] = hypervisor
	p.p["name"] = name
	p.p["ostypeid"] = ostypeid
	p.p["url"] = url
	p.p["zoneid"] = zoneid
	return p
}

// Registers an existing template into the CloudStack cloud.
func (s *TemplateService) RegisterTemplate(p *RegisterTemplateParams) (*RegisterTemplateResponse, error) {
	resp, err := s.cs.newRequest("registerTemplate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RegisterTemplateResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type RegisterTemplateResponse struct {
	Ostypeid              string            `json:"ostypeid,omitempty"`
	Account               string            `json:"account,omitempty"`
	Domain                string            `json:"domain,omitempty"`
	CrossZones            bool              `json:"crossZones,omitempty"`
	Created               string            `json:"created,omitempty"`
	Status                string            `json:"status,omitempty"`
	Zoneid                string            `json:"zoneid,omitempty"`
	Isready               bool              `json:"isready,omitempty"`
	Accountid             string            `json:"accountid,omitempty"`
	Hostname              string            `json:"hostname,omitempty"`
	Passwordenabled       bool              `json:"passwordenabled,omitempty"`
	Ispublic              bool              `json:"ispublic,omitempty"`
	Isfeatured            bool              `json:"isfeatured,omitempty"`
	Sshkeyenabled         bool              `json:"sshkeyenabled,omitempty"`
	Details               map[string]string `json:"details,omitempty"`
	Projectid             string            `json:"projectid,omitempty"`
	Bootable              bool              `json:"bootable,omitempty"`
	Domainid              string            `json:"domainid,omitempty"`
	Isdynamicallyscalable bool              `json:"isdynamicallyscalable,omitempty"`
	Id                    string            `json:"id,omitempty"`
	Name                  string            `json:"name,omitempty"`
	Hypervisor            string            `json:"hypervisor,omitempty"`
	Templatetag           string            `json:"templatetag,omitempty"`
	Checksum              string            `json:"checksum,omitempty"`
	Templatetype          string            `json:"templatetype,omitempty"`
	Sourcetemplateid      string            `json:"sourcetemplateid,omitempty"`
	Format                string            `json:"format,omitempty"`
	Displaytext           string            `json:"displaytext,omitempty"`
	Hostid                string            `json:"hostid,omitempty"`
	Removed               string            `json:"removed,omitempty"`
	Size                  int               `json:"size,omitempty"`
	Ostypename            string            `json:"ostypename,omitempty"`
	Project               string            `json:"project,omitempty"`
	Tags                  []struct {
		Value        string `json:"value,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Account      string `json:"account,omitempty"`
		Key          string `json:"key,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Project      string `json:"project,omitempty"`
	} `json:"tags,omitempty"`
	Zonename      string `json:"zonename,omitempty"`
	Isextractable bool   `json:"isextractable,omitempty"`
}

type UpdateTemplateParams struct {
	p map[string]interface{}
}

func (p *UpdateTemplateParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["bootable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("bootable", vv)
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["format"]; found {
		u.Set("format", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["isdynamicallyscalable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isdynamicallyscalable", vv)
	}
	if v, found := p.p["isrouting"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrouting", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["ostypeid"]; found {
		u.Set("ostypeid", v.(string))
	}
	if v, found := p.p["passwordenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("passwordenabled", vv)
	}
	if v, found := p.p["sortkey"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("sortkey", vv)
	}
	return u
}

func (p *UpdateTemplateParams) SetBootable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bootable"] = v
	return
}

func (p *UpdateTemplateParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
	return
}

func (p *UpdateTemplateParams) SetFormat(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["format"] = v
	return
}

func (p *UpdateTemplateParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *UpdateTemplateParams) SetIsdynamicallyscalable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isdynamicallyscalable"] = v
	return
}

func (p *UpdateTemplateParams) SetIsrouting(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrouting"] = v
	return
}

func (p *UpdateTemplateParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *UpdateTemplateParams) SetOstypeid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ostypeid"] = v
	return
}

func (p *UpdateTemplateParams) SetPasswordenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["passwordenabled"] = v
	return
}

func (p *UpdateTemplateParams) SetSortkey(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sortkey"] = v
	return
}

// You should always use this function to get a new UpdateTemplateParams instance,
// as then you are sure you have configured all required params
func (s *TemplateService) NewUpdateTemplateParams(id string) *UpdateTemplateParams {
	p := &UpdateTemplateParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates attributes of a template.
func (s *TemplateService) UpdateTemplate(p *UpdateTemplateParams) (*UpdateTemplateResponse, error) {
	resp, err := s.cs.newRequest("updateTemplate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateTemplateResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type UpdateTemplateResponse struct {
	Displaytext string `json:"displaytext,omitempty"`
	Size        int    `json:"size,omitempty"`
	Created     string `json:"created,omitempty"`
	Ispublic    bool   `json:"ispublic,omitempty"`
	Tags        []struct {
		Resourceid   string `json:"resourceid,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Value        string `json:"value,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Project      string `json:"project,omitempty"`
		Account      string `json:"account,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Key          string `json:"key,omitempty"`
	} `json:"tags,omitempty"`
	Bootable              bool              `json:"bootable,omitempty"`
	Isextractable         bool              `json:"isextractable,omitempty"`
	Id                    string            `json:"id,omitempty"`
	Zoneid                string            `json:"zoneid,omitempty"`
	Removed               string            `json:"removed,omitempty"`
	Projectid             string            `json:"projectid,omitempty"`
	Sourcetemplateid      string            `json:"sourcetemplateid,omitempty"`
	Checksum              string            `json:"checksum,omitempty"`
	Isdynamicallyscalable bool              `json:"isdynamicallyscalable,omitempty"`
	Passwordenabled       bool              `json:"passwordenabled,omitempty"`
	Account               string            `json:"account,omitempty"`
	Hostname              string            `json:"hostname,omitempty"`
	Ostypename            string            `json:"ostypename,omitempty"`
	Hostid                string            `json:"hostid,omitempty"`
	Isready               bool              `json:"isready,omitempty"`
	Ostypeid              string            `json:"ostypeid,omitempty"`
	Status                string            `json:"status,omitempty"`
	Accountid             string            `json:"accountid,omitempty"`
	Templatetag           string            `json:"templatetag,omitempty"`
	Hypervisor            string            `json:"hypervisor,omitempty"`
	Domain                string            `json:"domain,omitempty"`
	Format                string            `json:"format,omitempty"`
	Project               string            `json:"project,omitempty"`
	Name                  string            `json:"name,omitempty"`
	Templatetype          string            `json:"templatetype,omitempty"`
	Isfeatured            bool              `json:"isfeatured,omitempty"`
	Zonename              string            `json:"zonename,omitempty"`
	Details               map[string]string `json:"details,omitempty"`
	CrossZones            bool              `json:"crossZones,omitempty"`
	Domainid              string            `json:"domainid,omitempty"`
	Sshkeyenabled         bool              `json:"sshkeyenabled,omitempty"`
}

type CopyTemplateParams struct {
	p map[string]interface{}
}

func (p *CopyTemplateParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["destzoneid"]; found {
		u.Set("destzoneid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["sourcezoneid"]; found {
		u.Set("sourcezoneid", v.(string))
	}
	return u
}

func (p *CopyTemplateParams) SetDestzoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["destzoneid"] = v
	return
}

func (p *CopyTemplateParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *CopyTemplateParams) SetSourcezoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sourcezoneid"] = v
	return
}

// You should always use this function to get a new CopyTemplateParams instance,
// as then you are sure you have configured all required params
func (s *TemplateService) NewCopyTemplateParams(destzoneid string, id string) *CopyTemplateParams {
	p := &CopyTemplateParams{}
	p.p = make(map[string]interface{})
	p.p["destzoneid"] = destzoneid
	p.p["id"] = id
	return p
}

// Copies a template from one zone to another.
func (s *TemplateService) CopyTemplate(p *CopyTemplateParams) (*CopyTemplateResponse, error) {
	resp, err := s.cs.newRequest("copyTemplate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CopyTemplateResponse
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

type CopyTemplateResponse struct {
	JobID                 string            `json:"jobid,omitempty"`
	Format                string            `json:"format,omitempty"`
	Displaytext           string            `json:"displaytext,omitempty"`
	Checksum              string            `json:"checksum,omitempty"`
	Zoneid                string            `json:"zoneid,omitempty"`
	Hostname              string            `json:"hostname,omitempty"`
	Ostypename            string            `json:"ostypename,omitempty"`
	Ostypeid              string            `json:"ostypeid,omitempty"`
	Size                  int               `json:"size,omitempty"`
	Domain                string            `json:"domain,omitempty"`
	Sourcetemplateid      string            `json:"sourcetemplateid,omitempty"`
	Bootable              bool              `json:"bootable,omitempty"`
	Name                  string            `json:"name,omitempty"`
	Accountid             string            `json:"accountid,omitempty"`
	Account               string            `json:"account,omitempty"`
	Details               map[string]string `json:"details,omitempty"`
	Passwordenabled       bool              `json:"passwordenabled,omitempty"`
	Zonename              string            `json:"zonename,omitempty"`
	Hypervisor            string            `json:"hypervisor,omitempty"`
	Id                    string            `json:"id,omitempty"`
	Isready               bool              `json:"isready,omitempty"`
	Removed               string            `json:"removed,omitempty"`
	Isextractable         bool              `json:"isextractable,omitempty"`
	Created               string            `json:"created,omitempty"`
	CrossZones            bool              `json:"crossZones,omitempty"`
	Templatetag           string            `json:"templatetag,omitempty"`
	Status                string            `json:"status,omitempty"`
	Isdynamicallyscalable bool              `json:"isdynamicallyscalable,omitempty"`
	Isfeatured            bool              `json:"isfeatured,omitempty"`
	Templatetype          string            `json:"templatetype,omitempty"`
	Sshkeyenabled         bool              `json:"sshkeyenabled,omitempty"`
	Hostid                string            `json:"hostid,omitempty"`
	Tags                  []struct {
		Key          string `json:"key,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Value        string `json:"value,omitempty"`
		Project      string `json:"project,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Account      string `json:"account,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
	} `json:"tags,omitempty"`
	Domainid  string `json:"domainid,omitempty"`
	Projectid string `json:"projectid,omitempty"`
	Ispublic  bool   `json:"ispublic,omitempty"`
	Project   string `json:"project,omitempty"`
}

type DeleteTemplateParams struct {
	p map[string]interface{}
}

func (p *DeleteTemplateParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *DeleteTemplateParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *DeleteTemplateParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new DeleteTemplateParams instance,
// as then you are sure you have configured all required params
func (s *TemplateService) NewDeleteTemplateParams(id string) *DeleteTemplateParams {
	p := &DeleteTemplateParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a template from the system. All virtual machines using the deleted template will not be affected.
func (s *TemplateService) DeleteTemplate(p *DeleteTemplateParams) (*DeleteTemplateResponse, error) {
	resp, err := s.cs.newRequest("deleteTemplate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteTemplateResponse
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

type DeleteTemplateResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Success     bool   `json:"success,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
}

type ListTemplatesParams struct {
	p map[string]interface{}
}

func (p *ListTemplatesParams) toURLValues() url.Values {
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
	if v, found := p.p["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
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
	if v, found := p.p["showremoved"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("showremoved", vv)
	}
	if v, found := p.p["tags"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["templatefilter"]; found {
		u.Set("templatefilter", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListTemplatesParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListTemplatesParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListTemplatesParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
	return
}

func (p *ListTemplatesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListTemplatesParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
	return
}

func (p *ListTemplatesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListTemplatesParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
	return
}

func (p *ListTemplatesParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *ListTemplatesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListTemplatesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListTemplatesParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *ListTemplatesParams) SetShowremoved(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["showremoved"] = v
	return
}

func (p *ListTemplatesParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
	return
}

func (p *ListTemplatesParams) SetTemplatefilter(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templatefilter"] = v
	return
}

func (p *ListTemplatesParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new ListTemplatesParams instance,
// as then you are sure you have configured all required params
func (s *TemplateService) NewListTemplatesParams(templatefilter string) *ListTemplatesParams {
	p := &ListTemplatesParams{}
	p.p = make(map[string]interface{})
	p.p["templatefilter"] = templatefilter
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *TemplateService) GetTemplateID(name string, templatefilter string) (string, error) {
	p := &ListTemplatesParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name
	p.p["templatefilter"] = templatefilter

	l, err := s.ListTemplates(p)
	if err != nil {
		return "", err
	}

	if l.Count == 0 {
		return "", fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Templates[0].Id, nil
	}

	if l.Count > 1 {
		for _, v := range l.Templates {
			if v.Name == name {
				return v.Id, nil
			}
		}
	}
	return "", fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *TemplateService) GetTemplateByName(name string, templatefilter string) (*Template, int, error) {
	id, err := s.GetTemplateID(name, templatefilter)
	if err != nil {
		return nil, -1, err
	}

	r, count, err := s.GetTemplateByID(id, templatefilter)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *TemplateService) GetTemplateByID(id string, templatefilter string) (*Template, int, error) {
	p := &ListTemplatesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id
	p.p["templatefilter"] = templatefilter

	l, err := s.ListTemplates(p)
	if err != nil {
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.Templates[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Template UUID: %s!", id)
}

// List all public, private, and privileged templates.
func (s *TemplateService) ListTemplates(p *ListTemplatesParams) (*ListTemplatesResponse, error) {
	resp, err := s.cs.newRequest("listTemplates", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListTemplatesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListTemplatesResponse struct {
	Count     int         `json:"count"`
	Templates []*Template `json:"template"`
}

type Template struct {
	Project               string            `json:"project,omitempty"`
	Name                  string            `json:"name,omitempty"`
	CrossZones            bool              `json:"crossZones,omitempty"`
	Id                    string            `json:"id,omitempty"`
	Templatetag           string            `json:"templatetag,omitempty"`
	Ostypeid              string            `json:"ostypeid,omitempty"`
	Account               string            `json:"account,omitempty"`
	Checksum              string            `json:"checksum,omitempty"`
	Format                string            `json:"format,omitempty"`
	Size                  int               `json:"size,omitempty"`
	Ispublic              bool              `json:"ispublic,omitempty"`
	Templatetype          string            `json:"templatetype,omitempty"`
	Isdynamicallyscalable bool              `json:"isdynamicallyscalable,omitempty"`
	Domain                string            `json:"domain,omitempty"`
	Ostypename            string            `json:"ostypename,omitempty"`
	Hypervisor            string            `json:"hypervisor,omitempty"`
	Domainid              string            `json:"domainid,omitempty"`
	Passwordenabled       bool              `json:"passwordenabled,omitempty"`
	Sourcetemplateid      string            `json:"sourcetemplateid,omitempty"`
	Removed               string            `json:"removed,omitempty"`
	Hostid                string            `json:"hostid,omitempty"`
	Hostname              string            `json:"hostname,omitempty"`
	Created               string            `json:"created,omitempty"`
	Accountid             string            `json:"accountid,omitempty"`
	Projectid             string            `json:"projectid,omitempty"`
	Isready               bool              `json:"isready,omitempty"`
	Details               map[string]string `json:"details,omitempty"`
	Tags                  []struct {
		Key          string `json:"key,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Project      string `json:"project,omitempty"`
		Account      string `json:"account,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Value        string `json:"value,omitempty"`
	} `json:"tags,omitempty"`
	Zoneid        string `json:"zoneid,omitempty"`
	Isfeatured    bool   `json:"isfeatured,omitempty"`
	Displaytext   string `json:"displaytext,omitempty"`
	Bootable      bool   `json:"bootable,omitempty"`
	Zonename      string `json:"zonename,omitempty"`
	Sshkeyenabled bool   `json:"sshkeyenabled,omitempty"`
	Isextractable bool   `json:"isextractable,omitempty"`
	Status        string `json:"status,omitempty"`
}

type UpdateTemplatePermissionsParams struct {
	p map[string]interface{}
}

func (p *UpdateTemplatePermissionsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["accounts"]; found {
		vv := strings.Join(v.([]string), ", ")
		u.Set("accounts", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["isextractable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isextractable", vv)
	}
	if v, found := p.p["isfeatured"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isfeatured", vv)
	}
	if v, found := p.p["ispublic"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("ispublic", vv)
	}
	if v, found := p.p["op"]; found {
		u.Set("op", v.(string))
	}
	if v, found := p.p["projectids"]; found {
		vv := strings.Join(v.([]string), ", ")
		u.Set("projectids", vv)
	}
	return u
}

func (p *UpdateTemplatePermissionsParams) SetAccounts(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accounts"] = v
	return
}

func (p *UpdateTemplatePermissionsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *UpdateTemplatePermissionsParams) SetIsextractable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isextractable"] = v
	return
}

func (p *UpdateTemplatePermissionsParams) SetIsfeatured(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isfeatured"] = v
	return
}

func (p *UpdateTemplatePermissionsParams) SetIspublic(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ispublic"] = v
	return
}

func (p *UpdateTemplatePermissionsParams) SetOp(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["op"] = v
	return
}

func (p *UpdateTemplatePermissionsParams) SetProjectids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectids"] = v
	return
}

// You should always use this function to get a new UpdateTemplatePermissionsParams instance,
// as then you are sure you have configured all required params
func (s *TemplateService) NewUpdateTemplatePermissionsParams(id string) *UpdateTemplatePermissionsParams {
	p := &UpdateTemplatePermissionsParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a template visibility permissions. A public template is visible to all accounts within the same domain. A private template is visible only to the owner of the template. A priviledged template is a private template with account permissions added. Only accounts specified under the template permissions are visible to them.
func (s *TemplateService) UpdateTemplatePermissions(p *UpdateTemplatePermissionsParams) (*UpdateTemplatePermissionsResponse, error) {
	resp, err := s.cs.newRequest("updateTemplatePermissions", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateTemplatePermissionsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type UpdateTemplatePermissionsResponse struct {
	Success     string `json:"success,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
}

type ListTemplatePermissionsParams struct {
	p map[string]interface{}
}

func (p *ListTemplatePermissionsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *ListTemplatePermissionsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new ListTemplatePermissionsParams instance,
// as then you are sure you have configured all required params
func (s *TemplateService) NewListTemplatePermissionsParams(id string) *ListTemplatePermissionsParams {
	p := &ListTemplatePermissionsParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *TemplateService) GetTemplatePermissionByID(id string) (*TemplatePermission, int, error) {
	p := &ListTemplatePermissionsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id
	p.p["id"] = id

	l, err := s.ListTemplatePermissions(p)
	if err != nil {
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.TemplatePermissions[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for TemplatePermission UUID: %s!", id)
}

// List template visibility and all accounts that have permissions to view this template.
func (s *TemplateService) ListTemplatePermissions(p *ListTemplatePermissionsParams) (*ListTemplatePermissionsResponse, error) {
	resp, err := s.cs.newRequest("listTemplatePermissions", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListTemplatePermissionsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListTemplatePermissionsResponse struct {
	Count               int                   `json:"count"`
	TemplatePermissions []*TemplatePermission `json:"templatepermission"`
}

type TemplatePermission struct {
	Account    []string `json:"account,omitempty"`
	Ispublic   bool     `json:"ispublic,omitempty"`
	Id         string   `json:"id,omitempty"`
	Projectids []string `json:"projectids,omitempty"`
	Domainid   string   `json:"domainid,omitempty"`
}

type ExtractTemplateParams struct {
	p map[string]interface{}
}

func (p *ExtractTemplateParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["mode"]; found {
		u.Set("mode", v.(string))
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ExtractTemplateParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ExtractTemplateParams) SetMode(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["mode"] = v
	return
}

func (p *ExtractTemplateParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
	return
}

func (p *ExtractTemplateParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new ExtractTemplateParams instance,
// as then you are sure you have configured all required params
func (s *TemplateService) NewExtractTemplateParams(id string, mode string) *ExtractTemplateParams {
	p := &ExtractTemplateParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["mode"] = mode
	return p
}

// Extracts a template
func (s *TemplateService) ExtractTemplate(p *ExtractTemplateParams) (*ExtractTemplateResponse, error) {
	resp, err := s.cs.newRequest("extractTemplate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ExtractTemplateResponse
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

type ExtractTemplateResponse struct {
	JobID            string `json:"jobid,omitempty"`
	Id               string `json:"id,omitempty"`
	ExtractId        string `json:"extractId,omitempty"`
	Status           string `json:"status,omitempty"`
	Url              string `json:"url,omitempty"`
	Resultstring     string `json:"resultstring,omitempty"`
	Name             string `json:"name,omitempty"`
	Zonename         string `json:"zonename,omitempty"`
	Accountid        string `json:"accountid,omitempty"`
	Created          string `json:"created,omitempty"`
	Zoneid           string `json:"zoneid,omitempty"`
	State            string `json:"state,omitempty"`
	Storagetype      string `json:"storagetype,omitempty"`
	ExtractMode      string `json:"extractMode,omitempty"`
	Uploadpercentage int    `json:"uploadpercentage,omitempty"`
}

type PrepareTemplateParams struct {
	p map[string]interface{}
}

func (p *PrepareTemplateParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["templateid"]; found {
		u.Set("templateid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *PrepareTemplateParams) SetTemplateid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templateid"] = v
	return
}

func (p *PrepareTemplateParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new PrepareTemplateParams instance,
// as then you are sure you have configured all required params
func (s *TemplateService) NewPrepareTemplateParams(templateid string, zoneid string) *PrepareTemplateParams {
	p := &PrepareTemplateParams{}
	p.p = make(map[string]interface{})
	p.p["templateid"] = templateid
	p.p["zoneid"] = zoneid
	return p
}

// load template into primary storage
func (s *TemplateService) PrepareTemplate(p *PrepareTemplateParams) (*PrepareTemplateResponse, error) {
	resp, err := s.cs.newRequest("prepareTemplate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r PrepareTemplateResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type PrepareTemplateResponse struct {
	Ostypeid        string            `json:"ostypeid,omitempty"`
	Zoneid          string            `json:"zoneid,omitempty"`
	Details         map[string]string `json:"details,omitempty"`
	Hypervisor      string            `json:"hypervisor,omitempty"`
	Project         string            `json:"project,omitempty"`
	Checksum        string            `json:"checksum,omitempty"`
	Passwordenabled bool              `json:"passwordenabled,omitempty"`
	Size            int               `json:"size,omitempty"`
	Displaytext     string            `json:"displaytext,omitempty"`
	Account         string            `json:"account,omitempty"`
	Status          string            `json:"status,omitempty"`
	Created         string            `json:"created,omitempty"`
	Templatetag     string            `json:"templatetag,omitempty"`
	Bootable        bool              `json:"bootable,omitempty"`
	Templatetype    string            `json:"templatetype,omitempty"`
	Name            string            `json:"name,omitempty"`
	Accountid       string            `json:"accountid,omitempty"`
	Zonename        string            `json:"zonename,omitempty"`
	Isextractable   bool              `json:"isextractable,omitempty"`
	Domain          string            `json:"domain,omitempty"`
	Tags            []struct {
		Resourcetype string `json:"resourcetype,omitempty"`
		Account      string `json:"account,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Key          string `json:"key,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Project      string `json:"project,omitempty"`
		Value        string `json:"value,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
	} `json:"tags,omitempty"`
	Ostypename            string `json:"ostypename,omitempty"`
	Format                string `json:"format,omitempty"`
	Isfeatured            bool   `json:"isfeatured,omitempty"`
	Id                    string `json:"id,omitempty"`
	Domainid              string `json:"domainid,omitempty"`
	CrossZones            bool   `json:"crossZones,omitempty"`
	Removed               string `json:"removed,omitempty"`
	Ispublic              bool   `json:"ispublic,omitempty"`
	Hostid                string `json:"hostid,omitempty"`
	Sshkeyenabled         bool   `json:"sshkeyenabled,omitempty"`
	Hostname              string `json:"hostname,omitempty"`
	Sourcetemplateid      string `json:"sourcetemplateid,omitempty"`
	Isdynamicallyscalable bool   `json:"isdynamicallyscalable,omitempty"`
	Isready               bool   `json:"isready,omitempty"`
	Projectid             string `json:"projectid,omitempty"`
}

type UpgradeRouterTemplateParams struct {
	p map[string]interface{}
}

func (p *UpgradeRouterTemplateParams) toURLValues() url.Values {
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
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *UpgradeRouterTemplateParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *UpgradeRouterTemplateParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
	return
}

func (p *UpgradeRouterTemplateParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *UpgradeRouterTemplateParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *UpgradeRouterTemplateParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
	return
}

func (p *UpgradeRouterTemplateParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new UpgradeRouterTemplateParams instance,
// as then you are sure you have configured all required params
func (s *TemplateService) NewUpgradeRouterTemplateParams() *UpgradeRouterTemplateParams {
	p := &UpgradeRouterTemplateParams{}
	p.p = make(map[string]interface{})
	return p
}

// Upgrades router to use newer template
func (s *TemplateService) UpgradeRouterTemplate(p *UpgradeRouterTemplateParams) (*UpgradeRouterTemplateResponse, error) {
	resp, err := s.cs.newRequest("upgradeRouterTemplate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpgradeRouterTemplateResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type UpgradeRouterTemplateResponse struct {
	Jobid     string `json:"jobid,omitempty"`
	Jobstatus int    `json:"jobstatus,omitempty"`
}

type ListUcsTemplatesParams struct {
	p map[string]interface{}
}

func (p *ListUcsTemplatesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
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
	if v, found := p.p["ucsmanagerid"]; found {
		u.Set("ucsmanagerid", v.(string))
	}
	return u
}

func (p *ListUcsTemplatesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListUcsTemplatesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListUcsTemplatesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListUcsTemplatesParams) SetUcsmanagerid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ucsmanagerid"] = v
	return
}

// You should always use this function to get a new ListUcsTemplatesParams instance,
// as then you are sure you have configured all required params
func (s *TemplateService) NewListUcsTemplatesParams(ucsmanagerid string) *ListUcsTemplatesParams {
	p := &ListUcsTemplatesParams{}
	p.p = make(map[string]interface{})
	p.p["ucsmanagerid"] = ucsmanagerid
	return p
}

// List templates in ucs manager
func (s *TemplateService) ListUcsTemplates(p *ListUcsTemplatesParams) (*ListUcsTemplatesResponse, error) {
	resp, err := s.cs.newRequest("listUcsTemplates", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListUcsTemplatesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListUcsTemplatesResponse struct {
	Count        int            `json:"count"`
	UcsTemplates []*UcsTemplate `json:"ucstemplate"`
}

type UcsTemplate struct {
	Ucsdn string `json:"ucsdn,omitempty"`
}

type InstantiateUcsTemplateAndAssocaciateToBladeParams struct {
	p map[string]interface{}
}

func (p *InstantiateUcsTemplateAndAssocaciateToBladeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["bladeid"]; found {
		u.Set("bladeid", v.(string))
	}
	if v, found := p.p["profilename"]; found {
		u.Set("profilename", v.(string))
	}
	if v, found := p.p["templatedn"]; found {
		u.Set("templatedn", v.(string))
	}
	if v, found := p.p["ucsmanagerid"]; found {
		u.Set("ucsmanagerid", v.(string))
	}
	return u
}

func (p *InstantiateUcsTemplateAndAssocaciateToBladeParams) SetBladeid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bladeid"] = v
	return
}

func (p *InstantiateUcsTemplateAndAssocaciateToBladeParams) SetProfilename(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["profilename"] = v
	return
}

func (p *InstantiateUcsTemplateAndAssocaciateToBladeParams) SetTemplatedn(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templatedn"] = v
	return
}

func (p *InstantiateUcsTemplateAndAssocaciateToBladeParams) SetUcsmanagerid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ucsmanagerid"] = v
	return
}

// You should always use this function to get a new InstantiateUcsTemplateAndAssocaciateToBladeParams instance,
// as then you are sure you have configured all required params
func (s *TemplateService) NewInstantiateUcsTemplateAndAssocaciateToBladeParams(bladeid string, templatedn string, ucsmanagerid string) *InstantiateUcsTemplateAndAssocaciateToBladeParams {
	p := &InstantiateUcsTemplateAndAssocaciateToBladeParams{}
	p.p = make(map[string]interface{})
	p.p["bladeid"] = bladeid
	p.p["templatedn"] = templatedn
	p.p["ucsmanagerid"] = ucsmanagerid
	return p
}

// create a profile of template and associate to a blade
func (s *TemplateService) InstantiateUcsTemplateAndAssocaciateToBlade(p *InstantiateUcsTemplateAndAssocaciateToBladeParams) (*InstantiateUcsTemplateAndAssocaciateToBladeResponse, error) {
	resp, err := s.cs.newRequest("instantiateUcsTemplateAndAssocaciateToBlade", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r InstantiateUcsTemplateAndAssocaciateToBladeResponse
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

type InstantiateUcsTemplateAndAssocaciateToBladeResponse struct {
	JobID        string `json:"jobid,omitempty"`
	Bladedn      string `json:"bladedn,omitempty"`
	Hostid       string `json:"hostid,omitempty"`
	Profiledn    string `json:"profiledn,omitempty"`
	Id           string `json:"id,omitempty"`
	Ucsmanagerid string `json:"ucsmanagerid,omitempty"`
}
