package model

import (
	"github.com/golang/glog"
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
	glog.Info("SaveTaskInfo start")
	glog.Info("taskInfo=", taskInfo)
	glog.Info("SaveTaskInfo end")
	return
}
