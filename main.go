
package main

import(
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)
func main(){

	godotenv.Load(".env")
	token := os.Getenv("SLACK_AUTH_TOKEN")
	channelID := os.Getenv("SLACK_CHANNEL_ID")

	client := slack.New(token, slack.OptionDebug(true))

	attachment := slack.Attachment{
		Pretext: "Super Bot Message",
		Text:    "some text",
		Color: "#36a64f",
		Fields: []slack.AttachmentField{
			{
				Title: "Date",
				Value: time.Now().String(),
			},
		},
	}

	_, timeStamp, err := client.PostMessage(channelID, slack.MsgOptionAttachments(attachment))
	if err != nil{
		panic(err)
	}
	fmt.Printf("Message sent at %s", timeStamp)

	
}