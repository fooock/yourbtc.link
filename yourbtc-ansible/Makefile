EXTRA_VARS := @regtest.yaml
INVENTORY := localhost,
CONNECTION := local
TAGS :=

install:
	ansible-playbook -e $(EXTRA_VARS) -i $(INVENTORY) --connection=$(CONNECTION) playbook.yml

install-bitcoin:
	TAGS := bitcoin
	ansible-playbook -e $(EXTRA_VARS) -i $(INVENTORY) --connection=$(CONNECTION) playbook.yml --tags $(TAGS)

install-lnd:
	TAGS := lnd
	ansible-playbook -e $(EXTRA_VARS) -i $(INVENTORY) --connection=$(CONNECTION) playbook.yml --tags $(TAGS)
