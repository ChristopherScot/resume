FROM golang:1.22 as builder
WORKDIR ~/ChristopherScot/resume
COPY . . 
RUN go build -o resume 

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN apk --no-cache add curl
RUN rm -rf/var/cache/apk/*

RUN mkdir /app
COPY --from=builder /go/ChristopherScot/resume/resume /app
WORKDIR /app
CMD ["./resume"]