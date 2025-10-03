package infrastructure

import "time"

type InMemoryTime struct {
	expectedTime time.Time
}

func NewInMemoryTime(expectedTime time.Time) *InMemoryTime {
	return &InMemoryTime{
		expectedTime: expectedTime,
	}
}

func (i *InMemoryTime) Now() time.Time {
	return i.expectedTime
}
