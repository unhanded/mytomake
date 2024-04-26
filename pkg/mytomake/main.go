package mytomake

import (
	"github.com/unhanded/mytomake/internal/myto_api"
)

type Myto struct {
	DataDir string
}

func (m Myto) FromFile(filepath string) (string, error) {
	return myto_api.MytoFileExec(filepath, m.DataDir)
}

func (m Myto) FromBytes(data []byte) (string, error) {
	return myto_api.MytoDirectDataExec(data, m.DataDir)
}

func NewMyto(dataDir string) *Myto {
	return &Myto{
		DataDir: dataDir,
	}
}
