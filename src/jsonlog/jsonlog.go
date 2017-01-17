package jsonlog

import (
	"encoding/json"
	"github.com/golang/glog"
	"reflect"
)

func LoggerJsonObj(obj interface{}) (jsonData string, err error) {
	glog.V(10).Info("LoggerJsonObj start")
	var interfaceObj interface{} = obj
	entry := reflectField(interfaceObj)
	jsonDataByte, err := json.Marshal(entry)
	jsonData = string(jsonDataByte)
	glog.V(10).Infof("LoggerJsonObj jsonData=%v", jsonData)
	glog.V(10).Info("LoggerJsonObj end")
	return
}

func LoggerJsonArray(objs []interface{}) (jsonData string, err error) {
	glog.V(10).Info("LoggerJsonArray start")
	entryList := make([]map[string]interface{}, 0)
	for _, interfaceObj := range objs {
		entry := reflectField(interfaceObj)
		entryList = append(entryList, entry)
	}

	jsonDataByte, err := json.Marshal(entryList)
	jsonData = string(jsonDataByte)
	glog.V(10).Infof("LoggerJsonArray jsonData=%v", jsonData)
	glog.V(10).Info("LoggerJsonArray end")
	return
}

func reflectField(interfaceObj interface{}) (entry map[string]interface{}) {
	reflectObj := reflect.ValueOf(interfaceObj)
	typeOfType := reflectObj.Type()
	glog.V(10).Infof("ReflectField interfaceObj=%v", interfaceObj)
	glog.V(10).Infof("ReflectField reflectObj=%v", reflectObj)
	glog.V(10).Infof("ReflectField typeOfType=%v", typeOfType)
	entry = make(map[string]interface{})
	for i := 0; i < reflectObj.NumField(); i++ {
		field := reflectObj.Field(i)
		glog.V(10).Infof("ReflectField field=%v", field)
		var v interface{}
		val := field.Interface()
		b, ok := val.([]byte)
		if ok {
			v = string(b)
		} else {
			v = val
		}
		glog.V(10).Infof("ReflectField json_field=%v", typeOfType.Field(i).Name)
		glog.V(10).Infof("ReflectField json_value=%v", v)
		entry[typeOfType.Field(i).Name] = v
	}
	return
}
