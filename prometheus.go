package common

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"net/http"
)

const (
	PROMETHEUS_ADDR = "0.0.0.0"
)

func PrometheusBoot(port int) {
	http.Handle("/metrics", promhttp.Handler())

	go func() {
		err := http.ListenAndServe(fmt.Sprintf("%s:%d", PROMETHEUS_ADDR, port), nil)
		if err != nil {
			log.Fatal("启动失败")
		}
		log.Info("监控启动，端口为:", port)
	}()
}
