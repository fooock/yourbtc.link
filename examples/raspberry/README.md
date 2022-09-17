# Install services on a Raspberry

If you have your Raspberry up and running, clone this repository:

```shell
$ git clone https://github.com/fooock/yourbtc.link.git
```

Go to the `examples/raspberry` directory:

```shell
$ cd examples/raspberry
```

and execute this command:

```shell
$ ansible-playbook -e @mainnet.yaml -i localhost, --connection=local playbook.yml
```

>If you don't want to execute this inside your Raspberry and want to use your local
> machine to do this same job, you just need to change the inventory to match the IP
> of your Raspberry inside your network. Also, remove the `connection` flag.

Wait until finish!
