# Load Balancer

###### Configuration file `conf.json`
```json
{
  "name": "Load Balancer",
  "listener": ":80",
  "servers": [
    "0.0.0.0:8080",
    "0.0.0.0:8081",
    "0.0.0.0:8082",
    "0.0.0.0:8083",
    "0.0.0.0:8084"
  ]
}
```

## Installation

```shell
git clone https://github.com/beyrakIn/lb.git
cd lb
go build -o lb
./lb
```