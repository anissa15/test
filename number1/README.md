# Decimal to/from Binary Conversion

This application provide api to convert decimal to or from binary.
Run this app using this command : 

```bash
$ go run main.go
```

---

## Pre-requisites

This app is built using Go Programming Language, verify that you've installed Go by using this command : 

```bash
$ go version
```

If Go is not yet installed, please click bellow link and follow the steps to download and install Go : 
https://go.dev/doc/install

---

## List of Curl

### Decimal to Binary

```bash
curl -X POST http://localhost:8080/convert -d '{"type":"decimal","value":"19"}'
```

### Binary to Decimal

```bash
curl -X POST http://localhost:8080/convert -d '{"type":"binary","value":"10011"}'
```