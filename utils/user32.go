package utils


import (
    "fmt"
    // "log"
    "net/http"
    "strconv"
    "strings"
    "syscall"
    "time"
    // "unsafe"
)

// --- WinAPI 常量定义 ---
const (
    MOUSEEVENTF_MOVE       = 0x0001
    MOUSEEVENTF_LEFTDOWN   = 0x0002
    MOUSEEVENTF_LEFTUP     = 0x0004
    KEYEVENTF_KEYUP        = 0x0002
)

var (
    user32         = syscall.NewLazyDLL("user32.dll")
    procMouseEvent = user32.NewProc("mouse_event")
    procKeybdEvent = user32.NewProc("keybd_event")
		// TODO 建议使用 SendInput 替代 mouse_event 和 keybd_event 
		// procSendInput = user32.NewProc("SendInput")
    procSetCursorPos = user32.NewProc("SetCursorPos")
)


// 虚拟键码映射
var vkMap = map[string]byte{
// 原有字母键
    "A": 0x41, "B": 0x42, "C": 0x43, "D": 0x44, "E": 0x45, "F": 0x46, "G": 0x47,
    "H": 0x48, "I": 0x49, "J": 0x4A, "K": 0x4B, "L": 0x4C, "M": 0x4D, "N": 0x4E,
    "O": 0x4F, "P": 0x50, "Q": 0x51, "R": 0x52, "S": 0x53, "T": 0x54, "U": 0x55,
    "V": 0x56, "W": 0x57, "X": 0x58, "Y": 0x59, "Z": 0x5A,

    // 数字键 0-9 (基于 ASCII 值 0x30-0x39)
    "0": 0x30, "1": 0x31, "2": 0x32, "3": 0x33, "4": 0x34,
    "5": 0x35, "6": 0x36, "7": 0x37, "8": 0x38, "9": 0x39,

    // 原有特殊键
    "ENTER": 0x0D, "SPACE": 0x20, "TAB": 0x09, "ESC": 0x1B,
    "UP": 0x26, "DOWN": 0x28, "LEFT": 0x25, "RIGHT": 0x27,
    "CTRL": 0x11, "ALT": 0x12, "SHIFT": 0x10,

    // 功能键 F1-F12 [[6]]
    "F1": 0x70, "F2": 0x71, "F3": 0x72, "F4": 0x73, "F5": 0x74, "F6": 0x75,
    "F7": 0x76, "F8": 0x77, "F9": 0x78, "F10": 0x79, "F11": 0x7A, "F12": 0x7B,

    // Delete 键 (使用常见值)
    "DEL": 0x2E, // DELETE key

    // 标点符号 (基于 ASCII 值)
    "EXCLAMATION": 0x21, // ! (Exclamation mark)
    "QUOTATION": 0x22,   // " (Quotation mark)
    "HASH": 0x23,        // # (Number sign)
    "DOLLAR": 0x24,      // $ (Dollar sign)
    "PERCENT": 0x25,     // % (Percent sign)
    "AMPERSAND": 0x26,   // & (Ampersand)
    "APOSTROPHE": 0x27,  // ' (Apostrophe)
    "PARENLEFT": 0x28,   // ( (Left parenthesis)
    "PARENRIGHT": 0x29,  // ) (Right parenthesis)
    "ASTERISK": 0x2A,    // * (Asterisk)
    "PLUS": 0x2B,        // + (Plus sign)
    "COMMA": 0x2C,       // , (Comma) [[3]]
    "MINUS": 0x2D,       // - (Minus sign / Hyphen)
    "PERIOD": 0x2E,      // . (Period / Dot)
    "SLASH": 0x2F,       // / (Slash)
    "COLON": 0x3A,       // : (Colon)
    "SEMICOLON": 0x3B,   // ; (Semicolon)
    "LESSTHAN": 0x3C,    // < (Less-than sign)
    "EQUALS": 0x3D,      // = (Equals sign)
    "GREATERTHAN": 0x3E, // > (Greater-than sign)
    "QUESTION": 0x3F,    // ? (Question mark)
    "AT": 0x40,          // @ (At sign)
    "BRACKETLEFT": 0x5B, // [ (Left square bracket)
    "BACKSLASH": 0x5C,   // \ (Backslash)
    "BRACKETRIGHT": 0x5D, // ] (Right square bracket)
    "CARET": 0x5E,       // ^ (Caret)
    "UNDERSCORE": 0x5F,  // _ (Underscore)
    "GRAVE": 0x60,       // ` (Grave accent)
    "BRACELEFT": 0x7B,   // { (Left curly brace)
    "BAR": 0x7C,         // | (Vertical bar)
    "BRACERIGHT": 0x7D,  // } (Right curly brace)
    "TILDE": 0x7E,       // ~ (Tilde)
}

// --- 控制函数 ---
func MouseMove(x, y int) {
    procSetCursorPos.Call(uintptr(x), uintptr(y))
}

func MouseClick() {
    procMouseEvent.Call(uintptr(MOUSEEVENTF_LEFTDOWN), 0, 0, 0, 0)
    time.Sleep(50 * time.Millisecond)
    procMouseEvent.Call(uintptr(MOUSEEVENTF_LEFTUP), 0, 0, 0, 0)
}



// --- HTTP handlers ---
func HandleMouse(w http.ResponseWriter, r *http.Request) {
    xStr := r.URL.Query().Get("x")
    yStr := r.URL.Query().Get("y")
    t := r.URL.Query().Get("type")

    x, _ := strconv.Atoi(xStr)
    y, _ := strconv.Atoi(yStr)

    if t == "click" {
        if xStr != "" && yStr != "" {
            MouseMove(x, y)
            time.Sleep(100 * time.Millisecond)
        }
        MouseClick()
        fmt.Fprint(w, "Mouse clicked at", x, y)
    } else if t == "move" {
        MouseMove(x, y)
        fmt.Fprint(w, "Mouse moved")
    } else {
        w.WriteHeader(400)
        fmt.Fprint(w, "Invalid mouse type")
    }
}


// 模拟按下
func keyDown(vk byte) {
    procKeybdEvent.Call(uintptr(vk), 0, 0, 0)
}

// 模拟松开
func keyUp(vk byte) {
    procKeybdEvent.Call(uintptr(vk), 0, uintptr(KEYEVENTF_KEYUP), 0)
}

// 模拟一次按键（按下+松开）
func keyTap(vk byte) {
    keyDown(vk)
    time.Sleep(30 * time.Millisecond)
    keyUp(vk)
}

// 组合键处理
func pressCombo(target byte, ctrl, shift, alt bool, hasTarget bool) {
    if ctrl {
        keyDown(vkMap["CTRL"])
        time.Sleep(10 * time.Millisecond)
    }
    if shift {
        keyDown(vkMap["SHIFT"])
        time.Sleep(10 * time.Millisecond)
    }
    if alt {
        keyDown(vkMap["ALT"])
        time.Sleep(10 * time.Millisecond)
    }

    // 主键
    if hasTarget {
        keyTap(target)
    } else {
        // 如果没有目标键，仅仅按下/松开组合键一次（模拟切换）
        time.Sleep(60 * time.Millisecond)
    }

    // 松开组合键（反顺序）
    if alt {
        keyUp(vkMap["ALT"])
    }
    if shift {
        keyUp(vkMap["SHIFT"])
    }
    if ctrl {
        keyUp(vkMap["CTRL"])
    }
}

// --- HTTP handler ---
func HandleKey(w http.ResponseWriter, r *http.Request) {
    keyStr := strings.ToUpper(r.URL.Query().Get("key"))
    ctrl, _ := strconv.ParseBool(r.URL.Query().Get("ctrl"))
    shift, _ := strconv.ParseBool(r.URL.Query().Get("shift"))
    alt, _ := strconv.ParseBool(r.URL.Query().Get("alt"))

    hasTarget := keyStr != ""
    var vk byte
    var ok bool 

    if hasTarget {
        vk, ok = vkMap[keyStr]
        if !ok {
            w.WriteHeader(400)
            fmt.Fprintf(w, "Unknown key: %s", keyStr)
            return
        }
    }

    pressCombo(vk, ctrl, shift, alt, hasTarget)
    fmt.Fprintf(w, "Pressed key=%s ctrl=%v shift=%v alt=%v", keyStr, ctrl, shift, alt)
}