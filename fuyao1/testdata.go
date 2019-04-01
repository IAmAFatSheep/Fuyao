package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/fuyao/service"
)

func WriteData(i int, serviceSetup service.ServiceSetup) {
	fmt.Println("i_value:" + strconv.Itoa(i))
	fy := service.Fuyao{
		Name:     strconv.Itoa(i),
		CodeNo:   strconv.Itoa(i),
		BirthDay: time.Now(),
		Deadline: i,
		Photo:    strconv.Itoa(i),
	}

	_, err := serviceSetup.SaveInfo(fy)
	// msg, err := serviceSetup.SaveEdu(edu)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		// fmt.Println("信息发布成功, 交易编号为: " + msg)
	}
	ch <- 0
}
