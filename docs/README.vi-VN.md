# re-ASMR-Spider

[简体中文](README.zh-CN.md) | [繁體中文](README.zh-TW.md) | [English](../README.md) | [日本語](README.ja-JP.md) | [Русский](README.ru-RU.md) | [Deutsch](README.de-DE.md) | [Español](README.es-ES.md) | [Français](README.fr-FR.md) | [Português](README.pt-PT.md) | [Indonesia](README.id-ID.md) | Tiếng Việt | [हिन्दी](README.hi-IN.md) | [বাংলা](README.bn-BD.md) | [తెలుగు](README.te-IN.md) | [Türkçe](README.tr-TR.md) | [اردو](README.ur-PK.md) | [Hausa](README.ha-NG.md) | [Esperanto](README.eo.md) | [喵喵语](README.cat.md) | [火星语](README.martian.md)

---

Công cụ tải xuống âm thanh ASMR đơn giản và hiệu quả.

## Tính năng

- Tải xuống đa luồng với hỗ trợ tiếp tục
- Cơ chế thử lại thông minh để xử lý lỗi mạng
- Hiển thị tiến trình theo thời gian thực
- Xác minh kích thước tệp, tự động bỏ qua nội dung đã tải xuống
- Hỗ trợ đa ngôn ngữ (20+ ngôn ngữ)
- Hỗ trợ đa nền tảng (Windows, Linux, macOS, BSD)
- Cấu hình proxy

## Ảnh chụp màn hình
![Thử tải xuống](./images/asmr-spider-0.png)
![Thử lại tải xuống](./images/asmr-spider-1.png)

## Cài đặt

### Tải xuống Tệp nhị phân Đã biên dịch

Tải xuống phiên bản mới nhất từ trang [Releases](https://github.com/reuAC/re-asmr-spider/releases).

#### Hướng dẫn Lựa chọn Nền tảng

| Hệ điều hành | Kiến trúc CPU | Ví dụ Thiết bị/Tình huống | Tệp Tải xuống |
|-----------------|------------------|---------------------------|---------------|
| **Windows** | | | |
| Windows 11/10/8/7 | 64-bit Intel/AMD | Hầu hết PC và laptop hiện đại | `Windows_x86_64.zip` |
| Windows 11/10/8/7 | 32-bit Intel/AMD | PC cũ (trước 2010) | `Windows_i386.zip` |
| Windows 11 | ARM64 | Surface Pro X, laptop dựa trên ARM | `Windows_arm64.zip` |
| **Linux** | | | |
| Ubuntu, Debian, Fedora, CentOS | 64-bit Intel/AMD (x86_64) | Hầu hết máy chủ và máy tính để bàn | `Linux_x86_64.tar.gz` |
| Ubuntu, Debian, Fedora | 32-bit Intel/AMD (i386) | Hệ thống 32-bit cũ | `Linux_i386.tar.gz` |
| Raspberry Pi OS, Ubuntu | ARM64 (aarch64) | Raspberry Pi 3/4/5 (OS 64-bit), NVIDIA Jetson | `Linux_arm64.tar.gz` |
| Raspberry Pi OS | ARMv7 (32-bit) | Raspberry Pi 2/3/4 (OS 32-bit) | `Linux_armv7.tar.gz` |
| Raspberry Pi OS | ARMv6 (32-bit) | Raspberry Pi 1, Pi Zero | `Linux_armv6.tar.gz` |
| OpenWrt, LEDE | MIPS (big-endian) | Một số router (MediaTek, Atheros) | `Linux_mips.tar.gz` |
| OpenWrt, LEDE | MIPS64 (big-endian) | Router MIPS cao cấp | `Linux_mips64.tar.gz` |
| OpenWrt, LEDE | MIPSle (little-endian) | Một số router (Broadcom, Realtek) | `Linux_mipsle.tar.gz` |
| OpenWrt, LEDE | MIPS64le (little-endian) | Hệ thống MIPS64 little-endian | `Linux_mips64le.tar.gz` |
| Debian, Fedora | RISC-V 64-bit | Bo mạch SiFive, StarFive VisionFive | `Linux_riscv64.tar.gz` |
| **macOS** | | | |
| macOS 10.13+ | Intel (x86_64) | MacBook Pro, iMac (2006-2020) | `Darwin_x86_64.tar.gz` |
| macOS 11+ | Apple Silicon (M1/M2/M3) | MacBook Air/Pro M1/M2/M3, Mac Mini M1/M2 | `Darwin_arm64.tar.gz` |
| **BSD** | | | |
| FreeBSD | 64-bit Intel/AMD | Máy chủ và máy tính để bàn FreeBSD | `Freebsd_x86_64.tar.gz` |
| FreeBSD | 32-bit Intel/AMD | Hệ thống FreeBSD cũ | `Freebsd_i386.tar.gz` |
| FreeBSD | ARM64 | FreeBSD trên bo mạch ARM | `Freebsd_arm64.tar.gz` |
| OpenBSD | 64-bit Intel/AMD | Hệ thống OpenBSD | `Openbsd_x86_64.tar.gz` |
| OpenBSD | 32-bit Intel/AMD | Hệ thống OpenBSD cũ | `Openbsd_i386.tar.gz` |
| OpenBSD | ARM64 | OpenBSD trên bo mạch ARM | `Openbsd_arm64.tar.gz` |
| NetBSD | 64-bit Intel/AMD | Hệ thống NetBSD | `Netbsd_x86_64.tar.gz` |
| NetBSD | 32-bit Intel/AMD | Hệ thống NetBSD cũ | `Netbsd_i386.tar.gz` |
| NetBSD | ARM64 | NetBSD trên bo mạch ARM | `Netbsd_arm64.tar.gz` |

#### Cách Kiểm tra Hệ thống của Bạn

**Windows:**
- Nhấp chuột phải vào "This PC" > Properties
- Kiểm tra "System type" (64-bit hoặc 32-bit, dựa trên x64 hoặc ARM64)

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

Hoặc: Menu Apple > About This Mac > Chip (M1/M2/M3 = ARM64, Intel = x86_64)

### Biên dịch từ Mã nguồn

```bash
git clone https://github.com/reuAC/re-asmr-spider.git
cd re-asmr-spider
go build -o re-asmr-spider
```

## Sử dụng

### Chạy lần đầu

Chạy tệp thực thi. Tệp cấu hình mặc định `config.json` sẽ được tạo tự động.

```bash
./re-asmr-spider
```

### Cấu hình

Chỉnh sửa `config.json` hoặc sử dụng menu cấu hình tích hợp (tùy chọn 2):

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

**Tùy chọn Cấu hình:**

- `account` - Tên người dùng tài khoản ASMR.one
- `password` - Mật khẩu tài khoản ASMR.one
- `max_task` - Số lượng tác vụ tải xuống đồng thời tối đa
- `max_thread` - Số luồng trên mỗi lần tải xuống tệp
- `max_retry` - Số lần thử lại tối đa cho các lần tải xuống thất bại
- `language` - Ngôn ngữ giao diện (xem các ngôn ngữ được hỗ trợ bên dưới)
- `proxy` - Proxy HTTP/HTTPS (ví dụ: `http://127.0.0.1:7890`, để trống để vô hiệu hóa)

### Tải xuống Âm thanh

1. Chạy chương trình
2. Chọn tùy chọn `1. Start Download`
3. Nhập số RJ (ví dụ: `RJ373001`)
4. Để tải xuống nhiều tệp, phân tách các số RJ bằng dấu cách (ví dụ: `RJ373001 RJ123456 RJ789012`)

Các tệp tải xuống được lưu vào thư mục `downloads/`.

### Tiếp tục Tải xuống

Nếu quá trình tải xuống bị gián đoạn, chương trình sẽ phát hiện các tác vụ chưa hoàn thành khi chạy lần tiếp theo và nhắc bạn tiếp tục.

## Ngôn ngữ được Hỗ trợ

- Chinese (Simplified & Traditional)
- English
- Japanese
- German
- Spanish
- French
- Portuguese
- Russian
- Indonesian
- Vietnamese
- Hindi
- Bengali
- Telugu
- Turkish
- Urdu
- Hausa
- Esperanto
- Cat Language
- Martian

Chuyển đổi ngôn ngữ thông qua menu cấu hình (tùy chọn 4) hoặc bằng cách chỉnh sửa `config.json`.

## Nguồn Dữ liệu

Tất cả nội dung âm thanh đều có nguồn gốc từ [asmr.one](https://asmr.one).

## Giấy phép

MIT License

## Lời cảm ơn

Dựa trên [go-asmr-spider](https://github.com/DiheChen/go-asmr-spider) bởi DiheChen.

## Tuyên bố miễn trừ trách nhiệm

Công cụ này chỉ dành cho mục đích sử dụng cá nhân. Vui lòng tôn trọng người sáng tạo nội dung và hỗ trợ họ thông qua các kênh chính thức.
