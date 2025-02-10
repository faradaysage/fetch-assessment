# default port for the server
export PORT=8081

# Declare phony targets to ensure they always run.
.PHONY: generate format tidy build run test clean

# generate: Runs the code generation command using oapi-codegen.
generate:
	@echo "Generating API code from api.yml..."
	mkdir -p api
	go generate ./tools/tools.go

# format: Formats your Go source code using gofumpt.
# You can change this to "gofmt" if preferred.
format:
	@echo "Formatting code..."
	gofmt -l -w .

# tidy: Runs go mod tidy to clean up go.mod and go.sum.
tidy:
	@echo "Tidying dependencies..."
	go mod tidy

# build: Compiles the project into a binary.
build: tidy
	@echo "Building binary..."
	go build -o bin/fetch-assessment ./cmd/fetch-assessment

# run: Builds (if needed) and runs the project.
run: tidy
	@echo "Starting application..."
	go run ./cmd/fetch-assessment

# test: Runs all the tests in the project.
test: tidy
	@echo "Running tests..."
	go test ./...

# clean: Removes built artifacts.
clean:
	@echo "Cleaning up..."
	rm -rf bin
