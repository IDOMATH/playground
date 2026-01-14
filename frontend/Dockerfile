# Stage 1: Build the Angular app
FROM node:20-alpine AS build
WORKDIR /app

# Copy package files and install dependencies
COPY package*.json ./
RUN npm install

# Copy the rest of the application code
COPY . .
RUN npm run build -- --configuration production

# Stage 2: Serve the app with Nginx
FROM nginx:stable-alpine
# Copy the build output to replace the default nginx contents.
# Note: Replace 'your-app-name' with the actual name from angular.json
COPY --from=build /app/dist/frontend/browser /usr/share/nginx/html

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]