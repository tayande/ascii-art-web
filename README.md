# ASCII Art Web

A simple web application that converts text into ASCII art, built in Go. This project extends the original ASCII Art CLI tool by wrapping it in an HTTP server with a browser-based interface.

## Features
- Convert any text input into ASCII art directly from the browser
- Multiple banner styles (standard, shadow, thinkertoy)
- Clean, single-page interface with form-based input
- Server-side error handling for invalid input and unsupported characters

## How It Works
The app takes user input via an HTML form, passes it through the core ASCII art pipeline (`LoadBanner`, `ValidateInput`, `GenerateArt`), and renders the result back on the same page using Go's `html/template` package.

## Tech Stack
- Go (`net/http`, `html/template`)
- HTML/CSS (single template with conditional rendering)

## Running Locally
```bash
git clone https://github.com/your-username/ascii-art-web.git
cd ascii-art-web
go run main.go
```
Then open `http://localhost:8080` in your browser.

## Usage
1. Enter your text in the input field.
2. Choose a banner style.
3. Click submit to generate and view the ASCII art output.