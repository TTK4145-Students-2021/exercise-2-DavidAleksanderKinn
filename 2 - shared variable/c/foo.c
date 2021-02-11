#include <pthread.h>
#include <stdio.h>

int i = 0;
static pthread_mutex_t mutex = PTHREAD_MUTEX_INITIALIZER;

// Note the return type: void*
void* incrementingThreadFunction(){
    for (int j = 0; j < 1000000; j++) {
	// TODO: sync access to i
    pthread_mutex_lock(&mutex);
	i++;
    pthread_mutex_unlock(&mutex);
    }
    return NULL;
}

void* decrementingThreadFunction(){
    for (int j = 0; j < 1000000; j++) {
	// TODO: sync access to i
    pthread_mutex_lock(&mutex);
	i--;
    pthread_mutex_unlock(&mutex);
    }
    return NULL;
}
 //Mutex is used over semaphores since the resource (variable "i") should only be acsessed by one process at the time. 
 //Therefore it would be a safer option to have a lower accessibility then in semaphores.

int main(){
    pthread_t incrementingThread, decrementingThread;
    
    pthread_create(&incrementingThread, NULL, incrementingThreadFunction, NULL);
    pthread_create(&decrementingThread, NULL, decrementingThreadFunction, NULL);
    
    pthread_join(incrementingThread, NULL);
    pthread_join(decrementingThread, NULL);
    
    printf("The magic number is: %d\n", i);
    return 0;
}
