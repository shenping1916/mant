package yaml

var (
	//_POUND_SIGN        = "#"
	//_COLON             = ":"
	//_AMPERSAND         = "&"
	//_ASTERISK          = "*"
	//_HYPHEN            = "-"
	//_TILDE             = "~"
	//_VERTICAL_BAR      = []byte{124}        // "|"
	//_GREATER_THAN_SIGN = []byte{62}         // ">"
	_START = []byte{45, 45, 45} // "---"
	_END   = []byte{46, 46, 46} // "..."
)

var (
	strType   = "!!str"
	floatType = "!!float"
	setType   = "!!set"
)
