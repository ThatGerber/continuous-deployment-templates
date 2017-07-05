// Server Setings
variable "ami_id" {
  description = "The Ubuntu AMI to launch"
}

variable "vpc_id" {}
variable "vpc_subnet_id" {}

variable "db_password" {
  description = "The password to access the database"
}

variable "db_security_group_ids" {
  description = "The security groups to assign to the RDS instance"
}

variable "db_subnet_ids" {
  description = "The subnets to place the RDS instance in. Must be at least 2"
}
