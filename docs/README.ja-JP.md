# re-ASMR-Spider

[简体中文](README.zh-CN.md) | [繁體中文](README.zh-TW.md) | [English](../README.md) | 日本語

---

シンプルで効率的なASMR音声ダウンロードツール。

## 機能

- レジューム対応のマルチスレッドダウンロード
- ネットワークエラーを処理するスマートリトライメカニズム
- リアルタイム進捗表示
- ファイルサイズ検証、ダウンロード済みコンテンツを自動スキップ
- 多言語サポート
- マルチプラットフォームサポート（Windows、Linux、macOS、BSD）
- プロキシ設定
- カスタムフィルタリング
- コマンドラインモード

## スクリーンショット
![ダウンロード試行](./images/asmr-spider-0.png)
![再試行ダウンロード](./images/asmr-spider-1.png)
![カスタムフィルタリング](./images/asmr-spider-3.png)

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

### インタラクティブモード（デフォルト）

実行ファイルを実行してください。デフォルトの設定ファイル `config.json` が自動的に作成されます。

```bash
./re-asmr-spider
```

1. オプション `1. ダウンロード開始` を選択
2. RJ番号を入力（例：`RJ373001`）
3. 複数ダウンロードの場合は、RJ番号をスペースで区切って入力（例：`RJ373001 RJ123456 RJ789012`）

ダウンロードは `downloads/` ディレクトリに保存されます。

### コマンドラインモード

自動化やスクリプト実行のため、インタラクティブメニューなしで直接ダウンロード：

```bash
# 単一のRJ番号をダウンロード
./re-asmr-spider -download RJ123456

# 複数のRJ番号をダウンロード
./re-asmr-spider -download RJ123456,RJ789012,RJ345678

# カスタムアカウント認証情報を使用
./re-asmr-spider -download RJ123456 -account user@example.com -password mypass

# ダウンロードパラメータを設定
./re-asmr-spider -download RJ123456 -max-task 5 -max-thread 16 -buffer-size 16

# プロキシを使用
./re-asmr-spider -download RJ123456 -proxy http://127.0.0.1:7890

# カスタム設定ファイルを使用
./re-asmr-spider -config /path/to/config.json -download RJ123456

# ダウンロード優先順位
# flacに関する競合が発生した場合、flacのみダウンロード
# 例えば flac wav mp3 ogg の4つの形式で同名ファイルが存在する場合、flacのみダウンロード。
# ただし flac がない場合は wav をダウンロード。どちらもない場合は mp3 と ogg をダウンロード
./re-asmr-spider -download RJ123456 -format-priority flac,wav

# flacを優先的にダウンロードし、さらに lrc 字幕ファイルをダウンロード
./re-asmr-spider -download RJ123456 -format-priority flac,wav -include-formats lrc
```

**コマンドラインオプション：**

```
-download string      ダウンロードするRJ番号（カンマ区切り）
-config string        設定ファイルのパス（デフォルト: config.json）
-account string       ASMR.oneアカウントのユーザー名（設定を上書き）
-password string      ASMR.oneアカウントのパスワード（設定を上書き）
-max-task int         最大同時ダウンロードタスク数
-max-thread int       ファイルダウンロードあたりのスレッド数
-max-retry int        ダウンロード失敗時の最大リトライ回数
-buffer-size int      バッファサイズ（MB）、範囲 1-64、デフォルト 8
-proxy string         HTTP/HTTPSプロキシ（例：http://127.0.0.1:7890）
-version              バージョン情報を表示
-help                 ヘルプメッセージを表示
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
  "proxy": "",
  "buffer_size_mb": 8
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
- `buffer_size_mb` - ダウンロードバッファサイズ（MB）、範囲 1-64、デフォルト 8、VPS向けに最適化

### フォーマットフィルタリング

音声作品をダウンロードする際、同名で異なる形式のファイル（`.wav`、`.flac`、`.mp3` など）に遭遇することがあります。プログラムはインテリジェントなフォーマットフィルタリング機能を提供します：

**インタラクティブモード：**
- 形式の競合を自動検出
- ユーザーに形式を一つずつ選択するよう促します
- オプションには以下が含まれます：
  - すべての形式をダウンロード
  - 特定の形式を選択
  - 残りのすべての類似ファイルに選択を適用（バッチモード）
  - 残りのすべてのファイルですべての形式をダウンロード
- 形式の選択は保存され、ダウンロードが中断されても再開できます

### ダウンロードの再開

ダウンロードが中断された場合、次回実行時にプログラムが未完了のタスクを検出し、続行するかどうか確認します。形式選択の設定は保持されます。

## サポート言語

- 简体中文
- 繁體中文
- English
- 日本語

設定メニュー（オプション4）または `config.json` の編集で言語を切り替えられます。

## データソース

すべての音声コンテンツは [asmr.one](https://asmr.one) から取得しています。

## ライセンス

MIT License

## 謝辞

DiheChenの [go-asmr-spider](https://github.com/DiheChen/go-asmr-spider) をベースにしています。

## 免責事項

このツールは個人使用のみを目的としています。コンテンツクリエイターを尊重し、公式チャンネルを通じてサポートしてください。
