FROM node:20-alpine AS build

WORKDIR /app
ENV PATH /app/node_modules/.bin:$PATH

COPY package.json yarn.lock ./

RUN yarn install \
      --prefer-offline \
      --frozen-lockfile \
      --non-interactive

COPY . .

RUN yarn build

FROM caddy:2.6.4-alpine AS release

COPY Caddyfile /etc/caddy/Caddyfile
COPY --from=build /app/dist /srv
