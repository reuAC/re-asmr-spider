# re-ASMR-Spider

[简体中文](README.zh-CN.md) | [繁體中文](README.zh-TW.md) | [English](../README.md) | [日本語](README.ja-JP.md) | Русский | [Deutsch](README.de-DE.md) | [Español](README.es-ES.md) | [Français](README.fr-FR.md) | [Português](README.pt-PT.md) | [Indonesia](README.id-ID.md) | [Tiếng Việt](README.vi-VN.md) | [हिन्दी](README.hi-IN.md) | [বাংলা](README.bn-BD.md) | [తెలుగు](README.te-IN.md) | [Türkçe](README.tr-TR.md) | [اردو](README.ur-PK.md) | [Hausa](README.ha-NG.md) | [Esperanto](README.eo.md) | [喵喵语](README.cat.md) | [火星语](README.martian.md)

---

Простой и эффективный инструмент для загрузки ASMR аудио.

## Возможности

- Многопоточная загрузка с поддержкой возобновления
- Умный механизм повторных попыток для обработки сетевых ошибок
- Отображение прогресса в реальном времени
- Проверка размера файла, автоматический пропуск загруженного контента
- Поддержка нескольких языков (более 20 языков)
- Поддержка нескольких платформ (Windows, Linux, macOS, BSD)
- Настройка прокси

## Скриншоты
![Попытка загрузки](./images/asmr-spider-0.png)
![Повторная загрузка](./images/asmr-spider-1.png)

## Установка

### Загрузка готовых исполняемых файлов

Загрузите последнюю версию со страницы [Releases](https://github.com/reuAC/re-asmr-spider/releases).

#### Руководство по выбору платформы

| Операционная система | Архитектура процессора | Примеры устройств/сценариев | Файл для загрузки |
|-----------------|------------------|---------------------------|---------------|
| **Windows** | | | |
| Windows 11/10/8/7 | 64-битный Intel/AMD | Большинство современных ПК и ноутбуков | `Windows_x86_64.zip` |
| Windows 11/10/8/7 | 32-битный Intel/AMD | Старые ПК (до 2010 года) | `Windows_i386.zip` |
| Windows 11 | ARM64 | Surface Pro X, ноутбуки на ARM | `Windows_arm64.zip` |
| **Linux** | | | |
| Ubuntu, Debian, Fedora, CentOS | 64-битный Intel/AMD (x86_64) | Большинство серверов и настольных систем | `Linux_x86_64.tar.gz` |
| Ubuntu, Debian, Fedora | 32-битный Intel/AMD (i386) | Старые 32-битные системы | `Linux_i386.tar.gz` |
| Raspberry Pi OS, Ubuntu | ARM64 (aarch64) | Raspberry Pi 3/4/5 (64-битная ОС), NVIDIA Jetson | `Linux_arm64.tar.gz` |
| Raspberry Pi OS | ARMv7 (32-битный) | Raspberry Pi 2/3/4 (32-битная ОС) | `Linux_armv7.tar.gz` |
| Raspberry Pi OS | ARMv6 (32-битный) | Raspberry Pi 1, Pi Zero | `Linux_armv6.tar.gz` |
| OpenWrt, LEDE | MIPS (big-endian) | Некоторые роутеры (MediaTek, Atheros) | `Linux_mips.tar.gz` |
| OpenWrt, LEDE | MIPS64 (big-endian) | Высокопроизводительные MIPS роутеры | `Linux_mips64.tar.gz` |
| OpenWrt, LEDE | MIPSle (little-endian) | Некоторые роутеры (Broadcom, Realtek) | `Linux_mipsle.tar.gz` |
| OpenWrt, LEDE | MIPS64le (little-endian) | MIPS64 little-endian системы | `Linux_mips64le.tar.gz` |
| Debian, Fedora | RISC-V 64-битный | Платы SiFive, StarFive VisionFive | `Linux_riscv64.tar.gz` |
| **macOS** | | | |
| macOS 10.13+ | Intel (x86_64) | MacBook Pro, iMac (2006-2020) | `Darwin_x86_64.tar.gz` |
| macOS 11+ | Apple Silicon (M1/M2/M3) | MacBook Air/Pro M1/M2/M3, Mac Mini M1/M2 | `Darwin_arm64.tar.gz` |
| **BSD** | | | |
| FreeBSD | 64-битный Intel/AMD | Серверы и настольные системы FreeBSD | `Freebsd_x86_64.tar.gz` |
| FreeBSD | 32-битный Intel/AMD | Старые системы FreeBSD | `Freebsd_i386.tar.gz` |
| FreeBSD | ARM64 | FreeBSD на ARM платах | `Freebsd_arm64.tar.gz` |
| OpenBSD | 64-битный Intel/AMD | Системы OpenBSD | `Openbsd_x86_64.tar.gz` |
| OpenBSD | 32-битный Intel/AMD | Старые системы OpenBSD | `Openbsd_i386.tar.gz` |
| OpenBSD | ARM64 | OpenBSD на ARM платах | `Openbsd_arm64.tar.gz` |
| NetBSD | 64-битный Intel/AMD | Системы NetBSD | `Netbsd_x86_64.tar.gz` |
| NetBSD | 32-битный Intel/AMD | Старые системы NetBSD | `Netbsd_i386.tar.gz` |
| NetBSD | ARM64 | NetBSD на ARM платах | `Netbsd_arm64.tar.gz` |

#### Как проверить вашу систему

**Windows:**
- Щелкните правой кнопкой мыши "Этот компьютер" > Свойства
- Проверьте "Тип системы" (64-разрядная или 32-разрядная, x64 или ARM64)

**Linux:**
```bash
uname -m
# x86_64 -> 64-битный Intel/AMD
# i386/i686 -> 32-битный Intel/AMD
# aarch64 -> ARM64
# armv7l -> ARMv7
# armv6l -> ARMv6
# mips/mips64 -> MIPS
# riscv64 -> RISC-V 64-битный
```

**macOS:**
```bash
uname -m
# x86_64 -> Intel Mac
# arm64 -> Apple Silicon (M1/M2/M3)
```

Или: Меню Apple > Об этом Mac > Чип (M1/M2/M3 = ARM64, Intel = x86_64)

### Сборка из исходного кода

```bash
git clone https://github.com/reuAC/re-asmr-spider.git
cd re-asmr-spider
go build -o re-asmr-spider
```

## Использование

### Первый запуск

Запустите исполняемый файл. Файл конфигурации по умолчанию `config.json` будет создан автоматически.

```bash
./re-asmr-spider
```

### Конфигурация

Отредактируйте `config.json` или используйте встроенное меню конфигурации (опция 2):

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

**Параметры конфигурации:**

- `account` - Имя пользователя аккаунта ASMR.one
- `password` - Пароль аккаунта ASMR.one
- `max_task` - Максимальное количество одновременных задач загрузки
- `max_thread` - Количество потоков для загрузки одного файла
- `max_retry` - Максимальное количество повторных попыток для неудачных загрузок
- `language` - Язык интерфейса (см. поддерживаемые языки ниже)
- `proxy` - HTTP/HTTPS прокси (например, `http://127.0.0.1:7890`, оставьте пустым для отключения)

### Загрузка аудио

1. Запустите программу
2. Выберите опцию `1. Start Download`
3. Введите номер RJ (например, `RJ373001`)
4. Для нескольких загрузок разделите номера RJ пробелами (например, `RJ373001 RJ123456 RJ789012`)

Загрузки сохраняются в каталог `downloads/`.

### Возобновление загрузок

Если загрузка прервана, программа обнаружит незавершенные задачи при следующем запуске и предложит продолжить.

## Поддерживаемые языки

- Китайский (упрощенный и традиционный)
- Английский
- Японский
- Немецкий
- Испанский
- Французский
- Португальский
- Русский
- Индонезийский
- Вьетнамский
- Хинди
- Бенгальский
- Телугу
- Турецкий
- Урду
- Хауса
- Эсперанто
- Кошачий язык
- Марсианский

Переключите язык через меню конфигурации (опция 4) или отредактировав `config.json`.

## Источник данных

Весь аудио контент получен с [asmr.one](https://asmr.one).

## Лицензия

MIT License

## Благодарности

Основано на [go-asmr-spider](https://github.com/DiheChen/go-asmr-spider) от DiheChen.

## Отказ от ответственности

Этот инструмент предназначен только для личного использования. Пожалуйста, уважайте создателей контента и поддерживайте их через официальные каналы.
