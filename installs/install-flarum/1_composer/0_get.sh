#!/bin/bash
wget https://getcomposer.org/installer
php ./installer --check
php ./installer --install-dir=/opt/php-7.2.0/bin/ --filename=composer
