import { GetConnectedDevices } from "$lib/wails/wailsjs/go/app/App";
import type { connected } from "$lib/wails/wailsjs/go/models";
import type { PageLoad } from "./$types";

export const load: PageLoad = async ({}) => {
  let connectedDevices: connected.Connected[];
  try {
    connectedDevices = await GetConnectedDevices();
  } catch (e: any) {
    throw "Error searching connected devices: " + e.toString();
  }
  return {
    connectedDevices: connectedDevices ?? [],
  };
};
