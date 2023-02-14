#include <string>


#include "xyzsphinxbase/include/xyzsphinxbase/err.h"



const char *params125[] = {
"pocketsphinx_continuous",
"-alpha", "0.97",
"-backtrace", "yes",
"-beam", "1e-10000",
"-bestpath", "no",
"-cmn", "live",
//"-cmninit", "48.66,4.31,-7.12,5.61,-1.63,9.01,-4.65,-17.99,-16.52,-5.18,3.45,2.53,-1.34", //batch with nfft and wlen
"-cmninit", "54.97,4.93,-7.22,5.18,-1.72,9.32,-4.26,-18.37,-17.32,-6.05,2.84,1.84,-1.61",
"-dict", "/home/dbarbera/Data/art_db.phone",
"-dither", "no",
"-doublebw", "no",
"-featparams", "/usr/local/share/xyzpocketsphinx/model/art-en-us/en-us/feat.params",
"-frate", "125",
"-fsgusefiller", "no",
"-fwdflat", "no",
"-hmm","/usr/local/share/xyzpocketsphinx/model/art-en-us/en-us",
"-infile", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/allowed1_philip_fixed_trimmed.wav",
"-jsgf", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/forced_align_a3ecf04d-a77a-4269-9eb5-395f8dfbdd8a_allowed1_philip_fixed_trimmed.wav.jsgf",
"-logfn", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/log/e1e0d844-812b-496c-83fb-712de847f8a7_a3ecf04d_frate_125_debug_from_c_.log",
//"-logfn","/dev/null",
"-lpbeam", "1e-10000",
"-lponlybeam", "1e-10000",
"-lw", "6",
"-maxhmmpf", "-1",
"-maxwpf", "-1",
"-nfft", "256",
"-nwpen", "1",
"-pbeam", "1e-10000",
"-pip", "1",
"-pl_window", "0",
"-remove_dc", "no",
"-remove_noise", "yes",
"-remove_silence", "no",
"-topn", "4",
"-vad_postspeech", "25",
"-vad_prespeech", "5",
"-vad_startspeech", "8",
"-vad_threshold", "1",
"-wbeam", "1e-10000",
"-wip", "0.5",
"-wlen", "0.016"
};

const int params125_size=77;

const char *params72[] = {
"pocketsphinx_continuous",
"-alpha", "0.97",
"-backtrace", "yes",
"-beam", "1e-10000",
"-bestpath", "no",
"-cmn", "live",
//"-cmninit", "61.61,8.03,-6.54,4.13,-3.74,9.61,-5.77,-16.52,-13.85,-3.98,2.30,2.59,-1.94", //batch with nfft and wlen
"-cmninit", "60.41,7.96,-6.64,4.16,-3.60,9.51,-5.71,-16.39,-13.69,-3.84,2.33,2.42,-2.09",
"-dict", "/home/dbarbera/Data/art_db.phone",
"-dither", "yes",
"-doublebw", "yes",
"-featparams", "/usr/local/share/xyzpocketsphinx/model/art-en-us/en-us/feat.params",
"-frate", "72",
"-fsgusefiller", "no",
"-fwdflat", "no",
"-hmm","/usr/local/share/xyzpocketsphinx/model/art-en-us/en-us",
"-infile", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/allowed1_philip_fixed_trimmed.wav",
"-jsgf", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/forced_align_2668db47-d3ce-4760-ab4b-60b9b8a6c46e_allowed1_philip_fixed_trimmed.wav.jsgf",
"-logfn", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/log/d0c65f23-d9cc-4047-8f3c-3a91db3623ff_2668db47_frate_72_debug_from_c_.log",
//"-logfn","/dev/null",
"-lpbeam", "1e-10000",
"-lponlybeam", "1e-10000",
"-lw", "6",
"-maxhmmpf", "-1",
"-maxwpf", "-1",
"-nfft", "512",
"-nfilt", "25",
"-nwpen", "1",
"-pbeam", "1e-10000",
"-pip", "1.15",
"-pl_window", "0",
"-remove_dc", "no",
"-remove_noise", "yes",
"-remove_silence", "yes",
"-topn", "6",
"-vad_postspeech", "20",
"-vad_prespeech", "5",
"-vad_startspeech", "5",
"-vad_threshold", "1.5",
"-wbeam", "1e-10000",
"-wip", "0.25",
"-wlen", "0.032"
};
const int params72_size=77;

const char *params80[] = {
"pocketsphinx_continuous",
"-alpha", "0.97",
"-backtrace", "yes",
"-beam", "1e-10000",
"-bestpath", "no",
"-cmn", "live",
//"-cmninit", "60.83,7.43,-6.07,4.54,-3.88,9.48,-5.85,-16.81,-13.89,-3.82,2.39,2.53,-1.88", //batch with nfft and wlen
"-cmninit","60.34,7.33,-6.08,4.56,-3.81,9.44,-5.75,-16.69,-13.76,-3.75,2.37,2.45,-1.91",
"-dict", "/home/dbarbera/Data/art_db.phone",
"-dither", "yes",
"-doublebw", "yes",
"-featparams", "/usr/local/share/xyzpocketsphinx/model/art-en-us/en-us/feat.params",
"-frate", "80",
"-fsgusefiller", "no",
"-fwdflat", "no",
"-hmm","/usr/local/share/xyzpocketsphinx/model/art-en-us/en-us",
"-infile", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/allowed1_philip_fixed_trimmed.wav",
"-jsgf", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/forced_align_4311b957-22a7-446c-85d9-d154d4156d02_allowed1_philip_fixed_trimmed.wav.jsgf",
"-logfn", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/log/4f0e00fa-9d20-4096-8bb5-8aeedc110e52_4311b957_frate_80_debug_from_c_.log",
//"-logfn","/dev/null",
"-lpbeam", "1e-10000",
"-lponlybeam", "1e-10000",
"-lw", "6",
"-maxhmmpf", "-1",
"-maxwpf", "-1",
"-nfft", "512",
"-nwpen", "1",
"-pbeam", "1e-10000",
"-pip", "1.15",
"-pl_window", "0",
"-remove_dc", "no",
"-remove_noise", "yes",
"-remove_silence", "yes",
"-topn", "6",
"-vad_postspeech", "20",
"-vad_prespeech", "5",
"-vad_startspeech", "5",
"-vad_threshold", "1.5",
"-wbeam", "1e-10000",
"-wip", "0.25",
"-wlen", "0.028"
};
const int params80_size=77;

const char *params91[] = {
"pocketsphinx_continuous",
"-alpha", "0.97",
"-backtrace", "yes",
"-beam", "1e-10000",
"-bestpath", "no",
"-cmn", "live",
//"-cmninit", "55.22,5.33,-6.78,5.07,-2.13,9.10,-4.03,-17.60,-16.77,-6.25,2.43,1.62,-1.56", //batch with nfft and wlen
"-cmninit","55.57,5.37,-6.76,5.06,-1.99,9.18,-4.02,-17.64,-16.84,-6.30,2.55,1.68,-1.59",
"-dict", "/home/dbarbera/Data/art_db.phone",
"-dither", "no",
"-doublebw", "no",
"-featparams", "/usr/local/share/xyzpocketsphinx/model/art-en-us/en-us/feat.params",
"-frate", "91",
"-fsgusefiller", "no",
"-fwdflat", "no",
"-hmm","/usr/local/share/xyzpocketsphinx/model/art-en-us/en-us",
"-infile", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/allowed1_philip_fixed_trimmed.wav",
"-jsgf", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/forced_align_1eed3902-f7e5-444b-a4b8-29b5c47ea52e_allowed1_philip_fixed_trimmed.wav.jsgf",
"-logfn", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/log/f7a95619-d7c5-42a9-b548-561187b350da_1eed3902_frate_91_debug_from_c_.log",
//"-logfn","/dev/null",
"-lpbeam", "1e-10000",
"-lponlybeam", "1e-10000",
"-lw", "6",
"-maxhmmpf", "-1",
"-maxwpf", "-1",
"-nfft", "512",
"-nwpen", "1",
"-pbeam", "1e-10000",
"-pip", "1",
"-pl_window", "0",
"-remove_dc", "no",
"-remove_noise", "yes",
"-remove_silence", "no",
"-topn", "4",
"-vad_postspeech", "20",
"-vad_prespeech", "5",
"-vad_startspeech", "5",
"-vad_threshold", "0.5",
"-wbeam", "1e-10000",
"-wip", "0.5",
"-wlen", "0.024"
};
const int params91_size=77;

const char *params105[] = {
"pocketsphinx_continuous",
"-alpha", "0.97",
"-backtrace", "yes",
"-beam", "1e-10000",
"-bestpath", "no",
"-cmn", "live",
//"-cmninit", "53.92,4.73,-7.01,5.40,-1.92,8.97,-4.24,-17.95,-17.00,-6.15,2.58,1.61,-1.69",//batch with nftt and wlen
"-cmninit","55.29,4.86,-6.97,5.07,-1.77,9.28,-4.13,-18.09,-17.15,-6.12,2.82,1.79,-1.60",
"-dict", "/home/dbarbera/Data/art_db.phone",
"-dither", "no",
"-doublebw", "no",
"-featparams", "/usr/local/share/xyzpocketsphinx/model/art-en-us/en-us/feat.params",
"-frate", "105",
"-fsgusefiller", "no",
"-fwdflat", "no",
"-hmm","/usr/local/share/xyzpocketsphinx/model/art-en-us/en-us",
"-infile", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/allowed1_philip_fixed_trimmed.wav",
"-jsgf", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/forced_align_9ddb7131-fa08-4bc4-b44c-814b2ed9917e_allowed1_philip_fixed_trimmed.wav.jsgf",
"-logfn", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/log/c061830a-8106-4e06-ae16-18feb072ea45_9ddb7131_frate_105_debug_from_c_.log",
//"-logfn","/dev/null",
"-lpbeam", "1e-10000",
"-lponlybeam", "1e-10000",
"-lw", "6",
"-maxhmmpf", "-1",
"-maxwpf", "-1",
"-nfft", "512",
"-nwpen", "1",
"-pbeam", "1e-10000",
"-pip", "1",
"-pl_window", "0",
"-remove_dc", "no",
"-remove_noise", "yes",
"-remove_silence", "no",
"-topn", "4",
"-vad_postspeech", "20",
"-vad_prespeech", "5",
"-vad_startspeech", "5",
"-vad_threshold", "1.5",
"-wbeam", "1e-10000",
"-wip", "0.5",
"-wlen", "0.020"
};
const int params105_size=77;

//--------------------------- BATCH -------------------------

const char *batch_params72[] = {
"pocketsphinx_batch",
"-adcin", "yes",
"-alpha", "0.97",
"-beam", "1e-10000",
"-bestpath", "no",
"-cepdir", "/home/dbarbera/Repositories/c_xyx/test_data/allowed1_philip_allowed/audio",
"-cepext", ".wav",
//"-cmn","batch",
//"-cmninit","41.00,-5.29,-0.12,5.09,2.48,-4.07,-1.37,-1.78,-5.08,-2.05,-6.45,-1.42,1.17",
"-ctl", "/home/dbarbera/Repositories/c_xyx/test_data/allowed1_philip_allowed/ctl/ctl_allowed1_philip_fixed_trimmed.txt",
"-dict", "/home/dbarbera/Repositories/c_xyx/test_data/dictionaries/art_db.phone",
//"-dict", "/home/dbarbera/Repositories/c_xyx/test_data/dictionaries/art_db_v3_inference.phone",
"-dither", "yes",
"-doublebw", "yes",
//"-feat","1s_c_d_dd",
"-frate", "72",
"-fwdflat", "no",
"-hmm", "/usr/local/share/xyzpocketsphinx/model/art-en-us/en-us",
//"-hmm", "/home/dbarbera/Repositories/pronounce-experimental/Models/Bare/2022-04-13T15:31:50-085_Bare.ci_cont",
"-logfn", "/home/dbarbera/Repositories/c_xyx/test_data/allowed1_philip_allowed/log/7eb4a3dd-1fb5-4e64-b7f9-fb3cbbca5829_-frate_72_BATCH_xyz_plus.log",
"-lpbeam", "1e-10000",
"-lponlybeam", "1e-10000",
"-lw", "6",
"-maxhmmpf", "-1",
"-maxwpf", "-1",
//"-nfft", "512",
"-pbeam", "1e-10000",
"-pip", "1.15",
"-pl_window", "0",
"-remove_noise", "yes",
"-remove_silence", "yes",
//"-svspec", "0-12/13-25/26-38",
"-topn", "6",
"-vad_postspeech", "20",
"-vad_prespeech", "5",
"-vad_startspeech", "5",
"-vad_threshold", "1.5",
"-wbeam", "1e-10000",
"-wip", "0.25"//,
//"-wlen", "0.032"
};
const int batch_params72_size=63;

const char *batch_params80[] = {
"pocketsphinx_batch",
"-adcin", "yes",
"-alpha", "0.97",
"-beam", "1e-10000",
"-bestpath", "no",
"-cepdir", "/home/dbarbera/Repositories/c_xyx/test_data/allowed1_philip_allowed/audio",
"-cepext", ".wav",
"-ctl", "/home/dbarbera/Repositories/c_xyx/test_data/allowed1_philip_allowed/ctl/ctl_allowed1_philip_fixed_trimmed.txt",
"-dict", "/home/dbarbera/Repositories/c_xyx/test_data/dictionaries/art_db.phone",
"-dither", "yes",
"-doublebw", "yes",
"-frate", "80",
"-fwdflat", "no",
"-hmm", "/usr/local/share/xyzpocketsphinx/model/art-en-us/en-us",
"-logfn", "/home/dbarbera/Repositories/c_xyx/test_data/allowed1_philip_allowed/log/7eb4a3dd-1fb5-4e64-b7f9-fb3cbbca5829_-frate_80_BATCH_xyz_plus.log",
"-lpbeam", "1e-10000",
"-lponlybeam", "1e-10000",
"-lw", "6",
"-maxhmmpf", "-1",
"-maxwpf", "-1",
//"-nfft", "512",
"-pbeam", "1e-10000",
"-pip", "1.15",
"-pl_window", "0",
"-remove_noise", "yes",
"-remove_silence", "yes",
"-topn", "6",
"-vad_postspeech", "20",
"-vad_prespeech", "5",
"-vad_startspeech", "5",
"-vad_threshold", "1.5",
"-wbeam", "1e-10000",
"-wip", "0.25"//,
//"-wlen", "0.028"
};
const int batch_params80_size=63;

const char *batch_params91[] = {
"pocketsphinx_batch",
"-adcin", "yes",
"-alpha", "0.97",
"-beam", "1e-10000",
"-bestpath", "no",
"-cepdir", "/home/dbarbera/Repositories/c_xyx/test_data/allowed1_philip_allowed/audio",
"-cepext", ".wav",
"-ctl", "/home/dbarbera/Repositories/c_xyx/test_data/allowed1_philip_allowed/ctl/ctl_allowed1_philip_fixed_trimmed.txt",
"-dict", "/home/dbarbera/Repositories/c_xyx/test_data/dictionaries/art_db.phone",
"-dither", "no",
"-doublebw", "no",
"-frate", "91",
"-fwdflat", "no",
"-hmm", "/usr/local/share/xyzpocketsphinx/model/art-en-us/en-us",
"-logfn", "/home/dbarbera/Repositories/c_xyx/test_data/allowed1_philip_allowed/log/7eb4a3dd-1fb5-4e64-b7f9-fb3cbbca5829_-frate_91_BATCH_xyz_plus.log",
"-lpbeam", "1e-10000",
"-lponlybeam", "1e-10000",
"-lw", "6",
"-maxhmmpf", "-1",
"-maxwpf", "-1",
//"-nfft", "512",
"-pbeam", "1e-10000",
"-pip", "1",
"-pl_window", "0",
"-remove_noise", "yes",
"-remove_silence", "no",
"-topn", "4",
"-vad_postspeech", "20",
"-vad_prespeech", "5",
"-vad_startspeech", "5",
"-vad_threshold", "0.5",
"-wbeam", "1e-10000",
"-wip", "0.5"//,
//"-wlen", "0.024"
};
const int batch_params91_size=63;

const char *batch_params105[] = {
"pocketsphinx_batch",
"-adcin", "yes",
"-alpha", "0.97",
"-beam", "1e-10000",
"-bestpath", "no",
"-cepdir", "/home/dbarbera/Repositories/c_xyx/test_data/allowed1_philip_allowed/audio",
"-cepext", ".wav",
"-ctl", "/home/dbarbera/Repositories/c_xyx/test_data/allowed1_philip_allowed/ctl/ctl_allowed1_philip_fixed_trimmed.txt",
"-dict", "/home/dbarbera/Repositories/c_xyx/test_data/dictionaries/art_db.phone",
"-dither", "no",
"-doublebw", "no",
"-frate", "105",
"-fwdflat", "no",
"-hmm", "/usr/local/share/xyzpocketsphinx/model/art-en-us/en-us",
"-logfn", "/home/dbarbera/Repositories/c_xyx/test_data/allowed1_philip_allowed/log/7eb4a3dd-1fb5-4e64-b7f9-fb3cbbca5829_-frate_105_BATCH_xyz_plus.log",
"-lpbeam", "1e-10000",
"-lponlybeam", "1e-10000",
"-lw", "6",
"-maxhmmpf", "-1",
"-maxwpf", "-1",
//"-nfft", "512",
"-pbeam", "1e-10000",
"-pip", "1",
"-pl_window", "0",
"-remove_noise", "yes",
"-remove_silence", "no",
"-topn", "4",
"-vad_postspeech", "20",
"-vad_prespeech", "5",
"-vad_startspeech", "5",
"-vad_threshold", "1.5",
"-wbeam", "1e-10000",
"-wip", "0.5"//,
//"-wlen", "0.020"
};
const int batch_params105_size=63;

const char *batch_params125[] = {
"pocketsphinx_batch",
"-adcin", "yes",
"-alpha", "0.97",
"-beam", "1e-10000",
"-bestpath", "no",
"-cepdir", "/home/dbarbera/Repositories/c_xyx/test_data/allowed1_philip_allowed/audio",
"-cepext", ".wav",
"-ctl", "/home/dbarbera/Repositories/c_xyx/test_data/allowed1_philip_allowed/ctl/ctl_allowed1_philip_fixed_trimmed.txt",
"-dict", "/home/dbarbera/Repositories/c_xyx/test_data/dictionaries/art_db.phone",
"-dither", "no",
"-doublebw", "no",
"-frate", "125",
"-fwdflat", "no",
"-hmm", "/usr/local/share/xyzpocketsphinx/model/art-en-us/en-us",
"-logfn", "/home/dbarbera/Repositories/c_xyx/test_data/allowed1_philip_allowed/log/7eb4a3dd-1fb5-4e64-b7f9-fb3cbbca5829_-frate_125_BATCH_xyz_plus.log",
"-lpbeam", "1e-10000",
"-lponlybeam", "1e-10000",
"-lw", "6",
"-maxhmmpf", "-1",
"-maxwpf", "-1",
//"-nfft", "256",
"-pbeam", "1e-10000",
"-pip", "1",
"-pl_window", "0",
"-remove_noise", "yes",
"-remove_silence", "no",
"-topn", "4",
"-vad_postspeech", "20",
"-vad_prespeech", "5",
"-vad_startspeech", "8",
"-vad_threshold", "1",
"-wbeam", "1e-10000",
"-wip", "0.5"//,
//"-wlen", "0.016"
};
const int batch_params125_size=63;


void* create_buffer(int* bsize, const char* filename, const char* mode){
    FILE* file = NULL;
    file = fopen(filename, mode);
    if (file == NULL) {
        E_ERROR_SYSTEM("Failed to open %s for parsing", filename);
        return 0;
    }
    fseek(file, 0, SEEK_END);
    long fsize = ftell(file);
    *bsize=fsize;
    fseek(file, 0, SEEK_SET);  /* same as rewind(f); */

    
    void* contents = (void*)malloc(fsize + 1);
    fread(contents, 1, fsize, file);
    fclose(file);

    return contents;
}

void delete_buffer(void* buffer){
    free(buffer);
}

int number_parameters(char *params[]) {
//This function is unreliable: depends on the memory sometimes there won't be a NULL!!!!
    int i=0;
    int count = 0;
    while(params[i]!=NULL) {
        count++;
        i++;
    }

    return count;
}

char* get_value(char *params[], const char *key) {
    char *value=NULL;
    int i=0;
    int count = 0;
    while((params[i]!=NULL) && (strcmp(params[i], key) != 0)) {
        count++;
        i++;
    }
    if(params[i+1] != NULL){
        value = (char*)malloc(strlen(params[i+1])*sizeof(char)+1);
        //memcpy(value, params[i], strlen(params[i+1])*sizeof(char)+1);
        strcpy(value, params[i+1]);
      //  printf(" key: %s value: %s\n", key, value);
        return value;
    } else {
        return NULL;
    }
}

char* get_audiofile(const char *ctl_file, const char *audio_dir, const char *extension ) {
    FILE* file = NULL;
    char * line = NULL;
    size_t len = 0;
    ssize_t read;

    file = fopen(ctl_file, "r");
    if (file == NULL) {
        E_ERROR_SYSTEM("Failed to open %s for reading", ctl_file);
        return NULL;
    }

    if ((read = getline(&line, &len, file))==-1){
        E_ERROR_SYSTEM("Ctl file (%s) is empty.", ctl_file);
        return NULL;
    };

    char *audiofile = (char*)malloc(sizeof(char)*(len+strlen(audio_dir)+strlen(extension)+1+1));
    audiofile[0] = '\0';
    strcat(audiofile, audio_dir);
    strcat(audiofile,"/");
    strcat(audiofile, line);
    strcat(audiofile, extension);

    free(line);
    fclose(file);

    return audiofile;

}


struct BinaryData {
    void *buffer;
    int size;
};

struct CharData {
    char **p;
    int size;
};

struct ResultData {
    char result[512];
    int size;
};

struct PS_Data {
    struct BinaryData jsgf;
    struct BinaryData wav;
    struct CharData params;
    struct ResultData result;
};

struct PS_Batch_Data {
    struct BinaryData wav;
    struct CharData params;
    struct ResultData result;
};

void load_data(const char *p1[], const int p1_size, 
               const char *p2[], const int p2_size, 
               const char *p3[], const int p3_size,
               const char *p4[], const int p4_size,
               const char *p5[], const int p5_size,
               struct PS_Data data[5]) {
//void load_data(char *p1[], char *p2[], char *p3[], char *p4[], char *p5[], struct Data data[
    data[0].params.p=(char**)p1;
    data[0].params.size=p1_size;
    
    data[1].params.p=(char**)p2;
    data[1].params.size=p2_size;

    data[2].params.p=(char**)p3;
    data[2].params.size=p3_size; //number_parameters((char**)p3);

    data[3].params.p=(char**)p4;
    data[3].params.size=p4_size; //number_parameters((char**)p4);

    data[4].params.p=(char**)p5;
    data[4].params.size=p5_size; //number_parameters((char**)p5);

    for( int i =0; i< 5; i++){
        data[i].jsgf.buffer = create_buffer( &data[i].jsgf.size, get_value(data[i].params.p, "-jsgf"), "rb");
        data[i].wav.buffer  = create_buffer( &data[i].wav.size,  get_value(data[i].params.p, "-infile"), "rb");
        memset(data[i].result.result, 'a', sizeof(char)*512);
        data[i].result.size = 512;
    }
}


void load_data_batch(const char *p1[], const int p1_size, 
               const char *p2[], const int p2_size, 
               const char *p3[], const int p3_size,
               const char *p4[], const int p4_size,
               const char *p5[], const int p5_size,
               struct PS_Batch_Data data[5]) {
//void load_data(char *p1[], char *p2[], char *p3[], char *p4[], char *p5[], struct Data data[5]) {

    data[0].params.p=(char**)p1;
    data[0].params.size=p1_size; //number_parameters((char**)p1); 

    data[1].params.p=(char**)p2;
    data[1].params.size=p2_size;
    
    data[2].params.p=(char**)p3;
    data[2].params.size=p3_size; //number_parameters((char**)p3);

    data[3].params.p=(char**)p4;
    data[3].params.size=p4_size; //number_parameters((char**)p4);

    data[4].params.p=(char**)p5;
    data[4].params.size=p5_size; //number_parameters((char**)p5);

    for( int i = 0; i < 5; i++){
        //data[i].jsgf.buffer = create_buffer( &data[i].jsgf.size, get_value(data[i].params.p, "-jsgf"), "rb");
        char *ctl_file = get_value(data[i].params.p, "-ctl");
        char *cep_dir = get_value(data[i].params.p, "-cepdir");
        char *extension = get_value(data[i].params.p, "-cepext");
        char *ctl_audio_file = get_audiofile(ctl_file, cep_dir, extension);
        data[i].wav.buffer  = create_buffer( &data[i].wav.size,  ctl_audio_file, "rb");
        memset(data[i].result.result, 'a', sizeof(char)*512);
        data[i].result.size = 512;
    }

    //printf("working on it");

}

