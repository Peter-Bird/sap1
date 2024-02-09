package scr

// ANSI escape code constants for text color
const (
	Black   = "\x1b[30m"
	Red     = "\x1b[31m"
	Green   = "\x1b[32m"
	Yellow  = "\x1b[33m"
	Blue    = "\x1b[34m"
	Magenta = "\x1b[35m"
	Cyan    = "\x1b[36m"
	White   = "\x1b[37m"
)

// ANSI escape code constants for background color
const (
	BgBlack   = "\x1b[40m"
	BgRed     = "\x1b[41m"
	BgGreen   = "\x1b[42m"
	BgYellow  = "\x1b[43m"
	BgBlue    = "\x1b[44m"
	BgMagenta = "\x1b[45m"
	BgCyan    = "\x1b[46m"
	BgWhite   = "\x1b[47m"
)

// ANSI escape code constants for formatting
const (
	Bold       = "\x1b[1m"
	Underline  = "\x1b[4m"
	Italic     = "\x1b[3m"
	CrossedOut = "\x1b[9m"
	Reset      = "\x1b[0m"
	Blink      = "\x1b[5m"
	RapidBlink = "\x1b[6m"
	Reverse    = "\x1b[7m"
	Conceal    = "\x1b[8m"
	Normal     = "\x1b[22m"
	SlowBlink  = "\x1b[25m"
	Overline   = "\x1b[53m"
	Framed     = "\x1b[51m"
	Encircled  = "\x1b[52m"
)

// ANSI escape code constants for cursor commands
const (
	CursorHome     = "\x1b[H"
	CursorUp       = "\x1b[%dA"
	CursorDown     = "\x1b[%dB"
	CursorForward  = "\x1b[%dC"
	CursorBackward = "\x1b[%dD"
	CursorNextLine = "\x1b[%dE"
	CursorPrevLine = "\x1b[%dF"
	CursorColumn   = "\x1b[%dG"
	CursorPosition = "\x1b[%d;%dH"
	CursorSave     = "\x1b[s"
	CursorRestore  = "\x1b[u"
	CursorHide     = "\x1b[?25l"
	CursorShow     = "\x1b[?25h"
)

// ANSI escape code constants for screen and line manipulation
const (
	ClearScreen      = "\x1b[2J"
	ClearScreenHome  = "\x1b[H"
	ScrollUp         = "\x1b[%dS"
	ScrollDown       = "\x1b[%dT"
	ClearLine        = "\x1b[2K"
	ClearLineToEnd   = "\x1b[0K"
	ClearLineToStart = "\x1b[1K"

	Bell = "\x07"
)
