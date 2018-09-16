package objects

import (
	"github.com/tcoupin/rok4go/utils/log"
)

type Config struct {
	backend Backend
	global  *GlobalConfig
}

func (c *Config) SetBackend(b Backend) {
	c.backend = b
	b.GetGlobalConfig(c.global)
	c.GlobalConfig(true)
}

func (c *Config) GlobalConfig(force bool) *GlobalConfig {
	if c.global == nil || force {
		log.DEBUG("Global config is nil, create a new one")
		c.global = &GlobalConfig{}
		c.backend.GetGlobalConfig(c.global)
	}
	if change := DefaultGlobalConfig(c.global); change {
		err := c.backend.SetGlobalConfig(c.global)
		if err != nil {
			log.ERROR("Fail to update global config: %v", err)
		}
	}
	return c.global
}

func (c *Config) SetGlobalConfig(gc *GlobalConfig) error {
	log.DEBUG("Set Global Config")
	err := c.backend.SetGlobalConfig(gc)
	if err == nil {
		c.global = gc
	}
	return err
}
