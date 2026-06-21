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

### Clone the Repository

```bash
git clone https://github.com/EbraamSobhy/htmlrun.git
cd htmlrun
```

### Build

```bash
go build -o htmlrun
```

---

## Usage

Place your `index.html` file in a directory and run:

```bash
./htmlrun
```

The server will start on:

```text
http://localhost:8080
```

Example output:

```text
Serving: http://localhost:8080
```

Open the URL in your browser to view the project.

---

## Project Structure

```text
project/
├── index.html
├── style.css
├── script.js
└── assets/
```

Example:

```html
<!DOCTYPE html>
<html>
<head>
    <title>Hello World</title>
</head>
<body>
    <h1>Hello from HTMLRun!</h1>
</body>
</html>
```

Running:

```bash
./htmlrun
```

will serve the entire directory.

---

## Requirements

* Go 1.20 or newer

Check your version:

```bash
go version
```

---

## Current Behavior

The application:

1. Uses the current directory as the web root.
2. Serves all static files.
3. Listens on port `8080`.
4. Makes the project available at:

```text
http://localhost:8080
```

---

## Example

Start the server:

```bash
./htmlrun
```

Visit:

```text
http://localhost:8080
```

You should see your `index.html` page rendered in the browser.

---

## Future Roadmap

Planned features:

* Automatic browser opening
* Custom port support
* File watching and live reload
* Single-file mode
* HTTPS support
* Configuration file support
* Colored terminal output

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

