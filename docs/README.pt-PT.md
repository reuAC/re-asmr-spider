# re-ASMR-Spider

[简体中文](README.zh-CN.md) | [繁體中文](README.zh-TW.md) | [English](../README.md) | [日本語](README.ja-JP.md) | [Русский](README.ru-RU.md) | [Deutsch](README.de-DE.md) | [Español](README.es-ES.md) | [Français](README.fr-FR.md) | Português | [Indonesia](README.id-ID.md) | [Tiếng Việt](README.vi-VN.md) | [हिन्दी](README.hi-IN.md) | [বাংলা](README.bn-BD.md) | [తెలుగు](README.te-IN.md) | [Türkçe](README.tr-TR.md) | [اردو](README.ur-PK.md) | [Hausa](README.ha-NG.md) | [Esperanto](README.eo.md) | [喵喵语](README.cat.md) | [火星语](README.martian.md)

---

Uma ferramenta simples e eficiente para transferir ficheiros de áudio ASMR.

## Funcionalidades

- Transferência multithreaded com suporte para retomar
- Mecanismo de repetição inteligente para lidar com erros de rede
- Visualização do progresso em tempo real
- Verificação do tamanho do ficheiro, ignora automaticamente conteúdo transferido
- Suporte multilíngue (mais de 20 idiomas)
- Suporte multiplataforma (Windows, Linux, macOS, BSD)
- Configuração de proxy

## Instalação

### Transferir binários pré-compilados

Transfira a versão mais recente da página [Releases](https://github.com/reuAC/re-asmr-spider/releases).

#### Guia de seleção de plataforma

| Sistema operativo | Arquitetura CPU | Exemplos de dispositivos/Cenários | Ficheiro de transferência |
|-----------------|------------------|---------------------------|---------------|
| **Windows** | | | |
| Windows 11/10/8/7 | 64-bit Intel/AMD | A maioria dos PCs e portáteis modernos | `Windows_x86_64.zip` |
| Windows 11/10/8/7 | 32-bit Intel/AMD | PCs antigos (pré-2010) | `Windows_i386.zip` |
| Windows 11 | ARM64 | Surface Pro X, portáteis baseados em ARM | `Windows_arm64.zip` |
| **Linux** | | | |
| Ubuntu, Debian, Fedora, CentOS | 64-bit Intel/AMD (x86_64) | A maioria dos servidores e desktops | `Linux_x86_64.tar.gz` |
| Ubuntu, Debian, Fedora | 32-bit Intel/AMD (i386) | Sistemas antigos de 32 bits | `Linux_i386.tar.gz` |
| Raspberry Pi OS, Ubuntu | ARM64 (aarch64) | Raspberry Pi 3/4/5 (SO 64 bits), NVIDIA Jetson | `Linux_arm64.tar.gz` |
| Raspberry Pi OS | ARMv7 (32-bit) | Raspberry Pi 2/3/4 (SO 32 bits) | `Linux_armv7.tar.gz` |
| Raspberry Pi OS | ARMv6 (32-bit) | Raspberry Pi 1, Pi Zero | `Linux_armv6.tar.gz` |
| OpenWrt, LEDE | MIPS (big-endian) | Alguns routers (MediaTek, Atheros) | `Linux_mips.tar.gz` |
| OpenWrt, LEDE | MIPS64 (big-endian) | Routers MIPS de gama alta | `Linux_mips64.tar.gz` |
| OpenWrt, LEDE | MIPSle (little-endian) | Alguns routers (Broadcom, Realtek) | `Linux_mipsle.tar.gz` |
| OpenWrt, LEDE | MIPS64le (little-endian) | Sistemas MIPS64 little-endian | `Linux_mips64le.tar.gz` |
| Debian, Fedora | RISC-V 64-bit | Placas SiFive, StarFive VisionFive | `Linux_riscv64.tar.gz` |
| **macOS** | | | |
| macOS 10.13+ | Intel (x86_64) | MacBook Pro, iMac (2006-2020) | `Darwin_x86_64.tar.gz` |
| macOS 11+ | Apple Silicon (M1/M2/M3) | MacBook Air/Pro M1/M2/M3, Mac Mini M1/M2 | `Darwin_arm64.tar.gz` |
| **BSD** | | | |
| FreeBSD | 64-bit Intel/AMD | Servidores e desktops FreeBSD | `Freebsd_x86_64.tar.gz` |
| FreeBSD | 32-bit Intel/AMD | Sistemas FreeBSD antigos | `Freebsd_i386.tar.gz` |
| FreeBSD | ARM64 | FreeBSD em placas ARM | `Freebsd_arm64.tar.gz` |
| OpenBSD | 64-bit Intel/AMD | Sistemas OpenBSD | `Openbsd_x86_64.tar.gz` |
| OpenBSD | 32-bit Intel/AMD | Sistemas OpenBSD antigos | `Openbsd_i386.tar.gz` |
| OpenBSD | ARM64 | OpenBSD em placas ARM | `Openbsd_arm64.tar.gz` |
| NetBSD | 64-bit Intel/AMD | Sistemas NetBSD | `Netbsd_x86_64.tar.gz` |
| NetBSD | 32-bit Intel/AMD | Sistemas NetBSD antigos | `Netbsd_i386.tar.gz` |
| NetBSD | ARM64 | NetBSD em placas ARM | `Netbsd_arm64.tar.gz` |

#### Como verificar o seu sistema

**Windows:**
- Clique com o botão direito em "Este PC" > Propriedades
- Verifique "Tipo de sistema" (64 bits ou 32 bits, baseado em x64 ou baseado em ARM64)

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

Ou: menu Apple > Acerca deste Mac > Chip (M1/M2/M3 = ARM64, Intel = x86_64)

### Compilar a partir do código fonte

```bash
git clone https://github.com/reuAC/re-asmr-spider.git
cd re-asmr-spider
go build -o re-asmr-spider
```

## Utilização

### Primeira execução

Execute o ficheiro executável. Um ficheiro de configuração predefinido `config.json` será criado automaticamente.

```bash
./re-asmr-spider
```

### Configuração

Edite `config.json` ou use o menu de configuração integrado (opção 2):

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

**Opções de configuração:**

- `account` - Nome de utilizador da conta ASMR.one
- `password` - Palavra-passe da conta ASMR.one
- `max_task` - Número máximo de tarefas de transferência simultâneas
- `max_thread` - Número de threads por transferência de ficheiro
- `max_retry` - Número máximo de tentativas para transferências falhadas
- `language` - Idioma da interface (ver idiomas suportados abaixo)
- `proxy` - Proxy HTTP/HTTPS (por exemplo, `http://127.0.0.1:7890`, deixe vazio para desativar)

### Transferir áudio

1. Execute o programa
2. Selecione a opção `1. Start Download`
3. Introduza o número RJ (por exemplo, `RJ373001`)
4. Para múltiplas transferências, separe os números RJ com espaços (por exemplo, `RJ373001 RJ123456 RJ789012`)

As transferências são guardadas no diretório `downloads/`.

### Retomar transferências

Se uma transferência for interrompida, o programa detetará tarefas incompletas na próxima execução e solicitará que continue.

## Idiomas suportados

- Chinês (Simplificado e Tradicional)
- Inglês
- Japonês
- Alemão
- Espanhol
- Francês
- Português
- Russo
- Indonésio
- Vietnamita
- Hindi
- Bengali
- Telugu
- Turco
- Urdu
- Hausa
- Esperanto
- Linguagem de gato
- Marciano

Mude o idioma através do menu de configuração (opção 4) ou editando `config.json`.

## Fonte de dados

Todo o conteúdo de áudio é proveniente de [asmr.one](https://asmr.one).

## Licença

MIT License

## Agradecimentos

Baseado em [go-asmr-spider](https://github.com/DiheChen/go-asmr-spider) por DiheChen.

## Aviso legal

Esta ferramenta destina-se apenas a uso pessoal. Por favor, respeite os criadores de conteúdo e apoie-os através de canais oficiais.
