package main

import (
	"app/pkg/config"
	"fmt"

	"github.com/slack-go/slack"
)

func main() {
	api := slack.New(config.SLACK_BOT_TOKEN)
	channelArr := []string{
		config.CHANNEL_ID,
	}
	fileArr := []string{
		"file.txt",
	}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File: fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Println("		[*] ERROR ", err)	
		}
		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)
	}
}