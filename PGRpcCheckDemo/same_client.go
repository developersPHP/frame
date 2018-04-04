package main

import (
	_ "PGRpcCheckDemo/app/bigPackData"
	"PGRpcCheckDemo/app/controllers"
	"PGRpcCheckDemo/app/proto"
	"fmt"
	"framework/rpcclient/core"
	"putil/log"
	_ "time"
)

type MyBigPackData struct {
	client *rpcclient.RpcCall
}

func (dis *MyBigPackData) RpcRequest(req *rpcclient.RpcRecvReq, body []byte) {
	//plog.Debug(req)
	//plog.Debug(body)

	responsestr := "hi I'm Server return your data"
	dis.client.SendPacket(req, []byte(responsestr))
}

var (
	count int = 4
	chs       = make([]chan *rpcclient.RpcResponse, count)
)

func main() {
	client, err := rpcclient.NewRpcCall()
	if err != nil {
		plog.Fatal("fatal")
	}

	err = client.RpcInit("./conf/clientapp.conf")
	if err != nil {
		plog.Debug("rpc 初始化异常")
		return
	}

	mydisp := new(MyBigPackData)
	mydisp.client = client
	err = client.LaunchRpcClient(mydisp)

	//err = client.LaunchRpcClient("192.168.100.126", 9081, mydisp)
	if err != nil {
		plog.Fatal("lauch failed", err)
		return
	}
	plog.Debug("LaunchRpcClient succeed!!!")
	for i := 1; i < 4; i++ {
		request := new(RPCProto.BigPackDataRequest)
		request.ContainerReq = createBigData(i)
		req_bytes, err := request.Marshal()
		if err != nil {
			fmt.Println(req_bytes)
		}
		funname := fmt.Sprintf("bigdata.GetData_%d", i)
		plog.Debug("funname is ", funname)
		for j := 1; j < 100; j++ {
			res := client.SendAndRecvRespRpcMsg(funname, req_bytes, 5000, 0)
			if res.ReturnCode != 0 {
				//rpc返回结果异常
				//RPC_RESPONSE_COMPLET         = 0 //完成
				//	RPC_RESPONSE_TIMEOUT         = 1 //超时
				//	RPC_RESPONSE_SENDFAILED      = 2 //发送错误
				//	RPC_RESPONSE_NETERR          = 3 //发生网络错误
				//	RPC_RESPONSE_TARGET_NOTFOUND = 4 //Net层没有发现目标实例
				plog.Debug("rpc return code = ", res.ReturnCode, " return err = ", res.Err)
			} else {
				arithResp := new(RPCProto.BigPackDataResponse)
				arithResp.Unmarshal(res.Body)
				plog.Debug("return value  = ", arithResp)
			}
		}
	}

	var controlstr string
	fmt.Scanln(&controlstr)

	if controlstr == "q" || controlstr == "Q" {
		return
	}
}

func createBigData(num int) map[string]string {
	bigPack := new(controllers.BigPack)
	bigPack.Route = 5 + num
	bigPack.Container = make(map[string]string)
	container := bigPack.CreateBigPack()
	return container
}