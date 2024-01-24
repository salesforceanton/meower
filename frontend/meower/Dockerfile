FROM node:21-alpine

ENV PATH /app/node_modules/.bin:$PATH
ENV CHOKIDAR_USEPOLLING true
ENV WDS_SOCKET_PORT 0

WORKDIR /app

COPY package*.json ./

RUN npm install --silent
RUN npm install nodemon --save-dev

COPY . ./

CMD ["nodemon", "--exec", "npm", "start"]