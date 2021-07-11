## Reading Notes

### LogDir

See file _glog_file.go_.

LogDir is specified via `--log_dir`. If absent, the `os.TempDir()` is used, so
`$TMPDIR` is used on UNIX.

Log files are lazily created. The overheads are low.

### Logging Destination

```
Cond 1. logtostderr     = True   only stderr
Cond 2. logtostderr     = False
  Cond 2a. alsologtostderr = True   both stderr and log_dir
  Cond 2b. alsologtostderr = False
    Cond 2b_1 level >= stderrthreshold   both stderr and log_dir
    Cond 2b_2 level <  stderrthreshold   only log_dir
```

### Performance Optimization

- `loggingT.freeList` maintains a list of byte buffersfree. In addition, its
  lock is separate from the main mutex so buffers can be grabbed and printed to
  without holding the main lock, for better parallelization.
- `formatHeader` uses customized method to format header instead of printf

### Log Rotation

`syncBuffer.Write` rotates the log files when buffer is full.

### File name and Line Number

The file name and line number are obtained via getting callers' PC and
[CallersFrames](https://golang.org/pkg/runtime/#CallersFrames). See
[runtime.Callers](https://golang.org/pkg/runtime/#Callers)

### Go log and glog

` CopyStandardLogTo` arranges for messages written to the Go "log" package's
default logs to also appear in the Google logs for the named and lower
severities.`
