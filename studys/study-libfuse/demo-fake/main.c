/*
 * main.c
 * Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com>
 *
 * Distributed under terms of the GPL license.
 */

#include <fuse/fuse.h>
#include "operations.h"

const struct fuse_operations fake_ops = {
	.getattr = fake_getattr,
	.readlink = NULL,
	.getdir = NULL,
	.mknod = NULL,
	.mkdir = fake_mkdir,
	.unlink = NULL,
	.rmdir = fake_rmdir,
	.symlink = NULL,
	.rename = NULL,
	.link = NULL,
	.chmod = fake_chmod,
	.chown = fake_chown,
	.truncate = fake_truncate,
	.utime = NULL,

	.open = fake_open,
	.read = fake_read,
	.release = fake_release,
	.write = fake_write,

	.statfs = NULL,
	.flush = fake_flush,
	.fsync = fake_fsync,

	.setxattr = NULL,
	.getxattr = NULL,
	.listxattr = NULL,
	.removexattr = NULL,

	.opendir = fake_opendir,
	.readdir = fake_readdir,
	.releasedir = fake_releasedir,

	.fsyncdir = NULL,
	.init = NULL,
	.destroy = NULL,
	.access = fake_access,
	.create = NULL,
	.ftruncate = NULL,
	.fgetattr = NULL,
};

int main(int argc, char *argv[])
{
	if (!fuse_main(argc, argv, &fake_ops, NULL)){
		return -1;
	}
	return 0;
}
