# re-ASMR-Spider

[简体中文](docs/README.zh-CN.md) | [繁體中文](docs/README.zh-TW.md) | English | [日本語](docs/README.ja-JP.md)

---

A simple and efficient ASMR audio download tool.

## Features

- Multi-threaded download with resume support
- Smart retry mechanism for handling network errors
- Real-time progress display
- File size verification, automatically skip downloaded content
- Multi-language support (20+ languages)
- Multi-platform support (Windows, Linux, macOS, BSD)
- Proxy configuration

## Screenshots
![Download Attempt](docs/images/asmr-spider-0.png)
![Retry Download](docs/images/asmr-spider-1.png)

## Installation

### Download Pre-built Binaries

Download the latest release from the [Releases](https://github.com/reuAC/re-asmr-spider/releases) page.

#### Platform Selection Guide

| Operating System | CPU Architecture | Example Devices/Scenarios | Download File |
|-----------------|------------------|---------------------------|---------------|
| **Windows** | | | |
| Windows 11/10/8/7 | 64-bit Intel/AMD | Most modern PCs and laptops | `Windows_x86_64.zip` |
| Windows 11/10/8/7 | 32-bit Intel/AMD | Older PCs (pre-2010) | `Windows_i386.zip` |
| Windows 11 | ARM64 | Surface Pro X, ARM-based laptops | `Windows_arm64.zip` |
| **Linux** | | | |
| Ubuntu, Debian, Fedora, CentOS | 64-bit Intel/AMD (x86_64) | Most servers and desktops | `Linux_x86_64.tar.gz` |
| Ubuntu, Debian, Fedora | 32-bit Intel/AMD (i386) | Old 32-bit systems | `Linux_i386.tar.gz` |
| Raspberry Pi OS, Ubuntu | ARM64 (aarch64) | Raspberry Pi 3/4/5 (64-bit OS), NVIDIA Jetson | `Linux_arm64.tar.gz` |
| Raspberry Pi OS | ARMv7 (32-bit) | Raspberry Pi 2/3/4 (32-bit OS) | `Linux_armv7.tar.gz` |
| Raspberry Pi OS | ARMv6 (32-bit) | Raspberry Pi 1, Pi Zero | `Linux_armv6.tar.gz` |
| OpenWrt, LEDE | MIPS (big-endian) | Some routers (MediaTek, Atheros) | `Linux_mips.tar.gz` |
| OpenWrt, LEDE | MIPS64 (big-endian) | High-end MIPS routers | `Linux_mips64.tar.gz` |
| OpenWrt, LEDE | MIPSle (little-endian) | Some routers (Broadcom, Realtek) | `Linux_mipsle.tar.gz` |
| OpenWrt, LEDE | MIPS64le (little-endian) | MIPS64 little-endian systems | `Linux_mips64le.tar.gz` |
| Debian, Fedora | RISC-V 64-bit | SiFive boards, StarFive VisionFive | `Linux_riscv64.tar.gz` |
| **macOS** | | | |
| macOS 10.13+ | Intel (x86_64) | MacBook Pro, iMac (2006-2020) | `Darwin_x86_64.tar.gz` |
| macOS 11+ | Apple Silicon (M1/M2/M3) | MacBook Air/Pro M1/M2/M3, Mac Mini M1/M2 | `Darwin_arm64.tar.gz` |
| **BSD** | | | |
| FreeBSD | 64-bit Intel/AMD | FreeBSD servers and desktops | `Freebsd_x86_64.tar.gz` |
| FreeBSD | 32-bit Intel/AMD | Older FreeBSD systems | `Freebsd_i386.tar.gz` |
| FreeBSD | ARM64 | FreeBSD on ARM boards | `Freebsd_arm64.tar.gz` |
| OpenBSD | 64-bit Intel/AMD | OpenBSD systems | `Openbsd_x86_64.tar.gz` |
| OpenBSD | 32-bit Intel/AMD | Older OpenBSD systems | `Openbsd_i386.tar.gz` |
| OpenBSD | ARM64 | OpenBSD on ARM boards | `Openbsd_arm64.tar.gz` |
| NetBSD | 64-bit Intel/AMD | NetBSD systems | `Netbsd_x86_64.tar.gz` |
| NetBSD | 32-bit Intel/AMD | Older NetBSD systems | `Netbsd_i386.tar.gz` |
| NetBSD | ARM64 | NetBSD on ARM boards | `Netbsd_arm64.tar.gz` |

#### How to Check Your System

**Windows:**
- Right-click "This PC" > Properties
- Check "System type" (64-bit or 32-bit, x64-based or ARM64-based)

**Linux:**
```bash
uname -m
# x86_64 -> 64-bit Intel/AMD
# i386/i686 -> 32-bit Intel/AMD
# aarch64 -> ARM64
# armv7l -> ARMv7
# armv6l -> ARMv6
# mips/mips64 -> MIPS
# riscv64 -> RISC-V 64-bit
```

**macOS:**
```bash
uname -m
# x86_64 -> Intel Mac
# arm64 -> Apple Silicon (M1/M2/M3)
```

Or: Apple menu > About This Mac > Chip (M1/M2/M3 = ARM64, Intel = x86_64)

### Build from Source

```bash
git clone https://github.com/reuAC/re-asmr-spider.git
cd re-asmr-spider
go build -o re-asmr-spider
```

## Usage

### First Run

Run the executable. A default configuration file `config.json` will be created automatically.

```bash
./re-asmr-spider
```

### Configuration

Edit `config.json` or use the built-in configuration menu (option 2):

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

**Configuration Options:**

- `account` - ASMR.one account username
- `password` - ASMR.one account password
- `max_task` - Maximum concurrent download tasks
- `max_thread` - Number of threads per file download
- `max_retry` - Maximum retry attempts for failed downloads
- `language` - Interface language (see supported languages below)
- `proxy` - HTTP/HTTPS proxy (e.g., `http://127.0.0.1:7890`, leave empty to disable)

### Download Audio

1. Run the program
2. Select option `1. Start Download`
3. Enter RJ number (e.g., `RJ373001`)
4. For multiple downloads, separate RJ numbers with spaces (e.g., `RJ373001 RJ123456 RJ789012`)

Downloads are saved to the `downloads/` directory.

### Resume Downloads

If a download is interrupted, the program will detect incomplete tasks on the next run and prompt you to continue.

## Supported Languages

- Chinese (Simplified) - 简体中文
- Chinese (Traditional) - 繁體中文
- English
- Japanese - 日本語

Switch language via the configuration menu (option 4) or by editing `config.json`.

## Data Source

All audio content is sourced from [asmr.one](https://asmr.one).

## License

MIT License

## Acknowledgements

Based on [go-asmr-spider](https://github.com/DiheChen/go-asmr-spider) by DiheChen.

## Disclaimer

This tool is for personal use only. Please respect content creators and support them through official channels.
