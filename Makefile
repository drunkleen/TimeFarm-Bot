BUILD_DIR = ./bin
APP_NAME = "TimeFarmBot"

build: build-windows-arm build-windows-amd64 build-linux-arm build-linux-amd64
	@mkdir "$(BUILD_DIR)/configs"
	@touch "$(BUILD_DIR)/configs/query.conf"

build-windows-arm:
	@GOOS=windows GOARCH=arm go build -o $(BUILD_DIR)/$(APP_NAME)-windows-arm.exe

build-windows-amd64:
	@GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)-windows-x86_64.exe

build-linux-arm:
	@GOOS=linux GOARCH=arm go build -o $(BUILD_DIR)/$(APP_NAME)-linux-arm

build-linux-amd64:
	@GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)-linux-x86_64

clean:
	@rm -rf $(BUILD_DIR)
#
#run: build
#	@./bin/blumbot


run: build
	@sh -c '\
		OS=$$(uname -s | tr A-Z a-z); \
		ARCH=$$(uname -m); \
		case "$${OS}" in \
			linux) \
				case "$${ARCH}" in \
					x86_64) \
						echo "Running Linux AMD64 binary..."; \
						$(BUILD_DIR)/$(APP_NAME)-linux-x86_64 ;; \
					arm*|aarch64) \
						echo "Running Linux ARM binary..."; \
						$(BUILD_DIR)/$(APP_NAME)-linux-arm ;; \
				esac ;; \
			darwin) \
				case "$${ARCH}" in \
					x86_64) \
						echo "Running macOS AMD64 binary..."; \
						$(BUILD_DIR)/$(APP_NAME)-darwin-x86_64 ;; \
					arm*|aarch64) \
						echo "Running macOS ARM binary..."; \
						$(BUILD_DIR)/$(APP_NAME)-darwin-arm ;; \
				esac ;; \
			msys*|cygwin*|mingw*) \
				case "$${ARCH}" in \
					x86_64) \
						echo "Running Windows AMD64 binary..."; \
						$(BUILD_DIR)/$(APP_NAME)-windows-x86_64.exe ;; \
					arm*|aarch64) \
						echo "Running Windows ARM binary..."; \
						$(BUILD_DIR)/$(APP_NAME)-windows-arm.exe ;; \
				esac ;; \
			*) echo "Unsupported OS: $${OS}" ;; \
		esac'