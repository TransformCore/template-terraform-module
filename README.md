# Terraform Module Template

A base template for Terraform modules created at ENGINE Transformation which comes with default standards we use such as terratest, pre-commit terraform docs and a simple pipeline to make sure people have terraform fmt'd

## Requirements

- Terraform >=0.12.x
- Go >=1.5
- Pre-commit

## Testing

In order to run tests in the `test/` folder run the following command:
```bash
make test
```

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

| Name | Version |
|------|---------|
| terraform | >= 0.13 |
| aws | >= 3, < 4.0 |

## Providers

| Name | Version |
|------|---------|
| null | n/a |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| create | Conditional argument to create the resource | `bool` | `true` | no |

## Outputs

| Name | Description |
|------|-------------|
| this\_null\_resource\_id | The ID for the null resource |

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->

## License

See LICENSE for full details.
