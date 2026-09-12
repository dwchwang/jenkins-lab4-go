pipeline {
    agent any

    environment {
        GO_VERSION = '1.22.0'
        GO_ROOT = "${WORKSPACE}/go"
        PATH = "${WORKSPACE}/go/bin:${env.PATH}"
    }

    stages {

        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Setup Go') {
            steps {
                sh '''
                    set -e

                    echo "=== Setup Go ==="

                    if [ ! -x "$GO_ROOT/bin/go" ]; then
                        echo "Go not found. Installing Go ${GO_VERSION}..."

                        rm -rf "$GO_ROOT"
                        mkdir -p "$GO_ROOT"

                        curl -fL \
                            "https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz" \
                            -o /tmp/go.tar.gz

                        tar -xzf /tmp/go.tar.gz \
                            -C "$WORKSPACE"

                        rm -f /tmp/go.tar.gz
                    fi

                    echo "Go version:"
                    go version
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

                    go test -v \
                        -coverprofile=coverage.out \
                        ./...

                    echo "=== Coverage ==="
                    go tool cover -func=coverage.out
                '''
            }
        }

        stage('Build') {
            steps {
                sh '''
                    set -e

                    CGO_ENABLED=0 \
                    GOOS=linux \
                    GOARCH=amd64 \
                    go build -o myapp .

                    echo "=== Binary ==="
                    ls -lh myapp
                    file myapp
                '''
            }
        }
    }

    post {
        success {
            archiveArtifacts \
                artifacts: 'myapp,coverage.out', \
                fingerprint: true

            echo 'Build thành công, artifact đã được archive.'
        }

        failure {
            echo 'Pipeline FAILED.'
        }

        always {
            echo "Build number: ${BUILD_NUMBER}"
        }
    }
}
