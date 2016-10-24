module "rancher-server" {
  source = "./modules/rancher/server"

  vpc_id = "${module.vpc.vpc_id}"
  vpc_subnet_ids = [
    "${module.vpc.public_subnets}",
    "${module.vpc.private_subnets}"
  ]
  vpc_private_subnet = "10.0.1.0/24"
  vpc_public_subnet  = "10.0.101.0/24"

//  r53_zone_id = ""
//  domain = ""
//  public_ip = ""

}