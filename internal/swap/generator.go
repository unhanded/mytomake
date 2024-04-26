package swap

import (
	"fmt"

	"github.com/unhanded/mytomake/internal/cfg"
	"github.com/unhanded/mytomake/internal/core"
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
