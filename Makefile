
PLUGINS_DIR = $$HOME/.terraform.d/plugins

build:
	mkdir -p $(PLUGINS_DIR)
	go build -o $(PLUGINS_DIR)/terraform-provider-pastebin

clean:
	rm -rf $(PLUGINS_DIR)/terraform-provider-pastebin
