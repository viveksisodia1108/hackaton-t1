FROM dockerdeepak/workshop:golangbase

WORKDIR /app

COPY . .
RUN go mod download && cd cmd && CGO_ENABLED=0 GOOS=linux go build -o /transactions
EXPOSE 3000
CMD [ "/transactions" ]
