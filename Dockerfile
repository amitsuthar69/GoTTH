FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN curl -L https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64 -o /usr/local/bin/tailwindcss

RUN chmod +x /usr/local/bin/tailwindcss

RUN go install github.com/a-h/templ/cmd/templ@latest

RUN which templ && templ --version

RUN templ generate

RUN tailwindcss -i ./static/css/input.css -o ./static/css/style.css --minify

RUN go build -o main ./main.go

EXPOSE 8080

CMD ["./main"]