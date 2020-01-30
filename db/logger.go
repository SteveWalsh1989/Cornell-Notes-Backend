package db

import "fmt"

/*
 Logger:
 --------------------------------------------
 file contains functions which can be used to
 print helpful information to the console when
 the application is running.
*/

var sps = "\n//----------------------//\n//   "
var spe = "   //\n//   "

//LogTitle .. basic logging function that highlights a title to console for readability
func LogTitle(str string) {
	var i int
	fmt.Print("//---")
	for i = 0; i < 30; i++ {
		fmt.Print("-")
	}
	fmt.Print("----")
	fmt.Println("\n//  ", str)

}

//LogValue .. basic logging function that highlights a value to console for readability
func LogValue(str string, value string) {
	var i int
	fmt.Print("//---")
	for i = 0; i < 30; i++ {
		fmt.Print("-")
	}
	fmt.Print("----")
	fmt.Println("\n//  ", str+": "+value)

}

//LogDBConn .. basic logging function that highlights a value to console for readability
func LogDBConn(str string) {
	var i int
	fmt.Print("//---")
	for i = 0; i < 30; i++ {
		fmt.Print("-")
	}
	fmt.Print("----")
	fmt.Println("\n**  ", str)

}

//LogRequest .. logs URL requests
func LogRequest(method string, url string) {
	var i int
	fmt.Print("//---")
	for i = 0; i < 30; i++ {
		fmt.Print("-")
	}
	fmt.Print("----")
	fmt.Println("\n//  Request Method:", method)
	fmt.Println("//  Request URL:", url)

}

//LogRequestSuccess .. logs success
func LogRequestSuccess() {
	fmt.Println("\n//  Request Sent Successfully")

}

//Check ... for non nil errors and prints
func Check(err error) {
	if err != nil {
		fmt.Println("-- Error: ", err)
	}
}
