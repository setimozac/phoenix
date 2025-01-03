##STEP1
FROM golang:1.22 AS builder

WORKDIR /app

# copy the project dependencies
COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY cmd/main.go cmd/main.go
COPY internal/ internal/
COPY api/ api/

# build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o env-manager cmd/main.go


##STEP2
FROM gcr.io/distroless/static:nonroot

WORKDIR /

# copy the binary to the final image
COPY --from=builder /app/env-manager .

# start the operator
ENTRYPOINT [ "/env-manager" ]


