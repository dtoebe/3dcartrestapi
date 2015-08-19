package main

import (
	"fmt"

	"github.com/dtoebe/3dcartrestapi/receiver"
)

func main() {
	res := receiver.NewConf("https://www.yourstoreURL.com", "AppPrivateKey", "AppToken", 1 /*API Version Number Int */)

	status, header, body := res.GetData("Products")

	fmt.Println("Status:", status)
	fmt.Println("Header:", header)
	fmt.Println("Body:", string(body))
}
