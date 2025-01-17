---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "bowtie_resource_group Resource - terraform-provider-bowtie"
subcategory: ""
description: |-
  
---

# bowtie_resource_group (Resource)



## Example Usage

```terraform
resource "bowtie_resource_group" "example" {
  name      = "dev"
  resources = ["a2f99c6e-e9a2-401e-ab2d-3b62e02a2f5d"]
}

# Full Example
resource "bowtie_resource" "cidr" {
  name     = "example"
  protocol = "http"
  location = {
    cidr = "10.0.0.0/16"
  }
  ports = {
    collection = [80, 443]
  }

}

resource "bowtie_resource" "dns" {
  name     = "example"
  protocol = "https"
  location = {
    dns = "test.example.com"
  }
  ports = {
    collection = [443, 80, 8080]
  }
}

resource "bowtie_resource_group" "corp" {
  name      = "corp"
  resources = [bowtie_resource.corp.id]
}

resource "bowtie_resource_group" "dns" {
  name      = "dns record"
  resources = [bowtie_resource.dns.id]
  inherited = [bowtie_resource_group.corp.id]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `inherited` (List of String) The list of resource groups to include in this resource group
- `name` (String) The human readable name/description of the resource group
- `resources` (List of String) The resources that should directly be included in this resource group

### Read-Only

- `id` (String) The id for the resource group in the api

## Import

Import is supported using the following syntax:

```shell
terraform import bowtie_resource_group.example 47480e17-e7a2-4f7d-a0c0-3db8fd86c4ff
```
