package main

import (
	"context"
	"musicRoomBookingbot/application"
)

// 호스트 관리 API
// 사용자 API
func main() {
	ctx, cf := context.WithCancel(context.Background())
	app := application.GetInstance(ctx, cf)

	endConsume := make(chan bool)

	go app.ShutdownService()
	go app.StartService(endConsume)

	<-endConsume

	<-app.Ctx.Done()
}
