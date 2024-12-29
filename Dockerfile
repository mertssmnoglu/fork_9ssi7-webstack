FROM node:20.11.1-slim AS web_builder

WORKDIR /

COPY ./api/web ./

RUN npm install

RUN npm run build

FROM node:20.11.1-slim AS admin_builder

WORKDIR /

COPY ./api/admin ./

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
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go install github.com/a-h/templ/cmd/templ@latest

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
COPY --from=web_builder /static/dist /api/web/static/dist
COPY --from=web_builder /templates /api/web/templates
COPY --from=admin_builder /dist /api/admin/dist
EXPOSE $HTTP_PORT $RPC_PORT

CMD ["/main"]