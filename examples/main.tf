variable "api_id" {
  type = string
}

variable "api_key" {
  type = string
}

provider "victorops" {
  api_id  = var.api_id
  api_key = var.api_key
}

resource "victorops_user" "user_1" {
  first_name = "Test"
  last_name = "Terraform"
  username = "test_terraform"
  email = "test-terraform@omio.com"
}

resource "victorops_team" "team_1" {
  name = "terraform-test-team"
}
