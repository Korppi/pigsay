# pigsay
Inspired by cowsay. I made this to learn more about go and cobra.

# How to use
In builds/ folder you can find binaries for pigsay. For windows use pigsay.exe and for linux use pigsay. Download them, open terminal in folder you downloaded them and execute command 

```console
[teppo@vmi1490863 ~]$ ./pigsay Hello World
```

It should look like this
```console
[teppo@vmi1490863 ~]$ ./pigsay Hello World
   _____________
  /             \
  | Hello World |
  \_  __________/
    \|
         _/|________
        / o         \
       E,            |S
        \___________/
         WW       WW

```
It is possible that on linux you might need to give executable right to pigsay before it works:
```console
[teppo@vmi1490863 ~]$ sudo chmod +x pigsay
```

There is help command provided
```console
[teppo@vmi1490863 ~]$ ./pigsay --help
Make a pig say things!

Usage:
  pigsay [-e] <MESSAGE> [flags]

Examples:
pigsay
        Display help
pigsay Hello World
        Display pig saying "Hello World"
pigsay -e @ Hello World
        Display pig with custom eye @ saying "Hello World"
pigsay --eyes @ Hello World
        Display pig with custom eye @ saying "Hello World"


Flags:
  -e, --eyes string   give pig different eyes (default "o")
  -h, --help          help for pigsay
  -v, --version       version for pigsay

```

# Example video

[![asciicast](https://asciinema.org/a/PoIgoc2GUqiAtg7tN8zfzwWFg.svg)](https://asciinema.org/a/PoIgoc2GUqiAtg7tN8zfzwWFg)
