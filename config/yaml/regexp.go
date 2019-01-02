package yaml

import "regexp"

var (
	Regexp_Node                 = regexp.MustCompile(`^(\S+:.*)$`)
	Regexp_ChildNode            = regexp.MustCompile(`^(\s+.*:.*)$`)
	Regexp_KeyValuePair         = regexp.MustCompile(`^(.*:\s+[^(\*|>|&)]{1,})$`)
	Regexp_EndWithFold          = regexp.MustCompile(`^(.*:\s+>)$`)
	Regexp_EndWithVertical      = regexp.MustCompile(`^(.*:\s+\|)$`)
	Regexp_ChildEndwithVertical = regexp.MustCompile(`^(\s+.*:\s+\|)$`)
	Regexp_Anchor               = regexp.MustCompile(`^(.*:\s+&\S+)$`)
	Regexp_Asterisk             = regexp.MustCompile(`^(.*:\s+\*.*)$`)
	Regexp_Array                = regexp.MustCompile(`^(\s*\-.*)$`)
	Regexp_ArrayChild           = regexp.MustCompile(`^(\s*-\s+.*:.*)`)
)
