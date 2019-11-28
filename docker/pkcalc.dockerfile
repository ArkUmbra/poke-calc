# Referenced from https://codinglatte.com/posts/golang/golang-small-docker-images/

FROM golang as builder

WORKDIR /go/src/github.com/arkumbra/pkcalc

COPY . .

RUN go get .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# deployment image
FROM alpine:3.8

# copy ca-certificates from builder
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

#WORKDIR /bin/

COPY --from=builder /go/src/github.com/arkumbra/pkcalc/ .
RUN sh build.sh

CMD [ "./app" ]

EXPOSE 8080