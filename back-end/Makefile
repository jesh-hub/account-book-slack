.PHONY: deploy-dev deploy-prod

invoke-local:
	sls invoke local -f payments --verbose

invoke-remote:
	sls invoke -f payments --log

deploy:
	sls deploy --verbose