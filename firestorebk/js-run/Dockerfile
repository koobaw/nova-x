# Use the official Node.js image as a base image
FROM node:18

# Set the working directory in the container
WORKDIR /usr/src/app

# Copy package.json and package-lock.json to the working directory
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the job script to the working directory
COPY job.js .

# Define the command to run your application
CMD ["node", "job.js"]
