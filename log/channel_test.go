package log

import (
	"github.com/go-kratos/kratos/v2/log"
	"testing"
)

func TestShowOutput(t *testing.T) {
	t.Logf("stderr: %v, stdout: %v, file: %v", Stderr, Stdout, File)
}

func TestNewChannel(t *testing.T) {
	channel := NewChannel("test").PersistentKV("key", "val")

	logger := channel.NewLogger("level", "info")

	h := log.NewHelper(logger)

	h.Info("I am an info log")
	h.Debug("I am an debug log, may not show")

	// Cannot change the configuration, need to rebuild by NewLogger()
	h1 := log.NewHelper(channel.Level(log.LevelDebug).NewLogger())

	h1.Info("I am an info log")
	h1.Debug("I am an debug log, can show")

}

func TestChannel_File(t *testing.T) {
	// maxSize default: 100 MB
	channel := NewChannel("file").Level(log.LevelDebug).Output(File)

	h := log.NewHelper(channel.NewLogger())

	h.Debug("I am an debug log")
	h.Info("I am an info log")

}

func TestChannel_File_Options(b *testing.T) {
	channel := NewChannel("file_compress", &Options{
		Directory:  "../test/log",
		MaxSize:    1, // 1M
		MaxBackups: 10,
		Compress:   true,
	})

	channel.Level(log.LevelDebug).Output(File)

	h := log.NewHelper(channel.NewLogger())

	for i := 0; i < 50000; i++ {
		h.Debugf("index: %d, I am an debug log", i)
	}
}
