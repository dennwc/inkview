#include "_cgo_export.h"

int main_handler(int t, int p1, int p2) {
    return goMainHandler(t, p1, p2); 
}

void c_keyboard_handler(char* text) {
    goKeyboardHandler(text);    
}