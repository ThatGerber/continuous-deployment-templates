output "rds_instance_id" {
  description = "Output the ID of the RDS instance"
  value       = "${aws_db_instance.main_rds_instance.id}"
}

output "rds_instance_address" {
  description = "Output the Address of the RDS instance"
  value       = "${aws_db_instance.main_rds_instance.address}"
}

output "rds_instance_endpoint" {
  value = "${aws_db_instance.main_rds_instance.endpoint}"
}

output "rds_instance_port" {
  value = "${aws_db_instance.main_rds_instance.port}"
}

output "subnet_group_id" {
  description = "Output the ID of the Subnet Group"
  value       = "${aws_db_subnet_group.main_db_subnet_group.id}"
}
