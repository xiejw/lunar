## Reading Notes

### LogDir

See file _glog_file.go_.

LogDir is specified via `--log_dir`. If absent, the `os.TempDir()` is used, so
`$TMPDIR` is used on UNIX.

Log files are lazily created. The overheads are low.
