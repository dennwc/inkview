FROM 5keeve/pocketbook-sdk:6.3.0-b288-v1

LABEL maintainer="https://github.com/Skeeve"

ARG GOLANG_VERSION=17.2

RUN apt-get update && \
    apt-get install -y xz-utils git curl && \
    apt-get clean

# download Pocketbook SDK
ADD ./patches/* /tmp/

# download specified Go binary release that will act as a bootstrap compiler for Go toolchain
# download sources for that release and apply the patch
# build a new toolchain and remove an old one
RUN mkdir /gosrc \
 && curl https://dl.google.com/go/go1.17.2.linux-amd64.tar.gz | tar xzf - --directory=/ \
 && curl https://dl.google.com/go/go1.17.2.src.tar.gz         | tar xzf - --directory=/gosrc \
 && patch /gosrc/go/src/cmd/go/internal/work/exec.go < /tmp/go-pb.patch \
 && patch /gosrc/go/src/net/dnsconfig_unix.go < /tmp/dns-pb.patch \
 && cd /gosrc/go/src && GOROOT_BOOTSTRAP=/go ./make.bash  \
 && rm -r /go && mv /gosrc/go /go && rm -r /gosrc \
;
WORKDIR /app
VOLUME /app

ENTRYPOINT ["/go/bin/go"]
CMD ["build"]

ADD ./*.go ./*.c ./*.h /gopath/src/github.com/dennwc/inkview/

ARG CC=${SDK_BASE}/usr/bin/arm-obreey-linux-gnueabi-clang
ARG GOOS=linux
ARG GOARCH=arm
ARG GOARM=7
ARG CGO_ENABLED=1

ENV GOROOT=/go GOPATH=/gopath PATH="/go/bin:${SDK_BASE}/usr/bin:$PATH" CC=${CC} GOOS=${GOOS} GOARCH=${GOARCH} GOARM=${GOARM} CGO_ENABLED=${CGO_ENABLED}

RUN go get github.com/mattn/go-sqlite3
