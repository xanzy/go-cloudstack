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

	params := &cloudstack.ListAccountsParams{}
	params.SetListall(true)

	var response *cloudstack.ListAccountsResponse
	response, _ = cs.Account.ListAccounts(params)

	for i := 0; i < response.Count; i++ {
		fmt.Println(response.Accounts[i].Name)
	}
}
