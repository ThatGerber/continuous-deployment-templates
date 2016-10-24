module "rancher_cluster" {
  source = "./modules/rancher/cluster"

  rancher_host_url = ""
  rancher_host_token = ""
  vpc_id = "${module.vpc.vpc_id}"
  vpc_cidr_block = "10.0.0.0/16"
  subnet_ids = [
    "${module.vpc.public_subnets}",
    "${module.vpc.private_subnets}"
  ]

  rancher_asg_min_size = "1"
  rancher_asg_max_size = "4"
  rancher_asg_desired_capacity = "1"
  rancher_key_name = ""

  ami = "ami-2d39803a"
}