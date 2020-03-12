package main

import (
	"FYP_Backend/db"
	"fmt"
	"net/http"
	"strconv"
)

//CheckCookie .. checks cookies and increment count value on incoming requests
func CheckCookie(res http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("cn-cookie")

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "cn-cookie",
			Value: "0",
		}
	}

	count, err := strconv.Atoi(cookie.Value)
	db.Check(err)

	count++
	cookie.Value = strconv.Itoa(count)

	http.SetCookie(res, cookie)

	fmt.Println("// ", cookie)

}
