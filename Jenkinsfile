pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Check Go') {
            steps {
                sh '''
                    set -e
                    echo "=== Go ==="
                    which go
                    go version

                    echo "=== Go environment ==="
                    go env GOROOT
                    go env GOPATH
                    go env GOMOD
                '''
            }
        }

        stage('Vet') {
            steps {
                sh '''
                    set -e
                    go vet ./...
                '''
            }
        }

        stage('Test') {
            steps {
                sh '''
                    set -e
                    go test -v -coverprofile=coverage.out ./...

                    echo "=== Coverage ==="
                    go tool cover -func=coverage.out
                '''
            }
        }

        stage('Build') {
            steps {
                sh '''
                    set -e

                    CGO_ENABLED=0 go build -o myapp .

                    echo "=== Binary ==="
                    ls -lh myapp
                    file myapp
                '''
            }
        }
    }

    post {
        success {
            archiveArtifacts(
                artifacts: 'myapp,coverage.out',
                fingerprint: true
            )

            echo 'Build thành công, artifact đã được archive'
        }

        failure {
            echo 'Pipeline FAILED'
        }
    }
}
