package proxy

import "github.com/GaruGaru/flaw/flaws"

type Config struct {
	Bind   string
	Host   string
	Scheme string

	EnabledFlaws []flaws.FlawMiddleware
}
