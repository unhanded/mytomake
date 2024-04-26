package myto_api

import (
	"os"

	"github.com/unhanded/mytomake/internal/files"
	"github.com/unhanded/mytomake/internal/swap"
)

func MytoFileExec(fp string, dataDir string) (string, error) {
	b, bErr := os.ReadFile(fp)
	if bErr != nil {
		return "", bErr
	}

	res, err := MytoDirectDataExec(b, dataDir)
	if err != nil {
		return "", err
	}

	return res, nil
}

func MytoDirectDataExec(data []byte, dataDir string) (string, error) {
	tbf := files.MytoFile{}

	loadErr := tbf.FromBytes(data)
	if loadErr != nil {
		return "", loadErr
	}

	parseErr := tbf.Parse()
	if parseErr != nil {
		return "", parseErr
	}

	tbf.Config.DataDir = dataDir

	out, outErr := swap.Generate([]byte(tbf.TemplateSection), tbf.Config)
	if outErr != nil {
		return "", outErr
	}
	return string(out), nil
}
