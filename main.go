package main

import (
	"fmt"
	_ "log"
	_ "os/exec"
)

type UserLive struct {
	Id string
}
type CurrentLivesResponse struct {
	Data []UserLive `json:"data"`
}

func main() {
	fmt.Println("Server start")
	//cmd := exec.Command("/bin/sh", "-c", "cd E:\\twitch-messages\\twitch-message-bot; go run main.go")

	//cmd := exec.Command("go run E:\\twitch-messages\\twitch-message-bot\\main.go")
	//if err := cmd.Run(); err != nil {
	//	fmt.Println(err)
	//}
	//req := fasthttp.AcquireRequest()
	//req.SetRequestURI("https://api.twitch.tv/helix/streams")
	//req.Header.Add("Client-Id", "7vkqcgo3gt8z2uioc4hdyru7mczov2")
	//req.Header.Add("Authorization", "Bearer 8aehbnvpiwurj2ja780kc3m8902ol8")
	//
	//resp := fasthttp.AcquireResponse()
	//client := &fasthttp.Client{}
	//client.Do(req, resp)
	//
	//bodyBytes := resp.Body()
	//
	////println(string(bodyBytes))
	//var data CurrentLivesResponse
	//err := json.Unmarshal(bodyBytes, &data)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//itest := 1
	//for true {
	//	println("running: " + strconv.Itoa(itest))
	//	time.Sleep(5 * time.Second)
	//	itest++
	//}
}
