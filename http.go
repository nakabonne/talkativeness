package talkativeness

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func LogRequest(r *http.Request, body bool) {
	dump, err := httputil.DumpRequest(r, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\n>>Request message<<\n")
	fmt.Println(string(dump))
}

func LogResponse(resp *http.Response, body bool) {
	dump, err := httputil.DumpResponse(resp, body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\n>>Response message<<\n")
	log.Println(string(dump))
}
