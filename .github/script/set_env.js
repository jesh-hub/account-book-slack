module.exports = ({ github, context, core }) => {
  const eventName = context.eventName
  //  const { BRANCH_NAME } = process.env
  let deployStage = ''

  switch (eventName) {
    case 'push':
      deployStage = "prod"
      break
    case 'pull_request':
      deployStage = "dev"
      break
  }
  if (deployStage) {
    console.log(`Add OS Env var ==> DEPLOY_STAGE: ${deployStage}`)
    core.exportVariable('DEPLOY_STAGE', deployStage)
  } else {
    core.setFailed(`Can't set environment`)
  }
}