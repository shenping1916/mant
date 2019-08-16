package config

import "testing"

func TestNewConfig(t *testing.T) {
	c := NewConfig("yaml", "/Users/shenping/Project/golang/src/mant/config/yaml/cfg/test.yaml")
	t.Log(c.GetBool("test"))
	t.Log(c.GetBool("log.rotate"))
	t.Log(c.GetStringArray("customer.table"))
}

func BenchmarkNewConfig(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	c := NewConfig("yaml", "/Users/shenping/Project/golang/src/mant/config/yaml/cfg/test.yaml")
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		_, _ = c.GetString("customer.family")
		b.StopTimer()
	}
}
