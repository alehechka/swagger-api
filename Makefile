prepare:
	go install github.com/go-courier/husky/cmd/husky@latest
	go install github.com/swaggo/swag/cmd/swag@v1.8.2
	go install golang.org/x/tools/cmd/goimports@latest
	go mod download
	husky init
	swag init

build:
	swag init
	go build main.go