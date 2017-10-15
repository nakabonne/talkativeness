# talkativeness
A package that outputs frequently used logs

# Usage

```go
import "github.com/ryonakao/talkativeness"

func handler(w http.ResponseWriter, req *http.Request) {
	talkativeness.LogRequest(req, true)
}
```

A log is output

```
>>Request message<<

GET /room HTTP/1.1
Host: localhost:8080
Accept-Encoding: gzip, deflate, br
Accept-Language: ja,en-US;q=0.8,en;q=0.6
Cache-Control: no-cache
Connection: Upgrade
Origin: http://localhost:8080
Pragma: no-cache
Sec-Websocket-Extensions: permessage-deflate; client_max_window_bits
Sec-Websocket-Key: R1RGOCJKDkHksrbfEqPuTg==
Sec-Websocket-Version: 13
Upgrade: websocket
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36
```

# Installation

```
$ go get github.com/ryonakao/talkativeness
```
