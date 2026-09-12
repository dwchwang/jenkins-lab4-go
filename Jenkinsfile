pipeline {
    agent any

    environment {
        GO_VERSION = '1.22.0'
        GOROOT = "${WORKSPACE}/go-sdk/go"
        GOPATH = "${WORKSPACE}/gopath"
        PATH = "${WORKSPACE}/go-sdk/go/bin:${WORKSPACE}/gopath/bin:${PATH}"
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
                    if [ ! -x "${GOROOT}/bin/go" ]; then
                        mkdir -p "${WORKSPACE}/go-sdk"
                        curl -sL https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz \
                            | tar -C "${WORKSPACE}/go-sdk" -xz
                    fi
                    go version
                '''
            }
        }

        stage('Vet') {
            steps {
                sh 'go vet ./...'
            }
        }

        stage('Test') {
            steps {
                sh '''
                    go test -v -coverprofile=coverage.out ./...
                    go tool cover -func=coverage.out
                '''
            }
        }

        stage('Build') {
            steps {
                sh '''
                    CGO_ENABLED=0 go build -o myapp
                    ls -lh myapp
                '''
            }
        }
    }

    post {
        success {
            archiveArtifacts artifacts: 'myapp', fingerprint: true
            echo 'Build thành công, artifact đã lưu'
        }
    }
}
