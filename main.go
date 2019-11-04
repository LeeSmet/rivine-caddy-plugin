package main

import (
	"github.com/caddyserver/caddy"
)

func init() {
	caddy.RegisterPlugin("rivine_explorer_balancer", caddy.Plugin{
		ServerType: "http",
		Action:     setup,
	})
}

func setup(c *caddy.Controller) error {

	return nil
}
