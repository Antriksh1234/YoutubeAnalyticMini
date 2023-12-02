package youtube

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type Feedback struct {
	Sentiment string `json:"sentiment"`
	Summary   string `json:"summary"`
	YoutubeStats
}

type YoutubeStats struct {
	ChannelName  string `json:"channel"`
	Title        string `json:"title"`
	ThumbnailURL string `json:"thumbnailURL"`
	Views        uint64 `json:"views"`
	Likes        uint64 `json:"likes"`
	Dislikes     uint64 `jsno:"dislikes"`
}

var APIKey = os.Getenv("API_KEY")

// FetchComments Fetches the Youtube comments under a Youtube Video via its Video URL.
func FetchComments(videoURL string) ([]string, error) {
	// Extract video ID from the YouTube URL
	videoID, err := extractVideoID(videoURL)
	if err != nil {
		return nil, err
	}

	service, err := youtube.NewService(context.TODO(), option.WithAPIKey(APIKey))
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
		return nil, err
	}

	// Fetch comments using the video ID
	comments, err := FetchCommentsByVideoID(service, videoID)
	if err != nil {
		log.Printf("Error fetching comments: %v", err)
		return nil, err
	}

	return comments, nil
}

func ExtractStatsFromVideo(videoURL string) (YoutubeStats, error) {
	stats := YoutubeStats{}

	// Extract video ID from the YouTube URL
	videoID, err := extractVideoID(videoURL)
	if err != nil {
		log.Fatalf("Cannot extract video id from URL: %v", err)
		return stats, err
	}

	service, err := youtube.NewService(context.TODO(), option.WithAPIKey(APIKey))
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
		return stats, err
	}

	call := service.Videos.List([]string{"statistics", "snippet"}).Id(videoID)
	response, err := call.Do()
	if err != nil {
		log.Fatalf("Error making API call to retrieve video details: %v", err)
		return stats, err
	}

	if len(response.Items) == 0 {
		fmt.Println("Video not found")
		return stats, nil
	}

	statistics := response.Items[0].Statistics
	snippet := response.Items[0].Snippet

	stats.Views = statistics.ViewCount
	stats.Likes = statistics.LikeCount
	stats.Dislikes = statistics.DislikeCount
	stats.Title = snippet.Title
	stats.ThumbnailURL = snippet.Thumbnails.Default.Url
	stats.ChannelName = snippet.ChannelTitle

	return stats, nil
}

func extractVideoID(videoURL string) (string, error) {
	u, err := url.Parse(videoURL)
	if err != nil {
		return "", err
	}

	if u.Host == "youtu.be" {
		// Extract video ID from short-form URL
		videoID := strings.TrimPrefix(u.Path, "/")
		if videoID == "" {
			return "", fmt.Errorf("unable to extract video ID from URL")
		}
		return videoID, nil
	}

	queryParams, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return "", err
	}

	videoID := queryParams.Get("v")
	if videoID == "" {
		return "", fmt.Errorf("unable to extract video ID from URL")
	}

	return videoID, nil
}

func FetchCommentsByVideoID(service *youtube.Service, videoID string) ([]string, error) {
	// Retrieve the comments for the video
	call := service.CommentThreads.List([]string{"snippet"}).VideoId(videoID)
	response, err := call.Do()
	if err != nil {
		log.Printf("Error fetching comments for video %s: %v\n", videoID, err)
		return nil, err
	}

	// Extract comment text from the response
	var comments []string
	for _, item := range response.Items {
		comment := item.Snippet.TopLevelComment.Snippet.TextDisplay
		comments = append(comments, comment)
	}

	return comments, nil
}
