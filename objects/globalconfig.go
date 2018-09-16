package objects

import (
	"encoding/json"
	"github.com/tcoupin/rok4go/utils/log"
	"strings"
)

type GlobalConfig struct {
	Title    string
	Keywords []string
}

func DefaultGlobalConfig(c *GlobalConfig) (change bool) {
	log.TRACE("Applying default value for GlobalConfig:\n %v", c)
	change = false
	if len(c.Title) == 0 {
		c.Title = "Rok4Go Server"
		change = true
	}
	if len(c.Keywords) == 0 {
		c.Keywords = []string{"WMTS", "Rok4", "Rok4Go", "Golang"}
		change = true
	}
	if change {
		log.TRACE("Updated GlobalConfig:\n %v", c)
	}
	return
}

func (c *GlobalConfig) String() string {
	var str strings.Builder
	enc := json.NewEncoder(&str)
	enc.Encode(*c)
	return str.String()
}
