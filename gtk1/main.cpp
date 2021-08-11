#include <gtk/gtk.h>
#include "./gtk.hpp"
void printint(int v); 

int main(int argc, const char **argv) 
{
    printint(argc);
    testGtk ();
    return 1;
}