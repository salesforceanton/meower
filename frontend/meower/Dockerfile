FROM node:21-alpine

ENV PATH /app/node_modules/.bin:$PATH
ENV CHOKIDAR_USEPOLLING true

WORKDIR /app

COPY package*.json ./

RUN npm install --silent
RUN npm install nodemon --save-dev

COPY . ./

CMD ["nodemon", "--exec", "npm", "start"]