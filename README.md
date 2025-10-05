# Go-Templ-Fiber-Tailwind 🌠 🦫

#### A modern blog application built with a powerful Go stack, designed for speed and simplicity. This project serves blog posts written in Markdown from a local JSON file, providing a fast, server-rendered experience with a clean, responsive design.

## Features 🌟

  - Markdown-Powered: Blog posts are written in simple Markdown and stored in a allposts.json file.

  - Server-Side Rendering: Uses Templ for type-safe, component-based HTML rendering directly on the server.

  - High-Performance Backend: Built on Fiber, a Go web framework designed for high performance and low memory allocation.

  - Beautifully Styled: Styled with Tailwind CSS for a modern, utility-first design that is fully responsive.

  - Dynamic Content:

  - An archive page that lists all available blog posts.

  - Dynamic routes to render individual blog posts.

  - Fast Markdown Parsing: Leverages [Goldmark](https://www.github.com/yuin/goldmark) for extremely fast and extensible Markdown-to-HTML conversion.

## Tech Stack 🥞

- Backend: Go
- Web Framework: Fiber
- Templating: Templ
- Styling: Tailwind CSS
- Markdown Parser: Goldmark

## Screenshots 📸 
<div class="flex">
  <img width="400" height="360" alt="Home Page View" src="https://github.com/user-attachments/assets/08b92884-2585-40ed-ae17-169f1ad8146c" />
  <img width="400" height="380" alt="Blog Post View" src="https://github.com/user-attachments/assets/dbe014db-6910-4786-bf26-bdc6c3e8e536" />
</div>

<br>

## Project Structure 🛤️

```plaintext
.
├── handlers/         # Fiber request handlers
├── models/           # Go struct definitions (e.g., Post)
├── static/           # Static assets (CSS, images)
├── templates/        # Templ components (.templ files)
├── allposts.json     # Data source for blog posts
├── go.mod            # Go module dependencies
└── main.go           # Application entry point
```


## Getting Started

Follow these instructions to get a copy of the project up and running on your local machine.

## Prerequisites

#### Make sure you have the following installed:

- Go (version 1.21 or later)

- Templ CLI

- Tailwind CSS CLI (via npm)

<br>

## Installation 🤖

Clone the repository:

```bash
git clone https://github.com/your-username/templ-go.git
cd templ-go
```

Generate Templ Components:

The templ CLI watches for changes and generates Go code from your .templ files.

```bash
templ generate --watch
```
Install Go dependencies:

```bash
go mod tidy
```
Build Tailwind CSS:

##### Run the Tailwind CLI to build your input.css or just use the CDN 🌪️

```bash
npx tailwindcss -i ./static/css/input.css -o ./static/css/output.css --watch
```

##### CDN ⬇️
```html
<script>https://cdn.tailwindcss.com</script>
```

<br>

### Run the application: 🚀

Open a new terminal window and run the main Go program.

```
go run .
```
The server will be running on http://localhost:8080.



### Screenshots

<img width="764" height="802" alt="Blog Homepage" src="https://github.com/user-attachments/assets/7c325232-87a7-4f5d-a916-09fe87fcb251" />

<img width="764" height="802" alt="Blog Post View" src="https://github.com/user-attachments/assets/dbe014db-6910-4786-bf26-bdc6c3e8e536" />

License

This project is licensed under the MIT License - see the LICENSE file for details.
