pipeline {
    agent none

    stages {
        stage('Build') {
            agent any
            environment {
                GOPATH = "/tmp"
            }
            steps {
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
        stage('Test') {
            agent any
            environment {
                GOPATH = "/tmp"
            }
            steps {
                sh "go test -v ./..."
            }
        }
    }
}
