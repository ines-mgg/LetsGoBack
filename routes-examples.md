# pour /upload

curl -F "file=@main.go" http://localhost:8080/web/upload
curl -F "file=@main.go" http://localhost:8080/web/upload -v

# pour /upload-multiple

curl -F "files=@main.go" -F "files=@go.mod" http://localhost:8080/upload-multiple

# Type MIME interdit

curl -F "file=@logo.png" http://localhost:8080/web/upload -v

# Trop gros

dd if=/dev/zero of=big.txt bs=1M count=10
curl -F "file=@big.txt" http://localhost:8080/web/upload -v
