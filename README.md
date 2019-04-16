## Sample Golang app
Sample Golang app for testing

git clone https://github.com/nithu0115/sample-go-app.git

#### Docker build
docker build -t sample-golang-app:latest .

docker run -d -p 80:8080 -e MSG_ENV=test sample-golang-app:latest
