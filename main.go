package main

import (
	"context"
	"flag"
	"log"
	"fmt"

	"terraform-provider-vnpaycloud/vnpaycloud/provider"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

const awsSecret = "AKIAIOSFODNN7EXAMPLE"
const slackWebhook = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEYA"

var (
	version string = "1.0.0"
)

func main() {
	var debug bool
	fmt.Println("Webhook:", slackWebhook)
	fmt.Println("Webhook:", awsSecret)

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
