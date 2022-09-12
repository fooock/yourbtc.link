# Ansible Bitcoin role

Ansible role to install the [Bitcoin Core](https://bitcoincore.org/en/about/) client as a `systemd` service. By default,
it uses sane defaults and some hardening measures for the Systemd service.

### Requirements

>At this moment this role is only supported in Ubuntu operating systems. 

This role requires a user with `sudo` permissions to work properly.

### Testing

You can execute tests using `molecule`. Install the [`requirements.txt`](molecule) file depending on if you want
to execute tests through Docker or with a VM managed by Vagrant. For example, if you want to use a VM:

```bash
$ molecule test -s vagrant
```

### Variables

You can change some variables to install this role to fit your needs. The default values to install the 
Bitcoin node are the following ones:

| Name              	 | Value              	 |
|---------------------|----------------------|
| `bitcoin_user`    	 | `bitcoin`          	 |
| `bitcoin_group`   	 | `bitcoin`          	 |
| `bitcoin_version` 	 | `23.0`             	 |
| `bitcoin_arch`    	 | `x86_64-linux-gnu` 	 |

>If you want to install Bitcoin into a Raspberry you need to change the architecture to `aarch64-linux-gnu`.

By default, this installer uses `gpg` to verify the integrity and signature of the downloaded artifacts. This
behaviour is controlled by the `bitcoin_pgp_builders_pub_key` field. The content of this structure and default values
are the following ones:

| Name       	 | ID                                         	 |
|--------------|----------------------------------------------|
| `darosior` 	 | `590B7292695AFFA5B672CBB2E13FC145CD3F4304` 	 |
| `guggero`  	 | `F4FC70F07310028424EFC20A8E4256593F177720` 	 |
| `fanquake` 	 | `E777299FC265DD04793070EB944D35F9AC3DB76A` 	 |

If you only want to verify with one user, you should use something like this:

```yaml
bitcoin_pgp_builders_pub_key:
  - id: 590B7292695AFFA5B672CBB2E13FC145CD3F4304
    name: darosior
```

>I use the Guix attestations to verify the release. The data can be found on the [Bitcoin Github official repository](https://github.com/bitcoin-core/guix.sigs).
> If the release can't be trusted the role will fail the installation.

To configure the Bitcoin node, you can use the following variables:

| Name                   	 | Value           	 | Note                                             	 |
|--------------------------|-------------------|----------------------------------------------------|
| `bitcoin_data_dir`     	 | `/data/bitcoin` 	 | 	                                                  |
| `bitcoin_network`      	 | `main`          	 | Valid values are: `regtest`, `signet` and `test` 	 |
| `bitcoin_rpc_user`     	 | `yourbtc`       	 | 	                                                  |
| `bitcoin_rpc_password` 	 | `yourbtc`       	 | 	                                                  |
| `bitcoin_zmq_host`     	 | `127.0.0.1`     	 | 	                                                  |
