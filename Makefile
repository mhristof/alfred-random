MAKEFLAGS += --warn-undefined-variables
SHELL := bash
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := zip
.ONESHELL:


IMAGE := alfred-random

bin/alfred-random: $(shell find ./ -name '*.go')
	go build -o bin/alfred-random main.go


zip: info.plist Makefile alfred-random.sh bin/alfred-random 9A4E0276-BFF2-4386-9699-AB65E494C7B8.png icon.png
	zip -r $(IMAGE).alfredworkflow $^

clean:
	rm -rf $(IMAGE).alfredworkflow
.PHONY: clean

.PHONY: minor
minor:  ## Create a minor git tag and push it
	sed -i "" 's/$(shell semver current | tr -d 'v' )/$(shell semver -n | rev | cut -d ' ' -f1 | rev | tr -d 'v')/' info.plist
	make commitVersion
	semver
	git push --tags

.PHONY: patch
patch:  ## Create a patch git tag and push it
	sed -i "" 's/$(shell semver current | tr -d 'v' )/$(shell semver -p -n | rev | cut -d ' ' -f1 | rev | tr -d 'v')/' info.plist
	make commitVersion
	semver --patch
	git push --tags

.PHONY: commitVersion
commitVersion:
	git add info.plist
	git commit -m 'bumped version' info.plist

.PHONY:
help:           ## Show this help.
	@grep '.*:.*##' Makefile | grep -v grep  | sort | sed 's/:.* ##/:/g' | column -t -s:
