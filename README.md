# Device Virtual Go
[![Build Status](https://jenkins.edgexfoundry.org/view/EdgeX%20Foundry%20Project/job/edgexfoundry/job/device-virtual-go/job/main/badge/icon)](https://jenkins.edgexfoundry.org/view/EdgeX%20Foundry%20Project/job/edgexfoundry/job/device-virtual-go/job/main/) [![Code Coverage](https://codecov.io/gh/edgexfoundry/device-virtual-go/branch/main/graph/badge.svg?token=ll7zq2c3Q7)](https://codecov.io/gh/edgexfoundry/device-virtual-go) [![Code Coverage](https://codecov.io/gh/edgexfoundry/device-virtual-go/branch/master/graph/badge.svg?token=ll7zq2c3Q7)](https://codecov.io/gh/edgexfoundry/device-virtual-go) [![Go Report Card](https://goreportcard.com/badge/github.com/edgexfoundry/device-virtual-go)](https://goreportcard.com/report/github.com/edgexfoundry/device-virtual-go) [![GitHub Latest Dev Tag)](https://img.shields.io/github/v/tag/edgexfoundry/device-virtual-go?include_prereleases&sort=semver&label=latest-dev)](https://github.com/edgexfoundry/device-virtual-go/tags) ![GitHub Latest Stable Tag)](https://img.shields.io/github/v/tag/edgexfoundry/device-virtual-go?sort=semver&label=latest-stable) [![GitHub License](https://img.shields.io/github/license/edgexfoundry/device-virtual-go)](https://choosealicense.com/licenses/apache-2.0/) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/edgexfoundry/device-virtual-go) [![GitHub Pull Requests](https://img.shields.io/github/issues-pr-raw/edgexfoundry/device-virtual-go)](https://github.com/edgexfoundry/device-virtual-go/pulls) [![GitHub Contributors](https://img.shields.io/github/contributors/edgexfoundry/device-virtual-go)](https://github.com/edgexfoundry/device-virtual-go/contributors) [![GitHub Committers](https://img.shields.io/badge/team-committers-green)](https://github.com/orgs/edgexfoundry/teams/device-virtual-go-committers/members) [![GitHub Commit Activity](https://img.shields.io/github/commit-activity/m/edgexfoundry/device-virtual-go)](https://github.com/edgexfoundry/device-virtual-go/commits)

> **Warning**  
> The **main** branch of this repository contains work-in-progress development code for the upcoming release, and is **not guaranteed to be stable or working**.
> It is only compatible with the [main branch of edgex-compose](https://github.com/edgexfoundry/edgex-compose) which uses the Docker images built from the **main** branch of this repo and other repos.
>
> **The source for the latest release can be found at [Releases](https://github.com/edgexfoundry/device-virtual-go/releases).**

## Overview
The virtual device service simulates different kinds of [devices](https://docs.edgexfoundry.org/2.1/general/Definitions/#device) to generate events and readings to the [core data](https://docs.edgexfoundry.org/2.1/microservices/core/data/Ch-CoreData/) micro service, and users send commands and get responses through the [command and control](https://docs.edgexfoundry.org/2.1/microservices/core/command/Ch-Command/) micro service. These features of the virtual device services are useful when executing functional or performance tests without having any real devices.
## Usage
Users can refer to [the document](https://docs.edgexfoundry.org/2.1/microservices/device/virtual/Ch-VirtualDevice/) to learn how to use this device service.

## Build with NATS Messaging
Currently, the NATS Messaging capability (NATS MessageBus) is opt-in at build time.
This means that the published Docker image and Snaps do not include the NATS messaging capability.

The following make commands will build the local binary or local Docker image with NATS messaging
capability included.
```makefile
make build-nats
make docker-nats
```

The locally built Docker image can then be used in place of the published Docker image in your compose file.
See [Compose Builder](https://github.com/edgexfoundry/edgex-compose/tree/main/compose-builder#gen) `nat-bus` option to generate compose file for NATS and local dev images.

## Community
- Chat: https://edgexfoundry.slack.com
- Mailing lists: https://lists.edgexfoundry.org/mailman/listinfo

## License
[Apache-2.0](LICENSE)
