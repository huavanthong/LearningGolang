FROM node

# # dependencies
# RUN apt-get update && \
# 	apt-get -y --no-install-recommends install libnss3 \
# 	&& apt-get clean -qq && rm -rf /var/lib/apt/lists/*

RUN apt-get update && \
	apt-get -y --no-install-recommends install libgtkextra-dev libgconf2-dev libnss3 libasound2 \
	libxtst-dev libxss1 libdrm2 libgbm-dev libgtk-3-bin software-properties-common\
	&& apt-get clean -qq && rm -rf /var/lib/apt/lists/*
  
RUN mkdir -p /home/node/app/node_modules && chown -R node:node /home/node/app


WORKDIR /home/node/app

COPY package.json .

USER node
RUN npm install

COPY --chown=node:node . .

CMD ["npm", "start"]