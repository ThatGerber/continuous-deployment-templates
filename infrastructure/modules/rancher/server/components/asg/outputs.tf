output "elb_dns" {
  value = "${aws_elb.server.dns_name}"
}
