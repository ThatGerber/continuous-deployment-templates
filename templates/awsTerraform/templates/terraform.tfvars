{{if (eq .Var.tfBackend "s3")}}
bucket = "{{.Var.tfStateBucket}}"
key    = "infrastructure/terraform.tfstate"
region = "{{.Var.tfStateRegion}}"


{{else}}path = ".terraform/terraform.tfstate"
{{end}}
