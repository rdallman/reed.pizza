# build stage
FROM golang:1.14.2-alpine3.11 AS build-env
ENV D=/go/src/github.com/rdallman/worldwidereed
ADD . $D
RUN cd $D && go build -o wwwreed

FROM alpine:3.11
WORKDIR /app
COPY --from=build-env /go/src/github.com/rdallman/worldwidereed /app/
ENTRYPOINT ["/app/wwwreed"]
EXPOSE 80 443
