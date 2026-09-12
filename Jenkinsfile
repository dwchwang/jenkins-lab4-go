pipeline {
    agent any
    environment {
        // Định nghĩa đường dẫn Go và PATH ngay từ đầu để không phải gõ export lại ở từng stage
        PATH = "\(pipeline {
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
                    if [ ! -d "$WORKSPACE/go" ]; then
                        echo "Chưa có Go trong workspace, tiến hành tải về..."
                        curl -sL https://go.dev/dl/go1.22.0.linux-amd64.tar.gz | tar -C "$WORKSPACE" -xz
                    fi
                    export PATH="\(WORKSPACE/go/bin:\)PATH"
                    go version
                '''
            }
        }
        stage('Vet') {
            steps {
                sh '''
                    export PATH="\(WORKSPACE/go/bin:\)PATH"
                    go vet ./...
                '''
            }
        }
        stage('Test') {
            steps {
                sh '''
                    export PATH="\(WORKSPACE/go/bin:\)PATH"
                    go test -v -coverprofile=coverage.out ./...
                    go tool cover -func=coverage.out
                '''
            }
        }
        stage('Build') {
            steps {
                sh '''
                    export PATH="\(WORKSPACE/go/bin:\)PATH"
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
}{WORKSPACE}/go/bin:\){env.PATH}"
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
                    if ! command -v go &> /dev/null; then
                        echo "Chưa có Go, tiến hành tải về..."
                        curl -sL https://go.dev/dl/go1.22.0.linux-amd64.tar.gz | tar -C ${WORKSPACE} -xz
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
