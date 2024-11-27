prereq:
	go install github.com/go-critic/go-critic/cmd/gocritic@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/mathaou/termdbms@latest
