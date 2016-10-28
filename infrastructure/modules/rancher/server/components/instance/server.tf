resource "aws_instance" "server" {
  ami = "${var.ami_id}"
  key_name = "${var.key_name}"
  instance_type = "${var.instance_type}"
  vpc_security_group_ids = [
    "${aws_security_group.rancher.id}",
    "${var.security_group_ids}"
  ]
  subnet_id = "${var.subnet_id}"
  associate_public_ip_address = "${var.has_public_id}"
  user_data = "${var.config}"
  iam_instance_profile = "${var.instance_profile}"
  root_block_device {
    volume_type = "gp2"
    volume_size = "${var.root_volume_size}"
  }
  tags {
    Name = "rancher-${var.stack}-d0${var.environment}-v000"
  }
}

resource "aws_security_group" "rancher" {
  name = "rancher-server-d0${var.environment}"
  description = "Allow traffic to rancher server"
  vpc_id = "${var.vpc_id}"

  // HTTP
  ingress {
    from_port = "${var.port}"
    to_port = "${var.port}"
    protocol = "tcp"
    cidr_blocks = [
      "${var.access_cidr}"
    ]
  }
  ingress {
    from_port = 22
    to_port = 22
    protocol = "tcp"
    cidr_blocks = [
      "${var.access_cidr}"
    ]
  }
}
