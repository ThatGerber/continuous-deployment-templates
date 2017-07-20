data "template_file" "user_data" {
  template = "${file("${path.module}/user_data.tmpl")}"

  vars {
    image          = "${var.image}"
    docker_version = "${var.docker_version}"
    db_url         = "${var.db_url}"
    db_port        = "${var.db_port}"
    db_schema      = "${var.db_schema}"
    db_username    = "${var.db_username}"
    db_password    = "${var.db_password}"
  }
}
