variable "security_group_name" {
  description = <<EOH
    Name of security Group.
    Required unless security group ID is passed.
EOH

  default = false
}

variable "vpc_id" {
  description = <<EOH
    VPC to add the Security group in.
    Required unless security group ID is passed.
EOH

  default = false
}

variable "security_group_id" {
  description = <<EOH
    ID of the Security Group to add Rules to.
    Creates security group if none provided.
EOH

  default = false
}

variable "source_cidr_block" {
  description = <<EOH
    CIDR Block to allow traffic.
EOH
}

variable "use_ssh" {
  description = <<EOH
    Allow SSH Ingress (TCP Port 22)
EOH

  default = false
}

variable "use_http" {
  description = <<EOH
    Allow HTTP Ingress (TCP Port 80)
EOH

  default = false
}

variable "use_https" {
  description = <<EOH
    Allow HTTPS Ingress (TCP Port 443)
EOH

  default = false
}

variable "use_jmx" {
  description = <<EOH
    Allow JMX Ingress (TCP Port 7199)
EOH

  default = false
}

variable "use_http_alt" {
  description = <<EOH
    Allow HTTP Alt Ingress (TCP Port 8080)
EOH

  default = false
}

variable "use_mysql" {
  description = <<EOH
    Allow MySQL Ingress (TCP Port 3306)
EOH

  default = false
}

variable "use_postgres" {
  description = <<EOH
    Allow Postgres Ingress (TCP Port 5432)
EOH

  default = false
}
