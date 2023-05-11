FROM node:19

COPY ./nodejs /code
COPY ./garden.json /data/garden.json

WORKDIR /code

ENV GARDENS_JSON_FILE_PATH=/data/garden.json
RUN npm install

EXPOSE 3000

CMD ["node", "src/server.mjs"]