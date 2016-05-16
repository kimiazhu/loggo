// Description: loggo
// Author: ZHU HAIHUA
// Since: 2016-05-02 21:47
package loggo

import (
	"github.com/kimiazhu/slf4g"
)

func init() {
	slf4g.Loggers = make([]Loggo, 0)
	// load configuration
	LoadConfiguration()
}


type Loggo struct {

}

func (log *Loggo) Finest(arg0 interface{}, args ...interface{}) {

}

func (log *Loggo) Fine(arg0 interface{}, args ...interface{}) {

}

func (log *Loggo) Debug(arg0 interface{}, args ...interface{}) {

}

func (log *Loggo) Trace(arg0 interface{}, args ...interface{}) {

}

func (log *Loggo) Info(arg0 interface{}, args ...interface{}) {

}

func (log *Loggo) Warn(arg0 interface{}, args ...interface{}) {

}

func (log *Loggo) Error(arg0 interface{}, args ...interface{}) {

}

func (log *Loggo) Fatal(arg0 interface{}, args ...interface{}) {

}