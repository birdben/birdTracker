package model

import (
	"github.com/golang/glog"
	"jsonlog"
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
	glog.V(10).Infof("taskInfo=%v", taskInfo)

	jsonData, err := jsonlog.LoggerJsonObj(taskInfo)

	glog.Infof("SaveTaskInfo jsonData=%v", jsonData)
	glog.V(10).Info("SaveTaskInfo end")
	return
}

func BatchSaveTaskInfo(taskInfoList []TaskInfo) (err error) {
	glog.V(10).Info("BatchSaveTaskInfo start")
	glog.V(10).Infof("taskInfo=%v", taskInfoList)

	//jsonData, err := LoggerJsonArray(taskInfoList)
	var interfaceSlice []interface{} = make([]interface{}, len(taskInfoList))
	for i, d := range taskInfoList {
		interfaceSlice[i] = d
	}
	jsonData, err := jsonlog.LoggerJsonArray(interfaceSlice)

	glog.Infof("BatchSaveTaskInfo jsonData=%v", jsonData)
	glog.V(10).Info("BatchSaveTaskInfo end")
	return
}
