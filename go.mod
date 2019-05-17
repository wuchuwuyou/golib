module golib

go 1.12

replace golib/action => ./action

replace golib/file => ./file

replace golib/api => ./api

require (
	golang.org/x/mobile v0.0.0-20190509164839-32b2708ab171 // indirect
	golib/action v0.0.0-00010101000000-000000000000 // indirect
	golib/api v0.0.0-00010101000000-000000000000
	golib/file v0.0.0-00010101000000-000000000000
)
