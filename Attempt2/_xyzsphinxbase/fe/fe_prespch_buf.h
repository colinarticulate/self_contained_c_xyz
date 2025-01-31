/* -*- c-basic-offset: 4; indent-tabs-mode: nil -*- */
/* ====================================================================
 * Copyright (c) 2013 Carnegie Mellon University.  All rights 
 * reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 *
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer. 
 *
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in
 *    the documentation and/or other materials provided with the
 *    distribution.
 *
 * This work was supported in part by funding from the Defense Advanced 
 * Research Projects Agency and the National Science Foundation of the 
 * United States of America, and the CMU Sphinx Speech Consortium.
 *
 * THIS SOFTWARE IS PROVIDED BY CARNEGIE MELLON UNIVERSITY ``AS IS'' AND 
 * ANY EXPRESSED OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, 
 * THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR
 * PURPOSE ARE DISCLAIMED.  IN NO EVENT SHALL CARNEGIE MELLON UNIVERSITY
 * NOR ITS EMPLOYEES BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 * SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT 
 * LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, 
 * DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY 
 * THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT 
 * (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE 
 * OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 *
 * ====================================================================
 *
 */

/* Buffer that maintains both features and raw audio for the VAD implementation */

#ifndef FE_INTERNAL_H
#define FE_INTERNAL_H

#include "xyzsphinxbase/ckd_alloc.h"
#include "xyzsphinxbase/err.h"
#include "xyzsphinxbase/fe.h"

typedef struct prespch_buf_s prespch_buf_t;

// /* Creates prespeech buffer */
// prespch_buf_t *fe_prespch_init(int num_frames, int num_cepstra,
//                                int num_samples);

// /* Reads mfcc frame from prespeech buffer */
// int fe_prespch_read_cep(prespch_buf_t * prespch_buf, mfcc_t * fea);

// /* Writes mfcc frame to prespeech buffer */
// void fe_prespch_write_cep(prespch_buf_t * prespch_buf, mfcc_t * fea);

// /* Reads pcm frame from prespeech buffer */
// void fe_prespch_read_pcm(prespch_buf_t * prespch_buf, int16 *samples,
//                          int32 * samples_num);

// /* Writes pcm frame to prespeech buffer */
// void fe_prespch_write_pcm(prespch_buf_t * prespch_buf, int16 * samples);

// /* Resets read/write pointers for cepstrum buffer */
// void fe_prespch_reset_cep(prespch_buf_t * prespch_buf);

// /* Resets read/write pointer for pcm audio buffer */
// void fe_prespch_reset_pcm(prespch_buf_t * prespch_buf);

// /* Releases prespeech buffer */
// void fe_prespch_free(prespch_buf_t * prespch_buf);

// /* Returns number of accumulated frames */
// int32 fe_prespch_ncep(prespch_buf_t * prespch_buf);


struct prespch_buf_s {
    /* saved mfcc frames */
    mfcc_t **cep_buf;
    /* saved pcm audio */
    int16 *pcm_buf;

    /* flag for pcm buffer initialization */
    int16 cep_write_ptr;
    /* read pointer for cep buffer */
    int16 cep_read_ptr;
    /* Count */
    int16 ncep;
    

    /* flag for pcm buffer initialization */
    int16 pcm_write_ptr;
    /* read pointer for cep buffer */
    int16 pcm_read_ptr;
    /* Count */
    int16 npcm;
    
    /* frames amount in cep buffer */
    int16 num_frames_cep;
    /* frames amount in pcm buffer */
    int16 num_frames_pcm;
    /* filters amount */
    int16 num_cepstra;
    /* amount of fresh samples in frame */
    int16 num_samples;
};

prespch_buf_t *
fe_prespch_init(int num_frames, int num_cepstra, int num_samples)
{
    prespch_buf_t *prespch_buf;

    prespch_buf = (prespch_buf_t *) ckd_calloc(1, sizeof(prespch_buf_t));

    prespch_buf->num_cepstra = num_cepstra;
    prespch_buf->num_frames_cep = num_frames;
    prespch_buf->num_samples = num_samples;
    prespch_buf->num_frames_pcm = 0;

    prespch_buf->cep_write_ptr = 0;
    prespch_buf->cep_read_ptr = 0;
    prespch_buf->ncep = 0;
    
    prespch_buf->pcm_write_ptr = 0;
    prespch_buf->pcm_read_ptr = 0;
    prespch_buf->npcm = 0;

    prespch_buf->cep_buf = (mfcc_t **)
        ckd_calloc_2d(num_frames, num_cepstra,
                      sizeof(**prespch_buf->cep_buf));

    prespch_buf->pcm_buf = (int16 *)
        ckd_calloc(prespch_buf->num_frames_pcm * prespch_buf->num_samples,
                   sizeof(int16));

    return prespch_buf;
}


int
fe_prespch_read_cep(prespch_buf_t * prespch_buf, mfcc_t * feat)
{
    if (prespch_buf->ncep == 0)
        return 0;
    memcpy(feat, prespch_buf->cep_buf[prespch_buf->cep_read_ptr],
           sizeof(mfcc_t) * prespch_buf->num_cepstra);
    prespch_buf->cep_read_ptr = (prespch_buf->cep_read_ptr + 1) % prespch_buf->num_frames_cep;
    prespch_buf->ncep--;
    return 1;
}

void
fe_prespch_write_cep(prespch_buf_t * prespch_buf, mfcc_t * feat)
{
    memcpy(prespch_buf->cep_buf[prespch_buf->cep_write_ptr], feat,
           sizeof(mfcc_t) * prespch_buf->num_cepstra);
    prespch_buf->cep_write_ptr = (prespch_buf->cep_write_ptr + 1) % prespch_buf->num_frames_cep;
    if (prespch_buf->ncep < prespch_buf->num_frames_cep) {
        prespch_buf->ncep++;	
    } else {
        prespch_buf->cep_read_ptr = (prespch_buf->cep_read_ptr + 1) % prespch_buf->num_frames_cep;
    }
}

void
fe_prespch_read_pcm(prespch_buf_t * prespch_buf, int16 *samples,
                    int32 *samples_num)
{
    int i;
    int16 *cursample = samples;
    *samples_num = prespch_buf->npcm * prespch_buf->num_samples;
    for (i = 0; i < prespch_buf->npcm; i++) {
	memcpy(cursample, &prespch_buf->pcm_buf[prespch_buf->pcm_read_ptr * prespch_buf->num_samples],
	           prespch_buf->num_samples * sizeof(int16));
	prespch_buf->pcm_read_ptr = (prespch_buf->pcm_read_ptr + 1) % prespch_buf->num_frames_pcm;
    }
    prespch_buf->pcm_read_ptr = 0;
    prespch_buf->pcm_write_ptr = 0;    
    prespch_buf->npcm = 0;
    return;
}

void
fe_prespch_write_pcm(prespch_buf_t * prespch_buf, int16 * samples)
{
    int32 sample_ptr;

    sample_ptr = prespch_buf->pcm_write_ptr * prespch_buf->num_samples;
    memcpy(&prespch_buf->pcm_buf[sample_ptr], samples,
           prespch_buf->num_samples * sizeof(int16));

    prespch_buf->pcm_write_ptr = (prespch_buf->pcm_write_ptr + 1) % prespch_buf->num_frames_pcm;
    if (prespch_buf->npcm < prespch_buf->num_frames_pcm) {
        prespch_buf->npcm++;	
    } else {
        prespch_buf->pcm_read_ptr = (prespch_buf->pcm_read_ptr + 1) % prespch_buf->num_frames_pcm;
    }
}

void
fe_prespch_reset_cep(prespch_buf_t * prespch_buf)
{
    prespch_buf->cep_read_ptr = 0;
    prespch_buf->cep_write_ptr = 0;
    prespch_buf->ncep = 0;
}

void
fe_prespch_reset_pcm(prespch_buf_t * prespch_buf)
{
    prespch_buf->pcm_read_ptr = 0;
    prespch_buf->pcm_write_ptr = 0;
    prespch_buf->npcm = 0;
}

void
fe_prespch_free(prespch_buf_t * prespch_buf)
{
    if (!prespch_buf)
	return;
    if (prespch_buf->cep_buf)
        ckd_free_2d((void **) prespch_buf->cep_buf);
    if (prespch_buf->pcm_buf)
        ckd_free(prespch_buf->pcm_buf);
    ckd_free(prespch_buf);
}

int32 
fe_prespch_ncep(prespch_buf_t * prespch_buf)
{
    return prespch_buf->ncep;
}


#endif                          /* FE_INTERNAL_H */
