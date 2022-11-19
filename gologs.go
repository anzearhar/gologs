// Package gologs implements functions for logging messages to the terminal
// and files.
package gologs

import (
    "fmt"
    "time"
)

const (
    format = "2006-01-02 15:04:05.000"

    typeDebug = "DEBU"
    typeInformation = "INFO"
    typeNotice = "NOTI"
    typeWarning = "WARN"
    typeError = "ERRO"
    typeCritical = "CRIT"

    colorReset = "\x1b[0m"

    colorNil = ""
    colorForegroundBlack = "\x1b[30m"
    colorForegroundRed = "\x1b[31m"
    colorForegroundGreen = "\x1b[32m"
    colorForegroundYellow = "\x1b[33m"
    colorForegroundBlue = "\x1b[34m"
    colorForegroundMagenta = "\x1b[35m"
    colorForegroundCyan = "\x1b[36m"
    colorForegroundWhite = "\x1b[37m"
)

// log formats and prints the given message
func log(message string, logType string, colorForeground string, printInColor bool) {
    timestamp := "[" + time.Now().Format(format) + "] "
    logType = "[" + logType + "]"
    logPrefix := timestamp + logType
    if printInColor {
        logPrefix = colorForeground + logPrefix + colorReset
    }
    fmt.Println(logPrefix, message)
}

// Debug prints a colored debug message to stdout
func Debug(message string) {
    log(message, typeDebug, colorForegroundCyan, true);
}

// DebugUncolored prints a debug message to stdout
func DebugUncolored(message string) {
    log(message, typeDebug, colorNil, false);
}

// Info prints a colored information message to stdout
func Info(message string) {
    log(message, typeInformation, colorForegroundWhite, true);
}

// InfoUncolored prints an information message to stdout
func InfoUncolored(message string) {
    log(message, typeInformation, colorNil, false);
}

// Notice prints a colored notification message to stdout
func Notice(message string) {
    log(message, typeNotice, colorForegroundGreen, true);
}

// NoticeUncolored prints a notification message to stdout
func NoticeUncolored(message string) {
    log(message, typeNotice, colorNil, false);
}

// Warning prints a colored warning message to stdout
func Warning(message string) {
    log(message, typeWarning, colorForegroundYellow, true);
}

// WarningUncolored prints a warning message to stdout
func WarningUncolored(message string) {
    log(message, typeWarning, colorNil, false);
}

// Error prints a colored error message to stdout
func Error(message string) {
    log(message, typeError, colorForegroundRed, true);
}

// ErrorUncolored prints an error message to stdout
func ErrorUncolored(message string) {
    log(message, typeError, colorNil, false);
}

// Critical prints a colored critical message to stdout
func Critical(message string) {
    log(message, typeCritical, colorForegroundMagenta, true);
}

// CriticalUncolored prints a critical message to stdout
func CriticalUncolored(message string) {
    log(message, typeCritical, colorNil, false);
}
