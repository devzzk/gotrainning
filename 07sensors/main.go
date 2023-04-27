package main

import (
	sdk "github.com/sensorsdata/sa-sdk-go"
	"log"
)

func main() {
	sa_service_url := ""
	var consumer, err = sdk.InitDefaultConsumer(sa_service_url, 10000)
	if err != nil {
		log.Println(err)
	}
	sa := sdk.InitSensorsAnalytics(consumer, "default", false)
	defer sa.Close()
	properties := map[string]interface{}{
		"price": 12,
		"name":  "apple",
	}
	err = sa.Track("ABCDEFG1234567", "ViewProduct", properties, false)

}
