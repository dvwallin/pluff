# PLUFF
A simple app to compress files, specified in a file, into a specified target folder.

```
Usage of /tmp/go-build522348288/b001/exe/main:
  -targetFolder string
        the folder you want the archives to be stored in (default "$HOME/.pluff")
```

It automatically searches for any of the following:
```
~/.pluff.conf
~/pluff.conf
./.pluff.conf
./pluff.conf
/etc/pluff.conf
```

The config file `pluff.conf` is simply a text-file where you specify the files and folders you want in the archive. One folder/file per row.
For example:
```
/home/david/Images
/home/david/Documents
/home/david/my_todo.txt
```