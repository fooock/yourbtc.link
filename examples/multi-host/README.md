# Running on different hosts

In this example we are going to use two virtual machines. One will have `bitcoind` installed
and the other one will have `lnd`. Each one with its own resources but sharing the same private
network in order to allow communication succeed. If you want to test this using Vagrant just execute 
the following command:

```shell
$ vagrant up
```

This command will start up two virtual machines. You can check it using the following command:

```bash
$ vagrant global-status
id       name     provider   state   directory
-------------------------------------------------------------------------------------------
144cdd6  bitcoind virtualbox running /projects/yourbtc.link/examples/multi-host
c95c04b  lnd      virtualbox running /projects/yourbtc.link/examples/multi-host
```

After the command ends successfully, check that services are running properly inside the
virtual machine. You can connect to each one with using the machine name after `vagrant ssh`.

### Configuration

If you want to change the configuration you need to take into account that the most important
part is the networking since if configured wrong your LND node will be unable to start. As a 
summary, you need to set two static IPs. In our example they are:

| Service 	 | `lnd`          	 | `bitcoind`     	 |
|-----------|------------------|------------------|
| IP   	    | `192.168.56.5` 	 | `192.168.56.4` 	 |

>If you want to change these values to test, do it directly in the [Vagrantfile](Vagrantfile) and remember
> to update also the values in the [regtest.yaml](regtest.yaml) file.

### Check status

For example, to check the `bitcoind` service inside its own virtual machine first connect to the
machine using `vagrant ssh bitcoind` and then check the service status:

```bash
vagrant@yourbtc-bitcoind:~$ systemctl status bitcoind-regtest.service
● bitcoind-regtest.service - Bitcoin Core daemon for network Regtest
     Loaded: loaded (/lib/systemd/system/bitcoind-regtest.service; enabled; vendor preset: enabled)
     Active: active (running) since Sat 2022-09-17 16:17:35 UTC; 2min 2s ago
       Docs: https://github.com/fooock/yourbtc.link
    Process: 18291 ExecStart=/usr/local/bitcoin-core-23.0/bin/bitcoind -conf=/etc/bitcoin/regtest/bitcoin.conf (code=exited, status=0/SUCCESS)
   Main PID: 18294 (bitcoind)
      Tasks: 15 (limit: 1131)
     Memory: 38.7M
     CGroup: /system.slice/bitcoind-regtest.service
             └─18294 /usr/local/bitcoin-core-23.0/bin/bitcoind -conf=/etc/bitcoin/regtest/bitcoin.conf
```

To check the `lnd` service connect to the `lnd` machine and the the service status:

```bash
vagrant@yourbtc-lnd:~$ systemctl status lnd-regtest.service
● lnd-regtest.service - Lightning Network Daemon for network Regtest
     Loaded: loaded (/lib/systemd/system/lnd-regtest.service; enabled; vendor preset: enabled)
     Active: active (running) since Sat 2022-09-17 16:18:11 UTC; 7min ago
       Docs: https://github.com/fooock/yourbtc.link
   Main PID: 17263 (lnd)
     Status: "Wallet locked"
      Tasks: 8 (limit: 1131)
     Memory: 11.7M
     CGroup: /system.slice/lnd-regtest.service
             └─17263 /usr/local/lnd-v0.15.0-beta/lnd -C /etc/lnd/regtest/lnd.conf --lnddir=/data/lnd
```

>Note the `Wallet locked` status in the `lnd` service.

To destroy all resources just execute the following commands:

```bash
$ vagrant destroy
```
