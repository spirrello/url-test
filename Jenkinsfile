#!groovy


timestamps {
    node {
        def root = tool name: 'Go 1.12.6', type: 'go'
        // Export environment variables pointing to the directory where Go was installed
        withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
            sh 'go version'
            //sh "go test -v"


            stage('Checkout') {
                checkout scm

                sh "go test -v"
            }
        }
    }
}
