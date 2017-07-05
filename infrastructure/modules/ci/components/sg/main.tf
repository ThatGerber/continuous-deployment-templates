# Create SG or Append Rules to provided SGID

resource "aws_security_group" "main_security_group" {
  count       = "${var.security_group_id ? 0 : 1}"
  name        = "${var.security_group_name}"
  description = "tf-sg-${var.security_group_name}"
  vpc_id      = "${var.vpc_id}"
}

data "aws_security_group" "provided" {
  id = "${var.security_group_id || aws_security_group.main_security_group.id}"
}

// Allow any internal network flow.
resource "aws_security_group_rule" "ingress_any_any_self" {
  security_group_id = "${data.aws_security_group.provided.id}"
  from_port         = 0
  to_port           = 65535
  protocol          = "-1"
  self              = true
  type              = "ingress"
}

// SSH
resource "aws_security_group_rule" "ingress_tcp_22_self" {
  count             = "${var.use_ssh}"
  security_group_id = "${data.aws_security_group.provided.id}"
  from_port         = 22
  to_port           = 22
  protocol          = "tcp"
  cidr_blocks       = "${var.source_cidr_block}"
  type              = "ingress"
}

// HTTP
resource "aws_security_group_rule" "ingress_tcp_80_self" {
  count             = "${var.use_http}"
  security_group_id = "${data.aws_security_group.provided.id}"
  from_port         = 80
  to_port           = 80
  protocol          = "tcp"
  cidr_blocks       = "${var.source_cidr_block}"
  type              = "ingress"
}

// HTTPS
resource "aws_security_group_rule" "ingress_tcp_443_self" {
  count             = "${var.use_https}"
  security_group_id = "${data.aws_security_group.provided.id}"
  from_port         = 443
  to_port           = 443
  protocol          = "tcp"
  cidr_blocks       = "${var.source_cidr_block}"
  type              = "ingress"
}

// JMX
resource "aws_security_group_rule" "ingress_tcp_7199_self" {
  count             = "${var.use_jmx}"
  security_group_id = "${data.aws_security_group.provided.id}"
  from_port         = 7199
  to_port           = 7199
  protocol          = "tcp"
  cidr_blocks       = "${var.source_cidr_block}"
  type              = "ingress"
}

// HTTP Alt
resource "aws_security_group_rule" "ingress_tcp_8080_self" {
  count             = "${var.use_http_alt}"
  security_group_id = "${data.aws_security_group.provided.id}"
  from_port         = 443
  to_port           = 443
  protocol          = "tcp"
  cidr_blocks       = "${var.source_cidr_block}"
  type              = "ingress"
}

// Postgres
resource "aws_security_group_rule" "ingress_tcp_5432_self" {
  count             = "${var.use_postgres}"
  security_group_id = "${data.aws_security_group.provided.id}"
  from_port         = 5432
  to_port           = 5432
  protocol          = "tcp"
  cidr_blocks       = "${var.source_cidr_block}"
  type              = "ingress"
}

// MySQL
resource "aws_security_group_rule" "ingress_tcp_3306_self" {
  count             = "${var.use_mysql}"
  security_group_id = "${data.aws_security_group.provided.id}"
  from_port         = 3306
  to_port           = 3306
  protocol          = "tcp"
  cidr_blocks       = "${var.source_cidr_block}"
  type              = "ingress"
}
