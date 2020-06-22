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

resource "victorops_team" "team_1" {
  name = "test-team"
}
