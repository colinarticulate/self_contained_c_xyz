#include "a.h"
#include <b.h>

FFI_PLUGIN_EXPORT intptr_t sum(intptr_t a, intptr_t b) {return a+b;}

FFI_PLUGIN_EXPORT intptr_t sum_long_running(intptr_t a, intptr_t b) {
  // Simulate work.
#if _WIN32
  Sleep(5000);
#else
  usleep(5000 * 1000);
#endif
  return a+b;
}

FFI_PLUGIN_EXPORT intptr_t sum_long_running_from_b(intptr_t a, intptr_t b) {
  // Simulate work.
#if _WIN32
  Sleep(5000);
#else
  usleep(2*100 * 1000);
#endif
  return b_sum(a, b);
}