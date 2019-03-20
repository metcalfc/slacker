clean:
	rm -rf ./dist

snapshot:
	goreleaser --snapshot --skip-publish