# Copyright (c) HashiCorp, Inc.

resource "vnpaycloud_blockstorage_volume" "volume-01" {
  name        = "vnpaycloud-volume-example"
  description = "vnpaycloud-volume-example"
  size        = 20
  volume_type = "c1-standard"
}