FROM golang

EXPOSE 3000

# Set pwd to the go folder
WORKDIR /app

COPY . .

# Start process
CMD ["sh", "-c", "sleep 10 && go run main.go"]