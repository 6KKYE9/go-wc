# go-wc

统计文本的行数、词数、字节数，类似 `wc`。支持 `-l/-w/-c` 单独显示，默认三项都显示。

## 安装

```bash
go build -o go-wc.exe
```

## 用法

```bash
go-wc file.txt
go-wc -l file.txt       # 只数行
echo "a b c" | go-wc -w # 只数词
```

## 说明

零依赖纯 Go。没有结尾换行的一行也算一行，和常见 wc 行为一致。
