sudo yum update -y
sudo yum install -y golang
sudo yum install -y mysql
sudo amazon-linux-extras install nginx -y
sudo cp golang-app/nginx/golang-app.conf /etc/nginx/conf.d/golang-app.conf
cd /tmp
wget -O epel.rpm –nv https://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm
sudo yum install -y ./epel.rpm
sudo yum install python2-certbot.noarch
sudo yum install python-certbot-nginx
sudo certbot --nginx

