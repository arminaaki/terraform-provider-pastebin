
PLUGINS_DIR = $$HOME/.terraform.d/plugins

build:
	mkdir -p $(PLUGINS_DIR)
	go build -o $(PLUGINS_DIR)/terraform-provider-pastebin

test-dep:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint
	GO111MODULE=on go mod vendor
test: test-dep
	golangci-lint run


clean:
	rm -rf $(PLUGINS_DIR)/terraform-provider-pastebin
