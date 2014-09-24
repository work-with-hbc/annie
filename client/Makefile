.PHONY: test install

clients=coffee

test:
	$(foreach client, $(clients), cd $(client) && $(MAKE) test;)


install:
	$(foreach client, $(clients), cd $(client) && $(MAKE) install;)
