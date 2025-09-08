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
  go build -o app ./...

FROM alpine:edge AS runner
WORKDIR /root

EXPOSE 8080

COPY --from=build /root/app /root/app

CMD ["/root/app"]
