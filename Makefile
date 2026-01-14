build:
	@echo "Building for all platforms..."
	# Windows
	GOOS=windows GOARCH=amd64 go build -o bin/todo-windows.exe .
	# Linux
	GOOS=linux GOARCH=amd64 go build -o bin/todo-linux .
	# macOS
	GOOS=darwin GOARCH=arm64 go build -o bin/todo-mac .
	@echo "Done! Check the /bin folder."
