module github.com/peterszarvas94/goat/cli/goat

go 1.24.1

require (
	github.com/peterszarvas94/goat/pkg v0.2.0
	github.com/spf13/cobra v1.8.1
)

replace github.com/peterszarvas94/goat/pkg => ../../pkg

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)
