# Running services on the same host

>This example is to show how we can install `bitcoind` with `lnd` in the same host as
> a Raspberry or a virtual machine.

If you want to test this using Vagrant just execute the following command:

```shell
$ vagrant up
```

### Configuration

As you can see in the [regtest.yaml](regtest.yaml) file, the `lnd_group` and `bitcoin_group` share
the same value. This is because `lnd` needs to read the Bitcoin data, and for this to work, 
they need to share the same group to be able to read the required files. Note that this is only required
when installing both services in the same host, like in this example.

### Check status

After the script ends successfully, check that services are running properly inside the
virtual machine.

To check the `bitcoind` service:

```
vagrant@ubuntu-focal:~$ systemctl status bitcoind-regtest.service
● bitcoind-regtest.service - Bitcoin Core daemon for network Regtest
     Loaded: loaded (/lib/systemd/system/bitcoind-regtest.service; enabled; vendor preset: enabled)
     Active: active (running) since Sat 2022-09-17 15:32:28 UTC; 8min ago
       Docs: https://github.com/fooock/yourbtc.link
   Main PID: 20575 (bitcoind)
      Tasks: 15 (limit: 1131)
     Memory: 38.4M
     CGroup: /system.slice/bitcoind-regtest.service
             └─20575 /usr/local/bitcoin-core-23.0/bin/bitcoind -conf=/etc/bitcoin/regtest/bitcoin.conf
```

To check the `lnd` service:

```
vagrant@ubuntu-focal:~$ systemctl status lnd-regtest.service
● lnd-regtest.service - Lightning Network Daemon for network Regtest
     Loaded: loaded (/lib/systemd/system/lnd-regtest.service; enabled; vendor preset: enabled)
     Active: active (running) since Sat 2022-09-17 15:39:34 UTC; 59s ago
       Docs: https://github.com/fooock/yourbtc.link
   Main PID: 31725 (lnd)
     Status: "Wallet locked"
      Tasks: 7 (limit: 1131)
     Memory: 11.3M
     CGroup: /system.slice/lnd-regtest.service
             └─31725 /usr/local/lnd-v0.15.0-beta/lnd -C /etc/lnd/regtest/lnd.conf --lnddir=/data/lnd
```

>Note the `Wallet locked` status in the `lnd` service.

To destroy all resources just execute the following command:

```shell
$ vagrant destroy
```
