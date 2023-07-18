FROM dockerdeepak/workshop:golangbase

WORKDIR /app

COPY . .
RUN ls && go mod download && go mod tidy
ENTRYPOINT ["go", "run", "cmd/main.go", "-b", "0.0.0.0"]
