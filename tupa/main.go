package main

import "github.com/tupatech/tupa"

func main() {
	routeManager := func() {
		tupa.AddRoutes(nil, func() []tupa.RouteInfo {
			return []tupa.RouteInfo{
				{
					Path:   "/",
					Method: "GET",
					Handler: func(c *tupa.TupaContext) error {
						return tupa.WriteJSONHelper(c.Resp, 200, "Hello world! :D")
					},
				},
			}
		})
	}

	server := tupa.NewAPIServer(":6969", routeManager)
	server.New()
}
