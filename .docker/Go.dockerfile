FROM golang:1.21.0-alpine as golang

COPY . /tmp/factory

WORKDIR /tmp/factory

ENV GOCACHE=/tmp/.cache/go-build

RUN apk upgrade --no-cache && \
    apk add --no-cache bash curl git unzip && \
    echo -n "Downloading Go modules ... " && \
    go mod download && \
    echo "Done" && \
    go install github.com/jstemmer/go-junit-report@latest && \
    go install github.com/axw/gocov/gocov@latest && \
    go install github.com/matm/gocov-html/cmd/gocov-html@v1.2.0 && \
    find /go/pkg/mod -type d -exec chmod 755 {} \; && \
    find /go/pkg/mod/cache -type d -exec chmod 777 {} \; && \
    find ${GOCACHE} -type d -exec chmod 777 {} \;

# -----------------------------------------------------------------
FROM golang as devcontainer

RUN apk add --no-cache build-base gcc make && \
    go install github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest && \
    go install github.com/ramya-rao-a/go-outline@latest && \
    go install github.com/cweill/gotests/gotests@latest && \
    go install github.com/fatih/gomodifytags@latest && \
    go install github.com/josharian/impl@latest && \
    go install github.com/haya14busa/goplay/cmd/goplay@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install honnef.co/go/tools/cmd/staticcheck@latest && \
    go install golang.org/x/tools/gopls@latest && \
    ln -s /go/bin/dlv /go/bin/dlv-dap && \
    apk del build-base gcc && \
    touch /root/.bashrc