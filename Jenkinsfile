pipeline {
    agent any // 表示该任务在任何可用的 Jenkins 节点上运行

    environment {
        // 获取之前创建的 Docker 凭证 ID（例如：docker-hub-credentials-id）
        DOCKER_CREDENTIALS = credentials('docker-hub-credentials-id')
    }

    environment {
        // 定义 Docker 镜像的名称
        FRONTEND_IMAGE = 'tsvuetes_frontend'
        BACKEND_IMAGE = 'tsvuetes_backend'
        FRONTEND_DIR = './TSVuetes_frontend' // 前端代码所在目录

    }


    stages {
        stage('Docker Login') {
            steps {
                script {
                    // 使用凭证登录 Docker
                    sh "echo ${DOCKER_CREDENTIALS_USR}:${DOCKER_CREDENTIALS_PSW} | docker login -u ${DOCKER_CREDENTIALS_USR} --password-stdin"
                }
            }
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
                    // 确保 npm 安装了所有依赖
                    sh "cd $FRONTEND_DIR && npm install"
                    // 执行前端构建，生成 dist 目录
                    sh "cd $FRONTEND_DIR && npm run build"
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
                    sh 'docker run -d --name TSVuetes_backend -p 8081:8080 $BACKEND_IMAGE'
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