#ifdef __linux__ //unix // It seems we are not in linux is unix __linux or __linux__

#include <stdio.h>

//Linux
FILE *crossplatformfmemopen (void *__s, size_t __len, const char *__modes){
    // FILE *crossplatformfmemopen (void *__s, unsigned long __len, const char *__modes) {
    return fmemopen (__s, __len, __modes);
}


#elif defined(__ANDROID__) 


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

FILE* crossplatformfmemopen(void* buf, size_t size, const char* mode) {
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
       
    }
}

#elif defined(__APPLE__) || defined (__ARM_ARCH) //|| defined (TARGET_OS_IPHONE)//MacOS iOS
//
// Copyright 2012 Jeff Verkoeyen
// Originally ported from https://github.com/ingenuitas/python-tesseract/blob/master/fmemopen.c
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/mman.h>

struct fmem {
  size_t pos;
  size_t size;
  char *buffer;
};
typedef struct fmem fmem_t;

static int readfn(void *handler, char *buf, int size) {
  fmem_t *mem = handler;
  size_t available = mem->size - mem->pos;
  
  if (size > available) {
    size = available;
  }
  memcpy(buf, mem->buffer, sizeof(char) * size);
  mem->pos += size;
  
  return size;
}

static int writefn(void *handler, const char *buf, int size) {
  fmem_t *mem = handler;
  size_t available = mem->size - mem->pos;

  if (size > available) {
    size = available;
  }
  memcpy(mem->buffer, buf, sizeof(char) * size);
  mem->pos += size;

  return size;
}

static fpos_t seekfn(void *handler, fpos_t offset, int whence) {
  size_t pos;
  fmem_t *mem = handler;

  switch (whence) {
    case SEEK_SET: pos = offset; break;
    case SEEK_CUR: pos = mem->pos + offset; break;
    case SEEK_END: pos = mem->size + offset; break;
    default: return -1;
  }

  if (pos > mem->size) {
    return -1;
  }

  mem->pos = pos;
  return (fpos_t)pos;
}

static int closefn(void *handler) {
  free(handler);
  return 0;
}

FILE *crossplatformfmemopen(void *buf, size_t size, const char *mode) {
  // This data is released on fclose.
  fmem_t* mem = (fmem_t *) malloc(sizeof(fmem_t));

  // Zero-out the structure.
  memset(mem, 0, sizeof(fmem_t));

  mem->size = size;
  mem->buffer = buf;

  // funopen's man page: https://developer.apple.com/library/mac/#documentation/Darwin/Reference/ManPages/man3/funopen.3.html
  return funopen(mem, readfn, writefn, seekfn, closefn);
}

#endif