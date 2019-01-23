package log

import "testing"

func TestConsole(t *testing.T) {
	st := struct {
		A string
	}{A: "it is a struct"}

	l := GetLogger()
	l.SetConfig(LoggerConfig{
		LogLevel:  DEBUG,
		IsConsole: true,
	})
	l.Debug(`company`, "didi", "province", "beijing")
	l.AddTextPrefix("method", "resUser")
	l.Info("name", "xiaoli")
	l.Info("data", map[string]string{"info": "abc", "info1": "def"})
	l.Info("struct", st)
	// debug的不会输出
}
