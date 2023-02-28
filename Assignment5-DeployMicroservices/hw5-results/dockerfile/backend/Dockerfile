# Base image
FROM ubuntu:latest
# Directory
WORKDIR /dockerclass
# Executables
COPY backend /dockerclass
COPY config.txt /dockerclass
# Permission
RUN chmod +x /dockerclass/backend
# Run executables
ENTRYPOINT ["/dockerclass/backend"]

