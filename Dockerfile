FROM golang:1.16-alpine AS gobuild
WORKDIR /src
COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY  . ./

RUN go build

FROM node:16-alpine3.11 AS angularbuild
WORKDIR /usr/src/app
COPY ./ui/package*.json ./
RUN npm ci
COPY ./ui .
RUN npm run build

FROM golang:1.16-alpine
WORKDIR /app

COPY --from=angularbuild /usr/src/static ./static
COPY --from=gobuild /src/golculator .
COPY --from=gobuild /src/operations ./operations

CMD ["/app/golculator", "seed"]




