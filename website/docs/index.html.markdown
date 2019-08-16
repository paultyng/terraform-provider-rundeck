---
layout: "rundeck"
page_title: "Provider: Rundeck"
sidebar_current: "docs-rundeck-index"
description: |-
  The Rundeck provider configures projects, jobs and keys in Rundeck.
---

# Rundeck Provider

The Rundeck provider allows Terraform to create and configure Projects,
Jobs and Keys in [Rundeck](http://rundeck.org/). Rundeck is a tool
for runbook automation and execution of arbitrary management tasks,
allowing operators to avoid logging in to individual machines directly
via SSH.

The provider configuration block accepts the following arguments:

* ``url`` - (Required) The root URL of a Rundeck server. May alternatively be set via the
  ``RUNDECK_URL`` environment variable.

* ``api_version`` - (Optional) The API version of the server. Defaults to `14`, the
  minium supported version. May alternatively be set via the ``RUNDECK_API_VERSION``
  environment variable. 

* ``auth_token`` - (Required) The API auth token to use when making requests. May alternatively
  be set via the ``RUNDECK_AUTH_TOKEN`` environment variable.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
provider "rundeck" {
  url         = "http://rundeck.example.com/"
  api_version = "26"
  auth_token  = "abcd1234"
}

resource "rundeck_project" "anvils" {
  name        = "anvils"
  description = "Application for managing Anvils"

  ssh_key_storage_path = "${rundeck_private_key.anvils.path}"

  resource_model_source {
    type = "file"

    config = {
      format = "resourcexml"

      # This path is interpreted on the Rundeck server.
      file = "/var/rundeck/projects/anvils/resources.xml"
    }
  }
}

resource "rundeck_job" "bounceweb" {
  name              = "Bounce Web Servers"
  project_name      = "${rundeck_project.anvils.name}"
  node_filter_query = "tags: web"
  description       = "Restart the service daemons on all the web servers"

  command {
    shell_command = "sudo service anvils restart"
  }
}

resource "rundeck_public_key" "anvils" {
  path         = "anvils/id_rsa.pub"
  key_material = "ssh-rsa yada-yada-yada"
}

resource "rundeck_private_key" "anvils" {
  path         = "anvils/id_rsa"
  key_material = "${file(\"id_rsa.pub\")}"
}

resource "rundeck_password" "anvils" {
  path         = "anvils/admin_login"
  key_material = "${file(\"secret.txt\")}"
}

data "local_file" "acl" {
  filename = "${path.module}/acl.yaml"
}

resource "rundeck_acl_policy" "example" {
  name = "ExampleAcl.aclpolicy"

  policy = "${data.local_file.acl.content}"
}
```
