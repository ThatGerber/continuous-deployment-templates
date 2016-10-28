resource "template_file" "user_data" {
  template = "${file("${path.module}/user_data.tmpl")}"
  vars {
    image = "${var.image}"
    docker_version = "${var.docker_version}"
  }
}
