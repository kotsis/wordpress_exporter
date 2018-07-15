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
or
```sh
$ wordpress_exporter -host=127.0.0.1 -port=3306 -user=uuuu -db=dddd -tableprefix=wp_ -pass=xxxx
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
$ sudo docker-compose up -d
```

Now a wordpress is being served at : http://localhost:8000 where you must visit and create a user with a password.
Then you can login in WordPress and create posts, users etc.

Next you must start the wordpress_exporter
```sh
$ wordpress_exporter -port=33306 -db=wordpress -user=wordpress -pass=wordpress1234
```

You will see the metrics from those actions.

# Grafana
You can find a WordPress dashboard in $GOPATH/src/github.com/kotsis/wordress_exporter/wordpress_grafana.json

For it to work you must define in Grafana a new Prometheus data source as prom1
This must be the Prometheus instance that is scrapin metrics from wordpress_exporter.
Then you can import the above json file and start viewing the metrics.
