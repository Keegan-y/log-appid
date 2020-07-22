//其他数据类型转化为字节数组，再将字节数组转化为go的数据类型。
package main
import (
	"encoding/json"
	"fmt"
)
//type User struct{
//	Id int `json:"id"`
//	Name string `json:name`
//}

func main(){
	//字符串
	str_obj := `{
  "sync_info": {
      "latest_block_hash": "1414FC7313F85D7451831D029295F729F4CA3D07F18B21621BC18BF16392FBCD",
      "latest_app_hash": "198B627EF0CDE9144CFBC566986C56CAEA75F12A6D66A7345B32059D2B9E9760",
      "latest_block_height": "598995",
      "latest_block_time": "2020-01-29T17:01:11.450658236Z",
      "catching_up": true
    }
}`
	//var user User
	//json.Unmarshal([]byte(str_obj),&user)
	//字符串反解析为结构体

	//fmt.Printf("%v\n",user)
	var d map[string]interface{}
	//将字符串反解析为字典
	json.Unmarshal([]byte(str_obj),&d)
	//fmt.Printf("%v\n",d)
	//打印字典里的值
	//for key,value := range d{
	//	fmt.Printf("key:%v,value:%v\n",key,value)
	//}
	//key:sync_info,value是个字典，再次[]取值
	fmt.Printf("latest_block_height:%v\n%v\n",d["sync_info"].(map[string]interface{})["latest_block_height"],d["sync_info"].(map[string]interface{})["latest_block_height"],d["sync_info"].(map[string]interface{})["latest_block_height"].(string))
	//断言===空接口转化为具体的类型
	//Invalid operation: d["sync_info"][""] (type interface{} does not support indexing)
	// Inspection info: Reports invalid index and slice expressions
}
