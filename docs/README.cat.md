# re-ASMR-Spider 喵

[简体中文](README.zh-CN.md) | [繁體中文](README.zh-TW.md) | [English](../README.md) | [日本語](README.ja-JP.md) | [Русский](README.ru-RU.md) | [Deutsch](README.de-DE.md) | [Español](README.es-ES.md) | [Français](README.fr-FR.md) | [Português](README.pt-PT.md) | [Indonesia](README.id-ID.md) | [Tiếng Việt](README.vi-VN.md) | [हिन्दी](README.hi-IN.md) | [বাংলা](README.bn-BD.md) | [తెలుగు](README.te-IN.md) | [Türkçe](README.tr-TR.md) | [اردو](README.ur-PK.md) | [Hausa](README.ha-NG.md) | [Esperanto](README.eo.md) | 喵喵语 | [火星语](README.martian.md)

---

一个简单又高效的 ASMR 音频下载工具喵~

## 特性喵

- 支持多线程下载和断点续传喵
- 智能重试机制，处理网络错误喵
- 实时显示下载进度喵
- 文件大小验证，自动跳过已下载内容喵
- 支持多语言（20+ 种语言）喵
- 支持多平台（Windows、Linux、macOS、BSD）喵
- 代理配置喵

## 运行截图喵
![尝试下载作品喵](./images/asmr-spider-0.png)
![重试下载喵](./images/asmr-spider-1.png)

## 安装喵

### 下载预编译二进制文件喵

从 [Releases](https://github.com/reuAC/re-asmr-spider/releases) 页面下载最新版本喵~

#### 平台选择指南喵

| 操作系统 | CPU 架构 | 示例设备/场景 | 下载文件 |
|-----------------|------------------|---------------------------|---------------|
| **Windows 喵** | | | |
| Windows 11/10/8/7 | 64 位 Intel/AMD | 大多数现代 PC 和笔记本电脑喵 | `Windows_x86_64.zip` |
| Windows 11/10/8/7 | 32 位 Intel/AMD | 较旧的 PC（2010 年之前）喵 | `Windows_i386.zip` |
| Windows 11 | ARM64 | Surface Pro X、基于 ARM 的笔记本电脑喵 | `Windows_arm64.zip` |
| **Linux 喵** | | | |
| Ubuntu、Debian、Fedora、CentOS | 64 位 Intel/AMD (x86_64) | 大多数服务器和桌面喵 | `Linux_x86_64.tar.gz` |
| Ubuntu、Debian、Fedora | 32 位 Intel/AMD (i386) | 旧的 32 位系统喵 | `Linux_i386.tar.gz` |
| Raspberry Pi OS、Ubuntu | ARM64 (aarch64) | Raspberry Pi 3/4/5（64 位系统）、NVIDIA Jetson 喵 | `Linux_arm64.tar.gz` |
| Raspberry Pi OS | ARMv7 (32 位) | Raspberry Pi 2/3/4（32 位系统）喵 | `Linux_armv7.tar.gz` |
| Raspberry Pi OS | ARMv6 (32 位) | Raspberry Pi 1、Pi Zero 喵 | `Linux_armv6.tar.gz` |
| OpenWrt、LEDE | MIPS (big-endian) | 一些路由器（MediaTek、Atheros）喵 | `Linux_mips.tar.gz` |
| OpenWrt、LEDE | MIPS64 (big-endian) | 高端 MIPS 路由器喵 | `Linux_mips64.tar.gz` |
| OpenWrt、LEDE | MIPSle (little-endian) | 一些路由器（Broadcom、Realtek）喵 | `Linux_mipsle.tar.gz` |
| OpenWrt、LEDE | MIPS64le (little-endian) | MIPS64 小端系统喵 | `Linux_mips64le.tar.gz` |
| Debian、Fedora | RISC-V 64 位 | SiFive 开发板、StarFive VisionFive 喵 | `Linux_riscv64.tar.gz` |
| **macOS 喵** | | | |
| macOS 10.13+ | Intel (x86_64) | MacBook Pro、iMac（2006-2020）喵 | `Darwin_x86_64.tar.gz` |
| macOS 11+ | Apple Silicon (M1/M2/M3) | MacBook Air/Pro M1/M2/M3、Mac Mini M1/M2 喵 | `Darwin_arm64.tar.gz` |
| **BSD 喵** | | | |
| FreeBSD | 64 位 Intel/AMD | FreeBSD 服务器和桌面喵 | `Freebsd_x86_64.tar.gz` |
| FreeBSD | 32 位 Intel/AMD | 较旧的 FreeBSD 系统喵 | `Freebsd_i386.tar.gz` |
| FreeBSD | ARM64 | ARM 开发板上的 FreeBSD 喵 | `Freebsd_arm64.tar.gz` |
| OpenBSD | 64 位 Intel/AMD | OpenBSD 系统喵 | `Openbsd_x86_64.tar.gz` |
| OpenBSD | 32 位 Intel/AMD | 较旧的 OpenBSD 系统喵 | `Openbsd_i386.tar.gz` |
| OpenBSD | ARM64 | ARM 开发板上的 OpenBSD 喵 | `Openbsd_arm64.tar.gz` |
| NetBSD | 64 位 Intel/AMD | NetBSD 系统喵 | `Netbsd_x86_64.tar.gz` |
| NetBSD | 32 位 Intel/AMD | 较旧的 NetBSD 系统喵 | `Netbsd_i386.tar.gz` |
| NetBSD | ARM64 | ARM 开发板上的 NetBSD 喵 | `Netbsd_arm64.tar.gz` |

#### 如何检查你的系统喵

**Windows 喵：**
- 右键点击"此电脑" > 属性喵
- 查看"系统类型"（64 位或 32 位，基于 x64 或 ARM64）喵

**Linux 喵：**
```bash
uname -m
# x86_64 -> 64 位 Intel/AMD 喵
# i386/i686 -> 32 位 Intel/AMD 喵
# aarch64 -> ARM64 喵
# armv7l -> ARMv7 喵
# armv6l -> ARMv6 喵
# mips/mips64 -> MIPS 喵
# riscv64 -> RISC-V 64 位喵
```

**macOS 喵：**
```bash
uname -m
# x86_64 -> Intel Mac 喵
# arm64 -> Apple Silicon (M1/M2/M3) 喵
```

或者：Apple 菜单 > 关于本机 > 芯片（M1/M2/M3 = ARM64，Intel = x86_64）喵

### 从源代码构建喵

```bash
git clone https://github.com/reuAC/re-asmr-spider.git
cd re-asmr-spider
go build -o re-asmr-spider
```

## 使用喵

### 首次运行喵

运行可执行文件喵。会自动创建默认配置文件 `config.json` 喵~

```bash
./re-asmr-spider
```

### 配置喵

编辑 `config.json` 或使用内置配置菜单（选项 2）喵：

```json
{
  "account": "your_username",
  "password": "your_password",
  "max_task": 3,
  "max_thread": 8,
  "max_retry": 3,
  "language": "zh-CN",
  "proxy": ""
}
```

**配置选项喵：**

- `account` - ASMR.one 账户用户名喵
- `password` - ASMR.one 账户密码喵
- `max_task` - 最大并发下载任务数喵
- `max_thread` - 每个文件下载的线程数喵
- `max_retry` - 失败下载的最大重试次数喵
- `language` - 界面语言（见下方支持的语言）喵
- `proxy` - HTTP/HTTPS 代理（例如 `http://127.0.0.1:7890`，留空禁用）喵

### 下载音频喵

1. 运行程序喵
2. 选择选项 `1. Start Download` 喵
3. 输入 RJ 号（例如 `RJ373001`）喵
4. 对于多个下载，用空格分隔 RJ 号（例如 `RJ373001 RJ123456 RJ789012`）喵

下载的文件保存到 `downloads/` 目录喵~

### 恢复下载喵

如果下载中断，程序会在下次运行时检测到未完成的任务，并提示你继续喵~

## 支持的语言喵

- 中文（简体和繁体）喵
- 英语喵
- 日语喵
- 德语喵
- 西班牙语喵
- 法语喵
- 葡萄牙语喵
- 俄语喵
- 印尼语喵
- 越南语喵
- 印地语喵
- 孟加拉语喵
- 泰卢固语喵
- 土耳其语喵
- 乌尔都语喵
- 豪萨语喵
- 世界语喵
- 猫语喵
- 火星语喵

通过配置菜单（选项 4）或编辑 `config.json` 切换语言喵~

## 数据来源喵

所有音频内容均来自 [asmr.one](https://asmr.one) 喵~

## 许可证喵

MIT 许可证喵

## 致谢喵

基于 DiheChen 的 [go-asmr-spider](https://github.com/DiheChen/go-asmr-spider) 喵~

## 免责声明喵

此工具仅供个人使用喵。请尊重内容创作者，并通过官方渠道支持他们喵~
