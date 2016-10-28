module "db" {
  source = "../../../../app/db"
}

module "config" {
  source = "../../components/config_database"
  environment = "${var.enviornment}"
  stack = "${var.stack}"
  
  image = "${var.image}"
  docker_version = "${var.docker_version}"
  port = "${var.port}"
  
  db_schema = "${var.db_schema}"
  db_username = "${var.db_username}"
  db_password = "${var.db_password}"
  db_url = "${var.db_url}"
  db_port = "${var.db_port}"
}

module "server" {
  source = "../../components/instance"
  config = "${module.config.user_data}"
  
  // Required
  ami_id = "${var.ami_id}"
  vpc_id = "${var.vpc_id}"
  subnet_id = "${var.subnet_id}"
  
  // Optional
  
  environment = "${var.enviornment}"
  stack = "${var.stack}"
  
  key_name = "${var.key_name}"
  instance_profile = "${var.instance_profile}"
  instance_type = "${var.instance_type}"
  access_cidr = "${var.access_cidr}"
  security_group_ids = "${var.security_group_ids}"
  has_public_ip = "${var.has_public_ip}"
  root_volume_size = "${var.root_volume_size}"
  port = "${var.port}"
}
