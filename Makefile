.PHONY: clean

bio: bio.go
	GOOS=linux GOARCH=amd64 go build -v bio.go

clean:
	rm -f bio
