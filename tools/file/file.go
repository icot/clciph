package file

import (
	"fmt"
	"io/ioutil"

	log "github.com/Sirupsen/logrus"
)

func Load(filename string) []byte {

	log.Debug(fmt.Sprintf("Loading: %s", filename))
	buf, _ := ioutil.ReadFile(filename)
	return buf

}
