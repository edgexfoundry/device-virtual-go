# EdgeX Device Virtual Service Snap
[![edgex-device-virtual](https://snapcraft.io/edgex-device-virtual/badge.svg)](https://snapcraft.io/edgex-device-virtual)

This directory contains the snap packaging of the EdgeX Device Virtual device service.

The snap is built automatically and published on the Snap Store as [edgex-device-virtual].

For usage instructions, please refer to Device Virtual section in [Getting Started using Snaps][docs].

## Build from source
Execute the following command from the top-level directory of this repo:
```
snapcraft
```

This will create a snap package file with `.snap` extension. It can be installed locally by setting the `--dangerous` flag:
```bash
sudo snap install --dangerous <snap-file>
```

The [snapcraft overview](https://snapcraft.io/docs/snapcraft-overview) provides additional details.

### Obtain a Secret Store token
The `edgex-secretstore-token` snap slot makes it possible to automatically receive a token from a locally installed platform snap. Note that the **auto connection does NOT happen right** now because the snap publisher isn't same as the `edgexfoundry` platrform snap (i.e. Canonical).

If the snap is built and installed locally, the interface will not auto-connect. You can check the status of the connections by running the `snap connections edgex-device-virtual` command.

To manually connect and obtain a token:
```bash
sudo snap connect edgexfoundry:edgex-secretstore-token edgex-device-virtual:edgex-secretstore-token
```

Please refer [here][secret-store-token] for further information.

[edgex-device-virtual]: https://snapcraft.io/edgex-device-virtual
[docs]: https://docs.edgexfoundry.org/2.2/getting-started/Ch-GettingStartedSnapUsers/#device-virtual
[secret-store-token]: https://docs.edgexfoundry.org/2.2/getting-started/Ch-GettingStartedSnapUsers/#secret-store-token

