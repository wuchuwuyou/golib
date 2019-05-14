module golib

go 1.11

require (
	github.com/spf13/viper v1.3.2
	goLib/action v0.0.0-00010101000000-000000000000
	goLib/file v0.0.0
	golang.org/x/mobile v0.0.0-20190509164839-32b2708ab171 // indirect
	github.com/valyala/fasthttp v1.2.0
)

replace goLib/action => ./action

replace goLib/file => ./file
