package core

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
)

type MytoDict struct {
	DictName string
	DataDir  string
}

func (mtd MytoDict) Value() string {
	lookup, err := mtd.GetDict()
	if err != nil {
		return "DICT_ERROR"
	}
	rnd, rndErr := lookup.Random()
	if rndErr != nil {
		return "LOOKUP_ERROR"
	}
	return rnd
}

func (mtd MytoDict) GetDict() (MytoDictLookup, error) {
	expectedFilename := fmt.Sprintf("%s.mytodict", mtd.DictName)
	dictsDir := mtd.DataDir + "/dicts"
	contents, err := os.ReadDir(dictsDir)
	if err != nil {
		return MytoDictLookup{}, err
	}
	lookup := MytoDictLookup{}
	for _, f := range contents {
		if f.Name() == expectedFilename {
			lookup.Load(dictsDir + "/" + f.Name())
			return lookup, nil
		}
	}
	return MytoDictLookup{}, fmt.Errorf("DICTIONARY_ERROR")
}

type MytoDictLookup struct {
	DictName string
	Entries  []string
}

func (mtdl *MytoDictLookup) Load(fp string) error {
	f, err := os.Open(fp)
	if err != nil {
		return err
	}
	defer f.Close()

	fileInfo, statErr := f.Stat()

	if statErr != nil {
		return statErr
	}

	var dictBytes []byte = make([]byte, fileInfo.Size()+128)
	_, fileReadErr := f.Read(dictBytes)
	if fileReadErr != nil {
		return fileReadErr
	}
	return mtdl.enter(dictBytes)
}

func (mtdl *MytoDictLookup) enter(b []byte) error {
	items := bytes.Split(b, []byte("\n"))
	for _, item := range items {
		if len(item) > 0 {
			mtdl.Entries = append(mtdl.Entries, string(item))
		}
	}
	return nil
}

func (mtdl MytoDictLookup) Random() (string, error) {
	if len(mtdl.Entries) == 0 {
		return "", fmt.Errorf("DICTIONARY_EMPTY")
	}
	idx := rand.Int31n(int32(len(mtdl.Entries)))
	return mtdl.Entries[idx], nil
}
