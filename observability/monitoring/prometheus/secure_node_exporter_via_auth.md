make local folder with necessary permission

    sudo mkdir /etc/node_exporter/
    sudo touch /etc/node_exporter/config.yml
    sudo chown -R node_exporter:node_exporter /etc/node_exporter
    sudo chmod 755 /etc/node_exporter
    sudo chmod 644 /etc/node_exporter/config.yml

install and generate password hash

    sudo apt update
    sudo apt install apache2-utils -y

    htpasswd -nBC 10 "" | tr -d ':\n'; echo


add hash to config

    sudo vi /etc/node_exporter/config.yml
    basic_auth_users:
        prometheus: <hashed-password>

    
configure node exporter systemd service

    sudo vi /etc/systemd/system/node_exporter.service

    <!-- update with this -->

    ExecStart=/usr/local/bin/node_exporter --web.config.file=/etc/node_exporter/config.yml

proper permission for node exporter directory and file



reload and restart systemd

    sudo systemctl daemon-reload
    sudo systemctl restart node_exporter
    sudo systemctl enable node_exporter




prometheus.yml

    - job_name: 'node_exporter'
        static_configs:
          - targets: ['localhost:9100']
            basic_auth:
                username: prometheus
                password: secret-password
