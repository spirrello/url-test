#!groovy

/**
 * This is an example Jenkins pipeline Jenkinsfile for building and testing a private Github Go project that uses glide.sh.
 * This style of Pipeline works with the Jenkins Git plugin and multibranch pipeline projects. In a multi-branch project configuring
 * a webhook for push and pull request is needed from the Github repository to get builds to automatically start after a push.
 */
pipeline {
    // agent {
    //     label 'your-agent-goes-here'
    // }
    stages {
        stage('Your Build Stage Name Goes Here') {
            steps {
                script {
                    /**
                     * To be able to access this Jenkins `tool` the https://wiki.jenkins.io/display/JENKINS/Go+Plugin plugin is needed.
                     * With more recent versions of Jenkins the documentation for adding a `go` installation is out of date. To properly
                     * configure a go installation go to the Jenkins tools configuration (Manage Jenkins -> Global Tool Configuration)
                     * find the "Go" and "Go installations" section and click "Add Go". The `name` specified below should
                     * line up with the "Go installation" to be used.
                     */
                    def root = tool name: '1.12.6', type: 'go'

                    /**
                     * Add in GOPATH, GOROOT, GOBIN to the environment and add go to the path for jenkins.
                     * The environment variables are needed for glide to properly work and adding go to the path is required to
                     */
                    withEnv(["GOPATH=${env.WORKSPACE}/go", "GOROOT=${root}", "GOBIN=${root}/bin", "PATH+GO=${root}/bin"]) {
                        sh "mkdir -p ${env.WORKSPACE}/go/src"

                        echo "Installing glide for this step"
                        sh 'curl https://glide.sh/get | sh'

                        echo "Configuring git to use ssh rather than https for downloading private repositories"
                        // This configures git settings to allow for private repositories to be downloaded with glide.
                        sh "git config --local url.ssh://git@github.com/.insteadOf https://github.com/"

                        echo "Installing glide dependencies"
                        sh "glide install"

                        echo "Building Go Code"
                        sh "go build ..."

                    }
                }
            }
        }
        stage('Your Test Stage Name Goes Here') {
            steps {
                script {
                    def root = tool name: '1.12.6', type: 'go'
                    withEnv(["GOPATH=${env.WORKSPACE}/go", "GOROOT=${root}", "GOBIN=${root}/bin", "PATH+GO=${root}/bin"]) {
                        echo "Installing glide for this step"
                        sh 'curl https://glide.sh/get | sh'

                        echo "Testing Go Code"
                        /**
                         * Since glide is installed, glide novendor or nv for short can be taken advantage of to list all
                         * files to test sans vendored dependencies.
                         */
                        sh 'go test -v $(glide nv)'
                    }
                }
            }
        }
    }
    /**
     * This post step will always execute regardless of a build failing or passing to clean up the setting that allows glide
     * to install private dependencies from Github. When using `checkout scm` or the default Jenkins clone step for a git
     * multibranch pipeline this undo change is needed. If the url change is not undone it will fail subsequent builds because the
     * Jenkins Git plugin will fail to clone the repository correctly.
     */
    // post {
    //     always {
    //         script {
    //             echo "Undoing config for git to use ssh rather than https for downloading private repositories"
    //             sh "git config --local --unset url.ssh://git@github.com/.insteadOf https://github.com/"
    //         }
    //     }
    // }
}

// timestamps {
//     node {

//         stage('Checkout') {
//             checkout scm
//         }

//         stage('Build') {

//             echo 'Setting go path....'
//             sh './setgo.sh'

//             //sh './gradlew build -x test'
//             sh './gradlew build --no-daemon'

//             echo '$GOROOT'
//            //echo 'go version'

//         }

//         stage('Test') {

//             // unit tests
//             echo "Testing...."

//             //sh "go test -v"

//         }
//     }
// }