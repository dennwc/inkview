#include "_cgo_export.h"

int main_handler(int t, int p1, int p2) {
    return goMainHandler(t, p1, p2); 
}

void c_keyboard_handler(char* text) {
    goKeyboardHandler(text);    
}

void c_rotate_handler(int direction){
    goRotateHandler(direction);
}

void c_dialog_handler(int button){
    goDialogHandler(button);
}

void c_timeedit_handler(long newtime){
    goTimeEditHandler(newtime);
}