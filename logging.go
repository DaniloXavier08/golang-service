package main

import (
	"context"
	"fmt"
	"time"
)

type LoggigingService struct {
	next Service
}

func NewLoggingService(next Service) Service {
	return &LoggigingService{
		next: next,
	}
}

func (s *LoggigingService) GetCatFact(ctx context.Context) (fact *CatFact, err error) {
	defer func(start time.Time) {
		fmt.Printf("fact=%s err=%v tooke=%v\n", fact.Fact, err, time.Since(start))
	}(time.Now())

	return s.next.GetCatFact(ctx)
}
