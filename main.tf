terraform {
  required_providers {
    vnpaycloud = {
      source = "terraform-provider-vnpaycloud/vnpaycloud"
      version = "1.0.0"
    }
  }
}

provider "vnpaycloud" {
  auth_url = "aaa"
  application_credential_id = "aaaa"
  application_credential_name = "tf_dev_for_code"
  application_credential_secret = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
}

resource "vnpaycloud_volume" "volume_created_by_tf_thuannt" {
  name = "volume_created_by_tf_thuannt"
  description = "volume_created_by_tf_thuannt"
  size = 30
  volume_type = "c1-standard"
}