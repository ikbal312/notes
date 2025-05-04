requirements:

<!-- install ansible -->
    sudo apt-get update -y
    sudo apt install software-properties-common -y
    sudo apt-add-repository --yes --update ppa:ansible/ansible
    sudo apt-get install -y ansible


<!-- install aws sdk  and cli -->
    pip install boto3 botocore awscli
    
<!-- install ansible collection "amazon.aws" -->
    ansible-galaxy collection install amazon.aws



run:
    ansible-playbook create_infra.yml
