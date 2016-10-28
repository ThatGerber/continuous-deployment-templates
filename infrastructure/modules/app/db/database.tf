resource "aws_db_instance" "main_rds_instance" {
    identifier = "${var.rds_instance_name}"
    allocated_storage = "${var.rds_allocated_storage}"
    engine = "${var.rds_engine_type}"
    engine_version = "${var.rds_engine_version}"
    instance_class = "${var.rds_instance_class}"
    name = "${var.database_name}"
    username = "${var.database_user}"
    password = "${var.database_password}"
    vpc_security_group_ids = ["${var.rds_security_group_ids}"]
    db_subnet_group_name = "${aws_db_subnet_group.main_db_subnet_group.name}"
    parameter_group_name = "${var.db_parameter_group}"
    multi_az = "${var.rds_is_multi_az}"
    storage_type = "${var.rds_storage_type}"
    publicly_accessible = "${var.rds_is_public}"
}

resource "aws_db_subnet_group" "main_db_subnet_group" {
    name = "${var.rds_instance_name}-subnetgrp"
    description = "RDS subnet group"
    subnet_ids = ["${var.rds_subnets}"]
}
