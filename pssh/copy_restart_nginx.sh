pscp.pssh -H c2 -H c3 /etc/nginx/nginx.conf /home/isucon/nginx.conf
pssh -H c2 -H c3 -i -x -tt sudo mv ~/nginx.conf /etc/nginx/nginx.conf
pssh -H c2 -H c3 -i  -x -tt sudo service nginx restart
