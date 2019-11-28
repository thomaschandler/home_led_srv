

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
