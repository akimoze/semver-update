.PHONY: build
build:
	@for os in darwin linux windows; do \
		ext=""; \
		if [ "$$os" = "windows" ]; then ext=".exe"; fi; \
		for arch in amd64 arm64; do \
			GOOS=$$os GOARCH=$$arch go build -o "bin/semver-update-$$os-$$arch$$ext" ./main.go; echo "output bin/semver-update-$$os-$$arch$$ext" ; \
		done; \
	done

.PHONY: clean
clean:
	rm -rf bin/*
