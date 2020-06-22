# Lock Server

Goal of the lock server is to maintain a log of the locks acquired at any moment of time to ensure consistency for any service using the lock server.

## Spec

The lock server implements two main functions:
1. Acquire
2. Release

### Acquire

* The lock server grants the lock to one client at a time provided a lock is currently not acquired on the said acquiring object.
* Returns a lockToken, a hash based on the object and the client.

### Release

* The lock server release the lock on the object demanded. A lock token acts as an auth to ensure there are no malicious clients.
* A point to note here is that we should not let lock "release" for objects where the lock isnt acquired on. This might be a sign of malicious clients trying attacks to get access to objects in the DFS.
## Implementation

### The lock service
To make this lock server idea scalable and add multiple implementations of the lock service, we design the following:

```
type Data interface {
    ID() string
    ...
}

type LockService interface {
    Insert(Data) error
    Delete(Data) error
}
```

The data to be inserted is named `Data` for lack of a better name.

### Data structure for maintaining locks

```
type SafeLockMap struct {
    LockMap map[string]{}interface 
    Mutex   sync.Mutex 
}
```

### Acquire

```
func acquire(obj) Lock,err {
    // check for existing lock on obj, return if doesnt exist
    if !checkLock(obj) {
        return lock
    } 
}
```

### Release 

```
func release(obj, lockToken) err {
    // check for existing lock on object, return only if exists
    if checkLock(obj){
        return nil
    }
}
```


## Further

* The locks acquired by the client must be an authentication for the file system to verify that this is a legitimate client that has acquired a lock and no other application is already using it and I can safely provide access to this client.
* The files or "objects" must be a standard way of storage in the distributed file system.
* The lockserver is only accessed by the DFS. No client has direct access to the lockserver. 
* DFS's abstraction to this API must be something like this:

```
// on incoming client access to file `F`
lock,err := acquireLock(F)
if err != nil {
    // grant access to file
}else {
    // report / wait
}
```
