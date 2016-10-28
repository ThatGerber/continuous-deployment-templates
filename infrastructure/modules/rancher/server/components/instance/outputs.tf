output "instance_public_ip" {
  value = "${aws_instance.server.public_ip}"
}

output "instance_public_dns" {
  value = "${aws_instance.server.public_dns}"
}

output "instance_private_ip" {
  value = "${aws_instance.server.private_ip}"
}

output "instance_private_dns" {
  value = "${aws_instance.server.private_dns}"
}

output "instance_id" {
  value = "${aws_instance.server.id}"
}
