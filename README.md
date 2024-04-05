# GoTTH Stack

This is a GoTTH Stack boilerplate project to quickly spin-up a new Golang Web App with Database support.

This project is inspired from [TomDoesTech's](https://github.com/TomDoesTech) [GoTTH](https://github.com/TomDoesTech/GOTTH) stack, but with different pattern and an Integrated LibSQL (SQLite) Database Support.

## About the Stack

The GoTTH stack includes the following Languages and Technologies:

1. **Go** with Chi Router.
2. **TailwindCSS** Standalone Executable.
3. **Turso** LibSQL (SQLite for Production).
4. **HTMX** with Templ as Templating Engine.

## Installation Guide

### Chi Router

```bash
go get -u github.com/go-chi/chi/v5
```

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
