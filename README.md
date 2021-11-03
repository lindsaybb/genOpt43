# genOpt43

genOpt43 is an Iskratel PON utility program for converting TR069 ACS URL to a Hex-Encoded DHCP Option.
This value can enable the Innbox CPE/ONT to connect to this service at the point of Network Discovery.
This program can also decode the value in the format it generates for validation.

| Flag | Description |
| ------ | ------ |
| -h | Show this help |

# Example Input
```sh
http://10.5.100.205:7547
```

# Example Output
```sh
0x0118687474703a2f2f31302e352e3130302e3230353a37353437
```

# Example Input
```sh
0x0118687474703a2f2f31302e352e3130302e3230353a37353437
```

# Example Output
```sh
http://10.5.100.205:7547
```

# Example Input
```sh
https://acs.lab.local:10302
```

# Example Output
```sh
0x011b68747470733a2f2f6163732e6c61622e6c6f63616c3a3130333032
```



# On a Mikrotik DHCP Server:
```sh
/ip dhcp-server option add name=ACS code=43 value=0x011b68747470733a2f2f6163732e6c61622e6c6f63616c3a3130333032
/ip dhcp-server network set 0 dhcp-option=ACS
```
