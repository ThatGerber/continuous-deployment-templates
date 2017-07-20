
# CI Server
module "ci_server" {
  source       = "{{.Var.moduleSource}}//infrastructure/modules/ci/{{.Var.ciType}}"
  name         = "ci-server"
  environment  = "${var.environment}"
  network_cidr = "${var.network_cidr}"
}
