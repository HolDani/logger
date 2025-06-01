package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var FilePath = "log.txt"

const (
	LogError   = "ERROR"
	LogWarning = "WARNING"
	LogInfo    = "INFO"
	LogDebug   = "DEBUG"
)

func Info(data interface{}, mode string) {
	b, err := checkMode(mode)
	if err != nil {
		fmt.Println(err)
		return
	}
	if b {
		fmt.Println("[INFO] ", data)
	} else {
		logWriter(LogInfo, data)
	}

}

func Error(data interface{}, mode string) {
	b, err := checkMode(mode)
	if err != nil {
		fmt.Println(err)
		return
	}
	if b {
		fmt.Println("[ERROR] ", data)
	} else {
		logWriter(LogError, data)
	}
}

func Warning(data interface{}, mode string) {
	b, err := checkMode(mode)
	if err != nil {
		fmt.Println(err)
		return
	}
	if b {
		fmt.Println("[WARNING] ", data)
	} else {
		logWriter(LogWarning, data)
	}
}

func Debug(data interface{}, mode string) {
	b, err := checkMode(mode)
	if err != nil {
		fmt.Println(err)
		return
	}
	if b {
		fmt.Println("[DEBUG] ", data)
	} else {
		logWriter(LogDebug, data)
	}
}

func checkMode(mode string) (bool, error) {
	switch mode {
	case "stdout":
		return true, nil
	case "file":
		return false, nil
	default:
		return false, fmt.Errorf("Unsupported mode: %s", mode)
	}
}

func logWriter(level, data interface{}) {
	err := os.MkdirAll(filepath.Dir(FilePath), os.ModePerm)
	if err != nil {
		fmt.Println("Failed to create directories:", err)
		return
	}
	file, err := os.OpenFile(FilePath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	outData := fmt.Sprintf("[%s] %v [%s]\n", level, data, time.Now().Format("2006-01-02 15:04:05"))
	file.WriteString(outData)

	defer file.Close()
}
