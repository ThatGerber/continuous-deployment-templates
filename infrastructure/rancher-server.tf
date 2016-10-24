module "rancher-server" {
  source = "./modules/rancher/server"

  vpc_id = "${module.vpc.vpc_id}"
  vpc_subnet_id = "${module.vpc.public_subnets}"
  vpc_private_subnet = "10.0.1.0/24"
  vpc_public_subnet  = "10.0.101.0/24"

  vpc_sg = "${module.vpc.default_security_group_id}"
  r53_zone_id = "Z2UE7N9F6X6QW6"

}