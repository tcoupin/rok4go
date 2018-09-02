package log

import (
	"io/ioutil"
	"log"
	"os"
)

const (
	RedBold    = "\033[1;31m"
	YellowBold = "\033[1;33m"
	BlueBold   = "\033[1;34m"
	GreenBold  = "\033[1;32m"
	Bold       = "\033[1m"
	Reset      = "\033[0m"
	ResetBold  = "\033[22m"
)

type Logger struct {
	level    int
	logtrace *log.Logger
	logdebug *log.Logger
	loginfo  *log.Logger
	logwarn  *log.Logger
	logerror *log.Logger
}

const (
	LEVEL_TRACE = iota
	LEVEL_DEBUG
	LEVEL_INFO
	LEVEL_WARN
	LEVEL_ERROR
)

func (l Logger) trace(format string, v ...interface{}) {
	l.logtrace.Printf(format, v...)
}
func (l Logger) debug(format string, v ...interface{}) {
	l.logdebug.Printf(format, v...)

}
func (l Logger) info(format string, v ...interface{}) {
	l.loginfo.Printf(format, v...)
}
func (l Logger) warning(format string, v ...interface{}) {
	l.logwarn.Printf(format, v...)
}
func (l Logger) error(format string, v ...interface{}) {
	l.logerror.Printf(format, v...)
}

func (l *Logger) setLevel(level int) {
	l.level = level
	l.logtrace = l.initLevel(Reset+Bold+"[TRACE] "+ResetBold, LEVEL_TRACE)
	l.logdebug = l.initLevel(Reset+GreenBold+"[DEBUG] "+ResetBold, LEVEL_DEBUG)
	l.loginfo = l.initLevel(Reset+BlueBold+"[INFO]  "+ResetBold, LEVEL_INFO)
	l.logwarn = l.initLevel(Reset+YellowBold+"[WARN]  "+ResetBold, LEVEL_WARN)
	l.logerror = l.initLevel(Reset+RedBold+"[ERROR] "+ResetBold, LEVEL_ERROR)

}

func (l *Logger) initLevel(levelname string, level int) *log.Logger {
	if l.level > level {
		return log.New(ioutil.Discard, levelname, log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		return log.New(os.Stdout, levelname, log.Ldate|log.Ltime)
	}
}

var LOGGER *Logger

func SetLevel(level int) {
	if LOGGER == nil {
		LOGGER = &Logger{}
	}
	LOGGER.setLevel(level)
}

func TRACE(format string, v ...interface{}) {
	LOGGER.trace(format, v...)
}
func DEBUG(format string, v ...interface{}) {
	LOGGER.debug(format, v...)
}
func INFO(format string, v ...interface{}) {
	LOGGER.info(format, v...)
}
func WARNING(format string, v ...interface{}) {
	LOGGER.warning(format, v...)
}
func ERROR(format string, v ...interface{}) {
	LOGGER.error(format, v...)
}
