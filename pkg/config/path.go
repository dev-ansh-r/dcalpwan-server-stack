package config

import (
	"os"
	"path"
	"strings"
)

// WithConfigFileFlag is an option for the manager that automatically enables the config file flag
// and tries to infer it from the working directory, user home directory and user configuration directory.
// The Option can only be used once.
func WithConfigFileFlag(flag string) Option {
	return func(m *Manager) {
		if m.configFlag != "" {
			panic("WithConfigFileFlag option should only be used once")
		}

		m.configFlag = flag

		// Use the default from the config if set.
		if def := m.viper.GetStringSlice(flag); def != nil {
			m.viper.SetDefault(flag, def)
			return
		}

		file := m.name + ".yml"
		dotfile := "." + file

		var paths, envPaths []string
		if dir, err := os.Getwd(); err == nil {
			paths = append(paths, path.Join(dir, dotfile))
			envPaths = append(envPaths, dotfile)
		}
		if dir, err := os.UserHomeDir(); err == nil {
			paths = append(paths, path.Join(dir, dotfile))
			envPaths = append(envPaths, path.Join(dir, dotfile))
		}
		if dir, err := os.UserConfigDir(); err == nil {
			paths = append(paths, path.Join(dir, dotfile))
			envPaths = append(envPaths, path.Join(dir, dotfile))
		}

		m.defaultPaths = paths
		m.viper.SetDefault(flag, paths)

		if f := m.flags.Lookup(flag); f != nil {
			f.DefValue = "[" + strings.Join(envPaths, ",") + "]"
		} else {
			m.flags.StringSlice(flag, envPaths, "Location of the configuration file")
		}
	}
}
