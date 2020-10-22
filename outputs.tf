output "this_null_resource_id" {
  description = "The ID for the null resource"
  value = element(
    concat(
      null_resource.this.*.id,
      [""],
    ),
    0,
  )
}
