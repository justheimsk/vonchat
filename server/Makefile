TARGET = main
BUILD_DIR = build
SRC = cmd/api/main.go

all: $(TARGET)

$(TARGET): $(BUILD_DIR)
	go build -v -o $(BUILD_DIR)/$(TARGET) $(SRC)

$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

lint:
	$(shell go env GOPATH)/bin/revive -formatter friendly ./...

fmt:
	go fmt ./...

run:
	$(BUILD_DIR)/$(TARGET)
