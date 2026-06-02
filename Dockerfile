FROM node:20-alpine AS deps

WORKDIR /app

COPY package*.json ./
RUN npm ci --ignore-scripts

FROM deps AS build

COPY nest-cli.json tsconfig*.json ./
COPY src ./src
RUN npm run build

FROM node:20-alpine AS runtime

WORKDIR /app

ENV NODE_ENV=production

COPY package*.json ./
RUN npm ci --omit=dev --ignore-scripts \
  && npm cache clean --force

COPY --from=build /app/dist ./dist

EXPOSE 5000

USER node

CMD ["node", "dist/main.js"]
