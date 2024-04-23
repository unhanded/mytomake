package swap

import (
	"fmt"
	"strings"

	"github.com/unhanded/mytoman/internal/cfg"
	"github.com/unhanded/mytoman/internal/core"
	"github.com/unhanded/mytoman/internal/util"
)

func getMyto(t string, mcfg cfg.MytoConfig) (core.Myto, error) {
	if strings.HasPrefix(t, "@float") {
		args, err := util.TrimmedInnerSplit(t, "[", "]", ",")
		if err != nil {
			return nil, err
		}
		rArgs := util.RangeStringArgs{Min: args[0], Max: args[1]}
		return core.MytoFloat{Args: rArgs}, nil
	}

	if strings.HasPrefix(t, "@int") {
		args, err := util.TrimmedInnerSplit(t, "[", "]", ",")
		if err != nil {
			return nil, err
		}
		rArgs := util.RangeStringArgs{Min: args[0], Max: args[1]}
		return core.MytoInt{Args: rArgs}, nil
	}

	if strings.HasPrefix(t, "@dict") {
		args, err := util.TrimmedInnerSplit(t, "[", "]", ",")
		if err != nil {
			return nil, err
		}
		return core.MytoDict{DictName: args[0], DataDir: mcfg.DataDir}, nil
	}

	if strings.HasPrefix(t, "@uuid") {
		return core.MytoUUID{}, nil
	}
	return nil, fmt.Errorf("invalid func name")
}
