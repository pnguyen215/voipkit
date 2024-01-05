.PHONY: run build test tidy deps-upgrade deps-clean-cache

# ==============================================================================
# Start Main
run:
	go run main.go

build:
	go build main.go

# ==============================================================================
# Modules support
test:
	go test -cover ./...

tidy:
	go mod tidy
	go mod vendor

deps-upgrade:
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	go get -u -t -d -v ./...
	go mod tidy
	go mod vendor

deps-clean-cache:
	go clean -modcache
