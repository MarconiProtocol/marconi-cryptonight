# Normally you could just run "go build" or "go test" from within the
# marconi-cryptonight directory. But since the C code uses some
# instructions from Intel's AES instruction set, we need to whitelist
# the -maes compiler flag so that the golang compiler doesn't
# complain.

# We list the 'test' target first so it becomes the default if the
# user just runs 'make', since the 'test' target is probably more
# useful than the 'build' target, which only checks if the code
# compiles successfully.
test:
	CGO_CFLAGS_ALLOW=-maes go test -v

build:
	CGO_CFLAGS_ALLOW=-maes go build -v
