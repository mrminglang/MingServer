package main

import (
	"fmt"

	"gitlab.upchinaproduct.com/taf/tafgo/taf"
	"gitlab.upchinaproduct.com/taf/tafgo/taf/protocol/res/adminf"
)

func main() {
	comm := taf.NewCommunicator()
	obj := "MingApp.MingServer.MingHelloObj@tcp -h 127.0.0.1 -p 10014 -t 60000"
	app := new(adminf.AdminF)
	comm.StringToProxy(obj, app)
	ret, err := app.Notify("taf.dumpstack")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret)
}
