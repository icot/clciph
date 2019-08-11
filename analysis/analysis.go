package analysis

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	prettyjson "github.com/hokaccha/go-prettyjson"
	"github.com/icot/clciph/tools/file"
)

type Analysis struct {
	Cyphertext string
	Bytes      []byte
	Mapping    map[byte]byte
	Freqs      map[byte]float32
}

func getMapping(bytes []byte) map[byte]byte {
	return nil
}

func getFreqs(bytes []byte) map[byte]float32 {
	return nil
}

func Analyze(ciphertext []byte) *Analysis {
	a := new(Analysis)
	a.Cyphertext = string(ciphertext)
	a.Bytes = ciphertext
	a.Mapping = getMapping(a.Bytes)
	a.Freqs = getFreqs(a.Bytes)
	return a
}

func AnalyzeFile(filename string) *Analysis {
	log.Debug(fmt.Sprintf("Loading %s", filename))
	contents := file.Load(filename)
	log.Debug(contents)
	buf := Analyze(contents)
	str, _ := prettyjson.Marshal(buf)
	log.Debug("Dumping Analisis")
	log.Debug(string(str))
	return Analyze(contents)
}
