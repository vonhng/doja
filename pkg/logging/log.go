// username: vonhng
// create_time: 2020/1/1 - 23:42
// mail: vonhng.feng@gmail.com
package logging

import (
	"doja/pkg/config"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type Level int

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

var (
	F                  *os.File
	err                error
	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARNING", "ERROR", "FATALF"}
)

// setPrefix set the prefix of the log output
func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}
	logger.SetPrefix(logPrefix)
}

func SetUpLog() {
	logName := fmt.Sprintf("doja-%s.log", time.Now().Format("20060102"))
	fmt.Printf(config.Setting.Logging.LogPath)
	logFullPath := config.Setting.Logging.LogPath + "/" + logName

	F, err = os.OpenFile(logFullPath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("open %s failed: %v", logFullPath, err)
	}
	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

// Debug output logs at debug level
func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}

// Info output logs at info level
func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v)
}

// Warn output logs at warn level
func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v)
}

// Error output logs at error level
func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v)
}

// Fatal output logs at fatal level
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v)
}
