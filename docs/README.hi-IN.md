# re-ASMR-Spider

[简体中文](README.zh-CN.md) | [繁體中文](README.zh-TW.md) | [English](../README.md) | [日本語](README.ja-JP.md) | [Русский](README.ru-RU.md) | [Deutsch](README.de-DE.md) | [Español](README.es-ES.md) | [Français](README.fr-FR.md) | [Português](README.pt-PT.md) | [Indonesia](README.id-ID.md) | [Tiếng Việt](README.vi-VN.md) | हिन्दी | [বাংলা](README.bn-BD.md) | [తెలుగు](README.te-IN.md) | [Türkçe](README.tr-TR.md) | [اردو](README.ur-PK.md) | [Hausa](README.ha-NG.md) | [Esperanto](README.eo.md) | [喵喵语](README.cat.md) | [火星语](README.martian.md)

---

एक सरल और कुशल ASMR ऑडियो डाउनलोड टूल।

## विशेषताएं

- पुनः शुरू समर्थन के साथ मल्टी-थ्रेडेड डाउनलोड
- नेटवर्क त्रुटियों को संभालने के लिए स्मार्ट पुनः प्रयास तंत्र
- रियल-टाइम प्रगति प्रदर्शन
- फ़ाइल आकार सत्यापन, डाउनलोड की गई सामग्री को स्वचालित रूप से छोड़ें
- बहु-भाषा समर्थन (20+ भाषाएं)
- बहु-मंच समर्थन (Windows, Linux, macOS, BSD)
- प्रॉक्सी कॉन्फ़िगरेशन

## स्क्रीनशॉट
![डाउनलोड प्रयास](./images/asmr-spider-0.png)
![पुनः प्रयास डाउनलोड](./images/asmr-spider-1.png)

## इंस्टॉलेशन

### पूर्व-निर्मित बाइनरी डाउनलोड करें

[Releases](https://github.com/reuAC/re-asmr-spider/releases) पृष्ठ से नवीनतम रिलीज़ डाउनलोड करें।

#### प्लेटफ़ॉर्म चयन गाइड

| ऑपरेटिंग सिस्टम | CPU आर्किटेक्चर | उदाहरण डिवाइस/परिदृश्य | डाउनलोड फ़ाइल |
|-----------------|------------------|---------------------------|---------------|
| **Windows** | | | |
| Windows 11/10/8/7 | 64-bit Intel/AMD | अधिकांश आधुनिक PC और लैपटॉप | `Windows_x86_64.zip` |
| Windows 11/10/8/7 | 32-bit Intel/AMD | पुराने PC (2010 से पहले) | `Windows_i386.zip` |
| Windows 11 | ARM64 | Surface Pro X, ARM-आधारित लैपटॉप | `Windows_arm64.zip` |
| **Linux** | | | |
| Ubuntu, Debian, Fedora, CentOS | 64-bit Intel/AMD (x86_64) | अधिकांश सर्वर और डेस्कटॉप | `Linux_x86_64.tar.gz` |
| Ubuntu, Debian, Fedora | 32-bit Intel/AMD (i386) | पुरानी 32-bit सिस्टम | `Linux_i386.tar.gz` |
| Raspberry Pi OS, Ubuntu | ARM64 (aarch64) | Raspberry Pi 3/4/5 (64-bit OS), NVIDIA Jetson | `Linux_arm64.tar.gz` |
| Raspberry Pi OS | ARMv7 (32-bit) | Raspberry Pi 2/3/4 (32-bit OS) | `Linux_armv7.tar.gz` |
| Raspberry Pi OS | ARMv6 (32-bit) | Raspberry Pi 1, Pi Zero | `Linux_armv6.tar.gz` |
| OpenWrt, LEDE | MIPS (big-endian) | कुछ राउटर (MediaTek, Atheros) | `Linux_mips.tar.gz` |
| OpenWrt, LEDE | MIPS64 (big-endian) | हाई-एंड MIPS राउटर | `Linux_mips64.tar.gz` |
| OpenWrt, LEDE | MIPSle (little-endian) | कुछ राउटर (Broadcom, Realtek) | `Linux_mipsle.tar.gz` |
| OpenWrt, LEDE | MIPS64le (little-endian) | MIPS64 little-endian सिस्टम | `Linux_mips64le.tar.gz` |
| Debian, Fedora | RISC-V 64-bit | SiFive बोर्ड, StarFive VisionFive | `Linux_riscv64.tar.gz` |
| **macOS** | | | |
| macOS 10.13+ | Intel (x86_64) | MacBook Pro, iMac (2006-2020) | `Darwin_x86_64.tar.gz` |
| macOS 11+ | Apple Silicon (M1/M2/M3) | MacBook Air/Pro M1/M2/M3, Mac Mini M1/M2 | `Darwin_arm64.tar.gz` |
| **BSD** | | | |
| FreeBSD | 64-bit Intel/AMD | FreeBSD सर्वर और डेस्कटॉप | `Freebsd_x86_64.tar.gz` |
| FreeBSD | 32-bit Intel/AMD | पुरानी FreeBSD सिस्टम | `Freebsd_i386.tar.gz` |
| FreeBSD | ARM64 | ARM बोर्ड पर FreeBSD | `Freebsd_arm64.tar.gz` |
| OpenBSD | 64-bit Intel/AMD | OpenBSD सिस्टम | `Openbsd_x86_64.tar.gz` |
| OpenBSD | 32-bit Intel/AMD | पुरानी OpenBSD सिस्टम | `Openbsd_i386.tar.gz` |
| OpenBSD | ARM64 | ARM बोर्ड पर OpenBSD | `Openbsd_arm64.tar.gz` |
| NetBSD | 64-bit Intel/AMD | NetBSD सिस्टम | `Netbsd_x86_64.tar.gz` |
| NetBSD | 32-bit Intel/AMD | पुरानी NetBSD सिस्टम | `Netbsd_i386.tar.gz` |
| NetBSD | ARM64 | ARM बोर्ड पर NetBSD | `Netbsd_arm64.tar.gz` |

#### अपनी सिस्टम की जांच कैसे करें

**Windows:**
- "This PC" पर राइट-क्लिक करें > Properties
- "System type" की जांच करें (64-bit या 32-bit, x64-आधारित या ARM64-आधारित)

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

या: Apple मेनू > About This Mac > Chip (M1/M2/M3 = ARM64, Intel = x86_64)

### सोर्स से बिल्ड करें

```bash
git clone https://github.com/reuAC/re-asmr-spider.git
cd re-asmr-spider
go build -o re-asmr-spider
```

## उपयोग

### पहली बार चलाना

एक्जीक्यूटेबल चलाएं। एक डिफ़ॉल्ट कॉन्फ़िगरेशन फ़ाइल `config.json` स्वचालित रूप से बनाई जाएगी।

```bash
./re-asmr-spider
```

### कॉन्फ़िगरेशन

`config.json` संपादित करें या अंतर्निहित कॉन्फ़िगरेशन मेनू का उपयोग करें (विकल्प 2):

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

**कॉन्फ़िगरेशन विकल्प:**

- `account` - ASMR.one खाता उपयोगकर्ता नाम
- `password` - ASMR.one खाता पासवर्ड
- `max_task` - अधिकतम समवर्ती डाउनलोड कार्य
- `max_thread` - प्रति फ़ाइल डाउनलोड में थ्रेड की संख्या
- `max_retry` - विफल डाउनलोड के लिए अधिकतम पुनः प्रयास प्रयास
- `language` - इंटरफ़ेस भाषा (नीचे समर्थित भाषाएं देखें)
- `proxy` - HTTP/HTTPS प्रॉक्सी (उदा., `http://127.0.0.1:7890`, अक्षम करने के लिए खाली छोड़ें)

### ऑडियो डाउनलोड करें

1. प्रोग्राम चलाएं
2. विकल्प `1. Start Download` चुनें
3. RJ नंबर दर्ज करें (उदा., `RJ373001`)
4. एकाधिक डाउनलोड के लिए, RJ नंबरों को स्पेस से अलग करें (उदा., `RJ373001 RJ123456 RJ789012`)

डाउनलोड `downloads/` डायरेक्टरी में सहेजे जाते हैं।

### डाउनलोड फिर से शुरू करें

यदि डाउनलोड बाधित हो जाता है, तो प्रोग्राम अगली बार चलाने पर अधूरे कार्यों का पता लगाएगा और आपसे जारी रखने के लिए संकेत देगा।

## समर्थित भाषाएं

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

कॉन्फ़िगरेशन मेनू (विकल्प 4) के माध्यम से या `config.json` संपादित करके भाषा स्विच करें।

## डेटा स्रोत

सभी ऑडियो सामग्री [asmr.one](https://asmr.one) से प्राप्त की गई है।

## लाइसेंस

MIT License

## आभार

DiheChen द्वारा [go-asmr-spider](https://github.com/DiheChen/go-asmr-spider) पर आधारित।

## अस्वीकरण

यह टूल केवल व्यक्तिगत उपयोग के लिए है। कृपया सामग्री निर्माताओं का सम्मान करें और आधिकारिक चैनलों के माध्यम से उनका समर्थन करें।
