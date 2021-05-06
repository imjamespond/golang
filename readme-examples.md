
# Example requests
### oauth2.ClientCredentials 无法获取userId, 不适合生成jwt token
```
http://localhost:8001/token?grant_type=client_credentials&client_id=000000&client_secret=999999&scope=read
```
### oauth2.AuthorizationCode
client call
```
curl --location --request GET \
'http://192.168.1.101:8001/authorize-code?client_id=000000&redirect_uri=http%3A%2F%2Flocalhost%3A8080%2Foauth2%2FgetToken&response_type=code&state=12345&userId=1001' \
--header 'Authorization: Basic MDAwMDAwOjk5OTk5OQ==' -v

# redirect to app and call, redirect_uri must be matched the above redirect_uri
curl -i -X POST \
   -H "Content-Type:application/x-www-form-urlencoded" \
   -d "grant_type=authorization_code" \
   -d "code=CODE_RETURNED_BY_AUTHCODE..." \
   -d "state=23456" \
   -d "client_id=000000" \
   -d "client_secret=999999" \
   -d "redirect_uri=http://localhost:8080/oauth2/getToken" \
 'http://localhost:8001/token'
 
curl -i -X GET \
   -H "Authorization:Bearer ACCESS_TOKEN" \
 'http://localhost:8001/test'

```

### oauth2.Refreshing
```
curl -i -X POST \
   -H "Content-Type:application/x-www-form-urlencoded" \
   -d "grant_type=refresh_token" \
   -d "client_id=000000" \
   -d "client_secret=999999" \
   -d "refresh_token=REFRESH_TOKEN..." \
 'http://localhost:8001/token'
```