# Mutex and Channel basics

### What is an atomic operation?
> An operation (or set of operations) guranteed to run without interruptions.

### What is a semaphore?
> A variable used to synchronize concurrent execution.

### What is a mutex?
> A mutex is a mutually exclusive flag used to control access to a piece of code or resource. Use concept of ownership.

### What is the difference between a mutex and a binary semaphore?
> A mutex is used to provide exclusive access to a resource, but a binary semaphore is used for synchronizing execution of signaling the other threads. Semaphores requires cooperation, mutex does not. 

### What is a critical section?
> A group of instruction that need to execute atomically.

### What is the difference between race conditions and data races?
 > A race condition is a collective term for all timing related errors, where the instruction order affects the result. A datarace is when multiple instructions try to access the same memory location from different threads (without synchronization) and is therefore a race condition.

### List some advantages of using message passing over lock-based synchronization primitives.
> 1. Scalability, 2. avoid deadlocks, 3. Enforce synchronization (or no need for it) 4. Intutitive design, 5. Easier debugging.

### List some advantages of using lock-based synchronization primitives over message passing.
> 1. Lower memoryfootprint in embedded systems. 
