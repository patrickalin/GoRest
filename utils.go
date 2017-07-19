package http

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

func funcName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

func logFatal(err error, fct string, msg string, params ...string) {
	logrus.WithFields(logrus.Fields{
		"param": fmt.Sprintf(strings.Join(params[:], ",")),
		"error": err,
		"fct":   fct,
	}).Warn(msg)
	log.WithFields(logrus.Fields{
		"param": fmt.Sprintf(strings.Join(params[:], ",")),
		"error": err,
		"fct":   fct,
	}).Fatal(msg)
}

func logDebug(fct string, msg string, params ...string) {
	log.WithFields(logrus.Fields{
		"param": fmt.Sprintf(strings.Join(params[:], ",")),
		"fct":   fct,
	}).Debug(msg)
}

func checkErr(err error, fct string, msg string, params ...string) {
	if err != nil {
		logFatal(err, msg, fct, fmt.Sprintf(strings.Join(params[:], ",")))
	}
}
