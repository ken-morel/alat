import { writable } from "svelte/store";
import type { device, pair } from "$lib/wailsjs/go/models";

export const selectedDeviceForPairing = writable<device.DeviceInfo | null>(
  null,
);

export const selectedPairedDevice = writable<pair.Pair | null>(null);
