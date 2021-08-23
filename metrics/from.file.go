package metrics

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func GetMetricFromFile(filename string) []PromMetric {
	promSlice := []PromMetric{}
	metricTagsSlice := []string{}
	metricTags := map[string]string{}
	// задаем regex для парсинга метрик из файла
	r := regexp.MustCompile(`(?P<metric_name>^[a-z_]+?),?(?P<metric_tags>([a-z0-9_-]+=[a-z0-9_-]+,?)+)? (?P<metric_value>\d+)`)
	// вычитываем данные из файла
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Have read from file", filename)
	scanner := bufio.NewScanner(file)
	// парсим строки файла, создаем экземпляр метрики на каждую строку и создаем PromMetricSlice
	for scanner.Scan() {
		metricName := r.FindStringSubmatch(scanner.Text())[1]
		if r.FindStringSubmatch(scanner.Text())[2] != "" {
			// преобразуем строку в slice и далее в map, чтобы соответствовать struct
			metricTagsSlice = strings.Split(r.FindStringSubmatch(scanner.Text())[2], ",")
			for _, tag := range metricTagsSlice {
				intSlice := strings.Split(tag, "=")
				metricTags[intSlice[0]] = intSlice[1]
			}
		} else {
			metricTags = map[string]string{}
		}
		metricValue, _ := strconv.ParseFloat(r.FindStringSubmatch(scanner.Text())[4], 64)
		singleMetric := PromMetric{
			Name:  metricName,
			Tags:  metricTags,
			Value: metricValue,
		}
		promSlice = append(promSlice, singleMetric)
	}
	return promSlice
}
