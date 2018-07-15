# wordpress_exporter
Prometheus exporter for WordPress

# Install wordpress_exporter
```sh
$ go get github.com/kotsis/wordpress_exporter
```

# How to setup a docker WordPress for monitoring

# Usage of wordpress_exporter
```sh
$ wordpress_exporter -wpconfig=/path/to/wp-config
```
It starts serving metrics at http://localhost:8888/metrics

# Assumptions
For Prometheus to start scraping the metrics you have to edit /etc/prometheus/prometheus.yml and add:

```sh
  - job_name: 'wordpress'
    # metrics_path defaults to '/metrics'
    # scheme defaults to 'http'.
    static_configs:
    - targets: ['localhost:8888']
```

the above is valid if the exporter runs at the same host as prometheus service. If prometheus runs
in a docker container perhaps you will need to change localhost with the IP of the host system, something like 172.17.0.1

# WordPress service with docker-compose

Here is provided a quick WordPress service setup with docker-compose for testing the wordpress_exporter.
You can go in $GOPATH/src/github.com/kotsis/wordress_exporter and run:
```sh
$ docker-compose up -d
```

Now a wordpress is being server at :

# Grafana
You can find a WordPress dashboard in $GOPATH/src/github.com/kotsis/wordress_exporter/wordpress_grafana.json

For it to work you must define in Grafana a new Prometheus data source as prom1
This must be the Prometheus instance that is scrapin metrics from wordpress_exporter.
Then you can import the above json file and start viewing the metrics.
