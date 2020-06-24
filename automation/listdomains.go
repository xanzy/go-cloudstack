package main

import (
	"fmt"
	"github.com/xanzy/go-cloudstack/v2/cloudstack"
	Constants "github.com/xanzy/go-cloudstack/v2/constants"
)

func main()  {
	server := Constants.Server
	apikey := Constants.Apikey
	secretkey := Constants.Secretkey

	cs := cloudstack.NewAsyncClient(server, apikey, secretkey,false)

	params := &cloudstack.ListDomainsParams{}
	params.SetListall(true)
	//p, err := cs.Domain.ListDomains(params)
	var response *cloudstack.ListDomainsResponse
	response, _ = cs.Domain.ListDomains(params)

	for i := 0 ; i < response.Count ; i++ {
		fmt.Println(response.Domains[i].Name)
	}
}
