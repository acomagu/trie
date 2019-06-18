.PHONY: bench
bench:
	curl https://dumps.wikimedia.org/enwiki/latest/enwiki-latest-all-titles-in-ns0.gz -o bench/enwiki-latest-all-titles-in-ns0.gz
	(cd bench && go test -bench . -benchmem .)
