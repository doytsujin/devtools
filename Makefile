all: deps build

VERSION := v2

clean:
	rm -rf archive
	rm -rf bin/*
	rm -rf test.db

build:
	go build -o bin/aggregate ./cmd/aggregate/.
	go build -o bin/archive ./cmd/archive/.

archive: build
	./bin/archive -database="sqlite3" -connectionString="./test.db" -maxDurationMin=15 -archiveUrl="http://localhost:8100/%d-%02d-%02d-%d.json.gz"

docker-build:
	docker build --tag grafana/github-repo-metrics:${VERSION} .

docker-push:
	docker push grafana/github-repo-metrics:${VERSION}

archive-prod: build
	./bin/archive \
		-database="postgres" \
		-connectionString="postgresql://grafana:password@localhost:5432/grafana?sslmode=disable" \
		-maxDurationMin=5 \
		-archiveUrl="http://data.githubarchive.org/%d-%02d-%02d-%d.json.gz"

aggregate: build
	./bin/aggregate -database="sqlite3" -connectionString="./test2.db"

aggregate-prod: build
	./bin/aggregate \
		-database="postgres" \
		-connectionString="postgresql://grafana:password@localhost:5432/grafana?sslmode=disable" 