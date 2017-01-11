数据分为两类：

- 任务数据 : 每天计划任务时创建任务，每天回顾任务时根据规则给任务打分
- 时间数据 : 需要写爬虫，从Toggl爬取时间数据



### 任务数据结构

dailyTask：是否为日常任务（0：非日常任务，1：日常任务，是系统每天自动创建的）
priority：优先级（0：高，1：中，2：低）
score：得分（0-100之间）

```
{
  "createTime": "20170109",
  "dailyTask": 1,
  "finishDate": "",
  "planedDate": "20170109",
  "priority": 0,
  "score": 0,
  "taskName": "时间管理",
  "type": "计划和回顾"
}

{
  "createTime": "20170109",
  "dailyTask": 1,
  "finishDate": "20170109",
  "planedDate": "20170109",
  "priority": 0,
  "score": 0,
  "taskName": "时间管理",
  "type": "计划和回顾"
}
```

### 打分规则数据结构

每个type下有多个rules，每个type下的rules的score总和是100
period : 打分周期（0：临时，1：每天，2：每周，3：每月）

```
{
  "rules": [
    {
      "period": 0,
      "ruleName": "",
      "score": 100
    },
    {
      "period": 0,
      "ruleName": "",
      "score": 100
    }
  ],
  "type": "运动"
}
```

### 时间数据结构

```
{
  "client": "mobile",
  "costTime": "00:30:00",
  "endTime": "2017-01-10 00:30:00",
  "project": "运动",
  "startTime": "2017-01-10 00:00:00",
  "tag": "休息放松",
  "title": "体力锻炼"
}
```

### 启动项目步骤：

1. 执行build.sh脚本

```
$ sh build.sh
```

2. 执行startup.sh脚本

```
$ sh startup.sh
```
