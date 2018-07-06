FROM debian:stretch-slim

RUN apt-get update && apt-get install -y ca-certificates

ENV DB "mysql"
ENV CONNSTRING "grafana:password@tcp(localhost:3306)/grafana?collation=utf8mb4_unicode_ci"
ENV ARCHIVE_URL "https://data.githubarchive.org/%d-%02d-%02d-%d.json.gz"
ENV BIN "archive"

ADD bin/archive /
ADD bin/aggregate /
ADD run.sh /
WORKDIR /
CMD ["sh", "-c", "/run.sh" ]
