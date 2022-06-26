FROM tetafro/golang-gcc:latest AS dev

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o main
EXPOSE 8080
CMD ["./main"]