#!/bin/sh
php-cgi -b 127.0.0.1:9000
sudo systemctl start nginx
