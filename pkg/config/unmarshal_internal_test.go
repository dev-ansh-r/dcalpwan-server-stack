package config

import "io"

func (m *Manager) MergeConfig(in io.Reader) error {
	return m.mergeConfig(in)
}
