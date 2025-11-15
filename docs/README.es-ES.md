# re-ASMR-Spider

[简体中文](README.zh-CN.md) | [繁體中文](README.zh-TW.md) | [English](../README.md) | [日本語](README.ja-JP.md) | [Русский](README.ru-RU.md) | [Deutsch](README.de-DE.md) | Español | [Français](README.fr-FR.md) | [Português](README.pt-PT.md) | [Indonesia](README.id-ID.md) | [Tiếng Việt](README.vi-VN.md) | [हिन्दी](README.hi-IN.md) | [বাংলা](README.bn-BD.md) | [తెలుగు](README.te-IN.md) | [Türkçe](README.tr-TR.md) | [اردو](README.ur-PK.md) | [Hausa](README.ha-NG.md) | [Esperanto](README.eo.md) | [喵喵语](README.cat.md) | [火星语](README.martian.md)

---

Una herramienta simple y eficiente para descargar audio ASMR.

## Características

- Descarga multiproceso con soporte para reanudación
- Mecanismo de reintento inteligente para manejar errores de red
- Visualización del progreso en tiempo real
- Verificación del tamaño del archivo, omite automáticamente el contenido descargado
- Soporte multiidioma (más de 20 idiomas)
- Soporte multiplataforma (Windows, Linux, macOS, BSD)
- Configuración de proxy

## Capturas de pantalla
![Intento de descarga](./images/asmr-spider-0.png)
![Reintentar descarga](./images/asmr-spider-1.png)

## Instalación

### Descargar binarios precompilados

Descargue la última versión desde la página de [Releases](https://github.com/reuAC/re-asmr-spider/releases).

#### Guía de selección de plataforma

| Sistema operativo | Arquitectura de CPU | Dispositivos/Escenarios de ejemplo | Archivo de descarga |
|-----------------|------------------|---------------------------|---------------|
| **Windows** | | | |
| Windows 11/10/8/7 | 64-bit Intel/AMD | La mayoría de PCs y portátiles modernos | `Windows_x86_64.zip` |
| Windows 11/10/8/7 | 32-bit Intel/AMD | PCs antiguos (anteriores a 2010) | `Windows_i386.zip` |
| Windows 11 | ARM64 | Surface Pro X, portátiles basados en ARM | `Windows_arm64.zip` |
| **Linux** | | | |
| Ubuntu, Debian, Fedora, CentOS | 64-bit Intel/AMD (x86_64) | La mayoría de servidores y escritorios | `Linux_x86_64.tar.gz` |
| Ubuntu, Debian, Fedora | 32-bit Intel/AMD (i386) | Sistemas antiguos de 32 bits | `Linux_i386.tar.gz` |
| Raspberry Pi OS, Ubuntu | ARM64 (aarch64) | Raspberry Pi 3/4/5 (SO 64 bits), NVIDIA Jetson | `Linux_arm64.tar.gz` |
| Raspberry Pi OS | ARMv7 (32-bit) | Raspberry Pi 2/3/4 (SO 32 bits) | `Linux_armv7.tar.gz` |
| Raspberry Pi OS | ARMv6 (32-bit) | Raspberry Pi 1, Pi Zero | `Linux_armv6.tar.gz` |
| OpenWrt, LEDE | MIPS (big-endian) | Algunos routers (MediaTek, Atheros) | `Linux_mips.tar.gz` |
| OpenWrt, LEDE | MIPS64 (big-endian) | Routers MIPS de gama alta | `Linux_mips64.tar.gz` |
| OpenWrt, LEDE | MIPSle (little-endian) | Algunos routers (Broadcom, Realtek) | `Linux_mipsle.tar.gz` |
| OpenWrt, LEDE | MIPS64le (little-endian) | Sistemas MIPS64 little-endian | `Linux_mips64le.tar.gz` |
| Debian, Fedora | RISC-V 64-bit | Placas SiFive, StarFive VisionFive | `Linux_riscv64.tar.gz` |
| **macOS** | | | |
| macOS 10.13+ | Intel (x86_64) | MacBook Pro, iMac (2006-2020) | `Darwin_x86_64.tar.gz` |
| macOS 11+ | Apple Silicon (M1/M2/M3) | MacBook Air/Pro M1/M2/M3, Mac Mini M1/M2 | `Darwin_arm64.tar.gz` |
| **BSD** | | | |
| FreeBSD | 64-bit Intel/AMD | Servidores y escritorios FreeBSD | `Freebsd_x86_64.tar.gz` |
| FreeBSD | 32-bit Intel/AMD | Sistemas FreeBSD antiguos | `Freebsd_i386.tar.gz` |
| FreeBSD | ARM64 | FreeBSD en placas ARM | `Freebsd_arm64.tar.gz` |
| OpenBSD | 64-bit Intel/AMD | Sistemas OpenBSD | `Openbsd_x86_64.tar.gz` |
| OpenBSD | 32-bit Intel/AMD | Sistemas OpenBSD antiguos | `Openbsd_i386.tar.gz` |
| OpenBSD | ARM64 | OpenBSD en placas ARM | `Openbsd_arm64.tar.gz` |
| NetBSD | 64-bit Intel/AMD | Sistemas NetBSD | `Netbsd_x86_64.tar.gz` |
| NetBSD | 32-bit Intel/AMD | Sistemas NetBSD antiguos | `Netbsd_i386.tar.gz` |
| NetBSD | ARM64 | NetBSD en placas ARM | `Netbsd_arm64.tar.gz` |

#### Cómo verificar su sistema

**Windows:**
- Haga clic derecho en "Este equipo" > Propiedades
- Verifique "Tipo de sistema" (64 bits o 32 bits, basado en x64 o basado en ARM64)

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

O: menú Apple > Acerca de este Mac > Chip (M1/M2/M3 = ARM64, Intel = x86_64)

### Compilar desde el código fuente

```bash
git clone https://github.com/reuAC/re-asmr-spider.git
cd re-asmr-spider
go build -o re-asmr-spider
```

## Uso

### Primera ejecución

Ejecute el archivo ejecutable. Se creará automáticamente un archivo de configuración predeterminado `config.json`.

```bash
./re-asmr-spider
```

### Configuración

Edite `config.json` o use el menú de configuración integrado (opción 2):

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

**Opciones de configuración:**

- `account` - Nombre de usuario de la cuenta ASMR.one
- `password` - Contraseña de la cuenta ASMR.one
- `max_task` - Número máximo de tareas de descarga simultáneas
- `max_thread` - Número de hilos por descarga de archivo
- `max_retry` - Número máximo de reintentos para descargas fallidas
- `language` - Idioma de la interfaz (ver idiomas compatibles a continuación)
- `proxy` - Proxy HTTP/HTTPS (por ejemplo, `http://127.0.0.1:7890`, déjelo vacío para deshabilitarlo)

### Descargar audio

1. Ejecute el programa
2. Seleccione la opción `1. Start Download`
3. Ingrese el número RJ (por ejemplo, `RJ373001`)
4. Para múltiples descargas, separe los números RJ con espacios (por ejemplo, `RJ373001 RJ123456 RJ789012`)

Las descargas se guardan en el directorio `downloads/`.

### Reanudar descargas

Si una descarga se interrumpe, el programa detectará las tareas incompletas en la siguiente ejecución y le solicitará continuar.

## Idiomas compatibles

- Chino (Simplificado y Tradicional)
- Inglés
- Japonés
- Alemán
- Español
- Francés
- Portugués
- Ruso
- Indonesio
- Vietnamita
- Hindi
- Bengalí
- Telugu
- Turco
- Urdu
- Hausa
- Esperanto
- Idioma de gato
- Marciano

Cambie el idioma a través del menú de configuración (opción 4) o editando `config.json`.

## Fuente de datos

Todo el contenido de audio proviene de [asmr.one](https://asmr.one).

## Licencia

MIT License

## Agradecimientos

Basado en [go-asmr-spider](https://github.com/DiheChen/go-asmr-spider) de DiheChen.

## Descargo de responsabilidad

Esta herramienta es solo para uso personal. Por favor, respete a los creadores de contenido y apóyelos a través de canales oficiales.
