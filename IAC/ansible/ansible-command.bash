# Ansible Cheatsheet - Detailed
### Install 
    sudo apt-get update -y
    sudo apt install software-properties-common -y
    sudo apt-add-repository --yes --update ppa:ansible/ansible
    sudo apt-get install -y ansible
### connection test
    ansible all --key-file <your_key_file> -i inventory -m ping
    ansible all -m ping


### Ad-hoc Commands
    ansible <host-pattern> -i <inventory> -m <module> -a "<module-arguments>"
    ansible all --list-hosts
    ansible all -m gather_facts --limit <remote_server_ip>

### Elevated(superuser or root privileges) ad-hoc commands executed on remote nodes.
    ansible all -m apt -a "update_cache=yes" --become     # update the package index on a remote server
    ansible all -m apt -a name=vim-nox --become     # To install a package like vim-nox
    ansible all -m apt -a "upgrade=dist" --become   # Upgrade the distribution

### Directory
    /etc/ansible/           # Description: Default location for Ansible configuration files, including `ansible.cfg`, inventory files, and roles.
    /etc/ansible/roles      # Description: Default location for Ansible roles installed system-wide.
    ~/.ansible/             # Description: User-specific Ansible configuration directory.  Can override settings in `/etc/ansible/`.
    /usr/share/ansible/     # Description: Location where Ansible modules are stored.
    /var/log/ansible.log    # Description: Default location for Ansible log files (if logging is enabled).

### Environment Variables
    ANSIBLE_CONFIG
        # Description: Specifies the path to the Ansible configuration file.
        # Usage: `export ANSIBLE_CONFIG=/path/to/ansible.cfg`
    ANSIBLE_INVENTORY
        # Description: Specifies the path to the Ansible inventory file.
        # Usage: `export ANSIBLE_INVENTORY=/path/to/inventory.ini`
    ANSIBLE_HOST_KEY_CHECKING
        # Description: Controls whether Ansible should verify host keys.
        # Usage: `export ANSIBLE_HOST_KEY_CHECKING=False` (Disable host key checking - not recommended for production)
    ANSIBLE_LIBRARY
        # Description: Specifies the path to custom Ansible modules.
        # Usage: `export ANSIBLE_LIBRARY=/path/to/custom/modules`
    ANSIBLE_ROLES_PATH
        # Description: Specifies the path to custom Ansible roles.
        # Usage: `export ANSIBLE_ROLES_PATH=/path/to/custom/roles`
    ANSIBLE_VAULT_PASSWORD
        # Description: Specifies the password for Ansible Vault.  Less secure than using `--ask-vault-pass`.
        # Usage: `export ANSIBLE_VAULT_PASSWORD=your_vault_password`
    ANSIBLE_FORCE_COLOR
        # Description: Forces Ansible to use colored output, even when running in a non-interactive terminal.
        # Usage: `export ANSIBLE_FORCE_COLOR=True`
