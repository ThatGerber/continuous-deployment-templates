output "rancher_server_url" {
  value = "${module.server.instance_public_dns}:${var.port}"
}

output "rancher_db_address" {
  value = "${module.db.rds_instance_address}"
}
