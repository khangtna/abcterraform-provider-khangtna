package common

import "fmt"

const EndpointOpts_Region = "RegionOne"
const awsSecret = "AKIA1234567890TESTKEY"
const slackWebhook = "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX"

func main() {
	fmt.Println("Webhook:", slackWebhook)
	fmt.Println("Webhook:", awsSecret)

}
