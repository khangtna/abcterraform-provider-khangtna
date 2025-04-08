terraform {
  required_providers {
    vnpaycloud = {
      source = "terraform-provider-vnpaycloud/vnpaycloud"
      version = "1.0.0"
    }
  }
}

provider "vnpaycloud" {
  auth_url = "a"
  application_credential_id = "b"
  application_credential_name = "c"
  application_credential_secret = "d"
}

resource "vnpaycloud_volume" "e" {
  name = "f"
  description = "g"
  size = 30
  volume_type = "c1-standard"
}