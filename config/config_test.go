package config

import "testing"

func TestNewConfig(t *testing.T) {
	c := NewConfig("yaml", "/Users/shenping/Project/golang/src/mant/config/yaml/cfg/test.yaml")
	t.Log(c.GetString("customer.family"))
}
