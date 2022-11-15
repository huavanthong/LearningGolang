FROM node:latest AS node_builder

RUN mkdir -p /app/webreactjs

WORKDIR /app/webreactjs

COPY /app/package.json .

# Install npm 
RUN npm install

# Beacause node:latest does not use apk.
# Install bash for webreactjs container 
# Refer: [here](https://stackoverflow.com/questions/58210383/nodelatest-for-alpine-apk-not-found-because-sbin-is-not-on-path)
# RUN apk update && apk add bash

COPY . .

LABEL name="webreactjs" version="1.0"

EXPOSE 3000

CMD ["npm", "start"]