.PHONY: version

all: build-server

# Generate a new semantic versioning number
version:
	@./semver.sh
	git commit -am "Version update"
	git push

