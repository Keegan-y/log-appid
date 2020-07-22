package main

import "fmt"

func main(){
//	resp, err:=   http.Get("http://app-dd696c44-33d1-4716-8557-cd2ccb9e33a1.cls-08f70d3d-d9e8-4445-8a88-f62646577016.ankr.com/status")
//	//fmt.Printf("%v,%T",&resp,resp)//0xc000006028,*http.Response
//	if err != nil {
//		return
//	}
//	//panic: runtime error: invalid memory address or nil pointer dereference.所以添加if判断
//	defer resp.Body.Close()
//	obj, err:= ioutil.ReadAll(resp.Body)
//	if err != nil {
//		// handle error
//	}
//	fmt.Println(string(obj))
//	//fmt.Printf("%v",obj)
	//obj_0 := {"jsonrpc": "2.0"}go语言有这种数据定义吗？没有。
	//obj_0 := "'protocol_version': {"
	//fmt.Printf("%T",obj_0)
	obj := `{
  "jsonrpc": "2.0",
  "id": "",
  "result": {
    "node_info": {
      "protocol_version": {
        "p2p": "7",
        "block": "10",
        "app": "0"
      },
      "id": "5e8d14e4f3e134aa1ffb998e1ca1b6fd83e13fc0",
      "listen_addr": "tcp://0.0.0.0:26656",
      "network": "cosmoshub-3",
      "version": "0.32.7",
      "channels": "4020212223303800",
      "moniker": "ankr-test",
      "other": {
        "tx_index": "on",
        "rpc_address": "tcp://0.0.0.0:26657"
      }
    },
    "sync_info": {
      "latest_block_hash": "1414FC7313F85D7451831D029295F729F4CA3D07F18B21621BC18BF16392FBCD",
      "latest_app_hash": "198B627EF0CDE9144CFBC566986C56CAEA75F12A6D66A7345B32059D2B9E9760",
      "latest_block_height": "598995",
      "latest_block_time": "2020-01-29T17:01:11.450658236Z",
      "catching_up": true
    },
    "validator_info": {
      "address": "F5A53AEA21D8190FD9882BA5C3320B127399E5E1",
      "pub_key": {
        "type": "tendermint/PubKeyEd25519",
        "value": "blYkrmnAX4HsUf/nOyFh4cr09dXU7qT+TLigbFgSf2k="
      },
      "voting_power": "0"
    }
  }
}
`
	//``支持多行，不支持转义字符。""不支持多行，支持转义字符。而且使用过程中会有符号嵌套问题。
	fmt.Printf("%T",obj)//string。现在数据对象是go的string
	fmt.Printf("%v",obj)
	//如何通过字符串拿到块高？
	//通过循环取值字符串？
	for i:=0;i<=10;i++{
		fmt.Printf("%v",string(obj[i]))
	}
	//切片取值
	fmt.Printf("%v",string(obj[0:11]))//obj[6]115
	//结果都是{
	//  "jsonrp
	//将数据对象转化为字典/集合的形式取值。怎么做？找到自己合适的位置。


}
/*
{"app":{"id":"app-5f693ebd-8de1-4696-965b-921eb8e3d5f4","name":"Kusama-Test","creator":"271f5165-0822-47b3-8fba-1d5c4c55b549","team_id":"271f5165-0822-47b3-8fba-1d5c4c55b549","namespace":{"id":"ns-df1c83d3-0ce7-43f7-bade-4cf1632e0ce7","name":"","cluster_id":"cls-dec3c32b-4f06-462f-b827-dee931d39a72","cpu_limit":"2000","mem_limit":"8000","storage_limit":"80000","cpu_used":"0","mem_used":"0","storage_used":"0","status":"NAMESPACE_AVAILABLE","created_time":"2020-07-19T06:34:51.831Z","updated_time":"2020-07-19T06:34:51.831Z"},"chart":{"id":"kusama","name":"kusama","description":"Full Node","repository":"stable","icon_url":"https://assets.ankr.com/charts/icon-only/kusama.svg","version":"1.14.0","app_version":"1.14.0","app_hooks":[{"label":"apphook1=node-name","container":"kusama-node","action":"query node name","description":"query node name","command":"cat /polkadot/kusama/node_name.txt","args":{}},{"label":"apphook2=archive-mode","container":"kusama-node","action":"query node archive mode","description":"query node archive mode","command":"cat /polkadot/kusama/archive_mode.txt","args":{}},{"label":"apphook3=node-status","container":"kusama-node","action":"query node status","description":"query node status","command":"bash /polkadot/scripts/node_status.sh","args":{}},{"label":"apphook4=query-chain-properties","container":"kusama-node","action":"query chain properties","description":"query chain properties","command":"node /polkadot/scripts/query_chain_properties.js 2\u003e/dev/null","args":{}},{"label":"apphook5=query-node-status","container":"kusama-node","action":"query node syncing status","description":"query node syncing status","command":"node /polkadot/scripts/query_node_status.js 2\u003e/dev/null","args":{}}]},"status":"APP_RUNNING","resource":"==\u003e v1/Deployment\nNAME                                              AGE\napp-5f693ebd-8de1-4696-965b-921eb8e3d5f4-kusama   19h\n\n==\u003e v1/PersistentVolumeClaim\nNAME                                                  AGE\napp-5f693ebd-8de1-4696-965b-921eb8e3d5f4-kusama-pvc   19h\n\n==\u003e v1/Pod(related)\nNAME                                                              AGE\napp-5f693ebd-8de1-4696-965b-921eb8e3d5f4-kusama-7479b8c84dqv896   19h\n\n==\u003e v1/Service\nNAME                                                  AGE\napp-5f693ebd-8de1-4696-965b-921eb8e3d5f4-kusama-p2p   19h\napp-5f693ebd-8de1-4696-965b-921eb8e3d5f4-kusama-rpc   19h\n\n==\u003e v1beta1/Ingress\nNAME                                              AGE\napp-5f693ebd-8de1-4696-965b-921eb8e3d5f4-kusama   19h\n\n","note":"\n\nGet the Kusama RPC URL:\n\n  You should be able to access your new Kusama RPC through\n  http://app-5f693ebd-8de1-4696-965b-921eb8e3d5f4.cls-dec3c32b-4f06-462f-b827-dee931d39a72.ankr.com\n\n  See Websocket version by\n\n  ws://app-5f693ebd-8de1-4696-965b-921eb8e3d5f4.cls-dec3c32b-4f06-462f-b827-dee931d39a72.ankr.com/ws\n","endpoint":"app-5f693ebd-8de1-4696-965b-921eb8e3d5f4.cls-dec3c32b-4f06-462f-b827-dee931d39a72.ankr.com","node_ports":[],"events":[],"created_time":"2020-07-19T06:34:51.830Z","updated_time":"2020-07-20T02:31:48.013Z"}}

{
  "jsonrpc": "2.0",
  "id": "",
  "result": {
    "node_info": {
      "protocol_version": {
        "p2p": "7",
        "block": "10",
        "app": "0"
      },
      "id": "5e8d14e4f3e134aa1ffb998e1ca1b6fd83e13fc0",
      "listen_addr": "tcp://0.0.0.0:26656",
      "network": "cosmoshub-3",
      "version": "0.32.7",
      "channels": "4020212223303800",
      "moniker": "ankr-test",
      "other": {
        "tx_index": "on",
        "rpc_address": "tcp://0.0.0.0:26657"
      }
    },
    "sync_info": {
      "latest_block_hash": "1414FC7313F85D7451831D029295F729F4CA3D07F18B21621BC18BF16392FBCD",
      "latest_app_hash": "198B627EF0CDE9144CFBC566986C56CAEA75F12A6D66A7345B32059D2B9E9760",
      "latest_block_height": "598995",
      "latest_block_time": "2020-01-29T17:01:11.450658236Z",
      "catching_up": true
    },
    "validator_info": {
      "address": "F5A53AEA21D8190FD9882BA5C3320B127399E5E1",
      "pub_key": {
        "type": "tendermint/PubKeyEd25519",
        "value": "blYkrmnAX4HsUf/nOyFh4cr09dXU7qT+TLigbFgSf2k="
      },
      "voting_power": "0"
    }
  }
}

*/
