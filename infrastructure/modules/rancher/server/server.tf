resource "aws_instance" "rancher-server" {
  ami = "${var.ami}"
  instance_type = "m4.large"
  key_name = "rancher-server"
  security_groups = [
    "${aws_security_group.rancher.id}",
    "${aws_security_group.db_access.id}"]
  associate_public_ip_address = true
  user_data = "${data.template_file.rancher_server_user_data.rendered}"
  root_block_device {
    volume_type = "gp2"
    volume_size = "32"
  }
}

resource "aws_security_group" "rancher" {
  name = "rancher_server_internal"
  description = "Allow traffic to rancher instances"
  vpc_id = "${var.vpc_id}"

  // HTTP
  ingress {
    from_port = 80
    to_port = 80
    protocol = "tcp"
    self = true
    cidr_blocks = [
      "${var.vpc_private_subnet}",
      "${var.vpc_public_subnet}"
    ]
  }
  ingress {
    from_port = 22
    to_port = 22
    protocol = "tcp"
    cidr_blocks = [
      "${var.vpc_private_subnet}",
      "${var.vpc_public_subnet}"
    ]
  }
  ingress {
    from_port = 443
    to_port = 443
    protocol = "tcp"
    cidr_blocks = [
      "${var.vpc_private_subnet}",
      "${var.vpc_public_subnet}"
    ]
  }
  ingress {
    from_port = 8080
    to_port = 8080
    protocol = "tcp"
    cidr_blocks = [
      "${var.vpc_private_subnet}",
      "${var.vpc_public_subnet}"
    ]
  }

  //Rancher IPSEC
  ingress {
    from_port = 500
    to_port = 500
    protocol = "udp"
    self = true
  }
  egress {
    from_port = 500
    to_port = 500
    protocol = "udp"
    self = true
  }
  ingress {
    from_port = 4500
    to_port = 4500
    protocol = "udp"
    self = true
  }
  egress {
    from_port = 4500
    to_port = 4500
    protocol = "udp"
    self = true
  }
}

data "template_file" "rancher_server_user_data" {
  template = "${file("${path.module}/rancher_user_data.tmpl")}"
  vars {
    image = "rancher/server:v1.1.4"
    rdsUrl = "${aws_db_instance.rancher_db.address}"
    rdsPass = "b1w0rldw1d3"
  }
}

//IAM
resource "aws_iam_instance_profile" "rancher" {
  name = "rancher-profile"
  roles = ["${aws_iam_role.rancher.name}"]
}

resource "aws_iam_role" "rancher" {
  name = "rancher-role"
  path = "/"
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
      {
          "Action": "sts:AssumeRole",
          "Principal": {
             "Service": "ec2.amazonaws.com"
          },
          "Effect": "Allow",
          "Sid": ""
      }
  ]

}
EOF
}

//Route 53 Entry
/*resource "aws_route53_record" "rancher_dns" {
  zone_id = "${var.r53_zone_id}"
  name = "rancher.${var.domain}"
  type = "A"
  records = ["${var.public_ip}"]
}*/

//RDS Config
resource "aws_db_instance" "rancher_db" {
  identifier = "rancher-db"
  allocated_storage = 10
  engine = "mysql"
  engine_version = "5.6.23"
  instance_class = "db.t2.small"
  storage_type = "gp2"
  name = "rancher"
  username = "biroot"
  password = "b1w0rldw1d3"
  multi_az = false
  port = 3306
  backup_retention_period = 7
  backup_window = "02:00-03:00"
  maintenance_window = "sun:03:30-sun:04:30"
  publicly_accessible = false
  apply_immediately = false
  vpc_security_group_ids = [
    "${aws_security_group.db_access.id}"
  ]
  db_subnet_group_name = "${aws_db_subnet_group.db_subnet.id}"
}

resource "aws_security_group" "db_access" {
  name = "rancher_db_access"
  description = "Access to Rancher RDS db"
  ingress {
    from_port = 3306
    to_port = 3306
    protocol = "tcp"
    self = true
  }
  vpc_id = "${var.vpc_id}"
}

resource "aws_db_subnet_group" "db_subnet" {
  name = "rancher-dbs"
  description = "VPC DB subnet group for Rancher"
  subnet_ids = ["${var.vpc_subnet_ids}"]
}