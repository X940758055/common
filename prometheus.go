package common

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"net/http"
)

func PrometheusBoot(addr string, port int) {
	http.Handle("/metrics", promhttp.Handler())

	go func() {
		err := http.ListenAndServe(fmt.Sprintf("%s:%d", addr, port), nil)
		if err != nil {
			log.Fatal("启动失败:", err)
		}
		log.Info("监控启动，端口为:", port)
	}()
}
