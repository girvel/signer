FROM golang:1.24 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY main.go ./
COPY src ./src
COPY docs ./docs
RUN CGO_ENABLED=0 go build -o /bin/signer .

FROM scratch
COPY private.pem ./
COPY --from=build /bin/signer /bin/signer
CMD ["/bin/signer"]
