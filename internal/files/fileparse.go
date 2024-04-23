package files

import (
	"fmt"
	"strings"

	"github.com/unhanded/mytoman/internal/cfg"
)

type MytoFile struct {
	configSectionStr string
	TemplateSection  string
	Config           cfg.MytoConfig
}

func (tbf *MytoFile) FromBytes(data []byte) error {
	sections := strings.Split(string(data), "---\n")
	if len(sections) < 2 {
		return fmt.Errorf("could not split into two sections")
	}
	tbf.configSectionStr = sections[0]
	tbf.TemplateSection = sections[1]
	return nil
}

func (tbf *MytoFile) Parse() error {
	return tbf.parseConfig()
}

func (tbf *MytoFile) parseConfig() error {
	cfg, cfgErr := cfg.ParseConfig(tbf.configSectionStr)
	if cfgErr != nil {
		return cfgErr
	}
	tbf.Config = cfg
	return nil
}

func IsSplitter(ln string) bool {
	return strings.HasPrefix(ln, "---")
}
