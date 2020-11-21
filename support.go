package config

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// // SwitchBack ...
// func SwitchBack() {

// 	os.Chdir(CurrDIR)
// }

// PrintStructsPretty ...
func PrintStructsPretty(c interface{}) {
	fmt.Println(reflect.TypeOf(c), c)
	b, err := json.MarshalIndent(c, "", "    ")
	if err == nil {
		fmt.Println(reflect.TypeOf(c), " struct:")
		fmt.Println(string(b))
	}
}

// // SwitchInput ...
// func SwitchInput() {
// 	// GOPATH := os.Getenv("GOPATH")
// 	// InDIR = GOPATH + "/src/github.com/5gif/" + indir
// 	log.Println("Switching 2 Input : ", InDIR)
// 	os.Chdir(InDIR)
// }

// // SwitchOutput ...
// func SwitchOutput() {
// 	// GOPATH := os.Getenv("GOPATH")
// 	log.Println("Switching 2 Output : ", OutDIR)
// 	os.Chdir(OutDIR)
// }
