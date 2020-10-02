package models

import "testing"

func NewReview(stars int, comment string) *CreateReviewCMD {
	return &CreateReviewCMD{
		Stars:   stars,
		Comment: comment,
	}
}

func Test_withCorrectParams(t *testing.T) {
	r := NewReview(4, "The iphone X look good")
	err := r.validate()
	if err != nil {
		t.Error("The validation did not pass")
		t.Fail()
	}
}

func Test_shouldFailWithWrongNumberOfStars(t *testing.T) {
	r := NewReview(8, "The iphone X look really good")

	err := r.validate()

	if err != nil {
		t.Error("Should fail with more of 5 stars")
	}

}
