`fly -t ci.neighborly.tools set-pipeline -c ci/pipeline.yml -p golang-app --var "deploy_key=$(cat /Users/dmml/.ssh/golang_rsa)"`
