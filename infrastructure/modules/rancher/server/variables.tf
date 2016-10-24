//instance
variable ami{
  default = "ami-2d39803a"
}

//Route 53
//variable "r53_zone_id" {}
//variable "domain" {}
//variable "public_ip" {}

//RDS
variable "vpc_id" {}
variable "vpc_subnet_ids" {
  type = "list"
}
variable "vpc_public_subnet" {}
variable "vpc_private_subnet" {}

variable "db_master_password" {
  default = "b1w0rldw1d3"
}

output "address" {
  value = "${aws_db_instance.rancher_db.address}"
}
