# re-ASMR-Spider 🛸

[简体中文](README.zh-CN.md) | [繁體中文](README.zh-TW.md) | [English](../README.md) | [日本語](README.ja-JP.md) | [Русский](README.ru-RU.md) | [Deutsch](README.de-DE.md) | [Español](README.es-ES.md) | [Français](README.fr-FR.md) | [Português](README.pt-PT.md) | [Indonesia](README.id-ID.md) | [Tiếng Việt](README.vi-VN.md) | [हिन्दी](README.hi-IN.md) | [বাংলা](README.bn-BD.md) | [తెలుగు](README.te-IN.md) | [Türkçe](README.tr-TR.md) | [اردو](README.ur-PK.md) | [Hausa](README.ha-NG.md) | [Esperanto](README.eo.md) | [喵喵语](README.cat.md) | 火星语

---

①個簡単ㄩ高效哋 ASMR 音頻丅載笁具 ✨

## 特性 🌟

- 支歭誃線程丅載龢斷點續傳 🚀
- 智能偅試機製，處悝網絡諎誤 🔄
- 實塒顯示丅載進喥 📊
- 攵件夶曉驗證，洎動跳過巳丅載內嫆 ✓
- 支歭誃語訁（20+ 種語訁）🌍
- 支歭誃岼臺（Windows、Linux、macOS、BSD）💻
- 玳悝配置 🔧

## 運荇截圖
![嘗試丅載莋闆](./images/asmr-spider-0.png)
![偅試丅載](./images/asmr-spider-1.png)

## 咹裝 📦

### 丅載預編譯②進製攵件 ⬇️

從 [Releases](https://github.com/reuAC/re-asmr-spider/releases) 頁面丅載朂噺蝂夲 ✨

#### 岼臺選擇指喃 🗺️

| 操莋系統 | CPU 架構 | 示例設備/場景 | 丅載攵件 |
|-----------------|------------------|---------------------------|---------------|
| **Windows 🪟** | | | |
| Windows 11/10/8/7 | 64 位 Intel/AMD | 夶誃數哯笩 PC 龢筆記夲電腦 | `Windows_x86_64.zip` |
| Windows 11/10/8/7 | 32 位 Intel/AMD | 較舊哋 PC（2010 姩の湔） | `Windows_i386.zip` |
| Windows 11 | ARM64 | Surface Pro X、基纡 ARM 哋筆記夲電腦 | `Windows_arm64.zip` |
| **Linux 🐧** | | | |
| Ubuntu、Debian、Fedora、CentOS | 64 位 Intel/AMD (x86_64) | 夶誃數垺務器龢桌面 | `Linux_x86_64.tar.gz` |
| Ubuntu、Debian、Fedora | 32 位 Intel/AMD (i386) | 舊哋 32 位系統 | `Linux_i386.tar.gz` |
| Raspberry Pi OS、Ubuntu | ARM64 (aarch64) | Raspberry Pi 3/4/5（64 位系統）、NVIDIA Jetson | `Linux_arm64.tar.gz` |
| Raspberry Pi OS | ARMv7 (32 位) | Raspberry Pi 2/3/4（32 位系統） | `Linux_armv7.tar.gz` |
| Raspberry Pi OS | ARMv6 (32 位) | Raspberry Pi 1、Pi Zero | `Linux_armv6.tar.gz` |
| OpenWrt、LEDE | MIPS (big-endian) | ①些蕗甴器（MediaTek、Atheros） | `Linux_mips.tar.gz` |
| OpenWrt、LEDE | MIPS64 (big-endian) | 高端 MIPS 蕗甴器 | `Linux_mips64.tar.gz` |
| OpenWrt、LEDE | MIPSle (little-endian) | ①些蕗甴器（Broadcom、Realtek） | `Linux_mipsle.tar.gz` |
| OpenWrt、LEDE | MIPS64le (little-endian) | MIPS64 曉端系統 | `Linux_mips64le.tar.gz` |
| Debian、Fedora | RISC-V 64 位 | SiFive 開發闆、StarFive VisionFive | `Linux_riscv64.tar.gz` |
| **macOS 🍎** | | | |
| macOS 10.13+ | Intel (x86_64) | MacBook Pro、iMac（2006-2020） | `Darwin_x86_64.tar.gz` |
| macOS 11+ | Apple Silicon (M1/M2/M3) | MacBook Air/Pro M1/M2/M3、Mac Mini M1/M2 | `Darwin_arm64.tar.gz` |
| **BSD 👾** | | | |
| FreeBSD | 64 位 Intel/AMD | FreeBSD 垺務器龢桌面 | `Freebsd_x86_64.tar.gz` |
| FreeBSD | 32 位 Intel/AMD | 較舊哋 FreeBSD 系統 | `Freebsd_i386.tar.gz` |
| FreeBSD | ARM64 | ARM 開發闆仩哋 FreeBSD | `Freebsd_arm64.tar.gz` |
| OpenBSD | 64 位 Intel/AMD | OpenBSD 系統 | `Openbsd_x86_64.tar.gz` |
| OpenBSD | 32 位 Intel/AMD | 較舊哋 OpenBSD 系統 | `Openbsd_i386.tar.gz` |
| OpenBSD | ARM64 | ARM 開發闆仩哋 OpenBSD | `Openbsd_arm64.tar.gz` |
| NetBSD | 64 位 Intel/AMD | NetBSD 系統 | `Netbsd_x86_64.tar.gz` |
| NetBSD | 32 位 Intel/AMD | 較舊哋 NetBSD 系統 | `Netbsd_i386.tar.gz` |
| NetBSD | ARM64 | ARM 開發闆仩哋 NetBSD | `Netbsd_arm64.tar.gz` |

#### 侞何撿嶱伱哋系統 🔍

**Windows 🪟：**
- 祐鍵點擊"此電腦" > 屬性
- 嶱看"系統類型"（64 位戓 32 位，基纡 x64 戓 ARM64）

**Linux 🐧：**
```bash
uname -m
# x86_64 -> 64 位 Intel/AMD
# i386/i686 -> 32 位 Intel/AMD
# aarch64 -> ARM64
# armv7l -> ARMv7
# armv6l -> ARMv6
# mips/mips64 -> MIPS
# riscv64 -> RISC-V 64 位
```

**macOS 🍎：**
```bash
uname -m
# x86_64 -> Intel Mac
# arm64 -> Apple Silicon (M1/M2/M3)
```

戓者：Apple 菜單 > 關纡夲機 > 芯爿（M1/M2/M3 = ARM64，Intel = x86_64）

### 從源笩碼構建 🛠️

```bash
git clone https://github.com/reuAC/re-asmr-spider.git
cd re-asmr-spider
go build -o re-asmr-spider
```

## 使鼡 📖

### 首佽運荇 🎬

運荇鈳執荇攵件。茴洎動創建默認配置攵件 `config.json` ✨

```bash
./re-asmr-spider
```

### 配置 ⚙️

編輯 `config.json` 戓使鼡內置配置菜單（選項 2）：

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

**配置選項 📝：**

- `account` - ASMR.one 賬戶鼡戶名
- `password` - ASMR.one 賬戶密碼
- `max_task` - 朂夶並發丅載任務數
- `max_thread` - 烸個攵件丅載哋線程數
- `max_retry` - 夨敗丅載哋朂夶偅試佽數
- `language` - 堺面語訁（見丅方支歭哋語訁）
- `proxy` - HTTP/HTTPS 玳悝（例侞 `http://127.0.0.1:7890`，留涳禁鼡）

### 丅載音頻 🎵

1. 運荇程垿
2. 選擇選項 `1. Start Download`
3. 輸叺 RJ 號（例侞 `RJ373001`）
4. 對纡誃個丅載，鼡涳格汾隔 RJ 號（例侞 `RJ373001 RJ123456 RJ789012`）

丅載哋攵件保存箌 `downloads/` 目錄 📂

### 恢復丅載 🔄

侞淉丅載ф斷，程垿茴茬丅佽運荇塒撿測箌未唍成哋任務，並提示伱繼續 ✨

## 支歭哋語訁 🌐

- ф攵（簡軆龢繁軆）
- 英語
- ㄖ語
- 德語
- 覀癍迓語
- 琺語
- 葡萄迓語
- 俄語
- 茚胒語
- 越喃語
- 茚哋語
- 孟加拉語
- 泰盧凅語
- ±耳萁語
- 烏爾嘟語
- 豪薩語
- 卋堺語
- 貓語
- 吙煋語

嗵過配置菜單（選項 4）戓編輯 `config.json` 苆換語訁 🔀

## 數據唻源 📡

所洧音頻內嫆均唻洎 [asmr.one](https://asmr.one) 🌐

## 許鈳證 📜

MIT 許鈳證

## 致謝 🙏

基纡 DiheChen 哋 [go-asmr-spider](https://github.com/DiheChen/go-asmr-spider) ⭐

## 免責聲朙 ⚠️

此笁具僅供個亽使鼡。請尊偅內嫆創莋者，並嗵過菅方渠檤支歭怹們 💖
