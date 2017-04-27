variable "github_token" {}

provider "github" {
  token        = "${var.github_token}"
  organization = ""
}

resource "github_repository" "terraform-provider-mydns" {
  name        = "terraform-provider-mydns"
  description = "A plugin for terraform"
}
