# Jeep - Jira Issue Printer

## Requirements

* SSH forwarding server
* Public domain and IP address
* SSL certificate (Jira webhooks only supports HTTPS)
* Thermal printer
* Raspberry Pi Zero or any other Linux machine you can keep connected to the printer
* Jira

![Jira Issue Printer gif](https://orelfichman.com/wp-content/uploads/2022/06/jira-printer-1.gif =400x496)


## Setup

1. Make sure these settings are present in `/etc/ssh/sshd_config`:
```config
GatewayPorts clientspecified
AllowTcpForwarding yes
```
2. Set up your DNS to point to the publicly accessible server
3. Set up the [Jira webhook](https://developer.atlassian.com/server/jira/platform/webhooks/)
2. Download the latest version from the releases page and place it in a directory of your choosing
3. Connect the printer to your machine
4. make a directory named `config` where you placed the executable, and create & edit `config.yaml` to suit your needs
5. Perform SSH forwarding to the remote server:
```sh
ssh -R 0.0.0.0:<remote-port>:0.0.0.0<local-port> -N -f <user>@<host>
```
5. Launch the application (as root or a member of the `lp` group)
6. Hear that sweet, sweet sound of thermal printing once an issue is opened
