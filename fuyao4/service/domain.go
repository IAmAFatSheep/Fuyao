package service

import (
	"fmt"
	"time"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
)

type Fuyao struct {
	ObjectType string    `json:"docType"`
	Name       string    `json:"Name"`     // 商品名称
	CodeNo     string    `json:"CodeNo"`   // 编码号
	BirthDay   time.Time `json:"BirthDay"` // 时间戳
	Deadline   int       `json:"SaveTime"` // 保存时间(天)
	Photo      string    `json:"Photo"`    // 照片
}

type ServiceSetup struct {
	ChaincodeID string
	Client      *channel.Client
}

func regitserEvent(client *channel.Client, chaincodeID, eventID string) (fab.Registration, <-chan *fab.CCEvent) {

	reg, notifier, err := client.RegisterChaincodeEvent(chaincodeID, eventID)
	if err != nil {
		fmt.Printf("注册链码事件失败: %s", err)
	}
	return reg, notifier
}

func eventResult(notifier <-chan *fab.CCEvent, eventID string) error {
	select {
	case ccEvent := <-notifier:
		fmt.Printf("接收到链码事件: %v\n", ccEvent)
	case <-time.After(time.Second * 20):
		return fmt.Errorf("不能根据指定的事件ID接收到相应的链码事件(%s)", eventID)
	}
	return nil
}
