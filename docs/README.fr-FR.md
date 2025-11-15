# re-ASMR-Spider

[简体中文](README.zh-CN.md) | [繁體中文](README.zh-TW.md) | [English](../README.md) | [日本語](README.ja-JP.md) | [Русский](README.ru-RU.md) | [Deutsch](README.de-DE.md) | [Español](README.es-ES.md) | Français | [Português](README.pt-PT.md) | [Indonesia](README.id-ID.md) | [Tiếng Việt](README.vi-VN.md) | [हिन्दी](README.hi-IN.md) | [বাংলা](README.bn-BD.md) | [తెలుగు](README.te-IN.md) | [Türkçe](README.tr-TR.md) | [اردو](README.ur-PK.md) | [Hausa](README.ha-NG.md) | [Esperanto](README.eo.md) | [喵喵语](README.cat.md) | [火星语](README.martian.md)

---

Un outil simple et efficace pour télécharger des fichiers audio ASMR.

## Fonctionnalités

- Téléchargement multi-thread avec support de reprise
- Mécanisme de réessai intelligent pour gérer les erreurs réseau
- Affichage de la progression en temps réel
- Vérification de la taille des fichiers, ignore automatiquement le contenu téléchargé
- Support multilingue (plus de 20 langues)
- Support multiplateforme (Windows, Linux, macOS, BSD)
- Configuration du proxy

## Captures d'écran
![Tentative de téléchargement](./images/asmr-spider-0.png)
![Réessayer le téléchargement](./images/asmr-spider-1.png)

## Installation

### Télécharger les binaires précompilés

Téléchargez la dernière version depuis la page [Releases](https://github.com/reuAC/re-asmr-spider/releases).

#### Guide de sélection de plateforme

| Système d'exploitation | Architecture CPU | Exemples d'appareils/Scénarios | Fichier de téléchargement |
|-----------------|------------------|---------------------------|---------------|
| **Windows** | | | |
| Windows 11/10/8/7 | 64-bit Intel/AMD | La plupart des PC et ordinateurs portables modernes | `Windows_x86_64.zip` |
| Windows 11/10/8/7 | 32-bit Intel/AMD | PC anciens (avant 2010) | `Windows_i386.zip` |
| Windows 11 | ARM64 | Surface Pro X, ordinateurs portables basés sur ARM | `Windows_arm64.zip` |
| **Linux** | | | |
| Ubuntu, Debian, Fedora, CentOS | 64-bit Intel/AMD (x86_64) | La plupart des serveurs et bureaux | `Linux_x86_64.tar.gz` |
| Ubuntu, Debian, Fedora | 32-bit Intel/AMD (i386) | Anciens systèmes 32 bits | `Linux_i386.tar.gz` |
| Raspberry Pi OS, Ubuntu | ARM64 (aarch64) | Raspberry Pi 3/4/5 (OS 64 bits), NVIDIA Jetson | `Linux_arm64.tar.gz` |
| Raspberry Pi OS | ARMv7 (32-bit) | Raspberry Pi 2/3/4 (OS 32 bits) | `Linux_armv7.tar.gz` |
| Raspberry Pi OS | ARMv6 (32-bit) | Raspberry Pi 1, Pi Zero | `Linux_armv6.tar.gz` |
| OpenWrt, LEDE | MIPS (big-endian) | Certains routeurs (MediaTek, Atheros) | `Linux_mips.tar.gz` |
| OpenWrt, LEDE | MIPS64 (big-endian) | Routeurs MIPS haut de gamme | `Linux_mips64.tar.gz` |
| OpenWrt, LEDE | MIPSle (little-endian) | Certains routeurs (Broadcom, Realtek) | `Linux_mipsle.tar.gz` |
| OpenWrt, LEDE | MIPS64le (little-endian) | Systèmes MIPS64 little-endian | `Linux_mips64le.tar.gz` |
| Debian, Fedora | RISC-V 64-bit | Cartes SiFive, StarFive VisionFive | `Linux_riscv64.tar.gz` |
| **macOS** | | | |
| macOS 10.13+ | Intel (x86_64) | MacBook Pro, iMac (2006-2020) | `Darwin_x86_64.tar.gz` |
| macOS 11+ | Apple Silicon (M1/M2/M3) | MacBook Air/Pro M1/M2/M3, Mac Mini M1/M2 | `Darwin_arm64.tar.gz` |
| **BSD** | | | |
| FreeBSD | 64-bit Intel/AMD | Serveurs et bureaux FreeBSD | `Freebsd_x86_64.tar.gz` |
| FreeBSD | 32-bit Intel/AMD | Anciens systèmes FreeBSD | `Freebsd_i386.tar.gz` |
| FreeBSD | ARM64 | FreeBSD sur cartes ARM | `Freebsd_arm64.tar.gz` |
| OpenBSD | 64-bit Intel/AMD | Systèmes OpenBSD | `Openbsd_x86_64.tar.gz` |
| OpenBSD | 32-bit Intel/AMD | Anciens systèmes OpenBSD | `Openbsd_i386.tar.gz` |
| OpenBSD | ARM64 | OpenBSD sur cartes ARM | `Openbsd_arm64.tar.gz` |
| NetBSD | 64-bit Intel/AMD | Systèmes NetBSD | `Netbsd_x86_64.tar.gz` |
| NetBSD | 32-bit Intel/AMD | Anciens systèmes NetBSD | `Netbsd_i386.tar.gz` |
| NetBSD | ARM64 | NetBSD sur cartes ARM | `Netbsd_arm64.tar.gz` |

#### Comment vérifier votre système

**Windows :**
- Clic droit sur "Ce PC" > Propriétés
- Vérifiez "Type de système" (64 bits ou 32 bits, basé sur x64 ou basé sur ARM64)

**Linux :**
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

**macOS :**
```bash
uname -m
# x86_64 -> Intel Mac
# arm64 -> Apple Silicon (M1/M2/M3)
```

Ou : menu Apple > À propos de ce Mac > Puce (M1/M2/M3 = ARM64, Intel = x86_64)

### Compiler à partir du code source

```bash
git clone https://github.com/reuAC/re-asmr-spider.git
cd re-asmr-spider
go build -o re-asmr-spider
```

## Utilisation

### Première exécution

Exécutez le fichier exécutable. Un fichier de configuration par défaut `config.json` sera créé automatiquement.

```bash
./re-asmr-spider
```

### Configuration

Modifiez `config.json` ou utilisez le menu de configuration intégré (option 2) :

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

**Options de configuration :**

- `account` - Nom d'utilisateur du compte ASMR.one
- `password` - Mot de passe du compte ASMR.one
- `max_task` - Nombre maximum de tâches de téléchargement simultanées
- `max_thread` - Nombre de threads par téléchargement de fichier
- `max_retry` - Nombre maximum de tentatives pour les téléchargements échoués
- `language` - Langue de l'interface (voir les langues prises en charge ci-dessous)
- `proxy` - Proxy HTTP/HTTPS (par exemple, `http://127.0.0.1:7890`, laissez vide pour désactiver)

### Télécharger de l'audio

1. Exécutez le programme
2. Sélectionnez l'option `1. Start Download`
3. Entrez le numéro RJ (par exemple, `RJ373001`)
4. Pour plusieurs téléchargements, séparez les numéros RJ avec des espaces (par exemple, `RJ373001 RJ123456 RJ789012`)

Les téléchargements sont enregistrés dans le répertoire `downloads/`.

### Reprendre les téléchargements

Si un téléchargement est interrompu, le programme détectera les tâches incomplètes lors de la prochaine exécution et vous invitera à continuer.

## Langues prises en charge

- Chinois (Simplifié et Traditionnel)
- Anglais
- Japonais
- Allemand
- Espagnol
- Français
- Portugais
- Russe
- Indonésien
- Vietnamien
- Hindi
- Bengali
- Telugu
- Turc
- Ourdou
- Haoussa
- Espéranto
- Langue de chat
- Martien

Changez de langue via le menu de configuration (option 4) ou en modifiant `config.json`.

## Source de données

Tout le contenu audio provient de [asmr.one](https://asmr.one).

## Licence

MIT License

## Remerciements

Basé sur [go-asmr-spider](https://github.com/DiheChen/go-asmr-spider) par DiheChen.

## Clause de non-responsabilité

Cet outil est destiné à un usage personnel uniquement. Veuillez respecter les créateurs de contenu et les soutenir via les canaux officiels.
