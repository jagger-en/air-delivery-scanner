##
## Build stage
##
FROM golang:1.19-buster AS build

WORKDIR /app
COPY go.mod ./
COPY main.go ./
COPY ./web ./web
RUN go build -o /airdelscn-server

##
## Use distroless image: https://github.com/GoogleContainerTools/distroless
##
FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /airdelscn-server /airdelscn-server
COPY --from=build /app/web /web

EXPOSE 7070

USER nonroot:nonroot

ENTRYPOINT ["/airdelscn-server"]
