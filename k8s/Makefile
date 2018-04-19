.PHONY: create-secret
create-secret:
	kubectl create secret generic rssnotify-config-secret --from-file ../config.yaml

.PHONY: deploy
deploy:
	kubectl apply -f cronjob.yaml