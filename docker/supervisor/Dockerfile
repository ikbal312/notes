FROM ubuntu:latest
# Sets environment variable to prevent interactive prompts during installation
RUN DEBIAN_FRONTEND=noninteractive
RUN apt-get update \
    && apt-get install -y \
    apache2 \
    php \
    libapache2-mod-php \
    mysql-client \
    supervisor \
    # Cleans up the local repository of retrieved package files
    && apt-get clean \ 
    # Removes all package lists from /var/lib/apt/lists/
    && rm -rf /var/lib/apt/lists/*




RUN apt-get update \
    && apt-get -y install \
    mysql-server \
    # Cleans up the local repository of retrieved package files
    && apt-get clean \ 
    # Removes all package lists from /var/lib/apt/lists/
    && rm -rf /var/lib/apt/lists/*

# configure apache
RUN a2enmod rewrite

# configure supervisord
COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf

EXPOSE 80 3306

CMD [ "/usr/bin/supervisord","-c", "/etc/supervisor/conf.d/supervisord.conf" ]

