package main

import (
	"fmt"
	"net/http"
	"strconv"
)

//CheckCookie .. checks cookies and increment count value on incoming requests
func CheckCookie(res http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("my-cookie")

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "my-cookie",
			Value: "0",
		}
	}

	count, _ := strconv.Atoi(cookie.Value)
	count++
	cookie.Value = strconv.Itoa(count)

	http.SetCookie(res, cookie)

	fmt.Println(cookie)

}
