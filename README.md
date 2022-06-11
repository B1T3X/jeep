# Jeep - Jira Issue Printer

## Requirements

* SSH forwarding server
* Public domain and IP address
* SSL certificate (Jira webhooks only supports HTTPS)
* Thermal printer
* Raspberry Pi Zero or any other Linux machine you can keep connected to the printer
* Jira

## Setup

1. Make sure these settings are present in `/etc/ssh/sshd_config`:
```config
GatewayPorts clientspecified
AllowTcpForwarding yes
```
2. Download the latest version from the releases page and place it in
   a directory of your choosing
3. Connect the printer to your machine
4. make a directory named `config` where you placed the executable, and create & edit `config.yaml` to suit your needs
5. Perform SSH forwarding to the remote server:
```sh
ssh -R 0.0.0.0:<remote-port>:0.0.0.0<local-port> -N -f <user>@<host>
```
5. Launch the application
6. Hear that sweet, sweet sound of thermal printing once an issue is opened
