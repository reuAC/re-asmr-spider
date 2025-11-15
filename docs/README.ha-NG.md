# re-ASMR-Spider

[简体中文](README.zh-CN.md) | [繁體中文](README.zh-TW.md) | [English](../README.md) | [日本語](README.ja-JP.md) | [Русский](README.ru-RU.md) | [Deutsch](README.de-DE.md) | [Español](README.es-ES.md) | [Français](README.fr-FR.md) | [Português](README.pt-PT.md) | [Indonesia](README.id-ID.md) | [Tiếng Việt](README.vi-VN.md) | [हिन्दी](README.hi-IN.md) | [বাংলা](README.bn-BD.md) | [తెలుగు](README.te-IN.md) | [Türkçe](README.tr-TR.md) | [اردو](README.ur-PK.md) | Hausa | [Esperanto](README.eo.md) | [喵喵语](README.cat.md) | [火星语](README.martian.md)

---

Kayan aikin saukar da sauti na ASMR mai sauƙi da inganci.

## Fasalulluka

- Saukar da fayil da yawa tare da goyon bayan ci gaba
- Tsarin sake gwadawa mai wayo don magance matsalolin hanyar sadarwa
- Nuna ci gaba na lokaci-lokaci
- Tabbatar da girman fayil, tsallake abin da aka sauke ta atomatik
- Goyon bayan harsuna da yawa (20+ harsuna)
- Goyon bayan dandamali da yawa (Windows, Linux, macOS, BSD)
- Saitunan wakili

## Hotunan Allo
![Gwajin Sauke](./images/asmr-spider-0.png)
![Sake Gwadawa](./images/asmr-spider-1.png)

## Shigarwa

### Saukar da Binaries da aka Gina

Saukar da sabon saki daga shafin [Releases](https://github.com/reuAC/re-asmr-spider/releases).

#### Jagorar Zaɓin Dandamali

| Tsarin Aiki | Gine-ginen CPU | Misalai na Na'urori/Yanayi | Saukar da Fayil |
|-----------------|------------------|---------------------------|---------------|
| **Windows** | | | |
| Windows 11/10/8/7 | 64-bit Intel/AMD | Yawancin PC da kwamfyutocin kwakwalwa na zamani | `Windows_x86_64.zip` |
| Windows 11/10/8/7 | 32-bit Intel/AMD | Tsofaffin PC (kafin 2010) | `Windows_i386.zip` |
| Windows 11 | ARM64 | Surface Pro X, kwamfyutocin ARM | `Windows_arm64.zip` |
| **Linux** | | | |
| Ubuntu, Debian, Fedora, CentOS | 64-bit Intel/AMD (x86_64) | Yawancin sabar da tebur | `Linux_x86_64.tar.gz` |
| Ubuntu, Debian, Fedora | 32-bit Intel/AMD (i386) | Tsofaffin tsarin 32-bit | `Linux_i386.tar.gz` |
| Raspberry Pi OS, Ubuntu | ARM64 (aarch64) | Raspberry Pi 3/4/5 (64-bit OS), NVIDIA Jetson | `Linux_arm64.tar.gz` |
| Raspberry Pi OS | ARMv7 (32-bit) | Raspberry Pi 2/3/4 (32-bit OS) | `Linux_armv7.tar.gz` |
| Raspberry Pi OS | ARMv6 (32-bit) | Raspberry Pi 1, Pi Zero | `Linux_armv6.tar.gz` |
| OpenWrt, LEDE | MIPS (big-endian) | Wasu routers (MediaTek, Atheros) | `Linux_mips.tar.gz` |
| OpenWrt, LEDE | MIPS64 (big-endian) | Routers na MIPS masu girma | `Linux_mips64.tar.gz` |
| OpenWrt, LEDE | MIPSle (little-endian) | Wasu routers (Broadcom, Realtek) | `Linux_mipsle.tar.gz` |
| OpenWrt, LEDE | MIPS64le (little-endian) | Tsarin MIPS64 little-endian | `Linux_mips64le.tar.gz` |
| Debian, Fedora | RISC-V 64-bit | Allunan SiFive, StarFive VisionFive | `Linux_riscv64.tar.gz` |
| **macOS** | | | |
| macOS 10.13+ | Intel (x86_64) | MacBook Pro, iMac (2006-2020) | `Darwin_x86_64.tar.gz` |
| macOS 11+ | Apple Silicon (M1/M2/M3) | MacBook Air/Pro M1/M2/M3, Mac Mini M1/M2 | `Darwin_arm64.tar.gz` |
| **BSD** | | | |
| FreeBSD | 64-bit Intel/AMD | Sabar da tebur na FreeBSD | `Freebsd_x86_64.tar.gz` |
| FreeBSD | 32-bit Intel/AMD | Tsofaffin tsarin FreeBSD | `Freebsd_i386.tar.gz` |
| FreeBSD | ARM64 | FreeBSD akan allunan ARM | `Freebsd_arm64.tar.gz` |
| OpenBSD | 64-bit Intel/AMD | Tsarin OpenBSD | `Openbsd_x86_64.tar.gz` |
| OpenBSD | 32-bit Intel/AMD | Tsofaffin tsarin OpenBSD | `Openbsd_i386.tar.gz` |
| OpenBSD | ARM64 | OpenBSD akan allunan ARM | `Openbsd_arm64.tar.gz` |
| NetBSD | 64-bit Intel/AMD | Tsarin NetBSD | `Netbsd_x86_64.tar.gz` |
| NetBSD | 32-bit Intel/AMD | Tsofaffin tsarin NetBSD | `Netbsd_i386.tar.gz` |
| NetBSD | ARM64 | NetBSD akan allunan ARM | `Netbsd_arm64.tar.gz` |

#### Yadda Ake Bincika Tsarin Ku

**Windows:**
- Danna dama "This PC" > Properties
- Duba "System type" (64-bit ko 32-bit, x64-based ko ARM64-based)

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

Ko kuma: Menu na Apple > About This Mac > Chip (M1/M2/M3 = ARM64, Intel = x86_64)

### Gina daga Tushe

```bash
git clone https://github.com/reuAC/re-asmr-spider.git
cd re-asmr-spider
go build -o re-asmr-spider
```

## Amfani

### Gudu na Farko

Gudanar da fayil mai aiwatarwa. Za a ƙirƙiri fayil na daidaitawa na tsoho `config.json` ta atomatik.

```bash
./re-asmr-spider
```

### Daidaitawa

Gyara `config.json` ko yi amfani da menu na daidaitawa na ciki (zaɓi 2):

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

**Zaɓuɓɓukan Daidaitawa:**

- `account` - Sunan mai amfani na asusun ASMR.one
- `password` - Kalmar sirri na asusun ASMR.one
- `max_task` - Matsakaicin ayyukan saukar da lokaci ɗaya
- `max_thread` - Adadin zaren saukar da fayil ɗaya
- `max_retry` - Matsakaicin ƙoƙarin sake gwadawa don saukar da fayilolin da suka gaza
- `language` - Harshen mu'amala (duba harsuna da aka goyan baya a ƙasa)
- `proxy` - Wakili na HTTP/HTTPS (misali, `http://127.0.0.1:7890`, bar komai don kashe shi)

### Saukar da Sauti

1. Gudanar da shirin
2. Zaɓi zaɓi `1. Start Download`
3. Shigar da lambar RJ (misali, `RJ373001`)
4. Don saukar da fayiloli da yawa, raba lambobin RJ tare da sarari (misali, `RJ373001 RJ123456 RJ789012`)

Ana adana abubuwan da aka sauke zuwa kundin adireshi `downloads/`.

### Ci gaba da Saukarwa

Idan an katse saukarwa, shirin zai gano ayyukan da ba a cika ba a gudu na gaba kuma ya ɗaga ku don ci gaba.

## Harsuna da aka Goyan baya

- Sinanci (Mai sauƙi da Na gargajiya)
- Turanci
- Japananci
- Jamusanci
- Sifen
- Faransanci
- Fotigal
- Rashanci
- Indunusiya
- Harshen Vietnam
- Hindi
- Bengali
- Telugu
- Baturke
- Urdu
- Hausa
- Esperanto
- Harshen Cat
- Harshen Mars

Canja harshe ta menu na daidaitawa (zaɓi 4) ko ta gyara `config.json`.

## Tushen Bayanai

Dukkan abun ciki na sauti ana samun su daga [asmr.one](https://asmr.one).

## Lasisi

Lasisi na MIT

## Godiya

Dangane da [go-asmr-spider](https://github.com/DiheChen/go-asmr-spider) na DiheChen.

## Sanarwa

Wannan kayan aiki don amfani na mutum kaɗai ne. Da fatan za a girmama masu ƙirƙirar abun ciki kuma a goyi bayan su ta hanyoyin hukuma.
