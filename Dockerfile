FROM golang:1.17 as build

#ENV GOPROXY https://goproxy.io
ENV GO111MODULE on

# WORKDIR /go/cache
# ADD go.mod .
# ADD go.sum .
# RUN go mod download

WORKDIR /go/release
ADD . .
#RUN go mod tidy
RUN GOOS=linux CGO_ENABLED=0 go build -o heco-monitor

FROM alpine as prod
RUN apk update && apk add tzdata
WORKDIR /app

COPY --from=build /go/release/heco-monitor /app/
COPY --from=build /go/release/conf /app/conf

CMD ["./heco-monitor"]