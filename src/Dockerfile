FROM golang:1.23rc2-alpine
WORKDIR /app
COPY go.* ./
COPY . .
EXPOSE 5000
CMD [ "go", "run", "main.go" ]