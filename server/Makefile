TARGET = main
BUILD_DIR = build
SRC = cmd/api/main.go

all: $(TARGET)

$(TARGET): $(BUILD_DIR) tidy
	go build -o $(BUILD_DIR)/$(TARGET) $(SRC)

$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

tidy:
	go mod tidy

lint:
	$(shell go env GOPATH)/bin/revive -formatter friendly ./...

fmt:
	go fmt ./...
