//
// Copyright 2018, Sander van Harmelen
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
	"net/url"
	"strconv"
)

type AddImageStoreS3Params struct {
	p map[string]interface{}
}

func (p *AddImageStoreS3Params) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["accesskey"]; found {
		u.Set("accesskey", v.(string))
	}
	if v, found := p.p["bucket"]; found {
		u.Set("bucket", v.(string))
	}
	if v, found := p.p["connectiontimeout"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("connectiontimeout", vv)
	}
	if v, found := p.p["connectionttl"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("connectionttl", vv)
	}
	if v, found := p.p["endpoint"]; found {
		u.Set("endpoint", v.(string))
	}
	if v, found := p.p["maxerrorretry"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("maxerrorretry", vv)
	}
	if v, found := p.p["s3signer"]; found {
		u.Set("s3signer", v.(string))
	}
	if v, found := p.p["secretkey"]; found {
		u.Set("secretkey", v.(string))
	}
	if v, found := p.p["sockettimeout"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("sockettimeout", vv)
	}
	if v, found := p.p["usehttps"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("usehttps", vv)
	}
	if v, found := p.p["usetcpkeepalive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("usetcpkeepalive", vv)
	}
	return u
}

func (p *AddImageStoreS3Params) SetAccesskey(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accesskey"] = v
	return
}

func (p *AddImageStoreS3Params) SetBucket(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bucket"] = v
	return
}

func (p *AddImageStoreS3Params) SetConnectiontimeout(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["connectiontimeout"] = v
	return
}

func (p *AddImageStoreS3Params) SetConnectionttl(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["connectionttl"] = v
	return
}

func (p *AddImageStoreS3Params) SetEndpoint(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["endpoint"] = v
	return
}

func (p *AddImageStoreS3Params) SetMaxerrorretry(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["maxerrorretry"] = v
	return
}

func (p *AddImageStoreS3Params) SetS3signer(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["s3signer"] = v
	return
}

func (p *AddImageStoreS3Params) SetSecretkey(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["secretkey"] = v
	return
}

func (p *AddImageStoreS3Params) SetSockettimeout(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sockettimeout"] = v
	return
}

func (p *AddImageStoreS3Params) SetUsehttps(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["usehttps"] = v
	return
}

func (p *AddImageStoreS3Params) SetUsetcpkeepalive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["usetcpkeepalive"] = v
	return
}

// You should always use this function to get a new AddImageStoreS3Params instance,
// as then you are sure you have configured all required params
func (s *S3Service) NewAddImageStoreS3Params(accesskey string, bucket string, endpoint string, secretkey string) *AddImageStoreS3Params {
	p := &AddImageStoreS3Params{}
	p.p = make(map[string]interface{})
	p.p["accesskey"] = accesskey
	p.p["bucket"] = bucket
	p.p["endpoint"] = endpoint
	p.p["secretkey"] = secretkey
	return p
}

// Adds S3 Image Store
func (s *S3Service) AddImageStoreS3(p *AddImageStoreS3Params) (*AddImageStoreS3Response, error) {
	resp, err := s.cs.newRequest("addImageStoreS3", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddImageStoreS3Response
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type AddImageStoreS3Response struct {
	Details      []interface{} `json:"details"`
	Id           string        `json:"id"`
	Name         string        `json:"name"`
	Protocol     string        `json:"protocol"`
	Providername string        `json:"providername"`
	Scope        string        `json:"scope"`
	Url          string        `json:"url"`
	Zoneid       string        `json:"zoneid"`
	Zonename     string        `json:"zonename"`
}
