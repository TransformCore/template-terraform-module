resource "null_resource" "this" {
  count = var.create ? 1 : 0

  # An example resource with a conditional create
}
