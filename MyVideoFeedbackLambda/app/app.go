package app

import (
	"log"

	"github.com/WriteRightProject/WriteRightLambda/comprehend"
	"github.com/WriteRightProject/WriteRightLambda/youtube"
	"github.com/aws/aws-lambda-go/events"
)

type App struct {
}

func (app *App) GetFeedbackOfYoutubeVideo(event events.APIGatewayProxyRequest) youtube.Feedback {

	//Get the Video URL from the event
	videoURL := event.Body

	//Lets get the sentiment of the user from comments
	summary, sentiment := app.GetSentimentOfAudience(videoURL)

	//Also lets get the stats of the video
	stats := app.GetStatsForYoutubeVideo(videoURL)

	return youtube.Feedback{
		Summary:      summary,
		Sentiment:    sentiment,
		YoutubeStats: stats,
	}
}

func (a *App) GetSentimentOfAudience(videoURL string) (string, string) {
	comments, err := youtube.FetchComments(videoURL)

	if err != nil {
		log.Fatal("Could not get youtube comments!. Got an error: ", err.Error())
		return "N/A", "N/A"
	}

	summary, sentiment := comprehend.GetVideoFeedback(comments)

	log.Println(summary, sentiment)
	return summary, sentiment
}

func (a *App) GetStatsForYoutubeVideo(videoURL string) youtube.YoutubeStats {
	youtubeStats, err := youtube.ExtractStatsFromVideo(videoURL)

	if err != nil {
		log.Fatal("Could not get youtube stats for the video!. Got an error: ", err.Error())
	}

	log.Println(youtubeStats)
	return youtubeStats
}
