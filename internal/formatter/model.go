package formatter

type Color string

const (
	Red     Color = "\x1b[31m"
	Green   Color = "\x1b[32m"
	Yellow  Color = "\x1b[33m"
	Blue    Color = "\x1b[34m"
	Default Color = "\x1b[39m"

	Strike = "\x1b[9m"
	Reset  = "\x1b[0m"
)
