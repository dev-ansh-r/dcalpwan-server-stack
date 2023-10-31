package component

import "go.thethings.network/lorawan-stack/v3/pkg/task"

// RegisterTask registers a task, optionally with restart policy and backoff, to be started after the component started.
func (c *Component) RegisterTask(conf *task.Config) {
	c.taskConfigs = append(c.taskConfigs, conf)
}

// StartTask implements the TaskStarter interface.
func (c *Component) StartTask(conf *task.Config) {
	c.taskStarter.StartTask(conf)
}

func (c *Component) startTasks() {
	for _, conf := range c.taskConfigs {
		c.taskStarter.StartTask(conf)
	}
}
