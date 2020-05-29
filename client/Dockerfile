#init client
FROM node:alpine 

WORKDIR /usr/app

COPY package*.json ./

# RUN npm install
# RUN npm ci -qy
RUN yarn install

COPY . .

EXPOSE 3000

CMD ["yarn", "start"]