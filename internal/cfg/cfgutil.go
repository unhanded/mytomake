package cfg

import (
	"fmt"
	"strings"

	"github.com/unhanded/mytomake/internal/util"
)

type ConfigItem struct {
	Key      string
	FuncName string
	Args     []string
}

func parseCfgLine(ln string, lineN int) (ConfigItem, error) {
	ln = util.Untab(ln)
	if validationErr := validateConfigLine(ln); validationErr != nil {
		return ConfigItem{}, fmt.Errorf("%s in validation at config line %d\n\t%s", validationErr, lineN+1, ln)
	}

	parts, partsErr := util.TrimmedSplit(ln, ":")
	if partsErr != nil {
		return ConfigItem{}, fmt.Errorf("%s from from TrimmedSplit() at config line %d\n\t%s", partsErr, lineN+1, ln)
	}

	k := parts[0]
	tf := parts[1]
	tis, tisErr := util.TrimmedInnerSplit(tf, "[", "]", ",")
	if tisErr != nil {
		return ConfigItem{}, fmt.Errorf("%s from TrimmedInnerSplit() at config line %d\n\t%s", tisErr, lineN+1, ln)
	}

	cfgItem := ConfigItem{Key: k, FuncName: tf, Args: tis}
	return cfgItem, nil
}

func validateConfigLine(ln string) error {
	trimmed := strings.TrimSpace(ln)
	if util.IsEmptyLine(trimmed) {
		return fmt.Errorf("empty line")
	}

	if !canSplit(trimmed) {
		return fmt.Errorf("Unsplittable")
	}
	sections, sectionsErr := util.TrimmedSplit(trimmed, ":")
	if sectionsErr != nil {
		return sectionsErr
	}

	k := sections[0]
	if !keyIsOk(k) {
		return fmt.Errorf("invalid key")
	}

	tf := sections[1]
	if !isValidTemplateFuncCall(tf) {
		return fmt.Errorf("invalid func")
	}
	return nil
}

func isValidTemplateFuncCall(tf string) bool {
	if !strings.HasPrefix(tf, "@") {
		return false
	}
	if !(strings.Count(tf, "[") == strings.Count(tf, "]")) {
		return false
	}
	builtInFuncs := []string{"int", "float", "dict", "uuid"}
	for _, bif := range builtInFuncs {
		if strings.Contains(tf, bif) {
			return true
		}
	}
	return false
}

func keyIsOk(k string) bool {
	return (strings.Count(k, " ") == 0)
}

func canSplit(ln string) bool {
	return _hasExactlyOneSplitter(ln) && _isSplitInRange(ln)
}

func _hasExactlyOneSplitter(ln string) bool {
	return strings.Count(ln, ":") == 1
}

func _isSplitInRange(ln string) bool {
	length := len(ln)
	splitPoint := strings.Index(ln, ":")

	return (0 < splitPoint) && (splitPoint < length-1)
}
