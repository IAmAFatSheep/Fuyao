package main

import (
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

const DOC_TYPE = "FyObj"

// 保存信息
// args: Fuyao
func PutFy(stub shim.ChaincodeStubInterface, fy Fuyao) ([]byte, bool) {

	fy.ObjectType = DOC_TYPE

	b, err := json.Marshal(fy)
	if err != nil {
		return nil, false
	}

	// 保存信息状态
	err = stub.PutState(fy.CodeNo, b)
	if err != nil {
		return nil, false
	}

	return b, true
}

// 根据编码号查询信息状态
// args: codeId
func GetFyInfo(stub shim.ChaincodeStubInterface, codeId string) (Fuyao, bool) {
	var fy Fuyao
	// 根据编码号查询信息状态
	b, err := stub.GetState(codeId)
	if err != nil {
		return fy, false
	}

	if b == nil {
		return fy, false
	}

	// 对查询到的状态进行反序列化
	err = json.Unmarshal(b, &fy)
	if err != nil {
		return fy, false
	}

	// 返回结果
	return fy, true
}

// 添加信息
// args: FuyaoObject
// 编码号为 key, Fuyao 为 value
func (t *FuyaoChaincode) addInfo(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 2 {
		return shim.Error("给定的参数个数不符合要求")
	}

	var fy Fuyao
	err := json.Unmarshal([]byte(args[0]), &fy)
	if err != nil {
		return shim.Error("反序列化信息时发生错误")
	}

	// 查重: 编码号码必须唯一
	_, exist := GetFyInfo(stub, fy.CodeNo)
	if exist {
		return shim.Error("要添加的编码号已存在")
	}

	_, bl := PutFy(stub, fy)
	if !bl {
		return shim.Error("保存信息时发生错误")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("信息添加成功"))
}

// 根据编码号查询详情
// args: entityID
func (t *FuyaoChaincode) queryFyInfoByCodeID(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("给定的参数个数不符合要求")
	}

	// 根据编码号查询fy状态
	b, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("根据编码号查询信息失败")
	}

	if b == nil {
		return shim.Error("根据编码号没有查询到相关的信息")
	}

	// // 对查询到的状态进行反序列化
	// var fy Fuyao
	// err = json.Unmarshal(b, &fy)
	// if err != nil {
	// 	return shim.Error("反序列化fy信息失败")
	// }

	// // 返回
	// result, err := json.Marshal(edu)
	// if err != nil {
	// 	return shim.Error("序列化edu信息时发生错误")
	// }
	return shim.Success(b)
}

// 根据编码号删除信息
// args: codeId
func (t *FuyaoChaincode) delFy(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("给定的参数个数不符合要求")
	}

	err := stub.DelState(args[0])
	if err != nil {
		return shim.Error("删除信息时发生错误")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("信息删除成功"))
}
