data "aws_ami" "ubuntu_1604" {
  most_recent = true

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-trusty-16.04-amd64-server-*"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  owners = ["099720109477"] # Canonical
}

module "rancher-server" {
  source = "./modules/rancher/server/deployments/single-database"

  environment = "${var.environment}"
  stack       = "${var.stack}"

  vpc_id        = "${module.vpc.vpc_id}"
  vpc_subnet_id = ""

  db_password           = ""
  db_security_group_ids = ""
  db_subnet_ids         = ""

  ami_id         = "${data.aws_ami.ubuntu_1604.id}"
  docker_version = "1.12.6-0~xenial"
  instance_type  = "t2.micro"

  # Optional
  # image = ""
  # port = "8080"
  # key_name = ""
  # instance_profile = ""
  # access_cidr = ""
  # security_group_ids = []
  # has_public_ip = true
  # root_volume_size = "32"
  # db_schema = ""
  # db_username = ""
  # db_url = ""
  # db_port = ""
}
