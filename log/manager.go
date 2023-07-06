package log

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/zdz1715/gokratos/zmap"
)

type ManagerOption func(m *Manager)

func WithOutput(output Output) ManagerOption {
	return func(m *Manager) {
		m.SetOutput(output)
	}
}

func WithChannelOption(opts *Options) ManagerOption {
	return func(m *Manager) {
		m.SetOption(opts)
	}
}

func WithPersistentKV(kvs ...interface{}) ManagerOption {
	return func(m *Manager) {
		m.SetPersistentKV(kvs...)
	}
}

func WithLevel(level log.Level) ManagerOption {
	return func(m *Manager) {
		m.SetLevel(level)
	}
}

type Manager struct {
	level log.Level

	kvs []interface{}

	output             Output
	defaultChannelName string

	channels zmap.StringMap[*Channel]
	opts     *Options
}

func (m *Manager) SetLevel(level log.Level) *Manager {
	if level != m.level {
		m.level = level
	}
	return m
}

func (m *Manager) SetPersistentKV(kvs ...interface{}) *Manager {
	m.kvs = kvs
	return m
}

func (m *Manager) SetOption(opts *Options) *Manager {
	m.opts = opts
	return m
}

func (m *Manager) SetOutput(output Output) *Manager {
	if output != m.output {
		m.output = output
	}
	return m
}

func (m *Manager) Get(names ...string) log.Logger {
	name := m.defaultChannelName
	if len(names) > 0 && names[0] != "" {
		name = names[0]
	}
	if c, ok := m.channels.Get(name); ok {
		return c.GetLogger()
	}
	return log.GetLogger()
}

func (m *Manager) Reassign() {
	m.AddChannel(m.channels.Keys()...)
}

func (m *Manager) AddChannel(names ...string) {
	for _, name := range names {
		m.addSingleChannel(name)
	}
}

func (m *Manager) addSingleChannel(name string) {

	channel, ok := m.channels.Get(name)
	if !ok {
		channel = NewChannel(name, m.opts)
	}

	logger := channel.Level(m.level).Output(m.output).Option(m.opts).PersistentKV(m.kvs...).NewLogger()

	if name == m.defaultChannelName {
		// global log
		log.SetLogger(logger)
	}

	m.channels.Set(name, channel)
}

func NewManager(defaultChannelName string, opts ...ManagerOption) *Manager {
	m := &Manager{
		kvs:                make([]interface{}, 0),
		defaultChannelName: defaultChannelName,
	}

	for _, o := range opts {
		o(m)
	}

	m.addSingleChannel(defaultChannelName)

	return m
}
