.PHONY: invoke-local invoke-local deploy deploy-all

export APP_NAME=payments
export APP_NAME_USER=user_signup

invoke-local:
	sls invoke local -f ${APP_NAME} --verbose

invoke-remote:
	sls invoke -f ${APP_NAME} --log

deploy:
	sls deploy -f ${APP_NAME} --verbose
	sls deploy -f ${APP_NAME_USER} --verbose

deploy-all:
	sls deploy --verbose