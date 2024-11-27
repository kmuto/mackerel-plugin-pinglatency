package pinglatency

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	mp "github.com/mackerelio/go-mackerel-plugin"
	probing "github.com/prometheus-community/pro-bing"
)

// Plugin mackerel plugin
type Plugin struct {
	Count   int
	Timeout int
	Hosts   []string
	Verbose bool
}

// GraphDefinition interface for mackerelplugin
func (p *Plugin) GraphDefinition() map[string]mp.Graphs {
	graphdef := map[string]mp.Graphs{
		"ping.latency": {
			Label:   "Ping Latency",
			Unit:    "milliseconds",
			Metrics: []mp.Metrics{},
		},
		"ping.packet_loss": {
			Label:   "Ping Packet Loss",
			Unit:    "percentage",
			Metrics: []mp.Metrics{},
		},
	}

	for _, host := range p.Hosts {
		eh := strings.ReplaceAll(host, ".", "_")
		latencyGraph := graphdef["ping.latency"]
		latencyGraph.Metrics = append(latencyGraph.Metrics, mp.Metrics{Name: eh + ".avg", Label: host + " avg"}, mp.Metrics{Name: eh + ".min", Label: host + " min"}, mp.Metrics{Name: eh + ".max", Label: host + " max"}, mp.Metrics{Name: eh + ".stddev", Label: host + " stddev"})
		graphdef["ping.latency"] = latencyGraph
		lossGraph := graphdef["ping.packet_loss"]
		lossGraph.Metrics = append(lossGraph.Metrics, mp.Metrics{Name: eh + ".packet_loss", Label: host + " packet loss"})
		graphdef["ping.packet_loss"] = lossGraph
	}
	return graphdef
}

type pingResult struct {
	host       string
	avg        float64
	min        float64
	max        float64
	stddev     float64
	packetLoss float64
}

func (p *Plugin) ping(host string, channel chan *pingResult) {
	result := &pingResult{host: host}
	pinger, err := probing.NewPinger(host)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create pinger: %s\n", err)
		result.packetLoss = -1
		channel <- result
		return
	}

	pinger.Count = p.Count
	pinger.Timeout = time.Duration(p.Timeout) * time.Second
	pinger.SetPrivileged(true) // if false, use UDP socket
	pinger.OnFinish = func(stats *probing.Statistics) {
		result.packetLoss = stats.PacketLoss
		result.avg = float64(stats.AvgRtt.Microseconds()) / 1000
		result.min = float64(stats.MinRtt.Microseconds()) / 1000
		result.max = float64(stats.MaxRtt.Microseconds()) / 1000
		result.stddev = float64(stats.StdDevRtt.Microseconds()) / 1000

		if p.Verbose {
			fmt.Fprintf(os.Stderr, "--- %s ping statistics ---\n", stats.Addr)
			fmt.Fprintf(os.Stderr, "%d packets transmitted, %d received, %v%% packet loss\n", stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
			fmt.Fprintf(os.Stderr, "rtt min/avg/max/mdev = %v/%v/%v/%v\n", stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
		}
	}

	if p.Verbose {
		fmt.Fprintf(os.Stderr, "PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
	}
	err = pinger.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to run pinger: %s\n", err)
		result.packetLoss = -1
		channel <- result
		return
	}

	channel <- result
}

// FetchMetrics interface for mackerelplugin
func (p *Plugin) FetchMetrics() (map[string]float64, error) {
	channel := make(chan *pingResult)
	for _, host := range p.Hosts {
		go p.ping(host, channel)
	}

	results := make([]*pingResult, 0, len(p.Hosts))
	for i := 0; i < len(p.Hosts); i++ {
		results = append(results, <-channel)
	}

	ret := make(map[string]float64)
	for _, r := range results {
		if r.packetLoss == -1 {
			continue
		}
		eh := strings.ReplaceAll(r.host, ".", "_")
		ret[eh+".packet_loss"] = r.packetLoss
		if r.packetLoss == 100.0 {
			continue
		}
		ret[eh+".avg"] = r.avg
		ret[eh+".min"] = r.min
		ret[eh+".max"] = r.max
		ret[eh+".stddev"] = r.stddev
	}
	return ret, nil
}

// Do the plugin
func Do() {
	optCount := flag.Int("c", 3, "Number of ping packets")
	optTimeout := flag.Int("t", 5, "Timeout seconds for each ping packet")
	optVerbose := flag.Bool("v", false, "Verbose mode")
	flag.Parse()

	if len(flag.Args()) == 0 {
		flag.Usage()
		fmt.Fprintf(os.Stderr, "Hosts are not specified\n")
		os.Exit(1)
	}

	plugin := mp.NewMackerelPlugin(&Plugin{
		Count:   *optCount,
		Timeout: *optTimeout,
		Hosts:   flag.Args(),
		Verbose: *optVerbose,
	})
	plugin.Run()
}
