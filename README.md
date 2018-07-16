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

## Install
* `cd $GOPATH`
* `go get github.com/dvwallin/pluff`
* `go install github.com/dvwallin/pluff`

## License
MIT License

Copyright (c) 2018 David Satime Wallin

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
