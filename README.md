# re-ASMR-Spider

[简体中文](docs/README.zh-CN.md) | [繁體中文](docs/README.zh-TW.md) | English | [日本語](docs/README.ja-JP.md)

---

A simple and efficient ASMR audio download tool.

## Features

- Multi-threaded download with resume support
- Smart retry mechanism for handling network errors
- Real-time progress display
- File size verification, automatically skip downloaded content
- Multi-language support
- Multi-platform support (Windows, Linux, macOS, BSD)
- Proxy configuration
- Custom filtering
- Command-line mode

## Screenshots
![Download Attempt](docs/images/asmr-spider-0.png)
![Retry Download](docs/images/asmr-spider-1.png)
![Custom Filtering](docs/images/asmr-spider-3.png)

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

### Interactive Mode (Default)

Run the executable. A default configuration file `config.json` will be created automatically.

```bash
./re-asmr-spider
```

1. Select option `1. Start Download`
2. Enter RJ number (e.g., `RJ373001`)
3. For multiple downloads, separate RJ numbers with spaces (e.g., `RJ373001 RJ123456 RJ789012`)

Downloads are saved to the `downloads/` directory.

### Command-Line Mode

For automation and scripting, use CLI mode to download directly without interactive menus:

```bash
# Download a single RJ
./re-asmr-spider -download RJ123456

# Download multiple RJs
./re-asmr-spider -download RJ123456,RJ789012,RJ345678

# Use custom account credentials
./re-asmr-spider -download RJ123456 -account user@example.com -password mypass

# Configure download parameters
./re-asmr-spider -download RJ123456 -max-task 5 -max-thread 16 -buffer-size 16

# Use proxy
./re-asmr-spider -download RJ123456 -proxy http://127.0.0.1:7890

# Use custom config file
./re-asmr-spider -config /path/to/config.json -download RJ123456

# Download priority
# When flac conflicts occur, download only flac
# For example, if files exist in flac, wav, mp3, and ogg formats with same name, only flac will be downloaded.
# But if flac doesn't exist, wav will be downloaded. If neither exists, mp3 and ogg will be downloaded
./re-asmr-spider -download RJ123456 -format-priority flac,wav

# Download flac with priority, and additionally download lrc subtitle files
./re-asmr-spider -download RJ123456 -format-priority flac,wav -include-formats lrc
```

**Command-Line Options:**

```
-download string      RJ numbers to download (comma-separated)
-config string        Path to config file (default: config.json)
-account string       ASMR.one account username (overrides config)
-password string      ASMR.one account password (overrides config)
-max-task int         Maximum concurrent download tasks
-max-thread int       Number of threads per file download
-max-retry int        Maximum retry attempts for failed downloads
-buffer-size int      Buffer size in MB (1-64, default: 8)
-proxy string         HTTP/HTTPS proxy (e.g., http://127.0.0.1:7890)
-format-priority string
                      Format priority for resolving conflicting files with the same name, reducing downloads, comma-separated (e.g., flac,wav,mp3)
                      When files with the same name exist in multiple formats, only download the highest priority format
                      (Interactive mode will prompt user for each choice)
-include-formats string
                      Additional file extensions to download, comma-separated (e.g., lrc,jpg)
                      After resolving conflicts above, download all files with specified extensions
                      Can be used with -format-priority
-version              Show version information
-help                 Show help message
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
  "proxy": "",
  "buffer_size_mb": 8
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
- `buffer_size_mb` - Download buffer size in MB (1-64, default: 8, optimized for VPS)

### Format Filtering

When downloading audio works, you may encounter files with the same name but different formats (e.g., `.wav`, `.flac`, `.mp3`). The program provides intelligent format filtering:

**Interactive Mode:**
- Automatically detects format conflicts
- Prompts user for format selection one by one
- Options include:
  - Download all formats
  - Select specific formats
  - Apply selection to all remaining similar files (batch mode)
  - Download all formats for all remaining files
- Format selection is saved and can be resumed if download is interrupted

### Resume Downloads

If a download is interrupted, the program will detect incomplete tasks on the next run and prompt you to continue. Your format selection preferences will be preserved.

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
