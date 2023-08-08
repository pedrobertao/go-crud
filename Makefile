# Start
run: 
	go run main.go

# Build
build:
	CGO_ENABLED=0 GOOS=linux go build -o go-crud main.go

# Test
test: 
	go test -v ./...