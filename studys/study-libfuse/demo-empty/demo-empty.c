/*
 * main.c
 * Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com>
 *
 * Distributed under terms of the GPL license.
 */

#include <fuse/fuse.h>

int main(int argc, char *argv[])
{
	if (!fuse_main(argc, argv, NULL, NULL)){
		return -1;
	}
	return 0;
}
