FROM ubuntu:xenial

RUN apt-get update && \
    apt-get install -y xz-utils build-essential libc6-i386 wget nano && \
    apt-get clean

# download Pocketbook SDK
RUN wget https://storage.googleapis.com/dennwc-public/pbsdk-linux-1.1.0.deb -qO /tmp/pbsdk-linux.deb && \
    dpkg -i /tmp/pbsdk-linux.deb && \
    rm /tmp/pbsdk-linux.deb

ADD ./patches/* /tmp/

# download specified Go binary release that will act as a bootstrap compiler for Go toolchain
# download sources for that release and apply the patch
# build a new toolchain and remove an old one
RUN wget https://dl.google.com/go/go1.9.4.linux-amd64.tar.gz -qO /tmp/go.tar.gz && \
    tar -xf /tmp/go.tar.gz && \
    rm /tmp/go.tar.gz && \
    wget https://dl.google.com/go/go1.9.4.src.tar.gz -qO /tmp/go.tar.gz && \
    mkdir -p /gosrc && tar -xf /tmp/go.tar.gz -C /gosrc && \
    rm /tmp/go.tar.gz && \
    patch /gosrc/go/src/cmd/go/internal/work/build.go < /tmp/go-pb.patch && \
    patch /gosrc/go/src/net/dnsconfig_unix.go < /tmp/dns-pb.patch && \
    cd /gosrc/go/src && GOROOT_BOOTSTRAP=/go ./make.bash && \
    rm -r /go && mv /gosrc/go /go && rm -r /gosrc

WORKDIR /app
VOLUME /app

ADD build.sh /
ENTRYPOINT ["/build.sh"]

ADD ./* /gopath/src/github.com/dennwc/inkview/