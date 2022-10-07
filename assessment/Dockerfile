# Get Alpine Golang Image
FROM golang:1.19-alpine

# Use app directory as Workdir
WORKDIR /app

# Copy Go module
COPY go.mod ./
COPY go.sum ./

# Download module
RUN go mod download
# Copy Project file into workdir
COPY . ./
# build Go
RUN go build -o /e-commerce-go

EXPOSE 8080

#Execute Go Binary
CMD [ "/e-commerce-go" ]