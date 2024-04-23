package swap

import (
	"fmt"

	"github.com/unhanded/mytoman/internal/cfg"
	"github.com/unhanded/mytoman/internal/core"
)

func Generate(tmplSection []byte, cfg cfg.MytoConfig) ([]byte, error) {
	processed := string(tmplSection)
	for _, rp := range cfg.Items {
		blaster, blasterErr := getMyto(rp.FuncName, cfg)
		if blasterErr != nil {
			return nil, blasterErr
		}
		targetKey := fmt.Sprintf("@%s", rp.Key)
		newStr := core.Replace(blaster, processed, targetKey)
		processed = newStr
	}
	return []byte(processed), nil
}
