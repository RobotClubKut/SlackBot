package log

//errorの制御関係

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

//Error is error data structs.
type Error struct {
	Time    string
	Error   string
	Message string
}

//ErrorData is Error array.
type ErrorData struct {
	Errors []Error
}

//WriteErrorLog is output err log.
func WriteErrorLog(err error) {
	WriteErrorLogAndMessage(err, "")
}

//Terminate is abnormal exit.
func Terminate(err error) {
	WriteErrorLog(err)
	log.Fatalln(err)
}

//TerminateAndWriteMessage is abnormal exit.
func TerminateAndWriteMessage(err error, msg string) {
	WriteErrorLogAndMessage(err, msg)
	log.Fatalln(err)
}

//WriteErrorLogAndMessage is output err log.
func WriteErrorLogAndMessage(err error, msg string) {
	fileName := "../logs/error.json"
	bufError := ReadErrorLog()
	errData := Error{Time: time.Now().String(), Error: err.Error(), Message: msg}
	var w ErrorData
	if bufError == nil {
		var errors ErrorData
		errors.Errors = append(errors.Errors, errData)
		w = errors

	} else {
		bufError.Errors = append(bufError.Errors, errData)
		w = *bufError
	}

	js, _ := json.Marshal(w)
	ioutil.WriteFile(fileName, js, 0644)
}

//ReadErrorLog is read error.log file
func ReadErrorLog() *ErrorData {
	fileName := "../logs/error.json"
	bin, _ := ioutil.ReadFile(fileName)
	var js ErrorData
	err := json.Unmarshal(bin, &js)
	if err != nil {
		return nil
	}
	return &js
}
