package main

import (
	"fmt"

	"github.com/xanzy/go-cloudstack/v2/cloudstack"
	Constants "github.com/xanzy/go-cloudstack/v2/constants"
)

func main() {
	server := Constants.Server
	apikey := Constants.Apikey
	secretkey := Constants.Secretkey

	cs := cloudstack.NewAsyncClient(server, apikey, secretkey, false)

	p := cs.VirtualMachine.NewDeployVirtualMachineParams("<service offering id",
		"<template id>",
		"<zone id>")
	p.SetName("test-vm")
	p.SetDisplayname("test-vm")
	p.SetDiskofferingid("<>")

	r, err := cs.VirtualMachine.DeployVirtualMachine(p)
	if err != nil {
		fmt.Printf("Error creating the new instance %s:", err)
	}

	fmt.Printf("New virtual machine created: Name: %s, ID: %s", r.Name, r.Id)
}
