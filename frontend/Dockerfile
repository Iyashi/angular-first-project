FROM node:20-alpine AS base-stage
LABEL author="Jannik Lohaus <jannik@lo.haus>" \
  version="1.0" \
  description="Angular container for my first angluar project"
WORKDIR /app

FROM base-stage AS install-stage
COPY package*.json ./
RUN npm install --silent

FROM install-stage AS development-stage
ENV NODE_ENV=development
ENV FRONTEND_PORT=8080
EXPOSE 8080
VOLUME /app
COPY . .
CMD npm run start -- --host 0.0.0.0 --port ${FRONTEND_PORT} --public-host angular-first-project.local --hmr
#CMD ["npm", "run", "start", "--", "--port", "${FRONTEND_PORT}", "--public-host", "angular-first-project.local", "--hmr"]

FROM install-stage AS build-stage
ENV NODE_ENV=production
COPY . .
RUN npm run build

FROM nginx:alpine AS production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html
