# HTMLRun

**HTMLRun** is a lightweight, cross-platform command-line tool that instantly serves local HTML projects through a web browser.

Designed for developers, designers, and students, HTMLRun provides a fast way to preview static websites without installing Node.js, configuring build tools, or setting up a development environment.

---

## Features

* 🚀 Instant local web server
* 📄 Serves HTML, CSS, JavaScript, images, fonts, and other static assets
* ⚡ Lightweight and fast
* 🔧 Zero configuration required
* 🌐 Opens projects through a local browser URL
* 💻 Cross-platform support (Windows, macOS, Linux)
* 📦 Single executable with no external dependencies

---

## Installation

Download the latest binary for your operating system from the project's Releases page.

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

Download:

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

Output:

```text
Serving: index.html

HTML Preview:
http://localhost:8080/index.html
```

Open the displayed URL in your browser to preview the project.

---

## Example Project

```text
project/
├── index.html
├── styles.css
├── script.js
└── assets/
    └── logo.png
```

Run:

```bash
cd project
htmlrun index.html
```

HTMLRun automatically serves all files within the project directory.

---

## Why HTMLRun?

Modern frontend tooling can be excessive for simple static projects.

HTMLRun is ideal when you need to:

* Preview a static website quickly
* Test HTML/CSS prototypes
* Review frontend assignments
* Share local demos during development
* Serve landing pages without additional tooling

Unlike many alternatives, HTMLRun requires:

* No Node.js
* No npm packages
* No configuration files
* No framework setup

Simply run the executable and start previewing your project.

---

## Supported File Types

HTMLRun serves all static assets, including:

* HTML
* CSS
* JavaScript
* PNG, JPG, SVG, WebP
* Fonts
* JSON
* PDF
* Any other static file located within the project directory

---

## Building from Source

Requirements:

* Go 1.22+

Clone the repository:

```bash
git clone https://github.com/yourusername/htmlrun.git
cd htmlrun
```

Build:

```bash
go build -o htmlrun
```

Run:

```bash
./htmlrun index.html
```

---

## Roadmap

Planned features:

* Custom port support
* Automatic browser launch
* Directory serving mode
* Live reload support
* HTTPS local server option

---

## License

MIT License

Copyright (c) 2026

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, subject to the terms of the MIT License.

