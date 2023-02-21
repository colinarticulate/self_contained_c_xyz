#include <chrono>
#include <thread>
#include <pthread.h> 

#include "ps_plus.h"
#include "batch_plus.h"
#include "data.h"

using namespace std::chrono;


void sequential_encapsulated(PS_Data data[5]) {
    XYZ_PocketSphinx ps[5];

	high_resolution_clock::time_point start;
	high_resolution_clock::time_point end;
    high_resolution_clock::time_point starts[5];
	high_resolution_clock::time_point ends[5];

    start=high_resolution_clock::now();
    for(int i = 0; i< 5; i++) {
        starts[i]=high_resolution_clock::now();
        ps[i].init((void*)(data[i].jsgf.buffer), (size_t)data[i].jsgf.size, (void*)data[i].wav.buffer, (size_t)data[i].wav.size, (int)data[i].params.size, data[i].params.p);
        ps[i].init_recognition();
        ps[i].recognize_from_buffered_file();
        ps[i].terminate();
        ends[i]=high_resolution_clock::now();

        strcpy(data[i].result.result,ps[i]._result);
        data[i].result.size = ps[i]._result_size;
    }
   end=high_resolution_clock::now();

   double total=0;
    for(int i = 0; i< 5; i++) {
        int frate = atoi(get_value(data[i].params.p,"-frate"));
        printf("%d\t%s\t%lfms\n",frate,data[i].result.result,duration<double, std::milli>(ends[i] - starts[i]).count());
        total = total + duration<double, std::milli>(ends[i] - starts[i]).count();
    }

	auto dur_us = duration<double, std::micro>(end - start).count();
	auto dur_ms = duration<double, std::milli>(end - start).count();
	printf("Time: %lfus %lfms\n", dur_us, dur_ms);
    printf("Computation time on task: %lfms\n", total);
}

void process(XYZ_PocketSphinx *ps, PS_Data *data) {
        ps->init((void*)(data->jsgf.buffer), (size_t)data->jsgf.size, (void*)data->wav.buffer, (size_t)data->wav.size, (int)data->params.size, data->params.p);
        ps->init_recognition();
        ps->recognize_from_buffered_file();
        ps->terminate(); 
}

void process_batch(XYZ_Batch *ps, PS_Batch_Data *data) {
        ps->init((void*)data->wav.buffer, (size_t)data->wav.size, (int)data->params.size, data->params.p);
        ps->init_recognition();
        ps->process();
        ps->terminate(); 
}

void init(XYZ_PocketSphinx *ps, PS_Data *data) {
        ps->init((void*)(data->jsgf.buffer), (size_t)data->jsgf.size, (void*)data->wav.buffer, (size_t)data->wav.size, (int)data->params.size, data->params.p);
        ps->init_recognition();
}

void process_no_init(XYZ_PocketSphinx *ps) {
       // ps->init((void*)(data->jsgf.buffer), (size_t)data->jsgf.size, (void*)data->wav.buffer, (size_t)data->wav.size, (int)data->params.size, data->params.p);
        ps->init_recognition();
        ps->recognize_from_buffered_file();
        ps->terminate(); 
}

void parallel_encapsulated(PS_Data data[5]) {
    XYZ_PocketSphinx ps[5];
    //std::thread threads[5];

	high_resolution_clock::time_point start;
	high_resolution_clock::time_point end;
    // high_resolution_clock::time_point starts[5];
	// high_resolution_clock::time_point ends[5];

    start=high_resolution_clock::now();
    // for(int i = 0; i< 5; i++) {
    //     init(&ps[i], &data[i]);

    // }
    // printf("debug:");
    // for(int i = 0; i< 5; i++) {
        
    //     threads[i](process, &ps[i], &data[i]);
    //     //process(&ps[i],&data[i]);
    //     //process_no_init(&ps[i]);
        
    // }
    // for(int i = 0; i< 5; i++) {
    //     threads[i].join();
    // }

    std::thread t0(process, &ps[0],&data[0]);
    std::thread t1(process, &ps[1],&data[1]);
    std::thread t2(process, &ps[2],&data[2]);
    std::thread t3(process, &ps[3],&data[3]);
    std::thread t4(process, &ps[4],&data[4]);
    
    // int i=0;
    // ps[i].init((void*)(data[i].jsgf.buffer), (size_t)data[i].jsgf.size, (void*)data[i].wav.buffer, (size_t)data[i].wav.size, (int)data[i].params.size, data[i].params.p);
    // ps[i].init_recognition();
    // ps[i].recognize_from_buffered_file();
    // ps[i].terminate();
    // std::thread t0(process_no_init, &ps[i]);
    
    t0.join();
    t1.join();
    t2.join();
    t3.join();
    t4.join();
        


   end=high_resolution_clock::now();

    for(int i = 0; i< 5; i++) {
        strcpy(data[i].result.result,ps[i]._result);
        data[i].result.size = ps[i]._result_size;
    }

    //double total=0;
    for(int i = 0; i< 5; i++) {
        int frate = atoi(get_value(data[i].params.p,"-frate"));
        printf("%d\t%s\n",frate,data[i].result.result);
        // printf("%d\t%s\t\t%lfms\n",frate,data[i].result.result,duration<double, std::milli>(ends[i] - starts[i]).count());
        // total = total + duration<double, std::milli>(ends[i] - starts[i]).count();
    }

	auto dur_us = duration<double, std::micro>(end - start).count();
	auto dur_ms = duration<double, std::milli>(end - start).count();
	printf("Time: %lfus %lfms\n", dur_us, dur_ms);
    //printf("Computation time on task: %lfms\n", total);
}

void parallel_encapsulated_batch(PS_Batch_Data data[5]) {
    XYZ_Batch ps[5];
    //std::thread threads[5];

	high_resolution_clock::time_point start;
	high_resolution_clock::time_point end;
    // high_resolution_clock::time_point starts[5];
	// high_resolution_clock::time_point ends[5];

    start=high_resolution_clock::now();
    // for(int i = 0; i< 5; i++) {
    //     init(&ps[i], &data[i]);

    // }
    // printf("debug:");
    // for(int i = 0; i< 5; i++) {
        
    //     threads[i](process, &ps[i], &data[i]);
    //     //process(&ps[i],&data[i]);
    //     //process_no_init(&ps[i]);
        
    // }
    // for(int i = 0; i< 5; i++) {
    //     threads[i].join();
    // }

    std::thread t0(process_batch, &ps[0],&data[0]);
    std::thread t1(process_batch, &ps[1],&data[1]);
    std::thread t2(process_batch, &ps[2],&data[2]);
    std::thread t3(process_batch, &ps[3],&data[3]);
    std::thread t4(process_batch, &ps[4],&data[4]);
    
    // int i=0;
    // ps[i].init((void*)(data[i].jsgf.buffer), (size_t)data[i].jsgf.size, (void*)data[i].wav.buffer, (size_t)data[i].wav.size, (int)data[i].params.size, data[i].params.p);
    // ps[i].init_recognition();
    // ps[i].recognize_from_buffered_file();
    // ps[i].terminate();
    // std::thread t0(process_no_init, &ps[i]);
    
    t0.join();
    t1.join();
    t2.join();
    t3.join();
    t4.join();
        


   end=high_resolution_clock::now();

    for(int i = 0; i< 5; i++) {
        strcpy(data[i].result.result,ps[i]._result);
        data[i].result.size = ps[i]._result_size;
    }

    //double total=0;
    for(int i = 0; i< 5; i++) {
        int frate = atoi(get_value(data[i].params.p,"-frate"));
        printf("%d\t%s\n",frate,data[i].result.result);
        // printf("%d\t%s\t\t%lfms\n",frate,data[i].result.result,duration<double, std::milli>(ends[i] - starts[i]).count());
        // total = total + duration<double, std::milli>(ends[i] - starts[i]).count();
    }

	auto dur_us = duration<double, std::micro>(end - start).count();
	auto dur_ms = duration<double, std::milli>(end - start).count();
	printf("Time: %lfus %lfms\n", dur_us, dur_ms);
    //printf("Computation time on task: %lfms\n", total);
}


void sequential_encapsulated_batch(PS_Batch_Data data[5]) {
    XYZ_Batch ps[5];

	high_resolution_clock::time_point start;
	high_resolution_clock::time_point end;
    high_resolution_clock::time_point starts[5];
	high_resolution_clock::time_point ends[5];

    start=high_resolution_clock::now();
    for(int i = 0; i< 5; i++) {
        starts[i]=high_resolution_clock::now();  
        ps[i].init((void*)data[i].wav.buffer, (size_t)data[i].wav.size, (int)data[i].params.size, data[i].params.p);
        //ends[i]=high_resolution_clock::now(); 
        ps[i].init_recognition();
        
        ps[i].process();
        
        ps[i].terminate();
        
        ends[i]=high_resolution_clock::now();

        strcpy(data[i].result.result,ps[i]._result);
        data[i].result.size = ps[i]._result_size;
    }
   end=high_resolution_clock::now();

   double total=0;
    for(int i = 0; i< 5; i++) {
        int frate = atoi(get_value(data[i].params.p,"-frate"));
 
        printf("%d\t%s\t%lfms\n",frate,data[i].result.result,duration<double, std::milli>(ends[i] - starts[i]).count());
        total = total + duration<double, std::milli>(ends[i] - starts[i]).count();
    }

	auto dur_us = duration<double, std::micro>(end - start).count();
	auto dur_ms = duration<double, std::milli>(end - start).count();
	printf("Time: %lfus %lfms\n", dur_us, dur_ms);
    printf("Computation time on task: %lfms\n", total);
}

//Only using pthread library:
struct arg_struct {
    XYZ_PocketSphinx *arg1;
    PS_Data *arg2;
};

//void process(XYZ_PocketSphinx *ps, PS_Data *data) {
void *p_process(void *arguments) {
    struct arg_struct *args = (struct arg_struct *)arguments;

    args->arg1->init((void*)(args->arg2->jsgf.buffer), (size_t)args->arg2->jsgf.size, (void*)args->arg2->wav.buffer, (size_t)args->arg2->wav.size, (int)args->arg2->params.size, args->arg2->params.p);
    args->arg1->init_recognition();
    args->arg1->recognize_from_buffered_file();
    args->arg1->terminate(); 

    pthread_exit(NULL);
}

void parallel_encapsualted_with_pthreads(PS_Data data[5]){
    XYZ_PocketSphinx ps[5];

    pthread_t thread_ids[5];
    int max_calls=5;

    //init args:
    struct arg_struct args[5];
    for(int i = 0; i<max_calls; i++) {
      args[i].arg1 = &ps[i];
      args[i].arg2 = &data[i];
    };

	high_resolution_clock::time_point start;
	high_resolution_clock::time_point end;

    start=high_resolution_clock::now();
    for(int i = 0; i<max_calls; i++) {

        pthread_create(&thread_ids[i], NULL, &p_process, &args[i]); 
         
    };

    for(int i = 0; i<max_calls; i++) {
        pthread_join(thread_ids[i], NULL);
    }
    end=high_resolution_clock::now();

    for(int i = 0; i< 5; i++) {
        strcpy(data[i].result.result,ps[i]._result);
        data[i].result.size = ps[i]._result_size;
    }

    //double total=0;
    for(int i = 0; i< 5; i++) {
        int frate = atoi(get_value(data[i].params.p,"-frate"));
        printf("%d\t%s\n",frate,data[i].result.result);
        // printf("%d\t%s\t\t%lfms\n",frate,data[i].result.result,duration<double, std::milli>(ends[i] - starts[i]).count());
        // total = total + duration<double, std::milli>(ends[i] - starts[i]).count();
    }

	auto dur_us = duration<double, std::micro>(end - start).count();
	auto dur_ms = duration<double, std::milli>(end - start).count();
	printf("Time: %lfus %lfms\n", dur_us, dur_ms);
    //printf("Computation time on task: %lfms\n", total);

}

//Only using pthread library:
struct arg_batch_struct {
    XYZ_Batch *arg1;
    PS_Batch_Data *arg2;
};

void *p_process_batch(void *arguments) {
    struct arg_batch_struct *args = (struct arg_batch_struct *)arguments;

    args->arg1->init((void*)args->arg2->wav.buffer, (size_t)args->arg2->wav.size, (int)args->arg2->params.size, args->arg2->params.p);
    args->arg1->init_recognition();
    args->arg1->process();
    args->arg1->terminate(); 

    pthread_exit(NULL);
}

void parallel_encapsualted_batch_with_pthreads(PS_Batch_Data data[5]){
    XYZ_Batch ps[5];

    pthread_t thread_ids[5];
    int max_calls=5;

    //init args:
    struct arg_batch_struct args[5];
    for(int i = 0; i<max_calls; i++) {
      args[i].arg1 = &ps[i];
      args[i].arg2 = &data[i];
    };

	high_resolution_clock::time_point start;
	high_resolution_clock::time_point end;

    start=high_resolution_clock::now();
    for(int i = 0; i<max_calls; i++) {

        pthread_create(&thread_ids[i], NULL, &p_process_batch, &args[i]); 
         
    };

    for(int i = 0; i<max_calls; i++) {
        pthread_join(thread_ids[i], NULL);
    }
    end=high_resolution_clock::now();

    for(int i = 0; i< 5; i++) {
        strcpy(data[i].result.result,ps[i]._result);
        data[i].result.size = ps[i]._result_size;
    }

    //double total=0;
    for(int i = 0; i< 5; i++) {
        int frate = atoi(get_value(data[i].params.p,"-frate"));
        printf("%d\t%s\n",frate,data[i].result.result);
        // printf("%d\t%s\t\t%lfms\n",frate,data[i].result.result,duration<double, std::milli>(ends[i] - starts[i]).count());
        // total = total + duration<double, std::milli>(ends[i] - starts[i]).count();
    }

	auto dur_us = duration<double, std::micro>(end - start).count();
	auto dur_ms = duration<double, std::milli>(end - start).count();
	printf("Time: %lfus %lfms\n", dur_us, dur_ms);
    //printf("Computation time on task: %lfms\n", total);

}


int
main()
{
    struct PS_Data data[5];
    load_data(params125, params125_size,
              params72, params72_size, 
              params80, params80_size,
              params91, params91_size,
              params105, params105_size, 
              data);
    //load_data(params125, params125, params125, params125, params125, data);


    // DATA data(5);
    // data.load(params125);

    //sequential_encapsulated(data);
    parallel_encapsulated(data);
    //parallel_encapsualted_with_pthreads(data);

    // int i=4;
    // XYZ_PocketSphinx ps1;
    // ps1.init((void*)(data[i].jsgf.buffer), (size_t)data[i].jsgf.size, (void*)data[i].wav.buffer, (size_t)data[i].wav.size, (int)data[i].params.size, data[i].params.p);
    // ps1.init_recognition();
    // ps1.recognize_from_buffered_file();
    // printf("%s\n",ps1._result);
    // ps1.terminate();

    struct PS_Batch_Data batch_data[5];
    load_data_batch(batch_params72, batch_params72_size,
                    batch_params80, batch_params80_size,
                    batch_params91, batch_params91_size,
                    batch_params105, batch_params105_size,
                    batch_params125, batch_params125_size,
                    batch_data);
    
    //sequential_encapsulated_batch(batch_data);
    parallel_encapsulated_batch(batch_data);
    //parallel_encapsualted_batch_with_pthreads(batch_data);
    
    printf("Working on it...\n");

    return 0;
    

    
}


