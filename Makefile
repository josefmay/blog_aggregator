run:
	go build -o out && ./out

migrate:
	cd sql/schema
	goose postgres "postgres://josefmay:@localhost:5432/blogator" up
	cd blog_aggregator