variable "github_token" {}

provider "github" {
  token        = "${var.github_token}"
  organization = ""
}

resource "github_repository" "terraform-mydns" {
  name        = "terraform-mydns"
  description = "A plugin for terraform"
}
