# gexpose

A net tool that exposes local service to public.

[![Travis](https://travis-ci.com/net-byte/gexpose.svg?branch=main)](https://github.com/net-byte/gexpose)
[![Go Report Card](https://goreportcard.com/badge/github.com/net-byte/gexpose)](https://goreportcard.com/report/github.com/net-byte/gexpose)
![image](https://img.shields.io/badge/License-MIT-orange)
![image](https://img.shields.io/badge/License-Anti--996-red)

# Usage

```
Usage of ./gexpose:
  -server
        server mode
  -k string
        encryption key (default "Xn2r4u7x!A%D*G8")
  -l string
        local address (default ":9000")
  -p string
        proxy address (default ":8701")
  -s string
        server address (default ":8702")
  -e string
        expose address (default ":8703")
  -t int
        dial timeout in seconds (default 30)

```

## Build

```
sh scripts/build.sh
```

# License
[The MIT License (MIT)](https://raw.githubusercontent.com/net-byte/opensocks/main/LICENSE)
