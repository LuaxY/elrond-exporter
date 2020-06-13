package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/LuaxY/elrond-exporter/pkg/elrond"
)

var (
	port     string
	interval int
)

func main() {
	flag.StringVar(&port, "port", "8888", "metrics http port")
	flag.IntVar(&interval, "interval", 5, "interval in seconds")
	flag.Parse()

	ctx := context.Background()

	if err := initPrometheus(); err != nil {
		log.Fatal(err)
	}

	nodes, err := elrond.Discovery()

	if err != nil {
		log.Fatal(err)
	}

	ticker := time.NewTicker(5 * time.Second)

	go func() {
		for {
			for _, node := range nodes {
				if err := updateStatus(node); err != nil {
					log.Println(err)
				}

				if err := updateStatistics(node); err != nil {
					log.Println(err)
				}
			}

			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				continue
			}
		}
	}()

	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))

	log.Println("start listening on port " + port)
	log.Println("visit http://localhost:" + port + "/metrics")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
