provider "aws" {
  region = "us-west-2"
  profile = "opi-devops"
}

data "aws_availability_zones" "available" {}

data "aws_ami" "ubuntu" {
  most_recent = true
  filter {
    name = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-xenial-16.04-amd64-server*"]
  }
}

resource "aws_key_pair" "test" {
  key_name = "test-key"
  public_key = "${file("~/.ssh/id_rsa.pub")}"
}

module "network" {
  source = "github.com/terraform-community-modules/tf_aws_vpc"
  name = "engelman-devops-test-vpc"
  cidr = "10.0.0.0/16"
  private_subnets = ["10.0.1.0/24", "10.0.2.0/24"]
  public_subnets = ["10.0.0.0/24"]
  azs = ["${data.aws_availability_zones.available.names}"]
  enable_dns_hostnames = true
  enable_dns_support = true
  enable_nat_gateway = false
}

module "rancher" {
  source = "../../infrastructure/modules/rancher/server/deployments/single-database"
  vpc_id = "${module.network.vpc_id}"
  subnet_id = "${module.network.public_subnets[0]}"
  ami_id = "${data.aws_ami.ubuntu.id}"
  
  key_name = "${aws_key_pair.test.key_name}"
  
  db_subnet_ids = [
    "${module.network.private_subnets}"
  ]
  db_security_group_ids = [
    "${module.network.default_security_group_id}"
  ]
}

output "db_url" {
  value = "${module.rancher.}"
}
output "url" {
  value = "${module.rancher.rancher_server_dns}"
}
