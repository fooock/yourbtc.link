# Ansible roles

[![.github/workflows/ansible.yml](https://github.com/fooock/yourbtc.link/actions/workflows/ansible.yml/badge.svg)](https://github.com/fooock/yourbtc.link/actions/workflows/ansible.yml)

>:warning: You need to have knowledge in Ansible in order to run these scripts by hand.

Here you can find the Ansible roles used by the platform. At this moment the available ones are:

* [`bitcoind`](bitcoin-role)
* [`lnd`](lnd-role)

### Requirements

You need to have at least `Python 3.8` and `Ansible 2.10` to run the roles. Go to [`Installing Ansible`](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html)
if you need to install it in your system. For testing purposes you will also need Docker or Vagrant.
