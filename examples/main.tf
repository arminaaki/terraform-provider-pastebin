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

resource "pastebin_create_api" "pasteA" {
  api_dev_key           =  "" #'2cddd4dab41f10754e9dfd5dd6f9fbbf'; // your api_developer_key
  api_paste_code        =  "" #'just some random text you :)'; // your paste text
  api_paste_private     =  "" #'1'; // 0=public 1=unlisted 2=private
  api_paste_name        =  "" #'justmyfilename.php'; // name or title of your paste
  api_paste_expire_date =  "" #'10M';
  api_paste_format      =  "" #'php';
  api_user_key          =  "" #''; // if an invalid or expired api_user_key is used, an error will spawn. If no api_user_key is used, a guest paste will be created
  api_paste_name        =  "" #urlencode($api_paste_name);
  api_paste_code        =  "" #urlencode($api_paste_code);
}
