data "aws_vpc" "selected" {
  id = "${var.vpc_id}"
}

resource "aws_launch_configuration" "rancher" {
  image_id = "${var.host_ami_id}"
  instance_type = "${var.host_instance_type}"
  key_name = "${var.host_key_name}"
  iam_instance_profile = "${var.host_iam_instance_profile}"
  security_groups = [
    "${aws_security_group.rancher.id}",
    "${var.host_additional_security_group_ids}"
  ]
  ebs_optimized = true
  root_block_device {
    volume_type = "${var.host_root_volume_type}"
    volume_size = "${var.host_root_volume_size}"
    delete_on_termination = true
  }
  user_data = "${data.template_file.user_data.rendered}"
  lifecycle {
    create_before_destroy = true
  }
}

data "template_file" "user_data" {
  template = "${file("${path.module}/user_data.tmpl")}"
  vars {
    rancher_agent_image = "${var.rancher_agent_image}"
    rancher_api_url = "${var.rancher_api_url}"
    rancher_api_version = "${var.rancher_api_version}"
    rancher_environment_token = "${var.rancher_environment_token}"
    rancher_host_labels = "host.type=${var.type}"
  }
}

resource "aws_autoscaling_group" "rancher" {
  name = "${var.environment}_rancher-${var.type}"
  max_size = "${var.asg_max_size}"
  min_size = "${var.asg_min_size}"
  desired_capacity = "${var.asg_desired_capacity}"
  launch_configuration = "${aws_launch_configuration.rancher.id}"
  load_balancers = ["${var.asg_load_balancers}"]
  health_check_type = "EC2"
  health_check_grace_period = 300
  vpc_zone_identifier = [
    "${var.asg_subnet_ids}"
  ]
  tag {
    key = "Name"
    value = "${var.environment}_rancher-${var.type}"
    propagate_at_launch = true
  }
}

resource "aws_security_group" "rancher" {
  name = "${var.environment}_rancher-${var.type}"
  description = "Allow traffic to rancher instances"
  ingress {
    from_port = 500
    to_port = 500
    protocol = "udp"
    self = true
  }
  ingress {
    from_port = 4500
    to_port = 4500
    protocol = "udp"
    self = true
  }
  egress {
    from_port = 500
    to_port = 500
    protocol = "udp"
    self = true
  }
  egress {
    from_port = 4500
    to_port = 4500
    protocol = "udp"
    self = true
  }
  vpc_id = "${var.vpc_id}"
}
