package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/xanzy/go-cloudstack/v2/cloudstack"
	Constants "github.com/xanzy/go-cloudstack/v2/constants"
)

func main() {
	server := Constants.Server
	apikey := Constants.Apikey
	secretkey := Constants.Secretkey

	cs := cloudstack.NewAsyncClient(server, apikey, secretkey, false)

	params := &cloudstack.ListVirtualMachinesParams{}
	params.SetListall(true)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err == nil {
		text = strings.Trim(text, "\n")
		if text != "" {
			fmt.Printf("searching for vm %s\n", text)
			params.SetKeyword(text)
		}

	}

	var response *cloudstack.ListVirtualMachinesResponse
	response, _ = cs.VirtualMachine.ListVirtualMachines(params)

	for i := 0; i < response.Count; i++ {
		fmt.Println("Name is " + response.VirtualMachines[i].Name)
		fmt.Println("Id is " + response.VirtualMachines[i].Id)
	}

}
