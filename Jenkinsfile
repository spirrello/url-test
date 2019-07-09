#!groovy


timestamps {
    node {

        stage('Checkout') {
            checkout scm
        }

        stage('Build') {

            withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
                sh 'go version'
        }
            //sh './gradlew build -x test'
            sh './gradlew build --no-daemon'

            echo '$GOROOT'
           //echo 'go version'

        }

        stage('Test') {

            // unit tests
            echo "Testing...."

            //sh "go test -v"

        }
    }
}