FROM ubuntu:14.04

# Update apt-get
RUN apt-get update

# Install Go
RUN apt-get install -y golang

# Install Elasticsearch
RUN apt-get install -y elasticsearch

# Bundle app source
COPY . /src

# Install app dependencies
WORKDIR /src

RUN ./OwnLocal -p -s

EXPOSE 80

CMD /opt/elasticsearch/bin/elasticsearch && ./OwnLocal -p
