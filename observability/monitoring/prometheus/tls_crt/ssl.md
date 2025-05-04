# Generate a self signed  TLS Certificate for Node Exporter
    openssl req -new -newkey rsa:2048 -days 365 -nodes -x509 \
    -keyout node_exporter.key -out node_exporter.crt \
    -subj "/C=US/ST=California/L=Oakland/O=MyOrg/CN=localhost" \
    -addext "subjectAltName = DNS:localhost"



# exporter config
/etc/node_exporter/config.yml

    tls_server_config:
        cert_file: node_exporter.crt
        key_file: node_exporter.key

 /etc/systemd/system/node_exporter.service

    ExecStart=/usr/local/bin/node_exporter \
    --web.config.file=/etc/node_exporter/config.yml

# prometheus config
/etc/prometheus/prometheus.yml


# exporter node contains 
    crt and key file

# prometheus contains
    crt file


# copy give certificate to prometheus
    sudo cp /etc/node_exporter/node_exporter.crt /etc/prometheus/


# update prometheus.yml file
/etc/prometheus/prometheus.yml

    scrape_configs:
    - job_name: 'node_exporter'
        scheme: https
        tls_config:
            ca_file: /etc/prometheus/node_exporter.crt
            insecure_skip_verify: true
        static_configs:
          - targets: ['localhost:9100']
# reload prometheus 
    sudo systemctl daemon-reload
    sudo systemctl restart prometheus




# check
    curl  https://localhost:9100/metrics
    curl -k https://localhost:9100/metrics # k ignore self signed key
