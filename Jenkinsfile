#!groovy


timestamps {
    node {

        stage('Checkout') {
            checkout scm
        }

        stage('Build') {

            echo 'Setting go path....'
            sh './setgo.sh'

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