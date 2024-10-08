# GOLapsDump

## Overview

GOLapsDump is a fast and efficient tool written in Golang for extracting Laps Passwords.

## Getting Started

**Clone the repository:**
```bash
git clone https://github.com/k4ls3c/GOLapsDump.git
```
Compile the source code:
```
go build
```
## Usage
Execute GOLapsDump with the following command:
```
GOLapsDump.exe -u jorah.mormont -p Pa$$w0rd -d essos.local -l 192.168.56.12
```
```
GOLapsDump.exe -u jorah.mormont -H 92937945B518814341DE3F726500D4FF -d essos.local -l 192.168.56.12
```
![GOLapsDump](https://github.com/k4ls3c/GOLapsDump/assets/148506834/f4e40b08-b254-40ce-a5aa-f84ecd897a63)

Options
```
  -u    username for LDAP
  -p    password for LDAP
  -H    ntHash
  -l    LDAP server
  -d    Domain
  -port LDAP server port (default is 389)
  -o    Output file path
```
Example
```
GOLapsDump.exe -u jorah.mormont -p Pa$$w0rd -d essos.local -l 192.168.56.12 -o /path/to/output/file.txt
```
## Disclaimer

The author is not responsible for unauthorized use of this tool. Use responsibly and ensure compliance with legal and ethical standards.

## Reference
- https://github.com/kfosaaen/Get-LAPSPasswords/blob/master/Get-LAPSPasswords.ps1
- https://github.com/n00py/LAPSDumper
