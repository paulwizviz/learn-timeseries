package event

import (
	"context"
	"math/rand/v2"
)

type Weather struct {
	Country  string `json:"country"`
	Location string
	Temp     int `json:"temp"`
	Rain     int `json:"rain"`
}

func Generate(ctx context.Context) chan Weather {
	ch := make(chan Weather, 1)
	go func() {
		defer close(ch)
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
				ch <- Weather{
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

func randomTemp() int {
	r := rand.IntN(40)
	if rand.IntN(2) == 0 {
		return r - 30
	}
	return r
}
