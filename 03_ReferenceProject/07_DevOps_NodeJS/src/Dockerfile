FROM node:14

LABEL maintainer="Akın BİÇER <akin.bicer@outlook.com.tr>"
LABEL version="1.0"
LABEL description="This project was created to better understand DevOps with Docker and Kubernetes examples. The main purpose of the application is to see the IP address and machine name of the machine on which it is running."

EXPOSE 3000

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm install --only=production

COPY . .

CMD ["node", "index.js"]
