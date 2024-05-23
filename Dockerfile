# Build frontend dist.
FROM node:22-alpine AS frontend
WORKDIR /frontend-build
COPY . .
RUN apk update && apk add bash
WORKDIR /frontend-build/web
RUN corepack enable && pnpm i
RUN pnpm build

# Use nginx for distrubution.
FROM nginx:alpine
RUN apk update && apk add bash

WORKDIR /www
COPY ./nginx.conf /etc/nginx/nginx.conf
COPY --from=frontend /frontend-build/web/dist /www

EXPOSE 8080

# ENTRYPOINT [ "/bin/bash" ]
CMD ["nginx", "-g", "daemon off;"]
