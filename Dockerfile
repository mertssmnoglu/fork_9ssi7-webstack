### Node Stage
FROM node:20.11.1-slim AS node_builder

WORKDIR /

COPY ./api/web ./

RUN npm install

RUN npm run build

FROM golang:1.23-alpine AS builder
RUN apk update && apk add --no-cache ca-certificates

ENV CGO_ENABLED=0 GO111MODULE=on GOOS=linux

WORKDIR /

COPY go.* ./
RUN   --mount=type=cache,target=/go/pkg/mod \
    go mod download
COPY . .

# install templ
RUN go install github.com/a-h/templ/cmd/templ@latest

RUN templ generate ./api/web/templates/

COPY ./api/web/templates ./api/web/templates

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -o main ./cmd/main.go

FROM scratch

ENV HTTP_PORT=8080
ENV RPC_PORT=9090

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /main .
COPY --from=node_builder /static/dist /api/web/static/dist
COPY --from=node_builder /templates /api/web/templates
EXPOSE $HTTP_PORT $RPC_PORT

CMD ["/main"]