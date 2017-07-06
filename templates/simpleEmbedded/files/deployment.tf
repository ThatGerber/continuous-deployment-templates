variable "region" {}

variable "profile" {}

variable "ssh_public_key_path" {
  default = "~/.ssh/id_rsa.pub"
}

variable "environment" {
  default = "tools"
}

variable "stack" {
  default = "server"
}

variable "network_cidr" {
  default = "10.0.0.0/16"
}

provider "aws" {
  region  = "${var.region}"
  profile = "${var.profile}"
  allowed_account_ids = ["${var.aws_account_id}"]
}

data "aws_availability_zones" "available" {}

data "aws_ami" "ubuntu_xenial" {
  most_recent = true

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-xenial-16.04-amd64-server-*"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  owners = ["099720109477"] # Canonical
}

resource "aws_key_pair" "test" {
  key_name   = "test-key"
  public_key = "${file("${var.ssh_public_key_path}")}"
}

module "network" {
  source               = "github.com/terraform-community-modules/tf_aws_vpc"
  name                 = "rancher-${var.stack}-d0${var.environment}"
  cidr                 = "${var.network_cidr}"
  private_subnets      = []
  public_subnets       = ["10.0.0.0/16"]
  azs                  = ["${data.aws_availability_zones.available.names}"]
  enable_dns_hostnames = true
  enable_dns_support   = true
  enable_nat_gateway   = false
}

module "rancher" {
  source    = "{{.Variables.moduleSource}}/infrastructure/modules/rancher/server/deployments/standalone"
  vpc_id    = "${module.network.vpc_id}"
  subnet_id = "${module.network.public_subnets[0]}"
  ami_id    = "${data.aws_ami.ubuntu_xenial.image_id}"

  key_name = "${aws_key_pair.test.key_name}"
}

output "url" {
  value = "${module.rancher.rancher_server_url}"
}
