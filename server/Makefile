TARGET = main
BUILD_DIR = build
SRC = cmd/api/main.go

all: $(TARGET)

$(TARGET): $(BUILD_DIR)
	go build -v -o $(BUILD_DIR)/$(TARGET) $(SRC)

$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

run:
	DATABASE_DRIVER=postgres POSTGRES_HOST=localhost POSTGRES_PORT=5432 POSTGRES_DB=main POSTGRES_USER=root POSTGRES_PASSWORD=admin DEBUG=true $(BUILD_DIR)/$(TARGET)
