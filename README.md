# Device-Virtual-GO
The functions of Virtual Device Service GO just like the java version [device-virtual](https://github.com/edgexfoundry/device-virtual),
but supports more data types of random values:
* Bool
* Int8, Int16, Int32, Int64
* Uint8, Uint16, Uint32, Uint64
* Float32, Float64

This version of Virtual Device Service is implemented based on [Device SDK GO](https://github.com/edgexfoundry/device-sdk-go),
and leveraged [ql](https://godoc.org/modernc.org/ql) (an embedded SQL database engine) to simulate virtual resources.

# Docker compose file settings
Adding service:
```yaml
device-virtual:
image: edgexfoundry/docker-device-virtual-go:1.0.0
ports:
  - "49990:49990"
container_name: device-virtual
hostname: device-virtual
networks:
  - edgex-network
volumes:
  - db-data:/data/db
  - log-data:/edgex/logs
  - consul-config:/consul/config
  - consul-data:/consul/data
  - db-devicevirtual:/db # Mount ql database directory is optional
depends_on:
  - data
  - command
```
# How to use
For now, Virtual Device Service contains 4 pre-defined devices as random value generators:
* Random-Boolean-Generator01
* Random-Integer-Generator01
* Random-UnsignedInteger-Generator01
* Random-Float-Generator01

Use Core-Command Service APIs to find executable commands information:
* http://[host]:48082/api/v1/device/name/Random-Boolean-Generator01
* http://[host]:48082/api/v1/device/name/Random-Integer-Generator01
* http://[host]:48082/api/v1/device/name/Random-UnsignedInteger-Generator01
* http://[host]:48082/api/v1/device/name/Random-Float-Generator01

NOTE:
* The Enable_Randomization attribute of resource is automatically disabled when you use put command to set a specified value.
* The minimum and maximum values of resource can be specified in the device profile. Example:
>>```yaml
>>deviceResources:
>>  -
>>    name: "RandomValue_Int8"
>>    description: "Generate random int8 value"
>>    properties:
>>      value:
>>        { type: "Int8", readWrite: "R", minimum: "-100", maximum: "100", defaultValue: "0" }
>>      units:
>>        { type: "String", readWrite: "R", defaultValue: "random int8 value" }
>>```
# Manipulate virtual resources via command ql tool (optional)
1. Install [command ql](https://godoc.org/modernc.org/ql/ql)
2. Enter ql database directory:
> * If the Virtual Device Service runs in Docker container, use the following command to find the path of ql database directory: 
>> ```console
>> $ docker volume inspect edgex_temp_db-devicevirtual
>> ```
> * If the Virtual Device Service runs in dev mode, the ql database directory is under the driver directory. 

Command examples:
> * Query all data:
>>```console
>>$ ql -db deviceVirtual.db -fld "select * from VIRTUAL_RESOURCE"
>>```
> * Update Enable_Randomization:
>> ```console
>>$ ql -db deviceVirtual.db "update VIRTUAL_RESOURCE set ENABLE_RANDOMIZATION=false where DEVICE_NAME=\"Random-Integer-Generator01\" and DEVICE_RESOURCE_NAME=\"RandomValue_Int8\" "
>> ```
> * Update Value:
>> ```console
>>$ ql -db deviceVirtual.db "update VIRTUAL_RESOURCE set VALUE=\"26\" where DEVICE_NAME=\"Random-Integer-Generator01\" and DEVICE_RESOURCE_NAME=\"RandomValue_Int8\" "
>> ```
