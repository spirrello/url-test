#!groovy


timestamps {
    node {

        stage('Checkout') {
            checkout scm
        }

        stage('Build') {
            sh './gradlew build -x test'
            //sh './gradlew build --no-daemon'

            sh 'go version'
        }

        stage('Test') {

            // unit tests
            echo "Testing...."

            sh "go test -v"

        }
    }
}