package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	//json对象是什么
//	obj := `{
//	"jsonrpc":"2.0",
//	"id": "",
//	"result": {"id":"5e8d14e4f3e134aa1ffb998e1ca1b6fd83e13fc0"},
//	"sync_info": {"latest_block_height":"598995"}
//}`
/*
A:如何获取这些对象的内容？
有些app是通过点点点的方式，找到curl ,执行，结果。
有什么缺点？在哪里？
1.手动
2.爬虫代替手动存库(非手动)
***有没有什么接口可以调用这些数据？***
B:如何快速地获取这些对象的内容？
把这个对象的内容封装在专门的文件里，文件名是app的名字*/
	resp, _ :=   http.Get("http://app-dd696c44-33d1-4716-8557-cd2ccb9e33a1.cls-08f70d3d-d9e8-4445-8a88-f62646577016.ankr.com/status")
	//if err != nil {
	//	// handle error
	//}
	//fmt.Printf(resp)
	//defer resp.Body.Close()
	obj, _:= ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	// handle error
	//}
	//fmt.Printf("obj类型：%T\n",obj)
	////fmt.Printf(obj)//[]uint8要转化字符串
	//fmt.Println(string(obj)[3])
	fmt.Printf("obj类型：%T\n",obj)
	//a:=json.Marshal(obj)
	//obj := string(obj)
	//fmt.Printf("%v,%T",string(obj),string(obj))
	//结构体标签
	//type Appinfo struct{
	//	//Jsonrpc string `json:"jsonrpc"`
	//	//Id string `json:""`
	//	//Result map[string]string `json:"result"`//这个字段的数据类型是map
	//	Sync_info map[string]string `json:"sync_info"`
	//}
	//
	//var appinfo Appinfo
	//给结构体变量赋值
	//appinfo.Sync_info = obj//类型一致
	//err := json.Unmarshal([]byte(obj),&appinfo)
	//fmt.Printf("err===》%v\n",err)
	//fmt.Printf("latest_block_height:%v",appinfo.Sync_info["latest_block_height"])
	//值取不到。因为没给结构体变量赋值。50
	//将[]uint8转化为字典
}