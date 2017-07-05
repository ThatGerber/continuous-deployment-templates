output "vpc_id" {
  value = "${module.devops_vpc.vpc_id}"
}

output "public_subnets" {
  value = ["${module.devops_vpc.public_subnets}"]
}

output "public_subnet_ips" {
  value = ["${values(data.external.public_subnets.result)}"]
}

output "private_subnets" {
  value = ["${module.devops_vpc.private_subnets}"]
}

output "private_subnet_ips" {
  value = ["${values(data.external.private_subnets.result)}"]
}

output "database_subnets" {
  value = ["${module.devops_vpc.database_subnets}"]
}

output "database_subnet_group" {
  value = "${module.devops_vpc.database_subnet_group}"
}

output "elasticache_subnets" {
  value = ["${module.devops_vpc.elasticache_subnets}"]
}

output "elasticache_subnet_group" {
  value = "${module.devops_vpc.elasticache_subnet_group}"
}

output "default_security_group_id" {
  value = "${module.devops_vpc.default_security_group_id}"
}
