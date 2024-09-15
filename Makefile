# Build dependencies
GO = tinygo

# Whether to build for debugging instead of release
DEBUG = 0

# Compilation flags
GOFLAGS = -target ./target.json -panic trap
ifeq ($(DEBUG), 1)
	GOFLAGS += -opt 1
else
	GOFLAGS += -opt z -no-debug
endif

all:
	@mkdir -p build
	$(GO) build $(GOFLAGS) -o build/cart.wasm .

.PHONY: clean
clean:
	rm -rf build
