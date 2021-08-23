### Експортер метрик в формате prometheus


## Использование

В данный момент можно толко читать из файла. Список переменных среды необходимо выставить обязательно
1. env:"PROM_PORT" - порт на котором експортер будет слушать подключения
2. env:"PROM_SOURCE_TYPE" - тип источника(поддерживается только 'file')
3. env:"PROM_SOURCE" - путь к файлу из которого брать метрики
4. env:"PROM_RENEW" - время обновления файла в секундах
5. env:"PROM_TOKEN" - токен для доступа

## Пример

    export PROM_PORT=8080
    export PROM_SOURCE_TYPE=file
    export PROM_SOURCE=./metrics.txt
    export PROM_RENEW=5
    export PROM_TOKEN=securetoken

    chmod +x prom-exporter
    ./prom-exporter


# Особенности
Для получения метрик требуется отправить запрос на /metrics, в headers выставить "Token: $PROM_TOKEN"

В файле метрики должны выглядеть следующим образом

    uptime_metric,equipment=server,resets=25,test=ddd,tag=test 4


uptime_metric - название метрики
equipment=server,resets=25,test=ddd,tag=test - теги
4 - значение метрики