package jsonlog

import (
	"flag"
	"github.com/golang/glog"
	"testing"
)

type Person struct {
	Id      string
	Name    string
	Age     int32
	Company string
	Job     string
	Address string
}

func Test_LoggerJsonObj(t *testing.T) {
	flag.Parse()
	glog.Info("TestLoggerJsonObj start")

	person := Person{}
	person.Id = "1"
	person.Name = "birdben"
	person.Age = 20
	person.Company = "BAT"
	person.Job = "程序员"
	person.Address = "北京"

	jsonData, err := LoggerJsonObj(person)
	if err != nil {
		glog.Errorf("TestLoggerJsonObj error: %v", err)
	}
	glog.Infof("TestLoggerJsonObj jsonData=%v", jsonData)

	glog.Info("TestLoggerJsonObj end")
	glog.Flush()
	return
}

func Test_LoggerJsonArray(t *testing.T) {
	flag.Parse()
	glog.Info("TestLoggerJsonArray start")

	person1 := Person{}
	person1.Id = "1"
	person1.Name = "birdben"
	person1.Age = 20
	person1.Company = "BAT"
	person1.Job = "程序员"
	person1.Address = "北京"

	person2 := Person{}
	person2.Id = "2"
	person2.Name = "zhangsan"
	person2.Age = 30
	person2.Company = "银行"
	person2.Job = "业务员"
	person2.Address = "上海"

	var personList = make([]Person, 2)
	personList[0] = person1
	personList[1] = person2

	var interfaceSlice []interface{} = make([]interface{}, len(personList))
	for i, d := range personList {
		interfaceSlice[i] = d
	}
	jsonData, err := LoggerJsonArray(interfaceSlice)
	if err != nil {
		glog.Errorf("TestLoggerJsonArray error: %v", err)
	}
	glog.Infof("TestLoggerJsonArray jsonData=%v", jsonData)

	glog.Info("TestLoggerJsonArray end")
	glog.Flush()
	return
}
