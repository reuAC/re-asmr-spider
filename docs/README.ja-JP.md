# re-ASMR-Spider

[简体中文](README.zh-CN.md) | [繁體中文](README.zh-TW.md) | [English](../README.md) | 日本語

---

シンプルで効率的なASMR音声ダウンロードツール。

## 機能

- レジューム対応のマルチスレッドダウンロード
- ネットワークエラーを処理するスマートリトライメカニズム
- リアルタイム進捗表示
- ファイルサイズ検証、ダウンロード済みコンテンツを自動スキップ
- 多言語サポート（20言語以上）
- マルチプラットフォームサポート（Windows、Linux、macOS、BSD）
- プロキシ設定

## スクリーンショット
![ダウンロード試行](./images/asmr-spider-0.png)
![再試行ダウンロード](./images/asmr-spider-1.png)

## インストール

### ビルド済みバイナリのダウンロード

[リリース](https://github.com/reuAC/re-asmr-spider/releases)ページから最新リリースをダウンロードしてください。

#### プラットフォーム選択ガイド

| オペレーティングシステム | CPUアーキテクチャ | デバイス例/シナリオ | ダウンロードファイル |
|-----------------|------------------|---------------------------|---------------|
| **Windows** | | | |
| Windows 11/10/8/7 | 64ビット Intel/AMD | 最新のPCやノートパソコン | `Windows_x86_64.zip` |
| Windows 11/10/8/7 | 32ビット Intel/AMD | 古いPC（2010年以前） | `Windows_i386.zip` |
| Windows 11 | ARM64 | Surface Pro X、ARMベースのノートパソコン | `Windows_arm64.zip` |
| **Linux** | | | |
| Ubuntu、Debian、Fedora、CentOS | 64ビット Intel/AMD (x86_64) | ほとんどのサーバーとデスクトップ | `Linux_x86_64.tar.gz` |
| Ubuntu、Debian、Fedora | 32ビット Intel/AMD (i386) | 古い32ビットシステム | `Linux_i386.tar.gz` |
| Raspberry Pi OS、Ubuntu | ARM64 (aarch64) | Raspberry Pi 3/4/5（64ビットOS）、NVIDIA Jetson | `Linux_arm64.tar.gz` |
| Raspberry Pi OS | ARMv7（32ビット） | Raspberry Pi 2/3/4（32ビットOS） | `Linux_armv7.tar.gz` |
| Raspberry Pi OS | ARMv6（32ビット） | Raspberry Pi 1、Pi Zero | `Linux_armv6.tar.gz` |
| OpenWrt、LEDE | MIPS（ビッグエンディアン） | 一部のルーター（MediaTek、Atheros） | `Linux_mips.tar.gz` |
| OpenWrt、LEDE | MIPS64（ビッグエンディアン） | ハイエンドMIPSルーター | `Linux_mips64.tar.gz` |
| OpenWrt、LEDE | MIPSle（リトルエンディアン） | 一部のルーター（Broadcom、Realtek） | `Linux_mipsle.tar.gz` |
| OpenWrt、LEDE | MIPS64le（リトルエンディアン） | MIPS64リトルエンディアンシステム | `Linux_mips64le.tar.gz` |
| Debian、Fedora | RISC-V 64ビット | SiFiveボード、StarFive VisionFive | `Linux_riscv64.tar.gz` |
| **macOS** | | | |
| macOS 10.13以降 | Intel (x86_64) | MacBook Pro、iMac（2006-2020） | `Darwin_x86_64.tar.gz` |
| macOS 11以降 | Apple Silicon (M1/M2/M3) | MacBook Air/Pro M1/M2/M3、Mac Mini M1/M2 | `Darwin_arm64.tar.gz` |
| **BSD** | | | |
| FreeBSD | 64ビット Intel/AMD | FreeBSDサーバーとデスクトップ | `Freebsd_x86_64.tar.gz` |
| FreeBSD | 32ビット Intel/AMD | 古いFreeBSDシステム | `Freebsd_i386.tar.gz` |
| FreeBSD | ARM64 | ARMボード上のFreeBSD | `Freebsd_arm64.tar.gz` |
| OpenBSD | 64ビット Intel/AMD | OpenBSDシステム | `Openbsd_x86_64.tar.gz` |
| OpenBSD | 32ビット Intel/AMD | 古いOpenBSDシステム | `Openbsd_i386.tar.gz` |
| OpenBSD | ARM64 | ARMボード上のOpenBSD | `Openbsd_arm64.tar.gz` |
| NetBSD | 64ビット Intel/AMD | NetBSDシステム | `Netbsd_x86_64.tar.gz` |
| NetBSD | 32ビット Intel/AMD | 古いNetBSDシステム | `Netbsd_i386.tar.gz` |
| NetBSD | ARM64 | ARMボード上のNetBSD | `Netbsd_arm64.tar.gz` |

#### システムの確認方法

**Windows:**
- 「PC」を右クリック > プロパティ
- 「システムの種類」を確認（64ビットまたは32ビット、x64ベースまたはARM64ベース）

**Linux:**
```bash
uname -m
# x86_64 -> 64ビット Intel/AMD
# i386/i686 -> 32ビット Intel/AMD
# aarch64 -> ARM64
# armv7l -> ARMv7
# armv6l -> ARMv6
# mips/mips64 -> MIPS
# riscv64 -> RISC-V 64ビット
```

**macOS:**
```bash
uname -m
# x86_64 -> Intel Mac
# arm64 -> Apple Silicon (M1/M2/M3)
```

または：Appleメニュー > このMacについて > チップ（M1/M2/M3 = ARM64、Intel = x86_64）

### ソースからビルド

```bash
git clone https://github.com/reuAC/re-asmr-spider.git
cd re-asmr-spider
go build -o re-asmr-spider
```

## 使い方

### 初回実行

実行ファイルを実行してください。デフォルトの設定ファイル `config.json` が自動的に作成されます。

```bash
./re-asmr-spider
```

### 設定

`config.json` を編集するか、組み込みの設定メニュー（オプション2）を使用してください：

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

**設定オプション：**

- `account` - ASMR.oneアカウントのユーザー名
- `password` - ASMR.oneアカウントのパスワード
- `max_task` - 最大同時ダウンロードタスク数
- `max_thread` - ファイルダウンロードあたりのスレッド数
- `max_retry` - ダウンロード失敗時の最大リトライ回数
- `language` - インターフェース言語（以下のサポート言語を参照）
- `proxy` - HTTP/HTTPSプロキシ（例：`http://127.0.0.1:7890`、無効にする場合は空のまま）

### 音声のダウンロード

1. プログラムを実行
2. オプション `1. ダウンロード開始` を選択
3. RJ番号を入力（例：`RJ373001`）
4. 複数ダウンロードの場合は、RJ番号をスペースで区切って入力（例：`RJ373001 RJ123456 RJ789012`）

ダウンロードは `downloads/` ディレクトリに保存されます。

### ダウンロードの再開

ダウンロードが中断された場合、次回実行時にプログラムが未完了のタスクを検出し、続行するかどうか確認します。

## サポート言語

- 中国語（簡体字・繁体字）
- 英語
- 日本語
- ドイツ語
- スペイン語
- フランス語
- ポルトガル語
- ロシア語
- インドネシア語
- ベトナム語
- ヒンディー語
- ベンガル語
- テルグ語
- トルコ語
- ウルドゥー語
- ハウサ語
- エスペラント語
- 猫語
- 火星語

設定メニュー（オプション4）または `config.json` の編集で言語を切り替えられます。

## データソース

すべての音声コンテンツは [asmr.one](https://asmr.one) から取得しています。

## ライセンス

MIT License

## 謝辞

DiheChenの [go-asmr-spider](https://github.com/DiheChen/go-asmr-spider) をベースにしています。

## 免責事項

このツールは個人使用のみを目的としています。コンテンツクリエイターを尊重し、公式チャンネルを通じてサポートしてください。
