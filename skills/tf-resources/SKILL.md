---
name: tf-resources
description: Terraform resources - creation and lifecycle
sasmp_version: "1.3.0"
bonded_agent: tf-resources
bond_type: PRIMARY_BOND
---

# Terraform Resources Skill

## Resource Syntax

```hcl
resource "aws_instance" "web" {
  ami           = "ami-123456"
  instance_type = "t2.micro"
  
  tags = {
    Name = "WebServer"
  }
}

# Access attributes
output "instance_id" {
  value = aws_instance.web.id
}
```

## Meta-Arguments

```hcl
# count
resource "aws_instance" "server" {
  count         = 3
  ami           = "ami-123456"
  instance_type = "t2.micro"
  tags = {
    Name = "Server-${count.index}"
  }
}

# for_each
resource "aws_instance" "server" {
  for_each = toset(["web", "api", "db"])
  ami      = "ami-123456"
  tags = {
    Name = each.value
  }
}

# depends_on
resource "aws_instance" "web" {
  depends_on = [aws_db_instance.database]
}

# lifecycle
resource "aws_instance" "web" {
  lifecycle {
    create_before_destroy = true
    prevent_destroy       = true
    ignore_changes        = [tags]
  }
}
```

## Quick Reference

| Meta-Arg | Purpose |
|----------|---------|
| count | Create multiple |
| for_each | Create from map/set |
| depends_on | Explicit dependency |
| lifecycle | Control behavior |
| provider | Specify provider |

## Related
- tf-variables - Input variables
- tf-resources agent
