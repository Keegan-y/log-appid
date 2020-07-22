package main

import "fmt"

func main() {
	/*
	a := `{"name":"chalvern.github.io","full_name":"chalvern/chalvern.github.io","private":false,"owner":{"login":"chalvern","html_url":"https://github.com/chalvern"},"html_url":"https://github.com/chalvern/chalvern.github.io","description":"jingwei.link blog"}`

	// OwnerS 会被嵌套在 JingweiS 中
	// 在每个字段后的 `json:xxx` 即 tag，与 json 中的 key 对应
	type OwnerS struct {
		Login   string `json:"login"`
		HTMLURL string `json:"html_url"`
	}

	//  JingweiS 中嵌套了 OwnerS 用来对应 a 变量 json 字符串中的嵌套的 json 内容
	type JingweiS struct {
		Name        string `json:"name"`
		FullName    string `json:"full_name"`
		Private     bool   `json:"private"`
		Owner       OwnerS `json:"owner"`
		HTMLURL     string `json:"html_url"`
		Description string `json:"description"`
	}

	var jingwei JingweiS
	err := json.Unmarshal([]byte(a), &jingwei)
	fmt.Printf("%#v \n %#v \n", err, jingwei)
	// 接下来可以通过 “.” 运算符获取期望的值了
	fmt.Printf("%#v\n", jingwei.Owner.Login)

	 */
	fmt.Printf("%v", len("app-755ba73e-bd7b-4352-a8aa-b06422681695"))//40 34
}
