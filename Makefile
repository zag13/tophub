.PHONY: api
# generate api
api:
	find app -type d -depth 1 -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) api'

.PHONY: wire
# generate wire
wire:
	find app -type d -depth 1 -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) wire'

.PHONY: proto
# generate proto
proto:
	find app -type d -depth 1 -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) proto'

.PHONY: zag
zag:
	find app -type d -depth 1 -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) zag'