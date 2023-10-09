
package main

import(
	"fmt"
	"os"
	"time"
	"aalsaedi265/slack_bot/database"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	// "github.com/slack-go/slack/socketmode"
)
func main(){

	database.Connector();

	godotenv.Load(".env")
	token := os.Getenv("SLACK_AUTH_TOKEN")
	channelID := os.Getenv("SLACK_CHANNEL_ID")
	appToken := os.Getenv("SLACK_APP_TOKEN")

	client := slack.New(token, slack.OptionDebug(true), slack.OptionAppLevelToken(appToken))

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