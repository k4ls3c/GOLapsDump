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
GOLapsDump.exe -u user@na.domain.local -p Pa$$w0rd -d na.domain.local
```
Options
```
    -u string: Specify the username for LDAP.
    -p string: Provide the password for LDAP.
    -l string: Define the LDAP server (or domain).
    -d string: Set the domain.
    -port int: Specify the LDAP server port (default is 389).
    -o string: Define the output file path.
```
Example
```
GOLapsDump.exe -u user@na.domain.local -p Pa$$w0rd -d na.domain.local -o /path/to/output/file.txt
```
## Disclaimer

The author is not responsible for unauthorized use of this tool. Use responsibly and ensure compliance with legal and ethical standards.

## Reference
- https://github.com/kfosaaen/Get-LAPSPasswords/blob/master/Get-LAPSPasswords.ps1
