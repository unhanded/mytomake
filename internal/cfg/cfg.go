package cfg

import (
	"fmt"
	"strings"

	"github.com/unhanded/mytomake/internal/util"
)

type MytoConfig struct {
	Items   []ConfigItem
	DataDir string
}

func ParseConfig(content string) (MytoConfig, error) {
	items := make([]ConfigItem, 0)
	lines := strings.Split(content, "\n")
	for lineN, line := range lines {
		if util.IsLineComment(line) || util.IsEmptyLine(line) {
			continue
		}
		item, err := parseCfgLine(line, lineN)
		if err != nil {
			return MytoConfig{}, fmt.Errorf("error ocurred: %s", err)
		}
		items = append(items, item)
	}
	return MytoConfig{Items: items}, nil
}
