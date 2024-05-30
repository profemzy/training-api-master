FROM golang:1.22.3-alpine as dependencies

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

FROM dependencies AS build
COPY  . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/server/main.go

FROM golang:1.22.3-alpine
COPY --from=build /app .
CMD ["./app"]