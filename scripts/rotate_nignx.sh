#!/bin/sh

# init var
log_dir=/var/log/nginx/
log_name=access.log
dest=access_`date +%Y%m%d_%H-%M-%S`.log 

#rotate log
sudo cp $"$log_dir$log_name" $dest
sudo cp $"$log_dir$log_name" "$log_dir$dest"
sudo rm $"$log_dir$log_name"

#restart nginx
if type systemctl >/dev/null 2>&1; then
  sudo systemctl restart nginx
else
  sudo service nginx restart
fi
