FROM golang:1-alpine as build

RUN apk add --no-cache git

RUN go get -u github.com/Masterminds/glide
RUN git clone https://github.com/bitspill/flod /go/src/github.com/bitspill/flod
WORKDIR /go/src/github.com/bitspill/flod
RUN glide install
RUN go install . ./cmd/...


FROM alpine:latest
ARG network=mainnet

COPY --from=build /go/bin/flod /flod/bin/flod
COPY ./${network}.conf /flod/flod.conf

VOLUME ["/root/.flod/"]
ENTRYPOINT ["/flod/bin/flod", "--configfile=/flod/flod.conf"]
