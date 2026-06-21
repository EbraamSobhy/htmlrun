# HTMLRun

A lightweight CLI tool written in Go that serves a local HTML project and makes it accessible through a browser.

Perfect for quickly previewing static websites, prototypes, and frontend experiments without installing Node.js or setting up a development server.

## Features

* Instant local web server
* Serves `index.html` and related assets
* Lightweight and fast
* Zero configuration
* Works with HTML, CSS, JavaScript, images, and other static files
* Cross-platform (Windows, macOS, Linux)

---

## Installation

### Download CLI
*   [Linux (x64)](portfolio-cli-linux)
*   [Windows (x64)](portfolio-cli.exe)
*   [macOS (Intel)](portfolio-cli-macos)
*   [macOS (Apple Silicon)](portfolio-cli-macos-arm64)
### Linux (x64)

```bash
chmod +x htmlrun-linux-amd64
sudo mv htmlrun-linux-amd64 /usr/local/bin/htmlrun
```

### macOS (Intel)

```bash
chmod +x htmlrun-macos-amd64
sudo mv htmlrun-macos-amd64 /usr/local/bin/htmlrun
```

### macOS (Apple Silicon)

```bash
chmod +x htmlrun-macos-arm64
sudo mv htmlrun-macos-arm64 /usr/local/bin/htmlrun
```

### Windows (x64)

```text
htmlrun-windows-amd64.exe
```

Optionally add the executable directory to your system PATH.

---
## Usage

Place the executable anywhere in your PATH and run:

```bash
htmlrun index.html
```

### Output:

```bash
index.html http://localhost:8080

```

---

## Why HTMLRun?

Many frontend developers simply need a quick way to preview an HTML file without installing large toolchains or frameworks.

HTMLRun provides:

* No Node.js dependency
* No package installation
* No configuration files
* Minimal resource usage
* Fast startup times

---

## License

MIT License

Feel free to use, modify, and distribute this project.

