# 使用官方 Nginx 镜像作为基础镜像
FROM nginx:stable-perl

# 删除 Nginx 默认的静态页面
RUN rm -rf /usr/share/nginx/html/*

# 将打包后的 Vue 静态文件复制到 Nginx 默认的目录下
COPY ./dist /usr/share/nginx/html

# 暴露 80 端口，以便外界访问 Nginx 服务
EXPOSE 80

# 启动 Nginx
CMD ["nginx", "-g", "daemon off;"]
