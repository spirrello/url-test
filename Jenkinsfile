#!groovy


timestamps {
    // node {
    //     def root = tool name: 'Go 1.12.6', type: 'go'
    //     // Export environment variables pointing to the directory where Go was installed
    //     withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
    //         sh 'go version'
    //         //sh "go test -v"


    //         stage('Checkout') {
    //             checkout scm

    //             sh "go test -v"
    //         }
    //     }
    // }

    node {
    def root = tool name: 'Go 1.12.6', type: 'go'
    ws("${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/src/github.com/spirrello/url-test") {
        withEnv(["GOROOT=${root}", "GOPATH=${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/", "PATH+GO=${root}/bin"]) {
            env.PATH="${GOPATH}/bin:$PATH"

            stage 'Checkout'

            git url: 'https://github.com/spirrello/url-test.git'

            stage 'preTest'
            sh 'go version'
            //sh 'go get -u github.com/golang/dep/...'
            //sh 'dep init'

            stage 'Test'
            sh 'go vet'
            sh 'go test -cover'

            stage 'Build'
            sh 'go build .'

            stage 'Deploy'
            // Do nothing.
        }
    }
}
}
