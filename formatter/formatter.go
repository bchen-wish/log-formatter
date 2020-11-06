package formatter

import (
	"github.com/VertexC/log-formatter/formatter/general"
)

type Formatter interface {
	Init()
	Format(msg string) map[string]interface{}
}

type Config struct {
	Type       string         `yaml:"type"`
	GeneralCfg general.Config `yaml:"general"`
}

func New(config Config) Formatter {
	switch config.Type {
	case "general":
		formatter := new(general.Formatter)
		formatter.Init()
		formatter.SetConfig(config.GeneralCfg)
		return formatter
	}
	return nil
}

func Execute(config Config, inputCh chan interface{}, outputCh chan interface{}) {
	formatter := New(config)
	for {
		record := <-inputCh
		message := record.(map[string]interface{})["message"].(string)
		// FIXME: here kv should be map[string]string

		kvMap := formatter.Format(message)
		if kvMap == nil {
			kvMap = map[string]interface{}{"sourceData_": record}
			continue
		} else {
			kvMap["sourceData_"] = record
		}
		outputCh <- kvMap
	}
}
