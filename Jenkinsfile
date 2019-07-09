#!groovy


timestamps {
    node {

        stage('Checkout') {
            checkout scm
        }

        stage('Build') {

            echo "go version"

            echo "HOME: $HOME"

            echo "PWD: $PWD"

            echo 'Setting go path....'
            sh './setgo.sh'

            echo '$GOROOT'


            // //sh './gradlew build -x test'
            // sh './gradlew build --no-daemon'

        //     def root = tool name: '1.8', type: 'go'

        // /**
        //     * Add in GOPATH, GOROOT, GOBIN to the environment and add go to the path for jenkins.
        //     * The environment variables are needed for glide to properly work and adding go to the path is required to
        //     */
        //     withEnv(["GOPATH=${env.WORKSPACE}/go", "GOROOT=${root}", "GOBIN=${root}/bin", "PATH+GO=${root}/bin"]) {
        //         sh "mkdir -p ${env.WORKSPACE}/go/src"

        //         echo "Installing glide for this step"
        //         sh 'curl https://glide.sh/get | sh'

        //         echo "Configuring git to use ssh rather than https for downloading private repositories"
        //         // This configures git settings to allow for private repositories to be downloaded with glide.
        //         sh "git config --local url.ssh://git@github.com/.insteadOf https://github.com/"

        //         echo "Installing glide dependencies"
        //         sh "glide install"

        //         echo "Building Go Code"
        //         sh "go build ..."




            }

        stage('Test') {

            // unit tests
            echo "Testing...."

            //sh "go test -v"

        }
    }
}
