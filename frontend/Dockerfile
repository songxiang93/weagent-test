FROM node:18-alpine AS builder

WORKDIR /app

COPY package*.json ./

# Add --no-optional flag to avoid optional dependencies issues
RUN npm install --no-optional

COPY . .

RUN npm run build

FROM nginx:alpine

COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]