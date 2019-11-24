# Proto gen

```
go get -u github.com/golang/protobuf/protoc-gen-go
go build github.com/golang/protobuf/protoc-gen-go
set -gx PATH $GO_PATH/bin:$PATH
# From proto dir in home_led
./bin/protoc --go_out="$GO_PATH/src/github.com/thomaschandler/home_led_srv" led.proto
```

*TODO*: https://github.com/golang/protobuf#packages-and-input-paths Fix package
name at compile stage

# Building

```
go get github.com/tarm/serial
go get github.com/brutella/hc
go get github.com/brutella/hc/accessory
```

## For OpenWRT Based Router

```
env GOOS=linux GOARCH=arm GOARM=7 go build
```

scp resulting binary to router

On router:

```
opkg update && opkg install kmod-usb-acm
./home_led_srv &
```

Exit SSH session

# References

- [GOOS, GOARCH, GOARM](https://www.thepolyglotdeveloper.com/2017/04/cross-compiling-golang-applications-raspberry-pi/)
