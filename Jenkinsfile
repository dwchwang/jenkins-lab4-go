pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Setup Go') {
            steps {
                sh '''
                    # Cài Go nếu agent chưa có (dùng docker agent tốt hơn - lab 10)
                    which go || (
                        curl -sL https://go.dev/dl/go1.22.0.linux-amd64.tar.gz | tar -C /usr/local -xz
                    )
                    export PATH=$PATH:/usr/local/go/bin
                    go version
                '''
            }
        }

        stage('Vet') {
            steps {
                sh 'export PATH=$PATH:/usr/local/go/bin && go vet ./...'
            }
        }

        stage('Test') {
            steps {
                sh '''
                    export PATH=$PATH:/usr/local/go/bin
                    go test -v -coverprofile=coverage.out ./...
                    go tool cover -func=coverage.out
                '''
            }
        }

        stage('Build') {
            steps {
                sh '''
                    export PATH=$PATH:/usr/local/go/bin
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
