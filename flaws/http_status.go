package flaws

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

type HttpStatusCodeFlaw struct {
	Status int
}

func NewHttpStatusCode(status int) HttpStatusCodeFlaw {
	return HttpStatusCodeFlaw{Status: status}
}

func (l HttpStatusCodeFlaw) Middleware(w http.ResponseWriter, r *http.Request) FlawResult {
	logrus.Warnf("[INJECTED] HttpStatusCode %d in request %s", l.Status, r.RequestURI)
	w.WriteHeader(l.Status)
	return INTERRUPT
}

func (l HttpStatusCodeFlaw) Name() string {
	return "HttpStatusCode"
}
