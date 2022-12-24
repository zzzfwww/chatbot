# train

## 处理数据源
```shell
ls *.yml |awk  '{print "mv "$0" en-"$0""}' |bash
```

## 训练样本
```shell
cd /chatbot/cli/train
go run train.go -d ../data 
```
