#include "a.h"
#include <b.h>

intptr_t sum(intptr_t a, intptr_t b) {return a+b;}

intptr_t sum_long_running(intptr_t a, intptr_t b) {
  // Simulate work.
#if _WIN32
  Sleep(5000);
#else
  usleep(2*1000 * 1000);
#endif
  return sum(a, b);
}

intptr_t sum_long_running_from_b(intptr_t a, intptr_t b) {
  // Simulate work.
#if _WIN32
  Sleep(5000);
#else
  usleep(2*1000 * 1000);
#endif
  return b_sum(a, b);
}

int main(void){
  intptr_t a = 2;
  intptr_t b = 3;
  intptr_t sum = sum_long_running(a,b);
  printf("%ld\n",sum);
  intptr_t bsum = sum_long_running_from_b(a+1,b+1);
  printf("%ld\n",bsum);

  return 0;
}