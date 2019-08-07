package analysis

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
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

func Analyze(ciphertext string) Analisis {
	a := Analisis{}
	a.Cyphertext = ciphertext
	a.Bytes = []byte(ciphertext)
	a.Mapping = getMapping(a.Bytes)
	a.Freqs = getFreqs(a.Bytes)
	return a
}

func AnalyzeFile(filename string) Analisis {
	log.Debug(fmt.Sprintf("Loading %s", words))
	contents := file.Load(filename)
	log.Debug(contents)
	return Analyze(contents)
}
