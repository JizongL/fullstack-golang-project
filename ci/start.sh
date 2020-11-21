sudo yum update -y
sudo yum install -y golang
sudo yum install -y mysql
sudo amazon-linux-extras install nginx -y
sudo cp golang-app/nginx/golang-app.conf /etc/nginx/conf.d/golang-app.conf



