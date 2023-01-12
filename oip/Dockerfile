FROM golang:1.19-alpine as build

RUN apk add --no-cache git curl

# OIP daemon from oip5 branch 2023/01/12
ARG OIP_VERSION="5e1dc8c920ee5659295a5aee8a8a41534565ea7d"
RUN git clone https://github.com/bitspill/verbose-umbrella /oip
WORKDIR /oip
RUN git checkout ${OIP_VERSION}
# OIP daemon from main branch
# RUN git clone https://github.com/oipwg/oip /oip
# WORKDIR /oip

RUN go mod download

RUN go build -ldflags \
        "-X github.com/oipwg/oip/version.GitCommitHash=$(git rev-parse --short HEAD) \
        -X github.com/oipwg/oip/version.BuildDate=$(date +'%Y.%m.%d.%H%M%S') \
        -X github.com/oipwg/oip/version.BuiltBy=Docker \
        -X 'github.com/oipwg/oip/version.GoVersion=$(go version)' \
        -s -w" -o oipd github.com/oipwg/oip/cmd/oipd

FROM alpine:latest
ARG network=mainnet

COPY --from=build /oip/oipd /oip/bin/oipd
COPY ./config.${network}.yml /oip/config.yml

EXPOSE 1606
ENV LOG=*
ENTRYPOINT ["/oip/bin/oipd", "--appdir=/oip"]
