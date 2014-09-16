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
	"crypto/hmac"
	"crypto/sha1"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type CSError struct {
	ErrorCode   int    `json:"errorcode"`
	CSErrorCode int    `json:"cserrorcode"`
	ErrorText   string `json:"errortext"`
}

func (e *CSError) Error() error {
	return fmt.Errorf("CloudStack API error %d (CSExceptionErrorCode: %d): %s", e.ErrorCode, e.CSErrorCode, e.ErrorText)
}

type CloudStackClient struct {
	client  *http.Client // The http client for communicating
	baseURL string       // The base URL of the API
	apiKey  string       // Api key
	secret  string       // Secret key
	async   bool         // Wait for async calls to finish
	timeout int64        // Max waiting timeout in seconds for async jobs to finish; defaults to 60 seconds

	NetworkACL       *NetworkACLService
	LDAP             *LDAPService
	Firewall         *FirewallService
	Account          *AccountService
	SecurityGroup    *SecurityGroupService
	NetworkDevice    *NetworkDeviceService
	Project          *ProjectService
	GuestOS          *GuestOSService
	UCS              *UCSService
	SSH              *SSHService
	Limit            *LimitService
	ISO              *ISOService
	ImageStore       *ImageStoreService
	InternalLB       *InternalLBService
	S3               *S3Service
	Certificate      *CertificateService
	Cluster          *ClusterService
	Pool             *PoolService
	Address          *AddressService
	VPN              *VPNService
	Router           *RouterService
	Region           *RegionService
	DiskOffering     *DiskOfferingService
	Resourcetags     *ResourcetagsService
	APIDiscovery     *APIDiscoveryService
	VLAN             *VLANService
	Resourcemetadata *ResourcemetadataService
	SystemCapacity   *SystemCapacityService
	ServiceOffering  *ServiceOfferingService
	NetworkOffering  *NetworkOfferingService
	OvsElement       *OvsElementService
	Logout           *LogoutService
	CloudIdentifier  *CloudIdentifierService
	AutoScale        *AutoScaleService
	Template         *TemplateService
	Pod              *PodService
	StoragePool      *StoragePoolService
	Nic              *NicService
	Swift            *SwiftService
	LoadBalancer     *LoadBalancerService
	Host             *HostService
	Snapshot         *SnapshotService
	Event            *EventService
	Usage            *UsageService
	StratosphereSSP  *StratosphereSSPService
	NAT              *NATService
	Alert            *AlertService
	AffinityGroup    *AffinityGroupService
	VMGroup          *VMGroupService
	VPC              *VPCService
	User             *UserService
	SystemVM         *SystemVMService
	Zone             *ZoneService
	Domain           *DomainService
	Baremetal        *BaremetalService
	PortableIP       *PortableIPService
	Volume           *VolumeService
	NiciraNVP        *NiciraNVPService
	Hypervisor       *HypervisorService
	BigSwitchVNS     *BigSwitchVNSService
	Network          *NetworkService
	VirtualMachine   *VirtualMachineService
	Configuration    *ConfigurationService
	Asyncjob         *AsyncjobService
	Login            *LoginService
}

// Creates a new client for communicating with CloudStack
func newClient(apiurl string, apikey string, secret string, async bool, verifyssl bool) *CloudStackClient {
	cs := &CloudStackClient{
		client: &http.Client{
			Transport: &http.Transport{
				Proxy:           http.ProxyFromEnvironment,
				TLSClientConfig: &tls.Config{InsecureSkipVerify: !verifyssl}, // If verifyssl is true, skipping the verify should be false and vice versa
			},
		},
		baseURL: apiurl,
		apiKey:  apikey,
		secret:  secret,
		async:   async,
		timeout: 60,
	}
	cs.NetworkACL = NewNetworkACLService(cs)
	cs.LDAP = NewLDAPService(cs)
	cs.Firewall = NewFirewallService(cs)
	cs.Account = NewAccountService(cs)
	cs.SecurityGroup = NewSecurityGroupService(cs)
	cs.NetworkDevice = NewNetworkDeviceService(cs)
	cs.Project = NewProjectService(cs)
	cs.GuestOS = NewGuestOSService(cs)
	cs.UCS = NewUCSService(cs)
	cs.SSH = NewSSHService(cs)
	cs.Limit = NewLimitService(cs)
	cs.ISO = NewISOService(cs)
	cs.ImageStore = NewImageStoreService(cs)
	cs.InternalLB = NewInternalLBService(cs)
	cs.S3 = NewS3Service(cs)
	cs.Certificate = NewCertificateService(cs)
	cs.Cluster = NewClusterService(cs)
	cs.Pool = NewPoolService(cs)
	cs.Address = NewAddressService(cs)
	cs.VPN = NewVPNService(cs)
	cs.Router = NewRouterService(cs)
	cs.Region = NewRegionService(cs)
	cs.DiskOffering = NewDiskOfferingService(cs)
	cs.Resourcetags = NewResourcetagsService(cs)
	cs.APIDiscovery = NewAPIDiscoveryService(cs)
	cs.VLAN = NewVLANService(cs)
	cs.Resourcemetadata = NewResourcemetadataService(cs)
	cs.SystemCapacity = NewSystemCapacityService(cs)
	cs.ServiceOffering = NewServiceOfferingService(cs)
	cs.NetworkOffering = NewNetworkOfferingService(cs)
	cs.OvsElement = NewOvsElementService(cs)
	cs.Logout = NewLogoutService(cs)
	cs.CloudIdentifier = NewCloudIdentifierService(cs)
	cs.AutoScale = NewAutoScaleService(cs)
	cs.Template = NewTemplateService(cs)
	cs.Pod = NewPodService(cs)
	cs.StoragePool = NewStoragePoolService(cs)
	cs.Nic = NewNicService(cs)
	cs.Swift = NewSwiftService(cs)
	cs.LoadBalancer = NewLoadBalancerService(cs)
	cs.Host = NewHostService(cs)
	cs.Snapshot = NewSnapshotService(cs)
	cs.Event = NewEventService(cs)
	cs.Usage = NewUsageService(cs)
	cs.StratosphereSSP = NewStratosphereSSPService(cs)
	cs.NAT = NewNATService(cs)
	cs.Alert = NewAlertService(cs)
	cs.AffinityGroup = NewAffinityGroupService(cs)
	cs.VMGroup = NewVMGroupService(cs)
	cs.VPC = NewVPCService(cs)
	cs.User = NewUserService(cs)
	cs.SystemVM = NewSystemVMService(cs)
	cs.Zone = NewZoneService(cs)
	cs.Domain = NewDomainService(cs)
	cs.Baremetal = NewBaremetalService(cs)
	cs.PortableIP = NewPortableIPService(cs)
	cs.Volume = NewVolumeService(cs)
	cs.NiciraNVP = NewNiciraNVPService(cs)
	cs.Hypervisor = NewHypervisorService(cs)
	cs.BigSwitchVNS = NewBigSwitchVNSService(cs)
	cs.Network = NewNetworkService(cs)
	cs.VirtualMachine = NewVirtualMachineService(cs)
	cs.Configuration = NewConfigurationService(cs)
	cs.Asyncjob = NewAsyncjobService(cs)
	cs.Login = NewLoginService(cs)
	return cs
}

// Default non-async client. So for async calls you need to implement and check the async job result yourself. When using
// HTTPS with a self-signed certificate to connect to your CloudStack API, you would probably want to set 'verifyssl' to
// false so the call ignores the SSL errors/warnings.
func NewClient(apiurl string, apikey string, secret string, verifyssl bool) *CloudStackClient {
	cs := newClient(apiurl, apikey, secret, false, verifyssl)
	return cs
}

// For sync API calls this client behaves exactly the same as a standard client call, but for async API calls
// this client will wait until the async job is finished or until the configured AsyncTimeout is reached. When the async
// job finishes successfully it will return actual object received from the API and nil, but when the timout is
// reached it will return the initial object containing the async job ID for the running job and a warning.
func NewAsyncClient(apiurl string, apikey string, secret string, verifyssl bool) *CloudStackClient {
	cs := newClient(apiurl, apikey, secret, true, verifyssl)
	return cs
}

// When using the async client an api call will wait for the async call to finish before returning. The default is to poll for 60
// seconds, to check if the async job is finished.
func (cs *CloudStackClient) AsyncTimeout(timeoutInSeconds int64) {
	cs.timeout = timeoutInSeconds
}

// A helper function that you can use to get the result of a running async job. If the job is not finished within the configured
// timeout, the async job return a warning saying the timer has expired.
func (cs *CloudStackClient) GetAsyncJobResult(jobid string, timeout int64) (b json.RawMessage, warn error, err error) {
	currentTime := time.Now().Unix()
	for {
		p := cs.Asyncjob.NewQueryAsyncJobResultParams(jobid)
		r, err := cs.Asyncjob.QueryAsyncJobResult(p)
		if err != nil {
			return nil, nil, err
		}

		// Status 1 means the job is finished successfully
		if r.Jobstatus == 1 {
			b, err := getRawValue(r.Jobresult)
			if err != nil {
				return nil, nil, err
			}
			return b, nil, nil
		}

		// When the status is 2, the job has failed
		if r.Jobstatus == 2 {
			if r.Jobresulttype == "text" {
				return nil, nil, fmt.Errorf(string(r.Jobresult))
			} else {
				return nil, nil, fmt.Errorf("Undefined error: %s", string(r.Jobresult))
			}
		}

		if time.Now().Unix()-currentTime > timeout {
			return nil, fmt.Errorf("Timeout while waiting for async job to finish"), nil
		}
		time.Sleep(3 * time.Second)
	}
}

// Execute the request against a CS API. Will return the raw JSON data returned by the API and nil if
// no error occured. If the API returns an error the result will be nil and the HTTP error code and CS
// error details. If a processing (code) error occurs the result will be nil and the generated error
func (cs *CloudStackClient) newRequest(api string, params url.Values) (json.RawMessage, error) {
	params.Set("apikey", cs.apiKey)
	params.Set("command", api)
	params.Set("response", "json")

	// Generate signature for API call
	// * Serialize parameters and sort them by key, done by Encode
	// * Convert the entire argument string to lowercase
	// * Replace all instances of '+' to '%20'
	// * Calculate HMAC SHA1 of argument string with CloudStack secret
	// * URL encode the string and convert to base64
	s := params.Encode()
	s2 := strings.ToLower(s)
	s3 := strings.Replace(s2, "+", "%20", -1)
	mac := hmac.New(sha1.New, []byte(cs.secret))
	mac.Write([]byte(s3))
	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	signature = url.QueryEscape(signature)

	// Create the final URL before we issue the request
	url := cs.baseURL + "?" + s + "&signature=" + signature

	resp, err := cs.client.Get(url)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	// Need to get the raw value to make the result play nice
	b, err = getRawValue(b)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		var e CSError
		if err := json.Unmarshal(b, &e); err != nil {
			return nil, err
		}
		return nil, e.Error()
	}
	return b, nil
}

// Generic function to get the first raw value from a response as json.RawMessage
func getRawValue(b json.RawMessage) (json.RawMessage, error) {
	var m map[string]json.RawMessage
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	for _, v := range m {
		return v, nil
	}
	return nil, fmt.Errorf("Unable to extract the raw value from:\n\n%s\n\n", string(b))
}

type NetworkACLService struct {
	cs *CloudStackClient
}

func NewNetworkACLService(cs *CloudStackClient) *NetworkACLService {
	return &NetworkACLService{cs: cs}
}

type LDAPService struct {
	cs *CloudStackClient
}

func NewLDAPService(cs *CloudStackClient) *LDAPService {
	return &LDAPService{cs: cs}
}

type FirewallService struct {
	cs *CloudStackClient
}

func NewFirewallService(cs *CloudStackClient) *FirewallService {
	return &FirewallService{cs: cs}
}

type AccountService struct {
	cs *CloudStackClient
}

func NewAccountService(cs *CloudStackClient) *AccountService {
	return &AccountService{cs: cs}
}

type SecurityGroupService struct {
	cs *CloudStackClient
}

func NewSecurityGroupService(cs *CloudStackClient) *SecurityGroupService {
	return &SecurityGroupService{cs: cs}
}

type NetworkDeviceService struct {
	cs *CloudStackClient
}

func NewNetworkDeviceService(cs *CloudStackClient) *NetworkDeviceService {
	return &NetworkDeviceService{cs: cs}
}

type ProjectService struct {
	cs *CloudStackClient
}

func NewProjectService(cs *CloudStackClient) *ProjectService {
	return &ProjectService{cs: cs}
}

type GuestOSService struct {
	cs *CloudStackClient
}

func NewGuestOSService(cs *CloudStackClient) *GuestOSService {
	return &GuestOSService{cs: cs}
}

type UCSService struct {
	cs *CloudStackClient
}

func NewUCSService(cs *CloudStackClient) *UCSService {
	return &UCSService{cs: cs}
}

type SSHService struct {
	cs *CloudStackClient
}

func NewSSHService(cs *CloudStackClient) *SSHService {
	return &SSHService{cs: cs}
}

type LimitService struct {
	cs *CloudStackClient
}

func NewLimitService(cs *CloudStackClient) *LimitService {
	return &LimitService{cs: cs}
}

type ISOService struct {
	cs *CloudStackClient
}

func NewISOService(cs *CloudStackClient) *ISOService {
	return &ISOService{cs: cs}
}

type ImageStoreService struct {
	cs *CloudStackClient
}

func NewImageStoreService(cs *CloudStackClient) *ImageStoreService {
	return &ImageStoreService{cs: cs}
}

type InternalLBService struct {
	cs *CloudStackClient
}

func NewInternalLBService(cs *CloudStackClient) *InternalLBService {
	return &InternalLBService{cs: cs}
}

type S3Service struct {
	cs *CloudStackClient
}

func NewS3Service(cs *CloudStackClient) *S3Service {
	return &S3Service{cs: cs}
}

type CertificateService struct {
	cs *CloudStackClient
}

func NewCertificateService(cs *CloudStackClient) *CertificateService {
	return &CertificateService{cs: cs}
}

type ClusterService struct {
	cs *CloudStackClient
}

func NewClusterService(cs *CloudStackClient) *ClusterService {
	return &ClusterService{cs: cs}
}

type PoolService struct {
	cs *CloudStackClient
}

func NewPoolService(cs *CloudStackClient) *PoolService {
	return &PoolService{cs: cs}
}

type AddressService struct {
	cs *CloudStackClient
}

func NewAddressService(cs *CloudStackClient) *AddressService {
	return &AddressService{cs: cs}
}

type VPNService struct {
	cs *CloudStackClient
}

func NewVPNService(cs *CloudStackClient) *VPNService {
	return &VPNService{cs: cs}
}

type RouterService struct {
	cs *CloudStackClient
}

func NewRouterService(cs *CloudStackClient) *RouterService {
	return &RouterService{cs: cs}
}

type RegionService struct {
	cs *CloudStackClient
}

func NewRegionService(cs *CloudStackClient) *RegionService {
	return &RegionService{cs: cs}
}

type DiskOfferingService struct {
	cs *CloudStackClient
}

func NewDiskOfferingService(cs *CloudStackClient) *DiskOfferingService {
	return &DiskOfferingService{cs: cs}
}

type ResourcetagsService struct {
	cs *CloudStackClient
}

func NewResourcetagsService(cs *CloudStackClient) *ResourcetagsService {
	return &ResourcetagsService{cs: cs}
}

type APIDiscoveryService struct {
	cs *CloudStackClient
}

func NewAPIDiscoveryService(cs *CloudStackClient) *APIDiscoveryService {
	return &APIDiscoveryService{cs: cs}
}

type VLANService struct {
	cs *CloudStackClient
}

func NewVLANService(cs *CloudStackClient) *VLANService {
	return &VLANService{cs: cs}
}

type ResourcemetadataService struct {
	cs *CloudStackClient
}

func NewResourcemetadataService(cs *CloudStackClient) *ResourcemetadataService {
	return &ResourcemetadataService{cs: cs}
}

type SystemCapacityService struct {
	cs *CloudStackClient
}

func NewSystemCapacityService(cs *CloudStackClient) *SystemCapacityService {
	return &SystemCapacityService{cs: cs}
}

type ServiceOfferingService struct {
	cs *CloudStackClient
}

func NewServiceOfferingService(cs *CloudStackClient) *ServiceOfferingService {
	return &ServiceOfferingService{cs: cs}
}

type NetworkOfferingService struct {
	cs *CloudStackClient
}

func NewNetworkOfferingService(cs *CloudStackClient) *NetworkOfferingService {
	return &NetworkOfferingService{cs: cs}
}

type OvsElementService struct {
	cs *CloudStackClient
}

func NewOvsElementService(cs *CloudStackClient) *OvsElementService {
	return &OvsElementService{cs: cs}
}

type LogoutService struct {
	cs *CloudStackClient
}

func NewLogoutService(cs *CloudStackClient) *LogoutService {
	return &LogoutService{cs: cs}
}

type CloudIdentifierService struct {
	cs *CloudStackClient
}

func NewCloudIdentifierService(cs *CloudStackClient) *CloudIdentifierService {
	return &CloudIdentifierService{cs: cs}
}

type AutoScaleService struct {
	cs *CloudStackClient
}

func NewAutoScaleService(cs *CloudStackClient) *AutoScaleService {
	return &AutoScaleService{cs: cs}
}

type TemplateService struct {
	cs *CloudStackClient
}

func NewTemplateService(cs *CloudStackClient) *TemplateService {
	return &TemplateService{cs: cs}
}

type PodService struct {
	cs *CloudStackClient
}

func NewPodService(cs *CloudStackClient) *PodService {
	return &PodService{cs: cs}
}

type StoragePoolService struct {
	cs *CloudStackClient
}

func NewStoragePoolService(cs *CloudStackClient) *StoragePoolService {
	return &StoragePoolService{cs: cs}
}

type NicService struct {
	cs *CloudStackClient
}

func NewNicService(cs *CloudStackClient) *NicService {
	return &NicService{cs: cs}
}

type SwiftService struct {
	cs *CloudStackClient
}

func NewSwiftService(cs *CloudStackClient) *SwiftService {
	return &SwiftService{cs: cs}
}

type LoadBalancerService struct {
	cs *CloudStackClient
}

func NewLoadBalancerService(cs *CloudStackClient) *LoadBalancerService {
	return &LoadBalancerService{cs: cs}
}

type HostService struct {
	cs *CloudStackClient
}

func NewHostService(cs *CloudStackClient) *HostService {
	return &HostService{cs: cs}
}

type SnapshotService struct {
	cs *CloudStackClient
}

func NewSnapshotService(cs *CloudStackClient) *SnapshotService {
	return &SnapshotService{cs: cs}
}

type EventService struct {
	cs *CloudStackClient
}

func NewEventService(cs *CloudStackClient) *EventService {
	return &EventService{cs: cs}
}

type UsageService struct {
	cs *CloudStackClient
}

func NewUsageService(cs *CloudStackClient) *UsageService {
	return &UsageService{cs: cs}
}

type StratosphereSSPService struct {
	cs *CloudStackClient
}

func NewStratosphereSSPService(cs *CloudStackClient) *StratosphereSSPService {
	return &StratosphereSSPService{cs: cs}
}

type NATService struct {
	cs *CloudStackClient
}

func NewNATService(cs *CloudStackClient) *NATService {
	return &NATService{cs: cs}
}

type AlertService struct {
	cs *CloudStackClient
}

func NewAlertService(cs *CloudStackClient) *AlertService {
	return &AlertService{cs: cs}
}

type AffinityGroupService struct {
	cs *CloudStackClient
}

func NewAffinityGroupService(cs *CloudStackClient) *AffinityGroupService {
	return &AffinityGroupService{cs: cs}
}

type VMGroupService struct {
	cs *CloudStackClient
}

func NewVMGroupService(cs *CloudStackClient) *VMGroupService {
	return &VMGroupService{cs: cs}
}

type VPCService struct {
	cs *CloudStackClient
}

func NewVPCService(cs *CloudStackClient) *VPCService {
	return &VPCService{cs: cs}
}

type UserService struct {
	cs *CloudStackClient
}

func NewUserService(cs *CloudStackClient) *UserService {
	return &UserService{cs: cs}
}

type SystemVMService struct {
	cs *CloudStackClient
}

func NewSystemVMService(cs *CloudStackClient) *SystemVMService {
	return &SystemVMService{cs: cs}
}

type ZoneService struct {
	cs *CloudStackClient
}

func NewZoneService(cs *CloudStackClient) *ZoneService {
	return &ZoneService{cs: cs}
}

type DomainService struct {
	cs *CloudStackClient
}

func NewDomainService(cs *CloudStackClient) *DomainService {
	return &DomainService{cs: cs}
}

type BaremetalService struct {
	cs *CloudStackClient
}

func NewBaremetalService(cs *CloudStackClient) *BaremetalService {
	return &BaremetalService{cs: cs}
}

type PortableIPService struct {
	cs *CloudStackClient
}

func NewPortableIPService(cs *CloudStackClient) *PortableIPService {
	return &PortableIPService{cs: cs}
}

type VolumeService struct {
	cs *CloudStackClient
}

func NewVolumeService(cs *CloudStackClient) *VolumeService {
	return &VolumeService{cs: cs}
}

type NiciraNVPService struct {
	cs *CloudStackClient
}

func NewNiciraNVPService(cs *CloudStackClient) *NiciraNVPService {
	return &NiciraNVPService{cs: cs}
}

type HypervisorService struct {
	cs *CloudStackClient
}

func NewHypervisorService(cs *CloudStackClient) *HypervisorService {
	return &HypervisorService{cs: cs}
}

type BigSwitchVNSService struct {
	cs *CloudStackClient
}

func NewBigSwitchVNSService(cs *CloudStackClient) *BigSwitchVNSService {
	return &BigSwitchVNSService{cs: cs}
}

type NetworkService struct {
	cs *CloudStackClient
}

func NewNetworkService(cs *CloudStackClient) *NetworkService {
	return &NetworkService{cs: cs}
}

type VirtualMachineService struct {
	cs *CloudStackClient
}

func NewVirtualMachineService(cs *CloudStackClient) *VirtualMachineService {
	return &VirtualMachineService{cs: cs}
}

type ConfigurationService struct {
	cs *CloudStackClient
}

func NewConfigurationService(cs *CloudStackClient) *ConfigurationService {
	return &ConfigurationService{cs: cs}
}

type AsyncjobService struct {
	cs *CloudStackClient
}

func NewAsyncjobService(cs *CloudStackClient) *AsyncjobService {
	return &AsyncjobService{cs: cs}
}

type LoginService struct {
	cs *CloudStackClient
}

func NewLoginService(cs *CloudStackClient) *LoginService {
	return &LoginService{cs: cs}
}
