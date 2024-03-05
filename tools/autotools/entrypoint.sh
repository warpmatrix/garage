function dev_side() {
    aclocal # Set up an m4 environment
    autoconf # Generate configure from configure.ac
    automake --add-missing # Generate Makefile.in from Makefile.am
    ./configure # Generate Makefile from Makefile.in
    make distcheck # Use Makefile to build and test a tarball to distribute
}

function client_side() {
    ./configure # Generate Makefile from Makefile.in
    make # Use Makefile to build the program
    make install # Use Makefile to install the program
}
