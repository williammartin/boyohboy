FROM golang:1.10 AS builder

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/github.com/williammartin/boyohboy
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /boyohboy .

FROM scratch
COPY --from=builder /boyohboy ./
ENTRYPOINT ["./boyohboy"]
