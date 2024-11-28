pipeline {
    agent any  // 表示该任务在任何可用的 Jenkins 节点上运行

    environment {
        // 定义 Docker 镜像的名称
        FRONTEND_IMAGE = 'TSVuetes_frontend'
        BACKEND_IMAGE = 'TSVuetes_backend'
    }

    stages {
        stage('Checkout') {
            steps {
                // 拉取最新的代码
                checkout scm  // 这个会拉取 Git 仓库中的最新代码
            }
        }

        stage('Build Frontend') {
            steps {
                script {
                    // 在 Docker 容器中构建前端 Vue 项目
                    sh 'docker build -t $FRONTEND_IMAGE ./TSVuetes_frontend'
                }
            }
        }

        stage('Build Backend') {
            steps {
                script {
                    // 在 Docker 容器中构建后端 Go 项目
                    sh 'docker build -t $BACKEND_IMAGE ./TSVuetes_backend'
                }
            }
        }

        stage('Push Docker Images') {
            steps {
                script {
                    // 推送 Docker 镜像到 Docker Hub 或私有镜像仓库
                    sh 'docker push $FRONTEND_IMAGE'
                    sh 'docker push $BACKEND_IMAGE'
                }
            }
        }

        stage('Deploy Frontend') {
            steps {
                script {
                    // 停止并删除现有的前端容器
                    sh 'docker stop TSVuetes_frontend || true && docker rm TSVuetes_frontend || true'
                    // 启动新的前端容器
                    sh 'docker run -d --name TSVuetes_backend -p 3000:3000 $TSVuetes_backend'
                }
            }
        }

        stage('Deploy Backend') {
            steps {
                script {
                    // 停止并删除现有的后端容器
                    sh 'docker stop TSVuetes_backend || true && docker rm TSVuetes_backend || true'
                    // 启动新的后端容器
                    sh 'docker run -d --name TSVuetes_backend -p 8080:8080 $BACKEND_IMAGE'
                }
            }
        }
    }

    post {
        always {
            // 在构建后清理 Docker 镜像和容器
            sh 'docker system prune -f'
        }
    }
}