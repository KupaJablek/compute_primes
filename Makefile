BINARY_NAME = compute_primes
build:
	go build -o ${BINARY_NAME} .  

make clean:
	go clean
	rm ./BINARY_NAME