# Define the output binary names
BINARY_NAME=shr
LINUX_BINARY=$(BINARY_NAME)_linux
WINDOWS_BINARY=$(BINARY_NAME).exe

# Default target
all:
	@echo "Building for all platforms..."
	make build-linux
	make build-windows

# Build for Linux
build-linux:
	@echo "Building for Linux..."
	GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o $(LINUX_BINARY) .

# Build for Windows
build-windows:
	@echo "Building for Windows..."
	GOOS=windows GOARCH=amd64 go build -ldflags "-w -s" -o $(WINDOWS_BINARY) .

# Clean up build artifacts
clean:
	rm -f $(LINUX_BINARY) $(WINDOWS_BINARY)

# Install for linux
install-linux:
	@echo "Installing for Linux..."
	make build-linux
	chmod +x $(LINUX_BINARY)
	sudo cp $(LINUX_BINARY) /usr/local/bin/$(BINARY_NAME)

# Uninstall for linux
uninstall-linux:
	@echo "Uninstalling for Linux..."
	sudo rm -f /usr/local/bin/$(BINARY_NAME)

.PHONY:
	all build-linux build-windows clean
