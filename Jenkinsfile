#!groovy


timestamps {

        node {
        def root = tool name: 'Go 1.12.6', type: 'go'
            ws("${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/src/github.com/spirrello/url-test") {
                withEnv(["GOROOT=${root}", "GOPATH=${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/", "PATH+GO=${root}/bin"]) {
                    env.PATH="${GOPATH}/bin:$PATH"

                    stage('Checkout'){

                        git url: 'https://github.com/spirrello/url-test.git'

                        //sh './gradlew build --no-daemon'
                        sh './gradlew build -x test'

                    }


                    stage('Test'){

                        sh 'go test -v'

                        sh 'go run -race url-test.go'
                    }


                    stage('Build'){

                        sh "go build ."
                    }

                    // Do nothing.
                }
            }
        }
}
