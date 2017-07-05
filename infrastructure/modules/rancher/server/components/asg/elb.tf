resource "aws_elb" "server" {
  name = "rancher-${var.stack}-d0${var.environment}"

  security_groups = [
    "${aws_security_group.rancher-internal.id}",
    "${aws_security_group.rancher-external.id}",
    "${var.elb_security_group_ids}",
  ]

  subnets = [
    "${var.elb_subnet_ids}",
  ]

  listener {
    lb_port           = "${var.elb_lb_port}"
    lb_protocol       = "tcp"
    instance_port     = "${var.port}"
    instance_protocol = "tcp"
  }

  tags {
    Name = "rancher-${var.stack}-d0${var.environment}"
  }
}
