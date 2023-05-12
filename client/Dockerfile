FROM node:20-alpine AS build

WORKDIR /app

COPY package.json yarn.lock ./

RUN yarn install \
      --prefer-offline \
      --frozen-lockfile \
      --non-interactive \
      --production

COPY . .

RUN yarn build

FROM caddy:2.6.4-alpine AS release

COPY Caddyfile /etc/caddy/Caddyfile
COPY --from=build /app/dist /srv
