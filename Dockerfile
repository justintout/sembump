FROM golang:alpine as builder
LABEL AUTHOR Justin Tout <justin.tout@case.edu>

ENV PATH /go/bin:/usr/local/go/bin:$PATH
ENV GOPATH /go

RUN	apk add --no-cache \
    ca-certificates

COPY . /go/src/github.com/justintout/base

RUN set -x \
    && apk add --no-cache --virtual .build-deps \
    git \
    gcc \
    libc-dev \
    libgcc \
    make \
    && cd /go/src/github.com/justintout/base \
    && make static \
    && mv base /usr/bin/base \
    && apk del .build-deps \
    && rm -rf /go \
    && echo "Build complete."

FROM scratch

COPY --from=builder /usr/bin/base /usr/bin/base
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs

ENTRYPOINT [ "base" ]
CMD [ "--help" ]
