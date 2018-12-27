package yaml

import "regexp"

var (
	Regexp_Node                 = regexp.MustCompile(`(^\w+.*:.*)$`)
	Regexp_ChildNode            = regexp.MustCompile(`^(\s+.*:.*)$`)
	Regexp_KeyValuePair         = regexp.MustCompile(`^(.*:\s+[^(\*|>|&)]{1,})$`)
	Regexp_EndWithFold          = regexp.MustCompile(`^(.*:\s+>)$`)
	Regexp_EndwithVertical      = regexp.MustCompile(`^(.*:\s+\|)$`)
	Regexp_ChildEndwithVertical = regexp.MustCompile(`^(\s+.*:\s+\|)$`)
	Regexp_Anchor               = regexp.MustCompile(`^(.*:\s+&\S+)$`)
	Regexp_Array                = regexp.MustCompile(`^(\s*\-.*)$`)
)
