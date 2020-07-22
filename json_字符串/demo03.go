package main
import (
	"encoding/json"
	"fmt"
)

func main(){
	obj := `{"name":"jg","full_name":"yjg","sex":"fale","age":18}`
	type Info struct{
		Name string `json:"name"`
		Full_name string `json:"full_name"`
		Sex string `json:"sex"`
		Age int `json:"age"`
	}
	var yjg Info
	err := json.Unmarshal([]byte(obj),&yjg)
	fmt.Printf("%v\n %v\n",err,yjg)
}
