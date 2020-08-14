package main

// OAuth2
import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

func main() {
	// oauth2 configures a client that uses app credentials to keep a fresh token
	config := &clientcredentials.Config{
		ClientID:     "LJNv0LzKWrAqOqvz7S3XSz2Cd",
		ClientSecret: "SDXPOQVovqXtfGrMUlXfQnXOS06P5seUMgqbVu4J2AUZYWnQOh",
		TokenURL:     "https://api.twitter.com/oauth2/AAAAAAAAAAAAAAAAAAAAAH4lGwEAAAAAKKikVOZrFjohDRDYKjjfHossafU%3D7MoYnW7EB8IEr8mHVhd450hRiT8k6U8A7UycKkAwTHzYDd7PEi",
	}
	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth2.NoContext)

	// Twitter client
	client := twitter.NewClient(httpClient)

	tweets, resp, err := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
		Count: 20,
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(tweets)
	fmt.Println(resp)
}
