pipeline {
    agent any  // 表示该任务在任何可用的 Jenkins 节点上运行

    environment {
        // 定义 Docker 镜像的名称
        FRONTEND_IMAGE = 'tsvuetes_frontend'
        BACKEND_IMAGE = 'tsvuetes_backend'
        FRONTEND_DIR = './TSVuetes_frontend'  // 前端代码所在目录
    }

    stages {
        stage('Docker Login') {
            steps {
                // 使用 withCredentials 来安全地注入 Docker 凭证
                withCredentials([usernamePassword(credentialsId: 'docker-hub-credentials-id', usernameVariable: 'DOCKER_CREDENTIALS_USR', passwordVariable: 'DOCKER_CREDENTIALS_PSW')]) {
                    script {
                        // 使用注入的用户名和密码进行 Docker 登录
                        sh """
                            echo \$DOCKER_CREDENTIALS_PSW | docker login -u \$DOCKER_CREDENTIALS_USR --password-stdin
                        """
                    }
                }
            }
        }

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
                    sh 'docker build -t $FRONTEND_IMAGE $FRONTEND_DIR'
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
                    sh 'docker run -d --name TSVuetes_frontend -p 3000:3000 $FRONTEND_IMAGE'
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
