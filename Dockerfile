# Use the official Golang base image
FROM golang:1.17-alpine

# Set build arguments with default values
ARG DVWA_URL="https://192.168.4.10"
ARG DVWA_HOST="192.168.4.10"
ARG SHOP_URL="https://shop.corp.fabriclab.ca"
ARG FWB_URL="https://192.168.4.10/fwb/"
ARG SPEEDTEST_URL="http://speedtest.corp.fabriclab.ca"
ARG KALI_URL="https://flbr1kali01.fortiweb.fabriclab.ca"
ARG TOKEN="eyJ1c2VybmFtZSI6InVzZXJhcGkiLCJwYXNzd29yZCI6ImZhY2VMT0NLeWFybjY3ISJ9Cg=="
ARG FWB_MGT_IP="192.168.4.2"
ARG POLICY="DVWA_POLICY"
ARG USER_AGENT="FortiWeb Demo Tool"

# Install required packages; curl, jq
RUN apk update && apk add --no-cache curl jq

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code to the container
COPY . .

# Build the Go application with build arguments
RUN go build -ldflags="-X 'main.DVWA_URL=$DVWA_URL' \
    -X 'main.DVWA_HOST=$DVWA_HOST' \
    -X 'main.SHOP_URL=$SHOP_URL' \
    -X 'main.FWB_URL=$FWB_URL' \
    -X 'main.SPEEDTEST_URL=$SPEEDTEST_URL' \
    -X 'main.KALI_URL=$KALI_URL' \
    -X 'main.TOKEN=$TOKEN' \
    -X 'main.FWB_MGT_IP=$FWB_MGT_IP' \
    -X 'main.POLICY=$POLICY' \
    -X 'main.USER_AGENT=$USER_AGENT'" \
    -o darwin

# Set the container command to run the Go application
CMD ["./darwin"]
