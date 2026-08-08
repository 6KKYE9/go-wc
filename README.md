# go-wc

复制粘贴改格式改到手酸？这玩意儿一行就搞定。

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
