# [Ecdsa](https://learn.akamai.com/en-us/webhelp/iot/jwt-access-control/GUID-C3B1D111-E0B5-4B3B-9FF0-06D48CF40679.html)  
### Create a private key.
`openssl ecparam -list_curves`  
`openssl ecparam -genkey -name prime256v1 -noout -out ec256-private.pem`
### Create a public key by extracting it from the private key.
`openssl ec -in ec256-private.pem -pubout > ec256-public.pem`

# [Sign and verify](https://github.com/dgrijalva/jwt-go/tree/master/cmd/jwt)  

```shell
go run github.com/dgrijalva/jwt-go/cmd/jwt
# go build -o jwt ${GOPATH}/pkg/mod/github.com/dgrijalva/jwt-go@v3.2.0+incompatible/cmd/jwt/*.go
echo {\"someone\":\"foobar\"} | \
./jwt -key ./keypair/ec256-private.pem -alg ES256 -sign - | \
./jwt -key ./keypair/ec256-public.pem  -alg ES256 -verify -
```

## Test code
https://github.com/dgrijalva/jwt-go/blob/master/ecdsa_test.go
