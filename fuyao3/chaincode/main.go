package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type FuyaoChaincode struct {
}

func (t *FuyaoChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {

	return shim.Success(nil)
}

func (t *FuyaoChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	// 获取用户意图
	fun, args := stub.GetFunctionAndParameters()

	if fun == "addInfo" {
		return t.addInfo(stub, args) // 添加信息
	} else if fun == "queryFyInfoByCodeID" {
		return t.queryFyInfoByCodeID(stub, args) // 根据编码号查询信息
	} else if fun == "delFy" {
		return t.delFy(stub, args) // 根据编码号删除信息
	}

	return shim.Error("指定的函数名称错误")

}
func main() {
	err := shim.Start(new(FuyaoChaincode))
	if err != nil {
		fmt.Printf("启动FuyaoChaincode时发生错误: %s", err)
	}
}
