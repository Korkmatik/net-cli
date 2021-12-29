# net-cli
This is a cli tool to do basic network operations

```
NAME:
   Website Lookup CLI - Let's you query IPs, CNAMEs, MX records and name Servers!

USAGE:
   cli.exe [global options] command [command options] [arguments...]

COMMANDS:
   ns       Looks up the Name Servers for a Particular Host
   ip       Looks up the IP addresses for a particular host
   cname    Looks up the CNAME for a particular host
   mx       Looks up the MX records for a particular host
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

# Usage Examples

For example you can get the name servers for a particular host:

```
.\cli.exe ns --host google.de
```

[![nameservers.png](https://i.postimg.cc/fLw45QGV/nameservers.png)](https://postimg.cc/2qXX8M8Y)

You can also get the IP for a domain:

```
.\cli.exe ip --host google.de
```

[![ip.png](https://i.postimg.cc/zBdQjbwQ/ip.png)](https://postimg.cc/2bZTj5Kd)

You can also get the CNAMEs for a particular host:

```
.\cli.exe cname --host google.de
```

[![cname.png](https://i.postimg.cc/VN7y7Vm4/cname.png)](https://postimg.cc/5HFkypHF)

Also you can ask for the MX Records for a particular domain:

```
.\cli.exe mx --host google.com
```

[![mx.png](https://i.postimg.cc/RhpN4R7T/mx.png)](https://postimg.cc/9wTWByN4)

# Compiling Instructions

## Linux
You will need the Golang compiler for this. You can compile it with the following command:

```bash
git clone https://github.com/Korkmatik/net-cli.git
cd net-cli
go build cmd/my-cli/cli.go cmd/my-cli/actions.go
```

Or use the `build.sh` script inside this repository

## Windows

You will need to install the Golang compiler to compile it. Run the following commands inside Powershell:

```powershell
git clone https://github.com/Korkmatik/net-cli.git
cd net-cli
go build .\cmd\my-cli\cli.go .\cmd\my-cli\actions.go
```
