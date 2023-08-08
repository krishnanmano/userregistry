MONGODB_VERSION=6.0-ubi8

.PHONY: list
list:
	@sh -c "$(MAKE) -p no_targets__ 2>/dev/null | \
        awk -F':' '/^[a-zA-Z0-9][^\$$#\/\\t=]*:([^=]|$$)/ {split(\$$1,A,/ /);for(i in A)print A[i]}' | \
        grep -v Makefile | \
        grep -v '%' | \
        grep -v '__\$$' | \
        sort -u"

.PHONY: startdb
startdb:
	docker run --name mongodb -d -p 27017:27017 mongodb/mongodb-community-server:$(MONGODB_VERSION)

.PHONY: stopdb
stopdb:
	docker stop mongodb && docker rm mongodb
