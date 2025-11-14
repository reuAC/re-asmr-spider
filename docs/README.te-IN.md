# re-ASMR-Spider

[简体中文](README.zh-CN.md) | [繁體中文](README.zh-TW.md) | [English](../README.md) | [日本語](README.ja-JP.md) | [Русский](README.ru-RU.md) | [Deutsch](README.de-DE.md) | [Español](README.es-ES.md) | [Français](README.fr-FR.md) | [Português](README.pt-PT.md) | [Indonesia](README.id-ID.md) | [Tiếng Việt](README.vi-VN.md) | [हिन्दी](README.hi-IN.md) | [বাংলা](README.bn-BD.md) | తెలుగు | [Türkçe](README.tr-TR.md) | [اردو](README.ur-PK.md) | [Hausa](README.ha-NG.md) | [Esperanto](README.eo.md) | [喵喵语](README.cat.md) | [火星语](README.martian.md)

---

సరళమైన మరియు సమర్థవంతమైన ASMR ఆడియో డౌన్‌లోడ్ సాధనం.

## లక్షణాలు

- పునఃప్రారంభ మద్దతుతో మల్టీ-థ్రెడ్ డౌన్‌లోడ్
- నెట్‌వర్క్ లోపాలను నిర్వహించడానికి స్మార్ట్ రీట్రై మెకానిజం
- రియల్-టైమ్ ప్రోగ్రెస్ డిస్‌ప్లే
- ఫైల్ సైజ్ వెరిఫికేషన్, డౌన్‌లోడ్ చేసిన కంటెంట్‌ను స్వయంచాలకంగా దాటవేస్తుంది
- మల్టీ-లాంగ్వేజ్ మద్దతు (20+ భాషలు)
- మల్టీ-ప్లాట్‌ఫారమ్ మద్దతు (Windows, Linux, macOS, BSD)
- ప్రాక్సీ కాన్ఫిగరేషన్

## ఇన్‌స్టాలేషన్

### ముందుగా-నిర్మించిన బైనరీలను డౌన్‌లోడ్ చేయండి

[Releases](https://github.com/reuAC/re-asmr-spider/releases) పేజీ నుండి తాజా విడుదలను డౌన్‌లోడ్ చేయండి.

#### ప్లాట్‌ఫారమ్ ఎంపిక మార్గదర్శిని

| ఆపరేటింగ్ సిస్టమ్ | CPU ఆర్కిటెక్చర్ | ఉదాహరణ పరికరాలు/పరిస్థితులు | డౌన్‌లోడ్ ఫైల్ |
|-----------------|------------------|---------------------------|---------------|
| **Windows** | | | |
| Windows 11/10/8/7 | 64-బిట్ Intel/AMD | చాలా ఆధునిక PCలు మరియు ల్యాప్‌టాప్‌లు | `Windows_x86_64.zip` |
| Windows 11/10/8/7 | 32-బిట్ Intel/AMD | పాత PCలు (2010కు ముందు) | `Windows_i386.zip` |
| Windows 11 | ARM64 | Surface Pro X, ARM-ఆధారిత ల్యాప్‌టాప్‌లు | `Windows_arm64.zip` |
| **Linux** | | | |
| Ubuntu, Debian, Fedora, CentOS | 64-బిట్ Intel/AMD (x86_64) | చాలా సర్వర్లు మరియు డెస్క్‌టాప్‌లు | `Linux_x86_64.tar.gz` |
| Ubuntu, Debian, Fedora | 32-బిట్ Intel/AMD (i386) | పాత 32-బిట్ సిస్టమ్‌లు | `Linux_i386.tar.gz` |
| Raspberry Pi OS, Ubuntu | ARM64 (aarch64) | Raspberry Pi 3/4/5 (64-బిట్ OS), NVIDIA Jetson | `Linux_arm64.tar.gz` |
| Raspberry Pi OS | ARMv7 (32-బిట్) | Raspberry Pi 2/3/4 (32-బిట్ OS) | `Linux_armv7.tar.gz` |
| Raspberry Pi OS | ARMv6 (32-బిట్) | Raspberry Pi 1, Pi Zero | `Linux_armv6.tar.gz` |
| OpenWrt, LEDE | MIPS (big-endian) | కొన్ని రూటర్లు (MediaTek, Atheros) | `Linux_mips.tar.gz` |
| OpenWrt, LEDE | MIPS64 (big-endian) | హై-ఎండ్ MIPS రూటర్లు | `Linux_mips64.tar.gz` |
| OpenWrt, LEDE | MIPSle (little-endian) | కొన్ని రూటర్లు (Broadcom, Realtek) | `Linux_mipsle.tar.gz` |
| OpenWrt, LEDE | MIPS64le (little-endian) | MIPS64 little-endian సిస్టమ్‌లు | `Linux_mips64le.tar.gz` |
| Debian, Fedora | RISC-V 64-బిట్ | SiFive boards, StarFive VisionFive | `Linux_riscv64.tar.gz` |
| **macOS** | | | |
| macOS 10.13+ | Intel (x86_64) | MacBook Pro, iMac (2006-2020) | `Darwin_x86_64.tar.gz` |
| macOS 11+ | Apple Silicon (M1/M2/M3) | MacBook Air/Pro M1/M2/M3, Mac Mini M1/M2 | `Darwin_arm64.tar.gz` |
| **BSD** | | | |
| FreeBSD | 64-బిట్ Intel/AMD | FreeBSD సర్వర్లు మరియు డెస్క్‌టాప్‌లు | `Freebsd_x86_64.tar.gz` |
| FreeBSD | 32-బిట్ Intel/AMD | పాత FreeBSD సిస్టమ్‌లు | `Freebsd_i386.tar.gz` |
| FreeBSD | ARM64 | ARM బోర్డ్‌లలో FreeBSD | `Freebsd_arm64.tar.gz` |
| OpenBSD | 64-బిట్ Intel/AMD | OpenBSD సిస్టమ్‌లు | `Openbsd_x86_64.tar.gz` |
| OpenBSD | 32-బిట్ Intel/AMD | పాత OpenBSD సిస్టమ్‌లు | `Openbsd_i386.tar.gz` |
| OpenBSD | ARM64 | ARM బోర్డ్‌లలో OpenBSD | `Openbsd_arm64.tar.gz` |
| NetBSD | 64-బిట్ Intel/AMD | NetBSD సిస్టమ్‌లు | `Netbsd_x86_64.tar.gz` |
| NetBSD | 32-బిట్ Intel/AMD | పాత NetBSD సిస్టమ్‌లు | `Netbsd_i386.tar.gz` |
| NetBSD | ARM64 | ARM బోర్డ్‌లలో NetBSD | `Netbsd_arm64.tar.gz` |

#### మీ సిస్టమ్‌ను ఎలా తనిఖీ చేయాలి

**Windows:**
- "This PC"పై రైట్-క్లిక్ చేయండి > Properties
- "System type" తనిఖీ చేయండి (64-బిట్ లేదా 32-బిట్, x64-ఆధారిత లేదా ARM64-ఆధారిత)

**Linux:**
```bash
uname -m
# x86_64 -> 64-బిట్ Intel/AMD
# i386/i686 -> 32-బిట్ Intel/AMD
# aarch64 -> ARM64
# armv7l -> ARMv7
# armv6l -> ARMv6
# mips/mips64 -> MIPS
# riscv64 -> RISC-V 64-బిట్
```

**macOS:**
```bash
uname -m
# x86_64 -> Intel Mac
# arm64 -> Apple Silicon (M1/M2/M3)
```

లేదా: Apple మెను > About This Mac > Chip (M1/M2/M3 = ARM64, Intel = x86_64)

### సోర్స్ నుండి బిల్డ్ చేయండి

```bash
git clone https://github.com/reuAC/re-asmr-spider.git
cd re-asmr-spider
go build -o re-asmr-spider
```

## ఉపయోగం

### మొదటి రన్

ఎగ్జిక్యూటబుల్‌ను రన్ చేయండి. డిఫాల్ట్ కాన్ఫిగరేషన్ ఫైల్ `config.json` స్వయంచాలకంగా సృష్టించబడుతుంది.

```bash
./re-asmr-spider
```

### కాన్ఫిగరేషన్

`config.json` ఎడిట్ చేయండి లేదా బిల్ట్-ఇన్ కాన్ఫిగరేషన్ మెను (ఆప్షన్ 2) ఉపయోగించండి:

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

**కాన్ఫిగరేషన్ ఆప్షన్‌లు:**

- `account` - ASMR.one ఖాతా యూజర్‌నేమ్
- `password` - ASMR.one ఖాతా పాస్‌వర్డ్
- `max_task` - గరిష్ట ఏకకాల డౌన్‌లోడ్ టాస్క్‌లు
- `max_thread` - ఒక్కో ఫైల్ డౌన్‌లోడ్‌కు థ్రెడ్‌ల సంఖ్య
- `max_retry` - విఫలమైన డౌన్‌లోడ్‌ల కోసం గరిష్ట రీట్రై ప్రయత్నాలు
- `language` - ఇంటర్‌ఫేస్ భాష (క్రింద మద్దతు ఉన్న భాషలను చూడండి)
- `proxy` - HTTP/HTTPS ప్రాక్సీ (ఉదా., `http://127.0.0.1:7890`, నిలిపివేయడానికి ఖాళీగా ఉంచండి)

### ఆడియో డౌన్‌లోడ్ చేయండి

1. ప్రోగ్రామ్ రన్ చేయండి
2. ఆప్షన్ `1. Start Download` ఎంచుకోండి
3. RJ నంబర్ నమోదు చేయండి (ఉదా., `RJ373001`)
4. బహుళ డౌన్‌లోడ్‌ల కోసం, RJ నంబర్‌లను స్పేస్‌లతో వేరు చేయండి (ఉదా., `RJ373001 RJ123456 RJ789012`)

డౌన్‌లోడ్‌లు `downloads/` డైరెక్టరీలో సేవ్ చేయబడతాయి.

### డౌన్‌లోడ్‌లను పునఃప్రారంభించండి

డౌన్‌లోడ్ అంతరాయం కలిగితే, తదుపరి రన్‌లో ప్రోగ్రామ్ అసంపూర్ణ టాస్క్‌లను గుర్తిస్తుంది మరియు కొనసాగించమని మిమ్మల్ని అడుగుతుంది.

## మద్దతు ఉన్న భాషలు

- చైనీస్ (సరళీకృత & సాంప్రదాయ)
- ఇంగ్లీష్
- జపనీస్
- జర్మన్
- స్పానిష్
- ఫ్రెంచ్
- పోర్చుగీస్
- రష్యన్
- ఇండోనేషియన్
- వియత్నామీస్
- హిందీ
- బెంగాలీ
- తెలుగు
- టర్కిష్
- ఉర్దూ
- హౌసా
- ఎస్పెరాంటో
- పిల్లి భాష
- మార్టియన్

కాన్ఫిగరేషన్ మెను (ఆప్షన్ 4) ద్వారా లేదా `config.json` ఎడిట్ చేయడం ద్వారా భాషను మార్చండి.

## డేటా మూలం

అన్ని ఆడియో కంటెంట్ [asmr.one](https://asmr.one) నుండి పొందబడింది.

## లైసెన్స్

MIT License

## కృతజ్ఞతలు

DiheChen యొక్క [go-asmr-spider](https://github.com/DiheChen/go-asmr-spider) ఆధారంగా రూపొందించబడింది.

## నిరాకరణ

ఈ సాధనం వ్యక్తిగత ఉపయోగం కోసం మాత్రమే. దయచేసి కంటెంట్ సృష్టికర్తలను గౌరవించండి మరియు అధికారిక ఛానెల్‌ల ద్వారా వారికి మద్దతు ఇవ్వండి.
