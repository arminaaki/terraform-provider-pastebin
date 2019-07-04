provider "pastebin" {
  api_dev_key       = "${var.api_dev_key}"
  api_user_name     = "${var.api_user_name}"
  api_user_password = "${var.api_user_password}"
}

resource "pastebin_api_user_key" "api_key" {
  name = "NAME"
}

# output "API KEY" {
#   value = "${example_server.api_key.api_user_key}"
# }
