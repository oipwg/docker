FROM node:latest

RUN git clone https://github.com/oipwg/oip-hdmw-interface /usr/src/webwallet
WORKDIR /usr/src/webwallet
RUN npm install
RUN ./node_modules/.bin/next build

EXPOSE 7000
CMD [ "npm", "start" ]