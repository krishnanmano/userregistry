.PHONY: list
list:
	@sh -c "$(MAKE) -p no_targets__ 2>/dev/null | \
        awk -F':' '/^[a-zA-Z0-9][^\$$#\/\\t=]*:([^=]|$$)/ {split(\$$1,A,/ /);for(i in A)print A[i]}' | \
        grep -v Makefile | \
        grep -v '%' | \
        grep -v '__\$$' | \
        sort -u"

MONGODB_VERSION=6.0-ubi8
.PHONY: startDb
startDb:
	docker run --name mongodb -d -p 27017:27017 mongodb/mongodb-community-server:$(MONGODB_VERSION)

.PHONY: stopDb
stopDb:
	docker stop mongodb && docker rm mongodb

.PHONY: genCoverageReport
genCoverageReport:
	go test -v -coverprofile cover.out ./...
	go tool cover -html cover.out -o cover.html

.PHONY: runTest
runTest:
	go test -v -cover ./...

.PHONY: runIntegrationTest
runIntegrationTest:
	go test -v userregistry/integration_test

.PHONY: cleanUp
cleanUp:
	rm -rf cover.html cover.out
