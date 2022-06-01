.PHONY: invoke-local invoke-local deploy deploy-all

export APP_NAME=payments

invoke-local:
	sls invoke local -f ${APP_NAME} --verbose

invoke-remote:
	sls invoke -f ${APP_NAME} --log

deploy:
	sls deploy -f ${APP_NAME} --verbose

deploy-all:
	sls deploy --verbose