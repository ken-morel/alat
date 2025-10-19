import { writable } from "svelte/store";
import type { connected } from "./wails/wailsjs/go/models";
import { app } from "$lib/wails/wailsjs/go/models";

export const sendToDevices = writable<connected.Connected[]>([]);
