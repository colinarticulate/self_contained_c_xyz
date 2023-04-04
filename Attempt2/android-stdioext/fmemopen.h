//This is a chatgpt suggestion, let's hope it works

#define _GNU_SOURCE
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <fcntl.h>

#define READ_END 0
#define WRITE_END 1

struct mem_file {
    char* buf;
    size_t size;
    size_t pos;
    int fd[2];
};

static int mem_read(void* cookie, char* buf, int size) {
    struct mem_file* mf = (struct mem_file*)cookie;
    int to_read = size < (mf->size - mf->pos) ? size : (mf->size - mf->pos);
    if (to_read > 0) {
        memcpy(buf, mf->buf + mf->pos, to_read);
        mf->pos += to_read;
    }
    return to_read;
}

static int mem_write(void* cookie, const char* buf, int size) {
    struct mem_file* mf = (struct mem_file*)cookie;
    int to_write = size < (mf->size - mf->pos) ? size : (mf->size - mf->pos);
    if (to_write > 0) {
        memcpy(mf->buf + mf->pos, buf, to_write);
        mf->pos += to_write;
    }
    return to_write;
}

static fpos_t mem_seek(void* cookie, fpos_t offset, int whence) {
    struct mem_file* mf = (struct mem_file*)cookie;
    fpos_t new_pos;
    switch (whence) {
        case SEEK_SET:
            new_pos = offset;
            break;
        case SEEK_CUR:
            new_pos = mf->pos + offset;
            break;
        case SEEK_END:
            new_pos = mf->size + offset;
            break;
        default:
            return -1;
    }
    if (new_pos < 0 || new_pos > mf->size) {
        return -1;
    }
    mf->pos = new_pos;
    return new_pos;
}

static int mem_close(void* cookie) {
    struct mem_file* mf = (struct mem_file*)cookie;
    close(mf->fd[READ_END]);
    close(mf->fd[WRITE_END]);
    free(mf->buf);
    free(mf);
    return 0;
}

FILE* fmemopen(void* buf, size_t size, const char* mode) {
    if (mode[0] != 'r' && mode[0] != 'w' && mode[0] != 'a') {
        return NULL;
    }
    if (mode[1] != '\0') {
        return NULL;
    }
    struct mem_file* mf = (struct mem_file*)malloc(sizeof(struct mem_file));
    if (mf == NULL) {
        return NULL;
    }
    mf->buf = (char*)buf;
    mf->size = size;
    mf->pos = 0;
    if (pipe(mf->fd) == -1) {
        free(mf);
        return NULL;
    }
    int flags = O_CLOEXEC;
    if (mode[0] == 'w') {
        flags |= O_WRONLY;
    } else if (mode[0] == 'a') {
        flags |= O_WRONLY | O_APPEND;
    } else {
        flags |= O_RDONLY;
    }
    int fd = fcntl(mf->fd[READ_END], F_DUPFD_CLOEXEC, 0);
    if (fd == -1) {
        close(mf->fd[READ_END]);
        close(mf->fd[WRITE_END]);
       
