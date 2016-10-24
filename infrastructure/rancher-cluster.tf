module "rancher_cluster" {
  source = "./modules/rancher/cluster"

  rancher_host_url = "http://${module.rancher-server.rancher_dns}:8080/v1/scripts"
  rancher_host_token = "252568D35F5100C0D2C2:1477328400000:bPgdC4w1CzNgMz5D4bnQFHeQ8"
  vpc_id = "${module.vpc.vpc_id}"
  vpc_cidr_block = "10.0.0.0/16"
  subnet_ids = [
    "${module.vpc.public_subnets}",
    "${module.vpc.private_subnets}"
  ]
  security_groups = [
    "${module.vpc.default_security_group_id}",
    "${module.rancher-server.db_security_group}"
  ]

  rancher_asg_min_size = "1"
  rancher_asg_max_size = "4"
  rancher_asg_desired_capacity = "1"
  rancher_key_name = "rancher"

  ami = "ami-d732f0b7"
}