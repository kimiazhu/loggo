// Description: loggo
// Author: ZHU HAIHUA
// Since: 2016-05-02 21:46
package loggo

import (
	"os/exec"
	"os"
	"path/filepath"
)

func LoadConfiguration() {
	// auto load config from default position
	//Logger = NewDefaultLogger(DEBUG)
	file, _ := exec.LookPath(os.Args[0])
	dir := filepath.Dir(file)
	if _, err := os.Stat("log4go.xml"); !os.IsNotExist(err) {
		loadConfiguration("loggo.xml")
	} else if _, err := os.Stat(filepath.Join(dir, "/loggo.xml")); !os.IsNotExist(err) {
		loadConfiguration(filepath.Join(dir, "loggo.xml"))
	} else if _, err := os.Stat(filepath.Join(dir, "/conf/loggo.xml")); !os.IsNotExist(err) {
		loadConfiguration(filepath.Join(dir, "/conf/loggo.xml"))
	} else {
		//fmt.Fprintf(os.Stderr, "log4go config not found, exec dir is: %s, u need to load it by yourself.\n", dir)
	}
}

func loadConfiguration(filepath string) {

}