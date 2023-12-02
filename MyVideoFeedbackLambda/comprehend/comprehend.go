package comprehend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/comprehend"
)

var meaningcloudAPIKey = os.Getenv("MEANING_CLOUD_API_KEY")
var comprehendClient comprehend.Comprehend

func initComprehendClient() {
	sess, err := session.NewSession()

	if err != nil {
		log.Println("Could not establish session")
	}

	comprehendClient = *comprehend.New(sess)
}

// Provides us with the feedback of the audience via their comments
func GetVideoFeedback(comments []string) (string, string) {

	//Lets initialize the client first
	initComprehendClient()

	//Get the summary of cocomprehend
	summary := getAudienceFeedbackSummary(comments)

	//Get the sentiment of the audience via their comments
	sentiment, _ := getAudienceSentiment(comments)

	return summary, sentiment
}

func getAudienceSentiment(comments []string) (string, error) {
	text := strings.Join(comments, ".")

	//Fetching sentiment via AWS Comprehend
	result, err := comprehendClient.DetectSentiment(&comprehend.DetectSentimentInput{
		Text:         &text,
		LanguageCode: aws.String("en"), // Set the language code accordingly
	})

	if err != nil {
		log.Printf("Error while fetching sentiment: %s", err.Error())
		return "", err
	}

	// Return the sentiment
	return *result.Sentiment, nil
}

func getAudienceFeedbackSummary(comments []string) string {

	// Define the text to be summarized
	text := strings.Join(comments, ".")

	// Create the API request URL

	// Create the API request URL
	apiURL := "https://api.meaningcloud.com/summarization-1.0"

	params := url.Values{}
	params.Set("key", meaningcloudAPIKey)
	params.Set("txt", text)
	params.Set("sentences", "3")

	urlWithParams := fmt.Sprintf("%s?%s", apiURL, params.Encode())

	// Create the API request body
	requestBody := bytes.NewBuffer([]byte{})

	// Create the API request
	req, err := http.NewRequest("POST", urlWithParams, requestBody)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	// Send the API request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		}

		log.Fatalf("API error: %s\n", body)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Print the response body as a string
	fmt.Println(string(responseBody))

	// Parse the JSON response
	decoder := json.NewDecoder(resp.Body)
	var summary map[string]interface{}
	err = decoder.Decode(&summary)
	log.Println(summary)
	if err != nil {
		fmt.Println("Error occured!", err.Error())
		return ""
	}

	// Extract the summary text
	summaryText := ""

	return summaryText
}
