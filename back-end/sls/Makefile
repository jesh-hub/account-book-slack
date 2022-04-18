.PHONY: deploy-dev deploy-prod

invoke-local:
	sls invoke local -f payments -s dev --verbose

invoke-remote:
	sls invoke -f payments -s dev --log

deploy:
	sls deploy --stage dev --verbose