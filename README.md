# Installing `func` CLI on macOS and Ubuntu

The `func` CLI tool is essential for working with Knative functions. This guide provides installation instructions for macOS and Ubuntu.

## Installation Steps

### macOS
1. Install the Knative client:
   ```bash
   brew install knative/client/kn
2. Install the Knative kn-plugins:
   ```bash
   brew tap knative-extensions/kn-plugins
3. Install the func:
   ```bash
   brew install func


### ubuntu
1. Download the func binary for Ubuntu:
   ```bash
   wget https://github.com/knative/func/releases/download/knative-v1.16.1/func_linux_amd64
2. Make the binary executable:
   ```bash
   chmod +x func_linux_amd64
3. Move the binary to a directory in your PATH:
   ```bash
   sudo mv func_linux_amd64 /usr/local/bin/func
