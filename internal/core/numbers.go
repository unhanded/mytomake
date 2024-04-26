package core

import (
	"fmt"
	"math/rand"

	"github.com/unhanded/mytomake/internal/util"
)

type MytoFloat struct {
	Args util.RangeStringArgs
}

func (tbflt MytoFloat) Value() string {
	min, max, err := tbflt.Args.ToFloat()
	if err != nil {
		return err.Error()
	}

	r := rand.Float64()*(max-min) + min
	b := fmt.Sprintf("%f", r)

	return b
}

type MytoInt struct {
	Args util.RangeStringArgs
}

func (tbint MytoInt) Value() string {
	min, max, err := tbint.Args.ToInt()
	if err != nil {
		return err.Error()
	}

	r := rand.Int63n(max-min) + min
	b := fmt.Sprintf("%d", r)

	return b
}
