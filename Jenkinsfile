pipeline {
    agent any

    environment {
        GO = '/usr/bin/go'
    }

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
                    /usr/bin/go version
                    /usr/bin/go env GOROOT
                    /usr/bin/go env GOPATH
                    /usr/bin/go env GOMOD
                '''
            }
        }

        stage('Vet') {
            steps {
                sh '''
                    set -e
                    /usr/bin/go vet ./...
                '''
            }
        }

        stage('Test') {
            steps {
                sh '''
                    set -e

                    /usr/bin/go test -v \
                        -coverprofile=coverage.out \
                        ./...

                    echo "=== Coverage ==="
                    /usr/bin/go tool cover -func=coverage.out
                '''
            }
        }

        stage('Build') {
            steps {
                sh '''
                    set -e

                    CGO_ENABLED=0 \
                    /usr/bin/go build -o myapp .

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
