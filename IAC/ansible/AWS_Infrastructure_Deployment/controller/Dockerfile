FROM ubuntu:latest AS base
# Avoid prompts from apt
ENV DEBIAN_FRONTEND=noninteractive


RUN  apt update -y \
    &&  apt install -y --no-install-recommends openssh-server ansible sudo  \
    &&  mkdir /var/run/sshd \
    &&  useradd -m -s /bin/bash ansible \
    &&  usermod -aG sudo ansible \
    &&  sed -i 's/^#PasswordAuthentication yes/PasswordAuthentication no/g' /etc/ssh/sshd_config \
    &&  echo "PermitRootLogin no" >> /etc/ssh/sshd_config \
    &&  echo "AllowUsers ansible" >> /etc/ssh/sshd_config  \
    &&  mkdir -p /home/ansible/.ssh \
    &&  chown -R ansible:ansible /home/ansible/.ssh  \
    &&  chmod 700 /home/ansible/.ssh  \
    &&  touch /home/ansible/.ssh/authorized_keys \
    &&  chmod 600 /home/ansible/.ssh/authorized_keys \
    &&  echo "ansible ALL=(ALL) NOPASSWD:ALL" | tee /etc/sudoers.d/ansible \
    &&  chmod 0440 /etc/sudoers.d/ansible \
    && rm -rf /var/lib/apt/lists/*  # Clean up apt cache
COPY --chown=ansible:ansible ./data/ssh/id_rsa.pub /home/ansible/.ssh/authorized_keys

EXPOSE 22

ENTRYPOINT ["/usr/sbin/sshd", "-D"]



FROM base AS controller

ENV DEBIAN_FRONTEND=noninteractive
# Set the virtual environment as the default Python environment
ENV PATH="/opt/venv/bin:$PATH"
# Install AWS CLI
RUN apt update -y \
    && apt install -y --no-install-recommends python3-pip python3-venv \
    && python3 -m venv /opt/venv \
    && /opt/venv/bin/pip install boto3 botocore awscli \
    && rm -rf /var/lib/apt/lists/*  # Clean up apt cache
    
ENTRYPOINT ["/usr/sbin/sshd", "-D"]