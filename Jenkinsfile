pipeline {
    agent none

    stages {
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
                        sh 'go build'
                        archiveArtifacts artifacts: 'trippy'
                    }
                    dir('examples/draggablerect') {
                        sh 'go build'
                        archiveArtifacts artifacts: 'draggablerect'
                    }
                    dir('examples/manyrects') {
                        sh 'go build'
                        archiveArtifacts artifacts: 'manyrects'
                    }
                }
            }
        }
        stage('Test') {
            agent any
            environment {
                GOPATH = "$WORKSPACE"
            }
            steps {
                // Get this in the right place.
                dir('src/github.com/CrimsonAS/goggle') {
                    sh "go test -v ./..."
                }
            }
        }
    }
}
