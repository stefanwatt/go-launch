export function FuzzyFindDesktopEntry(arg1) {
    return new Promise((resolve, reject) => {
        const messageId = Math.random().toString(36).substring(2);
        const listener = (event) => {
            if (event.data.type === "result" && event.data.messageId === messageId) {
                window.removeEventListener("message", listener);
                resolve(event.data.result);
            }
        };
        window.addEventListener("message", listener);
        window.parent.postMessage({ type: "call", method: "'1'", args: Array.from(arguments), messageId: messageId }, "*");
        setTimeout(() => {
            window.removeEventListener("message", listener);
            reject(new Error("'1': Timeout waiting for response"));
        }, 5000);
    });
}
export function GetDesktopEntries()  {
    return new Promise((resolve, reject) => {
        const messageId = Math.random().toString(36).substring(2);
        const listener = (event) => {
            if (event.data.type === "result" && event.data.messageId === messageId) {
                window.removeEventListener("message", listener);
                resolve(event.data.result);
            }
        };
        window.addEventListener("message", listener);
        window.parent.postMessage({ type: "call", method: "'1'", args: Array.from(arguments), messageId: messageId }, "*");
        setTimeout(() => {
            window.removeEventListener("message", listener);
            reject(new Error("'1': Timeout waiting for response"));
        }, 5000);
    });
}
export function GetExternalUiPath()  {
    return new Promise((resolve, reject) => {
        const messageId = Math.random().toString(36).substring(2);
        const listener = (event) => {
            if (event.data.type === "result" && event.data.messageId === messageId) {
                window.removeEventListener("message", listener);
                resolve(event.data.result);
            }
        };
        window.addEventListener("message", listener);
        window.parent.postMessage({ type: "call", method: "'1'", args: Array.from(arguments), messageId: messageId }, "*");
        setTimeout(() => {
            window.removeEventListener("message", listener);
            reject(new Error("'1': Timeout waiting for response"));
        }, 5000);
    });
}
export function HideLauncher()  {
    return new Promise((resolve, reject) => {
        const messageId = Math.random().toString(36).substring(2);
        const listener = (event) => {
            if (event.data.type === "result" && event.data.messageId === messageId) {
                window.removeEventListener("message", listener);
                resolve(event.data.result);
            }
        };
        window.addEventListener("message", listener);
        window.parent.postMessage({ type: "call", method: "'1'", args: Array.from(arguments), messageId: messageId }, "*");
        setTimeout(() => {
            window.removeEventListener("message", listener);
            reject(new Error("'1': Timeout waiting for response"));
        }, 5000);
    });
}
export function LaunchApp(arg1) {
    return new Promise((resolve, reject) => {
        const messageId = Math.random().toString(36).substring(2);
        const listener = (event) => {
            if (event.data.type === "result" && event.data.messageId === messageId) {
                window.removeEventListener("message", listener);
                resolve(event.data.result);
            }
        };
        window.addEventListener("message", listener);
        window.parent.postMessage({ type: "call", method: "'1'", args: Array.from(arguments), messageId: messageId }, "*");
        setTimeout(() => {
            window.removeEventListener("message", listener);
            reject(new Error("'1': Timeout waiting for response"));
        }, 5000);
    });
}
