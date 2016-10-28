output "rancher_server_dns" {
  value = "${module.server.instance_public_dns}"
}
