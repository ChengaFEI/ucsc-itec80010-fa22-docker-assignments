# Base image
FROM ubuntu:latest
# Directory
WORKDIR /dockerclass
# Executables
COPY frontend /dockerclass
COPY config.txt /dockerclass
# Permission
RUN chmod +x /dockerclass/frontend
# Run executables
ENTRYPOINT ["/dockerclass/frontend"]

