provider "aws" {
  region = "us-west-2"
  profile = "bi-worldwide"
}

data "terraform_remote_state" "remote-state" {
  backend = "s3"
  config {
    bucket = "biw-tsgsandbox-hack"
    key = "ahaynssen/terraform.tfstate"
    region = "us-west-2"
  }
}
