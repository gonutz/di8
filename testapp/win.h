#include <Windows.h>

#define OK                     0
#define Error_RegisterClassEx -1
#define Error_CreateWindowEx  -2

LRESULT CALLBACK messageCallbackGo(HWND, UINT, WPARAM, LPARAM);

int openWindow(HWND* window);