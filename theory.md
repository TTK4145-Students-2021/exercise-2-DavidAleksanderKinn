Exercise 2 - Theory questions
-----------------------------

### What is an atomic operation?
> Atomic operations is when a opperation process is completly independent from the all other processes. For example when executing a RMW (Read, modify, Whrite) opperation, one wants to finish this process before any other processes writes on the same value. This is what went wrong in the code from the last Assignment.  

### What is a critical section?
> In a concurrent program it is common that diffrient processes access the same resources. This can cause problems if the accessablilty i'snt done with atomic operations. The segment of the code that accesses this shared resources are called a critical section.

### What is the difference between race conditions and data races?
> A race condition is when the output depends on the exetution ordering of the treads. A data race is when two or more threads access the same varable at the same time and one or more of the threads are writing to the varable. 

### What are the differences between semaphores, binary semaphores, and mutexes?
> In mutexes it is arranged a form of ownership where whatever process or thread decides when to wait, also decides then it is ok to post (in some sence one can think of the process having this ownership as a key-holder). A mutex is a class. In binary semaphores all processes and threads can decised when to wait and when to post (more simular to a signal than a lock). If the critical section can be shared by N numbers of processes at the same time, a semaphores can distrebute the resourses of each prosess. Unalike binary semaphores and mutexes, (counting) semaphores can have values outside a boolean range. Semaphores are integers. 

### What are the differences between channels (in Communicating Sequential Processes, or as used by Go, Rust), mailboxes (in the Actor model, or as used by Erlang, D, Akka), and queues (as used by Python)? 
> In channels threads can wait to send values before the resiver is ready. In a que the "select" functionality a channel have i'snt used. In mailboxes is dedicated for one resiver in contrast to the others. 

### List some advantages of using message passing over lock-based synchronization primitives.
> This is a way to easier to access resourse proirsation everywere in the code. This is great when the threads don't share the same variables, but can be risky if they do. 

### List some advantages of using lock-based synchronization primitives over message passing.
> The lock-based synchronization primitives is a safer way to handle when two or more prosesses shares the same variable. To program this kinds of primitives is often easier than programing message passing. 
