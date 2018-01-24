package Legacy

import (
	"github.com/autopilothq/lg"
	"github.com/pkg/errors"

	ct "github.com/autopilothq/banks/contract/types"
	proto "github.com/autopilothq/banks/protocol"
	"github.com/hen-chow/tweetBank/dataAccess/tweets"
	"github.com/hen-chow/tweetBank/resources/Tweets/protocol"
)

var (
	// ErrReadFailed indicates that the tweets couldn't be read
	// from Twitter API, or that the response was malformed in some way.
	ErrReadFailed = "Could not read the tweets, or " +
		"the response was badly formed"

	// ErrNoSearchTerm indicates that the caller omitted the spot id argument
	ErrNoSearchTerm = "The search term is missing, you must provide it"
)

// GetTweet is something that you should document
func GetTweet(request *proto.Request, log lg.Log, aux ct.Auxiliary) *proto.Response {
	getTweetReq := request.Body.(*protocol.GetTweetRequest)

	// The spot argument is required. Return an error if they caller has omitted it.
	if getTweetReq.SearchTerm == "" {
		code := proto.ResponseCode_ERR_INVALID_ARGS
		return request.MakeFatalResponse(code, ErrNoSearchTerm)
	}

	// setting a default search limit
	var limit string
	if getTweetReq.Limit == "" {
		limit = "10"
	} else {
		limit = getTweetReq.Limit
	}

	tweet, err := tweets.GetTweet(getTweetReq.SearchTerm, limit)
	if err != nil {
		code := proto.ResponseCode_ERR_READ_FAILED

		return request.MakeFatalResponse(code, errors.Wrap(err, ErrReadFailed).Error())
	}

	tweets := make([]*protocol.Tweet, 0, len(tweet.Statuses))

	for _, tweet := range tweet.Statuses {
		newTweet := &protocol.Tweet{
			TweetId:            tweet.Id,
			Text:               tweet.Text,
			CreatedAt:          tweet.CreatedAt,
			FavoriteCount:      int32(tweet.FavoriteCount),
			RetweetCount:       int32(tweet.RetweetCount),
			Lang:               tweet.Lang,
			UserName:           tweet.User.Name,
			UserLocation:       tweet.User.Location,
			UserFollowersCount: int32(tweet.User.FollowersCount),
		}

		tweets = append(tweets, newTweet)
	}

	response := request.MakeOkResponse()
	body := response.Body.(*protocol.GetTweetResponse)
	body.Tweets = tweets

	return response

	// tweets := []*Tweet{tweet}

	// return request.MakeOkResponse()

}
