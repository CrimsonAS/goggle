pipeline {
    agent none

    stages {
        // We can't actually build, because CentOS lacks SDL2..
        /*
        stage('Build') {
            agent any
            environment {
                GOPATH = "$WORKSPACE"
            }
            steps {
                // Get this in the right place.
                dir('src/github.com/CrimsonAS/goggle') {
                    checkout scm
                    dir('examples/trippy') {
                        sh 'go get' // ### we should vendor, probably
                        sh 'go build'
                        archiveArtifacts artifacts: 'trippy'
                    }
                    dir('examples/draggablerect') {
                        sh 'go get' // ### we should vendor, probably
                        sh 'go build'
                        archiveArtifacts artifacts: 'draggablerect'
                    }
                    dir('examples/manyrects') {
                        sh 'go get' // ### we should vendor, probably
                        sh 'go build'
                        archiveArtifacts artifacts: 'manyrects'
                    }
                }
            }
        }
        */
        stage('Test') {
            agent any
            environment {
                GOPATH = "$WORKSPACE"
            }
            steps {
                // Get this in the right place.
                dir('src/github.com/CrimsonAS/goggle') {
                    sh "go test -v ./animation"
                    sh "go test -v ./renderer/private"
                    sh "go test -v ./sg"
                }
            }
        }
    }
}
