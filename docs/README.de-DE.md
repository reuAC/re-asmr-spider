# re-ASMR-Spider

[简体中文](README.zh-CN.md) | [繁體中文](README.zh-TW.md) | [English](../README.md) | [日本語](README.ja-JP.md) | [Русский](README.ru-RU.md) | Deutsch | [Español](README.es-ES.md) | [Français](README.fr-FR.md) | [Português](README.pt-PT.md) | [Indonesia](README.id-ID.md) | [Tiếng Việt](README.vi-VN.md) | [हिन्दी](README.hi-IN.md) | [বাংলা](README.bn-BD.md) | [తెలుగు](README.te-IN.md) | [Türkçe](README.tr-TR.md) | [اردو](README.ur-PK.md) | [Hausa](README.ha-NG.md) | [Esperanto](README.eo.md) | [喵喵语](README.cat.md) | [火星语](README.martian.md)

---

Ein einfaches und effizientes Tool zum Herunterladen von ASMR-Audiodateien.

## Funktionen

- Multithreaded-Download mit Unterstützung für Fortsetzung
- Intelligenter Wiederholungsmechanismus zur Behandlung von Netzwerkfehlern
- Echtzeit-Fortschrittsanzeige
- Überprüfung der Dateigröße, heruntergeladene Inhalte werden automatisch übersprungen
- Mehrsprachige Unterstützung (über 20 Sprachen)
- Multiplattform-Unterstützung (Windows, Linux, macOS, BSD)
- Proxy-Konfiguration

## Screenshots
![Download-Versuch](./images/asmr-spider-0.png)
![Wiederholter Download](./images/asmr-spider-1.png)

## Installation

### Vorgefertigte Binärdateien herunterladen

Laden Sie die neueste Version von der [Releases](https://github.com/reuAC/re-asmr-spider/releases)-Seite herunter.

#### Plattform-Auswahlhilfe

| Betriebssystem | CPU-Architektur | Beispielgeräte/Szenarien | Download-Datei |
|-----------------|------------------|---------------------------|---------------|
| **Windows** | | | |
| Windows 11/10/8/7 | 64-bit Intel/AMD | Die meisten modernen PCs und Laptops | `Windows_x86_64.zip` |
| Windows 11/10/8/7 | 32-bit Intel/AMD | Ältere PCs (vor 2010) | `Windows_i386.zip` |
| Windows 11 | ARM64 | Surface Pro X, ARM-basierte Laptops | `Windows_arm64.zip` |
| **Linux** | | | |
| Ubuntu, Debian, Fedora, CentOS | 64-bit Intel/AMD (x86_64) | Die meisten Server und Desktops | `Linux_x86_64.tar.gz` |
| Ubuntu, Debian, Fedora | 32-bit Intel/AMD (i386) | Alte 32-Bit-Systeme | `Linux_i386.tar.gz` |
| Raspberry Pi OS, Ubuntu | ARM64 (aarch64) | Raspberry Pi 3/4/5 (64-Bit-OS), NVIDIA Jetson | `Linux_arm64.tar.gz` |
| Raspberry Pi OS | ARMv7 (32-bit) | Raspberry Pi 2/3/4 (32-Bit-OS) | `Linux_armv7.tar.gz` |
| Raspberry Pi OS | ARMv6 (32-bit) | Raspberry Pi 1, Pi Zero | `Linux_armv6.tar.gz` |
| OpenWrt, LEDE | MIPS (big-endian) | Einige Router (MediaTek, Atheros) | `Linux_mips.tar.gz` |
| OpenWrt, LEDE | MIPS64 (big-endian) | High-End-MIPS-Router | `Linux_mips64.tar.gz` |
| OpenWrt, LEDE | MIPSle (little-endian) | Einige Router (Broadcom, Realtek) | `Linux_mipsle.tar.gz` |
| OpenWrt, LEDE | MIPS64le (little-endian) | MIPS64-Little-Endian-Systeme | `Linux_mips64le.tar.gz` |
| Debian, Fedora | RISC-V 64-bit | SiFive-Boards, StarFive VisionFive | `Linux_riscv64.tar.gz` |
| **macOS** | | | |
| macOS 10.13+ | Intel (x86_64) | MacBook Pro, iMac (2006-2020) | `Darwin_x86_64.tar.gz` |
| macOS 11+ | Apple Silicon (M1/M2/M3) | MacBook Air/Pro M1/M2/M3, Mac Mini M1/M2 | `Darwin_arm64.tar.gz` |
| **BSD** | | | |
| FreeBSD | 64-bit Intel/AMD | FreeBSD-Server und -Desktops | `Freebsd_x86_64.tar.gz` |
| FreeBSD | 32-bit Intel/AMD | Ältere FreeBSD-Systeme | `Freebsd_i386.tar.gz` |
| FreeBSD | ARM64 | FreeBSD auf ARM-Boards | `Freebsd_arm64.tar.gz` |
| OpenBSD | 64-bit Intel/AMD | OpenBSD-Systeme | `Openbsd_x86_64.tar.gz` |
| OpenBSD | 32-bit Intel/AMD | Ältere OpenBSD-Systeme | `Openbsd_i386.tar.gz` |
| OpenBSD | ARM64 | OpenBSD auf ARM-Boards | `Openbsd_arm64.tar.gz` |
| NetBSD | 64-bit Intel/AMD | NetBSD-Systeme | `Netbsd_x86_64.tar.gz` |
| NetBSD | 32-bit Intel/AMD | Ältere NetBSD-Systeme | `Netbsd_i386.tar.gz` |
| NetBSD | ARM64 | NetBSD auf ARM-Boards | `Netbsd_arm64.tar.gz` |

#### So überprüfen Sie Ihr System

**Windows:**
- Rechtsklick auf "Dieser PC" > Eigenschaften
- Überprüfen Sie "Systemtyp" (64-Bit oder 32-Bit, x64-basiert oder ARM64-basiert)

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

Oder: Apple-Menü > Über diesen Mac > Chip (M1/M2/M3 = ARM64, Intel = x86_64)

### Aus dem Quellcode erstellen

```bash
git clone https://github.com/reuAC/re-asmr-spider.git
cd re-asmr-spider
go build -o re-asmr-spider
```

## Verwendung

### Erster Start

Führen Sie die ausführbare Datei aus. Eine Standard-Konfigurationsdatei `config.json` wird automatisch erstellt.

```bash
./re-asmr-spider
```

### Konfiguration

Bearbeiten Sie `config.json` oder verwenden Sie das integrierte Konfigurationsmenü (Option 2):

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

**Konfigurationsoptionen:**

- `account` - ASMR.one-Konto-Benutzername
- `password` - ASMR.one-Konto-Passwort
- `max_task` - Maximale Anzahl gleichzeitiger Download-Aufgaben
- `max_thread` - Anzahl der Threads pro Datei-Download
- `max_retry` - Maximale Anzahl von Wiederholungsversuchen für fehlgeschlagene Downloads
- `language` - Oberflächensprache (siehe unterstützte Sprachen unten)
- `proxy` - HTTP/HTTPS-Proxy (z.B. `http://127.0.0.1:7890`, leer lassen zum Deaktivieren)

### Audio herunterladen

1. Führen Sie das Programm aus
2. Wählen Sie Option `1. Start Download`
3. Geben Sie die RJ-Nummer ein (z.B. `RJ373001`)
4. Für mehrere Downloads trennen Sie die RJ-Nummern mit Leerzeichen (z.B. `RJ373001 RJ123456 RJ789012`)

Downloads werden im Verzeichnis `downloads/` gespeichert.

### Downloads fortsetzen

Wenn ein Download unterbrochen wird, erkennt das Programm beim nächsten Start unvollständige Aufgaben und fordert Sie auf, fortzufahren.

## Unterstützte Sprachen

- Chinesisch (Vereinfacht & Traditionell)
- Englisch
- Japanisch
- Deutsch
- Spanisch
- Französisch
- Portugiesisch
- Russisch
- Indonesisch
- Vietnamesisch
- Hindi
- Bengali
- Telugu
- Türkisch
- Urdu
- Hausa
- Esperanto
- Katzensprache
- Marsianisch

Wechseln Sie die Sprache über das Konfigurationsmenü (Option 4) oder durch Bearbeiten von `config.json`.

## Datenquelle

Alle Audioinhalte stammen von [asmr.one](https://asmr.one).

## Lizenz

MIT License

## Danksagungen

Basierend auf [go-asmr-spider](https://github.com/DiheChen/go-asmr-spider) von DiheChen.

## Haftungsausschluss

Dieses Tool ist nur für den persönlichen Gebrauch bestimmt. Bitte respektieren Sie die Content-Ersteller und unterstützen Sie sie über offizielle Kanäle.
