import { writable } from "svelte/store";
import type { connected } from "./wails/wailsjs/go/models";
import { GetConnectedDevices } from "$lib/wails/wailsjs/go/app/App";

export const connectedDevices = writable<connected.Connected[]>([]);
export default connectedDevices;

export const interval = setInterval(async () => {
  connectedDevices.set((await GetConnectedDevices()) || []);
}, 1000);
