# https://www.microfocus.com/documentation/host-access-for-the-cloud/2-4-2/user-guide/using-prometheus-grafana.html

FROM grafana/grafana:latest
ADD ./provisioning /etc/grafana/provisioning
ADD ./grafana.ini /etc/grafana/grafana.ini
ADD ./dashboards /var/lib/grafana/dashboards