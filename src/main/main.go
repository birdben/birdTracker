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

	taskInfo := model.TaskInfo{}
	taskInfo.Id = "1"
	taskInfo.TaskName = "GitHub"
	taskInfo.TypeName = "高效工作"
	taskInfo.Priority = 0
	taskInfo.DailyTask = 0
	taskInfo.CreateTime = dateStr
	taskInfo.PlanedDate = dateStr
	taskInfo.FinishDate = dateStr
	taskInfo.Score = 0

	model.SaveTaskInfo(taskInfo)
	glog.Info("main end")
	glog.Flush()
}
