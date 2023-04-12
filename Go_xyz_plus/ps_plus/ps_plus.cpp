//Based on xyz_plus.cpp
//https://github.com/DavidBarbera/xyz_plus/blob/v2.0.0/xyz_plus.cpp

#ifdef WIN32
#define EXPORT __declspec(dllexport)
#else
#define EXPORT extern "C" __attribute__((visibility("default"))) __attribute__((used))
#endif

#include <cstddef>
#include <cstdio>
#include <cstdlib>
#include <exception>
#include <stdexcept>

#include "ps_continuous.h"
#include "ps_batch.h"


int _ps_continuous_call(char* jsgf_buffer, int jsgf_buffer_size, char* audio_buffer, int audio_buffer_size, int argc, char *argv[], char* result, int rsize){

    XYZ_PocketSphinx ps;
    ps.init(jsgf_buffer, jsgf_buffer_size, audio_buffer, audio_buffer_size, argc, argv);
    ps.init_recognition();
    ps.recognize_from_buffered_file();
    ps.terminate();

    if (ps._result_size < rsize && strlen(ps._result)>0){

        for(int i=0;i<ps._result_size; i++){
            result[i]=(char)ps._result[i];
        }
    } 

    return ps._result_size;
 } 

EXPORT
char* ps_continuous_call(char* jsgf_buffer, int jsgf_buffer_size, char* audio_buffer, int audio_buffer_size, int argc, char *argv[], char* result, int rsize){
    char* wresult=NULL;
    try {
        _ps_continuous_call( jsgf_buffer,  jsgf_buffer_size,  audio_buffer,  audio_buffer_size,  argc, argv, result, rsize);
    } catch(std::exception &e) {
        wresult = strdup(e.what());
    }
    return wresult;
}

int _ps_batch_call(void* audio_buffer, int audio_buffer_size, int argc, char *argv[], char* result, int rsize){

    XYZ_Batch ps;
    ps.init(audio_buffer, audio_buffer_size, argc, argv);
    ps.init_recognition();
    ps.process();
    ps.terminate();

    if (ps._result_size < rsize && strlen(ps._result)>0){

        for(int i=0;i<ps._result_size; i++){
            result[i]=(char)ps._result[i];
        }
    } 

    return ps._result_size;
 } 

EXPORT
char* ps_batch_call(void* audio_buffer, int audio_buffer_size, int argc, char *argv[], char* result, int rsize){
    char* wresult=NULL;
    try {
        _ps_batch_call( audio_buffer,  audio_buffer_size,  argc, argv, result, rsize);
    } catch(std::exception &e) {
        wresult = strdup(e.what());
    }
    return wresult;
}
