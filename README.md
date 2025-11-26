![image](https://github.com/mytechnotalent/traffic-generator/blob/main/turbo-attack.png?raw=true)

## FREE Reverse Engineering Self-Study Course [HERE](https://github.com/mytechnotalent/Reverse-Engineering-Tutorial)

<br>

# turbo-attack
A turbo traffic generator pentesting tool to generate random traffic with random MAC and IP addresses in addition to random sequence numbers to a particular IP and port.

<br>

# Linux IPv4 (ARM64)
```bash
usage: sudo ./turbo-attack_010_linux_arm64 eth0 4 192.168.0.2 443 150000
```

# Linux IPv6 (ARM64)
```bash
usage: sudo ./turbo-attack_010_linux_arm64 eth0 6 fe80:0000:0000:0000:0000:0000:0000:0002 443 150000
```

# Linux IPv4 (AMD64)
```bash
usage: sudo ./turbo-attack_010_linux_amd64 eth0 4 192.168.0.2 443 150000
```

# Linux IPv6 (AMD64)
```bash
usage: sudo ./turbo-attack_010_linux_amd64 eth0 6 fe80:0000:0000:0000:0000:0000:0000:0002 443 150000
```

<br>

# Usage
* Designed for Kali Linux or Debian distros only.  This software uses Linux syscalls in order to optimize for speed.
* Tested within a Kali Linux ARM64 VMware Fusion VM on a Macbook Pro ARM64 machine.

<br>

# Terms Of Use
* Do NOT use this on any computer you do not own or are not allowed to run this on.<br>
* You may NEVER attempt to sell this, it is free and open source.<br>
* The authors and publishers assume no responsibility.<br>
* For educational purposes only.

<br>

# Run Tests
```bash
sudo /usr/local/go/bin/go test -v -cover ./...
sudo /usr/local/go/bin/go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

<br>

# License
[Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0)
