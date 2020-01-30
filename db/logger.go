package db

import "fmt"

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

//Check ... for non nil errors and prints
func Check(err error) {
	if err != nil {
		fmt.Println("-- Error: ", err)
	}
}
