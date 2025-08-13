import { GetPairedDevices } from "$lib/wailsjs/go/app/App.js";
import type { core } from "$lib/wailsjs/go/models.js";

export async function load({}) {
  return {
    pairedDevicesPromise: new Promise<core.DeviceInfo[]>((resolve) =>
      setTimeout(() => GetPairedDevices().then(resolve), 1000),
    ),
  };
}
