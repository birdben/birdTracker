package main

import (
	"flag"
	"model"
	"time"
	//"config"
	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	glog.Info("main start")
	now := time.Now()
	dateStr := now.Format("20160102")

	taskInfo1 := model.TaskInfo{}
	taskInfo1.Id = "1"
	taskInfo1.TaskName = "GitHub1"
	taskInfo1.TypeName = "高效工作1"
	taskInfo1.Priority = 0
	taskInfo1.DailyTask = 0
	taskInfo1.CreateTime = dateStr
	taskInfo1.PlanedDate = dateStr
	taskInfo1.FinishDate = dateStr
	taskInfo1.Score = 0

	taskInfo2 := model.TaskInfo{}
	taskInfo2.Id = "2"
	taskInfo2.TaskName = "GitHub2"
	taskInfo2.TypeName = "高效工作2"
	taskInfo2.Priority = 0
	taskInfo2.DailyTask = 0
	taskInfo2.CreateTime = dateStr
	taskInfo2.PlanedDate = dateStr
	taskInfo2.FinishDate = dateStr
	taskInfo2.Score = 0

	model.SaveTaskInfo(taskInfo1)

	var taskInfoList = make([]model.TaskInfo, 2)
	taskInfoList[0] = taskInfo1
	taskInfoList[1] = taskInfo2

	model.BatchSaveTaskInfo(taskInfoList)

	glog.Info("main end")
	glog.Flush()
}
