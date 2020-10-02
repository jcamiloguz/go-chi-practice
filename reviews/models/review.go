package models

import (
	"errors"
	"time"
)

const maxLengthInComments = 400

//Review represent an anon review from some website
type Review struct {
	ID      int64
	Stars   int       //1-5
	Comment string    // max 400chars
	date    time.Time // create at
}

//CreateReviewCMD command to create a new review
type CreateReviewCMD struct {
	Stars   int    `json:"stars"`
	Comment string `json:"comment"`
}

func (cmd *CreateReviewCMD) validate() error {
	if cmd.Stars < 1 || cmd.Stars > 5 {
		return errors.New("starts must be 1-5")
	}

	if len(cmd.Comment) > maxLengthInComments {
		return errors.New("comment must be less than 400 chars")
	}
	return nil
}
