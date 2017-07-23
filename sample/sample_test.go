package sample

import (
	"testing"

	"github.com/golang/mock/gomock"

	mock "github.com/kitakitabauer/gomock-sample/mock_sample"
)

func TestSample(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSample := mock.NewMockSample(ctrl)
	mockSample.EXPECT().Method("hoge").Return(1)

	t.Log("result: ", mockSample.Method("hoge"))
}

func TestSample2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSample := mock.NewMockSample(ctrl)

	gomock.InOrder(
		mockSample.EXPECT().Method("hoge").Return(1),
		mockSample.EXPECT().Method("fuga").Return(2),
	)
	/*
		mockSample.EXPECT().Mehotd("hoge").Return(1).After(
			mockSample.EXPECT().Method("fuga").Return(2),
			)
		)
	*/

	t.Log("result", mockSample.Method("hoge"))
	t.Log("result", mockSample.Method("fuga"))
}

func TestSample3(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSample := mock.NewMockSample(ctrl)

	mockSample.EXPECT().Method("hoge").Return(1).Times(2)

	// mockSample.EXPECT().Method("hoge").Return(1).AnyTimes()

	t.Log("result", mockSample.Method("hoge"))
	t.Log("result", mockSample.Method("hoge"))
}
