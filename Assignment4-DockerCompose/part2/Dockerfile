# Base image
FROM alpine:latest
# Directory
WORKDIR /scripts
# Shell script
COPY . .
# Add execute permission
RUN chmod +x fortune.sh
# Run command
CMD ["./fortune.sh"]

