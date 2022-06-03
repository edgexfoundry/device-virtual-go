
<a name="EdgeX Virtual Device Service (found in device-virtual-go) Changelog"></a>
## EdgeX Virtual Device Service
[Github repository](https://github.com/edgexfoundry/device-virtual-go)

## Change Logs for EdgeX Dependencies

- [device-sdk-go](https://github.com/edgexfoundry/device-sdk-go/blob/jakarta/CHANGELOG.md)

## [v2.1.1] - Jakarta - 2021-06-08 (Only compatible with the 2.x releases)

### Bug Fixes 🐛

- see SDK changelog link above

## [v2.0.0] Ireland - 2021-06-30  (Not Compatible with 1.x releases)
### Change Logs for EdgeX Dependencies
- [device-sdk-go](https://github.com/edgexfoundry/device-sdk-go/blob/v2.0.0/CHANGELOG.md)
- [go-mod-core-contracts](https://github.com/edgexfoundry/go-mod-core-contracts/blob/v2.0.0/CHANGELOG.md)

### Features ✨
- Enable using MessageBus as the default ([#dee740d](https://github.com/edgexfoundry/device-virtual-go/commits/dee740d))
- Add secure MessagBus capability ([#e8304ae](https://github.com/edgexfoundry/device-virtual-go/commits/e8304ae))
- Add Registry/Config Access token capability ([#182](https://github.com/edgexfoundry/device-virtual-go/issues/182)) ([#ade1702](https://github.com/edgexfoundry/device-virtual-go/commits/ade1702))
    ```
    BREAKING CHANGE:
    When run with the secure Edgex Stack now need to have the SecretStore configured, a Vault token created and run with EDGEX_SECURITY_SECRET_STORE=true.
    ```
- update driver implementation to reflect v2 profiles ([#0759054](https://github.com/edgexfoundry/device-virtual-go/commits/0759054))
- Remove Logging configuration ([#7c0b286](https://github.com/edgexfoundry/device-virtual-go/commits/7c0b286))
### Bug Fixes 🐛
- update separator for GET array value handler function ([#f5ae1f5](https://github.com/edgexfoundry/device-virtual-go/commits/f5ae1f5))
- update default service configuration ([#735eab6](https://github.com/edgexfoundry/device-virtual-go/commits/735eab6))
- pass correct argument in prepareVirtualResources ([#3f1af1c](https://github.com/edgexfoundry/device-virtual-go/commits/3f1af1c))
### Code Refactoring ♻
- remove unimplemented InitCmd/RemoveCmd configuration ([#db5966d](https://github.com/edgexfoundry/device-virtual-go/commits/db5966d))
- Change PublishTopicPrefix value to be 'edgex/events/device' ([#3806501](https://github.com/edgexfoundry/device-virtual-go/commits/3806501))
- Update to assign and uses new Port Assignments ([#a4c3f51](https://github.com/edgexfoundry/device-virtual-go/commits/a4c3f51))
    ```
    BREAKING CHANGE:
    Device Virtual default port number has changed to 59900
    ```
- Moved go mod tidy in dockerfile ([#e5c72d3](https://github.com/edgexfoundry/device-virtual-go/commits/e5c72d3))
- Update for new service key names and overrides for hyphen to underscore ([#f085b26](https://github.com/edgexfoundry/device-virtual-go/commits/f085b26))
    ```
    BREAKING CHANGE:
    Service key names used in configuration have changed.
    ```
- Updated to latest SDK and update MessageQue type to be `redis` ([#ff158f9](https://github.com/edgexfoundry/device-virtual-go/commits/ff158f9))
- consume v2 Device SDK ([#941086c](https://github.com/edgexfoundry/device-virtual-go/commits/941086c))
### Documentation 📖
- Add badges to readme ([#e3e4674](https://github.com/edgexfoundry/device-virtual-go/commits/e3e4674))
### Build 👷
- update Dockerfile to go 1.16 ([#8cf28ce](https://github.com/edgexfoundry/device-virtual-go/commits/8cf28ce))
- update go.mod to go 1.16 ([#587e06e](https://github.com/edgexfoundry/device-virtual-go/commits/587e06e))
### Continuous Integration 🔄
- update local docker image names ([#390274e](https://github.com/edgexfoundry/device-virtual-go/commits/390274e))

<a name="v1.3.1"></a>
## [v1.3.1] - 2021-02-02
### Code Refactoring ♻
- Upgrade SDK to v1.4.0 ([#dd6dddd](https://github.com/edgexfoundry/device-virtual-go/commits/dd6dddd))
### Build 👷
- update device-sdk-go to v1.3.1-dev.4 ([#d2603f8](https://github.com/edgexfoundry/device-virtual-go/commits/d2603f8))
- **deps:** Bump github.com/edgexfoundry/device-sdk-go ([#fb417ca](https://github.com/edgexfoundry/device-virtual-go/commits/fb417ca))
### Continuous Integration 🔄
- add semantic.yml for commit linting, update PR template to latest ([#c0dc29d](https://github.com/edgexfoundry/device-virtual-go/commits/c0dc29d))
- standardize dockerfiles ([#6351328](https://github.com/edgexfoundry/device-virtual-go/commits/6351328))

<a name="v1.3.0"></a>
## [v1.3.0] - 2020-11-18
### Features ✨
- Support array value type ([#56f7dc2](https://github.com/edgexfoundry/device-virtual-go/commits/56f7dc2))
### Doc
- Update top-level README ([#d268b7b](https://github.com/edgexfoundry/device-virtual-go/commits/d268b7b))
### Bug Fixes 🐛
- ReadWrite field of each device resource should be RW ([#b91c50e](https://github.com/edgexfoundry/device-virtual-go/commits/b91c50e))
### Code Refactoring ♻
- Upgrade SDK to v1.2.4-dev.34 ([#1077bc0](https://github.com/edgexfoundry/device-virtual-go/commits/1077bc0))
- update dockerfile to appropriately use ENTRYPOINT and CMD, closes[#125](https://github.com/edgexfoundry/device-virtual-go/issues/125) ([#ee911db](https://github.com/edgexfoundry/device-virtual-go/commits/ee911db))
### Build 👷
- update go-mod-core-contracts to 0.1.111 ([#7fc4ffb](https://github.com/edgexfoundry/device-virtual-go/commits/7fc4ffb))
- update device-sdk-go to 1.3.0 ([#b61769c](https://github.com/edgexfoundry/device-virtual-go/commits/b61769c))
- upgrade device-sdk-go ([#edf0204](https://github.com/edgexfoundry/device-virtual-go/commits/edf0204))
- upgrade to use Go1.15 ([#7a8becd](https://github.com/edgexfoundry/device-virtual-go/commits/7a8becd))
- **all:** Enable use of DependaBot to maintain Go dependencies ([#befc574](https://github.com/edgexfoundry/device-virtual-go/commits/befc574))

<a name="v1.2.3"></a>
## [v1.2.3] - 2020-08-19
### Bug Fixes 🐛
- service fails when run with read-only root file system ([#9874cd4](https://github.com/edgexfoundry/device-virtual-go/commits/9874cd4))

<a name="v1.2.2"></a>
## [v1.2.2] - 2020-07-09
### Doc
- update pr template to include dependencies section ([#e9454c0](https://github.com/edgexfoundry/device-virtual-go/commits/e9454c0))
### Bug Fixes 🐛
- High memory usage ([#02e176c](https://github.com/edgexfoundry/device-virtual-go/commits/02e176c))

<a name="v1.2.1"></a>
## [v1.2.1] - 2020-06-12
### Code Refactoring ♻
- remove binary autoevent ([#4f04737](https://github.com/edgexfoundry/device-virtual-go/commits/4f04737))
- upgrade SDK to v1.2.0 ([#01be24e](https://github.com/edgexfoundry/device-virtual-go/commits/01be24e))
