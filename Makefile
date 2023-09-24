.PHONY: run
run: run-api

.PHONY: run-api
run-api:
	docker compose up

.PHONY: it-api
it-api: #apiコンテナに接続
	docker exec -it api /bin/bash

.PHONY: it-db
it-db: #dbコンテナに接続
	docker exec -it db /bin/bash