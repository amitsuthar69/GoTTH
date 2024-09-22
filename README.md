# GoTTH Stack

The **GoTTH Stack** is designed to quickly spin up a Full-Stack Golang Web Application with database support.

## About the Stack

Includes following Languages and Technologies:

1. **Golang** and **default ServeMux**.
2. **TailwindCSS** Standalone Executable.
3. **Turso** LibSQL (SQLite for Production).
4. **HTMX** with **Templ** as Templating Engine.

## Run the project

```
air
```

## Installation Guide

### Air for Live Reload

```bash
go install github.com/cosmtrek/air@latest
```

### Templ

```bash
go install github.com/a-h/templ/cmd/templ@latest
```

You can add `templ` to your local bashrc or zshrc profile config to avoid absolute path while using `templ` command.

### TailwindCSS

1. Install the executable binary for your platform from the [latest release](https://github.com/tailwindlabs/tailwindcss/releases/latest) from GitHub in the root of project.

2. Make sure to give it executable permissions:

```bash
chmod +x ./tailwindcss
```

### Turso

```bash
go get github.com/tursodatabase/libsql-client-go/libsql
```

1. Then [Install the Turso CLI](https://docs.turso.tech/quickstart), create your database, generate the AuthToken & DB_URL and add them in the `.env` file.
2. You can refer this awesome Turso [Quickstart Go Docs](https://docs.turso.tech/sdk/go/quickstart) if you get stuck anywhere.

## .env

```
PORT=8080
DB_URL=libsql://yourDBName-tursoUserName.turso.io?authToken=yourAuthToken
```

## Basic Project Guide

```
GOTTH
│
├── /internal/                 # Private application and business logic
│   ├── /server/               # Server-specific logic
│   │   ├── server.go          # Core server logic
│   │   └── routes.go          # Route definitions
│   └── /database/             # Database-related logic
│       └── database.go        # Database connection and queries
│
├── /static/                   # Static assets (CSS, JS, images, etc.)
│   ├── /css/                  # CSS files
│   │   ├── input.css          # Tailwind input file
│   │   └── style.css          # Generated CSS file
│   └── /scripts/              # JavaScript files
│       └── htmx.min.js        # HTMX JavaScript library
│
├── /web/                      # Web UI-related handlers
│       ├── hello.templ        # Templating file for the hello page
│       └── hello.go           # Hello handler logic
│
├── main.go                    # Application entry point
├── Makefile                   # Build automation script
├── .air.toml                  # Air live reloading config file
├── tailwind.config.js         # Tailwind CSS configuration
└── tailwindcss                # Tailwind CSS binary or related file
```

---

Inspired from [TomDoesTech's](https://github.com/TomDoesTech) [GoTTH](https://github.com/TomDoesTech/GOTTH) stack, but with a simpler folder structure and an Integrated Turso LibSQL Database Support.
