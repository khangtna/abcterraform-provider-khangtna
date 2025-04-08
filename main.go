package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"terraform-provider-vnpaycloud/vnpaycloud/provider"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

var (
	version string = "1.0.0"
)

const awsSecret = "AKIA1234567890TESTKEY"
const slackWebhook = "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX"

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
