.PHONY: deploy-dev deploy-prod

invoke-local:
	sls invoke local -f payments -s dev --verbose

invoke-remote:
	sls invoke -f payments -s dev --log

deploy-dev:
	sls deploy --stage dev --verbose

deploy-prod:
	sls deploy --stage prod --verbose