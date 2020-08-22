
PLUGINS_DIR := $$HOME/.terraform.d/plugins
DIST_DIR    := pkg
BINARY_NAME := terraform-provider-pastebin

build:
	mkdir -p $(PLUGINS_DIR)
	go build -o $(PLUGINS_DIR)/$(BINARY_NAME)

clean:
	rm -rf $(PLUGINS_DIR)/$(BINARY_NAME)

