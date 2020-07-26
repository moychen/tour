# 工具集
使用cobra构建自己的工具集

## 数据库表结构转Struct
**go run main.go sql struct -t=编程语言类型 --username=用户名 --password=密码 --db=数据库名 --table=表名**
```
-t: 支持C++、Go;
```
```bash
$ go run main.go sql struct -t=GO --username=root --password=*** --db=evdata --table=csv2db
```


## 时间处理
### 获取当前时间
```bash
$ go run main.go time now
```
### 计算时间
**go run main.go time calc -c=起始时间 -d=时长**
```
-c: 支持timestamp和“2006-01-02 15:04:05”两种格式。
-d: 支持多种格式：
    如24h6m等。
````
```bash
$ go run main.go time calc -c='2020-07-26' -d=24h 
```

## 单词格式转换
**go run main.go word -s=单词内容 -m=模式**
```
-s: 单词内容
-m: 单词转换模式
	1：全部单词转为大写
	2：全部单词转为小写
	3：下划线单词转为大写驼峰单词
	4：下划线单词转为小写驼峰单词
	5：驼峰单词转为下划线单词
```
```bash
$ go run main.go word -s=hello -m=1
```