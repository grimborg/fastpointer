fastpointer
===========

Set the speed and sensitivity of the Thinkpad pointer

    fastpointer --help

    Usage of fastpointer:
      -sensitivity=250: Pointer sensitivity
      -speed=120: Pointer speed

To install and run without sudo:

    go build fastpointer.go
    sudo chown root:root fastpointer
    sudo chmod +s fastpointer
    sudo mv fastpointer /usr/local/bin
