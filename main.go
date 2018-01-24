package main

import (
	"github.com/hen-chow/tweetBank/cmd"

	// register protocols and resources
	_ "github.com/hen-chow/tweetBank/protocols"
	_ "github.com/hen-chow/tweetBank/resources"
)

func main() {
	cmd.RootCmd.Execute()
}
