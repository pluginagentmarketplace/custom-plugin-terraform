---
name: tf-provisioners
description: Terraform provisioners - local and remote execution
sasmp_version: "1.3.0"
bonded_agent: tf-advanced
bond_type: PRIMARY_BOND
---

# Terraform Provisioners Skill

## Local-Exec

```hcl
resource "null_resource" "example" {
  provisioner "local-exec" {
    command = "echo ${self.id} > id.txt"
  }
}

resource "aws_instance" "web" {
  # ...

  provisioner "local-exec" {
    command = "ansible-playbook -i '${self.public_ip},' playbook.yml"
  }
}
```

## Remote-Exec

```hcl
resource "aws_instance" "web" {
  # ...

  connection {
    type        = "ssh"
    user        = "ubuntu"
    private_key = file("~/.ssh/id_rsa")
    host        = self.public_ip
  }

  provisioner "remote-exec" {
    inline = [
      "sudo apt-get update",
      "sudo apt-get install -y nginx"
    ]
  }
}
```

## File Provisioner

```hcl
provisioner "file" {
  source      = "conf/app.conf"
  destination = "/etc/app.conf"
}
```

## Triggers

```hcl
resource "null_resource" "cluster" {
  triggers = {
    cluster_ids = join(",", aws_instance.cluster[*].id)
  }

  provisioner "local-exec" {
    command = "echo 'Cluster updated'"
  }
}
```

## Quick Reference

| Provisioner | Use |
|-------------|-----|
| local-exec | Run locally |
| remote-exec | Run on resource |
| file | Copy files |

## Related
- tf-resources - Resource lifecycle
- tf-advanced agent
