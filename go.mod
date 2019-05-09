module golib

go 1.12

require (
	github.com/spf13/viper v1.3.2
	goLib/action v0.0.0-00010101000000-000000000000
)

replace goLib/action => ./action
