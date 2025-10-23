import { writable } from "svelte/store";
import type { connected } from "./wails/wailsjs/go/models";
import { app } from "$lib/wails/wailsjs/go/models";
import { get } from "svelte/store";

export const sendingFiles = writable<app.SendFile[]>([]);
export const sendingDevices = writable<connected.Connected[]>([]);

export function isSendingTo(dev: connected.Connected): bool {
  return get(sendingDevices).filter((d) => d.info.id == dev.info.id).length > 0;
}
