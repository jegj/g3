install:
	go mod download
run:
	go run main.go
test:
	go test -v ./...
build: main.go
	go build -ldflags="-X 'main.Version=v1.0.0'" -o nvcly $<
vet:
	go vet main.go
vet_shadow:
	go vet -vettool=$(which shadow) main.go
staticcheck:
	staticcheck ./...
audit:
	govulncheck -mode binary -show verbose nvcly
gosec:
	gosec ./...
tidy:
	go mod tidy
