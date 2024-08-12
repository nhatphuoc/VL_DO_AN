FROM golang

EXPOSE 3000

# Set pwd to the go folder
WORKDIR /app

COPY ./common ./common
COPY ./schedule ./schedule
COPY ./environment ./environment
COPY ./food_drink ./food_drink
COPY ./gallery ./gallery
COPY ./home-data ./home-data
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum
COPY ./main.go ./main.go

# Start process
CMD ["go", "run", "main.go"]