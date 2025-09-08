FROM alpine:edge AS build
WORKDIR /root

RUN apk upgrade; \
    apk add go

RUN go env -w GOCACHE=/go-cache
RUN go env -w GOMODCACHE=/gomod-cache

COPY ./go.mod ./
RUN --mount=type=cache,target=/gomod-cache \
  go mod download

COPY ./ ./
RUN --mount=type=cache,target=/gomod-cache --mount=type=cache,target=/go-cache \
  go build -o out/ ./...

FROM alpine:edge AS runner
WORKDIR /root

RUN apk upgrade; apk add zellij helix helix-tree-sitter-vendor

EXPOSE 8080

COPY --from=build /root/out/* /root/

CMD ["/root/app"]
