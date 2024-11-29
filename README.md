# mackerel-plugin-pinglatency

[日本語](#説明)

## Description
**mackerel-plugin-pinglatency** shows ping packet loss and ping latencies on Mackerel graph.

## Synopsis
```
mackerel-plugin-pinglatency HOST1 HOST2 ...
```

will export metrics

```
ping.latency.HOST1.avg
ping.latency.HOST1.min
ping.latency.HOST1.max
ping.packet_loss.HOST1_packet_loss
ping.latency.HOST2.avg
ping.latency.HOST2.min
ping.latency.HOST2.max
ping.packet_loss.HOST2_packet_loss
...
```

## Usage
### Installation
On Linux:
```
sudo mkr plugin install kmuto/mackerel-plugin-pinglatency
```

On Windows (as Administrator):
```
"C:\Program Files\Mackerel\mackerel-agent\mkr.exe" plugin install kmuto/mackerel-plugin-pinglatency
```

### Definition in mackerel-agent.conf
On Linux:
```
[plugin.metrics.pinglatency]
command = ["/opt/mackerel-agent/plugins/bin/mackerel-plugin-pinglatency", "HOST"]
```

On Windows:
```
[plugin.metrics.pinglatency]
command = ["plugins\\bin\\mackerel-plugin-pinglatency.exe", "HOST"]
```

### Options
- `-c COUNT`: Number of ping packets (default: 5)
- `-t SECONDS`: Timeout seconds for ping (default: 15)
- `-V`: Verbose mode
- `-v`: Show version

## License
© 2024 Kenshi Muto

MIT License (see LICENSE file)

---

## 説明
**mackerel-plugin-pinglatency** は、pingパケットおよびpingレイテンシーをMackerelのグラフとして表示します。

## 概要
```
mackerel-plugin-pinglatency HOST1 HOST2 ...
```

以下のメトリックが書き出されます。

```
ping.latency.HOST1.avg
ping.latency.HOST1.min
ping.latency.HOST1.max
ping.packet_loss.HOST1_packet_loss
ping.latency.HOST2.avg
ping.latency.HOST2.min
ping.latency.HOST2.max
ping.packet_loss.HOST2_packet_loss
...
```

## 使い方
### プラグインのインストール
Linuxの場合:
```
sudo mkr plugin install kmuto/mackerel-plugin-pinglatency
```

Windowsの場合 (管理者権限):
```
"C:\Program Files\Mackerel\mackerel-agent\mkr.exe" plugin install kmuto/mackerel-plugin-pinglatency
```

### mackerel-agent.confでの設定
Linuxの場合:
```
[plugin.metrics.pinglatency]
command = ["/opt/mackerel-agent/plugins/bin/mackerel-plugin-pinglatency", "HOST"]
```

Windowsの場合:
```
[plugin.metrics.pinglatency]
command = ["plugins\\bin\\mackerel-plugin-pinglatency.exe", "HOST"]
```

### オプション
- `-c COUNT`: pingパケットの試行回数 (デフォルト: 5)
- `-t SECONDS`: pingのタイムアウト秒 (デフォルト: 15)
- `-V`: 冗長モード
- `-v`: バージョン表示

## ライセンス
© 2024 Kenshi Muto

MIT License (LICENSEファイルを参照)
