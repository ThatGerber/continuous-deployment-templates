resource "aws_security_group" "rancher-internal" {
  name        = "rancher-server-internal-d0${var.environment}"
  description = "Allow internal traffic to rancher server"
  vpc_id      = "${var.vpc_id}"

  // HTTP
  ingress {
    from_port = "${var.port}"
    to_port   = "${var.port}"
    protocol  = "tcp"
    self      = true
  }

  egress {
    from_port = "${var.port}"
    to_port   = "${var.port}"
    protocol  = "tcp"
    self      = true
  }
}

resource "aws_security_group" "rancher-external" {
  name        = "rancher-server-external-d0${var.environment}"
  description = "Allow traffic to rancher server ELB"
  vpc_id      = "${var.vpc_id}"

  // HTTP
  ingress {
    from_port = "${var.elb_lb_port}"
    to_port   = "${var.elb_lb_port}"
    protocol  = "tcp"

    cidr_blocks = [
      "${var.elb_access_cidr}",
    ]
  }
}
