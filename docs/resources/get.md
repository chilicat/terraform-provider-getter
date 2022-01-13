
---
page_title: "Download File"
description: |-
  Downloads a file
---

# get Resource/Data Source

Downloads a file


## Example Usage

```hcl
resource "getter_get" "myfile" {
  url = "https://server/my/file.txt"
  dst = "/my/dest"
}
```

## Argument Reference
* `url` - (Required) URL to download file from
* `dst` - (Required) Download destination
* `insecure` - (Optional) Set to True to disable TLS validation

## Attribute Reference
