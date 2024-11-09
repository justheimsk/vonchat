TARGET = main
BUILD_DIR = build
SRC = cmd/api/main.go

all: $(TARGET)

$(TARGET): $(BUILD_DIR)
	go build -v -o $(BUILD_DIR)/$(TARGET) $(SRC)

$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

run:
	SQLITE_PATH=./sqlite.db $(BUILD_DIR)/$(TARGET)
