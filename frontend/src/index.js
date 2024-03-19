/** @typedef {import('./app').GoLaunch.GoMethodRequest} GoMethodRequest */
import { FuzzyFindDesktopEntry, LaunchApp } from "$lib/wailsjs/go/main/App";

export async function handleIframeMessage(event) {
  /** @type {GoMethodRequest} */
  const data = event.data;
  const { method, messageId, args } = data;
  if (method === "FuzzyFindDesktopEntry") {
    const result = await FuzzyFindDesktopEntry(args[0]);
    console.log(`returning result to iframe`, result);
    event.source.postMessage({ messageId, result }, event.origin);
  }
  if (method === "LaunchApp") {
    await LaunchApp(args[0]);
    event.source.postMessage({ messageId }, event.origin);
  }
}

window.addEventListener("message", handleIframeMessage);
