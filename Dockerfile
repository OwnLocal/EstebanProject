FROM ubuntu:16.04

# Update apt-get
RUN apt-get update -y

# Install Go
RUN apt-get install -y golang

# Bundle app source
COPY . /src

# Install app dependencies
WORKDIR /src

EXPOSE 80

CMD ./OwnLocal -p
