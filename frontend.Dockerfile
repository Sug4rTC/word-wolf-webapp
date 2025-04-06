FROM node:20-alpine

WORKDIR /app/frontend

COPY ./word-wolf-frontend/package.json ./word-wolf-frontend/package-lock.json ./
RUN npm install

COPY ./word-wolf-frontend ./

EXPOSE 5173

CMD ["npm", "run", "dev", "--", "--host"]