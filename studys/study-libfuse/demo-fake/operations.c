/*
 * operations.c
 * Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com>
 *
 * Distributed under terms of the GPL license.
 */

#include "operations.h"
#include <stdio.h>

int fake_getattr(const char * input, struct stat * stat){
	printf("input: %s\n", input);
	return 0;
}

int fake_mkdir(const char * input, mode_t mode){
	printf("input: %s\n", input);
	return 0;
}

int fake_rmdir(const char * input){
	printf("input: %s\n", input);
	return 0;
}

int fake_chmod(const char * input, mode_t mode){
	printf("input: %s\n", input);
	return 0;
}

int fake_chown(const char * input, uid_t uid, gid_t gid){
	printf("input: %s\n", input);
	return 0;
}

int fake_truncate(const char * input, off_t off){
	printf("input: %s\n", input);
	return 0;
}

int fake_open(const char * input, struct fuse_file_info * info){
	printf("input: %s\n", input);
	return 0;
}

int fake_read(const char * input, char *data, size_t size, off_t off, struct fuse_file_info * info){
	printf("input: %s\n", input);
	return 0;
}

int fake_write(const char * input, const char *data, size_t size, off_t off, struct fuse_file_info *info){
	printf("input: %s\n", input);
	return 0;
}

int fake_release(const char * input, struct fuse_file_info * info){
	printf("input: %s\n", input);
	return 0;
}

int fake_flush(const char * input, struct fuse_file_info * info){
	printf("input: %s\n", input);
	return 0;
}

int fake_fsync(const char * input, int sync, struct fuse_file_info * info){
	printf("input: %s\n", input);
	return 0;
}

int fake_opendir(const char * input, struct fuse_file_info * info){
	printf("input: %s\n", input);
	return 0;
}

int fake_readdir(const char * input, void * m, fuse_fill_dir_t dir, off_t off, struct fuse_file_info *info){
	printf("input: %s\n", input);
	return 0;
}

int fake_releasedir(const char * input, struct fuse_file_info *info){
	printf("input: %s\n", input);
	return 0;
}

int fake_access(const char * input, int m){
	printf("input: %s\n", input);
	return 0;
}
