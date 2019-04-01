package flaws

import "net/http"

type FlawResult int

const (
	CONTINUE FlawResult = iota
	INTERRUPT
)

type Flaw interface {
	Middleware(w http.ResponseWriter, r *http.Request) FlawResult
	Name() string
}

type FlawMiddleware struct {
	Flaw      Flaw
	RunOption RunOption
}

func Of(flaw Flaw, runOption RunOption) FlawMiddleware {
	return FlawMiddleware{
		Flaw:      flaw,
		RunOption: runOption,
	}
}

func (f FlawMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if f.RunOption.CanRun() {
			result := f.Flaw.Middleware(w, r)
			if result == INTERRUPT {
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
