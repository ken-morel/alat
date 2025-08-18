import { writable } from "svelte/store";
import type { device } from "$lib/wailsjs/go/models";

export const selectedDeviceForPairing = writable<device.DeviceInfo | null>(
  null,
);
