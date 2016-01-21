#include "win.h"

LRESULT CALLBACK messageCallbackGo(HWND, UINT, WPARAM, LPARAM);

int openWindow(HWND* window) {
	WNDCLASSEX cls = {0};
	cls.cbSize = sizeof(WNDCLASSEX);
	cls.lpfnWndProc = messageCallbackGo;
	cls.lpszClassName = L"GoWindowClass";
	cls.hCursor = LoadCursor(0, IDC_ARROW);

	ATOM atom = RegisterClassEx(&cls);
	if (atom == 0) {
		return Error_RegisterClassEx;
	}

	*window = CreateWindowEx(0, L"GoWindowClass", "",
		WS_OVERLAPPEDWINDOW | WS_VISIBLE,
		300, 300, 640, 480, 0, 0, 0, 0);
	if (*window == 0) {
		return Error_CreateWindowEx;
	}

	return OK;
}
