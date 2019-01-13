package yaml

func (y *Yaml) Match(field string, fn handle) {
	ok := true
	switch ok {
	case Regexp_Node.MatchString(field),
		Regexp_ChildNode.MatchString(field),
		Regexp_KeyValuePair.MatchString(field),
		Regexp_EndWithFold.MatchString(field),
		Regexp_EndWithVertical.MatchString(field),
		Regexp_ChildEndwithVertical.MatchString(field),
		Regexp_Anchor.MatchString(field),
		Regexp_Asterisk.MatchString(field),
		Regexp_Array.MatchString(field),
		Regexp_ArrayChild.MatchString(field),
		Regexp_ArrayNode.MatchString(field):

		fn()
	}
}
