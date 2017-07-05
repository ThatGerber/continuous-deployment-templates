module "db" {
  source = "../../../../app/db"

  rds_instance_name     = "rancher-${var.stack}-d0${var.environment}"
  rds_allocated_storage = "5"
  rds_instance_class    = "db.t2.small"
  database_name         = "${var.db_schema}"
  database_user         = "${var.db_username}"
  database_password     = "${var.db_password}"

  rds_security_group_ids = [
    "${var.db_security_group_ids}",
  ]

  rds_subnets = [
    "${var.db_subnet_ids}",
  ]
}

module "config" {
  depends_on = ["${module.db}"]

  source      = "../../components/config_database"
  environment = "${var.enviornment}"
  stack       = "${var.stack}"

  image          = "${var.image}"
  docker_version = "${var.docker_version}"
  port           = "${var.port}"

  db_schema   = "${var.db_schema}"
  db_username = "${var.db_username}"
  db_password = "${var.db_password}"
  db_url      = "${module.db.rds_instance_endpoint}"
  db_port     = "${module.db.rds_instance_port}"
}

module "server" {
  source = "../../components/instance"

  depends_on = "${module.config}"

  config = "${module.config.user_data}"

  // Required
  ami_id    = "${var.ami_id}"
  vpc_id    = "${var.vpc_id}"
  subnet_id = "${var.vpc_subnet_id}"

  // Optional
  environment        = "${var.enviornment}"
  stack              = "${var.stack}"
  key_name           = "${var.key_name}"
  instance_profile   = "${var.instance_profile}"
  instance_type      = "${var.instance_type}"
  access_cidr        = "${var.access_cidr}"
  security_group_ids = "${var.security_group_ids}"
  has_public_ip      = "${var.has_public_ip}"
  root_volume_size   = "${var.root_volume_size}"
  port               = "${var.port}"
}
