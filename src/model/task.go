package model

import (
	"encoding/json"
	"github.com/golang/glog"
	"reflect"
)

// 注意这里定义struct内部变量一定要大写开头，否则外部引用model.TaskInfo是无法给struct内部属性赋值的
type TaskInfo struct {
	Id         string
	TaskName   string
	TypeName   string
	Priority   int32
	DailyTask  int32
	CreateTime string
	PlanedDate string
	FinishDate string
	Score      int32
}

func SaveTaskInfo(taskInfo TaskInfo) (err error) {
	glog.V(10).Info("SaveTaskInfo start")
	glog.V(10).Info("taskInfo=%v", taskInfo)

	jsonData, err := LoggerJsonObj(taskInfo)

	glog.Infof("SaveTaskInfo jsonData=%v", jsonData)
	glog.V(10).Info("SaveTaskInfo end")
	return
}

func BatchSaveTaskInfo(taskInfoList []TaskInfo) (err error) {
	glog.V(10).Info("BatchSaveTaskInfo start")
	glog.V(10).Info("taskInfo=%v", taskInfoList)

	//jsonData, err := LoggerJsonArray(taskInfoList)
	var interfaceSlice []interface{} = make([]interface{}, len(taskInfoList))
	for i, d := range taskInfoList {
		interfaceSlice[i] = d
	}
	jsonData, err := LoggerJsonArray(interfaceSlice)

	glog.Infof("BatchSaveTaskInfo jsonData=%v", jsonData)
	glog.V(10).Info("BatchSaveTaskInfo end")
	return
}

func LoggerJsonObj(obj interface{}) (jsonData string, err error) {
	glog.V(10).Info("LoggerJsonObj start")
	var interfaceObj interface{} = obj
	reflectObj := reflect.ValueOf(interfaceObj)
	typeOfType := reflectObj.Type()
	entry := make(map[string]interface{})
	for i := 0; i < reflectObj.NumField(); i++ {
		field := reflectObj.Field(i)
		var v interface{}
		val := field.Interface()
		b, ok := val.([]byte)
		if ok {
			v = string(b)
		} else {
			v = val
		}
		glog.V(10).Infof("LoggerJsonObj json_field=%v", typeOfType.Field(i).Name)
		glog.V(10).Infof("LoggerJsonObj json_value=%v", v)
		entry[typeOfType.Field(i).Name] = v
	}
	jsonDataByte, err := json.Marshal(entry)
	jsonData = string(jsonDataByte)
	glog.V(10).Infof("LoggerJsonObj jsonData=%v", jsonData)
	glog.V(10).Info("LoggerJsonObj end")
	return
}

func LoggerJsonArray(objs []interface{}) (jsonData string, err error) {
	glog.V(10).Info("LoggerJsonArray start")
	tableData := make([]map[string]interface{}, 0)
	for _, obj := range objs {
		reflectObj := reflect.ValueOf(obj)
		glog.V(10).Infof("LoggerJsonArray obj=%v", obj)
		glog.V(10).Infof("LoggerJsonArray reflectObj=%v", reflectObj)
		entry := make(map[string]interface{})
		for i := 0; i < reflectObj.NumField(); i++ {
			typeOfType := reflectObj.Type()
			field := reflectObj.Field(i)
			glog.V(10).Infof("LoggerJsonArray typeOfType=%v", typeOfType)
			glog.V(10).Infof("LoggerJsonArray field=%v", field)
			var v interface{}
			val := field.Interface()
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			glog.V(10).Infof("LoggerJsonArray json_field=%v", typeOfType.Field(i).Name)
			glog.V(10).Infof("LoggerJsonArray json_value=%v", v)
			entry[typeOfType.Field(i).Name] = v
		}
		tableData = append(tableData, entry)
	}

	jsonDataByte, err := json.Marshal(tableData)
	jsonData = string(jsonDataByte)
	glog.V(10).Infof("LoggerJsonArray jsonData=%v", jsonData)
	glog.V(10).Info("LoggerJsonArray end")
	return
}
