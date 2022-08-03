EXTRA_VARS := @../examples/regtest.yaml
INVENTORY := localhost,
CONNECTION := local
TAGS :=

install:
	ansible-playbook -e $(EXTRA_VARS) -i $(INVENTORY) --connection=$(CONNECTION) playbook.yml

install-bitcoin:
	$(eval TAGS=bitcoin)
	ansible-playbook -e $(EXTRA_VARS) -i $(INVENTORY) --connection=$(CONNECTION) playbook.yml --tags $(TAGS)

install-lnd:
	$(eval TAGS=lnd)
	ansible-playbook -e $(EXTRA_VARS) -i $(INVENTORY) --connection=$(CONNECTION) playbook.yml --tags $(TAGS)

#
# This directive helps executing things faster since it sets
# the default binary location with the version, network and we
# only need to pass the required command.
#
LND_VERSION := 0.15.0-beta
LND_NET := regtest
LND_C :=

lnd:
	/usr/local/lnd-$(LND_VERSION)/lncli -n $(LND_NET) $(LND_C)
