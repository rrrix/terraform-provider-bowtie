---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "bowtie_group Resource - terraform-provider-bowtie"
subcategory: ""
description: |-
  A user group used to assign policies for access to groups of users
---

# bowtie_group (Resource)

A user group used to assign policies for access to groups of users

## Example Usage

```terraform
resource "bowtie_group" "admins" {
  name = "admins"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the group

### Read-Only

- `id` (String) The group identifier
- `last_updated` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import bowtie_group.admins 47480e17-e7a2-4f7d-a0c0-3db8fd86c4ff
```
