FROM golang:1.23-alpine AS base-stage
WORKDIR /app

FROM base-stage AS build-stage
COPY . .
RUN go build -o /app/app-backend .

FROM alpine AS production-stage
ENV BACKEND_PORT=8081
EXPOSE 8081
COPY --from=build-stage /app/app-backend /usr/local/bin/app-backend
CMD app-backend

FROM base-stage AS development-stage
ENV BACKEND_PORT=8081
EXPOSE 8081
VOLUME /app
RUN go install github.com/mitranim/gow@latest
COPY . .
CMD gow run .
