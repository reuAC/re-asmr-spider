# re-ASMR-Spider

[简体中文](README.zh-CN.md) | [繁體中文](README.zh-TW.md) | [English](../README.md) | [日本語](README.ja-JP.md) | [Русский](README.ru-RU.md) | [Deutsch](README.de-DE.md) | [Español](README.es-ES.md) | [Français](README.fr-FR.md) | [Português](README.pt-PT.md) | [Indonesia](README.id-ID.md) | [Tiếng Việt](README.vi-VN.md) | [हिन्दी](README.hi-IN.md) | [বাংলা](README.bn-BD.md) | [తెలుగు](README.te-IN.md) | Türkçe | [اردو](README.ur-PK.md) | [Hausa](README.ha-NG.md) | [Esperanto](README.eo.md) | [喵喵语](README.cat.md) | [火星语](README.martian.md)

---

Basit ve verimli bir ASMR ses indirme aracı.

## Özellikler

- Devam etme desteğiyle çok iş parçacıklı indirme
- Ağ hatalarını yönetmek için akıllı yeniden deneme mekanizması
- Gerçek zamanlı ilerleme gösterimi
- Dosya boyutu doğrulama, indirilen içeriği otomatik olarak atlar
- Çoklu dil desteği (20+ dil)
- Çoklu platform desteği (Windows, Linux, macOS, BSD)
- Proxy yapılandırması

## Ekran Görüntüleri
![İndirme Denemesi](./images/asmr-spider-0.png)
![Yeniden Deneme](./images/asmr-spider-1.png)

## Kurulum

### Önceden Derlenmiş İkili Dosyaları İndirin

[Releases](https://github.com/reuAC/re-asmr-spider/releases) sayfasından en son sürümü indirin.

#### Platform Seçim Kılavuzu

| İşletim Sistemi | CPU Mimarisi | Örnek Cihazlar/Senaryolar | İndirme Dosyası |
|-----------------|------------------|---------------------------|---------------|
| **Windows** | | | |
| Windows 11/10/8/7 | 64-bit Intel/AMD | Çoğu modern PC ve dizüstü bilgisayar | `Windows_x86_64.zip` |
| Windows 11/10/8/7 | 32-bit Intel/AMD | Eski PC'ler (2010 öncesi) | `Windows_i386.zip` |
| Windows 11 | ARM64 | Surface Pro X, ARM tabanlı dizüstü bilgisayarlar | `Windows_arm64.zip` |
| **Linux** | | | |
| Ubuntu, Debian, Fedora, CentOS | 64-bit Intel/AMD (x86_64) | Çoğu sunucu ve masaüstü | `Linux_x86_64.tar.gz` |
| Ubuntu, Debian, Fedora | 32-bit Intel/AMD (i386) | Eski 32-bit sistemler | `Linux_i386.tar.gz` |
| Raspberry Pi OS, Ubuntu | ARM64 (aarch64) | Raspberry Pi 3/4/5 (64-bit OS), NVIDIA Jetson | `Linux_arm64.tar.gz` |
| Raspberry Pi OS | ARMv7 (32-bit) | Raspberry Pi 2/3/4 (32-bit OS) | `Linux_armv7.tar.gz` |
| Raspberry Pi OS | ARMv6 (32-bit) | Raspberry Pi 1, Pi Zero | `Linux_armv6.tar.gz` |
| OpenWrt, LEDE | MIPS (big-endian) | Bazı yönlendiriciler (MediaTek, Atheros) | `Linux_mips.tar.gz` |
| OpenWrt, LEDE | MIPS64 (big-endian) | Üst düzey MIPS yönlendiriciler | `Linux_mips64.tar.gz` |
| OpenWrt, LEDE | MIPSle (little-endian) | Bazı yönlendiriciler (Broadcom, Realtek) | `Linux_mipsle.tar.gz` |
| OpenWrt, LEDE | MIPS64le (little-endian) | MIPS64 little-endian sistemler | `Linux_mips64le.tar.gz` |
| Debian, Fedora | RISC-V 64-bit | SiFive boards, StarFive VisionFive | `Linux_riscv64.tar.gz` |
| **macOS** | | | |
| macOS 10.13+ | Intel (x86_64) | MacBook Pro, iMac (2006-2020) | `Darwin_x86_64.tar.gz` |
| macOS 11+ | Apple Silicon (M1/M2/M3) | MacBook Air/Pro M1/M2/M3, Mac Mini M1/M2 | `Darwin_arm64.tar.gz` |
| **BSD** | | | |
| FreeBSD | 64-bit Intel/AMD | FreeBSD sunucular ve masaüstler | `Freebsd_x86_64.tar.gz` |
| FreeBSD | 32-bit Intel/AMD | Eski FreeBSD sistemleri | `Freebsd_i386.tar.gz` |
| FreeBSD | ARM64 | ARM kartlarında FreeBSD | `Freebsd_arm64.tar.gz` |
| OpenBSD | 64-bit Intel/AMD | OpenBSD sistemleri | `Openbsd_x86_64.tar.gz` |
| OpenBSD | 32-bit Intel/AMD | Eski OpenBSD sistemleri | `Openbsd_i386.tar.gz` |
| OpenBSD | ARM64 | ARM kartlarında OpenBSD | `Openbsd_arm64.tar.gz` |
| NetBSD | 64-bit Intel/AMD | NetBSD sistemleri | `Netbsd_x86_64.tar.gz` |
| NetBSD | 32-bit Intel/AMD | Eski NetBSD sistemleri | `Netbsd_i386.tar.gz` |
| NetBSD | ARM64 | ARM kartlarında NetBSD | `Netbsd_arm64.tar.gz` |

#### Sisteminizi Nasıl Kontrol Edersiniz

**Windows:**
- "Bu Bilgisayar"a sağ tıklayın > Özellikler
- "Sistem türü"nü kontrol edin (64-bit veya 32-bit, x64 tabanlı veya ARM64 tabanlı)

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

Veya: Apple menüsü > Bu Mac Hakkında > Çip (M1/M2/M3 = ARM64, Intel = x86_64)

### Kaynaktan Derleyin

```bash
git clone https://github.com/reuAC/re-asmr-spider.git
cd re-asmr-spider
go build -o re-asmr-spider
```

## Kullanım

### İlk Çalıştırma

Yürütülebilir dosyayı çalıştırın. Varsayılan yapılandırma dosyası `config.json` otomatik olarak oluşturulacaktır.

```bash
./re-asmr-spider
```

### Yapılandırma

`config.json` dosyasını düzenleyin veya yerleşik yapılandırma menüsünü (seçenek 2) kullanın:

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

**Yapılandırma Seçenekleri:**

- `account` - ASMR.one hesap kullanıcı adı
- `password` - ASMR.one hesap şifresi
- `max_task` - Maksimum eşzamanlı indirme görevi
- `max_thread` - Dosya başına indirme iş parçacığı sayısı
- `max_retry` - Başarısız indirmeler için maksimum yeniden deneme sayısı
- `language` - Arayüz dili (aşağıdaki desteklenen dillere bakın)
- `proxy` - HTTP/HTTPS proxy (ör., `http://127.0.0.1:7890`, devre dışı bırakmak için boş bırakın)

### Ses İndirin

1. Programı çalıştırın
2. `1. Start Download` seçeneğini seçin
3. RJ numarasını girin (ör., `RJ373001`)
4. Birden fazla indirme için, RJ numaralarını boşluklarla ayırın (ör., `RJ373001 RJ123456 RJ789012`)

İndirmeler `downloads/` dizinine kaydedilir.

### İndirmeleri Devam Ettirin

Bir indirme kesilirse, program bir sonraki çalıştırmada tamamlanmamış görevleri algılayacak ve devam etmenizi isteyecektir.

## Desteklenen Diller

- Çince (Basitleştirilmiş ve Geleneksel)
- İngilizce
- Japonca
- Almanca
- İspanyolca
- Fransızca
- Portekizce
- Rusça
- Endonezce
- Vietnamca
- Hintçe
- Bengalce
- Telugu
- Türkçe
- Urduca
- Hausaca
- Esperanto
- Kedi Dili
- Marslı

Yapılandırma menüsü (seçenek 4) aracılığıyla veya `config.json` dosyasını düzenleyerek dili değiştirin.

## Veri Kaynağı

Tüm ses içeriği [asmr.one](https://asmr.one) adresinden alınmıştır.

## Lisans

MIT License

## Teşekkürler

DiheChen'in [go-asmr-spider](https://github.com/DiheChen/go-asmr-spider) projesine dayanmaktadır.

## Feragatname

Bu araç yalnızca kişisel kullanım içindir. Lütfen içerik üreticilerine saygı gösterin ve onları resmi kanallar aracılığıyla destekleyin.
