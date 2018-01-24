// Code generated by 'banks apply'; DO NOT EDIT.

package resources

import (
	"github.com/autopilothq/banks/contract"
	"github.com/autopilothq/banks/protocol/marshalable"
	"github.com/autopilothq/banks/types"

	"github.com/hen-chow/tweetBank/protocols"
	rodata "github.com/hen-chow/tweetBank/resourcefile"

	// register resources and load their protocols
	_ "github.com/hen-chow/tweetBank/resources/Tweets"
	protocolTweets "github.com/hen-chow/tweetBank/resources/Tweets/protocol"
)

func init() {
	lock, err := rodata.GetROLockfile()
	if err != nil {
		panic(err)
	}

	contract.Provides(
		lock.Resources,
		lock.Resource,
		protocols.Read,
		MsgIDFromReqBody,
	)
}

// MsgIDFromReqBody looks up a message ID for a request body type. It is used
// by banks dispatch.
func MsgIDFromReqBody(reqBody marshalable.Body) (types.MessageID, bool) {
	switch reqBody.(type) {

	case *protocolTweets.GetTweetRequest:
		return types.MessageID(65537), true

	default:
		return types.MessageID(0), false
	}
}