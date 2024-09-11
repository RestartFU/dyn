# Dyn: The Dynamic Linux Package Manager

Welcome to Dyn, a modern, user-friendly package manager for Linux that's designed to make your life easier. Whether you're a seasoned sysadmin or a curious beginner, Dyn streamlines your package management experience with simplicity and efficiency.

## Installation

### Prerequisites
- Ensure you have **Go** installed on your system.

### From Source
```sh
git clone --depth=1 https://github.com/RestartFU/dyn
cd dyn
sudo make install
```

### From Releases
```sh
redirect=$(curl -w "%{redirect_url}" -o /dev/null -s "https://github.com/RestartFU/dyn/releases/latest/download/dyn")
download_url=$(curl -w "%{redirect_url}" -o /dev/null -s $redirect)
curl -o dyn $download_url
chmod +x dyn
sudo mv dyn /usr/bin/dyn
```

## Quick Start

### Install a Package
```sh
dyn install <package_name>
```
**Example:**
```sh
dyn install discord
```

### Remove a Package
```sh
dyn remove <package_name>
```
**Example:**
```sh
dyn remove discord
```

### Update a Package
```sh
dyn update <package_name>
```
**Example:**
```sh
dyn update discord
```

### Fetch Dyn's Package Repository
```sh
dyn fetch
```

### Fetch and Install in One Go
For convenience, you can fetch and install a package simultaneously:
```sh
dyn fetch install discord
```

## Why Choose Dyn?

- **Speed**: Cloned with `--depth=1` for faster setup.
- **Simplicity**: Straightforward commands for all your package needs.
- **Efficiency**: Updates, installs, and removes packages with minimal fuss.

## Advanced Usage (TODO)

Dyn supports a variety of options for more complex operations. Here's a quick overview:

- **List Installed Packages**:
  ```sh
  dyn list
  ```

- **Search for Packages**:
  ```sh
  dyn search <keyword>
  ```

- **Show Package Information**:
  ```sh
  dyn info <package_name>
  ```

## Contributing

We welcome contributions! If you have ideas or find bugs, please:

1. Fork the repository.
2. Make your changes.
3. Submit a pull request.
