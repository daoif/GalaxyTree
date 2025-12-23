# 🎄 圣诞贺卡打包方案

## 项目结构

```
12.25/
├── main.go       # Go 服务器源码
├── go.mod        # Go 模块定义
├── index.html    # 主页面
└── png/
    └── g1.png    # 金币纹理图片
```

## 使用方法

### 开发模式（需安装 Go）

```bash
# 进入项目目录
cd e:\code\12.25

# 启动开发服务器
go run main.go
```

浏览器会自动打开贺卡页面。

### 编译打包（生成 exe）

```bash
# Windows 下编译（当前系统）
go build -ldflags="-s -w" -o christmas.exe main.go

# 跨平台编译
# Windows:
set GOOS=windows
set GOARCH=amd64
go build -ldflags="-s -w" -o christmas.exe main.go

# Mac:
set GOOS=darwin
set GOARCH=amd64
go build -ldflags="-s -w" -o christmas_mac main.go
```

### 分发

编译后只需分发 `christmas.exe` 一个文件即可！
- 双击运行
- 自动打开浏览器
- 无需其他依赖

## 安装 Go（如未安装）

下载地址：https://go.dev/dl/

下载 Windows 版本安装后重启终端即可使用。
