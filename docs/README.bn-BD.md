# re-ASMR-Spider

[简体中文](README.zh-CN.md) | [繁體中文](README.zh-TW.md) | [English](../README.md) | [日本語](README.ja-JP.md) | [Русский](README.ru-RU.md) | [Deutsch](README.de-DE.md) | [Español](README.es-ES.md) | [Français](README.fr-FR.md) | [Português](README.pt-PT.md) | [Indonesia](README.id-ID.md) | [Tiếng Việt](README.vi-VN.md) | [हिन्दी](README.hi-IN.md) | বাংলা | [తెలుగు](README.te-IN.md) | [Türkçe](README.tr-TR.md) | [اردو](README.ur-PK.md) | [Hausa](README.ha-NG.md) | [Esperanto](README.eo.md) | [喵喵语](README.cat.md) | [火星语](README.martian.md)

---

একটি সহজ এবং দক্ষ ASMR অডিও ডাউনলোড টুল।

## বৈশিষ্ট্য

- পুনরায় শুরু সমর্থন সহ মাল্টি-থ্রেডেড ডাউনলোড
- নেটওয়ার্ক ত্রুটি পরিচালনার জন্য স্মার্ট পুনরায় চেষ্টা প্রক্রিয়া
- রিয়েল-টাইম অগ্রগতি প্রদর্শন
- ফাইল সাইজ যাচাইকরণ, ডাউনলোড করা বিষয়বস্তু স্বয়ংক্রিয়ভাবে এড়িয়ে যায়
- বহু-ভাষা সমর্থন (২০+ ভাষা)
- বহু-প্ল্যাটফর্ম সমর্থন (Windows, Linux, macOS, BSD)
- প্রক্সি কনফিগারেশন

## স্ক্রিনশট
![ডাউনলোড প্রচেষ্টা](./images/asmr-spider-0.png)
![পুনরায় চেষ্টা ডাউনলোড](./images/asmr-spider-1.png)

## ইনস্টলেশন

### পূর্ব-নির্মিত বাইনারি ডাউনলোড করুন

[Releases](https://github.com/reuAC/re-asmr-spider/releases) পৃষ্ঠা থেকে সর্বশেষ রিলিজ ডাউনলোড করুন।

#### প্ল্যাটফর্ম নির্বাচন গাইড

| অপারেটিং সিস্টেম | CPU আর্কিটেকচার | উদাহরণ ডিভাইস/পরিস্থিতি | ডাউনলোড ফাইল |
|-----------------|------------------|---------------------------|---------------|
| **Windows** | | | |
| Windows 11/10/8/7 | 64-বিট Intel/AMD | বেশিরভাগ আধুনিক PC এবং ল্যাপটপ | `Windows_x86_64.zip` |
| Windows 11/10/8/7 | 32-বিট Intel/AMD | পুরানো PC (২০১০-এর আগে) | `Windows_i386.zip` |
| Windows 11 | ARM64 | Surface Pro X, ARM-ভিত্তিক ল্যাপটপ | `Windows_arm64.zip` |
| **Linux** | | | |
| Ubuntu, Debian, Fedora, CentOS | 64-বিট Intel/AMD (x86_64) | বেশিরভাগ সার্ভার এবং ডেস্কটপ | `Linux_x86_64.tar.gz` |
| Ubuntu, Debian, Fedora | 32-বিট Intel/AMD (i386) | পুরানো 32-বিট সিস্টেম | `Linux_i386.tar.gz` |
| Raspberry Pi OS, Ubuntu | ARM64 (aarch64) | Raspberry Pi 3/4/5 (64-বিট OS), NVIDIA Jetson | `Linux_arm64.tar.gz` |
| Raspberry Pi OS | ARMv7 (32-বিট) | Raspberry Pi 2/3/4 (32-বিট OS) | `Linux_armv7.tar.gz` |
| Raspberry Pi OS | ARMv6 (32-বিট) | Raspberry Pi 1, Pi Zero | `Linux_armv6.tar.gz` |
| OpenWrt, LEDE | MIPS (big-endian) | কিছু রাউটার (MediaTek, Atheros) | `Linux_mips.tar.gz` |
| OpenWrt, LEDE | MIPS64 (big-endian) | হাই-এন্ড MIPS রাউটার | `Linux_mips64.tar.gz` |
| OpenWrt, LEDE | MIPSle (little-endian) | কিছু রাউটার (Broadcom, Realtek) | `Linux_mipsle.tar.gz` |
| OpenWrt, LEDE | MIPS64le (little-endian) | MIPS64 little-endian সিস্টেম | `Linux_mips64le.tar.gz` |
| Debian, Fedora | RISC-V 64-বিট | SiFive boards, StarFive VisionFive | `Linux_riscv64.tar.gz` |
| **macOS** | | | |
| macOS 10.13+ | Intel (x86_64) | MacBook Pro, iMac (2006-2020) | `Darwin_x86_64.tar.gz` |
| macOS 11+ | Apple Silicon (M1/M2/M3) | MacBook Air/Pro M1/M2/M3, Mac Mini M1/M2 | `Darwin_arm64.tar.gz` |
| **BSD** | | | |
| FreeBSD | 64-বিট Intel/AMD | FreeBSD সার্ভার এবং ডেস্কটপ | `Freebsd_x86_64.tar.gz` |
| FreeBSD | 32-বিট Intel/AMD | পুরানো FreeBSD সিস্টেম | `Freebsd_i386.tar.gz` |
| FreeBSD | ARM64 | ARM বোর্ডে FreeBSD | `Freebsd_arm64.tar.gz` |
| OpenBSD | 64-বিট Intel/AMD | OpenBSD সিস্টেম | `Openbsd_x86_64.tar.gz` |
| OpenBSD | 32-বিট Intel/AMD | পুরানো OpenBSD সিস্টেম | `Openbsd_i386.tar.gz` |
| OpenBSD | ARM64 | ARM বোর্ডে OpenBSD | `Openbsd_arm64.tar.gz` |
| NetBSD | 64-বিট Intel/AMD | NetBSD সিস্টেম | `Netbsd_x86_64.tar.gz` |
| NetBSD | 32-বিট Intel/AMD | পুরানো NetBSD সিস্টেম | `Netbsd_i386.tar.gz` |
| NetBSD | ARM64 | ARM বোর্ডে NetBSD | `Netbsd_arm64.tar.gz` |

#### কীভাবে আপনার সিস্টেম পরীক্ষা করবেন

**Windows:**
- "This PC"-তে রাইট-ক্লিক করুন > Properties
- "System type" চেক করুন (64-বিট বা 32-বিট, x64-ভিত্তিক বা ARM64-ভিত্তিক)

**Linux:**
```bash
uname -m
# x86_64 -> 64-বিট Intel/AMD
# i386/i686 -> 32-বিট Intel/AMD
# aarch64 -> ARM64
# armv7l -> ARMv7
# armv6l -> ARMv6
# mips/mips64 -> MIPS
# riscv64 -> RISC-V 64-বিট
```

**macOS:**
```bash
uname -m
# x86_64 -> Intel Mac
# arm64 -> Apple Silicon (M1/M2/M3)
```

অথবা: Apple মেনু > About This Mac > Chip (M1/M2/M3 = ARM64, Intel = x86_64)

### সোর্স থেকে বিল্ড করুন

```bash
git clone https://github.com/reuAC/re-asmr-spider.git
cd re-asmr-spider
go build -o re-asmr-spider
```

## ব্যবহার

### প্রথম রান

এক্সিকিউটেবল চালান। একটি ডিফল্ট কনফিগারেশন ফাইল `config.json` স্বয়ংক্রিয়ভাবে তৈরি হবে।

```bash
./re-asmr-spider
```

### কনফিগারেশন

`config.json` সম্পাদনা করুন অথবা বিল্ট-ইন কনফিগারেশন মেনু (বিকল্প ২) ব্যবহার করুন:

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

**কনফিগারেশন অপশন:**

- `account` - ASMR.one অ্যাকাউন্ট ইউজারনেম
- `password` - ASMR.one অ্যাকাউন্ট পাসওয়ার্ড
- `max_task` - সর্বোচ্চ সমসাময়িক ডাউনলোড টাস্ক
- `max_thread` - প্রতি ফাইল ডাউনলোডে থ্রেডের সংখ্যা
- `max_retry` - ব্যর্থ ডাউনলোডের জন্য সর্বোচ্চ পুনরায় চেষ্টা
- `language` - ইন্টারফেস ভাষা (নিচে সমর্থিত ভাষা দেখুন)
- `proxy` - HTTP/HTTPS প্রক্সি (যেমন, `http://127.0.0.1:7890`, নিষ্ক্রিয় করতে খালি রাখুন)

### অডিও ডাউনলোড করুন

1. প্রোগ্রাম চালান
2. বিকল্প `1. Start Download` নির্বাচন করুন
3. RJ নম্বর লিখুন (যেমন, `RJ373001`)
4. একাধিক ডাউনলোডের জন্য, RJ নম্বরগুলি স্পেস দিয়ে আলাদা করুন (যেমন, `RJ373001 RJ123456 RJ789012`)

ডাউনলোডগুলি `downloads/` ডিরেক্টরিতে সংরক্ষিত হয়।

### ডাউনলোড পুনরায় শুরু করুন

যদি একটি ডাউনলোড বাধাগ্রস্ত হয়, প্রোগ্রামটি পরবর্তী রানে অসম্পূর্ণ টাস্ক সনাক্ত করবে এবং আপনাকে চালিয়ে যেতে বলবে।

## সমর্থিত ভাষা

- চীনা (সরলীকৃত ও ঐতিহ্যবাহী)
- ইংরেজি
- জাপানি
- জার্মান
- স্প্যানিশ
- ফরাসি
- পর্তুগিজ
- রাশিয়ান
- ইন্দোনেশীয়
- ভিয়েতনামী
- হিন্দি
- বাংলা
- তেলুগু
- তুর্কি
- উর্দু
- হাউসা
- এসপেরান্তো
- বিড়াল ভাষা
- মার্টিয়ান

কনফিগারেশন মেনু (বিকল্প ৪) এর মাধ্যমে বা `config.json` সম্পাদনা করে ভাষা পরিবর্তন করুন।

## ডেটা সোর্স

সমস্ত অডিও বিষয়বস্তু [asmr.one](https://asmr.one) থেকে সংগৃহীত।

## লাইসেন্স

MIT License

## কৃতজ্ঞতা

DiheChen-এর [go-asmr-spider](https://github.com/DiheChen/go-asmr-spider) এর উপর ভিত্তি করে তৈরি।

## দাবিত্যাগ

এই টুলটি শুধুমাত্র ব্যক্তিগত ব্যবহারের জন্য। অনুগ্রহ করে বিষয়বস্তু নির্মাতাদের সম্মান করুন এবং অফিসিয়াল চ্যানেলের মাধ্যমে তাদের সমর্থন করুন।
