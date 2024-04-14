FROM nginx:latest

WORKDIR /usr/share/nginx/html

COPY index.html .
COPY script.js .

EXPOSE 122

CMD ["nginx", "-g", "daemon off;"]