package main

import (
	"my_first_test/gomock/mock"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_mock_LearnInf(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	name := "haha"
	mockLearnInf := mock.NewMockLeranInf(ctrl)

	mockLearnInf.EXPECT().LearnGo(name).DoAndReturn(func(string) (string, error) {
		return "Learn Go is ok", nil
	}).Times(2)

	// mockLearnInf.EXPECT().LearnGo(name).Return("Learn Go is ok-case2", nil)

	output, err := mockLearnInf.LearnGo(name)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("output: %s\n", output)

	output2, err := mockLearnInf.LearnGo(name)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("output: %s\n", output2)

}
