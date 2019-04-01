package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/fuyao/sdkInit"
	"github.com/fuyao/service"
	"github.com/fuyao/web"
	controller "github.com/fuyao/web/appController"
)

const (
	configFile  = "config.yaml"
	initialized = false
	FyCC        = "chainhero"
)

var ch chan int = make(chan int)

func main() {

	initInfo := &sdkInit.InitInfo{

		ChannelID:     "chainhero",
		ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/fuyao/fixtures/artifacts/chainhero.channel.tx",

		OrgAdmin:       "Admin",
		OrgName:        "Org1",
		OrdererOrgName: "orderer1.hf.chainhero.io",

		ChaincodeID:     FyCC,
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath:   "github.com/fuyao/chaincode/",
		UserName:        "User1",
	}

	sdk, err := sdkInit.SetupSDK(configFile, initialized)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	defer sdk.Close()

	err = sdkInit.CreateChannel(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	channelClient, err := sdkInit.InstallAndInstantiateCC(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(channelClient)

	//===========================================//

	serviceSetup := service.ServiceSetup{
		ChaincodeID: FyCC,
		Client:      channelClient,
	}

	//tps测试
	runtime.GOMAXPROCS(1)
	t0 := time.Now()
	for i := 300; i < 400; i++ {
		go WriteData(i, serviceSetup)

	}
	for i := 300; i < 400; i++ {
		<-ch
	}
	endTime := time.Since(t0)
	fmt.Println("运行时间：", endTime)

	//===========================================//
	app := controller.Application{
		Setup: &serviceSetup,
	}
	web.WebStart(app)
}
