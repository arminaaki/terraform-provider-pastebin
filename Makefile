
PLUGINS_DIR := $$HOME/.terraform.d/plugins
DIST_DIR    := pkg
BINARY_NAME := terraform-provider-pastebin
VERSION     := v0.1.0

lintcheck:
	- npm install -g markdownlint-cli
	- markdownlint README.md

build:
	mkdir -p $(PLUGINS_DIR)
	go build -o $(PLUGINS_DIR)/$(BINARY_NAME)

test-dep:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint
	GO111MODULE=on go mod vendor
test: test-dep
	golangci-lint run


release:
	mkdir -p $(DIST_DIR)
	go get github.com/mitchellh/gox
	go get github.com/tcnksm/ghr
	gox --parallel=10 -os="linux darwin " -arch="amd64 386" -output="$(DIST_DIR)/${BINARY_NAME}-$(VERSION)-{{.OS}}-{{.Arch}}" .
	@ghr --username  arminaaki --token $(GITHUB_TOKEN) --replace --prerelease --debug $(VERSION) $(DIST_DIR)/

clean:
	rm -rf $(PLUGINS_DIR)/$(BINARY_NAME)

