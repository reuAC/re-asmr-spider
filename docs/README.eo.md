# re-ASMR-Spider

[简体中文](README.zh-CN.md) | [繁體中文](README.zh-TW.md) | [English](../README.md) | [日本語](README.ja-JP.md) | [Русский](README.ru-RU.md) | [Deutsch](README.de-DE.md) | [Español](README.es-ES.md) | [Français](README.fr-FR.md) | [Português](README.pt-PT.md) | [Indonesia](README.id-ID.md) | [Tiếng Việt](README.vi-VN.md) | [हिन्दी](README.hi-IN.md) | [বাংলা](README.bn-BD.md) | [తెలుగు](README.te-IN.md) | [Türkçe](README.tr-TR.md) | [اردو](README.ur-PK.md) | [Hausa](README.ha-NG.md) | Esperanto | [喵喵语](README.cat.md) | [火星语](README.martian.md)

---

Simpla kaj efika ASMR-aŭdio elŝuta ilo.

## Trajtoj

- Multfadena elŝuto kun rekomenca subteno
- Inteligenta reprova mekanismo por trakti reteroj erarojn
- Realtempa progresa montro
- Dosiergrandeca kontrolo, aŭtomate preterpasi elŝutitan enhavon
- Multlingva subteno (20+ lingvoj)
- Multplatforma subteno (Windows, Linux, macOS, BSD)
- Prokurila agordo

## Instalado

### Elŝuti Antaŭkonstruitajn Duumajn Dosierojn

Elŝutu la plej freŝan eldonon de la paĝo [Releases](https://github.com/reuAC/re-asmr-spider/releases).

#### Gvidilo pri Platformelekto

| Operaciumo | CPU-Arkitekturo | Ekzemplaj Aparatoj/Scenaroj | Elŝuta Dosiero |
|-----------------|------------------|---------------------------|---------------|
| **Windows** | | | |
| Windows 11/10/8/7 | 64-bita Intel/AMD | Plej modernaj PCoj kaj tekkomputiloj | `Windows_x86_64.zip` |
| Windows 11/10/8/7 | 32-bita Intel/AMD | Pli malnovaj PCoj (antaŭ 2010) | `Windows_i386.zip` |
| Windows 11 | ARM64 | Surface Pro X, ARM-bazitaj tekkomputiloj | `Windows_arm64.zip` |
| **Linux** | | | |
| Ubuntu, Debian, Fedora, CentOS | 64-bita Intel/AMD (x86_64) | Plej serviloj kaj labortabloj | `Linux_x86_64.tar.gz` |
| Ubuntu, Debian, Fedora | 32-bita Intel/AMD (i386) | Malnovaj 32-bitaj sistemoj | `Linux_i386.tar.gz` |
| Raspberry Pi OS, Ubuntu | ARM64 (aarch64) | Raspberry Pi 3/4/5 (64-bita OS), NVIDIA Jetson | `Linux_arm64.tar.gz` |
| Raspberry Pi OS | ARMv7 (32-bita) | Raspberry Pi 2/3/4 (32-bita OS) | `Linux_armv7.tar.gz` |
| Raspberry Pi OS | ARMv6 (32-bita) | Raspberry Pi 1, Pi Zero | `Linux_armv6.tar.gz` |
| OpenWrt, LEDE | MIPS (big-endian) | Iuj rutigistoj (MediaTek, Atheros) | `Linux_mips.tar.gz` |
| OpenWrt, LEDE | MIPS64 (big-endian) | Alt-nivelaj MIPS rutigistoj | `Linux_mips64.tar.gz` |
| OpenWrt, LEDE | MIPSle (little-endian) | Iuj rutigistoj (Broadcom, Realtek) | `Linux_mipsle.tar.gz` |
| OpenWrt, LEDE | MIPS64le (little-endian) | MIPS64 little-endian sistemoj | `Linux_mips64le.tar.gz` |
| Debian, Fedora | RISC-V 64-bita | SiFive-tabuloj, StarFive VisionFive | `Linux_riscv64.tar.gz` |
| **macOS** | | | |
| macOS 10.13+ | Intel (x86_64) | MacBook Pro, iMac (2006-2020) | `Darwin_x86_64.tar.gz` |
| macOS 11+ | Apple Silicon (M1/M2/M3) | MacBook Air/Pro M1/M2/M3, Mac Mini M1/M2 | `Darwin_arm64.tar.gz` |
| **BSD** | | | |
| FreeBSD | 64-bita Intel/AMD | FreeBSD-serviloj kaj labortabloj | `Freebsd_x86_64.tar.gz` |
| FreeBSD | 32-bita Intel/AMD | Pli malnovaj FreeBSD-sistemoj | `Freebsd_i386.tar.gz` |
| FreeBSD | ARM64 | FreeBSD sur ARM-tabuloj | `Freebsd_arm64.tar.gz` |
| OpenBSD | 64-bita Intel/AMD | OpenBSD-sistemoj | `Openbsd_x86_64.tar.gz` |
| OpenBSD | 32-bita Intel/AMD | Pli malnovaj OpenBSD-sistemoj | `Openbsd_i386.tar.gz` |
| OpenBSD | ARM64 | OpenBSD sur ARM-tabuloj | `Openbsd_arm64.tar.gz` |
| NetBSD | 64-bita Intel/AMD | NetBSD-sistemoj | `Netbsd_x86_64.tar.gz` |
| NetBSD | 32-bita Intel/AMD | Pli malnovaj NetBSD-sistemoj | `Netbsd_i386.tar.gz` |
| NetBSD | ARM64 | NetBSD sur ARM-tabuloj | `Netbsd_arm64.tar.gz` |

#### Kiel Kontroli Vian Sistemon

**Windows:**
- Dekstra-klaku "This PC" > Properties
- Kontrolu "System type" (64-bita aŭ 32-bita, x64-bazita aŭ ARM64-bazita)

**Linux:**
```bash
uname -m
# x86_64 -> 64-bita Intel/AMD
# i386/i686 -> 32-bita Intel/AMD
# aarch64 -> ARM64
# armv7l -> ARMv7
# armv6l -> ARMv6
# mips/mips64 -> MIPS
# riscv64 -> RISC-V 64-bita
```

**macOS:**
```bash
uname -m
# x86_64 -> Intel Mac
# arm64 -> Apple Silicon (M1/M2/M3)
```

Aŭ: Apple-menuo > About This Mac > Chip (M1/M2/M3 = ARM64, Intel = x86_64)

### Konstrui el Fontokodo

```bash
git clone https://github.com/reuAC/re-asmr-spider.git
cd re-asmr-spider
go build -o re-asmr-spider
```

## Uzado

### Unua Rulado

Rulu la plenumeblan dosieron. Defaŭlta agorda dosiero `config.json` estos kreita aŭtomate.

```bash
./re-asmr-spider
```

### Agordo

Redaktu `config.json` aŭ uzu la enkonstruitan agordan menuon (opcio 2):

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

**Agordaj Opcioj:**

- `account` - ASMR.one konta uzantnomo
- `password` - ASMR.one konta pasvorto
- `max_task` - Maksimuma samtempa elŝutaj taskoj
- `max_thread` - Nombro da fadenoj por unu dosiero elŝuto
- `max_retry` - Maksimuma reprovaj tentoj por malsukcesintaj elŝutoj
- `language` - Interfaca lingvo (vidu subtenatajn lingvojn sube)
- `proxy` - HTTP/HTTPS prokurilo (ekz., `http://127.0.0.1:7890`, lasu malplena por malŝalti)

### Elŝuti Aŭdion

1. Rulu la programon
2. Elektu opcion `1. Start Download`
3. Enigu RJ-numeron (ekz., `RJ373001`)
4. Por multoblaj elŝutoj, disigi RJ-numerojn per spacoj (ekz., `RJ373001 RJ123456 RJ789012`)

Elŝutoj estas konservitaj en la dosierujo `downloads/`.

### Rekomenci Elŝutojn

Se elŝuto estas interrompita, la programo detektos nekompletajn taskojn en la sekva rulado kaj petos vin daŭrigi.

## Subtenataj Lingvoj

- Ĉina (Simpligita & Tradicia)
- Angla
- Japana
- Germana
- Hispana
- Franca
- Portugala
- Rusa
- Indonezia
- Vjetnama
- Hindia
- Bengala
- Telugua
- Turka
- Urdua
- Haŭsa
- Esperanto
- Kata Lingvo
- Marsa

Ŝanĝu lingvon per la agorda menuo (opcio 4) aŭ redaktante `config.json`.

## Datumfonto

Ĉiuj aŭdio-enhavoj devenas de [asmr.one](https://asmr.one).

## Permesilo

MIT-Permesilo

## Dankoj

Bazita sur [go-asmr-spider](https://github.com/DiheChen/go-asmr-spider) de DiheChen.

## Malgarantio

Ĉi tiu ilo estas nur por persona uzo. Bonvolu respekti enhavkreantojn kaj subteni ilin per oficialaj kanaloj.
