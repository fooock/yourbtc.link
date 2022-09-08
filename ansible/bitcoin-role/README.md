# Ansible Bitcoin role

Ansible role to install the [Bitcoin Core](https://bitcoincore.org/en/about/) client as a `systemd` service.

### Requirements

>At this moment this role is only supported in Ubuntu systems. 

This role requires a user with `sudo` permissions to work properly.

### Testing

You can execute tests using `molecule`. Install the [`requirements.txt`](molecule) file depending if you want
to execute tests through Docker or with a VM managed by Vagrant.
