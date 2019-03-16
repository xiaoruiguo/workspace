#! /bin/sh
#
# clone.sh
# Copyright (C) 2018 lijiaocn <lijiaocn@foxmail.com>
#
# Distributed under terms of the GPL license.
#

#On Mac
brew install node watchman yarn
yarn global add bower

git clone 'https://github.com/rancher/ui'

# cd 'ui'
# ./scripts/update-dependencies
# yarn start
