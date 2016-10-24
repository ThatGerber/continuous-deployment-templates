/*variable "rancher_security_groups_ids" {
  type = "list"
  description = "Security Groups to apply to Rancher agent machines"
}*/

variable "rancher_host_url" {
  description = "Rancher Server URL"
}

variable "rancher_host_token" {
  description = "Rancher agent environment token"
}

variable "vpc_id" {
  description = "VPC ID to add resources in"
}

variable "vpc_cidr_block" {
  description = "The CIDR block of the VPC"
}

variable "subnet_ids" {
  type = "list"
  description = "The subnets to launch Rancher agents in"
}

variable "rancher_asg_min_size" {
  description = "The Rancher auto scaling group minimal size"
  default = 1
}

variable "rancher_asg_max_size" {
  description = "The Rancher auto scaling group minimal size"
  default = 4
}

variable "rancher_asg_desired_capacity" {
  description = "The Rancher auto scaling group minimal size"
  default = 1
}

variable "rancher_key_name" {
  description = "The Rancher auto scaling group minimal size"
}

variable "ami" {
  default = "ami-2d39803a"
}

resource "aws_launch_configuration" "rancher" {
  image_id = "${var.ami}"
  instance_type = "m4.large"
  key_name = "${var.rancher_key_name}"
  iam_instance_profile = "rancher-profile"
  security_groups = [
    "${aws_security_group.rancher.id}"
  ]
  associate_public_ip_address = true
  ebs_optimized = true
  ebs_block_device {
    device_name = "sdf"
    volume_type = "gp2"
    volume_size = "32"
    delete_on_termination = true
  }
  user_data = <<EOF
#cloud-config
packages:
  - ntp
  - linux-image-extra-virtual
  - linux-image-extra-3.13.0-67-generic
write_files:
  - path: /opt/docker/install.sh
    permissions: "0755"
    content: |
      #!/bin/bash
      DOCKER_VERSION=$1
      apt-key adv --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys 58118E89F3A912897C070ADBF76221572C52609D
      mkdir -p /etc/apt/sources.list.d/
      echo "deb https://apt.dockerproject.org/repo ubuntu-trusty main" > /etc/apt/sources.list.d/docker.list
      apt-get update
      apt-get install -y docker-engine=$DOCKER_VERSION
runcmd:
  - [ cloud-init-per, once, docker, /opt/docker/install.sh, 1.11.2-0~trusty ]
  - docker run -e CATTLE_AGENT_IP=`ec2metadata --local-ipv4` -d --privileged -v /var/run/docker.sock:/var/run/docker.sock rancher/agent:latest ${var.rancher_host_url}/${var.rancher_host_token}
EOF
  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_autoscaling_group" "rancher" {
  name = "rancher_rancher-hosts_default"
  max_size = "${var.rancher_asg_max_size}"
  min_size = "${var.rancher_asg_min_size}"
  desired_capacity = "${var.rancher_asg_desired_capacity}"
  launch_configuration = "${aws_launch_configuration.rancher.id}"
  health_check_type = "EC2"
  health_check_grace_period = 300
  vpc_zone_identifier = [
    "${var.subnet_ids}"
  ]
}

resource "aws_security_group" "rancher" {
  name = "rancher-agent"
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
  ingress {
    from_port = 0
    to_port = 0
    protocol = "-1"
    cidr_blocks = [
      "${var.vpc_cidr_block}"
    ]
  }
  egress {
    from_port = 443
    to_port = 443
    protocol = "tcp"
    cidr_blocks = [
      "0.0.0.0/0"
    ]
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
