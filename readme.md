# glisp-shell-demo

This is basic example code to demonstrate how you can create an interactive lispy shell by embedding [glisp](https://github.com/zhemao/glisp) within your golang program.

## Usage

    $ make
    $ ./glisp-shell-demo 
    glisp> (+ 1 1)
    > 2
    glisp> (defn add [a] (+ a 1))
    > ()
    glisp> (add 5)
    > 6
    glisp> 

## License

MIT
