package flaws

import (
	"github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
	"time"
)

type LatencyFlaw struct {
	Min int
	Max int
}

func NewRandomLatency(min int, max int) LatencyFlaw {
	return LatencyFlaw{
		Min: min,
		Max: max,
	}
}

func NewFixedLatency(value int) LatencyFlaw {
	return LatencyFlaw{
		Min: value,
		Max: value,
	}
}

func (l LatencyFlaw) Middleware(w http.ResponseWriter, r *http.Request) FlawResult {
	var sleepDuration = l.Max

	if l.Max != l.Min {
		sleepDuration = rand.Intn(l.Max-l.Min) + l.Min
	}

	logrus.Warnf("[INJECTED] latency of %dms in request %s", sleepDuration, r.URL.String())

	time.Sleep(time.Duration(sleepDuration) * time.Millisecond)
	return CONTINUE
}

func (l LatencyFlaw) Name() string {
	return "Latency"
}
