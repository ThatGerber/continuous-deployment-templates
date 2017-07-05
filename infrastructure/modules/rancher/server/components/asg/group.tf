resource "aws_launch_configuration" "server" {
  name                 = "rancher-${var.stack}-d0${var.environment}-v000"
  image_id             = "${var.ami_id}"
  instance_type        = "${var.instance_type}"
  iam_instance_profile = "${var.instance_profile}"
  key_name             = "${var.key_name}"

  security_groups = [
    "${aws_security_group.rancher-internal.id}",
    "${var.security_group_ids}",
  ]

  associate_public_ip_address = "${var.has_public_ip}"

  root_block_device {
    volume_type = "gp2"
    volume_size = "${var.root_volume_size}"
  }
}

resource "aws_autoscaling_group" "server" {
  name                 = "rancher-${var.stack}-d0${var.environment}-v000"
  max_size             = "${var.asg_max_size}"
  min_size             = "${var.asg_min_size}"
  desired_capacity     = "${var.asg_desired_capacity}"
  launch_configuration = "${aws_launch_configuration.server.id}"
  health_check_type    = "EC2"

  load_balancers = [
    "${aws_elb.server.id}",
  ]

  vpc_zone_identifier = [
    "${var.asg_subnet_ids}",
  ]

  tag {
    Name = "rancher-${var.stack}-d0${var.environment}-v000"
  }
}
