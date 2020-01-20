package config

import "github.com/martini-contrib/cors"

var CorsOptions = &cors.Options{AllowOrigins: []string{"*"}}
