# Embedded Configuration

This component creates the CloudInit UserData for an EC2 server to 
install the necessary requirements for running Rancher.

## Assumptions

This component assumes that target machine is running a version of Ubuntu.

## Installed Items

This component installs the following items on the server:

* Docker
* Docker AUFS support
* Rancher Server Docker image


Once installed, the Rancher image is launched and bound to the host machines port.
