import { GetPairedDevices } from "$lib/wailsjs/go/app/App.js";
import { pair } from "$lib/wailsjs/go/models.js";

export async function load({}) {
  return {
    pairedDevicesPromise: new Promise<pair.Pair[]>((resolve) =>
      setTimeout(() => GetPairedDevices().then(resolve), 1000),
    ),
  };
}
