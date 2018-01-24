package tweets

import (
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

func GetTweet(searchTerm string, limit string) (anaconda.SearchResponse, error) {
	consumerKey := os.Getenv("PLURALSIGHT_TWEET_TWITTER_CONSUMER_KEY")
	consumerSecret := os.Getenv("PLURALSIGHT_TWEET_TWITTER_CONSUMER_SECRET")
	accessToken := os.Getenv("PLURALSIGHT_TWEET_TWITTER_ACCESS_TOKEN")
	accessSecret := os.Getenv("PLURALSIGHT_TWEET_TWITTER_ACCESS_SECRET")

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessSecret)

	// setting a limit
	v := url.Values{}
	v.Set("count", limit)

	// Search request
	searchResult, err := api.GetSearch(searchTerm, v)

	// var tweetResp Response
	if err != nil {
		return searchResult, err
	}

	return searchResult, nil
}
