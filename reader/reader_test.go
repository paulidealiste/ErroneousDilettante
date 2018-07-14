package reader

import (
	"testing"
)

func TestFillNames(t *testing.T) {
	testreader := Reader{}
	testerr := testreader.FillNames("test.csv")
	if testerr != nil {
		t.Error("Something went wrong!")
	}
}

func TestFillSurnames(t *testing.T) {
	testreader := Reader{}
	testerr := testreader.FillSurnames("tesft.csv")
	if testerr != nil {
		t.Error("Something went wrong!", testerr)
	}
}

func TestFillReviews(t *testing.T) {
	testreader := Reader{}
	testerr := testreader.FillReviews("testkrook.csv")
	if testerr != nil {
		t.Error("Something went wrong!", testerr)
	}
}
