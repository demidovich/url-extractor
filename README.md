## url extractor

Utility for parsing http href tags from url. Skips static links. Converts relative paths to absolute.

#### Example of usage

```
go run url_extractor.go https://gmail.com

Request https://gmail.com
Total URLs found: 8

https://accounts.google.com/v3/signin/
https://accounts.google.com/v3/signin/identifier
https://gmail.com/signin/usernamerecovery?continue=https://mail.google.com/mail/u/0/&amp;dsh=S889468511:1657453848964432&amp;emr=1&amp;flowEntry=ServiceLogin&amp;flowName=WebLiteSignIn&amp;followup=https://mail.google.com/mail/u/0/&amp;ifkv=AX3vH38jdZVWkMjAg6cqMpP7yqa1Fn4a5kv6UOAFHUofMstggX0PSuV56i1y80cG_0az6klthXWW6g&amp;osid=1&amp;service=mail
https://support.google.com/accounts?p=signin_privatebrowsing&amp;hl=en-US
https://gmail.com/SignUp?continue=https://mail.google.com/mail/u/0/&amp;dsh=S889468511:1657453848964432&amp;emr=1&amp;flowEntry=ServiceLogin&amp;flowName=WebLiteSignIn&amp;followup=https://mail.google.com/mail/u/0/&amp;ifkv=AX3vH38jdZVWkMjAg6cqMpP7yqa1Fn4a5kv6UOAFHUofMstggX0PSuV56i1y80cG_0az6klthXWW6g&amp;osid=1&amp;service=mail
https://support.google.com/accounts?hl=en-US
https://accounts.google.com/TOS?loc=RU&amp;hl=en-US&amp;privacy=true
https://accounts.google.com/TOS?loc=RU&amp;hl=en-US
```
