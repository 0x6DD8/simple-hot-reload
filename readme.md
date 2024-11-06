# Simple Hot Reload

Simple Hot Reload (SHR) is a simple hot reload tool that watches a specified directory for file changes and automatically restarts a specified program when changes are detected.
Primarily build for quick cli java hot reload
## Features

- Watch a specified directory for file changes.
- Automatically restart a specified program when changes are detected.
- Should be able to run on any platform that supports Go. It has been tested only on Linux though.

## Installation

### Prerequisites

- Go 1.23.2 or later

### Build from Source

To build the Simple Hot Reload tool from source, follow these steps:

1. Clone the repository:

    ```sh
    git clone https://github.com/0x6DD8/simple-hot-reload.git
    cd simple-hot-reload
    ```

2. Build the binary for your platform:

    ```sh
    make build-linux   # For Linux
    make build-windows # For Windows
    ```

3. Optionally, you can move the binary to `/usr/local/bin` on Linux:

    ```sh
    sudo make install-linux
    ```

## Usage

To use the Simple Hot Reload, run the following command:

```sh
./shr <path-to-watch> <program-to-run> [args...]
```

### Example

```sh
./shr /path/to/project java main.java
```

This command will watch the `/path/to/project` directory for changes and automatically restart the `java main.java` command whenever a change is detected.

## Makefile Targets

- `make all`: Build the binary for all supported platforms.
- `make build-linux`: Build the binary for Linux.
- `make build-windows`: Build the binary for Windows.
- `make clean`: Clean up build artifacts.
- `make install-linux`: Install the Linux binary to `/usr/local/bin`.
- `make uninstall-linux`: Uninstall the Linux binary from `/usr/local/bin`.
