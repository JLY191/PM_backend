# 2023-2024秋冬 软件工程管理 后端

## 如何使用

### 后端端口：8080

### config.yaml修改

用户名和密码修改成本地MySQL的

![](images/conf.png)


### MySQL建库

连接MySQL，建立数据库trip2world
```sql
create database trip2world;
```

### 下载go

去https://go.dev/dl/下载go的二进制文件包，如果是Windows操作系统需要把go添加到环境变量

![](images/godow.png)

go环境变量添加参考https://blog.csdn.net/liu_chen_yang/article/details/132012969

### 打开Go Module

```bash
$ go env -w GO111MODULE=on
```

### 然后运行本项目

Windows：

```bash
$ go mod tidy
# 等待依赖下载完...
$ go run ./
# 或者go build编译后运行.exe
```

macOS：

```bash
$ go mod tidy
# 等待依赖下载完...
$ go run *.go
# 或者go build编译后运行.exec
```
### init.sql

运行后，表应当被建立了。然后需要插入数据，复制并用DataGrip、MySQL Workbench等执行
