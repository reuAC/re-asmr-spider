# re-ASMR-Spider

[简体中文](README.zh-CN.md) | 繁體中文 | [English](../README.md) | [日本語](README.ja-JP.md)

---

一個簡單高效的 ASMR 音訊下載工具。

## 功能特色

- 多執行緒下載，支援斷點續傳
- 智慧重試機制，自動處理網路錯誤
- 即時進度顯示
- 檔案大小驗證，自動跳過已下載的內容
- 多語言支援
- 多平台支援（Windows、Linux、macOS、BSD）
- 代理伺服器設定
- 自訂篩選
- 命令列模式

## 運行截圖
![嘗試下載作品](./images/asmr-spider-0.png)
![重試下載](./images/asmr-spider-1.png)
![自訂篩選](./images/asmr-spider-3.png)

## 安裝

### 下載預編譯執行檔

從 [Releases](https://github.com/reuAC/re-asmr-spider/releases) 頁面下載最新版本。

#### 平台選擇指南

| 作業系統 | CPU 架構 | 裝置範例 | 下載檔案 |
|---------|---------|---------|---------|
| **Windows** | | | |
| Windows 11/10/8/7 | 64 位元 Intel/AMD | 大多數現代 PC 和筆記型電腦 | `Windows_x86_64.zip` |
| Windows 11/10/8/7 | 32 位元 Intel/AMD | 較舊的 PC（2010 年前） | `Windows_i386.zip` |
| Windows 11 | ARM64 | Surface Pro X、ARM 筆記型電腦 | `Windows_arm64.zip` |
| **Linux** | | | |
| Ubuntu、Debian、Fedora、CentOS | 64 位元 Intel/AMD (x86_64) | 大多數伺服器和桌面 | `Linux_x86_64.tar.gz` |
| Ubuntu、Debian、Fedora | 32 位元 Intel/AMD (i386) | 舊的 32 位元系統 | `Linux_i386.tar.gz` |
| Raspberry Pi OS、Ubuntu | ARM64 (aarch64) | 樹莓派 3/4/5（64 位元系統）、NVIDIA Jetson | `Linux_arm64.tar.gz` |
| Raspberry Pi OS | ARMv7（32 位元） | 樹莓派 2/3/4（32 位元系統） | `Linux_armv7.tar.gz` |
| Raspberry Pi OS | ARMv6（32 位元） | 樹莓派 1、Pi Zero | `Linux_armv6.tar.gz` |
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
| FreeBSD | ARM64 | FreeBSD ARM 開發板 | `Freebsd_arm64.tar.gz` |
| OpenBSD | 64 位元 Intel/AMD | OpenBSD 系統 | `Openbsd_x86_64.tar.gz` |
| OpenBSD | 32 位元 Intel/AMD | 較舊的 OpenBSD 系統 | `Openbsd_i386.tar.gz` |
| OpenBSD | ARM64 | OpenBSD ARM 開發板 | `Openbsd_arm64.tar.gz` |
| NetBSD | 64 位元 Intel/AMD | NetBSD 系統 | `Netbsd_x86_64.tar.gz` |
| NetBSD | 32 位元 Intel/AMD | 較舊的 NetBSD 系統 | `Netbsd_i386.tar.gz` |
| NetBSD | ARM64 | NetBSD ARM 開發板 | `Netbsd_arm64.tar.gz` |

#### 如何查看系統資訊

**Windows：**
- 右鍵點選「本機」>「內容」
- 查看「系統類型」（64 位元或 32 位元，x64 或 ARM64）

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

### 互動模式

執行程式，將自動建立預設設定檔 `config.json`。

```bash
./re-asmr-spider
```

1. 選擇選項 `1. 開始下載`
2. 輸入 RJ 編號（例如：`RJ373001`）
3. 多個下載時，用空格分隔 RJ 編號（例如：`RJ373001 RJ123456 RJ789012`）

下載的檔案儲存在 `downloads/` 目錄。

### 命令列模式

用於自動化和腳本呼叫，無需互動式選單即可直接下載：

```bash
# 下載單一 RJ 編號
./re-asmr-spider -download RJ123456

# 下載多個 RJ 編號
./re-asmr-spider -download RJ123456,RJ789012,RJ345678

# 使用自訂帳號密碼
./re-asmr-spider -download RJ123456 -account user@example.com -password mypass

# 設定下載參數
./re-asmr-spider -download RJ123456 -max-task 5 -max-thread 16 -buffer-size 16

# 使用代理伺服器
./re-asmr-spider -download RJ123456 -proxy http://127.0.0.1:7890

# 使用自訂設定檔
./re-asmr-spider -config /path/to/config.json -download RJ123456

# 下載優先順序
# 當出現有關 flac 的衝突時，僅下載 flac
# 例如同時存在 flac wav mp3 ogg 四個格式的同名檔案，則僅下載 flac。
# 但若沒有 flac，則會下載 wav。都沒有時則會下載 mp3 與 ogg
./re-asmr-spider -download RJ123456 -format-priority flac,wav

# 優先下載 flac，並額外下載 lrc 字幕檔案
./re-asmr-spider -download RJ123456 -format-priority flac,wav -include-formats lrc
```

**命令列選項：**

```
-download string      要下載的 RJ 編號（逗號分隔）
-config string        設定檔路徑（預設：config.json）
-account string       ASMR.one 帳號使用者名稱（覆寫設定檔）
-password string      ASMR.one 帳號密碼（覆寫設定檔）
-max-task int         最大並行下載任務數
-max-thread int       單檔案下載執行緒數
-max-retry int        下載失敗的最大重試次數
-buffer-size int      緩衝區大小（MB），範圍 1-64，預設 8
-proxy string         HTTP/HTTPS 代理伺服器（例如：http://127.0.0.1:7890）
-format-priority string
                      格式優先順序，用於解決衝突的同名檔案，減少下載量，逗號分隔（例如：flac,wav,mp3）
                      當同名檔案存在多種格式時，只下載優先順序最高的格式
                      （互動模式會逐個詢問使用者選擇）
-include-formats string
                      額外下載的副檔名，逗號分隔（例如：lrc,jpg）
                      在上述衝突解決後，額外下載所有指定副檔名的檔案
                      可與 -format-priority 配合使用
-version              顯示版本資訊
-help                 顯示說明訊息
```

### 設定

編輯 `config.json` 或使用內建設定選單（選項 2）：

```json
{
  "account": "你的使用者名稱",
  "password": "你的密碼",
  "max_task": 3,
  "max_thread": 8,
  "max_retry": 3,
  "language": "zh-CN",
  "proxy": "",
  "buffer_size_mb": 8
}
```

**設定選項：**

- `account` - ASMR.one 帳號使用者名稱
- `password` - ASMR.one 帳號密碼
- `max_task` - 最大並行下載任務數
- `max_thread` - 單檔案下載執行緒數
- `max_retry` - 下載失敗的最大重試次數
- `language` - 介面語言（參見下方支援的語言列表）
- `proxy` - HTTP/HTTPS 代理伺服器（例如：`http://127.0.0.1:7890`，留空則停用）
- `buffer_size_mb` - 下載緩衝區大小（MB），範圍 1-64，預設 8，針對 VPS 最佳化

### 格式篩選

下載音訊作品時，可能會遇到同名不同格式的檔案（如 `.wav`、`.flac`、`.mp3`）。程式提供智慧格式篩選功能：

**互動模式：**
- 自動偵測格式衝突
- 逐個提示使用者選擇要下載的格式
- 選項包括：
  - 下載所有格式
  - 選擇特定格式
  - 將選擇套用到所有剩餘的同類檔案（批次模式）
  - 為所有剩餘檔案下載所有格式
- 格式選擇會被儲存，下載中斷後可繼續使用

### 恢復下載

如果下載被中斷，程式會在下次執行時偵測到未完成的任務並提示您繼續。您的格式選擇偏好將被保留。

## 支援的語言

- 簡體中文
- 繁體中文
- English
- 日本語

透過設定選單（選項 4）或編輯 `config.json` 切換語言。

## 資料來源

所有音訊內容來源於 [asmr.one](https://asmr.one)。

## 授權條款

MIT License

## 致謝

基於 DiheChen 的 [go-asmr-spider](https://github.com/DiheChen/go-asmr-spider) 開發。

## 免責聲明

此工具僅供個人使用。請尊重內容創作者，並透過官方管道支持他們。
