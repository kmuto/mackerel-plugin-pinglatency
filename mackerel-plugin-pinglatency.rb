#!/usr/bin/env ruby
# Copyright 2024 Kenshi Muto
#
# mackerel-ping-latency.rb
# pingを使ったパケットロス（%）、レイテンシー（min, max, avg, stddev）をMackerelのグラフにする
#
# 登録方法
# mackerel-agent.confに以下を記載する
#
# [plugin.metrics.対象ホスト名]
# command = ['置いた場所/mackerel-ping-latency.rb', '-H', '対象ホスト名']
# -cオプションでpingカウントを指定可能です（デフォルトは3）
#
require 'open3'
require 'optparse'
require 'json'

# グラフ定義作成
if ENV['MACKEREL_AGENT_PLUGIN_META'] == '1'
  meta = {
    :graphs => {
      'ping.loss' => {
        :label => 'Ping Loss',
        :unit => 'percentage',
        :metrics => [
          {
            :name => 'percent',
            :label => 'percent'
          }
        ]
      },
      'ping.latency' => {
        :label => 'Ping Latency',
        :unit => 'milliseconds',
        :metrics => [
          {
            :name => 'rttmin',
            :label => 'Min'
          },
          {
            :name => 'rttmax',
            :label => 'Max'
          },
          {
            :name => 'rttavg',
            :label => 'Average'
          },
          {
            :name => 'rttstddev',
            :label => 'STDDev'
          }
        ]
      }
    }
  }
  puts '# mackerel-agent-plugin'
  puts meta.to_json
  exit 0
end

# パラメータ
count = 3
host = 'localhost'

opt = OptionParser.new
opt.on('-c PINGCOUNT') { |v| count = v }
opt.on('-H HOSTNAME') { |v| host = v }
opt.parse!(ARGV)

o, status = Open3.capture2e("ping -c #{count} -q #{host}")
unless status.success?
  STDERR.puts o
  exit 1
end

# 結果パース
# PING localhost (127.0.0.1): 56 data bytes
# --- localhost ping statistics ---
# 5 packets transmitted, 5 packets received, 0.0% packet loss
# round-trip min/avg/max/stddev = 0.062/0.193/0.290/0.086 ms

info = {}

o.each_line do |l|
  if l =~ /([\d\.]+)% packet loss/
    info[:loss] = $1.to_f
  elsif l =~ /= ([\d\.]+)\/([\d\.]+)\/([\d\.]+)\/([\d\.]+) ms/
    info[:rttmin] = $1.to_f
    info[:rttavg] = $2.to_f
    info[:rttmax] = $3.to_f
    info[:rttstddev] = $4.to_f
  end
end

# 結果出力
epoch = Time.now.to_i

puts <<EOT
ping.loss.percent\t#{info[:loss]}\t#{epoch}
ping.latency.rttmin\t#{info[:rttmin]}\t#{epoch}
ping.latency.rttavg\t#{info[:rttavg]}\t#{epoch}
ping.latency.rttmax\t#{info[:rttmax]}\t#{epoch}
ping.latency.rttstddev\t#{info[:rttstddev]}\t#{epoch}
EOT
