// #ifdef __linux__ //unix // It seems we are not in linux is unix __linux or __linux__
#ifdef __WIN32__
#include <stdio.h>

//Linux
FILE *crossplatformfmemopen (void *__s, size_t __len, const char *__modes){
    // FILE *crossplatformfmemopen (void *__s, unsigned long __len, const char *__modes) {
    return fmemopen (__s, __len, __modes);
}


// #elif defined(ANDROID) || defined(__ANDROID__) || defined(__ANDROID_API__) //Android
#elif defined(__linux__)


#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <errno.h>

#define MEMSTREAM_BUFSIZE 1024

#define	__SWR	0x0008		/* OK to write */
#define	__SRD	0x0004		/* OK to read */
#define	__SRW	0x0010		/* open for reading & writing */
 */
typedef	struct __sFILE {
	unsigned char *_p;	/* current position in (some) buffer */
	int	_r;		/* read space left for getc() */
	int	_w;		/* write space left for putc() */
	short	_flags;		/* flags, below; this FILE is free if 0 */
	short	_file;		/* fileno, if Unix descriptor, else -1 */
	struct	__sbuf _bf;	/* the buffer (at least 1 byte, if !NULL) */
	int	_lbfsize;	/* 0 or -_bf._size, for inline putc */
	/* operations */
	void	*_cookie;	/* cookie passed to io functions */
	int	(*_close)(void *);
	int	(*_read)(void *, char *, int);
	fpos_t	(*_seek)(void *, fpos_t, int);
	int	(*_write)(void *, const char *, int);
	/* extension data, to avoid further ABI breakage */
	struct	__sbuf _ext;
	/* data for long sequences of ungetc() */
	unsigned char *_up;	/* saved _p when _p is doing ungetc data */
	int	_ur;		/* saved _r when _r is counting ungetc data */
	/* tricks to meet minimum requirements even when malloc() fails */
	unsigned char _ubuf[3];	/* guarantee an ungetc() buffer */
	unsigned char _nbuf[1];	/* guarantee a getc() buffer */
	/* separate buffer for fgetln() when line crosses buffer boundary */
	struct	__sbuf _lb;	/* buffer for fgetln() */
	/* Unix stdio files get aligned to block boundaries on fseek() */
	int	_blksize;	/* stat.st_blksize (may be != _bf._size) */
	fpos_t	_offset;	/* current lseek offset */
} FILE;

FILE	*__sfp(void);


struct memstream {
    char *buf;
    size_t size;
    size_t pos;
    int eof;
};

FILE *
funopen(const void *cookie, int (*readfn)(void *, char *, int),
	int (*writefn)(void *, const char *, int),
	fpos_t (*seekfn)(void *, fpos_t, int), int (*closefn)(void *))
{
	FILE *fp;
	int flags;
	if (readfn == NULL) {
		if (writefn == NULL) {		/* illegal */
			errno = EINVAL;
			return (NULL);
		} else
			flags = __SWR;		/* write only */
	} else {
		if (writefn == NULL)
			flags = __SRD;		/* read only */
		else
			flags = __SRW;		/* read-write */
	}
	if ((fp = __sfp()) == NULL)
		return (NULL);
	fp->_flags = flags;
	fp->_file = -1;
	fp->_cookie = (void *)cookie;		/* SAFE: cookie not modified */
	fp->_read = readfn;
	fp->_write = writefn;
	fp->_seek = seekfn;
	fp->_close = closefn;
	return (fp);
}


static int memstream_read(void *cookie, char *buf, int size)
{
    struct memstream *ms = (struct memstream *)cookie;
    int bytes_left = ms->size - ms->pos;
    int bytes_to_read = (size < bytes_left) ? size : bytes_left;

    if (bytes_to_read <= 0) {
        ms->eof = 1;
        return 0;
    }

    memcpy(buf, ms->buf + ms->pos, bytes_to_read);
    ms->pos += bytes_to_read;

    return bytes_to_read;
}

static int memstream_write(void *cookie, const char *buf, int size)
{
    struct memstream *ms = (struct memstream *)cookie;
    int bytes_left = MEMSTREAM_BUFSIZE - ms->pos;
    int bytes_to_write = (size < bytes_left) ? size : bytes_left;

    if (bytes_to_write <= 0) {
        return 0;
    }

    memcpy(ms->buf + ms->pos, buf, bytes_to_write);
    ms->pos += bytes_to_write;

    if (ms->pos > ms->size) {
        ms->size = ms->pos;
    }

    return bytes_to_write;
}

static fpos_t memstream_seek(void *cookie, fpos_t offset, int whence)
{
    struct memstream *ms = (struct memstream *)cookie;
    long pos;

    switch (whence) {
        case SEEK_SET:
            pos = offset;
            break;

        case SEEK_CUR:
            pos = ms->pos + offset;
            break;

        case SEEK_END:
            pos = ms->size + offset;
            break;

        default:
            return -1;
    }

    if (pos < 0 || pos > (long)ms->size) {
        return -1;
    }

    ms->pos = pos;

    return pos;
}

static int memstream_close(void *cookie)
{
    struct memstream *ms = (struct memstream *)cookie;

    free(ms->buf);
    free(ms);

    return 0;
}

FILE *fmemopen(void *buf, size_t size, const char *mode)
{
    struct memstream *ms;
    FILE *fp;
    int flags = 0;

    if (strchr(mode, 'r') != NULL) {
        flags |= 1;
    }

    if (strchr(mode, 'w') != NULL) {
        flags |= 2;
    }

    if (flags == 0) {
        return NULL;
    }

    ms = (struct memstream *)malloc(sizeof(struct memstream));
    if (ms == NULL) {
        return NULL;
    }

    ms->buf = (char *)buf;
    ms->size = size;
    ms->pos = 0;
    ms->eof = 0;

    fp = funopen(ms, memstream_read, memstream_write, memstream_seek, memstream_close);
    if (fp == NULL) {
        free(ms);
        return NULL;
    }

    if (flags == 1) {
        setbuf(fp, NULL);
    }

    return fp;
}





#elif defined(__APPLE__)//MacOS iOS
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