# Ansible playbook

>:warning: You need to have knowledge in Ansible in order to run these scripts by hand.

This playbook is used to install the dependencies needed by the [`yourbtc.link`](../README.md) 
platform. It contains two generic roles for [`bitcoind`](../yourbtc-ansible/bitcoin-role) and [`lnd`](../yourbtc-ansible/lnd-role).
You can tune it to fit your needs by importing the roles in a playbook and changing some variables as 
you can see below.

### Run playbook for a specific network

In order to run the project against one specific network you need to overwrite some variables. 
See the file `regtest.yaml` as an example. If you want to install in into a host, just use the 
following command:

```shell
make install EXTRA_VARS=@regtest.yaml
```
