#
# Copyright (C) 2019 IOTech Ltd
#
# SPDX-License-Identifier: Apache-2.0

FROM golang:1.11-alpine AS builder

LABEL license='SPDX-License-Identifier: Apache-2.0' \
  copyright='Copyright (c) 2019: IOTech'

RUN sed -e 's/dl-cdn[.]alpinelinux.org/nl.alpinelinux.org/g' -i~ /etc/apk/repositories

RUN apk add --update --no-cache make git

# set the working directory
WORKDIR /github.com/edgexfoundry/device-virtual

COPY . .

RUN make build

FROM scratch

ENV APP_PORT=49990
EXPOSE $APP_PORT

COPY --from=builder /github.com/edgexfoundry/device-virtual/cmd /

ENTRYPOINT ["/device-virtual","--profile=docker","--confdir=/res","--registry=consul://edgex-core-consul:8500"]
