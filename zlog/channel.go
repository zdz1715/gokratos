package zlog

import (
	"github.com/go-kratos/kratos/v2/log"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path/filepath"
)

const (
	defaultDirectory = "logs"
	defaultExt       = "log"
)

type Options struct {
	Directory string
	// default: log
	Ext string
	// MaxAge is the maximum number of days to retain old log files based on the
	// timestamp encoded in their filename.  Note that a day is defined as 24
	// hours and may not exactly correspond to calendar days due to daylight
	// savings, leap seconds, etc. The default is not to remove old log files
	// based on age.
	MaxAge int
	// MaxSize is the maximum size in megabytes of the log file before it gets rotated.
	// It defaults to 100 megabytes. unit: megabytes
	MaxSize int
	// MaxBackups is the maximum number of old log files to retain. The default is to retain all old log files
	MaxBackups int
	// Compress determines if the rotated log files should be compressed using gzip.
	// The default is not to perform compression.
	Compress bool
}

type Channel struct {
	opts   *Options
	logger log.Logger
	// default: Stderr
	output Output
	// default: LevelInfo
	level log.Level

	// Channel name
	name     string
	filePath string

	kvs []interface{}
}

func (c *Channel) PersistentKV(kvs ...interface{}) *Channel {
	c.kvs = kvs
	return c
}

func (c *Channel) Output(o Output) *Channel {
	c.output = o
	return c
}

func (c *Channel) Option(opts *Options) *Channel {
	c.opts = opts

	directory := defaultDirectory
	ext := defaultExt

	if opts != nil {
		if c.opts.Directory != "" {
			directory = c.opts.Directory
		}
		if c.opts.Ext != "" {
			ext = c.opts.Ext
		}
	}

	c.filePath = filepath.Join(directory, c.name+"."+ext)
	return c
}

func (c *Channel) Level(level log.Level) *Channel {
	c.level = level
	return c
}

func (c *Channel) GetName() string {
	return c.name
}

func (c *Channel) GetOutput() Output {
	return c.output
}

func (c *Channel) GetLevel() log.Level {
	return c.level
}

func (c *Channel) GetLogger() log.Logger {
	return c.logger
}

func (c *Channel) GetPersistentKV() []interface{} {
	return c.kvs
}

func (c *Channel) writer() io.Writer {
	switch c.output {
	case File:
		return &lumberjack.Logger{
			Filename:   c.filePath,
			MaxAge:     c.opts.MaxAge,
			MaxSize:    c.opts.MaxSize, //default: 100MB
			MaxBackups: c.opts.MaxBackups,
			LocalTime:  true,
			Compress:   c.opts.Compress,
		}
	case Stdout:
		return os.Stdout
	default:
		return os.Stderr
	}
}

func (c *Channel) NewLogger(kvs ...interface{}) log.Logger {
	logger := log.NewStdLogger(c.writer())

	if c.name != "" {
		logger = log.With(logger, "channel", c.name)
	}

	logger = log.NewFilter(
		logger,
		log.FilterLevel(c.level),
	)

	if len(c.kvs) > 0 {
		logger = log.With(logger, c.kvs...)
	}

	if len(kvs) > 0 {
		logger = log.With(logger, kvs...)
	}

	c.logger = logger

	return logger
}

func NewChannel(name string, opts ...*Options) *Channel {
	c := &Channel{
		name: name,
	}

	if len(opts) > 0 && opts[0] != nil {
		c.Option(opts[0])
	}

	return c
}
