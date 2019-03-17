package yaml

import "regexp"

var (
	Regexp_TopNode              = regexp.MustCompile(`^(\S+:.*)$`)
	Regexp_ChildNode            = regexp.MustCompile(`^(\s+\S+:\s+.*)$`)
	Regexp_Node                 = regexp.MustCompile(`^(\S+:)$`)
	Regexp_SecondNode           = regexp.MustCompile(`^(\s+\S+:)$`)
	Regexp_KeyValuePair         = regexp.MustCompile(`^(\S+:\s+[^\*|>|&]{1,})$`)
	Regexp_EndWithFold          = regexp.MustCompile(`^(.*:\s+>)$`)
	Regexp_EndWithVertical      = regexp.MustCompile(`^(.*:\s+\|)$`)
	Regexp_ChildEndWithVertical = regexp.MustCompile(`^(\s+.*:\s+\|)$`)
	Regexp_Anchor               = regexp.MustCompile(`^(.*:\s+&\S+)$`)
	Regexp_Asterisk             = regexp.MustCompile(`^(.*:\s+\*.*)$`)
	Regexp_Array                = regexp.MustCompile(`^(-)$`)
	Regexp_ArrayChild           = regexp.MustCompile(`^(\s*-\s+\S+)$`)
)
