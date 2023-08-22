run:
	@echo "Loading...."
	@go run main.go

push:
	git add . && git commit -m "update" && git push origin Farizzz