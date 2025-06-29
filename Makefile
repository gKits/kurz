tw:
	@npx @tailwindcss/cli -i input.css -o ./public/styles.css --watch

dev:
	@go tool templ generate -watch -proxyport=4040 -proxy="http://localhost:4000" -open-browser=false -cmd="go run ./cmd/server/main.go"
