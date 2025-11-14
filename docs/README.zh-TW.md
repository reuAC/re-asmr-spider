# re-ASMR-Spider

[简体中文](README.zh-CN.md) | 繁體中文 | [English](../README.md) | [日本語](README.ja-JP.md) | [Русский](README.ru-RU.md) | [Deutsch](README.de-DE.md) | [Español](README.es-ES.md) | [Français](README.fr-FR.md) | [Português](README.pt-PT.md) | [Indonesia](README.id-ID.md) | [Tiếng Việt](README.vi-VN.md) | [हिन्दी](README.hi-IN.md) | [বাংলা](README.bn-BD.md) | [తెలుగు](README.te-IN.md) | [Türkçe](README.tr-TR.md) | [اردو](README.ur-PK.md) | [Hausa](README.ha-NG.md) | [Esperanto](README.eo.md) | [喵喵语](README.cat.md) | [火星语](README.martian.md)

---

一個簡單高效的 ASMR 音訊下載工具。

## 功能特色

- 多執行緒下載，支援斷點續傳
- 智慧重試機制，處理網路錯誤
- 即時進度顯示
- 檔案大小驗證，自動跳過已下載內容
- 多語言支援（20+ 種語言）
- 多平台支援（Windows、Linux、macOS、BSD）
- 代理伺服器設定

## 安裝

### 下載預編譯執行檔

從 [Releases](https://github.com/reuAC/re-asmr-spider/releases) 頁面下載最新版本。

#### 平台選擇指南

| 作業系統 | CPU 架構 | 範例裝置/場景 | 下載檔案 |
|-----------------|------------------|---------------------------|---------------|
| **Windows** | | | |
| Windows 11/10/8/7 | 64 位元 Intel/AMD | 大多數現代 PC 和筆記型電腦 | `Windows_x86_64.zip` |
| Windows 11/10/8/7 | 32 位元 Intel/AMD | 較舊的 PC（2010 年前） | `Windows_i386.zip` |
| Windows 11 | ARM64 | Surface Pro X、ARM 架構筆記型電腦 | `Windows_arm64.zip` |
| **Linux** | | | |
| Ubuntu、Debian、Fedora、CentOS | 64 位元 Intel/AMD (x86_64) | 大多數伺服器和桌面 | `Linux_x86_64.tar.gz` |
| Ubuntu、Debian、Fedora | 32 位元 Intel/AMD (i386) | 舊的 32 位元系統 | `Linux_i386.tar.gz` |
| Raspberry Pi OS、Ubuntu | ARM64 (aarch64) | Raspberry Pi 3/4/5（64 位元作業系統）、NVIDIA Jetson | `Linux_arm64.tar.gz` |
| Raspberry Pi OS | ARMv7（32 位元） | Raspberry Pi 2/3/4（32 位元作業系統） | `Linux_armv7.tar.gz` |
| Raspberry Pi OS | ARMv6（32 位元） | Raspberry Pi 1、Pi Zero | `Linux_armv6.tar.gz` |
| OpenWrt、LEDE | MIPS（大端序） | 部分路由器（MediaTek、Atheros） | `Linux_mips.tar.gz` |
| OpenWrt、LEDE | MIPS64（大端序） | 高階 MIPS 路由器 | `Linux_mips64.tar.gz` |
| OpenWrt、LEDE | MIPSle（小端序） | 部分路由器（Broadcom、Realtek） | `Linux_mipsle.tar.gz` |
| OpenWrt、LEDE | MIPS64le（小端序） | MIPS64 小端序系統 | `Linux_mips64le.tar.gz` |
| Debian、Fedora | RISC-V 64 位元 | SiFive 開發板、StarFive VisionFive | `Linux_riscv64.tar.gz` |
| **macOS** | | | |
| macOS 10.13+ | Intel (x86_64) | MacBook Pro、iMac（2006-2020） | `Darwin_x86_64.tar.gz` |
| macOS 11+ | Apple Silicon (M1/M2/M3) | MacBook Air/Pro M1/M2/M3、Mac Mini M1/M2 | `Darwin_arm64.tar.gz` |
| **BSD** | | | |
| FreeBSD | 64 位元 Intel/AMD | FreeBSD 伺服器和桌面 | `Freebsd_x86_64.tar.gz` |
| FreeBSD | 32 位元 Intel/AMD | 較舊的 FreeBSD 系統 | `Freebsd_i386.tar.gz` |
| FreeBSD | ARM64 | ARM 開發板上的 FreeBSD | `Freebsd_arm64.tar.gz` |
| OpenBSD | 64 位元 Intel/AMD | OpenBSD 系統 | `Openbsd_x86_64.tar.gz` |
| OpenBSD | 32 位元 Intel/AMD | 較舊的 OpenBSD 系統 | `Openbsd_i386.tar.gz` |
| OpenBSD | ARM64 | ARM 開發板上的 OpenBSD | `Openbsd_arm64.tar.gz` |
| NetBSD | 64 位元 Intel/AMD | NetBSD 系統 | `Netbsd_x86_64.tar.gz` |
| NetBSD | 32 位元 Intel/AMD | 較舊的 NetBSD 系統 | `Netbsd_i386.tar.gz` |
| NetBSD | ARM64 | ARM 開發板上的 NetBSD | `Netbsd_arm64.tar.gz` |

#### 如何檢查您的系統

**Windows：**
- 右鍵點選「本機」>「內容」
- 檢查「系統類型」（64 位元或 32 位元、x64 架構或 ARM64 架構）

**Linux：**
```bash
uname -m
# x86_64 -> 64 位元 Intel/AMD
# i386/i686 -> 32 位元 Intel/AMD
# aarch64 -> ARM64
# armv7l -> ARMv7
# armv6l -> ARMv6
# mips/mips64 -> MIPS
# riscv64 -> RISC-V 64 位元
```

**macOS：**
```bash
uname -m
# x86_64 -> Intel Mac
# arm64 -> Apple Silicon (M1/M2/M3)
```

或：Apple 選單 > 關於這台 Mac > 晶片（M1/M2/M3 = ARM64，Intel = x86_64）

### 從原始碼建置

```bash
git clone https://github.com/reuAC/re-asmr-spider.git
cd re-asmr-spider
go build -o re-asmr-spider
```

## 使用方法

### 首次執行

執行程式，將自動建立預設設定檔 `config.json`。

```bash
./re-asmr-spider
```

### 設定

編輯 `config.json` 或使用內建設定選單（選項 2）：

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

**設定選項：**

- `account` - ASMR.one 帳號使用者名稱
- `password` - ASMR.one 帳號密碼
- `max_task` - 最大同時下載任務數
- `max_thread` - 每個檔案下載的執行緒數
- `max_retry` - 下載失敗的最大重試次數
- `language` - 介面語言（參見下方支援的語言）
- `proxy` - HTTP/HTTPS 代理伺服器（例如：`http://127.0.0.1:7890`，留空則停用）

### 下載音訊

1. 執行程式
2. 選擇選項 `1. Start Download`
3. 輸入 RJ 編號（例如：`RJ373001`）
4. 若要下載多個檔案，請用空格分隔 RJ 編號（例如：`RJ373001 RJ123456 RJ789012`）

下載的檔案將儲存至 `downloads/` 目錄。

### 續傳下載

若下載中斷，程式將在下次執行時偵測未完成的任務，並提示您繼續下載。

## 支援的語言

- 中文（簡體和繁體）
- 英語
- 日語
- 德語
- 西班牙語
- 法語
- 葡萄牙語
- 俄語
- 印尼語
- 越南語
- 印地語
- 孟加拉語
- 泰盧固語
- 土耳其語
- 烏爾都語
- 豪薩語
- 世界語
- 貓語
- 火星語

可透過設定選單（選項 4）或編輯 `config.json` 切換語言。

## 資料來源

所有音訊內容均來自 [asmr.one](https://asmr.one)。

## 授權條款

MIT License

## 致謝

基於 DiheChen 的 [go-asmr-spider](https://github.com/DiheChen/go-asmr-spider) 專案。

## 免責聲明

此工具僅供個人使用。請尊重內容創作者，並透過官方管道支持他們。
