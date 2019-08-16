package yaml

import (
	"strings"
)

const (
	RegexpTopNode              = 0x0
	RegexpNode                 = 0x1
	RegexpChildNode            = 0x2
	RegexpKeyValuePair         = 0x3
	RegexpEndWithFold          = 0x4
	RegexpEndWithVertical      = 0x5
	RegexpChildEndWithVertical = 0x6
	RegexpAnchor               = 0x7
	RegexpAsterisk             = 0x8
	RegexpArray                = 0x9
	RegexpArrayChild           = 0x10
	RegexpSecondNode           = 0x11
)

func (y *Yaml) Match(s *segment, array *List, m map[string]interface{}) {
	switch y.MatchHandle(s.key) {
	case RegexpKeyValuePair:
		// like(a:  b)
		y.KeyValuePair(s.key, m)
	case RegexpArray:
		/*
			        -
						part_no:   A4786
						descrip:   Water Bucket (Filled)
						price:     1.47
						quantity:  4

					-
						part_no:   E1628
						descrip:   High Heeled "Ruby" Slippers
						size:      8
						price:     133.7
						quantity:  1
		*/
		y.KeyArray(array, m)
	case RegexpArrayChild:
		/*
			like(server:
					 - 120.168.117.21
					 - 120.168.117.22
					 - 120.168.117.23)
		*/
		y.KeyArrayChild(s.key, array)
	case RegexpEndWithFold:
		// like(a: >)
		y.KeyFoldPair(s, y.Data)
		y.KeyFoldPair(s, m)
	case RegexpEndWithVertical:
		// like(a: |)
		y.KeyVerticalPair(s, y.Data)
		y.KeyVerticalPair(s, m)
	case RegexpAnchor:
		// like(a: &id001)
		y.KeyAnchor(s)
	case RegexpAsterisk:
		// like(a: *id001)
		y.KeyAsterisk(s.key)
	case RegexpNode:
		key := s.key
		if len(s.value) > 0 {
			m := make(map[string]interface{})
			array := make(List, 0, len(s.value))

			if s.value[0].(string) == "- " {
				y.ElementReverse(s.value)
			}

			for _, value := range s.value {
				switch value := value.(type) {
				case string:
					if y.MatchHandle(value) == RegexpSecondNode {
						if len(m) > 0 {
							_key := s.key

							y.Lock()
							for k, v := range m {
								s.key = _key + "." + k
								y.Data[s.key] = v
							}
							m = make(map[string]interface{})
							y.Unlock()
						}

						value = strings.TrimRight(strings.TrimSpace(value), ":")
						s.key = strings.TrimRight(key, ":") + "." + value
					} else {
						s.key = strings.TrimRight(s.key, ":")
						value = strings.TrimSpace(value)
					}

					// TODO: add comment
					y.Match(&segment{
						key: value,
					}, &array, m)
				}
			}

			y.Lock()
			if len(array) > 0 {
				y.Data[s.key] = array
				if len(m) > 0 {
					m = make(map[string]interface{})
				}
			}

			if len(m) > 0 {
				_key := s.key

				for k, v := range m {
					s.key = _key + "." + k
					y.Data[s.key] = v
				}
			}
			y.Unlock()
		}
	}
}

func (y *Yaml) ElementReverse(array List) {
	length := len(array)
	for i, j := 0, length-1; i < length/2; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
}

func (y *Yaml) MatchHandle(field string) int {
	switch {
	case Regexp_Node.MatchString(field):
		return RegexpNode
	case Regexp_KeyValuePair.MatchString(field):
		return RegexpKeyValuePair
	case Regexp_ChildNode.MatchString(field):
		return RegexpChildNode
	case Regexp_EndWithFold.MatchString(field):
		return RegexpEndWithFold
	case Regexp_EndWithVertical.MatchString(field):
		return RegexpEndWithVertical
	case Regexp_ChildEndWithVertical.MatchString(field):
		return RegexpChildEndWithVertical
	case Regexp_Anchor.MatchString(field):
		return RegexpAnchor
	case Regexp_Asterisk.MatchString(field):
		return RegexpAsterisk
	case Regexp_Array.MatchString(field):
		return RegexpArray
	case Regexp_ArrayChild.MatchString(field):
		return RegexpArrayChild
	case Regexp_SecondNode.MatchString(field):
		return RegexpSecondNode
	default:
		return RegexpTopNode
	}
}
