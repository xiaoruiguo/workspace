/*
 * operations.h
 * Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com>
 *
 * Distributed under terms of the GPL license.
 */

#ifndef OPERATIONS_H
#define OPERATIONS_H

#include <sys/stat.h>
#include <fuse/fuse.h>

int fake_getattr(const char *, struct stat *);
int fake_mkdir(const char * input, mode_t mode);
int fake_rmdir(const char * input);
int fake_chmod(const char * input, mode_t mode);
int fake_chown(const char * input, uid_t uid, gid_t gid);
int fake_truncate(const char * input, off_t off);
int fake_open(const char * input, struct fuse_file_info * info);
int fake_read(const char *, char *, size_t, off_t off, struct fuse_file_info * info);
int fake_write(const char *, const char *, size_t size, off_t off, struct fuse_file_info *info);
int fake_release(const char *input, struct fuse_file_info * info);
int fake_flush(const char* input, struct fuse_file_info *);
int fake_fsync(const char * input, int sync, struct fuse_file_info * info);
int fake_opendir(const char * input, struct fuse_file_info * info);
int fake_readdir(const char * input, void * m, fuse_fill_dir_t dir, off_t off, struct fuse_file_info *info);
int fake_releasedir(const char * input, struct fuse_file_info *info);
int fake_access(const char * input, int m);

#endif /* !OPERATIONS_H */
