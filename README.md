# mackerel-plugin-pinglatency

[日本語](#説明)

## Description
Show ping packet loss and ping latencies on Mackerel graph.

## Synopsis
```
mackerel-plugin-pinglatency HOST1 HOST2 ...
```

will export metrics

```
ping.latency.HOST1.avg
ping.latency.HOST1.min
ping.latency.HOST1.max
ping.latency.HOST1.stddev
ping.packet_loss.HOST1.packet_loss
ping.latency.HOST2.avg
ping.latency.HOST2.min
ping.latency.HOST2.max
ping.latency.HOST2.stddev
ping.packet_loss.HOST2.packet_loss
...
```

## Usage
### Installation
```
mkr plugin install kmuto/mackerel-plugin-pinglatency
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
- `-c COUNT`: Number of ping packets (default: 3)
- `-t SECONDS`: Timeout seconds for each ping packet (default: 5)
- `-v`: Verbose mode

## License
© 2024 Kenshi Muto

MIT License (see LICENSE file)

---

## 説明
pingパケットおよびpingレイテンシーをMackerelのグラフとして表示します。

## 概要
```
mackerel-plugin-pinglatency HOST1 HOST2 ...
```

以下のメトリックが書き出されます。

```
ping.latency.HOST1.avg
ping.latency.HOST1.min
ping.latency.HOST1.max
ping.latency.HOST1.stddev
ping.packet_loss.HOST1.packet_loss
ping.latency.HOST2.avg
ping.latency.HOST2.min
ping.latency.HOST2.max
ping.latency.HOST2.stddev
ping.packet_loss.HOST2.packet_loss
...
```

## 使い方
### プラグインのインストール
```
mkr plugin install kmuto/mackerel-plugin-pinglatency
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
- `-c COUNT`: pingパケットの試行回数 (デフォルト: 3)
- `-t SECONDS`: pingパケットのタイムアウト秒 (デフォルト: 5)
- `-v`: 冗長モード

## ライセンス
© 2024 Kenshi Muto

MIT License (LICENSEファイルを参照)
