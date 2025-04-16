package main

import (
	"context"
	"flag"
	"log"

	"terraform-provider-vnpaycloud/vnpaycloud/provider"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

const (
    awsAccessKeyID     = "123"
    awsSecret = "kHeUAwnSUizTWpSbyGAz4f+As5LshPIjvtpswqGb"
    awsRegion          = "us-west-2"
)

var (
	version string = "1.0.0"
)

func main() {
	var debug bool



	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/vnpaycloud/vnpaycloud",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), provider.New(version), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
