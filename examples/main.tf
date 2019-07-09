provider "pastebin" {
  api_dev_key       = "${var.api_dev_key}"
  api_user_name     = "${var.api_user_name}"
  api_user_password = "${var.api_user_password}"
}

resource "pastebin_api_user_key" "api_key" {
  name = "my_key"
}

output "API KEY" {
  value = "${pastebin_api_user_key.api_key.api_user_key}"
}

resource "pastebin_create_paste" "pasteA" {
  name                  = "pasteA"
  api_user_key          = "${pastebin_api_user_key.api_key.api_user_key}"
  api_dev_key           = "${var.api_dev_key}"
  api_paste_code        = "puts 'Hello World'"
  api_paste_name        = "main.rb"
  api_paste_expire_date = "10M"
  api_paste_format      = "ruby"
  api_paste_private     = "2"
}
