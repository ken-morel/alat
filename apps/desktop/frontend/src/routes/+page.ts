import { GetPairedDevices } from "$lib/wailsjs/go/app/App.js";
import { pair } from "$lib/wailsjs/go/models.js";

export async function load({}) {
  return {
    pairedDevicesPromise: GetPairedDevices(),
  };
}
