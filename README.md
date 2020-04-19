```
  ____ _____     _____ ____         ____ _     ___
 / ___/ _ \ \   / /_ _|  _ \       / ___| |   |_ _|
| |  | | | \ \ / / | || | | |_____| |   | |    | |
| |__| |_| |\ V /  | || |_| |_____| |___| |___ | |
 \____\___/  \_/  |___|____/       \____|_____|___|
```

Get COVID-19 stats on your terminal

## Description
This terminal command allow you to check stats of COVID-19 on your terminal.

This tools uses the API COVID19: [https://covid19api.com/](https://covid19api.com/)

## How to use it?
```bash
➜ go run covidcli.go --help
NAME:
   covidcli - Get COVID-19 stats on your terminal

USAGE:
   covidcli [global options] command [command options] [arguments...]

VERSION:
   2020.04.19

AUTHOR:
   Ignacio Peluffo <ipeluffo@gmail.com>

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --from value     (default: (*time.Time)(nil))
   --to value       (default: (*time.Time)(nil))
   --country value
   --help, -h       show help (default: false)
   --version, -v    print the version (default: false)
```

## Authors
Ignacio Peluffo: [https://github.com/ipeluffo](https://github.com/ipeluffo)

## License
Apache 2.0: [https://github.com/ipeluffo/covidcli/blob/master/LICENSE](https://github.com/ipeluffo/covidcli/blob/master/LICENSE)
