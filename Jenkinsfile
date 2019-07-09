#!groovy


timestamps {

        node {
        def root = tool name: 'Go 1.12.6', type: 'go'
            ws("${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/src/github.com/spirrello/url-test") {
                withEnv(["GOROOT=${root}", "GOPATH=${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/", "PATH+GO=${root}/bin"]) {
                    env.PATH="${GOPATH}/bin:$PATH"

                    stage('Checkout'){

                        git url: 'https://github.com/spirrello/url-test.git'

                        sh './gradlew build --no-daemon'

                    }

                    stage('preTest'){
                        sh 'go version'


                    }

                    stage('Test'){

                        sh 'go test -v'
                    }


                    stage('Build'){

                        sh "go build ."
                    }


                    stage('Archive'){

                        archiveArtifacts artifacts: 'url-test.zip'

                    }
                    // Do nothing.
                }
            }
        }
}
