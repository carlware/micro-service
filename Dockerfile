## -------------------------------------------------------------------------------------------------

FROM golang:1.14-alpine as builder

# Arguments
ARG GIT_ID
ARG GIT_TOKEN

ENV GRPC_HEALTH_PROBE_VERSION=v0.2.2
ENV GRPC_WEB_VERSION=1.0.4

RUN set -eux; \
    apk update && \
    apk add --no-cache make upx gcc musl-dev linux-headers git zip unzip wget curl && \
    curl -sSL https://github.com/grpc/grpc-web/releases/download/${GRPC_WEB_VERSION}/protoc-gen-grpc-web-${GRPC_WEB_VERSION}-linux-x86_64 -o /usr/local/bin/protoc-gen-grpc-web && \
    chmod +x /usr/local/bin/protoc-gen-grpc-web && \
    wget -q -O /usr/local/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /usr/local/bin/grpc_health_probe && \
    mkdir /src && \
    echo 'nobody:x:65534:' > /src/group.nobody && \
    echo 'nobody:x:65534:65534::/:' > /src/passwd.nobody

RUN adduser -D -g "" golang

# Force go modules
ENV GO111MODULE=on
ENV GO_VERSION="go version | awk '{print $3}'"

WORKDIR $GOPATH/src/workspace

# Copy project go module
RUN chown golang:golang $GOPATH/src/workspace
USER golang
COPY --chown=golang:golang . .

# Private repo
# RUN git config \
#     --global \
#     url."https://${GIT_ID}:${GIT_TOKEN}@gitlab.com/condomilux/go-common".insteadOf \
#     "https://gitlab.com/condomilux/go-common"

# download tools
RUN make -C tools

# Build final target
RUN make build

# Compress binaries
RUN set -eux; \
    upx -9 bin/* && \
    chmod +x bin/*

# ## -------------------------------------------------------------------------------------------------

# FROM gcr.io/distroless/base
FROM golang:alpine

# Arguments
ARG BUILD_DATE
ARG VERSION
ARG VCS_REF
ARG APP_NAME

# Metadata
LABEL com.condomilux.build-date=$BUILD_DATE
LABEL com.condomilux.name=$APP_NAME
LABEL com.condomilux.vcs-ref=$VCS_REF
LABEL com.condomilux.vendor="condomilux"
LABEL com.condomilux.version=$VERSION

COPY --from=builder /src/group.nobody /etc/group
COPY --from=builder /src/passwd.nobody /etc/passwd
USER nobody:nobody

COPY --from=builder /go/src/workspace/bin/app /bin/service
COPY --from=builder /go/src/workspace/migrations /go/migrations/
# COPY --from=builder /usr/local/bin/grpc_health_probe /bin/grpc_health_probe

WORKDIR /go

# CMD ["/bin/service"]
ENTRYPOINT ["/bin/service"]
