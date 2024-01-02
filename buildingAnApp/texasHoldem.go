package poker

import (
	"io"
	// "os"
	"time"
)

type TexasHoldem struct {
	alerter BlindAlerter
	store   PlayerStore
}

func (p *TexasHoldem) Start(numberOfPlayers int,to io.Writer) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute //for actual application
	// blindIncrement := time.Duration(5+numberOfPlayers) * time.Second //for test

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		p.alerter.ScheduleAlertAt(blindTime, blind,to)
		blindTime = blindTime + blindIncrement
	}
}

func (p *TexasHoldem) Finish(winner string) {
	p.store.RecordWin(winner)
}

func NewTexasHoldem(alerter BlindAlerter, store PlayerStore) *TexasHoldem {
	return &TexasHoldem{
		alerter: alerter,
		store:   store,
	}
}
