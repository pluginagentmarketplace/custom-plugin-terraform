---
name: tf-functions
description: Terraform built-in functions
sasmp_version: "1.3.0"
bonded_agent: tf-advanced
bond_type: PRIMARY_BOND
---

# Terraform Functions Skill

## String Functions

```hcl
# Common
format("Hello, %s!", var.name)
join("-", ["a", "b", "c"])
split(",", "a,b,c")
replace("hello", "l", "L")
trim("  hello  ", " ")
lower("HELLO")
upper("hello")
```

## Collection Functions

```hcl
# List/Set
length(var.list)
element(var.list, 0)
concat(list1, list2)
flatten([[1,2], [3,4]])
distinct([1, 2, 2, 3])
sort([3, 1, 2])

# Map
lookup(var.map, "key", "default")
merge(map1, map2)
keys(var.map)
values(var.map)
```

## Numeric/Logic

```hcl
min(1, 2, 3)
max(1, 2, 3)
abs(-10)
ceil(1.5)
floor(1.5)

coalesce("", "", "hello")
try(local.optional, "default")
can(local.maybe_error)
```

## Encoding/File

```hcl
base64encode("hello")
base64decode("aGVsbG8=")
jsonencode({ key = "value" })
jsondecode("{}")
yamlencode({ key = "value" })
file("script.sh")
templatefile("template.tpl", { name = "world" })
```

## Quick Reference

| Category | Functions |
|----------|-----------|
| String | format, join, split, replace |
| Collection | length, element, merge |
| Numeric | min, max, abs, ceil |
| Type | tostring, tolist, tomap |
| File | file, templatefile |

## Related
- tf-hcl - Expressions
- tf-advanced agent
