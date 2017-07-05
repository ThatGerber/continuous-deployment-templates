output "id" {
  value = "${data.aws_security_group.provided.id}"
}

output "name" {
  value = "${data.aws_security_group.provided.name}"
}

output "tags" {
  value = "${data.aws_security_group.provided.tags}"
}

output "vpc_id" {
  value = "${data.aws_security_group.provided.vpc_id}"
}

output "description" {
  value = "${data.aws_security_group.provided.description}"
}

output "arn" {
  value = "${data.aws_security_group.provided.arn}"
}
