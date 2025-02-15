package weather

import (
	"context"
	"math/rand/v2"
)

type Event struct {
	Country  string `json:"country"`
	Location string
	Temp     int `json:"temp"`
	Rain     int `json:"rain"`
}

func GenerateEvents(ctx context.Context) chan Event {
	ch := make(chan Event, 1)
	go func() {
		defer close(ch)
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
				ch <- Event{
					Country:  "UK",
					Location: "Zone A",
					Temp:     randomTemp(),
					Rain:     rand.IntN(5),
				}
			}
		}
	}()
	return ch
}

func GenerateTenEvents() []Event {
	weathers := []Event{}
	for v := range 10 {
		var w Event
		if v == 2 || v == 3 {
			w.Country = "DE"
			w.Location = "Zone A"
			w.Temp = randomTemp()
			w.Rain = rand.IntN(5)
		} else if v == 4 || v == 5 {
			w.Country = "UK"
			w.Location = "Zone A"
			w.Temp = randomTemp()
			w.Rain = rand.IntN(5)
		} else if v == 6 || v == 7 {
			w.Country = "FR"
			w.Location = "Zone A"
			w.Temp = randomTemp()
			w.Rain = rand.IntN(5)
		} else if v == 8 || v == 9 {
			w.Country = "SG"
			w.Location = "Zone A"
			w.Temp = randomTemp()
			w.Rain = rand.IntN(5)
		} else {
			w.Country = "US"
			w.Location = "Zone A"
			w.Temp = randomTemp()
			w.Rain = rand.IntN(5)
		}
		weathers = append(weathers, w)
	}
	return weathers
}

func randomTemp() int {
	r := rand.IntN(40)
	if rand.IntN(2) == 0 {
		return r - 30
	}
	return r
}
