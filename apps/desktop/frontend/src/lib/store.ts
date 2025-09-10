import { writable } from "svelte/store";
import type { connected } from "./wails/wailsjs/go/models";

export const connectedDevice = writable<connected.Connected | null>(null);
