FROM golang:1.16-alpine as build

ENV GO11MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    BASE_PATH=/go/src/app

WORKDIR $BASE_PATH

COPY . ./

RUN go mod download
RUN go build -o main .


FROM scratch as run

ENV BASE_PATH=/go/src/app

COPY --from=build $BASE_PATH/main /app/main

EXPOSE 8080

ENTRYPOINT ["/app/main"]
