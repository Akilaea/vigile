# Vigile

> 轻量级服务器监控面板 · 守夜人

Vigile 是一个轻量的服务器与网站监控平台：单二进制、开箱即用，支持 Docker 容器指标、历史数据、告警通知、多用户与 OAuth 登录。

名字取自拉丁词根 **vigil**（守夜、警戒）—— 盯着你的机器，不喧哗。

[![License](https://img.shields.io/badge/license-MIT-9944ee)](LICENSE)
[![Go](https://img.shields.io/badge/Go-1.26%2B-00ADD8?logo=go)](https://go.dev)

---

## 为什么是 Vigile

市面上主流方案要么太重（Grafana + Prometheus 全家桶），要么 agent 内存持续增长、几百台机器后开始吃紧。Vigile 的底盘是在**同机、同条件实测**后选定的——内存表现是硬指标：

| 组件 | **Vigile** | Nezha v2.3.15 | 差异 |
|---|---|---|---|
| 面板 / 服务端 RSS | **28.5 MB**（平稳） | 41.0 MB（平稳） | 少 **30%** |
| 单机 Agent RSS | **16.6 MB**（零漂移） | 20.2 MB（持续上涨） | 少 **18%** 且不漂移 |
| Agent 二进制 | **13.8 MB** | 18.24 MB | 小 24% |
| 面板二进制 | **32.4 MB** | 80.9 MB | 小 **60%** |

**规模化换算（100 台被监控机器）**

- Vigile：`28.5 + 100 × 16.6` ≈ **1688 MB**
- Nezha：`41.0 + 100 × 20.2` ≈ **2061 MB**
- → 省下约 **373 MB**

> 上表为 Windows 空闲态实测（无真实主机接入、无历史数据积累）。Nezha 的 VictoriaMetrics TSDB 默认为关闭状态；其 agent 的内存上涨趋势仅观测 25 秒，未验证是否收敛。
>
> 测试方法：预热 40 秒后每 5 秒采样一次 RSS，共 6 次。

**Vigile 自身的内存优化仍在进行中**——底盘选对了，但还没有开始针对性的改造。

---

## 架构

Vigile 由两个组件构成：

| 组件 | 说明 |
|---|---|
| **Hub**（面板） | 基于 [PocketBase](https://pocketbase.io/) 的 Web 应用，提供仪表盘、告警与系统管理。单二进制，前端已内嵌。 |
| **Agent** | 部署在每台被监控机器上，采集指标并通过 WebSocket 上报 Hub。 |

---

## 快速开始

### 从源码构建

需要 **Go 1.26+** 与 **Node.js**（用于构建前端）。

```bash
git clone https://github.com/Akilaea/vigile.git
cd vigile

# 1. 构建前端（产物会被 Go 内嵌进二进制）
cd internal/site && npm install && npm run build && cd ../..

# 2. 构建面板
go build -o ./build/vigile -ldflags "-w -s" ./internal/cmd/hub

# 3. 构建 Agent（以 Linux 为例）
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
  go build -o ./build/vigile-agent -ldflags "-w -s" ./internal/cmd/agent

# 4. 启动面板
./build/vigile serve --http=0.0.0.0:8090 --dir=./vigile_data
```

启动后访问 `http://localhost:8090`，创建管理员账号，然后在面板里添加系统、复制生成的密钥部署 Agent。

> **Windows 构建 Agent 注意**：Windows 版 Agent 需要 **.NET SDK**（用于编译 LibreHardwareMonitor 以读取温度/风扇传感器）。Linux 版无此依赖。

---

## 支持的指标

- **CPU** —— 主机及 Docker / Podman 容器
- **内存** —— 主机与容器，含 swap 与 ZFS ARC
- **磁盘** —— 容量与 I/O，支持多分区 / 多设备
- **网络** —— 主机与容器流量
- **负载** —— 1 / 5 / 15 分钟
- **温度 / 风扇** —— 主机传感器
- **GPU** —— NVIDIA / AMD / Intel 的占用与功耗
- **电池** —— 主机及部分外设
- **容器** —— Docker / Podman 容器状态与指标
- **S.M.A.R.T.** —— 磁盘健康、eMMC 磨损、Linux mdraid 阵列状态
- **ZFS** —— 池容量、健康度、I/O、scrub 状态、按数据集的用量
- **systemd / 软件包更新 / Wi-Fi / Btrfs**

---

## 路线图

- [ ] 远程终端（WebSSH）—— 类似 Nezha 的 `IOStream`，直接在面板里操作被监控机器
- [ ] 定时任务下发
- [ ] 内存占用的针对性优化
- [ ] 中文文档

---

## 归属与许可

Vigile 是 **[Beszel](https://github.com/henrygd/beszel)**（作者 henrygd，MIT 许可）的衍生作品。

上游的 MIT 许可证与版权声明已**原样保留**于 [`LICENSE`](LICENSE)。本项目的改动范围、以及**有意保留未改**的上游引用（如 `beszel.dev` 更新端点、源码注释中的上游 issue 链接）详见 [`NOTICE.md`](NOTICE.md)。

本项目与上游作者无隶属关系，亦未获其背书；上游不对本项目提供支持。

本衍生作品同样以 **MIT 许可证**发布。
