
<a name="EdgeX Virtual Device Service (found in device-virtual-go) Changelog"></a>
## EdgeX Virtual Device Service
[Github repository](https://github.com/edgexfoundry/device-virtual-go)

### Change Logs for EdgeX Dependencies
- [device-sdk-go](https://github.com/edgexfoundry/device-sdk-go/blob/main/CHANGELOG.md)
- [go-mod-core-contracts](https://github.com/edgexfoundry/go-mod-core-contracts/blob/main/CHANGELOG.md)
- [go-mod-bootstrap](https://github.com/edgexfoundry/go-mod-bootstrap/blob/main/CHANGELOG.md)  (indirect dependency)
- [go-mod-messaging](https://github.com/edgexfoundry/go-mod-messaging/blob/main/CHANGELOG.md) (indirect dependency)
- [go-mod-registry](https://github.com/edgexfoundry/go-mod-registry/blob/main/CHANGELOG.md)  (indirect dependency)
- [go-mod-secrets](https://github.com/edgexfoundry/go-mod-secrets/blob/main/CHANGELOG.md) (indirect dependency)
- [go-mod-configuration](https://github.com/edgexfoundry/go-mod-configuration/blob/main/CHANGELOG.md) (indirect dependency)

## [4.0.0] Odessa - 2025-03-12 (Only compatible with the 4.x releases)

### ✨  Features

- Add new build-noziti and docker-noziti targets into Makefile ([ca17556…](https://github.com/edgexfoundry/device-virtual-go/commit/ca17556e0757fa9eca4f5e6146377f7c67c5ecb3))
- Enable PIE support for ASLR and full RELRO ([3f31901…](https://github.com/edgexfoundry/device-virtual-go/commit/3f319016f8e11170ce0158ca92a15177c41c5652))
- Allow empty profileName in Device ([ad92c84…](https://github.com/edgexfoundry/device-virtual-go/commit/ad92c84b73dbd8840292c1a592621e653de09e3c))
- Update device-sdk to support openziti ([#467](https://github.com/edgexfoundry/device-virtual-go/issues/467)) ([14494e7…](https://github.com/edgexfoundry/device-virtual-go/commit/14494e7562879e41062c7461f1813d33e0fc95a7))

### ♻ Code Refactoring

- Update module to v4 ([6bfdd3f…](https://github.com/edgexfoundry/device-rest-go/commit/6bfdd3ff5f3e792eafad2b4c95b01495d5837e2e))
```text
BREAKING CHANGE: update go module to v4
```

### 🐛 Bug Fixes

- Only one ldflags flag is allowed ([ec1ec23…](https://github.com/edgexfoundry/device-rest-go/commit/ec1ec23295ae906f939fdcce9f1e5a4eef1babde))

### 👷 Build

- Upgrade to go-1.23, Linter1.61.0 and Alpine 3.20 ([21d3830…](https://github.com/edgexfoundry/device-rest-go/commit/21d3830f1720a37f1895b62ff4ca15d755f1aed1))

## [v3.1.0] Napa - 2023-11-15 (Only compatible with the 3.x releases)

### ✨  Features

- Remove snap packaging ([#440](https://github.com/edgexfoundry/device-virtual-go/issues/440)) ([a824126…](https://github.com/edgexfoundry/device-virtual-go/commit/a824126fc7aa73da39901a16544d301861d29328))
```text

BREAKING CHANGE: Remove snap packaging ([#440](https://github.com/edgexfoundry/device-virtual-go/issues/440))

```


### ♻ Code Refactoring

- Remove obsolete comments from config file ([#442](https://github.com/edgexfoundry/device-virtual-go/issues/442)) ([105c2ff…](https://github.com/edgexfoundry/device-virtual-go/commit/105c2ff86687b61d63b26fbbe95b93aee24690dc))
- Remove github.com/pkg/errors from Attribution.txt ([fdc055d…](https://github.com/edgexfoundry/device-virtual-go/commit/fdc055dafb82b7d393639588b2b3ea660dae6311))


### 🐛 Bug Fixes

- Add missing SDKVERSION variable in Makefile for version API ([#415](https://github.com/edgexfoundry/device-virtual-go/issues/415)) ([8bac9b2…](https://github.com/edgexfoundry/device-virtual-go/commit/8bac9b22eb31fd3371358eef67a168645ee27d12))


### 👷 Build

- Upgrade to go-1.21, Linter1.54.2 and Alpine 3.18 ([#427](https://github.com/edgexfoundry/device-virtual-go/issues/427)) ([b41150d…](https://github.com/edgexfoundry/device-virtual-go/commit/b41150d8f29f1474a4d4d35e5cf0bbe8d5784310))


### 🤖 Continuous Integration

- Add automated release workflow on tag creation ([0531bee…](https://github.com/edgexfoundry/device-virtual-go/commit/0531beef1411d7a1dfe9a18ad2a8589293894050))


## [v3.0.0] Minnesota - 2023-05-31 (Only compatible with the 3.x releases)

### Features ✨
- Allow min equal to max when generating random value ([#0fb4799](https://github.com/edgexfoundry/device-virtual-go/commits/0fb4799))
- Update for common config ([#339](https://github.com/edgexfoundry/device-virtual-go/pull/339))
    ```text
    BREAKING CHANGE: Configuration file is changed to remove common config settings
    ```
- Use latest SDK for MessageBus Request API ([#337](https://github.com/edgexfoundry/device-virtual-go/pull/337))
    ```text
    BREAKING CHANGE: Commands via MessageBus topic configuration are changed
    ```
- Remove ZeroMQ MessageBus capability ([#325](https://github.com/edgexfoundry/device-virtual-go/pull/325))
    ```text
    BREAKING CHANGE: ZeroMQ MessageBus capability no longer available
    ```

### Bug Fixes 🐛
- Update Discover func error message ([#915c2b7](https://github.com/edgexfoundry/device-virtual-go/commits/915c2b7))
- Update deviceprofile resources for UpdateDevice ([#2d7e124](https://github.com/edgexfoundry/device-virtual-go/commits/2d7e124))
- **snap:** Refactor to avoid conflicts with readonly config provider directory ([#354](https://github.com/edgexfoundry/device-virtual-go/issues/354)) ([#96a5dbd](https://github.com/edgexfoundry/device-virtual-go/commits/96a5dbd))

### Code Refactoring ♻
- Remove deprecated rand.Seed function ([#c9d512d](https://github.com/edgexfoundry/device-virtual-go/commits/c9d512d))
- Use integer for minimum and maximum properties ([#374](https://github.com/edgexfoundry/device-virtual-go/pull/374))
    ```text
    BREAKING CHANGE: Use integer for minimum and maximum properties
    ```
- Change configuration and devices files format to YAML ([#368](https://github.com/edgexfoundry/device-virtual-go/pull/368))
    ```text
    BREAKING CHANGE: Configuration files are now in YAML format, Default file name is now configuration.yaml
    ```
- Remove unused topic configuration ([#345](https://github.com/edgexfoundry/device-virtual-go/issues/345)) ([#d88ec88](https://github.com/edgexfoundry/device-virtual-go/commits/d88ec88))
- Refactor random value generation function ([#d35a6d7](https://github.com/edgexfoundry/device-virtual-go/commits/d35a6d7))
- **snap:** Update command and metadata sourcing ([#355](https://github.com/edgexfoundry/device-virtual-go/issues/355)) ([#34458d9](https://github.com/edgexfoundry/device-virtual-go/commits/34458d9))
- **snap:** Refactor and upgrade to edgex-snap-hooks v3 ([#328](https://github.com/edgexfoundry/device-virtual-go/issues/328)) ([#ade53ff](https://github.com/edgexfoundry/device-virtual-go/commits/ade53ff))

### Documentation 📖
- Add main branch Warning ([#405](https://github.com/edgexfoundry/device-virtual-go/issues/405)) ([#d55cfd1](https://github.com/edgexfoundry/device-virtual-go/commits/d55cfd1))

### Build 👷
- Update to Go 1.20, Alpine 3.17 and linter v1.51.2 ([#534be7e](https://github.com/edgexfoundry/device-virtual-go/commits/534be7e))

## [v2.3.0] Levski - 2022-11-09  (Only compatible with the 2.x releases)

### Features ✨

- Add new Service Metrics configuration ([#08ba88b](https://github.com/edgexfoundry/device-virtual-go/commits/08ba88b))
- Add NATS configuration and build option ([#302](https://github.com/edgexfoundry/device-virtual-go/issues/302)) ([#6354348](https://github.com/edgexfoundry/device-virtual-go/commits/6354348))
- Add commanding via message configuration ([#0b45d56](https://github.com/edgexfoundry/device-virtual-go/commits/0b45d56))
- Add go-winio to attribution (new SPIFFE dependency) ([#a7b7b7f](https://github.com/edgexfoundry/device-virtual-go/commits/a7b7b7f))
- **snap:** Add snap packaging ([#287](https://github.com/edgexfoundry/device-virtual-go/issues/287)) ([#dce4ce0](https://github.com/edgexfoundry/device-virtual-go/commits/dce4ce0))

### Bug Fixes 🐛

- **snap:** Remove duplicate file copying in install hook ([#311](https://github.com/edgexfoundry/device-virtual-go/issues/311)) ([#38745b3](https://github.com/edgexfoundry/device-virtual-go/commits/38745b3))
- **snap:** Set unique name for config interface ([#299](https://github.com/edgexfoundry/device-virtual-go/issues/299)) ([#b155924](https://github.com/edgexfoundry/device-virtual-go/commits/b155924))

### Code Refactoring ♻

- **snap:** edgex-snap-hooks related upgrade ([#290](https://github.com/edgexfoundry/device-virtual-go/issues/290)) ([#1d4e8f4](https://github.com/edgexfoundry/device-virtual-go/commits/1d4e8f4))

### Build 👷

- Upgrade to Go 1.18 and alpine 3.16 ([#294](https://github.com/edgexfoundry/device-virtual-go/issues/294)) ([#92de881](https://github.com/edgexfoundry/device-virtual-go/commits/92de881))

## [v2.2.0] Kamakura - 2022-05-11  (Only compatible with the 2.x releases)

### Features ✨
- Add MaxEventSize and ReadingUnits to configuration ([#1d794d6](https://github.com/edgexfoundry/device-virtual-go/commits/1d794d6))
- Enable security hardening ([#5ba56e1](https://github.com/edgexfoundry/device-virtual-go/commits/5ba56e1))
- **security:** Roll out delayed start configuration.toml scaffolding ([#01bd024](https://github.com/edgexfoundry/device-virtual-go/commits/01bd024))

### Bug Fixes 🐛
- **security:** Dependency version bump for device-sdk and go-mod-core-contracts ([#cf90458](https://github.com/edgexfoundry/device-virtual-go/commits/cf90458))

### Performance Improvements ⚡
- **app:** Use maps instead of in-RAM SQL DB ([#260](https://github.com/edgexfoundry/device-virtual-go/issues/260)) ([#261](https://github.com/edgexfoundry/device-virtual-go/issues/261)) ([#7f10fc8](https://github.com/edgexfoundry/device-virtual-go/commits/7f10fc8))

### Build 👷
- Make w/o cgo on Windows ([#07b7053](https://github.com/edgexfoundry/device-virtual-go/commits/07b7053))
- Update to latest SDK w/o ZMQ on windows ([#0843306](https://github.com/edgexfoundry/device-virtual-go/commits/0843306))
    ```
    BREAKING CHANGE:
    ZeroMQ no longer supported on native Windows for EdgeX
    MessageBus
    ```

### Continuous Integration 🔄
- gomod changes related for Go 1.17 ([#257b1e4](https://github.com/edgexfoundry/device-virtual-go/commits/257b1e4))
- Go 1.17 related changes ([#bf2a4df](https://github.com/edgexfoundry/device-virtual-go/commits/bf2a4df))

## [v2.1.0] Jakarta - 2021-11-18  (Only compatible with the 2.x releases)

### Features ✨
- Update configuration for new CORS and Secrets File settings ([#c0ef7e9](https://github.com/edgexfoundry/device-virtual-go/commits/c0ef7e9))

### Bug Fixes 🐛
- Update all TOML to use quote and not single-quote ([#7c8b3a8](https://github.com/edgexfoundry/device-virtual-go/commits/7c8b3a8))
- Use formatted versions of logging APIs and fine tune err messages ([#2332541](https://github.com/edgexfoundry/device-virtual-go/commits/2332541))
- Optimize defer statements ([#6648057](https://github.com/edgexfoundry/device-virtual-go/commits/6648057))
- Remove unnecessary device update logic ([#08e808d](https://github.com/edgexfoundry/device-virtual-go/commits/08e808d))

### Build 👷
- Update to use released SDK ([#72804ff](https://github.com/edgexfoundry/device-virtual-go/commits/72804ff))
- Update alpine base to 3.14 ([#b8ea1e8](https://github.com/edgexfoundry/device-virtual-go/commits/b8ea1e8))

### Continuous Integration 🔄
- Remove need for CI specific Dockerfile ([#085dd40](https://github.com/edgexfoundry/device-virtual-go/commits/085dd40))

## [v2.0.0] Ireland - 2021-06-30  (Only compatible with the 2.x releases)

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
