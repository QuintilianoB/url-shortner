FROM node:alpine

COPY frontend /frontend

WORKDIR /frontend

RUN npm install

EXPOSE 8080

CMD npm run serve