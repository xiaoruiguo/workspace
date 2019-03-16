#!/usr/bin/perl
#
# decode_url.sh
# Copyright (C) 2018 lijiaocn <lijiaocn@foxmail.com>
#
# Distributed under terms of the GPL license.
#


 
use URI::Escape;
 
my $encodedurl = $ARGV[0];

my $url = uri_unescape($encodedurl);
 
print "$url\n";
