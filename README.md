# iptv-proxy

[![Built with Claude Code](https://img.shields.io/badge/Built%20with-Claude%20Code-D97757?logo=anthropic&logoColor=white)](https://claude.com/claude-code)

一个极简的 IPTV 直播流反向代理。把本地端口的请求原样转发到上游 IPTV 源，
保留原始 path/query，并对直播流做即时 flush（不缓冲），适合在本地或路由器上
转发运营商组播/单播源。

## 工作原理

基于 Go 标准库的 `httputil.ReverseProxy`：

- 监听本地地址，将收到的请求转发到上游 IPTV 源；
- `FlushInterval: -1`，每次写入立即 flush，保证直播流低延迟、不积压；
- 不改写 path 与 query，仅替换目标 host。

## 安装

从 [Releases](../../releases) 下载对应平台的压缩包（提供 Linux amd64 / arm64）：

```bash
tar xzf iptv-proxy_*_linux_amd64.tar.gz
./iptv-proxy --help
```

或自行编译（需 Go 1.26+）：

```bash
go build -o iptv-proxy ./
```

## 使用

```bash
./iptv-proxy -listen 0.0.0.0:6610 -upstream xx.xxx.xx.xx:6610
```

| 参数        | 默认值              | 说明                    |
| ----------- | ------------------- | ----------------------- |
| `-listen`   | `0.0.0.0:6610`      | 本地监听地址 `IP:PORT`  |
| `-upstream` | `xx.xxx.xx.xx:6610` | 上游 IPTV 源 `IP:PORT`  |

> 请把 `-upstream` 替换为你自己的 IPTV 源地址。

启动后，把播放器里的源地址指向本机监听地址即可，例如
`http://<本机IP>:6610/<原始频道路径>`。

## 备注

仓库内的 `p.conf` 是等价功能的 nginx 配置参考（含 30x 重定向处理），
按需取用。

## 关于

本项目由 [Claude Code](https://claude.com/claude-code) 协助开发。
