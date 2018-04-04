package main

import (
	"PGRpcCheckDemo/app/bigPackData"
	"PGRpcCheckDemo/app/proto"
	"framework/rpcclient/core"
	"putil/log"
	"runtime"
	_ "strconv"
	"strings"

	_ "github.com/astaxie/beego/config"
)

var (
	quitFlag chan bool //退出标记
	svrtype  int64
	svrid    int64
	version  int32
	groupids []int64
	svrname  string
	RegFuncs []string //所有注册的方法= {"Add", "Multiply"}
	netIp    string
	netPort  int
)

type BigDataHandle struct {
	client *rpcclient.RpcCall
}

func (data *BigDataHandle) RpcRequest(req *rpcclient.RpcRecvReq, body []byte) {
	serviceAndMethodName := string(req.Rpchead.MethodName) //方法名称(eg：User.getUserInfo)
	var methodName string
	//转成纯方法名称
	pos := strings.LastIndex(serviceAndMethodName, ".")
	if pos > 0 && pos < len(serviceAndMethodName) {
		methodName = serviceAndMethodName[pos+1:] //避免越界
	}
	//	if len(methodName) == 0 {
	//		//参数异常的情况下
	//	}

	rt := []byte{} //最终返回的字节切片
	//TODO检查方法名称是否存在
	switch methodName {
	case "GetData":
		//接收参数并处理
		bigdata := new(bigPackData.BigPackData)
		bigDataReq := new(RPCProto.BigPackDataRequest)

		bigDataReq.Unmarshal(body) //解参数的pb包
		bigDataResp := bigdata.GetData(bigDataReq)
		plog.Debug("===mybigdata", methodName)
		rt, _ = bigDataResp.Marshal() //返回数据的pb格式化
	default:
		//rt := []byte{}
	}

	data.client.SendPacket(req, rt) //处理完毕之后返回数据！
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	quitFlag = make(chan bool)
	//实例化rpc
	client, err := rpcclient.NewRpcCall()
	if err != nil {
		plog.Fatal("fatal")
		return
	}

	err = client.RpcInit("./conf/app-test.conf")

	if err != nil {
		plog.Debug("rpc 初始化异常")
		return
	}

	mydisp := new(BigDataHandle)
	mydisp.client = client
	//设置本rpc服务的代理人（Net层）的ip和端口，并启动服务！
	err = client.LaunchRpcClient(mydisp)
	if err != nil {
		plog.Fatal("lauch failed", err)
		return
	}

	//控制退出
	_ = <-quitFlag //收消息退出！
}