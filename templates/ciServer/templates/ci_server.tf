
# CI Server
module "ci_server" {
  source       = "{{.moduleSource}}//infrastructure/modules/ci/{{.ciType}}"
  name         = "ci-server"
  environment  = "${var.environment}"
  network_cidr = "${var.network_cidr}"
}
