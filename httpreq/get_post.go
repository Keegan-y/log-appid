package main
import (
	"fmt"
	"io/ioutil"
	"net/http"
)
//func httpGet() {
//	resp, err :=   http.Get("http://www.01happy.com/demo/accept.php?id=1")
//	if err != nil {
//		// handle error
//	}
//
//	defer resp.Body.Close()
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		// handle error
//	}
//
//	fmt.Println(string(body))
//}

func main(){
	resp, err :=   http.Get("http://192.168.1.39:20251/app/fbcba2b8-c56b-4483-8fd7-038f2c632192/metrics")
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}