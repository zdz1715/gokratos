package zlog

import (
	"github.com/go-kratos/kratos/v2/log"
	"testing"
)

func TestManager(t *testing.T) {
	opts := []ManagerOption{
		WithOutput(Stdout),
		WithLevel(log.LevelInfo),
		WithPersistentKV("by", "manager"),
		WithChannelOption(nil),
	}
	m := NewManager("default", opts...)

	m.AddChannel("log_1", "log_2")

	h := log.NewHelper(m.Get())
	h1 := log.NewHelper(m.Get("log_1"))
	h2 := log.NewHelper(m.Get("log_2"))

	h.Info("i am default")
	h1.Info("i am log_1")
	h2.Info("i am log_2")

	m.SetLevel(log.LevelDebug).SetPersistentKV().Reassign()

	h = log.NewHelper(m.Get())
	h1 = log.NewHelper(m.Get("log_1"))
	h2 = log.NewHelper(m.Get("log_2"))

	h.Debug("i am default")
	h1.Debug("i am log_1")
	h2.Debug("i am log_2")
}

func TestManager_FILE(t *testing.T) {
	opts := []ManagerOption{
		WithOutput(File),
		WithLevel(log.LevelDebug),
		WithPersistentKV("by", "manager"),
		WithChannelOption(&Options{
			Directory: "../test/log1",
		}),
	}
	m := NewManager("default", opts...)

	m.AddChannel("log_1", "log_2")

	h := log.NewHelper(m.Get())
	h1 := log.NewHelper(m.Get("log_1"))
	h2 := log.NewHelper(m.Get("log_2"))

	h.Info("i am default")
	h1.Info("i am log_1")
	h2.Info("i am log_2")

	m.SetLevel(log.LevelDebug).SetPersistentKV().SetOption(&Options{
		Directory: "../test/log2",
	}).Reassign()

	h = log.NewHelper(m.Get())
	h1 = log.NewHelper(m.Get("log_1"))
	h2 = log.NewHelper(m.Get("log_2"))

	h.Debug("i am default")
	h1.Debug("i am log_1")
	h2.Debug("i am log_2")
}
