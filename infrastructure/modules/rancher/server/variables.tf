//instance
variable ami{
  default = "ami-d732f0b7"
}

//Route 53
variable "r53_zone_id" {}

//RDS
variable "vpc_id" {}
variable "vpc_subnet_id" {
  type = "list"
}
variable "vpc_public_subnet" {}
variable "vpc_private_subnet" {}

variable "db_master_password" {
  default = "b1w0rldw1d3"
}

variable "vpc_sg" {}
output "address" {
  value = "${aws_db_instance.rancher_db.address}"
}

output "rancher_dns" {
  value = "${aws_route53_record.rancher_dns.fqdn}"
}

output "db_security_group" {
  value = "${aws_security_group.db_access.id}"
}