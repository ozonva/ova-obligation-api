package saver

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/ozonva/ova-obligation-api/internal/entity"
	"github.com/ozonva/ova-obligation-api/internal/flusher"
	"github.com/stretchr/testify/assert"
)

func TestSaver_SaveByTimer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	flusherMock := flusher.NewMockFlusher(ctrl)

	entities := []entity.Obligation{{ID: 1}}
	flusherMock.EXPECT().Flush(entities).Return(nil).MinTimes(2)

	saver := NewSaver(5, flusherMock, *time.NewTicker(100 * time.Microsecond))
	saver.Save(entities[0])

	time.Sleep(10000 * time.Microsecond)

	saver.Save(entities[0])

	time.Sleep(10000 * time.Microsecond)

}

func TestSaver_SaveToClosedSaver(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	flusherMock := flusher.NewMockFlusher(ctrl)

	entities := []entity.Obligation{{ID: 1}}
	flusherMock.EXPECT().Flush(entities).Return(nil).MaxTimes(0)

	saver := NewSaver(5, flusherMock, *time.NewTicker(1 * time.Second))
	saver.Close()
	err := saver.Save(entities[0])

	assert.NotNil(t, err)
	assert.Error(t, err, "saver already closed")
	assert.ErrorIs(t, err, ErrorClosedSaver)
}

func TestSaver_CloseClosedSaver(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	flusherMock := flusher.NewMockFlusher(ctrl)

	entities := []entity.Obligation{{ID: 1}}
	flusherMock.EXPECT().Flush(entities).Return(nil).MaxTimes(0)

	saver := NewSaver(5, flusherMock, *time.NewTicker(1 * time.Second))
	saver.Close()
	err := saver.Close()

	assert.NotNil(t, err)
	assert.Error(t, err, "saver already closed")
	assert.ErrorIs(t, err, ErrorClosedSaver)
}
