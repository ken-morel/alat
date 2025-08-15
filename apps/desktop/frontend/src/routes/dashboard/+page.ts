import { GetPairedDevices } from "$lib/wailsjs/go/app/App.js";
import type { device } from "$lib/wailsjs/go/models.js";

export async function load({}) {
  return {
    pairedDevicesPromise: new Promise<device.DeviceInfo[]>((resolve) =>
      setTimeout(() => GetPairedDevices().then(resolve), 1000),
    ),
  };
}
