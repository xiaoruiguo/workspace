#! /bin/sh
#
# prepare_env_mac.sh
# Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com>
#
# Distributed under terms of the GPL license.
#

sudo chown -R $(whoami) $(brew --prefix)/*
brew install node
