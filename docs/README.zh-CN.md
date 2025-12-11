# re-ASMR-Spider

简体中文 | [繁體中文](README.zh-TW.md) | [English](../README.md) | [日本語](README.ja-JP.md)

---

一个简单高效的 ASMR 音声下载工具。

## 功能特性

- 多线程下载，支持断点续传
- 智能重试机制，自动处理网络错误
- 实时进度显示
- 文件大小校验，自动跳过已下载的内容
- 多语言支持（20+种语言）
- 多平台支持（Windows、Linux、macOS、BSD）
- 代理配置

## 运行截图
![尝试下载作品](./images/asmr-spider-0.png)
![重试下载](./images/asmr-spider-1.png)

## 安装

### 下载预编译二进制文件

从 [Releases](https://github.com/reuAC/re-asmr-spider/releases) 页面下载最新版本。

#### 平台选择指南

| 操作系统 | CPU 架构 | 设备示例 | 下载文件 |
|---------|---------|---------|---------|
| **Windows** | | | |
| Windows 11/10/8/7 | 64位 Intel/AMD | 大多数现代 PC 和笔记本 | `Windows_x86_64.zip` |
| Windows 11/10/8/7 | 32位 Intel/AMD | 老旧 PC（2010年前） | `Windows_i386.zip` |
| Windows 11 | ARM64 | Surface Pro X、ARM 笔记本 | `Windows_arm64.zip` |
| **Linux** | | | |
| Ubuntu、Debian、Fedora、CentOS | 64位 Intel/AMD (x86_64) | 大多数服务器和桌面 | `Linux_x86_64.tar.gz` |
| Ubuntu、Debian、Fedora | 32位 Intel/AMD (i386) | 老旧32位系统 | `Linux_i386.tar.gz` |
| Raspberry Pi OS、Ubuntu | ARM64 (aarch64) | 树莓派 3/4/5（64位系统）、NVIDIA Jetson | `Linux_arm64.tar.gz` |
| Raspberry Pi OS | ARMv7 (32位) | 树莓派 2/3/4（32位系统） | `Linux_armv7.tar.gz` |
| Raspberry Pi OS | ARMv6 (32位) | 树莓派 1、Pi Zero | `Linux_armv6.tar.gz` |
| OpenWrt、LEDE | MIPS（大端） | 部分路由器（MediaTek、Atheros） | `Linux_mips.tar.gz` |
| OpenWrt、LEDE | MIPS64（大端） | 高端 MIPS 路由器 | `Linux_mips64.tar.gz` |
| OpenWrt、LEDE | MIPSle（小端） | 部分路由器（Broadcom、Realtek） | `Linux_mipsle.tar.gz` |
| OpenWrt、LEDE | MIPS64le（小端） | MIPS64 小端系统 | `Linux_mips64le.tar.gz` |
| Debian、Fedora | RISC-V 64位 | SiFive 开发板、StarFive VisionFive | `Linux_riscv64.tar.gz` |
| **macOS** | | | |
| macOS 10.13+ | Intel (x86_64) | MacBook Pro、iMac（2006-2020） | `Darwin_x86_64.tar.gz` |
| macOS 11+ | Apple Silicon (M1/M2/M3) | MacBook Air/Pro M1/M2/M3、Mac Mini M1/M2 | `Darwin_arm64.tar.gz` |
| **BSD** | | | |
| FreeBSD | 64位 Intel/AMD | FreeBSD 服务器和桌面 | `Freebsd_x86_64.tar.gz` |
| FreeBSD | 32位 Intel/AMD | 老旧 FreeBSD 系统 | `Freebsd_i386.tar.gz` |
| FreeBSD | ARM64 | FreeBSD ARM 开发板 | `Freebsd_arm64.tar.gz` |
| OpenBSD | 64位 Intel/AMD | OpenBSD 系统 | `Openbsd_x86_64.tar.gz` |
| OpenBSD | 32位 Intel/AMD | 老旧 OpenBSD 系统 | `Openbsd_i386.tar.gz` |
| OpenBSD | ARM64 | OpenBSD ARM 开发板 | `Openbsd_arm64.tar.gz` |
| NetBSD | 64位 Intel/AMD | NetBSD 系统 | `Netbsd_x86_64.tar.gz` |
| NetBSD | 32位 Intel/AMD | 老旧 NetBSD 系统 | `Netbsd_i386.tar.gz` |
| NetBSD | ARM64 | NetBSD ARM 开发板 | `Netbsd_arm64.tar.gz` |

#### 如何查看系统信息

**Windows：**
- 右键"此电脑" > 属性
- 查看"系统类型"（64位或32位，x64或ARM64）

**Linux：**
```bash
uname -m
# x86_64 -> 64位 Intel/AMD
# i386/i686 -> 32位 Intel/AMD
# aarch64 -> ARM64
# armv7l -> ARMv7
# armv6l -> ARMv6
# mips/mips64 -> MIPS
# riscv64 -> RISC-V 64位
```

**macOS：**
```bash
uname -m
# x86_64 -> Intel Mac
# arm64 -> Apple Silicon (M1/M2/M3)
```

或：Apple 菜单 > 关于本机 > 芯片（M1/M2/M3 = ARM64，Intel = x86_64）

### 从源码构建

```bash
git clone https://github.com/reuAC/re-asmr-spider.git
cd re-asmr-spider
go build -o re-asmr-spider
```

## 使用方法

### 首次运行

运行可执行文件，将自动创建默认配置文件 `config.json`。

```bash
./re-asmr-spider
```

### 配置

编辑 `config.json` 或使用内置配置菜单（选项2）：

```json
{
  "account": "你的用户名",
  "password": "你的密码",
  "max_task": 3,
  "max_thread": 8,
  "max_retry": 3,
  "language": "zh-CN",
  "proxy": ""
}
```

**配置选项：**

- `account` - ASMR.one 账号用户名
- `password` - ASMR.one 账号密码
- `max_task` - 最大并发下载任务数
- `max_thread` - 单文件下载线程数
- `max_retry` - 失败下载的最大重试次数
- `language` - 界面语言（见下方支持的语言列表）
- `proxy` - HTTP/HTTPS 代理（例如：`http://127.0.0.1:7890`，留空则禁用）

### 下载音声

1. 运行程序
2. 选择选项 `1. 开始下载`
3. 输入 RJ 号（例如：`RJ373001`）
4. 多个下载时，用空格分隔 RJ 号（例如：`RJ373001 RJ123456 RJ789012`）

下载的文件保存在 `downloads/` 目录。

### 恢复下载

如果下载被中断，程序会在下次运行时检测到未完成的任务并提示您继续。

## 支持的语言

- 简体中文和繁体中文
- English
- 日本語
- Deutsch
- Español
- Français
- Português
- Русский
- Bahasa Indonesia
- Tiếng Việt
- हिन्दी
- বাংলা
- తెలుగు
- Türkçe
- اردو
- Hausa
- Esperanto
- 喵喵语
- 火星语

通过配置菜单（选项4）或编辑 `config.json` 切换语言。

## 数据来源

所有音声内容来源于 [asmr.one](https://asmr.one)。

## 许可证

MIT License

## 鸣谢

基于 DiheChen 的 [go-asmr-spider](https://github.com/DiheChen/go-asmr-spider) 开发。

## 免责声明

本工具仅供个人使用。请尊重内容创作者并通过官方渠道支持他们。
