import { WindowReloadApp } from "$lib/wailsjs/runtime/runtime";

export async function handleError() {
  WindowReloadApp();
}
// Only wails way to recover from error
