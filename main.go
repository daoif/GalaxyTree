package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

//go:embed index.html
//go:embed png/*
var content embed.FS

func main() {
	var listener net.Listener
	var err error

	// 试图优先监听 1225-1230 端口
	preferredPorts := []int{1225, 1226, 1227, 1228, 1229, 1230}
	for _, p := range preferredPorts {
		listener, err = net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", p))
		if err == nil {
			break // 成功拿到端口
		}
		fmt.Printf("⚠️ 端口 %d 被占用或不可用: %v\n", p, err)
	}

	// 如果所有优先端口都失败，回退到随机端口
	if listener == nil {
		fmt.Printf("⚠️ 所有优先端口 (1225-1230) 均不可用，尝试自动分配...\n")
		listener, err = net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			panic(fmt.Sprintf("无法启动服务器: %v", err))
		}
	}
	
	port := listener.Addr().(*net.TCPAddr).Port

	// 1. 优先处理外部配置文件请求 (允许用户分发 galaxy_tree_config.json)
	http.HandleFunc("/config.json", func(w http.ResponseWriter, r *http.Request) {
		// 尝试读取运行目录下的 json
		data, err := os.ReadFile("galaxy_tree_config.json")
		if err == nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
			return
		}
		// 找不到则返回 404 (前端会处理)
		http.NotFound(w, r)
	})

	// 2. 创建文件服务器 (Embed FS)
	http.Handle("/", http.FileServer(http.FS(content)))

	// 启动服务器
	addr := fmt.Sprintf("127.0.0.1:%d", port)
	url := fmt.Sprintf("http://%s", addr)

	fmt.Printf("🎄 圣诞贺卡服务器启动！\n")
	fmt.Printf("📍 地址: %s\n", url)
	fmt.Printf("按 Ctrl+C 退出\n\n")

	// 自动打开浏览器
	go openBrowser(url)

	// 启动 HTTP 服务 (使用已创建的 listener)
	if err := http.Serve(listener, nil); err != nil {
		panic(err)
	}
}

func openBrowser(url string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}
	cmd.Run()
}

// 用于获取子目录的辅助函数（如果需要）
func getSubFS(fsys embed.FS, dir string) (fs.FS, error) {
	return fs.Sub(fsys, dir)
}
