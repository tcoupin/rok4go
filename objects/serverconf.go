package objects

import (
	"github.com/tcoupin/rok4go/utils/log"
)

// Config provide access to all configuration parts
// Config use backend to get and serialize configs
type Config struct {
	backend Backend
	global  *GlobalConfig
}

// SetBackend associates config and backend.
// Force retrieving to cache config
func (c *Config) SetBackend(b Backend) {
	c.backend = b
	c.GlobalConfig(true)
}

// GlobalConfig return global config from cache or backend if nil or force
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

// SetGlobalConfig store GlobalConfig using backend
func (c *Config) SetGlobalConfig(gc *GlobalConfig) error {
	log.DEBUG("Set Global Config")
	err := c.backend.SetGlobalConfig(gc)
	if err == nil {
		c.global = gc
	}
	return err
}
