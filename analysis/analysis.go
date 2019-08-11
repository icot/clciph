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
	mapping := map[byte]byte{}
	for _, letter := range bytes {
		if letter != 0x20 {
			_, ok := mapping[letter]
			if !ok {
				mapping[letter] = letter
			}
		}
	}
	return mapping
}

func getFreqs(bytes []byte) map[byte]float32 {
	table := map[byte]float32{}
	for _, letter := range bytes {
		if letter != 0x20 {
			_, ok := table[letter]
			if ok {
				table[letter]++
			} else {
				table[letter] = 1
			}
		}
	}
	// Convert value to percentage
	for k, v := range table {
		table[k] = 100 * (v / float32(len(bytes)))
	}

	return table
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
