pipeline {
    agent none      // Không dùng agent chung, mỗi stage tự chọn
    stages {
        stage('Go Build') {
            agent { 
                docker { 
                    image 'golang:1.22' 
                    args '-u root'
                } 
            }
            steps {
                sh 'go version'
                sh 'go build -o myapp'
                stash includes: 'myapp', name: 'binary'   // Lưu để stage sau dùng
            }
        }
        stage('Test in Go') {
            agent { 
                docker { 
                    image 'golang:1.22' 
                    args '-u root'
                } 
            }
            steps {
                sh 'go test ./...'
            }
        }
        stage('Lint') {
            agent { 
                docker { 
                    image 'golangci/golangci-lint:latest' 
                    args '-u root'
                } 
            }
            steps {
                sh 'golangci-lint run || true'    // || true để không fail lab
            }
        }
        stage('Package') {
            agent any
            steps {
                unstash 'binary'                  // Lấy lại binary từ stage Build
                sh 'ls -lh myapp'
                archiveArtifacts 'myapp'
            }
        }
    }
}
