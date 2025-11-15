# re-ASMR-Spider

[简体中文](README.zh-CN.md) | [繁體中文](README.zh-TW.md) | [English](../README.md) | [日本語](README.ja-JP.md) | [Русский](README.ru-RU.md) | [Deutsch](README.de-DE.md) | [Español](README.es-ES.md) | [Français](README.fr-FR.md) | [Português](README.pt-PT.md) | [Indonesia](README.id-ID.md) | [Tiếng Việt](README.vi-VN.md) | [हिन्दी](README.hi-IN.md) | [বাংলা](README.bn-BD.md) | [తెలుగు](README.te-IN.md) | [Türkçe](README.tr-TR.md) | اردو | [Hausa](README.ha-NG.md) | [Esperanto](README.eo.md) | [喵喵语](README.cat.md) | [火星语](README.martian.md)

---

ایک سادہ اور موثر ASMR آڈیو ڈاؤن لوڈ ٹول۔

## خصوصیات

- دوبارہ شروع کرنے کی معاونت کے ساتھ ملٹی تھریڈڈ ڈاؤن لوڈ
- نیٹ ورک کی خرابیوں کو سنبھالنے کے لیے سمارٹ دوبارہ کوشش کا طریقہ کار
- ریئل ٹائم پیش رفت کی نمائش
- فائل سائز کی تصدیق، ڈاؤن لوڈ شدہ مواد کو خودکار طور پر چھوڑ دیتا ہے
- کثیر زبان کی معاونت (20+ زبانیں)
- کثیر پلیٹ فارم کی معاونت (Windows, Linux, macOS, BSD)
- پراکسی کنفیگریشن

## اسکرین شاٹس
![ڈاؤن لوڈ کی کوشش](./images/asmr-spider-0.png)
![دوبارہ کوشش](./images/asmr-spider-1.png)

## تنصیب

### پہلے سے تیار شدہ بائنریز ڈاؤن لوڈ کریں

[Releases](https://github.com/reuAC/re-asmr-spider/releases) صفحہ سے تازہ ترین ریلیز ڈاؤن لوڈ کریں۔

#### پلیٹ فارم کے انتخاب کی رہنمائی

| آپریٹنگ سسٹم | CPU آرکیٹیکچر | مثال کے آلات/منظرنامے | ڈاؤن لوڈ فائل |
|-----------------|------------------|---------------------------|---------------|
| **Windows** | | | |
| Windows 11/10/8/7 | 64-bit Intel/AMD | زیادہ تر جدید PC اور لیپ ٹاپ | `Windows_x86_64.zip` |
| Windows 11/10/8/7 | 32-bit Intel/AMD | پرانے PC (2010 سے پہلے) | `Windows_i386.zip` |
| Windows 11 | ARM64 | Surface Pro X, ARM پر مبنی لیپ ٹاپ | `Windows_arm64.zip` |
| **Linux** | | | |
| Ubuntu, Debian, Fedora, CentOS | 64-bit Intel/AMD (x86_64) | زیادہ تر سرورز اور ڈیسک ٹاپ | `Linux_x86_64.tar.gz` |
| Ubuntu, Debian, Fedora | 32-bit Intel/AMD (i386) | پرانے 32-bit سسٹم | `Linux_i386.tar.gz` |
| Raspberry Pi OS, Ubuntu | ARM64 (aarch64) | Raspberry Pi 3/4/5 (64-bit OS), NVIDIA Jetson | `Linux_arm64.tar.gz` |
| Raspberry Pi OS | ARMv7 (32-bit) | Raspberry Pi 2/3/4 (32-bit OS) | `Linux_armv7.tar.gz` |
| Raspberry Pi OS | ARMv6 (32-bit) | Raspberry Pi 1, Pi Zero | `Linux_armv6.tar.gz` |
| OpenWrt, LEDE | MIPS (big-endian) | کچھ راؤٹرز (MediaTek, Atheros) | `Linux_mips.tar.gz` |
| OpenWrt, LEDE | MIPS64 (big-endian) | ہائی اینڈ MIPS راؤٹرز | `Linux_mips64.tar.gz` |
| OpenWrt, LEDE | MIPSle (little-endian) | کچھ راؤٹرز (Broadcom, Realtek) | `Linux_mipsle.tar.gz` |
| OpenWrt, LEDE | MIPS64le (little-endian) | MIPS64 little-endian سسٹمز | `Linux_mips64le.tar.gz` |
| Debian, Fedora | RISC-V 64-bit | SiFive boards, StarFive VisionFive | `Linux_riscv64.tar.gz` |
| **macOS** | | | |
| macOS 10.13+ | Intel (x86_64) | MacBook Pro, iMac (2006-2020) | `Darwin_x86_64.tar.gz` |
| macOS 11+ | Apple Silicon (M1/M2/M3) | MacBook Air/Pro M1/M2/M3, Mac Mini M1/M2 | `Darwin_arm64.tar.gz` |
| **BSD** | | | |
| FreeBSD | 64-bit Intel/AMD | FreeBSD سرورز اور ڈیسک ٹاپ | `Freebsd_x86_64.tar.gz` |
| FreeBSD | 32-bit Intel/AMD | پرانے FreeBSD سسٹمز | `Freebsd_i386.tar.gz` |
| FreeBSD | ARM64 | ARM بورڈز پر FreeBSD | `Freebsd_arm64.tar.gz` |
| OpenBSD | 64-bit Intel/AMD | OpenBSD سسٹمز | `Openbsd_x86_64.tar.gz` |
| OpenBSD | 32-bit Intel/AMD | پرانے OpenBSD سسٹمز | `Openbsd_i386.tar.gz` |
| OpenBSD | ARM64 | ARM بورڈز پر OpenBSD | `Openbsd_arm64.tar.gz` |
| NetBSD | 64-bit Intel/AMD | NetBSD سسٹمز | `Netbsd_x86_64.tar.gz` |
| NetBSD | 32-bit Intel/AMD | پرانے NetBSD سسٹمز | `Netbsd_i386.tar.gz` |
| NetBSD | ARM64 | ARM بورڈز پر NetBSD | `Netbsd_arm64.tar.gz` |

#### اپنے سسٹم کی جانچ کیسے کریں

**Windows:**
- "This PC" پر دائیں کلک کریں > Properties
- "System type" چیک کریں (64-bit یا 32-bit، x64 پر مبنی یا ARM64 پر مبنی)

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

یا: Apple مینو > About This Mac > Chip (M1/M2/M3 = ARM64, Intel = x86_64)

### سورس سے بلڈ کریں

```bash
git clone https://github.com/reuAC/re-asmr-spider.git
cd re-asmr-spider
go build -o re-asmr-spider
```

## استعمال

### پہلا رن

ایگزیکیوٹیبل چلائیں۔ ایک ڈیفالٹ کنفیگریشن فائل `config.json` خودکار طور پر بنائی جائے گی۔

```bash
./re-asmr-spider
```

### کنفیگریشن

`config.json` میں ترمیم کریں یا بلٹ ان کنفیگریشن مینو (آپشن 2) استعمال کریں:

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

**کنفیگریشن کے اختیارات:**

- `account` - ASMR.one اکاؤنٹ کا صارف نام
- `password` - ASMR.one اکاؤنٹ کا پاس ورڈ
- `max_task` - زیادہ سے زیادہ بیک وقت ڈاؤن لوڈ ٹاسک
- `max_thread` - فی فائل ڈاؤن لوڈ میں تھریڈز کی تعداد
- `max_retry` - ناکام ڈاؤن لوڈز کے لیے زیادہ سے زیادہ دوبارہ کوشش کی تعداد
- `language` - انٹرفیس کی زبان (نیچے معاون زبانیں دیکھیں)
- `proxy` - HTTP/HTTPS پراکسی (مثلاً، `http://127.0.0.1:7890`، غیر فعال کرنے کے لیے خالی چھوڑیں)

### آڈیو ڈاؤن لوڈ کریں

1. پروگرام چلائیں
2. آپشن `1. Start Download` منتخب کریں
3. RJ نمبر درج کریں (مثلاً، `RJ373001`)
4. متعدد ڈاؤن لوڈز کے لیے، RJ نمبروں کو خالی جگہوں سے الگ کریں (مثلاً، `RJ373001 RJ123456 RJ789012`)

ڈاؤن لوڈز `downloads/` ڈائریکٹری میں محفوظ ہوتے ہیں۔

### ڈاؤن لوڈز دوبارہ شروع کریں

اگر ڈاؤن لوڈ میں رکاوٹ آتی ہے، تو پروگرام اگلے رن پر نامکمل ٹاسک کا پتہ لگائے گا اور آپ سے جاری رکھنے کے لیے کہے گا۔

## معاون زبانیں

- چینی (آسان اور روایتی)
- انگریزی
- جاپانی
- جرمن
- ہسپانوی
- فرانسیسی
- پرتگالی
- روسی
- انڈونیشیائی
- ویتنامی
- ہندی
- بنگالی
- تیلگو
- ترکی
- اردو
- ہاؤسا
- اسپرانٹو
- بلی کی زبان
- مریخی

کنفیگریشن مینو (آپشن 4) کے ذریعے یا `config.json` میں ترمیم کرکے زبان تبدیل کریں۔

## ڈیٹا کا ذریعہ

تمام آڈیو مواد [asmr.one](https://asmr.one) سے حاصل کیا جاتا ہے۔

## لائسنس

MIT License

## تشکرات

DiheChen کے [go-asmr-spider](https://github.com/DiheChen/go-asmr-spider) پر مبنی۔

## دستبرداری

یہ ٹول صرف ذاتی استعمال کے لیے ہے۔ براہ کرم مواد تخلیق کاروں کا احترام کریں اور سرکاری چینلز کے ذریعے ان کی حمایت کریں۔
