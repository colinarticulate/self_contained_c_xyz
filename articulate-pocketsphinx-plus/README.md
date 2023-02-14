# Articulate-pocketsphinx
Repo to create variations of CMU's pocketsphinx.  
This is the xyz version; our first prototype.  


## System installation  
Should work the same way as with the original pocketsphinx installation.  
[PocketSphinx 5prealpha] (https://github.com/cmusphinx/pocketsphinx)  
  

```
$ cd articulate-pocketsphinx
$ cd xyzsphinxbase
$ ./autogen.sh
$ make  
$ sudo make install
$ cd ..
$ cd xyzpocketsphinx
$ ./autogen.sh
$ make  
$ sudo make install  
$ cd /usr/local/lib 
$ sudo ldconfig  

```



