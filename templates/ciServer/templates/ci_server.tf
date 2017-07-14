
# CI Server
module "ci_server" {
  source       = "{{.Variables.moduleSource}}//infrastructure/modules/ci/{{.Variables.ciType}}"
  name         = "ci-server"
  environment  = "${var.environment}"
  network_cidr = "${var.network_cidr}"
}
