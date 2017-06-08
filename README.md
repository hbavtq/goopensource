# NTLM Proxy Example

start proxy (vendor directory) with user credentials in flags.

## Client

```shell
go run main.go
```

## Proxy

```shell
cd vendor/github.com/jbussdieker/go-ntlm-proxy
go run main.go -user=hbdev -domain=hbdev-PC -password=123456
```
