# re-ASMR-Spider

[简体中文](README.zh-CN.md) | [繁體中文](README.zh-TW.md) | [English](../README.md) | [日本語](README.ja-JP.md) | [Русский](README.ru-RU.md) | [Deutsch](README.de-DE.md) | [Español](README.es-ES.md) | [Français](README.fr-FR.md) | [Português](README.pt-PT.md) | Indonesia | [Tiếng Việt](README.vi-VN.md) | [हिन्दी](README.hi-IN.md) | [বাংলা](README.bn-BD.md) | [తెలుగు](README.te-IN.md) | [Türkçe](README.tr-TR.md) | [اردو](README.ur-PK.md) | [Hausa](README.ha-NG.md) | [Esperanto](README.eo.md) | [喵喵语](README.cat.md) | [火星语](README.martian.md)

---

Alat unduh audio ASMR yang sederhana dan efisien.

## Fitur

- Unduhan multi-thread dengan dukungan melanjutkan
- Mekanisme percobaan ulang cerdas untuk menangani kesalahan jaringan
- Tampilan progress real-time
- Verifikasi ukuran file, secara otomatis melewati konten yang sudah diunduh
- Dukungan multi-bahasa (20+ bahasa)
- Dukungan multi-platform (Windows, Linux, macOS, BSD)
- Konfigurasi proxy

## Instalasi

### Unduh Binari yang Sudah Dibuat

Unduh rilis terbaru dari halaman [Releases](https://github.com/reuAC/re-asmr-spider/releases).

#### Panduan Pemilihan Platform

| Sistem Operasi | Arsitektur CPU | Contoh Perangkat/Skenario | File Unduhan |
|-----------------|------------------|---------------------------|---------------|
| **Windows** | | | |
| Windows 11/10/8/7 | 64-bit Intel/AMD | Kebanyakan PC dan laptop modern | `Windows_x86_64.zip` |
| Windows 11/10/8/7 | 32-bit Intel/AMD | PC lama (sebelum 2010) | `Windows_i386.zip` |
| Windows 11 | ARM64 | Surface Pro X, laptop berbasis ARM | `Windows_arm64.zip` |
| **Linux** | | | |
| Ubuntu, Debian, Fedora, CentOS | 64-bit Intel/AMD (x86_64) | Kebanyakan server dan desktop | `Linux_x86_64.tar.gz` |
| Ubuntu, Debian, Fedora | 32-bit Intel/AMD (i386) | Sistem 32-bit lama | `Linux_i386.tar.gz` |
| Raspberry Pi OS, Ubuntu | ARM64 (aarch64) | Raspberry Pi 3/4/5 (OS 64-bit), NVIDIA Jetson | `Linux_arm64.tar.gz` |
| Raspberry Pi OS | ARMv7 (32-bit) | Raspberry Pi 2/3/4 (OS 32-bit) | `Linux_armv7.tar.gz` |
| Raspberry Pi OS | ARMv6 (32-bit) | Raspberry Pi 1, Pi Zero | `Linux_armv6.tar.gz` |
| OpenWrt, LEDE | MIPS (big-endian) | Beberapa router (MediaTek, Atheros) | `Linux_mips.tar.gz` |
| OpenWrt, LEDE | MIPS64 (big-endian) | Router MIPS kelas atas | `Linux_mips64.tar.gz` |
| OpenWrt, LEDE | MIPSle (little-endian) | Beberapa router (Broadcom, Realtek) | `Linux_mipsle.tar.gz` |
| OpenWrt, LEDE | MIPS64le (little-endian) | Sistem MIPS64 little-endian | `Linux_mips64le.tar.gz` |
| Debian, Fedora | RISC-V 64-bit | Board SiFive, StarFive VisionFive | `Linux_riscv64.tar.gz` |
| **macOS** | | | |
| macOS 10.13+ | Intel (x86_64) | MacBook Pro, iMac (2006-2020) | `Darwin_x86_64.tar.gz` |
| macOS 11+ | Apple Silicon (M1/M2/M3) | MacBook Air/Pro M1/M2/M3, Mac Mini M1/M2 | `Darwin_arm64.tar.gz` |
| **BSD** | | | |
| FreeBSD | 64-bit Intel/AMD | Server dan desktop FreeBSD | `Freebsd_x86_64.tar.gz` |
| FreeBSD | 32-bit Intel/AMD | Sistem FreeBSD lama | `Freebsd_i386.tar.gz` |
| FreeBSD | ARM64 | FreeBSD di board ARM | `Freebsd_arm64.tar.gz` |
| OpenBSD | 64-bit Intel/AMD | Sistem OpenBSD | `Openbsd_x86_64.tar.gz` |
| OpenBSD | 32-bit Intel/AMD | Sistem OpenBSD lama | `Openbsd_i386.tar.gz` |
| OpenBSD | ARM64 | OpenBSD di board ARM | `Openbsd_arm64.tar.gz` |
| NetBSD | 64-bit Intel/AMD | Sistem NetBSD | `Netbsd_x86_64.tar.gz` |
| NetBSD | 32-bit Intel/AMD | Sistem NetBSD lama | `Netbsd_i386.tar.gz` |
| NetBSD | ARM64 | NetBSD di board ARM | `Netbsd_arm64.tar.gz` |

#### Cara Memeriksa Sistem Anda

**Windows:**
- Klik kanan "This PC" > Properties
- Periksa "System type" (64-bit atau 32-bit, berbasis x64 atau ARM64)

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

Atau: Menu Apple > About This Mac > Chip (M1/M2/M3 = ARM64, Intel = x86_64)

### Build dari Source

```bash
git clone https://github.com/reuAC/re-asmr-spider.git
cd re-asmr-spider
go build -o re-asmr-spider
```

## Penggunaan

### Menjalankan Pertama Kali

Jalankan file executable. File konfigurasi default `config.json` akan dibuat secara otomatis.

```bash
./re-asmr-spider
```

### Konfigurasi

Edit `config.json` atau gunakan menu konfigurasi bawaan (opsi 2):

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

**Opsi Konfigurasi:**

- `account` - Nama pengguna akun ASMR.one
- `password` - Kata sandi akun ASMR.one
- `max_task` - Jumlah maksimum tugas unduhan bersamaan
- `max_thread` - Jumlah thread per unduhan file
- `max_retry` - Maksimum percobaan ulang untuk unduhan yang gagal
- `language` - Bahasa antarmuka (lihat bahasa yang didukung di bawah)
- `proxy` - Proxy HTTP/HTTPS (contoh: `http://127.0.0.1:7890`, kosongkan untuk menonaktifkan)

### Unduh Audio

1. Jalankan program
2. Pilih opsi `1. Start Download`
3. Masukkan nomor RJ (contoh: `RJ373001`)
4. Untuk beberapa unduhan, pisahkan nomor RJ dengan spasi (contoh: `RJ373001 RJ123456 RJ789012`)

Unduhan disimpan ke direktori `downloads/`.

### Melanjutkan Unduhan

Jika unduhan terputus, program akan mendeteksi tugas yang belum selesai saat dijalankan berikutnya dan akan menanyakan Anda untuk melanjutkan.

## Bahasa yang Didukung

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

Ganti bahasa melalui menu konfigurasi (opsi 4) atau dengan mengedit `config.json`.

## Sumber Data

Semua konten audio bersumber dari [asmr.one](https://asmr.one).

## Lisensi

MIT License

## Penghargaan

Berdasarkan [go-asmr-spider](https://github.com/DiheChen/go-asmr-spider) oleh DiheChen.

## Disclaimer

Alat ini hanya untuk penggunaan pribadi. Harap hormati pembuat konten dan dukung mereka melalui saluran resmi.
