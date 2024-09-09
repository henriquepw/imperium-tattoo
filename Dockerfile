FROM golang:alpine AS builder

WORKDIR /app
COPY . .
RUN wget -O tailwind https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.1/tailwindcss-linux-x64 && \
    chmod +x tailwind && ./tailwind -i ./static/css/input.css -o ./static/css/output.css --minify

RUN go install github.com/a-h/templ/cmd/templ@latest && templ generate

RUN go build -o main ./cmd/main.go

FROM alpine:latest
WORKDIR /app

COPY --from=builder app/static/ ./static/
COPY --from=builder app/main .
CMD ["./main"]
