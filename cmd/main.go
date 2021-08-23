package main

import (
	"fmt"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/jasonlvhit/gocron"
	"github.com/kuvalda989/prom-exporter/config"
	"github.com/kuvalda989/prom-exporter/metrics"
	log "github.com/sirupsen/logrus"
	"gopkg.in/macaron.v1"
)

// глобальная переменная, содержащая все актуальные метрики, обновляется через крон
var metricSlice []metrics.PromMetric

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
}

func renewMetrics(cfg config.Config) {
	// обработка ошибок
	defer func() {
		if panicInfo := recover(); panicInfo != nil {
			fmt.Printf("%v, %s", panicInfo, string(debug.Stack()))
			metricSlice = nil
		}
	}()
	// запуск обновления метрик
	metricSlice = metrics.GetMetrics(cfg)
}

func startCron(cfg config.Config) {
	// запуск cron джобы для периодического обновления списка метрик
	cronSeconds, _ := strconv.ParseUint(cfg.RenewTimeSeconds, 10, 64)
	cronSheduler := gocron.NewScheduler()
	log.Info("Create new cron with time ", cronSeconds)
	cronSheduler.Every(cronSeconds).Seconds().Do(renewMetrics, cfg)
	<-cronSheduler.Start()

}

func metricSliceToString(metricSlice []metrics.PromMetric) string {
	// создание строки из каждой метрики PromMetric
	promMetricString := ""
	metricString := ""
	for _, metric := range metricSlice {
		if len(metric.Tags) > 0 {
			metricString = fmt.Sprintf("%v{%v} %v\n", metric.Name, tagToString(metric.Tags), metric.Value)
		} else {
			metricString = fmt.Sprintf("%v %v\n", metric.Name, metric.Value)
		}
		promMetricString += metricString
	}
	return promMetricString
}

func tagToString(tag map[string]string) string {
	// конвертирую словарь в строку
	sumString := ""
	for key, value := range tag {
		sumString += fmt.Sprintf("%v=%q ", key, value)
	}
	return strings.TrimSpace(sumString)
}

func metricsHandler(ctx *macaron.Context) (int, string) {
	if metricSlice != nil && ctx.Req.Header.Get("Token") == config.Get().Token {
		return 200, metricSliceToString(metricSlice)
	} else if ctx.Req.Header.Get("Token") != config.Get().Token {
		return 403, "Access denied - wrong token"
	} else {
		return 500, "Something went wrong"
	}
}

func main() {
	// подгружаем конфигурационные переменные из env
	cfg := config.Get()
	// запускаем крон-джобу обновления метрик
	go startCron(cfg)

	webServer := macaron.Classic()
	webServer.Get("/metrics", metricsHandler)
	webServerPort, _ := strconv.Atoi(cfg.Port)
	webServer.Run("0.0.0.0", webServerPort)
}
